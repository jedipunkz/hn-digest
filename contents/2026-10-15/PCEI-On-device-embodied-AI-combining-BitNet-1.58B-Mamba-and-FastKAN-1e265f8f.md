---
source: "https://github.com/MtnMCG/PCEI"
hn_url: "https://news.ycombinator.com/item?id=50107233"
title: "PCEI – On-device embodied AI combining BitNet 1.58B, Mamba, and FastKAN"
article_title: "GitHub - MtnMCG/PCEI: PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar · GitHub"
image: "https://opengraph.githubassets.com/afbcc9d1339d9f6c94611d9e88866a9560ae5e40b599c480b24b9655727bb00c/MtnMCG/PCEI"
author: "mtnmcg"
captured_at: "2026-10-15T18:02:35Z"
capture_tool: "hn-digest"
hn_id: 50107233
score: 1
comments: 0
posted_at: "2026-10-15T17:39:21Z"
tags:
  - hacker-news
---

# PCEI – On-device embodied AI combining BitNet 1.58B, Mamba, and FastKAN

- HN: [50107233](https://news.ycombinator.com/item?id=50107233)
- Source: [github.com](https://github.com/MtnMCG/PCEI)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T17:39:21Z

## Translation

Title: PCEI – On-device embodied AI combining BitNet 1.58B, Mamba, and FastKAN
Article title: GitHub - MtnMCG/PCEI: PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar · GitHub
Description: PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar - MtnMCG/PCEI

Article text:
GitHub - MtnMCG/PCEI: PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
MtnMCG
PCEI Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
docker docker pcei_framework pcei_framework .env.example .env.example .gitignore .gitignore LICENSE LICENSE PCEI_Feasibility_Assessment.md PCEI_Feasibility_Assessment.md PCEI_Framework_Spec (2).txt PCEI_Framework_Spec (2).txt README.md README.md readme.txt readme.txt View all files Repository files navigation
PCEI: Physically Constrained Emergent Intelligence
An architectural thesis and empirical proof against monolithic Transformers, quadratic memory explosion, and brute-force hyperscaling.
🏛️ Executive Summary & The Case Against Hyperscaling
The modern artificial intelligence consensus is built on an unsustainable premise: that human-level intelligence will emerge simply by taking a quadratic-attention Transformer architecture, expanding it to hundreds of billions of parameters, and brute-forcing trillions of unstructured web tokens through multi-gigawatt datacenters.
This is an engineering and thermodynamic dead end.
The PCEI (Physically Constrained Emergent Intelligence) framework proposes a radical departure founded on biological physics, information theory, and hardware efficiency:
The Transformer Memorization Tax : Standard monolithic Transformers squander >80% of their parametric capacity functioning as lossy, static encyclopedic lookup tables rather than generalized reasoning engines. They scale up parameters not because reasoning requires it, but because unconstrained floating-point attention lacks the metabolic pressure to prune redundant facts.
The Quadratic Wall ($\mathcal{O}(L^2)$) : Global softmax attention requires compute that scales quadratically with sequence length and an exploding key-value ( $KV$ ) cache that chokes memory bandwidth during auto-regressive generation.
Biological & Physical Grounding : Real biological nervous systems operate under strict physical and metabolic constraints: linear state propagation, sparse ternary synaptic transmission ( $\ {-1, 0, 1}$ ), homeostatic field regulation, and offline sleep consolidation.
PCEI proves that when these physical constraints are engineered directly into the neural fabric, a sub-billion parameter model (750M) achieves syntactic fluency, deep latent reasoning, and multi-modal grounding on a single consumer GPU (24GB VRAM) or Apple Silicon Mac.
⚡ Comparative Architecture: Monolithic Transformer vs. PCEI
Architectural Dimension
Monolithic Hyperscale Transformer
PCEI (Physically Constrained Continuum)
Sequence Complexity
$\mathcal{O}(L^2)$ Quadratic Softmax Attention
$\mathcal{O}(L)$ Linear-Time Selective State Spaces (Mamba SSM)
Inference KV-Cache
Unbounded growth ( $\mathcal{O}(L)$ memory per user)
Constant $\mathcal{O}(1)$ recurrent state buffer
Synaptic Resolution
FP16 / BF16 (16-bit floating point multipliers)
1.58-bit Ternary Synaptic Gating ({-1, 0, +1})
Compute Core
Energy-intensive floating-point Matrix Multiply ( $FMUL$ )
Zero-Multiply Integer Addition ( $IADD$ ) Matrix Kernels
Context Unification
Dense $N \times N$ token-to-token cross-attention
Global Workspace Broadcast + Phase-Lagged Neural Synchrony
Memory Consolidation
None (Static post-training weights; continual forgetting)
Two-Stage Cognitive Lifecycle: Wakeful + Latent Dreaming Replay
Homeostasis & Stability
Unstable without massive LayerNorm / Weight Decay
Dual-Timescale Continuous Ambient Energy Field
Minimum Hardware Target
Server clusters (8x H100 80GB SXM)
Single Consumer GPU / Apple Silicon M-Series (24GB)
🧠 The Four Pillars of PCEI
┌──────────────────────────────────────────────┐
│ SENSORY & SYMBOLIC INPUT │
│ (Text, Vision Tokens, Auditory) │
└──────────────────────┬───────────────────────┘
▼
┌──────────────────────────────────────────────┐
│ SELECTIVE STATE-SPACE BACKBONE │
│ 24-Layer Linear Mamba SSM │
│ O(L) Compute · O(1) Memory │
└───────┬──────────────────────────────┬───────┘
│ │
┌───────────────────┴──────────┐ ┌──────────┴──────────────────┐
│ 1.58-BIT TERNARY GATING │ │ DUAL-TIMESCALE AMBIENT │
│ (BitNet) │ │ FIELD │
│ {-1, 0, +1} Synaptic Mesh │◄──────►│ Global Workspace Broadcast │
│ Biological Zero Inhibition │ │ Homeostatic Phase Coupling │
└───────────────────┬──────────┘ └──────────┬──────────────────┘
│ │
└──────────────┬───────────────┘
▼
┌──────────────────────────────────────────────┐
│ THE DREAMING CHAMBER │
│ Continuous Latent Diffusion Replay │
│ Eigenvector Contraction & Synaptic Prune │
└──────────────────────────────────────────────┘
1. Linear-Time Selective State Spaces (Mamba Backbone)
By replacing quadratic dot-product attention with hardware-aware selective state-space models (SSMs), PCEI enables sequence processing with $\mathcal{O}(L)$ compute complexity and constant $\mathcal{O}(1)$ recurrent memory during inference. The model processes continuous streams without memory fragmentation or context-window cliffs.
2. BitNet 1.58b Ternary Quantization
Synaptic weights are constrained to ternary values:
$$\mathbf{W} \in {-1, 0, +1}$$
Zero-Multiply Computing : Floating-point matrix multiplications ( $FMUL$ ) are replaced by pure integer addition and subtraction ( $IADD$ ), reducing matrix computation energy consumption by up to $10\times$ .
Active Zero Density : Comprehensive weight audits confirm an emergent distribution of 40.08% (-1) : 19.82% (0) : 40.09% (+1) . The $\sim 20%$ zero-density functions as biological inhibitory gating, eliminating the parameter bloat required by unconstrained networks.
3. Dual-Timescale Homeostatic Ambient Field
Selective SSMs historically suffer from context fragmentation in multi-turn dialogues because they lack global attention. PCEI solves this via an ambient scalar/vector potential field with dual timescales:
Fast Field ( $\tau_{\text{fast}}$ ) : Captures immediate conversational momentum and token affective resonance.
Slow Field ( $\tau_{\text{slow}}$ ) : Maintains homeostatic energy baselines and projects into all layers via a zero-initialized Global Workspace broadcast:
$$\mathbf{x} {\text{layer}} = \mathbf{x} {\text{layer}} + \tanh(\alpha) \cdot \mathbf{W} {\text{ambient}}(\mathbf{s} {\text{slow}})$$
This binds long-range subjects and predicates across turns without quadratic overhead.
4. Two-Stage Cognitive Lifecycle: The Dreaming Chamber
Hyperscale models suffer from catastrophic forgetting and plastic weight corruption because learning is purely feedforward. PCEI implements a biological two-phase lifecycle:
Wakeful Acquisition : Ingests continuous symbolic, mathematical, and grounded multimodal data.
Latent Dreaming Consolidation : An offline closed-loop diffusion process operates directly on continuous latent representations. Without external text, reverse diffusion denoises latent paths, prunes high-entropy synapses, and contracts the representational manifold onto its principal eigenvectors, systematically reducing held-out perplexity.
📚 Scientific Compendium: 11 Research Papers
The mathematical proofs, empirical telemetry, and ablations underpinning the PCEI architecture are published in full academic form within pcei_framework/docs/ :
PCEI/
├── LICENSE # MIT Non-Commercial License (MIT-NC)
├── README.md # Master architectural thesis & documentation
├── .env.example # Environment configuration template
├── pcei_framework/
│ ├── docs/ # Research compendium & system specifications
│ │ ├── ARCHITECTURE_SPECIFICATION.md
│ │ ├── CURRICULUM_SPECIFICATION.md
│ │ ├── MODEL_PRODUCTION_AND_WEIGHT_GRAFTING.md
│ │ ├── USER_OPERATIONS_GUIDE.md
│ │ └── RESEARCH_PAPER_01 ... 11.md
│ ├── src/ # Core neural engine
│ │ ├── mamba_text.py # Selective state-space backbone
│ │ ├── bitnet.py # Ternary {-1, 0, +1} quantized linear layers
│ │ ├── ambient_field.py # Dual-timescale homeostatic energy regulator
│ │ ├── dreaming_chamber.py # Closed-loop latent diffusion consolidation
│ │ ├── multimodal_cortex.py # Sensory cross-modal integration
│ │ └── global_workspace.py # Phase-lagged neural synchrony broadcast
│ ├── tools/ # Operational & diagnostic tools
│ │ ├── resonance_chat.py # Interactive CLI inference engine
│ │ ├── teacher_generator.py # High-density synthetic curriculum generator
│ │ └── probe_generation.py # Live latent probing & diagnostic suite
│ ├── scripts/ # Cloud synchronization & telemetry daemons
│ └── data/ # Curated reasoning curricula & exemplars
└── docker/ # Reproducible container environments
🚀 Quickstart & Getting Started
PyTorch 2.2+ (with CUDA for NVIDIA GPUs or MPS for Apple Silicon)
macOS (Apple Silicon M1/M2/M3/M4) or Linux (Ubuntu 22.04+)
# Clone the repository
git clone https://github.com/MtnMCG/PCEI.git
cd PCEI/pcei_framework
# Create and activate virtual environment
python3 -m venv venv
source venv/bin/activate
# Install dependencies
pip install -r requirements.txt
# Copy environment configuration
cp ../.env.example .env
3. Interactive Resonance Chat (Local Inference)
Run conversational inference utilizing the 24-layer Mamba TextStack with continuous ambient field modulation and latent pondering:
python3 tools/resonance_chat.py \
--checkpoint checkpoints/pcei_750m_reasoning_assistant_latest.pt \
--device mps # or --device cuda
4. Running the Dreaming Chamber (Latent Consolidation)
Execute offline sleep consolidation to contract representational manifolds:
python3 src/dreaming_chamber.py \
--checkpoint checkpoints/pcei_750m_reasoning_assistant_latest.pt \
--cycles 500 \
--denoise_steps 20
📜 License & Intellectual Property
This project is licensed under the MIT Non-Commercial License (MIT-NC) .
Academic, Educational & Personal Research : Freely granted without fee or restriction. You are welcome to inspect, study, experiment with, and modify the code for non-commercial research.
Commercial Use & Monetization : Any commercial use, including deployment in commercial services, incorporation into proprietary products, SaaS offerings, or training proprietary commercial models, is strictly prohibited without prior express written permission from the copyright holder.
For commercial licensing, enterprise partnerships, or inquiries, please open an issue or reach out to the project maintainer.
If you use the PCEI architecture, theoretical papers, or empirical methodologies in your research, pleas

[truncated]

## Original Extract

PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar - MtnMCG/PCEI

GitHub - MtnMCG/PCEI: PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
MtnMCG
PCEI Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
PCEI or Physically Constrained Emergent Intelligence is a new machine learning and large language model architecture that is memory efficient and bio-similar
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
docker docker pcei_framework pcei_framework .env.example .env.example .gitignore .gitignore LICENSE LICENSE PCEI_Feasibility_Assessment.md PCEI_Feasibility_Assessment.md PCEI_Framework_Spec (2).txt PCEI_Framework_Spec (2).txt README.md README.md readme.txt readme.txt View all files Repository files navigation
PCEI: Physically Constrained Emergent Intelligence
An architectural thesis and empirical proof against monolithic Transformers, quadratic memory explosion, and brute-force hyperscaling.
🏛️ Executive Summary & The Case Against Hyperscaling
The modern artificial intelligence consensus is built on an unsustainable premise: that human-level intelligence will emerge simply by taking a quadratic-attention Transformer architecture, expanding it to hundreds of billions of parameters, and brute-forcing trillions of unstructured web tokens through multi-gigawatt datacenters.
This is an engineering and thermodynamic dead end.
The PCEI (Physically Constrained Emergent Intelligence) framework proposes a radical departure founded on biological physics, information theory, and hardware efficiency:
The Transformer Memorization Tax : Standard monolithic Transformers squander >80% of their parametric capacity functioning as lossy, static encyclopedic lookup tables rather than generalized reasoning engines. They scale up parameters not because reasoning requires it, but because unconstrained floating-point attention lacks the metabolic pressure to prune redundant facts.
The Quadratic Wall ($\mathcal{O}(L^2)$) : Global softmax attention requires compute that scales quadratically with sequence length and an exploding key-value ( $KV$ ) cache that chokes memory bandwidth during auto-regressive generation.
Biological & Physical Grounding : Real biological nervous systems operate under strict physical and metabolic constraints: linear state propagation, sparse ternary synaptic transmission ( $\ {-1, 0, 1}$ ), homeostatic field regulation, and offline sleep consolidation.
PCEI proves that when these physical constraints are engineered directly into the neural fabric, a sub-billion parameter model (750M) achieves syntactic fluency, deep latent reasoning, and multi-modal grounding on a single consumer GPU (24GB VRAM) or Apple Silicon Mac.
⚡ Comparative Architecture: Monolithic Transformer vs. PCEI
Architectural Dimension
Monolithic Hyperscale Transformer
PCEI (Physically Constrained Continuum)
Sequence Complexity
$\mathcal{O}(L^2)$ Quadratic Softmax Attention
$\mathcal{O}(L)$ Linear-Time Selective State Spaces (Mamba SSM)
Inference KV-Cache
Unbounded growth ( $\mathcal{O}(L)$ memory per user)
Constant $\mathcal{O}(1)$ recurrent state buffer
Synaptic Resolution
FP16 / BF16 (16-bit floating point multipliers)
1.58-bit Ternary Synaptic Gating ({-1, 0, +1})
Compute Core
Energy-intensive floating-point Matrix Multiply ( $FMUL$ )
Zero-Multiply Integer Addition ( $IADD$ ) Matrix Kernels
Context Unification
Dense $N \times N$ token-to-token cross-attention
Global Workspace Broadcast + Phase-Lagged Neural Synchrony
Memory Consolidation
None (Static post-training weights; continual forgetting)
Two-Stage Cognitive Lifecycle: Wakeful + Latent Dreaming Replay
Homeostasis & Stability
Unstable without massive LayerNorm / Weight Decay
Dual-Timescale Continuous Ambient Energy Field
Minimum Hardware Target
Server clusters (8x H100 80GB SXM)
Single Consumer GPU / Apple Silicon M-Series (24GB)
🧠 The Four Pillars of PCEI
┌──────────────────────────────────────────────┐
│ SENSORY & SYMBOLIC INPUT │
│ (Text, Vision Tokens, Auditory) │
└──────────────────────┬───────────────────────┘
▼
┌──────────────────────────────────────────────┐
│ SELECTIVE STATE-SPACE BACKBONE │
│ 24-Layer Linear Mamba SSM │
│ O(L) Compute · O(1) Memory │
└───────┬──────────────────────────────┬───────┘
│ │
┌───────────────────┴──────────┐ ┌──────────┴──────────────────┐
│ 1.58-BIT TERNARY GATING │ │ DUAL-TIMESCALE AMBIENT │
│ (BitNet) │ │ FIELD │
│ {-1, 0, +1} Synaptic Mesh │◄──────►│ Global Workspace Broadcast │
│ Biological Zero Inhibition │ │ Homeostatic Phase Coupling │
└───────────────────┬──────────┘ └──────────┬──────────────────┘
│ │
└──────────────┬───────────────┘
▼
┌──────────────────────────────────────────────┐
│ THE DREAMING CHAMBER │
│ Continuous Latent Diffusion Replay │
│ Eigenvector Contraction & Synaptic Prune │
└──────────────────────────────────────────────┘
1. Linear-Time Selective State Spaces (Mamba Backbone)
By replacing quadratic dot-product attention with hardware-aware selective state-space models (SSMs), PCEI enables sequence processing with $\mathcal{O}(L)$ compute complexity and constant $\mathcal{O}(1)$ recurrent memory during inference. The model processes continuous streams without memory fragmentation or context-window cliffs.
2. BitNet 1.58b Ternary Quantization
Synaptic weights are constrained to ternary values:
$$\mathbf{W} \in {-1, 0, +1}$$
Zero-Multiply Computing : Floating-point matrix multiplications ( $FMUL$ ) are replaced by pure integer addition and subtraction ( $IADD$ ), reducing matrix computation energy consumption by up to $10\times$ .
Active Zero Density : Comprehensive weight audits confirm an emergent distribution of 40.08% (-1) : 19.82% (0) : 40.09% (+1) . The $\sim 20%$ zero-density functions as biological inhibitory gating, eliminating the parameter bloat required by unconstrained networks.
3. Dual-Timescale Homeostatic Ambient Field
Selective SSMs historically suffer from context fragmentation in multi-turn dialogues because they lack global attention. PCEI solves this via an ambient scalar/vector potential field with dual timescales:
Fast Field ( $\tau_{\text{fast}}$ ) : Captures immediate conversational momentum and token affective resonance.
Slow Field ( $\tau_{\text{slow}}$ ) : Maintains homeostatic energy baselines and projects into all layers via a zero-initialized Global Workspace broadcast:
$$\mathbf{x} {\text{layer}} = \mathbf{x} {\text{layer}} + \tanh(\alpha) \cdot \mathbf{W} {\text{ambient}}(\mathbf{s} {\text{slow}})$$
This binds long-range subjects and predicates across turns without quadratic overhead.
4. Two-Stage Cognitive Lifecycle: The Dreaming Chamber
Hyperscale models suffer from catastrophic forgetting and plastic weight corruption because learning is purely feedforward. PCEI implements a biological two-phase lifecycle:
Wakeful Acquisition : Ingests continuous symbolic, mathematical, and grounded multimodal data.
Latent Dreaming Consolidation : An offline closed-loop diffusion process operates directly on continuous latent representations. Without external text, reverse diffusion denoises latent paths, prunes high-entropy synapses, and contracts the representational manifold onto its principal eigenvectors, systematically reducing held-out perplexity.
📚 Scientific Compendium: 11 Research Papers
The mathematical proofs, empirical telemetry, and ablations underpinning the PCEI architecture are published in full academic form within pcei_framework/docs/ :
PCEI/
├── LICENSE # MIT Non-Commercial License (MIT-NC)
├── README.md # Master architectural thesis & documentation
├── .env.example # Environment configuration template
├── pcei_framework/
│ ├── docs/ # Research compendium & system specifications
│ │ ├── ARCHITECTURE_SPECIFICATION.md
│ │ ├── CURRICULUM_SPECIFICATION.md
│ │ ├── MODEL_PRODUCTION_AND_WEIGHT_GRAFTING.md
│ │ ├── USER_OPERATIONS_GUIDE.md
│ │ └── RESEARCH_PAPER_01 ... 11.md
│ ├── src/ # Core neural engine
│ │ ├── mamba_text.py # Selective state-space backbone
│ │ ├── bitnet.py # Ternary {-1, 0, +1} quantized linear layers
│ │ ├── ambient_field.py # Dual-timescale homeostatic energy regulator
│ │ ├── dreaming_chamber.py # Closed-loop latent diffusion consolidation
│ │ ├── multimodal_cortex.py # Sensory cross-modal integration
│ │ └── global_workspace.py # Phase-lagged neural synchrony broadcast
│ ├── tools/ # Operational & diagnostic tools
│ │ ├── resonance_chat.py # Interactive CLI inference engine
│ │ ├── teacher_generator.py # High-density synthetic curriculum generator
│ │ └── probe_generation.py # Live latent probing & diagnostic suite
│ ├── scripts/ # Cloud synchronization & telemetry daemons
│ └── data/ # Curated reasoning curricula & exemplars
└── docker/ # Reproducible container environments
🚀 Quickstart & Getting Started
PyTorch 2.2+ (with CUDA for NVIDIA GPUs or MPS for Apple Silicon)
macOS (Apple Silicon M1/M2/M3/M4) or Linux (Ubuntu 22.04+)
# Clone the repository
git clone https://github.com/MtnMCG/PCEI.git
cd PCEI/pcei_framework
# Create and activate virtual environment
python3 -m venv venv
source venv/bin/activate
# Install dependencies
pip install -r requirements.txt
# Copy environment configuration
cp ../.env.example .env
3. Interactive Resonance Chat (Local Inference)
Run conversational inference utilizing the 24-layer Mamba TextStack with continuous ambient field modulation and latent pondering:
python3 tools/resonance_chat.py \
--checkpoint checkpoints/pcei_750m_reasoning_assistant_latest.pt \
--device mps # or --device cuda
4. Running the Dreaming Chamber (Latent Consolidation)
Execute offline sleep consolidation to contract representational manifolds:
python3 src/dreaming_chamber.py \
--checkpoint checkpoints/pcei_750m_reasoning_assistant_latest.pt \
--cycles 500 \
--denoise_steps 20
📜 License & Intellectual Property
This project is licensed under the MIT Non-Commercial License (MIT-NC) .
Academic, Educational & Personal Research : Freely granted without fee or restriction. You are welcome to inspect, study, experiment with, and modify the code for non-commercial research.
Commercial Use & Monetization : Any commercial use, including deployment in commercial services, incorporation into proprietary products, SaaS offerings, or training proprietary commercial models, is strictly prohibited without prior express written permission from the copyright holder.
For commercial licensing, enterprise partnerships, or inquiries, please open an issue or reach out to the project maintainer.
If you use the PCEI architecture, theoretical papers, or empirical methodologies in your research, pleas

[truncated]
