---
source: "https://fazamhd.com/mental-models/llm/"
hn_url: "https://news.ycombinator.com/item?id=49664371"
title: "How LLMs work, with interactive simulations"
article_title: "Large Language Models | Faza"
image: "https://fazamhd.com/og/mental-models/llm.png"
author: "faza"
captured_at: "2026-09-11T20:33:45Z"
capture_tool: "hn-digest"
hn_id: 49664371
score: 1
comments: 0
posted_at: "2026-09-11T19:50:15Z"
tags:
  - hacker-news
---

# How LLMs work, with interactive simulations

- HN: [49664371](https://news.ycombinator.com/item?id=49664371)
- Source: [fazamhd.com](https://fazamhd.com/mental-models/llm/)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T19:50:15Z

## Translation

Title: How LLMs work, with interactive simulations
Article title: Large Language Models | Faza
Description: How predicting the next token across vast amounts of text turns statistical patterns into conversation, reasoning, and useful behavior.

Article text:
Large Language Models | Faza Skip to content Faza models labs work photos about ← All mental models Large Language Models
How predicting the next token across vast amounts of text turns statistical patterns into conversation, reasoning, and useful behavior.
20 sections 111 min core reading 25 optional deep dives (+23 min) Read
Questions we'll explore in this article.
01 How can predicting the next token produce reasoning and fluent conversation? Deconstruct an AI assistant into three distinct layers, trace the complete generation pipeline, and see why predicting varied text rewards useful representations of the world.
02 What happens inside a transformer during a single forward pass? Follow text through subword tokenization, high-dimensional vector embeddings, query-key-value self-attention, and deep residual transformer layers.
03 How are models trained, understood internally, and turned into assistants? Explore massive pre-training on next-token loss, the internal circuits discovered by mechanistic interpretability, post-training with RL and verifiable rewards, and in-context learning.
04 How do models run in production and what dictates serving costs? Understand benchmark evaluation and calibration, prefill vs decode stages, sampling controls, KV caching, memory bandwidth bottlenecks, speculative decoding, and quantization.
05 How are modern architectures scaling beyond dense text decoders? See how Mixture of Experts (MoE) decouples active compute from stored capacity, how hybrid attention compresses memory, and how vision patches integrate into the residual stream.
06 How does surrounding software connect models to tools and application context? Understand function calling protocols, error compounding in agents, RAG knowledge retrieval, prompt injection vulnerabilities, and multi-layered guardrails.
07 Why do models fail, what boundaries remain open, and what mental model should you keep? Diagnose hallucinations and failure modes across layers, survey open research frontiers, and synthesize the predictor, loop, and product into an engineering mental model.
Today, we take for granted what a single prompt can do. We use it to draft an email or ask for a recipe, but just as easily expect it to write complex software, prove mathematical theorems, or distill dozens of clinical trials. In just a few years, the conversation has moved from laughing at early hallucinations to debating whether frontier models are conscious.
For decades, artificial intelligence meant systems restricted to a single domain, like Deep Blue in chess, Watson on Jeopardy! , or AlphaGo in Go. A system that mastered chess was useless for anything else. ChatGPT showed the world, for the first time, that general capability could emerge from something as simple as predicting the next word.
We have long trusted software to route aircraft, predict weather, and trade markets without ever imagining a mind behind the code. A language model feels completely different. It talks back, adapts to our tone, and handles an astonishing variety of tasks, making it almost impossible not to treat it like a highly intelligent person.
But that fluency hides a bizarre contradiction, the same system that reasons through hard problems can struggle with the basics, or make bafflingly simple mistakes. These models are very much a work in progress, but they are already remarkably powerful. Building an accurate mental model clears up common misunderstandings, saves us endless frustration, and helps us use them far more effectively.
Try finishing each of these three lines in your head:
Two days after Monday is ? Wednesday
The first comes almost instantly, without thinking. The second requires counting two days ahead. The third demands an understanding of programming. You might see these as completely different kinds of thinking, some pure recollection, some logic, some code. But to a language model, all three are simply next-word prediction, even if arriving at each one requires fundamentally different computations.
A language model doesn’t stop after picking a single word. It appends what it just chose to the text, feeds the entire sequence back into itself, and predicts what comes next, building the response one piece at a time.
That simple loop is behind every response an AI has ever generated. We will use this exact sentence to trace the entire journey, from raw characters and vectors to attention, training, and reasoning.
How does the simple goal of predicting the next word produce such impressive abilities?
That is what I want to help you understand in this article. Follow the main text for the core mental model, or open the optional deep dives along the way for more details.
What is the actual system behind these AI models?
How can tools like ChatGPT, Claude, or any other AI assistant perform such a wide variety of tasks, seeming to understand our requests and respond with deep knowledge in seconds?
For the autoregressive assistants explored here, a useful model is three nested layers :
The Learned Predictor: At the center sits a large mathematical function with billions of parameters, the kind of network explored in the deep learning article. It is completely stateless , possessing no memory between runs and no independent agency. A single forward pass simply computes next-token probability scores for the vocabulary, and the system uses the probabilities at the final position to predict what comes next.
The Generation Loop: A program around the model runs the loop. It takes your prompt, runs the model once, samples a token from those probabilities, appends that new token to the sequence, and feeds the expanded sequence back in. Repeating this loop dozens of times per second is what creates the stream of text flowing across your screen.
The Product System: The application you actually interact with, whether a web UI (ChatGPT, Claude), an agentic harness (Claude Code, Codex), or an IDE integration (Cursor, Antigravity). The product manages your conversation history, applies content filters, fetches search results or documents, and executes tools.
expand close ↔ Swipe sideways to inspect
product system generation loop search results earlier messages tool result prompt every token, every call learned predictor fixed weights next token append it, call the model again tool call permissions content filters output checks reply Only the innermost box is the trained neural network, running once per token with no memory of prior calls. Everything else is traditional software. The model merely emits text requesting an action, and the outer product checks permissions, runs the tool, and feeds the output back into the context as more tokens.
Understanding this three-layer architecture is important, most confusion about AI comes from attributing to the inner model what the surrounding software is actually doing.
The model does not “remember” what you said three messages ago, the outer software simply pasted earlier messages back into the context window. The model does not “browse the internet” or “edit files”, it just generates tokens requesting a search or an action, and the software parses these action tokens, checks permissions, executes the tool, and pastes the result back in as more tokens.
Once you understand that the core engine is simply a next-token predictor, generating text piece by piece (where a token is a word, number, subword, or punctuation mark) and feeding it back into itself (a process called autoregression ), the fundamental question becomes, how does predicting the next token produce reasoning, conversation, and useful behavior?
The End-to-End Generation Pipeline
Before diving into the individual components, here is a high-level overview of how a single token is computed through the complete loop:
Tokenize: A tokenizer converts the input text into a sequence of integer token IDs.
Embed: Each ID retrieves a learned starting vector. Position information is incorporated into the computation. Some architectures add position vectors, while others rotate Queries and Keys.
Transform: Repeated attention and feed-forward blocks update every token’s vector using the preceding context.
Read out: The final vector at the last position produces an unnormalized score ( logit ) for every token in the vocabulary. Softmax converts these scores into probabilities.
Sample: The host program selects one token from the probability distribution based on its sampling settings.
Append: That selected token is appended to the context sequence.
Repeat: The model runs another forward pass for the next token, conditioned on the updated sequence.
tokens = tokenize(prompt) # Step 1: convert text into token IDs
while len (tokens) < limit:
logits = model(tokens)[ - 1 ] # Steps 2–4: compute scores at the final position
token = sample(logits) # Step 5: pick one token (temperature, top-p, etc.)
if token in STOP_TOKENS : # stop if end-of-sequence token is chosen
break
tokens.append(token) # Steps 6–7: append token to feed back into the next pass
return detokenize(tokens) # convert token IDs back into text
The tokenizer is external to the learned network. The model performs steps 2 through 4, embedding lookup, transformer computation, and readout into vocabulary scores. The surrounding generation loop controls token selection, repetition, and termination. This architectural boundary explains why changing temperature does not alter the model weights, why tool execution requires external software, and why an early error cascades into subsequent generations.
The interactive simulation below traces this complete pipeline end-to-end. Treat it as a visual map, the rest of the article breaks down each component in detail.
expand close ↔ Swipe sideways to inspect
output logits final state → scores
decode & append choose → repeat
49 – 95 96 embedding The 464 chemical 5931 symbol 6194 for 329 potassium 33156 is 318 one score per vocabulary entry softmax → one distribution K + 4.0 0 % P + 1.0 0 % Na + 0.7 0 % Ka + 0.3 0 % Pot -0.4 0 % k -0.4 0 % 50,257 Text becomes the model's alphabet. A tokenizer converts the prompt into integer token IDs from a fixed vocabulary. Raw characters never enter the neural network directly.
Each token ID retrieves a learned vector representation from an embedding table. Positional information is combined with this vector so the model knows sequence order.
All token positions pass through each layer together. Within each layer, attention gathers context from preceding tokens, and feed-forward networks process each position individually, both adding updates into the residual stream.
While every position computes a final state, generation inspects only the last one. That state produces an unnormalized score (logit) for every vocabulary token, and softmax turns them into probabilities.
The top candidate (the token 'K') is selected and appended to the input sequence. The loop runs again, computing only the newest token position while reading cached context from earlier positions.
This simulation depicts a GPT-3 sized baseline (96 layers over a 50,257 entry vocabulary). While modern models feature larger vocabularies and varied depths, the core execution mechanism remains identical. Token IDs and embedding rows are drawn from GPT-2's open weights (which shares GPT-3's vocabulary); internal activation states and attention weights are simplified for illustration.
Stage 1 of 5 : tokenize . A tokenizer converts the prompt into integer token IDs from a fixed vocabulary. Raw characters never enter the neural network directly.
That gives us the big picture view of the system. But it naturally raises the core question, why does predicting the next token produce outputs that can write code, solve mathematical proofs, and reason about the world?
The Predictability of Language
The foundational idea behind modern language models was published decades before they became practical, and it was originally aimed at a completely different problem.
In 1948, mat

[truncated]

## Original Extract

How predicting the next token across vast amounts of text turns statistical patterns into conversation, reasoning, and useful behavior.

Large Language Models | Faza Skip to content Faza models labs work photos about ← All mental models Large Language Models
How predicting the next token across vast amounts of text turns statistical patterns into conversation, reasoning, and useful behavior.
20 sections 111 min core reading 25 optional deep dives (+23 min) Read
Questions we'll explore in this article.
01 How can predicting the next token produce reasoning and fluent conversation? Deconstruct an AI assistant into three distinct layers, trace the complete generation pipeline, and see why predicting varied text rewards useful representations of the world.
02 What happens inside a transformer during a single forward pass? Follow text through subword tokenization, high-dimensional vector embeddings, query-key-value self-attention, and deep residual transformer layers.
03 How are models trained, understood internally, and turned into assistants? Explore massive pre-training on next-token loss, the internal circuits discovered by mechanistic interpretability, post-training with RL and verifiable rewards, and in-context learning.
04 How do models run in production and what dictates serving costs? Understand benchmark evaluation and calibration, prefill vs decode stages, sampling controls, KV caching, memory bandwidth bottlenecks, speculative decoding, and quantization.
05 How are modern architectures scaling beyond dense text decoders? See how Mixture of Experts (MoE) decouples active compute from stored capacity, how hybrid attention compresses memory, and how vision patches integrate into the residual stream.
06 How does surrounding software connect models to tools and application context? Understand function calling protocols, error compounding in agents, RAG knowledge retrieval, prompt injection vulnerabilities, and multi-layered guardrails.
07 Why do models fail, what boundaries remain open, and what mental model should you keep? Diagnose hallucinations and failure modes across layers, survey open research frontiers, and synthesize the predictor, loop, and product into an engineering mental model.
Today, we take for granted what a single prompt can do. We use it to draft an email or ask for a recipe, but just as easily expect it to write complex software, prove mathematical theorems, or distill dozens of clinical trials. In just a few years, the conversation has moved from laughing at early hallucinations to debating whether frontier models are conscious.
For decades, artificial intelligence meant systems restricted to a single domain, like Deep Blue in chess, Watson on Jeopardy! , or AlphaGo in Go. A system that mastered chess was useless for anything else. ChatGPT showed the world, for the first time, that general capability could emerge from something as simple as predicting the next word.
We have long trusted software to route aircraft, predict weather, and trade markets without ever imagining a mind behind the code. A language model feels completely different. It talks back, adapts to our tone, and handles an astonishing variety of tasks, making it almost impossible not to treat it like a highly intelligent person.
But that fluency hides a bizarre contradiction, the same system that reasons through hard problems can struggle with the basics, or make bafflingly simple mistakes. These models are very much a work in progress, but they are already remarkably powerful. Building an accurate mental model clears up common misunderstandings, saves us endless frustration, and helps us use them far more effectively.
Try finishing each of these three lines in your head:
Two days after Monday is ? Wednesday
The first comes almost instantly, without thinking. The second requires counting two days ahead. The third demands an understanding of programming. You might see these as completely different kinds of thinking, some pure recollection, some logic, some code. But to a language model, all three are simply next-word prediction, even if arriving at each one requires fundamentally different computations.
A language model doesn’t stop after picking a single word. It appends what it just chose to the text, feeds the entire sequence back into itself, and predicts what comes next, building the response one piece at a time.
That simple loop is behind every response an AI has ever generated. We will use this exact sentence to trace the entire journey, from raw characters and vectors to attention, training, and reasoning.
How does the simple goal of predicting the next word produce such impressive abilities?
That is what I want to help you understand in this article. Follow the main text for the core mental model, or open the optional deep dives along the way for more details.
What is the actual system behind these AI models?
How can tools like ChatGPT, Claude, or any other AI assistant perform such a wide variety of tasks, seeming to understand our requests and respond with deep knowledge in seconds?
For the autoregressive assistants explored here, a useful model is three nested layers :
The Learned Predictor: At the center sits a large mathematical function with billions of parameters, the kind of network explored in the deep learning article. It is completely stateless , possessing no memory between runs and no independent agency. A single forward pass simply computes next-token probability scores for the vocabulary, and the system uses the probabilities at the final position to predict what comes next.
The Generation Loop: A program around the model runs the loop. It takes your prompt, runs the model once, samples a token from those probabilities, appends that new token to the sequence, and feeds the expanded sequence back in. Repeating this loop dozens of times per second is what creates the stream of text flowing across your screen.
The Product System: The application you actually interact with, whether a web UI (ChatGPT, Claude), an agentic harness (Claude Code, Codex), or an IDE integration (Cursor, Antigravity). The product manages your conversation history, applies content filters, fetches search results or documents, and executes tools.
expand close ↔ Swipe sideways to inspect
product system generation loop search results earlier messages tool result prompt every token, every call learned predictor fixed weights next token append it, call the model again tool call permissions content filters output checks reply Only the innermost box is the trained neural network, running once per token with no memory of prior calls. Everything else is traditional software. The model merely emits text requesting an action, and the outer product checks permissions, runs the tool, and feeds the output back into the context as more tokens.
Understanding this three-layer architecture is important, most confusion about AI comes from attributing to the inner model what the surrounding software is actually doing.
The model does not “remember” what you said three messages ago, the outer software simply pasted earlier messages back into the context window. The model does not “browse the internet” or “edit files”, it just generates tokens requesting a search or an action, and the software parses these action tokens, checks permissions, executes the tool, and pastes the result back in as more tokens.
Once you understand that the core engine is simply a next-token predictor, generating text piece by piece (where a token is a word, number, subword, or punctuation mark) and feeding it back into itself (a process called autoregression ), the fundamental question becomes, how does predicting the next token produce reasoning, conversation, and useful behavior?
The End-to-End Generation Pipeline
Before diving into the individual components, here is a high-level overview of how a single token is computed through the complete loop:
Tokenize: A tokenizer converts the input text into a sequence of integer token IDs.
Embed: Each ID retrieves a learned starting vector. Position information is incorporated into the computation. Some architectures add position vectors, while others rotate Queries and Keys.
Transform: Repeated attention and feed-forward blocks update every token’s vector using the preceding context.
Read out: The final vector at the last position produces an unnormalized score ( logit ) for every token in the vocabulary. Softmax converts these scores into probabilities.
Sample: The host program selects one token from the probability distribution based on its sampling settings.
Append: That selected token is appended to the context sequence.
Repeat: The model runs another forward pass for the next token, conditioned on the updated sequence.
tokens = tokenize(prompt) # Step 1: convert text into token IDs
while len (tokens) < limit:
logits = model(tokens)[ - 1 ] # Steps 2–4: compute scores at the final position
token = sample(logits) # Step 5: pick one token (temperature, top-p, etc.)
if token in STOP_TOKENS : # stop if end-of-sequence token is chosen
break
tokens.append(token) # Steps 6–7: append token to feed back into the next pass
return detokenize(tokens) # convert token IDs back into text
The tokenizer is external to the learned network. The model performs steps 2 through 4, embedding lookup, transformer computation, and readout into vocabulary scores. The surrounding generation loop controls token selection, repetition, and termination. This architectural boundary explains why changing temperature does not alter the model weights, why tool execution requires external software, and why an early error cascades into subsequent generations.
The interactive simulation below traces this complete pipeline end-to-end. Treat it as a visual map, the rest of the article breaks down each component in detail.
expand close ↔ Swipe sideways to inspect
output logits final state → scores
decode & append choose → repeat
49 – 95 96 embedding The 464 chemical 5931 symbol 6194 for 329 potassium 33156 is 318 one score per vocabulary entry softmax → one distribution K + 4.0 0 % P + 1.0 0 % Na + 0.7 0 % Ka + 0.3 0 % Pot -0.4 0 % k -0.4 0 % 50,257 Text becomes the model's alphabet. A tokenizer converts the prompt into integer token IDs from a fixed vocabulary. Raw characters never enter the neural network directly.
Each token ID retrieves a learned vector representation from an embedding table. Positional information is combined with this vector so the model knows sequence order.
All token positions pass through each layer together. Within each layer, attention gathers context from preceding tokens, and feed-forward networks process each position individually, both adding updates into the residual stream.
While every position computes a final state, generation inspects only the last one. That state produces an unnormalized score (logit) for every vocabulary token, and softmax turns them into probabilities.
The top candidate (the token 'K') is selected and appended to the input sequence. The loop runs again, computing only the newest token position while reading cached context from earlier positions.
This simulation depicts a GPT-3 sized baseline (96 layers over a 50,257 entry vocabulary). While modern models feature larger vocabularies and varied depths, the core execution mechanism remains identical. Token IDs and embedding rows are drawn from GPT-2's open weights (which shares GPT-3's vocabulary); internal activation states and attention weights are simplified for illustration.
Stage 1 of 5 : tokenize . A tokenizer converts the prompt into integer token IDs from a fixed vocabulary. Raw characters never enter the neural network directly.
That gives us the big picture view of the system. But it naturally raises the core question, why does predicting the next token produce outputs that can write code, solve mathematical proofs, and reason about the world?
The Predictability of Language
The foundational idea behind modern language models was published decades before they became practical, and it was originally aimed at a completely different problem.
In 1948, mat

[truncated]
