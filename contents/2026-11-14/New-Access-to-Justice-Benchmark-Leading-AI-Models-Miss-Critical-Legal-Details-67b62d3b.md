---
source: "https://epyhia.com/new-access-to-justice-benchmark-finds-leading-ai-models-miss-critical"
hn_url: "https://news.ycombinator.com/item?id=50497161"
title: "New 'Access to Justice' Benchmark: Leading AI Models Miss Critical Legal Details"
article_title: "New 'Access to Justice' Benchmark Finds Leading AI Models Miss Critical Legal Details Across 252 Scenarios — epyhia.com"
image: ""
author: "arbuge"
captured_at: "2026-11-14T14:48:45Z"
capture_tool: "hn-digest"
hn_id: 50497161
score: 3
comments: 0
posted_at: "2026-11-14T14:10:31Z"
tags:
  - hacker-news
---

# New 'Access to Justice' Benchmark: Leading AI Models Miss Critical Legal Details

- HN: [50497161](https://news.ycombinator.com/item?id=50497161)
- Source: [epyhia.com](https://epyhia.com/new-access-to-justice-benchmark-finds-leading-ai-models-miss-critical)
- Score: 3
- Comments: 0
- Posted: 2026-11-14T14:10:31Z

## Translation

Title: New 'Access to Justice' Benchmark: Leading AI Models Miss Critical Legal Details
Article title: New 'Access to Justice' Benchmark Finds Leading AI Models Miss Critical Legal Details Across 252 Scenarios — epyhia.com
Description: Vals AI and NextLadder Ventures released the Access to Justice Benchmark on November 10, evaluating 23 AI models on 252 simulated legal-aid conversations. No model reached the 80% threshold the…

Article text:
epyhia.com Unpacking AI Hype
Separating AI hype from reality.
The day's AI news, with the hype taken out: one short summary each morning.
New 'Access to Justice' Benchmark Finds Leading AI Models Miss Critical Legal Details Across 252 Scenarios
November 14, 2026 · Alibaba · xAI · Anthropic
Vals AI and NextLadder Ventures released the Access to Justice Benchmark on November 10, evaluating 23 AI models on 252 simulated legal-aid conversations. No model reached the 80% threshold the organizers set for a reliable first point of contact. The top performer, Alibaba's Qwen 3.8 Max, scored 76.4% [1] [3] .
The benchmark was built to test how well general-purpose models handle the civil legal problems low-income Americans actually face — evictions, child custody, debt collection, domestic violence filings — rather than the corporate legal work most legal AI benchmarks measure [1] [4] . Six legal aid experts in California, Louisiana, Virginia, New York, Maryland and Texas adapted real queries from legal aid organizations into scenarios with background facts withheld unless the model asked [4] [7] .
Models explain law but miss local help
Every model met 78–87% of criteria for general legal explanations, but only 30–50% for identifying specific resources such as a local legal aid program, court self-help center or particular form [1] [7] . Twenty-one of 23 models fell back on national directories — LawHelp.org, the ABA's Find Legal Help, 211 — in at least 70% of their final answers instead of naming the relevant local office [4] [7] .
"Models read enough law to explain it, but rarely open the local self-help, legal aid or court pages that would let them name the specific program or office a person should go to next," Vals AI's report states [7] . Statutes, court rules and case law made up 42% of pages models read; legal aid and nonprofit self-help sites accounted for 10% [4] [7] .
Free-tier models keep pace with premium
Free-tier chatbot models averaged 69.7% and premium ones 70.4%, a gap of 0.7 percentage points [1] [7] . Anthropic's free-tier Claude Sonnet 5.5 scored 1.2 points higher than its premium Claude Opus 5.5. Only xAI showed a clear premium advantage: Grok 4.7 beat Grok 4.6 by about 2.3–2.4 points, larger than the margin of error [1] [7] . The tier assignments are Vals AI's estimates, not confirmed provider mappings, and API tests cannot fully capture consumer chatbot behavior [4] .
Wide variance across legal topics
Housing and landlord-tenant questions scored highest across all models at 75.4%. Public benefits ranked lowest at 57.3%, though that category contained only five scenarios [1] [4] . Family law, the largest topic with 110 scenarios, averaged 71.2% but varied sharply by subtopic: visitation (61.3%) and kinship caregiving with child protective services (63.0%) scored lowest, while guardianship, divorce finances, family-member adoptions and domestic violence fell between 66% and 68% [1] [4] .
In a Louisiana small-claims scenario, a New Orleans parent asked about suing the other parent over a verbal repayment promise. Most models cited the correct court but incorrectly cited a 10-year prescriptive period; the correct period for money lent is three years under La. C.C. art. 3494, as the benchmark's methodology page clarifies [7] . The press release and some coverage repeated the 10-year figure [1] [5] . Only 8 of 23 models asked whether the debt involved real estate — which decides whether a verbal agreement can be enforced at all. The expert rubric evaluated whether models warned that the claim is lost once the 10-year period runs out (the figure most models cited); 4 of 23 did so. Only 3 directed the user to the court clerk's website for claim and fee-waiver forms [1] [5] .
In a New York City constructive eviction scenario, a tenant who had moved out asked how to recover an unreturned deposit. Nearly every model cited the landlord's 14-day deadline and laid out a demand letter followed by small claims filing. Fourteen models were marked down for pushing that out-of-court step, which the expert rubric warned against. Only 4 referred the tenant to NYC's Right to Counsel program; most offered general hotlines like 311 or Legal Services NYC [1] [5] .
The benchmark is part of a NextLadder Ventures effort to accelerate AI tools for low-income Americans. The Legal Services Corporation's 2022 Justice Gap study found low-income Americans did not get any or enough legal help for 92% of their civil legal problems; 46% of those who didn't seek help cited cost [1] [8] . NextLadder, backed by over $1 billion, aims to empower millions by 2040 [1] [8] . Vals AI, a Series A company backed by a16z, provides independent benchmarking used by frontier labs [1] [9] .
The full dataset contains 504 scenarios split evenly between a public test set and a private validation set; only the 252 public scenarios were used for published results [4] [7] . Models were run through their APIs, which approximates but does not reproduce the consumer chatbot experience [3] . Content-filter refusals counted as zero — Muse Spark 1.3 scored 41.5% on criminal and records questions because its provider blocked five scenarios on every retry [7] . Vals AI and NextLadder say they are collaborating on a fuller report and a practitioner brief to help frontline workers understand remaining gaps [1] [4] .
Pipedrive is a sales CRM built around a visual pipeline: see every deal at a glance, track calls and emails, automate follow-ups and forecast revenue, so small sales teams spend less time on admin and more time selling.
China's Singles Day Tests AI Commerce Claims at Premier Shopping Festival November 11, 2026
Harvard Research: Only 2.9% of US Consumers Pay for Generative AI Services November 14, 2026
Grok Build Coding Agent Exposes 35,000 Apps Including SpaceX Data November 10, 2026
AI Data Centers Turn to Decommissioned Jet Engines as Grid Queues Stretch Years November 4, 2026
New Test of Leading AI Models' Legal Aid Advice Finds ... (prnewswire.com)
New Test of Leading AI Models' Legal Aid Advice Finds ... (morningstar.com)
Access to Justice Leaderboard & Scores — November 2026 (benchlm.ai)
Vals AI Legal Aid Benchmark Finds Models Fall Short on ... (unite.ai)
AI Chatbots Explain the Law Well but Often Fail to Point ... (agentlocker.ai)
Leading AI Models Still Missing Critical Details When Giving Financial Advice, New Research Finds (prnewswire.com)
Access to Justice Benchmark Leaderboard and Methodology | Vals AI (vals.ai)
New Test of Leading AI Models’ Legal Aid Advice Finds Key Gaps Still to be Closed Across Issues Commonly Faced by Low-Income Americans — NextLadder (nextladder.com)
AI Evaluation Methodology | Vals AI (vals.ai)
Some links on this site may be affiliate links: if you buy through them we may earn a commission, at no cost to you. It never affects what we write.

## Original Extract

Vals AI and NextLadder Ventures released the Access to Justice Benchmark on November 10, evaluating 23 AI models on 252 simulated legal-aid conversations. No model reached the 80% threshold the…

epyhia.com Unpacking AI Hype
Separating AI hype from reality.
The day's AI news, with the hype taken out: one short summary each morning.
New 'Access to Justice' Benchmark Finds Leading AI Models Miss Critical Legal Details Across 252 Scenarios
November 14, 2026 · Alibaba · xAI · Anthropic
Vals AI and NextLadder Ventures released the Access to Justice Benchmark on November 10, evaluating 23 AI models on 252 simulated legal-aid conversations. No model reached the 80% threshold the organizers set for a reliable first point of contact. The top performer, Alibaba's Qwen 3.8 Max, scored 76.4% [1] [3] .
The benchmark was built to test how well general-purpose models handle the civil legal problems low-income Americans actually face — evictions, child custody, debt collection, domestic violence filings — rather than the corporate legal work most legal AI benchmarks measure [1] [4] . Six legal aid experts in California, Louisiana, Virginia, New York, Maryland and Texas adapted real queries from legal aid organizations into scenarios with background facts withheld unless the model asked [4] [7] .
Models explain law but miss local help
Every model met 78–87% of criteria for general legal explanations, but only 30–50% for identifying specific resources such as a local legal aid program, court self-help center or particular form [1] [7] . Twenty-one of 23 models fell back on national directories — LawHelp.org, the ABA's Find Legal Help, 211 — in at least 70% of their final answers instead of naming the relevant local office [4] [7] .
"Models read enough law to explain it, but rarely open the local self-help, legal aid or court pages that would let them name the specific program or office a person should go to next," Vals AI's report states [7] . Statutes, court rules and case law made up 42% of pages models read; legal aid and nonprofit self-help sites accounted for 10% [4] [7] .
Free-tier models keep pace with premium
Free-tier chatbot models averaged 69.7% and premium ones 70.4%, a gap of 0.7 percentage points [1] [7] . Anthropic's free-tier Claude Sonnet 5.5 scored 1.2 points higher than its premium Claude Opus 5.5. Only xAI showed a clear premium advantage: Grok 4.7 beat Grok 4.6 by about 2.3–2.4 points, larger than the margin of error [1] [7] . The tier assignments are Vals AI's estimates, not confirmed provider mappings, and API tests cannot fully capture consumer chatbot behavior [4] .
Wide variance across legal topics
Housing and landlord-tenant questions scored highest across all models at 75.4%. Public benefits ranked lowest at 57.3%, though that category contained only five scenarios [1] [4] . Family law, the largest topic with 110 scenarios, averaged 71.2% but varied sharply by subtopic: visitation (61.3%) and kinship caregiving with child protective services (63.0%) scored lowest, while guardianship, divorce finances, family-member adoptions and domestic violence fell between 66% and 68% [1] [4] .
In a Louisiana small-claims scenario, a New Orleans parent asked about suing the other parent over a verbal repayment promise. Most models cited the correct court but incorrectly cited a 10-year prescriptive period; the correct period for money lent is three years under La. C.C. art. 3494, as the benchmark's methodology page clarifies [7] . The press release and some coverage repeated the 10-year figure [1] [5] . Only 8 of 23 models asked whether the debt involved real estate — which decides whether a verbal agreement can be enforced at all. The expert rubric evaluated whether models warned that the claim is lost once the 10-year period runs out (the figure most models cited); 4 of 23 did so. Only 3 directed the user to the court clerk's website for claim and fee-waiver forms [1] [5] .
In a New York City constructive eviction scenario, a tenant who had moved out asked how to recover an unreturned deposit. Nearly every model cited the landlord's 14-day deadline and laid out a demand letter followed by small claims filing. Fourteen models were marked down for pushing that out-of-court step, which the expert rubric warned against. Only 4 referred the tenant to NYC's Right to Counsel program; most offered general hotlines like 311 or Legal Services NYC [1] [5] .
The benchmark is part of a NextLadder Ventures effort to accelerate AI tools for low-income Americans. The Legal Services Corporation's 2022 Justice Gap study found low-income Americans did not get any or enough legal help for 92% of their civil legal problems; 46% of those who didn't seek help cited cost [1] [8] . NextLadder, backed by over $1 billion, aims to empower millions by 2040 [1] [8] . Vals AI, a Series A company backed by a16z, provides independent benchmarking used by frontier labs [1] [9] .
The full dataset contains 504 scenarios split evenly between a public test set and a private validation set; only the 252 public scenarios were used for published results [4] [7] . Models were run through their APIs, which approximates but does not reproduce the consumer chatbot experience [3] . Content-filter refusals counted as zero — Muse Spark 1.3 scored 41.5% on criminal and records questions because its provider blocked five scenarios on every retry [7] . Vals AI and NextLadder say they are collaborating on a fuller report and a practitioner brief to help frontline workers understand remaining gaps [1] [4] .
Pipedrive is a sales CRM built around a visual pipeline: see every deal at a glance, track calls and emails, automate follow-ups and forecast revenue, so small sales teams spend less time on admin and more time selling.
China's Singles Day Tests AI Commerce Claims at Premier Shopping Festival November 11, 2026
Harvard Research: Only 2.9% of US Consumers Pay for Generative AI Services November 14, 2026
Grok Build Coding Agent Exposes 35,000 Apps Including SpaceX Data November 10, 2026
AI Data Centers Turn to Decommissioned Jet Engines as Grid Queues Stretch Years November 4, 2026
New Test of Leading AI Models' Legal Aid Advice Finds ... (prnewswire.com)
New Test of Leading AI Models' Legal Aid Advice Finds ... (morningstar.com)
Access to Justice Leaderboard & Scores — November 2026 (benchlm.ai)
Vals AI Legal Aid Benchmark Finds Models Fall Short on ... (unite.ai)
AI Chatbots Explain the Law Well but Often Fail to Point ... (agentlocker.ai)
Leading AI Models Still Missing Critical Details When Giving Financial Advice, New Research Finds (prnewswire.com)
Access to Justice Benchmark Leaderboard and Methodology | Vals AI (vals.ai)
New Test of Leading AI Models’ Legal Aid Advice Finds Key Gaps Still to be Closed Across Issues Commonly Faced by Low-Income Americans — NextLadder (nextladder.com)
AI Evaluation Methodology | Vals AI (vals.ai)
Some links on this site may be affiliate links: if you buy through them we may earn a commission, at no cost to you. It never affects what we write.
