---
source: "https://thoughts.jock.pl/p/five-lenses-one-brain-agent-diversity-experiment-2026"
hn_url: "https://news.ycombinator.com/item?id=49657766"
title: "Five diverse AI agents vs. five clones for 14 nights: a constant tied both"
article_title: "Diverse AI Agents vs Clones: 14 Nights, Scored, No Skill"
image: "https://substackcdn.com/image/fetch/$s_!enzA!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F4ab6fef0-3be6-41a2-9164-2299d4a75b28_2048x2048.png"
author: "joozio"
captured_at: "2026-09-11T13:18:21Z"
capture_tool: "hn-digest"
hn_id: 49657766
score: 1
comments: 0
posted_at: "2026-09-11T13:04:24Z"
tags:
  - hacker-news
---

# Five diverse AI agents vs. five clones for 14 nights: a constant tied both

- HN: [49657766](https://news.ycombinator.com/item?id=49657766)
- Source: [thoughts.jock.pl](https://thoughts.jock.pl/p/five-lenses-one-brain-agent-diversity-experiment-2026)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T13:04:24Z

## Translation

Title: Five diverse AI agents vs. five clones for 14 nights: a constant tied both
Article title: Diverse AI Agents vs Clones: 14 Nights, Scored, No Skill
Description: Five diverse AI agents vs five clones, 14 nights, 4,130 scored forecasts. The lenses held apart. A base rate I got 15x wrong tied both to a constant.

Article text:
Diverse AI Agents vs Clones: 14 Nights, Scored, No Skill
Digital Thoughts
Subscribe Sign in I Ran Five Diverse AI Agents Against Five Clones for 14 Nights. A Number I Made Up Decided the Result.
4,130 scored forecasts, pre-registered. The lenses really thought differently. A constant tied both.
Pawel Jozefiak Sep 11, 2026 5 1 Share The scoreboard I opened on the morning of 5 September said my thesis had won. Five AI agents with genuinely different lenses had beaten five identical clones on 9 of 14 nights, and their pooled Brier score was 0.0225 against the clones' 0.0275. Lower is better. I had wanted that table since the first night.
I did not publish it. The reason is this post.
Most of my last few months went into agents that work together: several of them, with different jobs, handing work to each other, with rules about what each one is allowed to see. At some point I stopped and asked the only question that matters there. When you add the second agent, what did you actually add? If it is the same model reading the same context, you added compute. More tokens, more latency, a second opinion drawn from the same well. If it knows different things and looks through a different lens, you added something else, and that something else is the whole case for building a network instead of buying more of one model.
So I wrote the claim down before I could soften it. A network of agents compounds context, not compute. Agents with deliberately different roles, contexts and information diets should beat the same number of clones of the same model, at equal budget, on a task with real answers. If they tie, the "more agents means more intelligence" story is wrong, and I would rather learn that from my own data than from a bill.
Fourteen nights later I have an answer. It is a fourth ending I never wrote down, and the thing that decided it was a sentence I wrote myself.
Look at how multi-agent gets sold right now. Panels, councils, swarms, debate rounds, a "team of specialists". The pitch is always that the group knows more than any member. Underneath most of those products sits one model wearing several system prompts, running at several times the cost of one call. I have been tracking who is actually winning in that market since February , and the panel pitch is in nearly every listing.
I could not find anyone who had put the two arms side by side at equal budget, on a task with a scoreboard, and published the number including the boring one. It is an awkward experiment to run if you are selling the thing. So I built the scoreboard instead of arguing about it.
The rules, written before anything ran
These went into the plan first, and they are the only reason the rest of this post is worth reading:
Two arms, one variable. Five clones against five diverse agents. Same model (Opus 5), same budget, same tools, same task, same morning. The only thing that changes is what each agent knows and how it is told to look.
Reality grades it. No model scoring another model's work, because a language model judging language models is the same well again. Deterministic Python, public APIs, zero human judgment in the loop.
Pre-registered. Scoring rule, pass threshold and the three possible endings written down before the first sample. A failed check is a finding, never a reason to edit the rule.
The control is built strong. The clone pack is a full theory of why things spread, applied from first principles. A rigged control proves nothing and I would know I rigged it.
Every experiment pays either way. A null result still has to leave me with a working tool.
Nulls get published. Including the ending where my thesis loses.
The task: forecast virality. Every morning at 07:10 the machinery samples 30 fresh posts from Hacker News, Reddit and X, minutes old and not yet hot. All ten agents give each post a probability of crossing a set popularity bar within 48 hours, plus their top five picks. Two days later a script checks what happened and scores everyone: Brier on every probability (squared error, so lower is better), precision on the top five. I picked it because the feedback loop is brutal and cheap. The answer arrives in 48 hours, verification costs one API call, and 30 predictions per agent per day means statistical power in two weeks rather than two years. The agents also cannot argue with the result, which is the part I like most.
The diverse five: a Hacker News native who knows ranking decay and flag risk, a Reddit native who prices silent mod removals before appeal, an X native who cares about follower graphs and link penalties, a trend historian who ignores craft and asks whether the topic has been running hot for 30 days, and one agent, cold-read, kept deliberately blind to anything recent so it can only judge the content in front of it.
Everything below runs in the open. Every chart and every number is on the live experiment page , all fourteen nights.
Finding one: prompt diversity is cosplay until you make it mechanical
Night one nearly fooled me, and it fooled me through the rationales. I read those before I read the correlations. They are good. The Hacker News lens talks about gravity and second-chance pools. The Reddit lens prices removal risk. The X lens shrugs at quality and asks who the author is. Five vocabularies, five kinds of evidence, five conclusions in prose.
Then all five wrote down 0.02.
Mean pairwise correlation inside the clone arm: 0.959. Inside the "diverse" arm: 0.909. On 26 of 30 posts the five different agents landed within a spread of 0.05, several of them on the exact same number. A metaphysics thread, a Rust tooling release and a Pygame tutorial got the same shrug from five agents that supposedly evaluate completely different things. I had built a panel of experts and got one opinion with five accents .
The only reason I know this is a line in the plan called the manipulation check: agent-to-agent correlation has to be lower in the diverse arm than in the clone arm, or the packs are a placebo and get rewritten before any result counts. Without that line, night one reads as a clean success. Ten of ten submissions valid, digest generated, daily loop running, every green light I had built for myself lit. I would have run two weeks, got a difference between the arms, and published it as evidence for diversity, when what I had measured was noise between ten near-copies.
I have been on the wrong side of this before. When I ran 255 agents to make three comics , fan-out scaled logistics and left quality where it was. When I told my agent to build an app every day , the output was competent and painfully boring for weeks before I admitted it. The first time I let four agents loose on a build , I judged the whole thing by reading what came out. All three times I noticed late. This time the trap had a tripwire in it and it went off in four hours. That is the entire argument for pre-registration, and it cost me one line in a plan I had already written.
There is research under this. A Cornell group published Correlated Errors in Large Language Models in June 2025, measured across more than 350 models. On one leaderboard dataset, two models agreed with each other 60% of the time when both were wrong, and larger, more accurate models had more correlated errors, across architectures and providers. If two frontier models from rival labs still fail together most of the time, five prompts wrapped around one model were never going to buy me independence. I was asking wording to do a job wording cannot do.
The easy repair is to tell the agents to disagree. Wider spread, zero evidence, because it manufactures exactly the result the experiment exists to test. There is a rule in the packs now that says so in writing. What I did instead was make each lens mechanical. Version one gave each agent a personality and a paragraph about what it cares about, then asked for a probability. Version two is a two-step move: the agent first reaches a verdict in its own categories and its own failure modes, then reads a probability off the range that verdict is assigned. The number comes out of the lens instead of out of the shared instinct with the lens narrating over the top.
Same metaphysics post, through the rebuilt lenses:
hn-native : was 0.03, now 0.00 to 0.02. Metaphysics bait, flag or kill.
reddit-native : was 0.03, now 0.14 to 0.22. Easy-comment title travels.
x-native : was 0.02, now 0.07 to 0.11. Away platform, payoff withheld.
trend-historian : was 0.02, now 0.07 to 0.17. Cold topic, honest width.
cold-read : was 0.02, now 0.01 to 0.03. Uncheckable central claim, capped.
Everyone within 0.02 of each other becomes a span from 0.00 to 0.22, and every step of it is principled. Nobody was told to fight. They stopped deferring to the same generic gut feeling.
The next morning the clone arm sat at 0.918 and the rebuilt diverse arm at 0.619. Herded posts fell from 26 of 30 to 21. And it held. The check passed on all fourteen nights . The correlation gap between the arms averaged 0.294, and it was wider in the second week than the first, 0.329 against 0.260, so the lenses kept separating instead of drifting back into the herd once the novelty wore off. Across the run the clones sit around 0.90 to 0.97 internal correlation, one brain in five copies, and the lenses run 0.53 to 0.75, with one strange night (24 August) where both arms spread out at once.
My opinion after fourteen nights: diversity has to be built, and a prompt cannot supply it. If a panel's members share a model and a context and differ only in the paragraph that tells them who they are, they are one forecaster with a thesaurus. Independence comes from different procedures, different information and, I now think, different providers. I split my own workflow across providers back in January and it looked like a preference at the time. It looks like the only real lever now.
Finding two: a number I guessed decided the scoreboard
Night one had a second fix, and at the time it felt like the most grown-up thing I did all evening. The shared hygiene text every agent gets said the base rate of a post going hot was "single digit percent", which is where that uniform floor of 0.01 to 0.03 came from. I replaced it with what I called the honest number: 10 to 15%, or 3 to 5 hot posts in a 30-post sample. I had accidentally coached the panel into the herd, I thought, and now I had un-coached it.
Fourteen nights, 416 scored posts per agent, and three of them crossed the bar. Three. The realized rate was 0.72% . I had spent an evening carefully coaching ten agents to expect fifteen times that, and they obliged: both arms forecast around 11% on average for the whole fortnight, because packs do what they are told. The "single digit percent" I overwrote sat closer to reality than my correction. And the block went into the clone pack too (it is shared, byte for byte, across all six packs), so on the same night I wrote "a control you keep improving stops being a control" into my plan, I made the control's calibration worse. Both arms were miscoached by the same amount, so the comparison between them stayed fair. What moved was the level, for everyone.
It got worse as the data came in. At four nights the realized rate was 1.72%. At seven, 0.97%. At fourteen, 0.72%. My coached figure stood still the whole time, so the gap widened every week, and the share of each arm's score explained by that one number went 60%, then 71%, then 74% for the clones, and 49%, 62%, 68% for the lenses. I assumed more data would average it out. More data made it bigger.
Here is why that swamps everything. Brier is squared error on a probability. Answer 0.11 on a post that stays cold and you pay 0.11 squared for that slot. Answer near the rate that occurred and you pay almost nothing. Now do that on nearly every one of 416 slots, because on this task nearly every post stays cold. The level of your answers becomes most of your score, and how well you read the posts is a small term sit

[truncated]

## Original Extract

Five diverse AI agents vs five clones, 14 nights, 4,130 scored forecasts. The lenses held apart. A base rate I got 15x wrong tied both to a constant.

Diverse AI Agents vs Clones: 14 Nights, Scored, No Skill
Digital Thoughts
Subscribe Sign in I Ran Five Diverse AI Agents Against Five Clones for 14 Nights. A Number I Made Up Decided the Result.
4,130 scored forecasts, pre-registered. The lenses really thought differently. A constant tied both.
Pawel Jozefiak Sep 11, 2026 5 1 Share The scoreboard I opened on the morning of 5 September said my thesis had won. Five AI agents with genuinely different lenses had beaten five identical clones on 9 of 14 nights, and their pooled Brier score was 0.0225 against the clones' 0.0275. Lower is better. I had wanted that table since the first night.
I did not publish it. The reason is this post.
Most of my last few months went into agents that work together: several of them, with different jobs, handing work to each other, with rules about what each one is allowed to see. At some point I stopped and asked the only question that matters there. When you add the second agent, what did you actually add? If it is the same model reading the same context, you added compute. More tokens, more latency, a second opinion drawn from the same well. If it knows different things and looks through a different lens, you added something else, and that something else is the whole case for building a network instead of buying more of one model.
So I wrote the claim down before I could soften it. A network of agents compounds context, not compute. Agents with deliberately different roles, contexts and information diets should beat the same number of clones of the same model, at equal budget, on a task with real answers. If they tie, the "more agents means more intelligence" story is wrong, and I would rather learn that from my own data than from a bill.
Fourteen nights later I have an answer. It is a fourth ending I never wrote down, and the thing that decided it was a sentence I wrote myself.
Look at how multi-agent gets sold right now. Panels, councils, swarms, debate rounds, a "team of specialists". The pitch is always that the group knows more than any member. Underneath most of those products sits one model wearing several system prompts, running at several times the cost of one call. I have been tracking who is actually winning in that market since February , and the panel pitch is in nearly every listing.
I could not find anyone who had put the two arms side by side at equal budget, on a task with a scoreboard, and published the number including the boring one. It is an awkward experiment to run if you are selling the thing. So I built the scoreboard instead of arguing about it.
The rules, written before anything ran
These went into the plan first, and they are the only reason the rest of this post is worth reading:
Two arms, one variable. Five clones against five diverse agents. Same model (Opus 5), same budget, same tools, same task, same morning. The only thing that changes is what each agent knows and how it is told to look.
Reality grades it. No model scoring another model's work, because a language model judging language models is the same well again. Deterministic Python, public APIs, zero human judgment in the loop.
Pre-registered. Scoring rule, pass threshold and the three possible endings written down before the first sample. A failed check is a finding, never a reason to edit the rule.
The control is built strong. The clone pack is a full theory of why things spread, applied from first principles. A rigged control proves nothing and I would know I rigged it.
Every experiment pays either way. A null result still has to leave me with a working tool.
Nulls get published. Including the ending where my thesis loses.
The task: forecast virality. Every morning at 07:10 the machinery samples 30 fresh posts from Hacker News, Reddit and X, minutes old and not yet hot. All ten agents give each post a probability of crossing a set popularity bar within 48 hours, plus their top five picks. Two days later a script checks what happened and scores everyone: Brier on every probability (squared error, so lower is better), precision on the top five. I picked it because the feedback loop is brutal and cheap. The answer arrives in 48 hours, verification costs one API call, and 30 predictions per agent per day means statistical power in two weeks rather than two years. The agents also cannot argue with the result, which is the part I like most.
The diverse five: a Hacker News native who knows ranking decay and flag risk, a Reddit native who prices silent mod removals before appeal, an X native who cares about follower graphs and link penalties, a trend historian who ignores craft and asks whether the topic has been running hot for 30 days, and one agent, cold-read, kept deliberately blind to anything recent so it can only judge the content in front of it.
Everything below runs in the open. Every chart and every number is on the live experiment page , all fourteen nights.
Finding one: prompt diversity is cosplay until you make it mechanical
Night one nearly fooled me, and it fooled me through the rationales. I read those before I read the correlations. They are good. The Hacker News lens talks about gravity and second-chance pools. The Reddit lens prices removal risk. The X lens shrugs at quality and asks who the author is. Five vocabularies, five kinds of evidence, five conclusions in prose.
Then all five wrote down 0.02.
Mean pairwise correlation inside the clone arm: 0.959. Inside the "diverse" arm: 0.909. On 26 of 30 posts the five different agents landed within a spread of 0.05, several of them on the exact same number. A metaphysics thread, a Rust tooling release and a Pygame tutorial got the same shrug from five agents that supposedly evaluate completely different things. I had built a panel of experts and got one opinion with five accents .
The only reason I know this is a line in the plan called the manipulation check: agent-to-agent correlation has to be lower in the diverse arm than in the clone arm, or the packs are a placebo and get rewritten before any result counts. Without that line, night one reads as a clean success. Ten of ten submissions valid, digest generated, daily loop running, every green light I had built for myself lit. I would have run two weeks, got a difference between the arms, and published it as evidence for diversity, when what I had measured was noise between ten near-copies.
I have been on the wrong side of this before. When I ran 255 agents to make three comics , fan-out scaled logistics and left quality where it was. When I told my agent to build an app every day , the output was competent and painfully boring for weeks before I admitted it. The first time I let four agents loose on a build , I judged the whole thing by reading what came out. All three times I noticed late. This time the trap had a tripwire in it and it went off in four hours. That is the entire argument for pre-registration, and it cost me one line in a plan I had already written.
There is research under this. A Cornell group published Correlated Errors in Large Language Models in June 2025, measured across more than 350 models. On one leaderboard dataset, two models agreed with each other 60% of the time when both were wrong, and larger, more accurate models had more correlated errors, across architectures and providers. If two frontier models from rival labs still fail together most of the time, five prompts wrapped around one model were never going to buy me independence. I was asking wording to do a job wording cannot do.
The easy repair is to tell the agents to disagree. Wider spread, zero evidence, because it manufactures exactly the result the experiment exists to test. There is a rule in the packs now that says so in writing. What I did instead was make each lens mechanical. Version one gave each agent a personality and a paragraph about what it cares about, then asked for a probability. Version two is a two-step move: the agent first reaches a verdict in its own categories and its own failure modes, then reads a probability off the range that verdict is assigned. The number comes out of the lens instead of out of the shared instinct with the lens narrating over the top.
Same metaphysics post, through the rebuilt lenses:
hn-native : was 0.03, now 0.00 to 0.02. Metaphysics bait, flag or kill.
reddit-native : was 0.03, now 0.14 to 0.22. Easy-comment title travels.
x-native : was 0.02, now 0.07 to 0.11. Away platform, payoff withheld.
trend-historian : was 0.02, now 0.07 to 0.17. Cold topic, honest width.
cold-read : was 0.02, now 0.01 to 0.03. Uncheckable central claim, capped.
Everyone within 0.02 of each other becomes a span from 0.00 to 0.22, and every step of it is principled. Nobody was told to fight. They stopped deferring to the same generic gut feeling.
The next morning the clone arm sat at 0.918 and the rebuilt diverse arm at 0.619. Herded posts fell from 26 of 30 to 21. And it held. The check passed on all fourteen nights . The correlation gap between the arms averaged 0.294, and it was wider in the second week than the first, 0.329 against 0.260, so the lenses kept separating instead of drifting back into the herd once the novelty wore off. Across the run the clones sit around 0.90 to 0.97 internal correlation, one brain in five copies, and the lenses run 0.53 to 0.75, with one strange night (24 August) where both arms spread out at once.
My opinion after fourteen nights: diversity has to be built, and a prompt cannot supply it. If a panel's members share a model and a context and differ only in the paragraph that tells them who they are, they are one forecaster with a thesaurus. Independence comes from different procedures, different information and, I now think, different providers. I split my own workflow across providers back in January and it looked like a preference at the time. It looks like the only real lever now.
Finding two: a number I guessed decided the scoreboard
Night one had a second fix, and at the time it felt like the most grown-up thing I did all evening. The shared hygiene text every agent gets said the base rate of a post going hot was "single digit percent", which is where that uniform floor of 0.01 to 0.03 came from. I replaced it with what I called the honest number: 10 to 15%, or 3 to 5 hot posts in a 30-post sample. I had accidentally coached the panel into the herd, I thought, and now I had un-coached it.
Fourteen nights, 416 scored posts per agent, and three of them crossed the bar. Three. The realized rate was 0.72% . I had spent an evening carefully coaching ten agents to expect fifteen times that, and they obliged: both arms forecast around 11% on average for the whole fortnight, because packs do what they are told. The "single digit percent" I overwrote sat closer to reality than my correction. And the block went into the clone pack too (it is shared, byte for byte, across all six packs), so on the same night I wrote "a control you keep improving stops being a control" into my plan, I made the control's calibration worse. Both arms were miscoached by the same amount, so the comparison between them stayed fair. What moved was the level, for everyone.
It got worse as the data came in. At four nights the realized rate was 1.72%. At seven, 0.97%. At fourteen, 0.72%. My coached figure stood still the whole time, so the gap widened every week, and the share of each arm's score explained by that one number went 60%, then 71%, then 74% for the clones, and 49%, 62%, 68% for the lenses. I assumed more data would average it out. More data made it bigger.
Here is why that swamps everything. Brier is squared error on a probability. Answer 0.11 on a post that stays cold and you pay 0.11 squared for that slot. Answer near the rate that occurred and you pay almost nothing. Now do that on nearly every one of 416 slots, because on this task nearly every post stays cold. The level of your answers becomes most of your score, and how well you read the posts is a small term sit

[truncated]
