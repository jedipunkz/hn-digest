---
source: "https://www.amazon.science/blog/can-we-change-an-ai-models-architecture-without-retraining-it-from-scratch"
hn_url: "https://news.ycombinator.com/item?id=50075655"
title: "Can we change an AI model's architecture without retraining it from scratch?"
article_title: "Can we change an AI model’s architecture without retraining it from scratch? - Amazon Science"
image: "https://cdn.amazon.science/dims4/default/34de22a/2147483647/strip/true/crop/2270x1192+0+43/resize/1200x630!/quality/90/?url=https%3A%2F%2Famzn-science-production-science.s3.us-east-1.amazonaws.com%2Fscience%2F48%2Ffc%2Fcf86b92b484f9da3927d4c06f537%2Fstate-composition-16x9.png"
author: "gmays"
captured_at: "2026-10-13T17:58:04Z"
capture_tool: "hn-digest"
hn_id: 50075655
score: 1
comments: 0
posted_at: "2026-10-13T17:42:21Z"
tags:
  - hacker-news
---

# Can we change an AI model's architecture without retraining it from scratch?

- HN: [50075655](https://news.ycombinator.com/item?id=50075655)
- Source: [www.amazon.science](https://www.amazon.science/blog/can-we-change-an-ai-models-architecture-without-retraining-it-from-scratch)
- Score: 1
- Comments: 0
- Posted: 2026-10-13T17:42:21Z

## Translation

Title: Can we change an AI model's architecture without retraining it from scratch?
Article title: Can we change an AI model’s architecture without retraining it from scratch? - Amazon Science
Description: “Priming” reuses pretrained Transformer weights to build models with more efficient architectures at less than 1% of the original training cost and with throughputs up to twice as high.

Article text:
Can we change an AI model’s architecture without retraining it from scratch? - Amazon Science
Close
Close
Research
Research
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Nova
Try Amazon’s frontier foundation models.
Amazon Nova
Try Amazon’s frontier foundation models.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Research
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
News & blog
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
Collaborations
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Nova
Try Amazon’s frontier foundation models.
Amazon Nova
Try Amazon’s frontier foundation models.
Careers
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Can we change an AI model’s architecture without retraining it from scratch?
“Priming” reuses pretrained Transformer weights to build models with more efficient architectures at less than 1% of the original training cost and with throughputs up to twice as high.
A new AWS paper introduces priming, a method that reduces pretraining costs for hybrid architectures by over 99% by leveraging pretrained Transformer weights, saving trillions of tokens and hundreds of thousands of GPU-hours.
A 32-billion-parameter Gated KalmaNet (GKA)-based hybrid model matches its source Transformer on reasoning benchmarks while decoding up to twice as quickly with 128,000 tokens of context.
State space models consistently outperform sliding-window attention (SWS) even when SWA receives twice the memory budget, and GKA's variable-iteration mechanism enables 10% faster decoding with only ~1% accuracy loss.
AI agents are rapidly scaling: longer sessions, deeper reasoning chains, more autonomy. But the Transformer architecture underneath them still pays a memory tax that grows with the number of tokens it processes. As agents become more capable, that cost scales with them.
Now, there’s a proliferation of hybrid architectures that use some combination of methods to achieve a better balance of memory, speed, and quality. Those methods include sparse attention, which selects a subset of past tokens to consider when selecting the next token; sliding-window attention, which preserves only the most recent tokens; and state space models, which compress information from the full history into a fixed-size state.
These hybrid architectures can be understood as points in a sequence-model design space: to differing degrees, they mix information while traversing a sequence and passing through different states. What distinguishes them from each other is how they retain and use the past.
Until now, exploring this design space has required training every new architecture from scratch, using trillions of tokens and hundreds of thousands of GPU-hours. This has made controlled comparisons of hybrid architectures extremely expensive.
In a new paper, “ Priming: Hybrid state space models from pre-trained Transformers ,” our team at Amazon Web Services addresses this problem. We introduce priming, a method that leverages pretrained Transformer weights as a starting point to reduce the pretraining cost of hybrid architectures by more than 99%.
With priming, model developers can ask questions that were previously too expensive to answer at scale: Which attention layers can be replaced by more-efficient layers? How many attention layers can be replaced without sacrificing quality? Which state space layer works best? And what does the trade-off between speed and accuracy look like?
In our paper, we compare hybrid models that use three different types of state space models (Gated KalmaNet, Gated DeltaNet, and Mamba-2). Some of our findings:
Not all state space models are equal. Under controlled conditions, Gated KalmaNet (GKA) performs best, followed by Gated DeltaNet (GDN) and then Mamba-2. GKA uniquely offers variable-length inference computation without retraining and gives the user direct control over its position on the speed-accuracy frontier.
A 32-billion-parameter hybrid model reaches each accuracy threshold faster than an equivalently sized Transformer, converging to similar final accuracy in less wall clock time. The dataset is the AIME 2025 math competition benchmark (hard subset). Accuracy at time t is the fraction of rollouts for which a correct final answer has been produced by time t .
A good lossy summary of everything beats a perfect recording of the most recent past. State space models consistently outperform sliding-window attention (SWA), even when SWA gets twice the memory budget.
Same quality, twice the throughput. A 32-billion-parameter GKA-based hybrid model matches its source Transformer on reasoning benchmarks while decoding up to twice as quickly with 128,000 tokens of context and reaching correct answers in as little as two-thirds of the wall-clock time.
To foster further research on hybrid architectures, we have released a zoo of primed hybrid models at the eight-billion-parameter and 32-billion-parameter scales for long-context reasoning and instruction following, together with training code, optimized kernels, and a vLLM-serving plugin under the Apache 2.0 License.
A Transformer’s attention mechanism — which determines how relevant previously seen tokens are to the token currently being processed — relies on three different vectors, which take their names from database design: the query, the key, and the value. For each token, the query captures what that token is “looking for,” the key captures what it “offers,” and the value carries the actual information it will pass along. Relevance is computed by comparing the current token’s query against every other token’s key; the closer the match, the more that token’s value contributes to the output of attention.
Each of these vectors is produced by multiplying the token by a learned weight projection matrix. The Transformer outputs one token at a time, and each output token becomes a new input, so the context grows one token at a time as well. To save computation, the Transformer thus stores the key and value projections of each new input in a “KV cache”.
Two kinds of memory layers: Eidetic versus fading
The attention mechanism gives Transformers what we call eidetic memory : verbatim recall of every past token. The possibility of recalling over an LLM’s entire context window is what made in-context learning, long-document reasoning, and retrieval-augmented generation possible. But it comes at a cost that scales with the number of input tokens.
However, not all attention layers do the same job. Some perform long-range retrieval across tens of thousands of tokens. Others operate more locally, tracking the current clause, copying nearby tokens, and the like. The local ones don't need a growing cache: a compact, fixed-size summary of the past — what we call fading memory — suffices.
Hybrid models mix both: attention layers (eidetic, growing cache) for the layers that need global recall and state space layers (fading, fixed size) for the rest. The result is a model with a reduced memory footprint that still has access to verbatim recall where it matters.
Priming: Knowledge isn't locked to the architecture
Exploring hybrids’ design space has been impractical because each hybrid configuration would normally require pretraining from scratch. Priming sidesteps this by reusing the stored knowledge in a pretrained Transformer's weights.
Priming is possible because a Transformer is itself a special case of a state space realization . The outputs of both attention and SSM layers can be expressed as the result of multiplying the value sequence by a “mixing” matrix that decides how much each past token contributes to each current token. Attention computes the mixing matrix by multiplying the query and key matrices; an SSM produces a similar output through a recurrent state of fixed dimension that all past information must pass through.
Priming takes advantage of this structural similarity: if an SSM with a fixed-dimensional state can produce a mixing matrix close enough to the attention layer's, the query, key, and value projections from the source Transformer can be reused as the SSM's own input projections, with only the SSM-specific parameters needing to be learned.
Whether the approximation holds depends on whether the attention layer performs complex, long-range retrievals or simple, local operations. We distinguish between the two using a cheap empirical proxy: repla

[truncated]

## Original Extract

“Priming” reuses pretrained Transformer weights to build models with more efficient architectures at less than 1% of the original training cost and with throughputs up to twice as high.

Can we change an AI model’s architecture without retraining it from scratch? - Amazon Science
Close
Close
Research
Research
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Nova
Try Amazon’s frontier foundation models.
Amazon Nova
Try Amazon’s frontier foundation models.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Research
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
Research areas
Automated reasoning
Information and knowledge management
Operations research and optimization
Search and information retrieval
Security, privacy, and abuse prevention
Our scientific contributions
Publications
Research from our scientists and collaborators.
Conferences
Our experts present and discuss cutting-edge research at scientific meetings globally.
News & blog
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
The latest from Amazon researchers
Amazon Science Blog
Technical deep-dives and perspectives from our scientists.
News
Research milestones and recent achievements.
Collaborations
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Research Awards
Overview
Amazon Nova AI Challenge
Overview
Research collaborations
Overview
University of California, Los Angeles
University of Illinois Urbana-Champaign
University of Southern California
Amazon Nova
Try Amazon’s frontier foundation models.
Amazon Nova
Try Amazon’s frontier foundation models.
Careers
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Careers
Explore our open roles.
Amazon Scholars
Faculty research opportunities on industry-scale technical challenges.
Postdoctoral Science Program
Early-career research opportunities alongside experienced industry scientists.
Can we change an AI model’s architecture without retraining it from scratch?
“Priming” reuses pretrained Transformer weights to build models with more efficient architectures at less than 1% of the original training cost and with throughputs up to twice as high.
A new AWS paper introduces priming, a method that reduces pretraining costs for hybrid architectures by over 99% by leveraging pretrained Transformer weights, saving trillions of tokens and hundreds of thousands of GPU-hours.
A 32-billion-parameter Gated KalmaNet (GKA)-based hybrid model matches its source Transformer on reasoning benchmarks while decoding up to twice as quickly with 128,000 tokens of context.
State space models consistently outperform sliding-window attention (SWS) even when SWA receives twice the memory budget, and GKA's variable-iteration mechanism enables 10% faster decoding with only ~1% accuracy loss.
AI agents are rapidly scaling: longer sessions, deeper reasoning chains, more autonomy. But the Transformer architecture underneath them still pays a memory tax that grows with the number of tokens it processes. As agents become more capable, that cost scales with them.
Now, there’s a proliferation of hybrid architectures that use some combination of methods to achieve a better balance of memory, speed, and quality. Those methods include sparse attention, which selects a subset of past tokens to consider when selecting the next token; sliding-window attention, which preserves only the most recent tokens; and state space models, which compress information from the full history into a fixed-size state.
These hybrid architectures can be understood as points in a sequence-model design space: to differing degrees, they mix information while traversing a sequence and passing through different states. What distinguishes them from each other is how they retain and use the past.
Until now, exploring this design space has required training every new architecture from scratch, using trillions of tokens and hundreds of thousands of GPU-hours. This has made controlled comparisons of hybrid architectures extremely expensive.
In a new paper, “ Priming: Hybrid state space models from pre-trained Transformers ,” our team at Amazon Web Services addresses this problem. We introduce priming, a method that leverages pretrained Transformer weights as a starting point to reduce the pretraining cost of hybrid architectures by more than 99%.
With priming, model developers can ask questions that were previously too expensive to answer at scale: Which attention layers can be replaced by more-efficient layers? How many attention layers can be replaced without sacrificing quality? Which state space layer works best? And what does the trade-off between speed and accuracy look like?
In our paper, we compare hybrid models that use three different types of state space models (Gated KalmaNet, Gated DeltaNet, and Mamba-2). Some of our findings:
Not all state space models are equal. Under controlled conditions, Gated KalmaNet (GKA) performs best, followed by Gated DeltaNet (GDN) and then Mamba-2. GKA uniquely offers variable-length inference computation without retraining and gives the user direct control over its position on the speed-accuracy frontier.
A 32-billion-parameter hybrid model reaches each accuracy threshold faster than an equivalently sized Transformer, converging to similar final accuracy in less wall clock time. The dataset is the AIME 2025 math competition benchmark (hard subset). Accuracy at time t is the fraction of rollouts for which a correct final answer has been produced by time t .
A good lossy summary of everything beats a perfect recording of the most recent past. State space models consistently outperform sliding-window attention (SWA), even when SWA gets twice the memory budget.
Same quality, twice the throughput. A 32-billion-parameter GKA-based hybrid model matches its source Transformer on reasoning benchmarks while decoding up to twice as quickly with 128,000 tokens of context and reaching correct answers in as little as two-thirds of the wall-clock time.
To foster further research on hybrid architectures, we have released a zoo of primed hybrid models at the eight-billion-parameter and 32-billion-parameter scales for long-context reasoning and instruction following, together with training code, optimized kernels, and a vLLM-serving plugin under the Apache 2.0 License.
A Transformer’s attention mechanism — which determines how relevant previously seen tokens are to the token currently being processed — relies on three different vectors, which take their names from database design: the query, the key, and the value. For each token, the query captures what that token is “looking for,” the key captures what it “offers,” and the value carries the actual information it will pass along. Relevance is computed by comparing the current token’s query against every other token’s key; the closer the match, the more that token’s value contributes to the output of attention.
Each of these vectors is produced by multiplying the token by a learned weight projection matrix. The Transformer outputs one token at a time, and each output token becomes a new input, so the context grows one token at a time as well. To save computation, the Transformer thus stores the key and value projections of each new input in a “KV cache”.
Two kinds of memory layers: Eidetic versus fading
The attention mechanism gives Transformers what we call eidetic memory : verbatim recall of every past token. The possibility of recalling over an LLM’s entire context window is what made in-context learning, long-document reasoning, and retrieval-augmented generation possible. But it comes at a cost that scales with the number of input tokens.
However, not all attention layers do the same job. Some perform long-range retrieval across tens of thousands of tokens. Others operate more locally, tracking the current clause, copying nearby tokens, and the like. The local ones don't need a growing cache: a compact, fixed-size summary of the past — what we call fading memory — suffices.
Hybrid models mix both: attention layers (eidetic, growing cache) for the layers that need global recall and state space layers (fading, fixed size) for the rest. The result is a model with a reduced memory footprint that still has access to verbatim recall where it matters.
Priming: Knowledge isn't locked to the architecture
Exploring hybrids’ design space has been impractical because each hybrid configuration would normally require pretraining from scratch. Priming sidesteps this by reusing the stored knowledge in a pretrained Transformer's weights.
Priming is possible because a Transformer is itself a special case of a state space realization . The outputs of both attention and SSM layers can be expressed as the result of multiplying the value sequence by a “mixing” matrix that decides how much each past token contributes to each current token. Attention computes the mixing matrix by multiplying the query and key matrices; an SSM produces a similar output through a recurrent state of fixed dimension that all past information must pass through.
Priming takes advantage of this structural similarity: if an SSM with a fixed-dimensional state can produce a mixing matrix close enough to the attention layer's, the query, key, and value projections from the source Transformer can be reused as the SSM's own input projections, with only the SSM-specific parameters needing to be learned.
Whether the approximation holds depends on whether the attention layer performs complex, long-range retrievals or simple, local operations. We distinguish between the two using a cheap empirical proxy: repla

[truncated]
