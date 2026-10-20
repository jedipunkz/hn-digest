---
source: "https://engineering.monday.com/rebuilding-monday-api-edge-for-agents/"
hn_url: "https://news.ycombinator.com/item?id=50170482"
title: "Rebuilding the monday.com API edge for AI Agents"
article_title: "One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents - monday AI engineering"
image: "https://engineering.monday.com/wp-content/uploads/2026/10/Rebuilding_the_API_Edge-scaled.jpeg"
author: "jensneuse"
captured_at: "2026-10-20T18:57:57Z"
capture_tool: "hn-digest"
hn_id: 50170482
score: 1
comments: 0
posted_at: "2026-10-20T18:22:18Z"
tags:
  - hacker-news
---

# Rebuilding the monday.com API edge for AI Agents

- HN: [50170482](https://news.ycombinator.com/item?id=50170482)
- Source: [engineering.monday.com](https://engineering.monday.com/rebuilding-monday-api-edge-for-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T18:22:18Z

## Translation

Title: Rebuilding the monday.com API edge for AI Agents
Article title: One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents - monday AI engineering
Description: How AI agents grew our API traffic 2.5x, and how we rebuilt the GraphQL edge to handle it with 92% less CPU

Article text:
One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents - monday AI engineering monday.com
One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents
“Move every overdue task to next week and let the owners know.”
It takes a few seconds to type. For the AI agent that carries it out, it’s dozens of calls to the monday.com API: find the board, read every item, check the dates, update each task, notify each owner. Now multiply that by every agent, in every workspace.
Every one of those calls passes through the API edge before it reaches any business logic. The edge checks who’s calling, what they’re allowed to do, and how much they can ask for, then routes the GraphQL query to the subgraphs that own the data. (We wrote about how our teams build and share those subgraphs in From API Chaos to Collaborative Graph .)
At the start of 2026, that layer was reaching its limits. Agents were calling the API far more than we had planned for. Keeping up meant repeatedly increasing CPU and memory resources, and the edge started taking a growing share of our on-call time. So we rebuilt it, with one clear target: the edge should add no more than 100 ms to a request at p99.
Compared with a year ago, the edge now handles 2.5x the traffic on 92% less CPU and 80% less memory, and the time it adds to a request at p99 is 11x lower.
As our customers connected AI tools to their workspaces and monday’s own agents started doing real work, weekly requests through the edge grew 2.5x in a year, from about 1 billion to about 2.5 billion.
A router that kept asking for permission
The old edge was an Apollo Router with a Node.js coprocessor attached. The router knew how to plan and run GraphQL queries, but every business rule (authentication, rate limits, complexity budgets, API versions) lived in the coprocessor, a separate Node.js service. So at several points in every request, the router had to stop, send the request to the coprocessor over HTTP, and wait for its answer before it could continue.
Follow one request through that setup, and the problems show up one by one.
Every hop parsed the request again. At each stage, the router serialized the request context, sent it over HTTP, and Node.js parsed it back. JSON parsing is CPU-heavy and blocks the event loop while it runs, so the bigger the payload, the longer everything else in that process waits. Several round trips meant the same request was parsed multiple times. We optimized this path again and again, but even passing context between the router and the coprocessor cost milliseconds we didn’t have.
Rejections came too late. Requests that a rate limit should have rejected immediately were only stopped deep inside the router’s lifecycle, after they had already been parsed and passed back and forth. That meant spending resources on requests that were never going to succeed.
Business logic was tangled with the router. Our rules lived in coprocessor hooks, shaped around Apollo Router’s request lifecycle and internals. There was no clear boundary between deciding whether a request may run and routing it, so even a small rule change meant reasoning about the whole router. Development was slower and more bug-prone than it should have been.
All three had the same root cause: business logic sitting inside the router.
We split the edge into two layers, each with one job.
The gateway decides. Authentication, authorization, rate limits, complexity budgets, versioning, agent detection, and observability all run here as a pipeline of small, independent steps. If a request fails a check, the gateway rejects it on the spot, and the router never sees it.
The router routes. It plans each query, fetches each part from the subgraph that owns it, and merges the results. The one rule it enforces is the complexity budget the gateway set, because estimating a query’s cost needs the query plan. Beyond that, the router holds no business logic, so it stays as simple as possible.
Today, that router is WunderGraph Cosmo , an open source GraphQL router written in Go. Speed wasn’t the main reason we picked it. We picked it because it’s easy to modify, profile, and adapt, and because we could work closely with the people who built it. Over many long sessions with the WunderGraph team, we brought production profiles, bug reports, and fix ideas, and worked through them together until they shipped in Cosmo. We also run our own fork of Cosmo, with a few changes that reduce its memory footprint under our traffic.
With both layers in Go, we profile, patch, and extend the entire request path with a single set of tools, and Go gives us more direct control over memory: we decide how data is laid out, stream it through rather than holding it, and reuse buffers rather than allocating new ones. Most of the savings came from many small fixes, not one big one. For example, file uploads now stream straight through instead of being buffered. Each response is held in memory once, not three times. Federated tracing data (FTV1) is parsed directly from the wire format, which halved its share of router CPU usage.
Moving a public API without anyone noticing
Two GraphQL routers can both follow the GraphQL federation spec and still answer the same query differently. Before any customer depended on Cosmo, we wanted to find every one of those differences on real traffic.
The gateway paid off immediately. Because it sat in front of the old router, it could safely send a sample of read-only queries to Cosmo. Since it happened in the background, shadowing did not affect latency. Writes were never copied, and feature flags controlled exactly which accounts were shadowed and at what volume.
Comparing two routers naively is mostly noise: timestamps, ordering, anything non-deterministic. So we used a technique popularized by Diffy , Twitter’s open source tool: each shadowed query also went to the old router a second time. Anything that differed between the two old responses was noise by definition; whatever remained was a real difference.
The real differences turned out to be edge cases in validation, null handling, error shapes, and query cost calculation, where Apollo and Cosmo implement the spec differently. We fixed each one in the gateway or in our fork before any customer moved.
Only then did real traffic move, gradually and behind feature flags. Moving an account back was a flag flip. All of monday.com’s API traffic, more than 2 billion requests a week, now runs on a completely new edge. The switch went smoothly, and no customer had to change a thing.
The time the edge adds to a request dropped from about 490 ms at p99 to 44 ms today! The median request now spends 9 ms in the edge, down from about 80
The line drops in two steps. At cutover, the router stopped waiting on coprocessor round trips, and the edge’s p99 fell from about 490 ms to about 200 ms. In September, we reworked authorization, which further reduced latency: the rare check that misses the cache now takes about 4 ms, down from about 130 ms at median. That took roughly another 150 ms off the p99, bringing it well under our 100 ms target.
The edge also got far cheaper to run: 92% less CPU and 80% less memory than a year ago, at 2.5x the traffic.
Even at 10x today’s traffic, the edge would use less CPU than the old edge needed a year ago. We believe agents will become the main consumers of the monday.com API , so we’re optimizing the edge for agentic flows: many small, fast calls made by software on behalf of a person. That’s where we see work heading, and we want monday.com to be an AI-ready platform that agents can build on.
If you’re building agents on top of monday.com, the API documentation is the place to start. And if problems like this one sound like fun, we’re hiring .
How we measured: incoming GraphQL requests per week across all production regions, public and internal, counted at the old router before the migration and at the gateway after the migration. CPU and memory are the total usage of every production edge component (router, coprocessor, and gateway), averaged over two full Tuesdays each year: Sep 23 and 30, 2025, and Sep 15 and 22, 2026. Latency is the time the edge spends on a request, compared over the same weekday window, 10:00–17:30 UTC: three Wednesdays on the old edge (Apr 8, Apr 22, and Jun 10, 2026) against Oct 7, 2026. The chart shows weekly p99 since March 2026, when we started recording this metric.
All Rights Reserved © monday.com +1 (201) 778-4567
Derech Menachem Begin 52, Tel Aviv-Yafo

## Original Extract

How AI agents grew our API traffic 2.5x, and how we rebuilt the GraphQL edge to handle it with 92% less CPU

One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents - monday AI engineering monday.com
One prompt, dozens of API calls: rebuilding the monday.com API edge for AI Agents
“Move every overdue task to next week and let the owners know.”
It takes a few seconds to type. For the AI agent that carries it out, it’s dozens of calls to the monday.com API: find the board, read every item, check the dates, update each task, notify each owner. Now multiply that by every agent, in every workspace.
Every one of those calls passes through the API edge before it reaches any business logic. The edge checks who’s calling, what they’re allowed to do, and how much they can ask for, then routes the GraphQL query to the subgraphs that own the data. (We wrote about how our teams build and share those subgraphs in From API Chaos to Collaborative Graph .)
At the start of 2026, that layer was reaching its limits. Agents were calling the API far more than we had planned for. Keeping up meant repeatedly increasing CPU and memory resources, and the edge started taking a growing share of our on-call time. So we rebuilt it, with one clear target: the edge should add no more than 100 ms to a request at p99.
Compared with a year ago, the edge now handles 2.5x the traffic on 92% less CPU and 80% less memory, and the time it adds to a request at p99 is 11x lower.
As our customers connected AI tools to their workspaces and monday’s own agents started doing real work, weekly requests through the edge grew 2.5x in a year, from about 1 billion to about 2.5 billion.
A router that kept asking for permission
The old edge was an Apollo Router with a Node.js coprocessor attached. The router knew how to plan and run GraphQL queries, but every business rule (authentication, rate limits, complexity budgets, API versions) lived in the coprocessor, a separate Node.js service. So at several points in every request, the router had to stop, send the request to the coprocessor over HTTP, and wait for its answer before it could continue.
Follow one request through that setup, and the problems show up one by one.
Every hop parsed the request again. At each stage, the router serialized the request context, sent it over HTTP, and Node.js parsed it back. JSON parsing is CPU-heavy and blocks the event loop while it runs, so the bigger the payload, the longer everything else in that process waits. Several round trips meant the same request was parsed multiple times. We optimized this path again and again, but even passing context between the router and the coprocessor cost milliseconds we didn’t have.
Rejections came too late. Requests that a rate limit should have rejected immediately were only stopped deep inside the router’s lifecycle, after they had already been parsed and passed back and forth. That meant spending resources on requests that were never going to succeed.
Business logic was tangled with the router. Our rules lived in coprocessor hooks, shaped around Apollo Router’s request lifecycle and internals. There was no clear boundary between deciding whether a request may run and routing it, so even a small rule change meant reasoning about the whole router. Development was slower and more bug-prone than it should have been.
All three had the same root cause: business logic sitting inside the router.
We split the edge into two layers, each with one job.
The gateway decides. Authentication, authorization, rate limits, complexity budgets, versioning, agent detection, and observability all run here as a pipeline of small, independent steps. If a request fails a check, the gateway rejects it on the spot, and the router never sees it.
The router routes. It plans each query, fetches each part from the subgraph that owns it, and merges the results. The one rule it enforces is the complexity budget the gateway set, because estimating a query’s cost needs the query plan. Beyond that, the router holds no business logic, so it stays as simple as possible.
Today, that router is WunderGraph Cosmo , an open source GraphQL router written in Go. Speed wasn’t the main reason we picked it. We picked it because it’s easy to modify, profile, and adapt, and because we could work closely with the people who built it. Over many long sessions with the WunderGraph team, we brought production profiles, bug reports, and fix ideas, and worked through them together until they shipped in Cosmo. We also run our own fork of Cosmo, with a few changes that reduce its memory footprint under our traffic.
With both layers in Go, we profile, patch, and extend the entire request path with a single set of tools, and Go gives us more direct control over memory: we decide how data is laid out, stream it through rather than holding it, and reuse buffers rather than allocating new ones. Most of the savings came from many small fixes, not one big one. For example, file uploads now stream straight through instead of being buffered. Each response is held in memory once, not three times. Federated tracing data (FTV1) is parsed directly from the wire format, which halved its share of router CPU usage.
Moving a public API without anyone noticing
Two GraphQL routers can both follow the GraphQL federation spec and still answer the same query differently. Before any customer depended on Cosmo, we wanted to find every one of those differences on real traffic.
The gateway paid off immediately. Because it sat in front of the old router, it could safely send a sample of read-only queries to Cosmo. Since it happened in the background, shadowing did not affect latency. Writes were never copied, and feature flags controlled exactly which accounts were shadowed and at what volume.
Comparing two routers naively is mostly noise: timestamps, ordering, anything non-deterministic. So we used a technique popularized by Diffy , Twitter’s open source tool: each shadowed query also went to the old router a second time. Anything that differed between the two old responses was noise by definition; whatever remained was a real difference.
The real differences turned out to be edge cases in validation, null handling, error shapes, and query cost calculation, where Apollo and Cosmo implement the spec differently. We fixed each one in the gateway or in our fork before any customer moved.
Only then did real traffic move, gradually and behind feature flags. Moving an account back was a flag flip. All of monday.com’s API traffic, more than 2 billion requests a week, now runs on a completely new edge. The switch went smoothly, and no customer had to change a thing.
The time the edge adds to a request dropped from about 490 ms at p99 to 44 ms today! The median request now spends 9 ms in the edge, down from about 80
The line drops in two steps. At cutover, the router stopped waiting on coprocessor round trips, and the edge’s p99 fell from about 490 ms to about 200 ms. In September, we reworked authorization, which further reduced latency: the rare check that misses the cache now takes about 4 ms, down from about 130 ms at median. That took roughly another 150 ms off the p99, bringing it well under our 100 ms target.
The edge also got far cheaper to run: 92% less CPU and 80% less memory than a year ago, at 2.5x the traffic.
Even at 10x today’s traffic, the edge would use less CPU than the old edge needed a year ago. We believe agents will become the main consumers of the monday.com API , so we’re optimizing the edge for agentic flows: many small, fast calls made by software on behalf of a person. That’s where we see work heading, and we want monday.com to be an AI-ready platform that agents can build on.
If you’re building agents on top of monday.com, the API documentation is the place to start. And if problems like this one sound like fun, we’re hiring .
How we measured: incoming GraphQL requests per week across all production regions, public and internal, counted at the old router before the migration and at the gateway after the migration. CPU and memory are the total usage of every production edge component (router, coprocessor, and gateway), averaged over two full Tuesdays each year: Sep 23 and 30, 2025, and Sep 15 and 22, 2026. Latency is the time the edge spends on a request, compared over the same weekday window, 10:00–17:30 UTC: three Wednesdays on the old edge (Apr 8, Apr 22, and Jun 10, 2026) against Oct 7, 2026. The chart shows weekly p99 since March 2026, when we started recording this metric.
All Rights Reserved © monday.com +1 (201) 778-4567
Derech Menachem Begin 52, Tel Aviv-Yafo
