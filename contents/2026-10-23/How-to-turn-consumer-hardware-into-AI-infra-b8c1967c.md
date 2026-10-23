---
source: "https://neurogrid.cc/en"
hn_url: "https://news.ycombinator.com/item?id=50205001"
title: "How to turn consumer hardware into AI infra"
article_title: "Neurogrid Community Cloud — The Inference Marketplace Built by the People, for the People"
image: "https://neurogrid.cc/_next/static/immutable/media/neurogrid.1u4xndkygjv1u.png"
author: "cheyodev"
captured_at: "2026-10-23T07:03:04Z"
capture_tool: "hn-digest"
hn_id: 50205001
score: 3
comments: 0
posted_at: "2026-10-23T06:04:30Z"
tags:
  - hacker-news
---

# How to turn consumer hardware into AI infra

- HN: [50205001](https://news.ycombinator.com/item?id=50205001)
- Source: [neurogrid.cc](https://neurogrid.cc/en)
- Score: 3
- Comments: 0
- Posted: 2026-10-23T06:04:30Z

## Translation

Title: How to turn consumer hardware into AI infra
Article title: Neurogrid Community Cloud — The Inference Marketplace Built by the People, for the People
Description: Run language models on underutilized GPUs from people around the world. Affordable, accessible inference powered by community-owned hardware.

Article text:
Neurogrid Community Cloud — The Inference Marketplace Built by the People, for the People 🎉 Early suppliers: host a GPU now and get an extra $5 USD on every payout for your first 3 months. Become a host → Neurogrid Platform Use Cases Earnings NeuroChat New nrgrd-tui Manifest Join Open Console Neurogrid Community Cloud The Inference Marketplace Built by the People, for the People.
Run language models on underutilized GPUs from people around the world. Affordable, accessible inference powered by community-owned hardware.
Sign up with Google or GitHub in seconds.
Neurogrid connects people with underutilized GPUs to people who need affordable AI inference.
GPU owners contribute idle hardware to the network and earn from the workloads running on their machines. AI users reach distributed compute without owning expensive infrastructure.
The marketplace is the economic connection between the two.
02 — Use cases Works with the tools you already use.
Point any app that accepts an endpoint and an API key at Neurogrid. Placeholder examples below.
An IDE assistant sends completions and chat through Neurogrid instead of a single hosted model.
+------------------------------+
| editor.ts [saved] |
+------------------------------+
| fn suggest() { |
| // ng:: deepseek-r1 |
| return completion; |
| } |
+------------------------------+ ng config set model deepseek-r1
ng config set endpoint https://api.neurogrid.dev/v1 New Chat app — NeuroChat NeuroChat
Neurogrid's own chat app. Talk to open-source models running on the network right in your browser: no account, no setup, nothing to install.
Your chats are stored only in your browser.
Terminal client — nrgrd-tui nrgrd-tui
Neurogrid's own terminal client. Browse the models available on the network, see which providers are online, and run inference straight from the shell.
███╗ ██╗██████╗ ██████╗ ██████╗ ██████╗
████╗ ██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗
██╔██╗ ██║██████╔╝██║ ███╗██████╔╝██║ ██║
██║╚██╗██║██╔══██╗██║ ██║██╔══██╗██║ ██║
██║ ╚████║██║ ██║╚██████╔╝██║ ██║██████╔╝
╚═╝ ╚═══╝╚═╝ ╚═╝ ╚═════╝ ╚═╝ ╚═╝╚═════╝ nrgrd # interactive TUI
nrgrd "explain this repo"
nrgrd --model qwen3-coder "fix the failing tests" Chat gateway — OpenClaw OpenClaw
OpenClaw is a self-hosted gateway that connects Discord, Slack, Telegram and other chat apps to AI coding agents. Point its model config at a Neurogrid deployment to run those agents on community GPUs instead of a centralized provider.
+------------------------------+
| openclaw [] |
+------------------------------+
| > run qwen3-32b |
| :: routed -> supply/0142 |
| ........... |
| #### streaming |
| |
+------------------------------+ agents: {
defaults: {
model: { primary: "openai/qwen3-32b" },
baseUrl: "https://api.neurogrid.dev/v1"
}
} Self-hosted UI — Open WebUI Open WebUI
A self-hosted chat interface, pointed at a Neurogrid deployment through its OpenAI-compatible connection settings. Each conversation runs on whichever community GPU the marketplace matched.
+---------+--------------------+
| threads | open webui |
+---------+--------------------+
| . chat | user: ... |
| . chat | ng: #### |
| . chat | ------------ |
| + | supply/0087 |
+---------+--------------------+ OPENAI_API_BASE_URL=https://api.neurogrid.dev/v1
OPENAI_API_KEY=ng-•••••••••••••••• Any API-key app Endpoint + API key
Anything that speaks an OpenAI-compatible API works: agents, IDE plugins, scripts, internal tools.
app neurogrid
+------+ ...> +--------+
| key | | match |
+------+ <... +--------+
| |
...+ +...
o o
gpu gpu curl https://api.neurogrid.dev/v1/chat/completions \
-H "Authorization: Bearer $NG_KEY" \
-d '{"model":"llama-3.3-70b"}' Batch workloads Pipelines & agents
Queue long-running batch jobs across many providers at once and pay only for the compute that runs.
queue providers
[####] --> o o o o o
[###.] --> o o o o o
[##..] --> o o o o o
[#...] . . . . .
24 running ng jobs submit batch.jsonl \
--model mistral-small \
--concurrency 24 Coding agent IDE plugin
An IDE assistant sends completions and chat through Neurogrid instead of a single hosted model.
+------------------------------+
| editor.ts [saved] |
+------------------------------+
| fn suggest() { |
| // ng:: deepseek-r1 |
| return completion; |
| } |
+------------------------------+ ng config set model deepseek-r1
ng config set endpoint https://api.neurogrid.dev/v1 New Chat app — NeuroChat NeuroChat
Neurogrid's own chat app. Talk to open-source models running on the network right in your browser: no account, no setup, nothing to install.
Your chats are stored only in your browser.
Models on the network The open models currently served by community hardware, by lab.
OpenAI Mistral AI Meta Qwen DeepSeek Gemma GLM Lower cost More accessible compute Distributed infrastructure Community-owned capacity Better hardware utilization A new income stream for GPU owners 03 — Distribution Compute shouldn’t live in a handful of buildings.
Traditional AI infrastructure concentrates enormous amounts of computing power inside massive data centers. That infrastructure is expensive to build, expensive to operate, and often leaves enormous amounts of consumer hardware sitting idle.
Neurogrid takes a different approach: use the hardware that already exists.
Connect your underutilized GPU to Neurogrid.
The marketplace routes compatible workloads to available hardware.
AI workloads run on distributed community infrastructure.
GPU providers earn from the compute their hardware contributes.
Turn idle hardware into income.
Most GPUs sit unused for large parts of the day. Contribute that unused capacity and earn from workloads that would otherwise require centralized infrastructure.
For AI users Inference without the infrastructure bill.
Access the compute you need without owning the infrastructure required to produce it. Route your workloads through the marketplace and pay for what runs.
04 — Straight answers Before you connect a machine.
Handing your hardware to a network is a real decision. These are the questions people actually ask, answered without spin.
Isolated inference workloads, nothing else. Each job runs contained, without access to the rest of your system.
No. Intercepting, logging or retaining the contents of inference passing through your hardware is grounds for permanent suspension. It is written into the Manifest .
It depends on your hardware and on real demand. We would rather show you nothing than publish an earnings figure we cannot stand behind.
Bring your GPU. Bring your workloads. Help build a more accessible compute network.
Sign up with Google or GitHub in seconds.

## Original Extract

Run language models on underutilized GPUs from people around the world. Affordable, accessible inference powered by community-owned hardware.

Neurogrid Community Cloud — The Inference Marketplace Built by the People, for the People 🎉 Early suppliers: host a GPU now and get an extra $5 USD on every payout for your first 3 months. Become a host → Neurogrid Platform Use Cases Earnings NeuroChat New nrgrd-tui Manifest Join Open Console Neurogrid Community Cloud The Inference Marketplace Built by the People, for the People.
Run language models on underutilized GPUs from people around the world. Affordable, accessible inference powered by community-owned hardware.
Sign up with Google or GitHub in seconds.
Neurogrid connects people with underutilized GPUs to people who need affordable AI inference.
GPU owners contribute idle hardware to the network and earn from the workloads running on their machines. AI users reach distributed compute without owning expensive infrastructure.
The marketplace is the economic connection between the two.
02 — Use cases Works with the tools you already use.
Point any app that accepts an endpoint and an API key at Neurogrid. Placeholder examples below.
An IDE assistant sends completions and chat through Neurogrid instead of a single hosted model.
+------------------------------+
| editor.ts [saved] |
+------------------------------+
| fn suggest() { |
| // ng:: deepseek-r1 |
| return completion; |
| } |
+------------------------------+ ng config set model deepseek-r1
ng config set endpoint https://api.neurogrid.dev/v1 New Chat app — NeuroChat NeuroChat
Neurogrid's own chat app. Talk to open-source models running on the network right in your browser: no account, no setup, nothing to install.
Your chats are stored only in your browser.
Terminal client — nrgrd-tui nrgrd-tui
Neurogrid's own terminal client. Browse the models available on the network, see which providers are online, and run inference straight from the shell.
███╗ ██╗██████╗ ██████╗ ██████╗ ██████╗
████╗ ██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗
██╔██╗ ██║██████╔╝██║ ███╗██████╔╝██║ ██║
██║╚██╗██║██╔══██╗██║ ██║██╔══██╗██║ ██║
██║ ╚████║██║ ██║╚██████╔╝██║ ██║██████╔╝
╚═╝ ╚═══╝╚═╝ ╚═╝ ╚═════╝ ╚═╝ ╚═╝╚═════╝ nrgrd # interactive TUI
nrgrd "explain this repo"
nrgrd --model qwen3-coder "fix the failing tests" Chat gateway — OpenClaw OpenClaw
OpenClaw is a self-hosted gateway that connects Discord, Slack, Telegram and other chat apps to AI coding agents. Point its model config at a Neurogrid deployment to run those agents on community GPUs instead of a centralized provider.
+------------------------------+
| openclaw [] |
+------------------------------+
| > run qwen3-32b |
| :: routed -> supply/0142 |
| ........... |
| #### streaming |
| |
+------------------------------+ agents: {
defaults: {
model: { primary: "openai/qwen3-32b" },
baseUrl: "https://api.neurogrid.dev/v1"
}
} Self-hosted UI — Open WebUI Open WebUI
A self-hosted chat interface, pointed at a Neurogrid deployment through its OpenAI-compatible connection settings. Each conversation runs on whichever community GPU the marketplace matched.
+---------+--------------------+
| threads | open webui |
+---------+--------------------+
| . chat | user: ... |
| . chat | ng: #### |
| . chat | ------------ |
| + | supply/0087 |
+---------+--------------------+ OPENAI_API_BASE_URL=https://api.neurogrid.dev/v1
OPENAI_API_KEY=ng-•••••••••••••••• Any API-key app Endpoint + API key
Anything that speaks an OpenAI-compatible API works: agents, IDE plugins, scripts, internal tools.
app neurogrid
+------+ ...> +--------+
| key | | match |
+------+ <... +--------+
| |
...+ +...
o o
gpu gpu curl https://api.neurogrid.dev/v1/chat/completions \
-H "Authorization: Bearer $NG_KEY" \
-d '{"model":"llama-3.3-70b"}' Batch workloads Pipelines & agents
Queue long-running batch jobs across many providers at once and pay only for the compute that runs.
queue providers
[####] --> o o o o o
[###.] --> o o o o o
[##..] --> o o o o o
[#...] . . . . .
24 running ng jobs submit batch.jsonl \
--model mistral-small \
--concurrency 24 Coding agent IDE plugin
An IDE assistant sends completions and chat through Neurogrid instead of a single hosted model.
+------------------------------+
| editor.ts [saved] |
+------------------------------+
| fn suggest() { |
| // ng:: deepseek-r1 |
| return completion; |
| } |
+------------------------------+ ng config set model deepseek-r1
ng config set endpoint https://api.neurogrid.dev/v1 New Chat app — NeuroChat NeuroChat
Neurogrid's own chat app. Talk to open-source models running on the network right in your browser: no account, no setup, nothing to install.
Your chats are stored only in your browser.
Models on the network The open models currently served by community hardware, by lab.
OpenAI Mistral AI Meta Qwen DeepSeek Gemma GLM Lower cost More accessible compute Distributed infrastructure Community-owned capacity Better hardware utilization A new income stream for GPU owners 03 — Distribution Compute shouldn’t live in a handful of buildings.
Traditional AI infrastructure concentrates enormous amounts of computing power inside massive data centers. That infrastructure is expensive to build, expensive to operate, and often leaves enormous amounts of consumer hardware sitting idle.
Neurogrid takes a different approach: use the hardware that already exists.
Connect your underutilized GPU to Neurogrid.
The marketplace routes compatible workloads to available hardware.
AI workloads run on distributed community infrastructure.
GPU providers earn from the compute their hardware contributes.
Turn idle hardware into income.
Most GPUs sit unused for large parts of the day. Contribute that unused capacity and earn from workloads that would otherwise require centralized infrastructure.
For AI users Inference without the infrastructure bill.
Access the compute you need without owning the infrastructure required to produce it. Route your workloads through the marketplace and pay for what runs.
04 — Straight answers Before you connect a machine.
Handing your hardware to a network is a real decision. These are the questions people actually ask, answered without spin.
Isolated inference workloads, nothing else. Each job runs contained, without access to the rest of your system.
No. Intercepting, logging or retaining the contents of inference passing through your hardware is grounds for permanent suspension. It is written into the Manifest .
It depends on your hardware and on real demand. We would rather show you nothing than publish an earnings figure we cannot stand behind.
Bring your GPU. Bring your workloads. Help build a more accessible compute network.
Sign up with Google or GitHub in seconds.
