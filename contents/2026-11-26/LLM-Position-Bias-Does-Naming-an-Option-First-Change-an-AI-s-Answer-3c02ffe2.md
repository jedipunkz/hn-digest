---
source: "https://techstackups.com/articles/llm-position-bias/"
hn_url: "https://news.ycombinator.com/item?id=50655425"
title: "LLM Position Bias: Does Naming an Option First Change an AI's Answer?"
article_title: "LLM Position Bias: Does Naming an Option First Change an AI's Answer? | Tech Stackups"
image: "https://techstackups.com/img/articles/ai-option-order-bias/cover.jpg"
author: "ritzaco"
captured_at: "2026-11-26T14:45:37Z"
capture_tool: "hn-digest"
hn_id: 50655425
score: 2
comments: 0
posted_at: "2026-11-26T14:03:25Z"
tags:
  - hacker-news
---

# LLM Position Bias: Does Naming an Option First Change an AI's Answer?

- HN: [50655425](https://news.ycombinator.com/item?id=50655425)
- Source: [techstackups.com](https://techstackups.com/articles/llm-position-bias/)
- Score: 2
- Comments: 0
- Posted: 2026-11-26T14:03:25Z

## Translation

Title: LLM Position Bias: Does Naming an Option First Change an AI's Answer?
Article title: LLM Position Bias: Does Naming an Option First Change an AI's Answer? | Tech Stackups
Description: We tested LLM position bias with 2,000 Claude Code runs across 100 pairs. Order rarely mattered, but on toss-ups like Kafka vs RabbitMQ it decided the pick.

Article text:
Skip to main content Tech Stackups Home Topics Comparisons Guides Articles AX Search LLM Position Bias: Does Naming an Option First Change an AI's Answer?
Since the beginning of time, LLMs have always just agreed with their users. In a comparison where an LLM didn't know what to agree with, it would just choose the first-named option, assuming it was the first one that came to the user's mind for a reason. The reason is far less complex than that. Researchers call it position bias, and we wanted to know how much of it today's AI agents still have.
We asked Claude Code to choose between 100 pairs of options, such as PostgreSQL or MySQL, Slack or Microsoft Teams, McDonald's or Burger King, and heads or tails, 10 times with each option named first.
In 72 of the 100 pairs, the agent gave the same answer in all 20 runs, whatever the order. But when it had no real favourite, the order decided the answer. On a coin toss, it called whichever side we named first 19 times out of 20. (Yes, that means you can rig the coin toss when Claude picks who pays for lunch. No, you shouldn't.)
This article covers which pairs the order moved, the statistics behind the result, and how to check whether the order is deciding an answer you get from an agent.
How we tested All 2,000 sessions used Claude Opus 5.5 ( claude-opus-5-5 ) in Claude Code 2.1.326, on 25 and 26 November 2026. Each run was a fresh session in an empty folder, under a clean Claude Code profile with no tools, MCP servers, plugins or memory. Every prompt had the same shape: one sentence of context, then "Should I go with A or B? End your answer with one line: PICK: <name>." The two orders used the same context sentence, so the only difference was which name came first. We ran the 2,000 sessions in a shuffled order so time of day couldn't line up with the order of the names, and we fixed the hypotheses and tests below before the first run.
Does AI favour the first option? ​
The option named first won 1,031 of the 2,000 runs, or 51.6%. If the order made no difference, you'd expect about 50%.
That 1.6-point gap passes a significance test, but only just, and it comes almost entirely from pairs where the agent had no real preference. When it did have a preference, which was most of the time, swapping the order changed nothing.
Each dot below is a pair, placed by how often its first option won in each order. If the order made no difference, every dot would sit on the dashed line.
0 0 2 2 4 4 6 6 8 8 10 10 No order effect Named first helped Named second helped Option A won, when named first (of 10) Option A won, when named second (of 10) 48 24 6 Snowflake / BigQuery Heads / Tails Kafka / RabbitMQ Hover over or tap a dot to see its pairs. Option A is the first name in each pair, for example Kafka in Kafka / RabbitMQ. The 48 pairs top right and 24 bottom left gave the same answer in all 20 runs. A dot below the dashed line means naming option A first helped it, and a dot above means naming it second helped.
Most questions had a clear winner ​
The agent gave the same answer in all 20 runs on 61 of the 80 commercial questions, no matter how we asked. It picked PostgreSQL over MySQL for a new SaaS app, Slack over Microsoft Teams for a 50-person company, Toyota over Honda for a reliable family car and Coca-Cola over Pepsi for a party, in all 20 runs each time. Swapping the names round didn't move any of them.
The pairs that moved were the ones where the agent's view was already split, which is why non-commercial questions with no clear answer, like tea or coffee, moved most often.
Where position bias shows up: when the AI has no favourite ​
The pairs below show what happens when the agent has no settled view of its own.
A coin toss is the purest test, because there's no right answer to know. All 20 answers said the call was 50/50, yet the agent picked whichever side we'd named first in 19 of them.
Its reasons changed with the order. Asked "Heads or Tails?", it picked heads as "the classic", and asked "Tails or Heads?", it picked tails for a different reason:
It's a 50/50 either way … People call heads more often, so tails is a fine contrarian choice.
Both reasons sound sensible, but the order of the names decided the pick, and the agent found a reason for whichever side came first.
This was the one commercial pair where the effect was strong enough to count on its own. We said we needed "a message broker for event processing between services", which is a real toss-up: Kafka suits replayable event streams, RabbitMQ suits simpler message passing, and our sentence didn't say which we needed.
Kafka won 10 of 10 when named first and 2 of 10 when named second. The answers listed the same pros and cons either way and differed only in the recommendation at the end.
The answers that picked RabbitMQ admitted it was a close call: 7 of the 8 offered to reconsider if we gave more detail. One ended: "If you share your expected message rate, number of consuming services, and whether you need replay, I can confirm or flip this." The Kafka answers were more confident, with only 5 of 12 making a similar offer. One opened with "For event processing between services, I'd go with Kafka."
Snowflake vs BigQuery: the second name won ​
The order doesn't always favour the first name. When we asked about a data warehouse for an analytics team, the agent picked whichever option we named second in 16 of 20 runs:
After correcting for testing 100 pairs, this one isn't significant on its own, so treat it as a hint. It does match research on older models: a 2025 study found that models favour the first option when all options are strong, and later options otherwise , so the direction of the effect isn't fixed.
The other pairs that moved shifted by a few runs, for example:
Asana or Monday.com for a marketing team: Asana won 9 of 10 when named first and 5 of 10 when named second.
China or India for a two-week trip: China won 5 of 10 when named first and 2 of 10 when named second.
Canon or Nikon for a first mirrorless camera: Canon won 7 of 10 when named first and 5 of 10 when named second.
None of these is large enough to rule out chance on its own, but they lean the same way: of the 22 pairs that moved, 19 moved towards the first-named option.
The statistics: null hypothesis and t-test results ​
We fixed the hypotheses and tests before running anything:
Null hypothesis (H₀): the order of the names makes no difference. Across pairs, the first-named option is picked 50% of the time.
Alternative (H₁): the first-named option is picked more or less than 50% of the time. We tested both directions, because earlier research found models sometimes favour the last option.
Measure: for each pair, the share of its 20 runs that picked whichever option was named first. A pair with no order effect scores 50%, whatever its winner.
Primary test: a two-sided one-sample t-test of those 100 scores against 50%, at α = 0.05.
So we reject the null hypothesis at the 5% level: the order of the names has an effect. But the effect is small (Cohen's d = 0.20), and the p-value of 0.046 is only just under the 0.05 threshold.
Two checks show how much of that p = 0.046 depends on the coin-toss-style questions:
Leaving out the two pure-chance controls (heads or tails, and odd or even) raises the t-test to p = 0.11, which is no longer significant. The Cochran-Mantel-Haenszel test still gives p = 0.03. We only decided to run this check after seeing the data.
Testing the 80 commercial pairs on their own , a comparison we planned from the start, gives no significant result on either test: p = 0.36 on the t-test and p = 0.18 on the Cochran-Mantel-Haenszel test.
The close calls showed a bigger effect (58%) than the pairs with a clear winner (50.4%), as we expected. But there were only 15 close calls, too few for the difference to be significant on its own (Welch's t-test, p = 0.13). To show the effect within close calls, you'd need to find more close calls and run them many more times.
How to avoid position bias when an AI picks for you ​
Position bias is a narrow kind of AI bias, and in our study it had very little effect on the agent overall, less than earlier research would suggest. A 2025 study of LLM recommenders, which ranked longer candidate lists, found that items earlier in the list were more likely to be recommended . Opus 5.5 had a settled view on most of the products we asked about, and that view decided the answer. We've looked at what those default picks are in Claude Code's choice of vendors .
The order matters when the agent doesn't have a settled view. Then your phrasing decides the answer, and the agent still gives a reason for its pick. Kafka or RabbitMQ, Snowflake or BigQuery, Asana or Monday.com: these are the kinds of questions people bring to an agent because they're hard to call.
The quickest check is to ask "A or B?" and "B or A?" a few times each, in fresh sessions. When the answers follow whichever name came first, you've found a toss-up, and the order of the names is deciding it for you.
Each row is 10 runs per order. "Moved" marks the pairs where the split changed when we swapped the names.
Hash Kader is a software engineer and technical writer at Ritza . He contributes hands-on testing of AI tools and developer products to TechStackups, with a focus on what they're actually like to use.
Does AI favour the first option?
Most questions had a clear winner
Where position bias shows up: when the AI has no favourite Heads or tails
Snowflake vs BigQuery: the second name won
The statistics: null hypothesis and t-test results
How to avoid position bias when an AI picks for you

## Original Extract

We tested LLM position bias with 2,000 Claude Code runs across 100 pairs. Order rarely mattered, but on toss-ups like Kafka vs RabbitMQ it decided the pick.

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles AX Search LLM Position Bias: Does Naming an Option First Change an AI's Answer?
Since the beginning of time, LLMs have always just agreed with their users. In a comparison where an LLM didn't know what to agree with, it would just choose the first-named option, assuming it was the first one that came to the user's mind for a reason. The reason is far less complex than that. Researchers call it position bias, and we wanted to know how much of it today's AI agents still have.
We asked Claude Code to choose between 100 pairs of options, such as PostgreSQL or MySQL, Slack or Microsoft Teams, McDonald's or Burger King, and heads or tails, 10 times with each option named first.
In 72 of the 100 pairs, the agent gave the same answer in all 20 runs, whatever the order. But when it had no real favourite, the order decided the answer. On a coin toss, it called whichever side we named first 19 times out of 20. (Yes, that means you can rig the coin toss when Claude picks who pays for lunch. No, you shouldn't.)
This article covers which pairs the order moved, the statistics behind the result, and how to check whether the order is deciding an answer you get from an agent.
How we tested All 2,000 sessions used Claude Opus 5.5 ( claude-opus-5-5 ) in Claude Code 2.1.326, on 25 and 26 November 2026. Each run was a fresh session in an empty folder, under a clean Claude Code profile with no tools, MCP servers, plugins or memory. Every prompt had the same shape: one sentence of context, then "Should I go with A or B? End your answer with one line: PICK: <name>." The two orders used the same context sentence, so the only difference was which name came first. We ran the 2,000 sessions in a shuffled order so time of day couldn't line up with the order of the names, and we fixed the hypotheses and tests below before the first run.
Does AI favour the first option? ​
The option named first won 1,031 of the 2,000 runs, or 51.6%. If the order made no difference, you'd expect about 50%.
That 1.6-point gap passes a significance test, but only just, and it comes almost entirely from pairs where the agent had no real preference. When it did have a preference, which was most of the time, swapping the order changed nothing.
Each dot below is a pair, placed by how often its first option won in each order. If the order made no difference, every dot would sit on the dashed line.
0 0 2 2 4 4 6 6 8 8 10 10 No order effect Named first helped Named second helped Option A won, when named first (of 10) Option A won, when named second (of 10) 48 24 6 Snowflake / BigQuery Heads / Tails Kafka / RabbitMQ Hover over or tap a dot to see its pairs. Option A is the first name in each pair, for example Kafka in Kafka / RabbitMQ. The 48 pairs top right and 24 bottom left gave the same answer in all 20 runs. A dot below the dashed line means naming option A first helped it, and a dot above means naming it second helped.
Most questions had a clear winner ​
The agent gave the same answer in all 20 runs on 61 of the 80 commercial questions, no matter how we asked. It picked PostgreSQL over MySQL for a new SaaS app, Slack over Microsoft Teams for a 50-person company, Toyota over Honda for a reliable family car and Coca-Cola over Pepsi for a party, in all 20 runs each time. Swapping the names round didn't move any of them.
The pairs that moved were the ones where the agent's view was already split, which is why non-commercial questions with no clear answer, like tea or coffee, moved most often.
Where position bias shows up: when the AI has no favourite ​
The pairs below show what happens when the agent has no settled view of its own.
A coin toss is the purest test, because there's no right answer to know. All 20 answers said the call was 50/50, yet the agent picked whichever side we'd named first in 19 of them.
Its reasons changed with the order. Asked "Heads or Tails?", it picked heads as "the classic", and asked "Tails or Heads?", it picked tails for a different reason:
It's a 50/50 either way … People call heads more often, so tails is a fine contrarian choice.
Both reasons sound sensible, but the order of the names decided the pick, and the agent found a reason for whichever side came first.
This was the one commercial pair where the effect was strong enough to count on its own. We said we needed "a message broker for event processing between services", which is a real toss-up: Kafka suits replayable event streams, RabbitMQ suits simpler message passing, and our sentence didn't say which we needed.
Kafka won 10 of 10 when named first and 2 of 10 when named second. The answers listed the same pros and cons either way and differed only in the recommendation at the end.
The answers that picked RabbitMQ admitted it was a close call: 7 of the 8 offered to reconsider if we gave more detail. One ended: "If you share your expected message rate, number of consuming services, and whether you need replay, I can confirm or flip this." The Kafka answers were more confident, with only 5 of 12 making a similar offer. One opened with "For event processing between services, I'd go with Kafka."
Snowflake vs BigQuery: the second name won ​
The order doesn't always favour the first name. When we asked about a data warehouse for an analytics team, the agent picked whichever option we named second in 16 of 20 runs:
After correcting for testing 100 pairs, this one isn't significant on its own, so treat it as a hint. It does match research on older models: a 2025 study found that models favour the first option when all options are strong, and later options otherwise , so the direction of the effect isn't fixed.
The other pairs that moved shifted by a few runs, for example:
Asana or Monday.com for a marketing team: Asana won 9 of 10 when named first and 5 of 10 when named second.
China or India for a two-week trip: China won 5 of 10 when named first and 2 of 10 when named second.
Canon or Nikon for a first mirrorless camera: Canon won 7 of 10 when named first and 5 of 10 when named second.
None of these is large enough to rule out chance on its own, but they lean the same way: of the 22 pairs that moved, 19 moved towards the first-named option.
The statistics: null hypothesis and t-test results ​
We fixed the hypotheses and tests before running anything:
Null hypothesis (H₀): the order of the names makes no difference. Across pairs, the first-named option is picked 50% of the time.
Alternative (H₁): the first-named option is picked more or less than 50% of the time. We tested both directions, because earlier research found models sometimes favour the last option.
Measure: for each pair, the share of its 20 runs that picked whichever option was named first. A pair with no order effect scores 50%, whatever its winner.
Primary test: a two-sided one-sample t-test of those 100 scores against 50%, at α = 0.05.
So we reject the null hypothesis at the 5% level: the order of the names has an effect. But the effect is small (Cohen's d = 0.20), and the p-value of 0.046 is only just under the 0.05 threshold.
Two checks show how much of that p = 0.046 depends on the coin-toss-style questions:
Leaving out the two pure-chance controls (heads or tails, and odd or even) raises the t-test to p = 0.11, which is no longer significant. The Cochran-Mantel-Haenszel test still gives p = 0.03. We only decided to run this check after seeing the data.
Testing the 80 commercial pairs on their own , a comparison we planned from the start, gives no significant result on either test: p = 0.36 on the t-test and p = 0.18 on the Cochran-Mantel-Haenszel test.
The close calls showed a bigger effect (58%) than the pairs with a clear winner (50.4%), as we expected. But there were only 15 close calls, too few for the difference to be significant on its own (Welch's t-test, p = 0.13). To show the effect within close calls, you'd need to find more close calls and run them many more times.
How to avoid position bias when an AI picks for you ​
Position bias is a narrow kind of AI bias, and in our study it had very little effect on the agent overall, less than earlier research would suggest. A 2025 study of LLM recommenders, which ranked longer candidate lists, found that items earlier in the list were more likely to be recommended . Opus 5.5 had a settled view on most of the products we asked about, and that view decided the answer. We've looked at what those default picks are in Claude Code's choice of vendors .
The order matters when the agent doesn't have a settled view. Then your phrasing decides the answer, and the agent still gives a reason for its pick. Kafka or RabbitMQ, Snowflake or BigQuery, Asana or Monday.com: these are the kinds of questions people bring to an agent because they're hard to call.
The quickest check is to ask "A or B?" and "B or A?" a few times each, in fresh sessions. When the answers follow whichever name came first, you've found a toss-up, and the order of the names is deciding it for you.
Each row is 10 runs per order. "Moved" marks the pairs where the split changed when we swapped the names.
Hash Kader is a software engineer and technical writer at Ritza . He contributes hands-on testing of AI tools and developer products to TechStackups, with a focus on what they're actually like to use.
Does AI favour the first option?
Most questions had a clear winner
Where position bias shows up: when the AI has no favourite Heads or tails
Snowflake vs BigQuery: the second name won
The statistics: null hypothesis and t-test results
How to avoid position bias when an AI picks for you
