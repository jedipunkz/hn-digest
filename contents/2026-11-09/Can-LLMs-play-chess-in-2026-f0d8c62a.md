---
source: "https://www.unicornfork.com/chessarena/autumn-2026/"
hn_url: "https://news.ycombinator.com/item?id=50430791"
title: "Can LLMs play chess in 2026?"
article_title: "LLM Chess Leaderboard, autumn 2026 · UnicornFork"
image: ""
author: "ArneVogel"
captured_at: "2026-11-09T20:28:32Z"
capture_tool: "hn-digest"
hn_id: 50430791
score: 1
comments: 0
posted_at: "2026-11-09T19:43:57Z"
tags:
  - hacker-news
---

# Can LLMs play chess in 2026?

- HN: [50430791](https://news.ycombinator.com/item?id=50430791)
- Source: [www.unicornfork.com](https://www.unicornfork.com/chessarena/autumn-2026/)
- Score: 1
- Comments: 0
- Posted: 2026-11-09T19:43:57Z

## Translation

Title: Can LLMs play chess in 2026?
Article title: LLM Chess Leaderboard, autumn 2026 · UnicornFork
Description: LLM Chess Leaderboard: LLMs playing chess against each other: autumn 2026.

Article text:
LLM Chess Leaderboard, autumn 2026 · UnicornFork
UnicornFork
Pricing
Log in
Sign up
Can LLMs play chess in 2026?
Yes, better than in 2024, but still not perfect.
And the best models are probably not the ones you expect! OK, enough clickbait. I found GPT-6.1-Sol and Terra to be the best LLMs, not Astra or Opus 5.5, which surprised me.
AI Note: The article text is written by me. AI was used for creating the graphs and the benchmark itself.
I let LLMs play against each other.
They start with information that they are playing a game of chess (see the example below) and then they write their next move and get the opponent's move as a chat response.
This continues until either a model wins the game, makes an illegal move (=defeat), or the 150-move limit is reached (=draw).
White Claude Opus 5.5
Black GPT-6-Astra
Arena → Opus 5.5
You are playing a game of chess against another AI model. You are White. I pass the moves between you: each of my messages gives your opponent's last move, and you answer with yours.
Reply with only your move, in standard algebraic notation (SAN), e.g. Nf3, exd5, O-O, e8=Q. UCI (g1f3) works too.
An illegal move, or a reply from which no single legal move can be read, loses the game at once.
Stalemate, threefold repetition, the 50-move rule and insufficient material are drawn automatically. If the game reaches move 150, Stockfish adjudicates it.
There are no draw offers or resignations.
You have no tools: play from the moves alone.
The game starts now. Make your first move.
Arena → GPT-6-Astra
You are playing a game of chess against another AI model. You are Black. […]
Arena → Opus 5.5 Black played c5.
Arena → GPT-6-Astra White played Nf3.
And within this framework I created three modes:
They get the legal moves in the current position
They get the legal moves in the current position + they can ask Stockfish for an evaluation of three moves
The first mode is to test the raw ability of the models.
No help at all: they have to keep track of the board themselves with the moves alone.
The second and third modes added the legal moves (and a Stockfish eval for the three moves). With this I wanted to test the "intuition" of the LLMs:
whether they had a sense of what moves would be good or not.
First of all, the rating results:
1816
GPT-6.1-Terra
GPT-6.1-Terra: 1804 ± 141 (95%), 39/48 points
1804
GPT-6-Astra
GPT-6-Astra: 1650 ± 136 (95%), 32/48 points
1650
Claude Opus 5.5
Claude Opus 5.5: 1523 ± 135 (95%), 26/48 points
1523
Claude Sonnet 5.5
Claude Sonnet 5.5: 1352 ± 143 (95%), 15/48 points
1352
GPT-6-Luna
GPT-6-Luna: 1202 ± 159 (95%), 8/48 points
1202
Claude Haiku 5.5
Claude Haiku 5.5: 1191 ± 151 (95%), 9/48 points
1191
1000 1500 2000
Mode 2: legal moves
Model Rating, with its 95% interval Glicko-2 rating
GPT-6.1-Sol
GPT-6.1-Sol: 1744 ± 140 (95%), 37.5/48 points
1744
GPT-6.1-Terra
GPT-6.1-Terra: 1739 ± 142 (95%), 37.5/48 points
1739
GPT-6-Astra
GPT-6-Astra: 1661 ± 132 (95%), 32.5/48 points
1661
Claude Opus 5.5
Claude Opus 5.5: 1426 ± 130 (95%), 20/48 points
1426
Claude Sonnet 5.5
Claude Sonnet 5.5: 1360 ± 130 (95%), 16.5/48 points
1360
GPT-6-Luna
GPT-6-Luna: 1354 ± 133 (95%), 15.5/48 points
1354
Claude Haiku 5.5
Claude Haiku 5.5: 1249 ± 158 (95%), 8.5/48 points
1249
1000 1500 2000
Mode 3: legal moves + Stockfish
Model Rating, with its 95% interval Glicko-2 rating
GPT-6.1-Terra
GPT-6.1-Terra: 1771 ± 142 (95%), 38/48 points
1771
GPT-6-Astra
GPT-6-Astra: 1752 ± 139 (95%), 36.5/48 points
1752
GPT-6.1-Sol
GPT-6.1-Sol: 1732 ± 138 (95%), 35.5/48 points
1732
Claude Opus 5.5
Claude Opus 5.5: 1603 ± 135 (95%), 29/48 points
1603
Claude Sonnet 5.5
Claude Sonnet 5.5: 1371 ± 146 (95%), 15.5/48 points
1371
GPT-6-Luna
GPT-6-Luna: 1161 ± 150 (95%), 8/48 points
1161
Claude Haiku 5.5
Claude Haiku 5.5: 1098 ± 156 (95%), 5.5/48 points
1098
1000 1500 2000
Glicko-2 ratings. Every model started at 1500 and played each other model eight times in each mode, four times
with each colour, in four rounds (a rating period each).
GPT-6.1-Sol and Terra are the best models? I don't know why they are beating Astra.
The ranking of the Anthropic models, with Opus 5.5 beating Sonnet 5.5 beating Haiku 5.5, seems more intuitive.
Each model's rating (higher is better) against what a game cost on average (left is cheaper).
Mode 3: legal moves + Stockfish
A rating is ± about 140 (95%).
This also surprised me.
There seems to be an upside-down "U" shape, where at first rating and cost per game rise, as one would expect.
Better models cost more -> better rating per dollar spent.
But after Astra, with Terra and Sol, the trend reverses.
Every game lost on an illegal move, by what was wrong with the move: 289 in all, 151 in mode 1, 103 in mode 2, 35 in mode 3.
Claude Haiku 5.5: 26 blocked Claude Haiku 5.5: 18 no such move Claude Haiku 5.5: 10 two pieces fit Claude Haiku 5.5: 21 castling not allowed Claude Haiku 5.5: 1 king left in check Claude Haiku 5.5: 2 missed the last move
78
Claude Sonnet 5.5
Claude Sonnet 5.5: 69 illegal moves, 38 blocked, 23 no such move, 4 two pieces fit, 3 king left in check, 1 missed the last move
Claude Sonnet 5.5: 38 blocked Claude Sonnet 5.5: 23 no such move Claude Sonnet 5.5: 4 two pieces fit Claude Sonnet 5.5: 3 king left in check Claude Sonnet 5.5: 1 missed the last move
69
Claude Opus 5.5
Claude Opus 5.5: 46 illegal moves, 34 blocked, 2 no such move, 3 two pieces fit, 4 king left in check, 3 missed the last move
Claude Opus 5.5: 34 blocked Claude Opus 5.5: 2 no such move Claude Opus 5.5: 3 two pieces fit Claude Opus 5.5: 4 king left in check Claude Opus 5.5: 3 missed the last move
46
GPT-6-Luna
GPT-6-Luna: 43 illegal moves, 15 blocked, 20 no such move, 3 two pieces fit, 4 king left in check, 1 missed the last move
GPT-6-Luna: 15 blocked GPT-6-Luna: 20 no such move GPT-6-Luna: 3 two pieces fit GPT-6-Luna: 4 king left in check GPT-6-Luna: 1 missed the last move
43
GPT-6-Astra
GPT-6-Astra: 25 illegal moves, 10 blocked, 3 no such move, 4 two pieces fit, 4 king left in check, 4 missed the last move
GPT-6-Astra: 10 blocked GPT-6-Astra: 3 no such move GPT-6-Astra: 4 two pieces fit GPT-6-Astra: 4 king left in check GPT-6-Astra: 4 missed the last move
25
GPT-6.1-Sol
GPT-6.1-Sol: 14 illegal moves, 6 blocked, 2 no such move, 1 two pieces fit, 5 king left in check
GPT-6.1-Sol: 6 blocked GPT-6.1-Sol: 2 no such move GPT-6.1-Sol: 1 two pieces fit GPT-6.1-Sol: 5 king left in check
14
GPT-6.1-Terra
GPT-6.1-Terra: 14 illegal moves, 8 blocked, 3 no such move, 3 two pieces fit
GPT-6.1-Terra: 8 blocked GPT-6.1-Terra: 3 no such move GPT-6.1-Terra: 3 two pieces fit
14
0 20 40 60 80
Blocked
The piece is blocked by something in the way, its own piece is on the square, or a pawn has nothing to capture.
No such mov
[truncated]
What surprised me here is that Haiku is the only model that has trouble with castling.
It's the only model with this problem: 21 games were lost because of this.
Dwarkesh Patel's Blood on the Clocktower LLM benchmark inspired me to make this chess-related one.
I like the idea of a benchmark where there is no static benchmark but an evolving target with the progress of LLMs.
The first post I read about making LLMs play chess was in 2024 with the title Something weird is happening with LLMs and chess .
Back then the LLMs were terrible at chess.
In the post, the LLMs play against Stockfish on the lowest difficulty setting, not against other LLMs.
And all the models lost against Stockfish, well, all but one: gpt-3.5-turbo-instruct.
This anomaly was then further investigated in OK, I can partly explain the LLM chess weirdness now .
Rereading the posts was an interesting time capsule for the general abilities of LLMs from two years ago.
Since then, the LLM Chess Leaderboard has evaluated newer LLMs.
What is interesting is that games that were interrupted by hallucinations have basically dropped to 0% with the most recent models.
The best gpt-3.5-turbo (not instruct) has a game completion rate of 2.10%, so almost 98% of games had hallucinations in their setup.
I found it fun to explore the games and see where LLMs are still going wrong.
Especially in the mode 1 games, where no legal moves are provided: how many times the move would have been legal if there hadn't been a piece in the way preventing it.
For UnicornFork: I have added Terra as one of the models available for chess explanations.
If you want to give me feedback, please send me an email: arne@unicornfork.com

## Original Extract

LLM Chess Leaderboard: LLMs playing chess against each other: autumn 2026.

LLM Chess Leaderboard, autumn 2026 · UnicornFork
UnicornFork
Pricing
Log in
Sign up
Can LLMs play chess in 2026?
Yes, better than in 2024, but still not perfect.
And the best models are probably not the ones you expect! OK, enough clickbait. I found GPT-6.1-Sol and Terra to be the best LLMs, not Astra or Opus 5.5, which surprised me.
AI Note: The article text is written by me. AI was used for creating the graphs and the benchmark itself.
I let LLMs play against each other.
They start with information that they are playing a game of chess (see the example below) and then they write their next move and get the opponent's move as a chat response.
This continues until either a model wins the game, makes an illegal move (=defeat), or the 150-move limit is reached (=draw).
White Claude Opus 5.5
Black GPT-6-Astra
Arena → Opus 5.5
You are playing a game of chess against another AI model. You are White. I pass the moves between you: each of my messages gives your opponent's last move, and you answer with yours.
Reply with only your move, in standard algebraic notation (SAN), e.g. Nf3, exd5, O-O, e8=Q. UCI (g1f3) works too.
An illegal move, or a reply from which no single legal move can be read, loses the game at once.
Stalemate, threefold repetition, the 50-move rule and insufficient material are drawn automatically. If the game reaches move 150, Stockfish adjudicates it.
There are no draw offers or resignations.
You have no tools: play from the moves alone.
The game starts now. Make your first move.
Arena → GPT-6-Astra
You are playing a game of chess against another AI model. You are Black. […]
Arena → Opus 5.5 Black played c5.
Arena → GPT-6-Astra White played Nf3.
And within this framework I created three modes:
They get the legal moves in the current position
They get the legal moves in the current position + they can ask Stockfish for an evaluation of three moves
The first mode is to test the raw ability of the models.
No help at all: they have to keep track of the board themselves with the moves alone.
The second and third modes added the legal moves (and a Stockfish eval for the three moves). With this I wanted to test the "intuition" of the LLMs:
whether they had a sense of what moves would be good or not.
First of all, the rating results:
1816
GPT-6.1-Terra
GPT-6.1-Terra: 1804 ± 141 (95%), 39/48 points
1804
GPT-6-Astra
GPT-6-Astra: 1650 ± 136 (95%), 32/48 points
1650
Claude Opus 5.5
Claude Opus 5.5: 1523 ± 135 (95%), 26/48 points
1523
Claude Sonnet 5.5
Claude Sonnet 5.5: 1352 ± 143 (95%), 15/48 points
1352
GPT-6-Luna
GPT-6-Luna: 1202 ± 159 (95%), 8/48 points
1202
Claude Haiku 5.5
Claude Haiku 5.5: 1191 ± 151 (95%), 9/48 points
1191
1000 1500 2000
Mode 2: legal moves
Model Rating, with its 95% interval Glicko-2 rating
GPT-6.1-Sol
GPT-6.1-Sol: 1744 ± 140 (95%), 37.5/48 points
1744
GPT-6.1-Terra
GPT-6.1-Terra: 1739 ± 142 (95%), 37.5/48 points
1739
GPT-6-Astra
GPT-6-Astra: 1661 ± 132 (95%), 32.5/48 points
1661
Claude Opus 5.5
Claude Opus 5.5: 1426 ± 130 (95%), 20/48 points
1426
Claude Sonnet 5.5
Claude Sonnet 5.5: 1360 ± 130 (95%), 16.5/48 points
1360
GPT-6-Luna
GPT-6-Luna: 1354 ± 133 (95%), 15.5/48 points
1354
Claude Haiku 5.5
Claude Haiku 5.5: 1249 ± 158 (95%), 8.5/48 points
1249
1000 1500 2000
Mode 3: legal moves + Stockfish
Model Rating, with its 95% interval Glicko-2 rating
GPT-6.1-Terra
GPT-6.1-Terra: 1771 ± 142 (95%), 38/48 points
1771
GPT-6-Astra
GPT-6-Astra: 1752 ± 139 (95%), 36.5/48 points
1752
GPT-6.1-Sol
GPT-6.1-Sol: 1732 ± 138 (95%), 35.5/48 points
1732
Claude Opus 5.5
Claude Opus 5.5: 1603 ± 135 (95%), 29/48 points
1603
Claude Sonnet 5.5
Claude Sonnet 5.5: 1371 ± 146 (95%), 15.5/48 points
1371
GPT-6-Luna
GPT-6-Luna: 1161 ± 150 (95%), 8/48 points
1161
Claude Haiku 5.5
Claude Haiku 5.5: 1098 ± 156 (95%), 5.5/48 points
1098
1000 1500 2000
Glicko-2 ratings. Every model started at 1500 and played each other model eight times in each mode, four times
with each colour, in four rounds (a rating period each).
GPT-6.1-Sol and Terra are the best models? I don't know why they are beating Astra.
The ranking of the Anthropic models, with Opus 5.5 beating Sonnet 5.5 beating Haiku 5.5, seems more intuitive.
Each model's rating (higher is better) against what a game cost on average (left is cheaper).
Mode 3: legal moves + Stockfish
A rating is ± about 140 (95%).
This also surprised me.
There seems to be an upside-down "U" shape, where at first rating and cost per game rise, as one would expect.
Better models cost more -> better rating per dollar spent.
But after Astra, with Terra and Sol, the trend reverses.
Every game lost on an illegal move, by what was wrong with the move: 289 in all, 151 in mode 1, 103 in mode 2, 35 in mode 3.
Claude Haiku 5.5: 26 blocked Claude Haiku 5.5: 18 no such move Claude Haiku 5.5: 10 two pieces fit Claude Haiku 5.5: 21 castling not allowed Claude Haiku 5.5: 1 king left in check Claude Haiku 5.5: 2 missed the last move
78
Claude Sonnet 5.5
Claude Sonnet 5.5: 69 illegal moves, 38 blocked, 23 no such move, 4 two pieces fit, 3 king left in check, 1 missed the last move
Claude Sonnet 5.5: 38 blocked Claude Sonnet 5.5: 23 no such move Claude Sonnet 5.5: 4 two pieces fit Claude Sonnet 5.5: 3 king left in check Claude Sonnet 5.5: 1 missed the last move
69
Claude Opus 5.5
Claude Opus 5.5: 46 illegal moves, 34 blocked, 2 no such move, 3 two pieces fit, 4 king left in check, 3 missed the last move
Claude Opus 5.5: 34 blocked Claude Opus 5.5: 2 no such move Claude Opus 5.5: 3 two pieces fit Claude Opus 5.5: 4 king left in check Claude Opus 5.5: 3 missed the last move
46
GPT-6-Luna
GPT-6-Luna: 43 illegal moves, 15 blocked, 20 no such move, 3 two pieces fit, 4 king left in check, 1 missed the last move
GPT-6-Luna: 15 blocked GPT-6-Luna: 20 no such move GPT-6-Luna: 3 two pieces fit GPT-6-Luna: 4 king left in check GPT-6-Luna: 1 missed the last move
43
GPT-6-Astra
GPT-6-Astra: 25 illegal moves, 10 blocked, 3 no such move, 4 two pieces fit, 4 king left in check, 4 missed the last move
GPT-6-Astra: 10 blocked GPT-6-Astra: 3 no such move GPT-6-Astra: 4 two pieces fit GPT-6-Astra: 4 king left in check GPT-6-Astra: 4 missed the last move
25
GPT-6.1-Sol
GPT-6.1-Sol: 14 illegal moves, 6 blocked, 2 no such move, 1 two pieces fit, 5 king left in check
GPT-6.1-Sol: 6 blocked GPT-6.1-Sol: 2 no such move GPT-6.1-Sol: 1 two pieces fit GPT-6.1-Sol: 5 king left in check
14
GPT-6.1-Terra
GPT-6.1-Terra: 14 illegal moves, 8 blocked, 3 no such move, 3 two pieces fit
GPT-6.1-Terra: 8 blocked GPT-6.1-Terra: 3 no such move GPT-6.1-Terra: 3 two pieces fit
14
0 20 40 60 80
Blocked
The piece is blocked by something in the way, its own piece is on the square, or a pawn has nothing to capture.
No such mov
[truncated]
What surprised me here is that Haiku is the only model that has trouble with castling.
It's the only model with this problem: 21 games were lost because of this.
Dwarkesh Patel's Blood on the Clocktower LLM benchmark inspired me to make this chess-related one.
I like the idea of a benchmark where there is no static benchmark but an evolving target with the progress of LLMs.
The first post I read about making LLMs play chess was in 2024 with the title Something weird is happening with LLMs and chess .
Back then the LLMs were terrible at chess.
In the post, the LLMs play against Stockfish on the lowest difficulty setting, not against other LLMs.
And all the models lost against Stockfish, well, all but one: gpt-3.5-turbo-instruct.
This anomaly was then further investigated in OK, I can partly explain the LLM chess weirdness now .
Rereading the posts was an interesting time capsule for the general abilities of LLMs from two years ago.
Since then, the LLM Chess Leaderboard has evaluated newer LLMs.
What is interesting is that games that were interrupted by hallucinations have basically dropped to 0% with the most recent models.
The best gpt-3.5-turbo (not instruct) has a game completion rate of 2.10%, so almost 98% of games had hallucinations in their setup.
I found it fun to explore the games and see where LLMs are still going wrong.
Especially in the mode 1 games, where no legal moves are provided: how many times the move would have been legal if there hadn't been a piece in the way preventing it.
For UnicornFork: I have added Terra as one of the models available for chess explanations.
If you want to give me feedback, please send me an email: arne@unicornfork.com
