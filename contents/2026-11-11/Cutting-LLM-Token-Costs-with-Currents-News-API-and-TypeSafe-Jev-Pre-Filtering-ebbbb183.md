---
source: "https://currentsapi.services/en/blog/cut-llm-token-costs-typesafe-jev-prefiltering"
hn_url: "https://news.ycombinator.com/item?id=50450911"
title: "Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering"
article_title: "Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering"
image: "https://currentsapi.services/img/Cover-Image.jpg"
author: "gradientpenalty"
captured_at: "2026-11-11T08:55:17Z"
capture_tool: "hn-digest"
hn_id: 50450911
score: 2
comments: 0
posted_at: "2026-11-11T08:17:35Z"
tags:
  - hacker-news
---

# Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering

- HN: [50450911](https://news.ycombinator.com/item?id=50450911)
- Source: [currentsapi.services](https://currentsapi.services/en/blog/cut-llm-token-costs-typesafe-jev-prefiltering)
- Score: 2
- Comments: 0
- Posted: 2026-11-11T08:17:35Z

## Translation

Title: Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering
Description: Screening news with TypeSafe Jev before LLM synthesis cut input tokens 16–18% at a 0.70 threshold; blind judges could not tell five of six cells from raw.

Article text:
Platform
Introduction
Statistics -->
Sources
Solutions
Use Cases
Compare Alternatives
Pricing Comparison
Features Comparison
Pricing
Developers
Documentation
FAQ
Blog
◐
Comfort off
Log in
Start free
◐
Comfort off
Log in
Start free
250 requests/day · no card
Blog
Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering
Screening news with TypeSafe Jev before LLM synthesis cut input tokens 16–18% at a 0.70 threshold; blind judges could not tell five of six cells from raw.
Currents News API keyword search is built for recall. A query for a geopolitical development or a stock ticker returns every publication that mentions the search terms, including wire duplicates, boilerplate press releases, and generic sector watchlists. An agent that summarizes this output pays frontier input prices for every article, on-topic or not, on every scheduled run.
The two tiers differ by roughly 50–70× in input-token price. Frontier synthesis input currently runs $2.00–3.00 per million tokens; TypeSafe Jev classification input is $0.042 per million tokens. The practical answer is to move the relevance decision upstream: screen each candidate title and snippet with Jev, and send only retained candidates to the synthesis model.
We built this pipeline end-to-end and measured it on live data: 2,000 candidate articles across two real news pools, 60 briefings from three synthesis models, and a blind panel that scored filtered versus raw briefings without knowing which was which. At the default 0.70 relevance threshold, screening cut downstream input tokens by 16–18%; five of six blind-panel quality cells were statistically indistinguishable from raw, and the sixth showed a small negative delta. Stricter thresholds screened out 30–52% of candidates. At frontier pricing, this nets roughly $10–17 per 1,000 screened 200-article batches per synthesis model.
The cost of raw keyword retrieval
One 200-article candidate batch costs a frontier model roughly 35,000–45,000 prompt tokens, and a full 1,000-article pool costs roughly 174,000–223,000 per pass. A scheduled daily pipeline pays that on every run. The cost is not only financial. Tangential commentary and syndication duplicates also consume the model's attention budget and slow generation, because time-to-first-token scales with context length.
Three questions determine whether screening helps: how much of a real keyword-search pool is irrelevant, how many tokens removing it saves, and whether the briefing degrades when filtered. This article measures all three on live data.
Follow along with a free key. Free covers 250 requests/day with a 30-day lookback.
The core idea: screen candidates before synthesis
TypeSafe Jev is a classification API priced at $0.042 per million input tokens, evaluated here with typesafe-sdk 0.6.0 . For each candidate, the pipeline sends a compact screening prompt containing the title plus a description snippet and one Boolean relevance question through the Noul primitive. Jev returns a probability between 0 and 1. Candidates at or above the threshold pass through with headline, metadata, and Currents description. The rest never reach the synthesis model.
In practice, this means the relevance decision is made at classification cost instead of synthesis cost.
How the benchmark was measured
The benchmark used two live news pools of 1,000 candidates each, pulled from Currents at benchmark time. No hand-picked sample.
The first pool was geopolitics: U.S.–Greenland sovereignty and Arctic defense developments. The second was financial intelligence: AST SpaceMobile ($ASTS) constellation deployment, technical milestones, and capital structure. Each 1,000-candidate topic pool was assembled from 12 topic-specific query variants across monthly windows and deduplicated by article ID; individual articles were not manually selected.
Each topic was split into five batches of 200 candidates, and each batch was synthesized twice by the same model, once from the raw packet and once from the Jev-filtered packet. That produced 60 briefings (2 topics × 5 batches × 3 models × 2 arms). Synthesis models: claude-sonnet-4-6 , gemini-3.8-flash , glm-5.3-flash .
A blind panel of four judge models, none of them the author of a scored briefing, scored each pair on coverage, faithfulness, unsupported-claims count, and usefulness. Labels were randomized so judges could not identify the filtered briefing. We collected 82 judge verdicts across 30 pairs. The Claude judge completed 12 of 20 assignments before its provider quota interrupted the session, and that partial coverage is reflected in the panel aggregates below.
Every token count in this article comes from provider usage fields: Jev screening tokens, synthesis prompt tokens, and output tokens. No tokenizer estimates. Synthesis ran at temperature 0 to reduce sampling variability; each filtered/raw comparison used the same model and prompt structure, differing only in the candidate packet.
Quality deltas are aggregated per author-model as filtered − raw and reported with 95% confidence intervals from a batch-level bootstrap (10 batches total).
Measured token reduction on live keyword pools
Input-token reduction at the default threshold
Across both pools, Jev retained about 82% of candidates and cut synthesis input by roughly a sixth. (Prompt tokens in the table are per screened batch, summed across the three synthesis models.) Screening itself averaged 546 input tokens per article, for 1,092,430 tokens across the 2,000 screenings. That works out to about $0.000023 per article at Jev's $0.042/MTok rate.
These results suggest live keyword pools for precise queries contain less noise than commonly assumed. When the query names a specific entity or event, most returned articles are on-topic, so the filtering gain at the default threshold is moderate but consistent.
Candidate reduction across threshold settings
Jev returns a probability for every article, so the cutoff is a configurable parameter. The 0.70 default retains about 82% of candidates. Higher thresholds remove more candidates:
At 0.85 the same pools shed 30–35% of candidates; at 0.95, over half. Higher cutoffs trade edge-case coverage for savings, so the threshold should match the recall margin the application requires.
Measured briefing quality after filtering
Token savings are only useful if briefing quality is maintained. The blind panel therefore scored every pair with author and judge roles kept separate and labels randomized.
Panel utility deltas (filtered − raw) per synthesis model, both topics:
Five of six cells cross zero: the panel could not reliably distinguish filtered from raw. The exception is small and negative; glm-5.3-flash briefings on the Greenland topic scored −1.2 points under filtering, a detectable but minor dip on that cell. In three of the six cells, judges slightly preferred the filtered version.
When filtering reduces cost, and when it increases cost
Screening adds a new cost term; the net effect depends on synthesis pricing. Screening a 200-article batch costs about $0.0046 (200 articles × 546 tokens × $0.042/MTok), or $4.59 per 1,000 screened batches.
Net savings per 1,000 screened batches at the 0.70 threshold, priced per synthesis model. Measured token removal per run: 7,104 tokens for claude-sonnet-4-6 (the benchmark author); 6,689 for the six-cell pool mean, which the GPT rows use because no GPT model was a benchmark author:
At frontier price points ($2.00–3.00/MTok), screening nets roughly $10–17 per 1,000 screened batches for a single-model pipeline; a pipeline that briefs with all three benchmark models from one screening realizes roughly three times the LLM savings against the same Jev cost. At the current mini-class price (GPT-5 Mini, $0.25/MTok), the 0.70 threshold is net negative (−$2.92); a stricter threshold would be required, and the exact break-even point depends on the token lengths of the discarded candidates.
The practical implication is that semantic pre-filtering is a frontier-model optimization. Frontier models are the tier where input length translates most directly into cost.
Example implementation in Python
The pipeline fits in one script: search Currents, screen each candidate with Jev, brief with Claude. page_size is set to 20 for compatibility with Currents API free-tier developer accounts. Paid tiers accommodate higher candidate volumes, and screening calls parallelize cleanly.
import os
import requests
from typesafe_sdk import TypeSafeClient , Noul
import anthropic
# 1. Initialize Clients
CURRENTS_API_KEY = os . environ [ "CURRENTS_API_KEY" ]
TYPESAFE_API_KEY = os . environ [ "TYPESAFE_API_KEY" ]
ANTHROPIC_API_KEY = os . environ [ "ANTHROPIC_API_KEY" ]
typesafe = TypeSafeClient ( api_key = TYPESAFE_API_KEY )
claude = anthropic . Anthropic ( api_key = ANTHROPIC_API_KEY )
# 2. Retrieve Candidate Articles from Currents News API
# Note: Currents API supports Bearer token authentication.
query = "AST SpaceMobile"
response = requests . get (
"https://api.currentsapi.services/v1/search" ,
headers = { "Authorization" : f "Bearer { CURRENTS_API_KEY } " },
params = {
"keywords" : query ,
"language" : "en" ,
"page_size" : 20 ,
},
timeout = 10 ,
)
response . raise_for_status ()
candidates = response . json () . get ( "news" , [])
# 3. Pre-Filtering with TypeSafe System One (Jev)
relevance_question = Noul (
instructions = (
"Does this article provide substantive primary reporting, financial analysis, "
"or technical updates directly about AST SpaceMobile ($ASTS)? Answer NO to generic "
"watchlists, unrelated company news, or shareholder litigation solicitations."
)
)
retained_articles = []
for article in candidates :
title = article . get ( "title" , "" ) . strip ()
summary = article . get ( "description" , "" ) . strip ()[: 800 ]
state_text = f "Title: { title } \n Summary: { summary } "
judgment = typesafe . system_one (
state = state_text ,
questions = { "is_relevant" : relevance_question },
)
prob = judgment . nouls [ "is_relevant" ] . noul
if prob >= 0.70 :
retained_articles . append ( article )
print ( f "Retained { len ( retained_articles ) } of { len ( candidates ) } candidates." )
# 4. Synthesize Briefing with Frontier Model
context_blocks = []
for idx , art in enumerate ( retained_articles , 1 ):
context_blocks . append (
f "--- Article { idx } --- \n "
f "Title: { art . get ( 'title' ) } \n "
f "URL: { art . get ( 'url' ) } \n "
f "Publishe
[truncated]
Jev outputs calibrated probabilities, not deterministic guarantees. On these pools, 0.70 was conservative: five of six blind-panel cells were statistically indistinguishable from raw, and one showed a small negative delta. Applications that cannot tolerate false negatives in open-domain retrieval should use 0.50–0.60 and accept smaller savings.
The threshold only matters relative to the synthesis model's price. The economics table above is the decision tool: frontier synthesis justifies 0.70; at mini-class prices 0.70 is net negative and a stricter threshold would be required — the exact break-even depends on the token lengths of the discarded candidates — so mini-class pipelines may not need the filter at all.
If the pipeline relies on provider prompt caching with static context prefixes, check whether dynamic pre-filtering changes what is cached. For multi-entity monitoring, where each user watches a different query, there is no cache to preserve and filtering is a direct saving.
Screening verdicts can be cached across scheduled runs for repeated (article, question) pairs. The measured numbers are all cold-start.
Keyword retrieval returns recall-oriented results, and unfiltered synthesis bills all of them at frontier input prices. The measured two-tier setup (Currents News API for recall, TypeSafe Jev for semantic pre-filtering) cut briefing input by 16–18% at a conservative threshold, and stricter thresholds screened out 30–52% of candidates, with five of six blind-panel quality cells statistically indistinguishable from ra

[truncated]

## Original Extract

Screening news with TypeSafe Jev before LLM synthesis cut input tokens 16–18% at a 0.70 threshold; blind judges could not tell five of six cells from raw.

Platform
Introduction
Statistics -->
Sources
Solutions
Use Cases
Compare Alternatives
Pricing Comparison
Features Comparison
Pricing
Developers
Documentation
FAQ
Blog
◐
Comfort off
Log in
Start free
◐
Comfort off
Log in
Start free
250 requests/day · no card
Blog
Cutting LLM Token Costs with Currents News API and TypeSafe Jev Pre-Filtering
Screening news with TypeSafe Jev before LLM synthesis cut input tokens 16–18% at a 0.70 threshold; blind judges could not tell five of six cells from raw.
Currents News API keyword search is built for recall. A query for a geopolitical development or a stock ticker returns every publication that mentions the search terms, including wire duplicates, boilerplate press releases, and generic sector watchlists. An agent that summarizes this output pays frontier input prices for every article, on-topic or not, on every scheduled run.
The two tiers differ by roughly 50–70× in input-token price. Frontier synthesis input currently runs $2.00–3.00 per million tokens; TypeSafe Jev classification input is $0.042 per million tokens. The practical answer is to move the relevance decision upstream: screen each candidate title and snippet with Jev, and send only retained candidates to the synthesis model.
We built this pipeline end-to-end and measured it on live data: 2,000 candidate articles across two real news pools, 60 briefings from three synthesis models, and a blind panel that scored filtered versus raw briefings without knowing which was which. At the default 0.70 relevance threshold, screening cut downstream input tokens by 16–18%; five of six blind-panel quality cells were statistically indistinguishable from raw, and the sixth showed a small negative delta. Stricter thresholds screened out 30–52% of candidates. At frontier pricing, this nets roughly $10–17 per 1,000 screened 200-article batches per synthesis model.
The cost of raw keyword retrieval
One 200-article candidate batch costs a frontier model roughly 35,000–45,000 prompt tokens, and a full 1,000-article pool costs roughly 174,000–223,000 per pass. A scheduled daily pipeline pays that on every run. The cost is not only financial. Tangential commentary and syndication duplicates also consume the model's attention budget and slow generation, because time-to-first-token scales with context length.
Three questions determine whether screening helps: how much of a real keyword-search pool is irrelevant, how many tokens removing it saves, and whether the briefing degrades when filtered. This article measures all three on live data.
Follow along with a free key. Free covers 250 requests/day with a 30-day lookback.
The core idea: screen candidates before synthesis
TypeSafe Jev is a classification API priced at $0.042 per million input tokens, evaluated here with typesafe-sdk 0.6.0 . For each candidate, the pipeline sends a compact screening prompt containing the title plus a description snippet and one Boolean relevance question through the Noul primitive. Jev returns a probability between 0 and 1. Candidates at or above the threshold pass through with headline, metadata, and Currents description. The rest never reach the synthesis model.
In practice, this means the relevance decision is made at classification cost instead of synthesis cost.
How the benchmark was measured
The benchmark used two live news pools of 1,000 candidates each, pulled from Currents at benchmark time. No hand-picked sample.
The first pool was geopolitics: U.S.–Greenland sovereignty and Arctic defense developments. The second was financial intelligence: AST SpaceMobile ($ASTS) constellation deployment, technical milestones, and capital structure. Each 1,000-candidate topic pool was assembled from 12 topic-specific query variants across monthly windows and deduplicated by article ID; individual articles were not manually selected.
Each topic was split into five batches of 200 candidates, and each batch was synthesized twice by the same model, once from the raw packet and once from the Jev-filtered packet. That produced 60 briefings (2 topics × 5 batches × 3 models × 2 arms). Synthesis models: claude-sonnet-4-6 , gemini-3.8-flash , glm-5.3-flash .
A blind panel of four judge models, none of them the author of a scored briefing, scored each pair on coverage, faithfulness, unsupported-claims count, and usefulness. Labels were randomized so judges could not identify the filtered briefing. We collected 82 judge verdicts across 30 pairs. The Claude judge completed 12 of 20 assignments before its provider quota interrupted the session, and that partial coverage is reflected in the panel aggregates below.
Every token count in this article comes from provider usage fields: Jev screening tokens, synthesis prompt tokens, and output tokens. No tokenizer estimates. Synthesis ran at temperature 0 to reduce sampling variability; each filtered/raw comparison used the same model and prompt structure, differing only in the candidate packet.
Quality deltas are aggregated per author-model as filtered − raw and reported with 95% confidence intervals from a batch-level bootstrap (10 batches total).
Measured token reduction on live keyword pools
Input-token reduction at the default threshold
Across both pools, Jev retained about 82% of candidates and cut synthesis input by roughly a sixth. (Prompt tokens in the table are per screened batch, summed across the three synthesis models.) Screening itself averaged 546 input tokens per article, for 1,092,430 tokens across the 2,000 screenings. That works out to about $0.000023 per article at Jev's $0.042/MTok rate.
These results suggest live keyword pools for precise queries contain less noise than commonly assumed. When the query names a specific entity or event, most returned articles are on-topic, so the filtering gain at the default threshold is moderate but consistent.
Candidate reduction across threshold settings
Jev returns a probability for every article, so the cutoff is a configurable parameter. The 0.70 default retains about 82% of candidates. Higher thresholds remove more candidates:
At 0.85 the same pools shed 30–35% of candidates; at 0.95, over half. Higher cutoffs trade edge-case coverage for savings, so the threshold should match the recall margin the application requires.
Measured briefing quality after filtering
Token savings are only useful if briefing quality is maintained. The blind panel therefore scored every pair with author and judge roles kept separate and labels randomized.
Panel utility deltas (filtered − raw) per synthesis model, both topics:
Five of six cells cross zero: the panel could not reliably distinguish filtered from raw. The exception is small and negative; glm-5.3-flash briefings on the Greenland topic scored −1.2 points under filtering, a detectable but minor dip on that cell. In three of the six cells, judges slightly preferred the filtered version.
When filtering reduces cost, and when it increases cost
Screening adds a new cost term; the net effect depends on synthesis pricing. Screening a 200-article batch costs about $0.0046 (200 articles × 546 tokens × $0.042/MTok), or $4.59 per 1,000 screened batches.
Net savings per 1,000 screened batches at the 0.70 threshold, priced per synthesis model. Measured token removal per run: 7,104 tokens for claude-sonnet-4-6 (the benchmark author); 6,689 for the six-cell pool mean, which the GPT rows use because no GPT model was a benchmark author:
At frontier price points ($2.00–3.00/MTok), screening nets roughly $10–17 per 1,000 screened batches for a single-model pipeline; a pipeline that briefs with all three benchmark models from one screening realizes roughly three times the LLM savings against the same Jev cost. At the current mini-class price (GPT-5 Mini, $0.25/MTok), the 0.70 threshold is net negative (−$2.92); a stricter threshold would be required, and the exact break-even point depends on the token lengths of the discarded candidates.
The practical implication is that semantic pre-filtering is a frontier-model optimization. Frontier models are the tier where input length translates most directly into cost.
Example implementation in Python
The pipeline fits in one script: search Currents, screen each candidate with Jev, brief with Claude. page_size is set to 20 for compatibility with Currents API free-tier developer accounts. Paid tiers accommodate higher candidate volumes, and screening calls parallelize cleanly.
import os
import requests
from typesafe_sdk import TypeSafeClient , Noul
import anthropic
# 1. Initialize Clients
CURRENTS_API_KEY = os . environ [ "CURRENTS_API_KEY" ]
TYPESAFE_API_KEY = os . environ [ "TYPESAFE_API_KEY" ]
ANTHROPIC_API_KEY = os . environ [ "ANTHROPIC_API_KEY" ]
typesafe = TypeSafeClient ( api_key = TYPESAFE_API_KEY )
claude = anthropic . Anthropic ( api_key = ANTHROPIC_API_KEY )
# 2. Retrieve Candidate Articles from Currents News API
# Note: Currents API supports Bearer token authentication.
query = "AST SpaceMobile"
response = requests . get (
"https://api.currentsapi.services/v1/search" ,
headers = { "Authorization" : f "Bearer { CURRENTS_API_KEY } " },
params = {
"keywords" : query ,
"language" : "en" ,
"page_size" : 20 ,
},
timeout = 10 ,
)
response . raise_for_status ()
candidates = response . json () . get ( "news" , [])
# 3. Pre-Filtering with TypeSafe System One (Jev)
relevance_question = Noul (
instructions = (
"Does this article provide substantive primary reporting, financial analysis, "
"or technical updates directly about AST SpaceMobile ($ASTS)? Answer NO to generic "
"watchlists, unrelated company news, or shareholder litigation solicitations."
)
)
retained_articles = []
for article in candidates :
title = article . get ( "title" , "" ) . strip ()
summary = article . get ( "description" , "" ) . strip ()[: 800 ]
state_text = f "Title: { title } \n Summary: { summary } "
judgment = typesafe . system_one (
state = state_text ,
questions = { "is_relevant" : relevance_question },
)
prob = judgment . nouls [ "is_relevant" ] . noul
if prob >= 0.70 :
retained_articles . append ( article )
print ( f "Retained { len ( retained_articles ) } of { len ( candidates ) } candidates." )
# 4. Synthesize Briefing with Frontier Model
context_blocks = []
for idx , art in enumerate ( retained_articles , 1 ):
context_blocks . append (
f "--- Article { idx } --- \n "
f "Title: { art . get ( 'title' ) } \n "
f "URL: { art . get ( 'url' ) } \n "
f "Publishe
[truncated]
Jev outputs calibrated probabilities, not deterministic guarantees. On these pools, 0.70 was conservative: five of six blind-panel cells were statistically indistinguishable from raw, and one showed a small negative delta. Applications that cannot tolerate false negatives in open-domain retrieval should use 0.50–0.60 and accept smaller savings.
The threshold only matters relative to the synthesis model's price. The economics table above is the decision tool: frontier synthesis justifies 0.70; at mini-class prices 0.70 is net negative and a stricter threshold would be required — the exact break-even depends on the token lengths of the discarded candidates — so mini-class pipelines may not need the filter at all.
If the pipeline relies on provider prompt caching with static context prefixes, check whether dynamic pre-filtering changes what is cached. For multi-entity monitoring, where each user watches a different query, there is no cache to preserve and filtering is a direct saving.
Screening verdicts can be cached across scheduled runs for repeated (article, question) pairs. The measured numbers are all cold-start.
Keyword retrieval returns recall-oriented results, and unfiltered synthesis bills all of them at frontier input prices. The measured two-tier setup (Currents News API for recall, TypeSafe Jev for semantic pre-filtering) cut briefing input by 16–18% at a conservative threshold, and stricter thresholds screened out 30–52% of candidates, with five of six blind-panel quality cells statistically indistinguishable from ra

[truncated]
