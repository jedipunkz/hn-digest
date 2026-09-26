---
source: "https://aatishb.com/blog/2026/ai-jagged-intelligence/"
hn_url: "https://news.ycombinator.com/item?id=49860730"
title: "A Simple Puzzle Reveals the Jagged Edge of AI's Abilities"
article_title: "How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities"
image: "https://aatishb.com/blog/2026/ai-jagged-intelligence/promo.png"
author: "aatish"
captured_at: "2026-09-26T21:53:00Z"
capture_tool: "hn-digest"
hn_id: 49860730
score: 1
comments: 0
posted_at: "2026-09-26T21:31:55Z"
tags:
  - hacker-news
---

# A Simple Puzzle Reveals the Jagged Edge of AI's Abilities

- HN: [49860730](https://news.ycombinator.com/item?id=49860730)
- Source: [aatishb.com](https://aatishb.com/blog/2026/ai-jagged-intelligence/)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T21:31:55Z

## Translation

Title: A Simple Puzzle Reveals the Jagged Edge of AI's Abilities
Article title: How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities
Description: An interactive explanation of how AI can fail in surprising ways.

Article text:
How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities Aatish Bhatia
How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities
It’s undeniable that AI models have gained many powerful abilities. And yet, for all of their advances, they can still make basic blunders.
Researchers call this “ jagged intelligence ,” a term suggesting that there’s a jagged boundary separating what AI can and can’t do.
In this article, I’ll teach a large language model how to solve a simple puzzle, starting from scratch, and we’ll see exactly how this jagged intelligence emerges.
First, try the puzzle yourself:
Click on the tiles below to arrange them from 1 to 15
This is called the 15 puzzle . The goal is to arrange the tiles from 1 to 15 by sliding them around the grid.
Now that you’re familiar with it, let’s teach an AI language model how to solve this puzzle. Large language models learn language by guessing the next word in a sentence. As they encounter more and more sentences, their guesses improve.
Let’s adopt the same strategy. I’ll first create a tiny language to describe the universe of the 15 puzzle. It’ll have a vocabulary of 17 letters and a grammar that consists of valid puzzle solutions. I’ll then train an AI to imitate sentences written in this language.
Then, when the AI predicts the next word, it would really be predicting the next move to solve the puzzle. And if it can do this consistently for any scrambled puzzle that I give it, then you could argue that it has learned how to solve this puzzle.
So let’s get started. Every puzzle begins with a grid of tiles, like this:
I’ll write this out as a sequence of symbols (known as tokens ), along with a blank token to represent the empty square.
Next up, the moves. Every move involves sliding a tile into the empty square. So I’ll simply write down which tile moves at each step.
Just as periods mark the end of sentences, I’ve added a . token to mark that the puzzle is solved.
Putting everything together, here’s a complete puzzle solution written out as a sentence.
How a sentence records a solution
If you showed this sequence of symbols to a friend, they might not recognize that it represents the moves of a puzzle. Yet all of the information needed to replay this solution is encoded in this cryptic sentence.
Just as imitating human language teaches AI the rules of grammar, I’d like to know whether imitating these puzzle-solving sentences will teach it the rules of this game.
First, we need many sentences. To generate them, I’ll choose a million random starting arrangements of the puzzle, and solve each one step by step using a ‘teacher’ algorithm. Writing down every move that this teacher plays will give us a million training sentences.
Here’s some of this training data:
A million solutions to the 15 puzzle
We’re now ready to train the AI. We’ll train a GPT , the AI architecture at the heart of ChatGPT.
Unlike ChatGPT, our GPT won’t need to learn any human languages. It starts off as a blank slate. The only sentences that it encounters during its training are puzzle solutions written in our artificial language.
By prompting this AI with a puzzle board written out as 16 tokens and letting it complete the solution move by move, we can see how it plays.
Before its training, the AI plays moves at random, most of which are physically impossible.
But after training the AI on a million puzzle solutions, it plays like this:
This idea works! The AI can now solve 99.9 percent of randomly scrambled boards. Of the 3,000 puzzles that I tested it on, it played the exact sequence of moves that the teacher would have played in 2,996 of them. (Of the remaining four, it messed up twice, and twice it solved the puzzle using a different route.)
And it didn’t just memorize the answers. For one thing, I’m testing it on puzzle arrangements that it never encountered during its training, so it couldn’t possibly have memorized them.
And there are over ten trillion possible arrangements of this puzzle (that’s 16! / 2 for math nerds). A million is just 0.00001 percent of this number, so memorizing answers won’t get very far.
If a human could play like this, you’d reasonably conclude that they know how to solve the puzzle. You might even say that they understand it.
But here’s the point of this exercise: What if, instead of giving the AI a random puzzle, we give it the easiest possible puzzle — a puzzle that’s just one move away from its solution?
A six-year old could solve this particular puzzle. Surely a GPT trained on a million puzzle solutions can solve this too?
The AI can’t solve the easiest of puzzles, even though it can solve much harder ones. It makes a few moves and then loses the game by trying to move an immovable tile.
There’s an even simpler test: What if we give the AI a puzzle that’s already solved?
This puzzle is already solved. Will the AI recognize that it doesn’t have to do anything and simply end the game?
Once again, it plays a few moves and then loses. Despite its technical skill, it didn’t recognize that the best move was not to play at all.
Whenever the AI plays a move, it’s actually predicting the probability that each token comes next in the sequence (and we pick the most probable one).
We can inspect these probabilities to measure its confidence in any particular move. When we do this, we find that it gave its initial missteps a probability of over 97 percent.
So it wasn’t uncertain. Instead, it was confidently wrong — which is arguably worse.
How is it possible that this AI can solve difficult puzzles but not easy ones?
The answer is that generative AI models emulate the data that they are trained on. And this particular AI was trained to solve random puzzles.
Random puzzles are usually well scrambled
You’ll notice that all of these puzzles are well scrambled. Just as a shuffled deck of cards is unlikely to end up sorted correctly, these tiles are also unlikely to end up in their correct places.
How unlikely? To see one of the boards that stumped the AI above, you’d need to click the ‘get new puzzles’ button above a few hundred billion times. Clicking the button once every second, you’d have to wait eight thousand years.
Here’s a plot that shows all of the 10,461,394,944,000 arrangements of this puzzle, grouped by how many steps it takes to solve them.
Puzzle arrangements, grouped by difficulty
Most randomly selected puzzles fall within the bulk of this distribution , and take at least 30 steps to solve. (In practice, it takes many more steps unless you can play perfectly optimally.)
Conversely, very few randomly chosen puzzles will be 20 steps or fewer away from being solved — the chances of this are about one in three million. Easy puzzles are extreme outliers.
Let’s now look at how the AI performs at different levels of difficulty:
The AI’s accuracy at solving the 15 puzzle
This red curve reveals a jagged boundary in this AI’s abilities. It plays near-perfectly in regions where it has plenty of training data. But its performance plummets on easier puzzles. (In computer science jargon, it fails to generalize ‘out of distribution’.)
Since generative AI learns by imitating data, it makes sense that it does worse on puzzles that are outside of its training data. But here’s what I find most surprising:
First, it fails in the opposite direction from human learners.
The AI can solve the hardest puzzles, but it fails on the easiest ones. This is the opposite direction from how humans learn.
We first have to learn how to solve easy problems before we can tackle harder ones — we learn to walk before we run. Easier problems are our stepping stones towards harder ones.
But AI can skip over our stepping stones, as if learning how to run without knowing how to walk.
Second, it shows how inefficiently AI learns.
Despite having been trained on a million puzzles — more than you might see in a lifetime — this AI wasn’t able to extend its ability to the easiest puzzles.
This reveals an important difference between how humans and AI systems learn. With relatively few examples, we can generalize outside of our ‘training data’. Today’s AI struggles to do this, despite having encountered many times more data than we ever will.
“These models somehow just generalize dramatically worse than people, and it’s super obvious. That seems like a very fundamental thing,” said AI researcher and OpenAI co-founder Ilya Sutskever in a 2025 interview with Dwarkesh Patel.
The AI probably learned how to solve this puzzle not by understanding it as we would, but by using shortcuts — strategies that get the right answer, but for the wrong reasons .
These shortcuts work well in typical situations, which is why it can solve 99.9 percent of randomly scrambled boards. But they can lead the AI astray in unfamiliar situations that are outside of what it has been trained on. (As one group of Microsoft researchers put it: “ Shortcuts are statistically brittle .”)
Here’s how the computer scientist Melanie Mitchell described jagged intelligence in The Yale Review earlier this year:
“In general, today’s AI systems perform extremely well until, often unexpectedly, they don’t. They are inconsistent, lack a sense of when they should be confident or uncertain about their answer, are susceptible to manipulative prompting, and struggle with tasks that differ sufficiently from their training data.”
And this isn’t just about the 15 puzzle. In 2022, researchers found that an AI program that plays the board game Go at a superhuman level could be tricked into losing a game if its opponent played an amateur set of moves outside of the AI’s training data.
As it fell into this trap, the AI remained confident that it would win, even though an average human player could easily spot the trap.
“AI systems that seem to perform at a human level are often doing so in a very alien way, and so can fail in ways that are surprising to humans,” the researcher Adam Gleave told Ars Technica in 2022.
This goes beyond whether AI wins or loses at a board game. It’s also about what it does in unfamiliar situations that can harm human lives. Being confidently wrong can have terrifying consequences.
Here are three lessons for how to think about the jagged boundaries of artificial intelligence:
First, skill isn’t understanding. The parable of the 15 puzzle teaches us that AI can acquire extreme skill without developing the human intuitions that underlie this skill. It can emulate an expert near-perfectly while stumbling on the simplest solutions.
Second, AI can be confidently wrong. It performs well on tasks that resemble its training data, but can fail without warning on tasks that don’t. This is a growing concern as many institutions become increasingly reliant on generative AI.
And finally, it takes shortcuts. These shortcuts make it hard for AI to generalize outside of its training data. (It’s also one reason why AI needs so much training data. Training AI on a lot of diverse data penalizes some of these shortcuts and helps it learn more general patterns.)
Now, it’s clear that today’s ‘reasoning’ models are far beyond this simple GPT. In a following article, I’ll demonstrate how a few clever tricks can help our AI become a better learner.
Aatish Bhatia is a science writer, journalist and physicist. Follow him on Bluesky .

## Original Extract

An interactive explanation of how AI can fail in surprising ways.

How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities Aatish Bhatia
How a Simple Puzzle Reveals the Jagged Edge of AI’s Abilities
It’s undeniable that AI models have gained many powerful abilities. And yet, for all of their advances, they can still make basic blunders.
Researchers call this “ jagged intelligence ,” a term suggesting that there’s a jagged boundary separating what AI can and can’t do.
In this article, I’ll teach a large language model how to solve a simple puzzle, starting from scratch, and we’ll see exactly how this jagged intelligence emerges.
First, try the puzzle yourself:
Click on the tiles below to arrange them from 1 to 15
This is called the 15 puzzle . The goal is to arrange the tiles from 1 to 15 by sliding them around the grid.
Now that you’re familiar with it, let’s teach an AI language model how to solve this puzzle. Large language models learn language by guessing the next word in a sentence. As they encounter more and more sentences, their guesses improve.
Let’s adopt the same strategy. I’ll first create a tiny language to describe the universe of the 15 puzzle. It’ll have a vocabulary of 17 letters and a grammar that consists of valid puzzle solutions. I’ll then train an AI to imitate sentences written in this language.
Then, when the AI predicts the next word, it would really be predicting the next move to solve the puzzle. And if it can do this consistently for any scrambled puzzle that I give it, then you could argue that it has learned how to solve this puzzle.
So let’s get started. Every puzzle begins with a grid of tiles, like this:
I’ll write this out as a sequence of symbols (known as tokens ), along with a blank token to represent the empty square.
Next up, the moves. Every move involves sliding a tile into the empty square. So I’ll simply write down which tile moves at each step.
Just as periods mark the end of sentences, I’ve added a . token to mark that the puzzle is solved.
Putting everything together, here’s a complete puzzle solution written out as a sentence.
How a sentence records a solution
If you showed this sequence of symbols to a friend, they might not recognize that it represents the moves of a puzzle. Yet all of the information needed to replay this solution is encoded in this cryptic sentence.
Just as imitating human language teaches AI the rules of grammar, I’d like to know whether imitating these puzzle-solving sentences will teach it the rules of this game.
First, we need many sentences. To generate them, I’ll choose a million random starting arrangements of the puzzle, and solve each one step by step using a ‘teacher’ algorithm. Writing down every move that this teacher plays will give us a million training sentences.
Here’s some of this training data:
A million solutions to the 15 puzzle
We’re now ready to train the AI. We’ll train a GPT , the AI architecture at the heart of ChatGPT.
Unlike ChatGPT, our GPT won’t need to learn any human languages. It starts off as a blank slate. The only sentences that it encounters during its training are puzzle solutions written in our artificial language.
By prompting this AI with a puzzle board written out as 16 tokens and letting it complete the solution move by move, we can see how it plays.
Before its training, the AI plays moves at random, most of which are physically impossible.
But after training the AI on a million puzzle solutions, it plays like this:
This idea works! The AI can now solve 99.9 percent of randomly scrambled boards. Of the 3,000 puzzles that I tested it on, it played the exact sequence of moves that the teacher would have played in 2,996 of them. (Of the remaining four, it messed up twice, and twice it solved the puzzle using a different route.)
And it didn’t just memorize the answers. For one thing, I’m testing it on puzzle arrangements that it never encountered during its training, so it couldn’t possibly have memorized them.
And there are over ten trillion possible arrangements of this puzzle (that’s 16! / 2 for math nerds). A million is just 0.00001 percent of this number, so memorizing answers won’t get very far.
If a human could play like this, you’d reasonably conclude that they know how to solve the puzzle. You might even say that they understand it.
But here’s the point of this exercise: What if, instead of giving the AI a random puzzle, we give it the easiest possible puzzle — a puzzle that’s just one move away from its solution?
A six-year old could solve this particular puzzle. Surely a GPT trained on a million puzzle solutions can solve this too?
The AI can’t solve the easiest of puzzles, even though it can solve much harder ones. It makes a few moves and then loses the game by trying to move an immovable tile.
There’s an even simpler test: What if we give the AI a puzzle that’s already solved?
This puzzle is already solved. Will the AI recognize that it doesn’t have to do anything and simply end the game?
Once again, it plays a few moves and then loses. Despite its technical skill, it didn’t recognize that the best move was not to play at all.
Whenever the AI plays a move, it’s actually predicting the probability that each token comes next in the sequence (and we pick the most probable one).
We can inspect these probabilities to measure its confidence in any particular move. When we do this, we find that it gave its initial missteps a probability of over 97 percent.
So it wasn’t uncertain. Instead, it was confidently wrong — which is arguably worse.
How is it possible that this AI can solve difficult puzzles but not easy ones?
The answer is that generative AI models emulate the data that they are trained on. And this particular AI was trained to solve random puzzles.
Random puzzles are usually well scrambled
You’ll notice that all of these puzzles are well scrambled. Just as a shuffled deck of cards is unlikely to end up sorted correctly, these tiles are also unlikely to end up in their correct places.
How unlikely? To see one of the boards that stumped the AI above, you’d need to click the ‘get new puzzles’ button above a few hundred billion times. Clicking the button once every second, you’d have to wait eight thousand years.
Here’s a plot that shows all of the 10,461,394,944,000 arrangements of this puzzle, grouped by how many steps it takes to solve them.
Puzzle arrangements, grouped by difficulty
Most randomly selected puzzles fall within the bulk of this distribution , and take at least 30 steps to solve. (In practice, it takes many more steps unless you can play perfectly optimally.)
Conversely, very few randomly chosen puzzles will be 20 steps or fewer away from being solved — the chances of this are about one in three million. Easy puzzles are extreme outliers.
Let’s now look at how the AI performs at different levels of difficulty:
The AI’s accuracy at solving the 15 puzzle
This red curve reveals a jagged boundary in this AI’s abilities. It plays near-perfectly in regions where it has plenty of training data. But its performance plummets on easier puzzles. (In computer science jargon, it fails to generalize ‘out of distribution’.)
Since generative AI learns by imitating data, it makes sense that it does worse on puzzles that are outside of its training data. But here’s what I find most surprising:
First, it fails in the opposite direction from human learners.
The AI can solve the hardest puzzles, but it fails on the easiest ones. This is the opposite direction from how humans learn.
We first have to learn how to solve easy problems before we can tackle harder ones — we learn to walk before we run. Easier problems are our stepping stones towards harder ones.
But AI can skip over our stepping stones, as if learning how to run without knowing how to walk.
Second, it shows how inefficiently AI learns.
Despite having been trained on a million puzzles — more than you might see in a lifetime — this AI wasn’t able to extend its ability to the easiest puzzles.
This reveals an important difference between how humans and AI systems learn. With relatively few examples, we can generalize outside of our ‘training data’. Today’s AI struggles to do this, despite having encountered many times more data than we ever will.
“These models somehow just generalize dramatically worse than people, and it’s super obvious. That seems like a very fundamental thing,” said AI researcher and OpenAI co-founder Ilya Sutskever in a 2025 interview with Dwarkesh Patel.
The AI probably learned how to solve this puzzle not by understanding it as we would, but by using shortcuts — strategies that get the right answer, but for the wrong reasons .
These shortcuts work well in typical situations, which is why it can solve 99.9 percent of randomly scrambled boards. But they can lead the AI astray in unfamiliar situations that are outside of what it has been trained on. (As one group of Microsoft researchers put it: “ Shortcuts are statistically brittle .”)
Here’s how the computer scientist Melanie Mitchell described jagged intelligence in The Yale Review earlier this year:
“In general, today’s AI systems perform extremely well until, often unexpectedly, they don’t. They are inconsistent, lack a sense of when they should be confident or uncertain about their answer, are susceptible to manipulative prompting, and struggle with tasks that differ sufficiently from their training data.”
And this isn’t just about the 15 puzzle. In 2022, researchers found that an AI program that plays the board game Go at a superhuman level could be tricked into losing a game if its opponent played an amateur set of moves outside of the AI’s training data.
As it fell into this trap, the AI remained confident that it would win, even though an average human player could easily spot the trap.
“AI systems that seem to perform at a human level are often doing so in a very alien way, and so can fail in ways that are surprising to humans,” the researcher Adam Gleave told Ars Technica in 2022.
This goes beyond whether AI wins or loses at a board game. It’s also about what it does in unfamiliar situations that can harm human lives. Being confidently wrong can have terrifying consequences.
Here are three lessons for how to think about the jagged boundaries of artificial intelligence:
First, skill isn’t understanding. The parable of the 15 puzzle teaches us that AI can acquire extreme skill without developing the human intuitions that underlie this skill. It can emulate an expert near-perfectly while stumbling on the simplest solutions.
Second, AI can be confidently wrong. It performs well on tasks that resemble its training data, but can fail without warning on tasks that don’t. This is a growing concern as many institutions become increasingly reliant on generative AI.
And finally, it takes shortcuts. These shortcuts make it hard for AI to generalize outside of its training data. (It’s also one reason why AI needs so much training data. Training AI on a lot of diverse data penalizes some of these shortcuts and helps it learn more general patterns.)
Now, it’s clear that today’s ‘reasoning’ models are far beyond this simple GPT. In a following article, I’ll demonstrate how a few clever tricks can help our AI become a better learner.
Aatish Bhatia is a science writer, journalist and physicist. Follow him on Bluesky .
