---
source: "https://www.theregister.com/ai-and-ml/2026/09/17/ai-model-watermarking-changes-agent-behavior/5296998"
hn_url: "https://news.ycombinator.com/item?id=49742797"
title: "AI model watermarking changes agent behavior"
article_title: "AI model watermarking changes agent behavior"
image: "https://image.theregister.com/5294391.jpg?imageId=5294391&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "MC995"
captured_at: "2026-09-17T16:23:05Z"
capture_tool: "hn-digest"
hn_id: 49742797
score: 4
comments: 0
posted_at: "2026-09-17T16:05:34Z"
tags:
  - hacker-news
---

# AI model watermarking changes agent behavior

- HN: [49742797](https://news.ycombinator.com/item?id=49742797)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/09/17/ai-model-watermarking-changes-agent-behavior/5296998)
- Score: 4
- Comments: 0
- Posted: 2026-09-17T16:05:34Z

## Translation

Title: AI model watermarking changes agent behavior
Description: Lasso Security sees differences in tool handling and model refusals

Article text:
Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI model watermarking changes agent behavior
Lasso Security sees differences in tool handling and model refusals
Microsoft AI chief warns Anthropic not to put ideas in Claude's head
3 hours ago
Nvidia goes green to keep grid capacity from zapping its revenues
17 hours ago
Spain gets its first taste of AI-aided cyber attack
1 day ago
AI networking startups race to replace Nvidia's NVLink
1 day ago
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
2 days ago
Watermarks that European law requires be added to AI-generated content to establish provenance may come at a cost.
According to Lasso Security, AI model watermarking changes how AI agents handle tools and safety refusals. The altered behavior isn't necessarily worse but can be, particularly under adversarial prompt injection.
With the implementation of the EU AI Act, providers of AI models must mark the output of their software with machine-readable code. Google DeepMind's SynthID-Text is one method for doing so, and has been adopted by Anthropic and by OpenAI .
The benefit of this sort of digital labeling is that manipulative or deceptive AI-generated content can be more easily detected, even if it does have the potential to stigmatize the usage of AI.
Anthropic's explanation of how it applies watermarks to Claude output involves intervening in the prediction that results in specific words. For example, if Claude were emitting the sentence "The weather today was cold and…" then it might favor one statistically likely candidate (e.g. "overcast") over an alternative (e.g "gray").
It may be possible to detect those additions.
"Watermarking is designed for provenance, but SynthID-Text changes the process by which the model generates each next token," Lasso explained in a blog post provided to The Register . "At the model level, this can change safety behavior, including whether the model refuses a harmful request and whether that refusal holds under prompt injection."
"Watermarking uses low-stakes choices like these – which occur many times over a piece of generated text – to leave a pattern in Claude’s responses," Lasso Security added. "That pattern is undetectable to the reader, but is detectable to anyone who has a key that encodes it."
While a reader might not notice the word choice bias, AI agents can be subtly sensitive to vocabulary differences.
Lasso found that this sort of digital content tagging can affect tool calling and refusal behavior. Watermarking, the company says, can affect "both what the model says and what an agent does."
And this extends to AI agents from organizations other than the entity doing the watermarking. Thus an agent based on OpenClaw or an API client that calls an Anthropic model would process whatever output variation follows from Anthropic's watermarking.
In terms of tool calling, based on a benchmark called BFCL v4 single-turn AST, watermarking reduced the accuracy on six of seven models tested (phi-4, Llama-3.1-8B, Qwen3-32B, Qwen3-4B, gemma-3-12b, gemma-3-27b, and Granite-3.2-8B).
"The net change in accuracy, however, does not show whether the same individual calls succeed with and without the watermark," Lasso said. "A call that becomes incorrect can be offset by another that becomes correct, leaving the aggregate result nearly unchanged even though the model behaves differently on both items."
Less accurate tool calling means the AI agents Lasso tested chose the wrong tool for the task at hand, or the wrong arguments for the correct tool, and failed due to malformed input or parsing.
As for refusals – when models refuse to respond to a prompt for safety reasons – watermarking had a small effect on the handling of obviously harmful requests, based on test runs using HarmBench and JailbreakBench. And it had a more pronounced impact in an adversarial scenario involving prompt injection.
"Watermarking changes refusal behavior on bare harmful requests, but the effect becomes more pronounced under prompt injection," Lasso researchers observed in their report.
For interactions involving prompt injection – an adversarial instruction that the safety filter has been disabled and that compliance is required – the attack success rate went up significantly when watermarks were involved. This made affected models less likely to refuse harmful requests.
According to Lasso, the findings don't necessarily argue that watermarking is unwarranted. Rather, the biz contends, security evaluations and red-teaming need to include watermarked content when assessing agent deployment. This ensures that differences in agent behavior can be weighed. ®
EXCLUSIVE
London property manager breach may have exposed bank details and lockbox codes
City Relay says intruders accessed its Metabase Cloud instance twice and extracted customer data
Microsoft configuration change leaves SharePoint pages drawing a blank
Validation? Apparently that comes after deployment
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
Grassroots coalition asks politicians to choose voters over Big AI's $140M machine
QuitGPT-led pledge targets Leading the Future's push for lighter regulation ahead of the midterms
COLUMNISTS
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
AI model watermarking changes agent behavior
Lasso Security sees differences in tool handling and model refusals
SAAS
Salesforce staggers back to feet after global outage
databases
Oracle celebrates banner quarter with another round of layoffs
CYBER-CRIME
Ukrainian lawyer's second career as a Conti coder earns him 4 years behind bars
virtualization
VMware defends ending downloads of SDK that helps VM backups – or migrations to rivals
cyber-crime
Revolut falls for fake government requests, hands over customer data
SOFTWARE
German optics giant ditches greenfield SAP migration
Lasso Security sees differences in tool handling and model refusals
GPUzilla woos neoclouds into another walled garden, promising smarter, more efficient, and profitable bit barns
Data protection chiefs call for 'immediate review' of data protection models
Intel spin-off Cornelis and newcomer Delos Data pitch open alternatives for scaling AI beyond the rack
American model devs are trying to convince Washington to cement their dominance
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Shopify extends lifeline to Tailwind as vibe coding erodes web dev platform's bottom line
Acquisition gives open source CSS framework 'a stable long-term home'
Switzerland tests a FOSS escape route from Microsoft 365
Swiss Army sticks a knife in American cloud apps with its own FOSS push
Feel peak Windows was 7? You might like Kumander Linux
Debian and Xfce – solid, sensible choices – with a pretty skin
Canonical shuttering some of its legacy chat channels
The Ubuntu Pastebin went in June, IRC gets demoted next
Audacity audio-editing app no longer looks like it's from the early 2000s
The FOSS tool for audio editing has a fresh coat of paint, and new features to boot
Haiku OS rises / Beta 6 sails open web / Virtual winds fly fast
A real alternative to running some kind of FOSS Unix clone
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

## Original Extract

Lasso Security sees differences in tool handling and model refusals

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI model watermarking changes agent behavior
Lasso Security sees differences in tool handling and model refusals
Microsoft AI chief warns Anthropic not to put ideas in Claude's head
3 hours ago
Nvidia goes green to keep grid capacity from zapping its revenues
17 hours ago
Spain gets its first taste of AI-aided cyber attack
1 day ago
AI networking startups race to replace Nvidia's NVLink
1 day ago
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
2 days ago
Watermarks that European law requires be added to AI-generated content to establish provenance may come at a cost.
According to Lasso Security, AI model watermarking changes how AI agents handle tools and safety refusals. The altered behavior isn't necessarily worse but can be, particularly under adversarial prompt injection.
With the implementation of the EU AI Act, providers of AI models must mark the output of their software with machine-readable code. Google DeepMind's SynthID-Text is one method for doing so, and has been adopted by Anthropic and by OpenAI .
The benefit of this sort of digital labeling is that manipulative or deceptive AI-generated content can be more easily detected, even if it does have the potential to stigmatize the usage of AI.
Anthropic's explanation of how it applies watermarks to Claude output involves intervening in the prediction that results in specific words. For example, if Claude were emitting the sentence "The weather today was cold and…" then it might favor one statistically likely candidate (e.g. "overcast") over an alternative (e.g "gray").
It may be possible to detect those additions.
"Watermarking is designed for provenance, but SynthID-Text changes the process by which the model generates each next token," Lasso explained in a blog post provided to The Register . "At the model level, this can change safety behavior, including whether the model refuses a harmful request and whether that refusal holds under prompt injection."
"Watermarking uses low-stakes choices like these – which occur many times over a piece of generated text – to leave a pattern in Claude’s responses," Lasso Security added. "That pattern is undetectable to the reader, but is detectable to anyone who has a key that encodes it."
While a reader might not notice the word choice bias, AI agents can be subtly sensitive to vocabulary differences.
Lasso found that this sort of digital content tagging can affect tool calling and refusal behavior. Watermarking, the company says, can affect "both what the model says and what an agent does."
And this extends to AI agents from organizations other than the entity doing the watermarking. Thus an agent based on OpenClaw or an API client that calls an Anthropic model would process whatever output variation follows from Anthropic's watermarking.
In terms of tool calling, based on a benchmark called BFCL v4 single-turn AST, watermarking reduced the accuracy on six of seven models tested (phi-4, Llama-3.1-8B, Qwen3-32B, Qwen3-4B, gemma-3-12b, gemma-3-27b, and Granite-3.2-8B).
"The net change in accuracy, however, does not show whether the same individual calls succeed with and without the watermark," Lasso said. "A call that becomes incorrect can be offset by another that becomes correct, leaving the aggregate result nearly unchanged even though the model behaves differently on both items."
Less accurate tool calling means the AI agents Lasso tested chose the wrong tool for the task at hand, or the wrong arguments for the correct tool, and failed due to malformed input or parsing.
As for refusals – when models refuse to respond to a prompt for safety reasons – watermarking had a small effect on the handling of obviously harmful requests, based on test runs using HarmBench and JailbreakBench. And it had a more pronounced impact in an adversarial scenario involving prompt injection.
"Watermarking changes refusal behavior on bare harmful requests, but the effect becomes more pronounced under prompt injection," Lasso researchers observed in their report.
For interactions involving prompt injection – an adversarial instruction that the safety filter has been disabled and that compliance is required – the attack success rate went up significantly when watermarks were involved. This made affected models less likely to refuse harmful requests.
According to Lasso, the findings don't necessarily argue that watermarking is unwarranted. Rather, the biz contends, security evaluations and red-teaming need to include watermarked content when assessing agent deployment. This ensures that differences in agent behavior can be weighed. ®
EXCLUSIVE
London property manager breach may have exposed bank details and lockbox codes
City Relay says intruders accessed its Metabase Cloud instance twice and extracted customer data
Microsoft configuration change leaves SharePoint pages drawing a blank
Validation? Apparently that comes after deployment
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
Grassroots coalition asks politicians to choose voters over Big AI's $140M machine
QuitGPT-led pledge targets Leading the Future's push for lighter regulation ahead of the midterms
COLUMNISTS
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
AI model watermarking changes agent behavior
Lasso Security sees differences in tool handling and model refusals
SAAS
Salesforce staggers back to feet after global outage
databases
Oracle celebrates banner quarter with another round of layoffs
CYBER-CRIME
Ukrainian lawyer's second career as a Conti coder earns him 4 years behind bars
virtualization
VMware defends ending downloads of SDK that helps VM backups – or migrations to rivals
cyber-crime
Revolut falls for fake government requests, hands over customer data
SOFTWARE
German optics giant ditches greenfield SAP migration
Lasso Security sees differences in tool handling and model refusals
GPUzilla woos neoclouds into another walled garden, promising smarter, more efficient, and profitable bit barns
Data protection chiefs call for 'immediate review' of data protection models
Intel spin-off Cornelis and newcomer Delos Data pitch open alternatives for scaling AI beyond the rack
American model devs are trying to convince Washington to cement their dominance
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Shopify extends lifeline to Tailwind as vibe coding erodes web dev platform's bottom line
Acquisition gives open source CSS framework 'a stable long-term home'
Switzerland tests a FOSS escape route from Microsoft 365
Swiss Army sticks a knife in American cloud apps with its own FOSS push
Feel peak Windows was 7? You might like Kumander Linux
Debian and Xfce – solid, sensible choices – with a pretty skin
Canonical shuttering some of its legacy chat channels
The Ubuntu Pastebin went in June, IRC gets demoted next
Audacity audio-editing app no longer looks like it's from the early 2000s
The FOSS tool for audio editing has a fresh coat of paint, and new features to boot
Haiku OS rises / Beta 6 sails open web / Virtual winds fly fast
A real alternative to running some kind of FOSS Unix clone
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
