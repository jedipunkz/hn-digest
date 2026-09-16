---
source: "https://www.qawolf.com/blog/cost-of-ai-slop-in-software-development-report"
hn_url: "https://news.ycombinator.com/item?id=49723767"
title: "Cost of AI Slop in Software Development Report"
article_title: "The Cost of AI Slop in Software Development Report | QA Wolf"
image: "https://cdn.prod.website-files.com/6260298eca091b3621cf1890/6aa0887ea15a9d0ea014affb_og_blog_post%20(4).png"
author: "TheAnkurTyagi"
captured_at: "2026-09-16T09:42:43Z"
capture_tool: "hn-digest"
hn_id: 49723767
score: 2
comments: 0
posted_at: "2026-09-16T08:56:23Z"
tags:
  - hacker-news
---

# Cost of AI Slop in Software Development Report

- HN: [49723767](https://news.ycombinator.com/item?id=49723767)
- Source: [www.qawolf.com](https://www.qawolf.com/blog/cost-of-ai-slop-in-software-development-report)
- Score: 2
- Comments: 0
- Posted: 2026-09-16T08:56:23Z

## Translation

Title: Cost of AI Slop in Software Development Report
Article title: The Cost of AI Slop in Software Development Report | QA Wolf
Description: AI slop costs software teams $1.17 trillion a year. We added up the price of unverified AI code across 6 categories — from wasted tokens to churn.

Article text:
The Cost of AI Slop in Software Development Report | QA Wolf
Platform Mapping AI Autonomously outline your entire app in minutes Automation AI Turn prompts into deterministic Playwright & Appium Run Infra Run web and mobile tests in parallel Service Resources Docs Community Changelog Blog Pricing Terms of Service Privacy Policy Careers Contact QA Wolf is a hybrid platform & service that helps software teams ship better software faster by taking QA completely off their plate.
The Cost of AI Slop in Software Development Report
Developers have long suspected AI slop was expensive. But nobody had put a number to it. We did — and that number might surprise you: they start around $1.17 TRILLION.
And that's only a partial accounting of the total costs.
This piece adds up the true cost of AI-generated code that skipped proper verification, accumulated in production, and is now quietly compounding.
Wondering how we came to it? Read on.
Turning AI slop into a dollar figure
We thought this project would take an afternoon. Instead, it took weeks. Not because the data about the impact of AI slop doesn't exist but because too much of it exists but it’s all fragmented and partial.
Over the last couple of years, study after study has measured various small parts of the problem. There's research on how much AI-generated code gets thrown away almost as soon as it's written. There's also research on how often that code ships with security holes in it. And there's giant macro estimates for what bad software costs the economy as a whole, irrespective of whether it’s written by AI.
But none of these studies came with dollar figures attached. What follows is our attempt to quantify the costs in any way that we could. You might disagree with the way we calculated one figure or another but we’ve done what we could to be as rigorous as possible. We also showed our work so you can plug in your own values or estimates.
In the end, we decided to look at the cost of AI slop across 6 different categories:
Verification of AI-generated code
Change failure rates/fixes in production
Lost revenue from customer churn
EntelligenceAI recently released a study that looked at their customers’ engineering data and found that for every $1 of AI tokens:
$0.44 was going to fixing bugs in generated content.
$0.27 was spent on rework due to bad generated code.
That gave us a total of $0.71 per dollar of AI tokens or 71% of all direct AI spend that can be directly counted towards the costs of AI slop.
How much a company is actually spending on AI slop per engineer, however, can vary widely depending on what tool they’re using. For AI coding tools like GitHub Copilot that are sold via flat rate subscriptions, the cost of AI slop per engineer on their Pro+ plan that costs $39 per month would be 71% of that or $27.69 per engineer per month and $332.28 per engineer per year.
On the other end of the spectrum are the AI slop costs of consumption based code generation tools. For example, take Anthropic’s average reported Claude Code costs per developer of $6 to $12 a day or $180 to $360 per month. Earlier this year, they listed that amount in their docs but have since removed it, likely because the average is now higher. But, as that’s the most recent data we have, we decided to use that for our calculations.
At that rate, AI slop makes up $127.80 to $255.60 per month in AI costs or $1,533.60 to $3,067.20 per year per developer. With many companies reporting power users racking up as much as $500 to $2,000+ per month, that would cost $355 to $1,420 per month or $4,260 to $17,040 per year.
But just calculating the per engineer cost doesn’t provide a sense of what the total costs of AI slop in generation could be. According to Mordor Intelligence , global spending on AI code generation tools is expected to reach $9.35 billion annually by the end of 2026. Following the above ratio, you could say that 71% of that or $6.64 billion per year is wasted on AI slop.
Additional time it takes to validate and ship AI code
Generating code has never been faster. Verifying it has never been slower.
Reviewing AI-generated code takes more effort than reviewing human written code. Developers have to trace unfamiliar logic, hunt subtle bugs, and untangle large blocks of output nobody remembers writing.
In Harness's 2026 State of Engineering Excellence Report , most developers said that since their teams adopted AI they now spend 30% more time reviewing code — and 28% reported their review time climbing by more than 30%.
Engineers were already spending up to five hours a week reviewing code on average. A 30% increase on five hours adds about an hour and a half a week, pushing review to roughly six and a half hours.
Testing is likely the same story, however try as we might, we couldn't find studies quantifying how much extra time is spent testing and debugging AI code, so we can't quantify that for our analysis. We do know that around 45% of developers say debugging AI-generated code takes longer than fixing human-written code, according to Stack Overflow's 2025 Developer Survey. But since no studies exist on how long it previously took on average to debug code or how much extra time devs are spending, there's no way to objectively measure that so that remains a gap in our analysis.
According to Rippling , the average US software engineer earns about $148,000, or roughly $71 an hour. That extra hour and a half of verification a week costs about $107 per engineer a week, or roughly $426 a month. Over a year, that's about $5,538 per engineer. At a company with 50 engineers, that's roughly $277,000 every year.
Developers are spending more time fixing AI code after it's shipped but it’s hard to pinpoint exactly how much extra time they’re spending.
A 2026 study found 60% of software leaders reported quality issues in the past year because code creation outpaced testing capacity.
A study by GitClear found the percentage of new code requiring revision within two weeks grew from 5.5% in 2020 to 7.9% by 2024, highlighting the hidden time engineers spend fixing logically flawed AI implementations.
A 2026 study by Cortex found a 30% increase in change failure rates.
However, none of those studies quantify the number of extra bugs or quality issues those percentages amount to and it’s impossible to definitively pinpoint it from the data provided.
But we also know that the cost of a defect scales with how late it's caught. The most commonly cited data on this is research by the Systems Sciences Institute at IBM, where in a 1995 study, they found that the cost to fix an error in production was 100x more than fixing one caught during the design stage.
But other research has found much lower increases in time and cost. We decided it was better to be conservative. So, for that reason, we used a 2002 study by the National Institute of Standard Technology (NIST) for our calculations. It measured the cost of fixing bugs in production and found that it took 15 hours compared to five hours if it was found at the coding stage. That represents three times extra effort or 10 hours more per bug.
Conservatively, we also estimated a company might see three to five extra bugs per month that are found in production due to AI slop, representing 30 to 50 hours of extra work per month spent fixing those bugs. Some companies might see more and some might see less.
Using the average salary data for a software engineer from Rippling that we shared above, that would translate into $2,312 a month and $27,750 per year at the lower figure of 30 hours per month and $3,854 a month or $46,250 per year for the higher figure of 50 extra hours.
Over the last several years, there’s been a rise of AI code related incidents and that upwards trajectory is unlikely to change soon. In March 2026, Amazon held a mandatory all-hands after internal documents surfaced describing a "trend of incidents" with a "high blast radius" linked to "Gen-AI assisted changes." Amazon is far from alone.
According to the ThousandEyes blog , global outages climbed from 1,382 in January 2025 to 1,595 in February, then spiked to 2,110 in March.
In a 2026 report , 100% of the technology leaders surveyed claimed their company had experienced AI-related downtime
So, what’s the cost of an incident? It varies depending on the size of the company and the nature of an incident (i.e. complete downtime vs a slowed application) but research shows it costs an average of $15,000 per minute or $900,000 per hour, according to Splunk and Cisco's Hidden Costs of Downtime 2026 report.
It’s hard, however, to determine how much extra downtime companies are experiencing directly due to AI. So, we decided to be conservative and predict just one to five AI-related incidents a year at an incident length of 175 minutes , the average found in research by PagerDuty, at a cost of up to $15,000 per minute.
If teams had just one extra incident per year, that could cost them up to $2,635,000 depending on the size of their company and the severity of the incident. If they had five extra incidents, that could cost the company up to $13,125,000 depending on the same factors.
Larger companies or companies that had more incidents due to AI code, could see costs significantly higher than that.
AI slop debt is the new tech debt.
And technical debt is expensive. CISQ’s 2022 report pegged the cost of accumulated software debt at $1.52 trillion for U.S. codebases. That’s because technical debt forces teams to spend a large share of their time on maintenance and rework instead of new product features.
Developers are already feeling the pain with a 2026 study by Sonar Source finding that 40% of developers believe AI has increased debt by generating unnecessary or duplicative code.
Here’s what we know about what’s driving AI slop debt:
Recent academic research comparing human-written code to AI-generated code found that large language models produce code with significantly lower lexical diversity but with much higher structural repetition. This repetitive, pattern-based uniformity severely increases long-term maintainability issues.
Gitclear found that refactored code fell from about 21% of changed lines in 2022 to 3.8% in 2026. What’s more, the number of codeblocks with five or more duplicated lines increased 8x. That’s because assistants make it easy to tab in a new block and are unlikely to suggest reusing an existing function, partly because of limited context size. But that has huge implications on maintenance.
Security debt is also set to become a big issue. Veracode tested 100+ LLMs across 80 tasks and found AI introduced security vulnerabilities in 45% of cases, choosing the insecure option nearly half the time when given the choice.
The problem with all those stats is that they only tell us there’s a problem but don't quantify how big of a problem it is. One academic study that is often cited as finding increases in tech debt due to AI coding agents found more specifically that AI coding agents increased static analysis warnings by 30% and code complexity by 41%. While both arguably are potential technical debt, it doesn’t represent all of the types of technical debt that could be created so we don’t know what percentage of the total technical debt those figures represent.
In order to get to some kind of figure, we decided to take those figures as representative of technical debt as a whole and multiply them by the 2022 level of technical debt from the CISQ report, that would mean AI increased the costs of U.S. technical debt by over $456 billion and then again by $623 billion, adding up to a total of $1.07 trillion if one were to measure them cumulatively.
However, that’s a very big number and would make up the majority of our costs so we decided to be more conservative and assume technical debt increased by between $456 billion and $623 billion.
Customer churn due to poor quality applications
It wouldn’t be a surprise if companies were facing custom

[truncated]

## Original Extract

AI slop costs software teams $1.17 trillion a year. We added up the price of unverified AI code across 6 categories — from wasted tokens to churn.

The Cost of AI Slop in Software Development Report | QA Wolf
Platform Mapping AI Autonomously outline your entire app in minutes Automation AI Turn prompts into deterministic Playwright & Appium Run Infra Run web and mobile tests in parallel Service Resources Docs Community Changelog Blog Pricing Terms of Service Privacy Policy Careers Contact QA Wolf is a hybrid platform & service that helps software teams ship better software faster by taking QA completely off their plate.
The Cost of AI Slop in Software Development Report
Developers have long suspected AI slop was expensive. But nobody had put a number to it. We did — and that number might surprise you: they start around $1.17 TRILLION.
And that's only a partial accounting of the total costs.
This piece adds up the true cost of AI-generated code that skipped proper verification, accumulated in production, and is now quietly compounding.
Wondering how we came to it? Read on.
Turning AI slop into a dollar figure
We thought this project would take an afternoon. Instead, it took weeks. Not because the data about the impact of AI slop doesn't exist but because too much of it exists but it’s all fragmented and partial.
Over the last couple of years, study after study has measured various small parts of the problem. There's research on how much AI-generated code gets thrown away almost as soon as it's written. There's also research on how often that code ships with security holes in it. And there's giant macro estimates for what bad software costs the economy as a whole, irrespective of whether it’s written by AI.
But none of these studies came with dollar figures attached. What follows is our attempt to quantify the costs in any way that we could. You might disagree with the way we calculated one figure or another but we’ve done what we could to be as rigorous as possible. We also showed our work so you can plug in your own values or estimates.
In the end, we decided to look at the cost of AI slop across 6 different categories:
Verification of AI-generated code
Change failure rates/fixes in production
Lost revenue from customer churn
EntelligenceAI recently released a study that looked at their customers’ engineering data and found that for every $1 of AI tokens:
$0.44 was going to fixing bugs in generated content.
$0.27 was spent on rework due to bad generated code.
That gave us a total of $0.71 per dollar of AI tokens or 71% of all direct AI spend that can be directly counted towards the costs of AI slop.
How much a company is actually spending on AI slop per engineer, however, can vary widely depending on what tool they’re using. For AI coding tools like GitHub Copilot that are sold via flat rate subscriptions, the cost of AI slop per engineer on their Pro+ plan that costs $39 per month would be 71% of that or $27.69 per engineer per month and $332.28 per engineer per year.
On the other end of the spectrum are the AI slop costs of consumption based code generation tools. For example, take Anthropic’s average reported Claude Code costs per developer of $6 to $12 a day or $180 to $360 per month. Earlier this year, they listed that amount in their docs but have since removed it, likely because the average is now higher. But, as that’s the most recent data we have, we decided to use that for our calculations.
At that rate, AI slop makes up $127.80 to $255.60 per month in AI costs or $1,533.60 to $3,067.20 per year per developer. With many companies reporting power users racking up as much as $500 to $2,000+ per month, that would cost $355 to $1,420 per month or $4,260 to $17,040 per year.
But just calculating the per engineer cost doesn’t provide a sense of what the total costs of AI slop in generation could be. According to Mordor Intelligence , global spending on AI code generation tools is expected to reach $9.35 billion annually by the end of 2026. Following the above ratio, you could say that 71% of that or $6.64 billion per year is wasted on AI slop.
Additional time it takes to validate and ship AI code
Generating code has never been faster. Verifying it has never been slower.
Reviewing AI-generated code takes more effort than reviewing human written code. Developers have to trace unfamiliar logic, hunt subtle bugs, and untangle large blocks of output nobody remembers writing.
In Harness's 2026 State of Engineering Excellence Report , most developers said that since their teams adopted AI they now spend 30% more time reviewing code — and 28% reported their review time climbing by more than 30%.
Engineers were already spending up to five hours a week reviewing code on average. A 30% increase on five hours adds about an hour and a half a week, pushing review to roughly six and a half hours.
Testing is likely the same story, however try as we might, we couldn't find studies quantifying how much extra time is spent testing and debugging AI code, so we can't quantify that for our analysis. We do know that around 45% of developers say debugging AI-generated code takes longer than fixing human-written code, according to Stack Overflow's 2025 Developer Survey. But since no studies exist on how long it previously took on average to debug code or how much extra time devs are spending, there's no way to objectively measure that so that remains a gap in our analysis.
According to Rippling , the average US software engineer earns about $148,000, or roughly $71 an hour. That extra hour and a half of verification a week costs about $107 per engineer a week, or roughly $426 a month. Over a year, that's about $5,538 per engineer. At a company with 50 engineers, that's roughly $277,000 every year.
Developers are spending more time fixing AI code after it's shipped but it’s hard to pinpoint exactly how much extra time they’re spending.
A 2026 study found 60% of software leaders reported quality issues in the past year because code creation outpaced testing capacity.
A study by GitClear found the percentage of new code requiring revision within two weeks grew from 5.5% in 2020 to 7.9% by 2024, highlighting the hidden time engineers spend fixing logically flawed AI implementations.
A 2026 study by Cortex found a 30% increase in change failure rates.
However, none of those studies quantify the number of extra bugs or quality issues those percentages amount to and it’s impossible to definitively pinpoint it from the data provided.
But we also know that the cost of a defect scales with how late it's caught. The most commonly cited data on this is research by the Systems Sciences Institute at IBM, where in a 1995 study, they found that the cost to fix an error in production was 100x more than fixing one caught during the design stage.
But other research has found much lower increases in time and cost. We decided it was better to be conservative. So, for that reason, we used a 2002 study by the National Institute of Standard Technology (NIST) for our calculations. It measured the cost of fixing bugs in production and found that it took 15 hours compared to five hours if it was found at the coding stage. That represents three times extra effort or 10 hours more per bug.
Conservatively, we also estimated a company might see three to five extra bugs per month that are found in production due to AI slop, representing 30 to 50 hours of extra work per month spent fixing those bugs. Some companies might see more and some might see less.
Using the average salary data for a software engineer from Rippling that we shared above, that would translate into $2,312 a month and $27,750 per year at the lower figure of 30 hours per month and $3,854 a month or $46,250 per year for the higher figure of 50 extra hours.
Over the last several years, there’s been a rise of AI code related incidents and that upwards trajectory is unlikely to change soon. In March 2026, Amazon held a mandatory all-hands after internal documents surfaced describing a "trend of incidents" with a "high blast radius" linked to "Gen-AI assisted changes." Amazon is far from alone.
According to the ThousandEyes blog , global outages climbed from 1,382 in January 2025 to 1,595 in February, then spiked to 2,110 in March.
In a 2026 report , 100% of the technology leaders surveyed claimed their company had experienced AI-related downtime
So, what’s the cost of an incident? It varies depending on the size of the company and the nature of an incident (i.e. complete downtime vs a slowed application) but research shows it costs an average of $15,000 per minute or $900,000 per hour, according to Splunk and Cisco's Hidden Costs of Downtime 2026 report.
It’s hard, however, to determine how much extra downtime companies are experiencing directly due to AI. So, we decided to be conservative and predict just one to five AI-related incidents a year at an incident length of 175 minutes , the average found in research by PagerDuty, at a cost of up to $15,000 per minute.
If teams had just one extra incident per year, that could cost them up to $2,635,000 depending on the size of their company and the severity of the incident. If they had five extra incidents, that could cost the company up to $13,125,000 depending on the same factors.
Larger companies or companies that had more incidents due to AI code, could see costs significantly higher than that.
AI slop debt is the new tech debt.
And technical debt is expensive. CISQ’s 2022 report pegged the cost of accumulated software debt at $1.52 trillion for U.S. codebases. That’s because technical debt forces teams to spend a large share of their time on maintenance and rework instead of new product features.
Developers are already feeling the pain with a 2026 study by Sonar Source finding that 40% of developers believe AI has increased debt by generating unnecessary or duplicative code.
Here’s what we know about what’s driving AI slop debt:
Recent academic research comparing human-written code to AI-generated code found that large language models produce code with significantly lower lexical diversity but with much higher structural repetition. This repetitive, pattern-based uniformity severely increases long-term maintainability issues.
Gitclear found that refactored code fell from about 21% of changed lines in 2022 to 3.8% in 2026. What’s more, the number of codeblocks with five or more duplicated lines increased 8x. That’s because assistants make it easy to tab in a new block and are unlikely to suggest reusing an existing function, partly because of limited context size. But that has huge implications on maintenance.
Security debt is also set to become a big issue. Veracode tested 100+ LLMs across 80 tasks and found AI introduced security vulnerabilities in 45% of cases, choosing the insecure option nearly half the time when given the choice.
The problem with all those stats is that they only tell us there’s a problem but don't quantify how big of a problem it is. One academic study that is often cited as finding increases in tech debt due to AI coding agents found more specifically that AI coding agents increased static analysis warnings by 30% and code complexity by 41%. While both arguably are potential technical debt, it doesn’t represent all of the types of technical debt that could be created so we don’t know what percentage of the total technical debt those figures represent.
In order to get to some kind of figure, we decided to take those figures as representative of technical debt as a whole and multiply them by the 2022 level of technical debt from the CISQ report, that would mean AI increased the costs of U.S. technical debt by over $456 billion and then again by $623 billion, adding up to a total of $1.07 trillion if one were to measure them cumulatively.
However, that’s a very big number and would make up the majority of our costs so we decided to be more conservative and assume technical debt increased by between $456 billion and $623 billion.
Customer churn due to poor quality applications
It wouldn’t be a surprise if companies were facing custom

[truncated]
