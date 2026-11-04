---
source: "https://github.com/flitai-io/gpuworks"
hn_url: "https://news.ycombinator.com/item?id=50367630"
title: "Show HN: GPUWorks – A 40 MB GPU cluster manager for vLLM/SGLang, no K8s needed"
article_title: "GitHub - flitai-io/gpuworks: A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. · GitHub"
image: "https://opengraph.githubassets.com/fd3be2ccf6c0d91103e5dc46a4ca1dbd8bfe0930d2d41440422a5f10cc16875d/flitai-io/gpuworks"
author: "wangweinoo1"
captured_at: "2026-11-04T16:54:26Z"
capture_tool: "hn-digest"
hn_id: 50367630
score: 2
comments: 0
posted_at: "2026-11-04T16:09:54Z"
tags:
  - hacker-news
---

# Show HN: GPUWorks – A 40 MB GPU cluster manager for vLLM/SGLang, no K8s needed

- HN: [50367630](https://news.ycombinator.com/item?id=50367630)
- Source: [github.com](https://github.com/flitai-io/gpuworks)
- Score: 2
- Comments: 0
- Posted: 2026-11-04T16:09:54Z

## Translation

Title: Show HN: GPUWorks – A 40 MB GPU cluster manager for vLLM/SGLang, no K8s needed
Article title: GitHub - flitai-io/gpuworks: A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. · GitHub
Description: A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. - flitai-io/gpuworks

Article text:
GitHub - flitai-io/gpuworks: A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
flitai-io
gpuworks Public Notifications You must be signed in to change notification settings
Star 61 ( 61 ) You must be signed in to star a repository
A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
24 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
1,024 Commits 1,024 Commits Folders and files
.github .github LICENSES LICENSES build build cmd cmd docs-en docs-en docs-site docs-site docs docs internal internal pkg pkg scripts scripts secrets secrets test test ui ui .gitattributes .gitattributes .gitignore .gitignore .gitlab-ci.yml .gitlab-ci.yml .golangci.yml .golangci.yml .readthedocs.yaml .readthedocs.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COMMUNITY_GUIDELINES.md COMMUNITY_GUIDELINES.md COMMUNITY_GUIDELINES.zh-CN.md COMMUNITY_GUIDELINES.zh-CN.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING.zh-CN.md CONTRIBUTING.zh-CN.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md SECURITY.zh-CN.md SECURITY.zh-CN.md go.mod go.mod go.sum go.sum known-issues.md known-issues.md requirements-e2e.txt requirements-e2e.txt View all files Repository files navigation
A 40 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators.
GPUWorks is an ultra-lightweight GPU management and Model-as-a-Service (MaaS) platform. The whole control plane, web
UI included, is a 40 MB download. Join your own GPU servers, rented cloud instances or containers, or a Mac with one
command, pick a model, and GPUWorks estimates the memory it needs, picks the GPUs, starts vLLM, SGLang or any engine
an inference template defines (Strata, ninfer, …), and serves it behind one OpenAI-compatible endpoint.
NVIDIA, AMD Instinct and Ryzen AI Max, Huawei Ascend, T-Head PPU, Hygon DCU and Apple Silicon run side by side in the
same cluster.
The pool of capacity is the core: bare-metal servers, public cloud instances and private cloud VMs today, with
Kubernetes clusters and Slurm partitions joining next. On top of it, GPUWorks serves models now and will run
fine-tuning and training jobs on the same accelerators.
Verified model × accelerator combinations deploy in one click. The built-in catalog ships 1,235 launch specs
for 392 models: image, engine, flags and pinned weight revisions for each GPU type. The combinations we have run
on real NVIDIA, AMD, Ascend, PPU and Apple hardware come with launch settings that are known to work, so you pick a
model and press deploy instead of hunting for flags. Qwen, DeepSeek, GLM, Kimi, MiniMax, Llama, Gemma, gpt-oss,
plus embedding, rerank, speech, image and video models.
Heterogeneous accelerators, one scheduler. NVIDIA, AMD, Ascend, PPU, Hygon and Apple Silicon nodes join the same
cluster; the scheduler matches each deployment to hardware its launch spec supports.
Public and private cloud in one pool. Sync instances from RunPod, CompShare, AutoDL, Vast.ai and Alibaba Cloud,
or VMs from ZStack Cloud, install the agent in one click, and schedule them together with your own servers.
Ultra-lightweight. The server and the agent are each one static binary ( CGO_ENABLED=0 ); the server embeds the
web UI (English / 中文) and the agents for every platform, and keeps its state in one SQLite file. The agent dials
out over WebSocket, so nodes need no inbound ports (or the server dials in, for nodes that cannot reach out). In a
drill on Alibaba Cloud, one server ran 1,003 rented GPU machines across 7 regions: over 500 replicas placed and
running within 4 minutes, and 2,000 concurrent requests through the gateway with the server peaking at 9 CPU cores
and 1.6 GB of memory.
Memory estimate and explained placement. Before anything starts, you see how much GPU memory weights, KV
cache and engine overhead need, which GPUs the scheduler picked and why the others were rejected.
Share a GPU between models. dGPU puts a hard memory quota on each instance so several models fit on one card
(NVIDIA via HAMi-core, Ascend via vCANN, PPU via HGGC, Hygon via hy-smi virtual devices). Sleep/wake lets vLLM models take turns
on the same GPUs on NVIDIA, Ascend, PPU and Hygon.
Scale out. Replicas across nodes, multi-node tensor/pipeline parallel, and prefill/decode disaggregation
(vLLM + NIXL on NVIDIA, vLLM-Ascend + Mooncake on Ascend, SGLang or vLLM + Mooncake on PPU, within one host or over RDMA).
OpenAI-compatible gateway. Chat, completions, embeddings, rerank, audio, images, video and realtime under
/v1 ; /v1/messages and /v1/responses pass through. Least-request, prefix-affinity or failover routing, API
keys, per-key token usage.
See what is going on. GPU and engine metrics, searchable logs, per-request traces with a timeline (OTLP export),
a web terminal into hosts and containers, and a built-in benchmark that sweeps concurrency and finds the SLA limit.
AI assistant and MCP. Ask questions about the cluster in the UI, or connect Claude Code / any MCP client to the
read-only MCP endpoint.
Air-gapped ready. Import images and models from archives or an HTTP file server, and sync models between nodes.
1. Install the server on a Linux host (x86_64 or arm64) or an Apple Silicon Mac:
bash <( curl -fsSL https://dl.gpuworks.ai/get.sh )
The script picks the fastest download source (including mirrors in mainland China), verifies the checksum and runs
the installer, which sets up a service on port 8800 and prints the URL and the admin password. Add --quiet to
accept every default. To check each step yourself, follow the manual install .
2. Add a GPU node. Open the web UI, go to Compute nodes → Add node , and run the generated command on the GPU
host:
curl -fsSL http:// < server > :8800/install.sh | sh -s -- --server http:// < server > :8800 --token < node-token >
3. Deploy a model. Open Deploy , pick a model such as Qwen/Qwen3-1.7B and a recommended inference template,
review the configuration evaluation (memory estimate and placement), and deploy. Then create an inference token under
API Tokens and call the model by its public model name:
curl http:// < server > :8800/v1/chat/completions \
-H " Authorization: Bearer gw-... " -H " Content-Type: application/json " \
-d ' {"model": "qwen3-1.7b", "messages": [{"role": "user", "content": "Hello!"}]} '
The full walkthrough is in Getting started . For a Mac, see
Run your first local model on a Mac ; for air-gapped sites, see
Offline deployment .
Install with an AI coding agent
Paste this into Claude Code, Codex, Cursor or a similar agent that has shell access to your server:
Install the GPUWorks control plane on this host and verify it.
1. Run `bash <(curl -fsSL https://dl.gpuworks.ai/get.sh) --quiet`. It downloads the latest build, verifies its
checksum and installs a service on port 8800 (use sudo when asked). If I have no sudo or no systemd, append
`--dir ~/gpuworks-server --no-systemd` instead.
2. Verify with `curl -fsS http://<listen-address>/healthz` (expect "status":"ok") and show me the install summary:
URL, install directory, data directory and the initial admin password location.
3. Do not open firewall ports, change other services, or join GPU nodes. Tell me which address GPU nodes should
use to reach the server; I will add nodes from the web UI.
Supported hardware
Accelerator
Whole-GPU deploy
dGPU (shared GPU)
Sleep/wake
Multi-node
PD disaggregation
NVIDIA (CUDA)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
✅ vLLM + NIXL
AMD (ROCm: Instinct MI300 series, Ryzen AI Max)
✅ Docker or containerd
—
—
✅ vLLM
—
Huawei Ascend (CANN)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM-Ascend
✅
✅ vLLM-Ascend + Mooncake
T-Head PPU (HGGC)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
✅ SGLang + Mooncake (same node or over RDMA), vLLM + Mooncake (RDMA)
Hygon DCU (DTK)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
—
Apple Silicon (M1–M5, macOS 15+)
✅ vllm-metal, MLX, llama.cpp as native processes
—
—
—
—
✅ verified on real hardware · — not supported or not validated. Sleep/wake beyond NVIDIA is
verified on Ascend 910C, PPU ZW810E and Hygon K100; the tested versions and limits are in
choosing features .
GPUs are detected with the vendor CLIs ( nvidia-smi , npu-smi , ppu-smi , hy-smi ) or, for AMD, amdgpu sysfs and KFD, so there is no NVML or cgo
dependency. Per-combination status lives in the support matrix and the
Ascend matrix .
Engines: vLLM (incl. vLLM-Ascend and vLLM-Omni), SGLang, and on Mac vllm-metal, MLX-VLM, MLX-Audio and
llama.cpp. Built-in images are published as gpuworks/gpuworks-runner on Docker Hub, Quay and Huawei SWR
( image families ); you can also bring your own image.
The built-in catalog pins every model to fixed Hugging Face and ModelScope revisions and ships launch specs per
engine and accelerator. The deploy page recommends the specs that match the model and the target accelerator.
Catalog totals: 392 models from 77 organizations, 1,235 launch specs (745 NVIDIA, 216 PPU, 209 Ascend, 41 Hygon,
15 AMD, 9 Apple). Models outside the catalog can be registered from Hugging Face, ModelScope, a local directory, NFS or an
HTTP file server and deployed with your own arguments.
Overview of a production cluster with NVIDIA, PPU, Ascend and Hygon GPUs
Scheduling analysis : models taking turns on the same GPUs, card by card
Monitoring in production: latency, request shapes, concurrency and KV cache over 7 days
Usage analytics : tokens by hour, API key and model
PD disaggregation on Ascend 910C: KV hand-off, latency breakdown, requests in flight
GPU health checks , diagnosed by the AI assistant
Network test : HCCS card-to-card reachability matrix
Benchmark report: throughput, TTFT, latency
Model catalog with launch specs per vendor and engine
Playground against any deployment
Architecture
flowchart LR
user["Browser / CLI / MCP client"] -->|"Web UI, REST API"| server
app["Apps (OpenAI SDK)"] -->|"/v1/* + API key"| server
subgraph server["gpuworks server (one binary)"]
direction TB
gw["OpenAI gateway"]
ctl["Controller + scheduler<br/>memory estimate"]
db[("SQLite<br/>metrics · logs · traces")]
end
subgraph node1["GPU node"]
a1["gpuworks-agent"] --> e1["vLLM / SGLang<br/>Docker · nerdctl · runveil"]
end
subgraph node2["Mac"]
a2["gpuworks-agent"] --> e2["vllm-metal / MLX / llama.cpp"]
end
a1 -->|"outbou

[truncated]

## Original Extract

A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. - flitai-io/gpuworks

GitHub - flitai-io/gpuworks: A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
flitai-io
gpuworks Public Notifications You must be signed in to change notification settings
Star 61 ( 61 ) You must be signed in to star a repository
A ~45 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
24 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
1,024 Commits 1,024 Commits Folders and files
.github .github LICENSES LICENSES build build cmd cmd docs-en docs-en docs-site docs-site docs docs internal internal pkg pkg scripts scripts secrets secrets test test ui ui .gitattributes .gitattributes .gitignore .gitignore .gitlab-ci.yml .gitlab-ci.yml .golangci.yml .golangci.yml .readthedocs.yaml .readthedocs.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COMMUNITY_GUIDELINES.md COMMUNITY_GUIDELINES.md COMMUNITY_GUIDELINES.zh-CN.md COMMUNITY_GUIDELINES.zh-CN.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING.zh-CN.md CONTRIBUTING.zh-CN.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md SECURITY.zh-CN.md SECURITY.zh-CN.md go.mod go.mod go.sum go.sum known-issues.md known-issues.md requirements-e2e.txt requirements-e2e.txt View all files Repository files navigation
A 40 MB MaaS platform: one cluster across public and private clouds, on any mix of accelerators.
GPUWorks is an ultra-lightweight GPU management and Model-as-a-Service (MaaS) platform. The whole control plane, web
UI included, is a 40 MB download. Join your own GPU servers, rented cloud instances or containers, or a Mac with one
command, pick a model, and GPUWorks estimates the memory it needs, picks the GPUs, starts vLLM, SGLang or any engine
an inference template defines (Strata, ninfer, …), and serves it behind one OpenAI-compatible endpoint.
NVIDIA, AMD Instinct and Ryzen AI Max, Huawei Ascend, T-Head PPU, Hygon DCU and Apple Silicon run side by side in the
same cluster.
The pool of capacity is the core: bare-metal servers, public cloud instances and private cloud VMs today, with
Kubernetes clusters and Slurm partitions joining next. On top of it, GPUWorks serves models now and will run
fine-tuning and training jobs on the same accelerators.
Verified model × accelerator combinations deploy in one click. The built-in catalog ships 1,235 launch specs
for 392 models: image, engine, flags and pinned weight revisions for each GPU type. The combinations we have run
on real NVIDIA, AMD, Ascend, PPU and Apple hardware come with launch settings that are known to work, so you pick a
model and press deploy instead of hunting for flags. Qwen, DeepSeek, GLM, Kimi, MiniMax, Llama, Gemma, gpt-oss,
plus embedding, rerank, speech, image and video models.
Heterogeneous accelerators, one scheduler. NVIDIA, AMD, Ascend, PPU, Hygon and Apple Silicon nodes join the same
cluster; the scheduler matches each deployment to hardware its launch spec supports.
Public and private cloud in one pool. Sync instances from RunPod, CompShare, AutoDL, Vast.ai and Alibaba Cloud,
or VMs from ZStack Cloud, install the agent in one click, and schedule them together with your own servers.
Ultra-lightweight. The server and the agent are each one static binary ( CGO_ENABLED=0 ); the server embeds the
web UI (English / 中文) and the agents for every platform, and keeps its state in one SQLite file. The agent dials
out over WebSocket, so nodes need no inbound ports (or the server dials in, for nodes that cannot reach out). In a
drill on Alibaba Cloud, one server ran 1,003 rented GPU machines across 7 regions: over 500 replicas placed and
running within 4 minutes, and 2,000 concurrent requests through the gateway with the server peaking at 9 CPU cores
and 1.6 GB of memory.
Memory estimate and explained placement. Before anything starts, you see how much GPU memory weights, KV
cache and engine overhead need, which GPUs the scheduler picked and why the others were rejected.
Share a GPU between models. dGPU puts a hard memory quota on each instance so several models fit on one card
(NVIDIA via HAMi-core, Ascend via vCANN, PPU via HGGC, Hygon via hy-smi virtual devices). Sleep/wake lets vLLM models take turns
on the same GPUs on NVIDIA, Ascend, PPU and Hygon.
Scale out. Replicas across nodes, multi-node tensor/pipeline parallel, and prefill/decode disaggregation
(vLLM + NIXL on NVIDIA, vLLM-Ascend + Mooncake on Ascend, SGLang or vLLM + Mooncake on PPU, within one host or over RDMA).
OpenAI-compatible gateway. Chat, completions, embeddings, rerank, audio, images, video and realtime under
/v1 ; /v1/messages and /v1/responses pass through. Least-request, prefix-affinity or failover routing, API
keys, per-key token usage.
See what is going on. GPU and engine metrics, searchable logs, per-request traces with a timeline (OTLP export),
a web terminal into hosts and containers, and a built-in benchmark that sweeps concurrency and finds the SLA limit.
AI assistant and MCP. Ask questions about the cluster in the UI, or connect Claude Code / any MCP client to the
read-only MCP endpoint.
Air-gapped ready. Import images and models from archives or an HTTP file server, and sync models between nodes.
1. Install the server on a Linux host (x86_64 or arm64) or an Apple Silicon Mac:
bash <( curl -fsSL https://dl.gpuworks.ai/get.sh )
The script picks the fastest download source (including mirrors in mainland China), verifies the checksum and runs
the installer, which sets up a service on port 8800 and prints the URL and the admin password. Add --quiet to
accept every default. To check each step yourself, follow the manual install .
2. Add a GPU node. Open the web UI, go to Compute nodes → Add node , and run the generated command on the GPU
host:
curl -fsSL http:// < server > :8800/install.sh | sh -s -- --server http:// < server > :8800 --token < node-token >
3. Deploy a model. Open Deploy , pick a model such as Qwen/Qwen3-1.7B and a recommended inference template,
review the configuration evaluation (memory estimate and placement), and deploy. Then create an inference token under
API Tokens and call the model by its public model name:
curl http:// < server > :8800/v1/chat/completions \
-H " Authorization: Bearer gw-... " -H " Content-Type: application/json " \
-d ' {"model": "qwen3-1.7b", "messages": [{"role": "user", "content": "Hello!"}]} '
The full walkthrough is in Getting started . For a Mac, see
Run your first local model on a Mac ; for air-gapped sites, see
Offline deployment .
Install with an AI coding agent
Paste this into Claude Code, Codex, Cursor or a similar agent that has shell access to your server:
Install the GPUWorks control plane on this host and verify it.
1. Run `bash <(curl -fsSL https://dl.gpuworks.ai/get.sh) --quiet`. It downloads the latest build, verifies its
checksum and installs a service on port 8800 (use sudo when asked). If I have no sudo or no systemd, append
`--dir ~/gpuworks-server --no-systemd` instead.
2. Verify with `curl -fsS http://<listen-address>/healthz` (expect "status":"ok") and show me the install summary:
URL, install directory, data directory and the initial admin password location.
3. Do not open firewall ports, change other services, or join GPU nodes. Tell me which address GPU nodes should
use to reach the server; I will add nodes from the web UI.
Supported hardware
Accelerator
Whole-GPU deploy
dGPU (shared GPU)
Sleep/wake
Multi-node
PD disaggregation
NVIDIA (CUDA)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
✅ vLLM + NIXL
AMD (ROCm: Instinct MI300 series, Ryzen AI Max)
✅ Docker or containerd
—
—
✅ vLLM
—
Huawei Ascend (CANN)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM-Ascend
✅
✅ vLLM-Ascend + Mooncake
T-Head PPU (HGGC)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
✅ SGLang + Mooncake (same node or over RDMA), vLLM + Mooncake (RDMA)
Hygon DCU (DTK)
✅ Docker, containerd or runveil (built-in rootless runtime)
✅ memory quota at MiB granularity
✅ vLLM
✅
—
Apple Silicon (M1–M5, macOS 15+)
✅ vllm-metal, MLX, llama.cpp as native processes
—
—
—
—
✅ verified on real hardware · — not supported or not validated. Sleep/wake beyond NVIDIA is
verified on Ascend 910C, PPU ZW810E and Hygon K100; the tested versions and limits are in
choosing features .
GPUs are detected with the vendor CLIs ( nvidia-smi , npu-smi , ppu-smi , hy-smi ) or, for AMD, amdgpu sysfs and KFD, so there is no NVML or cgo
dependency. Per-combination status lives in the support matrix and the
Ascend matrix .
Engines: vLLM (incl. vLLM-Ascend and vLLM-Omni), SGLang, and on Mac vllm-metal, MLX-VLM, MLX-Audio and
llama.cpp. Built-in images are published as gpuworks/gpuworks-runner on Docker Hub, Quay and Huawei SWR
( image families ); you can also bring your own image.
The built-in catalog pins every model to fixed Hugging Face and ModelScope revisions and ships launch specs per
engine and accelerator. The deploy page recommends the specs that match the model and the target accelerator.
Catalog totals: 392 models from 77 organizations, 1,235 launch specs (745 NVIDIA, 216 PPU, 209 Ascend, 41 Hygon,
15 AMD, 9 Apple). Models outside the catalog can be registered from Hugging Face, ModelScope, a local directory, NFS or an
HTTP file server and deployed with your own arguments.
Overview of a production cluster with NVIDIA, PPU, Ascend and Hygon GPUs
Scheduling analysis : models taking turns on the same GPUs, card by card
Monitoring in production: latency, request shapes, concurrency and KV cache over 7 days
Usage analytics : tokens by hour, API key and model
PD disaggregation on Ascend 910C: KV hand-off, latency breakdown, requests in flight
GPU health checks , diagnosed by the AI assistant
Network test : HCCS card-to-card reachability matrix
Benchmark report: throughput, TTFT, latency
Model catalog with launch specs per vendor and engine
Playground against any deployment
Architecture
flowchart LR
user["Browser / CLI / MCP client"] -->|"Web UI, REST API"| server
app["Apps (OpenAI SDK)"] -->|"/v1/* + API key"| server
subgraph server["gpuworks server (one binary)"]
direction TB
gw["OpenAI gateway"]
ctl["Controller + scheduler<br/>memory estimate"]
db[("SQLite<br/>metrics · logs · traces")]
end
subgraph node1["GPU node"]
a1["gpuworks-agent"] --> e1["vLLM / SGLang<br/>Docker · nerdctl · runveil"]
end
subgraph node2["Mac"]
a2["gpuworks-agent"] --> e2["vllm-metal / MLX / llama.cpp"]
end
a1 -->|"outbou

[truncated]
