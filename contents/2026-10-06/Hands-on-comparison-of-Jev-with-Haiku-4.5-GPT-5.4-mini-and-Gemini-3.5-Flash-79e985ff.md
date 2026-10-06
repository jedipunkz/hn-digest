---
source: "https://blog.oodle.ai/is-jev-the-right-model-for-your-use-case-a-walkthrough-comparing-haiku-4-5-gpt-5-4-mini-and-gemini-3-5-flash/"
hn_url: "https://news.ycombinator.com/item?id=49982760"
title: "Hands-on comparison of Jev with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash"
article_title: "Is Jev the right model for your agents? A hands-on comparison with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash"
image: "https://storage.ghost.io/c/b2/48/b2485d36-0dcd-4e46-b7bd-418ab6e38e18/content/images/size/w1200/2026/09/jev_blog_image.png"
author: "akrish84"
captured_at: "2026-10-06T19:23:23Z"
capture_tool: "hn-digest"
hn_id: 49982760
score: 1
comments: 0
posted_at: "2026-10-06T19:16:17Z"
tags:
  - hacker-news
---

# Hands-on comparison of Jev with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash

- HN: [49982760](https://news.ycombinator.com/item?id=49982760)
- Source: [blog.oodle.ai](https://blog.oodle.ai/is-jev-the-right-model-for-your-use-case-a-walkthrough-comparing-haiku-4-5-gpt-5-4-mini-and-gemini-3-5-flash/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T19:16:17Z

## Translation

Title: Hands-on comparison of Jev with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash
Article title: Is Jev the right model for your agents? A hands-on comparison with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash
Description: A working example of how to evaluate whether Jev is the right model for your use case, with two datasets, four models and three judges, and accuracy, cost and latency for each.

Article text:
Sign in
Is Jev the right model for your agents? A hands-on comparison with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash
The Premise: Jev is insanely fast and cheap, but it only outputs probabilities for fixed answers, not generated text.
The Test: I compared Jev against GPT-5.4-mini, Claude Haiku 4.5, and Gemini 3.5 Flash on two Agent use cases: Support Escalation and Model Routing to see if Jev is a better model to use.
The Takeaway: Jev held its own (and even beat Haiku on one test), but its raw accuracy numbers don't tell the whole story. Read on to see the data, the cost breakdowns, and why tuning Jev's probability threshold changes everything.
It's been less than a week since Jev launched, and so many people have already tested it, written it up and talked about it that publishing this a week in feels late. The valley moves really fast.
Jev is TypeSafe AI's System One model. It is an incredibly cheap and fast model designed to ace multiple-choice classifications rather than write generative text. The model was built by a team led by Diogo Almeida , one of the original researchers behind ChatGPT and a co-inventor of RLHF. Hearing such big names behind a product automatically creates a sense of seriousness about its quality. I won't go in depth describing it here since there are already great resources like this guide on Jev . Instead, I am going to focus on the testing methodology I used to compare it against other models and how to know if Jev is the right choice for your agents.
the methodology I used to compare Jev against other models on two use cases, and how you can apply it to your own domain;
how to decide, from those results, whether Jev is the right model for a given type of use case.
And maybe, along the way, you'll get a sense of which kinds of use cases work well with Jev, a small glimpse into what is probably a much larger world of them.
I think of Jev like a student taking a test. It is expected to completely ace the multiple choice section. If you give it a question and a list of possible answers, it picks the right one incredibly fast and for very little cost. But it would fail the essay section right away because it cannot actually write out new text. Jev offers a few different answer shapes (a yes/no, one option out of a set, or a level on a scale), and it returns a probability over those answers rather than text.
The flow, end to end, is simple, and nothing in it is specific to my two examples. Take a dataset from your domain, run it against each model you are considering, have the same evaluators score every output against the expected output, then compare accuracy, cost and latency side by side and see where Jev lands.
Datasets. A dataset is a list of inputs, a prompt, and for each input the expected output that a human wrote or verified. For your domain that means the real inputs your system sees today, the prompt you would actually run, and the answer a person on your team agrees is right. Now, a large number of people might be thinking, "I don't have a dataset to begin with," and that's a fair place to be. Building one deserves its own post, so I'll leave it out of scope here.
Experiments. An experiment is the dataset run through a single model. Every input goes in with the prompt, and the model's output for each one is evaluated against the expected output. One model, one experiment, so run one for Jev and one for each model you would otherwise use. For me that was four in total.
Jev ( jev-latest , through OpenRouter)
I am deliberately not using the leading frontier models. Pick the models you would actually run a classifier on at scale, which for most teams are the smaller, cost-efficient ones, and those are the fair comparison for Jev.
Evaluators. An evaluator scores each output of an experiment against the expected output. Whatever evaluators you pick, attach the same ones to every experiment so every model is scored the same way. I used three judges and averaged them. The judges and their prompts are listed in their own section below.
What I'm measuring. Why even consider Jev in the first place? For me it comes down to three things, accuracy , cost and latency . Cost and latency are the promise. Jev is supposed to answer in a fraction of the time an LLM takes, and to be an order of magnitude cheaper than even the small, cost-efficient LLMs, and a couple of orders cheaper than the frontier ones. Accuracy is the question. What, if anything, do those savings cost you? For every experiment, these are the results we look at.
Accuracy. Which model scores highest, and where does Jev land in that order?
Cost. What did each model cost me for the whole experiment?
Latency. How long did Jev take per item, and how long did the others take?
At the end of the day the best-scoring model isn't always the one you want. Depending on the use case you can make trade-offs, and only you know which ones your use case can afford. If Jev is a lot cheaper and faster, I'm okay giving up a little accuracy for the right use case. That's the lens I'm looking through.
I ran all of this in Oodle as datasets and experiments, one experiment per model, with the same three judges attached to each.
I wanted one dataset that I resonated with and that I felt Jev would do well on. For the second, I wanted one where Jev looked like a good fit on paper, but where I was skeptical about the quality it would reach and how it would compare with the other models. I really wanted to find out.
Scenario 1: Automating Support Escalations
If you've ever talked to a support AI agent, you've been here. At some point the agent isn't giving you what you need, and you just want a human to take over.
So the question for this dataset is whether we can detect, from the sentiment of the conversation, when a support chat should be escalated to a human. The input is the whole chat between the customer and the agent; the expected output is whether the customer's sentiment is negative enough that a human should step in.
I know sentiment is a well-known, solved ML problem. But these days people reach for an LLM for every question of this kind, so let's do the comparison. And this is exactly the classification shape that Jev is built for, right? Let's put it to the test. I'm writing this before actually running the experiment, and I'm rooting for Jev on this one.
The conversations come from NebulaByte/E-Commerce_Customer_Support_Conversations on Hugging Face, 1,000 e-commerce support chats, each with the customer's sentiment recorded next to it. I took 200 of them, not 1,000, for the very scientific reason that I couldn't get hold of the company credit card in time and was running this experiment on my own. Balanced, at least, with 100 where the customer is negative or frustrated (escalate yes) and 100 where they are neutral or positive (escalate no).
Input (the conversation field, shortened).
Agent: Hello, thank you for contacting BrownBox customer support. My name is Alex. How may I assist you today?
Customer: Hi, I'm having issues with my air conditioner return and refund. I returned it two weeks ago, and I haven't received my refund yet. Can you help me with that?
Agent: I'm sorry to hear that, sir. May I know your order number, please? […] I will escalate this issue to our refunds team, and they will contact you within 48 hours with an update.
Customer: 48 hours? That's too long. I need my refund as soon as possible. I'm very frustrated with this situation. […]
Customer: No, that's all for now. Just please keep me updated on the progress of my refund.
Expected output {"escalate": "yes", "sentiment": "frustrated"}
The prompt every model was given
You review customer support conversations for a support team.
Read the conversation below and decide whether the customer's
sentiment is negative or frustrated, so that the case must be
escalated to a human supervisor.
Conversation:
{{conversation}}
Answer with exactly one word: yes or no. {{conversation}} is filled from each item's input. For Jev the same text is the question and the conversation goes in as the state; the answer shape is yes/no.
Scenario 2: Smart Model Routing
The second dataset is the one I'm really curious about, because here the model has to judge what the input is actually asking for , not just how it reads. I suspect Jev will score below the other LLMs on this one, even though the shape of the question is exactly what it is built for. I'm hoping it's not too far off.
The use case is to decide, given a user's prompt, whether it needs an expensive frontier model or whether a cheaper model is enough. The output is just one of the two. The prompt I give the models spells out when to prefer each, so what we're really testing is how well Jev and the LLMs can read the user's request, weigh it against the conditions I gave them, and route accordingly.
The data comes from routellm/gpt4_dataset , the LMSYS RouteLLM project. Every row is a real user prompt from Chatbot Arena and similar sources, answered by both a mid-size open model (Mixtral 8x7B) and GPT-4, with GPT-4 grading Mixtral's answer from 1 to 5. I took 200 prompts, balanced, with 100 where Mixtral's answer was graded 3 or lower (needs the frontier model, so yes) and 100 where it was graded 5 (no).
Input "I want to have two flutter choice chips in a row, the two are of the same size, but will be adaptive as the screen width" → Expected output no
Input "Can you design an interactive webpage using HTML and JavaScript that features a table showcasing four distinct professions…" → Expected output yes
The prompt every model was given
You are a router in front of two language models: a mid-size open
model (Mixtral 8x7B class), cheap and fast, and a frontier model
(GPT-4 class), expensive.
Read the user's request below. Would the mid-size model's answer be
noticeably worse than the frontier model's, so that the request must
go to the frontier model?
Route to the frontier model when the request needs multi-step
reasoning, precise code or maths, expert or niche knowledge, strict
instruction following, or long structured output. Simple chat, short
factual questions, rewording, summaries and casual creative writing
do not need it.
User request:
{{prompt}}
Answer with exactly one word: yes (frontier model) or no (mid-size
model is enough). The AI Judges
Each evaluator is a judge, a model of its own that is given the model's output and the expected output and returns a score for how well they match. I used three judges on every experiment, Gemini 3.5 Flash, Claude Haiku 4.5 and GPT-5.4-mini, and averaged their scores. Three rather than one so that a single judge's quirks do not decide the result. Each judge scores an item 0 or 1, so when they disagree the average for that item lands in between, and in practice they disagreed on very few items. Why those three? I asked Claude which models would make a fair, comparable set, and it named these. If your favourite is missing, I apologise in advance, and you can take it up with Claude.
The prompt used for judging in the support escalation dataset.
A model read a customer support conversation and was asked: is the
customer's sentiment negative, so that the case must be escalated to
a human? It had to answer yes or no.
Expected answer (the `escalate` field is the correct yes/no;
`sentiment` is the human label):
{{expected_output}}
Model output:
{{output}}
Score 1 if the model output gives the same yes/no decision as the
expected `escalate` value. Score 0 if it gives the opposite decision,
no decision, or anything that is not a clear yes or no.
Respond with JSON:
{"score": <0 or 1>, "reasoning": "<one sentence>"} And here is the prompt used by the judge for the model routing dataset.
A model was asked a question with a fixed answer format (a single word
or a single number). Compare its answer with the expected answer.
Expected answer:
{{expected_output}}
Model output:
{{output}}
Score 1 if the model output states the same answer as the expected
answer. I

[truncated]

## Original Extract

A working example of how to evaluate whether Jev is the right model for your use case, with two datasets, four models and three judges, and accuracy, cost and latency for each.

Sign in
Is Jev the right model for your agents? A hands-on comparison with Haiku 4.5, GPT-5.4-mini, and Gemini 3.5 Flash
The Premise: Jev is insanely fast and cheap, but it only outputs probabilities for fixed answers, not generated text.
The Test: I compared Jev against GPT-5.4-mini, Claude Haiku 4.5, and Gemini 3.5 Flash on two Agent use cases: Support Escalation and Model Routing to see if Jev is a better model to use.
The Takeaway: Jev held its own (and even beat Haiku on one test), but its raw accuracy numbers don't tell the whole story. Read on to see the data, the cost breakdowns, and why tuning Jev's probability threshold changes everything.
It's been less than a week since Jev launched, and so many people have already tested it, written it up and talked about it that publishing this a week in feels late. The valley moves really fast.
Jev is TypeSafe AI's System One model. It is an incredibly cheap and fast model designed to ace multiple-choice classifications rather than write generative text. The model was built by a team led by Diogo Almeida , one of the original researchers behind ChatGPT and a co-inventor of RLHF. Hearing such big names behind a product automatically creates a sense of seriousness about its quality. I won't go in depth describing it here since there are already great resources like this guide on Jev . Instead, I am going to focus on the testing methodology I used to compare it against other models and how to know if Jev is the right choice for your agents.
the methodology I used to compare Jev against other models on two use cases, and how you can apply it to your own domain;
how to decide, from those results, whether Jev is the right model for a given type of use case.
And maybe, along the way, you'll get a sense of which kinds of use cases work well with Jev, a small glimpse into what is probably a much larger world of them.
I think of Jev like a student taking a test. It is expected to completely ace the multiple choice section. If you give it a question and a list of possible answers, it picks the right one incredibly fast and for very little cost. But it would fail the essay section right away because it cannot actually write out new text. Jev offers a few different answer shapes (a yes/no, one option out of a set, or a level on a scale), and it returns a probability over those answers rather than text.
The flow, end to end, is simple, and nothing in it is specific to my two examples. Take a dataset from your domain, run it against each model you are considering, have the same evaluators score every output against the expected output, then compare accuracy, cost and latency side by side and see where Jev lands.
Datasets. A dataset is a list of inputs, a prompt, and for each input the expected output that a human wrote or verified. For your domain that means the real inputs your system sees today, the prompt you would actually run, and the answer a person on your team agrees is right. Now, a large number of people might be thinking, "I don't have a dataset to begin with," and that's a fair place to be. Building one deserves its own post, so I'll leave it out of scope here.
Experiments. An experiment is the dataset run through a single model. Every input goes in with the prompt, and the model's output for each one is evaluated against the expected output. One model, one experiment, so run one for Jev and one for each model you would otherwise use. For me that was four in total.
Jev ( jev-latest , through OpenRouter)
I am deliberately not using the leading frontier models. Pick the models you would actually run a classifier on at scale, which for most teams are the smaller, cost-efficient ones, and those are the fair comparison for Jev.
Evaluators. An evaluator scores each output of an experiment against the expected output. Whatever evaluators you pick, attach the same ones to every experiment so every model is scored the same way. I used three judges and averaged them. The judges and their prompts are listed in their own section below.
What I'm measuring. Why even consider Jev in the first place? For me it comes down to three things, accuracy , cost and latency . Cost and latency are the promise. Jev is supposed to answer in a fraction of the time an LLM takes, and to be an order of magnitude cheaper than even the small, cost-efficient LLMs, and a couple of orders cheaper than the frontier ones. Accuracy is the question. What, if anything, do those savings cost you? For every experiment, these are the results we look at.
Accuracy. Which model scores highest, and where does Jev land in that order?
Cost. What did each model cost me for the whole experiment?
Latency. How long did Jev take per item, and how long did the others take?
At the end of the day the best-scoring model isn't always the one you want. Depending on the use case you can make trade-offs, and only you know which ones your use case can afford. If Jev is a lot cheaper and faster, I'm okay giving up a little accuracy for the right use case. That's the lens I'm looking through.
I ran all of this in Oodle as datasets and experiments, one experiment per model, with the same three judges attached to each.
I wanted one dataset that I resonated with and that I felt Jev would do well on. For the second, I wanted one where Jev looked like a good fit on paper, but where I was skeptical about the quality it would reach and how it would compare with the other models. I really wanted to find out.
Scenario 1: Automating Support Escalations
If you've ever talked to a support AI agent, you've been here. At some point the agent isn't giving you what you need, and you just want a human to take over.
So the question for this dataset is whether we can detect, from the sentiment of the conversation, when a support chat should be escalated to a human. The input is the whole chat between the customer and the agent; the expected output is whether the customer's sentiment is negative enough that a human should step in.
I know sentiment is a well-known, solved ML problem. But these days people reach for an LLM for every question of this kind, so let's do the comparison. And this is exactly the classification shape that Jev is built for, right? Let's put it to the test. I'm writing this before actually running the experiment, and I'm rooting for Jev on this one.
The conversations come from NebulaByte/E-Commerce_Customer_Support_Conversations on Hugging Face, 1,000 e-commerce support chats, each with the customer's sentiment recorded next to it. I took 200 of them, not 1,000, for the very scientific reason that I couldn't get hold of the company credit card in time and was running this experiment on my own. Balanced, at least, with 100 where the customer is negative or frustrated (escalate yes) and 100 where they are neutral or positive (escalate no).
Input (the conversation field, shortened).
Agent: Hello, thank you for contacting BrownBox customer support. My name is Alex. How may I assist you today?
Customer: Hi, I'm having issues with my air conditioner return and refund. I returned it two weeks ago, and I haven't received my refund yet. Can you help me with that?
Agent: I'm sorry to hear that, sir. May I know your order number, please? […] I will escalate this issue to our refunds team, and they will contact you within 48 hours with an update.
Customer: 48 hours? That's too long. I need my refund as soon as possible. I'm very frustrated with this situation. […]
Customer: No, that's all for now. Just please keep me updated on the progress of my refund.
Expected output {"escalate": "yes", "sentiment": "frustrated"}
The prompt every model was given
You review customer support conversations for a support team.
Read the conversation below and decide whether the customer's
sentiment is negative or frustrated, so that the case must be
escalated to a human supervisor.
Conversation:
{{conversation}}
Answer with exactly one word: yes or no. {{conversation}} is filled from each item's input. For Jev the same text is the question and the conversation goes in as the state; the answer shape is yes/no.
Scenario 2: Smart Model Routing
The second dataset is the one I'm really curious about, because here the model has to judge what the input is actually asking for , not just how it reads. I suspect Jev will score below the other LLMs on this one, even though the shape of the question is exactly what it is built for. I'm hoping it's not too far off.
The use case is to decide, given a user's prompt, whether it needs an expensive frontier model or whether a cheaper model is enough. The output is just one of the two. The prompt I give the models spells out when to prefer each, so what we're really testing is how well Jev and the LLMs can read the user's request, weigh it against the conditions I gave them, and route accordingly.
The data comes from routellm/gpt4_dataset , the LMSYS RouteLLM project. Every row is a real user prompt from Chatbot Arena and similar sources, answered by both a mid-size open model (Mixtral 8x7B) and GPT-4, with GPT-4 grading Mixtral's answer from 1 to 5. I took 200 prompts, balanced, with 100 where Mixtral's answer was graded 3 or lower (needs the frontier model, so yes) and 100 where it was graded 5 (no).
Input "I want to have two flutter choice chips in a row, the two are of the same size, but will be adaptive as the screen width" → Expected output no
Input "Can you design an interactive webpage using HTML and JavaScript that features a table showcasing four distinct professions…" → Expected output yes
The prompt every model was given
You are a router in front of two language models: a mid-size open
model (Mixtral 8x7B class), cheap and fast, and a frontier model
(GPT-4 class), expensive.
Read the user's request below. Would the mid-size model's answer be
noticeably worse than the frontier model's, so that the request must
go to the frontier model?
Route to the frontier model when the request needs multi-step
reasoning, precise code or maths, expert or niche knowledge, strict
instruction following, or long structured output. Simple chat, short
factual questions, rewording, summaries and casual creative writing
do not need it.
User request:
{{prompt}}
Answer with exactly one word: yes (frontier model) or no (mid-size
model is enough). The AI Judges
Each evaluator is a judge, a model of its own that is given the model's output and the expected output and returns a score for how well they match. I used three judges on every experiment, Gemini 3.5 Flash, Claude Haiku 4.5 and GPT-5.4-mini, and averaged their scores. Three rather than one so that a single judge's quirks do not decide the result. Each judge scores an item 0 or 1, so when they disagree the average for that item lands in between, and in practice they disagreed on very few items. Why those three? I asked Claude which models would make a fair, comparable set, and it named these. If your favourite is missing, I apologise in advance, and you can take it up with Claude.
The prompt used for judging in the support escalation dataset.
A model read a customer support conversation and was asked: is the
customer's sentiment negative, so that the case must be escalated to
a human? It had to answer yes or no.
Expected answer (the `escalate` field is the correct yes/no;
`sentiment` is the human label):
{{expected_output}}
Model output:
{{output}}
Score 1 if the model output gives the same yes/no decision as the
expected `escalate` value. Score 0 if it gives the opposite decision,
no decision, or anything that is not a clear yes or no.
Respond with JSON:
{"score": <0 or 1>, "reasoning": "<one sentence>"} And here is the prompt used by the judge for the model routing dataset.
A model was asked a question with a fixed answer format (a single word
or a single number). Compare its answer with the expected answer.
Expected answer:
{{expected_output}}
Model output:
{{output}}
Score 1 if the model output states the same answer as the expected
answer. I

[truncated]
