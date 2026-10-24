---
source: "https://twitter.com/OpenAI/status/2113766354386084338"
hn_url: "https://news.ycombinator.com/item?id=50216917"
title: "OpenAI found a critical flaw in Brotli (used by Node.js, Bun and Envoy)"
article_title: "OpenAI on X: \"During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.\nThe flaw could let an attacker run their own cod…\n[truncated]"
image: "https://pbs.twimg.com/card_img/2113162888793694208/F12C7Kzo?format=webp&name=medium"
author: "shenli3514"
captured_at: "2026-10-24T00:18:28Z"
capture_tool: "hn-digest"
hn_id: 50216917
score: 1
comments: 0
posted_at: "2026-10-23T23:46:57Z"
tags:
  - hacker-news
---

# OpenAI found a critical flaw in Brotli (used by Node.js, Bun and Envoy)

- HN: [50216917](https://news.ycombinator.com/item?id=50216917)
- Source: [twitter.com](https://twitter.com/OpenAI/status/2113766354386084338)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T23:46:57Z

## Translation

Title: OpenAI found a critical flaw in Brotli (used by Node.js, Bun and Envoy)
Article title: OpenAI on X: "During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own cod…
[truncated]
Description: During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected…

Article text:
OpenAI on X: "During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected systems by sending or replying with malicious compressed data, without any need to create a user account. That is, an unpatched server that accepts compressed input could be taken over by a malicious client, and an unpatched client that accepts compressed input (such as “curl --compressed”) could be taken over by a malicious server.
This root cause issue was actually fixed in the Brotli source code in February, but it wasn’t known at the time that it gave rise to a security vulnerability and there was no more recent release of Brotli. Since Brotli is built into other software, any software which includes this library needs to take the update too, followed by the users of that software. We disclosed the vulnerability to Brotli and other downstream projects such as Node.js, which released or are currently releasing security updates, and Brotli cut new release versions and published a CVE.
More details on the patch and mitigations below:
https://t.co/Y1Y72aHxrW"
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. OpenAI @OpenAI During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected systems by sending or replying with malicious compressed data, without any need to create a user account. That is, an unpatched server that accepts compressed input could be taken over by a malicious client, and an unpatched client that accepts compressed input (such as “curl --compressed”) could be taken over by a malicious server.
This root cause issue was actually fixed in the Brotli source code in February, but it wasn’t known at the time that it gave rise to a security vulnerability and there was no more recent release of Brotli. Since Brotli is built into other software, any software which includes this library needs to take the update too, followed by the users of that software. We disclosed the vulnerability to Brotli and other downstream projects such as Node.js, which released or are currently releasing security updates, and Brotli cut new release versions and published a CVE.
More details on the patch and mitigations below:
patchtheplanet.org/brotli OpenAI @OpenAI 4h Over the past several months, we have substantially hardened the security of our research training environment. We’ve removed live internet access for all frontier models during training and built an offline world for them to use instead—a cached version of the web where they can Show more 10:55 PM · Oct 23, 2026 · 206.1K Views 173 136 1.5K 378
Sanjay @sanjaygpts 1h this is absolutely a black october for internet 1 32 4.2K
Alek @AlekVectis 1h fixed in february, released never. i finally have something in common with brotli 15 3K
David Stark @stark4833 1h Glad you’re finding and fixing this stuff. Now if only some of that energy could go into listening to users for a change, and giving us 4o/4.1 back in GPT too. #4oForAll 1 11 1.8K

## Original Extract

During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected…

OpenAI on X: "During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected systems by sending or replying with malicious compressed data, without any need to create a user account. That is, an unpatched server that accepts compressed input could be taken over by a malicious client, and an unpatched client that accepts compressed input (such as “curl --compressed”) could be taken over by a malicious server.
This root cause issue was actually fixed in the Brotli source code in February, but it wasn’t known at the time that it gave rise to a security vulnerability and there was no more recent release of Brotli. Since Brotli is built into other software, any software which includes this library needs to take the update too, followed by the users of that software. We disclosed the vulnerability to Brotli and other downstream projects such as Node.js, which released or are currently releasing security updates, and Brotli cut new release versions and published a CVE.
More details on the patch and mitigations below:
https://t.co/Y1Y72aHxrW"
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. OpenAI @OpenAI During this red-teaming exercise, we also found a vulnerability that allows for unauthenticated remote code execution in Brotli, a widely-deployed compression library included in many pieces of software such as Node.js, Bun, and Envoy.
The flaw could let an attacker run their own code on affected systems by sending or replying with malicious compressed data, without any need to create a user account. That is, an unpatched server that accepts compressed input could be taken over by a malicious client, and an unpatched client that accepts compressed input (such as “curl --compressed”) could be taken over by a malicious server.
This root cause issue was actually fixed in the Brotli source code in February, but it wasn’t known at the time that it gave rise to a security vulnerability and there was no more recent release of Brotli. Since Brotli is built into other software, any software which includes this library needs to take the update too, followed by the users of that software. We disclosed the vulnerability to Brotli and other downstream projects such as Node.js, which released or are currently releasing security updates, and Brotli cut new release versions and published a CVE.
More details on the patch and mitigations below:
patchtheplanet.org/brotli OpenAI @OpenAI 4h Over the past several months, we have substantially hardened the security of our research training environment. We’ve removed live internet access for all frontier models during training and built an offline world for them to use instead—a cached version of the web where they can Show more 10:55 PM · Oct 23, 2026 · 206.1K Views 173 136 1.5K 378
Sanjay @sanjaygpts 1h this is absolutely a black october for internet 1 32 4.2K
Alek @AlekVectis 1h fixed in february, released never. i finally have something in common with brotli 15 3K
David Stark @stark4833 1h Glad you’re finding and fixing this stuff. Now if only some of that energy could go into listening to users for a change, and giving us 4o/4.1 back in GPT too. #4oForAll 1 11 1.8K
