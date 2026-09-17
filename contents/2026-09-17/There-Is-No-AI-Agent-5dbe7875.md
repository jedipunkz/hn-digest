---
source: "https://zak.im/there-is-no-ai-agent/"
hn_url: "https://news.ycombinator.com/item?id=49743011"
title: "There Is No AI Agent"
article_title: "There Is No AI Agent"
image: ""
author: "pzk1"
captured_at: "2026-09-17T16:22:52Z"
capture_tool: "hn-digest"
hn_id: 49743011
score: 1
comments: 0
posted_at: "2026-09-17T16:19:41Z"
tags:
  - hacker-news
---

# There Is No AI Agent

- HN: [49743011](https://news.ycombinator.com/item?id=49743011)
- Source: [zak.im](https://zak.im/there-is-no-ai-agent/)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T16:19:41Z

## Translation

Title: There Is No AI Agent
Description: On delegation, liability, and machine worship.

Article text:
Sign in
Subscribe
By ZAKIM
in
AI
—
17 Sep 2026
There Is No AI Agent
On delegation, liability, and machine worship.
Autonomous systems are not creatures; they are software executing human instructions with access to capital. Calling them "agents" obscures legal liability and erases the human labor behind the weights.
Generative chat models fail at backend automation because sequential strings are slow, hallucinate schemas, and carry no calibrated confidence.
Real automation treats models as software: fast, typed, parallel decision heads embedded in deterministic code. The open-source community is already reproducing this.
Classical commercial law solved delegation across long distances centuries ago through shared-risk equity ( qirâd ) and the strict elimination of contractual uncertainty ( gharar ).
What modern tech calls agent autonomy is mostly gharar with branding. The engineering and the law point in the same direction: bound the output, sign the record, keep liability human.
On a Tuesday afternoon in Kuala Lumpur, I was debugging a webhook receiver that had silently dropped thirty-two transactions during a mid-day traffic spike.
The payment gateway had cleared the funds, but the partner's callback endpoint timed out after four seconds, leaving thirty-two customers charged and zero orders written to the fulfillment database. I wrote a SQL script in BigQuery to isolate the orphaned gateway transaction IDs, cross-referenced them against the bank settlement logs, and queued them into an idempotent retry worker.
A few terminals over, a Playwright automation script was spinning up in a container. Twice a month, it logs into a state statutory zakat portal to submit eight hundred salary-deduction records, handling session timeouts, CSRF tokens, and multipart uploads that used to take hours of manual keying.
I build and maintain automated backend pipelines and payment infrastructure for platforms handling religious financial obligations: zakat, charitable endowments, and religious welfare.
When you spend your working life inside the plumbing of financial obligations, you develop a visceral allergy to Silicon Valley mysticism.
Right now, the technology industry is in the grip of a theological fever. Venture capitalists and model labs speak about "autonomous AI agents" in the hushed, reverent tones once reserved for demiurges. We are told that we are on the verge of creating autonomous synthetic entities: digital persons that will negotiate contracts, allocate capital, and manage infrastructure without human intervention.
This framing is an illusion. Worse, it is an evasion.
In an interview on StarTalk, Jaron Lanier cut through the fog with a single, devastating sentence:
"Part of the ideology that makes AI into a creature instead of a collaboration also wants to think of bits as being this ethereal thing that's free and infinite. You show me a bit that didn't involve work. You show me a bit that didn't disperse heat... Information is physical or it's nothing. You can think of large language model AI as a whole bunch of people whose work was combined into this single document. But why do people want it to be a creature? You're a young man. You think you're the center of the universe. Of course you think you're making God." [1]
Lanier's critique cuts to the bone: There is no AI. There are only people, whose labor has been scraped and compiled into a statistical mirror.
When software executives personify the model, when they call a next-token predictor an "agent" and pretend it possesses intention or will, they are performing an ancient sleight of hand. They erase the millions of human writers, programmers, and labeling workers whose work made the weights possible: what Lanier calls the crisis of data dignity. And they erase their own legal and financial liability when the system fails.
If an algorithm hallucinates a bad medical triage recommendation, drops a customer's payment, or signs a ruinous commercial lease, the vendor shrugs: The agent made an autonomous decision. The model is a non-deterministic black box.
In classical Islamic thought, this evasion has a precise name: talbis (counterfeiting: dressing up an illusion to obscure a material reality).
A machine has no ruh (soul), no iradah (moral will), and no dhimmah (legal personhood). It cannot bear an amanah (a sacred trust), because an inanimate artifact cannot be held to moral account before God or man. To attribute agency to silicon is not futurism. It is idolatry.
Yet the practical demand for automation has never been greater.
If we reject the Silicon Valley fantasy of the AI "creature," how do we actually build automated systems that move money, process sensitive records, and execute workflows at scale?
The answer requires two things that Silicon Valley currently lacks:
An honest software engineering discipline that recognizes models as software modules rather than digital beings: compiler targets, fuzzy classifiers, and probabilistic decision gates embedded into deterministic code.
A classical commercial legal architecture that has already solved the problem of delegating high-stakes, long-distance agency without central state monopolies.
The technical foundation is already emerging across the engineering landscape.
The legal and institutional architecture was engineered fourteen hundred years ago in the commercial republics of Mecca and Medina.
1. Where Is the Automation? The Model Is Just Software
For four years, the tech industry has operated under a bizarre assumption: that the path to business automation runs through chat interfaces and open-ended generative text.
Diogo Almeida, an early researcher at OpenAI who helped build the reinforcement learning methods behind ChatGPT, asked the obvious question when introducing TypeSafe AI:
"Models have been superhuman at chat for years, so where is all the automation?" [2]
If large language models are so brilliant, why are businesses still manually reconciling spreadsheets, copying data between portals, and paying rooms full of people to check forms?
The bottleneck is architectural.
Generative large language models are autoregressive token predictors. They output strings: one character or word at a time, each conditioned on the last. Strings are flexible for conversation and open-ended drafting.
For software automation, strings are a disaster.
When you embed a generative LLM inside a production software loop, four things go wrong:
Latency : Generating tokens sequentially takes anywhere from 3 to 300 seconds. That is acceptable if a human is reading a chatbot response. It is fatal inside an automated backend workflow.
Hallucination and Type Errors : Strings have no guardrails. A model instructed to output JSON can drop a bracket, invent an imaginary enum value, hallucinate an API parameter, or inject a conversational refusal into a payload. In a deep software dependency chain, a single malformed token crashes the pipeline.
Epistemic Dishonesty : LLMs are notoriously overconfident. A model will state a false claim with the exact same linguistic certainty as a verified fact. It cannot tell software when it is guessing.
Economics : Autoregressive sampling consumes massive GPU compute, driving token costs to levels that obliterate the margins of the business running it.
This is why most "AI agent" startups stall at the demo stage. They build wrappers around open-ended chat prompts, discover that the agent derails three percent of the time, and realize that a three percent catastrophic failure rate makes a system un-deployable in banking, logistics, or healthcare.
The realization now taking root across the engineering community is simple: An LLM is not an agent. It is just software.
It is an array of floating-point numbers performing mathematical function approximation. When you stop asking it to be a conscious digital employee, you begin treating it as an engineering primitive: a fast, fuzzy classifier that converts unstructured data into typed program state.
We are seeing this architectural pivot break out across both commercial research and open-source systems:
TypeSafe's System One Models (Jev) : Instead of generating sequential strings, TypeSafe built models trained specifically for parallel structured decisions. The model takes unstructured program state in, and outputs typed probabilistic decisions in 70 milliseconds, with zero hallucinations and zero type errors, because the output space is mathematically constrained to predefined schemas. [2]
Browser-Use's Jev-Ultrafast : In real-world browser automation, the team at Browser-Use demonstrated how clean this separation becomes. Rather than letting an LLM generate brittle natural-language scripts or slow visual coordinates, they indexed the DOM controls into an action space. A fast structured decision head picks the operation (click, select, wait) and element index in a single network round-trip. In their published measurements for jev-ultrafast , browser protocol calls dropped from 1,092 to 101, and median task time from 9.450s to 7.092s, with a live Google Flights search completing in 7.1 seconds. Generative text models are only summoned if actual text needs to be typed into an input field. [3]
Open-source one-pass scoring ( jevlike and open-jev ) : The open-source community proved immediately that this pattern is not a closed vendor secret. Independent implementations like Vinny Larouge's jevlike showed that an option-attention head paired with a small, frozen open-weights encoder (such as Qwen 2.5 0.5B) can evaluate discrete choices in a single forward pass in under ten milliseconds, at roughly 100x the speed of a small decoder forced to write 400 tokens. [4] Dasein Labs' open-jev replicated the System One contract on local Apple Silicon via MLX, scoring structured queries at roughly 86 milliseconds per example on consumer hardware without emitting a single hallucinated token. [5]
What does this wave of experimentation prove?
It proves that we are in the earliest stages of demystifying the model.
When you treat machine learning as software, you decompose the problem properly: - Deterministic code handles control flow, state persistence, database transactions, and network retries. - Fast, typed classification heads handle fuzzy categorical decisions and intent routing. - Generative strings are reserved strictly for communicating with human beings.
This is what automation actually looks like. It is not an autonomous creature wandering through the internet. It is a strictly bounded, type-safe function call embedded inside human-authored code.
2. The Ancient Agency Problem: Delegation Across the Desert
Once you strip away the science-fiction terminology and recognize that an "agent" is simply delegated software executing instructions with capital and access, the problem transforms.
It ceases to be a machine learning problem. It becomes an institutional agency problem .
How does a principal delegate resources and operational authority to an agent who operates outside direct supervision, in an environment fraught with uncertainty, without allowing the agent to embezzle capital or shirk responsibility?
In his work Early Islam and the Birth of Capitalism , Benedikt Koehler demonstrates that this was the foundational economic problem of early Arabia. [6]
Consider the sixth-century Meccan caravan.
Mecca sat at the intersection of global trade routes, but had no agricultural base and no manufacturing. Its entire economy depended on organizing massive, capital-intensive trade caravans to Yemen, Syria, and Abyssinia. A single caravan could aggregate millions of dirhams of capital from hundreds of individual investors, assemble thousands of camels, and march across thousands of miles of hostile desert for six months.
The caravan was completely disconnected from the investors who funded it. There were no telegraphs, no banks, no police, and no sovereign state standing over the desert. Once the caravan departed Mecca, the managin

[truncated]

## Original Extract

On delegation, liability, and machine worship.

Sign in
Subscribe
By ZAKIM
in
AI
—
17 Sep 2026
There Is No AI Agent
On delegation, liability, and machine worship.
Autonomous systems are not creatures; they are software executing human instructions with access to capital. Calling them "agents" obscures legal liability and erases the human labor behind the weights.
Generative chat models fail at backend automation because sequential strings are slow, hallucinate schemas, and carry no calibrated confidence.
Real automation treats models as software: fast, typed, parallel decision heads embedded in deterministic code. The open-source community is already reproducing this.
Classical commercial law solved delegation across long distances centuries ago through shared-risk equity ( qirâd ) and the strict elimination of contractual uncertainty ( gharar ).
What modern tech calls agent autonomy is mostly gharar with branding. The engineering and the law point in the same direction: bound the output, sign the record, keep liability human.
On a Tuesday afternoon in Kuala Lumpur, I was debugging a webhook receiver that had silently dropped thirty-two transactions during a mid-day traffic spike.
The payment gateway had cleared the funds, but the partner's callback endpoint timed out after four seconds, leaving thirty-two customers charged and zero orders written to the fulfillment database. I wrote a SQL script in BigQuery to isolate the orphaned gateway transaction IDs, cross-referenced them against the bank settlement logs, and queued them into an idempotent retry worker.
A few terminals over, a Playwright automation script was spinning up in a container. Twice a month, it logs into a state statutory zakat portal to submit eight hundred salary-deduction records, handling session timeouts, CSRF tokens, and multipart uploads that used to take hours of manual keying.
I build and maintain automated backend pipelines and payment infrastructure for platforms handling religious financial obligations: zakat, charitable endowments, and religious welfare.
When you spend your working life inside the plumbing of financial obligations, you develop a visceral allergy to Silicon Valley mysticism.
Right now, the technology industry is in the grip of a theological fever. Venture capitalists and model labs speak about "autonomous AI agents" in the hushed, reverent tones once reserved for demiurges. We are told that we are on the verge of creating autonomous synthetic entities: digital persons that will negotiate contracts, allocate capital, and manage infrastructure without human intervention.
This framing is an illusion. Worse, it is an evasion.
In an interview on StarTalk, Jaron Lanier cut through the fog with a single, devastating sentence:
"Part of the ideology that makes AI into a creature instead of a collaboration also wants to think of bits as being this ethereal thing that's free and infinite. You show me a bit that didn't involve work. You show me a bit that didn't disperse heat... Information is physical or it's nothing. You can think of large language model AI as a whole bunch of people whose work was combined into this single document. But why do people want it to be a creature? You're a young man. You think you're the center of the universe. Of course you think you're making God." [1]
Lanier's critique cuts to the bone: There is no AI. There are only people, whose labor has been scraped and compiled into a statistical mirror.
When software executives personify the model, when they call a next-token predictor an "agent" and pretend it possesses intention or will, they are performing an ancient sleight of hand. They erase the millions of human writers, programmers, and labeling workers whose work made the weights possible: what Lanier calls the crisis of data dignity. And they erase their own legal and financial liability when the system fails.
If an algorithm hallucinates a bad medical triage recommendation, drops a customer's payment, or signs a ruinous commercial lease, the vendor shrugs: The agent made an autonomous decision. The model is a non-deterministic black box.
In classical Islamic thought, this evasion has a precise name: talbis (counterfeiting: dressing up an illusion to obscure a material reality).
A machine has no ruh (soul), no iradah (moral will), and no dhimmah (legal personhood). It cannot bear an amanah (a sacred trust), because an inanimate artifact cannot be held to moral account before God or man. To attribute agency to silicon is not futurism. It is idolatry.
Yet the practical demand for automation has never been greater.
If we reject the Silicon Valley fantasy of the AI "creature," how do we actually build automated systems that move money, process sensitive records, and execute workflows at scale?
The answer requires two things that Silicon Valley currently lacks:
An honest software engineering discipline that recognizes models as software modules rather than digital beings: compiler targets, fuzzy classifiers, and probabilistic decision gates embedded into deterministic code.
A classical commercial legal architecture that has already solved the problem of delegating high-stakes, long-distance agency without central state monopolies.
The technical foundation is already emerging across the engineering landscape.
The legal and institutional architecture was engineered fourteen hundred years ago in the commercial republics of Mecca and Medina.
1. Where Is the Automation? The Model Is Just Software
For four years, the tech industry has operated under a bizarre assumption: that the path to business automation runs through chat interfaces and open-ended generative text.
Diogo Almeida, an early researcher at OpenAI who helped build the reinforcement learning methods behind ChatGPT, asked the obvious question when introducing TypeSafe AI:
"Models have been superhuman at chat for years, so where is all the automation?" [2]
If large language models are so brilliant, why are businesses still manually reconciling spreadsheets, copying data between portals, and paying rooms full of people to check forms?
The bottleneck is architectural.
Generative large language models are autoregressive token predictors. They output strings: one character or word at a time, each conditioned on the last. Strings are flexible for conversation and open-ended drafting.
For software automation, strings are a disaster.
When you embed a generative LLM inside a production software loop, four things go wrong:
Latency : Generating tokens sequentially takes anywhere from 3 to 300 seconds. That is acceptable if a human is reading a chatbot response. It is fatal inside an automated backend workflow.
Hallucination and Type Errors : Strings have no guardrails. A model instructed to output JSON can drop a bracket, invent an imaginary enum value, hallucinate an API parameter, or inject a conversational refusal into a payload. In a deep software dependency chain, a single malformed token crashes the pipeline.
Epistemic Dishonesty : LLMs are notoriously overconfident. A model will state a false claim with the exact same linguistic certainty as a verified fact. It cannot tell software when it is guessing.
Economics : Autoregressive sampling consumes massive GPU compute, driving token costs to levels that obliterate the margins of the business running it.
This is why most "AI agent" startups stall at the demo stage. They build wrappers around open-ended chat prompts, discover that the agent derails three percent of the time, and realize that a three percent catastrophic failure rate makes a system un-deployable in banking, logistics, or healthcare.
The realization now taking root across the engineering community is simple: An LLM is not an agent. It is just software.
It is an array of floating-point numbers performing mathematical function approximation. When you stop asking it to be a conscious digital employee, you begin treating it as an engineering primitive: a fast, fuzzy classifier that converts unstructured data into typed program state.
We are seeing this architectural pivot break out across both commercial research and open-source systems:
TypeSafe's System One Models (Jev) : Instead of generating sequential strings, TypeSafe built models trained specifically for parallel structured decisions. The model takes unstructured program state in, and outputs typed probabilistic decisions in 70 milliseconds, with zero hallucinations and zero type errors, because the output space is mathematically constrained to predefined schemas. [2]
Browser-Use's Jev-Ultrafast : In real-world browser automation, the team at Browser-Use demonstrated how clean this separation becomes. Rather than letting an LLM generate brittle natural-language scripts or slow visual coordinates, they indexed the DOM controls into an action space. A fast structured decision head picks the operation (click, select, wait) and element index in a single network round-trip. In their published measurements for jev-ultrafast , browser protocol calls dropped from 1,092 to 101, and median task time from 9.450s to 7.092s, with a live Google Flights search completing in 7.1 seconds. Generative text models are only summoned if actual text needs to be typed into an input field. [3]
Open-source one-pass scoring ( jevlike and open-jev ) : The open-source community proved immediately that this pattern is not a closed vendor secret. Independent implementations like Vinny Larouge's jevlike showed that an option-attention head paired with a small, frozen open-weights encoder (such as Qwen 2.5 0.5B) can evaluate discrete choices in a single forward pass in under ten milliseconds, at roughly 100x the speed of a small decoder forced to write 400 tokens. [4] Dasein Labs' open-jev replicated the System One contract on local Apple Silicon via MLX, scoring structured queries at roughly 86 milliseconds per example on consumer hardware without emitting a single hallucinated token. [5]
What does this wave of experimentation prove?
It proves that we are in the earliest stages of demystifying the model.
When you treat machine learning as software, you decompose the problem properly: - Deterministic code handles control flow, state persistence, database transactions, and network retries. - Fast, typed classification heads handle fuzzy categorical decisions and intent routing. - Generative strings are reserved strictly for communicating with human beings.
This is what automation actually looks like. It is not an autonomous creature wandering through the internet. It is a strictly bounded, type-safe function call embedded inside human-authored code.
2. The Ancient Agency Problem: Delegation Across the Desert
Once you strip away the science-fiction terminology and recognize that an "agent" is simply delegated software executing instructions with capital and access, the problem transforms.
It ceases to be a machine learning problem. It becomes an institutional agency problem .
How does a principal delegate resources and operational authority to an agent who operates outside direct supervision, in an environment fraught with uncertainty, without allowing the agent to embezzle capital or shirk responsibility?
In his work Early Islam and the Birth of Capitalism , Benedikt Koehler demonstrates that this was the foundational economic problem of early Arabia. [6]
Consider the sixth-century Meccan caravan.
Mecca sat at the intersection of global trade routes, but had no agricultural base and no manufacturing. Its entire economy depended on organizing massive, capital-intensive trade caravans to Yemen, Syria, and Abyssinia. A single caravan could aggregate millions of dirhams of capital from hundreds of individual investors, assemble thousands of camels, and march across thousands of miles of hostile desert for six months.
The caravan was completely disconnected from the investors who funded it. There were no telegraphs, no banks, no police, and no sovereign state standing over the desert. Once the caravan departed Mecca, the managin

[truncated]
