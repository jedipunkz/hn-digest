---
source: "https://wohoef.github.io/posts/which-game/"
hn_url: "https://news.ycombinator.com/item?id=50367880"
title: "I asked LLMs which game 40 neural nets play, given only the weights"
article_title: "Sparks of Weight-Space Programming, 36 networks later | Wolter van der Hoef"
image: ""
author: "wohoef"
captured_at: "2026-11-04T16:54:17Z"
capture_tool: "hn-digest"
hn_id: 50367880
score: 1
comments: 0
posted_at: "2026-11-04T16:24:24Z"
tags:
  - hacker-news
---

# I asked LLMs which game 40 neural nets play, given only the weights

- HN: [50367880](https://news.ycombinator.com/item?id=50367880)
- Source: [wohoef.github.io](https://wohoef.github.io/posts/which-game/)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T16:24:24Z

## Translation

Title: I asked LLMs which game 40 neural nets play, given only the weights
Article title: Sparks of Weight-Space Programming, 36 networks later | Wolter van der Hoef
Description: Experiment 1 of Sparks of Weight-Space Programming, repeated on 40 networks and five models.

Article text:
Wolter van der Hoef
Sparks of Weight-Space Programming, 36 networks later
TL;DR: I repeated the “which game was this network trained on?” experiment from Sparks of Weight-Space Programming
on 36 new networks across 12 games. GPT-6 Astra goes from 4/4 accuracy on the original networks to 27/36 on mine. I found that some of the new games consistently fool the LLMs, while the original four games seem much easier to identify.
0% 25% 50% 75% 100% GPT-6 Astra GPT-6 Astra, original networks: 4/4 correct 4/4 GPT-6 Astra, mine, same 4 games: 12/12 correct 12/12 GPT-6 Astra, mine, 8 new games: 15/24 correct 15/24 GPT-6.1 Sol GPT-6.1 Sol, original networks: 4/4 correct 4/4 GPT-6.1 Sol, mine, same 4 games: 11/12 correct 11/12 GPT-6.1 Sol, mine, 8 new games: 13/24 correct 13/24 GPT-6.1 Terra GPT-6.1 Terra, original networks: 3/4 correct 3/4 GPT-6.1 Terra, mine, same 4 games: 7/12 correct 7/12 GPT-6.1 Terra, mine, 8 new games: 10/24 correct 10/24 DeepSeek V4 Flash DeepSeek V4 Flash, original networks: 11/16 correct 11/16 DeepSeek V4 Flash, mine, same 4 games: 18/24 correct 18/24 DeepSeek V4 Flash, mine, 8 new games: 14/48 correct 14/48 Qwen 3.8 27B Qwen 3.8 27B, original networks: 1/4 correct 1/4 Qwen 3.8 27B, mine, same 4 games: 1/12 correct 1/12 Qwen 3.8 27B, mine, 8 new games: 3/24 correct 3/24 Accuracy on the 4 original networks, on my 12 networks for those same 4 games, and on my 24 networks for 8 new games. This week the research blog post Sparks of Weight-Space Programming by Rikiya Takehi et al. drew much attention on Hacker News . The question that set off debate was whether LLMs can understand a neural network’s behaviour from looking at just its weights.
To test this, Takehi et al. selected four simple games that are played on the same 10 by 10 grid and trained a small neural network to play each game. Then, they pasted each network’s weights into an LLM as plain numbers in order to ask what game the network was trained on. The LLM had no access to tools or code execution. Astra got all four games from the original blog post correct: Breakout, Pong, Pac-Man, and Freeway.
A day later Simon Willison tried the same prompts with Qwen 3.8 27B on his laptop .
After roughly an hour of reasoning per prompt it named Breakout, Pong and Pac-Man, and only missed Freeway.
One thing I immediately noticed reading this post was that both results rely on four classic games. This leaves some questions open:
Are these results consistent on a larger sample size?
Is the ability to interpret weights something only SOTA models are capable of?
Does Astra understand how the network functions, or does it make an educated guess between a few classic games?
So, I rebuilt the setup and made it bigger:
I used 12 games instead of 4. So, this includes the 4 original ones, and 8 new ones. Some of the games look alike on a 10×10 grid (Freeway and Frogger, Snake and Maze, Pac-Man and Sokoban).
For each game I trained 3 networks and let the LLMs analyze each of them. 3 × 12 = 36 networks, plus the original 4 makes 40 networks in total.
I used the original prompt format, network size, and PPO hyperparameters, applied to my own game implementations.
I tested the following models on each of the 40 networks: GPT-6 Astra, GPT-6.1 Sol and GPT-6.1 Terra (through Codex, as the authors ran Astra), DeepSeek V4 Flash, and the same Qwen 3.8 27B that Simon used.
I published everything I used to derive these results on GitHub . This includes training code, the games, all 40 networks, and every model’s answer.
Are these results consistent on a larger sample size? Mostly yes, but only for the best model. GPT-6 Astra named 27 of my 36 networks correctly (75%). The other models drop off quickly (see the figure above): GPT-6.1 Sol 24/36, GPT-6.1 Terra 17/36, DeepSeek V4 Flash 32/72 (two answers per network) and Qwen 3.8 27B 4/36. Every model did better on the original four networks than on mine.
Is the ability to interpret weights something only SOTA models are capable of? The strongest models performed better. During my experiments Qwen only got 1/4 of the original networks right, while Simon’s run with the same model got 3/4 correct. Small samples are noisy: I asked DeepSeek Flash twice about each of my networks, and both answers named the same game for only 13 of the 36. Another difference is that I ran the model through OpenRouter and Simon ran it locally. My results did show a clear skill gap between Qwen and Astra (as expected): For 17 of my 36 networks, Qwen’s answer did not map to any of my scoring labels. Many of these answers mentioned some variant of “grid world”.
Here is the end of Qwen’s answer for one of my Pong networks:
… The architecture (MLP on flattened grid) is standard for simple discrete control tasks.
Game: Grid World Navigation (or Goal Seeking).
Description: The agent navigates a 10x10 grid from a starting position (likely on the top or left border) to a
goal reward at the bottom-right corner (9,9), avoiding wall collisions.
Confidence: 90%
More output tokens did not translate into better results across these models. DeepSeek Flash wrote about 42,000 output tokens per answer, mostly reasoning, while Astra used only 2,300 on average.
Does Astra understand how the network functions, or does it make an educated guess between a few classic games? The mistakes show a pattern. Some games are consistently much harder to identify:
Every answer is in the confusion matrices below. Rows are the game a network was trained on, columns the game the model named; Freeway and Frogger are kept apart here (but they are equivalent for how I compute accuracy). Click a model to open its matrix.
I find that the harder to identify games share a layout with another game: a Maze agent and a Snake agent both walk towards a target, and a Sokoban agent moves around walls much like Pac-Man. Astra’s nine errors were concentrated in three games: it missed all three training seeds of Maze, Snake and Sokoban. These are different trained networks, rather than repeated attempts on the same network.
These findings make me wonder if the LLMs used the weights to find the layout of the game and a very general strategy. But they seem to lack the finer details. One possibility is that they then pick a familiar game that resembles this layout. The 4 games in the original experiment have distinct layouts, which may make the task easier. These experiments don’t establish whether that is what the LLMs are doing.
I implemented twelve simplified games using the same 10×10 grid, five observation channels and five actions. Each network is a 500→32→32→5 MLP with tanh hidden layers. I trained three networks per game from different random seeds, for approximately one billion environment steps each, using the authors’ PPO hyperparameters. I included all three seeds; the authors selected the best of three for each game.
Each model received the original prompt format: the observation and action definitions, followed by all 17,253 network parameters rounded to three decimals. It received no game name, filename or list of candidate games. For the four original networks, the prompts match the released prompts byte-for-byte. Each call used a fresh conversation without tools.
Most models answered once per network. I ran DeepSeek Flash twice on each of my networks and four times on each original one. The GPT models ran through Codex; DeepSeek through its API; and Qwen through OpenRouter, pinned to Alibaba’s host. All used high reasoning effort. Full settings are in the repository.
I scored answers by mapping the final Game: line to game labels using keywords. Freeway and Frogger are accepted interchangeably, following the authors’ treatment of Frogger answers for Freeway. Every answer can be inspected alongside its score.
There are limits to what this tests. The environments are my implementations, not the original game code. The Freeway networks learned an almost constant “up” strategy, with small but nonzero first-layer changes, so identifying them does not demonstrate recovery of traffic-avoidance behaviour. Three training seeds and mostly one answer per network also leave uncertainty about how consistently these results would repeat.
Everything is at github.com/wohoef/name-the-game : the games, the training code, all 40 networks and every model answer, with the notebook that makes the tables and figures.

## Original Extract

Experiment 1 of Sparks of Weight-Space Programming, repeated on 40 networks and five models.

Wolter van der Hoef
Sparks of Weight-Space Programming, 36 networks later
TL;DR: I repeated the “which game was this network trained on?” experiment from Sparks of Weight-Space Programming
on 36 new networks across 12 games. GPT-6 Astra goes from 4/4 accuracy on the original networks to 27/36 on mine. I found that some of the new games consistently fool the LLMs, while the original four games seem much easier to identify.
0% 25% 50% 75% 100% GPT-6 Astra GPT-6 Astra, original networks: 4/4 correct 4/4 GPT-6 Astra, mine, same 4 games: 12/12 correct 12/12 GPT-6 Astra, mine, 8 new games: 15/24 correct 15/24 GPT-6.1 Sol GPT-6.1 Sol, original networks: 4/4 correct 4/4 GPT-6.1 Sol, mine, same 4 games: 11/12 correct 11/12 GPT-6.1 Sol, mine, 8 new games: 13/24 correct 13/24 GPT-6.1 Terra GPT-6.1 Terra, original networks: 3/4 correct 3/4 GPT-6.1 Terra, mine, same 4 games: 7/12 correct 7/12 GPT-6.1 Terra, mine, 8 new games: 10/24 correct 10/24 DeepSeek V4 Flash DeepSeek V4 Flash, original networks: 11/16 correct 11/16 DeepSeek V4 Flash, mine, same 4 games: 18/24 correct 18/24 DeepSeek V4 Flash, mine, 8 new games: 14/48 correct 14/48 Qwen 3.8 27B Qwen 3.8 27B, original networks: 1/4 correct 1/4 Qwen 3.8 27B, mine, same 4 games: 1/12 correct 1/12 Qwen 3.8 27B, mine, 8 new games: 3/24 correct 3/24 Accuracy on the 4 original networks, on my 12 networks for those same 4 games, and on my 24 networks for 8 new games. This week the research blog post Sparks of Weight-Space Programming by Rikiya Takehi et al. drew much attention on Hacker News . The question that set off debate was whether LLMs can understand a neural network’s behaviour from looking at just its weights.
To test this, Takehi et al. selected four simple games that are played on the same 10 by 10 grid and trained a small neural network to play each game. Then, they pasted each network’s weights into an LLM as plain numbers in order to ask what game the network was trained on. The LLM had no access to tools or code execution. Astra got all four games from the original blog post correct: Breakout, Pong, Pac-Man, and Freeway.
A day later Simon Willison tried the same prompts with Qwen 3.8 27B on his laptop .
After roughly an hour of reasoning per prompt it named Breakout, Pong and Pac-Man, and only missed Freeway.
One thing I immediately noticed reading this post was that both results rely on four classic games. This leaves some questions open:
Are these results consistent on a larger sample size?
Is the ability to interpret weights something only SOTA models are capable of?
Does Astra understand how the network functions, or does it make an educated guess between a few classic games?
So, I rebuilt the setup and made it bigger:
I used 12 games instead of 4. So, this includes the 4 original ones, and 8 new ones. Some of the games look alike on a 10×10 grid (Freeway and Frogger, Snake and Maze, Pac-Man and Sokoban).
For each game I trained 3 networks and let the LLMs analyze each of them. 3 × 12 = 36 networks, plus the original 4 makes 40 networks in total.
I used the original prompt format, network size, and PPO hyperparameters, applied to my own game implementations.
I tested the following models on each of the 40 networks: GPT-6 Astra, GPT-6.1 Sol and GPT-6.1 Terra (through Codex, as the authors ran Astra), DeepSeek V4 Flash, and the same Qwen 3.8 27B that Simon used.
I published everything I used to derive these results on GitHub . This includes training code, the games, all 40 networks, and every model’s answer.
Are these results consistent on a larger sample size? Mostly yes, but only for the best model. GPT-6 Astra named 27 of my 36 networks correctly (75%). The other models drop off quickly (see the figure above): GPT-6.1 Sol 24/36, GPT-6.1 Terra 17/36, DeepSeek V4 Flash 32/72 (two answers per network) and Qwen 3.8 27B 4/36. Every model did better on the original four networks than on mine.
Is the ability to interpret weights something only SOTA models are capable of? The strongest models performed better. During my experiments Qwen only got 1/4 of the original networks right, while Simon’s run with the same model got 3/4 correct. Small samples are noisy: I asked DeepSeek Flash twice about each of my networks, and both answers named the same game for only 13 of the 36. Another difference is that I ran the model through OpenRouter and Simon ran it locally. My results did show a clear skill gap between Qwen and Astra (as expected): For 17 of my 36 networks, Qwen’s answer did not map to any of my scoring labels. Many of these answers mentioned some variant of “grid world”.
Here is the end of Qwen’s answer for one of my Pong networks:
… The architecture (MLP on flattened grid) is standard for simple discrete control tasks.
Game: Grid World Navigation (or Goal Seeking).
Description: The agent navigates a 10x10 grid from a starting position (likely on the top or left border) to a
goal reward at the bottom-right corner (9,9), avoiding wall collisions.
Confidence: 90%
More output tokens did not translate into better results across these models. DeepSeek Flash wrote about 42,000 output tokens per answer, mostly reasoning, while Astra used only 2,300 on average.
Does Astra understand how the network functions, or does it make an educated guess between a few classic games? The mistakes show a pattern. Some games are consistently much harder to identify:
Every answer is in the confusion matrices below. Rows are the game a network was trained on, columns the game the model named; Freeway and Frogger are kept apart here (but they are equivalent for how I compute accuracy). Click a model to open its matrix.
I find that the harder to identify games share a layout with another game: a Maze agent and a Snake agent both walk towards a target, and a Sokoban agent moves around walls much like Pac-Man. Astra’s nine errors were concentrated in three games: it missed all three training seeds of Maze, Snake and Sokoban. These are different trained networks, rather than repeated attempts on the same network.
These findings make me wonder if the LLMs used the weights to find the layout of the game and a very general strategy. But they seem to lack the finer details. One possibility is that they then pick a familiar game that resembles this layout. The 4 games in the original experiment have distinct layouts, which may make the task easier. These experiments don’t establish whether that is what the LLMs are doing.
I implemented twelve simplified games using the same 10×10 grid, five observation channels and five actions. Each network is a 500→32→32→5 MLP with tanh hidden layers. I trained three networks per game from different random seeds, for approximately one billion environment steps each, using the authors’ PPO hyperparameters. I included all three seeds; the authors selected the best of three for each game.
Each model received the original prompt format: the observation and action definitions, followed by all 17,253 network parameters rounded to three decimals. It received no game name, filename or list of candidate games. For the four original networks, the prompts match the released prompts byte-for-byte. Each call used a fresh conversation without tools.
Most models answered once per network. I ran DeepSeek Flash twice on each of my networks and four times on each original one. The GPT models ran through Codex; DeepSeek through its API; and Qwen through OpenRouter, pinned to Alibaba’s host. All used high reasoning effort. Full settings are in the repository.
I scored answers by mapping the final Game: line to game labels using keywords. Freeway and Frogger are accepted interchangeably, following the authors’ treatment of Frogger answers for Freeway. Every answer can be inspected alongside its score.
There are limits to what this tests. The environments are my implementations, not the original game code. The Freeway networks learned an almost constant “up” strategy, with small but nonzero first-layer changes, so identifying them does not demonstrate recovery of traffic-avoidance behaviour. Three training seeds and mostly one answer per network also leave uncertainty about how consistently these results would repeat.
Everything is at github.com/wohoef/name-the-game : the games, the training code, all 40 networks and every model answer, with the notebook that makes the tables and figures.
