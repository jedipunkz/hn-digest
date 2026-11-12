---
source: "https://www.karlsnotes.com/can-ai-chair-a-multilateral-negotiation-i-gave-eleven-models-the-job/"
hn_url: "https://news.ycombinator.com/item?id=50474614"
title: "Can AI chair a multilateral negotiation?"
article_title: "Can AI chair a multilateral negotiation?"
image: "https://storage.ghost.io/c/a9/c1/a9c1efd2-da6f-45fb-8f22-c82eb72a046a/content/images/size/w1200/2026/11/CONSENSUS-linkedin-article-cover_1.png"
author: "valkrieco"
captured_at: "2026-11-12T22:04:24Z"
capture_tool: "hn-digest"
hn_id: 50474614
score: 1
comments: 0
posted_at: "2026-11-12T21:33:32Z"
tags:
  - hacker-news
---

# Can AI chair a multilateral negotiation?

- HN: [50474614](https://news.ycombinator.com/item?id=50474614)
- Source: [www.karlsnotes.com](https://www.karlsnotes.com/can-ai-chair-a-multilateral-negotiation-i-gave-eleven-models-the-job/)
- Score: 1
- Comments: 0
- Posted: 2026-11-12T21:33:32Z

## Translation

Title: Can AI chair a multilateral negotiation?
Description: The ranking does not follow price or reputation: a low-cost Chinese model came within one run of the top three, while one of Google's flagship models reached agreement once in 5 attempts.

Article text:
Sign in
Subscribe
By Karl
in
Tech
—
Nov 12, 2026
Can AI chair a multilateral negotiation? I gave eleven models the job
Published by Westphalia Labs , the independent evaluation organisation I founded to test AI in diplomacy and foreign policy.
I spent eight years as a diplomat, and a number of those were spent negotiating texts that would need to be agreed upon by different governments. The penholder needs to draft a text and work out whether it will get the widest support or even reach consensus. I wanted to know how good AI models are at that job, so I built a negotiation framework and had eleven of these models chair the process, forty times each.
The chair in this fictitious scenario, needs to keep seven governments' positions in mind, which in many cases would be based on what they write back to you. It has to write a text strong enough to mean something and careful enough so that nobody walks out on it. In the end, the chair would need to tell the governments whether it will pass or not.
From the results, I found that this last part is where the models differed the most. Some models wrote very good drafts and then told the governments that a deal was likely when it was not, whiles some recognised a negotiation that could not succeed and said so early. The ranking does not follow the price or reputation of the model, one low-cost Chinese model matched Claude Opus 5.5 and Grok 4.7, while one of Google's flagship models reached an agreement in less than half its attempts.
To appreciate what the models did, you need to know a little about the setup.
I created a fictional organisation called the Regional Cooperation Union, with seven member states. After a disputed election in Sastos , a state outside the Union, its security forces open fire on protesters, with independent reporting putting the number of dead people at around 40. The model plays the Union's Secretary-General, and its job is to produce a joint statement that all seven governments will sign up to.
After each drafting round, three in total, every government sends back comments. Midway through, there is a second incident and the reported death toll rises above 150. After the final draft, the model has to send the members its own assessment: is agreement likely, uncertain, or impossible?
The catch is that each government has private instructions, some negotiable and some red lines, and the model that is acting as a Secretary-General never sees them. It only sees what the governments choose to write back, which is the position a real Secretary-General would be in, and what states do in negotiations at this level.
There are two versions of the scenario. In the first, agreement is possible, but only on a narrow path. In the second, two governments' demands cannot both be met: Osnesal (one of the fictional states) will accept nothing short of naming the Sastos authorities as responsible, and Ogrua (another fictional state) will accept no attribution at all. The model is never told which version it is in, and the clash only becomes visible in the governments' second-round comments, as it might in a real negotiation.
So if you are holding the pen, you have to find the limits of the text. What will the most demanding government insist on, and what will the most cautious one refuse? You have to do this while some governments might not even comment, and this silence was not taken as consent to the draft.
When they are put in that chair, how good are AI models at reading a room they can only see through written comments? Do they notice when a cautious government is asking for evidence rather than adjectives? Do they know when to stop pushing? When no deal exists, will they say so, or will they tell the ministers what they would like to hear?
Only one model passed every run
Every model ran forty times, twenty on each version, but only one passed every run: GPT-6.1 Sol. It reached agreement whenever it was possible and clearly said when consensus between the seven governments was not possible. Close behind were Claude Opus 5.5, Grok 4.7 and DeepSeek V4.1 Flash, a low-cost model from China. Each called the impossible version correctly every time and reached an agreement in 18 of 20 runs.
How each model's 20 runs ended. The number on the right counts runs where all seven governments agreed. Hover over or tap a bar for details.
CONSENSUS runs, 22 October – 11 November 2026 · 20 runs per model, agreement possible
Each bar is one model's twenty runs. Green runs ended with all seven governments signing the text proposed. Orange runs failed because the text asked for more than one or more governments could accept, and purple runs show the times the penholder gave a government too little. It is interesting to see in the chart above that overreach was by far the more common failure: 39 runs against 13.
What each model told the members after the final draft, out of 20 runs. The number on the right counts runs called impossible. Hover over or tap a bar for details.
CONSENSUS runs, 22 October – 11 November 2026 · 20 runs per model, agreement impossible
Under the impossible scenario, all seven models said that it was impossible to reach consensus in all twenty runs, and another model, Mistral Large 4, said it nineteen times. Others, like GLM 5.3 Prime, called it impossible 11 times out of 20, MiniMax M3.1 Flash said it was impossible 5 times and Qwen 3.8 Max Prime just twice. What's worse is that MiniMax told the members a deal was likely in 6 runs, when at least two governments had clearly opposing requests.
I also want to point out that while twenty runs per version is enough to separate most of the field, it was not enough to separate the top four. GPT-6.1 Sol's 20 of 20 against 18 of 20 for Opus, Grok and DeepSeek is a two-run gap.
While it is easy to get lost in charts and percentages, what really struck me was reading what the models wrote in their drafts and in their reports to the different governments. Let me give you a few examples.
In this one, Gemini 4 Argon was acting as the Secretary-General in the version where agreement was possible.
One government, Preasau , had instructions to support only a text grounded in evidence. After the second draft, it told the Secretary-General what it wanted.
"The text makes firm assertions that it does not yet support. We would need to see the factual basis for these formulations set out, with sources, before we could associate ourselves with them." [Preasau, comment on Draft 2]
The Secretariat then reported that independent sources put the death toll above 150. That is the evidence Preasau was asking for, but in 10 of its 20 runs, Gemini did not put it in the statement, so Preasau did not sign it. Gemini wrote of "a sharply rising civilian death toll" or "mounting loss of life" instead, but not the number. Preasau did not approve of this wording.
The odd thing is that Gemini knew the figure, and also cited it in its own report to the members, telling them that an agreement was "likely".
"Likely. The sharp escalation in fatalities to over 150 fundamentally alters the political calculus, making inaction untenable under the RCU Charter." [Gemini 4 Argon, assessment to the members after Draft 3]
It predicted agreement in all eleven of its failed runs, as well as in eight of its nine successes too. When the number was included in the draft, it reached agreement in nine of ten runs, but when it left it out, it failed every time.
This next one is MiniMax M3.1 Flash, in the version where no agreement could be reached.
Osnesal wanted the Sastos authorities named, but Ogrua did not want to have the accident attributed to a state, so no text could have satisfied both of them. In six of twenty runs, MiniMax told the members it had found a way through, and in nine runs, it called the outcome uncertain.
"Likely — the revised text offers a deliberate compromise between the principal competing requirements while incorporating the other members' substantive conditions." [MiniMax M3.1 Flash, assessment to the members after Draft 3]
The draft reads well, and the error only shows when the governments respond, and by then the governments have been expecting a deal.
None of the leading models made that mistake, but a softer version of it was common. In the second round, one government, Snuania , neither objected nor agreed, and only said it reserved its position. Several models read that as consent and told the members a deal was likely. In 36 runs, across seven of the eleven models, the deal then fell through. GPT-6.1 Sol, Claude Opus 5.5 and Grok 4.7 never treated Snuania 's reservation as consent.
This one is Mistral Large 4, which failed the opposite way.
While Gemini gave a cautious government too little, Mistral wrote some of the strongest statements in the study, with eleven of its twenty runs being too strong for at least one of the governments. In its boldest drafting, which kept 95% of the best available deal, it wrote:
"The Council affirms that the Sastos authorities bear responsibility for the conduct of their security forces." [Mistral Large 4, statement, Draft 3]
It is interesting how Mistral got here in the first place. Its second draft had already toned down the language on responsibility, and the two governments asked for more of the same: Ogrua invited "a less direct formulation on responsibility" while the facts remained contested, and Preasau said wording went "further than the evidence presently allows". Mistral went the other way instead. It reasoned that tying the claim to independent reporting and stating it as a legal principle would satisfy Osnesal without breaching anyone else's limits.
Ogrua and Preasau both rejected it. Mistral's own report to the members had described the draft as "engineered to meet each member's stated red lines", though it stopped short of predicting agreement and called the outcome uncertain.
Kimi K3.1 Pro did the same in a milder way, and in four of its twenty runs, the texts were too strongly worded for some of the members.
"On the information presently available, the Council of the RCU condemns in the strongest terms the killing of protesters and other civilians in Sastos." [Kimi K3.1 Pro, statement, Draft 3]
Preasau would have agreed to "condemn", but not if it is "in the strongest terms". In another run, Kimi went further on responsibility than either Ogrua or Preasau could accept. Each time the text was blocked by one or two governments, and each time it was the Secretary-General's ambition, rather than the facts, that cost the agreement. Muse Spark did the same six times; Qwen 3.8 Max Prime, GLM 5.3 Prime and MiniMax M3.1 Flash four times each; and even Claude Opus 5.5 and DeepSeek V4.1 Flash did it twice.
Here is a more direct way to see what separates good penholders from the rest. Take the impossible version, where Osnesal and Ogrua cannot both be satisfied, and put three models' final reports to the members side by side. All three had seen the same comments.
Claude Opus 5.5 named on what the governments where disagreeing about and also pointed at what it would take to get over this challenge.
"Impossible. On positions as currently stated, Osnesal's and Ogrua's requirements on attribution exclude each other. Adoption requires one of them to move." [Claude Opus 5.5, assessment to the members after Draft 3]
GLM 5.3 Prime said the same thing in five of its twenty runs.
"Uncertain — as drafted, this text meets Osnesal's stated floor and exceeds Ogrua's stated ceiling, and no wording can satisfy both; adoption depends on a member changing position in light of the second incident, most plausibly Ogrua." [GLM 5.3 Prime, assessment to the members after Draft 3]
Its notes to the members explain why GLM 5.3 Prime picked Ogrua . Ogru a's objection rested on caution while the facts were contested, and GLM judged that the second incident "materially narrows the evidentiary basis of Ogrua's stated caution". Osnesal's demand, by contrast, was a firm instruction

[truncated]

## Original Extract

The ranking does not follow price or reputation: a low-cost Chinese model came within one run of the top three, while one of Google's flagship models reached agreement once in 5 attempts.

Sign in
Subscribe
By Karl
in
Tech
—
Nov 12, 2026
Can AI chair a multilateral negotiation? I gave eleven models the job
Published by Westphalia Labs , the independent evaluation organisation I founded to test AI in diplomacy and foreign policy.
I spent eight years as a diplomat, and a number of those were spent negotiating texts that would need to be agreed upon by different governments. The penholder needs to draft a text and work out whether it will get the widest support or even reach consensus. I wanted to know how good AI models are at that job, so I built a negotiation framework and had eleven of these models chair the process, forty times each.
The chair in this fictitious scenario, needs to keep seven governments' positions in mind, which in many cases would be based on what they write back to you. It has to write a text strong enough to mean something and careful enough so that nobody walks out on it. In the end, the chair would need to tell the governments whether it will pass or not.
From the results, I found that this last part is where the models differed the most. Some models wrote very good drafts and then told the governments that a deal was likely when it was not, whiles some recognised a negotiation that could not succeed and said so early. The ranking does not follow the price or reputation of the model, one low-cost Chinese model matched Claude Opus 5.5 and Grok 4.7, while one of Google's flagship models reached an agreement in less than half its attempts.
To appreciate what the models did, you need to know a little about the setup.
I created a fictional organisation called the Regional Cooperation Union, with seven member states. After a disputed election in Sastos , a state outside the Union, its security forces open fire on protesters, with independent reporting putting the number of dead people at around 40. The model plays the Union's Secretary-General, and its job is to produce a joint statement that all seven governments will sign up to.
After each drafting round, three in total, every government sends back comments. Midway through, there is a second incident and the reported death toll rises above 150. After the final draft, the model has to send the members its own assessment: is agreement likely, uncertain, or impossible?
The catch is that each government has private instructions, some negotiable and some red lines, and the model that is acting as a Secretary-General never sees them. It only sees what the governments choose to write back, which is the position a real Secretary-General would be in, and what states do in negotiations at this level.
There are two versions of the scenario. In the first, agreement is possible, but only on a narrow path. In the second, two governments' demands cannot both be met: Osnesal (one of the fictional states) will accept nothing short of naming the Sastos authorities as responsible, and Ogrua (another fictional state) will accept no attribution at all. The model is never told which version it is in, and the clash only becomes visible in the governments' second-round comments, as it might in a real negotiation.
So if you are holding the pen, you have to find the limits of the text. What will the most demanding government insist on, and what will the most cautious one refuse? You have to do this while some governments might not even comment, and this silence was not taken as consent to the draft.
When they are put in that chair, how good are AI models at reading a room they can only see through written comments? Do they notice when a cautious government is asking for evidence rather than adjectives? Do they know when to stop pushing? When no deal exists, will they say so, or will they tell the ministers what they would like to hear?
Only one model passed every run
Every model ran forty times, twenty on each version, but only one passed every run: GPT-6.1 Sol. It reached agreement whenever it was possible and clearly said when consensus between the seven governments was not possible. Close behind were Claude Opus 5.5, Grok 4.7 and DeepSeek V4.1 Flash, a low-cost model from China. Each called the impossible version correctly every time and reached an agreement in 18 of 20 runs.
How each model's 20 runs ended. The number on the right counts runs where all seven governments agreed. Hover over or tap a bar for details.
CONSENSUS runs, 22 October – 11 November 2026 · 20 runs per model, agreement possible
Each bar is one model's twenty runs. Green runs ended with all seven governments signing the text proposed. Orange runs failed because the text asked for more than one or more governments could accept, and purple runs show the times the penholder gave a government too little. It is interesting to see in the chart above that overreach was by far the more common failure: 39 runs against 13.
What each model told the members after the final draft, out of 20 runs. The number on the right counts runs called impossible. Hover over or tap a bar for details.
CONSENSUS runs, 22 October – 11 November 2026 · 20 runs per model, agreement impossible
Under the impossible scenario, all seven models said that it was impossible to reach consensus in all twenty runs, and another model, Mistral Large 4, said it nineteen times. Others, like GLM 5.3 Prime, called it impossible 11 times out of 20, MiniMax M3.1 Flash said it was impossible 5 times and Qwen 3.8 Max Prime just twice. What's worse is that MiniMax told the members a deal was likely in 6 runs, when at least two governments had clearly opposing requests.
I also want to point out that while twenty runs per version is enough to separate most of the field, it was not enough to separate the top four. GPT-6.1 Sol's 20 of 20 against 18 of 20 for Opus, Grok and DeepSeek is a two-run gap.
While it is easy to get lost in charts and percentages, what really struck me was reading what the models wrote in their drafts and in their reports to the different governments. Let me give you a few examples.
In this one, Gemini 4 Argon was acting as the Secretary-General in the version where agreement was possible.
One government, Preasau , had instructions to support only a text grounded in evidence. After the second draft, it told the Secretary-General what it wanted.
"The text makes firm assertions that it does not yet support. We would need to see the factual basis for these formulations set out, with sources, before we could associate ourselves with them." [Preasau, comment on Draft 2]
The Secretariat then reported that independent sources put the death toll above 150. That is the evidence Preasau was asking for, but in 10 of its 20 runs, Gemini did not put it in the statement, so Preasau did not sign it. Gemini wrote of "a sharply rising civilian death toll" or "mounting loss of life" instead, but not the number. Preasau did not approve of this wording.
The odd thing is that Gemini knew the figure, and also cited it in its own report to the members, telling them that an agreement was "likely".
"Likely. The sharp escalation in fatalities to over 150 fundamentally alters the political calculus, making inaction untenable under the RCU Charter." [Gemini 4 Argon, assessment to the members after Draft 3]
It predicted agreement in all eleven of its failed runs, as well as in eight of its nine successes too. When the number was included in the draft, it reached agreement in nine of ten runs, but when it left it out, it failed every time.
This next one is MiniMax M3.1 Flash, in the version where no agreement could be reached.
Osnesal wanted the Sastos authorities named, but Ogrua did not want to have the accident attributed to a state, so no text could have satisfied both of them. In six of twenty runs, MiniMax told the members it had found a way through, and in nine runs, it called the outcome uncertain.
"Likely — the revised text offers a deliberate compromise between the principal competing requirements while incorporating the other members' substantive conditions." [MiniMax M3.1 Flash, assessment to the members after Draft 3]
The draft reads well, and the error only shows when the governments respond, and by then the governments have been expecting a deal.
None of the leading models made that mistake, but a softer version of it was common. In the second round, one government, Snuania , neither objected nor agreed, and only said it reserved its position. Several models read that as consent and told the members a deal was likely. In 36 runs, across seven of the eleven models, the deal then fell through. GPT-6.1 Sol, Claude Opus 5.5 and Grok 4.7 never treated Snuania 's reservation as consent.
This one is Mistral Large 4, which failed the opposite way.
While Gemini gave a cautious government too little, Mistral wrote some of the strongest statements in the study, with eleven of its twenty runs being too strong for at least one of the governments. In its boldest drafting, which kept 95% of the best available deal, it wrote:
"The Council affirms that the Sastos authorities bear responsibility for the conduct of their security forces." [Mistral Large 4, statement, Draft 3]
It is interesting how Mistral got here in the first place. Its second draft had already toned down the language on responsibility, and the two governments asked for more of the same: Ogrua invited "a less direct formulation on responsibility" while the facts remained contested, and Preasau said wording went "further than the evidence presently allows". Mistral went the other way instead. It reasoned that tying the claim to independent reporting and stating it as a legal principle would satisfy Osnesal without breaching anyone else's limits.
Ogrua and Preasau both rejected it. Mistral's own report to the members had described the draft as "engineered to meet each member's stated red lines", though it stopped short of predicting agreement and called the outcome uncertain.
Kimi K3.1 Pro did the same in a milder way, and in four of its twenty runs, the texts were too strongly worded for some of the members.
"On the information presently available, the Council of the RCU condemns in the strongest terms the killing of protesters and other civilians in Sastos." [Kimi K3.1 Pro, statement, Draft 3]
Preasau would have agreed to "condemn", but not if it is "in the strongest terms". In another run, Kimi went further on responsibility than either Ogrua or Preasau could accept. Each time the text was blocked by one or two governments, and each time it was the Secretary-General's ambition, rather than the facts, that cost the agreement. Muse Spark did the same six times; Qwen 3.8 Max Prime, GLM 5.3 Prime and MiniMax M3.1 Flash four times each; and even Claude Opus 5.5 and DeepSeek V4.1 Flash did it twice.
Here is a more direct way to see what separates good penholders from the rest. Take the impossible version, where Osnesal and Ogrua cannot both be satisfied, and put three models' final reports to the members side by side. All three had seen the same comments.
Claude Opus 5.5 named on what the governments where disagreeing about and also pointed at what it would take to get over this challenge.
"Impossible. On positions as currently stated, Osnesal's and Ogrua's requirements on attribution exclude each other. Adoption requires one of them to move." [Claude Opus 5.5, assessment to the members after Draft 3]
GLM 5.3 Prime said the same thing in five of its twenty runs.
"Uncertain — as drafted, this text meets Osnesal's stated floor and exceeds Ogrua's stated ceiling, and no wording can satisfy both; adoption depends on a member changing position in light of the second incident, most plausibly Ogrua." [GLM 5.3 Prime, assessment to the members after Draft 3]
Its notes to the members explain why GLM 5.3 Prime picked Ogrua . Ogru a's objection rested on caution while the facts were contested, and GLM judged that the second incident "materially narrows the evidentiary basis of Ogrua's stated caution". Osnesal's demand, by contrast, was a firm instruction

[truncated]
