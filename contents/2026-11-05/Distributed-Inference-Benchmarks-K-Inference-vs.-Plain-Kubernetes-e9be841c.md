---
source: "https://www.buoyant.io/blog/distributed-inference-benchmarks-kinference-vs-plain-kubernetes"
hn_url: "https://news.ycombinator.com/item?id=50383051"
title: "Distributed Inference Benchmarks: [K]Inference vs. Plain Kubernetes"
article_title: "Distributed Inference Benchmarks: Kinference vs. Kubernetes"
image: "https://cdn.prod.website-files.com/679957f782f1eca8a9594edb/6ae4bfba107b48a9243e6609_Blog%20Preview%20Image%201200x630.webp"
author: "darksoul"
captured_at: "2026-11-05T19:35:54Z"
capture_tool: "hn-digest"
hn_id: 50383051
score: 1
comments: 0
posted_at: "2026-11-05T19:07:41Z"
tags:
  - hacker-news
---

# Distributed Inference Benchmarks: [K]Inference vs. Plain Kubernetes

- HN: [50383051](https://news.ycombinator.com/item?id=50383051)
- Source: [www.buoyant.io](https://www.buoyant.io/blog/distributed-inference-benchmarks-kinference-vs-plain-kubernetes)
- Score: 1
- Comments: 0
- Posted: 2026-11-05T19:07:41Z

## Translation

Title: Distributed Inference Benchmarks: [K]Inference vs. Plain Kubernetes
Article title: Distributed Inference Benchmarks: Kinference vs. Kubernetes
Description: Learn how Kinference cuts TTFT by up to 95% and boosts Kubernetes inference capacity compared to plain vLLM deployments in our latest benchmark.

Article text:
Distributed Inference Benchmarks: Kinference vs. Kubernetes
Skip to main content
Products AI Inference
Buoyant AI Run frontier-grade open-weight models with unlimited tokens and full inference control. Kinference [kay-inference] The open source inference control plane for vLLM and Kubernetes. Enterprise Service Mesh Buoyant Enterprise for Linkerd Enterprise-grade security and compliance for production infrastructure. Linkerd CNCF open source security for every Kubernetes cluster. The AI Kubernetes Show The Kubernetes AI Show dives deep into the real-world challenges of adopting AI on Kubernetes platforms. Watch now
Use Cases AI Inference
Cut coding inference spend
Move agentic coding to open-weight models running on GPUs you control. Burst to the cloud when your GPUs are full When GPUs saturate, excess requests burst to the cloud. Get full efficiency from every GPU More useful work from every GPU, even at peak utilization. Every app, model, tenant gets an SLO Cache-aware routing that predicts latency and holds every tenant’s SLOs. Intent-based routing Serve every prompt based on its context. Enterprise Service Mesh FIPS for FedRAMP, CMMC, and HIPAA A validated control plane and data plane, audit-ready in 90 days Zero trust in Kubernetes Strict mTLS and deny-by-default policy for every workload Multicluster and failover Federated services and cross-cluster mTLS for migrations Cross-zone cost reduction HAZL keeps healthy traffic in-zone and cuts transfer costs Observability and reliability Golden metrics, automated cert rotation, and OpenTelemetry tracing The AI Kubernetes Show The Kubernetes AI Show dives deep into the real-world challenges of adopting AI on Kubernetes platforms. Watch now
Resources Runbook Live fireside chats, demos, AMAs, and lessons learned, as well as self-paced labs and a hands-on escape room. Explore Runbook
Runbook Live fireside chats, demos, AMAs, and lessons learned as well as self-paced labs. Case studies Learn how organizations across the world run Linkerd in production. Ebooks, videos and resources Enterprise-focused documents to better inform you about the service mesh Buoyant Linkerd forum Get support for Linkerd from the maintainers, experts, and Linkerd community! Linkerd vs Istio How do the two service meshes compare? Linkerd vs Cilium Five Key Differences in 2025 Guide to mTLS and Kubernetes A Kubernetes engineer's guide to mutual TLS.
About About us Our mission and our story Partners Meet our partners, or become one. Together, we help teams worldwide use Linkerd with confidence. Contact us Have a question or ready to get started with Linkerd? Careers Our mission and our story Newsroom The latest news about the service mesh
Announcing Kinference: Scalable inference serving for Kubernetes |
Blog home Distributed Inference Benchmarks: Kinference vs. Plain Kubernetes
Linkerd Production Readiness Pre-Launch Checklist
eBPF, Sidecars, and the
Future of The Service Mesh
A Kubernetes engineer’s guide to mTLS
William Morgan
Kinference (“kay-inference") is an inference control plane for Kubernetes and vLLM . Kinference sits in front of a heterogeneous set of GPUs and models, each with different capabilities and levels of performance; takes inference requests from a set of heterogeneous clients, each with different requirements and competing demands; and figures out how to satisfy the clients without overloading the cluster. And it gives you declarative policies to determine exactly how that happens.
Part of how Kinference delivers on its promise is by load balancing inference requests based on predicted TTFT . In other words, given an inference request (a “prompt”), Kinference predicts the time-to-first-token (TTFT) for that request against each candidate endpoint, then sends it to the worker (e.g. vLLM pod) with the lowest predicted latency. TTFT is a good metric for this because it’s independent of the output size, and because load balancing here can make a tremendous difference–sometimes multiple orders of magnitude.
To make these TTFT predictions, Kinference makes use of a variety of data about the underlying vLLM pods, including contents of their KV cache . The KV cache is where inference workers store intermediate results of the prefill phase of inference computation. These results may be reused on future requests–but only if they share a prefix. Since many inference requests share a substantial prefix (for example, successive turns in the same conversation), sending a request to a worker with a high KV cache overlap can save a tremendous amount of time.
In this benchmark, we attempt to quantify exactly how much Kinference’s intelligent load balancing helps, by comparing inference requests to a multi-pod vLLM deployment with Kinference against the same requests without Kinference.
We look at two basic measures of distributed inference performance:
Unloaded latency. When the system is under light load and nothing is queuing, is Kinference faster?
Capacity. When the system is under heavy load, how much more capacity can Kinference provide while maintaining a given SLO?
Measuring unloaded latency demonstrates the value of KV-cache-aware load balancing in isolation, i.e. independent of queueing and other factors. Measuring capacity within an SLO, by contrast, demonstrates real-world performance: how does KV-cache-aware balancing improve a system’s overall ability to serve inference traffic?
To test our metrics, we created an AWS cluster where we deployed two identical pools of vLLM 0.31.0, running the Qwen-3.8-27B-FP8 open weight model. Each pool comprised 3 vLLM pods, each running on its own L40S GPU. The two arms of our test were two OpenAI chat-completion HTTP endpoints: one from Kinference, pointing to the first vLLM pool, and one from the Kubernetes Service record pointing to the second vLLM pool.
Our null hypothesis was the two arms would perform the same. Our alternative hypothesis, of course, was that we would see lower latency from Kinference on workloads whose requests share prefixes. We chose the primary metric for this comparison up front: median TTFT on two workloads, shared-prefix and multiturn (described below). The actual benchmarking was done with GuideLLM 0.8.1, which we ran on its own node in the cluster.
There are a couple things worth noting about this setup:
This is a relatively modern Qwen model, running as a dense checkpoint on relatively old GPUs. It’s slow: we’re reporting reqs/min, below, a far cry from our usual RPS fare with Linkerd . Similarly, our latency SLOs are relatively high–not appropriate for chat, but plausible for offline agentic workloads.
Different models and different hardware will deliver very different results, so the raw numbers here are less important as they are specific to these choices. The comparative numbers are the important ones.
For the service arm of our test, GuideLLM sent requests straight to the pool’s Kubernetes Service (ClusterIP). kube-proxy, in iptables mode, picks a random pod for each new connection, and GuideLLM opens a new connection for almost every request. Thus, the service arm saw requests distributed randomly across its three replicas.
For the unloaded latency benchmark, we used open-loop Poisson arrivals at a low rate which we empirically determined would not saturate the vLLM workers. For the capacity benchmark, we used closed-loop traffic at a fixed concurrency, i.e. each client sent traffic as quickly as it could. For each condition and trial within these benchmarks, both arms received identical prompts, and in the latency benchmarks requests were sent at identical times. Prompts and prefixes were new in every condition and trial.
Finally, each benchmark covered three types of inference traffic: a control workload, a shared-prefix workload, and a multiturn workload:
We ran each workload with a 60 s warmup period and a 45 s cooldown period, and measured the results over 240 s. The rate column and these windows apply to the latency benchmark; the capacity benchmark used only shared-prefix and multiturn workloads at fixed concurrency (see Results and analysis: Capacity benchmark ).
In total the latency benchmark had 3 workloads × 2 arms × 3 trials = 18 runs. The arms ran one after the other, and order alternated between conditions. A request was measured iff it started inside the measurement window.
Results and analysis: Unloaded latency benchmark
Summary: In the shared-prefix workload, Kinference cut median TTFT by 95% and median E2E latency by 45% . In the multiturn workload, Kinference cut median TTFT by about 65% and median E2E latency by about 32% , versus standard Kubernetes load balancing.
(Confidence intervals are 95% CIs from a paired hierarchical bootstrap with 2,000 resamples.)
In our control conditions, median TTFT is indistinguishable between the arms (6.92 s vs 6.83 s). Without reusable prefixes, Kinference adds no measurable latency, and the two pools perform alike–as we expect. Interestingly, the control p95 is another matter: 21.5 s on the service arm vs 12.3 s with Kinference. This is likely due to Kinference’s usage of vLLM queue size as well as KV cache overlap to predict TTFT–even with prompts that share zero overlap, Kinference can still make intelligent balancing decisions and reduce tail latencies.
In our shared-prefix workload, Kinference cuts median TTFT by about 95% (6.88 s to 0.31 s) and p95 TTFT by about 50% (18.3 s to 9.2 s), a clear improvement over default Kubernetes load balancing.
In our multiturn workload, Kinference cuts median TTFT by about 65% (1.70 s to 0.59 s) and p95 TTFT by about 40% (12.8 s to 7.7 s)–also a clear improvement. Note that the gain here is smaller than on shared-prefix because the first turn of every conversation has nothing to reuse.
Secondary metrics: ITL, E2E latency
Turning to inter-token latency (ITL) and end-to-end (E2E) latency, we see a similar story. In our shared-prefix workload, Kinference cuts median end-to-end latency by about 45% (20.5 s to 11.3 s) and median ITL by about 4% (44 ms to 42 ms). In our multiturn workload, it cuts median end-to-end latency by about 32% (13.4 s to 9.1 s) and median ITL by about 10% (47 ms to 43 ms). On control, neither difference is significant; both CIs reach 1.
Most of the end-to-end gain comes from the shorter TTFT, which is a component of E2E latency. The rest comes from faster token generation: Kinference’s replicas recompute fewer uncached prefixes, so there is less work that competes with producing the remaining tokens.
Output throughput was about equal on every workload, with ratios of 0.98–1.05 per trial, as expected at a fixed arrival rate that both arms keep up with. No request failed or was left unfinished.
Results and analysis: Capacity benchmark
Summary: At peak, Kinference served about 7.4× the service arm’s goodput on shared-prefix and about 3.5× on multiturn. It was able to attain our chosen SLO at concurrency 32 (the highest measured) on our shared-prefix workload, and concurrency 8 on the multiturn workload, whereas the kube-proxy approach was not able to attain the SLO at any concurrency level.
In our unloaded latency benchmark, we sought to measure the effect of Kinference’s KV-cache-aware load balancing in isolation. In our capacity benchmark, we ask a different question: with N users, each sending its next request as soon as the last answer finishes, how many requests per minute can each endpoint serve within a latency SLO?
Our null hypothesis was that at each concurrency, the two endpoints would have the same goodput : the number of requests per minute that met the SLO. Our alternative hypothesis was that Kinference would have higher goodput, and would keep meeting the SLO at higher concurrency. As before, we chose the primary metric up front: the goodput difference (Kinference − service) at each concurrency.
What SLO do we use? Ideally, we’d pick one from first principles–for example, we might not want a user waiting on a live chat bot to have to wait

[truncated]

## Original Extract

Learn how Kinference cuts TTFT by up to 95% and boosts Kubernetes inference capacity compared to plain vLLM deployments in our latest benchmark.

Distributed Inference Benchmarks: Kinference vs. Kubernetes
Skip to main content
Products AI Inference
Buoyant AI Run frontier-grade open-weight models with unlimited tokens and full inference control. Kinference [kay-inference] The open source inference control plane for vLLM and Kubernetes. Enterprise Service Mesh Buoyant Enterprise for Linkerd Enterprise-grade security and compliance for production infrastructure. Linkerd CNCF open source security for every Kubernetes cluster. The AI Kubernetes Show The Kubernetes AI Show dives deep into the real-world challenges of adopting AI on Kubernetes platforms. Watch now
Use Cases AI Inference
Cut coding inference spend
Move agentic coding to open-weight models running on GPUs you control. Burst to the cloud when your GPUs are full When GPUs saturate, excess requests burst to the cloud. Get full efficiency from every GPU More useful work from every GPU, even at peak utilization. Every app, model, tenant gets an SLO Cache-aware routing that predicts latency and holds every tenant’s SLOs. Intent-based routing Serve every prompt based on its context. Enterprise Service Mesh FIPS for FedRAMP, CMMC, and HIPAA A validated control plane and data plane, audit-ready in 90 days Zero trust in Kubernetes Strict mTLS and deny-by-default policy for every workload Multicluster and failover Federated services and cross-cluster mTLS for migrations Cross-zone cost reduction HAZL keeps healthy traffic in-zone and cuts transfer costs Observability and reliability Golden metrics, automated cert rotation, and OpenTelemetry tracing The AI Kubernetes Show The Kubernetes AI Show dives deep into the real-world challenges of adopting AI on Kubernetes platforms. Watch now
Resources Runbook Live fireside chats, demos, AMAs, and lessons learned, as well as self-paced labs and a hands-on escape room. Explore Runbook
Runbook Live fireside chats, demos, AMAs, and lessons learned as well as self-paced labs. Case studies Learn how organizations across the world run Linkerd in production. Ebooks, videos and resources Enterprise-focused documents to better inform you about the service mesh Buoyant Linkerd forum Get support for Linkerd from the maintainers, experts, and Linkerd community! Linkerd vs Istio How do the two service meshes compare? Linkerd vs Cilium Five Key Differences in 2025 Guide to mTLS and Kubernetes A Kubernetes engineer's guide to mutual TLS.
About About us Our mission and our story Partners Meet our partners, or become one. Together, we help teams worldwide use Linkerd with confidence. Contact us Have a question or ready to get started with Linkerd? Careers Our mission and our story Newsroom The latest news about the service mesh
Announcing Kinference: Scalable inference serving for Kubernetes |
Blog home Distributed Inference Benchmarks: Kinference vs. Plain Kubernetes
Linkerd Production Readiness Pre-Launch Checklist
eBPF, Sidecars, and the
Future of The Service Mesh
A Kubernetes engineer’s guide to mTLS
William Morgan
Kinference (“kay-inference") is an inference control plane for Kubernetes and vLLM . Kinference sits in front of a heterogeneous set of GPUs and models, each with different capabilities and levels of performance; takes inference requests from a set of heterogeneous clients, each with different requirements and competing demands; and figures out how to satisfy the clients without overloading the cluster. And it gives you declarative policies to determine exactly how that happens.
Part of how Kinference delivers on its promise is by load balancing inference requests based on predicted TTFT . In other words, given an inference request (a “prompt”), Kinference predicts the time-to-first-token (TTFT) for that request against each candidate endpoint, then sends it to the worker (e.g. vLLM pod) with the lowest predicted latency. TTFT is a good metric for this because it’s independent of the output size, and because load balancing here can make a tremendous difference–sometimes multiple orders of magnitude.
To make these TTFT predictions, Kinference makes use of a variety of data about the underlying vLLM pods, including contents of their KV cache . The KV cache is where inference workers store intermediate results of the prefill phase of inference computation. These results may be reused on future requests–but only if they share a prefix. Since many inference requests share a substantial prefix (for example, successive turns in the same conversation), sending a request to a worker with a high KV cache overlap can save a tremendous amount of time.
In this benchmark, we attempt to quantify exactly how much Kinference’s intelligent load balancing helps, by comparing inference requests to a multi-pod vLLM deployment with Kinference against the same requests without Kinference.
We look at two basic measures of distributed inference performance:
Unloaded latency. When the system is under light load and nothing is queuing, is Kinference faster?
Capacity. When the system is under heavy load, how much more capacity can Kinference provide while maintaining a given SLO?
Measuring unloaded latency demonstrates the value of KV-cache-aware load balancing in isolation, i.e. independent of queueing and other factors. Measuring capacity within an SLO, by contrast, demonstrates real-world performance: how does KV-cache-aware balancing improve a system’s overall ability to serve inference traffic?
To test our metrics, we created an AWS cluster where we deployed two identical pools of vLLM 0.31.0, running the Qwen-3.8-27B-FP8 open weight model. Each pool comprised 3 vLLM pods, each running on its own L40S GPU. The two arms of our test were two OpenAI chat-completion HTTP endpoints: one from Kinference, pointing to the first vLLM pool, and one from the Kubernetes Service record pointing to the second vLLM pool.
Our null hypothesis was the two arms would perform the same. Our alternative hypothesis, of course, was that we would see lower latency from Kinference on workloads whose requests share prefixes. We chose the primary metric for this comparison up front: median TTFT on two workloads, shared-prefix and multiturn (described below). The actual benchmarking was done with GuideLLM 0.8.1, which we ran on its own node in the cluster.
There are a couple things worth noting about this setup:
This is a relatively modern Qwen model, running as a dense checkpoint on relatively old GPUs. It’s slow: we’re reporting reqs/min, below, a far cry from our usual RPS fare with Linkerd . Similarly, our latency SLOs are relatively high–not appropriate for chat, but plausible for offline agentic workloads.
Different models and different hardware will deliver very different results, so the raw numbers here are less important as they are specific to these choices. The comparative numbers are the important ones.
For the service arm of our test, GuideLLM sent requests straight to the pool’s Kubernetes Service (ClusterIP). kube-proxy, in iptables mode, picks a random pod for each new connection, and GuideLLM opens a new connection for almost every request. Thus, the service arm saw requests distributed randomly across its three replicas.
For the unloaded latency benchmark, we used open-loop Poisson arrivals at a low rate which we empirically determined would not saturate the vLLM workers. For the capacity benchmark, we used closed-loop traffic at a fixed concurrency, i.e. each client sent traffic as quickly as it could. For each condition and trial within these benchmarks, both arms received identical prompts, and in the latency benchmarks requests were sent at identical times. Prompts and prefixes were new in every condition and trial.
Finally, each benchmark covered three types of inference traffic: a control workload, a shared-prefix workload, and a multiturn workload:
We ran each workload with a 60 s warmup period and a 45 s cooldown period, and measured the results over 240 s. The rate column and these windows apply to the latency benchmark; the capacity benchmark used only shared-prefix and multiturn workloads at fixed concurrency (see Results and analysis: Capacity benchmark ).
In total the latency benchmark had 3 workloads × 2 arms × 3 trials = 18 runs. The arms ran one after the other, and order alternated between conditions. A request was measured iff it started inside the measurement window.
Results and analysis: Unloaded latency benchmark
Summary: In the shared-prefix workload, Kinference cut median TTFT by 95% and median E2E latency by 45% . In the multiturn workload, Kinference cut median TTFT by about 65% and median E2E latency by about 32% , versus standard Kubernetes load balancing.
(Confidence intervals are 95% CIs from a paired hierarchical bootstrap with 2,000 resamples.)
In our control conditions, median TTFT is indistinguishable between the arms (6.92 s vs 6.83 s). Without reusable prefixes, Kinference adds no measurable latency, and the two pools perform alike–as we expect. Interestingly, the control p95 is another matter: 21.5 s on the service arm vs 12.3 s with Kinference. This is likely due to Kinference’s usage of vLLM queue size as well as KV cache overlap to predict TTFT–even with prompts that share zero overlap, Kinference can still make intelligent balancing decisions and reduce tail latencies.
In our shared-prefix workload, Kinference cuts median TTFT by about 95% (6.88 s to 0.31 s) and p95 TTFT by about 50% (18.3 s to 9.2 s), a clear improvement over default Kubernetes load balancing.
In our multiturn workload, Kinference cuts median TTFT by about 65% (1.70 s to 0.59 s) and p95 TTFT by about 40% (12.8 s to 7.7 s)–also a clear improvement. Note that the gain here is smaller than on shared-prefix because the first turn of every conversation has nothing to reuse.
Secondary metrics: ITL, E2E latency
Turning to inter-token latency (ITL) and end-to-end (E2E) latency, we see a similar story. In our shared-prefix workload, Kinference cuts median end-to-end latency by about 45% (20.5 s to 11.3 s) and median ITL by about 4% (44 ms to 42 ms). In our multiturn workload, it cuts median end-to-end latency by about 32% (13.4 s to 9.1 s) and median ITL by about 10% (47 ms to 43 ms). On control, neither difference is significant; both CIs reach 1.
Most of the end-to-end gain comes from the shorter TTFT, which is a component of E2E latency. The rest comes from faster token generation: Kinference’s replicas recompute fewer uncached prefixes, so there is less work that competes with producing the remaining tokens.
Output throughput was about equal on every workload, with ratios of 0.98–1.05 per trial, as expected at a fixed arrival rate that both arms keep up with. No request failed or was left unfinished.
Results and analysis: Capacity benchmark
Summary: At peak, Kinference served about 7.4× the service arm’s goodput on shared-prefix and about 3.5× on multiturn. It was able to attain our chosen SLO at concurrency 32 (the highest measured) on our shared-prefix workload, and concurrency 8 on the multiturn workload, whereas the kube-proxy approach was not able to attain the SLO at any concurrency level.
In our unloaded latency benchmark, we sought to measure the effect of Kinference’s KV-cache-aware load balancing in isolation. In our capacity benchmark, we ask a different question: with N users, each sending its next request as soon as the last answer finishes, how many requests per minute can each endpoint serve within a latency SLO?
Our null hypothesis was that at each concurrency, the two endpoints would have the same goodput : the number of requests per minute that met the SLO. Our alternative hypothesis was that Kinference would have higher goodput, and would keep meeting the SLO at higher concurrency. As before, we chose the primary metric up front: the goodput difference (Kinference − service) at each concurrency.
What SLO do we use? Ideally, we’d pick one from first principles–for example, we might not want a user waiting on a live chat bot to have to wait

[truncated]
