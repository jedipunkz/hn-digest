---
source: "https://arstechnica.com/ai/2026/09/exclusive-open-chinese-models-close-gap-with-silicon-valleys-frontier-ai-models/"
hn_url: "https://news.ycombinator.com/item?id=49711722"
title: "Paying for frontier AI models buys 4-month head start at 5x the cost"
article_title: "Exclusive: Paying for frontier AI models buys 4-month head start at 5x the cost - Ars Technica"
image: "https://cdn.arstechnica.net/wp-content/uploads/2026/09/China-US-AI-race-1152x648.jpg"
author: "QuantumNoodle"
captured_at: "2026-09-15T13:11:26Z"
capture_tool: "hn-digest"
hn_id: 49711722
score: 1
comments: 0
posted_at: "2026-09-15T12:47:32Z"
tags:
  - hacker-news
---

# Paying for frontier AI models buys 4-month head start at 5x the cost

- HN: [49711722](https://news.ycombinator.com/item?id=49711722)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/09/exclusive-open-chinese-models-close-gap-with-silicon-valleys-frontier-ai-models/)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T12:47:32Z

## Translation

Title: Paying for frontier AI models buys 4-month head start at 5x the cost
Article title: Exclusive: Paying for frontier AI models buys 4-month head start at 5x the cost - Ars Technica
Description: Ars previewed Mozilla’s report on how cheap open models caught up on capability.

Article text:
Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Mind the gap
Exclusive: Paying for frontier AI models buys 4-month head start at 5x the cost
Ars previewed Mozilla’s report on how cheap open models caught up on capability.
14
Credit:
Imen Ben Youssef / Hans Lucas / AFP via Getty Images
Credit:
Imen Ben Youssef / Hans Lucas / AFP via Getty Images
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
The performance gap between frontier AI models from US tech companies and the best open-weights models from Chinese companies has closed to just 4.4 months, according to a Mozilla report. That explains why many companies are shifting to the significantly cheaper open models for routine work—and helps reveal a narrow band of workloads where frontier models are worth the cost.
Most organizations should ideally be using open models as the default for the majority of their work, according to the latest State of Open Source AI report from Mozilla, published on September 15 and shared with Ars prior to publication. The report highlights how a leading open model, Moonshot AI’s Kimi K3, achieves a composite AI performance score on the Artificial Analysis Intelligence Index that is just three points behind Anthropic’s Fable 5 closed frontier model, all while costing just 30 percent of the latter.
“Closed earns its premium in a few places: expert professional work, high-intensity retrieval, and long context,” Raffi Krikorian, chief technology officer at Mozilla, said in an email to Ars. “We see the decision to pay for closed as workload-specific rather than organization-specific.”
The open-weights AI models allow anyone to download the main model components and run the models on their own computers, but developers still typically withhold vital information such as training data, the data pipeline, and training code. By comparison, US tech companies such as Anthropic and OpenAI mostly offer closed frontier models that keep everything proprietary, requiring customers to pay more for access.
Organizations still pay for closed frontier models because they work out of the box and come bundled with “compliance packaging, support, and accountability,” whereas many organizations lack the staff to run open-weights models well, Krikorian explained.
But the steadily narrowing performance gap between open and closed frontier models, coupled with the cost-effectiveness of open models, is now widely recognized by many companies. For example, the delivery company DoorDash has been using Kimi for routine work while reserving Fable for more difficult tasks that would normally take human experts longer to complete.
Unsurprisingly, this has accelerated the usage of open models since Mozilla’s inaugural State of Open Source AI report was published on July 14.
The narrowing performance gap is being measured in several ways. For example, the research nonprofit METR has defined an AI model’s time horizon as the length of tasks—measured by how long human experts require to complete them—that can be handled by AI models with a “reliable” 50 percent success rate. That time horizon has been doubling on a set cadence that has accelerated over time.
The best closed model can currently do a job that is 1.7 times as long as the longest job that the best open model can reliably finish.
“If the open frontier can handle a seven-hour job, the closed frontier can handle a 12-hour one,” Krikorian told Ars. “In four months, the open model handles the 12-hour job, and the closed one handles something around 20.”
Tasks requiring between eight and 12 hours are typically the ones that a closed frontier model can do and the open models cannot handle yet, Krikorian said. No models are generally capable of reliably doing tasks longer than 12 hours, whereas tasks taking less than eight hours can be done by either type of model and could be handed off to the cheaper open models.
There are many nuances and caveats for such comparisons. For example, closed frontier models often come with their own harness —the software layer that helps the models access various tools and memory to perform agent-like actions. Harnesses custom-built by AI labs for their models can sometimes boost performance on various tasks, whereas the models may perform less well with third-party harnesses.
To level the playing field, the benchmarking company Vals AI has been evaluating different open and closed models by using its own neutral harness. When every model ran on the same harness in the Terminal-Bench 2.1 evaluation , the open-weights model GLM 5.2 from the Chinese company Z.ai (Zhipu AI) scored within a point of Anthropic’s Claude Opus 4.7 and 4.8 while costing about five times less per completed task.
In other words, paying for closed frontier models buys about a four-month head start at about five times the per-task cost—but only when currently looking at tasks taking between eight and 12 hours.
“Pay when that head start is worth it, something like a deadline that lands before the open frontier catches up would be here,” Krikorian told Ars. “Routine work you’ll still be doing next quarter is not, because you’ll be able to do it for a fifth of the cost soon, and the model won’t be the bottleneck anyway.”
The surging popularity of open-weights models can be seen on OpenRouter , the AI gateway and marketplace that gives developers access to hundreds of different models from multiple providers. The Mozilla report noted that eight of the top 10 models ranked by token volumes in August 2026 provide open weights.
However, open models are still lagging behind closed frontier models in revenue. A paper by Frank Nagle and Daniel Yue for the Linux Foundation found that open models earned just 4 percent of overall revenue compared to closed models getting 96 percent of revenue—although they were looking at data from May through September of 2025.
“In the last year, open models have exploded, so we do expect revenue to have shifted,” Krikorian said.
The current reality is that “most of the open models the world runs on are Chinese,” Krikorian said. “The Chinese labs are running the same playbook the Americans ran with Android—give it away, but own the ecosystem around it,” he explained.
But Krikorian is not so concerned about “Chinese models” as he is about the risk of concentration. At the moment, the best open models are concentrated in China, whereas the best closed frontier models come from US companies. He argued for US and European labs to start competing in the “same open lane, so that no single country sets the world’s defaults.”
“The uncomfortable truth is that the plural ecosystem around open-weight AI is largely funded by Chinese capital right now,” Krikorian said. “That’s a plurality and a concentration at the same time.”
The history of development for open source software such as Linux may offer some lessons for how an “alternative coalition” could come together, Krikorian said. Such open source software infrastructure was funded by a coalition of organizations that “each needed the commodity layer to exist,” including neutral foundations.
Any coalition for building more open AI models would likely consist of “institutions with a mission rather than a market,” as opposed to frontier AI labs, Krikorian said. He described it as follows:
We need public compute programs funding fully open reference models. Switzerland’s national compute producing Apertus is an example. Foundations need to hold the same neutral ground that they held for the internet’s open protocols, extended into models, harnesses, and agentic standards. Companies that benefit from commodity models have to participate. And philanthropy has to cover what none of the others will—evaluation and audit infrastructure, especially.
Krikorian also called for the alternative coalition to go even further than open-weights models by embracing more of the open source approach. “It’s hard to fully trust a model with decisions if you can’t tell how it was trained or what it was evaluated against,” he said.
14 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
"Offensively cheap": Solar power is looking up
2.
Steam Frame: The Ars Technica review
3.
For the first time, the US military confirms it has deployed weapons in orbit
4.
Rocket Lab is seeing red about NASA's decision on a Mars spacecraft
5.
Apple releases iOS 27, macOS Golden Gate 27 with Siri AI and Liquid Glass refinements
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.

## Original Extract

Ars previewed Mozilla’s report on how cheap open models caught up on capability.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Mind the gap
Exclusive: Paying for frontier AI models buys 4-month head start at 5x the cost
Ars previewed Mozilla’s report on how cheap open models caught up on capability.
14
Credit:
Imen Ben Youssef / Hans Lucas / AFP via Getty Images
Credit:
Imen Ben Youssef / Hans Lucas / AFP via Getty Images
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
The performance gap between frontier AI models from US tech companies and the best open-weights models from Chinese companies has closed to just 4.4 months, according to a Mozilla report. That explains why many companies are shifting to the significantly cheaper open models for routine work—and helps reveal a narrow band of workloads where frontier models are worth the cost.
Most organizations should ideally be using open models as the default for the majority of their work, according to the latest State of Open Source AI report from Mozilla, published on September 15 and shared with Ars prior to publication. The report highlights how a leading open model, Moonshot AI’s Kimi K3, achieves a composite AI performance score on the Artificial Analysis Intelligence Index that is just three points behind Anthropic’s Fable 5 closed frontier model, all while costing just 30 percent of the latter.
“Closed earns its premium in a few places: expert professional work, high-intensity retrieval, and long context,” Raffi Krikorian, chief technology officer at Mozilla, said in an email to Ars. “We see the decision to pay for closed as workload-specific rather than organization-specific.”
The open-weights AI models allow anyone to download the main model components and run the models on their own computers, but developers still typically withhold vital information such as training data, the data pipeline, and training code. By comparison, US tech companies such as Anthropic and OpenAI mostly offer closed frontier models that keep everything proprietary, requiring customers to pay more for access.
Organizations still pay for closed frontier models because they work out of the box and come bundled with “compliance packaging, support, and accountability,” whereas many organizations lack the staff to run open-weights models well, Krikorian explained.
But the steadily narrowing performance gap between open and closed frontier models, coupled with the cost-effectiveness of open models, is now widely recognized by many companies. For example, the delivery company DoorDash has been using Kimi for routine work while reserving Fable for more difficult tasks that would normally take human experts longer to complete.
Unsurprisingly, this has accelerated the usage of open models since Mozilla’s inaugural State of Open Source AI report was published on July 14.
The narrowing performance gap is being measured in several ways. For example, the research nonprofit METR has defined an AI model’s time horizon as the length of tasks—measured by how long human experts require to complete them—that can be handled by AI models with a “reliable” 50 percent success rate. That time horizon has been doubling on a set cadence that has accelerated over time.
The best closed model can currently do a job that is 1.7 times as long as the longest job that the best open model can reliably finish.
“If the open frontier can handle a seven-hour job, the closed frontier can handle a 12-hour one,” Krikorian told Ars. “In four months, the open model handles the 12-hour job, and the closed one handles something around 20.”
Tasks requiring between eight and 12 hours are typically the ones that a closed frontier model can do and the open models cannot handle yet, Krikorian said. No models are generally capable of reliably doing tasks longer than 12 hours, whereas tasks taking less than eight hours can be done by either type of model and could be handed off to the cheaper open models.
There are many nuances and caveats for such comparisons. For example, closed frontier models often come with their own harness —the software layer that helps the models access various tools and memory to perform agent-like actions. Harnesses custom-built by AI labs for their models can sometimes boost performance on various tasks, whereas the models may perform less well with third-party harnesses.
To level the playing field, the benchmarking company Vals AI has been evaluating different open and closed models by using its own neutral harness. When every model ran on the same harness in the Terminal-Bench 2.1 evaluation , the open-weights model GLM 5.2 from the Chinese company Z.ai (Zhipu AI) scored within a point of Anthropic’s Claude Opus 4.7 and 4.8 while costing about five times less per completed task.
In other words, paying for closed frontier models buys about a four-month head start at about five times the per-task cost—but only when currently looking at tasks taking between eight and 12 hours.
“Pay when that head start is worth it, something like a deadline that lands before the open frontier catches up would be here,” Krikorian told Ars. “Routine work you’ll still be doing next quarter is not, because you’ll be able to do it for a fifth of the cost soon, and the model won’t be the bottleneck anyway.”
The surging popularity of open-weights models can be seen on OpenRouter , the AI gateway and marketplace that gives developers access to hundreds of different models from multiple providers. The Mozilla report noted that eight of the top 10 models ranked by token volumes in August 2026 provide open weights.
However, open models are still lagging behind closed frontier models in revenue. A paper by Frank Nagle and Daniel Yue for the Linux Foundation found that open models earned just 4 percent of overall revenue compared to closed models getting 96 percent of revenue—although they were looking at data from May through September of 2025.
“In the last year, open models have exploded, so we do expect revenue to have shifted,” Krikorian said.
The current reality is that “most of the open models the world runs on are Chinese,” Krikorian said. “The Chinese labs are running the same playbook the Americans ran with Android—give it away, but own the ecosystem around it,” he explained.
But Krikorian is not so concerned about “Chinese models” as he is about the risk of concentration. At the moment, the best open models are concentrated in China, whereas the best closed frontier models come from US companies. He argued for US and European labs to start competing in the “same open lane, so that no single country sets the world’s defaults.”
“The uncomfortable truth is that the plural ecosystem around open-weight AI is largely funded by Chinese capital right now,” Krikorian said. “That’s a plurality and a concentration at the same time.”
The history of development for open source software such as Linux may offer some lessons for how an “alternative coalition” could come together, Krikorian said. Such open source software infrastructure was funded by a coalition of organizations that “each needed the commodity layer to exist,” including neutral foundations.
Any coalition for building more open AI models would likely consist of “institutions with a mission rather than a market,” as opposed to frontier AI labs, Krikorian said. He described it as follows:
We need public compute programs funding fully open reference models. Switzerland’s national compute producing Apertus is an example. Foundations need to hold the same neutral ground that they held for the internet’s open protocols, extended into models, harnesses, and agentic standards. Companies that benefit from commodity models have to participate. And philanthropy has to cover what none of the others will—evaluation and audit infrastructure, especially.
Krikorian also called for the alternative coalition to go even further than open-weights models by embracing more of the open source approach. “It’s hard to fully trust a model with decisions if you can’t tell how it was trained or what it was evaluated against,” he said.
14 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
"Offensively cheap": Solar power is looking up
2.
Steam Frame: The Ars Technica review
3.
For the first time, the US military confirms it has deployed weapons in orbit
4.
Rocket Lab is seeing red about NASA's decision on a Mars spacecraft
5.
Apple releases iOS 27, macOS Golden Gate 27 with Siri AI and Liquid Glass refinements
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
