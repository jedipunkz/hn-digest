---
source: "https://www.cncf.io/blog/2026/10/19/how-ai-workloads-are-reshaping-kubernetes-autoscaling/"
hn_url: "https://news.ycombinator.com/item?id=50309663"
title: "What's happening in Kubernetes to support AI workloads"
article_title: "How AI workloads are reshaping Kubernetes autoscaling | CNCF"
image: "https://www.cncf.io/wp-content/uploads/2026/10/Blog-Default-50.png"
author: "asaiacai"
captured_at: "2026-10-30T20:27:23Z"
capture_tool: "hn-digest"
hn_id: 50309663
score: 1
comments: 0
posted_at: "2026-10-30T20:24:53Z"
tags:
  - hacker-news
---

# What's happening in Kubernetes to support AI workloads

- HN: [50309663](https://news.ycombinator.com/item?id=50309663)
- Source: [www.cncf.io](https://www.cncf.io/blog/2026/10/19/how-ai-workloads-are-reshaping-kubernetes-autoscaling/)
- Score: 1
- Comments: 0
- Posted: 2026-10-30T20:24:53Z

## Translation

Title: What's happening in Kubernetes to support AI workloads
Article title: How AI workloads are reshaping Kubernetes autoscaling | CNCF
Description: Kubernetes has quietly run much of the world’s containerized software on a division of labor between two control loops that rarely had to talk to each other.

Article text:
Skip to content
Accessibility
help
KubeCon + CloudNativeCon NA 2026 · Nov 9 - 12 · Salt Lake City · REGISTER NOW
Membership Hub – For Current Members Learn about all the benefits of being a CNCF Member
Technical Oversight Committee The TOC defines CNCF’s technical vision and provides experienced technical leadership to the cloud native community
Governing Board The GB is responsible for marketing, business oversight, and budget decisions for CNCF
End User Technical Advisory Board The End User TAB serves as the voice of the end users in CNCF community decisions
Ambassadors Meet our Ambassadors—experienced practitioners passionate about helping others learn about cloud native technologies
Graduated Projects considered stable, widely adopted, and production ready, attracting thousands of contributors
Incubating Projects used successfully in production by a small number users with a healthy pool of contributors
Sandbox Experimental projects not yet widely tested in production on the bleeding edge of technology
Archived Projects that have reached the end of their lifecycle and have become inactive
Project Metrics View metrics of CNCF projects moving through maturity levels
Contribute Join the 150K+ folx in #TeamCloudNative who’ve contributed their expertise to CNCF hosted projects
Services for CNCF Projects CNCF services for our open source projects – from marketing to legal services
Cloud Native Landscape A comprehensive categorical overview of projects and product offerings in the cloud native space
Project Journey Reports Showing how CNCF has impacted the progress and growth of various graduated projects
Project Tools Quick links to tools and resources for your CNCF project
Latest Project Journey Reports
Training Overview Find your path to cloud native success with training and certificates from the pioneer of cloud-native technology
Certifications Get certified by the only authoritative source for cloud-native certification accepted by industry
Courses Learn the basics of cloud native or train for a certification with training courses built in collaboration with the Linux Foundation
Kubestronaut Program Uplevel your Kubernetes skills
Kubernetes Training Partners Find a qualified KCNTP to prepare for your next certification
Certified Kubernetes Software conformance ensures your versions of CNCF projects support the required APIs
Cloud Native Network Function Certification (Beta) CNF Certification ensures applications demonstrate cloud native best practices
Kubernetes Certified Service Provider KCSPs have deep experience helping enterprises successfully adopt cloud native technologies
Enroll your company as a CNCF End User and save more than $10K in training and conference costs
End User Community Join our vendor-neutral community using cloud native technologies to build products and services
Events Meet #TeamCloudNative and CNCF staff at events around the world
Case Studies Read real-world case studies about the impact cloud native projects are having on organizations around the world
Humans of Cloud Native Read stories of amazing individuals and their contributions
The Cloud Native Heroes Challenge Be a cloud native hero! Help us defeat patent trolls to earn swag and prizes
Online Programs Watch our free online programs for the latest insights into cloud native technologies and projects
Community Groups Join #TeamCloudNative at events and meetups near you
Phippy & Friends Phippy explains core cloud native concepts in simple terms through stories perfect for all ages
Cloud Native Glossary Explore cloud native concepts in clear and simple language – no technical knowledge required!
Blog Catch up on the latest happenings and technical insights from #TeamCloudNative
Announcements Media releases and official CNCF announcements
News CNCF projects and #TeamCloudNative in the media
Reports Read transparent, in-depth reports on our organization, events, and projects
Search
CNCF
Blog
/
Member Post
How AI workloads are reshaping Kubernetes autoscaling
Posted on October 19, 2026
by Sajjan Gundapuneedi and Neelendra Bhandari, AWS
CNCF projects highlighted in this post
Kubernetes has quietly run much of the world’s containerized software on a division of labor between two control loops that rarely had to talk to each other. The scheduler places each pod on a node that already exists; the autoscaler adds nodes when pods have nowhere to fit. For a decade that separation was a sound design, because compute was largely interchangeable and either loop could work without knowing much about the other’s decisions. AI workloads, and the accelerators they run on, changed the terms; the clean split still serves most work well, and the pace of AI is what makes its edges worth revisiting.
It surfaces in a situation anyone running GPUs will recognize. A pod that should schedule simply will not, even though the cluster is not truly out of capacity, and the delay can leave costly accelerator capacity unused. The instinct is to assume the region has run dry, and sometimes it has. In other cases, the capacity is available and the pod still stays pending, because the node the autoscaler is willing to add is not one the scheduler is willing to use. The autoscaler answers the question it was built for, whether to add a node, when the real question is whether that node is even the right shape for the work. The rest of this piece examines how the community has closed part of it, how it plans to close the rest, and how the autoscaler is likely to grow into a more scheduler-aware partner.
The hard problems, and how the community is answering them
Several of the challenges accelerators bring share a thread. Each is something both loops now have to understand the same way. Not all are new or unique to AI; some were long familiar to high-performance and batch computing, and accelerators made them everyone’s problem. Three show how far the shared answers have come.
Reservations, capacity committed before the work arrives
This one is the most familiar. Because supply is tight and idle accelerators are expensive, teams may reserve capacity ahead of demand. When a scale-up draws from a reservation, the autoscaler must consider not only whether to grow but also whether the resulting node uses capacity the team has already committed to and is paying for. Reservations are not an autoscaling invention. They are a long-standing capacity practice that an autoscaler may need to account for, because a node can be launched against committed capacity rather than relying on unreserved supply.
Topology, the interconnect between devices, not just their count
Topology is about how devices are wired to one another, not simply how far apart they sit. An accelerator leans on the parts that feed it, including the CPU and memory in its local NUMA domain, the network interface that shares its PCIe path, and the high-bandwidth links to its peer accelerators; it runs best when those pieces share a locality so traffic does not cross a slower hop. That locality is a hierarchy, from devices on the same bus inside a server, to servers in a rack joined by a fast accelerator fabric, to racks linked across a data-center network, each tier a step slower than the one beneath it. The largest models span many servers at once, so “keep these close” reaches well beyond a single node. Historically, Kubernetes commonly represented accelerator resources as counts and lacked a portable, structured way to express many device relationships; a node advertised a number, a pod asked for some, and the scheduler matched the two integers. Dynamic Resource Allocation supplies what was missing. Instead of a bare count, a workload files a structured claim describing what it needs, including a relationship like an accelerator together with the network interface that feeds it. The scheduler reads those claims directly, and DRA-aware autoscalers are starting to use the same information when they simulate capacity, though support still varies by implementation.
The cost of a blind boundary. A DRA-aware autoscaler can account for a topology claim when selecting capacity; without that support it may provision an incompatible node.
Sharing, putting a whole device to work, not part of it
A modern accelerator is powerful enough that handing an entire one to a small job wastes most of it, so a single device can be partitioned into isolated instances that run concurrently. Some runtimes also time-slice devices, but that is a separate execution mechanism and not what consumable capacity models. Partitionable devices let one device be offered as smaller slices on demand, and consumable capacity goes finer, letting several workloads each draw a measured portion, so a request can be satisfied by a fraction of a device rather than requiring an entire device. The scheduler tracks those allocations directly, and DRA-aware autoscalers are starting to use the same model when they weigh a prospective node, though support still varies. Finer slicing can leave leftover fragments, but sensible minimum and step sizes help keep that smaller than the all-or-nothing waste it replaced.
Efficiency moved inside the device. Partitioning splits one device into isolated fixed instances; consumable capacity lets several claims draw measured shares at once.
Topology and sharing matter most for what they share. Together, DRA and its extensions give the scheduler and the autoscaler the same language for hardware, so the two are far less likely to disagree about what a viable node even looks like. That shared vocabulary is real progress, and it is also where the next problem begins.
Why a shared vocabulary is not enough
A common language lets the two loops agree on what a node is; it does not make them act as one, and some problems need exactly that. The clearest is gang scheduling, an all-or-nothing case that is not new but has grown far more acute under AI and tight capacity. A large job needs every pod to be scheduled together or none at all, and this is where shared facts alone fall short.
This part is being addressed, but only on one side. Through workload-aware scheduling , whose Workload and PodGroup APIs are beta upstream and remain disabled by default, the scheduler is learning to treat a group as one all-or-nothing unit, holding its pods until enough can be placed together so a half-placed job no longer strands hardware. The autoscaler was not part of that change. It still reasons pod by pod, so it can add a node or two for a few pending members of a gang and waste that capacity while the job never starts.
A companion mechanism, the provisioning request , is the current bridge. It lets a workload ask the autoscaler for a whole group’s capacity at once rather than pod by pod. With a configured admission check, Kueue reserves quota, creates the request, waits for capacity to be reported ready, and then admits the workload. Atomic scale-up is best effort and depends on provider support. It remains a bolt-on across a seam rather than a shared understanding, and it points at the deeper pattern. A capable autoscaler already reasons a little like a scheduler, predicting where the scheduler would place pods to decide whether a new node would help. The clearest sign is duplicated effort already baked in. The autoscaler simulates the scheduler’s placement to decide what to launch, and once the node exists the scheduler makes its own placement decision, repeating much of the same reasoning.
What the community is building next
This is why the two loops, which still do different jobs, increasingly have to be designed together, and that convergence is becoming an active design discussion. The workload-aware scheduling effort has signaled that its scope is meant to reach beyond the scheduler toward autoscaling over time, and a public proposal in the SIG-Autoscaling community now sketches integration APIs that let the two speak directly. One would let the scheduler see capacity that could be provisioned rathe

[truncated]

## Original Extract

Kubernetes has quietly run much of the world’s containerized software on a division of labor between two control loops that rarely had to talk to each other.

Skip to content
Accessibility
help
KubeCon + CloudNativeCon NA 2026 · Nov 9 - 12 · Salt Lake City · REGISTER NOW
Membership Hub – For Current Members Learn about all the benefits of being a CNCF Member
Technical Oversight Committee The TOC defines CNCF’s technical vision and provides experienced technical leadership to the cloud native community
Governing Board The GB is responsible for marketing, business oversight, and budget decisions for CNCF
End User Technical Advisory Board The End User TAB serves as the voice of the end users in CNCF community decisions
Ambassadors Meet our Ambassadors—experienced practitioners passionate about helping others learn about cloud native technologies
Graduated Projects considered stable, widely adopted, and production ready, attracting thousands of contributors
Incubating Projects used successfully in production by a small number users with a healthy pool of contributors
Sandbox Experimental projects not yet widely tested in production on the bleeding edge of technology
Archived Projects that have reached the end of their lifecycle and have become inactive
Project Metrics View metrics of CNCF projects moving through maturity levels
Contribute Join the 150K+ folx in #TeamCloudNative who’ve contributed their expertise to CNCF hosted projects
Services for CNCF Projects CNCF services for our open source projects – from marketing to legal services
Cloud Native Landscape A comprehensive categorical overview of projects and product offerings in the cloud native space
Project Journey Reports Showing how CNCF has impacted the progress and growth of various graduated projects
Project Tools Quick links to tools and resources for your CNCF project
Latest Project Journey Reports
Training Overview Find your path to cloud native success with training and certificates from the pioneer of cloud-native technology
Certifications Get certified by the only authoritative source for cloud-native certification accepted by industry
Courses Learn the basics of cloud native or train for a certification with training courses built in collaboration with the Linux Foundation
Kubestronaut Program Uplevel your Kubernetes skills
Kubernetes Training Partners Find a qualified KCNTP to prepare for your next certification
Certified Kubernetes Software conformance ensures your versions of CNCF projects support the required APIs
Cloud Native Network Function Certification (Beta) CNF Certification ensures applications demonstrate cloud native best practices
Kubernetes Certified Service Provider KCSPs have deep experience helping enterprises successfully adopt cloud native technologies
Enroll your company as a CNCF End User and save more than $10K in training and conference costs
End User Community Join our vendor-neutral community using cloud native technologies to build products and services
Events Meet #TeamCloudNative and CNCF staff at events around the world
Case Studies Read real-world case studies about the impact cloud native projects are having on organizations around the world
Humans of Cloud Native Read stories of amazing individuals and their contributions
The Cloud Native Heroes Challenge Be a cloud native hero! Help us defeat patent trolls to earn swag and prizes
Online Programs Watch our free online programs for the latest insights into cloud native technologies and projects
Community Groups Join #TeamCloudNative at events and meetups near you
Phippy & Friends Phippy explains core cloud native concepts in simple terms through stories perfect for all ages
Cloud Native Glossary Explore cloud native concepts in clear and simple language – no technical knowledge required!
Blog Catch up on the latest happenings and technical insights from #TeamCloudNative
Announcements Media releases and official CNCF announcements
News CNCF projects and #TeamCloudNative in the media
Reports Read transparent, in-depth reports on our organization, events, and projects
Search
CNCF
Blog
/
Member Post
How AI workloads are reshaping Kubernetes autoscaling
Posted on October 19, 2026
by Sajjan Gundapuneedi and Neelendra Bhandari, AWS
CNCF projects highlighted in this post
Kubernetes has quietly run much of the world’s containerized software on a division of labor between two control loops that rarely had to talk to each other. The scheduler places each pod on a node that already exists; the autoscaler adds nodes when pods have nowhere to fit. For a decade that separation was a sound design, because compute was largely interchangeable and either loop could work without knowing much about the other’s decisions. AI workloads, and the accelerators they run on, changed the terms; the clean split still serves most work well, and the pace of AI is what makes its edges worth revisiting.
It surfaces in a situation anyone running GPUs will recognize. A pod that should schedule simply will not, even though the cluster is not truly out of capacity, and the delay can leave costly accelerator capacity unused. The instinct is to assume the region has run dry, and sometimes it has. In other cases, the capacity is available and the pod still stays pending, because the node the autoscaler is willing to add is not one the scheduler is willing to use. The autoscaler answers the question it was built for, whether to add a node, when the real question is whether that node is even the right shape for the work. The rest of this piece examines how the community has closed part of it, how it plans to close the rest, and how the autoscaler is likely to grow into a more scheduler-aware partner.
The hard problems, and how the community is answering them
Several of the challenges accelerators bring share a thread. Each is something both loops now have to understand the same way. Not all are new or unique to AI; some were long familiar to high-performance and batch computing, and accelerators made them everyone’s problem. Three show how far the shared answers have come.
Reservations, capacity committed before the work arrives
This one is the most familiar. Because supply is tight and idle accelerators are expensive, teams may reserve capacity ahead of demand. When a scale-up draws from a reservation, the autoscaler must consider not only whether to grow but also whether the resulting node uses capacity the team has already committed to and is paying for. Reservations are not an autoscaling invention. They are a long-standing capacity practice that an autoscaler may need to account for, because a node can be launched against committed capacity rather than relying on unreserved supply.
Topology, the interconnect between devices, not just their count
Topology is about how devices are wired to one another, not simply how far apart they sit. An accelerator leans on the parts that feed it, including the CPU and memory in its local NUMA domain, the network interface that shares its PCIe path, and the high-bandwidth links to its peer accelerators; it runs best when those pieces share a locality so traffic does not cross a slower hop. That locality is a hierarchy, from devices on the same bus inside a server, to servers in a rack joined by a fast accelerator fabric, to racks linked across a data-center network, each tier a step slower than the one beneath it. The largest models span many servers at once, so “keep these close” reaches well beyond a single node. Historically, Kubernetes commonly represented accelerator resources as counts and lacked a portable, structured way to express many device relationships; a node advertised a number, a pod asked for some, and the scheduler matched the two integers. Dynamic Resource Allocation supplies what was missing. Instead of a bare count, a workload files a structured claim describing what it needs, including a relationship like an accelerator together with the network interface that feeds it. The scheduler reads those claims directly, and DRA-aware autoscalers are starting to use the same information when they simulate capacity, though support still varies by implementation.
The cost of a blind boundary. A DRA-aware autoscaler can account for a topology claim when selecting capacity; without that support it may provision an incompatible node.
Sharing, putting a whole device to work, not part of it
A modern accelerator is powerful enough that handing an entire one to a small job wastes most of it, so a single device can be partitioned into isolated instances that run concurrently. Some runtimes also time-slice devices, but that is a separate execution mechanism and not what consumable capacity models. Partitionable devices let one device be offered as smaller slices on demand, and consumable capacity goes finer, letting several workloads each draw a measured portion, so a request can be satisfied by a fraction of a device rather than requiring an entire device. The scheduler tracks those allocations directly, and DRA-aware autoscalers are starting to use the same model when they weigh a prospective node, though support still varies. Finer slicing can leave leftover fragments, but sensible minimum and step sizes help keep that smaller than the all-or-nothing waste it replaced.
Efficiency moved inside the device. Partitioning splits one device into isolated fixed instances; consumable capacity lets several claims draw measured shares at once.
Topology and sharing matter most for what they share. Together, DRA and its extensions give the scheduler and the autoscaler the same language for hardware, so the two are far less likely to disagree about what a viable node even looks like. That shared vocabulary is real progress, and it is also where the next problem begins.
Why a shared vocabulary is not enough
A common language lets the two loops agree on what a node is; it does not make them act as one, and some problems need exactly that. The clearest is gang scheduling, an all-or-nothing case that is not new but has grown far more acute under AI and tight capacity. A large job needs every pod to be scheduled together or none at all, and this is where shared facts alone fall short.
This part is being addressed, but only on one side. Through workload-aware scheduling , whose Workload and PodGroup APIs are beta upstream and remain disabled by default, the scheduler is learning to treat a group as one all-or-nothing unit, holding its pods until enough can be placed together so a half-placed job no longer strands hardware. The autoscaler was not part of that change. It still reasons pod by pod, so it can add a node or two for a few pending members of a gang and waste that capacity while the job never starts.
A companion mechanism, the provisioning request , is the current bridge. It lets a workload ask the autoscaler for a whole group’s capacity at once rather than pod by pod. With a configured admission check, Kueue reserves quota, creates the request, waits for capacity to be reported ready, and then admits the workload. Atomic scale-up is best effort and depends on provider support. It remains a bolt-on across a seam rather than a shared understanding, and it points at the deeper pattern. A capable autoscaler already reasons a little like a scheduler, predicting where the scheduler would place pods to decide whether a new node would help. The clearest sign is duplicated effort already baked in. The autoscaler simulates the scheduler’s placement to decide what to launch, and once the node exists the scheduler makes its own placement decision, repeating much of the same reasoning.
What the community is building next
This is why the two loops, which still do different jobs, increasingly have to be designed together, and that convergence is becoming an active design discussion. The workload-aware scheduling effort has signaled that its scope is meant to reach beyond the scheduler toward autoscaling over time, and a public proposal in the SIG-Autoscaling community now sketches integration APIs that let the two speak directly. One would let the scheduler see capacity that could be provisioned rathe

[truncated]
