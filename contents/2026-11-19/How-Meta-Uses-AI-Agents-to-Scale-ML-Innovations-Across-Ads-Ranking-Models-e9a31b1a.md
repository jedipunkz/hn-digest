---
source: "https://engineering.fb.com/2026/11/18/ml-applications/ai-agents-to-scale-ml-innovations-across-ads-ranking-models/"
hn_url: "https://news.ycombinator.com/item?id=50561912"
title: "How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models"
article_title: "How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models - Engineering at Meta"
image: "https://engineering.fb.com/wp-content/uploads/2026/11/Meta-Scaling-Techniques-image-2-2.png"
author: "gmays"
captured_at: "2026-11-19T16:21:26Z"
capture_tool: "hn-digest"
hn_id: 50561912
score: 2
comments: 0
posted_at: "2026-11-19T16:06:28Z"
tags:
  - hacker-news
---

# How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models

- HN: [50561912](https://news.ycombinator.com/item?id=50561912)
- Source: [engineering.fb.com](https://engineering.fb.com/2026/11/18/ml-applications/ai-agents-to-scale-ml-innovations-across-ads-ranking-models/)
- Score: 2
- Comments: 0
- Posted: 2026-11-19T16:06:28Z

## Translation

Title: How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models
Article title: How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models - Engineering at Meta
Description: Our previous post in this series introduced the Ranking Engineer Agent (REA), an autonomous system that generates hypotheses, designs experiments, and proves new machine learning (ML) techniques on…

Article text:
How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models - Engineering at Meta
Skip to content
Infrastructure Systems
Core Infra
Physical Infrastructure
Connectivity
Video Engineering & AR/VR
Video Engineering
Artificial Intelligence
ML Applications
How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models
Our previous post in this series introduced the Ranking Engineer Agent (REA) , an autonomous system that generates hypotheses, designs experiments, and proves new machine learning (ML) techniques on ads ranking models.
Now, we’re sharing how we propagate ranking model improvements like architectural advancements and training optimizations efficiently across a large and diverse fleet of ads ranking models.
This post walks through this process that we call “technique propagation” and how we’re optimizing the process using AI agents.
Since agents began driving execution in March 2026, we estimate AI agents handle over 60% of all technique propagation. Measured against the same work before agents, early results show 3x greater engineering productivity, a median 4x reduction in training compute across select combinations, and an estimated $45M+ in compute savings by the end of the year.
Meta’s ads recommendation system is powered by a large number of distinct ads ranking models, each serving different product surfaces and advertiser objectives, and each with its own architecture, training data, and serving constraints.
Across this portfolio, ML researchers and engineers are continuously advancing the state-of-the-art through innovations we call “techniques.” A technique is anything that makes a ranking model better such as architectural advancements like Wukong factorization layers, Interformer interaction networks, richer data representations, training optimizations, or efficiency improvements in training and inference.
However, because each model has its own architecture, data, and training and serving constraints, a technique proven on one model almost never copy-pastes onto the next. At the same time, speed matters because ML techniques have a shelf life. An innovation that delivers gains today may be superseded within a couple of model development cycles. The faster proven techniques reach the full model fleet, the more value they deliver before the next generation of techniques arrive.
So how do we propagate a technique to a diverse fleet of models as quickly as possible, while minimizing engineering bandwidth and compute?
Propagating a growing inventory of proven techniques across a large fleet of diverse models creates a combinatorially expanding space of technique-model combinations to explore, where both engineering bandwidth and compute become significant constraints.
The standard way to adapt a technique is to have engineers who are domain experts manually execute adaptation. An engineer studies how a technique interacts with a specific model’s architecture, generalizes it, and tunes it to that model’s constraints. This produces the strongest results, and is still the process for our highest-revenue models, but it does not scale to all of the models in the space. Each adaptation is effort-intensive and demands deep familiarity with both the technique and the target model. So throughput is capped by engineering bandwidth while the opportunity for adaptation grows much faster.
The second constraint on propagation is compute. Every adaptation attempt requires a training run, and with a finite compute budget the question is not only which techniques to try but how to navigate the combinatorial space of configurations without spending the budget on attempts (most of which will produce suboptimal gains).
Both constraints point to the same requirement: To propagate techniques to a large number of diverse models quickly, propagation has to become standardized, automated, and able to make judgment calls.
We started by standardizing model design and automating propagation: a template-driven approach organizes models into a small set of reusable and composable building blocks, so a technique validated once at the template level propagates automatically to every model built from that template. Over the last 2 years, through standardization and automation, we have shipped 5,000+ technique-model combinations across our large, diverse fleet of models while reducing per-model engineering time by 92%. You can read more about it in our post on Standard Model Templates (SMT) .
Standardization and automation handle the deterministic steps, but successful and efficient propagation still requires making judgment calls, such as which technique is worth propagating to which model, how wide a hyperparameter search should be, why a model is not showing gains, or how to detect and resolve failures. Those decisions still required significant human bandwidth, which prevented us from scaling further and from reaching a state where all techniques reached all applicable models within a few weeks of being proven.
That is where AI agents come in.
Scaling Technique Propagation With AI Agents
Our technique propagation system is built around three capabilities:
Discovery: Keep a live view of which techniques have been proven and estimate how well each might transfer to generate a prioritized set of technique-model candidates, ranked by expected impact.
Generalization: Generalize a proven technique so it works on models it was not originally designed for. For models built from shared templates, a technique is generalized once per template and reaches every model built from it. Beyond what templates cover, each model needs its own adaptation, which is where the deepest technical reasoning happens.
Adoption: Turn validated adaptations into shipped models, many at a time: launching experiments, managing asynchronous multi-day training runs, monitoring convergence, diagnosing and recovering from failures, and assembling proposals that can land.
Today, AI agents supercharge all the above capabilities by taking on the judgment-heavy, non-deterministic decision making and bringing deep intelligence to the system. They combine LLM-based reasoning with specialized ML tooling to handle the work of deciding which adaptations to try, diagnosing why something failed and automatically mitigating it, navigating hyperparameter search spaces efficiently, and managing capacity across hundreds of concurrent jobs. The agents share a common knowledge base, an accumulated record of prior training runs, model architectures, and technique outcomes, so they build on what has already been learned rather than starting from scratch.
Discovery answers two questions: What proven techniques exist, and how applicable are they to each model in the model fleet? Until recently, domain experts answered both, working through the assessments one model at a time. They tracked research and results across the organization, spotted the techniques worth propagating, and relied on their domain knowledge to judge applicability.
Agents aren’t constrained by the same knowledge or engineering bandwidth limitations as domain experts. Instead of relying on what any one expert knows, they draw on the shared knowledge base, and they make assessments across the entire model fleet at once rather than one model at a time. They systematically identify which proven techniques exist, estimate their transferability across the model fleet, and produce prioritized technique lists ranked by expected impact. Engineers now only have to review the agents’ assessments and make a decision from there.
How Agents Make Generalization More Computationally Efficient
Generalization is the process of re-implementing a technique proven on one model so it applies to models with different architectures, feature representations, training data, and serving constraints. The core idea usually transfers, but the implementation almost never does. The customization needed depends on the technique and the target model, ranging from retuning hyperparameters, to adapting features to what the target model already consumes or resizing the architecture to fit its inference budget, to redesigning the technique outright. Deciding what a given model needs, and how far to go, is a judgment call.
Depending on whether a model is built from standardized templates or customized individually, generalization comes in two flavors, and agents assist with both:
Template generalization: For models built from standardized templates , agents assist with building, evaluating, and improving the templates themselves. They propose template optimization hypotheses from proven techniques, launch experiments across representative model sets, evaluate results, diagnose failures, and iterate through multiple rounds, helping manage the multi-day loops that template development requires. Validating a technique once at the template level costs a fraction of the compute of adapting each model separately.
Bespoke generalization: For our highest-revenue models, where we still customize a technique for each model, agents progressively learn while working alongside ML engineers during the initial application of a technique to a small number of models. As the engineer makes adaptation decisions such as adjusting configurations, tuning hyperparameters, or resolving failures, the agent tracks every change and outcome, encoding what works, what doesn’t, and which decisions are model-specific versus generalizable. With that grounding, the agent proposes and tests adaptations for the remaining models, with the engineer reviewing and steering rather than driving each one.
Drawing on millions of prior training runs and each target model’s own properties, agents start from the configurations most likely to work, use the results of each round of training runs to narrow the hyperparameter search space for the next, and decide which experiments are worth continuing and which to stop early. As a result, when measured across select technique-model combinations, we have seen a median 4x reduction in training compute.
How Agents Accelerate Adoption
At the adoption stage, agents ship generalized techniques model-by-model. They launch experiments, manage asynchronous multi-day training runs, fine-tune each model, monitor convergence, recover from failures, and assemble validated proposals.
Fine-tuning happens differently depending on how the technique arrived. For techniques applied individually, agents help find the best configuration of technique bundles for that particular model. For models built from a shared template, an auto-research loop searches on top of the template, tuning each model for the gains the shared configuration alone does not deliver.
When failures occur, whether from an infrastructure issue, training divergence, or a data pipeline error, integrated debugging and recovery diagnose the root cause and apply targeted fixes or retries.
Running hundreds of these workflows at once raises the question of which jobs get priority, and which runs are worth continuing or stopping early. Getting the most out of a fixed compute budget is a global optimization problem, and it is where we are currently investing. Agents monitor concurrent training jobs, rank them by revenue importance, technique potential, and deadline proximity, and reclaim capacity from stalled or low-value runs, so GPUs stay pointed at the highest-value work rather than sitting idle or queued behind lower-priority experiments.
In the past, an engineer had to watch each run, triage every failure by hand, and restart it, which capped each of them at a handful of models. Now they can keep oversight at the decision points, approving plans and setting priorities at a 3x greater productivity (each engineer can now handle 3x more technique-model adaptations), compressing the time from a proven technique to a shipped model.
How We Evaluate Agent-Driven Propagation
Techniques propagated by agents are evaluated the same way as those prop

[truncated]

## Original Extract

Our previous post in this series introduced the Ranking Engineer Agent (REA), an autonomous system that generates hypotheses, designs experiments, and proves new machine learning (ML) techniques on…

How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models - Engineering at Meta
Skip to content
Infrastructure Systems
Core Infra
Physical Infrastructure
Connectivity
Video Engineering & AR/VR
Video Engineering
Artificial Intelligence
ML Applications
How Meta Uses AI Agents to Scale ML Innovations Across Ads Ranking Models
Our previous post in this series introduced the Ranking Engineer Agent (REA) , an autonomous system that generates hypotheses, designs experiments, and proves new machine learning (ML) techniques on ads ranking models.
Now, we’re sharing how we propagate ranking model improvements like architectural advancements and training optimizations efficiently across a large and diverse fleet of ads ranking models.
This post walks through this process that we call “technique propagation” and how we’re optimizing the process using AI agents.
Since agents began driving execution in March 2026, we estimate AI agents handle over 60% of all technique propagation. Measured against the same work before agents, early results show 3x greater engineering productivity, a median 4x reduction in training compute across select combinations, and an estimated $45M+ in compute savings by the end of the year.
Meta’s ads recommendation system is powered by a large number of distinct ads ranking models, each serving different product surfaces and advertiser objectives, and each with its own architecture, training data, and serving constraints.
Across this portfolio, ML researchers and engineers are continuously advancing the state-of-the-art through innovations we call “techniques.” A technique is anything that makes a ranking model better such as architectural advancements like Wukong factorization layers, Interformer interaction networks, richer data representations, training optimizations, or efficiency improvements in training and inference.
However, because each model has its own architecture, data, and training and serving constraints, a technique proven on one model almost never copy-pastes onto the next. At the same time, speed matters because ML techniques have a shelf life. An innovation that delivers gains today may be superseded within a couple of model development cycles. The faster proven techniques reach the full model fleet, the more value they deliver before the next generation of techniques arrive.
So how do we propagate a technique to a diverse fleet of models as quickly as possible, while minimizing engineering bandwidth and compute?
Propagating a growing inventory of proven techniques across a large fleet of diverse models creates a combinatorially expanding space of technique-model combinations to explore, where both engineering bandwidth and compute become significant constraints.
The standard way to adapt a technique is to have engineers who are domain experts manually execute adaptation. An engineer studies how a technique interacts with a specific model’s architecture, generalizes it, and tunes it to that model’s constraints. This produces the strongest results, and is still the process for our highest-revenue models, but it does not scale to all of the models in the space. Each adaptation is effort-intensive and demands deep familiarity with both the technique and the target model. So throughput is capped by engineering bandwidth while the opportunity for adaptation grows much faster.
The second constraint on propagation is compute. Every adaptation attempt requires a training run, and with a finite compute budget the question is not only which techniques to try but how to navigate the combinatorial space of configurations without spending the budget on attempts (most of which will produce suboptimal gains).
Both constraints point to the same requirement: To propagate techniques to a large number of diverse models quickly, propagation has to become standardized, automated, and able to make judgment calls.
We started by standardizing model design and automating propagation: a template-driven approach organizes models into a small set of reusable and composable building blocks, so a technique validated once at the template level propagates automatically to every model built from that template. Over the last 2 years, through standardization and automation, we have shipped 5,000+ technique-model combinations across our large, diverse fleet of models while reducing per-model engineering time by 92%. You can read more about it in our post on Standard Model Templates (SMT) .
Standardization and automation handle the deterministic steps, but successful and efficient propagation still requires making judgment calls, such as which technique is worth propagating to which model, how wide a hyperparameter search should be, why a model is not showing gains, or how to detect and resolve failures. Those decisions still required significant human bandwidth, which prevented us from scaling further and from reaching a state where all techniques reached all applicable models within a few weeks of being proven.
That is where AI agents come in.
Scaling Technique Propagation With AI Agents
Our technique propagation system is built around three capabilities:
Discovery: Keep a live view of which techniques have been proven and estimate how well each might transfer to generate a prioritized set of technique-model candidates, ranked by expected impact.
Generalization: Generalize a proven technique so it works on models it was not originally designed for. For models built from shared templates, a technique is generalized once per template and reaches every model built from it. Beyond what templates cover, each model needs its own adaptation, which is where the deepest technical reasoning happens.
Adoption: Turn validated adaptations into shipped models, many at a time: launching experiments, managing asynchronous multi-day training runs, monitoring convergence, diagnosing and recovering from failures, and assembling proposals that can land.
Today, AI agents supercharge all the above capabilities by taking on the judgment-heavy, non-deterministic decision making and bringing deep intelligence to the system. They combine LLM-based reasoning with specialized ML tooling to handle the work of deciding which adaptations to try, diagnosing why something failed and automatically mitigating it, navigating hyperparameter search spaces efficiently, and managing capacity across hundreds of concurrent jobs. The agents share a common knowledge base, an accumulated record of prior training runs, model architectures, and technique outcomes, so they build on what has already been learned rather than starting from scratch.
Discovery answers two questions: What proven techniques exist, and how applicable are they to each model in the model fleet? Until recently, domain experts answered both, working through the assessments one model at a time. They tracked research and results across the organization, spotted the techniques worth propagating, and relied on their domain knowledge to judge applicability.
Agents aren’t constrained by the same knowledge or engineering bandwidth limitations as domain experts. Instead of relying on what any one expert knows, they draw on the shared knowledge base, and they make assessments across the entire model fleet at once rather than one model at a time. They systematically identify which proven techniques exist, estimate their transferability across the model fleet, and produce prioritized technique lists ranked by expected impact. Engineers now only have to review the agents’ assessments and make a decision from there.
How Agents Make Generalization More Computationally Efficient
Generalization is the process of re-implementing a technique proven on one model so it applies to models with different architectures, feature representations, training data, and serving constraints. The core idea usually transfers, but the implementation almost never does. The customization needed depends on the technique and the target model, ranging from retuning hyperparameters, to adapting features to what the target model already consumes or resizing the architecture to fit its inference budget, to redesigning the technique outright. Deciding what a given model needs, and how far to go, is a judgment call.
Depending on whether a model is built from standardized templates or customized individually, generalization comes in two flavors, and agents assist with both:
Template generalization: For models built from standardized templates , agents assist with building, evaluating, and improving the templates themselves. They propose template optimization hypotheses from proven techniques, launch experiments across representative model sets, evaluate results, diagnose failures, and iterate through multiple rounds, helping manage the multi-day loops that template development requires. Validating a technique once at the template level costs a fraction of the compute of adapting each model separately.
Bespoke generalization: For our highest-revenue models, where we still customize a technique for each model, agents progressively learn while working alongside ML engineers during the initial application of a technique to a small number of models. As the engineer makes adaptation decisions such as adjusting configurations, tuning hyperparameters, or resolving failures, the agent tracks every change and outcome, encoding what works, what doesn’t, and which decisions are model-specific versus generalizable. With that grounding, the agent proposes and tests adaptations for the remaining models, with the engineer reviewing and steering rather than driving each one.
Drawing on millions of prior training runs and each target model’s own properties, agents start from the configurations most likely to work, use the results of each round of training runs to narrow the hyperparameter search space for the next, and decide which experiments are worth continuing and which to stop early. As a result, when measured across select technique-model combinations, we have seen a median 4x reduction in training compute.
How Agents Accelerate Adoption
At the adoption stage, agents ship generalized techniques model-by-model. They launch experiments, manage asynchronous multi-day training runs, fine-tune each model, monitor convergence, recover from failures, and assemble validated proposals.
Fine-tuning happens differently depending on how the technique arrived. For techniques applied individually, agents help find the best configuration of technique bundles for that particular model. For models built from a shared template, an auto-research loop searches on top of the template, tuning each model for the gains the shared configuration alone does not deliver.
When failures occur, whether from an infrastructure issue, training divergence, or a data pipeline error, integrated debugging and recovery diagnose the root cause and apply targeted fixes or retries.
Running hundreds of these workflows at once raises the question of which jobs get priority, and which runs are worth continuing or stopping early. Getting the most out of a fixed compute budget is a global optimization problem, and it is where we are currently investing. Agents monitor concurrent training jobs, rank them by revenue importance, technique potential, and deadline proximity, and reclaim capacity from stalled or low-value runs, so GPUs stay pointed at the highest-value work rather than sitting idle or queued behind lower-priority experiments.
In the past, an engineer had to watch each run, triage every failure by hand, and restart it, which capped each of them at a handful of models. Now they can keep oversight at the decision points, approving plans and setting priorities at a 3x greater productivity (each engineer can now handle 3x more technique-model adaptations), compressing the time from a proven technique to a shipped model.
How We Evaluate Agent-Driven Propagation
Techniques propagated by agents are evaluated the same way as those prop

[truncated]
