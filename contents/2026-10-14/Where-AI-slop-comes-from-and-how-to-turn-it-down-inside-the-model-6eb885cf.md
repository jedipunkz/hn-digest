---
source: "https://100toolkit.com/blog/where-ai-slop-comes-from/"
hn_url: "https://news.ycombinator.com/item?id=50092019"
title: "Where AI slop comes from, and how to turn it down inside the model"
article_title: "Where AI slop comes from, and how to turn it down inside the model"
image: ""
author: "possiblelion"
captured_at: "2026-10-14T18:20:16Z"
capture_tool: "hn-digest"
hn_id: 50092019
score: 1
comments: 0
posted_at: "2026-10-14T17:57:27Z"
tags:
  - hacker-news
---

# Where AI slop comes from, and how to turn it down inside the model

- HN: [50092019](https://news.ycombinator.com/item?id=50092019)
- Source: [100toolkit.com](https://100toolkit.com/blog/where-ai-slop-comes-from/)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T17:57:27Z

## Translation

Title: Where AI slop comes from, and how to turn it down inside the model
Description: AI slop is a direction inside the model. Pretraining learns it, preference tuning switches it on, and zeroing it out makes it worse.

Article text:
← 100toolkit.com ← erik@100toolkit.com Email LinkedIn
just you hide
Where AI slop comes from, and how to turn it down inside the model
The style is learned in pretraining and switched on by post-training, and it sits along one direction inside the model that can be pushed back without breaking it.
The em dashes, "it's not X, it's Y" reframes and lists of three in AI writing move together along one direction inside the model, which we can measure and push on.
The base model, trained only on web text, already contains that direction but writes near the plain end, while the assistant version of the same model writes 2.5 times sloppier.
In OLMo 2, where every training stage is public, supervised fine-tuning brings in the rhetorical moves and preference tuning doubles the score with buzzwords and bullet points, the same things its preference data marks as better.
The "not X, but Y" feature rises a word before the model writes it, and it points at the Chinese equivalent 而是 as readily as at "merely".
Zeroing out the direction, the usual way to remove a behavior, raises slop by up to 89%, while pushing the other way cuts it 85% with a prompt, and a small fine-tune keeps an 83% cut with no prompt.
A demo of the end result before the explanation: Qwen3-4B answering one prompt, and the same model after the fine-tune described at the end. Underlines are what our detector counts; ordinary human writing scores close to zero. Click any word you would cut, and everyone reading right now sees your strike.
Slop is one direction inside the model
A language model represents what it is writing as long vectors of numbers. We had Qwen3-4B, an open 4-billion-parameter assistant, write the same 100 pieces three ways: in its normal voice, as an over-the-top AI assistant and as a plain-spoken person. The average difference between the over-the-top and the plain versions points in one direction. It ranks the sloppy version higher in all 64 hand-written test pairs we tried, and on the model's everyday output it tracks the slop score with a correlation of up to 0.82.
You can also ask a direction which words it pushes the model toward. These are the two ends of this one.
The slop end is a ceremonial register in which things are eternal or ushered in. The plain end is made of the approximate words people use when they talk, like "really", "pretty", "basically" and "stuff".
Tropes are stored as concepts across languages
Each trope also has its own direction, only loosely related to the others. Reading them out shows the model storing the rhetorical move itself, in whatever language it comes in.
The model also commits to a trope early. Its internal slop signal jumps by about one standard deviation on the word before it writes "not" in a "not X, but Y" sentence. Nothing similar happens before an em dash, which has its own direction, almost unrelated to the rest.
Pretraining learns it, post-training switches it on
We ran the same analysis on Qwen's base model, the version trained only to predict internet text, before any assistant training. The slop direction is already there, nearly identical (cosine similarity up to 0.97), and it separates slop from plain writing perfectly, even though the base model does not write that way by default.
The model learned this register from human writing on the web , such as marketing copy and LinkedIn posts, and assistant training made it the default voice.
Preference tuning brings the buzzwords
Most labs don't publish the stages between base and finished model. The Allen Institute does for OLMo 2, along with the data used at each stage, so we measured every checkpoint.
Supervised fine-tuning, where the model learns from example conversations, brings in the rhetorical structures: "not X, but Y" and lists of three roughly triple. Preference tuning, where it learns which of two answers people preferred, doubles the slop score, almost entirely through buzzwords and formatting. Reinforcement learning on math and instruction-following changes little.
Because the preference data is public, we scored it too. Across 3,000 pairs, the answer labelled "better" was the sloppier one 57% of the time.
OLMo almost never uses em dashes at any stage, and they are rare in its training data. Qwen writes about 20 per 1000 words, so Qwen's em dash comes from something in its own pipeline that we can't see.
Zeroing it out makes things worse
The standard way to remove a behavior found like this is to zero out its direction, which is how people strip refusals from open models. For slop it backfired and added up to 89% more, because plain writing sits well below zero on this axis and setting the model to zero moves it toward slop. Pushing a fixed amount the other way works, within a narrow range. Drag the dial.
Banning words can't fix sentence structure
Earlier work on slop bans or penalizes the offending tokens. That removes the em dash and "delve" and leaves everything built from ordinary words, like "not X, but Y" and lists of three, while steering changes the register those sentences come from.
A small fine-tune keeps the effect with no prompt
Steering needs code running inside the model, so to make it permanent the steered model wrote 206 clean answers, we added 92 careful rewrites of the original model's own answers, and trained a small adapter (under 1% of the parameters) on them. With no special prompt it cuts slop 83%, scores 100% on our factual quiz and drifts less on ordinary text than the steering it learned from.
The cleaned-up models kept inventing prices and percentages, and at first we blamed the steering. Counting the numbers each setup introduces showed the cause was our own style prompt, which asked for "concrete details, numbers, names", and the model obeyed by making them up. The adapter trained on the rewrites introduces fewer numbers than the original model.
Does this hold for frontier models?
We can't look inside closed models, so we can't show the same direction exists there. Two parts of the story should carry over: any model trained on web text has seen this register, and any lab that uses preference tuning on human ratings is exposed to the same pull toward formatting and buzzwords, unless its raters or reward models push the other way. We have no explanation yet for Qwen's em dash. The next experiments are a blind test with human readers, a larger open model, and an independent AI-text detector as a second judge.
Models: Qwen3-4B-Instruct-2507 and its base model; OLMo 2 1B at the base, SFT, DPO, RLVR and final stages. OLMo's base repository defaults to a checkpoint that is not the parent of its chat models (their embeddings are unrelated), so we used the final pretraining branch, which matches exactly.
The slop score comes from a regex-based detector we built, so it is best read as a comparison between setups. Every result also reports drift on human prose, accuracy on a 30-question quiz and a flag for broken text.
Steering adds a fixed vector to the model's internal state at layers 12 to 23 of 36. The direction is the average difference between exaggerated-AI and plain versions of the same answers.
The 92 rewrites were produced by a stronger model following a written style guide, not by human writers. Code, prompts and outputs will be released with the paper.
Live cursors and shared strikes connect readers directly over WebRTC (via Trystero and public Nostr relays). Peers can see each other's IP addresses; "hide me" disconnects you. Nothing is stored.

## Original Extract

AI slop is a direction inside the model. Pretraining learns it, preference tuning switches it on, and zeroing it out makes it worse.

← 100toolkit.com ← erik@100toolkit.com Email LinkedIn
just you hide
Where AI slop comes from, and how to turn it down inside the model
The style is learned in pretraining and switched on by post-training, and it sits along one direction inside the model that can be pushed back without breaking it.
The em dashes, "it's not X, it's Y" reframes and lists of three in AI writing move together along one direction inside the model, which we can measure and push on.
The base model, trained only on web text, already contains that direction but writes near the plain end, while the assistant version of the same model writes 2.5 times sloppier.
In OLMo 2, where every training stage is public, supervised fine-tuning brings in the rhetorical moves and preference tuning doubles the score with buzzwords and bullet points, the same things its preference data marks as better.
The "not X, but Y" feature rises a word before the model writes it, and it points at the Chinese equivalent 而是 as readily as at "merely".
Zeroing out the direction, the usual way to remove a behavior, raises slop by up to 89%, while pushing the other way cuts it 85% with a prompt, and a small fine-tune keeps an 83% cut with no prompt.
A demo of the end result before the explanation: Qwen3-4B answering one prompt, and the same model after the fine-tune described at the end. Underlines are what our detector counts; ordinary human writing scores close to zero. Click any word you would cut, and everyone reading right now sees your strike.
Slop is one direction inside the model
A language model represents what it is writing as long vectors of numbers. We had Qwen3-4B, an open 4-billion-parameter assistant, write the same 100 pieces three ways: in its normal voice, as an over-the-top AI assistant and as a plain-spoken person. The average difference between the over-the-top and the plain versions points in one direction. It ranks the sloppy version higher in all 64 hand-written test pairs we tried, and on the model's everyday output it tracks the slop score with a correlation of up to 0.82.
You can also ask a direction which words it pushes the model toward. These are the two ends of this one.
The slop end is a ceremonial register in which things are eternal or ushered in. The plain end is made of the approximate words people use when they talk, like "really", "pretty", "basically" and "stuff".
Tropes are stored as concepts across languages
Each trope also has its own direction, only loosely related to the others. Reading them out shows the model storing the rhetorical move itself, in whatever language it comes in.
The model also commits to a trope early. Its internal slop signal jumps by about one standard deviation on the word before it writes "not" in a "not X, but Y" sentence. Nothing similar happens before an em dash, which has its own direction, almost unrelated to the rest.
Pretraining learns it, post-training switches it on
We ran the same analysis on Qwen's base model, the version trained only to predict internet text, before any assistant training. The slop direction is already there, nearly identical (cosine similarity up to 0.97), and it separates slop from plain writing perfectly, even though the base model does not write that way by default.
The model learned this register from human writing on the web , such as marketing copy and LinkedIn posts, and assistant training made it the default voice.
Preference tuning brings the buzzwords
Most labs don't publish the stages between base and finished model. The Allen Institute does for OLMo 2, along with the data used at each stage, so we measured every checkpoint.
Supervised fine-tuning, where the model learns from example conversations, brings in the rhetorical structures: "not X, but Y" and lists of three roughly triple. Preference tuning, where it learns which of two answers people preferred, doubles the slop score, almost entirely through buzzwords and formatting. Reinforcement learning on math and instruction-following changes little.
Because the preference data is public, we scored it too. Across 3,000 pairs, the answer labelled "better" was the sloppier one 57% of the time.
OLMo almost never uses em dashes at any stage, and they are rare in its training data. Qwen writes about 20 per 1000 words, so Qwen's em dash comes from something in its own pipeline that we can't see.
Zeroing it out makes things worse
The standard way to remove a behavior found like this is to zero out its direction, which is how people strip refusals from open models. For slop it backfired and added up to 89% more, because plain writing sits well below zero on this axis and setting the model to zero moves it toward slop. Pushing a fixed amount the other way works, within a narrow range. Drag the dial.
Banning words can't fix sentence structure
Earlier work on slop bans or penalizes the offending tokens. That removes the em dash and "delve" and leaves everything built from ordinary words, like "not X, but Y" and lists of three, while steering changes the register those sentences come from.
A small fine-tune keeps the effect with no prompt
Steering needs code running inside the model, so to make it permanent the steered model wrote 206 clean answers, we added 92 careful rewrites of the original model's own answers, and trained a small adapter (under 1% of the parameters) on them. With no special prompt it cuts slop 83%, scores 100% on our factual quiz and drifts less on ordinary text than the steering it learned from.
The cleaned-up models kept inventing prices and percentages, and at first we blamed the steering. Counting the numbers each setup introduces showed the cause was our own style prompt, which asked for "concrete details, numbers, names", and the model obeyed by making them up. The adapter trained on the rewrites introduces fewer numbers than the original model.
Does this hold for frontier models?
We can't look inside closed models, so we can't show the same direction exists there. Two parts of the story should carry over: any model trained on web text has seen this register, and any lab that uses preference tuning on human ratings is exposed to the same pull toward formatting and buzzwords, unless its raters or reward models push the other way. We have no explanation yet for Qwen's em dash. The next experiments are a blind test with human readers, a larger open model, and an independent AI-text detector as a second judge.
Models: Qwen3-4B-Instruct-2507 and its base model; OLMo 2 1B at the base, SFT, DPO, RLVR and final stages. OLMo's base repository defaults to a checkpoint that is not the parent of its chat models (their embeddings are unrelated), so we used the final pretraining branch, which matches exactly.
The slop score comes from a regex-based detector we built, so it is best read as a comparison between setups. Every result also reports drift on human prose, accuracy on a 30-question quiz and a flag for broken text.
Steering adds a fixed vector to the model's internal state at layers 12 to 23 of 36. The direction is the average difference between exaggerated-AI and plain versions of the same answers.
The 92 rewrites were produced by a stronger model following a written style guide, not by human writers. Code, prompts and outputs will be released with the paper.
Live cursors and shared strikes connect readers directly over WebRTC (via Trystero and public Nostr relays). Peers can see each other's IP addresses; "hide me" disconnects you. Nothing is stored.
