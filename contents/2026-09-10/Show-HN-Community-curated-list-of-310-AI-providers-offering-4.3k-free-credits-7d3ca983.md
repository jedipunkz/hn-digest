---
source: "https://www.uprouter.online/"
hn_url: "https://news.ycombinator.com/item?id=49641128"
title: "Show HN: Community curated list of 310 AI providers offering $4.3k free credits"
article_title: "Free AI API Key: 310 AI API routers, free tiers & one-key gateway | UPROUTER"
image: "https://uprouter.online/opengraph-image?8989f339e6ef9d7d"
author: "opensrcme"
captured_at: "2026-09-10T10:45:51Z"
capture_tool: "hn-digest"
hn_id: 49641128
score: 3
comments: 0
posted_at: "2026-09-10T10:00:54Z"
tags:
  - hacker-news
---

# Show HN: Community curated list of 310 AI providers offering $4.3k free credits

- HN: [49641128](https://news.ycombinator.com/item?id=49641128)
- Source: [www.uprouter.online](https://www.uprouter.online/)
- Score: 3
- Comments: 0
- Posted: 2026-09-10T10:00:54Z

## Translation

Title: Show HN: Community curated list of 310 AI providers offering $4.3k free credits
Article title: Free AI API Key: 310 AI API routers, free tiers & one-key gateway | UPROUTER
Description: Find a free AI API key for any major model. 310 AI API routers & 490 models compared on free credits, rate limits & live status — 215 free tiers worth $4.3k. No pay-to-rank; route it all through one OpenAI- and Claude-compatible gateway.
HN text: Hey people! I have been compiling a directory of AI API providers which provide a combined total of $4.3k in free API credits and I'm super proud of it so far. It serves as a community-curated directory. Everyone can review, upvote and edit API providers to keep the directory honest and accurate. After tons of hours I've managed to compile a database of 310 API providers. The directory also allows you to link all 310 API providers, to create one OpenAI- and Claude-compatible gateway then route any model, through any vendor - via a single URL with failover aliases, daily spend caps and per-request compute metering you can audit line by line. I have not found another project that offers an organized, community-curated system for AI API providers on the same scale as this one. There is a submission feature so you can submit your favorite AI API providers for free.

Article text:
Free AI API Key: 310 AI API routers, free tiers & one-key gateway | UPROUTER UPROUTER . ONLINE Directory Models Compare Journal Blog Earn Submit feedback Feedback >_ search ⌘K Sign in Open Connect live uprouter — AI API router and gateway
Find the best AI API routers.
Use them all with one API key.
UPROUTER is an OpenAI- and Claude-compatible gateway: plug in your API keys from 310 routers (eg. OpenRouter, AWS Bedrock, Alibaba Cloud), then route every request with any model, through any vendor — via a single URL with failover aliases, daily spend caps and per-request compute metering you can audit line by line.
✓ BYOK keys — AES-256-GCM encrypted
✓ failover aliases, cheapest-first
✓ daily caps with soft-limit headers
✓ per-request compute metering
$ curl https://uprouter.online/api/connect/v1/chat/completions \
-H " Authorization: Bearer upr_live_… " \
-d '{ "model": "sandbox-echo" '}'
← 200 OK · 412ms · 0.0002 compute · routed via sandbox-echo
audit log: tokens in/out, latency, cost — per request
{
"id": "chatcmpl_9f2…" ,
"model": "sandbox-echo" ,
"routed_via": "sandbox-echo" ,
"usage": { "compute": 0.0002 , "tokens": "42/310" }
} endpoint /api/connect/v1 auth Bearer upr_live_… metering compute / request > routers 310 > free-tier 215 > connect 211 > live up 116 → > models 490 > free quota 215 refer a developer → + 50 compute
they get 10 — you get 50 /earn#refer →
$ 310 routers & providers · 490 models tracked · $4.3k free credits
Every entry answers to the community
The index is expansive; the community keeps it honest. Votes, reviews and verified corrections feed straight back into rankings and data confidence.
Vote Upvote routers that deliver, downvote the ones that don’t — the score re-ranks the index.
Review Evidence-based reviews with verification tags. Ratings, not vibes.
Verify Corrections that survive review earn compute — and flip the entry to verified.
Every router page has an edit button — propose corrections to descriptions, prices and free tiers. Editors review each proposal before it goes live, and accepted edits earn compute.
Propose Open a router page, hit the edit button, and detail your proposed corrections.
Review Editors rigorously verify every proposed correction against primary sources.
Live Accepted edits are published instantly, and you earn compute for your work.
$ 310 entr ies match — showing 1 – 24 · sort: Free credits ↓
Official free tier up 279 ms updated 10h ago inference.api.nscale.com source · nscale.com/services/ai-services overview Nscale is a European AI hyperscaler providing Serverless Inference for large-scale AI deployment. It offers an OpenAI-compatible API to leading open models like Llama and DeepSeek. New users receive a $5 signup credit to start prototyping without initial cost.
New users get a one-time $5 free credit to explore the models; after that you add a card to buy more. The core is pay-as-you-go (e.g. GPT-OSS 120B ~$0.1 in / $0.4 out per 1M tokens), with no rate limits and no cold starts.
Free relay up 385 ms updated 10h ago api.bluesminds.com source · api.bluesminds.com overview BluesMinds is a primarily-paid AI relay reselling access to Claude and GPT at significant discounts. It offers a large free credit bonus ($500) for accounts linked to aged GitHub profiles, and provides free daily calls for specialized models like GPT-4.1.
Up to ~$500 in credits for aged GitHub accounts per community reports.
Paid API unknown updated 3d ago cloud.google.com source · cloud.google.com/vertex-ai/generative-ai/pricing overview Google Cloud's enterprise-grade AI platform (Vertex AI) providing managed access to Gemini models and third-party models like Llama and Mistral. Offers advanced features like grounding with Google Search, model tuning, and provisioned throughput for predictable performance.
No public free-tier information found in the source reference; treat free-tier availability as unverified rather than confirmed absent.
Free relay up 997 ms updated 10h ago agentrouter.org source · agent-router.org/pricing overview Agent Router is an open-source, public-welfare AI API aggregation platform designed to unify access to agents and MCP tools. It is 'forever free' for both users and developers, providing a discovery registry for skills (ClawHub) and a live A2A directory.
Offers $100-$200 free credit bonus for new accounts to bootstrap agent usage.
Paid API up 65 ms updated 10h ago aws.amazon.com source · aws.amazon.com/bedrock/pricing overview Amazon Bedrock is a fully managed service that offers a choice of high-performing foundation models from leading AI companies like Anthropic, Meta, and Mistral via a single API. It is primarily pay-as-you-go with various service tiers (Standard, Flex, Priority).
Bedrock itself has no permanent free tier and bills per use from the first call; new AWS accounts opened after 2025-07-15 receive $200 in credits (expire in 6 months) usable across services including Bedrock.
Search / Audio API unknown updated 3d ago deepgram.com source · deepgram.com/pricing overview Deepgram is a specialized speech-to-text and text-to-speech API provider. It offers high-performance audio models like Nova-3 and Flux, with real-time streaming and batch processing capabilities. Authentication is via API key created in its own dashboard.
$200 free credit on sign-up with no expiration.
Official free tier down 6002 ms updated 10h ago console.mistral.ai source · mistral.ai/pricing overview Mistral AI's official platform (La Plateforme) provides access to their open and proprietary models via an OpenAI-compatible API. It features a free 'Experiment' tier for developers and pay-as-you-go pricing for production workloads.
Community reported approximate monthly value of the Experiment free tier for developers.
Official free tier up 718 ms updated 10h ago token.sensenova.cn source · sensenova.ai/token-plan overview token.sensenova.cn is the official Token Plan for SenseTime's SenseNova multimodal AI platform. During its public beta, it offers generous free quotas for models like SenseNova 6.8 Flash Lite and U1 Fast, suitable for complex office workflows.
Free public beta (extended to end of July): 1,500 calls/5h per model (verified: 60k points/5h on token-plan page); ~7,200 calls/day ~ $200/mo equivalent at ~1K tokens/call. Phone signup, up to 20 API keys.
Free product up 1223 ms updated 10h ago longcat.chat source · longcat.chat/platform overview LongCat is Meituan's large model API platform (longcat.chat/platform), offering OpenAI/Anthropic-compatible endpoints. During its public beta phase in 2026, it provides generous daily free token allowances for models including LongCat-Flash and LongCat-2.0, with a context window of 131K tokens.
Free during public beta: initial ~5M tokens/day (per API changelog), activity can raise it up to ~120M/day; ~$150/mo equivalent at blended $1/1M tokens. Phone signup, key in ~3 min; paid top-ups not yet supported.
Official free tier up 19 ms updated 10h ago console.x.ai source · x.ai/api overview The xAI Console is the official developer interface for Elon Musk's Grok models. It offers a usage-based API with prepaid credits, prioritizing performance and direct access to their flagship large language models.
Community reported trial credit amount for new developer accounts.
Search / Audio API unknown updated 2d ago you.com source · you.com/pricing overview You.com provides a specialized Web Search API for AI agents and LLM applications. It offers real-time web intelligence and structured search results (snippets, full text). Features a consumption-based pricing model and generous onboarding credits for new developers.
No public free-tier information found in the source reference; treat free-tier availability as unverified rather than confirmed absent.
Free relay up 204 ms updated 10h ago api.gemai.cc source · api.gemai.cc overview Hajimi API (api.gemai.cc) is a Chinese third-party relay aggregation platform. It specializes in routing requests to Claude, Gemini, and GPT models through a unified OpenAI-compatible endpoint, primarily catering to users requiring simplified access to international models.
New users get 100 units of balance by default (community figure ~¥500 (≈$74.39)), QQ-email signup only; there is a daily check-in farming crowd; 'limited perk' zero-priced models can be delisted anytime.
Official free tier up 1972 ms updated 10h ago modelscope.cn source · modelscope.ai/docs/model-service/API-Inference/limits overview ModelScope (modelscope.cn), backed by Alibaba Cloud, is an open-source model community and inference platform. It provides a free API tier allowing users to make up to 2,000 OpenAI-compatible calls per day across a wide range of open-source models (Qwen, DeepSeek, GLM).
2,000 API-Inference calls/day per user (verified, aggregated across models; ~200-500/day per model), resets UTC+8 00:00, 429 when exceeded; requires Alibaba Cloud binding. ~$60/mo equivalent at ~1K tokens/call.
Search / Audio API unknown updated 2d ago gladia.io source · gladia.io/pricing overview Gladia is a speech-to-text API provider offering high-accuracy asynchronous and real-time transcription across 100+ languages. It features speaker diarization and automatic language detection, billed via a prepaid credit wallet system with a significant free starting grant.
50€ free credits (~$55 USD) one-time grant upon signup.
Free relay up 1320 ms updated 10h ago anyrouter.top source · anyrouter.top overview AnyRouter (anyrouter.top) is a public-welfare AI API relay primarily serving the mainland Chinese developer community. It specializes in routing Claude Code and other frontier models for free, supported by community sign-in bonuses and referral credits.
Grants ~$50-$100 in free credits on sign-up based on community reports.
Monitor / Directory unknown updated 3d ago context7.com source · context7.com/plans overview Context7 (powered by Upstash) is a documentation search API for LLMs, providing up-to-date context from libraries and codebases. It offers an anonymous free tier and a free API key for higher limits to support coding agents.
Value based on 5,000 included calls at $10 per 1,000 calls overage rate.
Official free tier up 200 ms updated 10h ago console.groq.com source · console.groq.com/docs/rate-limits overview Groq Cloud provides ultra-fast LLM inference using its proprietary LPU (Language Processing Unit) hardware. It offers an OpenAI-compatible API with a substantial free tier for developers to test and build applications at scale.
Community reported approximate monthly value of free tier rate limits for individual developers.
Official free tier unknown updated 3d ago pioneer.ai source · pioneer.ai/pricing overview Pioneer AI, by Fastino Labs, provides an inference API for specialized tasks like data extraction and classification. It offers seat-based subscriptions with included platform credits and priority support.
$75 free usage credits — no credit card required
Commercial aggregator up 602 ms updated 10h ago aerolink.lat source · aerolink.lat/pricing overview AeroLink (aerolink.lat) is a commercial AI API gateway and relay service that routes multiple coding assistants and automation tools through a single key. It offers tiered subscription plans with rolling usage windows and claims access to high-end models like Claude Opus 4.8.
Starter plan offers a 1-week free trial included in the promo.
Official free tier up updated 3d ago app.baseten.co source · baseten.co overview Baseten is a robust model inference and deployment platform that enables enterprises to run custom and open-weights models in production. It offers high scalability, dedicated GPU infrastructure, and a developer-friendly API for deploying models from a comprehensive library.
New workspaces get $30 in trial credits billed by compute time (confirmed by pricing FAQ and 2026 credit trackers); converts to pay-as-you-go once spent. Startup program (up to $25K) is separate.
Official free tier up updated 3d ago cerebras.ai so

[truncated]

## Original Extract

Find a free AI API key for any major model. 310 AI API routers & 490 models compared on free credits, rate limits & live status — 215 free tiers worth $4.3k. No pay-to-rank; route it all through one OpenAI- and Claude-compatible gateway.

Hey people! I have been compiling a directory of AI API providers which provide a combined total of $4.3k in free API credits and I'm super proud of it so far. It serves as a community-curated directory. Everyone can review, upvote and edit API providers to keep the directory honest and accurate. After tons of hours I've managed to compile a database of 310 API providers. The directory also allows you to link all 310 API providers, to create one OpenAI- and Claude-compatible gateway then route any model, through any vendor - via a single URL with failover aliases, daily spend caps and per-request compute metering you can audit line by line. I have not found another project that offers an organized, community-curated system for AI API providers on the same scale as this one. There is a submission feature so you can submit your favorite AI API providers for free.

Free AI API Key: 310 AI API routers, free tiers & one-key gateway | UPROUTER UPROUTER . ONLINE Directory Models Compare Journal Blog Earn Submit feedback Feedback >_ search ⌘K Sign in Open Connect live uprouter — AI API router and gateway
Find the best AI API routers.
Use them all with one API key.
UPROUTER is an OpenAI- and Claude-compatible gateway: plug in your API keys from 310 routers (eg. OpenRouter, AWS Bedrock, Alibaba Cloud), then route every request with any model, through any vendor — via a single URL with failover aliases, daily spend caps and per-request compute metering you can audit line by line.
✓ BYOK keys — AES-256-GCM encrypted
✓ failover aliases, cheapest-first
✓ daily caps with soft-limit headers
✓ per-request compute metering
$ curl https://uprouter.online/api/connect/v1/chat/completions \
-H " Authorization: Bearer upr_live_… " \
-d '{ "model": "sandbox-echo" '}'
← 200 OK · 412ms · 0.0002 compute · routed via sandbox-echo
audit log: tokens in/out, latency, cost — per request
{
"id": "chatcmpl_9f2…" ,
"model": "sandbox-echo" ,
"routed_via": "sandbox-echo" ,
"usage": { "compute": 0.0002 , "tokens": "42/310" }
} endpoint /api/connect/v1 auth Bearer upr_live_… metering compute / request > routers 310 > free-tier 215 > connect 211 > live up 116 → > models 490 > free quota 215 refer a developer → + 50 compute
they get 10 — you get 50 /earn#refer →
$ 310 routers & providers · 490 models tracked · $4.3k free credits
Every entry answers to the community
The index is expansive; the community keeps it honest. Votes, reviews and verified corrections feed straight back into rankings and data confidence.
Vote Upvote routers that deliver, downvote the ones that don’t — the score re-ranks the index.
Review Evidence-based reviews with verification tags. Ratings, not vibes.
Verify Corrections that survive review earn compute — and flip the entry to verified.
Every router page has an edit button — propose corrections to descriptions, prices and free tiers. Editors review each proposal before it goes live, and accepted edits earn compute.
Propose Open a router page, hit the edit button, and detail your proposed corrections.
Review Editors rigorously verify every proposed correction against primary sources.
Live Accepted edits are published instantly, and you earn compute for your work.
$ 310 entr ies match — showing 1 – 24 · sort: Free credits ↓
Official free tier up 279 ms updated 10h ago inference.api.nscale.com source · nscale.com/services/ai-services overview Nscale is a European AI hyperscaler providing Serverless Inference for large-scale AI deployment. It offers an OpenAI-compatible API to leading open models like Llama and DeepSeek. New users receive a $5 signup credit to start prototyping without initial cost.
New users get a one-time $5 free credit to explore the models; after that you add a card to buy more. The core is pay-as-you-go (e.g. GPT-OSS 120B ~$0.1 in / $0.4 out per 1M tokens), with no rate limits and no cold starts.
Free relay up 385 ms updated 10h ago api.bluesminds.com source · api.bluesminds.com overview BluesMinds is a primarily-paid AI relay reselling access to Claude and GPT at significant discounts. It offers a large free credit bonus ($500) for accounts linked to aged GitHub profiles, and provides free daily calls for specialized models like GPT-4.1.
Up to ~$500 in credits for aged GitHub accounts per community reports.
Paid API unknown updated 3d ago cloud.google.com source · cloud.google.com/vertex-ai/generative-ai/pricing overview Google Cloud's enterprise-grade AI platform (Vertex AI) providing managed access to Gemini models and third-party models like Llama and Mistral. Offers advanced features like grounding with Google Search, model tuning, and provisioned throughput for predictable performance.
No public free-tier information found in the source reference; treat free-tier availability as unverified rather than confirmed absent.
Free relay up 997 ms updated 10h ago agentrouter.org source · agent-router.org/pricing overview Agent Router is an open-source, public-welfare AI API aggregation platform designed to unify access to agents and MCP tools. It is 'forever free' for both users and developers, providing a discovery registry for skills (ClawHub) and a live A2A directory.
Offers $100-$200 free credit bonus for new accounts to bootstrap agent usage.
Paid API up 65 ms updated 10h ago aws.amazon.com source · aws.amazon.com/bedrock/pricing overview Amazon Bedrock is a fully managed service that offers a choice of high-performing foundation models from leading AI companies like Anthropic, Meta, and Mistral via a single API. It is primarily pay-as-you-go with various service tiers (Standard, Flex, Priority).
Bedrock itself has no permanent free tier and bills per use from the first call; new AWS accounts opened after 2025-07-15 receive $200 in credits (expire in 6 months) usable across services including Bedrock.
Search / Audio API unknown updated 3d ago deepgram.com source · deepgram.com/pricing overview Deepgram is a specialized speech-to-text and text-to-speech API provider. It offers high-performance audio models like Nova-3 and Flux, with real-time streaming and batch processing capabilities. Authentication is via API key created in its own dashboard.
$200 free credit on sign-up with no expiration.
Official free tier down 6002 ms updated 10h ago console.mistral.ai source · mistral.ai/pricing overview Mistral AI's official platform (La Plateforme) provides access to their open and proprietary models via an OpenAI-compatible API. It features a free 'Experiment' tier for developers and pay-as-you-go pricing for production workloads.
Community reported approximate monthly value of the Experiment free tier for developers.
Official free tier up 718 ms updated 10h ago token.sensenova.cn source · sensenova.ai/token-plan overview token.sensenova.cn is the official Token Plan for SenseTime's SenseNova multimodal AI platform. During its public beta, it offers generous free quotas for models like SenseNova 6.8 Flash Lite and U1 Fast, suitable for complex office workflows.
Free public beta (extended to end of July): 1,500 calls/5h per model (verified: 60k points/5h on token-plan page); ~7,200 calls/day ~ $200/mo equivalent at ~1K tokens/call. Phone signup, up to 20 API keys.
Free product up 1223 ms updated 10h ago longcat.chat source · longcat.chat/platform overview LongCat is Meituan's large model API platform (longcat.chat/platform), offering OpenAI/Anthropic-compatible endpoints. During its public beta phase in 2026, it provides generous daily free token allowances for models including LongCat-Flash and LongCat-2.0, with a context window of 131K tokens.
Free during public beta: initial ~5M tokens/day (per API changelog), activity can raise it up to ~120M/day; ~$150/mo equivalent at blended $1/1M tokens. Phone signup, key in ~3 min; paid top-ups not yet supported.
Official free tier up 19 ms updated 10h ago console.x.ai source · x.ai/api overview The xAI Console is the official developer interface for Elon Musk's Grok models. It offers a usage-based API with prepaid credits, prioritizing performance and direct access to their flagship large language models.
Community reported trial credit amount for new developer accounts.
Search / Audio API unknown updated 2d ago you.com source · you.com/pricing overview You.com provides a specialized Web Search API for AI agents and LLM applications. It offers real-time web intelligence and structured search results (snippets, full text). Features a consumption-based pricing model and generous onboarding credits for new developers.
No public free-tier information found in the source reference; treat free-tier availability as unverified rather than confirmed absent.
Free relay up 204 ms updated 10h ago api.gemai.cc source · api.gemai.cc overview Hajimi API (api.gemai.cc) is a Chinese third-party relay aggregation platform. It specializes in routing requests to Claude, Gemini, and GPT models through a unified OpenAI-compatible endpoint, primarily catering to users requiring simplified access to international models.
New users get 100 units of balance by default (community figure ~¥500 (≈$74.39)), QQ-email signup only; there is a daily check-in farming crowd; 'limited perk' zero-priced models can be delisted anytime.
Official free tier up 1972 ms updated 10h ago modelscope.cn source · modelscope.ai/docs/model-service/API-Inference/limits overview ModelScope (modelscope.cn), backed by Alibaba Cloud, is an open-source model community and inference platform. It provides a free API tier allowing users to make up to 2,000 OpenAI-compatible calls per day across a wide range of open-source models (Qwen, DeepSeek, GLM).
2,000 API-Inference calls/day per user (verified, aggregated across models; ~200-500/day per model), resets UTC+8 00:00, 429 when exceeded; requires Alibaba Cloud binding. ~$60/mo equivalent at ~1K tokens/call.
Search / Audio API unknown updated 2d ago gladia.io source · gladia.io/pricing overview Gladia is a speech-to-text API provider offering high-accuracy asynchronous and real-time transcription across 100+ languages. It features speaker diarization and automatic language detection, billed via a prepaid credit wallet system with a significant free starting grant.
50€ free credits (~$55 USD) one-time grant upon signup.
Free relay up 1320 ms updated 10h ago anyrouter.top source · anyrouter.top overview AnyRouter (anyrouter.top) is a public-welfare AI API relay primarily serving the mainland Chinese developer community. It specializes in routing Claude Code and other frontier models for free, supported by community sign-in bonuses and referral credits.
Grants ~$50-$100 in free credits on sign-up based on community reports.
Monitor / Directory unknown updated 3d ago context7.com source · context7.com/plans overview Context7 (powered by Upstash) is a documentation search API for LLMs, providing up-to-date context from libraries and codebases. It offers an anonymous free tier and a free API key for higher limits to support coding agents.
Value based on 5,000 included calls at $10 per 1,000 calls overage rate.
Official free tier up 200 ms updated 10h ago console.groq.com source · console.groq.com/docs/rate-limits overview Groq Cloud provides ultra-fast LLM inference using its proprietary LPU (Language Processing Unit) hardware. It offers an OpenAI-compatible API with a substantial free tier for developers to test and build applications at scale.
Community reported approximate monthly value of free tier rate limits for individual developers.
Official free tier unknown updated 3d ago pioneer.ai source · pioneer.ai/pricing overview Pioneer AI, by Fastino Labs, provides an inference API for specialized tasks like data extraction and classification. It offers seat-based subscriptions with included platform credits and priority support.
$75 free usage credits — no credit card required
Commercial aggregator up 602 ms updated 10h ago aerolink.lat source · aerolink.lat/pricing overview AeroLink (aerolink.lat) is a commercial AI API gateway and relay service that routes multiple coding assistants and automation tools through a single key. It offers tiered subscription plans with rolling usage windows and claims access to high-end models like Claude Opus 4.8.
Starter plan offers a 1-week free trial included in the promo.
Official free tier up updated 3d ago app.baseten.co source · baseten.co overview Baseten is a robust model inference and deployment platform that enables enterprises to run custom and open-weights models in production. It offers high scalability, dedicated GPU infrastructure, and a developer-friendly API for deploying models from a comprehensive library.
New workspaces get $30 in trial credits billed by compute time (confirmed by pricing FAQ and 2026 credit trackers); converts to pay-as-you-go once spent. Startup program (up to $25K) is separate.
Official free tier up updated 3d ago cerebras.ai so

[truncated]
