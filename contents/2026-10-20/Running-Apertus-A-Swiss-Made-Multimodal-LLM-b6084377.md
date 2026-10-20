---
source: "https://blog.x504.dev/p/running-apertus-a-swiss-made-multimodal"
hn_url: "https://news.ycombinator.com/item?id=50162920"
title: "Running Apertus: A Swiss-Made Multimodal LLM"
article_title: "Running Apertus: A Swiss-Made Multimodal LLM - by x504"
image: "https://substackcdn.com/image/fetch/$s_!CaJH!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F799812ad-01ab-40a6-bc5d-f2c2db95f7fa_950x520.png"
author: "asteriskeng"
captured_at: "2026-10-20T09:12:09Z"
capture_tool: "hn-digest"
hn_id: 50162920
score: 1
comments: 0
posted_at: "2026-10-20T08:24:52Z"
tags:
  - hacker-news
---

# Running Apertus: A Swiss-Made Multimodal LLM

- HN: [50162920](https://news.ycombinator.com/item?id=50162920)
- Source: [blog.x504.dev](https://blog.x504.dev/p/running-apertus-a-swiss-made-multimodal)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T08:24:52Z

## Translation

Title: Running Apertus: A Swiss-Made Multimodal LLM
Article title: Running Apertus: A Swiss-Made Multimodal LLM - by x504
Description: In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.

Article text:
Running Apertus: A Swiss-Made Multimodal LLM - by x504
Deep Life Learning
Subscribe Sign in Running Apertus: A Swiss-Made Multimodal LLM
In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.
x504 Oct 20, 2026 Share Nowadays in enterprise software engineering there are two primary models of using LLMs and provision them to a fleet of users:
Commercial Public APIs: Using closed-source models, such as Claude Opus , either directly through vendors like Anthropic or via cloud providers offering managed access to these models, such as Amazon Bedrock and Microsoft Azure AI Foundry.
One of the downsides is privacy as companies are sending private information and proprietary code.
Other problem the developers are facing is deprecation . The vendors tends to deprecate old models. Such models are used by applications in production that were carefully instrumented and designed to enhance the integration with these models. Forcing developers to go through a complete evaluation, benchmark and Blue/Green deployment process before the model is complete deprecated.
There are more concerns. For instance, today it remains unclear whether the judgments done by LLMs are implicitly shaped by geopolitical cues.
Image found in the internet - October 2026
On-Premise Clusters: Deploying open-weights models, such as DeepSeek, Qwen, AMALIA (A Multimodal focused in Portuguese languague and the Portuguese culture ), APERTUS (Fully Open Foundation Model for Sovereign AI developed in Switzerland) on internal hardware, such as NVIDIA. These deployments often involved aggressive quantization techniques (e.g., NVFP4) to optimize the underlaying hardware achieving higher throughput and lower latency compared with the base model while keeping all data interaction within the corporate network.
There are of course challenges when running self-hosted LLMs. Few of them are:
Computing. You have to own your own GPUs.
Knowledge. You have to know how to run an inference-distributed cluster and maintain it.
Operability. As everything you own, you have to maintain, monitor and fix when problems appears
I strongly believe we should reduce our dependency on external LLM vendors. They could change their terms, pricing, or policies at any time, potentially putting our workloads and data at risk. We need to achieve AI and LLM sovereignty by maintaining control over our infrastructure, workloads, and data. The internal data is one of the company’s most valuable assets, and we must ensure it remains private and under our control.
In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.
Before running anything or make any hardware decisions, we need to understand the task at hand. This starts with getting familiar with the characteristics of the LLM we want to run.
We are going to run the 8B (8-billion-parameter) version of APERTUS . Given that the model uses the BF16 (bfloat16) data type, we will need at least 32 GB of GPU VRAM to run it comfortably.
This is a multilingual model optimized for smaller GPUs. You can learn more about it here .
In order to enable thinking mode , we would need to set --reasoning-parser and --default-chat-template-kwargs.enable_thinking .
As today (October 2026) the Apertus 1.5 release has not been completely merged in the Transformers and vLLM libraries hence we will need to use at least --tool-call-parser as extra Argument when running vLLM.
The accelerator we are going to use is 1 x Nvidia RTX A6000 with 48GB VRAM. Which is enough to run the model we describe above for our own purposes.
The framework we are going to use to run inference is the very well stablished vLLM .
vLLM comes with a lot of built-in ready optimizations that allow run faster inference servers. Some of its features are:
Efficient management of attention key and value memory with PagedAttention .
Continuous batching of incoming requests, chunked prefill, prefix caching
We are running on an Ubuntu Server 24.04 LTS. Let’s install few pre-requisites packages. For the sake of this post we are going to use the vLLM Production-stack which is a reference stack for production deployments.
# Initial OS packages updates
apt-get update
# Clone production-stack repository
git clone https://github.com/vllm-project/production-stack.git
cd production-stack/utils
# install kubectl and verify it is installed
bash install-kubectl.sh
kubectl version --client
# Expected similar output:
Client Version: v1.37.1
Kustomize Version: v5.8.1
### Install helm and verify it is installed
bash install-helm.sh
helm version
# Expected similar output:
version.BuildInfo{Version:"v3.22.1", GitCommit:"d6c1a0e4027caa6abc7892e7ea9f0cda7447baa3", GitTreeState:"clean", GoVersion:"go1.26.9"}
### Instal helm diff plugin
# Source: https://github.com/databus23/helm-diff
helm plugin install https://github.com/databus23/helm-diff
# Expected similar output:
Installed plugin: diff
### Installing Minikube with GPU support
sudo apt remove minikube
bash install-minikube-cluster.sh
#### Verifying GPU Configuration
minikube status
# Expected similar output:
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
kubectl describe nodes | grep -i gpu
# Expected similar output:
nvidia.com/gpu: 1
# Run a test pod in kubernetes
kubectl run gpu-test --image=nvidia/cuda:12.2.0-runtime-ubuntu22.04 --restart=Never -- nvidia-smi
# Wait for the pod to be in phase: completed, then:
kubectl logs gpu-test
# Expected similar output:
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 570.195.03 Driver Version: 570.195.03 CUDA Version: 12.8 |
|-----------------------------------------+------------------------+----------------------+
| GPU Name Persistence-M | Bus-Id Disp.A | Volatile Uncorr. ECC |
| Fan Temp Perf Pwr:Usage/Cap | Memory-Usage | GPU-Util Compute M. |
| | | MIG M. |
|=========================================+========================+======================|
| 0 NVIDIA RTX A6000 On | 00000000:00:06.0 Off | Off |
| 30% 34C P8 35W / 300W | 1MiB / 49140MiB | 0% Default |
| | | N/A |
+---------------------
[truncated]
Easy. At this point we have a Kubernetes cluster using minikube with the Nvidia plugin installed in the cluster and ready to run any GPU workload. Let’s now deploy APERTUS with vLLM.
### Deploy vLLM Instance (model: vllm_apertus_1.5_release)
helm repo add vllm https://vllm-project.github.io/production-stack
# Create a file apertus.yaml with the following (or similar) content
# DO NOT FORGET to set your Hugging Face token in 'hf_token'
servingEngineSpec:
runtimeClassName: ""
modelSpec:
- name: "apertus-8b"
repository: "ghcr.io/swiss-ai/vllm_apertus_1.5_release"
tag: "latest-amd64"
modelURL: "swiss-ai/Apertus-v1.5-8B"
replicaCount: 1
requestCPU: 10
requestMemory: "16Gi"
requestGPU: 1
pvcStorage: "50Gi"
pvcAccessMode:
- ReadWriteOnce
vllmConfig:
enableChunkedPrefill: true
enablePrefixCaching: true
maxModelLen: 16384
extraArgs: ["--gpu-memory-utilization", "0.9", "--enable-auto-tool-choice", "--tool-call-parser", "apertus"]
hf_token: hf_1234..
# See first what is going to be installed
helm diff upgrade --install apertus vllm/vllm-stack -f apertus.yaml
# Then deploy
helm upgrade --install apertus vllm/vllm-stack -f apertus.yaml
# Confirm the helm chart was installed
helm list
NAME NAMESPACE REVISION UPDATED STATUS CHART APP VERSION
apertus default 1 2026-10-19 19:34:35.173857134 +0000 UTC deployed vllm-stack-0.1.13
# List the pods an you should see two pods being created
kubectl get pods
NAME READY STATUS RESTARTS AGE
apertus-apertus-8b-deployment-vllm-5c6f6bf484-7twjx 0/1 ContainerCreating 0 26s
apertus-deployment-router-76c84b85fb-c49sx 0/1 ContainerCreating 0 26s
# You could inspect the pod: apertus-8b-deployment-vllm and wills ee it is fetching the docker image
# which contains the model weights. it will take few minutes.
spec.containers{vllm}: Pulling image "ghcr.io/swiss-ai/vllm_apertus_1.5_release:latest-amd64" We need to wait until the model is pulled from HuggingFace.
Then the apertus deployment pod is going to change phase to running . I took few screenshots to give you an idea of the startup process. This will take few minutes to the engine to load all model weights into GPU and start the engine server.
Inference server is ready when you see:
Starting vLLM server on http://0.0.0.0:8000
Available routes are:
Route: /openapi.json, Methods: GET, HEAD
Route: /docs, Methods: GET, HEAD
Route: /docs/oauth2-redirect, Methods: GET, HEAD
...
Let’s confirm our inference server is working properly and ready to take prompts. An easy way is to expose the kubernetes service to the localhost of the Ubuntu server machine and then in another terminal run a the curl command below
# Expose vLLM server to port 30080 on localhost
kubectl port-forward svc/apertus-router-service 30080:80
# Then get the models available on our vLLM server
curl -o- http://localhost:30080/v1/models
You should get a response similar to this:
You can see model swiss-ai/Apertus-v1.5-8B is available! 🎉
Let’s now instal a tiny LLM CLI tool which comes handy to use for sending prompts and streaming the output back to the user in the terminal.
This is not the final interface you would like to give to users. This is only for educational purposes.
As a final interface to users you can use any of these open source solutions:
As a second test let’s open 6 terminals and send the following prompt to see our LLM working!.
Prompt: “Describe kubernetes in 200000 words“
We can confirm our vLLM deployment is working and ready to serve to users. Let’s now use few benchmarking tools to figure out the limits of our deployment ( Model-Hardware-Workload )
For educational purposes we are going to start with vLLM bench . This is an “old“ way provided by vLLm team to benchmark inference systems. it is written in Rust and is basically a client sending different and multiple prompts of different input-len at different rates to your inference server.
In order to take advantage that vLLM libraries are already installed in the deployment Kubernetes pod, we can run vllm-bench inside the pod.
# List the pods
kubectl get pods
# Exec into the apertus deployment pod
kubectl exec -it apertus-apertus-8b-deployment-vllm-6df55cbc8-gpcgl -- /bin/bash
# Run a benchmark with the following configuration
vllm bench serve \
--backend openai \
--base-url http://localhost:8000 \
--endpoint /v1/completions \
--model swiss-ai/Apertus-v1.5-8B \
--dataset-name random \
--num-prompts 500 \
--random-input-len 512 \
--random-output-len 1024 \
--request-rate 5 \
--max-concurrency 32 \
--temperature=0 At the end of the run you will get a report in the terminal. It could as well generate a JSON report for generating graphs afterwards. The screenshot below shows you while running the benchmark.
These metrics are very important, specially TTFT, ITL, TPOT and Request throughput.
Let’s run it again changing input-len from 512 to 1024. We expect the ITL, TPOT won’t change but the TTFT should increase as we are sending a longest prompt. We can confirm this in the results:
Now let’s change the request rate from 5 to 15 and max concurrency from 32 to 64. Now we expect ITL and TPOT be affected as we will be more close to compute-bound.
Here I use a basic configuration for benchmarking the LLM but you could use different datasets, even a custom one and way more advanced configurations.
Let’s now use a more modern benchmarking tool and which is the one recommended as today by the vLLM maintainers.
GuideLLM allows you to evaluate how LLMs perform under real workloads and configurations. It simulates end-to-end interactions with OpenAI-compatible and vLLM-native servers, generates workload patterns that reflect production usage, and produces detailed reports that help teams understand system behavior, resource needs, and operational limits
Again, this is a basic command that illustrate h

[truncated]

## Original Extract

In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.

Running Apertus: A Swiss-Made Multimodal LLM - by x504
Deep Life Learning
Subscribe Sign in Running Apertus: A Swiss-Made Multimodal LLM
In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.
x504 Oct 20, 2026 Share Nowadays in enterprise software engineering there are two primary models of using LLMs and provision them to a fleet of users:
Commercial Public APIs: Using closed-source models, such as Claude Opus , either directly through vendors like Anthropic or via cloud providers offering managed access to these models, such as Amazon Bedrock and Microsoft Azure AI Foundry.
One of the downsides is privacy as companies are sending private information and proprietary code.
Other problem the developers are facing is deprecation . The vendors tends to deprecate old models. Such models are used by applications in production that were carefully instrumented and designed to enhance the integration with these models. Forcing developers to go through a complete evaluation, benchmark and Blue/Green deployment process before the model is complete deprecated.
There are more concerns. For instance, today it remains unclear whether the judgments done by LLMs are implicitly shaped by geopolitical cues.
Image found in the internet - October 2026
On-Premise Clusters: Deploying open-weights models, such as DeepSeek, Qwen, AMALIA (A Multimodal focused in Portuguese languague and the Portuguese culture ), APERTUS (Fully Open Foundation Model for Sovereign AI developed in Switzerland) on internal hardware, such as NVIDIA. These deployments often involved aggressive quantization techniques (e.g., NVFP4) to optimize the underlaying hardware achieving higher throughput and lower latency compared with the base model while keeping all data interaction within the corporate network.
There are of course challenges when running self-hosted LLMs. Few of them are:
Computing. You have to own your own GPUs.
Knowledge. You have to know how to run an inference-distributed cluster and maintain it.
Operability. As everything you own, you have to maintain, monitor and fix when problems appears
I strongly believe we should reduce our dependency on external LLM vendors. They could change their terms, pricing, or policies at any time, potentially putting our workloads and data at risk. We need to achieve AI and LLM sovereignty by maintaining control over our infrastructure, workloads, and data. The internal data is one of the company’s most valuable assets, and we must ensure it remains private and under our control.
In this post, I’ll show you how easy it is to run APERTUS on your own server using the vLLM framework.
Before running anything or make any hardware decisions, we need to understand the task at hand. This starts with getting familiar with the characteristics of the LLM we want to run.
We are going to run the 8B (8-billion-parameter) version of APERTUS . Given that the model uses the BF16 (bfloat16) data type, we will need at least 32 GB of GPU VRAM to run it comfortably.
This is a multilingual model optimized for smaller GPUs. You can learn more about it here .
In order to enable thinking mode , we would need to set --reasoning-parser and --default-chat-template-kwargs.enable_thinking .
As today (October 2026) the Apertus 1.5 release has not been completely merged in the Transformers and vLLM libraries hence we will need to use at least --tool-call-parser as extra Argument when running vLLM.
The accelerator we are going to use is 1 x Nvidia RTX A6000 with 48GB VRAM. Which is enough to run the model we describe above for our own purposes.
The framework we are going to use to run inference is the very well stablished vLLM .
vLLM comes with a lot of built-in ready optimizations that allow run faster inference servers. Some of its features are:
Efficient management of attention key and value memory with PagedAttention .
Continuous batching of incoming requests, chunked prefill, prefix caching
We are running on an Ubuntu Server 24.04 LTS. Let’s install few pre-requisites packages. For the sake of this post we are going to use the vLLM Production-stack which is a reference stack for production deployments.
# Initial OS packages updates
apt-get update
# Clone production-stack repository
git clone https://github.com/vllm-project/production-stack.git
cd production-stack/utils
# install kubectl and verify it is installed
bash install-kubectl.sh
kubectl version --client
# Expected similar output:
Client Version: v1.37.1
Kustomize Version: v5.8.1
### Install helm and verify it is installed
bash install-helm.sh
helm version
# Expected similar output:
version.BuildInfo{Version:"v3.22.1", GitCommit:"d6c1a0e4027caa6abc7892e7ea9f0cda7447baa3", GitTreeState:"clean", GoVersion:"go1.26.9"}
### Instal helm diff plugin
# Source: https://github.com/databus23/helm-diff
helm plugin install https://github.com/databus23/helm-diff
# Expected similar output:
Installed plugin: diff
### Installing Minikube with GPU support
sudo apt remove minikube
bash install-minikube-cluster.sh
#### Verifying GPU Configuration
minikube status
# Expected similar output:
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
kubectl describe nodes | grep -i gpu
# Expected similar output:
nvidia.com/gpu: 1
# Run a test pod in kubernetes
kubectl run gpu-test --image=nvidia/cuda:12.2.0-runtime-ubuntu22.04 --restart=Never -- nvidia-smi
# Wait for the pod to be in phase: completed, then:
kubectl logs gpu-test
# Expected similar output:
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 570.195.03 Driver Version: 570.195.03 CUDA Version: 12.8 |
|-----------------------------------------+------------------------+----------------------+
| GPU Name Persistence-M | Bus-Id Disp.A | Volatile Uncorr. ECC |
| Fan Temp Perf Pwr:Usage/Cap | Memory-Usage | GPU-Util Compute M. |
| | | MIG M. |
|=========================================+========================+======================|
| 0 NVIDIA RTX A6000 On | 00000000:00:06.0 Off | Off |
| 30% 34C P8 35W / 300W | 1MiB / 49140MiB | 0% Default |
| | | N/A |
+---------------------
[truncated]
Easy. At this point we have a Kubernetes cluster using minikube with the Nvidia plugin installed in the cluster and ready to run any GPU workload. Let’s now deploy APERTUS with vLLM.
### Deploy vLLM Instance (model: vllm_apertus_1.5_release)
helm repo add vllm https://vllm-project.github.io/production-stack
# Create a file apertus.yaml with the following (or similar) content
# DO NOT FORGET to set your Hugging Face token in 'hf_token'
servingEngineSpec:
runtimeClassName: ""
modelSpec:
- name: "apertus-8b"
repository: "ghcr.io/swiss-ai/vllm_apertus_1.5_release"
tag: "latest-amd64"
modelURL: "swiss-ai/Apertus-v1.5-8B"
replicaCount: 1
requestCPU: 10
requestMemory: "16Gi"
requestGPU: 1
pvcStorage: "50Gi"
pvcAccessMode:
- ReadWriteOnce
vllmConfig:
enableChunkedPrefill: true
enablePrefixCaching: true
maxModelLen: 16384
extraArgs: ["--gpu-memory-utilization", "0.9", "--enable-auto-tool-choice", "--tool-call-parser", "apertus"]
hf_token: hf_1234..
# See first what is going to be installed
helm diff upgrade --install apertus vllm/vllm-stack -f apertus.yaml
# Then deploy
helm upgrade --install apertus vllm/vllm-stack -f apertus.yaml
# Confirm the helm chart was installed
helm list
NAME NAMESPACE REVISION UPDATED STATUS CHART APP VERSION
apertus default 1 2026-10-19 19:34:35.173857134 +0000 UTC deployed vllm-stack-0.1.13
# List the pods an you should see two pods being created
kubectl get pods
NAME READY STATUS RESTARTS AGE
apertus-apertus-8b-deployment-vllm-5c6f6bf484-7twjx 0/1 ContainerCreating 0 26s
apertus-deployment-router-76c84b85fb-c49sx 0/1 ContainerCreating 0 26s
# You could inspect the pod: apertus-8b-deployment-vllm and wills ee it is fetching the docker image
# which contains the model weights. it will take few minutes.
spec.containers{vllm}: Pulling image "ghcr.io/swiss-ai/vllm_apertus_1.5_release:latest-amd64" We need to wait until the model is pulled from HuggingFace.
Then the apertus deployment pod is going to change phase to running . I took few screenshots to give you an idea of the startup process. This will take few minutes to the engine to load all model weights into GPU and start the engine server.
Inference server is ready when you see:
Starting vLLM server on http://0.0.0.0:8000
Available routes are:
Route: /openapi.json, Methods: GET, HEAD
Route: /docs, Methods: GET, HEAD
Route: /docs/oauth2-redirect, Methods: GET, HEAD
...
Let’s confirm our inference server is working properly and ready to take prompts. An easy way is to expose the kubernetes service to the localhost of the Ubuntu server machine and then in another terminal run a the curl command below
# Expose vLLM server to port 30080 on localhost
kubectl port-forward svc/apertus-router-service 30080:80
# Then get the models available on our vLLM server
curl -o- http://localhost:30080/v1/models
You should get a response similar to this:
You can see model swiss-ai/Apertus-v1.5-8B is available! 🎉
Let’s now instal a tiny LLM CLI tool which comes handy to use for sending prompts and streaming the output back to the user in the terminal.
This is not the final interface you would like to give to users. This is only for educational purposes.
As a final interface to users you can use any of these open source solutions:
As a second test let’s open 6 terminals and send the following prompt to see our LLM working!.
Prompt: “Describe kubernetes in 200000 words“
We can confirm our vLLM deployment is working and ready to serve to users. Let’s now use few benchmarking tools to figure out the limits of our deployment ( Model-Hardware-Workload )
For educational purposes we are going to start with vLLM bench . This is an “old“ way provided by vLLm team to benchmark inference systems. it is written in Rust and is basically a client sending different and multiple prompts of different input-len at different rates to your inference server.
In order to take advantage that vLLM libraries are already installed in the deployment Kubernetes pod, we can run vllm-bench inside the pod.
# List the pods
kubectl get pods
# Exec into the apertus deployment pod
kubectl exec -it apertus-apertus-8b-deployment-vllm-6df55cbc8-gpcgl -- /bin/bash
# Run a benchmark with the following configuration
vllm bench serve \
--backend openai \
--base-url http://localhost:8000 \
--endpoint /v1/completions \
--model swiss-ai/Apertus-v1.5-8B \
--dataset-name random \
--num-prompts 500 \
--random-input-len 512 \
--random-output-len 1024 \
--request-rate 5 \
--max-concurrency 32 \
--temperature=0 At the end of the run you will get a report in the terminal. It could as well generate a JSON report for generating graphs afterwards. The screenshot below shows you while running the benchmark.
These metrics are very important, specially TTFT, ITL, TPOT and Request throughput.
Let’s run it again changing input-len from 512 to 1024. We expect the ITL, TPOT won’t change but the TTFT should increase as we are sending a longest prompt. We can confirm this in the results:
Now let’s change the request rate from 5 to 15 and max concurrency from 32 to 64. Now we expect ITL and TPOT be affected as we will be more close to compute-bound.
Here I use a basic configuration for benchmarking the LLM but you could use different datasets, even a custom one and way more advanced configurations.
Let’s now use a more modern benchmarking tool and which is the one recommended as today by the vLLM maintainers.
GuideLLM allows you to evaluate how LLMs perform under real workloads and configurations. It simulates end-to-end interactions with OpenAI-compatible and vLLM-native servers, generates workload patterns that reflect production usage, and produces detailed reports that help teams understand system behavior, resource needs, and operational limits
Again, this is a basic command that illustrate h

[truncated]
