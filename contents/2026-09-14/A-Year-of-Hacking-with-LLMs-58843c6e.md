---
source: "https://sites.google.com/site/zhiniangpeng/blogs/Hacking-with-LLMs-Eng"
hn_url: "https://news.ycombinator.com/item?id=49700571"
title: "A Year of Hacking with LLMs"
article_title: "A Year of Hacking with LLMs"
image: ""
author: "homarp"
captured_at: "2026-09-14T18:02:37Z"
capture_tool: "hn-digest"
hn_id: 49700571
score: 1
comments: 0
posted_at: "2026-09-14T17:22:39Z"
tags:
  - hacker-news
---

# A Year of Hacking with LLMs

- HN: [49700571](https://news.ycombinator.com/item?id=49700571)
- Source: [sites.google.com](https://sites.google.com/site/zhiniangpeng/blogs/Hacking-with-LLMs-Eng)
- Score: 1
- Comments: 0
- Posted: 2026-09-14T17:22:39Z

## Translation

Title: A Year of Hacking with LLMs
Description: This blog is a summary of my talk at the Offbyone 2026 cybersecurity conference. It records some of my thoughts as a cybersecurity researcher after spending a year using LLMs for research.
Slides: https://github.com/edwardzpeng/presentations/tree/main/offbyone%202026
Chinese version:

Article text:
A Year of Hacking with LLMs Search this site Embedded Files Skip to main content Skip to navigation Home
This blog is a summary of my talk at the Offbyone 2026 cybersecurity conference. It records some of my thoughts as a cybersecurity researcher after spending a year using LLMs for research.
Slides: https://github.com/edwardzpeng/presentations/tree/main/offbyone%202026
Chinese version: https://sites.google.com/site/zhiniangpeng/blogs/Hacking-with-LLMs
Over the past year, I have spent a great deal of time on one thing: trying to make LLM genuinely participate in cybersecurity research. The results have continued to surprise me.
In 2023, our first attempt to use an LLM for bug huntingfailed. By the end of 2025, we were already using LLMs with agents to analyze vulnerabilities, generate PoCs, and validate them on real devices. After entering 2026, some vulnerability research had even become close to highly automated. I then applied the same methods to cryptography and other cybersecurity projects, where the performance of LLMs was equally astonishing.
## First Attempt: Bug hunting with Code Llama
Our idea at the time was very straightforward: if an LLM could understand code, then after seeing enough vulnerable code, it should be able to determine whether a piece of code contained a vulnerability and identify the specific bug.
We collected approximately 40,000 Python vulnerability samples, 37,000 Java vulnerability samples, and 40,000 clean-code samples to fine-tune Code Llama. We hoped that, given a piece of code as input, it would output either “vulnerable,” together with the specific vulnerability, or “not vulnerable.”
It performed very well on the benchmark. Among 24 test classes, eight achieved a score of 100%, while 23 scored above 90%.
Then we tested it on real-world code. The result was almost useless.
The model produced many findings that looked like vulnerabilities but were not actually exploitable. The output was noisy, non-actionable, and impossible to validate. The beautiful benchmark numbers did not translate into real vulnerabilities.
Looking back, the reasons for the failure are now quite clear:
1. **The target was wrong.** Simple LLM-based recognition of vulnerability patterns is not the same as the process of vulnerability research.
2. **The capability was wrong.** More SFT data does not mean deeper reasoning.
3. **It overfit CVE patterns.** Common code features in CVEs are not the same as the security invariants that a system actually needs to preserve.
4. **There was not enough context.** A code snippet is not the complete context of the target system.
5. **There was no threat model.** Without security boundaries, adversary capabilities, entry points, and security goals, it is difficult to determine whether a behavior is truly a vulnerability.
6. **It could not validate its findings.** The model could not run, debug, and verify its own conclusions.
7. **There was no continuous research process.** A single question and answer is not research. Research requires a cycle of hypothesis, evidence, and revision.
Simply put, we treated vulnerability discovery as a classification problem.
Our idea was still too simple, and the LLM ecosystem was not mature enough. At the time, it could not yet be used effectively for vulnerability research.
## What Had Changed by the End of 2025?
Two years later, the situation was completely different.
First, the models themselves had become much stronger. Models such as GPT-5.2 and Opus 4.5 had much better reasoning capabilities and no longer merely recognized superficial patterns in code.
Second, context windows had become longer. More complete code repository, call chains, object lifetimes, and trust boundaries can be put into the context. The model no longer saw only an isolated function; it saw a whole system.
Agent workflows had also matured. Approaches such as ReAct, Plan-and-Execute, Chain of Thought, and Multi-Agent systems allowed models to break complex tasks into many steps and adjust their plans based on intermediate results.
Finally, more tools had become available. MCP, CLI, A2A, skills, and various deterministic analysis and execution scripts allowed models to operate on real targets instead of merely returning an answer in a chat box.
The real breakthrough, therefore, was not only a better model. It was a more complete system built around the model.
## Second Attempt: LLM + Agent
At the end of 2025, we tried again to use an LLM for bug hunting. This time, the target was access-control vulnerabilities in customized Android systems.
This was an ideal scenario for testing an agent.
The Android ecosystem today includes not only phones, but also watches, bands, TVs, in-vehicle systems, and many kinds of IoT devices. Different vendors add large numbers of OEM services to their ROMs. These services usually have elevated privileges while exposing interfaces to ordinary apps.
In theory, those interfaces should protect sensitive data and privileged operations through mechanisms such as permissions, UIDs, signatures, and AppOps. In complex customized code, however, gaps frequently appear in permission validation. A low-privilege third-party app may consequently gain access to device identifiers, personal data, or system settings.
This research was conducted between December 2025 and January 2026 by X0ev, Wh1tc from Kunlun Lab and me.
### Preparing Tools for the Agent
We did not simply ask the model to “read the entire firmware and find vulnerabilities.” Instead, we first prepared a collection of deterministic tools:
- **Firmware Unpacking:** Extract APKs from the system, vendor, and product partitions.
- **JADX Decompilation:** Convert system APKs into readable Java code.
- **Manifest Scan:** Organize permissions, actions, and components.
- **Indexing:** Index file paths, packages, and classes.
- **Task Generation:** Select exported, high-risk system components.
- **ADB and Logcat:** Invoke real interfaces and collect access-control evidence.
The tools themselves are deterministic. The skill determines when and in what order the agent should use them.
For high-risk APIs, we asked the agent to complete two reasoning steps.
The first step was to understand the behavior: which identities, parameters, and system states could trigger a sensitive operation?
The second step was to determine authorization: did this operation exceed the privileges that the caller should have? If so, it became a candidate vulnerability.
These two questions may look simple, but they distinguish between something that merely “looks like a vulnerability” and something that genuinely violates a security boundary.
### Validation on Real Devices
A candidate vulnerability was not enough. The agent needed to generate a PoC, run it on a real device, collect Logcat output, and then revise and retry the PoC based on the failure results.
Common causes of failure included an incorrect call order, a missing required state, unmet data dependencies, and incorrect permissions or runtime environments. The agent’s job was not to produce the correct answer in a single attempt, but to continue revising its work after a failure.
Using Codex GPT-5.2 Extra High, we eventually tested real devices from OPPO, HONOR, Xiaomi, Samsung, Google, and other vendors. The results were:
- 110 confirmed vulnerabilities;
- an average cost of less than one US dollar per PoC;
- approximately two iterations per PoC on average.
This time, the LLM was no longer merely a chatbot.
In 2023, we asked it, “Does this code look like a known vulnerability pattern?”
By the end of 2025, we were asking an agent to understand a system, investigate a bug candidate, and validate it with a PoC. The AI used tools, performed scoped analysis, and generated reports. Researchers selected attack surfaces, built the harnesses and skills, assessed the value of the results, and took responsibility for the final conclusions.
This is vulnerability research conducted jointly by researchers and AI.
## From Patches to Vulnerability Variants: Highly Automated Bug Hunting
In the previous research, a human still needed to select the attack surface and build a harness and skill around a particular security invariant.
With a stronger model and agent, could AI complete this part as well?
If we had enough data, especially patches, public write-ups, and PoCs, the answer was beginning to become “yes.”
We built Diffract. This is a joint work with Yunpeng Tian, @___2st.
The core Diffract workflow is:
1. Collect public vulnerability write-ups and PoCs.
2. Build a foundational knowledge base and skill.
4. Extract vulnerability root causes.
5. Abstract and preserve vulnerability patterns.
6. Search for variants in other code.
8. Feed new results back into the knowledge base and skill.
This is no longer a one-off question and answer. It is a research loop capable of accumulating knowledge.
Using Claude Opus 4.6 through 4.8 on Microsoft Windows Kernel and Services, we obtained:
- more than 200 confirmed vulnerabilities;
- RCE, LPE, and information-leak vulnerabilities across different modules.
In fact, this has already become a form of token economics. In the short term, the more tokens you can burn, the more vulnerabilities you can find, because the backlog of legacy vulnerabilities is simply enormous. Windows vulnerability research has become quite boring in 2026. I hope Microsoft can quickly clear this huge backlog of vulnerabilities. Although they actually doing a great job, this month almost 1,000 vulnerabilities fixed.
### The State of Vulnerability Research in 2026
By 2026, my assessment of LLM-based bug hunting was:
- With previous-generation models, bug hunting works as long as the harness is sufficiently good. For example, we found many vulnerabilities using GPT-5.2 and GLM-5.2.
- With a sufficiently capable model, nearly fully automated bug hunting has become possible. Opus 4.6, Kimi K3, DeepSeek V4, GPT-5.4, and GLM 5.3 are already good enough.
- Today, with a SOTA model, all you need to do is make a wish to the model, and it can find vulnerabilities. There is no longer any barrier to entry.
- These three changes happened within six months. Things are changing extremely quickly.
- Bug hunting has evolved into a form of token economics.
In the past, the primary cost of bug hunting was researcher's time. Model capability, tokens, tools, and computing resources are now becoming part of a new cost structure.
This article does not discuss vulnerability exploitation, but the situation is basically the same.
### How to Use AI for Bug Hunting
How to use AI depends on what model you have. The main difference is the level at which the researcher needs to provide help:
- If the model is less capable, the researcher needs to break the task into smaller steps, prepare the context, tools, and tests, and frequently check and correct its work. Much of the work is actually using engineering methods to make up for the model's limitations and ensure that the whole process runs correctly.
- If the model is good enough, it can usually run the full process independently. The researcher mainly monitors the research process, provides background information, system access, the runtime environment, or debugging experience at key moments, and judges the final results.
- If you use a SOTA model, routine analysis and validation often need little help. The researcher should focus on where the model failed to find vulnerabilities: check what it overlooked, add runtime states or cross-module relationships that it could not see, change the search and validation methods, and let it continue trying where it previously failed.
Therefore, the stronger the model, the less the researcher needs to help it complete the process, and the more the researcher needs to help it push beyond the process's current capability boundary.
### About Vulnerability Exploitation
In fact, vulnerability exploitation is also being dramatically accelerated by AI, but that is not the foc

[truncated]

## Original Extract

This blog is a summary of my talk at the Offbyone 2026 cybersecurity conference. It records some of my thoughts as a cybersecurity researcher after spending a year using LLMs for research.
Slides: https://github.com/edwardzpeng/presentations/tree/main/offbyone%202026
Chinese version:

A Year of Hacking with LLMs Search this site Embedded Files Skip to main content Skip to navigation Home
This blog is a summary of my talk at the Offbyone 2026 cybersecurity conference. It records some of my thoughts as a cybersecurity researcher after spending a year using LLMs for research.
Slides: https://github.com/edwardzpeng/presentations/tree/main/offbyone%202026
Chinese version: https://sites.google.com/site/zhiniangpeng/blogs/Hacking-with-LLMs
Over the past year, I have spent a great deal of time on one thing: trying to make LLM genuinely participate in cybersecurity research. The results have continued to surprise me.
In 2023, our first attempt to use an LLM for bug huntingfailed. By the end of 2025, we were already using LLMs with agents to analyze vulnerabilities, generate PoCs, and validate them on real devices. After entering 2026, some vulnerability research had even become close to highly automated. I then applied the same methods to cryptography and other cybersecurity projects, where the performance of LLMs was equally astonishing.
## First Attempt: Bug hunting with Code Llama
Our idea at the time was very straightforward: if an LLM could understand code, then after seeing enough vulnerable code, it should be able to determine whether a piece of code contained a vulnerability and identify the specific bug.
We collected approximately 40,000 Python vulnerability samples, 37,000 Java vulnerability samples, and 40,000 clean-code samples to fine-tune Code Llama. We hoped that, given a piece of code as input, it would output either “vulnerable,” together with the specific vulnerability, or “not vulnerable.”
It performed very well on the benchmark. Among 24 test classes, eight achieved a score of 100%, while 23 scored above 90%.
Then we tested it on real-world code. The result was almost useless.
The model produced many findings that looked like vulnerabilities but were not actually exploitable. The output was noisy, non-actionable, and impossible to validate. The beautiful benchmark numbers did not translate into real vulnerabilities.
Looking back, the reasons for the failure are now quite clear:
1. **The target was wrong.** Simple LLM-based recognition of vulnerability patterns is not the same as the process of vulnerability research.
2. **The capability was wrong.** More SFT data does not mean deeper reasoning.
3. **It overfit CVE patterns.** Common code features in CVEs are not the same as the security invariants that a system actually needs to preserve.
4. **There was not enough context.** A code snippet is not the complete context of the target system.
5. **There was no threat model.** Without security boundaries, adversary capabilities, entry points, and security goals, it is difficult to determine whether a behavior is truly a vulnerability.
6. **It could not validate its findings.** The model could not run, debug, and verify its own conclusions.
7. **There was no continuous research process.** A single question and answer is not research. Research requires a cycle of hypothesis, evidence, and revision.
Simply put, we treated vulnerability discovery as a classification problem.
Our idea was still too simple, and the LLM ecosystem was not mature enough. At the time, it could not yet be used effectively for vulnerability research.
## What Had Changed by the End of 2025?
Two years later, the situation was completely different.
First, the models themselves had become much stronger. Models such as GPT-5.2 and Opus 4.5 had much better reasoning capabilities and no longer merely recognized superficial patterns in code.
Second, context windows had become longer. More complete code repository, call chains, object lifetimes, and trust boundaries can be put into the context. The model no longer saw only an isolated function; it saw a whole system.
Agent workflows had also matured. Approaches such as ReAct, Plan-and-Execute, Chain of Thought, and Multi-Agent systems allowed models to break complex tasks into many steps and adjust their plans based on intermediate results.
Finally, more tools had become available. MCP, CLI, A2A, skills, and various deterministic analysis and execution scripts allowed models to operate on real targets instead of merely returning an answer in a chat box.
The real breakthrough, therefore, was not only a better model. It was a more complete system built around the model.
## Second Attempt: LLM + Agent
At the end of 2025, we tried again to use an LLM for bug hunting. This time, the target was access-control vulnerabilities in customized Android systems.
This was an ideal scenario for testing an agent.
The Android ecosystem today includes not only phones, but also watches, bands, TVs, in-vehicle systems, and many kinds of IoT devices. Different vendors add large numbers of OEM services to their ROMs. These services usually have elevated privileges while exposing interfaces to ordinary apps.
In theory, those interfaces should protect sensitive data and privileged operations through mechanisms such as permissions, UIDs, signatures, and AppOps. In complex customized code, however, gaps frequently appear in permission validation. A low-privilege third-party app may consequently gain access to device identifiers, personal data, or system settings.
This research was conducted between December 2025 and January 2026 by X0ev, Wh1tc from Kunlun Lab and me.
### Preparing Tools for the Agent
We did not simply ask the model to “read the entire firmware and find vulnerabilities.” Instead, we first prepared a collection of deterministic tools:
- **Firmware Unpacking:** Extract APKs from the system, vendor, and product partitions.
- **JADX Decompilation:** Convert system APKs into readable Java code.
- **Manifest Scan:** Organize permissions, actions, and components.
- **Indexing:** Index file paths, packages, and classes.
- **Task Generation:** Select exported, high-risk system components.
- **ADB and Logcat:** Invoke real interfaces and collect access-control evidence.
The tools themselves are deterministic. The skill determines when and in what order the agent should use them.
For high-risk APIs, we asked the agent to complete two reasoning steps.
The first step was to understand the behavior: which identities, parameters, and system states could trigger a sensitive operation?
The second step was to determine authorization: did this operation exceed the privileges that the caller should have? If so, it became a candidate vulnerability.
These two questions may look simple, but they distinguish between something that merely “looks like a vulnerability” and something that genuinely violates a security boundary.
### Validation on Real Devices
A candidate vulnerability was not enough. The agent needed to generate a PoC, run it on a real device, collect Logcat output, and then revise and retry the PoC based on the failure results.
Common causes of failure included an incorrect call order, a missing required state, unmet data dependencies, and incorrect permissions or runtime environments. The agent’s job was not to produce the correct answer in a single attempt, but to continue revising its work after a failure.
Using Codex GPT-5.2 Extra High, we eventually tested real devices from OPPO, HONOR, Xiaomi, Samsung, Google, and other vendors. The results were:
- 110 confirmed vulnerabilities;
- an average cost of less than one US dollar per PoC;
- approximately two iterations per PoC on average.
This time, the LLM was no longer merely a chatbot.
In 2023, we asked it, “Does this code look like a known vulnerability pattern?”
By the end of 2025, we were asking an agent to understand a system, investigate a bug candidate, and validate it with a PoC. The AI used tools, performed scoped analysis, and generated reports. Researchers selected attack surfaces, built the harnesses and skills, assessed the value of the results, and took responsibility for the final conclusions.
This is vulnerability research conducted jointly by researchers and AI.
## From Patches to Vulnerability Variants: Highly Automated Bug Hunting
In the previous research, a human still needed to select the attack surface and build a harness and skill around a particular security invariant.
With a stronger model and agent, could AI complete this part as well?
If we had enough data, especially patches, public write-ups, and PoCs, the answer was beginning to become “yes.”
We built Diffract. This is a joint work with Yunpeng Tian, @___2st.
The core Diffract workflow is:
1. Collect public vulnerability write-ups and PoCs.
2. Build a foundational knowledge base and skill.
4. Extract vulnerability root causes.
5. Abstract and preserve vulnerability patterns.
6. Search for variants in other code.
8. Feed new results back into the knowledge base and skill.
This is no longer a one-off question and answer. It is a research loop capable of accumulating knowledge.
Using Claude Opus 4.6 through 4.8 on Microsoft Windows Kernel and Services, we obtained:
- more than 200 confirmed vulnerabilities;
- RCE, LPE, and information-leak vulnerabilities across different modules.
In fact, this has already become a form of token economics. In the short term, the more tokens you can burn, the more vulnerabilities you can find, because the backlog of legacy vulnerabilities is simply enormous. Windows vulnerability research has become quite boring in 2026. I hope Microsoft can quickly clear this huge backlog of vulnerabilities. Although they actually doing a great job, this month almost 1,000 vulnerabilities fixed.
### The State of Vulnerability Research in 2026
By 2026, my assessment of LLM-based bug hunting was:
- With previous-generation models, bug hunting works as long as the harness is sufficiently good. For example, we found many vulnerabilities using GPT-5.2 and GLM-5.2.
- With a sufficiently capable model, nearly fully automated bug hunting has become possible. Opus 4.6, Kimi K3, DeepSeek V4, GPT-5.4, and GLM 5.3 are already good enough.
- Today, with a SOTA model, all you need to do is make a wish to the model, and it can find vulnerabilities. There is no longer any barrier to entry.
- These three changes happened within six months. Things are changing extremely quickly.
- Bug hunting has evolved into a form of token economics.
In the past, the primary cost of bug hunting was researcher's time. Model capability, tokens, tools, and computing resources are now becoming part of a new cost structure.
This article does not discuss vulnerability exploitation, but the situation is basically the same.
### How to Use AI for Bug Hunting
How to use AI depends on what model you have. The main difference is the level at which the researcher needs to provide help:
- If the model is less capable, the researcher needs to break the task into smaller steps, prepare the context, tools, and tests, and frequently check and correct its work. Much of the work is actually using engineering methods to make up for the model's limitations and ensure that the whole process runs correctly.
- If the model is good enough, it can usually run the full process independently. The researcher mainly monitors the research process, provides background information, system access, the runtime environment, or debugging experience at key moments, and judges the final results.
- If you use a SOTA model, routine analysis and validation often need little help. The researcher should focus on where the model failed to find vulnerabilities: check what it overlooked, add runtime states or cross-module relationships that it could not see, change the search and validation methods, and let it continue trying where it previously failed.
Therefore, the stronger the model, the less the researcher needs to help it complete the process, and the more the researcher needs to help it push beyond the process's current capability boundary.
### About Vulnerability Exploitation
In fact, vulnerability exploitation is also being dramatically accelerated by AI, but that is not the foc

[truncated]
