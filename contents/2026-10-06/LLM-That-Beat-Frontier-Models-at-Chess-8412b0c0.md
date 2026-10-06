---
source: "https://www.aidancooper.co.uk/the-2023-llm-that-beat-frontier-models-at-chess/"
hn_url: "https://news.ycombinator.com/item?id=49982787"
title: "LLM That Beat Frontier Models at Chess"
article_title: "The 2023 LLM That Beat Frontier Models at Chess"
image: "https://storage.ghost.io/c/36/28/362843d9-d91b-4b54-b9b1-bf5e6c6c5d54/content/images/size/w1200/2026/09/Screenshot-2026-09-27-at-16.13.43.png"
author: "speckx"
captured_at: "2026-10-06T19:23:19Z"
capture_tool: "hn-digest"
hn_id: 49982787
score: 1
comments: 0
posted_at: "2026-10-06T19:19:23Z"
tags:
  - hacker-news
---

# LLM That Beat Frontier Models at Chess

- HN: [49982787](https://news.ycombinator.com/item?id=49982787)
- Source: [www.aidancooper.co.uk](https://www.aidancooper.co.uk/the-2023-llm-that-beat-frontier-models-at-chess/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T19:19:23Z

## Translation

Title: LLM That Beat Frontier Models at Chess
Article title: The 2023 LLM That Beat Frontier Models at Chess
Description: The result says as much about game representation as it does about model capability.

Article text:
Sign in
Subscribe
Sep 27, 2026
4 min read
LLMs
The 2023 LLM That Beat Frontier Models at Chess
I recently repeated an experiment that I've run a few times over the past of couple years: pitting GPT-3.5 Turbo Instruct against frontier LLMs at chess. GPT-3.5 remains undefeated, winning all four games against GPT-5.6 Sol and Claude Opus 5. Last year , it beat o3, Grok 4 and Gemini 2.5 Pro six games to zero.
Ten games, ten wins. OpenAI is going to shut down access to GPT-3.5 Turbo Instruct tomorrow , so now seems like the time to explore why this 2023 LLM retires undefeated.
GPT-3.5 Turbo Instruct is instruction-tuned, but uses a completion interface - it continues the supplied text directly. For chess, this means supplying GPT-3.5 with an unfinished game represented as PGN , which it then continues with the next move.
[Event "World Championship"]
[Site "Moscow"]
[Date "2025.01.01"]
[Round "1"]
[White "Kasparov, Garry"]
[Black "Carlsen, Magnus"]
[Result "0-1"]
[WhiteElo "2800"]
[BlackElo "2800"]
1. e4 e5
2. Nf3
The PGN game prompt that GPT-3.5 Turbo Instruct continues each turn. To prime the model to play the best possible move, the headers describe an elite-level game, and declare a win for the side to move. GPT-5.6 Sol and Claude Opus 5 played through their consumer chat apps at medium thinking effort. After one setup message, every message in both directions was a single move in SAN, with no move numbers, no board, and no second chances. An illegal move lost the game.
The 2025 opponents played through the API instead. Each turn, they received the full PGN inside an instruction asking for the next move.
All four of the 2026 games ended between moves 20 and 22, and none by checkmate. Three times, the frontier model tried to capture with a bishop or rook that was already off the board. The fourth time, Opus 5 was in check with exactly two legal moves and chose neither, trying instead to take a defended knight with its king.
The frontier models suffered two failure modes:
forgetting where the pieces are
knowing where they are but playing badly anyway
The 2026 games are mostly the first kind. Three of the four illegal moves use a piece that no longer exists. Opus 5's king move (game 02) is a slightly different mistake: the knight really was on c6, but so was the pawn defending it. The format didn't help. Each move arrives in its own message, with no numbers and no board, so the model has to reassemble the game from a long transcript of terse messages every turn. GPT-3.5 was handed the complete, numbered record.
The 2025 games were decided on the board. Four of the six ended in checkmate. The other two ended in illegal moves from positions already beyond saving. A full PGN every turn didn't stop illegal moves either.
o3 is the most revealing opponent. My PGN2FEN benchmark asks models to convert a PGN move list into the resulting board position, and o3 scored above 90% at every game length. Reading the board was not its problem.
GPT-3.5 Turbo Instruct’s chess ability comes with a major limitation: it only plays well when given carefully formatted PGN . Every other representation I’ve tested yields poor play. Even minor whitespace discrepancies in the PGN break its chess performance.
A likely explanation is its training data. Continuing a realistic game record means predicting moves consistent with the game so far, and in records of strong games, those moves are mostly good ones. Enough of that could produce competent play without anyone setting out to teach chess.
The 2025 games tried the key in newer locks. The frontier models got the same PGN and lost all six, though the key they got was slightly different - GPT-3.5 continued the record, while the others were handed it inside an instruction and asked to pick a move.
There are two ways to read these results. The first is that the newer models can still play well, but chat training means a game record alone no longer draws that skill out. If so, a completion-style prompt should get better chess from them, although chat-only APIs make that hard to test directly.
The second is that the skill itself is weaker. o3 looks like evidence for this: it scored above 90% on PGN2FEN and still lost, so misreading the board wasn't the problem. But PGN2FEN only shows that o3 can rebuild a board when asked, not that it tracks one while choosing moves. A skill that only responds to the right format would produce the same result.
Separating the two means a completion-style test on the frontier models. But GPT-3.5 Turbo Instruct won't be around for the rematch.

## Original Extract

The result says as much about game representation as it does about model capability.

Sign in
Subscribe
Sep 27, 2026
4 min read
LLMs
The 2023 LLM That Beat Frontier Models at Chess
I recently repeated an experiment that I've run a few times over the past of couple years: pitting GPT-3.5 Turbo Instruct against frontier LLMs at chess. GPT-3.5 remains undefeated, winning all four games against GPT-5.6 Sol and Claude Opus 5. Last year , it beat o3, Grok 4 and Gemini 2.5 Pro six games to zero.
Ten games, ten wins. OpenAI is going to shut down access to GPT-3.5 Turbo Instruct tomorrow , so now seems like the time to explore why this 2023 LLM retires undefeated.
GPT-3.5 Turbo Instruct is instruction-tuned, but uses a completion interface - it continues the supplied text directly. For chess, this means supplying GPT-3.5 with an unfinished game represented as PGN , which it then continues with the next move.
[Event "World Championship"]
[Site "Moscow"]
[Date "2025.01.01"]
[Round "1"]
[White "Kasparov, Garry"]
[Black "Carlsen, Magnus"]
[Result "0-1"]
[WhiteElo "2800"]
[BlackElo "2800"]
1. e4 e5
2. Nf3
The PGN game prompt that GPT-3.5 Turbo Instruct continues each turn. To prime the model to play the best possible move, the headers describe an elite-level game, and declare a win for the side to move. GPT-5.6 Sol and Claude Opus 5 played through their consumer chat apps at medium thinking effort. After one setup message, every message in both directions was a single move in SAN, with no move numbers, no board, and no second chances. An illegal move lost the game.
The 2025 opponents played through the API instead. Each turn, they received the full PGN inside an instruction asking for the next move.
All four of the 2026 games ended between moves 20 and 22, and none by checkmate. Three times, the frontier model tried to capture with a bishop or rook that was already off the board. The fourth time, Opus 5 was in check with exactly two legal moves and chose neither, trying instead to take a defended knight with its king.
The frontier models suffered two failure modes:
forgetting where the pieces are
knowing where they are but playing badly anyway
The 2026 games are mostly the first kind. Three of the four illegal moves use a piece that no longer exists. Opus 5's king move (game 02) is a slightly different mistake: the knight really was on c6, but so was the pawn defending it. The format didn't help. Each move arrives in its own message, with no numbers and no board, so the model has to reassemble the game from a long transcript of terse messages every turn. GPT-3.5 was handed the complete, numbered record.
The 2025 games were decided on the board. Four of the six ended in checkmate. The other two ended in illegal moves from positions already beyond saving. A full PGN every turn didn't stop illegal moves either.
o3 is the most revealing opponent. My PGN2FEN benchmark asks models to convert a PGN move list into the resulting board position, and o3 scored above 90% at every game length. Reading the board was not its problem.
GPT-3.5 Turbo Instruct’s chess ability comes with a major limitation: it only plays well when given carefully formatted PGN . Every other representation I’ve tested yields poor play. Even minor whitespace discrepancies in the PGN break its chess performance.
A likely explanation is its training data. Continuing a realistic game record means predicting moves consistent with the game so far, and in records of strong games, those moves are mostly good ones. Enough of that could produce competent play without anyone setting out to teach chess.
The 2025 games tried the key in newer locks. The frontier models got the same PGN and lost all six, though the key they got was slightly different - GPT-3.5 continued the record, while the others were handed it inside an instruction and asked to pick a move.
There are two ways to read these results. The first is that the newer models can still play well, but chat training means a game record alone no longer draws that skill out. If so, a completion-style prompt should get better chess from them, although chat-only APIs make that hard to test directly.
The second is that the skill itself is weaker. o3 looks like evidence for this: it scored above 90% on PGN2FEN and still lost, so misreading the board wasn't the problem. But PGN2FEN only shows that o3 can rebuild a board when asked, not that it tracks one while choosing moves. A skill that only responds to the right format would produce the same result.
Separating the two means a completion-style test on the frontier models. But GPT-3.5 Turbo Instruct won't be around for the rematch.
