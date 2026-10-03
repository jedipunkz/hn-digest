---
source: "https://epoch.ai/publications/estimating-the-agent-population"
hn_url: "https://news.ycombinator.com/item?id=49944481"
title: "How many AI agents could run on the AI chips shipped through 2027?"
article_title: "How many AI agents could we run? | Epoch AI"
image: "https://epoch.ai/assets/images/posts/2026/estimating-the-agent-population/t-estimating-the-agent-population.jpg"
author: "iphonecorridor"
captured_at: "2026-10-03T15:03:51Z"
capture_tool: "hn-digest"
hn_id: 49944481
score: 1
comments: 0
posted_at: "2026-10-03T14:17:36Z"
tags:
  - hacker-news
---

# How many AI agents could run on the AI chips shipped through 2027?

- HN: [49944481](https://news.ycombinator.com/item?id=49944481)
- Source: [epoch.ai](https://epoch.ai/publications/estimating-the-agent-population)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T14:17:36Z

## Translation

Title: How many AI agents could run on the AI chips shipped through 2027?
Article title: How many AI agents could we run? | Epoch AI
Description: Memory shipped through 2027 could run 33–171 million concurrent frontier-model agents, or billions using efficient open models. Epoch AI estimates inference capacity from HBM supply, serving benchmarks, and agent-hour costs.

Article text:
How many AI agents could we run? | Epoch AI
Latest Our work Featured Trends in AI
Our work Featured Trends in AI Data on AI Capabilities & benchmarking Publications Papers & reports Data Insights Newsletter Podcast All publications Data explorers AI capabilities AI models AI data centers AI chip owners AI companies Polling on AI usage All data explorers Our benchmarks Epoch Capabilities Index MirrorCode FrontierMath: Open Problems Earthborne Rangers
Navigate by topic AI progress Scaling Software progress Open models Capabilities & benchmarks Math Industry Leading companies Finances Geopolitics Infrastructure Chips Data centers Energy Impacts Adoption & use Economic impact Future of AI
About Who we are About us Our team Transparency Engage with us Careers Press Work with our experts Donate
How many AI agents could run on the AI chips shipped through 2027?
How many AI agents could run on the AI chips shipped through 2027?
Tens to hundreds of millions running today's most capable models, or billions running cheaper ones.
Code and data
Cite By Jason Li Contents
Potential agent capacity and spending
How we estimate agent capacity
Estimating agent sessions per GPU today
Estimating future capacity from high-bandwidth memory
Conclusion: what the totals imply for AI demand
Appendix A. Additional benchmark and trace data
Appendix B. Capacity calculations and sensitivities
Appendix C. HBM specifications
Potential agent capacity and spending
How we estimate agent capacity
Estimating agent sessions per GPU today
Estimating future capacity from high-bandwidth memory
Conclusion: what the totals imply for AI demand
Appendix A. Additional benchmark and trace data
Appendix B. Capacity calculations and sensitivities
Appendix C. HBM specifications
AI companies are spending hundreds of billions of dollars a year on chips and data centers, on the premise that those chips will run AI agents to do work that people do today.
How many agents could this hardware buildout actually support?
AI chips shipped through 2027 could run tens to hundreds of millions of concurrent frontier-model agents. Running nonstop, these agents would supply as many weekly working hours as about 140–720 million full-time employees.
More efficient models could potentially support billions of agents on the same hardware. Applying DeepSeek V4 Pro serving benchmarks to the projected hardware supply yields approximately 1.9 billion concurrent agents supplying as many weekly working hours as 8 billion people each working 40 hours.
Even modest use of this capacity would require a massive increase in global demand for AI. Using 20% of our central capacity estimate would imply $2.6–5.3 trillion a year in API-equivalent spending, against roughly $1 trillion in developer revenue by end-2027 at fivefold annual growth.
Hourly agent spending varies substantially across models and harnesses. In our analysis of agent traces, Codex workloads averaged roughly $16–18 per hour of continuous agent activity, compared with $24–50 for Claude Code workloads.
Potential agent capacity and spending
Anthropic’s Dario Amodei has described a future “ country of geniuses in a datacenter ”, but how many AI agents could future data centers actually support? That scale matters for AI’s potential impact on the economy and labor force.
We find that hardware using high-bandwidth memory (HBM) shipped during 2025–27 could eventually support tens to hundreds of millions of concurrent frontier-model agents, assuming full deployment and allocation to these workloads. HBM shipped during 2025–26 could support 16–56 million concurrent agents once deployed. Including shipments through 2027 raises that estimate to about 30–170 million. 1
But unlike humans, an AI agent can work all 168 hours each week, 4.2 times the 40-hour workweek for a full-time employee. Therefore, these agents could work as many weekly hours as about 67–240 million people from hardware shipments through 2026, and about 140–720 million from shipments through 2027. For scale, the United States has a population of 342 million and an estimated 100 million knowledge workers. These comparisons count working hours alone. Agents can also produce output much faster than humans, though the quality of that output varies.
Even if we use only 20% of the capacity from memory shipped through 2027, the implied spending at API prices would be $2.6–5.3 trillion a year once the hardware is deployed. 2
For comparison, if model developers’ revenues keep growing fivefold each year, their combined annualized revenue would reach roughly $1 trillion by the end of 2027.
Demand could fall behind this potential supply, creating an overabundance of capacity.
The key uncertainty is whether sustained, rapid growth in demand for AI services will justify the investment.
How we estimate agent capacity
We estimate potential concurrent agents (\(A\)): how many agents could run at once on the memory shipped during 2025–27, assuming full deployment and allocation to the modeled workload.
We use these capacity estimates to derive working-hour and API-equivalent spending figures under the stated operating-time, pricing, allocation, and utilization assumptions.
Our estimate of potential concurrent agents is built from two terms:
Effective hardware supply (\(E\)): measured in GB300-equivalent inference units, analogous to FLOP-based H100 equivalents.
For these workloads, memory capacity constrains concurrency and memory bandwidth constrains streaming speed.
We count high-bandwidth memory (HBM) shipped from 2025 onward, including HBM3E and newer generations, in units of 288 GB, matching a GB300 GPU.
We then adjust for the concurrency that newer hardware can support:
where \(H_3\) and \(H_4\) are cumulative HBM3E and HBM4/4E supply (GB).
\(u\) is the ratio of agent sessions per GB on HBM4/4E systems to agent sessions per GB on HBM3E systems.
Serving capacity (\(c\)): concurrent active agent sessions per GB300 equivalent.
for closed models.
\(S\) is API spending per active agent-hour ($/agent-hour), using durations adjusted to remove identified human waits and cap other idle gaps.
\(G\) is GPU rental cost ($/GB300-hour).
\(K\) is API-equivalent revenue divided by reference serving cost.
For open models, we use benchmarked concurrency.
Main assumptions: \(S = \$30/\text{agent-hour}\), \(G = \$5/\text{GB300-hour}\), \({K = 5\text{–}10\times}\), and \({u = 2\times}\). We also test \({u = 1\times}\) and \(4\times\).
Open-model benchmarks use P90 streaming-speed targets of 50 and 100 output tokens per second per user, with 200 as a sensitivity. We hold current model and workload requirements fixed.
Estimating agent sessions per GPU today
We use “agent” as shorthand for an agentic workload running within a harness such as Codex or Claude Code. An agent session includes model calls and tool use, rather than continuous token generation.
We draw on two sources: SemiAnalysis’s AgentX serving benchmark for open models, and TraceLab , a public dataset of logged agent sessions, for closed models.
AgentX counts a main agent and its subagents as one session tree; TraceLab’s accounting groups do not always capture that complete tree. We use “agent” and “agent session” interchangeably when discussing capacity. A continuous agent session includes the time spent waiting for tool calls. We estimate the continuous working time by removing time spent waiting for human input and also capping unidentified idle time. One hour of this adjusted activity counts as one agent-hour.
For open models, serving benchmarks directly measure how many concurrent agent sessions the hardware supports at a given output speed. For closed models, we estimate concurrency from hourly spending and serving-cost assumptions.
Open models: serving benchmarks measure concurrency directly
For open models, we use the benchmark data from SemiAnalysis’s InferenceX AgentX . The AgentX benchmark uses a dataset of Claude Code agent session traces collected by SemiAnalysis. The benchmark replay uses synthetic text while preserving request lengths, shared context, and the timing and structure of model calls. Further details in the AgentX methodology .
Concurrency for AgentX is defined by the number of agent sessions launched. We divide the concurrency by the total number of GPUs to derive a metric of concurrent agent sessions per GPU. For prefill-decode disaggregated configurations, we count the number of combined GPUs. Figures 2–3 show the published configurations and the agent sessions per GPU.
Speed targets: 50 and 100 tokens per second per user
We use 50 and 100 TPS/user as round reference points for output speed in tokens per second. 200 TPS/user is also given in the appendix to test a more demanding scenario. For reference, both OpenAI and Anthropic tend to serve their frontier models around 50–70 TPS. In the benchmark data, P90 interactivity describes the output speed in TPS/user for the slower end of the distribution. This excludes the time to first token (TTFT) and does not measure end-to-end latency (see Appendix A).
The figures use the AgentX snapshot dated in Table 2 and cover seven models. We excluded any preview data not run directly through the InferenceX public repo.
Concurrency per GPU at the main speed targets
Closed frontier models: concurrency inferred from API spending
Since closed frontier models lack the architecture details needed for transparent hardware benchmarks, we instead look at the API costs of a continuously running agent. We analyzed TraceLab’s dataset of Codex and Claude Code agent sessions. We adjusted for human delays (agent waiting on human response) and divided spending by agent working time to normalize to an hourly rate. Using the API prices listed in Table A3, we calculated the API-equivalent spending per agent-hour.
We then estimate how much of that spending covers serving costs, using an assumed markup ratio of API revenue to serving cost. Comparing the resulting cost per agent-hour with the rental cost of a GB300 GPU gives an estimate of how many concurrent agents each GPU could support.
$30 per agent-hour sits between Codex and Claude Code costs
Hourly spending varies across models and harnesses. Under Figure 4’s retained-cache 4 and five-minute gap-cap assumptions, TraceLab’s pooled rates are $18.2/hour for GPT-5.5 , $15.5 for GPT-5.6 Sol, $24.3 for Opus 4.8, and $50.2 for Fable 5. We choose $30/hour as a round reference point which sits slightly towards the higher end. Future models may go up in price as with the jumps to Fable for Anthropic and GPT-6 Astra for OpenAI or they may go down due to competition and efficiency improvements. Figure 8 goes through a range of prices from $10 to $100 per agent-hour.
Let \(S\) be API spending per agent-hour and \(K\) the ratio of API-equivalent revenue to serving cost at our reference GPU rental price. \({K = 10\times}\) means $10 of API billing for every $1 of that cost.
The implied serving cost per agent-hour is \(S/K\). Dividing the GPU-hour rental price, \(G\), by this cost gives agent sessions per GB300 equivalent:
Table 1 uses \(S = \$30\) per agent-hour and \(G\) = $5.00/GB300-hour , from SemiAnalysis’s Rent – 3 Year Commit tier. At \({K = 10\times}\), the implied serving cost is $3 per agent-hour. A $5 GPU-hour supports 1.67 concurrent agent sessions.
Table 1. Closed-model assumptions at $30 per agent-hour and $5 per GB300-hour. Appendix B tests higher revenue/cost multiples.
Our main case uses \({K = 5\text{–}10\times}\), giving 0.833–1.667 agent sessions per GB300 at $30 per agent-hour. Table 2 shows the ratio \(K\) at around 4–10× for open models on AgentX at 50 TPS/user. 5
Open-model benchmarks imply revenue/cost ratios of 2–10×
We compare the assumed multiples with what each open-model benchmark would earn at its model’s API rates, using theoretical cache-hit rates. At each P90 speed target, we select the GPU with the lowest three-year rental cost per concurrent agent-hour.
Each speed column reports \

[truncated]

## Original Extract

Memory shipped through 2027 could run 33–171 million concurrent frontier-model agents, or billions using efficient open models. Epoch AI estimates inference capacity from HBM supply, serving benchmarks, and agent-hour costs.

How many AI agents could we run? | Epoch AI
Latest Our work Featured Trends in AI
Our work Featured Trends in AI Data on AI Capabilities & benchmarking Publications Papers & reports Data Insights Newsletter Podcast All publications Data explorers AI capabilities AI models AI data centers AI chip owners AI companies Polling on AI usage All data explorers Our benchmarks Epoch Capabilities Index MirrorCode FrontierMath: Open Problems Earthborne Rangers
Navigate by topic AI progress Scaling Software progress Open models Capabilities & benchmarks Math Industry Leading companies Finances Geopolitics Infrastructure Chips Data centers Energy Impacts Adoption & use Economic impact Future of AI
About Who we are About us Our team Transparency Engage with us Careers Press Work with our experts Donate
How many AI agents could run on the AI chips shipped through 2027?
How many AI agents could run on the AI chips shipped through 2027?
Tens to hundreds of millions running today's most capable models, or billions running cheaper ones.
Code and data
Cite By Jason Li Contents
Potential agent capacity and spending
How we estimate agent capacity
Estimating agent sessions per GPU today
Estimating future capacity from high-bandwidth memory
Conclusion: what the totals imply for AI demand
Appendix A. Additional benchmark and trace data
Appendix B. Capacity calculations and sensitivities
Appendix C. HBM specifications
Potential agent capacity and spending
How we estimate agent capacity
Estimating agent sessions per GPU today
Estimating future capacity from high-bandwidth memory
Conclusion: what the totals imply for AI demand
Appendix A. Additional benchmark and trace data
Appendix B. Capacity calculations and sensitivities
Appendix C. HBM specifications
AI companies are spending hundreds of billions of dollars a year on chips and data centers, on the premise that those chips will run AI agents to do work that people do today.
How many agents could this hardware buildout actually support?
AI chips shipped through 2027 could run tens to hundreds of millions of concurrent frontier-model agents. Running nonstop, these agents would supply as many weekly working hours as about 140–720 million full-time employees.
More efficient models could potentially support billions of agents on the same hardware. Applying DeepSeek V4 Pro serving benchmarks to the projected hardware supply yields approximately 1.9 billion concurrent agents supplying as many weekly working hours as 8 billion people each working 40 hours.
Even modest use of this capacity would require a massive increase in global demand for AI. Using 20% of our central capacity estimate would imply $2.6–5.3 trillion a year in API-equivalent spending, against roughly $1 trillion in developer revenue by end-2027 at fivefold annual growth.
Hourly agent spending varies substantially across models and harnesses. In our analysis of agent traces, Codex workloads averaged roughly $16–18 per hour of continuous agent activity, compared with $24–50 for Claude Code workloads.
Potential agent capacity and spending
Anthropic’s Dario Amodei has described a future “ country of geniuses in a datacenter ”, but how many AI agents could future data centers actually support? That scale matters for AI’s potential impact on the economy and labor force.
We find that hardware using high-bandwidth memory (HBM) shipped during 2025–27 could eventually support tens to hundreds of millions of concurrent frontier-model agents, assuming full deployment and allocation to these workloads. HBM shipped during 2025–26 could support 16–56 million concurrent agents once deployed. Including shipments through 2027 raises that estimate to about 30–170 million. 1
But unlike humans, an AI agent can work all 168 hours each week, 4.2 times the 40-hour workweek for a full-time employee. Therefore, these agents could work as many weekly hours as about 67–240 million people from hardware shipments through 2026, and about 140–720 million from shipments through 2027. For scale, the United States has a population of 342 million and an estimated 100 million knowledge workers. These comparisons count working hours alone. Agents can also produce output much faster than humans, though the quality of that output varies.
Even if we use only 20% of the capacity from memory shipped through 2027, the implied spending at API prices would be $2.6–5.3 trillion a year once the hardware is deployed. 2
For comparison, if model developers’ revenues keep growing fivefold each year, their combined annualized revenue would reach roughly $1 trillion by the end of 2027.
Demand could fall behind this potential supply, creating an overabundance of capacity.
The key uncertainty is whether sustained, rapid growth in demand for AI services will justify the investment.
How we estimate agent capacity
We estimate potential concurrent agents (\(A\)): how many agents could run at once on the memory shipped during 2025–27, assuming full deployment and allocation to the modeled workload.
We use these capacity estimates to derive working-hour and API-equivalent spending figures under the stated operating-time, pricing, allocation, and utilization assumptions.
Our estimate of potential concurrent agents is built from two terms:
Effective hardware supply (\(E\)): measured in GB300-equivalent inference units, analogous to FLOP-based H100 equivalents.
For these workloads, memory capacity constrains concurrency and memory bandwidth constrains streaming speed.
We count high-bandwidth memory (HBM) shipped from 2025 onward, including HBM3E and newer generations, in units of 288 GB, matching a GB300 GPU.
We then adjust for the concurrency that newer hardware can support:
where \(H_3\) and \(H_4\) are cumulative HBM3E and HBM4/4E supply (GB).
\(u\) is the ratio of agent sessions per GB on HBM4/4E systems to agent sessions per GB on HBM3E systems.
Serving capacity (\(c\)): concurrent active agent sessions per GB300 equivalent.
for closed models.
\(S\) is API spending per active agent-hour ($/agent-hour), using durations adjusted to remove identified human waits and cap other idle gaps.
\(G\) is GPU rental cost ($/GB300-hour).
\(K\) is API-equivalent revenue divided by reference serving cost.
For open models, we use benchmarked concurrency.
Main assumptions: \(S = \$30/\text{agent-hour}\), \(G = \$5/\text{GB300-hour}\), \({K = 5\text{–}10\times}\), and \({u = 2\times}\). We also test \({u = 1\times}\) and \(4\times\).
Open-model benchmarks use P90 streaming-speed targets of 50 and 100 output tokens per second per user, with 200 as a sensitivity. We hold current model and workload requirements fixed.
Estimating agent sessions per GPU today
We use “agent” as shorthand for an agentic workload running within a harness such as Codex or Claude Code. An agent session includes model calls and tool use, rather than continuous token generation.
We draw on two sources: SemiAnalysis’s AgentX serving benchmark for open models, and TraceLab , a public dataset of logged agent sessions, for closed models.
AgentX counts a main agent and its subagents as one session tree; TraceLab’s accounting groups do not always capture that complete tree. We use “agent” and “agent session” interchangeably when discussing capacity. A continuous agent session includes the time spent waiting for tool calls. We estimate the continuous working time by removing time spent waiting for human input and also capping unidentified idle time. One hour of this adjusted activity counts as one agent-hour.
For open models, serving benchmarks directly measure how many concurrent agent sessions the hardware supports at a given output speed. For closed models, we estimate concurrency from hourly spending and serving-cost assumptions.
Open models: serving benchmarks measure concurrency directly
For open models, we use the benchmark data from SemiAnalysis’s InferenceX AgentX . The AgentX benchmark uses a dataset of Claude Code agent session traces collected by SemiAnalysis. The benchmark replay uses synthetic text while preserving request lengths, shared context, and the timing and structure of model calls. Further details in the AgentX methodology .
Concurrency for AgentX is defined by the number of agent sessions launched. We divide the concurrency by the total number of GPUs to derive a metric of concurrent agent sessions per GPU. For prefill-decode disaggregated configurations, we count the number of combined GPUs. Figures 2–3 show the published configurations and the agent sessions per GPU.
Speed targets: 50 and 100 tokens per second per user
We use 50 and 100 TPS/user as round reference points for output speed in tokens per second. 200 TPS/user is also given in the appendix to test a more demanding scenario. For reference, both OpenAI and Anthropic tend to serve their frontier models around 50–70 TPS. In the benchmark data, P90 interactivity describes the output speed in TPS/user for the slower end of the distribution. This excludes the time to first token (TTFT) and does not measure end-to-end latency (see Appendix A).
The figures use the AgentX snapshot dated in Table 2 and cover seven models. We excluded any preview data not run directly through the InferenceX public repo.
Concurrency per GPU at the main speed targets
Closed frontier models: concurrency inferred from API spending
Since closed frontier models lack the architecture details needed for transparent hardware benchmarks, we instead look at the API costs of a continuously running agent. We analyzed TraceLab’s dataset of Codex and Claude Code agent sessions. We adjusted for human delays (agent waiting on human response) and divided spending by agent working time to normalize to an hourly rate. Using the API prices listed in Table A3, we calculated the API-equivalent spending per agent-hour.
We then estimate how much of that spending covers serving costs, using an assumed markup ratio of API revenue to serving cost. Comparing the resulting cost per agent-hour with the rental cost of a GB300 GPU gives an estimate of how many concurrent agents each GPU could support.
$30 per agent-hour sits between Codex and Claude Code costs
Hourly spending varies across models and harnesses. Under Figure 4’s retained-cache 4 and five-minute gap-cap assumptions, TraceLab’s pooled rates are $18.2/hour for GPT-5.5 , $15.5 for GPT-5.6 Sol, $24.3 for Opus 4.8, and $50.2 for Fable 5. We choose $30/hour as a round reference point which sits slightly towards the higher end. Future models may go up in price as with the jumps to Fable for Anthropic and GPT-6 Astra for OpenAI or they may go down due to competition and efficiency improvements. Figure 8 goes through a range of prices from $10 to $100 per agent-hour.
Let \(S\) be API spending per agent-hour and \(K\) the ratio of API-equivalent revenue to serving cost at our reference GPU rental price. \({K = 10\times}\) means $10 of API billing for every $1 of that cost.
The implied serving cost per agent-hour is \(S/K\). Dividing the GPU-hour rental price, \(G\), by this cost gives agent sessions per GB300 equivalent:
Table 1 uses \(S = \$30\) per agent-hour and \(G\) = $5.00/GB300-hour , from SemiAnalysis’s Rent – 3 Year Commit tier. At \({K = 10\times}\), the implied serving cost is $3 per agent-hour. A $5 GPU-hour supports 1.67 concurrent agent sessions.
Table 1. Closed-model assumptions at $30 per agent-hour and $5 per GB300-hour. Appendix B tests higher revenue/cost multiples.
Our main case uses \({K = 5\text{–}10\times}\), giving 0.833–1.667 agent sessions per GB300 at $30 per agent-hour. Table 2 shows the ratio \(K\) at around 4–10× for open models on AgentX at 50 TPS/user. 5
Open-model benchmarks imply revenue/cost ratios of 2–10×
We compare the assumed multiples with what each open-model benchmark would earn at its model’s API rates, using theoretical cache-hit rates. At each P90 speed target, we select the GPU with the lowest three-year rental cost per concurrent agent-hour.
Each speed column reports \

[truncated]
