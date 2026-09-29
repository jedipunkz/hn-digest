---
source: "https://getunblocked.com/blog/adaptive-routing-inference-providers/"
hn_url: "https://news.ycombinator.com/item?id=49897322"
title: "Routing LLM traffic across inference providers with TCP-style congestion control"
article_title: "Routing LLM traffic across inference providers by cost, speed and reliability | Unblocked"
image: "https://cdn.sanity.io/images/31mw1ch6/production/4977228b3abcc05dd292872beb2ab7e3c0c29fde-1600x836.png"
author: "dennispi"
captured_at: "2026-09-29T17:45:42Z"
capture_tool: "hn-digest"
hn_id: 49897322
score: 2
comments: 0
posted_at: "2026-09-29T17:43:35Z"
tags:
  - hacker-news
---

# Routing LLM traffic across inference providers with TCP-style congestion control

- HN: [49897322](https://news.ycombinator.com/item?id=49897322)
- Source: [getunblocked.com](https://getunblocked.com/blog/adaptive-routing-inference-providers/)
- Score: 2
- Comments: 0
- Posted: 2026-09-29T17:43:35Z

## Translation

Title: Routing LLM traffic across inference providers with TCP-style congestion control
Article title: Routing LLM traffic across inference providers by cost, speed and reliability | Unblocked
Description: Providers now sell the same open-weight models at very different prices and speeds. We describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.

Article text:
Routing LLM traffic across inference providers by cost, speed and reliability | Unblocked Products Outcomes Resources Customers Security Docs Pricing Log In Book a Demo Products Context Engine Turn scattered data into curated context for agents. Unblocked Code Hand off a task. Get back a verified draft pull request. AI Code Review High-signal PR feedback grounded in how your system works. Coding Agents and MCP Give your coding agents a shared understanding of your org. Developer Q&A Get answers without digging through tools or interrupting your team. How we decide what your agent needs to know Scoping retrieval, resolving conflicts, and enforcing permissions.
Outcomes Mergeable code Give agents the context to write code that fits your system and passes review. Save tokens Give agents the full picture up front, so they spend fewer tokens finding it. Reduce rework Help agents produce a better first draft with fewer correction cycles. Cloudbeds “Unblocked is the institutional memory you were never given access to.”
Resources Blog Perspectives on context engineering and AI-native development. Open Source Context engineering tools, built by us and open to everyone. Videos Talks, podcasts, and demos from the team building the context layer. Events Join live webinars, workshops, and conversations in person. AI Adoption Assessment Find your team’s context maturity and get a personalized plan to level up. Context Maturity Guide Where your team sits today, and what it takes to reach the next level.
All Articles Routing LLM traffic across inference providers by cost, speed and reliability
Providers now sell the same open-weight models at very different prices and speeds. We describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.
The market for open-weight inference is heating up, and providers now sell the same models at very different prices and speeds. In this post we describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.
Earlier this summer we wrote about moving most of our agent loops from Anthropic models to GLM , an open-weight model that several inference providers sell. At the end of that post we were sending all GLM traffic to Baseten and using Fireworks as a failover.
We have since replaced that setup with an adaptive router. It picks a provider for each request based on the cost and speed we measure in production, and it moves traffic away from a provider that starts returning errors. This post covers why we built it, how it works, the bugs we found after we turned it on, and what it has done for us. The numbers come from our production token ledger and logs.
The router does not handle every customer. Customers who require their data to be processed in a specific region still go to Anthropic or OpenAI models.
Fixed routers need babysitting #
Baseten, Fireworks and CoreWeave all serve GLM 5.2. They charge different prices, run at different speeds, and have outages at different times.
We started with round-robin. In our last full week on it, Fireworks served 51% of GLM 5.2 tasks and Baseten served 49%. Fireworks' prices were 25% higher than Baseten's and Baseten was faster, so half of our traffic cost more and took longer for no benefit.
We then switched to a fixed order, with Baseten first and Fireworks used only when Baseten failed. In the first full week on the fixed order, Baseten served 98.5% of tasks and Fireworks served 1.5%. That was cheaper and faster, but we still had three problems.
Changing the order needed a code change and a deploy, and someone had to decide which provider was better that week.
Our circuit breaker counted a 429 as a failure. Five 429s in a minute removed Baseten from the pool for the whole fleet, even when Baseten only wanted us to slow down a little.
Outages required manually removing a provider from the pool to reduce the cost of failovers.
A third provider, CoreWeave, had list prices about 45% lower than Baseten's. We had no data on how it performed, so we did not know where to put it in the order. We could have sampled it on a low percentage of traffic, but that again would have required a deployment.
We wanted routing that used the cheapest provider by default, took speed into account, moved traffic gradually when a provider returned rate-limit errors, and did all of this without a human stepping in.
We'll first explain how the router works today, then we'll discuss the iterations it took to get here.
We measure cost and speed per task, not per model call. A task is one question or one code review, and it usually makes four to six model calls. We record every call in a ledger with the provider, token counts, duration, time to first token and task ID, and then add the calls up by task.
We do this because the task is what the customer pays for and waits for, and the cost of a task depends on more than the price per token. A provider that fails a call makes the task retry it, so the same task costs more tokens there. A provider with a worse cache hit rate bills more input tokens at full price for the same prompt. A provider that returns weaker output can make the agent loop take more turns. None of that shows up in the price of a single call. Speed works the same way: a customer feels the whole task, not one call inside it.
Every five minutes a job summarizes the ledger for each provider and task type and publishes the result to Redis. Each service keeps a copy in memory, so picking a provider does not make a network call. If the summary is too old to trust, for example because the sampling job stops working, the router uses the fixed order. We exclude tasks that used more than one provider, because we cannot assign their cost to either one.
Every provider serves the same model, so they differ in three ways: cost, speed and reliability. We handle reliability separately, which the next section explains. The score covers cost and speed.
Copy text score = 0.7 × (lowest cost / this provider's cost)
+ 0.3 × (fastest time / this provider's time) We try the provider with the highest score first, and the others in score order if it fails.
Why 0.7 and 0.3. The model is the same whichever provider serves it, so cost is the main reason to prefer one provider. Speed still matters, because people wait for these tasks. We chose the weights by deciding how much extra we would pay for speed. With 0.7 and 0.3, a faster provider can only come first if it costs less than 1.75 times the cheapest one, so we pay at most 75% more for speed.
In this example, B costs 10% more and is a third faster, and it comes first.
Comparing apples to apples. A task takes longer when the answer is longer. We estimate how long each provider would take to write an answer of average length, using its time to first token and its tokens per second. Otherwise a provider that happened to serve longer answers would look slower than it is.
We also apply three rules after scoring.
A provider that is more than three times slower than the fastest ranks last, whatever it costs. Cost is worth more than twice as much as speed, so without this rule a provider with a low enough cost than the alternative would come first however slow it was. A lower bound on speed, proportional to the fastest provider, is a simple rule that provides a minimum quality of service.
To account for jitter, a provider must overtake the current first choice provider by 5% higher to replace it. The measurements change a little every five minutes. Without a margin, two providers with nearly equal scores would swap places on every update, and every swap sends traffic to a provider that has none of our prompts cached.
A conversation stays pinned to the provider it started with for five minutes. Each turn of a conversation sends the whole history again. The provider that served the last turn has that history cached, so the next turn is cheaper and faster there. We use five minutes because that is roughly how long a prompt stays cached.
Errors are not part of the score. If they were, a provider that was cheap enough would still rank first while it was failing.
Each provider has a permitted request rate, stored in Redis and shared by the fleet. The rate changes with each result.
A 429 means the provider is asking for less traffic, so we halve the rate, as TCP does when it detects congestion. Other failures get a smaller cut, because the circuit breaker already handles a provider that is down. Recovery is slow on purpose, so a provider that has just failed takes minutes to get back to full traffic. This scheme is called additive increase and multiplicative decrease, and it converges to a stable rate whatever rate it starts from. When the first-choice provider is over its rate, the request goes to the next provider in score order. We never drop a request. The circuit breaker from the earlier post still handles a provider that is completely down.
Here is what that looks like in production. One evening Baseten returned two bursts of 429s about two hours apart. Within a few minutes of each one the router had cut Baseten's permitted rate and moved most requests to Fireworks. When the errors stopped, the rate climbed back and the traffic followed. No request failed for lack of a provider.
50% 75% 100% Baseten success rate, last 15 minutes Success rate 54% 0 5 10 Baseten permitted request rate (requests per second) Requests per second 0 15 30 GLM 5.3 input tokens per second by provider (thousands) Thousands per second 20:00 21:00 22:00 23:00 00:00 First burst of 429s Errors stop, traffic returns Second burst Time (UTC) Baseten Fireworks Times in UTC. Token rate is a six-minute rolling average. Success rate fell to 0% at 23:24; the chart floors it at 50%. Every provider gets 5% of traffic, forever # A provider that receives no traffic produces no measurements, so we would never find out that it had improved. We send at least 5% of requests to every provider in the pool. We chose 5% because the cost is small: at worst one request in twenty goes to a slower or more expensive provider. To mitigate customer impact, we recommend running parallel calls to the sampling pool from real work serving customers on the best provider. It's important to use real traffic so that sampling is reflective of the current state of customers.
We keep exploring a provider even after we have plenty of measurements for it. A provider's price and speed can change at any time, and the only way to notice is to keep sending it some traffic.
We built the off switch before we turned it on #
We first ran the router in production without acting on its output. For a week it logged what it would have chosen while the fixed order kept serving, and we compared the two.
In steady state it agreed with the fixed order, which was correct, because Baseten was better than Fireworks on both cost and speed. The useful signal was what it did when something went wrong. During that week Baseten returned bursts of rate-limit errors, and each time the router's backoff cut the traffic it would have sent there and then recovered as the errors cleared. Those incidents would have been resolved ahead of any intervention by the team, and after a week of sampling real traffic we were confident the router would work.
To mitigate the final risk of turning it on for customers we built an off switch first. A single Postgres row turns adaptive routing on or off. We change it from our admin console and every service applies the change within about five seconds. If the row is missing or cannot be read, routing uses the fixed order. We turned adaptive routing on in production a week after the shadow run, and it has been on since.
The first iteration wasn't completely smooth. Almost right away we encountered several issues that didn't present themselves while running in "shadow" mode.
Ri

[truncated]

## Original Extract

Providers now sell the same open-weight models at very different prices and speeds. We describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.

Routing LLM traffic across inference providers by cost, speed and reliability | Unblocked Products Outcomes Resources Customers Security Docs Pricing Log In Book a Demo Products Context Engine Turn scattered data into curated context for agents. Unblocked Code Hand off a task. Get back a verified draft pull request. AI Code Review High-signal PR feedback grounded in how your system works. Coding Agents and MCP Give your coding agents a shared understanding of your org. Developer Q&A Get answers without digging through tools or interrupting your team. How we decide what your agent needs to know Scoping retrieval, resolving conflicts, and enforcing permissions.
Outcomes Mergeable code Give agents the context to write code that fits your system and passes review. Save tokens Give agents the full picture up front, so they spend fewer tokens finding it. Reduce rework Help agents produce a better first draft with fewer correction cycles. Cloudbeds “Unblocked is the institutional memory you were never given access to.”
Resources Blog Perspectives on context engineering and AI-native development. Open Source Context engineering tools, built by us and open to everyone. Videos Talks, podcasts, and demos from the team building the context layer. Events Join live webinars, workshops, and conversations in person. AI Adoption Assessment Find your team’s context maturity and get a personalized plan to level up. Context Maturity Guide Where your team sits today, and what it takes to reach the next level.
All Articles Routing LLM traffic across inference providers by cost, speed and reliability
Providers now sell the same open-weight models at very different prices and speeds. We describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.
The market for open-weight inference is heating up, and providers now sell the same models at very different prices and speeds. In this post we describe an adaptive router that keeps us on the cheapest provider that is fast and healthy, and moves our traffic when that changes, without an engineer involved.
Earlier this summer we wrote about moving most of our agent loops from Anthropic models to GLM , an open-weight model that several inference providers sell. At the end of that post we were sending all GLM traffic to Baseten and using Fireworks as a failover.
We have since replaced that setup with an adaptive router. It picks a provider for each request based on the cost and speed we measure in production, and it moves traffic away from a provider that starts returning errors. This post covers why we built it, how it works, the bugs we found after we turned it on, and what it has done for us. The numbers come from our production token ledger and logs.
The router does not handle every customer. Customers who require their data to be processed in a specific region still go to Anthropic or OpenAI models.
Fixed routers need babysitting #
Baseten, Fireworks and CoreWeave all serve GLM 5.2. They charge different prices, run at different speeds, and have outages at different times.
We started with round-robin. In our last full week on it, Fireworks served 51% of GLM 5.2 tasks and Baseten served 49%. Fireworks' prices were 25% higher than Baseten's and Baseten was faster, so half of our traffic cost more and took longer for no benefit.
We then switched to a fixed order, with Baseten first and Fireworks used only when Baseten failed. In the first full week on the fixed order, Baseten served 98.5% of tasks and Fireworks served 1.5%. That was cheaper and faster, but we still had three problems.
Changing the order needed a code change and a deploy, and someone had to decide which provider was better that week.
Our circuit breaker counted a 429 as a failure. Five 429s in a minute removed Baseten from the pool for the whole fleet, even when Baseten only wanted us to slow down a little.
Outages required manually removing a provider from the pool to reduce the cost of failovers.
A third provider, CoreWeave, had list prices about 45% lower than Baseten's. We had no data on how it performed, so we did not know where to put it in the order. We could have sampled it on a low percentage of traffic, but that again would have required a deployment.
We wanted routing that used the cheapest provider by default, took speed into account, moved traffic gradually when a provider returned rate-limit errors, and did all of this without a human stepping in.
We'll first explain how the router works today, then we'll discuss the iterations it took to get here.
We measure cost and speed per task, not per model call. A task is one question or one code review, and it usually makes four to six model calls. We record every call in a ledger with the provider, token counts, duration, time to first token and task ID, and then add the calls up by task.
We do this because the task is what the customer pays for and waits for, and the cost of a task depends on more than the price per token. A provider that fails a call makes the task retry it, so the same task costs more tokens there. A provider with a worse cache hit rate bills more input tokens at full price for the same prompt. A provider that returns weaker output can make the agent loop take more turns. None of that shows up in the price of a single call. Speed works the same way: a customer feels the whole task, not one call inside it.
Every five minutes a job summarizes the ledger for each provider and task type and publishes the result to Redis. Each service keeps a copy in memory, so picking a provider does not make a network call. If the summary is too old to trust, for example because the sampling job stops working, the router uses the fixed order. We exclude tasks that used more than one provider, because we cannot assign their cost to either one.
Every provider serves the same model, so they differ in three ways: cost, speed and reliability. We handle reliability separately, which the next section explains. The score covers cost and speed.
Copy text score = 0.7 × (lowest cost / this provider's cost)
+ 0.3 × (fastest time / this provider's time) We try the provider with the highest score first, and the others in score order if it fails.
Why 0.7 and 0.3. The model is the same whichever provider serves it, so cost is the main reason to prefer one provider. Speed still matters, because people wait for these tasks. We chose the weights by deciding how much extra we would pay for speed. With 0.7 and 0.3, a faster provider can only come first if it costs less than 1.75 times the cheapest one, so we pay at most 75% more for speed.
In this example, B costs 10% more and is a third faster, and it comes first.
Comparing apples to apples. A task takes longer when the answer is longer. We estimate how long each provider would take to write an answer of average length, using its time to first token and its tokens per second. Otherwise a provider that happened to serve longer answers would look slower than it is.
We also apply three rules after scoring.
A provider that is more than three times slower than the fastest ranks last, whatever it costs. Cost is worth more than twice as much as speed, so without this rule a provider with a low enough cost than the alternative would come first however slow it was. A lower bound on speed, proportional to the fastest provider, is a simple rule that provides a minimum quality of service.
To account for jitter, a provider must overtake the current first choice provider by 5% higher to replace it. The measurements change a little every five minutes. Without a margin, two providers with nearly equal scores would swap places on every update, and every swap sends traffic to a provider that has none of our prompts cached.
A conversation stays pinned to the provider it started with for five minutes. Each turn of a conversation sends the whole history again. The provider that served the last turn has that history cached, so the next turn is cheaper and faster there. We use five minutes because that is roughly how long a prompt stays cached.
Errors are not part of the score. If they were, a provider that was cheap enough would still rank first while it was failing.
Each provider has a permitted request rate, stored in Redis and shared by the fleet. The rate changes with each result.
A 429 means the provider is asking for less traffic, so we halve the rate, as TCP does when it detects congestion. Other failures get a smaller cut, because the circuit breaker already handles a provider that is down. Recovery is slow on purpose, so a provider that has just failed takes minutes to get back to full traffic. This scheme is called additive increase and multiplicative decrease, and it converges to a stable rate whatever rate it starts from. When the first-choice provider is over its rate, the request goes to the next provider in score order. We never drop a request. The circuit breaker from the earlier post still handles a provider that is completely down.
Here is what that looks like in production. One evening Baseten returned two bursts of 429s about two hours apart. Within a few minutes of each one the router had cut Baseten's permitted rate and moved most requests to Fireworks. When the errors stopped, the rate climbed back and the traffic followed. No request failed for lack of a provider.
50% 75% 100% Baseten success rate, last 15 minutes Success rate 54% 0 5 10 Baseten permitted request rate (requests per second) Requests per second 0 15 30 GLM 5.3 input tokens per second by provider (thousands) Thousands per second 20:00 21:00 22:00 23:00 00:00 First burst of 429s Errors stop, traffic returns Second burst Time (UTC) Baseten Fireworks Times in UTC. Token rate is a six-minute rolling average. Success rate fell to 0% at 23:24; the chart floors it at 50%. Every provider gets 5% of traffic, forever # A provider that receives no traffic produces no measurements, so we would never find out that it had improved. We send at least 5% of requests to every provider in the pool. We chose 5% because the cost is small: at worst one request in twenty goes to a slower or more expensive provider. To mitigate customer impact, we recommend running parallel calls to the sampling pool from real work serving customers on the best provider. It's important to use real traffic so that sampling is reflective of the current state of customers.
We keep exploring a provider even after we have plenty of measurements for it. A provider's price and speed can change at any time, and the only way to notice is to keep sending it some traffic.
We built the off switch before we turned it on #
We first ran the router in production without acting on its output. For a week it logged what it would have chosen while the fixed order kept serving, and we compared the two.
In steady state it agreed with the fixed order, which was correct, because Baseten was better than Fireworks on both cost and speed. The useful signal was what it did when something went wrong. During that week Baseten returned bursts of rate-limit errors, and each time the router's backoff cut the traffic it would have sent there and then recovered as the errors cleared. Those incidents would have been resolved ahead of any intervention by the team, and after a week of sampling real traffic we were confident the router would work.
To mitigate the final risk of turning it on for customers we built an off switch first. A single Postgres row turns adaptive routing on or off. We change it from our admin console and every service applies the change within about five seconds. If the row is missing or cannot be read, routing uses the fixed order. We turned adaptive routing on in production a week after the shadow run, and it has been on since.
The first iteration wasn't completely smooth. Almost right away we encountered several issues that didn't present themselves while running in "shadow" mode.
Ri

[truncated]
