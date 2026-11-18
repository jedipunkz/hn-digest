---
source: "https://www.lesswrong.com/posts/yQmzPMMeBLmGuZokr/i-m-worried-about-llms-accelerating-hardware-for-cheap-asi"
hn_url: "https://news.ycombinator.com/item?id=50547646"
title: "I'm worried about LLMs accelerating hardware for cheap ASI"
article_title: "I'm worried about LLMs accelerating hardware for cheap ASI — LessWrong"
image: "https://res.cloudinary.com/lesswrong-2-0/image/upload/v1654295382/new_mississippi_river_fjdmww.jpg"
author: "speckx"
captured_at: "2026-11-18T16:14:04Z"
capture_tool: "hn-digest"
hn_id: 50547646
score: 1
comments: 0
posted_at: "2026-11-18T15:40:11Z"
tags:
  - hacker-news
---

# I'm worried about LLMs accelerating hardware for cheap ASI

- HN: [50547646](https://news.ycombinator.com/item?id=50547646)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/yQmzPMMeBLmGuZokr/i-m-worried-about-llms-accelerating-hardware-for-cheap-asi)
- Score: 1
- Comments: 0
- Posted: 2026-11-18T15:40:11Z

## Translation

Title: I'm worried about LLMs accelerating hardware for cheap ASI
Article title: I'm worried about LLMs accelerating hardware for cheap ASI — LessWrong
Description: Basically: • * You could probably make a superintelligent AI with some currently unknown new paradigm for <$10k if you had the right hardware and al…

Article text:
Login I'm worried about LLMs accelerating hardware for cheap ASI — LessWrong AI Frontpage 15
I'm worried about LLMs accelerating hardware for cheap ASI
You could probably make a superintelligent AI with some currently unknown new paradigm for <$10k if you had the right hardware and algorithms
Plausibly this would be much easier and faster with better hardware
LLMs will make developers of novel-paradigm AI systems really good at hardware creation and design
This seems bad, because those novel-paradigm AIs seem unlikely to be aligned with human values
ASI could in principle be created cheaply
I think it's very plausible that with the right hardware and algorithms someone could make a wildly superintelligent AI model extremely cheaply (<$10k in energy cost), for roughly the reasons described in this post (although I think it underestimates how useful hardware improvements will be for building ASI). Basically, I think that given that you can train a human from nothing to a 20-year-old for less than $10k worth of electricity, you should expect to be able to train an AI to a similar level for the same amount of electricity with the right hardware and algorithms. In fact you should be able to do that much more efficiently, since you only care about intelligence and don't need to worry about powering a whole human body. And if you're making your own brain it's probably pretty doable to just make it bigger.
I don't think we'll get to this kind of novel ASI through iteratively improving current AI paradigms, but instead using something that's very different than current LLMs. To be clear, I also think that scaling LLMs could get us to superintelligence. I'm worried that ASI will be created through some novel paradigm first though.
It's hard to speculate about the alignment properties of an AI created through a training process that's unknown to me! However, it seems unlikely to me that if an ASI is created through novel, extremely cheap methods it would be aligned with human values. With today's LLM scaling, AI labs spend a lot of time trying to make their models aligned to human values, and iteratively work on aligning current models in the hope that techniques that they perfect on current models will scale. But if we're in a world where we get ASI through some cheap novel method, it seems less likely to be the result of iterative improvement and more likely to be the result of someone's first try at a new paradigm, which seems to me like it has much lower odds of happening to be aligned to human values on the first shot.
Better hardware might be very helpful for novel-paradigm ASI
If cheap novel-paradigm ASI is possible, why hasn't it happened yet? I think it's pretty plausible it's just that nobody has figured out the right mechanism yet, and we're just 1-2 algorithmic developments on the level of the transformer away from cheap novel-paradigm ASI. It seems pretty bad if this is the case, since in those worlds it would be nearly impossible to prevent the creation of a probably-misaligned superintelligence.
But another possibility is that many (but maybe not all?) of the ways of creating ASI through a novel paradigm are blocked on hardware improvements, and if we had the right hardware we'd get novel-paradigm ASI much faster by unlocking more routes to getting there. This seems pretty plausible to me - historically, deep learning was pretty bottlenecked by running on CPUs, and using GPUs to train deep neural networks might have contributed a lot to starting the deep learning revolution.
To be clear, I'm not sure if having better hardware is actually the blocker to novel-paradigm ASI! But it seems like it might be, and it seems important to think through what the consequences would be if that were the case.
It seems plausible to me that none of the current AI accelerator chips used for training models are well-suited to any paradigms for really cheap ASI. Human brains seem like they do computations in ways that are very different than AI accelerator chips; it seems to me like replicating the extreme cost-efficiency of humans might require creating similar hardware. AI research generally focuses on designing models that can be trained and run on existing hardware; the creation of new hardware for training AIs would probably allow for AI systems that work much differently than current ones.
Doing deep learning on GPUs was part of what started the deep learning revolution, by making it much faster to train models. It seems to me that we might end up in a situation where new hardware developments might significantly speed up AI development again.
LLMs will accelerate hardware development
LLMs are getting better at everything; it seems very plausible that they will rapidly speed up hardware development in the near future, which could be very useful for overcoming hardware bottlenecks to novel AI systems. LLMs are also getting better at FPGA programming [1] . It's less clear to me how useful current FPGAs would be for novel-paradigm AI, but it's plausible to me that they would be useful at least for developing hardware, and might be flexible enough to actually be able to run some kinds of novel-paradigm AI.
Isn't hardware manufacturing expensive?
You might object that manufacturing hardware is really expensive. Large-scale chip manufacturing on the latest semiconductor process is indeed very expensive, and the number of actors involved there is fairly small and manageable. But I don't think this is too much of a hurdle because:
It doesn't need to be on a very large scale: once you have one instance of your superintelligent AI you can use it to start an intelligence explosion
It doesn't need to be on the latest semiconductor process: this is helpful for building AI, but I doubt it would really be required
FPGAs could likely substitute for at least some part of the hardware needed
Here are some solutions I thought of; I'm not sure how good these ideas are:
AI labs could avoid making models too good at hardware design, and implement safeguards against using them for developing new hardware and FPGA programming [2]
Strengthening export controls on FPGAs (similar to those for Blackwell GPUs)
Put limits on chip fabrication and ban new kinds of AI accelerator chips from being produced. (Maybe using AI to review chips?)
Race to create an aligned superintelligence first (using more expensive conventional methods) and use it to prevent the creation of other superintelligences.
It seems to me like we could plausibly be in a world where preventing the creation of ASI is impossible without either extremely authoritarian rules or some kind of pivotal act, which is unfortunate. But there might be some low-hanging fruit when it comes to preventing ASI from being developed through novel methods!
I'm pretty uncertain about everything in this post - please let me know if you think I'm missing something here; I probably am!
^ I'm not sure if that's the best benchmark for FPGA programming (or at least the kind of FPGA-related capabilities I'm worried about)
^ Anthropic has some frontier LLM safeguards currently, although in my brief testing they don't seem to block AI-related FPGA programming

## Original Extract

Basically: • * You could probably make a superintelligent AI with some currently unknown new paradigm for <$10k if you had the right hardware and al…

Login I'm worried about LLMs accelerating hardware for cheap ASI — LessWrong AI Frontpage 15
I'm worried about LLMs accelerating hardware for cheap ASI
You could probably make a superintelligent AI with some currently unknown new paradigm for <$10k if you had the right hardware and algorithms
Plausibly this would be much easier and faster with better hardware
LLMs will make developers of novel-paradigm AI systems really good at hardware creation and design
This seems bad, because those novel-paradigm AIs seem unlikely to be aligned with human values
ASI could in principle be created cheaply
I think it's very plausible that with the right hardware and algorithms someone could make a wildly superintelligent AI model extremely cheaply (<$10k in energy cost), for roughly the reasons described in this post (although I think it underestimates how useful hardware improvements will be for building ASI). Basically, I think that given that you can train a human from nothing to a 20-year-old for less than $10k worth of electricity, you should expect to be able to train an AI to a similar level for the same amount of electricity with the right hardware and algorithms. In fact you should be able to do that much more efficiently, since you only care about intelligence and don't need to worry about powering a whole human body. And if you're making your own brain it's probably pretty doable to just make it bigger.
I don't think we'll get to this kind of novel ASI through iteratively improving current AI paradigms, but instead using something that's very different than current LLMs. To be clear, I also think that scaling LLMs could get us to superintelligence. I'm worried that ASI will be created through some novel paradigm first though.
It's hard to speculate about the alignment properties of an AI created through a training process that's unknown to me! However, it seems unlikely to me that if an ASI is created through novel, extremely cheap methods it would be aligned with human values. With today's LLM scaling, AI labs spend a lot of time trying to make their models aligned to human values, and iteratively work on aligning current models in the hope that techniques that they perfect on current models will scale. But if we're in a world where we get ASI through some cheap novel method, it seems less likely to be the result of iterative improvement and more likely to be the result of someone's first try at a new paradigm, which seems to me like it has much lower odds of happening to be aligned to human values on the first shot.
Better hardware might be very helpful for novel-paradigm ASI
If cheap novel-paradigm ASI is possible, why hasn't it happened yet? I think it's pretty plausible it's just that nobody has figured out the right mechanism yet, and we're just 1-2 algorithmic developments on the level of the transformer away from cheap novel-paradigm ASI. It seems pretty bad if this is the case, since in those worlds it would be nearly impossible to prevent the creation of a probably-misaligned superintelligence.
But another possibility is that many (but maybe not all?) of the ways of creating ASI through a novel paradigm are blocked on hardware improvements, and if we had the right hardware we'd get novel-paradigm ASI much faster by unlocking more routes to getting there. This seems pretty plausible to me - historically, deep learning was pretty bottlenecked by running on CPUs, and using GPUs to train deep neural networks might have contributed a lot to starting the deep learning revolution.
To be clear, I'm not sure if having better hardware is actually the blocker to novel-paradigm ASI! But it seems like it might be, and it seems important to think through what the consequences would be if that were the case.
It seems plausible to me that none of the current AI accelerator chips used for training models are well-suited to any paradigms for really cheap ASI. Human brains seem like they do computations in ways that are very different than AI accelerator chips; it seems to me like replicating the extreme cost-efficiency of humans might require creating similar hardware. AI research generally focuses on designing models that can be trained and run on existing hardware; the creation of new hardware for training AIs would probably allow for AI systems that work much differently than current ones.
Doing deep learning on GPUs was part of what started the deep learning revolution, by making it much faster to train models. It seems to me that we might end up in a situation where new hardware developments might significantly speed up AI development again.
LLMs will accelerate hardware development
LLMs are getting better at everything; it seems very plausible that they will rapidly speed up hardware development in the near future, which could be very useful for overcoming hardware bottlenecks to novel AI systems. LLMs are also getting better at FPGA programming [1] . It's less clear to me how useful current FPGAs would be for novel-paradigm AI, but it's plausible to me that they would be useful at least for developing hardware, and might be flexible enough to actually be able to run some kinds of novel-paradigm AI.
Isn't hardware manufacturing expensive?
You might object that manufacturing hardware is really expensive. Large-scale chip manufacturing on the latest semiconductor process is indeed very expensive, and the number of actors involved there is fairly small and manageable. But I don't think this is too much of a hurdle because:
It doesn't need to be on a very large scale: once you have one instance of your superintelligent AI you can use it to start an intelligence explosion
It doesn't need to be on the latest semiconductor process: this is helpful for building AI, but I doubt it would really be required
FPGAs could likely substitute for at least some part of the hardware needed
Here are some solutions I thought of; I'm not sure how good these ideas are:
AI labs could avoid making models too good at hardware design, and implement safeguards against using them for developing new hardware and FPGA programming [2]
Strengthening export controls on FPGAs (similar to those for Blackwell GPUs)
Put limits on chip fabrication and ban new kinds of AI accelerator chips from being produced. (Maybe using AI to review chips?)
Race to create an aligned superintelligence first (using more expensive conventional methods) and use it to prevent the creation of other superintelligences.
It seems to me like we could plausibly be in a world where preventing the creation of ASI is impossible without either extremely authoritarian rules or some kind of pivotal act, which is unfortunate. But there might be some low-hanging fruit when it comes to preventing ASI from being developed through novel methods!
I'm pretty uncertain about everything in this post - please let me know if you think I'm missing something here; I probably am!
^ I'm not sure if that's the best benchmark for FPGA programming (or at least the kind of FPGA-related capabilities I'm worried about)
^ Anthropic has some frontier LLM safeguards currently, although in my brief testing they don't seem to block AI-related FPGA programming
