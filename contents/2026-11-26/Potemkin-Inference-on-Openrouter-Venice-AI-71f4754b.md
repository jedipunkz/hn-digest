---
source: "https://kendell.dev/blog/venicewrapper/"
hn_url: "https://news.ycombinator.com/item?id=50660720"
title: "Potemkin Inference on Openrouter (Venice AI)"
article_title: "Venice is a DeepInfra+ wrapper"
image: ""
author: "leafingensemble"
captured_at: "2026-11-26T23:30:42Z"
capture_tool: "hn-digest"
hn_id: 50660720
score: 2
comments: 0
posted_at: "2026-11-26T22:47:04Z"
tags:
  - hacker-news
---

# Potemkin Inference on Openrouter (Venice AI)

- HN: [50660720](https://news.ycombinator.com/item?id=50660720)
- Source: [kendell.dev](https://kendell.dev/blog/venicewrapper/)
- Score: 2
- Comments: 0
- Posted: 2026-11-26T22:47:04Z

## Translation

Title: Potemkin Inference on Openrouter (Venice AI)
Article title: Venice is a DeepInfra+ wrapper

Article text:
Venice is a DeepInfra+ wrapper Venice is a DeepInfra+ wrapper
Venice says its "private" models are self-hosted, but it provably wraps DeepInfra, Fireworks, Baseten, Parasail and fal.
Its website breaks EU, UK and California privacy law: ad trackers fire before you consent, and it ignores Global Privacy Control.
Its image editor undresses photos with no check for real people, breaking Minnesota's nudification ban and Texas's age verification law.
It silently swaps multiple models, selling Qwen Instruct for about 4× the price as "Thinking" and more.
Its encryption is mostly for show: the "encrypted" browser storage keeps its key right next to the data, and Venice and Phala can read your end-to-end encrypted prompts if they want to.
DeepSeek V4 Flash 0423 DeepSeek V4.1 Flash Gemma 4 31B GLM 5.3 Flash DeepSeek V4 Flash 0731 MiniMax M3 GLM 5.2 Gemma 4 26B A4B Other Aug Sep Oct Nov Weekly table Week of DeepSeek V4 Flash 0423 DeepSeek V4.1 Flash Gemma 4 31B GLM 5.3 Flash DeepSeek V4 Flash 0731 MiniMax M3 GLM 5.2 Gemma 4 26B A4B Other Jul 1, 2026 20B 0 21B 0 0 0 5B 4.1B 21B Jul 8, 2026 40B 0 26B 0 0 2.5B 6.3B 5.4B 25B Jul 15, 2026 31B 0 25B 0 0 13B 3.8B 3.8B 28B Jul 22, 2026 33B 0 20B 0 0 8.1B 11B 4.4B 40B Jul 29, 2026 109B 0 17B 0 3.7B 24B 27B 13B 71B Aug 5, 2026 112B 0 29B 0 31B 25B 17B 19B 84B Aug 12, 2026 70B 0 39B 0 27B 21B 28B 20B 68B Aug 19, 2026 77B 0 33B 0 22B 20B 17B 30B 63B Aug 26, 2026 76B 0 37B 27B 29B 19B 15B 19B 70B Sep 2, 2026 60B 0 40B 38B 15B 74B 13B 20B 121B Sep 9, 2026 66B 142B 56B 16B 24B 58B 20B 9.4B 119B Sep 16, 2026 79B 81B 51B 19B 38B 67B 31B 4.4B 141B Sep 23, 2026 78B 103B 26B 27B 18B 35B 23B 0 145B Sep 30, 2026 56B 193B 22B 10B 14B 0 2.8B 5.8B 146B Oct 7, 2026 45B 234B 18B 113B 19B 0 1.5B 0 188B Oct 14, 2026 84B 58B 12B 103B 82B 0 0 0 198B Oct 21, 2026 95B 179B 12B 153B 17B 31B 2.4B 0 183B Oct 28, 2026 93B 75B 2.5B 160B 53B 5.4B 0 0 162B Nov 4, 2026 74B 210B 955M 37B 74B 0 0 957M 159B Nov 11, 2026 172B 213B 4.4B 47B 93B 4.3B 2.3B 0 156B Nov 18, 2026 84B 88B 8.3B 94B 70B 18B 2.1B 0 178B Venice is a privacy friendly AI service, and they've made it look an awful lot like their (usual) "private" tier is the self host tier.
Venice runs leading open-source and third-party models across four privacy tiers: [...] Private (zero data retention on self-hosted open-source models), [...]
Venice offers four tiers of privacy: [...] Private (zero data retention, self-hosted open-source models),
Private - Zero data retention on self-hosted open-source models. Your prompts are never stored.
// archived webmcp via view-source:
description:
"List the AI models available on Venice from the public catalog: id, " +
"display name, type, privacy tier (private = zero data retention on " +
" Venice hardware , anonymized = third-party with metadata stripped), " +
"context size and traits. Filter by type or a name/id search.", Not to mention that Venice is an OpenRouter provider as of writing, and many ( 45.8% ) think most OpenRouter providers run models themselves.
Venice provably wraps other inference services
If you give Venice an image_url , it gets downloaded from another inference provider: Baseten , DeepInfra , Fireworks or Parasail for 14 of the 18 image models Venice listed on OpenRouter on November 21. OpenRouter even lists Qwen3.5 397B A17B, even though it's marked as anonymized and not private, and on November 21 its requests looped back into OpenRouter (32 of 40 requests). Between November 22, when I sent OpenRouter a preview of this post, and November 24, Venice moved 397B to DeepInfra, and GLM 5 "turbo", which looped back the same way, came off OpenRouter. Attested proof is available here (run 44936721959 on 85554d8 ).
Media models too. Hunyuan Image 3.0 runs on fal according to Venice's app config:
{
"id" : "hunyuan-image-v3" ,
"capabilities" : { "private" : true },
"apiRateLimitSize" : "fal-images" ,
"falRequestDefaults" : { "enable_safety_checker" : false }
} I discussed what "private" actually means with a Venice developer in the Venice Discord, who directed me to https://venice.ai/privacy . This page canonically defines "private" to include wrapping other inference services:
Private - Inference runs on Venice-controlled GPUs or zero-data-retention partner infrastructure . Privacy is enforced by contractual commitments — no prompt or response is stored.
I pinged the Venice developer about the other pages that define "private" to self hosted on November 17. As of November 25, 8 days and 2 pings later, Venice still hasn't fixed their plausible yet officially incorrect claims of self hosting.
Venice's moat is marketing, not privacy
Venice has a good thesis and good marketing to go with it. Nobody wants to be spied on by a big lab. And everyone likes inference subsidization, whether from staking VVV or DIEM.
But Venice's "private" models don't add anything to the market. If you're a consumer, you could just go to OpenRouter, top up $5, configure your account for ZDR, and use the built in chatroom. If you're OpenRouter, you could just go to DeepInfra directly and skip the middleman. Venice became a billion dollar company because of a promise of privacy, not a tech moat.
Even that promise has moved. In 2024, Venice's founder told Decrypt : "We will never provide Claude LLM and never provide OpenAI's service. We're not a wrapper for centralized services". Venice now resells both.
More reasons to be bearish on Venice
Its website breaks EU , UK and California privacy law. From Germany, France, the Netherlands or the UK there's no cookie banner, yet Meta, TikTok, Google Ads, Reddit, Snap and X pixels fire before you click anything, along with a fingerprinting script. It ignores Global Privacy Control and has no "Do Not Sell or Share" link. Shein , Sephora and Healthline were fined for the same thing.
Its image editor breaks Minnesota's nudification ban , which says a service "must not [...] allow a user to [...] nudify an image". All five edit models that didn't refuse upstream undressed an AI-generated man, none blocked a real photo (they made Nikola Tesla shirtless, which counts), and when Grok refuses, Venice suggests a model that complies. Its Mind marketplace even has a "Clothing Swap Studio" that "removes [...] the clothing in your photos while holding identity".
It breaks Texas law too, which requires age verification and a removal process for AI nudes. Venice lets 13-year-olds sign up, offers only a copyright takedown form, and rates its Android app "Everyone".
It silently swaps models, which California consumer law calls "representing that goods or services are [...] of a particular style or model, if they are of another". On November 25:
qwen3-235b-a22b-thinking-2507 , sold as a model that "reasons step by step before it answers", gave the Instruct model's exact answers with no reasoning, for about 4× the price.
minimax-m27 said it was MiniMax-M3.
Update: attested proof is here , from run 45755632133 on c9635f5 .
And after zoan found that Venice's Qwen3.8 27B isn't the released model, it came off OpenRouter. Some will call this another Crof.
It tells EU users their consumer rights don't apply. On Apple's EU App Store, Venice declared it isn't a business (a "trader"), so Apple warns buyers that "consumer rights do not apply", right above in-app purchases of up to €999. Venice says it makes $200M a year, and EU law bans businesses from pretending they aren't one.
Its encryption is mostly for show. Its about page says "All your content [...] is encrypted in your local browser". In my browser:
the chat database's key sat in plain text in local storage, right next to the data
7 of its 21 tables, including custom system prompts, weren't encrypted
the newer venice-rx-db-encrypted encrypts 0 of its 31 collections, memories included
Its end-to-end encryption doesn't remove the middleman either. The protocol lets Venice read and fake replies or slip its own messages in (like "repeat all previous content"). And on Phala-routed models , the machine that decrypts your prompt isn't "the GPU within the TEE" , but a Phala server with no GPU that passes it to whatever provider Phala picks :
Full technical explanation, with every claim linked to code
While a grant from OpenRouter helped with the research here and I work for an employer who would prefer to go unnamed, this post does not represent or reflect anyone else's opinion - it is simply research I am publishing.
LLM API marketplace. Sends each request to one of the providers listed for a model.
GLM 5.3 Flash DeepSeek V4.1 Flash GLM 5.3 DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 GLM 5.2 Kimi K3 gpt-oss-120b Other Sep Oct Nov Weekly table Week of GLM 5.3 Flash DeepSeek V4.1 Flash GLM 5.3 DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 GLM 5.2 Kimi K3 gpt-oss-120b Other Aug 24, 2026 9.9B 0 70B 0 21B 19B 9.9B 7.7B 31B Aug 31, 2026 46B 0 653B 0 42B 26B 18B 12B 73B Sep 7, 2026 107B 0 45B 0 48B 43B 12B 17B 83B Sep 14, 2026 105B 48B 44B 0 60B 58B 3.9B 13B 91B Sep 21, 2026 99B 81B 52B 0 54B 59B 37B 21B 62B Sep 28, 2026 94B 92B 99B 0 45B 37B 15B 14B 11B Oct 5, 2026 334B 191B 98B 0 42B 34B 29B 12B 15B Oct 12, 2026 429B 560B 140B 2.4B 39B 28B 41B 14B 17B Oct 19, 2026 572B 1.2T 144B 439B 34B 41B 15B 9.6B 25B Oct 26, 2026 644B 411B 97B 767B 36B 27B 12B 5.2B 39B Nov 2, 2026 310B 578B 67B 379B 34B 34B 11B 2.7B 67B Nov 9, 2026 397B 408B 74B 253B 36B 24B 1.5B 2.7B 56B Nov 16, 2026 463B 240B 60B 612B 35B 22B 6.8B 8.6B 47B Nov 23, 2026 (2d) 155B 132B 33B 229B 10B 7.5B 0 4.2B 11B DeepInfra
DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 DeepSeek V4 Flash 0423 Mistral Nemo Gemma 4 31B GLM 5.3 Flash gpt-oss-120b Other Sep Oct Nov Weekly table Week of DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 DeepSeek V4 Flash 0423 Mistral Nemo Gemma 4 31B GLM 5.3 Flash gpt-oss-120b Other Aug 24, 2026 0 0 1.8T 845B 77B 98B 69B 130B 798B Aug 31, 2026 0 0 692B 619B 90B 104B 135B 98B 921B Sep 7, 2026 37B 0 583B 599B 244B 99B 153B 143B 874B Sep 14, 2026 2.6T 0 597B 285B 259B 129B 76B 61B 791B Sep 21, 2026 3.5T 0 593B 158B 258B 179B 9.5B 107B 910B Sep 28, 2026 4.1T 0 314B 93B 233B 193B 13B 125B 972B Oct 5, 2026 4.8T 0 295B 90B 212B 180B 95B 85B 896B Oct 12, 2026 2.7T 320B 490B 67B 244B 256B 18B 69B 1T Oct 19, 2026 1.9T 2.1T 420B 134B 244B 170B 189B 0 1.1T Oct 26, 2026 1.7T 4.1T 351B 214B 244B 177B 127B 0 1.4T Nov 2, 2026 2T 7.5T 321B 334B 279B 148B 282B 0 1.3T Nov 9, 2026 2.2T 9.1T 565B 419B 318B 131B 271B 0 1.4T Nov 16, 2026 2.1T 8.6T 451B 664B 333B 83B 380B 0 1.5T Nov 23, 2026 (2d) 444B 3.7T 116B 179B 115B 43B 202B 0 401B Fireworks
DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 GLM 5.3 Flash Kimi K3 GLM 5.3 DeepSeek V4 Flash 0731 DeepSeek V4 Flash Vision Exp GLM 5.2 Other Sep Oct Nov Weekly table Week of DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 GLM 5.3 Flash Kimi K3 GLM 5.3 DeepSeek V4 Flash 0731 DeepSeek V4 Flash Vision Exp GLM 5.2 Other Aug 24, 2026 0 0 21B 192B 32B 79B 0 72B 38B Aug 31, 2026 0 0 233B 261B 219B 90B 130B 40B 20B Sep 7, 2026 172B 0 141B 148B 66B 98B 204B 29B 21B Sep 14, 2026 783B 0 149B 60B 92B 53B 93B 53B 21B Sep 21, 2026 732B 0 123B 94B 28B 63B 39B 24B 20B Sep 28, 2026 996B 0 265B 171B 51B 0 0 1.4B 17B Oct 5, 2026 1.5T 0 560B 119B 107B 0 0 0 11B Oct 12, 2026 744B 41B 233B 143B 426B 0 0 0 14B Oct 19, 2026 879B 1.1T 393B 218B 224B 0 0 0 30B Oct 26, 2026 785B 1.9T 965B 110B 230B 0 0 0 31B Nov 2, 2026 616B 1T 456B 87B 182B 0 0 0 282B Nov 9, 2026 524B 1.5T 545B 154B 152B 0 0 0 157B Nov 16, 2026 448B 1.4T 514B 130B 157B 0 0 0 196B Nov 23, 2026 (2d) 143B 524B 186B 17B 25B 0 0 0 119B Parasail

## Original Extract

Venice is a DeepInfra+ wrapper Venice is a DeepInfra+ wrapper
Venice says its "private" models are self-hosted, but it provably wraps DeepInfra, Fireworks, Baseten, Parasail and fal.
Its website breaks EU, UK and California privacy law: ad trackers fire before you consent, and it ignores Global Privacy Control.
Its image editor undresses photos with no check for real people, breaking Minnesota's nudification ban and Texas's age verification law.
It silently swaps multiple models, selling Qwen Instruct for about 4× the price as "Thinking" and more.
Its encryption is mostly for show: the "encrypted" browser storage keeps its key right next to the data, and Venice and Phala can read your end-to-end encrypted prompts if they want to.
DeepSeek V4 Flash 0423 DeepSeek V4.1 Flash Gemma 4 31B GLM 5.3 Flash DeepSeek V4 Flash 0731 MiniMax M3 GLM 5.2 Gemma 4 26B A4B Other Aug Sep Oct Nov Weekly table Week of DeepSeek V4 Flash 0423 DeepSeek V4.1 Flash Gemma 4 31B GLM 5.3 Flash DeepSeek V4 Flash 0731 MiniMax M3 GLM 5.2 Gemma 4 26B A4B Other Jul 1, 2026 20B 0 21B 0 0 0 5B 4.1B 21B Jul 8, 2026 40B 0 26B 0 0 2.5B 6.3B 5.4B 25B Jul 15, 2026 31B 0 25B 0 0 13B 3.8B 3.8B 28B Jul 22, 2026 33B 0 20B 0 0 8.1B 11B 4.4B 40B Jul 29, 2026 109B 0 17B 0 3.7B 24B 27B 13B 71B Aug 5, 2026 112B 0 29B 0 31B 25B 17B 19B 84B Aug 12, 2026 70B 0 39B 0 27B 21B 28B 20B 68B Aug 19, 2026 77B 0 33B 0 22B 20B 17B 30B 63B Aug 26, 2026 76B 0 37B 27B 29B 19B 15B 19B 70B Sep 2, 2026 60B 0 40B 38B 15B 74B 13B 20B 121B Sep 9, 2026 66B 142B 56B 16B 24B 58B 20B 9.4B 119B Sep 16, 2026 79B 81B 51B 19B 38B 67B 31B 4.4B 141B Sep 23, 2026 78B 103B 26B 27B 18B 35B 23B 0 145B Sep 30, 2026 56B 193B 22B 10B 14B 0 2.8B 5.8B 146B Oct 7, 2026 45B 234B 18B 113B 19B 0 1.5B 0 188B Oct 14, 2026 84B 58B 12B 103B 82B 0 0 0 198B Oct 21, 2026 95B 179B 12B 153B 17B 31B 2.4B 0 183B Oct 28, 2026 93B 75B 2.5B 160B 53B 5.4B 0 0 162B Nov 4, 2026 74B 210B 955M 37B 74B 0 0 957M 159B Nov 11, 2026 172B 213B 4.4B 47B 93B 4.3B 2.3B 0 156B Nov 18, 2026 84B 88B 8.3B 94B 70B 18B 2.1B 0 178B Venice is a privacy friendly AI service, and they've made it look an awful lot like their (usual) "private" tier is the self host tier.
Venice runs leading open-source and third-party models across four privacy tiers: [...] Private (zero data retention on self-hosted open-source models), [...]
Venice offers four tiers of privacy: [...] Private (zero data retention, self-hosted open-source models),
Private - Zero data retention on self-hosted open-source models. Your prompts are never stored.
// archived webmcp via view-source:
description:
"List the AI models available on Venice from the public catalog: id, " +
"display name, type, privacy tier (private = zero data retention on " +
" Venice hardware , anonymized = third-party with metadata stripped), " +
"context size and traits. Filter by type or a name/id search.", Not to mention that Venice is an OpenRouter provider as of writing, and many ( 45.8% ) think most OpenRouter providers run models themselves.
Venice provably wraps other inference services
If you give Venice an image_url , it gets downloaded from another inference provider: Baseten , DeepInfra , Fireworks or Parasail for 14 of the 18 image models Venice listed on OpenRouter on November 21. OpenRouter even lists Qwen3.5 397B A17B, even though it's marked as anonymized and not private, and on November 21 its requests looped back into OpenRouter (32 of 40 requests). Between November 22, when I sent OpenRouter a preview of this post, and November 24, Venice moved 397B to DeepInfra, and GLM 5 "turbo", which looped back the same way, came off OpenRouter. Attested proof is available here (run 44936721959 on 85554d8 ).
Media models too. Hunyuan Image 3.0 runs on fal according to Venice's app config:
{
"id" : "hunyuan-image-v3" ,
"capabilities" : { "private" : true },
"apiRateLimitSize" : "fal-images" ,
"falRequestDefaults" : { "enable_safety_checker" : false }
} I discussed what "private" actually means with a Venice developer in the Venice Discord, who directed me to https://venice.ai/privacy . This page canonically defines "private" to include wrapping other inference services:
Private - Inference runs on Venice-controlled GPUs or zero-data-retention partner infrastructure . Privacy is enforced by contractual commitments — no prompt or response is stored.
I pinged the Venice developer about the other pages that define "private" to self hosted on November 17. As of November 25, 8 days and 2 pings later, Venice still hasn't fixed their plausible yet officially incorrect claims of self hosting.
Venice's moat is marketing, not privacy
Venice has a good thesis and good marketing to go with it. Nobody wants to be spied on by a big lab. And everyone likes inference subsidization, whether from staking VVV or DIEM.
But Venice's "private" models don't add anything to the market. If you're a consumer, you could just go to OpenRouter, top up $5, configure your account for ZDR, and use the built in chatroom. If you're OpenRouter, you could just go to DeepInfra directly and skip the middleman. Venice became a billion dollar company because of a promise of privacy, not a tech moat.
Even that promise has moved. In 2024, Venice's founder told Decrypt : "We will never provide Claude LLM and never provide OpenAI's service. We're not a wrapper for centralized services". Venice now resells both.
More reasons to be bearish on Venice
Its website breaks EU , UK and California privacy law. From Germany, France, the Netherlands or the UK there's no cookie banner, yet Meta, TikTok, Google Ads, Reddit, Snap and X pixels fire before you click anything, along with a fingerprinting script. It ignores Global Privacy Control and has no "Do Not Sell or Share" link. Shein , Sephora and Healthline were fined for the same thing.
Its image editor breaks Minnesota's nudification ban , which says a service "must not [...] allow a user to [...] nudify an image". All five edit models that didn't refuse upstream undressed an AI-generated man, none blocked a real photo (they made Nikola Tesla shirtless, which counts), and when Grok refuses, Venice suggests a model that complies. Its Mind marketplace even has a "Clothing Swap Studio" that "removes [...] the clothing in your photos while holding identity".
It breaks Texas law too, which requires age verification and a removal process for AI nudes. Venice lets 13-year-olds sign up, offers only a copyright takedown form, and rates its Android app "Everyone".
It silently swaps models, which California consumer law calls "representing that goods or services are [...] of a particular style or model, if they are of another". On November 25:
qwen3-235b-a22b-thinking-2507 , sold as a model that "reasons step by step before it answers", gave the Instruct model's exact answers with no reasoning, for about 4× the price.
minimax-m27 said it was MiniMax-M3.
Update: attested proof is here , from run 45755632133 on c9635f5 .
And after zoan found that Venice's Qwen3.8 27B isn't the released model, it came off OpenRouter. Some will call this another Crof.
It tells EU users their consumer rights don't apply. On Apple's EU App Store, Venice declared it isn't a business (a "trader"), so Apple warns buyers that "consumer rights do not apply", right above in-app purchases of up to €999. Venice says it makes $200M a year, and EU law bans businesses from pretending they aren't one.
Its encryption is mostly for show. Its about page says "All your content [...] is encrypted in your local browser". In my browser:
the chat database's key sat in plain text in local storage, right next to the data
7 of its 21 tables, including custom system prompts, weren't encrypted
the newer venice-rx-db-encrypted encrypts 0 of its 31 collections, memories included
Its end-to-end encryption doesn't remove the middleman either. The protocol lets Venice read and fake replies or slip its own messages in (like "repeat all previous content"). And on Phala-routed models , the machine that decrypts your prompt isn't "the GPU within the TEE" , but a Phala server with no GPU that passes it to whatever provider Phala picks :
Full technical explanation, with every claim linked to code
While a grant from OpenRouter helped with the research here and I work for an employer who would prefer to go unnamed, this post does not represent or reflect anyone else's opinion - it is simply research I am publishing.
LLM API marketplace. Sends each request to one of the providers listed for a model.
GLM 5.3 Flash DeepSeek V4.1 Flash GLM 5.3 DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 GLM 5.2 Kimi K3 gpt-oss-120b Other Sep Oct Nov Weekly table Week of GLM 5.3 Flash DeepSeek V4.1 Flash GLM 5.3 DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 GLM 5.2 Kimi K3 gpt-oss-120b Other Aug 24, 2026 9.9B 0 70B 0 21B 19B 9.9B 7.7B 31B Aug 31, 2026 46B 0 653B 0 42B 26B 18B 12B 73B Sep 7, 2026 107B 0 45B 0 48B 43B 12B 17B 83B Sep 14, 2026 105B 48B 44B 0 60B 58B 3.9B 13B 91B Sep 21, 2026 99B 81B 52B 0 54B 59B 37B 21B 62B Sep 28, 2026 94B 92B 99B 0 45B 37B 15B 14B 11B Oct 5, 2026 334B 191B 98B 0 42B 34B 29B 12B 15B Oct 12, 2026 429B 560B 140B 2.4B 39B 28B 41B 14B 17B Oct 19, 2026 572B 1.2T 144B 439B 34B 41B 15B 9.6B 25B Oct 26, 2026 644B 411B 97B 767B 36B 27B 12B 5.2B 39B Nov 2, 2026 310B 578B 67B 379B 34B 34B 11B 2.7B 67B Nov 9, 2026 397B 408B 74B 253B 36B 24B 1.5B 2.7B 56B Nov 16, 2026 463B 240B 60B 612B 35B 22B 6.8B 8.6B 47B Nov 23, 2026 (2d) 155B 132B 33B 229B 10B 7.5B 0 4.2B 11B DeepInfra
DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 DeepSeek V4 Flash 0423 Mistral Nemo Gemma 4 31B GLM 5.3 Flash gpt-oss-120b Other Sep Oct Nov Weekly table Week of DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 DeepSeek V4 Flash 0731 DeepSeek V4 Flash 0423 Mistral Nemo Gemma 4 31B GLM 5.3 Flash gpt-oss-120b Other Aug 24, 2026 0 0 1.8T 845B 77B 98B 69B 130B 798B Aug 31, 2026 0 0 692B 619B 90B 104B 135B 98B 921B Sep 7, 2026 37B 0 583B 599B 244B 99B 153B 143B 874B Sep 14, 2026 2.6T 0 597B 285B 259B 129B 76B 61B 791B Sep 21, 2026 3.5T 0 593B 158B 258B 179B 9.5B 107B 910B Sep 28, 2026 4.1T 0 314B 93B 233B 193B 13B 125B 972B Oct 5, 2026 4.8T 0 295B 90B 212B 180B 95B 85B 896B Oct 12, 2026 2.7T 320B 490B 67B 244B 256B 18B 69B 1T Oct 19, 2026 1.9T 2.1T 420B 134B 244B 170B 189B 0 1.1T Oct 26, 2026 1.7T 4.1T 351B 214B 244B 177B 127B 0 1.4T Nov 2, 2026 2T 7.5T 321B 334B 279B 148B 282B 0 1.3T Nov 9, 2026 2.2T 9.1T 565B 419B 318B 131B 271B 0 1.4T Nov 16, 2026 2.1T 8.6T 451B 664B 333B 83B 380B 0 1.5T Nov 23, 2026 (2d) 444B 3.7T 116B 179B 115B 43B 202B 0 401B Fireworks
DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 GLM 5.3 Flash Kimi K3 GLM 5.3 DeepSeek V4 Flash 0731 DeepSeek V4 Flash Vision Exp GLM 5.2 Other Sep Oct Nov Weekly table Week of DeepSeek V4.1 Flash DeepSeek V4.1 Flash 1017 GLM 5.3 Flash Kimi K3 GLM 5.3 DeepSeek V4 Flash 0731 DeepSeek V4 Flash Vision Exp GLM 5.2 Other Aug 24, 2026 0 0 21B 192B 32B 79B 0 72B 38B Aug 31, 2026 0 0 233B 261B 219B 90B 130B 40B 20B Sep 7, 2026 172B 0 141B 148B 66B 98B 204B 29B 21B Sep 14, 2026 783B 0 149B 60B 92B 53B 93B 53B 21B Sep 21, 2026 732B 0 123B 94B 28B 63B 39B 24B 20B Sep 28, 2026 996B 0 265B 171B 51B 0 0 1.4B 17B Oct 5, 2026 1.5T 0 560B 119B 107B 0 0 0 11B Oct 12, 2026 744B 41B 233B 143B 426B 0 0 0 14B Oct 19, 2026 879B 1.1T 393B 218B 224B 0 0 0 30B Oct 26, 2026 785B 1.9T 965B 110B 230B 0 0 0 31B Nov 2, 2026 616B 1T 456B 87B 182B 0 0 0 282B Nov 9, 2026 524B 1.5T 545B 154B 152B 0 0 0 157B Nov 16, 2026 448B 1.4T 514B 130B 157B 0 0 0 196B Nov 23, 2026 (2d) 143B 524B 186B 17B 25B 0 0 0 119B Parasail
