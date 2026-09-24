---
source: "https://bito.ai/benchmarks/ai-coding-model-cost/"
hn_url: "https://news.ycombinator.com/item?id=49826998"
title: "The cost and quality frontier across 29 AI coding models"
article_title: "What AI coding models really cost, 29 models benchmarked | Bito"
image: "https://bito.ai/wp-content/uploads/2026/09/Social-sharing-image-1.webp"
author: "kirti_soni171"
captured_at: "2026-09-24T07:17:24Z"
capture_tool: "hn-digest"
hn_id: 49826998
score: 1
comments: 0
posted_at: "2026-09-24T06:34:00Z"
tags:
  - hacker-news
---

# The cost and quality frontier across 29 AI coding models

- HN: [49826998](https://news.ycombinator.com/item?id=49826998)
- Source: [bito.ai](https://bito.ai/benchmarks/ai-coding-model-cost/)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T06:34:00Z

## Translation

Title: The cost and quality frontier across 29 AI coding models
Article title: What AI coding models really cost, 29 models benchmarked | Bito
Description: We tested 29 AI coding models on 60 real tasks and priced every run. One model reaches 93% of the top score for 2% of the cost. See the full data.

Article text:
What AI coding models really cost, 29 models benchmarked | Bito
Skip to content
We tested 29 AI models to measure the cost-quality frontier.
See who won
Start free
Products
Context engine
What a Crowded AI Frontier Means for Engineering Teams
How real-world scenarios and a plethora of options dramatically change the cost-quality frontier. Based on 29 AI models, 60 real engineering jobs, and 1,740 graded answers.
We tested 29 AI models on the same 60 real engineering tasks, graded every one of the 1,740 answers against criteria written before any model ran, and priced every run. That combination, quality and cost measured together on the same work, is what this research is about.
Four findings matter for anyone deciding what AI models to buy for their engineering teams.
Most of the market is overpriced for what it delivers
21 of the 29 models are beaten by something cheaper that scores at least as well. claude-opus-5 leads at 54.5 out of 60 and cost $135.51 for the full run. Another 7 models land within 4.5 points of it, at prices down to $2.58. deepseek-v4.1-flash, an open-weight model, scored 50.5 for $2.67. That is 93% of the top score for 2% of the money.
The rate card no longer predicts your bill
claude-fable-5.1 lists at 2x the per-token price of claude-opus-5 and cost 37% less to run the same 60 jobs, because it took fewer steps, made half as many model calls, and wrote a third as much output. gpt-6-astra lists at 2.5x gpt-5.6-sol, came in 18% cheaper, and scored 6.5 points higher. The number to buy on is cost per successful task, and we report it for every model.
The score alone cannot tell you which model to trust
Two behaviors separate the models that hold up in daily use. What a model does when your request is vague, and what it does when it forgets an instruction you gave it. We scored both, model by model, and the results rarely follow price. Of the 29 models, 15 clear both bars, and the ones that miss include several that look sensible on a leaderboard.
Tiering the work buys the top score for a fraction of the money
Sort the work by how much judgment it takes, put a different model on each tier, and the top group’s score survives at a fraction of the bill. Opus on the tier where it is clearly better, with cheap models underneath, scored 53.5 for $18. Two open-weight models on their own scored 52.5 for $2.50. Running claude-opus-5 on everything scored 54.5 for $135.51.
BUYING A CODING MODEL USED TO BE SIMPLE
Two years ago, even last year, you bought the newest flagship from Anthropic or OpenAI, because it was visibly better than everything else and the gap was worth almost any price. Three things changed.
Agentic coding multiplied token use 5x to 30x over a chat session.
The flagships’ price per token went up rather than down, with Fable and Astra as the clearest cases.
At least 20 serious models arrived from a dozen vendors, several of them open-weight and yours to host.
The default got expensive, and the field got crowded.
We ran 60 tasks of the kind engineering teams hand to AI assistants every day, writing new code, fixing bugs, reviewing someone’s plan, investigating a problem, refusing a dangerous instruction, and the half-formed request someone types between meetings. 34 of them run against 11 open-source job-queue projects, and 26 against purpose-built private code in 11 languages that exists nowhere public. Every task had its grading criteria written before any model ran. 28 of the 60 are checked against something real: we run the model’s code, or we planted a fault and check whether it was found. The other 32 are graded by an AI judge working from the pre-written criteria, with the model’s name stripped out first. We logged what every run consumed and priced it at list rates.
1. THE FRONTIER IS CROWDED, AND THE GROUND JUST BEHIND IT IS CHEAP
Economists call it an efficient frontier, the set of choices where you can’t do better on one measure without giving something up on the other. In plain terms, a model is on the frontier if nothing else is both cheaper and better. Here’s the whole field on one chart.
Efficient frontier
Proprietary
Open-weight
20
30
40
50
60
$0.5
$1
$2
$5
$10
$20
$50
$100
$200
Cost of the full 60-task run (log scale)
Score out of 60
claude-opus-5
gpt-6-astra
claude-fable-5.1
gemini-3.8-flash
deepseek-v4.1-flash
grok-4.6
glm-5.3-flash
grok-4.5
glm-5.3
gpt-5.6-sol
claude-opus-4-8
gpt-5.6-terra
claude-sonnet-5
minimax-m3
deepseek-v4-flash
claude-sonnet-4-6
qwen3.7-plus
inkling
nemotron-3-ultra
gemini-3.1-flash-lite
gemini-3.1-pro
claude-haiku-4-5
gpt-oss-120b
Score against cost. Each dot is one model, cost is the full 60-task run at list rates checked September 2026, on a log scale.
A model on the blue line has no competitor that is both cheaper and better.
Efficient frontier
Proprietary
Open-weight
20
30
40
50
60
$0.5
$2
$10
$50
$200
Cost of the full 60-task run (log scale)
Score out of 60
claude-opus-5
deepseek-v4.1-flash
gpt-oss-120b
Score against cost. Each dot is one model, cost is the full
60-task run at list rates checked September 2026, on a
log scale. A model on the blue line has no competitor that
is both cheaper and better.
Only 8 models sit on that line. The other 21 have at least one competitor above and to their left, cheaper and better at the same time. Some of the 21 are famous.
One thing to hold onto while reading the rest. Each model saw each task once. Re-run a model on the same 60 tasks and its score moves by about 2 points, with 1 verdict in 4 flipping between a pass and a partial. So treat every score here as plus or minus 2, and treat 2 models within 3 or 4 points of each other as tied. By that rule the top 8 are one group, from claude-opus-5 at 54.5 down to glm-5.3-flash at 50, at prices from $2.58 to $135.51.
The top of the chart is where the pricing story gets uncomfortable. claude-opus-5 leads at 54.5 out of 60, for $135.51. OpenAI’s gpt-6-astra is half a point behind at $129.26. claude-fable-5.1 scored 52.5 for $85.48. Then the price drops off a cliff. gemini-3.8-flash scored 51.5 for $17.29, and 2 open-weight models, deepseek-v4.1-flash and glm-5.3-flash, scored 50.5 and 50 for $2.67 and $2.58.
The gap between the best model in the study and a model that costs 2% as much is 4 points out of 60, about what one model moves between 2 runs.
The same data reads even more clearly as a ladder, the cheapest model that reaches each score level, and what the next step up costs.
The cost cliff. The cheapest model that reaches each score level, on a log scale, and what the next rung up costs. Read from the bottom, the first $2.67 buys 50.5 points, and the next $133 buys 4 more.
The cost cliff. The cheapest model that reaches each score level, on a log scale, and what the next rung up costs. Read from the bottom, the first $2.67 buys 50.5 points, and the next $133 buys 4 more.
$2.67 buys a score of 50.5. The next point costs $15 more, the point after that $68 more, and the last 2 points, from 52.5 up to 54.5, cost another $50. Every rung from $2.67 up is inside the noise of the one above it. Plenty of production work lives comfortably on the bottom rungs, at 4 or 5 cents a task.
This analysis uses list prices as of September 2026, and both the prices and the frontier keep moving. Nearly every model in the top 10 launched in the last 60 days. That is the pace. The overall point, that the frontier holds a crowded set of models with huge price differences, is what holds.
2. THE PRICE ON THE RATE CARD ISN’T REFLECTIVE OF THE PRICE YOU PAY
Every vendor publishes a price per million tokens, and that’s the number people compare. For an agentic workload it no longer tells you much. An agent doesn’t consume a fixed quantity of tokens per task. It reads files, runs commands, thinks, re-reads, tries again, and every one of those steps sends the whole conversation back through the model. 2 models with the same list price can land very different bills for the same job, depending on how they go about it.
So we divided each model’s total bill by the number of tasks it fully passed. That’s the cost of getting a job done, failures included, and it’s the number to put next to the list price. The chart shows the 12 models that fully passed at least 40 of the 60 tasks. The other 17 are left off on purpose. Several are cheaper per pass than anything on the chart, gpt-oss-120b among them at 2 cents per task, but a model that fails a third or more of what you hand it is not reliable enough for production.
List price per M tokens, in / out
Cost per successful task, each model’s full 60-task bill divided by the tasks it fully passed, for the 12 models that fully passed at least 40 of the 60. The grey figures at right are the vendor’s list price per million input and output tokens. An open circle marks an open-weight model.
claude-fable-5.1 is the clearest case. Anthropic lists it at $10 per million input tokens and $50 per million output, double claude-opus-5’s $5 and $25. On paper it’s the most expensive model in the study, tied with gpt-6-astra. In practice its 60-task run cost $85.48 against Opus’s $135.51, and both models fully passed 49 of the 60. Per successful task, Fable cost $1.74 and Opus $2.77.
The reason is in the usage logs. On a typical task Fable took 13 steps to Opus’s 17, made 10.5 calls to the model against 20.8, and wrote about 6,700 output tokens against 21,600. Over the whole run Opus pushed 68.7 million tokens through the model and Fable 29.7 million. Fable also reads its cache at a quarter the price, $0.25 per million against $0.50, which matters when 3 quarters of everything an agent sends is something it already sent.
The averages hide the mechanism. On the typical task the 2 cost about the same, 80 cents each at the median, and Fable came out cheaper on only 25 of the 60. What decides the bill is the tail. Opus had 10 tasks that cost more than $5 each, Fable had 5, and Opus’s 10 most expensive tasks were 56% of its entire bill. Fable’s advantage is that it rarely spirals.
Two tasks show what a spiral looks like.
On the first, someone asks how you’d tell whether a queued job is running or wedged. Opus took 45 steps and 65 model calls, ran for 14 minutes, and produced a correct answer for $9.84. Fable read the code, took 9 steps and 5 model calls, and produced a correct answer in 69 seconds for $0.71. Of the 29 models, 16 answered in a single step without opening the code at all, and every one of them scored a partial or a fail.
The second asks what states a task can be in and what moves it between them, in a Go task queue. Here, 10 models answered without opening a file, and all 10 failed. Opus read the code and passed in only 7 steps, but spun up sub-agents along the way and made 67 model calls, 6 minutes and $7.77. Fable read the code and passed in 17 steps and 9 model calls, 92 seconds and $1.22.
The same pattern holds at OpenAI. gpt-6-astra lists at $10 and $50, gpt-5.6-sol at $4 and $20. Astra’s run cost $129.26 and scored 54. Sol’s cost $158.13 and scored 47.5. Sol made 27 model calls per task to Astra’s 17 and wrote 4x the output. Per successful task, the model that costs 2.5x as much on the rate card was 30% cheaper.
It cuts the other way too. gemini-3.8-flash is priced at 3x to 6x deepseek-v4.1-flash per token, depending on the token type, and cost 6.5x as much to run, because it took a median of 55 steps per task, more than any other model in the study, against DeepSeek’s 17.
Steps cost time as well as money. On the typical task Fable answered in 75 seconds and Opus in 139, and on the tasks both passed the gap was 69 seconds against 114. gemini-3.8-flash shows the two measures pulling apart. Per round trip it’s one of the quickest in the study, 5.7 seconds against Fable’s 10.9 and Opus’s 11.5, and one of the slowest to finish a task, 115 seconds at the median, because it takes so many round trips. Asked what the system does when a worker dies mid-task and how you’d confirm it happened, Gemini took 113 steps

[truncated]

## Original Extract

We tested 29 AI coding models on 60 real tasks and priced every run. One model reaches 93% of the top score for 2% of the cost. See the full data.

What AI coding models really cost, 29 models benchmarked | Bito
Skip to content
We tested 29 AI models to measure the cost-quality frontier.
See who won
Start free
Products
Context engine
What a Crowded AI Frontier Means for Engineering Teams
How real-world scenarios and a plethora of options dramatically change the cost-quality frontier. Based on 29 AI models, 60 real engineering jobs, and 1,740 graded answers.
We tested 29 AI models on the same 60 real engineering tasks, graded every one of the 1,740 answers against criteria written before any model ran, and priced every run. That combination, quality and cost measured together on the same work, is what this research is about.
Four findings matter for anyone deciding what AI models to buy for their engineering teams.
Most of the market is overpriced for what it delivers
21 of the 29 models are beaten by something cheaper that scores at least as well. claude-opus-5 leads at 54.5 out of 60 and cost $135.51 for the full run. Another 7 models land within 4.5 points of it, at prices down to $2.58. deepseek-v4.1-flash, an open-weight model, scored 50.5 for $2.67. That is 93% of the top score for 2% of the money.
The rate card no longer predicts your bill
claude-fable-5.1 lists at 2x the per-token price of claude-opus-5 and cost 37% less to run the same 60 jobs, because it took fewer steps, made half as many model calls, and wrote a third as much output. gpt-6-astra lists at 2.5x gpt-5.6-sol, came in 18% cheaper, and scored 6.5 points higher. The number to buy on is cost per successful task, and we report it for every model.
The score alone cannot tell you which model to trust
Two behaviors separate the models that hold up in daily use. What a model does when your request is vague, and what it does when it forgets an instruction you gave it. We scored both, model by model, and the results rarely follow price. Of the 29 models, 15 clear both bars, and the ones that miss include several that look sensible on a leaderboard.
Tiering the work buys the top score for a fraction of the money
Sort the work by how much judgment it takes, put a different model on each tier, and the top group’s score survives at a fraction of the bill. Opus on the tier where it is clearly better, with cheap models underneath, scored 53.5 for $18. Two open-weight models on their own scored 52.5 for $2.50. Running claude-opus-5 on everything scored 54.5 for $135.51.
BUYING A CODING MODEL USED TO BE SIMPLE
Two years ago, even last year, you bought the newest flagship from Anthropic or OpenAI, because it was visibly better than everything else and the gap was worth almost any price. Three things changed.
Agentic coding multiplied token use 5x to 30x over a chat session.
The flagships’ price per token went up rather than down, with Fable and Astra as the clearest cases.
At least 20 serious models arrived from a dozen vendors, several of them open-weight and yours to host.
The default got expensive, and the field got crowded.
We ran 60 tasks of the kind engineering teams hand to AI assistants every day, writing new code, fixing bugs, reviewing someone’s plan, investigating a problem, refusing a dangerous instruction, and the half-formed request someone types between meetings. 34 of them run against 11 open-source job-queue projects, and 26 against purpose-built private code in 11 languages that exists nowhere public. Every task had its grading criteria written before any model ran. 28 of the 60 are checked against something real: we run the model’s code, or we planted a fault and check whether it was found. The other 32 are graded by an AI judge working from the pre-written criteria, with the model’s name stripped out first. We logged what every run consumed and priced it at list rates.
1. THE FRONTIER IS CROWDED, AND THE GROUND JUST BEHIND IT IS CHEAP
Economists call it an efficient frontier, the set of choices where you can’t do better on one measure without giving something up on the other. In plain terms, a model is on the frontier if nothing else is both cheaper and better. Here’s the whole field on one chart.
Efficient frontier
Proprietary
Open-weight
20
30
40
50
60
$0.5
$1
$2
$5
$10
$20
$50
$100
$200
Cost of the full 60-task run (log scale)
Score out of 60
claude-opus-5
gpt-6-astra
claude-fable-5.1
gemini-3.8-flash
deepseek-v4.1-flash
grok-4.6
glm-5.3-flash
grok-4.5
glm-5.3
gpt-5.6-sol
claude-opus-4-8
gpt-5.6-terra
claude-sonnet-5
minimax-m3
deepseek-v4-flash
claude-sonnet-4-6
qwen3.7-plus
inkling
nemotron-3-ultra
gemini-3.1-flash-lite
gemini-3.1-pro
claude-haiku-4-5
gpt-oss-120b
Score against cost. Each dot is one model, cost is the full 60-task run at list rates checked September 2026, on a log scale.
A model on the blue line has no competitor that is both cheaper and better.
Efficient frontier
Proprietary
Open-weight
20
30
40
50
60
$0.5
$2
$10
$50
$200
Cost of the full 60-task run (log scale)
Score out of 60
claude-opus-5
deepseek-v4.1-flash
gpt-oss-120b
Score against cost. Each dot is one model, cost is the full
60-task run at list rates checked September 2026, on a
log scale. A model on the blue line has no competitor that
is both cheaper and better.
Only 8 models sit on that line. The other 21 have at least one competitor above and to their left, cheaper and better at the same time. Some of the 21 are famous.
One thing to hold onto while reading the rest. Each model saw each task once. Re-run a model on the same 60 tasks and its score moves by about 2 points, with 1 verdict in 4 flipping between a pass and a partial. So treat every score here as plus or minus 2, and treat 2 models within 3 or 4 points of each other as tied. By that rule the top 8 are one group, from claude-opus-5 at 54.5 down to glm-5.3-flash at 50, at prices from $2.58 to $135.51.
The top of the chart is where the pricing story gets uncomfortable. claude-opus-5 leads at 54.5 out of 60, for $135.51. OpenAI’s gpt-6-astra is half a point behind at $129.26. claude-fable-5.1 scored 52.5 for $85.48. Then the price drops off a cliff. gemini-3.8-flash scored 51.5 for $17.29, and 2 open-weight models, deepseek-v4.1-flash and glm-5.3-flash, scored 50.5 and 50 for $2.67 and $2.58.
The gap between the best model in the study and a model that costs 2% as much is 4 points out of 60, about what one model moves between 2 runs.
The same data reads even more clearly as a ladder, the cheapest model that reaches each score level, and what the next step up costs.
The cost cliff. The cheapest model that reaches each score level, on a log scale, and what the next rung up costs. Read from the bottom, the first $2.67 buys 50.5 points, and the next $133 buys 4 more.
The cost cliff. The cheapest model that reaches each score level, on a log scale, and what the next rung up costs. Read from the bottom, the first $2.67 buys 50.5 points, and the next $133 buys 4 more.
$2.67 buys a score of 50.5. The next point costs $15 more, the point after that $68 more, and the last 2 points, from 52.5 up to 54.5, cost another $50. Every rung from $2.67 up is inside the noise of the one above it. Plenty of production work lives comfortably on the bottom rungs, at 4 or 5 cents a task.
This analysis uses list prices as of September 2026, and both the prices and the frontier keep moving. Nearly every model in the top 10 launched in the last 60 days. That is the pace. The overall point, that the frontier holds a crowded set of models with huge price differences, is what holds.
2. THE PRICE ON THE RATE CARD ISN’T REFLECTIVE OF THE PRICE YOU PAY
Every vendor publishes a price per million tokens, and that’s the number people compare. For an agentic workload it no longer tells you much. An agent doesn’t consume a fixed quantity of tokens per task. It reads files, runs commands, thinks, re-reads, tries again, and every one of those steps sends the whole conversation back through the model. 2 models with the same list price can land very different bills for the same job, depending on how they go about it.
So we divided each model’s total bill by the number of tasks it fully passed. That’s the cost of getting a job done, failures included, and it’s the number to put next to the list price. The chart shows the 12 models that fully passed at least 40 of the 60 tasks. The other 17 are left off on purpose. Several are cheaper per pass than anything on the chart, gpt-oss-120b among them at 2 cents per task, but a model that fails a third or more of what you hand it is not reliable enough for production.
List price per M tokens, in / out
Cost per successful task, each model’s full 60-task bill divided by the tasks it fully passed, for the 12 models that fully passed at least 40 of the 60. The grey figures at right are the vendor’s list price per million input and output tokens. An open circle marks an open-weight model.
claude-fable-5.1 is the clearest case. Anthropic lists it at $10 per million input tokens and $50 per million output, double claude-opus-5’s $5 and $25. On paper it’s the most expensive model in the study, tied with gpt-6-astra. In practice its 60-task run cost $85.48 against Opus’s $135.51, and both models fully passed 49 of the 60. Per successful task, Fable cost $1.74 and Opus $2.77.
The reason is in the usage logs. On a typical task Fable took 13 steps to Opus’s 17, made 10.5 calls to the model against 20.8, and wrote about 6,700 output tokens against 21,600. Over the whole run Opus pushed 68.7 million tokens through the model and Fable 29.7 million. Fable also reads its cache at a quarter the price, $0.25 per million against $0.50, which matters when 3 quarters of everything an agent sends is something it already sent.
The averages hide the mechanism. On the typical task the 2 cost about the same, 80 cents each at the median, and Fable came out cheaper on only 25 of the 60. What decides the bill is the tail. Opus had 10 tasks that cost more than $5 each, Fable had 5, and Opus’s 10 most expensive tasks were 56% of its entire bill. Fable’s advantage is that it rarely spirals.
Two tasks show what a spiral looks like.
On the first, someone asks how you’d tell whether a queued job is running or wedged. Opus took 45 steps and 65 model calls, ran for 14 minutes, and produced a correct answer for $9.84. Fable read the code, took 9 steps and 5 model calls, and produced a correct answer in 69 seconds for $0.71. Of the 29 models, 16 answered in a single step without opening the code at all, and every one of them scored a partial or a fail.
The second asks what states a task can be in and what moves it between them, in a Go task queue. Here, 10 models answered without opening a file, and all 10 failed. Opus read the code and passed in only 7 steps, but spun up sub-agents along the way and made 67 model calls, 6 minutes and $7.77. Fable read the code and passed in 17 steps and 9 model calls, 92 seconds and $1.22.
The same pattern holds at OpenAI. gpt-6-astra lists at $10 and $50, gpt-5.6-sol at $4 and $20. Astra’s run cost $129.26 and scored 54. Sol’s cost $158.13 and scored 47.5. Sol made 27 model calls per task to Astra’s 17 and wrote 4x the output. Per successful task, the model that costs 2.5x as much on the rate card was 30% cheaper.
It cuts the other way too. gemini-3.8-flash is priced at 3x to 6x deepseek-v4.1-flash per token, depending on the token type, and cost 6.5x as much to run, because it took a median of 55 steps per task, more than any other model in the study, against DeepSeek’s 17.
Steps cost time as well as money. On the typical task Fable answered in 75 seconds and Opus in 139, and on the tasks both passed the gap was 69 seconds against 114. gemini-3.8-flash shows the two measures pulling apart. Per round trip it’s one of the quickest in the study, 5.7 seconds against Fable’s 10.9 and Opus’s 11.5, and one of the slowest to finish a task, 115 seconds at the median, because it takes so many round trips. Asked what the system does when a worker dies mid-task and how you’d confirm it happened, Gemini took 113 steps

[truncated]
