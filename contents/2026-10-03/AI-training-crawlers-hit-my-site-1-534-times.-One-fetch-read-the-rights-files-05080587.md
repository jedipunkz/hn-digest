---
source: "https://juanlentino.com/notes/the-rights-files-nobody-reads/"
hn_url: "https://news.ycombinator.com/item?id=49944544"
title: "AI training crawlers hit my site 1,534 times. One fetch read the rights files"
article_title: "Do AI crawlers read opt-out files? What the server logs show"
image: "https://juanlentino.com/wp-content/uploads/sn-og/post-2071.png?v=1789736196"
author: "juanlentino"
captured_at: "2026-10-03T15:03:48Z"
capture_tool: "hn-digest"
hn_id: 49944544
score: 1
comments: 0
posted_at: "2026-10-03T14:24:44Z"
tags:
  - hacker-news
---

# AI training crawlers hit my site 1,534 times. One fetch read the rights files

- HN: [49944544](https://news.ycombinator.com/item?id=49944544)
- Source: [juanlentino.com](https://juanlentino.com/notes/the-rights-files-nobody-reads/)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T14:24:44Z

## Translation

Title: AI training crawlers hit my site 1,534 times. One fetch read the rights files
Article title: Do AI crawlers read opt-out files? What the server logs show
Description: Server-side telemetry: over fourteen days, declared AI-training crawlers fetched this site 1,534 times. One fetch touched a rights file. It was not GPTBot.

Article text:
Do AI crawlers read opt-out files? What the server logs show Skip to content Skip to content Juan Lentino Light Home
Fourteen days of machine readership
A measurement, not a blind spot
This site counts its machine readers. A small sensor at the edge classifies every automated fetch by crawler family and by the kind of surface it touched, then stores the counts and nothing else. No IP addresses, no paths beyond a coarse surface class, no record of humans at all. It was built to answer one narrow question: when a crawler that openly identifies as an AI-training agent reads this site, does it ever consult the terms it is bound by?
The terms are published where the standards say to publish them. A text-and-data-mining reservation sits at /.well-known/tdmrep.json , following the W3C's TDM Reservation Protocol. A machine-readable license sits at /license.xml . A policy page states the same position in plain language, and every HTML page this site serves carries the meta tags and response headers that point to it. Even robots.txt, the one file every crawler reads first, carries a Content-Signal line declaring ai-train=no and a License line pointing at the license file. Discovery is not the obstacle here. The pointers ride on every response the site sends.
Fourteen days of machine readership
The sensor first ran on July 28, 2026. From then to August 10 it recorded 17,490 automated reads. Most of that is the ordinary background hum of the web, generic bots and uptime probes and search engines going about their rounds. Inside it, 1,534 fetches came from crawlers that declare themselves as AI-training agents in their own user-agent strings. The distribution is lumpy. August 8 alone accounts for 841 of them, nearly all from OpenAI's GPTBot working through the site's pages. On a median day the figure was fifteen.
Across those fourteen days, exactly one of those 1,534 fetches touched a rights file. It came from OAI-SearchBot, the crawler OpenAI runs for its search product and states is not used for training. GPTBot, the one the reservation is written for, made hundreds of reads of the site's prose and never once asked for the terms attached to it.
The files are not unread in absolute terms. The sensor counted eighty fetches over the period, and for the stretch where the edge logs still name the requester, the answer is this site's own monitoring rather than anybody's crawler. The interesting part is that eighty understates it. Most of that tooling does not announce itself as a machine at all, so the sensor never counted it in the first place. The rights files have an audience, and it is the site checking its own work.
A measurement, not a blind spot
A number this small deserves suspicion, so the instrument is worth describing. The sensor runs on the same edge worker that serves the rights files, and it records the visit before the response is written, so a crawler cannot reach any of the three files without crossing it. If the sensor broke, the panel it feeds would report an error rather than a clean count. That distinction is load-bearing. A null says the watch failed; a number says the watch was on. This number is one.
These files are not decoration. In the EU, the text-and-data-mining exception lets a rights holder reserve their work from mining, provided the reservation is machine-readable, and TDMRep exists so that reservation has a standard address. The whole legal mechanism assumes the miners look. On this site, across these fourteen days, the one that matters did not. One site and fourteen days make an anecdote rather than a study, and the instrumentation is exactly what makes the anecdote worth writing down.
The publishing side of this arrangement is complete. The reservation is declared, the license is up, and the pointers travel on every page. The reading side has not appeared. Until it does, a machine-readable rights file is a statement for the record, not a channel to the machines it addresses.
Check your inbox or spam folder to confirm your subscription.
See it on the public Bitcoin ledger (mempool.space) →
Download proof (.ots)
Git ledger
Verify it yourself
2026.09.10 04 MIN Market harm names no track
2026.10.01 03 MIN Before you opt in to an AI music platform
2026.09.06 04 MIN There is no test that returns human
2026.07.03 03 MIN The court found the floor
2026.08.02 Better models erase the evidence

## Original Extract

Server-side telemetry: over fourteen days, declared AI-training crawlers fetched this site 1,534 times. One fetch touched a rights file. It was not GPTBot.

Do AI crawlers read opt-out files? What the server logs show Skip to content Skip to content Juan Lentino Light Home
Fourteen days of machine readership
A measurement, not a blind spot
This site counts its machine readers. A small sensor at the edge classifies every automated fetch by crawler family and by the kind of surface it touched, then stores the counts and nothing else. No IP addresses, no paths beyond a coarse surface class, no record of humans at all. It was built to answer one narrow question: when a crawler that openly identifies as an AI-training agent reads this site, does it ever consult the terms it is bound by?
The terms are published where the standards say to publish them. A text-and-data-mining reservation sits at /.well-known/tdmrep.json , following the W3C's TDM Reservation Protocol. A machine-readable license sits at /license.xml . A policy page states the same position in plain language, and every HTML page this site serves carries the meta tags and response headers that point to it. Even robots.txt, the one file every crawler reads first, carries a Content-Signal line declaring ai-train=no and a License line pointing at the license file. Discovery is not the obstacle here. The pointers ride on every response the site sends.
Fourteen days of machine readership
The sensor first ran on July 28, 2026. From then to August 10 it recorded 17,490 automated reads. Most of that is the ordinary background hum of the web, generic bots and uptime probes and search engines going about their rounds. Inside it, 1,534 fetches came from crawlers that declare themselves as AI-training agents in their own user-agent strings. The distribution is lumpy. August 8 alone accounts for 841 of them, nearly all from OpenAI's GPTBot working through the site's pages. On a median day the figure was fifteen.
Across those fourteen days, exactly one of those 1,534 fetches touched a rights file. It came from OAI-SearchBot, the crawler OpenAI runs for its search product and states is not used for training. GPTBot, the one the reservation is written for, made hundreds of reads of the site's prose and never once asked for the terms attached to it.
The files are not unread in absolute terms. The sensor counted eighty fetches over the period, and for the stretch where the edge logs still name the requester, the answer is this site's own monitoring rather than anybody's crawler. The interesting part is that eighty understates it. Most of that tooling does not announce itself as a machine at all, so the sensor never counted it in the first place. The rights files have an audience, and it is the site checking its own work.
A measurement, not a blind spot
A number this small deserves suspicion, so the instrument is worth describing. The sensor runs on the same edge worker that serves the rights files, and it records the visit before the response is written, so a crawler cannot reach any of the three files without crossing it. If the sensor broke, the panel it feeds would report an error rather than a clean count. That distinction is load-bearing. A null says the watch failed; a number says the watch was on. This number is one.
These files are not decoration. In the EU, the text-and-data-mining exception lets a rights holder reserve their work from mining, provided the reservation is machine-readable, and TDMRep exists so that reservation has a standard address. The whole legal mechanism assumes the miners look. On this site, across these fourteen days, the one that matters did not. One site and fourteen days make an anecdote rather than a study, and the instrumentation is exactly what makes the anecdote worth writing down.
The publishing side of this arrangement is complete. The reservation is declared, the license is up, and the pointers travel on every page. The reading side has not appeared. Until it does, a machine-readable rights file is a statement for the record, not a channel to the machines it addresses.
Check your inbox or spam folder to confirm your subscription.
See it on the public Bitcoin ledger (mempool.space) →
Download proof (.ots)
Git ledger
Verify it yourself
2026.09.10 04 MIN Market harm names no track
2026.10.01 03 MIN Before you opt in to an AI music platform
2026.09.06 04 MIN There is no test that returns human
2026.07.03 03 MIN The court found the floor
2026.08.02 Better models erase the evidence
