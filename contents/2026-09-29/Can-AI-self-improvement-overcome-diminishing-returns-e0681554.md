---
source: "https://www.rameznaam.com/p/ai-rsi-isnt-leading-to-super-intelligence"
hn_url: "https://news.ycombinator.com/item?id=49888101"
title: "Can AI self-improvement overcome diminishing returns?"
article_title: "Can AI self-improvement overcome diminishing returns?"
image: "https://substackcdn.com/image/fetch/$s_!cT_g!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F59d5d33b-af61-4e3c-946e-e14a227f0b04_4000x2374.jpeg"
author: "thevises"
captured_at: "2026-09-29T04:34:36Z"
capture_tool: "hn-digest"
hn_id: 49888101
score: 1
comments: 0
posted_at: "2026-09-29T04:06:31Z"
tags:
  - hacker-news
---

# Can AI self-improvement overcome diminishing returns?

- HN: [49888101](https://news.ycombinator.com/item?id=49888101)
- Source: [www.rameznaam.com](https://www.rameznaam.com/p/ai-rsi-isnt-leading-to-super-intelligence)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T04:06:31Z

## Translation

Title: Can AI self-improvement overcome diminishing returns?
Description: AI Self-Improvement (RSI) doesn't guarantee a fast takeoff to Artificial Super Intelligence (ASI). Everything in AI faces steep diminishing returns.

Article text:
Can AI self-improvement overcome diminishing returns?
Subscribe Sign in Can AI self-improvement overcome diminishing returns?
AI self-improvement may not be powerful enough to overcome diminishing returns. Even RSI may not lead to a fast takeoff to ASI (superintelligence).
Ramez Naam Sep 27, 2026 44 2 6 Share This is a cross-post of my guest post on Noahpinion , which Noah Smith kindly invited me to write. Thanks to Noah for the many conversations, and for encouraging and needling me to actually write this.
1. AI is Helping Improve Itself
AI is already helping improve itself. The question is whether even fully autonomous recursive self-improvement (RSI) would cause a runaway intelligence explosion.
This Substack is reader-supported. To receive new posts and support my work, consider becoming a free or paid subscriber.
The theory is that each generation of AI could build a better successor, faster than the last generation did. That could lead to a “fast takeoff,” with capabilities surging to artificial superintelligence (ASI) in a year, months, or even days.
Here’s my take: Given our best current data, the AI self-improvement loop would need to be roughly 5–10× stronger to sustain itself, let alone run away. I’ll explain this math in section 8 . I expect incredibly rapid AI progress by the standards of nearly any other technology. But the evidence we have doesn’t suggest a sudden explosion to incomprehensible superintelligence anytime soon.
I could be wrong. Forecasters have repeatedly underestimated AI progress ! I could well be next. One thing that’s clear is that we need better data . For now, let’s work with what we can measure, and stay open to breakthroughs that could change the picture.
How Strong Is the Feedback Loop?
Figure 1. How strong is the self-improvement loop? Model .
Here’s the case, with links to each part:
Narrow Superintelligence Is Here Today
From ECI Scores to METR Task Horizons
The Real World Looks Nothing Like METR or AI 2027
Impressive AI Numbers → Sharp Diminishing Returns
We’re Not Seeing Signs of Acceleration
Keeping Up the Pace Takes Exponentially More Resources
Better AI May Be Needed Just to Maintain the Pace
Progress Gets Harder; Ideas Get Harder to Find
The Current Feedback Loop Doesn’t Look Strong Enough
OpenAI’s Data Shows How Weak the Loop Is
What Could Accelerate Progress?
We Need More Data to Track This Well
Key charts: The feedback loop · Measured vs. forecast progress · Diminishing returns
People use “recursive self-improvement” to mean everything from AI boosting the productivity of human researchers to AI bootstrapping itself to incomprehensible intelligence. Here’s my taxonomy: productivity gains (Type 1), increasing autonomy while still facing diminishing returns (Types 2–4), and a runaway loop to superintelligence if we can ever find accelerating returns (Type 5).
Figure 2. Five types of AI self-improvement.
We’ve made real progress on Types 1 and 2: AI helps both researchers and engineers inside of AI companies, and powerful models can train and improve smaller ones. We haven’t yet seen clear evidence for Type 3 (though Alibaba just made some strong claims ) and certainly not for Type 4. I do expect autonomous self-improvement to arrive at some point. I’m skeptical that it leads to Type 5 - runaway super-intelligence - without a major conceptual breakthrough.
There are plenty of other definitions of RSI, which can be a bit confusing. Weco’s four levels of RSI are close to mine. For a broader tour of all the things people mean when they say ‘RSI’, read Tom Cunningham’s comprehensive guide .
We Already Have Narrow Superintelligence
I do expect narrow superintelligence in highly verifiable domains . Think chess, Go, formal math, parts of computer science and coding. Highly verifiable domains are largely formal and structured types of work where machines can generate unlimited training data, with perfect or near-perfect verification of correct vs incorrect, and do so entirely in software without waiting on the physical world or humans. That’s an ideal setting for AI learning.
Figure 3. What makes a domain highly verifiable?
In fact, we already have narrow superintelligence in game playing. We’re seeing it happen now in the most formal parts of math, in particular in proofs and in finding counter-examples that disprove major conjectures. For example, OpenAI recently reported an AI-generated proof resolving the Navier–Stokes existence and smoothness problem . Parts of software development are also extremely verifiable, while others are a bit less crisp (such as understanding what humans want).
That isn’t the same as broad superintelligence. Even our most powerful models need far more training data than humans, struggle to learn reliably from ongoing experience, and fail in surprising ways on tasks people find straightforward. Superhuman math doesn’t automatically mean superhuman judgment everywhere else.
3. Real AI Research is Harder than Benchmarks or Forecasts
Benchmarks and forecasts suggest that AI models should reliably succeed at coding tasks that take humans hours, without human help. The real world is messier. OpenAI’s internal data shows much shorter stretches of autonomous work on research tasks.
In its Research Acceleration / RSI report , OpenAI showed how often its models completed tasks with and without human help, grouped by how long a human would need to do the work.
Figure 4. OpenAI’s internal research tasks. Source .
Even on tasks that would take a human less than 15 minutes, OpenAI’s models succeeded without human intervention only 86% of the time. The estimated task length at 80% success was roughly 15 minutes over the first seven months of the year. July’s results were similar to the whole period average.
Fully autonomous RSI would require an AI to string together a great many research tasks reliably, stretching out over complex tasks that humans need weeks or months to accomplish. OpenAI’s data suggests that we aren’t close.
Anthropic also released a graph showing how Claude accelerates AI research. It shows that internal AI models collaborate on or even lead more than 90% of R&D tasks. That’s objectively impressive. At the same time, the graph reports zero cases of AI autonomously completing AI R&D tasks.
Figure 5. Claude’s role in internal AI R&D. Source .
These are incredible tools. But they still need skilled people to set direction and get them back on track.
The Gap Between Benchmarks and Reality
For years, METR has been publishing a chart showing what length of coding task (measured in human hours to complete) best-in-class AI models can achieve. It’s been called the most important graph in AI . METR’s Mythos Preview evaluation estimated that the model could succeed at 80% of coding tasks that took humans three hours.
Figure 6. METR’s 80% task horizons. Source .
From ECI Scores to METR Task Horizons
Epoch’s own rule of thumb is that every five additional points of ECI (their overall benchmark of AI capability) correspond to roughly a doubling of METR’s task horizon. Using that formula, we’d expect GPT 5.6 Sol and GPT 6 Astra to be 80% successful at completing tasks of around 4 hours and 11 hours of human length, respectively.
Another estimate (a forecast) of AI task length comes from the AI 2027 scenario, which estimated that by July 2026, frontier AIs would be 80% successful accomplishing tasks of around 11 hours. Fairly similar.
The AI 2027 Tracker charts all of these.
Figure 7. The AI 2027 Tracker. Source .
Inside OpenAI, though, the July research-task horizon at 80% success was roughly 15 minutes.
Measured Progress vs. AI 2027 and ECI-extrapolated METR
Figure 8. Forecasts, benchmarks, and real AI research. Tracker · OpenAI .
A four-hour benchmark horizon is about 16 times longer than OpenAI’s research horizon. AI 2027’s 11-hour forecast is about 44 times longer. Of course, the tasks being performed by researchers at OpenAI aren’t the same as those in the METR benchmark. So we should expect some discrepancy. This, however, goes well beyond that.
Actual AI research at OpenAI is an order of magnitude or more harder than metrics, benchmarks, or forecasts suggest. That should make us wary of relying too much on benchmarks, or of saying that future scenarios like AI 2027 are ‘on track.’ The authors of the related AI 2040 project still describe AI 2027 as roughly the future they expect, and say reality is tracking closer to it than even they expected. That’s not what we see from within OpenAI. This isn’t an apples-to-apples comparison, but the difference is remarkable. AI 2027 appears to be substantially over-optimistic in this regard.
In January of this year, Nathan Witkin made a case that the METR graph was exaggerating progress . The real world data suggests that at least some of his critiques were correct. The gap between benchmarks, forecasts, and data gleaned from actual use of AI should influence our expectations about the future.
4. The Sharp Diminishing Returns to Impressive AI Numbers
OpenAI’s report also shows impressive increases in AI token usage, in compute spend per researcher, and in lines of code written. But these aren’t results. They’re intermediate measures. How much progress do they actually drive?
Researchers used 124x more tokens per person. Engineers shipped roughly 7x as many lines of code per person. Researchers ran 1.6x as many experiments per researcher vs OpenAI’s 2025 whole year average.
Figure 9. Token use inside OpenAI. Source .
Figure 10. Experiment pace inside OpenAI. Source .
From More Tokens to More Experiments
Figure 11. From tokens to code to experiments. Source .
More tokens and code don’t tell us much on their own. The 1.6× experiment pace is closer to useful research output. Even that doesn’t mean AI is improving 1.6× faster.
An enormous increase in AI output has accompanied a much smaller increase in experiments run.
This isn’t a controlled experiment. We don’t know what would happen if researchers switched back to an older model. But it gives us a useful view of AI-assisted research inside a frontier lab.
It’s not just OpenAI. Anthropic reports that their engineers are now producing 8x as many lines of code per person as they did in 2024 - somewhat similar to OpenAI. Anthropic also sees significant diminishing returns between productivity and AI progress. Here’s a direct quote from its Mythos Preview system card :
“ Productivity uplift does not translate one-for-one to capabilities progress . We surveyed technical staff on the productivity uplift they experience from Claude Mythos Preview relative to zero AI assistance. The distribution is wide and the geometric mean is on the order of 4× . […] We estimate that reaching 2× on overall progress via this channel would require uplift roughly an order of magnitude larger than what we observe . ”- Anthropic, Claude Mythos Preview System Card; emphasis mine
Translation: To double the pace of AI progress, Anthropic estimates that AI would need to increase the productivity of their employees by roughly a factor of 40 relative to no AI assistance.
Figure 12. Anthropic’s productivity-to-progress estimate. Source .
This is an estimate, not a measurement of progress. Even the 4× productivity figure comes from an opt-in survey of 130 Anthropic staff . I put more weight on OpenAI’s logged experiments, though the two sources measure different things.
We don’t yet know how much those extra experiments are accelerating AI improvement, if at all. In general, there are also steeply diminishing returns of more experiments in most branches of science. That means that a 60% increase in experiment pace could be on the order of a 10% boost to AI improvement pace. (A power law exponent of 0.2, for those who want to do the math.) That’s speculation for now. We’ll learn more as the labs publish results.
Test Time Compute Also Has Diminishing Returns
What about giving the same AI model more time to think?
That scales badly also. In OpenAI’s re

[truncated]

## Original Extract

AI Self-Improvement (RSI) doesn't guarantee a fast takeoff to Artificial Super Intelligence (ASI). Everything in AI faces steep diminishing returns.

Can AI self-improvement overcome diminishing returns?
Subscribe Sign in Can AI self-improvement overcome diminishing returns?
AI self-improvement may not be powerful enough to overcome diminishing returns. Even RSI may not lead to a fast takeoff to ASI (superintelligence).
Ramez Naam Sep 27, 2026 44 2 6 Share This is a cross-post of my guest post on Noahpinion , which Noah Smith kindly invited me to write. Thanks to Noah for the many conversations, and for encouraging and needling me to actually write this.
1. AI is Helping Improve Itself
AI is already helping improve itself. The question is whether even fully autonomous recursive self-improvement (RSI) would cause a runaway intelligence explosion.
This Substack is reader-supported. To receive new posts and support my work, consider becoming a free or paid subscriber.
The theory is that each generation of AI could build a better successor, faster than the last generation did. That could lead to a “fast takeoff,” with capabilities surging to artificial superintelligence (ASI) in a year, months, or even days.
Here’s my take: Given our best current data, the AI self-improvement loop would need to be roughly 5–10× stronger to sustain itself, let alone run away. I’ll explain this math in section 8 . I expect incredibly rapid AI progress by the standards of nearly any other technology. But the evidence we have doesn’t suggest a sudden explosion to incomprehensible superintelligence anytime soon.
I could be wrong. Forecasters have repeatedly underestimated AI progress ! I could well be next. One thing that’s clear is that we need better data . For now, let’s work with what we can measure, and stay open to breakthroughs that could change the picture.
How Strong Is the Feedback Loop?
Figure 1. How strong is the self-improvement loop? Model .
Here’s the case, with links to each part:
Narrow Superintelligence Is Here Today
From ECI Scores to METR Task Horizons
The Real World Looks Nothing Like METR or AI 2027
Impressive AI Numbers → Sharp Diminishing Returns
We’re Not Seeing Signs of Acceleration
Keeping Up the Pace Takes Exponentially More Resources
Better AI May Be Needed Just to Maintain the Pace
Progress Gets Harder; Ideas Get Harder to Find
The Current Feedback Loop Doesn’t Look Strong Enough
OpenAI’s Data Shows How Weak the Loop Is
What Could Accelerate Progress?
We Need More Data to Track This Well
Key charts: The feedback loop · Measured vs. forecast progress · Diminishing returns
People use “recursive self-improvement” to mean everything from AI boosting the productivity of human researchers to AI bootstrapping itself to incomprehensible intelligence. Here’s my taxonomy: productivity gains (Type 1), increasing autonomy while still facing diminishing returns (Types 2–4), and a runaway loop to superintelligence if we can ever find accelerating returns (Type 5).
Figure 2. Five types of AI self-improvement.
We’ve made real progress on Types 1 and 2: AI helps both researchers and engineers inside of AI companies, and powerful models can train and improve smaller ones. We haven’t yet seen clear evidence for Type 3 (though Alibaba just made some strong claims ) and certainly not for Type 4. I do expect autonomous self-improvement to arrive at some point. I’m skeptical that it leads to Type 5 - runaway super-intelligence - without a major conceptual breakthrough.
There are plenty of other definitions of RSI, which can be a bit confusing. Weco’s four levels of RSI are close to mine. For a broader tour of all the things people mean when they say ‘RSI’, read Tom Cunningham’s comprehensive guide .
We Already Have Narrow Superintelligence
I do expect narrow superintelligence in highly verifiable domains . Think chess, Go, formal math, parts of computer science and coding. Highly verifiable domains are largely formal and structured types of work where machines can generate unlimited training data, with perfect or near-perfect verification of correct vs incorrect, and do so entirely in software without waiting on the physical world or humans. That’s an ideal setting for AI learning.
Figure 3. What makes a domain highly verifiable?
In fact, we already have narrow superintelligence in game playing. We’re seeing it happen now in the most formal parts of math, in particular in proofs and in finding counter-examples that disprove major conjectures. For example, OpenAI recently reported an AI-generated proof resolving the Navier–Stokes existence and smoothness problem . Parts of software development are also extremely verifiable, while others are a bit less crisp (such as understanding what humans want).
That isn’t the same as broad superintelligence. Even our most powerful models need far more training data than humans, struggle to learn reliably from ongoing experience, and fail in surprising ways on tasks people find straightforward. Superhuman math doesn’t automatically mean superhuman judgment everywhere else.
3. Real AI Research is Harder than Benchmarks or Forecasts
Benchmarks and forecasts suggest that AI models should reliably succeed at coding tasks that take humans hours, without human help. The real world is messier. OpenAI’s internal data shows much shorter stretches of autonomous work on research tasks.
In its Research Acceleration / RSI report , OpenAI showed how often its models completed tasks with and without human help, grouped by how long a human would need to do the work.
Figure 4. OpenAI’s internal research tasks. Source .
Even on tasks that would take a human less than 15 minutes, OpenAI’s models succeeded without human intervention only 86% of the time. The estimated task length at 80% success was roughly 15 minutes over the first seven months of the year. July’s results were similar to the whole period average.
Fully autonomous RSI would require an AI to string together a great many research tasks reliably, stretching out over complex tasks that humans need weeks or months to accomplish. OpenAI’s data suggests that we aren’t close.
Anthropic also released a graph showing how Claude accelerates AI research. It shows that internal AI models collaborate on or even lead more than 90% of R&D tasks. That’s objectively impressive. At the same time, the graph reports zero cases of AI autonomously completing AI R&D tasks.
Figure 5. Claude’s role in internal AI R&D. Source .
These are incredible tools. But they still need skilled people to set direction and get them back on track.
The Gap Between Benchmarks and Reality
For years, METR has been publishing a chart showing what length of coding task (measured in human hours to complete) best-in-class AI models can achieve. It’s been called the most important graph in AI . METR’s Mythos Preview evaluation estimated that the model could succeed at 80% of coding tasks that took humans three hours.
Figure 6. METR’s 80% task horizons. Source .
From ECI Scores to METR Task Horizons
Epoch’s own rule of thumb is that every five additional points of ECI (their overall benchmark of AI capability) correspond to roughly a doubling of METR’s task horizon. Using that formula, we’d expect GPT 5.6 Sol and GPT 6 Astra to be 80% successful at completing tasks of around 4 hours and 11 hours of human length, respectively.
Another estimate (a forecast) of AI task length comes from the AI 2027 scenario, which estimated that by July 2026, frontier AIs would be 80% successful accomplishing tasks of around 11 hours. Fairly similar.
The AI 2027 Tracker charts all of these.
Figure 7. The AI 2027 Tracker. Source .
Inside OpenAI, though, the July research-task horizon at 80% success was roughly 15 minutes.
Measured Progress vs. AI 2027 and ECI-extrapolated METR
Figure 8. Forecasts, benchmarks, and real AI research. Tracker · OpenAI .
A four-hour benchmark horizon is about 16 times longer than OpenAI’s research horizon. AI 2027’s 11-hour forecast is about 44 times longer. Of course, the tasks being performed by researchers at OpenAI aren’t the same as those in the METR benchmark. So we should expect some discrepancy. This, however, goes well beyond that.
Actual AI research at OpenAI is an order of magnitude or more harder than metrics, benchmarks, or forecasts suggest. That should make us wary of relying too much on benchmarks, or of saying that future scenarios like AI 2027 are ‘on track.’ The authors of the related AI 2040 project still describe AI 2027 as roughly the future they expect, and say reality is tracking closer to it than even they expected. That’s not what we see from within OpenAI. This isn’t an apples-to-apples comparison, but the difference is remarkable. AI 2027 appears to be substantially over-optimistic in this regard.
In January of this year, Nathan Witkin made a case that the METR graph was exaggerating progress . The real world data suggests that at least some of his critiques were correct. The gap between benchmarks, forecasts, and data gleaned from actual use of AI should influence our expectations about the future.
4. The Sharp Diminishing Returns to Impressive AI Numbers
OpenAI’s report also shows impressive increases in AI token usage, in compute spend per researcher, and in lines of code written. But these aren’t results. They’re intermediate measures. How much progress do they actually drive?
Researchers used 124x more tokens per person. Engineers shipped roughly 7x as many lines of code per person. Researchers ran 1.6x as many experiments per researcher vs OpenAI’s 2025 whole year average.
Figure 9. Token use inside OpenAI. Source .
Figure 10. Experiment pace inside OpenAI. Source .
From More Tokens to More Experiments
Figure 11. From tokens to code to experiments. Source .
More tokens and code don’t tell us much on their own. The 1.6× experiment pace is closer to useful research output. Even that doesn’t mean AI is improving 1.6× faster.
An enormous increase in AI output has accompanied a much smaller increase in experiments run.
This isn’t a controlled experiment. We don’t know what would happen if researchers switched back to an older model. But it gives us a useful view of AI-assisted research inside a frontier lab.
It’s not just OpenAI. Anthropic reports that their engineers are now producing 8x as many lines of code per person as they did in 2024 - somewhat similar to OpenAI. Anthropic also sees significant diminishing returns between productivity and AI progress. Here’s a direct quote from its Mythos Preview system card :
“ Productivity uplift does not translate one-for-one to capabilities progress . We surveyed technical staff on the productivity uplift they experience from Claude Mythos Preview relative to zero AI assistance. The distribution is wide and the geometric mean is on the order of 4× . […] We estimate that reaching 2× on overall progress via this channel would require uplift roughly an order of magnitude larger than what we observe . ”- Anthropic, Claude Mythos Preview System Card; emphasis mine
Translation: To double the pace of AI progress, Anthropic estimates that AI would need to increase the productivity of their employees by roughly a factor of 40 relative to no AI assistance.
Figure 12. Anthropic’s productivity-to-progress estimate. Source .
This is an estimate, not a measurement of progress. Even the 4× productivity figure comes from an opt-in survey of 130 Anthropic staff . I put more weight on OpenAI’s logged experiments, though the two sources measure different things.
We don’t yet know how much those extra experiments are accelerating AI improvement, if at all. In general, there are also steeply diminishing returns of more experiments in most branches of science. That means that a 60% increase in experiment pace could be on the order of a 10% boost to AI improvement pace. (A power law exponent of 0.2, for those who want to do the math.) That’s speculation for now. We’ll learn more as the labs publish results.
Test Time Compute Also Has Diminishing Returns
What about giving the same AI model more time to think?
That scales badly also. In OpenAI’s re

[truncated]
