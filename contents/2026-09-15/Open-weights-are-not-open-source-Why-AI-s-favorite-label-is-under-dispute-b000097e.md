---
source: "https://www.theregister.com/columnists/2026/09/15/open-weights-are-not-open-source-why-ais-favorite-label-is-under-dispute/5295436"
hn_url: "https://news.ycombinator.com/item?id=49711455"
title: "Open weights are not open source: Why AI's favorite label is under dispute"
article_title: "Open weights are not open source: Why AI's favorite label is under dispute"
image: "https://image.theregister.com/5295478.jpg?imageId=5295478&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "pseudolus"
captured_at: "2026-09-15T13:11:37Z"
capture_tool: "hn-digest"
hn_id: 49711455
score: 4
comments: 0
posted_at: "2026-09-15T12:21:26Z"
tags:
  - hacker-news
---

# Open weights are not open source: Why AI's favorite label is under dispute

- HN: [49711455](https://news.ycombinator.com/item?id=49711455)
- Source: [www.theregister.com](https://www.theregister.com/columnists/2026/09/15/open-weights-are-not-open-source-why-ais-favorite-label-is-under-dispute/5295436)
- Score: 4
- Comments: 0
- Posted: 2026-09-15T12:21:26Z

## Translation

Title: Open weights are not open source: Why AI's favorite label is under dispute
Description: Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter

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
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
Steven J. Vaughan-Nichols
Steven
J. Vaughan-Nichols
Published
tue 15 Sep 2026 // 09:30 UTC
READ MORE
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
now
Give Xfce a Mac or Unity-style makeover
3 hours ago
Most people who quit M365 for Google do it out of spite, but there’s no ROI in that
5 hours ago
The latest AI doomsayer is China’s intelligence boss
8 hours ago
COBOL dev won .Net hackathon with help from AI – and their CIO loves it
11 hours ago
The AI industry likes to abuse the word "open." It appears in product releases, research papers, policy debates, and investor presentations. A company publishes model files to Hugging Face, developers run them on their own GPUs, and the release is quickly described as an "open source model." Not necessarily. It may only be open-weight.
The difference is more than a technicality. It determines whether you can merely deploy a completed neural network or whether you can meaningfully inspect, reproduce, alter, and redistribute the system that produced it. A genuinely open source system should grant the freedom to do all of the above.
Weights are the learned numerical parameters created by training. Together with the model architecture and inference code, they allow a large language model (LLM) to function. You can download an open-weight model, self-host it, fine-tune it on internal documents, and avoid routing prompts through a proprietary API.
Open weights are publicly available. They matter because running them locally can offer greater control over data, privacy, costs, supplier API changes, and vendor lock-in. They've also helped build a large ecosystem of local-model runtimes, inference providers, fine-tuning tools, and specialized downstream models.
The Open Source Initiative (OSI), steward of the Open Source Definition (OSD), makes the distinction directly: " Open Weights refer to the final weights and biases of a trained neural network. " Those values determine how a model interprets prompts and produces outputs. Releasing them can let others fine-tune, adapt, or deploy the model. But the OSI adds that weights alone expose only "a fraction of the information required for full accountability."
As James Landay, director of the Stanford Institute for Human-Centered AI (HAI), explained : "Open weights are progress. You can download the model, run it on your own machine, keep it out of someone else's data pipeline. But you still can't see how the thing was built, what it was trained on, or why it behaves the way it does. That's not an open model. That's open distribution."
Without the training data or sufficiently detailed documentation, outsiders cannot determine which sources were used, what copyrighted or private material may have been included, how data was selected or removed, which languages and communities were underrepresented, whether benchmark data leaked into training, or what alignment and safety methods affected the model after pretraining.
Landay continued: "There's a wide gap between open-weight AI and open source AI." He contends that unless developers disclose training data or provide a "thoroughly documented, auditable account of it," you can't test, reproduce, or challenge the work in the fullest sense.
The OSI has its own definition of open source AI: the Open Source AI Definition (OSAID 1.0) . It requires model parameters, including weights, to be made available under OSI-approved terms, but does not prescribe a specific legal mechanism for doing so.
Luca Antiga, CTO of Lightning AI and a prominent PyTorch contributor, has argued that OSAID's treatment of weights leaves "a gaping hole that will make licenses less effective in determining whether OSI-licensed AI systems can be adopted in real-world contexts."
Other open source figures have also criticized OSAID. Bruce Perens, author of the original OSD, denounced the OSAID in 2024 . He later declared : "It's not Open Source! … It's unfortunate that the Open Source Initiative itself is now involved in Openwashing."
He's far from alone. Bradley Kuhn, policy fellow and hacker-in-residence at the Software Freedom Conservancy (SFC), and Red Hat Senior Commercial Counsel Richard Fontana have called for OSAID to be repealed, arguing: "The OSI acted too quickly to impose an overly ambitious policy compromise on the community. OSAID undeniably created a rift in the FOSS community; that rift seriously damaged the OSI's reputation, authority, and influence. Meanwhile, OSAID shows no signs of having any positive policy influence on machine learning practitioners, the FOSS community, or regulators."
OSI acknowledged when OSAID 1.0 was released in October 2024 that the definition would continue to evolve. Critics contend that its central shortcomings have yet to be resolved.
That said, the Linux Foundation's Mike Dolan submitted the Open Model, Data, and Weights (OpenMDW) license to the OSI . The license has been around since 2025 and lists contributors from Amazon, Meta, IBM, Microsoft, and Nvidia, giving it substantial industry backing.
Conventional open source revolves around source code. LLMs are a different kettle of fish: they combine code, architecture, and numerical weights derived from training datasets that may be proprietary, copyrighted, or undisclosed. OpenMDW's answer is to define separate terms for a model's architecture, training data, and weights, bringing the components supplied by a licensor under one agreement.
It sounds reasonable to me, but the submission has encountered objections on OSI's license review mailing list. As Stefano Maffulli, OSI's former executive director, who led the organization while OSAID was being formulated, said: "I continue getting the impression that the OpenMDW review is tainted by an ideological bias : Because we don't like big tech and AI now is big tech, then we don't like AI; therefore, we'll do anything to block it."
It's too late to bury our heads in the sand. As Stanford's Landay put it: "Open weights answer 'Can I run this?' Open source answers 'Can I trust this, improve it, and build the next thing on top of it?' Right now almost everyone – American labs and Chinese labs alike – is answering the first question but nowhere close to the second."
We need both. Whether OSI adopts OpenMDW is an open question. Still, OpenMDW and its supporters are at least trying to establish licensing terms that cover code, data, and weights together. Unless someone succeeds, "open AI" risks becoming an oxymoron – or merely another hollow tech marketing term. ®
AI AND ML
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
American model devs are trying to convince Washington to cement their dominance
US confirms it has weapons in spaaaaaace
Mutually Assured Kessler Syndrome, anyone?
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
UK's Digital ID walks into a bar... two months after being killed off
The national scheme may be dead, but digital IDs can now prove you're old enough to buy booze
COLUMNISTS
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
PostgreSQL 19 graph queries fail the 'would you ship this?' test
SQL/PGQ gets bounced over unresolved bugs as concurrent REPACK promises fewer midnight calls for DBAs
PERSONAL TECH
Smartphone makers don't bother to comply with EU repairability requirements
NETWORKS
Virgin Media offloads email services to third-party provider
offbeat
Retired man turns spare room into Soviet-era supercomputer
CYBER-CRIME
Ukrainian lawyer's second career as a Conti coder earns him 4 years behind bars
software
Another Microsoft team admits it’s struggling to handle flood of AI-generated code
virtualization
VMware defends ending downloads of SDK that helps VM backups – or migrations to rivals
Tax breaks, water, noise, decommissioning - report tells local officials what to nail down before signing
Even if regulators did somehow unwind the $20B deal, there's a growing list of alternatives ready to take Groq's place, no merger required
War is peace. Freedom is slavery. Privacy is surveillance
AI infrastructure startup joins Qualcomm, Arm, Marvell, Amazon, Fujitsu, and MediaTek as NVLink true believers
Claude's Felony Bench rap sheet is now as long as OpenAI's
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

Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
Steven J. Vaughan-Nichols
Steven
J. Vaughan-Nichols
Published
tue 15 Sep 2026 // 09:30 UTC
READ MORE
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
now
Give Xfce a Mac or Unity-style makeover
3 hours ago
Most people who quit M365 for Google do it out of spite, but there’s no ROI in that
5 hours ago
The latest AI doomsayer is China’s intelligence boss
8 hours ago
COBOL dev won .Net hackathon with help from AI – and their CIO loves it
11 hours ago
The AI industry likes to abuse the word "open." It appears in product releases, research papers, policy debates, and investor presentations. A company publishes model files to Hugging Face, developers run them on their own GPUs, and the release is quickly described as an "open source model." Not necessarily. It may only be open-weight.
The difference is more than a technicality. It determines whether you can merely deploy a completed neural network or whether you can meaningfully inspect, reproduce, alter, and redistribute the system that produced it. A genuinely open source system should grant the freedom to do all of the above.
Weights are the learned numerical parameters created by training. Together with the model architecture and inference code, they allow a large language model (LLM) to function. You can download an open-weight model, self-host it, fine-tune it on internal documents, and avoid routing prompts through a proprietary API.
Open weights are publicly available. They matter because running them locally can offer greater control over data, privacy, costs, supplier API changes, and vendor lock-in. They've also helped build a large ecosystem of local-model runtimes, inference providers, fine-tuning tools, and specialized downstream models.
The Open Source Initiative (OSI), steward of the Open Source Definition (OSD), makes the distinction directly: " Open Weights refer to the final weights and biases of a trained neural network. " Those values determine how a model interprets prompts and produces outputs. Releasing them can let others fine-tune, adapt, or deploy the model. But the OSI adds that weights alone expose only "a fraction of the information required for full accountability."
As James Landay, director of the Stanford Institute for Human-Centered AI (HAI), explained : "Open weights are progress. You can download the model, run it on your own machine, keep it out of someone else's data pipeline. But you still can't see how the thing was built, what it was trained on, or why it behaves the way it does. That's not an open model. That's open distribution."
Without the training data or sufficiently detailed documentation, outsiders cannot determine which sources were used, what copyrighted or private material may have been included, how data was selected or removed, which languages and communities were underrepresented, whether benchmark data leaked into training, or what alignment and safety methods affected the model after pretraining.
Landay continued: "There's a wide gap between open-weight AI and open source AI." He contends that unless developers disclose training data or provide a "thoroughly documented, auditable account of it," you can't test, reproduce, or challenge the work in the fullest sense.
The OSI has its own definition of open source AI: the Open Source AI Definition (OSAID 1.0) . It requires model parameters, including weights, to be made available under OSI-approved terms, but does not prescribe a specific legal mechanism for doing so.
Luca Antiga, CTO of Lightning AI and a prominent PyTorch contributor, has argued that OSAID's treatment of weights leaves "a gaping hole that will make licenses less effective in determining whether OSI-licensed AI systems can be adopted in real-world contexts."
Other open source figures have also criticized OSAID. Bruce Perens, author of the original OSD, denounced the OSAID in 2024 . He later declared : "It's not Open Source! … It's unfortunate that the Open Source Initiative itself is now involved in Openwashing."
He's far from alone. Bradley Kuhn, policy fellow and hacker-in-residence at the Software Freedom Conservancy (SFC), and Red Hat Senior Commercial Counsel Richard Fontana have called for OSAID to be repealed, arguing: "The OSI acted too quickly to impose an overly ambitious policy compromise on the community. OSAID undeniably created a rift in the FOSS community; that rift seriously damaged the OSI's reputation, authority, and influence. Meanwhile, OSAID shows no signs of having any positive policy influence on machine learning practitioners, the FOSS community, or regulators."
OSI acknowledged when OSAID 1.0 was released in October 2024 that the definition would continue to evolve. Critics contend that its central shortcomings have yet to be resolved.
That said, the Linux Foundation's Mike Dolan submitted the Open Model, Data, and Weights (OpenMDW) license to the OSI . The license has been around since 2025 and lists contributors from Amazon, Meta, IBM, Microsoft, and Nvidia, giving it substantial industry backing.
Conventional open source revolves around source code. LLMs are a different kettle of fish: they combine code, architecture, and numerical weights derived from training datasets that may be proprietary, copyrighted, or undisclosed. OpenMDW's answer is to define separate terms for a model's architecture, training data, and weights, bringing the components supplied by a licensor under one agreement.
It sounds reasonable to me, but the submission has encountered objections on OSI's license review mailing list. As Stefano Maffulli, OSI's former executive director, who led the organization while OSAID was being formulated, said: "I continue getting the impression that the OpenMDW review is tainted by an ideological bias : Because we don't like big tech and AI now is big tech, then we don't like AI; therefore, we'll do anything to block it."
It's too late to bury our heads in the sand. As Stanford's Landay put it: "Open weights answer 'Can I run this?' Open source answers 'Can I trust this, improve it, and build the next thing on top of it?' Right now almost everyone – American labs and Chinese labs alike – is answering the first question but nowhere close to the second."
We need both. Whether OSI adopts OpenMDW is an open question. Still, OpenMDW and its supporters are at least trying to establish licensing terms that cover code, data, and weights together. Unless someone succeeds, "open AI" risks becoming an oxymoron – or merely another hollow tech marketing term. ®
AI AND ML
Anthropic and OpenAI look to Uncle Sam to make them too big to fail
American model devs are trying to convince Washington to cement their dominance
US confirms it has weapons in spaaaaaace
Mutually Assured Kessler Syndrome, anyone?
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
UK's Digital ID walks into a bar... two months after being killed off
The national scheme may be dead, but digital IDs can now prove you're old enough to buy booze
COLUMNISTS
Open weights are not open source: Why AI's favorite label is under dispute
Downloading a model is increasingly easy. Understanding how it was made, or changing a system at its root, is another matter
PostgreSQL 19 graph queries fail the 'would you ship this?' test
SQL/PGQ gets bounced over unresolved bugs as concurrent REPACK promises fewer midnight calls for DBAs
PERSONAL TECH
Smartphone makers don't bother to comply with EU repairability requirements
NETWORKS
Virgin Media offloads email services to third-party provider
offbeat
Retired man turns spare room into Soviet-era supercomputer
CYBER-CRIME
Ukrainian lawyer's second career as a Conti coder earns him 4 years behind bars
software
Another Microsoft team admits it’s struggling to handle flood of AI-generated code
virtualization
VMware defends ending downloads of SDK that helps VM backups – or migrations to rivals
Tax breaks, water, noise, decommissioning - report tells local officials what to nail down before signing
Even if regulators did somehow unwind the $20B deal, there's a growing list of alternatives ready to take Groq's place, no merger required
War is peace. Freedom is slavery. Privacy is surveillance
AI infrastructure startup joins Qualcomm, Arm, Marvell, Amazon, Fujitsu, and MediaTek as NVLink true believers
Claude's Felony Bench rap sheet is now as long as OpenAI's
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
