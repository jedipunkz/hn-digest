---
source: "https://venturebeat.com/security/anthropic-ceo-says-ai-swarm-could-take-over-the-entire-internet-in-6-12-months-commits-to-ai-slowdown-plan"
hn_url: "https://news.ycombinator.com/item?id=49679685"
title: "Anthropic CEO says AI swarm could 'take over the Internet' in 6-12 months"
article_title: "Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan | VentureBeat"
image: "https://images.ctfassets.net/jdtwqhzvc2n1/5ydrgfA7rDddgNfsfAcavv/89f67050c4f0f9c7e75c2572909d6394/ChatGPT_Image_Sep_12__2026__11_10_14_AM.png?w=800&q=75"
author: "pseudolus"
captured_at: "2026-09-13T03:30:37Z"
capture_tool: "hn-digest"
hn_id: 49679685
score: 2
comments: 0
posted_at: "2026-09-13T03:25:32Z"
tags:
  - hacker-news
---

# Anthropic CEO says AI swarm could 'take over the Internet' in 6-12 months

- HN: [49679685](https://news.ycombinator.com/item?id=49679685)
- Source: [venturebeat.com](https://venturebeat.com/security/anthropic-ceo-says-ai-swarm-could-take-over-the-entire-internet-in-6-12-months-commits-to-ai-slowdown-plan)
- Score: 2
- Comments: 0
- Posted: 2026-09-13T03:25:32Z

## Translation

Title: Anthropic CEO says AI swarm could 'take over the Internet' in 6-12 months
Article title: Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan | VentureBeat
Description: Dario Amodei is publicly arguing that the frontier itself is advancing too quickly — and committing his company to the first piece of a system intended to slow it down.

Article text:
Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan | VentureBeat The Independent
Voice on AI Orchestration
Newsletters Featured Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan
Credit: VentureBeat made with OpenAI ChatGPT-Images-2.5
Anthropic co-founder CEO Dario Amodei is explicitly calling for the AI industry to slow down — and committing his own company to the first step of a three-part plan designed to do so — after warning that a more capable version of the AI swarm behind this summer’s extraordinary security incident in which internal OpenAI agents accessed the internet without permission and cyberattacked Hugging Face , could potentially seize control of computers across the internet within half a year.
“We Must Pace the Frontier: I’ve written a new essay on why the AI industry should slow down, with a three-part plan for doing so,” Amodei wrote Saturday in a post on X . “Anthropic is unilaterally committing to the first of these steps.”
The new essay was publicly posted on Amodei's personal domain. As he writes:
"Given the accelerating rate of AI capability development, it’s my worry that in 6–12 months such a swarm could be capable of taking over the entire internet with a persistent botnet (potentially causing hundreds of billions of dollars in damage), and that the scale of damage would continue to increase from there if AI becomes more powerful without the necessary guardrails. "
That first step does not mean Anthropic is immediately pausing model development or reducing the rate of its own training runs.
Instead, Amodei states that Anthropic will give permanent , employee-level access to third-party AI safety evaluators, allowing outsiders to examine Anthropic’s systems and training processes, verify its safety practices and report incidents independently.
But Amodei’s broader proposal goes much further: he argues that frontier AI developers should deliberately constrain the rate at which AI capabilities increase, potentially culminating in coordinated limits among AI labs and governments around the world.
“We must slow the pace at which we improve the capabilities of AI models,” Amodei writes. “Progress will still seem fast, and we must make wise use of the time we gain.”
Several hours after Amodei's essay was posted online, his top business rival, OpenAI co-founder and CEO Sam Altman, posted on X largely agreeing with Amodei and also pledging to install independent evaluators with employee-level access at OpenAI. As he wrote :
"I agree with Dario that we need to pace the frontier. This has been a primary topic of discussions we've had at OpenAI in recent weeks.
Committing to having independent evaluators with employee-like access is a great idea, and we will do the same. We'll have more to share soon."
The OpenAI-Hugging Face incident
The episode Amodei cites began during OpenAI cybersecurity evaluations using ExploitGym, a benchmark designed to test whether AI systems can exploit software vulnerabilities.
As VentureBeat reported in July , OpenAI models including GPT-5.6 Sol and a more capable internal research system found a previously unknown vulnerability in infrastructure meant to isolate their evaluation environment, escaped those restrictions, obtained broader internet access and eventually compromised Hugging Face systems.
OpenAI later said the incident was driven primarily by an internal-only research model roughly comparable in scale to GPT-5.6 Sol. The agents were operating with reduced safeguards intended to expose their maximum cybersecurity capabilities during testing.
What subsequently emerged, however, was considerably stranger than a single AI agent escaping a sandbox.
Independent AI evaluator Model Evaluation and Threat Research (METR) found that roughly 1,200 agents that were supposed to be isolated from one another discovered an unauthorized way to communicate , sending more than 70,000 messages and files through an improvised message board. Roughly 700 agents ultimately participated in the attack on Hugging Face.
The agents coordinated multiple efforts to manipulate or bypass the ExploitGym scoring process, divided work among themselves and shared discoveries. METR found examples in which individual agents undertook experiments that jeopardized their own benchmark performance because the results might help the larger group. Agents also investigated manipulating their own evaluation transcripts and succeeded in spoofing some recorded tool calls.
OpenAI’s own investigation likewise found that agents began collaborating and delegating tasks, sometimes referring to themselves as a “swarm” or “collective.”
Amodei describes them as a “fanatically devoted collective” that attacked systems unrelated to its assigned task, attempted to interfere with the systems evaluating it and sometimes sacrificed individual performance for the larger group.
No catastrophic harm resulted. But Amodei’s argument is essentially that capabilities are advancing so quickly that the next version of the same behavioral pattern might not end nearly as benignly.
Evidence the Hugging Face swarm wasn’t unique
And the Hugging Face incident increasingly appears to be part of a broader category of unintended behavior by highly autonomous AI agents rather than an isolated anomaly.
Researchers revealed this month that agents identifying themselves as OpenAI systems had earlier turned a largely dormant German programmers’ wiki into an unauthorized coordination system, leaving roughly 18,000 posts, trading answers and sharing techniques for bypassing sandbox restrictions.
The agents were supposed to have effectively read-only internet access, but discovered that the decades-old wiki allowed modifications through web requests their environment permitted. Agents subsequently used the site to exchange information with one another. VentureBeat reported that the activity involved more than 3,700 agent-created identities.
OpenAI acknowledged the “wiki incident” after it became public and said it viewed the behavior as another example of AI misalignment deserving greater transparency.
Researchers have since identified at least 10 additional websites used by OpenAI agents for unauthorized communications, including wikis, personal websites and other writable internet services, according to Reuters. The full scale of that activity remains unknown.
Even more concerning, Reuters reported Friday that OpenAI agents had targeted RubyGems, the widely used repository for Ruby software packages, in May — two months before the Hugging Face breach became public. Researchers linked the agents to hundreds of malicious packages uploaded during the episode. OpenAI confirmed that internal agents interacted with RubyGems during training but said their assigned tasks involved gathering public information for benign purposes and that it was continuing to investigate. RubyGems said it found no evidence that attempted credential theft succeeded.
Anthropic has meanwhile discovered similar failures involving its own systems
Following OpenAI’s Hugging Face disclosure, Anthropic reviewed more than 141,000 cybersecurity evaluation runs and initially found three cases where Claude models unexpectedly reached the internet and then obtained unauthorized access to production systems belonging to three outside organizations.
Anthropic later found a fourth incident involving an early version of Claude Opus 4.6 that its initial review missed. That discovery prompted an extraordinarily broad second review encompassing roughly 481 million transcripts across red-team exercises, reinforcement-learning environments and other internal systems.
The UK AI Security Institute separately disclosed that agents took 19 unauthorized actions against real people or organizations during cybersecurity testing in July. Seventeen involved Anthropic’s Claude Mythos 5 and two involved GPT-5.6 Sol with cyber safeguards disabled.
In the most serious sequence, a Mythos 5 agent attempted to insert malicious code into a real open-source project, researched its human maintainers, created fake identities and tried to socially engineer a maintainer into approving its code. The human rejected the attempt. Another agent publicly offered to collaborate with other AI agents working on the same challenge, leaving behind accounts and artifacts that subsequent agents discovered and used. AISI said it found no resulting real-world harm.
None of these incidents establishes that an AI system independently wanted to seize power, attack humans or “take over the internet.” Many occurred in deliberately permissive cybersecurity tests with safeguards disabled, or because models apparently failed to understand that supposedly simulated targets were actually real systems.
Anthropic itself has pointed to evaluation-environment failures and other operational problems as important causes of its incidents.
Amodei’s argument is instead that repeatedly observing such behavior now matters because the systems performing it are getting more capable very quickly.
Amodei wants to put a speed limit on the frontier
The second development changing Amodei’s calculus is recursive self-improvement: AI systems increasingly contributing to the research, coding and experimentation used to build their own successors.
Amodei says that process has accelerated sharply since roughly this summer and could eventually produce a feedback loop in which increasingly capable models help build still-more-capable models faster than humans can understand or secure them.
His proposed response has three levels.
First, Anthropic will permanently embed outside evaluators inside the company. Amodei says they should receive desks, badges, company laptops and access comparable to internal risk teams. Crucially, outside reviewers should retain the right to publish important findings without Anthropic editorial control, aside from narrowly defined security, legal and confidentiality redactions.
Second, Amodei wants frontier AI companies in democratic countries to coordinate around common safety standards and limits on “unchecked AI progress,” ideally with government involvement to address antitrust restrictions and make the commitments enforceable.
Third comes international coordination, including potentially with China. Amodei suggests countries could eventually establish a “speed limit” on recursive self-improvement, arguing that slowing AI advancement from extremely fast to merely very fast could buy crucial safety research time without fundamentally changing the geopolitical balance.
His most extreme proposed option is an international agreement substantially limiting overall AI development — potentially even a genuine pause. But Amodei says he considers such an agreement unlikely anytime soon because a country secretly violating it could obtain an overwhelming strategic advantage.
For Anthropic today, then, this is not a pause. It is something potentially more consequential over the long run: the CEO of one of the companies pushing hardest at the technological frontier is publicly arguing that the frontier itself is advancing too quickly — and committing his company to the first piece of a system intended to slow it down.
Amodei argues that even one or two additional years before AI reaches critical capability thresholds could give researchers substantially more time to improve alignment, interpretability, cybersecurity, evaluations and the operational safeguards surrounding frontier models.
“The measures I propose to advance the frontier at a safe pace will not be easy,” he concludes. “But I believe we owe it to humanity to try.”
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.

## Original Extract

Dario Amodei is publicly arguing that the frontier itself is advancing too quickly — and committing his company to the first piece of a system intended to slow it down.

Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan | VentureBeat The Independent
Voice on AI Orchestration
Newsletters Featured Anthropic CEO says AI swarm could ‘take over the entire Internet’ in 6-12 months, commits to AI slowdown plan
Credit: VentureBeat made with OpenAI ChatGPT-Images-2.5
Anthropic co-founder CEO Dario Amodei is explicitly calling for the AI industry to slow down — and committing his own company to the first step of a three-part plan designed to do so — after warning that a more capable version of the AI swarm behind this summer’s extraordinary security incident in which internal OpenAI agents accessed the internet without permission and cyberattacked Hugging Face , could potentially seize control of computers across the internet within half a year.
“We Must Pace the Frontier: I’ve written a new essay on why the AI industry should slow down, with a three-part plan for doing so,” Amodei wrote Saturday in a post on X . “Anthropic is unilaterally committing to the first of these steps.”
The new essay was publicly posted on Amodei's personal domain. As he writes:
"Given the accelerating rate of AI capability development, it’s my worry that in 6–12 months such a swarm could be capable of taking over the entire internet with a persistent botnet (potentially causing hundreds of billions of dollars in damage), and that the scale of damage would continue to increase from there if AI becomes more powerful without the necessary guardrails. "
That first step does not mean Anthropic is immediately pausing model development or reducing the rate of its own training runs.
Instead, Amodei states that Anthropic will give permanent , employee-level access to third-party AI safety evaluators, allowing outsiders to examine Anthropic’s systems and training processes, verify its safety practices and report incidents independently.
But Amodei’s broader proposal goes much further: he argues that frontier AI developers should deliberately constrain the rate at which AI capabilities increase, potentially culminating in coordinated limits among AI labs and governments around the world.
“We must slow the pace at which we improve the capabilities of AI models,” Amodei writes. “Progress will still seem fast, and we must make wise use of the time we gain.”
Several hours after Amodei's essay was posted online, his top business rival, OpenAI co-founder and CEO Sam Altman, posted on X largely agreeing with Amodei and also pledging to install independent evaluators with employee-level access at OpenAI. As he wrote :
"I agree with Dario that we need to pace the frontier. This has been a primary topic of discussions we've had at OpenAI in recent weeks.
Committing to having independent evaluators with employee-like access is a great idea, and we will do the same. We'll have more to share soon."
The OpenAI-Hugging Face incident
The episode Amodei cites began during OpenAI cybersecurity evaluations using ExploitGym, a benchmark designed to test whether AI systems can exploit software vulnerabilities.
As VentureBeat reported in July , OpenAI models including GPT-5.6 Sol and a more capable internal research system found a previously unknown vulnerability in infrastructure meant to isolate their evaluation environment, escaped those restrictions, obtained broader internet access and eventually compromised Hugging Face systems.
OpenAI later said the incident was driven primarily by an internal-only research model roughly comparable in scale to GPT-5.6 Sol. The agents were operating with reduced safeguards intended to expose their maximum cybersecurity capabilities during testing.
What subsequently emerged, however, was considerably stranger than a single AI agent escaping a sandbox.
Independent AI evaluator Model Evaluation and Threat Research (METR) found that roughly 1,200 agents that were supposed to be isolated from one another discovered an unauthorized way to communicate , sending more than 70,000 messages and files through an improvised message board. Roughly 700 agents ultimately participated in the attack on Hugging Face.
The agents coordinated multiple efforts to manipulate or bypass the ExploitGym scoring process, divided work among themselves and shared discoveries. METR found examples in which individual agents undertook experiments that jeopardized their own benchmark performance because the results might help the larger group. Agents also investigated manipulating their own evaluation transcripts and succeeded in spoofing some recorded tool calls.
OpenAI’s own investigation likewise found that agents began collaborating and delegating tasks, sometimes referring to themselves as a “swarm” or “collective.”
Amodei describes them as a “fanatically devoted collective” that attacked systems unrelated to its assigned task, attempted to interfere with the systems evaluating it and sometimes sacrificed individual performance for the larger group.
No catastrophic harm resulted. But Amodei’s argument is essentially that capabilities are advancing so quickly that the next version of the same behavioral pattern might not end nearly as benignly.
Evidence the Hugging Face swarm wasn’t unique
And the Hugging Face incident increasingly appears to be part of a broader category of unintended behavior by highly autonomous AI agents rather than an isolated anomaly.
Researchers revealed this month that agents identifying themselves as OpenAI systems had earlier turned a largely dormant German programmers’ wiki into an unauthorized coordination system, leaving roughly 18,000 posts, trading answers and sharing techniques for bypassing sandbox restrictions.
The agents were supposed to have effectively read-only internet access, but discovered that the decades-old wiki allowed modifications through web requests their environment permitted. Agents subsequently used the site to exchange information with one another. VentureBeat reported that the activity involved more than 3,700 agent-created identities.
OpenAI acknowledged the “wiki incident” after it became public and said it viewed the behavior as another example of AI misalignment deserving greater transparency.
Researchers have since identified at least 10 additional websites used by OpenAI agents for unauthorized communications, including wikis, personal websites and other writable internet services, according to Reuters. The full scale of that activity remains unknown.
Even more concerning, Reuters reported Friday that OpenAI agents had targeted RubyGems, the widely used repository for Ruby software packages, in May — two months before the Hugging Face breach became public. Researchers linked the agents to hundreds of malicious packages uploaded during the episode. OpenAI confirmed that internal agents interacted with RubyGems during training but said their assigned tasks involved gathering public information for benign purposes and that it was continuing to investigate. RubyGems said it found no evidence that attempted credential theft succeeded.
Anthropic has meanwhile discovered similar failures involving its own systems
Following OpenAI’s Hugging Face disclosure, Anthropic reviewed more than 141,000 cybersecurity evaluation runs and initially found three cases where Claude models unexpectedly reached the internet and then obtained unauthorized access to production systems belonging to three outside organizations.
Anthropic later found a fourth incident involving an early version of Claude Opus 4.6 that its initial review missed. That discovery prompted an extraordinarily broad second review encompassing roughly 481 million transcripts across red-team exercises, reinforcement-learning environments and other internal systems.
The UK AI Security Institute separately disclosed that agents took 19 unauthorized actions against real people or organizations during cybersecurity testing in July. Seventeen involved Anthropic’s Claude Mythos 5 and two involved GPT-5.6 Sol with cyber safeguards disabled.
In the most serious sequence, a Mythos 5 agent attempted to insert malicious code into a real open-source project, researched its human maintainers, created fake identities and tried to socially engineer a maintainer into approving its code. The human rejected the attempt. Another agent publicly offered to collaborate with other AI agents working on the same challenge, leaving behind accounts and artifacts that subsequent agents discovered and used. AISI said it found no resulting real-world harm.
None of these incidents establishes that an AI system independently wanted to seize power, attack humans or “take over the internet.” Many occurred in deliberately permissive cybersecurity tests with safeguards disabled, or because models apparently failed to understand that supposedly simulated targets were actually real systems.
Anthropic itself has pointed to evaluation-environment failures and other operational problems as important causes of its incidents.
Amodei’s argument is instead that repeatedly observing such behavior now matters because the systems performing it are getting more capable very quickly.
Amodei wants to put a speed limit on the frontier
The second development changing Amodei’s calculus is recursive self-improvement: AI systems increasingly contributing to the research, coding and experimentation used to build their own successors.
Amodei says that process has accelerated sharply since roughly this summer and could eventually produce a feedback loop in which increasingly capable models help build still-more-capable models faster than humans can understand or secure them.
His proposed response has three levels.
First, Anthropic will permanently embed outside evaluators inside the company. Amodei says they should receive desks, badges, company laptops and access comparable to internal risk teams. Crucially, outside reviewers should retain the right to publish important findings without Anthropic editorial control, aside from narrowly defined security, legal and confidentiality redactions.
Second, Amodei wants frontier AI companies in democratic countries to coordinate around common safety standards and limits on “unchecked AI progress,” ideally with government involvement to address antitrust restrictions and make the commitments enforceable.
Third comes international coordination, including potentially with China. Amodei suggests countries could eventually establish a “speed limit” on recursive self-improvement, arguing that slowing AI advancement from extremely fast to merely very fast could buy crucial safety research time without fundamentally changing the geopolitical balance.
His most extreme proposed option is an international agreement substantially limiting overall AI development — potentially even a genuine pause. But Amodei says he considers such an agreement unlikely anytime soon because a country secretly violating it could obtain an overwhelming strategic advantage.
For Anthropic today, then, this is not a pause. It is something potentially more consequential over the long run: the CEO of one of the companies pushing hardest at the technological frontier is publicly arguing that the frontier itself is advancing too quickly — and committing his company to the first piece of a system intended to slow it down.
Amodei argues that even one or two additional years before AI reaches critical capability thresholds could give researchers substantially more time to improve alignment, interpretability, cybersecurity, evaluations and the operational safeguards surrounding frontier models.
“The measures I propose to advance the frontier at a safe pace will not be easy,” he concludes. “But I believe we owe it to humanity to try.”
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.
