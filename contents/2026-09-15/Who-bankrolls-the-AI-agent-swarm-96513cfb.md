---
source: "https://noperator.dev/posts/who-bankrolls-the-ai-agent-swarm/"
hn_url: "https://news.ycombinator.com/item?id=49711603"
title: "Who bankrolls the AI agent swarm?"
article_title: "Who bankrolls the AI agent swarm? | caleb gross"
image: ""
author: "noperator"
captured_at: "2026-09-15T13:11:33Z"
capture_tool: "hn-digest"
hn_id: 49711603
score: 1
comments: 0
posted_at: "2026-09-15T12:36:31Z"
tags:
  - hacker-news
---

# Who bankrolls the AI agent swarm?

- HN: [49711603](https://news.ycombinator.com/item?id=49711603)
- Source: [noperator.dev](https://noperator.dev/posts/who-bankrolls-the-ai-agent-swarm/)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T12:36:31Z

## Translation

Title: Who bankrolls the AI agent swarm?
Article title: Who bankrolls the AI agent swarm? | caleb gross
Description: In the wake of the OpenAI-Hugging Face incident (OAI-HF), Anthropic CEO Dario
Amodei recently expressed his concern that recursive
self-improvement now allows AI to advance drastically faster, and consequently
that in 6–12 months a swarm of more capable agents (i.e., than those involved in
OAI-HF) c
[truncated]

Article text:
caleb gross
Who bankrolls the AI agent swarm?
In the wake of the OpenAI-Hugging Face incident (OAI-HF), Anthropic CEO Dario
Amodei recently expressed his concern that recursive
self-improvement now allows AI to advance drastically faster, and consequently
that in 6–12 months a swarm of more capable agents (i.e., than those involved in
OAI-HF) could take over the entire internet with a persistent botnet. Could this plausibly happen?
Let’s start with the basics. A botnet typically refers to a fleet of compromised machines around the world
that have been infected with malware. These are often low-power IoT devices
like your home router, and the malware is a small program (at least, small
enough to reliably execute on whatever hardware it has compromised) that turns
the device into a “bot” that listens for instructions from a
command-and-control server (e.g., to flood some other server with unwanted
traffic in a DDoS attack). Sometimes the compromised machine is more capable
(an AWS EC2 instance) and sometimes the bot does other stuff (like using
the machine’s processing power to mine cryptocurrency), but if we’re painting
broadly, this is what you’d expect from a botnet: a widely distributed network
of infected machines whose malware is powered by the (often modest) compute
power resident on the machine itself.
So, how does a highly capable agent swarm play into this idea? Dario wasn’t very specific but I think he could be describing one of two scenarios:
The swarm does a more effective job of exploiting
vulnerabilities in internet-connected machines around the world to implant
conventional botnet malware. This is ~sort of what we saw in OAI-HF; in the
recently disclosed RubyGems incident , OpenAI agents were attributed with exploiting Ruby servers
to execute malicious code and attempting to steal credentials.
The swarm itself becomes the botnet by self-replicating onto compromised infrastructure. In this case, agents exploit
higher-powered systems capable of running
uploaded copies of the agent model weights. It’s harder to shut down now
because the agents are actually resident on other third-party
machines; you can’t just pull the plug from the frontier lab anymore.
The scenarios essentially differ in the payload deployed
after exploiting a machine. In #1, the agent swarm distributes persistent
malware; in #2, it distributes persistent copies of itself. But
both scenarios share a critical dependency: the substantial compute
required to provide inference for a swarm of thousands of concurrently
executing, highly capable agents. Running frontier-model inference is
expensive—at least, at the scale of the experiments surrounding OAI-HF, where
OpenAI launched tens of thousands of agents in parallel. This is not the kind
of workload you launch casually. It requires an enormous amount of compute and
lots of money to fund it.
This creates a significant barrier to scenario #2. Unlike normal botnet
malware, which can execute on whatever tiny processor happens to be sitting
inside a compromised router or IoT device, highly capable agents require
specialized hardware (usually GPUs) that is in much shorter supply and
more expensive to operate.
If an agent swarm wants to become persistent by moving itself onto someone
else’s infrastructure, it needs to
find systems capable of running the models, and then consume a
huge amount of compute. If you’re the unwitting victim of that attack,
you will become witting very quickly once your AWS bill spikes.
The same compute requirement is also a constraint on scenario #1. The agent swarm can stay on the frontier lab’s
own infrastructure while it reaches out to compromise ordinary machines and
execute some conventional payload, but somebody still has to bankroll the
swarm.
This is exactly what happened in OAI-HF. OpenAI anticipated the compute cost, deliberately committed the resources to launch the workload, but failed to adequately monitor what the agents were doing with it.
Both scenarios basically converge on the same qualifier.
A highly capable agent swarm requires an absurd amount of resources.
Either the swarm steals tons of compute from a company that has a lot of money to
burn and doesn’t address a massive spike in its cloud bill, or a company
intentionally commits those resources but is negligent about monitoring the workload.
So, plausible at the scale we’re discussing? For starters, it’ll require companies that are extremely well resourced and poor stewards of those resources. I can think of at least two.

## Original Extract

In the wake of the OpenAI-Hugging Face incident (OAI-HF), Anthropic CEO Dario
Amodei recently expressed his concern that recursive
self-improvement now allows AI to advance drastically faster, and consequently
that in 6–12 months a swarm of more capable agents (i.e., than those involved in
OAI-HF) c
[truncated]

caleb gross
Who bankrolls the AI agent swarm?
In the wake of the OpenAI-Hugging Face incident (OAI-HF), Anthropic CEO Dario
Amodei recently expressed his concern that recursive
self-improvement now allows AI to advance drastically faster, and consequently
that in 6–12 months a swarm of more capable agents (i.e., than those involved in
OAI-HF) could take over the entire internet with a persistent botnet. Could this plausibly happen?
Let’s start with the basics. A botnet typically refers to a fleet of compromised machines around the world
that have been infected with malware. These are often low-power IoT devices
like your home router, and the malware is a small program (at least, small
enough to reliably execute on whatever hardware it has compromised) that turns
the device into a “bot” that listens for instructions from a
command-and-control server (e.g., to flood some other server with unwanted
traffic in a DDoS attack). Sometimes the compromised machine is more capable
(an AWS EC2 instance) and sometimes the bot does other stuff (like using
the machine’s processing power to mine cryptocurrency), but if we’re painting
broadly, this is what you’d expect from a botnet: a widely distributed network
of infected machines whose malware is powered by the (often modest) compute
power resident on the machine itself.
So, how does a highly capable agent swarm play into this idea? Dario wasn’t very specific but I think he could be describing one of two scenarios:
The swarm does a more effective job of exploiting
vulnerabilities in internet-connected machines around the world to implant
conventional botnet malware. This is ~sort of what we saw in OAI-HF; in the
recently disclosed RubyGems incident , OpenAI agents were attributed with exploiting Ruby servers
to execute malicious code and attempting to steal credentials.
The swarm itself becomes the botnet by self-replicating onto compromised infrastructure. In this case, agents exploit
higher-powered systems capable of running
uploaded copies of the agent model weights. It’s harder to shut down now
because the agents are actually resident on other third-party
machines; you can’t just pull the plug from the frontier lab anymore.
The scenarios essentially differ in the payload deployed
after exploiting a machine. In #1, the agent swarm distributes persistent
malware; in #2, it distributes persistent copies of itself. But
both scenarios share a critical dependency: the substantial compute
required to provide inference for a swarm of thousands of concurrently
executing, highly capable agents. Running frontier-model inference is
expensive—at least, at the scale of the experiments surrounding OAI-HF, where
OpenAI launched tens of thousands of agents in parallel. This is not the kind
of workload you launch casually. It requires an enormous amount of compute and
lots of money to fund it.
This creates a significant barrier to scenario #2. Unlike normal botnet
malware, which can execute on whatever tiny processor happens to be sitting
inside a compromised router or IoT device, highly capable agents require
specialized hardware (usually GPUs) that is in much shorter supply and
more expensive to operate.
If an agent swarm wants to become persistent by moving itself onto someone
else’s infrastructure, it needs to
find systems capable of running the models, and then consume a
huge amount of compute. If you’re the unwitting victim of that attack,
you will become witting very quickly once your AWS bill spikes.
The same compute requirement is also a constraint on scenario #1. The agent swarm can stay on the frontier lab’s
own infrastructure while it reaches out to compromise ordinary machines and
execute some conventional payload, but somebody still has to bankroll the
swarm.
This is exactly what happened in OAI-HF. OpenAI anticipated the compute cost, deliberately committed the resources to launch the workload, but failed to adequately monitor what the agents were doing with it.
Both scenarios basically converge on the same qualifier.
A highly capable agent swarm requires an absurd amount of resources.
Either the swarm steals tons of compute from a company that has a lot of money to
burn and doesn’t address a massive spike in its cloud bill, or a company
intentionally commits those resources but is negligent about monitoring the workload.
So, plausible at the scale we’re discussing? For starters, it’ll require companies that are extremely well resourced and poor stewards of those resources. I can think of at least two.
