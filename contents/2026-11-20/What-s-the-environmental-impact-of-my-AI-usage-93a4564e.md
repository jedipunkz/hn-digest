---
source: "https://substack.hadley.nz/p/whats-the-environmental-impact-of"
hn_url: "https://news.ycombinator.com/item?id=50575419"
title: "What's the environmental impact of my AI usage?"
article_title: "What’s the environmental impact of my AI usage?"
image: "https://substackcdn.com/image/fetch/$s_!SGWk!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F7485c002-a64e-4b6d-a526-816a6fe55fd7_1200x600.png"
author: "doener"
captured_at: "2026-11-20T15:09:59Z"
capture_tool: "hn-digest"
hn_id: 50575419
score: 1
comments: 1
posted_at: "2026-11-20T14:49:35Z"
tags:
  - hacker-news
---

# What's the environmental impact of my AI usage?

- HN: [50575419](https://news.ycombinator.com/item?id=50575419)
- Source: [substack.hadley.nz](https://substack.hadley.nz/p/whats-the-environmental-impact-of)
- Score: 1
- Comments: 1
- Posted: 2026-11-20T14:49:35Z

## Translation

Title: What's the environmental impact of my AI usage?
Article title: What’s the environmental impact of my AI usage?
Description: I spent nearly $6,000 on LLM tokens in the last six months — so how guilty should I feel about the planet?

Article text:
What’s the environmental impact of my AI usage?
AI for data science
Subscribe Sign in What’s the environmental impact of my AI usage?
Hadley Wickham Nov 20, 2026 15 2 1 Share In my last post I tried to estimate what tokens really cost with my favourite open model, Kimi K3. That analysis has a nice side effect: it also tells us the amount of power we draw per token, allowing me to estimate the environmental impact of my personal AI usage. Similar to last week, I’m estimating the marginal environmental impact of an additional token, ignoring the cost of training the model and manufacturing the GPUs.
I’ll first figure out the environmental impact per million tokens (Mtok). To apply these numbers to my AI usage generally, I make the simplifying assumption that impact scales with cost to estimate impact per dollar. Then I’ll combine that with my personal usage data to estimate my AI impact, and look at how it compares to some of my other major environmental impacts.
We’ll first estimate power consumption, then use that to estimate water usage and CO2 generation.
Like last week, we’re going to run Kimi K3 on a Supermicro 8U GPU SuperServer, which has 8 GB300 GPUs. Each GB300 has a “thermal design power” of ~1.4 kW: this is the amount of sustained heat a chip puts out, and a reasonable proxy for peak power draw because most of the power used turns into heat. So 8 GPUs yield ~11.2 kW, then CPUs, networking, and storage add another ~1.5 kW, giving a total load of ~12.7 kW.
Now we need to adjust this by the “power usage effectiveness” (PUE) of the colocation facility. This is the ratio between the total power used by the server facility, and the amount used by the individual servers, and is predominantly determined by cooling efficiency. This varies a lot based on the data center, but luckily AWS , Google and Microsoft publish their data, with PUEs ranging from 1.09 to 1.17. I’m going to be conservative and assume that the new LLM providers are maybe not quite this good, and have a PUE of 1.25. This yields ~15.9 kW facility draw at 100% usage.
But GPUs still draw ~30% of power at idle, so running below full utilisation uses more energy per token: +10% at 75% load, +30% at 50%, +90% at 25%. Like last week, I’ll assume the 75% case as my default, giving a total draw of 17.5 kW.
I’ll use the same throughput numbers as last week, from InferenceX : ~6,000 tokens per second per GPU, i.e. ~48,000 tokens/s ≈ 124B tokens/month. Putting it together, at 75% load, the server processes 93B tokens while drawing 17.5 kW. 17.5 kW × 24 h × 30 days = 12,600 kWh for 93B tokens, which is ~0.14 kWh/Mtok.
Now that we have kWh/Mtok, we can figure out the water usage, based on the “water usage effectiveness” (WUE). That’s liters of water used for cooling divided by server power usage in kWh. Note the denominator: it’s server power usage, not facility usage, so we’ll have to divide back by PUE in our calculations.
We only have WUE data from AWS and Microsoft (same sources as above), with values of 0.12 and 0.27 respectively. Again, we’ll take the top end and add a little padding, and assume a WUE of 0.33 L/kWh. So that’s 0.14 kWh/Mtok ÷ 1.25 PUE × 0.33 L/kWh WUE = ~0.037 L/Mtok of on-site water.
We also need to include the water used by power generation, which this study from 2018 estimates at an average of 2.2 L/kWh. Those numbers are probably overestimates since solar and wind power don’t require water for cooling, and are a bigger share of the mix today, which fits our general theme of conservative estimates.
Power generation has to cover the full facility draw, so that’s 0.14 kWh/Mtok × 2.2 L/kWh = ~0.31 L/Mtok of additional water usage, bringing the total water usage to ~0.35 L/Mtok.
Carbon generation is simple given power consumption. We can use eGRID data from 2023 which gives a US average of 0.348 kg CO2/kWh. So that’s 0.14 kWh/Mtok × 0.348 kg CO2/kWh = 0.049 kg CO2/Mtok, which we’ll round up to ~0.05 kg CO2/Mtok .
Now I need to make a big assumption to apply what we know about Kimi K3 to other models: the environmental impact is proportional to cost. My justification is that electricity is the primary marginal cost of serving a token (reasonable) and that models are priced rationally (probably not very reasonable, but the best we can do). So my assumption is that that a model that costs 10× more per token likely consumes roughly 10× more energy per token. This lets me estimate the impact of closed models (Fable, Opus, GPT) that I don’t otherwise have data for.
The blended price I pay for Kimi K3 is $0.69/Mtok (input + cached + output tokens, at my ~90% cache rate) which gives per-dollar impacts of: 0.14 kWh/Mtok → ~0.20 kWh/$, 0.35 L/Mtok → ~0.51 L/$, and 0.05 kg CO2/Mtok → ~0.07 kg CO2/$.
Now let’s connect this to my usage. I used AgentsView to collect all of my AI coding usage from May to November, ~6.3 months. This covers coding agents, not web chat or ellmer, but I’m confident that this is the vast majority of my usage. My total spend was $5,871 , mostly across 8 models:
Model Input Cached Output Cost
(MTok) (MTok) (MTok)
---------------- ------ ------- ------- ------
Fable 5 0.3 1,440 7.3 $2,510
Opus 5 0.05 1,573 5.8 $1,221
Opus 4.8 0.7 754 4.5 $687
Kimi K3 75.3 852 7.0 $645
Opus 5.5 0.01 368 3.4 $245
GPT 5.6 15.3 261 1.5 $237
Sonnet 5 14.9 390 2.9 $190
GPT 5.5 2.7 100 0.3 $74
Everything else 4.0 113 1.7 $63 You I can estimate the environmental impact of my usage, and what it would scale up to per year:
Energy Water Carbon
(kWh) (L) (kg CO2)
---------- -------- -------- --------
Per Mtok 0.14 0.35 0.05
Per dollar ~0.20 ~0.51 ~0.07
My usage ~1,200 ~3,000 ~410
Annualised ~2,200 ~5,700 ~780 Let’s contextualise that with some other of my environmental impacts:
Energy: our house uses ~2300 kWh/month in the summer and ~1300 kWh/month winter. There’s such a big difference because Houston is hot, and despite having a very well insulated house and minimal cooling (by US standards), it’s still expensive. So that’s a total of ~21,500 kWh/year.
Water: our irrigation system uses 400 gallons × 3 days/week ≈ 4,500 L/week ≈ 240,000 L/year. This is what our landscapers suggested, and I did not really think about, but looking at it now feels like something I should change.
Carbon: we use 100% solar power and have electric vehicles, so I think my primary carbon impact is from flying (which I do a lot). In the last 12 months, I flew on 45 flights, for a total of ~137,000 km and ~203 hours in the air, an estimated ~25,400 kg CO2.
Water Energy Carbon
(L) (kWh) (kg CO2)
-------------- -------- ------- --------
My lifestyle 240,000 21,500 25,400
Coding agents 5,700 2,200 780
% 2% 10% 3% So overall, the carbon and water impacts of my AI usage are marginal, and the energy impact is small. But that’s my own usage, where (at least in the last few months) I carefully monitor costs with Posit Assistant, and I’m very aware of what I’m spending.
What if you used subscriptions instead? Last post I noted that a $200/month Claude Max plan gets you somewhere between $2,000 and $10,000 worth of tokens. Assuming I leaned into that heavily and used $5,000 worth of tokens every month, that’s ~12,000 kWh of energy per year, ~31,000 L of water, and ~4,200 kg of CO2. So the biggest environmental impacts come from subscriptions, which hide the underlying costs, and this is especially true for folks who use multiple subscriptions.
So overall, I don’t feel too worried about the environmental impact of my use of coding agents. But I do feel bad about how much water I’m using for irrigation, and I’ll look into reducing that. I also continue to feel bad about the carbon impact of my flights, but I’m not sure what to do about that, except continuing to donate to climate charities .
My monthly spend has also dropped considerably, as I have become more aware of how much different models cost. About half my total spend was on Fable, and I now use models of this calibre much more sparingly. I think you can get a lot done with $200-$500 of AI / month, and I think for most people (or at least most Americans), that’s not going to significantly increase your environmental impact.
What do you all think? Does my analysis seem plausible? Are there ways I could improve it? Am I too generous towards my own environment impact?
15 2 1 Share Previous Discussion about this post Comments Restacks John McIntire 43m Interesting analysis. The main lesson is obvious ...
On your water use, Houston has ... 1500 mm annual average rainfall ... maybe you could do some primitive water harvesting (stick a barrel under your gutters) and/or switch to a desert / zero lawn garden ... ? My Chinese associates tell me there are lectures on "Zeriscaping basics" in the Houston area.
Like Reply Share 1 reply 1 more comment... Top Latest Discussions Ready for more?

## Original Extract

I spent nearly $6,000 on LLM tokens in the last six months — so how guilty should I feel about the planet?

What’s the environmental impact of my AI usage?
AI for data science
Subscribe Sign in What’s the environmental impact of my AI usage?
Hadley Wickham Nov 20, 2026 15 2 1 Share In my last post I tried to estimate what tokens really cost with my favourite open model, Kimi K3. That analysis has a nice side effect: it also tells us the amount of power we draw per token, allowing me to estimate the environmental impact of my personal AI usage. Similar to last week, I’m estimating the marginal environmental impact of an additional token, ignoring the cost of training the model and manufacturing the GPUs.
I’ll first figure out the environmental impact per million tokens (Mtok). To apply these numbers to my AI usage generally, I make the simplifying assumption that impact scales with cost to estimate impact per dollar. Then I’ll combine that with my personal usage data to estimate my AI impact, and look at how it compares to some of my other major environmental impacts.
We’ll first estimate power consumption, then use that to estimate water usage and CO2 generation.
Like last week, we’re going to run Kimi K3 on a Supermicro 8U GPU SuperServer, which has 8 GB300 GPUs. Each GB300 has a “thermal design power” of ~1.4 kW: this is the amount of sustained heat a chip puts out, and a reasonable proxy for peak power draw because most of the power used turns into heat. So 8 GPUs yield ~11.2 kW, then CPUs, networking, and storage add another ~1.5 kW, giving a total load of ~12.7 kW.
Now we need to adjust this by the “power usage effectiveness” (PUE) of the colocation facility. This is the ratio between the total power used by the server facility, and the amount used by the individual servers, and is predominantly determined by cooling efficiency. This varies a lot based on the data center, but luckily AWS , Google and Microsoft publish their data, with PUEs ranging from 1.09 to 1.17. I’m going to be conservative and assume that the new LLM providers are maybe not quite this good, and have a PUE of 1.25. This yields ~15.9 kW facility draw at 100% usage.
But GPUs still draw ~30% of power at idle, so running below full utilisation uses more energy per token: +10% at 75% load, +30% at 50%, +90% at 25%. Like last week, I’ll assume the 75% case as my default, giving a total draw of 17.5 kW.
I’ll use the same throughput numbers as last week, from InferenceX : ~6,000 tokens per second per GPU, i.e. ~48,000 tokens/s ≈ 124B tokens/month. Putting it together, at 75% load, the server processes 93B tokens while drawing 17.5 kW. 17.5 kW × 24 h × 30 days = 12,600 kWh for 93B tokens, which is ~0.14 kWh/Mtok.
Now that we have kWh/Mtok, we can figure out the water usage, based on the “water usage effectiveness” (WUE). That’s liters of water used for cooling divided by server power usage in kWh. Note the denominator: it’s server power usage, not facility usage, so we’ll have to divide back by PUE in our calculations.
We only have WUE data from AWS and Microsoft (same sources as above), with values of 0.12 and 0.27 respectively. Again, we’ll take the top end and add a little padding, and assume a WUE of 0.33 L/kWh. So that’s 0.14 kWh/Mtok ÷ 1.25 PUE × 0.33 L/kWh WUE = ~0.037 L/Mtok of on-site water.
We also need to include the water used by power generation, which this study from 2018 estimates at an average of 2.2 L/kWh. Those numbers are probably overestimates since solar and wind power don’t require water for cooling, and are a bigger share of the mix today, which fits our general theme of conservative estimates.
Power generation has to cover the full facility draw, so that’s 0.14 kWh/Mtok × 2.2 L/kWh = ~0.31 L/Mtok of additional water usage, bringing the total water usage to ~0.35 L/Mtok.
Carbon generation is simple given power consumption. We can use eGRID data from 2023 which gives a US average of 0.348 kg CO2/kWh. So that’s 0.14 kWh/Mtok × 0.348 kg CO2/kWh = 0.049 kg CO2/Mtok, which we’ll round up to ~0.05 kg CO2/Mtok .
Now I need to make a big assumption to apply what we know about Kimi K3 to other models: the environmental impact is proportional to cost. My justification is that electricity is the primary marginal cost of serving a token (reasonable) and that models are priced rationally (probably not very reasonable, but the best we can do). So my assumption is that that a model that costs 10× more per token likely consumes roughly 10× more energy per token. This lets me estimate the impact of closed models (Fable, Opus, GPT) that I don’t otherwise have data for.
The blended price I pay for Kimi K3 is $0.69/Mtok (input + cached + output tokens, at my ~90% cache rate) which gives per-dollar impacts of: 0.14 kWh/Mtok → ~0.20 kWh/$, 0.35 L/Mtok → ~0.51 L/$, and 0.05 kg CO2/Mtok → ~0.07 kg CO2/$.
Now let’s connect this to my usage. I used AgentsView to collect all of my AI coding usage from May to November, ~6.3 months. This covers coding agents, not web chat or ellmer, but I’m confident that this is the vast majority of my usage. My total spend was $5,871 , mostly across 8 models:
Model Input Cached Output Cost
(MTok) (MTok) (MTok)
---------------- ------ ------- ------- ------
Fable 5 0.3 1,440 7.3 $2,510
Opus 5 0.05 1,573 5.8 $1,221
Opus 4.8 0.7 754 4.5 $687
Kimi K3 75.3 852 7.0 $645
Opus 5.5 0.01 368 3.4 $245
GPT 5.6 15.3 261 1.5 $237
Sonnet 5 14.9 390 2.9 $190
GPT 5.5 2.7 100 0.3 $74
Everything else 4.0 113 1.7 $63 You I can estimate the environmental impact of my usage, and what it would scale up to per year:
Energy Water Carbon
(kWh) (L) (kg CO2)
---------- -------- -------- --------
Per Mtok 0.14 0.35 0.05
Per dollar ~0.20 ~0.51 ~0.07
My usage ~1,200 ~3,000 ~410
Annualised ~2,200 ~5,700 ~780 Let’s contextualise that with some other of my environmental impacts:
Energy: our house uses ~2300 kWh/month in the summer and ~1300 kWh/month winter. There’s such a big difference because Houston is hot, and despite having a very well insulated house and minimal cooling (by US standards), it’s still expensive. So that’s a total of ~21,500 kWh/year.
Water: our irrigation system uses 400 gallons × 3 days/week ≈ 4,500 L/week ≈ 240,000 L/year. This is what our landscapers suggested, and I did not really think about, but looking at it now feels like something I should change.
Carbon: we use 100% solar power and have electric vehicles, so I think my primary carbon impact is from flying (which I do a lot). In the last 12 months, I flew on 45 flights, for a total of ~137,000 km and ~203 hours in the air, an estimated ~25,400 kg CO2.
Water Energy Carbon
(L) (kWh) (kg CO2)
-------------- -------- ------- --------
My lifestyle 240,000 21,500 25,400
Coding agents 5,700 2,200 780
% 2% 10% 3% So overall, the carbon and water impacts of my AI usage are marginal, and the energy impact is small. But that’s my own usage, where (at least in the last few months) I carefully monitor costs with Posit Assistant, and I’m very aware of what I’m spending.
What if you used subscriptions instead? Last post I noted that a $200/month Claude Max plan gets you somewhere between $2,000 and $10,000 worth of tokens. Assuming I leaned into that heavily and used $5,000 worth of tokens every month, that’s ~12,000 kWh of energy per year, ~31,000 L of water, and ~4,200 kg of CO2. So the biggest environmental impacts come from subscriptions, which hide the underlying costs, and this is especially true for folks who use multiple subscriptions.
So overall, I don’t feel too worried about the environmental impact of my use of coding agents. But I do feel bad about how much water I’m using for irrigation, and I’ll look into reducing that. I also continue to feel bad about the carbon impact of my flights, but I’m not sure what to do about that, except continuing to donate to climate charities .
My monthly spend has also dropped considerably, as I have become more aware of how much different models cost. About half my total spend was on Fable, and I now use models of this calibre much more sparingly. I think you can get a lot done with $200-$500 of AI / month, and I think for most people (or at least most Americans), that’s not going to significantly increase your environmental impact.
What do you all think? Does my analysis seem plausible? Are there ways I could improve it? Am I too generous towards my own environment impact?
15 2 1 Share Previous Discussion about this post Comments Restacks John McIntire 43m Interesting analysis. The main lesson is obvious ...
On your water use, Houston has ... 1500 mm annual average rainfall ... maybe you could do some primitive water harvesting (stick a barrel under your gutters) and/or switch to a desert / zero lawn garden ... ? My Chinese associates tell me there are lectures on "Zeriscaping basics" in the Houston area.
Like Reply Share 1 reply 1 more comment... Top Latest Discussions Ready for more?
