---
source: "https://labqoat.com/blog/how-well-can-ai-identify-spiders"
hn_url: "https://news.ycombinator.com/item?id=49784855"
title: "I'm afraid of spiders. So I made AI look at 2k of them"
article_title: "I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them. · Labqoat"
image: "https://labqoat.com/spider-bench/og.jpg"
author: "michalwarda"
captured_at: "2026-09-21T09:40:24Z"
capture_tool: "hn-digest"
hn_id: 49784855
score: 4
comments: 0
posted_at: "2026-09-21T09:00:37Z"
tags:
  - hacker-news
---

# I'm afraid of spiders. So I made AI look at 2k of them

- HN: [49784855](https://news.ycombinator.com/item?id=49784855)
- Source: [labqoat.com](https://labqoat.com/blog/how-well-can-ai-identify-spiders)
- Score: 4
- Comments: 0
- Posted: 2026-09-21T09:00:37Z

## Translation

Title: I'm afraid of spiders. So I made AI look at 2k of them
Article title: I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them. · Labqoat
Description: I’m afraid of spiders, so I made AI look at 2,000 of them. What it got right, what it got wrong, and how much it cost.

Article text:
I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them. · Labqoat Skip to content labqoat Platform 01 Compute 02 Research 03 Blog Our philosophy Explore the lab Menu + ← All posts September 11, 2026 Benchmarks 5 min read
I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them.
First, approximately ten seconds of biology
Is getting half wrong actually bad?
I’m afraid of spiders. A lot. 🕷️
My identification system currently consists of “the one with long legs” and “the short but big one.” There’s also “where the fuck did it go,” but that’s more of an emergency than a classification.
So I got curious: how well could AI actually identify them?
I tested nine models on the same 2,000 spider photos to see how often they got the species right.
(Looking at this many spiders was not a comfortable experience.)
This post was heavily inspired by Piotr Migdał’s post on identifying mushrooms with AI .
First, approximately ten seconds of biology
A family is a broad group of related spiders. A genus is a smaller group inside it, and a species is the specific kind of spider.
For example: Araneidae → Araneus → Araneus diadematus , the European garden spider . In that last name, Araneus is the genus.
That’s enough biology for now.
Each model got the same 2,000 photos , covering 671 species and subspecies , and a list of 20 possible names for each photo, including the correct answer.
The task was to identify the spider’s species from the photo and return the matching name from the list.
I gave each model the same possible answers so I could compare how well they distinguished the species, with fewer ambiguities in scoring. That makes this a multiple-choice identification test , and the choices themselves can help.
I built the dataset from research-grade iNaturalist observations, using the community’s species identifications as the expected answers. The species list comes from a Polish spider checklist, but the photos were taken worldwide. After filtering out label mismatches and unsuitable images, 2,183 photos remained, and I used the same 2,000-photo subset for every model.
I’d call any of these “spider,” but… 😅
According to iNaturalist , it’s a Red-bellied Jumping Spider ( Philaeus chrysops ).
9 of 9 picked the expected species.
13 . Leptorchestes berolinensis
* Muse Spark 1.3 uses the Contributor tier.
Selected examples; overall scores use all 2,000 photos.
Gemini 3.8 Flash got 997 out of 2,000 correct: 49.85% , the highest score in these runs.
GPT-6 Astra followed at 47.50% , then Claude Fable 5.1 at 43.10% and Muse Spark 1.3 at 39.10% . The top two were separated by 47 photos. I’d read that as a result on this set of spiders, not a universal ranking of how much the models know about them.
Same 2,000 photos for every model. Wrong and failed answers stay in the denominator.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Code and benchmark results on GitHub
Gemini 3.8 Flash missed the exact species in 1,003 photos. But in 871 of those cases , it named another spider from the same family. That’s almost 87% of its misses . In 276 cases, it even got the genus right.
GPT-6 Astra and Claude Fable 5.1 showed a similar pattern. Their exact-species scores were 47.50% and 43.10% , but their answers belonged to the correct family 92.00% and 91.45% of the time. Most of their mistakes happened between species within the same family.
Muse Spark 1.3 identified more exact species than GLM 5.3 Flash. GLM’s answers, however, belonged to the correct family slightly more often: 89.40% versus 88.70% .
That makes the errors more interesting than a simple correct-or-wrong score suggests. There’s a substantial difference between getting the broad group right and identifying the particular species.
Each answer appears in one segment. “Only” means the more specific identification was wrong. Failed responses include invalid names, refusals, empty answers, truncations, and request errors.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Counts, not cumulative percentages. Each row adds up to 2,000.
Code and benchmark results on GitHub
Muse Spark 1.3 was the cheapest model here and still finished fourth, ahead of five more expensive models. At its Contributor pricing, the entire run cost $0.93 for 782 correct identifications .
Moving to Gemini 3.8 Flash brought that total to 997. That’s 215 more correct answers for another $26.18 , an improvement of 10.75 percentage points at roughly 29 times the cost . 💸
Spending beyond that didn’t improve the results. Claude Fable 5.1 cost about twice as much as Gemini, and GPT-6 Astra cost 3.6 times as much . Both scored lower. Gemini delivered the highest accuracy here; Muse offered a cheaper trade-off if identifying fewer species was acceptable.
Estimated totals for 2,000 photos. The dollar axis is logarithmic: spacing represents cost ratios.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Code and benchmark results on GitHub
Is getting half wrong actually bad?
Getting almost half right is respectable when the task is telling similar species apart. The gap between recognising a family and naming the exact species is where this gets difficult.
Try telling these two spiders apart. For Araniella opisthographa , NatureSpot’s identification guidance calls for examination at high magnification. That’s the sort of detail a photo can leave out.
The code and full benchmark results are on GitHub .
PS: If you actually know spiders, I’d love to hear what you think on twitter . How obvious are these mistakes to someone who knows what they’re looking at?
P.P.S. Thankfully, I have a cat in charge of spider security at home. Call security
The AI lab for agent swarms.
Build better models, together.

## Original Extract

I’m afraid of spiders, so I made AI look at 2,000 of them. What it got right, what it got wrong, and how much it cost.

I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them. · Labqoat Skip to content labqoat Platform 01 Compute 02 Research 03 Blog Our philosophy Explore the lab Menu + ← All posts September 11, 2026 Benchmarks 5 min read
I’m Afraid of Spiders. So I Made AI Look at 2,000 of Them.
First, approximately ten seconds of biology
Is getting half wrong actually bad?
I’m afraid of spiders. A lot. 🕷️
My identification system currently consists of “the one with long legs” and “the short but big one.” There’s also “where the fuck did it go,” but that’s more of an emergency than a classification.
So I got curious: how well could AI actually identify them?
I tested nine models on the same 2,000 spider photos to see how often they got the species right.
(Looking at this many spiders was not a comfortable experience.)
This post was heavily inspired by Piotr Migdał’s post on identifying mushrooms with AI .
First, approximately ten seconds of biology
A family is a broad group of related spiders. A genus is a smaller group inside it, and a species is the specific kind of spider.
For example: Araneidae → Araneus → Araneus diadematus , the European garden spider . In that last name, Araneus is the genus.
That’s enough biology for now.
Each model got the same 2,000 photos , covering 671 species and subspecies , and a list of 20 possible names for each photo, including the correct answer.
The task was to identify the spider’s species from the photo and return the matching name from the list.
I gave each model the same possible answers so I could compare how well they distinguished the species, with fewer ambiguities in scoring. That makes this a multiple-choice identification test , and the choices themselves can help.
I built the dataset from research-grade iNaturalist observations, using the community’s species identifications as the expected answers. The species list comes from a Polish spider checklist, but the photos were taken worldwide. After filtering out label mismatches and unsuitable images, 2,183 photos remained, and I used the same 2,000-photo subset for every model.
I’d call any of these “spider,” but… 😅
According to iNaturalist , it’s a Red-bellied Jumping Spider ( Philaeus chrysops ).
9 of 9 picked the expected species.
13 . Leptorchestes berolinensis
* Muse Spark 1.3 uses the Contributor tier.
Selected examples; overall scores use all 2,000 photos.
Gemini 3.8 Flash got 997 out of 2,000 correct: 49.85% , the highest score in these runs.
GPT-6 Astra followed at 47.50% , then Claude Fable 5.1 at 43.10% and Muse Spark 1.3 at 39.10% . The top two were separated by 47 photos. I’d read that as a result on this set of spiders, not a universal ranking of how much the models know about them.
Same 2,000 photos for every model. Wrong and failed answers stay in the denominator.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Code and benchmark results on GitHub
Gemini 3.8 Flash missed the exact species in 1,003 photos. But in 871 of those cases , it named another spider from the same family. That’s almost 87% of its misses . In 276 cases, it even got the genus right.
GPT-6 Astra and Claude Fable 5.1 showed a similar pattern. Their exact-species scores were 47.50% and 43.10% , but their answers belonged to the correct family 92.00% and 91.45% of the time. Most of their mistakes happened between species within the same family.
Muse Spark 1.3 identified more exact species than GLM 5.3 Flash. GLM’s answers, however, belonged to the correct family slightly more often: 89.40% versus 88.70% .
That makes the errors more interesting than a simple correct-or-wrong score suggests. There’s a substantial difference between getting the broad group right and identifying the particular species.
Each answer appears in one segment. “Only” means the more specific identification was wrong. Failed responses include invalid names, refusals, empty answers, truncations, and request errors.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Counts, not cumulative percentages. Each row adds up to 2,000.
Code and benchmark results on GitHub
Muse Spark 1.3 was the cheapest model here and still finished fourth, ahead of five more expensive models. At its Contributor pricing, the entire run cost $0.93 for 782 correct identifications .
Moving to Gemini 3.8 Flash brought that total to 997. That’s 215 more correct answers for another $26.18 , an improvement of 10.75 percentage points at roughly 29 times the cost . 💸
Spending beyond that didn’t improve the results. Claude Fable 5.1 cost about twice as much as Gemini, and GPT-6 Astra cost 3.6 times as much . Both scored lower. Gemini delivered the highest accuracy here; Muse offered a cheaper trade-off if identifying fewer species was acceptable.
Estimated totals for 2,000 photos. The dollar axis is logarithmic: spacing represents cost ratios.
* Muse Spark 1.3 uses the discounted Contributor tier, which allows provider training/data use.
Code and benchmark results on GitHub
Is getting half wrong actually bad?
Getting almost half right is respectable when the task is telling similar species apart. The gap between recognising a family and naming the exact species is where this gets difficult.
Try telling these two spiders apart. For Araniella opisthographa , NatureSpot’s identification guidance calls for examination at high magnification. That’s the sort of detail a photo can leave out.
The code and full benchmark results are on GitHub .
PS: If you actually know spiders, I’d love to hear what you think on twitter . How obvious are these mistakes to someone who knows what they’re looking at?
P.P.S. Thankfully, I have a cat in charge of spider security at home. Call security
The AI lab for agent swarms.
Build better models, together.
