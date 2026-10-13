---
source: "https://developer.nvidia.com/blog/nvidia-runai-is-now-open-source-as-part-of-nvidia-dsx-ai-factory-platform/"
hn_url: "https://news.ycombinator.com/item?id=50075257"
title: "Nvidia Run:AI Is Now Open Source as Part of Nvidia DSX AI Factory Platform"
article_title: "NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform | NVIDIA Technical Blog"
image: "https://developer-blogs.nvidia.com/wp-content/uploads/2026/01/run-ai.jpg"
author: "wondersky"
captured_at: "2026-10-13T17:58:25Z"
capture_tool: "hn-digest"
hn_id: 50075257
score: 1
comments: 0
posted_at: "2026-10-13T17:17:03Z"
tags:
  - hacker-news
---

# Nvidia Run:AI Is Now Open Source as Part of Nvidia DSX AI Factory Platform

- HN: [50075257](https://news.ycombinator.com/item?id=50075257)
- Source: [developer.nvidia.com](https://developer.nvidia.com/blog/nvidia-runai-is-now-open-source-as-part-of-nvidia-dsx-ai-factory-platform/)
- Score: 1
- Comments: 0
- Posted: 2026-10-13T17:17:03Z

## Translation

Title: Nvidia Run:AI Is Now Open Source as Part of Nvidia DSX AI Factory Platform
Article title: NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform | NVIDIA Technical Blog
Description: Today, NVIDIA is making the full NVIDIA Run:ai platform available as open source software, giving customers, developers, and ecosystem partners access to…

Article text:
NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Data Center / Cloud
NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform
Accelerate innovation across the DSX ecosystem and bring AI factory solutions to market faster with the open foundation for AI workload and GPU orchestration.
NVIDIA is releasing the full NVIDIA Run:ai platform as open source software, expanding the NVIDIA DSX AI Factory Platform with a foundation for AI workload and GPU orchestration.
The release includes four modular components that can be adopted independently: KAI Scheduler for GPU-aware Kubernetes scheduling, GPU Fractioning Operator for CUDA-native GPU sharing, KAI Resource Management for organizational quota and hierarchy, and Workload Map for a consistent workload-management contract.
Eight NVIDIA DSX ecosystem partners including Crusoe, Mirantis, Nebius, OpenNebula Systems, Rafay, Saturn Cloud, Spectro Cloud, and vCluster are integrating the open source platform into their offerings.
Related open source projects Model Streamer and Grove complement the platform by accelerating model loading and orchestrating multicomponent inference workloads.
Explore the NVIDIA Run:ai full platform repository to inspect code and begin adoption.
Review the KAI Scheduler project for GPU-aware Kubernetes scheduling capabilities.
Try the GPU Fractioning Operator to enable CUDA-native GPU fractions on Kubernetes clusters.
Today, NVIDIA is making the full NVIDIA Run:ai platform available as open source software, giving customers, developers, and ecosystem partners access to production-grade AI workload and GPU orchestration capabilities. This release expands the NVIDIA DSX AI Factory Platform by establishing NVIDIA Run:ai as an open foundation that enables the broader ecosystem to bring differentiated solutions to market faster in response to evolving customer needs.
Over the past several years, NVIDIA Run:ai has helped organizations run AI workloads on Kubernetes at scale by dynamically allocating GPU resources and managing workloads across the full lifecycle, from data processing and training to inference.
By open sourcing NVIDIA Run:ai platform capabilities as part of NVIDIA DSX AI Factory, it is now possible for every NVIDIA DSX ecosystem partner and AI cloud to increase workload throughput, improve GPU utilization, and operate shared AI infrastructure more efficiently and effectively.
This post details the open source release, which includes the complete NVIDIA Run:ai platform and four components that can also be adopted independently: KAI Scheduler, GPU Fractioning Operator , KAI Resource Management , Workload Map . Together with additional platform code, these components form the integrated NVIDIA Run:ai platform. Platform teams and ecosystem partners can adopt NVIDIA Run:ai as an integrated whole, build on it for their own products and services, or integrate individual components based on their requirements.
How to adopt open source NVIDIA Run:ai
NVIDIA Run:ai can be adopted directly as open-source software or through NVIDIA DSX ecosystem partner offerings.
Deploy the open source platform directly
The complete source code for NVIDIA Run:ai is available in the NVIDIA Run:ai repository within the NVIDIA DSX AI Factory GitHub organization . Teams can inspect the code, contribute to its development, build on top of the platform or adopt its modular components independently.
The complete open source NVIDIA Run:ai platform is also available on NVIDIA NGC as a Community Product, with a Helm chart for installing the integrated stack.
Note: The NVIDIA Run:ai source code is governed by its applicable open source licenses. Downloading the Community Product from NGC requires acceptance of the applicable NVIDIA Software License Agreement and Product-Specific Terms for NVIDIA AI Products. No paid NVIDIA Run:ai subscription or NVIDIA support entitlement is required for the Community Product, and NVIDIA enterprise support is not included.
Adopt open source NVIDIA Run:ai through NVIDIA DSX ecosystem partners
Open source NVIDIA Run:ai gives NVIDIA DSX ecosystem partners direct access to AI workload and GPU orchestration innovations they can use in their own NVIDIA DSX AI factory solutions. By adopting the complete platform or integrating individual components, partners can bring those innovations to market faster and respond more quickly to evolving customer and operational requirements.
Crusoe integrates NVIDIA Run:ai with Crusoe Managed Kubernetes, combining GPU workload orchestration with managed Kubernetes and GPU infrastructure.
Mirantis extends k0rdent AI with NVIDIA Run:ai to deliver a lifecycle-managed, Metal-to-Model AI factory stack that automates infrastructure provisioning and maximizes GPU utilization.
Nebius offers NVIDIA Run:ai as a one-click application deployment in Nebius AI Cloud, giving customers a streamlined way to deploy and operate the complete platform on Nebius Managed Service for Kubernetes.
OpenNebula Systems integrates NVIDIA Run:ai with the OpenNebula AI Cloud Platform through OneKS, combining GPU workload orchestration with infrastructure provisioning and lifecycle management for open and sovereign AI clouds.
Rafay is expanding its NVIDIA Run:ai integration to deliver it as a fully managed service. Rafay automates the entire stack, from GPU infrastructure and clusters to the NVIDIA Run:ai control plane and tenants, so AI cloud providers and enterprises can offer it through their own self-service catalogs.
Saturn Cloud integrates NVIDIA Run:ai to help AI cloud operators launch white-label AI inference services through its Token Factory platform, combining GPU orchestration with dedicated and multi-tenant model serving, per-token metering, and automated billing.
Spectro Cloud integrates NVIDIA Run:ai into PaletteAI, enabling customers to deploy it through PaletteAI Studio and manage it with the same lifecycle automation used across their AI infrastructure.
vCluster provides NVIDIA Run:ai through pre-validated deployment templates that stand up a complete environment per tenant, allowing operators to strengthen tenant boundaries while NVIDIA Run:ai continues to schedule across one GPU fleet.
For NVIDIA Enterprise AI Factory deployments, NVIDIA Run:ai continues to be available through NVIDIA-supported commercial offerings including as a standalone product, through NVIDIA Mission Control and NVIDIA AI Enterprise, and select cloud marketplaces. Availability, licensing, and support terms vary by offering.
For teams that already have a platform and need to fill a specific capability gap, the following sections describe when and how to use each modular component independently.
Modular open source components of the NVIDIA Run:ai platform
This section details the four components included in the open source release: KAI Scheduler, GPU Fractioning Operator, KAI Resource Management, and Workload Map. With additional platform code, these components form the integrated NVIDIA Run:ai platform. You can adopt NVIDIA Run:ai as an integrated whole or integrate individual components based on your unique requirements.
KAI Scheduler: The scheduling engine behind NVIDIA Run:ai
KAI Scheduler is the Kubernetes-native scheduler that powers the NVIDIA Run:ai platform. NVIDIA released KAI Scheduler as open source software under the Apache 2.0 license and contributed it to CNCF as a Sandbox project. KAI remains the upstream engine behind NVIDIA Run:ai, and all scheduler development happens in the community repository, so every improvement flows through the same codebase. NVIDIA continues to support the project’s maintainers and work with the broader ecosystem to advance AI workload scheduling at scale.
KAI Scheduler is purpose-built for GPU workloads at scale, coordinating:
Gang scheduling with prescheduling in-memory simulation, so multipod jobs are placed all-or-nothing and aren’t evicted needlessly
Fair-share (including time-based fair-share) and quota allocation that keeps resource shares balanced across teams
Topology awareness and bin-packing at the scheduling-decision level
GPU sharing, which allows multiple workloads to share single or multiple GPUs (MIG, MPS, time sharing, HaMi, CUDA 13.4 MPS fractioning)
Dynamic Resource Allocation (DRA) support
To learn more about key features, visit the kai-scheduler/KAI-Scheduler GitHub repo.
Use KAI Scheduler for diverse AI workloads that require fair access to cluster resources and topology-aware placement. It can be deployed in an existing cluster alongside the current software stack without requiring the complete NVIDIA Run:ai platform.
Follow the project, ask questions in the #kai-scheduler channel on CNCF Slack, and contribute on the kai-scheduler/KAI-Scheduler GitHub repo.
GPU Fractioning: An operator for GPU fractions on Kubernetes
The NVIDIA Run:ai platform has long supported fractional GPU allocation, enabling multiple workloads to share a GPU with hard memory isolation. Fractions can be requested on demand and set up on the fly, with no per-node setup from an administrator in advance. It is one of the most widely used capabilities among ML teams running workloads that do not require the full memory of a GPU, including interactive notebooks and inference for smaller models. NVIDIA Run:ai pioneered that experience with a proprietary implementation.
GPU Fractioning Operator brings CUDA-native GPU fractioning to Kubernetes and the proprietary implementation has been retired in favor of the native capability available in CUDA 13.4 . The CUDA driver enforces a hard limit on GPU memory for each workload, and each workload sees only the GPU memory allocated to its fraction. What is still needed is a clean way to consume it on Kubernetes. This is what GPU Fractioning Operator provides: an operator that enables CUDA GPU fractions on a cluster, along with a metrics exporter for per-pod visibility into fractions in use. It is a Kubernetes integration layer that allows workloads to use CUDA-native fractioning, and it is scheduler-agnostic.
Use the GPU Fractioning Operator for workloads that can share underutilized GPU capacity, including notebooks, development sessions, and smaller models. GPU Fractioning Operator integrates with KAI Scheduler to translate available GPU capacity into an actual placement decision, while remaining compatible with other schedulers.
GPU Fractioning Operator brings CUDA 13.4-native GPU fractioning to Kubernetes and integrates with KAI Scheduler. To get started, refer to the GPU Fractioning Operator documentation .
KAI Resource Management: The organizational layer
While the scheduler decides how work is placed, KAI Resource Management governs who gets what:
Projects and departments: The organizational structure, or “org layer,” that maps resources to the teams and initiatives that consume them.
Resource management with an organizational hierarchy: The ability to allocate, guarantee, and share GPU capacity across that org structure, with fair-share and quota semantics.
Node pools: Grouping and targeting of heterogeneous hardware so workloads land on the right resources.
Together, these let platform teams move beyond raw Kubernetes scheduling and manage GPUs the way organizations actually operate: by team, by department, by initiative, and by resource guarantee.
Use KAI Resource Management for heterogeneous clusters shared by multiple teams that require defined shares of capacity. For example, a platform team could guarantee 50% of the cluster to research while preventing any one team from monopolizing the remaining capacity. If you are already tracking quota entitlements in a spreadsheet or hand-maintaining namespaces per team, this is the layer that replaces it. It can be combined with KAI Scheduler in your clusters.
Read more about how these concepts map to real deployments in the KAI Resource Management documentation .
Workload Map: A wo

[truncated]

## Original Extract

Today, NVIDIA is making the full NVIDIA Run:ai platform available as open source software, giving customers, developers, and ecosystem partners access to…

NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Data Center / Cloud
NVIDIA Run:ai Is Now Open Source as Part of NVIDIA DSX AI Factory Platform
Accelerate innovation across the DSX ecosystem and bring AI factory solutions to market faster with the open foundation for AI workload and GPU orchestration.
NVIDIA is releasing the full NVIDIA Run:ai platform as open source software, expanding the NVIDIA DSX AI Factory Platform with a foundation for AI workload and GPU orchestration.
The release includes four modular components that can be adopted independently: KAI Scheduler for GPU-aware Kubernetes scheduling, GPU Fractioning Operator for CUDA-native GPU sharing, KAI Resource Management for organizational quota and hierarchy, and Workload Map for a consistent workload-management contract.
Eight NVIDIA DSX ecosystem partners including Crusoe, Mirantis, Nebius, OpenNebula Systems, Rafay, Saturn Cloud, Spectro Cloud, and vCluster are integrating the open source platform into their offerings.
Related open source projects Model Streamer and Grove complement the platform by accelerating model loading and orchestrating multicomponent inference workloads.
Explore the NVIDIA Run:ai full platform repository to inspect code and begin adoption.
Review the KAI Scheduler project for GPU-aware Kubernetes scheduling capabilities.
Try the GPU Fractioning Operator to enable CUDA-native GPU fractions on Kubernetes clusters.
Today, NVIDIA is making the full NVIDIA Run:ai platform available as open source software, giving customers, developers, and ecosystem partners access to production-grade AI workload and GPU orchestration capabilities. This release expands the NVIDIA DSX AI Factory Platform by establishing NVIDIA Run:ai as an open foundation that enables the broader ecosystem to bring differentiated solutions to market faster in response to evolving customer needs.
Over the past several years, NVIDIA Run:ai has helped organizations run AI workloads on Kubernetes at scale by dynamically allocating GPU resources and managing workloads across the full lifecycle, from data processing and training to inference.
By open sourcing NVIDIA Run:ai platform capabilities as part of NVIDIA DSX AI Factory, it is now possible for every NVIDIA DSX ecosystem partner and AI cloud to increase workload throughput, improve GPU utilization, and operate shared AI infrastructure more efficiently and effectively.
This post details the open source release, which includes the complete NVIDIA Run:ai platform and four components that can also be adopted independently: KAI Scheduler, GPU Fractioning Operator , KAI Resource Management , Workload Map . Together with additional platform code, these components form the integrated NVIDIA Run:ai platform. Platform teams and ecosystem partners can adopt NVIDIA Run:ai as an integrated whole, build on it for their own products and services, or integrate individual components based on their requirements.
How to adopt open source NVIDIA Run:ai
NVIDIA Run:ai can be adopted directly as open-source software or through NVIDIA DSX ecosystem partner offerings.
Deploy the open source platform directly
The complete source code for NVIDIA Run:ai is available in the NVIDIA Run:ai repository within the NVIDIA DSX AI Factory GitHub organization . Teams can inspect the code, contribute to its development, build on top of the platform or adopt its modular components independently.
The complete open source NVIDIA Run:ai platform is also available on NVIDIA NGC as a Community Product, with a Helm chart for installing the integrated stack.
Note: The NVIDIA Run:ai source code is governed by its applicable open source licenses. Downloading the Community Product from NGC requires acceptance of the applicable NVIDIA Software License Agreement and Product-Specific Terms for NVIDIA AI Products. No paid NVIDIA Run:ai subscription or NVIDIA support entitlement is required for the Community Product, and NVIDIA enterprise support is not included.
Adopt open source NVIDIA Run:ai through NVIDIA DSX ecosystem partners
Open source NVIDIA Run:ai gives NVIDIA DSX ecosystem partners direct access to AI workload and GPU orchestration innovations they can use in their own NVIDIA DSX AI factory solutions. By adopting the complete platform or integrating individual components, partners can bring those innovations to market faster and respond more quickly to evolving customer and operational requirements.
Crusoe integrates NVIDIA Run:ai with Crusoe Managed Kubernetes, combining GPU workload orchestration with managed Kubernetes and GPU infrastructure.
Mirantis extends k0rdent AI with NVIDIA Run:ai to deliver a lifecycle-managed, Metal-to-Model AI factory stack that automates infrastructure provisioning and maximizes GPU utilization.
Nebius offers NVIDIA Run:ai as a one-click application deployment in Nebius AI Cloud, giving customers a streamlined way to deploy and operate the complete platform on Nebius Managed Service for Kubernetes.
OpenNebula Systems integrates NVIDIA Run:ai with the OpenNebula AI Cloud Platform through OneKS, combining GPU workload orchestration with infrastructure provisioning and lifecycle management for open and sovereign AI clouds.
Rafay is expanding its NVIDIA Run:ai integration to deliver it as a fully managed service. Rafay automates the entire stack, from GPU infrastructure and clusters to the NVIDIA Run:ai control plane and tenants, so AI cloud providers and enterprises can offer it through their own self-service catalogs.
Saturn Cloud integrates NVIDIA Run:ai to help AI cloud operators launch white-label AI inference services through its Token Factory platform, combining GPU orchestration with dedicated and multi-tenant model serving, per-token metering, and automated billing.
Spectro Cloud integrates NVIDIA Run:ai into PaletteAI, enabling customers to deploy it through PaletteAI Studio and manage it with the same lifecycle automation used across their AI infrastructure.
vCluster provides NVIDIA Run:ai through pre-validated deployment templates that stand up a complete environment per tenant, allowing operators to strengthen tenant boundaries while NVIDIA Run:ai continues to schedule across one GPU fleet.
For NVIDIA Enterprise AI Factory deployments, NVIDIA Run:ai continues to be available through NVIDIA-supported commercial offerings including as a standalone product, through NVIDIA Mission Control and NVIDIA AI Enterprise, and select cloud marketplaces. Availability, licensing, and support terms vary by offering.
For teams that already have a platform and need to fill a specific capability gap, the following sections describe when and how to use each modular component independently.
Modular open source components of the NVIDIA Run:ai platform
This section details the four components included in the open source release: KAI Scheduler, GPU Fractioning Operator, KAI Resource Management, and Workload Map. With additional platform code, these components form the integrated NVIDIA Run:ai platform. You can adopt NVIDIA Run:ai as an integrated whole or integrate individual components based on your unique requirements.
KAI Scheduler: The scheduling engine behind NVIDIA Run:ai
KAI Scheduler is the Kubernetes-native scheduler that powers the NVIDIA Run:ai platform. NVIDIA released KAI Scheduler as open source software under the Apache 2.0 license and contributed it to CNCF as a Sandbox project. KAI remains the upstream engine behind NVIDIA Run:ai, and all scheduler development happens in the community repository, so every improvement flows through the same codebase. NVIDIA continues to support the project’s maintainers and work with the broader ecosystem to advance AI workload scheduling at scale.
KAI Scheduler is purpose-built for GPU workloads at scale, coordinating:
Gang scheduling with prescheduling in-memory simulation, so multipod jobs are placed all-or-nothing and aren’t evicted needlessly
Fair-share (including time-based fair-share) and quota allocation that keeps resource shares balanced across teams
Topology awareness and bin-packing at the scheduling-decision level
GPU sharing, which allows multiple workloads to share single or multiple GPUs (MIG, MPS, time sharing, HaMi, CUDA 13.4 MPS fractioning)
Dynamic Resource Allocation (DRA) support
To learn more about key features, visit the kai-scheduler/KAI-Scheduler GitHub repo.
Use KAI Scheduler for diverse AI workloads that require fair access to cluster resources and topology-aware placement. It can be deployed in an existing cluster alongside the current software stack without requiring the complete NVIDIA Run:ai platform.
Follow the project, ask questions in the #kai-scheduler channel on CNCF Slack, and contribute on the kai-scheduler/KAI-Scheduler GitHub repo.
GPU Fractioning: An operator for GPU fractions on Kubernetes
The NVIDIA Run:ai platform has long supported fractional GPU allocation, enabling multiple workloads to share a GPU with hard memory isolation. Fractions can be requested on demand and set up on the fly, with no per-node setup from an administrator in advance. It is one of the most widely used capabilities among ML teams running workloads that do not require the full memory of a GPU, including interactive notebooks and inference for smaller models. NVIDIA Run:ai pioneered that experience with a proprietary implementation.
GPU Fractioning Operator brings CUDA-native GPU fractioning to Kubernetes and the proprietary implementation has been retired in favor of the native capability available in CUDA 13.4 . The CUDA driver enforces a hard limit on GPU memory for each workload, and each workload sees only the GPU memory allocated to its fraction. What is still needed is a clean way to consume it on Kubernetes. This is what GPU Fractioning Operator provides: an operator that enables CUDA GPU fractions on a cluster, along with a metrics exporter for per-pod visibility into fractions in use. It is a Kubernetes integration layer that allows workloads to use CUDA-native fractioning, and it is scheduler-agnostic.
Use the GPU Fractioning Operator for workloads that can share underutilized GPU capacity, including notebooks, development sessions, and smaller models. GPU Fractioning Operator integrates with KAI Scheduler to translate available GPU capacity into an actual placement decision, while remaining compatible with other schedulers.
GPU Fractioning Operator brings CUDA 13.4-native GPU fractioning to Kubernetes and integrates with KAI Scheduler. To get started, refer to the GPU Fractioning Operator documentation .
KAI Resource Management: The organizational layer
While the scheduler decides how work is placed, KAI Resource Management governs who gets what:
Projects and departments: The organizational structure, or “org layer,” that maps resources to the teams and initiatives that consume them.
Resource management with an organizational hierarchy: The ability to allocate, guarantee, and share GPU capacity across that org structure, with fair-share and quota semantics.
Node pools: Grouping and targeting of heterogeneous hardware so workloads land on the right resources.
Together, these let platform teams move beyond raw Kubernetes scheduling and manage GPUs the way organizations actually operate: by team, by department, by initiative, and by resource guarantee.
Use KAI Resource Management for heterogeneous clusters shared by multiple teams that require defined shares of capacity. For example, a platform team could guarantee 50% of the cluster to research while preventing any one team from monopolizing the remaining capacity. If you are already tracking quota entitlements in a spreadsheet or hand-maintaining namespaces per team, this is the layer that replaces it. It can be combined with KAI Scheduler in your clusters.
Read more about how these concepts map to real deployments in the KAI Resource Management documentation .
Workload Map: A wo

[truncated]
