---
source: "https://steadily.ca/newsletter/edition/predictable-ai-skills-with-scripts"
hn_url: "https://news.ycombinator.com/item?id=50291514"
title: "Make your AI skills more predictable with scripts"
article_title: "Make your AI skills more predictable with scripts | Steadily"
image: "https://assets.steadily.ca/images/sharing/edition/2/V3YLzC7DNIE9Eb6QPZnpfQ74qCsKFHVxIZjBOFXG_M0.png"
author: "dhruvio"
captured_at: "2026-10-29T16:39:38Z"
capture_tool: "hn-digest"
hn_id: 50291514
score: 1
comments: 0
posted_at: "2026-10-29T16:19:41Z"
tags:
  - hacker-news
---

# Make your AI skills more predictable with scripts

- HN: [50291514](https://news.ycombinator.com/item?id=50291514)
- Source: [steadily.ca](https://steadily.ca/newsletter/edition/predictable-ai-skills-with-scripts)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T16:19:41Z

## Translation

Title: Make your AI skills more predictable with scripts
Article title: Make your AI skills more predictable with scripts | Steadily
Description: How can we make AI skills more predictable? This edition explains how scripts fit into skills, and details an example of how a lead-generation skill can be improved with the addition of three scripts. Scripts give us more control over how our agents behave, making them more predictable.

Article text:
Pricing FAQ Contact Us
Sunday Newsletter
Book a Call
SUNDAY NEWSLETTER #002 · PUBLISHED 2026-10-18
Make your AI skills more predictable with scripts
Dhruv Dang
Principal Software Engineer, Steadily
This Sunday's newsletter is about making your AI skills more predictable with scripts. We'll explore what skills are and how scripts fit into them. We'll walk through an example of a lead-generation skill that starts off as a single prompt. Then, we'll improve it by adding three scripts alongside its instructions, giving us more control over how our agent behaves, ultimately making it more predictable.
Scripts are small computer programs that perform specific tasks.
Skills are a quick way to package instructions and scripts to augment an AI agent’s abilities.
Bundling scripts with a skill lets you define exactly how your agent should complete certain tasks with code.
Depending on the use case, a skill that combines instructions and scripts together can yield more predictable outcomes than one without any bundled scripts.
An AI skill gives an agent the resources it needs to complete a specific job. For example, a "website-design" skill gives an agent the ability to design websites.
A skill is implemented as a folder containing a SKILL.md file with a name, description, and instructions. It can also contain scripts, reference documents, and other resources. Agents can load the instructions when they're relevant and run the bundled scripts when needed.
For example, the docx skill by Anthropic bundles scripts and instructions to give agents the ability to edit Microsoft Word documents.
What are the benefits of skills?
Customization. Skills let you shape your agent around the work you actually need done.
Predictability. Because a skill can bundle scripts, and scripts are programs you write and control, you can move repeatable operations out of the agent's judgement and into code. The more of a task you define in scripts, the more reliably the agent performs it.
Example: A Better Lead Generation Skill with Scripts
Let's imagine we run a HealthTech company that serves Canadian pharmacies with 10 to 100 employees, and we use a skill to find potential customers. Let's refactor it.
Here's how our find-leads skill's instructions are defined in SKILL.md :
---
name: find-leads
description: Find and save potential customers that match our target criteria.
---
Find Canadian pharmacies with 10 to 100 employees.
Search the web for potential customers. For each candidate, collect its name, website, country, employee count, available contact details, and the source URLs supporting those details.
Save companies that match our criteria to our CRM.
While this skill explains the job, it doesn't define how to:
Save leads so we can follow up with them later.
We could simply add more instructions, but we’d still be relying on the agent to fill in any gaps with its own judgement. If we make these tasks explicit in code, the agent no longer has to guess how we want them done.
So, let’s expand our skill by giving it three scripts:
find-leads/
├── SKILL.md
└── scripts/
├── web_search.js
├── qualify_lead.js
└── store_lead.js
Each script handles a specific task:
web_search.js : Search the web based on a query provided by the agent, and return results in a consistent format. We can limit which search providers and websites are used in this script. Our agent takes the result of this script and infers structured leads to pass to qualify_lead.js .
qualify_lead.js : Qualify leads based on explicit criteria: must be a pharmacy, must have between 10 and 100 employees and must be headquartered in Canada. The script returns either “qualified,” “disqualified” or “unknown.” The agent then decides whether to clarify “unknown” leads with us, and passes the resulting qualified leads to store_lead.js .
store_lead.js : Save leads to our CRM in a specific collection that we can then review and use for cold outreach, outputting whether leads saved successfully. The script can also skip leads for customers we already have a relationship with by checking the CRM.
As you can see, the introduction of these scripts gives us a lot more control over how the agent acts in its lead generation process by letting us encode business rules in each of them. However, we still need to instruct our agent to use these scripts in SKILL.md . Here’s an updated set of instructions doing just that:
---
name: find-leads
description: Find and save potential customers using the bundled lead-generation scripts.
---
Find Canadian pharmacies with 10 to 100 employees. Use the bundled scripts for searching, qualification, and storage.
Run the commands below from this skill's directory.
1. Search with: `./scripts/web_search.js "<search query>"`
2. Investigate candidates and prepare a JSON file for each lead. Include its name, website, country, employee count, available contact details, and supporting source URLs. Leave unknown fields null.
3. Qualify each candidate with: `./scripts/qualify_lead.js "<path to lead JSON file>"`
4. Only save candidates classified as qualified, using: `./scripts/store_lead.js "<path to lead JSON file>"`
5. Report what was saved, disqualified, or left unknown. Explain any script failures instead of claiming success.
Notice how the updated instructions read much more like a defined process. The agent still does what it's good at: choosing search queries, investigating candidates, and turning messy web pages into structured lead details. The scripts handle the steps where we've already decided how the work should be done. The result is a more predictable lead generation workflow, and one we can improve over time. As we measure the quality of the leads it produces, we can refine the scripts accordingly.
This version is also easier to change than the original. Each part of the process now lives in one place. To add an external lead generation service, we update web_search.js . To start targeting pharmacies in the US, we update qualify_lead.js . If we switch CRMs, we update store_lead.js . The skill is more efficient, too: the agent no longer spends tokens and time working out how to search, qualify, and save leads, because the scripts already do that work.
Note: I chose not to show the implementation of this example’s scripts in this edition as I felt it was out of scope. But if you’re interested in building a system like this for your business and would like to collaborate, just reply to this email and let me know.
To apply this concept in your own AI practice, have a look through your skills, prompt library, or chat history. Find a task you often repeat and ask: which parts should work the same way every time? Here are a few ideas:
Research: Let the agent interpret findings, but use a script to assemble a final research document in a fixed format that collates findings and citations and preserves source material.
Social media: Let the agent draft posts, but use scripts to validate scheduling details, submit approved posts, and retrieve engagement data.
AI coding: Use scripts to create a project, migration, or landing page from a template instead of rebuilding the same scaffold each time.
Agent memory: Let the model propose what to remember, but put memory storage and retrieval rules in code.
You can also build agents by combining skills. For example, an agent focused on social media marketing could research topics, draft campaigns, schedule posts, deploy custom landing pages, and save campaign performance to memory.
Which skills would have the greatest impact on your business? How could you break each one down into a process, and which steps could you implement as scripts?
Skills let you customize what your agent can do. Bundling scripts with a skill gives you more control over how specific tasks are performed. In our lead-generation example, the agent handles work that benefits from judgement, such as researching and interpreting candidates, while three scripts manage searching, qualification, and storage. This made the skill’s behaviour more predictable and, ultimately, the agent more reliable.
One practical lesson in AI product execution you can read in under 10 minutes. Delivered every other Sunday morning.
I accept the Privacy Policy and Terms of Service .
Subscribe Connect with us
Contact Us
LinkedIn
X (Twitter)
About
Terms of Service
Privacy Policy
© 2026 Steadily Consulting Inc.

## Original Extract

How can we make AI skills more predictable? This edition explains how scripts fit into skills, and details an example of how a lead-generation skill can be improved with the addition of three scripts. Scripts give us more control over how our agents behave, making them more predictable.

Pricing FAQ Contact Us
Sunday Newsletter
Book a Call
SUNDAY NEWSLETTER #002 · PUBLISHED 2026-10-18
Make your AI skills more predictable with scripts
Dhruv Dang
Principal Software Engineer, Steadily
This Sunday's newsletter is about making your AI skills more predictable with scripts. We'll explore what skills are and how scripts fit into them. We'll walk through an example of a lead-generation skill that starts off as a single prompt. Then, we'll improve it by adding three scripts alongside its instructions, giving us more control over how our agent behaves, ultimately making it more predictable.
Scripts are small computer programs that perform specific tasks.
Skills are a quick way to package instructions and scripts to augment an AI agent’s abilities.
Bundling scripts with a skill lets you define exactly how your agent should complete certain tasks with code.
Depending on the use case, a skill that combines instructions and scripts together can yield more predictable outcomes than one without any bundled scripts.
An AI skill gives an agent the resources it needs to complete a specific job. For example, a "website-design" skill gives an agent the ability to design websites.
A skill is implemented as a folder containing a SKILL.md file with a name, description, and instructions. It can also contain scripts, reference documents, and other resources. Agents can load the instructions when they're relevant and run the bundled scripts when needed.
For example, the docx skill by Anthropic bundles scripts and instructions to give agents the ability to edit Microsoft Word documents.
What are the benefits of skills?
Customization. Skills let you shape your agent around the work you actually need done.
Predictability. Because a skill can bundle scripts, and scripts are programs you write and control, you can move repeatable operations out of the agent's judgement and into code. The more of a task you define in scripts, the more reliably the agent performs it.
Example: A Better Lead Generation Skill with Scripts
Let's imagine we run a HealthTech company that serves Canadian pharmacies with 10 to 100 employees, and we use a skill to find potential customers. Let's refactor it.
Here's how our find-leads skill's instructions are defined in SKILL.md :
---
name: find-leads
description: Find and save potential customers that match our target criteria.
---
Find Canadian pharmacies with 10 to 100 employees.
Search the web for potential customers. For each candidate, collect its name, website, country, employee count, available contact details, and the source URLs supporting those details.
Save companies that match our criteria to our CRM.
While this skill explains the job, it doesn't define how to:
Save leads so we can follow up with them later.
We could simply add more instructions, but we’d still be relying on the agent to fill in any gaps with its own judgement. If we make these tasks explicit in code, the agent no longer has to guess how we want them done.
So, let’s expand our skill by giving it three scripts:
find-leads/
├── SKILL.md
└── scripts/
├── web_search.js
├── qualify_lead.js
└── store_lead.js
Each script handles a specific task:
web_search.js : Search the web based on a query provided by the agent, and return results in a consistent format. We can limit which search providers and websites are used in this script. Our agent takes the result of this script and infers structured leads to pass to qualify_lead.js .
qualify_lead.js : Qualify leads based on explicit criteria: must be a pharmacy, must have between 10 and 100 employees and must be headquartered in Canada. The script returns either “qualified,” “disqualified” or “unknown.” The agent then decides whether to clarify “unknown” leads with us, and passes the resulting qualified leads to store_lead.js .
store_lead.js : Save leads to our CRM in a specific collection that we can then review and use for cold outreach, outputting whether leads saved successfully. The script can also skip leads for customers we already have a relationship with by checking the CRM.
As you can see, the introduction of these scripts gives us a lot more control over how the agent acts in its lead generation process by letting us encode business rules in each of them. However, we still need to instruct our agent to use these scripts in SKILL.md . Here’s an updated set of instructions doing just that:
---
name: find-leads
description: Find and save potential customers using the bundled lead-generation scripts.
---
Find Canadian pharmacies with 10 to 100 employees. Use the bundled scripts for searching, qualification, and storage.
Run the commands below from this skill's directory.
1. Search with: `./scripts/web_search.js "<search query>"`
2. Investigate candidates and prepare a JSON file for each lead. Include its name, website, country, employee count, available contact details, and supporting source URLs. Leave unknown fields null.
3. Qualify each candidate with: `./scripts/qualify_lead.js "<path to lead JSON file>"`
4. Only save candidates classified as qualified, using: `./scripts/store_lead.js "<path to lead JSON file>"`
5. Report what was saved, disqualified, or left unknown. Explain any script failures instead of claiming success.
Notice how the updated instructions read much more like a defined process. The agent still does what it's good at: choosing search queries, investigating candidates, and turning messy web pages into structured lead details. The scripts handle the steps where we've already decided how the work should be done. The result is a more predictable lead generation workflow, and one we can improve over time. As we measure the quality of the leads it produces, we can refine the scripts accordingly.
This version is also easier to change than the original. Each part of the process now lives in one place. To add an external lead generation service, we update web_search.js . To start targeting pharmacies in the US, we update qualify_lead.js . If we switch CRMs, we update store_lead.js . The skill is more efficient, too: the agent no longer spends tokens and time working out how to search, qualify, and save leads, because the scripts already do that work.
Note: I chose not to show the implementation of this example’s scripts in this edition as I felt it was out of scope. But if you’re interested in building a system like this for your business and would like to collaborate, just reply to this email and let me know.
To apply this concept in your own AI practice, have a look through your skills, prompt library, or chat history. Find a task you often repeat and ask: which parts should work the same way every time? Here are a few ideas:
Research: Let the agent interpret findings, but use a script to assemble a final research document in a fixed format that collates findings and citations and preserves source material.
Social media: Let the agent draft posts, but use scripts to validate scheduling details, submit approved posts, and retrieve engagement data.
AI coding: Use scripts to create a project, migration, or landing page from a template instead of rebuilding the same scaffold each time.
Agent memory: Let the model propose what to remember, but put memory storage and retrieval rules in code.
You can also build agents by combining skills. For example, an agent focused on social media marketing could research topics, draft campaigns, schedule posts, deploy custom landing pages, and save campaign performance to memory.
Which skills would have the greatest impact on your business? How could you break each one down into a process, and which steps could you implement as scripts?
Skills let you customize what your agent can do. Bundling scripts with a skill gives you more control over how specific tasks are performed. In our lead-generation example, the agent handles work that benefits from judgement, such as researching and interpreting candidates, while three scripts manage searching, qualification, and storage. This made the skill’s behaviour more predictable and, ultimately, the agent more reliable.
One practical lesson in AI product execution you can read in under 10 minutes. Delivered every other Sunday morning.
I accept the Privacy Policy and Terms of Service .
Subscribe Connect with us
Contact Us
LinkedIn
X (Twitter)
About
Terms of Service
Privacy Policy
© 2026 Steadily Consulting Inc.
