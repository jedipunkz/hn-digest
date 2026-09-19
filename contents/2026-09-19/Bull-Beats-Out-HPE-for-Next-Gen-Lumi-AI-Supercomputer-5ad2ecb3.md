---
source: "https://www.nextplatform.com/hpc/2026/09/17/bull-beats-out-hpe-for-next-gen-lumi-ai-supercomputer/5297292"
hn_url: "https://news.ycombinator.com/item?id=49765950"
title: "Bull Beats Out HPE for Next-Gen Lumi AI Supercomputer"
article_title: "Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer"
image: "https://image.nextplatform.com/5297294.jpg?imageId=5297294&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "rbanffy"
captured_at: "2026-09-19T12:55:17Z"
capture_tool: "hn-digest"
hn_id: 49765950
score: 1
comments: 0
posted_at: "2026-09-19T12:19:20Z"
tags:
  - hacker-news
---

# Bull Beats Out HPE for Next-Gen Lumi AI Supercomputer

- HN: [49765950](https://news.ycombinator.com/item?id=49765950)
- Source: [www.nextplatform.com](https://www.nextplatform.com/hpc/2026/09/17/bull-beats-out-hpe-for-next-gen-lumi-ai-supercomputer/5297292)
- Score: 1
- Comments: 0
- Posted: 2026-09-19T12:19:20Z

## Translation

Title: Bull Beats Out HPE for Next-Gen Lumi AI Supercomputer
Article title: Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer
Description: In many ways, Europe is more traditional than America, and this is particularly so when it comes to ...

Article text:
Jump to main content
Search
More topics
All the latest news, from all sections
Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer
In many ways, Europe is more traditional than America, and this is particularly so when it comes to GenAI infrastructure. While all of the national supercomputing labs have built systems that can do both traditional HPC simulation and modeling, in Europe AI efforts to custom tune GenAI and traditional machine learning models are focused on the national labs, not in AI model builders.
Of course, that is because the major AI model builders are based on the United States and China. But no matter. Our point is that with the national labs leading both HPC and GenAI efforts across Europe, there are grownups in the room who are applying GenAI to scientific challenges funded by governments while also providing compute allocations to businesses in Europe that are trying to build AI models as well as use HPC tools to design products.
The national labs are not financed by private equity or the makers of GPU accelerators or the big clouds, so they have to make do with less – two or three orders of magnitude less money and less compute. But that’s fine. They are still getting more compute capacity than they ever had access to, and they can go to sleep at night knowing that they are unequivocally trying to do good in the world.
So it is with the upcoming next generation Lumi-AI supercomputer – which will be one component in what is being called the Lumi AI factory to borrow a term coined by Nvidia – that is being funded by the EuroHPC Joint Undertaking and that will be installed at the Center for Scientific Computation (CSC) facilities in Kajaani, Finland. (Lumi is Finnish for “snow on the ground,” but the mascot of the machine is a white wolf. They could have named the machine “susi,” which is wolf in Finnish. . . .)
The original Lumi hybrid supercomputer cluster, a mix of CPU-only, CPU-GPU, big memory CPU, object storage, Lustre parallel file system storage, and flash storage nodes, was first unveiled way back in October 2020 . That hybrid machine was expected to have a total of 552 petaflops of aggregate computing at FP64 precision. The Lumi-C partition, with over 200,000 CPUs cores based on AMD Epyc processors would have about 2 petaflops of FP64 compute, with the Lumi-G partition based on AMD Instinct processors having about 550 petaflops of double precision floating point processing. The Lumi cluster of clusters would also have 30 PB of Ceph object storage, 80 PB of Lustre storage, and 7 PB of flash storage. The nodes within each partition as well as linking those partitions together was based on the “Rosetta” 200 Gb/sec Slingshot-10 interconnect, with Hewlett Packard Enterprise as the main contractor for the system.
The whole shebang was given a $237 million budget by EuroHPC, which worked out to about $431 per teraflops. The Lumi-C partition was delivered in the summer of 2021, and the Lumi-G GPU partition was delivered at the end of 2021. These partitions were operational in June 2022. The Lumi-G partition was ranked as the number 11 supercomputer in the Top500 rankings in June of this year, so it is no slouch even though it was four years old at the time.
The Lumi-C partition has 1,536 nodes, each of which are equipped with a pair of 64-core “Milan” Epyc 7763 processors spinning at 2.45 GHz. That comes to a total of 196,860 cores, to be precise. This partition is ranked number 262 on the June Top500 rankings, with a peak theoretical performance of 7.29 petaflops and a sustained Linpack performance of 6.3 petaflops, which is a very respectable 82.6 percent computational efficiency in a 1.22 megawatt power envelope. That works out to 5.18 gigaflops per watt.
This machine uses a 200 Gb/sec Slingshot 11 interconnect to lash the nodes together.
The Lumi-G partition has 2,978 nodes, with each one having a single 64-core “Trento” Epyc 7A53 custom processor with four “Aldebaran” MI250X GPU accelerators hanging off that. (The Trento chip was custom built for the “Frontier” supercomputer at Oak Ridge National Laboratory in the US, and provided memory coherence across AMD CPUs and GPUs before that was officially commercially available.) That is a total of 2,978 Trento CPUs, with 190,592 cores, and 11,912 Aldebaran GPUs, with a total of 2.62 million compute units and 167.7 million streaming processors. The nodes in the Lumi-G partition are linked by a 200 Gb/sec Slingshot-11 interconnect. This partition of the Lumi machine has a peak theoretical performance of 531.5 petaflops at FP64 precision, and rated 379.7 petaflops on the Linpack test, which is a computational efficiency of 71.4 percent. This partition weighs in at 7.1 megawatts and delivers 53.43 gigaflops per watt on Linpack, which is more than 10X better power efficiency than the CPU-only partition.
What I know about the future Lumi-AI partition, which is really an upgrade to the Lumi-G partition that I assume will stay in production given the value of those AMD GPUs and the scarcity of GPU capacity on the world, is that it will offer 10X the performance on AI workloads and “around 2X” the performance on FP64 HPC workloads.
It will not be hard to beat the AI performance of the MI250X, which supported INT4 and INT8 data formats as well as FP16 formats, but the performance was the same for all three. So INT4 and INT8 were not precisely native on the Aldebaran GPUs, but grafted on top of FP16, which was fine for that time. But with the Altair GPUs that are being packaged as the MI430X aimed at both HPC and AI workloads (unlike the top-end Altair MI455X, which is aimed at AI inference and training), FP4 and FP8 processing is native, so you get a multiplicative effect as you trim down from FP16. So 4X of that 10X is coming from the shift from FP16 down to FP4. So that leaves 2.5X more raw performance at FP16 comparing the Altair GPU used in the MI430X to the Aldebaran GPU used in the MI250X.
This means, of course, that the node count in the Lumi-AI partition should be a lot smaller, even if the devices themselves will certainly burn more electricity and therefore generate more heat.
Let’s figure out that node count on Lumi-AI.
The MI250X is rated at 47.9 teraflops on its vector units and 95.7 teraflops on its matrix units. The MI430X has a peak of 288 teraflops on FP64 workloads, and AMD was not precise about this being on vector or matrix units. It doesn’t have to be because the MI430X delivers 288 teraflops on both vector and matrix units. If you do the math on the Lumi-G partition, it looks like CSC Finland was counting the vector FP64 performance (47.9 teraflops times four GPUs per node times 2,978 nodes is a peak of 570.6 petaflops. (We don’t know why this doesn’t match the 531.5 Rpeak on the Linpack test.)
If you want 2X performance on the GPUs at FP64 vector for Lumi-AI, then you only need 3,962 MI430X GPUs, and that works out to a mere 991 nodes, a 66.7 percent reduction in node count for the AI partition. (This breaks the exascale barrier, with 1.14 exaflops peak, which is psychologically important. Each node will, however, burn about 2X the energy, at 5 kilowatts for the Venice-Altair 1x4 board compared to 2.5 kilowatts for the Trento-Aldebaran 1x4 board. And that means the Lumi-AI partition will probably consume about 5 megawatts, assuming 991 nodes, which is a 33 percent reduction compared to the Lumi-G partition.
The weird thing is that when we do the math on this number of nodes to derive the AI performance at FP64, we only get an 8X performance increase. You need 1,240 nodes with four MI430X GPUs in Lumi-AI to get a 10X increase over Lumi-G at 4-bit precision. That might mean that CSC Finland will get an extra 25 percent more FP64 vector performance than it is talking about.
I’m going with 1,240 nodes with 4,960 MI430X GPUs for my guess, which is a 58.4 percent reduction in GPU partition node count and about 6.2 megawatts of system power for Lumi-AI. Both Bull, the primary contractor for the Lumi-AI partition, and CSC Finland were clear that AI performance would be 10X higher and HPC performance would be “about 2X.” Why HPC centers are imprecise in their statements is a mystery, and it is counterintuitive for an organization that lives by precision. Go figure.
Using a 256-core “Venice” Epyc 9006 processor, the Lumi-AI partition will have 317,440 cores, which is an appreciable number of compute units to help orchestrate those GPUs and which is 66.7 percent higher than the core count in the current Lumi-G partition.
Given all of this, and a price tag of €387.8 million for the Lumi-AI partition, converting that to US dollars yields a cost of $445.1 million and against a peak FP64 performance of 1.43 exaflops for the Lumi-AI partition, you get a $312 per teraflops at peak. This is a 27.6 percent lower cost per teraflops than the $431 per teraflops peak that the entire Lumi cluster cost when it was installed in 2021. It is not known if the Lumi CPU-only partition and the various storage clusters will also be upgraded under the Lumi-AI contract, but we presume this to be the case. Which makes this an apples-to-apples overall system comparison.
That CSC Finland and its EuroHPC funding arm went with Bull over HPE in this deal is not a big surprise, really. EuroHPC wants to buy European as much as possible. Moreover, while Slingshot is an Ethernet variant that can be possibly modified to adhere to Ultra Ethernet Consortium hardware, Bulls BXI protocol is being adapted to run on UEC hardware and make use of its various features. Which means Bull customers using the BullSequana XH3500 systems will have a choice of switchery when they want to do scale out clustering.
The Lumi-AI GPU partition is expected to be installed in the second half of 2027 in a new datacenter located in Kajaani. Finland, Czechia, Denmark, Estonia, Norway, and Poland are kicking in money alongside the European Union’s EuroHPC JU. The precise amounts everyone is paying was not disclosed.
The Lumi AI factory will not only include this new AMD Venice-Altair cluster, but will also have a quantum computing element dubbed Lumi-IQ. The quantum system is being built by IQM Quantum Computers, a Finnish company that has raised €600 million in venture capital so far and that has aspirations to be Europe’s top quantum computing supplier. The company is preparing to go public and has a valuation of $1.8 billion at this time.
The plan is to installed a Halocene H4 system at CSC Finland in 2027, which will have 150 physical cubits and support 5 logical qubits. Upgrades will be done in 2028 to improve error correction and to improve the reliability of logical qubits. In 2029, CSC Finland will get an upgrade to a Halocene H5 quantum computer, which will have an unknown number of physical qubits but will deliver at least 9 logical qubits. That should imply 270 physical qubits to deliver those local qubits.
Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer
Infleqtion And Nvidia Reduce Physical-to-Logical Qubit Ratio in Quantum Computers
How To Smash The Memory Wall Plaguing High Performance Systems
Operationalize AI At Scale With HPE And Nvidia
Despite Hefty Capital Investments, The Cloud Starts Paying Off For Oracle
Broadcom Rides Rocketing Trend For Custom AI Accelerators
Startup d-Matrix Will Pair Its Raptor Memory-Based XPU To Nvidia Rackscale Iron
HPE Sees No Customer Pushback On Rising Datacenter Equipment Costs
COMPUTE
Dell Says AI Will Drive 75 Percent Of Datacenter Demand By 2030
Nvidia Expands Its Open Source AI Presence With $12.9 Billion Hugging Face Buy
Optics Still Driving Marvell’s AI Business More Than Custom Chips
VMware Intros Private AI Cloud, AI Factory As Workloads Shift To On-Prem
In The Long Run, Nvidia NVSwitch Is The InfiniBand Of Scale Up AI Networks
Peeling Apart That Supposed $120 Billion Chip Deal Google Inked With Marvell
COMPUTE
The Network Is More Of Nvidia’s Computer Than It Ever Was At Sun
STORE
You Probably

[truncated]

## Original Extract

In many ways, Europe is more traditional than America, and this is particularly so when it comes to ...

Jump to main content
Search
More topics
All the latest news, from all sections
Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer
In many ways, Europe is more traditional than America, and this is particularly so when it comes to GenAI infrastructure. While all of the national supercomputing labs have built systems that can do both traditional HPC simulation and modeling, in Europe AI efforts to custom tune GenAI and traditional machine learning models are focused on the national labs, not in AI model builders.
Of course, that is because the major AI model builders are based on the United States and China. But no matter. Our point is that with the national labs leading both HPC and GenAI efforts across Europe, there are grownups in the room who are applying GenAI to scientific challenges funded by governments while also providing compute allocations to businesses in Europe that are trying to build AI models as well as use HPC tools to design products.
The national labs are not financed by private equity or the makers of GPU accelerators or the big clouds, so they have to make do with less – two or three orders of magnitude less money and less compute. But that’s fine. They are still getting more compute capacity than they ever had access to, and they can go to sleep at night knowing that they are unequivocally trying to do good in the world.
So it is with the upcoming next generation Lumi-AI supercomputer – which will be one component in what is being called the Lumi AI factory to borrow a term coined by Nvidia – that is being funded by the EuroHPC Joint Undertaking and that will be installed at the Center for Scientific Computation (CSC) facilities in Kajaani, Finland. (Lumi is Finnish for “snow on the ground,” but the mascot of the machine is a white wolf. They could have named the machine “susi,” which is wolf in Finnish. . . .)
The original Lumi hybrid supercomputer cluster, a mix of CPU-only, CPU-GPU, big memory CPU, object storage, Lustre parallel file system storage, and flash storage nodes, was first unveiled way back in October 2020 . That hybrid machine was expected to have a total of 552 petaflops of aggregate computing at FP64 precision. The Lumi-C partition, with over 200,000 CPUs cores based on AMD Epyc processors would have about 2 petaflops of FP64 compute, with the Lumi-G partition based on AMD Instinct processors having about 550 petaflops of double precision floating point processing. The Lumi cluster of clusters would also have 30 PB of Ceph object storage, 80 PB of Lustre storage, and 7 PB of flash storage. The nodes within each partition as well as linking those partitions together was based on the “Rosetta” 200 Gb/sec Slingshot-10 interconnect, with Hewlett Packard Enterprise as the main contractor for the system.
The whole shebang was given a $237 million budget by EuroHPC, which worked out to about $431 per teraflops. The Lumi-C partition was delivered in the summer of 2021, and the Lumi-G GPU partition was delivered at the end of 2021. These partitions were operational in June 2022. The Lumi-G partition was ranked as the number 11 supercomputer in the Top500 rankings in June of this year, so it is no slouch even though it was four years old at the time.
The Lumi-C partition has 1,536 nodes, each of which are equipped with a pair of 64-core “Milan” Epyc 7763 processors spinning at 2.45 GHz. That comes to a total of 196,860 cores, to be precise. This partition is ranked number 262 on the June Top500 rankings, with a peak theoretical performance of 7.29 petaflops and a sustained Linpack performance of 6.3 petaflops, which is a very respectable 82.6 percent computational efficiency in a 1.22 megawatt power envelope. That works out to 5.18 gigaflops per watt.
This machine uses a 200 Gb/sec Slingshot 11 interconnect to lash the nodes together.
The Lumi-G partition has 2,978 nodes, with each one having a single 64-core “Trento” Epyc 7A53 custom processor with four “Aldebaran” MI250X GPU accelerators hanging off that. (The Trento chip was custom built for the “Frontier” supercomputer at Oak Ridge National Laboratory in the US, and provided memory coherence across AMD CPUs and GPUs before that was officially commercially available.) That is a total of 2,978 Trento CPUs, with 190,592 cores, and 11,912 Aldebaran GPUs, with a total of 2.62 million compute units and 167.7 million streaming processors. The nodes in the Lumi-G partition are linked by a 200 Gb/sec Slingshot-11 interconnect. This partition of the Lumi machine has a peak theoretical performance of 531.5 petaflops at FP64 precision, and rated 379.7 petaflops on the Linpack test, which is a computational efficiency of 71.4 percent. This partition weighs in at 7.1 megawatts and delivers 53.43 gigaflops per watt on Linpack, which is more than 10X better power efficiency than the CPU-only partition.
What I know about the future Lumi-AI partition, which is really an upgrade to the Lumi-G partition that I assume will stay in production given the value of those AMD GPUs and the scarcity of GPU capacity on the world, is that it will offer 10X the performance on AI workloads and “around 2X” the performance on FP64 HPC workloads.
It will not be hard to beat the AI performance of the MI250X, which supported INT4 and INT8 data formats as well as FP16 formats, but the performance was the same for all three. So INT4 and INT8 were not precisely native on the Aldebaran GPUs, but grafted on top of FP16, which was fine for that time. But with the Altair GPUs that are being packaged as the MI430X aimed at both HPC and AI workloads (unlike the top-end Altair MI455X, which is aimed at AI inference and training), FP4 and FP8 processing is native, so you get a multiplicative effect as you trim down from FP16. So 4X of that 10X is coming from the shift from FP16 down to FP4. So that leaves 2.5X more raw performance at FP16 comparing the Altair GPU used in the MI430X to the Aldebaran GPU used in the MI250X.
This means, of course, that the node count in the Lumi-AI partition should be a lot smaller, even if the devices themselves will certainly burn more electricity and therefore generate more heat.
Let’s figure out that node count on Lumi-AI.
The MI250X is rated at 47.9 teraflops on its vector units and 95.7 teraflops on its matrix units. The MI430X has a peak of 288 teraflops on FP64 workloads, and AMD was not precise about this being on vector or matrix units. It doesn’t have to be because the MI430X delivers 288 teraflops on both vector and matrix units. If you do the math on the Lumi-G partition, it looks like CSC Finland was counting the vector FP64 performance (47.9 teraflops times four GPUs per node times 2,978 nodes is a peak of 570.6 petaflops. (We don’t know why this doesn’t match the 531.5 Rpeak on the Linpack test.)
If you want 2X performance on the GPUs at FP64 vector for Lumi-AI, then you only need 3,962 MI430X GPUs, and that works out to a mere 991 nodes, a 66.7 percent reduction in node count for the AI partition. (This breaks the exascale barrier, with 1.14 exaflops peak, which is psychologically important. Each node will, however, burn about 2X the energy, at 5 kilowatts for the Venice-Altair 1x4 board compared to 2.5 kilowatts for the Trento-Aldebaran 1x4 board. And that means the Lumi-AI partition will probably consume about 5 megawatts, assuming 991 nodes, which is a 33 percent reduction compared to the Lumi-G partition.
The weird thing is that when we do the math on this number of nodes to derive the AI performance at FP64, we only get an 8X performance increase. You need 1,240 nodes with four MI430X GPUs in Lumi-AI to get a 10X increase over Lumi-G at 4-bit precision. That might mean that CSC Finland will get an extra 25 percent more FP64 vector performance than it is talking about.
I’m going with 1,240 nodes with 4,960 MI430X GPUs for my guess, which is a 58.4 percent reduction in GPU partition node count and about 6.2 megawatts of system power for Lumi-AI. Both Bull, the primary contractor for the Lumi-AI partition, and CSC Finland were clear that AI performance would be 10X higher and HPC performance would be “about 2X.” Why HPC centers are imprecise in their statements is a mystery, and it is counterintuitive for an organization that lives by precision. Go figure.
Using a 256-core “Venice” Epyc 9006 processor, the Lumi-AI partition will have 317,440 cores, which is an appreciable number of compute units to help orchestrate those GPUs and which is 66.7 percent higher than the core count in the current Lumi-G partition.
Given all of this, and a price tag of €387.8 million for the Lumi-AI partition, converting that to US dollars yields a cost of $445.1 million and against a peak FP64 performance of 1.43 exaflops for the Lumi-AI partition, you get a $312 per teraflops at peak. This is a 27.6 percent lower cost per teraflops than the $431 per teraflops peak that the entire Lumi cluster cost when it was installed in 2021. It is not known if the Lumi CPU-only partition and the various storage clusters will also be upgraded under the Lumi-AI contract, but we presume this to be the case. Which makes this an apples-to-apples overall system comparison.
That CSC Finland and its EuroHPC funding arm went with Bull over HPE in this deal is not a big surprise, really. EuroHPC wants to buy European as much as possible. Moreover, while Slingshot is an Ethernet variant that can be possibly modified to adhere to Ultra Ethernet Consortium hardware, Bulls BXI protocol is being adapted to run on UEC hardware and make use of its various features. Which means Bull customers using the BullSequana XH3500 systems will have a choice of switchery when they want to do scale out clustering.
The Lumi-AI GPU partition is expected to be installed in the second half of 2027 in a new datacenter located in Kajaani. Finland, Czechia, Denmark, Estonia, Norway, and Poland are kicking in money alongside the European Union’s EuroHPC JU. The precise amounts everyone is paying was not disclosed.
The Lumi AI factory will not only include this new AMD Venice-Altair cluster, but will also have a quantum computing element dubbed Lumi-IQ. The quantum system is being built by IQM Quantum Computers, a Finnish company that has raised €600 million in venture capital so far and that has aspirations to be Europe’s top quantum computing supplier. The company is preparing to go public and has a valuation of $1.8 billion at this time.
The plan is to installed a Halocene H4 system at CSC Finland in 2027, which will have 150 physical cubits and support 5 logical qubits. Upgrades will be done in 2028 to improve error correction and to improve the reliability of logical qubits. In 2029, CSC Finland will get an upgrade to a Halocene H5 quantum computer, which will have an unknown number of physical qubits but will deliver at least 9 logical qubits. That should imply 270 physical qubits to deliver those local qubits.
Bull Beats Out HPE For Next-Gen Lumi AI Supercomputer
Infleqtion And Nvidia Reduce Physical-to-Logical Qubit Ratio in Quantum Computers
How To Smash The Memory Wall Plaguing High Performance Systems
Operationalize AI At Scale With HPE And Nvidia
Despite Hefty Capital Investments, The Cloud Starts Paying Off For Oracle
Broadcom Rides Rocketing Trend For Custom AI Accelerators
Startup d-Matrix Will Pair Its Raptor Memory-Based XPU To Nvidia Rackscale Iron
HPE Sees No Customer Pushback On Rising Datacenter Equipment Costs
COMPUTE
Dell Says AI Will Drive 75 Percent Of Datacenter Demand By 2030
Nvidia Expands Its Open Source AI Presence With $12.9 Billion Hugging Face Buy
Optics Still Driving Marvell’s AI Business More Than Custom Chips
VMware Intros Private AI Cloud, AI Factory As Workloads Shift To On-Prem
In The Long Run, Nvidia NVSwitch Is The InfiniBand Of Scale Up AI Networks
Peeling Apart That Supposed $120 Billion Chip Deal Google Inked With Marvell
COMPUTE
The Network Is More Of Nvidia’s Computer Than It Ever Was At Sun
STORE
You Probably

[truncated]
