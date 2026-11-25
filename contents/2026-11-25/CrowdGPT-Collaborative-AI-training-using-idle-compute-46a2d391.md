---
source: "https://www.crowdgpt.net/"
hn_url: "https://news.ycombinator.com/item?id=50642047"
title: "CrowdGPT – Collaborative AI training using idle compute"
article_title: "CrowdGPT"
image: ""
author: "Vxtzq"
captured_at: "2026-11-25T13:56:23Z"
capture_tool: "hn-digest"
hn_id: 50642047
score: 3
comments: 1
posted_at: "2026-11-25T13:29:40Z"
tags:
  - hacker-news
---

# CrowdGPT – Collaborative AI training using idle compute

- HN: [50642047](https://news.ycombinator.com/item?id=50642047)
- Source: [www.crowdgpt.net](https://www.crowdgpt.net/)
- Score: 3
- Comments: 1
- Posted: 2026-11-25T13:29:40Z

## Translation

Title: CrowdGPT – Collaborative AI training using idle compute
Article title: CrowdGPT

Article text:
View roadmap
×
C
CrowdGPT
home
install
about
status
leaderboard
submit
beta
contact
English
Français
Deutsch
Español
Türkçe
简体中文
source
C
CrowdGPT
home
install
about
status
leaderboard
submit
beta
contact
Source on GitHub
AI owned by everyone .
CrowdGPT is focusing on creating the world's first 0$, datacenterless and 100% open AI. Run the app on your computer, and the AI gets better.
Huggingface
Discord
—
Epoch
—
GPUs currently active
—
Tokens processed
—
Held-out loss
A new paradigm for training AI.
A single company training a frontier model uses thousands of A100s in a datacenter. Most people don't have that hardware. But millions of people have gaming PCs with RTX cards sitting idle. CrowdGPT uses that compute, and the model that gets trained belongs to everyone who helped build it.
Your GPU trains on 10MB data chunks streamed from the public CrowdGPT dataset on HuggingFace; a dataset anyone can inspect or contribute to. No personal files, browser history, or local data are ever accessed or transmitted.
Trains on consumer GPUs that are already powered on. CrowdGPT requires no new datacenters, no additional hardware manufacturing, which is a huge tradeoff of traditional AI.
The client is a single, human-readable Python file that anyone can audit and modify. Model weights are downloadable, 100% free and open on HuggingFace. View server logs
The client is a single open-source Python file which is human-readable. It downloads public training data online, trains locally and uploads only the model update. It cannot access your files or compromise your machine.
The server is protected too. Every update is validated for size, NaN values and norm bounds before it is cross-validated by peers, and repeat offenders are banned from the project.
Read the source before you run it.
Every contribution, even the smallest, makes the model better. Run a client, star the repo, or open an issue.
Star on GitHub
home / install
Install the client
Install everything with one command. Even if you don't have python/git installed.
~6 GB of free disk space. PyTorch alone is around 2.5 GB.
A discrete GPU is strongly recommended. NVIDIA (CUDA), AMD (ROCm / DirectML), Apple Silicon (MPS) and Intel Arc (XPU) are all supported.
A stable internet connection for streaming model weights and dataset chunks.
CPU-only mode works, but is extremely slow.
The installer handles everything else: GPU detection, the correct PyTorch wheel, Python and Git if missing, and an isolated virtual environment.
Copy this into your terminal. It downloads the installer, auto-detects your hardware, and starts the client.
Detect hardware. Installs the right dependencies for your Graphics Card.
Install dependencies. Python and Git are installed if missing, then PyTorch from the correct index and the project's requirements; everything is inside an isolated virtual environment so nothing touches your system Python.
Launch. Downloads the model weights from HuggingFace, streams a dataset chunk, and starts training. You can stop at any time.
Four steps turn a small slice of online, curated text into one shared model.
Diagram unavailable
How CrowdGPT coordinates a bunch of independent GPUs to train a single language model.
1. Assignment
When your device connects, the server gives it one 10MB slice of the public CrowdGPT dataset and the latest model weights. The dataset is open and curated; weights come from HuggingFace and the slice streams in over HTTP Range requests.
Your GPU trains on that curated data, then keeps only the difference: the "delta", between the original weights and the new ones. Nothing personal on your computer is read.
The client compresses the delta and sends it back. The server validates it and rejects it if:
Valid deltas are averaged into the global model on a fixed schedule; every new version is published to HuggingFace so anyone can test or audit it.
Real-time metrics of the CrowdGPT federated training network.
See which device types are actively training vs. waiting for more pool members.
What we have shipped, what we are building right now, and what comes next.
A 1B model pretrained entirely by the community on consumer GPUs, released openly on Hugging Face.
The next base model. Same pipeline, rebuilt from scratch with a stricter aggregation and cross-check.
See what's happening under the hood
The coordinator writes everything it does to a plain stdout log: shards assigned, updates received, validation failures, aggregation rounds. You can view the whole history with the click of a button.
Contributions: total model updates submitted across all contributors.
Tokens Processed: cumulative tokens trained on by every GPU in the swarm.
Active Clients: GPUs that have pinged the coordinator in the last 10 minutes.
Model Loss: cross-entropy loss of the most recently aggregated update.
Stats refresh automatically every 15 seconds from the HuggingFace CDN. The global model weights are pushed to HuggingFace every 4 hours after aggregated updates complete.
The top contributors powering the CrowdGPT training run. Rankings are based on total tokens processed.
Help grow the CrowdGPT dataset. Submit text or files to the training queue. All submissions are reviewed by maintainers before being added to the public dataset.
Plain text, documents, or any text corpus (.txt, .md, .json, .jsonl, .csv).
Public domain, permissively licensed, or your own original work.
No personal information, no copyrighted material you don't own, no malicious content.
The dataset is fully open and mirrored on HuggingFace. Nothing you submit here is used for training until it has been reviewed by a maintainer.
Questions, bug reports, or interested in running your own coordinator? Reach out.
The best place to report bugs, request features, or discuss technical details of the federated logic.
For business inquiries, partnership opportunities, or sensitive security disclosures.
The primary coordinator API is proudly hosted and maintained by Xevex.
All trained model checkpoints and configuration files are publicly available on HuggingFace.

## Original Extract

View roadmap
×
C
CrowdGPT
home
install
about
status
leaderboard
submit
beta
contact
English
Français
Deutsch
Español
Türkçe
简体中文
source
C
CrowdGPT
home
install
about
status
leaderboard
submit
beta
contact
Source on GitHub
AI owned by everyone .
CrowdGPT is focusing on creating the world's first 0$, datacenterless and 100% open AI. Run the app on your computer, and the AI gets better.
Huggingface
Discord
—
Epoch
—
GPUs currently active
—
Tokens processed
—
Held-out loss
A new paradigm for training AI.
A single company training a frontier model uses thousands of A100s in a datacenter. Most people don't have that hardware. But millions of people have gaming PCs with RTX cards sitting idle. CrowdGPT uses that compute, and the model that gets trained belongs to everyone who helped build it.
Your GPU trains on 10MB data chunks streamed from the public CrowdGPT dataset on HuggingFace; a dataset anyone can inspect or contribute to. No personal files, browser history, or local data are ever accessed or transmitted.
Trains on consumer GPUs that are already powered on. CrowdGPT requires no new datacenters, no additional hardware manufacturing, which is a huge tradeoff of traditional AI.
The client is a single, human-readable Python file that anyone can audit and modify. Model weights are downloadable, 100% free and open on HuggingFace. View server logs
The client is a single open-source Python file which is human-readable. It downloads public training data online, trains locally and uploads only the model update. It cannot access your files or compromise your machine.
The server is protected too. Every update is validated for size, NaN values and norm bounds before it is cross-validated by peers, and repeat offenders are banned from the project.
Read the source before you run it.
Every contribution, even the smallest, makes the model better. Run a client, star the repo, or open an issue.
Star on GitHub
home / install
Install the client
Install everything with one command. Even if you don't have python/git installed.
~6 GB of free disk space. PyTorch alone is around 2.5 GB.
A discrete GPU is strongly recommended. NVIDIA (CUDA), AMD (ROCm / DirectML), Apple Silicon (MPS) and Intel Arc (XPU) are all supported.
A stable internet connection for streaming model weights and dataset chunks.
CPU-only mode works, but is extremely slow.
The installer handles everything else: GPU detection, the correct PyTorch wheel, Python and Git if missing, and an isolated virtual environment.
Copy this into your terminal. It downloads the installer, auto-detects your hardware, and starts the client.
Detect hardware. Installs the right dependencies for your Graphics Card.
Install dependencies. Python and Git are installed if missing, then PyTorch from the correct index and the project's requirements; everything is inside an isolated virtual environment so nothing touches your system Python.
Launch. Downloads the model weights from HuggingFace, streams a dataset chunk, and starts training. You can stop at any time.
Four steps turn a small slice of online, curated text into one shared model.
Diagram unavailable
How CrowdGPT coordinates a bunch of independent GPUs to train a single language model.
1. Assignment
When your device connects, the server gives it one 10MB slice of the public CrowdGPT dataset and the latest model weights. The dataset is open and curated; weights come from HuggingFace and the slice streams in over HTTP Range requests.
Your GPU trains on that curated data, then keeps only the difference: the "delta", between the original weights and the new ones. Nothing personal on your computer is read.
The client compresses the delta and sends it back. The server validates it and rejects it if:
Valid deltas are averaged into the global model on a fixed schedule; every new version is published to HuggingFace so anyone can test or audit it.
Real-time metrics of the CrowdGPT federated training network.
See which device types are actively training vs. waiting for more pool members.
What we have shipped, what we are building right now, and what comes next.
A 1B model pretrained entirely by the community on consumer GPUs, released openly on Hugging Face.
The next base model. Same pipeline, rebuilt from scratch with a stricter aggregation and cross-check.
See what's happening under the hood
The coordinator writes everything it does to a plain stdout log: shards assigned, updates received, validation failures, aggregation rounds. You can view the whole history with the click of a button.
Contributions: total model updates submitted across all contributors.
Tokens Processed: cumulative tokens trained on by every GPU in the swarm.
Active Clients: GPUs that have pinged the coordinator in the last 10 minutes.
Model Loss: cross-entropy loss of the most recently aggregated update.
Stats refresh automatically every 15 seconds from the HuggingFace CDN. The global model weights are pushed to HuggingFace every 4 hours after aggregated updates complete.
The top contributors powering the CrowdGPT training run. Rankings are based on total tokens processed.
Help grow the CrowdGPT dataset. Submit text or files to the training queue. All submissions are reviewed by maintainers before being added to the public dataset.
Plain text, documents, or any text corpus (.txt, .md, .json, .jsonl, .csv).
Public domain, permissively licensed, or your own original work.
No personal information, no copyrighted material you don't own, no malicious content.
The dataset is fully open and mirrored on HuggingFace. Nothing you submit here is used for training until it has been reviewed by a maintainer.
Questions, bug reports, or interested in running your own coordinator? Reach out.
The best place to report bugs, request features, or discuss technical details of the federated logic.
For business inquiries, partnership opportunities, or sensitive security disclosures.
The primary coordinator API is proudly hosted and maintained by Xevex.
All trained model checkpoints and configuration files are publicly available on HuggingFace.
