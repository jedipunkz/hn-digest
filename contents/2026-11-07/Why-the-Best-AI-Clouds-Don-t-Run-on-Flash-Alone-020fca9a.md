---
source: "https://www.backblaze.com/blog/why-the-best-ai-clouds-dont-run-on-flash-alone/"
hn_url: "https://news.ycombinator.com/item?id=50407455"
title: "Why the Best AI Clouds Don't Run on Flash Alone"
article_title: "Why AI Clouds Need a Tiered Storage Architecture"
image: "https://www.backblaze.com/blog/wp-content/uploads/2026/03/Q126-0009-Blog-Header-1440x820-1.png"
author: "gmays"
captured_at: "2026-11-07T19:55:22Z"
capture_tool: "hn-digest"
hn_id: 50407455
score: 1
comments: 0
posted_at: "2026-11-07T19:50:34Z"
tags:
  - hacker-news
---

# Why the Best AI Clouds Don't Run on Flash Alone

- HN: [50407455](https://news.ycombinator.com/item?id=50407455)
- Source: [www.backblaze.com](https://www.backblaze.com/blog/why-the-best-ai-clouds-dont-run-on-flash-alone/)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T19:50:34Z

## Translation

Title: Why the Best AI Clouds Don't Run on Flash Alone
Article title: Why AI Clouds Need a Tiered Storage Architecture
Description: Learn how tiered AI storage helps neoclouds control costs, avoid hyperscaler lock-in, and retain data across the entire AI lifecycle.

Article text:
Cloud Storage Tips, stories, insights, investigations, and how-tos about cloud storage.
Backing Up Topics related to backing up data, including backup strategies and techniques, and technical details and insights about backup.
Business Lab Bootstrapping, start-up tales and challenges, marketing mayhem, and more.
Backblaze Bits Topics related to Backblaze products and releases, as well as Backblaze new hires, job postings, things around the office, irreverent humor, and more.
Tech Lab Articles that explore our technical content more deeply, and are often geared towards our developer audience.
Partner News Your space to read up on our latest partnership announcements, partner content, and more.
Why the Best AI Clouds Don’t Run on Flash Alone
September 24, 2026
by
David Johnson
// No Comments
Neoclouds are built around GPUs, and GPUs can only run as fast as the data coming in. At the speeds AI training demands, conventional cloud storage solutions can’t keep up.
That’s where flash storage comes in. Flash is fast and reliable, purpose-built for workloads that need data delivered at extreme speed and low latency. But much of the storage work that surrounds active training—saving checkpoints, storing datasets, keeping finished models—doesn’t need flash-grade speed (or flash-grade prices).
That’s an issue for neoclouds. Their infrastructure is built around flash. So for storage that doesn’t need flash, they often point customers to hyperscaler object storage solutions instead. It’s cheaper, it works, and it keeps the neocloud focused on what it does best.
The problem is data gravity : the more data you have stored somewhere, the harder it becomes to leave. Hyperscalers make this worse by design. They let you put data in for free but charge prohibitively expensive egress fees to take it out. So once your data is entrenched on a hyperscaler’s platform, the cost of leaving keeps climbing. And wherever your data lives, that’s where you tend to buy more services. So while the neocloud keeps getting paid for GPU computing , the hyperscaler quietly takes over the broader account. What started as a sensible storage decision ends up with the hyperscaler owning everything except the GPU contract.
This article explains how a sensible storage decision can slowly hand a customer to a hyperscaler, and how a tiered storage architecture, with flash where it belongs and cheaper high-capacity storage everywhere else, lets neoclouds compete for the whole account.
Each lifecycle stage carries a different storage profile
To see why flash alone isn’t the answer, it helps to walk through what each stage of the AI lifecycle actually needs from storage.
It starts with data ingestion . Before training begins, raw datasets have to be cleansed, labeled, and prepared. This stage needs high-capacity storage that can move large volumes of data efficiently. Speed matters, but not to the degree that flash is necessary.
Active training is the stage where flash earns its place. Data moves through several storage tiers before it reaches the GPU, but at that final step, the model needs it delivered continuously and at very high speed . Only flash is fast enough to keep thousands of GPUs running without interruption.
Checkpointing runs throughout training. Because AI training is a synchronous process with every node in the cluster working in lockstep, a single component failure can bring the entire training run to a halt. And at the scale of a large modern cluster, failures are not rare . The failure rate also scales exponentially with cluster size. Checkpoints are the solution: frequent snapshots of the model’s current state that allow training to resume from where it left off. At large scale, a checkpoint may need to capture 15TB of data in under five seconds, repeated every few minutes, around the clock. This stage needs storage that is fast enough to keep up with that pace, and reliable enough that nothing gets lost.
Once training finishes, the outputs —model weights, logs, and metadata—go into longer-term storage. Fine-tuning experiments and additional training runs add more over time. This stage needs capacity and durability, not speed.
Finally, inference : serving a deployed model requires repeated access to model weights on demand. The data volumes are far smaller than training, and while latency-sensitive deployments may still benefit from flash, the storage demands are far less intensive.
Flash storage is purpose-built for only one of these stages (active training). At every other stage, what’s needed is a cost-effective, high-capacity storage tier, which is precisely what neoclouds are currently sending to hyperscalers. ( NVIDIA’s certified storage program validates storage solutions separately for each of these stages, recognizing that a single storage architecture can’t serve all of them well.)
Flash is the wrong tool for most of the job
Flash storage is expensive because it’s built for speed. But you pay for that speed whether you need it or not. Using flash to store datasets, checkpoints, and finished model artifacts is like shipping all of your mail overnight when most of it could go by ground; you’re paying a premium for a capability you don’t need.
And there’s a second problem beyond the cost. Flash that’s tied up storing checkpoints and datasets isn’t available for GPUs, and GPUs are what neoclouds sell. Every terabyte of flash used in the wrong place is capacity taken away from the core product.
The economics have always favored a tiered approach: fast storage where speed is needed, cheaper high-capacity storage everywhere else. In fact, Google, Meta, and Microsoft have run tiered storage architectures for years (Meta’s engineering team describes their approach in four explicit tiers, each selected for the job it does best).
But building out a full storage tier below flash is a significant undertaking, and most neoclouds have understandably stayed focused on their core GPU product. With SSD prices up 257% in under a year, that gap has become too expensive to ignore.
How a tiered stack maps to the lifecycle
The solution is to follow the same logic the hyperscalers have used for years: match the storage to the workload, not the other way around.
For active training, flash is still the answer. VAST, WEKA, and DDN handle this tier with fast, GPU-adjacent storage that keeps training running at full speed.
For everything else, B2 Neo can fill the gap. Datasets live here before training starts. Checkpoints write here continuously during training. Finished model weights and records accumulate here over time. The storage is fast and durable, but priced for capacity workloads rather than flash workloads. B2 Neo connects directly to the hot tier, so data moves between them without a detour through a hyperscaler.
The best part is the neocloud doesn’t have to build any of this. B2 Neo is available as a white-label product, deployed under the neocloud’s own brand. The customer only sees one seamless platform; their data stays on the neocloud’s infrastructure from raw dataset to deployed model and never touches a hyperscaler.
Flash is the right storage for one stage of the AI lifecycle. At every other stage, it’s an expensive tool for a job that doesn’t need it.
Neoclouds that only serve the active training tier leave a gap, and customers often fill that gap by sending their data to a hyperscaler. Once it’s there, egress fees keep it there. The neocloud ends up handing the broader customer relationship to whoever holds the data.
The good news is that closing the gap doesn’t require building anything from scratch. The tiered architecture that Google and Meta run at scale is available to neoclouds today as white-label infrastructure.
In this market, the GPU contract is just the beginning. Keeping the data is what keeps the customer.
Hit a wall trying to scale your storage infrastructure? B2 Neo gets you the rest of the way there, no rebuild required. Let’s talk .
David Johnson is the Director of Product Marketing at Backblaze, where he specializes in cloud backup and archiving for businesses. With extensive experience building the product marketing function at Vultr, he brings deep knowledge of the cloud infrastructure industry to Backblaze. David's passion for technology means his basement is a mini data center, filled with homelab projects where he spends his free time enhancing his knowledge of the industry and becoming a better informed expert on all things backup and archive. Connect with him on LinkedIn.
AI Workloads Are Breaking Old Storage Assumptions
November 5, 2026
by
Maddie Presland
Scaling AI Data Storage and Delivery with NVIDIA AIStore and B2
November 3, 2026
by
Jeronimo De Leon
The Mac Backup Problem Most IT Teams Don’t See Coming
October 27, 2026
by
Kari Wilson

## Original Extract

Learn how tiered AI storage helps neoclouds control costs, avoid hyperscaler lock-in, and retain data across the entire AI lifecycle.

Cloud Storage Tips, stories, insights, investigations, and how-tos about cloud storage.
Backing Up Topics related to backing up data, including backup strategies and techniques, and technical details and insights about backup.
Business Lab Bootstrapping, start-up tales and challenges, marketing mayhem, and more.
Backblaze Bits Topics related to Backblaze products and releases, as well as Backblaze new hires, job postings, things around the office, irreverent humor, and more.
Tech Lab Articles that explore our technical content more deeply, and are often geared towards our developer audience.
Partner News Your space to read up on our latest partnership announcements, partner content, and more.
Why the Best AI Clouds Don’t Run on Flash Alone
September 24, 2026
by
David Johnson
// No Comments
Neoclouds are built around GPUs, and GPUs can only run as fast as the data coming in. At the speeds AI training demands, conventional cloud storage solutions can’t keep up.
That’s where flash storage comes in. Flash is fast and reliable, purpose-built for workloads that need data delivered at extreme speed and low latency. But much of the storage work that surrounds active training—saving checkpoints, storing datasets, keeping finished models—doesn’t need flash-grade speed (or flash-grade prices).
That’s an issue for neoclouds. Their infrastructure is built around flash. So for storage that doesn’t need flash, they often point customers to hyperscaler object storage solutions instead. It’s cheaper, it works, and it keeps the neocloud focused on what it does best.
The problem is data gravity : the more data you have stored somewhere, the harder it becomes to leave. Hyperscalers make this worse by design. They let you put data in for free but charge prohibitively expensive egress fees to take it out. So once your data is entrenched on a hyperscaler’s platform, the cost of leaving keeps climbing. And wherever your data lives, that’s where you tend to buy more services. So while the neocloud keeps getting paid for GPU computing , the hyperscaler quietly takes over the broader account. What started as a sensible storage decision ends up with the hyperscaler owning everything except the GPU contract.
This article explains how a sensible storage decision can slowly hand a customer to a hyperscaler, and how a tiered storage architecture, with flash where it belongs and cheaper high-capacity storage everywhere else, lets neoclouds compete for the whole account.
Each lifecycle stage carries a different storage profile
To see why flash alone isn’t the answer, it helps to walk through what each stage of the AI lifecycle actually needs from storage.
It starts with data ingestion . Before training begins, raw datasets have to be cleansed, labeled, and prepared. This stage needs high-capacity storage that can move large volumes of data efficiently. Speed matters, but not to the degree that flash is necessary.
Active training is the stage where flash earns its place. Data moves through several storage tiers before it reaches the GPU, but at that final step, the model needs it delivered continuously and at very high speed . Only flash is fast enough to keep thousands of GPUs running without interruption.
Checkpointing runs throughout training. Because AI training is a synchronous process with every node in the cluster working in lockstep, a single component failure can bring the entire training run to a halt. And at the scale of a large modern cluster, failures are not rare . The failure rate also scales exponentially with cluster size. Checkpoints are the solution: frequent snapshots of the model’s current state that allow training to resume from where it left off. At large scale, a checkpoint may need to capture 15TB of data in under five seconds, repeated every few minutes, around the clock. This stage needs storage that is fast enough to keep up with that pace, and reliable enough that nothing gets lost.
Once training finishes, the outputs —model weights, logs, and metadata—go into longer-term storage. Fine-tuning experiments and additional training runs add more over time. This stage needs capacity and durability, not speed.
Finally, inference : serving a deployed model requires repeated access to model weights on demand. The data volumes are far smaller than training, and while latency-sensitive deployments may still benefit from flash, the storage demands are far less intensive.
Flash storage is purpose-built for only one of these stages (active training). At every other stage, what’s needed is a cost-effective, high-capacity storage tier, which is precisely what neoclouds are currently sending to hyperscalers. ( NVIDIA’s certified storage program validates storage solutions separately for each of these stages, recognizing that a single storage architecture can’t serve all of them well.)
Flash is the wrong tool for most of the job
Flash storage is expensive because it’s built for speed. But you pay for that speed whether you need it or not. Using flash to store datasets, checkpoints, and finished model artifacts is like shipping all of your mail overnight when most of it could go by ground; you’re paying a premium for a capability you don’t need.
And there’s a second problem beyond the cost. Flash that’s tied up storing checkpoints and datasets isn’t available for GPUs, and GPUs are what neoclouds sell. Every terabyte of flash used in the wrong place is capacity taken away from the core product.
The economics have always favored a tiered approach: fast storage where speed is needed, cheaper high-capacity storage everywhere else. In fact, Google, Meta, and Microsoft have run tiered storage architectures for years (Meta’s engineering team describes their approach in four explicit tiers, each selected for the job it does best).
But building out a full storage tier below flash is a significant undertaking, and most neoclouds have understandably stayed focused on their core GPU product. With SSD prices up 257% in under a year, that gap has become too expensive to ignore.
How a tiered stack maps to the lifecycle
The solution is to follow the same logic the hyperscalers have used for years: match the storage to the workload, not the other way around.
For active training, flash is still the answer. VAST, WEKA, and DDN handle this tier with fast, GPU-adjacent storage that keeps training running at full speed.
For everything else, B2 Neo can fill the gap. Datasets live here before training starts. Checkpoints write here continuously during training. Finished model weights and records accumulate here over time. The storage is fast and durable, but priced for capacity workloads rather than flash workloads. B2 Neo connects directly to the hot tier, so data moves between them without a detour through a hyperscaler.
The best part is the neocloud doesn’t have to build any of this. B2 Neo is available as a white-label product, deployed under the neocloud’s own brand. The customer only sees one seamless platform; their data stays on the neocloud’s infrastructure from raw dataset to deployed model and never touches a hyperscaler.
Flash is the right storage for one stage of the AI lifecycle. At every other stage, it’s an expensive tool for a job that doesn’t need it.
Neoclouds that only serve the active training tier leave a gap, and customers often fill that gap by sending their data to a hyperscaler. Once it’s there, egress fees keep it there. The neocloud ends up handing the broader customer relationship to whoever holds the data.
The good news is that closing the gap doesn’t require building anything from scratch. The tiered architecture that Google and Meta run at scale is available to neoclouds today as white-label infrastructure.
In this market, the GPU contract is just the beginning. Keeping the data is what keeps the customer.
Hit a wall trying to scale your storage infrastructure? B2 Neo gets you the rest of the way there, no rebuild required. Let’s talk .
David Johnson is the Director of Product Marketing at Backblaze, where he specializes in cloud backup and archiving for businesses. With extensive experience building the product marketing function at Vultr, he brings deep knowledge of the cloud infrastructure industry to Backblaze. David's passion for technology means his basement is a mini data center, filled with homelab projects where he spends his free time enhancing his knowledge of the industry and becoming a better informed expert on all things backup and archive. Connect with him on LinkedIn.
AI Workloads Are Breaking Old Storage Assumptions
November 5, 2026
by
Maddie Presland
Scaling AI Data Storage and Delivery with NVIDIA AIStore and B2
November 3, 2026
by
Jeronimo De Leon
The Mac Backup Problem Most IT Teams Don’t See Coming
October 27, 2026
by
Kari Wilson
