---
source: "https://productnow.ai/blogs/scaling-reliable-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=50291714"
title: "Scaling Reliable AI Agents"
article_title: "Scaling Reliable AI Agents | ProductNow"
image: "https://productnow.ai/Blogs/scaling-reliable-ai-agents.webp"
author: "kadhirvelm"
captured_at: "2026-10-29T16:39:35Z"
capture_tool: "hn-digest"
hn_id: 50291714
score: 1
comments: 0
posted_at: "2026-10-29T16:29:26Z"
tags:
  - hacker-news
---

# Scaling Reliable AI Agents

- HN: [50291714](https://news.ycombinator.com/item?id=50291714)
- Source: [productnow.ai](https://productnow.ai/blogs/scaling-reliable-ai-agents)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T16:29:26Z

## Translation

Title: Scaling Reliable AI Agents
Article title: Scaling Reliable AI Agents | ProductNow
Description: Running an agent once is easy. Running it thousands of times a day is a different sport. How ProductNow chases 9s of reliability: shrink what the agent does, let the schema drive the retry loop, and make search the interface.

Article text:
Scaling Reliable AI Agents | ProductNow Home How it works Careers Our blog Log in Get in touch Home How it works Careers Our blog Log in Get in touch All posts Scaling Reliable AI Agents by Kadhir Manickam (4 minutes) <section data-section-id='3906425a-7937-4822-aba9-0ab0678941c9'><p><span style="white-space: pre-wrap;">Running an agent once is easy. Running it thousands of times a day is a different sport.</span></p><p><br></p><p><span style="white-space: pre-wrap;">LLMs are non-deterministic, they do weird stuff a lot. That's easy to hide in a demo and brutal in production, where you're chasing 9s of reliability and every weird edge eventually shows up.</span></p><p><br></p><p><span style="white-space: pre-wrap;">The tail of failures is so wide that two things get hard fast: reproducing an issue, and hillclimbing your way past it. There are too many one-off failures to replay by hand, so we run in high volume and watch the probabilities instead. When a failure rate stops moving, we know the issue is ironed out.</span></p><p><br></p><p><span style="white-space: pre-wrap;">At ProductNow, a few principles get us there. They all point the same way: ask the model to do as little as possible, and let deterministic engineering carry the rest.</span></p><p><br></p></section>
<section data-section-id='bb25fde0-22db-44dc-bd24-ae48b64e2813'><p><b><strong style="white-space: pre-wrap;">1) Shrink what the agents do</strong></b></p><p><br></p><p><span style="white-space: pre-wrap;">The first principle: put determinism everywhere the model does not need to think.</span></p><p><br></p><p><span style="white-space: pre-wrap;">In practice, every part of the system that touches the outside world is a deterministic script. When we fetch information from an external system, a script does it, with a strict schema on its inputs and outputs. When we write back downstream, same thing. That leaves the agent in the middle doing the minimum: it assembles context inside ProductNow and acts on
[truncated]
The Company Brain: What It Buys You and What It Takes
Every company runs on knowledge it can't find. What a company brain is, the four ways it pays for itself, the six failure modes that sink homegrown attempts, and how ProductNow answers each one.
Every company now has thousands of agents trying to get work done, and almost none of them know what the organization actually knows. Organizational context is a big data problem with its own 4Vs. What a context engine is, the five attributes that define one, and why ProductNow built it as a plain document editor.
AI-native data integration without an ontology
Every data integration effort historically required an ontology — slow to build, locking clients out of their own system, never truly self-service. AI changes that equation: continuous ingestion into living documents, agent-managed pruning, and a deep graph — integrations in 24 hours, not months.
One shared source of truth from all your systems, people, and agents, so you can act with confidence.
© 2026 ProductNow. All rights reserved.

## Original Extract

Running an agent once is easy. Running it thousands of times a day is a different sport. How ProductNow chases 9s of reliability: shrink what the agent does, let the schema drive the retry loop, and make search the interface.

Scaling Reliable AI Agents | ProductNow Home How it works Careers Our blog Log in Get in touch Home How it works Careers Our blog Log in Get in touch All posts Scaling Reliable AI Agents by Kadhir Manickam (4 minutes) <section data-section-id='3906425a-7937-4822-aba9-0ab0678941c9'><p><span style="white-space: pre-wrap;">Running an agent once is easy. Running it thousands of times a day is a different sport.</span></p><p><br></p><p><span style="white-space: pre-wrap;">LLMs are non-deterministic, they do weird stuff a lot. That's easy to hide in a demo and brutal in production, where you're chasing 9s of reliability and every weird edge eventually shows up.</span></p><p><br></p><p><span style="white-space: pre-wrap;">The tail of failures is so wide that two things get hard fast: reproducing an issue, and hillclimbing your way past it. There are too many one-off failures to replay by hand, so we run in high volume and watch the probabilities instead. When a failure rate stops moving, we know the issue is ironed out.</span></p><p><br></p><p><span style="white-space: pre-wrap;">At ProductNow, a few principles get us there. They all point the same way: ask the model to do as little as possible, and let deterministic engineering carry the rest.</span></p><p><br></p></section>
<section data-section-id='bb25fde0-22db-44dc-bd24-ae48b64e2813'><p><b><strong style="white-space: pre-wrap;">1) Shrink what the agents do</strong></b></p><p><br></p><p><span style="white-space: pre-wrap;">The first principle: put determinism everywhere the model does not need to think.</span></p><p><br></p><p><span style="white-space: pre-wrap;">In practice, every part of the system that touches the outside world is a deterministic script. When we fetch information from an external system, a script does it, with a strict schema on its inputs and outputs. When we write back downstream, same thing. That leaves the agent in the middle doing the minimum: it assembles context inside ProductNow and acts on
[truncated]
The Company Brain: What It Buys You and What It Takes
Every company runs on knowledge it can't find. What a company brain is, the four ways it pays for itself, the six failure modes that sink homegrown attempts, and how ProductNow answers each one.
Every company now has thousands of agents trying to get work done, and almost none of them know what the organization actually knows. Organizational context is a big data problem with its own 4Vs. What a context engine is, the five attributes that define one, and why ProductNow built it as a plain document editor.
AI-native data integration without an ontology
Every data integration effort historically required an ontology — slow to build, locking clients out of their own system, never truly self-service. AI changes that equation: continuous ingestion into living documents, agent-managed pruning, and a deep graph — integrations in 24 hours, not months.
One shared source of truth from all your systems, people, and agents, so you can act with confidence.
© 2026 ProductNow. All rights reserved.
