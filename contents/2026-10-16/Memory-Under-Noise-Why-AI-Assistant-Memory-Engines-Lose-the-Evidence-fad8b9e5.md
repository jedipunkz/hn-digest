---
source: "https://www.aai-labs.com/en/research/memory-under-noise"
hn_url: "https://news.ycombinator.com/item?id=50116468"
title: "Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence"
article_title: "Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows | AAI Labs"
image: ""
author: "tasubotadas"
captured_at: "2026-10-16T10:50:21Z"
capture_tool: "hn-digest"
hn_id: 50116468
score: 1
comments: 0
posted_at: "2026-10-16T10:06:26Z"
tags:
  - hacker-news
---

# Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence

- HN: [50116468](https://news.ycombinator.com/item?id=50116468)
- Source: [www.aai-labs.com](https://www.aai-labs.com/en/research/memory-under-noise)
- Score: 1
- Comments: 0
- Posted: 2026-10-16T10:06:26Z

## Translation

Title: Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence
Article title: Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows | AAI Labs
Description: We benchmarked five AI memory systems. Hybrid RAG won on accuracy and cost; Hindsight was the strongest dedicated engine. Here is where each system fails and what we recommend for assistants with growing conversation histories.

Article text:
Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows | AAI Labs Main Services Cases Research Team Company More Contact Loading… Services AI for Energy
UAB Taikomasis dirbtinis intelektas © 2026
Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows
We benchmarked five AI memory systems. Hybrid RAG won on accuracy and cost; Hindsight was the strongest dedicated engine. Here is where each system fails and what we recommend for assistants with growing conversation histories.
We benchmarked five open-source memory systems to answer a practical question: which one should power an assistant that needs to remember a growing conversation history? Our recommendation is hybrid RAG over the original transcripts. It was the most accurate option at every noise level, and its total run cost was a fraction of every dedicated memory engine we tested. A more elaborate memory architecture needs to earn its place by beating that baseline.
What worked, what failed, and what we would choose
The results below use an 8K-token context budget with 90% unrelated history. Costs cover the full 1,000-question run, including memory building and answering. Graphiti’s result is at 50% noise because we did not run it at 90%.
Hybrid RAG: our default choice
96.8% accuracy for $1.70. Chunking the raw history and combining keyword and vector search beat all five memory engines. Accuracy fell only 1.7 percentage points as noise grew from 0% to 90%. It preserves the original wording and attribution, then gives the answer model enough relevant text to work with. For conversational recall, this is where we would start.
Hindsight: the strongest dedicated memory engine
89.9% accuracy for $20.39. Its combination of semantic, keyword, graph and temporal retrieval kept relevant evidence available as history grew. At a 2K context budget, it was statistically tied with RAG. That makes it the most compelling dedicated engine in this comparison when context is tight. The trade-off is operational cost: roughly 12 times RAG’s total run cost, and median retrieval latency rising from 2.8 to 15.6 seconds as the store grew.
OpenViking: compact and stable, with a recall ceiling
82.7% accuracy for $20.73. It stayed close to its 84.3% clean-history score while returning under 1,000 context tokens on average. Its fixed quotas per memory category keep context compact, but leave useful evidence out even before much noise is added. Choose this trade-off deliberately if compact context matters; we would not choose it to maximise recall.
Mem0: lower cost, but facts get crowded out
81.5% accuracy for $7.03, down from 90.9% without noise. It was the least expensive dedicated engine at 90% noise, but its top-20 memory limit became a bottleneck as unrelated history accumulated. In persona conversations, extraction also dropped speaker names, weakening answers about who said what. It deserves consideration when a small set of extracted facts is enough; we would not trust the default retrieval for detailed recall over a long history.
Honcho: configuration makes a substantial difference
69.7% accuracy for $18.28 after memory building finished. The earlier unfinished build scored 66.3%. Tuning retrieval to fill each person’s share of the context budget with similar observations raised accuracy to 77.5%. Its default mix of similar and recent observations gives too much space to recency as history grows. Our recommendation: finish the background build and tune retrieval before relying on it. Even after tuning, it remained behind RAG and Hindsight here.
Graphiti: a poor default fit for conversational recall
53.6% accuracy at 50% noise for $19.84. Returning only ten graph facts often failed to supply enough evidence for multi-part answers. We would not pick this default configuration for question answering over conversation history. Use a graph because your application needs explicit relationships and temporal structure, then verify that its retrieval can answer your actual questions.
The failure pattern: memory exists, but the answer never sees enough of it
The important distinction was what retrieval returned. RAG and Hindsight filled the available budget with relevance-ranked text. Fixed item counts and slots reserved for recent memories let irrelevant history crowd out the evidence. Simply pasting in the latest history also failed: the full-context baseline fell from 98.0% to 52.4% once older evidence no longer fit.
Extraction creates another weak point. The extraction engines scored 6–18 percentage points worse on facts mentioned once than on repeated facts; RAG stayed around 98–99% in that comparison. Missing context is not the whole story either: answer terms appeared in many wrong Mem0 and Hindsight answers, pointing to incomplete, misattributed or mis-dated memories. A fact being stored somewhere does not mean an assistant can use it correctly.
Our takeaway: make the memory layer prove its value
Keep the original history and establish a hybrid RAG baseline first. Add a memory engine when you can demonstrate a benefit your product needs.
Inspect what reaches the answer model. Test facts mentioned once, speaker attribution and older conversations as the history grows. Increase fixed retrieval limits or remove recency bias when they hide relevant evidence.
Budget for ingestion as well as queries. Building memory consumed 90–98% of total run costs for the four extraction engines at 90% noise. Many short sessions can be more expensive than fewer long ones.
We tested Hindsight, Mem0, Honcho, OpenViking and Graphiti against hybrid RAG, full-context and no-memory baselines on 1,000 questions drawn from four conversation and email corpora. Unrelated history grew from 0% to 50% to 90%, with the model, embedder, answer prompt and judge held constant. These recommendations concern conversational question answering; the study does not measure every engine’s editing, inspection or other product features. Small differences such as Mem0 versus OpenViking are not a reliable basis for a ranking.
Tadas Subonis, AAI Labs. The full 14-page paper is in English.

## Original Extract

We benchmarked five AI memory systems. Hybrid RAG won on accuracy and cost; Hindsight was the strongest dedicated engine. Here is where each system fails and what we recommend for assistants with growing conversation histories.

Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows | AAI Labs Main Services Cases Research Team Company More Contact Loading… Services AI for Energy
UAB Taikomasis dirbtinis intelektas © 2026
Memory Under Noise: Why AI-Assistant Memory Engines Lose the Evidence as a User’s History Grows
We benchmarked five AI memory systems. Hybrid RAG won on accuracy and cost; Hindsight was the strongest dedicated engine. Here is where each system fails and what we recommend for assistants with growing conversation histories.
We benchmarked five open-source memory systems to answer a practical question: which one should power an assistant that needs to remember a growing conversation history? Our recommendation is hybrid RAG over the original transcripts. It was the most accurate option at every noise level, and its total run cost was a fraction of every dedicated memory engine we tested. A more elaborate memory architecture needs to earn its place by beating that baseline.
What worked, what failed, and what we would choose
The results below use an 8K-token context budget with 90% unrelated history. Costs cover the full 1,000-question run, including memory building and answering. Graphiti’s result is at 50% noise because we did not run it at 90%.
Hybrid RAG: our default choice
96.8% accuracy for $1.70. Chunking the raw history and combining keyword and vector search beat all five memory engines. Accuracy fell only 1.7 percentage points as noise grew from 0% to 90%. It preserves the original wording and attribution, then gives the answer model enough relevant text to work with. For conversational recall, this is where we would start.
Hindsight: the strongest dedicated memory engine
89.9% accuracy for $20.39. Its combination of semantic, keyword, graph and temporal retrieval kept relevant evidence available as history grew. At a 2K context budget, it was statistically tied with RAG. That makes it the most compelling dedicated engine in this comparison when context is tight. The trade-off is operational cost: roughly 12 times RAG’s total run cost, and median retrieval latency rising from 2.8 to 15.6 seconds as the store grew.
OpenViking: compact and stable, with a recall ceiling
82.7% accuracy for $20.73. It stayed close to its 84.3% clean-history score while returning under 1,000 context tokens on average. Its fixed quotas per memory category keep context compact, but leave useful evidence out even before much noise is added. Choose this trade-off deliberately if compact context matters; we would not choose it to maximise recall.
Mem0: lower cost, but facts get crowded out
81.5% accuracy for $7.03, down from 90.9% without noise. It was the least expensive dedicated engine at 90% noise, but its top-20 memory limit became a bottleneck as unrelated history accumulated. In persona conversations, extraction also dropped speaker names, weakening answers about who said what. It deserves consideration when a small set of extracted facts is enough; we would not trust the default retrieval for detailed recall over a long history.
Honcho: configuration makes a substantial difference
69.7% accuracy for $18.28 after memory building finished. The earlier unfinished build scored 66.3%. Tuning retrieval to fill each person’s share of the context budget with similar observations raised accuracy to 77.5%. Its default mix of similar and recent observations gives too much space to recency as history grows. Our recommendation: finish the background build and tune retrieval before relying on it. Even after tuning, it remained behind RAG and Hindsight here.
Graphiti: a poor default fit for conversational recall
53.6% accuracy at 50% noise for $19.84. Returning only ten graph facts often failed to supply enough evidence for multi-part answers. We would not pick this default configuration for question answering over conversation history. Use a graph because your application needs explicit relationships and temporal structure, then verify that its retrieval can answer your actual questions.
The failure pattern: memory exists, but the answer never sees enough of it
The important distinction was what retrieval returned. RAG and Hindsight filled the available budget with relevance-ranked text. Fixed item counts and slots reserved for recent memories let irrelevant history crowd out the evidence. Simply pasting in the latest history also failed: the full-context baseline fell from 98.0% to 52.4% once older evidence no longer fit.
Extraction creates another weak point. The extraction engines scored 6–18 percentage points worse on facts mentioned once than on repeated facts; RAG stayed around 98–99% in that comparison. Missing context is not the whole story either: answer terms appeared in many wrong Mem0 and Hindsight answers, pointing to incomplete, misattributed or mis-dated memories. A fact being stored somewhere does not mean an assistant can use it correctly.
Our takeaway: make the memory layer prove its value
Keep the original history and establish a hybrid RAG baseline first. Add a memory engine when you can demonstrate a benefit your product needs.
Inspect what reaches the answer model. Test facts mentioned once, speaker attribution and older conversations as the history grows. Increase fixed retrieval limits or remove recency bias when they hide relevant evidence.
Budget for ingestion as well as queries. Building memory consumed 90–98% of total run costs for the four extraction engines at 90% noise. Many short sessions can be more expensive than fewer long ones.
We tested Hindsight, Mem0, Honcho, OpenViking and Graphiti against hybrid RAG, full-context and no-memory baselines on 1,000 questions drawn from four conversation and email corpora. Unrelated history grew from 0% to 50% to 90%, with the model, embedder, answer prompt and judge held constant. These recommendations concern conversational question answering; the study does not measure every engine’s editing, inspection or other product features. Small differences such as Mem0 versus OpenViking are not a reliable basis for a ranking.
Tadas Subonis, AAI Labs. The full 14-page paper is in English.
