---
source: "https://research.nvidia.com/labs/lpr/vflash/"
hn_url: "https://news.ycombinator.com/item?id=50398514"
title: "VFlash: Enabling System-1 VLMs for Physical AI"
article_title: "VFlash: Enabling System-1 VLMs for Physical AI · NVIDIA Research"
image: ""
author: "gmays"
captured_at: "2026-11-06T21:18:43Z"
capture_tool: "hn-digest"
hn_id: 50398514
score: 2
comments: 0
posted_at: "2026-11-06T21:04:04Z"
tags:
  - hacker-news
---

# VFlash: Enabling System-1 VLMs for Physical AI

- HN: [50398514](https://news.ycombinator.com/item?id=50398514)
- Source: [research.nvidia.com](https://research.nvidia.com/labs/lpr/vflash/)
- Score: 2
- Comments: 0
- Posted: 2026-11-06T21:04:04Z

## Translation

Title: VFlash: Enabling System-1 VLMs for Physical AI
Article title: VFlash: Enabling System-1 VLMs for Physical AI · NVIDIA Research
Description: VFlash enables System-1 VLMs for Physical AI through fast, accurate parallel visual planning, accelerating robot and autonomous driving inference.

Article text:
Skip to the method
VFlash
Overview Method Accuracy Primitives Robot control Physical AI Team BibTeX
V F l as h Enabling System-1 VLMs
for Physical AI.
Perception →
Visual plans
→ Action
A generalist decoding framework for fast, accurate visual planning. Decode complete primitives in parallel for faster robot and driving inference.
One parallel draft per visual primitive.
3D waypoint error vs. the AR target.
Boxes, trajectories, dense motion & text.
System-1 visual planning for Physical AI.
Embodied models do not necessarily think in text; they can also think in visual plans. Predicting where to grasp, how to move, and how a scene will evolve bridges perception and action for robots and autonomous vehicles. However, agents must continually update these plans, and autoregressive generation of coordinate tokens makes detailed visual plans increasingly expensive to generate.
We introduce VFlash, a generalist decoding framework for Physical AI that adapts VLM speculative decoding around a central insight: the known structure of visual primitives can be exploited to achieve spatial accuracy without reproducing the autoregressive target’s exact tokens. A lightweight drafter trained on annotated geometry can generate complete visual primitives in a parallel draft and achieve more accurate robot trajectories than its autoregressive target.
Applying speculative decoding with a fixed block size is not optimal for Physical AI. Generalist decoding must accommodate rich output types, known structures, and lengths specified on demand.
Typical outputs span both text, such as thinking and sub-task planning, and visual primitives including boxes, points, depth, and 2D/3D trajectories for manipulation and driving.
Decoding outputs
are structured .
Points, bounding boxes, and trajectories have known geometric structures. This structure lets blocks cover entire primitives while assigning coordinate roles to token positions.
Decoding requirements
are versatile .
A generalist decoder must support variable output lengths. NAVSIM requires 8 waypoints over 4 seconds, Waymo 20 over 5 seconds, and Alpamayo 64 over 6.4 seconds.
VFlash supports rich Physical AI outputs with systematic parallel decoding designs. A generic token window does not capture where one such primitive begins or ends.
A tool call declares the output length, while the primitive type specifies the coordinate layout. Using frozen VLM context, a lightweight drafter predicts the entire primitive in one parallel pass. The target VLM then processes the plan to update its context and extract features used for downstream action.
Plan the gripper’s next motion.
Encode the observation and instruction once to build the target context.
Training and decoding blocks should match complete primitive boundaries , allowing the drafter to learn and predict the full geometry together rather than fragments split by fixed token windows. VFlash introduces a structured tool call to declare the length of a variable-size primitive; its type determines the coordinate layout. Flash Primitive Decoding allocates a block for the complete geometry and fills the entire primitive in one parallel draft conditioned on frozen VLM features .
Each output position needs a known role before parallel prediction begins. Different coordinate values can occupy different numbers of text tokens and shift later positions. We use dedicated coordinate tokens or fixed-width numeric fields to keep waypoint, axis, and delimiter positions consistent. This consistent layout supports parallel prediction with constraints applied to each position.
The target VLM remains frozen while we train a lightweight drafter on annotated geometry. The drafter has five non-causal transformer layers and is trained with token cross-entropy on annotated primitives . It predicts complete visual primitives in parallel without requiring token-level agreement . Sub-task text is decoded speculatively and verified by the target VLM.
A drafter that surpasses
its autoregressive target.
VFlash reduces all six trajectory error metrics across ShareRobot-Bench and VABench-V relative to autoregressive (AR) decoding. Across three robot datasets and three target variants, mean 3D displacement error (ADE) decreases by 42.6% relative to the autoregressive target .
Principled parallelism across visual tasks.
VFlash predicts complete 2D/3D trajectories in parallel. Matching training and decoding blocks to primitive boundaries lets the drafter learn and predict the full geometry together. Across all evaluated 2D/3D robot settings, trajectory errors decrease, including average 3D waypoint error.
With frozen MolmoMotion-AR, VFlash reduces DAVIS average displacement error by 27% and achieves comparable overall quality on dense 3D motion forecasting. Single-rollout generation is approximately 25× faster on H100 at batch size one. VFlash drafts each future frame of query points as one block.
On COCO and Dense200, VFlash delivers 2.6–2.8× more boxes per second , with mean F1 close to the autoregressive target. Flash Primitive Decoding predicts all coordinates of a box in a single parallel draft. Each block covers one box, and subsequent blocks handle additional objects.
The same drafter retains verified decoding for sub-task text. Text uses 16-position blocks: one accepted anchor and 15 proposals. The target verifies the block in one pass and keeps the longest matching prefix plus its own next token. Both methods produce an identical sub-task description for the example shown.
Visual plans improve downstream control.
The downstream world action models and diffusion action heads differ in architecture, but share a common design: they cross-attend to VLM features encoding the generated visual plans. We obtain these features by passing VFlash’s predicted primitives through the target VLM. The action models are trained for robot manipulation or driving, while the VLM remains frozen.
MolmoSpaces · world action model
The robot policy is a world action model (WAM) obtained by post-training Cosmos-Predict2.5 for control. It keeps its camera and robot-state inputs and receives the frozen VLM’s question, sub-task, and trajectory features through its existing cross-attention pathway. The WAM predicts 20-step action chunks and replans every three chunks. VFlash reaches 70.2% average success versus 62.6% for the baseline.
With the same WAM architecture, autoregressive 2D+3D plans yield 61.0% success versus VFlash’s 70.2% . Success falls to 63.7% when image features replace trajectory features while retaining the generated sub-task, demonstrating the value of visual geometry beyond language .
Checkmarks identify which VLM states condition the WAM. All variants use the same WAM architecture and training budget. Success is averaged equally across the four suites.
LIBERO family · diffusion action head
Gains remain under perturbations.
A PIv3 action head reads frozen VLM features through cross-attention. The baseline encodes images and instructions using Qwen3-VL-4B. Our VLM supplies states through the predicted sub-task or full plan, predicting a sub-task and ten-waypoint 2D+3D trajectory from primary-camera and wrist-camera images before each eight-action chunk.
Accelerating Visual Planning across Physical AI
Physical AI is latency-sensitive. VFlash pairs a target VLM with a compact drafter. The runtime manages drafting, verification where required, KV-cache updates, and downstream feature extraction. Measurements use customized TensorRT-Edge-LLM runtimes and TensorRT engines on Jetson Thor, DGX Spark, and H100.
Efficient visual planning for autonomous driving.
VFlash goes beyond robot manipulation. With a frozen planning VLM and a ReCogDrive-derived action head, it achieves competitive driving performance on NAVSIM and Waymo while substantially reducing decoding latency . The same variable-length decoder also supports the trajectories used by Alpamayo.
Autoregressive decoding accounts for 75–87% of measured model-stage latency on Thor . VFlash reduces this stage by 12–42× , yielding 5× faster robot inference and over 3× faster NAVSIM driving inference on Jetson Thor and DGX Spark.
Measurements use batch size one and process one request at a time. CUDA events measure model-stage latency, including target feature preparation. Speedups compare AR and VFlash on the same task and hardware platform.
Precision and hardware. Batch size one, one request at a time. Vision, target, and draft engines use FP16 with FP16 KV caches, without weight or KV quantization. Action-model precision is configured separately: DiT weights are BF16, with FP16 attention in the optimized robot WAM.
Primitive and feature passes. VFlash uses verified 16-position blocks for text, then drafts a whole visual primitive. One cached target prefill over the accepted primitive produces downstream features and is included in decoding time. AR, DFlash, and VFlash receive the same supplied tool-call JSON in the measured pipelines.
Action-model optimizations. On Thor and Spark, the WAM uses a custom head-dimension-128 CuTe DSL attention kernel; the H100 baseline uses native TensorRT attention. Driving latency is measured with ten DDIM steps to match the comparison planner. The paper evaluates driving quality with five DDIM steps.
Timing boundary. Each request begins with fresh vision computation and model state. CUDA events measure the instrumented stages, including host launch gaps. Model loading, engine building, graph capture, warmup, file loading, CPU preprocessing outside the measured modules, and output detokenization are excluded. Totals sum the stage means and report model-pipeline latency rather than full sensor-to-actuator latency.
Sampling. Thor robotics uses 24 records and ten passes after five warmups; Spark and H100 robotics use one pass over the same 24 records. Driving uses five warmups followed by ten passes over five records.
1 NVIDIA · 2 Johns Hopkins University · 3 Hong Kong Polytechnic University · 4 Purdue University · 5 The University of Hong Kong
* Equal contribution. † Work done during an internship at NVIDIA. ‡ Corresponding author.
@misc{chou2026vflash,
title = {{VFlash}: Fast, Accurate Visual Planning via Flash Primitive Decoding},
author = {Chou, Yu-Cheng and Wang, Shihao and Huang, Jin and Huang, Chi-Pin and Yang, Chiao-An and Fan, Xianzhe and Wang, Jing and Li, Zhiqi and Kuang, Yuanguo and Zheng, Ruijie and Bjorck, Johan and Liu, Sifei and Su, Hang and Badki, Abhishek and Zhang, Lei and Yuille, Alan and Fan, Linxi and Kautz, Jan and Yu, Zhiding},
year = {2026},
url = {https://research.nvidia.com/labs/lpr/vflash/}
}
VFlash · NVIDIA Learning and Perception Research Back to top ↑ Privacy Policy Manage My Privacy Do Not Sell or Share My Data Terms of Service Accessibility Corporate Policies Contact © 2026 NVIDIA Corporation

## Original Extract

VFlash enables System-1 VLMs for Physical AI through fast, accurate parallel visual planning, accelerating robot and autonomous driving inference.

Skip to the method
VFlash
Overview Method Accuracy Primitives Robot control Physical AI Team BibTeX
V F l as h Enabling System-1 VLMs
for Physical AI.
Perception →
Visual plans
→ Action
A generalist decoding framework for fast, accurate visual planning. Decode complete primitives in parallel for faster robot and driving inference.
One parallel draft per visual primitive.
3D waypoint error vs. the AR target.
Boxes, trajectories, dense motion & text.
System-1 visual planning for Physical AI.
Embodied models do not necessarily think in text; they can also think in visual plans. Predicting where to grasp, how to move, and how a scene will evolve bridges perception and action for robots and autonomous vehicles. However, agents must continually update these plans, and autoregressive generation of coordinate tokens makes detailed visual plans increasingly expensive to generate.
We introduce VFlash, a generalist decoding framework for Physical AI that adapts VLM speculative decoding around a central insight: the known structure of visual primitives can be exploited to achieve spatial accuracy without reproducing the autoregressive target’s exact tokens. A lightweight drafter trained on annotated geometry can generate complete visual primitives in a parallel draft and achieve more accurate robot trajectories than its autoregressive target.
Applying speculative decoding with a fixed block size is not optimal for Physical AI. Generalist decoding must accommodate rich output types, known structures, and lengths specified on demand.
Typical outputs span both text, such as thinking and sub-task planning, and visual primitives including boxes, points, depth, and 2D/3D trajectories for manipulation and driving.
Decoding outputs
are structured .
Points, bounding boxes, and trajectories have known geometric structures. This structure lets blocks cover entire primitives while assigning coordinate roles to token positions.
Decoding requirements
are versatile .
A generalist decoder must support variable output lengths. NAVSIM requires 8 waypoints over 4 seconds, Waymo 20 over 5 seconds, and Alpamayo 64 over 6.4 seconds.
VFlash supports rich Physical AI outputs with systematic parallel decoding designs. A generic token window does not capture where one such primitive begins or ends.
A tool call declares the output length, while the primitive type specifies the coordinate layout. Using frozen VLM context, a lightweight drafter predicts the entire primitive in one parallel pass. The target VLM then processes the plan to update its context and extract features used for downstream action.
Plan the gripper’s next motion.
Encode the observation and instruction once to build the target context.
Training and decoding blocks should match complete primitive boundaries , allowing the drafter to learn and predict the full geometry together rather than fragments split by fixed token windows. VFlash introduces a structured tool call to declare the length of a variable-size primitive; its type determines the coordinate layout. Flash Primitive Decoding allocates a block for the complete geometry and fills the entire primitive in one parallel draft conditioned on frozen VLM features .
Each output position needs a known role before parallel prediction begins. Different coordinate values can occupy different numbers of text tokens and shift later positions. We use dedicated coordinate tokens or fixed-width numeric fields to keep waypoint, axis, and delimiter positions consistent. This consistent layout supports parallel prediction with constraints applied to each position.
The target VLM remains frozen while we train a lightweight drafter on annotated geometry. The drafter has five non-causal transformer layers and is trained with token cross-entropy on annotated primitives . It predicts complete visual primitives in parallel without requiring token-level agreement . Sub-task text is decoded speculatively and verified by the target VLM.
A drafter that surpasses
its autoregressive target.
VFlash reduces all six trajectory error metrics across ShareRobot-Bench and VABench-V relative to autoregressive (AR) decoding. Across three robot datasets and three target variants, mean 3D displacement error (ADE) decreases by 42.6% relative to the autoregressive target .
Principled parallelism across visual tasks.
VFlash predicts complete 2D/3D trajectories in parallel. Matching training and decoding blocks to primitive boundaries lets the drafter learn and predict the full geometry together. Across all evaluated 2D/3D robot settings, trajectory errors decrease, including average 3D waypoint error.
With frozen MolmoMotion-AR, VFlash reduces DAVIS average displacement error by 27% and achieves comparable overall quality on dense 3D motion forecasting. Single-rollout generation is approximately 25× faster on H100 at batch size one. VFlash drafts each future frame of query points as one block.
On COCO and Dense200, VFlash delivers 2.6–2.8× more boxes per second , with mean F1 close to the autoregressive target. Flash Primitive Decoding predicts all coordinates of a box in a single parallel draft. Each block covers one box, and subsequent blocks handle additional objects.
The same drafter retains verified decoding for sub-task text. Text uses 16-position blocks: one accepted anchor and 15 proposals. The target verifies the block in one pass and keeps the longest matching prefix plus its own next token. Both methods produce an identical sub-task description for the example shown.
Visual plans improve downstream control.
The downstream world action models and diffusion action heads differ in architecture, but share a common design: they cross-attend to VLM features encoding the generated visual plans. We obtain these features by passing VFlash’s predicted primitives through the target VLM. The action models are trained for robot manipulation or driving, while the VLM remains frozen.
MolmoSpaces · world action model
The robot policy is a world action model (WAM) obtained by post-training Cosmos-Predict2.5 for control. It keeps its camera and robot-state inputs and receives the frozen VLM’s question, sub-task, and trajectory features through its existing cross-attention pathway. The WAM predicts 20-step action chunks and replans every three chunks. VFlash reaches 70.2% average success versus 62.6% for the baseline.
With the same WAM architecture, autoregressive 2D+3D plans yield 61.0% success versus VFlash’s 70.2% . Success falls to 63.7% when image features replace trajectory features while retaining the generated sub-task, demonstrating the value of visual geometry beyond language .
Checkmarks identify which VLM states condition the WAM. All variants use the same WAM architecture and training budget. Success is averaged equally across the four suites.
LIBERO family · diffusion action head
Gains remain under perturbations.
A PIv3 action head reads frozen VLM features through cross-attention. The baseline encodes images and instructions using Qwen3-VL-4B. Our VLM supplies states through the predicted sub-task or full plan, predicting a sub-task and ten-waypoint 2D+3D trajectory from primary-camera and wrist-camera images before each eight-action chunk.
Accelerating Visual Planning across Physical AI
Physical AI is latency-sensitive. VFlash pairs a target VLM with a compact drafter. The runtime manages drafting, verification where required, KV-cache updates, and downstream feature extraction. Measurements use customized TensorRT-Edge-LLM runtimes and TensorRT engines on Jetson Thor, DGX Spark, and H100.
Efficient visual planning for autonomous driving.
VFlash goes beyond robot manipulation. With a frozen planning VLM and a ReCogDrive-derived action head, it achieves competitive driving performance on NAVSIM and Waymo while substantially reducing decoding latency . The same variable-length decoder also supports the trajectories used by Alpamayo.
Autoregressive decoding accounts for 75–87% of measured model-stage latency on Thor . VFlash reduces this stage by 12–42× , yielding 5× faster robot inference and over 3× faster NAVSIM driving inference on Jetson Thor and DGX Spark.
Measurements use batch size one and process one request at a time. CUDA events measure model-stage latency, including target feature preparation. Speedups compare AR and VFlash on the same task and hardware platform.
Precision and hardware. Batch size one, one request at a time. Vision, target, and draft engines use FP16 with FP16 KV caches, without weight or KV quantization. Action-model precision is configured separately: DiT weights are BF16, with FP16 attention in the optimized robot WAM.
Primitive and feature passes. VFlash uses verified 16-position blocks for text, then drafts a whole visual primitive. One cached target prefill over the accepted primitive produces downstream features and is included in decoding time. AR, DFlash, and VFlash receive the same supplied tool-call JSON in the measured pipelines.
Action-model optimizations. On Thor and Spark, the WAM uses a custom head-dimension-128 CuTe DSL attention kernel; the H100 baseline uses native TensorRT attention. Driving latency is measured with ten DDIM steps to match the comparison planner. The paper evaluates driving quality with five DDIM steps.
Timing boundary. Each request begins with fresh vision computation and model state. CUDA events measure the instrumented stages, including host launch gaps. Model loading, engine building, graph capture, warmup, file loading, CPU preprocessing outside the measured modules, and output detokenization are excluded. Totals sum the stage means and report model-pipeline latency rather than full sensor-to-actuator latency.
Sampling. Thor robotics uses 24 records and ten passes after five warmups; Spark and H100 robotics use one pass over the same 24 records. Driving uses five warmups followed by ten passes over five records.
1 NVIDIA · 2 Johns Hopkins University · 3 Hong Kong Polytechnic University · 4 Purdue University · 5 The University of Hong Kong
* Equal contribution. † Work done during an internship at NVIDIA. ‡ Corresponding author.
@misc{chou2026vflash,
title = {{VFlash}: Fast, Accurate Visual Planning via Flash Primitive Decoding},
author = {Chou, Yu-Cheng and Wang, Shihao and Huang, Jin and Huang, Chi-Pin and Yang, Chiao-An and Fan, Xianzhe and Wang, Jing and Li, Zhiqi and Kuang, Yuanguo and Zheng, Ruijie and Bjorck, Johan and Liu, Sifei and Su, Hang and Badki, Abhishek and Zhang, Lei and Yuille, Alan and Fan, Linxi and Kautz, Jan and Yu, Zhiding},
year = {2026},
url = {https://research.nvidia.com/labs/lpr/vflash/}
}
VFlash · NVIDIA Learning and Perception Research Back to top ↑ Privacy Policy Manage My Privacy Do Not Sell or Share My Data Terms of Service Accessibility Corporate Policies Contact © 2026 NVIDIA Corporation
