---
source: "https://epoch.ai/publications/ai-systems-are-improving-rapidly-at-slay-the-spire-and-its-sequel"
hn_url: "https://news.ycombinator.com/item?id=50103102"
title: "How well can AI play Slay the Spire?"
article_title: "How well can AI play Slay the Spire? | Epoch AI"
image: "https://epoch.ai/assets/images/posts/2026/ai-systems-are-improving-rapidly-at-slay-the-spire-and-its-sequel/soc-fig2.png"
author: "agentdev001"
captured_at: "2026-10-15T13:06:52Z"
capture_tool: "hn-digest"
hn_id: 50103102
score: 1
comments: 0
posted_at: "2026-10-15T13:05:41Z"
tags:
  - hacker-news
---

# How well can AI play Slay the Spire?

- HN: [50103102](https://news.ycombinator.com/item?id=50103102)
- Source: [epoch.ai](https://epoch.ai/publications/ai-systems-are-improving-rapidly-at-slay-the-spire-and-its-sequel)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T13:05:41Z

## Translation

Title: How well can AI play Slay the Spire?
Article title: How well can AI play Slay the Spire? | Epoch AI
Description: Months ago, frontier models could barely beat Slay the Spire. Now they win 90%+ at the easiest setting and up to 20% at the hardest. Results from Epoch AI’s new SpireClimb benchmark.

Article text:
How well can AI play Slay the Spire? | Epoch AI
Latest Our work Featured Trends in AI
Our work Featured Trends in AI Data on AI Capabilities & benchmarking Publications Papers & reports Data Insights Newsletter Podcast All publications Data explorers AI capabilities AI models AI data centers AI chip sales AI companies Polling on AI usage All data explorers Our benchmarks Epoch Capabilities Index MirrorCode FrontierMath: Open Problems Epoch Automation Reports Earthborne Rangers
Navigate by topic AI progress Scaling Software progress Open models Capabilities & benchmarks Math Industry Leading companies Finances Geopolitics Infrastructure Chips Data centers Energy Impacts Adoption & use Economic impact Future of AI
About Who we are About us Our team Transparency Engage with us Careers Press Work with our experts Donate
AI systems are improving rapidly at Slay the Spire and its sequel
AI systems are improving rapidly at Slay the Spire and its sequel
In just six months, frontier models have gone from struggling to beat the games at all to over 90% win rates at base difficulty—even when playing through a visual computer-use harness.
Cite
Benchmark page By Benjamin Ou and Simon Chen Contents
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Some frontier models can beat Slay the Spire 1 on maximum difficulty
More models can beat Slay the Spire 2 on maximum difficulty
Elicitation gaps likely remain
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Some frontier models can beat Slay the Spire 1 on maximum difficulty
More models can beat Slay the Spire 2 on maximum difficulty
Elicitation gaps likely remain
Slay the Spire (StS1) and its recent sequel, Slay the Spire 2 (StS2), are difficult games.
Since the original’s release in 2017, the game has attracted a dedicated community of players.
Top-tier players play the games on their maximum difficulty settings and chase “win streaks” that require high win rates to sustain consecutive victories.
Slay the Spire’s difficulty and popularity have more recently drawn attention from AI researchers.
In its release of Claude Fable 5 in June 2026, Anthropic noted that Fable 5 “reached the game’s final act three times more often” than Claude Opus 4.8, a phrasing that indicated the model still could not reliably win games.
METR researcher Ajeya Cotra used win rates for Slay the Spire 2 as part of her forecasts for AI capabilities in 2026 .
As a single-player game with randomness, hidden information, and a huge decision tree, Slay the Spire has long remained resistant to being conquered by both classic machine learning techniques and large language models (LLMs).
Here, we present findings from our new benchmark, SpireClimb , that show this is rapidly changing.
Frontier models can now beat Slay the Spire 1 at base difficulty (“Ascension 0”) at least 90% of the time. Just six months ago, they struggled to beat the game at Ascension 0 at all.
They can also beat Slay the Spire 1’s maximum difficulty (“Ascension 20 Heart”) , with GPT-6 Astra achieving a 17% win rate and GPT-6.1 Sol achieving a 20% win rate.
Models also perform better at StS2’s hardest setting (Ascension 10) than at comparable difficulty tiers in the original. Frontier models reached win rates above 25%, reflecting StS2’s lower overall difficulty in its early-access version.
These win rates remain far below those of top human players , who regularly achieve at least 80% win rates at maximum difficulty in both games.
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Prior efforts at using Slay the Spire as an LLM benchmark have been scattershot, with substantial variation among harnesses used, affordances allowed, and models tested.
The Orak paper in June 2025 found that Gemini 2.5 Pro with a text harness could “defeat the second boss in two out of three runs”, which likely corresponds to a 0% overall win rate.
Results from StSBench , which also uses a text harness, show that by June 2026, frontier model win rates had become high enough at Ascension 0 to score wins given only a single attempt.
More recently, Epoch’s own informal experiments with the Codex Desktop computer-use harness found GPT-5.6 was able to beat the moderate difficulty level Ascension 8 in August 2026, which likely corresponds to an Ascension 0 win rate well above 50%.
Given this state, we believe there’s value in a consistent framework for evaluating current and future models’ performance at both StS1 and StS2.
Additionally, we’re interested in evaluating models’ computer use abilities, which are also rapidly improving.
Our goal is to roughly measure the performance an unskilled user could achieve by opening an LLM desktop app on their computer, typing in “play Slay the Spire for me”, and letting it run.
To that end, we present our findings here as part of a new Epoch benchmark, SpireClimb.
We aim to evaluate models on both games at maximum difficulty in a single-agent visual computer-use harness based on Inspect ReAct . 1
We sample a fixed set of arbitrarily chosen seeds, with each model given 15 independent single-shot attempts at winning a run for each character in both games.
Models are allowed to take notes to persist information across context compaction, and can consult the Slay the Spire wiki so they have fair access to information an experienced player would have.
We target the maximum difficulty for both games because we’ve found that, under these conditions, frontier and near-frontier models can already beat the easiest difficulty level, Ascension 0, on a vision harness the vast majority of the time.
All models we tested showed substantial pre-existing knowledge about StS1 in their reasoning, frequently recalling specific facts and numeric values, upcoming threats to plan around, and common community opinions about the strength of various cards and strategies.
This aligns with our expectations, given that StS1 was released nearly a decade ago; older models also display similar levels of knowledge.
We therefore believe this recent spike in performance can be attributed to some combination of recent improvements in computer use capability, general game-playing ability, and perhaps also specific reinforcement learning against StS1.
Some frontier models can beat Slay the Spire 1 on maximum difficulty
While reliably beating Ascension 0 is a significant achievement for frontier models, particularly through a visual computer-use harness, the attention of top players and dedicated fans has always been focused on StS1’s maximum difficulty, “Ascension 20 Heart”.
This setting asks the player to not only play at the maximum Ascension level, but to defeat the optional final boss, “The Heart”.
The Heart is substantially more difficult than any other encounter in the game, and players who can reliably beat Ascension 0 can spend many dozens or hundreds of hours on their attempts to beat it.
That said, beating The Heart is far from impossible.
Strong StS1 players regularly show win rates above 50% , with “XecnaR” (widely considered to be the single best player) showing win rates from 80% to 95% depending on the character.
This level of performance requires intimate knowledge of the game’s many details, with every decision balancing the short-term gains needed to survive each encounter with the long-term gains needed to defeat the final boss.
Just a few months ago, frontier models showed no sign of approaching these capabilities, with few even bothering to try measuring performance at this difficulty level.
But with the release of GPT-6 Astra, followed by GPT-6.1 Sol, this has changed.
While still far below the level of top human players, a 20% win rate—up from 0% a few months ago—is a major jump in capability.
Notably, Sol’s performance differs between characters in a way that roughly mirrors XecnaR’s statistics, performing better with Ironclad and Watcher than Silent and Defect.
This likely reflects the inherent gap in power level between the characters, with Watcher in particular broadly considered the strongest character.
We aren’t sure why Sol’s performance with Silent is extremely low.
GPT-6 Astra, despite having a slightly lower overall win rate, managed to score 1 win with Silent out of 15 samples, for a 7% win rate.
This is statistically indistinguishable from Sol’s performance, so frontier models may simply find the character unusually challenging.
More models can beat Slay the Spire 2 on maximum difficulty
If frontier labs are using games like StS1 as training environments for their models, we might question whether these results reflect an increase in general capabilities.
As previously mentioned, LLMs seem to know a lot about strong StS1 strategies, so what happens if they play a game they know less about?
Testing models against StS2 can provide insight into this question, albeit with some trade-offs and limitations.
The sequel’s main advantage here is that it is new enough to be largely outside frontier models’ training data, giving us a clearer look at how well their strategy game skills generalize to a less familiar game.
However, benchmarks require a fixed task to be interpretable, whereas StS2 is in early access and regularly receives content updates.
We will not incorporate future content updates to the version of StS2 we use for this benchmark, instead freezing it at v0.107.1 . 2
Additionally, as StS2 currently lacks an equivalent to the first game’s highly difficult final boss, The Heart, it is broadly considered much easier than the first game by most players. 3
Top player data is less available here than in StS1, but community aggregation sites like SpireBird and Spire Codex show that hundreds of players can sustain win rates above 80% in StS2’s maximum difficulty, Ascension 10.
We found no consensus best player with readily available published data, but we did find video evidence of a player from China, “白夕Seal”, achieving 53 consecutive Ascension 10 wins rotating between characters.
While deducing a precise win rate from a potentially cherry-picked win streak is difficult, naively this corresponds to a win rate of at least 90%, likely closer to 95%.
As such, all evidence points to StS2’s maximum difficulty being easier than that of the first game.
Accordingly, the top models’ win rates here are about twice as high.
Breaking down the top model’s performance by character, we might expect to see the sequel’s new characters, Regent and Necrobinder, perform worse given their plausible lack of weight in training data.
However, of the two, only Necrobinder shows a noticeably lower win rate.
Compared to StS1, no clear pattern emerges here that is consistent with the measured strength of each character in StS2, which, on Spire Codex , shows Silent with a relatively high win rate and Defect with a relatively low win rate among high-ranking players.
Our relatively low per-character sample sizes also make drawing a firm conclusion difficult.
Regardless, we interpret these results as evidence that frontier models do not take a significant penalty in playing a less familiar game.
Elicitation gaps likely remain
Our choice to present the game to models through a vision-based computer-use harness likely fails to elicit the full extent of models’ pure reasoning skills in playing StS1 and StS2.
Interface mistakes still occur, and requiring models to play the game through its graphical interface may encourage quicker decisions at the expense of more deliberate reasoning.
Regardless, we think this benchmark is a useful measure of how models perform when playing in conditions similar to those faced by human players.
As model capabilities improve rapidly across multiple dimensions, benchmarks that capture the full breadth of their skills are increasingly important.
A model that can strategically navigate the interface of complex games while coherently pursuing its goals may also be better positioned to automate existing human tasks end-to-end, without needing custom harnesses o

[truncated]

## Original Extract

Months ago, frontier models could barely beat Slay the Spire. Now they win 90%+ at the easiest setting and up to 20% at the hardest. Results from Epoch AI’s new SpireClimb benchmark.

How well can AI play Slay the Spire? | Epoch AI
Latest Our work Featured Trends in AI
Our work Featured Trends in AI Data on AI Capabilities & benchmarking Publications Papers & reports Data Insights Newsletter Podcast All publications Data explorers AI capabilities AI models AI data centers AI chip sales AI companies Polling on AI usage All data explorers Our benchmarks Epoch Capabilities Index MirrorCode FrontierMath: Open Problems Epoch Automation Reports Earthborne Rangers
Navigate by topic AI progress Scaling Software progress Open models Capabilities & benchmarks Math Industry Leading companies Finances Geopolitics Infrastructure Chips Data centers Energy Impacts Adoption & use Economic impact Future of AI
About Who we are About us Our team Transparency Engage with us Careers Press Work with our experts Donate
AI systems are improving rapidly at Slay the Spire and its sequel
AI systems are improving rapidly at Slay the Spire and its sequel
In just six months, frontier models have gone from struggling to beat the games at all to over 90% win rates at base difficulty—even when playing through a visual computer-use harness.
Cite
Benchmark page By Benjamin Ou and Simon Chen Contents
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Some frontier models can beat Slay the Spire 1 on maximum difficulty
More models can beat Slay the Spire 2 on maximum difficulty
Elicitation gaps likely remain
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Some frontier models can beat Slay the Spire 1 on maximum difficulty
More models can beat Slay the Spire 2 on maximum difficulty
Elicitation gaps likely remain
Slay the Spire (StS1) and its recent sequel, Slay the Spire 2 (StS2), are difficult games.
Since the original’s release in 2017, the game has attracted a dedicated community of players.
Top-tier players play the games on their maximum difficulty settings and chase “win streaks” that require high win rates to sustain consecutive victories.
Slay the Spire’s difficulty and popularity have more recently drawn attention from AI researchers.
In its release of Claude Fable 5 in June 2026, Anthropic noted that Fable 5 “reached the game’s final act three times more often” than Claude Opus 4.8, a phrasing that indicated the model still could not reliably win games.
METR researcher Ajeya Cotra used win rates for Slay the Spire 2 as part of her forecasts for AI capabilities in 2026 .
As a single-player game with randomness, hidden information, and a huge decision tree, Slay the Spire has long remained resistant to being conquered by both classic machine learning techniques and large language models (LLMs).
Here, we present findings from our new benchmark, SpireClimb , that show this is rapidly changing.
Frontier models can now beat Slay the Spire 1 at base difficulty (“Ascension 0”) at least 90% of the time. Just six months ago, they struggled to beat the game at Ascension 0 at all.
They can also beat Slay the Spire 1’s maximum difficulty (“Ascension 20 Heart”) , with GPT-6 Astra achieving a 17% win rate and GPT-6.1 Sol achieving a 20% win rate.
Models also perform better at StS2’s hardest setting (Ascension 10) than at comparable difficulty tiers in the original. Frontier models reached win rates above 25%, reflecting StS2’s lower overall difficulty in its early-access version.
These win rates remain far below those of top human players , who regularly achieve at least 80% win rates at maximum difficulty in both games.
LLMs have rapidly improved at Slay the Spire 1’s base difficulty
Prior efforts at using Slay the Spire as an LLM benchmark have been scattershot, with substantial variation among harnesses used, affordances allowed, and models tested.
The Orak paper in June 2025 found that Gemini 2.5 Pro with a text harness could “defeat the second boss in two out of three runs”, which likely corresponds to a 0% overall win rate.
Results from StSBench , which also uses a text harness, show that by June 2026, frontier model win rates had become high enough at Ascension 0 to score wins given only a single attempt.
More recently, Epoch’s own informal experiments with the Codex Desktop computer-use harness found GPT-5.6 was able to beat the moderate difficulty level Ascension 8 in August 2026, which likely corresponds to an Ascension 0 win rate well above 50%.
Given this state, we believe there’s value in a consistent framework for evaluating current and future models’ performance at both StS1 and StS2.
Additionally, we’re interested in evaluating models’ computer use abilities, which are also rapidly improving.
Our goal is to roughly measure the performance an unskilled user could achieve by opening an LLM desktop app on their computer, typing in “play Slay the Spire for me”, and letting it run.
To that end, we present our findings here as part of a new Epoch benchmark, SpireClimb.
We aim to evaluate models on both games at maximum difficulty in a single-agent visual computer-use harness based on Inspect ReAct . 1
We sample a fixed set of arbitrarily chosen seeds, with each model given 15 independent single-shot attempts at winning a run for each character in both games.
Models are allowed to take notes to persist information across context compaction, and can consult the Slay the Spire wiki so they have fair access to information an experienced player would have.
We target the maximum difficulty for both games because we’ve found that, under these conditions, frontier and near-frontier models can already beat the easiest difficulty level, Ascension 0, on a vision harness the vast majority of the time.
All models we tested showed substantial pre-existing knowledge about StS1 in their reasoning, frequently recalling specific facts and numeric values, upcoming threats to plan around, and common community opinions about the strength of various cards and strategies.
This aligns with our expectations, given that StS1 was released nearly a decade ago; older models also display similar levels of knowledge.
We therefore believe this recent spike in performance can be attributed to some combination of recent improvements in computer use capability, general game-playing ability, and perhaps also specific reinforcement learning against StS1.
Some frontier models can beat Slay the Spire 1 on maximum difficulty
While reliably beating Ascension 0 is a significant achievement for frontier models, particularly through a visual computer-use harness, the attention of top players and dedicated fans has always been focused on StS1’s maximum difficulty, “Ascension 20 Heart”.
This setting asks the player to not only play at the maximum Ascension level, but to defeat the optional final boss, “The Heart”.
The Heart is substantially more difficult than any other encounter in the game, and players who can reliably beat Ascension 0 can spend many dozens or hundreds of hours on their attempts to beat it.
That said, beating The Heart is far from impossible.
Strong StS1 players regularly show win rates above 50% , with “XecnaR” (widely considered to be the single best player) showing win rates from 80% to 95% depending on the character.
This level of performance requires intimate knowledge of the game’s many details, with every decision balancing the short-term gains needed to survive each encounter with the long-term gains needed to defeat the final boss.
Just a few months ago, frontier models showed no sign of approaching these capabilities, with few even bothering to try measuring performance at this difficulty level.
But with the release of GPT-6 Astra, followed by GPT-6.1 Sol, this has changed.
While still far below the level of top human players, a 20% win rate—up from 0% a few months ago—is a major jump in capability.
Notably, Sol’s performance differs between characters in a way that roughly mirrors XecnaR’s statistics, performing better with Ironclad and Watcher than Silent and Defect.
This likely reflects the inherent gap in power level between the characters, with Watcher in particular broadly considered the strongest character.
We aren’t sure why Sol’s performance with Silent is extremely low.
GPT-6 Astra, despite having a slightly lower overall win rate, managed to score 1 win with Silent out of 15 samples, for a 7% win rate.
This is statistically indistinguishable from Sol’s performance, so frontier models may simply find the character unusually challenging.
More models can beat Slay the Spire 2 on maximum difficulty
If frontier labs are using games like StS1 as training environments for their models, we might question whether these results reflect an increase in general capabilities.
As previously mentioned, LLMs seem to know a lot about strong StS1 strategies, so what happens if they play a game they know less about?
Testing models against StS2 can provide insight into this question, albeit with some trade-offs and limitations.
The sequel’s main advantage here is that it is new enough to be largely outside frontier models’ training data, giving us a clearer look at how well their strategy game skills generalize to a less familiar game.
However, benchmarks require a fixed task to be interpretable, whereas StS2 is in early access and regularly receives content updates.
We will not incorporate future content updates to the version of StS2 we use for this benchmark, instead freezing it at v0.107.1 . 2
Additionally, as StS2 currently lacks an equivalent to the first game’s highly difficult final boss, The Heart, it is broadly considered much easier than the first game by most players. 3
Top player data is less available here than in StS1, but community aggregation sites like SpireBird and Spire Codex show that hundreds of players can sustain win rates above 80% in StS2’s maximum difficulty, Ascension 10.
We found no consensus best player with readily available published data, but we did find video evidence of a player from China, “白夕Seal”, achieving 53 consecutive Ascension 10 wins rotating between characters.
While deducing a precise win rate from a potentially cherry-picked win streak is difficult, naively this corresponds to a win rate of at least 90%, likely closer to 95%.
As such, all evidence points to StS2’s maximum difficulty being easier than that of the first game.
Accordingly, the top models’ win rates here are about twice as high.
Breaking down the top model’s performance by character, we might expect to see the sequel’s new characters, Regent and Necrobinder, perform worse given their plausible lack of weight in training data.
However, of the two, only Necrobinder shows a noticeably lower win rate.
Compared to StS1, no clear pattern emerges here that is consistent with the measured strength of each character in StS2, which, on Spire Codex , shows Silent with a relatively high win rate and Defect with a relatively low win rate among high-ranking players.
Our relatively low per-character sample sizes also make drawing a firm conclusion difficult.
Regardless, we interpret these results as evidence that frontier models do not take a significant penalty in playing a less familiar game.
Elicitation gaps likely remain
Our choice to present the game to models through a vision-based computer-use harness likely fails to elicit the full extent of models’ pure reasoning skills in playing StS1 and StS2.
Interface mistakes still occur, and requiring models to play the game through its graphical interface may encourage quicker decisions at the expense of more deliberate reasoning.
Regardless, we think this benchmark is a useful measure of how models perform when playing in conditions similar to those faced by human players.
As model capabilities improve rapidly across multiple dimensions, benchmarks that capture the full breadth of their skills are increasingly important.
A model that can strategically navigate the interface of complex games while coherently pursuing its goals may also be better positioned to automate existing human tasks end-to-end, without needing custom harnesses o

[truncated]
