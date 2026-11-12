---
source: "https://twitter.com/willdepue/status/2120992457630326841"
hn_url: "https://news.ycombinator.com/item?id=50474813"
title: "Top Chinese AI labs built their newest models on Nvidia's Blackwell chips"
article_title: "will depue on X: \"I have a major result to share.\nI’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.\nThis sounds insane, b… /\n[truncated]"
image: "https://pbs.twimg.com/media/HW9GgJ_WcAA9m3I?format=webp&name=large"
author: "leroy_masochist"
captured_at: "2026-11-12T22:04:20Z"
capture_tool: "hn-digest"
hn_id: 50474813
score: 6
comments: 2
posted_at: "2026-11-12T21:49:52Z"
tags:
  - hacker-news
---

# Top Chinese AI labs built their newest models on Nvidia's Blackwell chips

- HN: [50474813](https://news.ycombinator.com/item?id=50474813)
- Source: [twitter.com](https://twitter.com/willdepue/status/2120992457630326841)
- Score: 6
- Comments: 2
- Posted: 2026-11-12T21:49:52Z

## Translation

Title: Top Chinese AI labs built their newest models on Nvidia's Blackwell chips
Article title: will depue on X: "I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, b… /
[truncated]
Description: I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered th…

Article text:
will depue on X: "I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, b… / X Post
will depue on X: "I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered that almost every AI model carries a hidden signature of the GPU that trained it. When a model is first set up, the starting values of its weights are laid out in a pattern that depends on the exact GPU. Each combination of GPU and seed number produces a unique pattern—a sort of fingerprint of the hardware used to create the model. Traces of that pattern survive training into the published model.
Again, this evidence isn’t based on any insider information or leaking. It’s based on a forensic analysis of the model weights released by the companies themselves.
I searched more than six trillion combinations of chip layout and seed number and matched them to open-weights models. In doing so, I found conclusive evidence that six Chinese AI models—including DeepSeek-V4-Pro, Kimi K3, and MiniMax-M3—were initialized on NVIDIA Blackwell systems. Four of the models match Grace Blackwell, the GB200 and GB300 parts in NVIDIA’s flagship liquid-cooled rack-scale systems.
I then verified these results on real hardware. Running unmodified PyTorch with the seeds I’d recovered, a Blackwell B300 reproduced DeepSeek-V4-Flash’s untouched starting weights bit for bit, and a rented GB200 reproduced Kimi K3’s to a correlation of 0.999995. A Hopper H100 reproduced none of DeepSeek-V4-Flash’s, and changing the seed by one broke the match. (DeepSeek’s seed is 953049583 by the way.) The method also was verified to match 24 known open-weights models to the hardware that developers said they used.
It is exceedingly unlikely that these results are a coincidence. Across six trillion combinations, the best accidental match should reach only about 8 sigma (i.e., eight standard deviations abov
[truncated]
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
will depue @willdepue I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered that almost every AI model carries a hidden signature of the GPU that trained it. When a model is first set up, the starting values of its weights are laid out in a pattern that depends on the exact GPU. Each combination of GPU and seed number produces a unique pattern—a sort of fingerprint of the hardware used to create the model. Traces of that pattern survive training into the published model.
Again, this evidence isn’t based on any insider information or leaking. It’s based on a forensic analysis of the model weights released by the companies themselves.
I searched more than six trillion combinations of chip layout and seed number and matched them to open-weights models. In doing so, I found conclusive evidence that six Chinese AI models—including DeepSeek-V4-Pro, Kimi K3, and MiniMax-M3—were initialized on NVIDIA Blackwell systems. Four of the models match Grace Blackwell, the GB200 and GB300 parts in NVIDIA’s flagship liquid-cooled rack-scale systems.
I then verified these results on real hardware. Running unmodified PyTorch with the seeds I’d recovered, a Blackwell B300 reproduced DeepSeek-V4-Flash’s untouched starting weights bit for bit, and a rented GB200 reproduced Kimi K3’s to a correlation of 0.999995. A Hopper H100 reproduced none of DeepSeek-V4-Flash’s, and changing the seed by one broke the match. (DeepSeek’s seed is 953049583 by the way.) The method also was verified to match 24 known open-weights models to the hardware that developers said they used.
It is exceedingly unlikely that these results are a coincidence. Across six trillion combinations, the best accidental match should reach only about 8 sigma (i.e., eight standard deviations
[truncated]
will depue @willdepue 34m If you’re unfamiliar with how machine learning works: an AI model is truly just a very long list (billions, sometimes trillions) of numbers called weights, which, loosely, when multiplied and added together produce readable text.
Training adjusts those numbers over and over Show more 2 1 155 11K will depue @willdepue 34m A second test with a photo of me. I took a model’s ordinary random starting weights and arranged them once so that they show the photo. Then I trained the model and kept the arrangement fixed. The photo fades but the sky and the treeline can still be made out at the end. GIF 1 102 8K will depue @willdepue 34m The random numbers a model starts with aren’t truly random.
They come from a formula that starts from one starting number called the seed, and using the same seed gives the same numbers every time. Engineers set it up this way so they can repeat an experiment.
To fill billions Show more GIF 1 1 91 6.6K will depue @willdepue 34m That means I can take a published model, pick a chip and a seed, generate what the starting weights would have been and compare them to the real ones. With the right chip they line up and with the wrong chip they don’t.
This is that search on DeepSeek-V4-Flash. Every processor count gives noise until 148 (Blackwell), where the weights match. GIF 2 2 74 5.8K will depue @willdepue 34m Critically, MiniMax and Ant have not been publicly tied to Blackwell before. MiniMax’s models match Hopper through M2.7 in April and Blackwell from M3 in June. Ant’s own GitHub has benchmark logs from a GB300 dated March 31. 1 2 67 5.4K will depue @willdepue 34m Of course, this tells us nothing about where these chips physically are. They could have been smuggled into China or the labs are using them abroad through gaps that Congress has not closed. In either case, it’s clear that the export bans are deeply failing. 1 1 83 5.1K will depue @willdepue 34m Unfortunately, it’s of note that now that this is public any lab
[truncated]
Molly Cantillon @mollycantillon 30m Validates what I think is an open secret! Nice job Will 2 42 3.8K

## Original Extract

I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered th…

will depue on X: "I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, b… / X Post
will depue on X: "I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered that almost every AI model carries a hidden signature of the GPU that trained it. When a model is first set up, the starting values of its weights are laid out in a pattern that depends on the exact GPU. Each combination of GPU and seed number produces a unique pattern—a sort of fingerprint of the hardware used to create the model. Traces of that pattern survive training into the published model.
Again, this evidence isn’t based on any insider information or leaking. It’s based on a forensic analysis of the model weights released by the companies themselves.
I searched more than six trillion combinations of chip layout and seed number and matched them to open-weights models. In doing so, I found conclusive evidence that six Chinese AI models—including DeepSeek-V4-Pro, Kimi K3, and MiniMax-M3—were initialized on NVIDIA Blackwell systems. Four of the models match Grace Blackwell, the GB200 and GB300 parts in NVIDIA’s flagship liquid-cooled rack-scale systems.
I then verified these results on real hardware. Running unmodified PyTorch with the seeds I’d recovered, a Blackwell B300 reproduced DeepSeek-V4-Flash’s untouched starting weights bit for bit, and a rented GB200 reproduced Kimi K3’s to a correlation of 0.999995. A Hopper H100 reproduced none of DeepSeek-V4-Flash’s, and changing the seed by one broke the match. (DeepSeek’s seed is 953049583 by the way.) The method also was verified to match 24 known open-weights models to the hardware that developers said they used.
It is exceedingly unlikely that these results are a coincidence. Across six trillion combinations, the best accidental match should reach only about 8 sigma (i.e., eight standard deviations abov
[truncated]
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
will depue @willdepue I have a major result to share.
I’ve found conclusive forensic evidence that the top 4 Chinese AI labs—DeepSeek, Moonshot, MiniMax and Ant Group—have built their newest models on NVIDIA’s most advanced Blackwell chips, which are barred from export to China.
This sounds insane, but I discovered that almost every AI model carries a hidden signature of the GPU that trained it. When a model is first set up, the starting values of its weights are laid out in a pattern that depends on the exact GPU. Each combination of GPU and seed number produces a unique pattern—a sort of fingerprint of the hardware used to create the model. Traces of that pattern survive training into the published model.
Again, this evidence isn’t based on any insider information or leaking. It’s based on a forensic analysis of the model weights released by the companies themselves.
I searched more than six trillion combinations of chip layout and seed number and matched them to open-weights models. In doing so, I found conclusive evidence that six Chinese AI models—including DeepSeek-V4-Pro, Kimi K3, and MiniMax-M3—were initialized on NVIDIA Blackwell systems. Four of the models match Grace Blackwell, the GB200 and GB300 parts in NVIDIA’s flagship liquid-cooled rack-scale systems.
I then verified these results on real hardware. Running unmodified PyTorch with the seeds I’d recovered, a Blackwell B300 reproduced DeepSeek-V4-Flash’s untouched starting weights bit for bit, and a rented GB200 reproduced Kimi K3’s to a correlation of 0.999995. A Hopper H100 reproduced none of DeepSeek-V4-Flash’s, and changing the seed by one broke the match. (DeepSeek’s seed is 953049583 by the way.) The method also was verified to match 24 known open-weights models to the hardware that developers said they used.
It is exceedingly unlikely that these results are a coincidence. Across six trillion combinations, the best accidental match should reach only about 8 sigma (i.e., eight standard deviations
[truncated]
will depue @willdepue 34m If you’re unfamiliar with how machine learning works: an AI model is truly just a very long list (billions, sometimes trillions) of numbers called weights, which, loosely, when multiplied and added together produce readable text.
Training adjusts those numbers over and over Show more 2 1 155 11K will depue @willdepue 34m A second test with a photo of me. I took a model’s ordinary random starting weights and arranged them once so that they show the photo. Then I trained the model and kept the arrangement fixed. The photo fades but the sky and the treeline can still be made out at the end. GIF 1 102 8K will depue @willdepue 34m The random numbers a model starts with aren’t truly random.
They come from a formula that starts from one starting number called the seed, and using the same seed gives the same numbers every time. Engineers set it up this way so they can repeat an experiment.
To fill billions Show more GIF 1 1 91 6.6K will depue @willdepue 34m That means I can take a published model, pick a chip and a seed, generate what the starting weights would have been and compare them to the real ones. With the right chip they line up and with the wrong chip they don’t.
This is that search on DeepSeek-V4-Flash. Every processor count gives noise until 148 (Blackwell), where the weights match. GIF 2 2 74 5.8K will depue @willdepue 34m Critically, MiniMax and Ant have not been publicly tied to Blackwell before. MiniMax’s models match Hopper through M2.7 in April and Blackwell from M3 in June. Ant’s own GitHub has benchmark logs from a GB300 dated March 31. 1 2 67 5.4K will depue @willdepue 34m Of course, this tells us nothing about where these chips physically are. They could have been smuggled into China or the labs are using them abroad through gaps that Congress has not closed. In either case, it’s clear that the export bans are deeply failing. 1 1 83 5.1K will depue @willdepue 34m Unfortunately, it’s of note that now that this is public any lab
[truncated]
Molly Cantillon @mollycantillon 30m Validates what I think is an open secret! Nice job Will 2 42 3.8K
