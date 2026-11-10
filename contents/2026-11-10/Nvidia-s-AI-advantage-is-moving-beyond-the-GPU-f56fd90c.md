---
source: "https://techcrunch.com/2026/08/29/nvidias-ai-advantage-is-moving-beyond-the-gpu/"
hn_url: "https://news.ycombinator.com/item?id=50435539"
title: "Nvidia's AI advantage is moving beyond the GPU"
article_title: "Nvidia’s AI advantage is moving beyond the GPU | TechCrunch"
image: "https://techcrunch.com/wp-content/uploads/2026/08/jense-nvidi-chip-GettyImages-2266485392.jpg?resize=1200,800"
author: "gmays"
captured_at: "2026-11-10T03:13:39Z"
capture_tool: "hn-digest"
hn_id: 50435539
score: 1
comments: 0
posted_at: "2026-11-10T02:59:27Z"
tags:
  - hacker-news
---

# Nvidia's AI advantage is moving beyond the GPU

- HN: [50435539](https://news.ycombinator.com/item?id=50435539)
- Source: [techcrunch.com](https://techcrunch.com/2026/08/29/nvidias-ai-advantage-is-moving-beyond-the-gpu/)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T02:59:27Z

## Translation

Title: Nvidia's AI advantage is moving beyond the GPU
Article title: Nvidia’s AI advantage is moving beyond the GPU | TechCrunch
Description: The new generation of data center systems is increasing efficiency with smarter traffic control instead of just more processor cycles.

Article text:
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Nvidia’s AI advantage is moving beyond the GPU
Before this week, the dominant story about Nvidia went something like this: For the first few years of the AI boom, Nvidia was the only source for state-of-the-art GPUs, which became immensely profitable as the industry scaled out. In the last few years, hyperscalers like Amazon and Google have started building their own chips, and Nvidia is no longer the only game in town, leading many investors to wonder how durable its advantage really is.
It’s a compelling story, and mostly true. After growing its market cap 10x between the start of 2023 and mid-2025, Nvidia shares have been on a more modest trajectory for the past year, driven by concerns about GPU competition.
A new narrative has taken shape since the company’s earnings on Wednesday and investors are starting to realize that Nvidia’s advantage goes far beyond GPUs. As AI’s compute grows into the gigawatt scale, orchestration has become an increasingly complex task. Not surprisingly, Nvidia has built much of the state-of-the-art hardware needed to handle it, giving the company a huge advantage in the systems that surround the GPU even as it sees increased competition on the GPUs themselves.
For all the talk of compute as a commodity , it’s still incredibly difficult to operate a mega-scale data center at peak efficiency — and as deployments get bigger and faster, that challenge is only growing.
You can see some of this just by looking at the details of what Nvidia is actually selling. The company is currently rolling out its Vera Rubin architecture, which pairs the Rubin GPU with a collection of other units, including the Vera CPU, the Groq 3 LPX inference accelerator, and similar racks for storage and networking.
Over the past week, I’ve been talking to folks at Nvidia about what those systems actually do, and the results have been surprising. Like the Rubin GPU, they’re extremely specialized systems, but instead of churning through tokens, they’re making sure everything outside the GPU works as efficiently as possible. If the GPU is the engine, these are the rest of the car.
The Vera CPU in particular is focused on the problem of orchestrating data. “Vera is important because there’s only so much memory that you can put in a single server or any sort of compute platform,” Jason Hardy, Nvidia’s VP of storage technology, told me.
As data centers have scaled up computing power, memory capacity has scaled up, too, which is why companies like Micron have gotten rich in the second wave of the infrastructure boom. But getting that data to the GPU at the right time isn’t straightforward — and as companies look to drive tokens-per-watt lower and lower, they’re realizing how important that kind of traffic direction is.
“We saw upwards of 3x improvement in these operations, where the Vera CPU is allowing for acceleration,” Hardy said. “So now we can use our flash to its fullest potential, because we can get all that performance out of it without bottlenecking.”
You can see versions of the same problem outside of Nvidia. When OpenAI developed its Jalapeño chip, a major focus was avoiding these challenges entirely by minimizing the amount of data that needs to be moved around.
“We designed Jalapeño to minimize data movement and communication delays,” the company said in a blog post earlier this month . “Its large domain allows the entire workload to remain within one connected system, minimizing data movement and helping the complete request stay fast and efficient from beginning to end.”
It’s a different approach, avoiding data movement entirely by conducting a workload within one integrated chip. But the overall logic is the same, increasing efficiency with smarter traffic control instead of just more processor cycles. That in turn opens up a whole new layer of infrastructure for companies to compete over.
This new focus on data orchestration isn’t automatically a win for Nvidia. The company will have to compete with rival chipmakers and hyperscalers just as it has with GPUs. But the competition has moved to a new layer, where building a rival GPU matters less than being able to make the entire system work efficiently.
And at least in the early stages, Nvidia looks to have a commanding lead.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
2027
San Francisco
Get an Early Peek at the Next TechCrunch Disrupt
Disrupt, Startup Battlefield, and more will return next year, and you can be at the front of the line.
Sign up for our waitlist
Musk says Tesla will refund some Cybertruck deposits, defends decision to end production
Tesla will stop making the Cybertruck in 2027
Claude can now analyze your spending with its new Personal Finance tool
Harbinger raises $240M after expanding into defense
Almost no one uses AI agents to book their Uber, company says
Meta acquires AI coding startup Conductor
Data centers are turning to old jet engines as the wait for power grows

## Original Extract

The new generation of data center systems is increasing efficiency with smarter traffic control instead of just more processor cycles.

TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Nvidia’s AI advantage is moving beyond the GPU
Before this week, the dominant story about Nvidia went something like this: For the first few years of the AI boom, Nvidia was the only source for state-of-the-art GPUs, which became immensely profitable as the industry scaled out. In the last few years, hyperscalers like Amazon and Google have started building their own chips, and Nvidia is no longer the only game in town, leading many investors to wonder how durable its advantage really is.
It’s a compelling story, and mostly true. After growing its market cap 10x between the start of 2023 and mid-2025, Nvidia shares have been on a more modest trajectory for the past year, driven by concerns about GPU competition.
A new narrative has taken shape since the company’s earnings on Wednesday and investors are starting to realize that Nvidia’s advantage goes far beyond GPUs. As AI’s compute grows into the gigawatt scale, orchestration has become an increasingly complex task. Not surprisingly, Nvidia has built much of the state-of-the-art hardware needed to handle it, giving the company a huge advantage in the systems that surround the GPU even as it sees increased competition on the GPUs themselves.
For all the talk of compute as a commodity , it’s still incredibly difficult to operate a mega-scale data center at peak efficiency — and as deployments get bigger and faster, that challenge is only growing.
You can see some of this just by looking at the details of what Nvidia is actually selling. The company is currently rolling out its Vera Rubin architecture, which pairs the Rubin GPU with a collection of other units, including the Vera CPU, the Groq 3 LPX inference accelerator, and similar racks for storage and networking.
Over the past week, I’ve been talking to folks at Nvidia about what those systems actually do, and the results have been surprising. Like the Rubin GPU, they’re extremely specialized systems, but instead of churning through tokens, they’re making sure everything outside the GPU works as efficiently as possible. If the GPU is the engine, these are the rest of the car.
The Vera CPU in particular is focused on the problem of orchestrating data. “Vera is important because there’s only so much memory that you can put in a single server or any sort of compute platform,” Jason Hardy, Nvidia’s VP of storage technology, told me.
As data centers have scaled up computing power, memory capacity has scaled up, too, which is why companies like Micron have gotten rich in the second wave of the infrastructure boom. But getting that data to the GPU at the right time isn’t straightforward — and as companies look to drive tokens-per-watt lower and lower, they’re realizing how important that kind of traffic direction is.
“We saw upwards of 3x improvement in these operations, where the Vera CPU is allowing for acceleration,” Hardy said. “So now we can use our flash to its fullest potential, because we can get all that performance out of it without bottlenecking.”
You can see versions of the same problem outside of Nvidia. When OpenAI developed its Jalapeño chip, a major focus was avoiding these challenges entirely by minimizing the amount of data that needs to be moved around.
“We designed Jalapeño to minimize data movement and communication delays,” the company said in a blog post earlier this month . “Its large domain allows the entire workload to remain within one connected system, minimizing data movement and helping the complete request stay fast and efficient from beginning to end.”
It’s a different approach, avoiding data movement entirely by conducting a workload within one integrated chip. But the overall logic is the same, increasing efficiency with smarter traffic control instead of just more processor cycles. That in turn opens up a whole new layer of infrastructure for companies to compete over.
This new focus on data orchestration isn’t automatically a win for Nvidia. The company will have to compete with rival chipmakers and hyperscalers just as it has with GPUs. But the competition has moved to a new layer, where building a rival GPU matters less than being able to make the entire system work efficiently.
And at least in the early stages, Nvidia looks to have a commanding lead.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
2027
San Francisco
Get an Early Peek at the Next TechCrunch Disrupt
Disrupt, Startup Battlefield, and more will return next year, and you can be at the front of the line.
Sign up for our waitlist
Musk says Tesla will refund some Cybertruck deposits, defends decision to end production
Tesla will stop making the Cybertruck in 2027
Claude can now analyze your spending with its new Personal Finance tool
Harbinger raises $240M after expanding into defense
Almost no one uses AI agents to book their Uber, company says
Meta acquires AI coding startup Conductor
Data centers are turning to old jet engines as the wait for power grows
