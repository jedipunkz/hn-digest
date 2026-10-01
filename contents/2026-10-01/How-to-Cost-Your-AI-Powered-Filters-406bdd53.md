---
source: "https://fsdatalab.github.io/blog/ai-filter-cost-estimates/"
hn_url: "https://news.ycombinator.com/item?id=49925844"
title: "How to Cost Your AI-Powered Filters"
article_title: "How to Cost Your AI-Powered Filters | Full Stack Data Lab"
image: "https://fsdatalab.github.io/assets/blog/ai-filter-cost-estimates/social-preview.png"
author: "shreya_shankar"
captured_at: "2026-10-01T20:10:25Z"
capture_tool: "hn-digest"
hn_id: 49925844
score: 2
comments: 0
posted_at: "2026-10-01T19:12:53Z"
tags:
  - hacker-news
---

# How to Cost Your AI-Powered Filters

- HN: [49925844](https://news.ycombinator.com/item?id=49925844)
- Source: [fsdatalab.github.io](https://fsdatalab.github.io/blog/ai-filter-cost-estimates/)
- Score: 2
- Comments: 0
- Posted: 2026-10-01T19:12:53Z

## Translation

Title: How to Cost Your AI-Powered Filters
Article title: How to Cost Your AI-Powered Filters | Full Stack Data Lab
Description: How fast could your AI-powered filters run? We do the math for an H100 and show how filter order changes the answer.

Article text:
How to Cost Your AI-Powered Filters | Full Stack Data Lab
Full Stack Data
Blog
How to Cost Your AI-Powered Filters
Background.
AI-powered filters.
Roofline Model & Speed of Light.
Cost Model for One Filter.
Workload and Notation.
Cost Model for a Conjunction of Filters.
KV Reuse and HBM Capacity.
Example: IMDB Query.
Filter Ordering.
SoL Estimate for the Ordered Conjunction.
We recently released Quail 1 , an execution engine for AI-SQL, which extends SQL with functions that call LLMs. To choose between query plans, Quail needs to estimate how long each AI operation will take. But estimating AI operation latency is not straightforward! In this article, we discuss: How can we estimate the latency of an AI-SQL query on a given LLM and GPU?
One approach is to profile the system by running representative queries and fitting a model to the measurements. But such a profile would have to be redone for every new model, GPU, or workload, and might not even be accurate — which would be a massive headache. Instead, in Quail, we use the roofline model to estimate the latency of an AI operation. 2 That is, we compute the arithmetic and memory traffic needed for a query, then divide by the GPU’s peak compute throughput and memory bandwidth, to obtain a speed-of-light (SoL) estimate, or a lower bound on how quickly a query could run based on hardware limits.
In this article, we’ll walk through how to derive a SoL estimate for AI-SQL queries that only contain filter operators. Well go through:
The GPU and transformer concepts behind our cost model ( Section 2 ).
How to estimate latency for one AI-powered filter ( Section 3 ).
How to cost a conjunction of filters and choose a filter order ( Section 4 ).
An example using Qwen3-4B-fp8 on an H100, with an interactive playground to explore different filter orders ( Sections 5 and 6 ).
Here, we first define AI-powered filters. We then describe the hardware and model choices used in our examples (H100 GPU and Qwen3-4B), giving an overview of the parts of a transformer forward pass that affect cost. Finally, we introduce the roofline model we use to estimate SoL latency.
We define AI-powered filters and introduce the single-filter query and the conjunction of filters used throughout the article.
AI-SQL filter. A regular SQL filter evaluates a condition such as price > 10 . An AI-SQL filter instead asks an LLM to decide whether each row or document satisfies a condition written in natural language. To execute an AI-SQL filter, an LLM is invoked on each row or document and returns a single token, true or false.
Figure 1 displays the general prompt structure for the AI-powered filter query; its structure is similar to that of a SQL query. Within each invocation of the operator there exists a preamble , “DOCUMENT:\n”, the actual document , and the filter instruction “Evaluate TRUE or FALSE for the following question: […]”. The output of each predicate is a single-token output of true or false. Only the surviving documents pass through to subsequent filters. Figure 1 shows a query with a conjunction of filters: \(F_1\) (mentions a positive aspect) → \(F_2\) (discusses the ending) → \(F_3\) (mentions a named actor), all three run over the reviews table. We use the first predicate as the single AI-powered filter example in Section 3 .
The preamble and filter instruction add tokens to each request, so their lengths affect the cost.
Our examples assume an NVIDIA H100 SXM GPU. Its Hopper architecture introduced FP8 Tensor Cores as part of a design aimed at accelerating transformer models. 3 FP8 is an eight-bit floating-point number format.
Figure 2 shows the parts of the H100 that affect our cost model. High-bandwidth memory (HBM) is the GPU’s main memory. It stores the model’s weight matrices, embedding table, and key-value (KV) cache , which holds attention vectors from previously processed tokens for reuse. We explain the vectors in Section 2.3. Data read from HBM passes through the shared L2 cache before being sent to one of the GPU’s 132 streaming multiprocessors (SMs) . Within each SM, registers and the combined shared memory and L1 cache hold small tiles of data close to the Tensor Cores. Each SM has four Tensor Cores, which perform the matrix multiplications used by the model. The H100 has a 50 MB L2 cache shared by all of its SMs.
Our cost model tracks two hardware costs: arithmetic and HBM traffic. For arithmetic, it divides the number of floating-point operations (FLOPs) by the Tensor Core throughput. For HBM traffic, it divides the number of bytes read from or written to HBM by the HBM bandwidth. The model does not account for each intermediate cache level separately.
Table 1 maps the H100 specifications to the symbols used in our equations. Arithmetic throughput is the number of FLOPs that the GPU can perform per second, and memory bandwidth is the number of bytes that it can transfer from HBM per second. \(\Pi_{bf16}\) and \(\Pi_{fp8}\) are the peak arithmetic throughput rates for BF16 (a 16-bit floating-point number format) and FP8 operations, while \(\beta\) is the peak HBM bandwidth. We use the FP8 rate for the projection and MLP matrix multiplications, and the BF16 rate for attention in our FlashAttention 4 setup.
Table 1. H100 SXM hardware constants.
A forward pass processes input tokens through the model’s layers. Each filter processes its input and predicts one token, true or false. Within each layer, projections transform token vectors through matrix multiplication, attention combines information from the current and earlier tokens, and a multilayer perceptron (MLP) applies further transformations to each token independently. Our cost model counts the arithmetic and memory traffic of each component within every layer. We use Qwen3-4B FP8 5 in our examples. 6 The loaded model occupies about 4.5 GB in HBM.
Before the first layer, the model uses an embedding table to map each token ID to a hidden state , a vector of 2,560 values in Qwen3-4B. Each layer transforms the hidden state before passing it to the next layer.
Figure 3 shows the execution order within one Qwen3-4B layer. The model first computes query (Q), key (K), and value (V) projections, then performs attention and an output projection. The MLP follows.
Projections. The model’s Q, K, and V projection matrices transform each token’s hidden state into Q, K, and V vectors used by attention. After attention, the output projection uses another matrix multiplication to produce a vector with the same width as the hidden state. Fortunately, projections can be highly parallelized: each projection applies the same matrix to every token independently, so the GPU can process many tokens at once. Processing more tokens together amortizes the cost of reading the weights from HBM, since the GPU can reuse the same weights across many tokens.
Attention. Attention combines the V vectors of the current token and earlier tokens. Comparing the current token’s Q vector with each token’s K vector determines how much each V vector contributes. Unfortunately, each token must be compared with itself and all previous tokens, so doubling the number of tokens roughly quadruples the attention arithmetic.
Models have multiple attention heads , and the attention calculation above runs separately for each head. Each head uses its own Q vector and computes its result independently, so the GPU can run the heads in parallel.
To reduce the memory needed for K and V, Qwen3-4B uses grouped query attention (GQA) . 7 Its 32 query heads are arranged in eight groups of four, and the heads in each group share the same K and V vectors. Sharing reduces the K and V projection matrices and KV cache to one quarter the size of full multi-head attention , where each of the 32 heads has its own K and V vectors. 8
In Section 4 , we will consider queries with multiple filters on the same document and explain how the filters can reuse the document’s cached K and V vectors ( KV cache ).
MLP. After attention and the output projection, the multilayer perceptron (MLP) further transforms each token’s hidden state. It first multiplies the hidden state by the gate and up matrices to produce two wider vectors. The SwiGLU activation first transforms the gate vector, then uses it to scale each element of the up vector. 9 The down matrix then returns the result to the original hidden-state width. Like projections, the MLP operates on each token independently, so many tokens can run in parallel and share the weights read from HBM. Its arithmetic grows with the number of tokens, rather than the number of token pairs as in attention.
The resulting hidden states pass to the next transformer layer, which repeats the projections, attention, and MLP. Qwen3-4B has 36 layers. After the final layer, the model uses the same embedding table to convert the final hidden state into scores for output tokens. Table 2 lists the matrix dimensions and parameter counts we need to calculate the arithmetic and memory traffic across all layers.
Table 2. Qwen3-4B architecture constants.
The parameter count \(P\) includes only projection and MLP weights. The embedding table contains \(\text{vocab}\times d_{\text{model}}\) values, where \(\text{vocab}\) is the number of distinct tokens the model can read or output and \(d_{\text{model}}\) is the hidden-state width.
2.4 Roofline Model & Speed of Light
For each component of a forward pass, we need to estimate the time spent on arithmetic and memory transfers. The compute time is the number of FLOPs divided by the GPU’s arithmetic throughput. The memory transfer time is the number of bytes read from or written to HBM divided by the HBM bandwidth.
Sometimes arithmetic takes longer; other times, memory transfers take longer. We use the roofline model to estimate latency from both times. The model assumes that arithmetic and memory transfers overlap, so the component’s latency is the longer of the two:
Here, \(T\) is the estimated latency, \(\Pi\) is the arithmetic throughput in FLOP/s, and \(\beta\) is the HBM bandwidth in bytes/s. Using the GPU’s peak arithmetic throughput and memory bandwidth gives a speed-of-light (SoL) latency estimate for the operation. The estimate is a lower bound because an implementation may not sustain the peak rates or overlap arithmetic and memory transfers perfectly.
The roofline model predates modern LLMs. Williams et al. introduced it in 2009 to relate arithmetic and memory traffic to hardware performance. We learned a lot about applying it to transformer inference from blog posts by Kipply 10 , Fergus Finn 11 , Ben Mayer 12 , and Modal 13 , which we recommend reading.
If memory transfers take longer, the operation is memory-bound . If arithmetic takes longer, it is compute-bound . To determine which case applies, we compare the operation’s operational intensity , \(I\), with the GPU’s ridge point , \(I^*\):
Here, \(I\) is the arithmetic performed per byte transferred from HBM. \(I^*\) is the intensity at which compute time and memory transfer time are equal. Below the ridge point, the operation is memory-bound. Above it, the operation is compute-bound.
The H100 has different arithmetic throughput rates for FP8 and BF16, so each has a different ridge point. Our projections and MLP use FP8, while attention uses BF16. For the H100 rates in Table 1 :
Here, \(I^*\) is the FP8 ridge point used for projections and the MLP, and \(I^*_{bf16}\) is the BF16 ridge point used for attention.
Figure 4 shows how operational intensity limits arithmetic throughput. In the sloped region, HBM bandwidth is the limit: performing more arithmetic for each byte transferred allows higher throughput. Once an operation reaches the ridge point, arithmetic throughput is the limit, so the plot becomes horizontal.
In Section 3.5 , we will calculate how many tokens we need to process before arithmetic, rather than memory transfers, determines each component’s latency.
We estimate one filter’s latency by counting the arith

[truncated]

## Original Extract

How fast could your AI-powered filters run? We do the math for an H100 and show how filter order changes the answer.

How to Cost Your AI-Powered Filters | Full Stack Data Lab
Full Stack Data
Blog
How to Cost Your AI-Powered Filters
Background.
AI-powered filters.
Roofline Model & Speed of Light.
Cost Model for One Filter.
Workload and Notation.
Cost Model for a Conjunction of Filters.
KV Reuse and HBM Capacity.
Example: IMDB Query.
Filter Ordering.
SoL Estimate for the Ordered Conjunction.
We recently released Quail 1 , an execution engine for AI-SQL, which extends SQL with functions that call LLMs. To choose between query plans, Quail needs to estimate how long each AI operation will take. But estimating AI operation latency is not straightforward! In this article, we discuss: How can we estimate the latency of an AI-SQL query on a given LLM and GPU?
One approach is to profile the system by running representative queries and fitting a model to the measurements. But such a profile would have to be redone for every new model, GPU, or workload, and might not even be accurate — which would be a massive headache. Instead, in Quail, we use the roofline model to estimate the latency of an AI operation. 2 That is, we compute the arithmetic and memory traffic needed for a query, then divide by the GPU’s peak compute throughput and memory bandwidth, to obtain a speed-of-light (SoL) estimate, or a lower bound on how quickly a query could run based on hardware limits.
In this article, we’ll walk through how to derive a SoL estimate for AI-SQL queries that only contain filter operators. Well go through:
The GPU and transformer concepts behind our cost model ( Section 2 ).
How to estimate latency for one AI-powered filter ( Section 3 ).
How to cost a conjunction of filters and choose a filter order ( Section 4 ).
An example using Qwen3-4B-fp8 on an H100, with an interactive playground to explore different filter orders ( Sections 5 and 6 ).
Here, we first define AI-powered filters. We then describe the hardware and model choices used in our examples (H100 GPU and Qwen3-4B), giving an overview of the parts of a transformer forward pass that affect cost. Finally, we introduce the roofline model we use to estimate SoL latency.
We define AI-powered filters and introduce the single-filter query and the conjunction of filters used throughout the article.
AI-SQL filter. A regular SQL filter evaluates a condition such as price > 10 . An AI-SQL filter instead asks an LLM to decide whether each row or document satisfies a condition written in natural language. To execute an AI-SQL filter, an LLM is invoked on each row or document and returns a single token, true or false.
Figure 1 displays the general prompt structure for the AI-powered filter query; its structure is similar to that of a SQL query. Within each invocation of the operator there exists a preamble , “DOCUMENT:\n”, the actual document , and the filter instruction “Evaluate TRUE or FALSE for the following question: […]”. The output of each predicate is a single-token output of true or false. Only the surviving documents pass through to subsequent filters. Figure 1 shows a query with a conjunction of filters: \(F_1\) (mentions a positive aspect) → \(F_2\) (discusses the ending) → \(F_3\) (mentions a named actor), all three run over the reviews table. We use the first predicate as the single AI-powered filter example in Section 3 .
The preamble and filter instruction add tokens to each request, so their lengths affect the cost.
Our examples assume an NVIDIA H100 SXM GPU. Its Hopper architecture introduced FP8 Tensor Cores as part of a design aimed at accelerating transformer models. 3 FP8 is an eight-bit floating-point number format.
Figure 2 shows the parts of the H100 that affect our cost model. High-bandwidth memory (HBM) is the GPU’s main memory. It stores the model’s weight matrices, embedding table, and key-value (KV) cache , which holds attention vectors from previously processed tokens for reuse. We explain the vectors in Section 2.3. Data read from HBM passes through the shared L2 cache before being sent to one of the GPU’s 132 streaming multiprocessors (SMs) . Within each SM, registers and the combined shared memory and L1 cache hold small tiles of data close to the Tensor Cores. Each SM has four Tensor Cores, which perform the matrix multiplications used by the model. The H100 has a 50 MB L2 cache shared by all of its SMs.
Our cost model tracks two hardware costs: arithmetic and HBM traffic. For arithmetic, it divides the number of floating-point operations (FLOPs) by the Tensor Core throughput. For HBM traffic, it divides the number of bytes read from or written to HBM by the HBM bandwidth. The model does not account for each intermediate cache level separately.
Table 1 maps the H100 specifications to the symbols used in our equations. Arithmetic throughput is the number of FLOPs that the GPU can perform per second, and memory bandwidth is the number of bytes that it can transfer from HBM per second. \(\Pi_{bf16}\) and \(\Pi_{fp8}\) are the peak arithmetic throughput rates for BF16 (a 16-bit floating-point number format) and FP8 operations, while \(\beta\) is the peak HBM bandwidth. We use the FP8 rate for the projection and MLP matrix multiplications, and the BF16 rate for attention in our FlashAttention 4 setup.
Table 1. H100 SXM hardware constants.
A forward pass processes input tokens through the model’s layers. Each filter processes its input and predicts one token, true or false. Within each layer, projections transform token vectors through matrix multiplication, attention combines information from the current and earlier tokens, and a multilayer perceptron (MLP) applies further transformations to each token independently. Our cost model counts the arithmetic and memory traffic of each component within every layer. We use Qwen3-4B FP8 5 in our examples. 6 The loaded model occupies about 4.5 GB in HBM.
Before the first layer, the model uses an embedding table to map each token ID to a hidden state , a vector of 2,560 values in Qwen3-4B. Each layer transforms the hidden state before passing it to the next layer.
Figure 3 shows the execution order within one Qwen3-4B layer. The model first computes query (Q), key (K), and value (V) projections, then performs attention and an output projection. The MLP follows.
Projections. The model’s Q, K, and V projection matrices transform each token’s hidden state into Q, K, and V vectors used by attention. After attention, the output projection uses another matrix multiplication to produce a vector with the same width as the hidden state. Fortunately, projections can be highly parallelized: each projection applies the same matrix to every token independently, so the GPU can process many tokens at once. Processing more tokens together amortizes the cost of reading the weights from HBM, since the GPU can reuse the same weights across many tokens.
Attention. Attention combines the V vectors of the current token and earlier tokens. Comparing the current token’s Q vector with each token’s K vector determines how much each V vector contributes. Unfortunately, each token must be compared with itself and all previous tokens, so doubling the number of tokens roughly quadruples the attention arithmetic.
Models have multiple attention heads , and the attention calculation above runs separately for each head. Each head uses its own Q vector and computes its result independently, so the GPU can run the heads in parallel.
To reduce the memory needed for K and V, Qwen3-4B uses grouped query attention (GQA) . 7 Its 32 query heads are arranged in eight groups of four, and the heads in each group share the same K and V vectors. Sharing reduces the K and V projection matrices and KV cache to one quarter the size of full multi-head attention , where each of the 32 heads has its own K and V vectors. 8
In Section 4 , we will consider queries with multiple filters on the same document and explain how the filters can reuse the document’s cached K and V vectors ( KV cache ).
MLP. After attention and the output projection, the multilayer perceptron (MLP) further transforms each token’s hidden state. It first multiplies the hidden state by the gate and up matrices to produce two wider vectors. The SwiGLU activation first transforms the gate vector, then uses it to scale each element of the up vector. 9 The down matrix then returns the result to the original hidden-state width. Like projections, the MLP operates on each token independently, so many tokens can run in parallel and share the weights read from HBM. Its arithmetic grows with the number of tokens, rather than the number of token pairs as in attention.
The resulting hidden states pass to the next transformer layer, which repeats the projections, attention, and MLP. Qwen3-4B has 36 layers. After the final layer, the model uses the same embedding table to convert the final hidden state into scores for output tokens. Table 2 lists the matrix dimensions and parameter counts we need to calculate the arithmetic and memory traffic across all layers.
Table 2. Qwen3-4B architecture constants.
The parameter count \(P\) includes only projection and MLP weights. The embedding table contains \(\text{vocab}\times d_{\text{model}}\) values, where \(\text{vocab}\) is the number of distinct tokens the model can read or output and \(d_{\text{model}}\) is the hidden-state width.
2.4 Roofline Model & Speed of Light
For each component of a forward pass, we need to estimate the time spent on arithmetic and memory transfers. The compute time is the number of FLOPs divided by the GPU’s arithmetic throughput. The memory transfer time is the number of bytes read from or written to HBM divided by the HBM bandwidth.
Sometimes arithmetic takes longer; other times, memory transfers take longer. We use the roofline model to estimate latency from both times. The model assumes that arithmetic and memory transfers overlap, so the component’s latency is the longer of the two:
Here, \(T\) is the estimated latency, \(\Pi\) is the arithmetic throughput in FLOP/s, and \(\beta\) is the HBM bandwidth in bytes/s. Using the GPU’s peak arithmetic throughput and memory bandwidth gives a speed-of-light (SoL) latency estimate for the operation. The estimate is a lower bound because an implementation may not sustain the peak rates or overlap arithmetic and memory transfers perfectly.
The roofline model predates modern LLMs. Williams et al. introduced it in 2009 to relate arithmetic and memory traffic to hardware performance. We learned a lot about applying it to transformer inference from blog posts by Kipply 10 , Fergus Finn 11 , Ben Mayer 12 , and Modal 13 , which we recommend reading.
If memory transfers take longer, the operation is memory-bound . If arithmetic takes longer, it is compute-bound . To determine which case applies, we compare the operation’s operational intensity , \(I\), with the GPU’s ridge point , \(I^*\):
Here, \(I\) is the arithmetic performed per byte transferred from HBM. \(I^*\) is the intensity at which compute time and memory transfer time are equal. Below the ridge point, the operation is memory-bound. Above it, the operation is compute-bound.
The H100 has different arithmetic throughput rates for FP8 and BF16, so each has a different ridge point. Our projections and MLP use FP8, while attention uses BF16. For the H100 rates in Table 1 :
Here, \(I^*\) is the FP8 ridge point used for projections and the MLP, and \(I^*_{bf16}\) is the BF16 ridge point used for attention.
Figure 4 shows how operational intensity limits arithmetic throughput. In the sloped region, HBM bandwidth is the limit: performing more arithmetic for each byte transferred allows higher throughput. Once an operation reaches the ridge point, arithmetic throughput is the limit, so the plot becomes horizontal.
In Section 3.5 , we will calculate how many tokens we need to process before arithmetic, rather than memory transfers, determines each component’s latency.
We estimate one filter’s latency by counting the arith

[truncated]
