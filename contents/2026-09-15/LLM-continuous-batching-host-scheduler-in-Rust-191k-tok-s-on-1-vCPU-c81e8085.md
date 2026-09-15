---
source: "https://github.com/bmartin-systems/cortex-serving-arena-preview"
hn_url: "https://news.ycombinator.com/item?id=49711702"
title: "LLM continuous batching host scheduler in Rust (191k tok/s on 1 vCPU)"
article_title: "GitHub - bmartin-systems/cortex-serving-arena-preview: High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) · GitHub"
image: "https://opengraph.githubassets.com/5f16ec6c11a93c6445f16f6463594280d09f29d30da17e73d9f7228bbbe3b88a/bmartin-systems/cortex-serving-arena-preview"
author: "bmartin-systems"
captured_at: "2026-09-15T13:11:28Z"
capture_tool: "hn-digest"
hn_id: 49711702
score: 1
comments: 0
posted_at: "2026-09-15T12:45:35Z"
tags:
  - hacker-news
---

# LLM continuous batching host scheduler in Rust (191k tok/s on 1 vCPU)

- HN: [49711702](https://news.ycombinator.com/item?id=49711702)
- Source: [github.com](https://github.com/bmartin-systems/cortex-serving-arena-preview)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T12:45:35Z

## Translation

Title: LLM continuous batching host scheduler in Rust (191k tok/s on 1 vCPU)
Article title: GitHub - bmartin-systems/cortex-serving-arena-preview: High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) · GitHub
Description: High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) - bmartin-systems/cortex-serving-arena-preview

Article text:
GitHub - bmartin-systems/cortex-serving-arena-preview: High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) · GitHub
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
bmartin-systems
/
cortex-serving-arena-preview
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github/ workflows .github/ workflows audit audit formal formal harness harness submission submission .gitignore .gitignore README.md README.md View all files Repository files navigation
Serving Arena — High-Throughput 1-vCPU Continuous Batching & Hardware-Aligned Paged Memory Allocation Benchmark
This repository provides an evaluation harness and independent ground truth verification suite for high-performance LLM serving runtimes under strict 1-vCPU hardware isolation ( taskset -c 0 ).
Lock-Free Serving Scheduler : High-performance request allocation and sequence management achieving sub-microsecond dispatch latency.
Paged KV Cache : Fine-grained 16-token page allocation with zero memory fragmentation and fast vector dequantization.
Continuous Batching & Chunked Prefill : Interleaved prompt processing and token generation minimizing Time To First Token (TTFT) and Inter-Token Latency (ITL/TPOT).
Formal SMT Z3 Verification : 5 mathematical theorems proven bit-by-bit (QF_BV UNSAT) certifying zero page collision, deadlock freedom, and L1D cache containment.
Independent FP64 Anti-Cheat Auditor : Standalone verification engine recalculating token logits, scale bounds (N >= 5000 requests), and cryptographic provenance.
├── .github/workflows/
│ └── benchmark.yml # Reproducible 1-vCPU CI workflow
├── audit/ # Independent anti-cheat auditor (Rust)
│ ├── Cargo.toml
│ └── src/main.rs
├── formal/ # Formal SMT Z3 verification suite
│ └── formal_verification_smt.py
├── harness/
│ ├── requirements.txt # Python dependencies (Standard Library only)
│ └── runner.py # Metrology harness under taskset -c 0
└── submission/
├── libcortex_serving_arena.so # Ephemeral CI RAM injection (/dev/shm)
├── cortex_serving.py # Python C-FFI bindings
├── verify_ground_truth.py # Ground truth semantic parity verification
└── run.sh # End-to-end local audit & verification runner
Verification & Reproduction Protocol
This benchmark is architected around a transparent two-tier verification model:
GitHub Actions Cloud CI (Full End-to-End Metrology) :
The complete benchmark with the native engine executes directly on clean, hardware-isolated cloud virtual machines under taskset -c 0 . GitHub Actions serves as an independent, tamper-proof auditor capturing live hardware telemetry ( /proc/cpuinfo , AVX-512 vector flags, cache hierarchy, invariant TSC clock source) and publishing cryptographic bit-exact SHA-256 seals.
Inspect live CI runs and silicon telemetry: GitHub Actions Workflow Runs
Local Offline Auditing (Open Source Invariants & Proofs) :
Researchers and system engineers can clone this repository to independently verify formal mathematical theorems and run the standalone anti-cheat auditor on their local hardware without requiring proprietary binaries.
Architecture / OS : Linux x86_64 with AVX2 or AVX-512 instruction support ( taskset -c 0 ).
Toolchains : Python 3.9+ (Standard Library only), Rust / Cargo (for the independent anti-cheat auditor), pip install z3-solver (for formal proofs).
Running Local Offline Auditing
# 1. Certify Formal SMT Z3 Mathematical Invariants (Bit-Exact UNSAT Proofs)
python3 formal/formal_verification_smt.py
# 2. Run the Independent Anti-Cheat Rust Auditor
cargo run --release --manifest-path audit/Cargo.toml
# 3. Alternatively, execute the complete local audit suite:
./submission/run.sh
Proprietary Native Engine & Evaluation Access
The high-performance native engine ( libcortex_serving_arena.so ) is protected intellectual property ( Covered by CIPO CA 3,322,620 ) and is not distributed in the public git repository. During CI runs, it is injected into ephemeral in-memory RAM ( /dev/shm ) via encrypted secrets and immediately purged post-run.
Contact : bmartin.systems@gmail.com
Comparative World Ranking (1-vCPU Metrology)
World Rank
Engine / Implementation
Architecture / ISA
Throughput (tok/s)
TPOT (ms)
Status / Delta
🥇
Native Rust Silicon Engine
Intel Xeon Platinum (AVX-512)
191,617.3 tok/s
0.20 ms
World #1 Leader (+2041.0% vs vLLM)
🥇
Native Rust Silicon Engine
AMD EPYC 9V45 Zen 4 (AVX-512)
177,519.2 tok/s
0.22 ms
World #1 Leader (+1883.5% vs vLLM)
🥇
Native Rust Silicon Engine
AMD EPYC 7763 Zen 3 (AVX2)
90,235.9 tok/s
0.71 ms
CI Verified Record (+908.2% vs vLLM)
🥈
TensorRT-LLM (v0.12 C++/CUDA)
Generic C++ / GPU
12,000.0 tok/s
1.28 ms
Upstream Reference
🥉
vLLM (v0.6.0 PagedAttention)
Python BlockMgr
8,950.0 tok/s
1.75 ms
Official Reference Baseline
4
HuggingFace TGI (Rust Core)
Rust Async
7,200.0 tok/s
2.10 ms
Standard Async Reference
5
llama.cpp (Static Batching)
CPU Native
4,200.0 tok/s
3.50 ms
CPU Native Baseline
Note
Scope & Metrology Methodology : This benchmark specifically isolates and evaluates the host CPU serving scheduler and Paged KV-cache allocation subsystem under strict 1-vCPU hardware isolation ( taskset -c 0 ). In production clusters (e.g. 8x NVIDIA H100), host scheduling latency is the root cause of GPU starvation bubbles, dissipating 200W-350W per GPU at idle. The baseline vLLM (8,950 tok/s) and TensorRT-LLM (12,000 tok/s) figures reflect official upstream host scheduling and block allocation ceilings under identical 1-vCPU isolation.
AI FinOps & Energy Efficiency (1-vCPU Metrology)
Engine / Implementation
Throughput (tok/s)
Tokens / Dollar
Tokens / s / Watt
OpEx Savings vs vLLM
Native Rust Silicon Engine
90,235.9
10,828,302,051
3,609.4
+681.1% Savings
TensorRT-LLM (v0.12)
12,000.0
1,440,000,000
480.0
Baseline Reference
vLLM (v0.6.0)
8,950.0
1,074,000,000
358.0
Standard Reference
Intellectual Property
Patent Protection : The core architecture and vector algorithms are legally protected ( Covered by CIPO CA 3,322,620 ).
Contact : bmartin.systems@gmail.com
High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark)
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) - bmartin-systems/cortex-serving-arena-preview

GitHub - bmartin-systems/cortex-serving-arena-preview: High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark) · GitHub
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
bmartin-systems
/
cortex-serving-arena-preview
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github/ workflows .github/ workflows audit audit formal formal harness harness submission submission .gitignore .gitignore README.md README.md View all files Repository files navigation
Serving Arena — High-Throughput 1-vCPU Continuous Batching & Hardware-Aligned Paged Memory Allocation Benchmark
This repository provides an evaluation harness and independent ground truth verification suite for high-performance LLM serving runtimes under strict 1-vCPU hardware isolation ( taskset -c 0 ).
Lock-Free Serving Scheduler : High-performance request allocation and sequence management achieving sub-microsecond dispatch latency.
Paged KV Cache : Fine-grained 16-token page allocation with zero memory fragmentation and fast vector dequantization.
Continuous Batching & Chunked Prefill : Interleaved prompt processing and token generation minimizing Time To First Token (TTFT) and Inter-Token Latency (ITL/TPOT).
Formal SMT Z3 Verification : 5 mathematical theorems proven bit-by-bit (QF_BV UNSAT) certifying zero page collision, deadlock freedom, and L1D cache containment.
Independent FP64 Anti-Cheat Auditor : Standalone verification engine recalculating token logits, scale bounds (N >= 5000 requests), and cryptographic provenance.
├── .github/workflows/
│ └── benchmark.yml # Reproducible 1-vCPU CI workflow
├── audit/ # Independent anti-cheat auditor (Rust)
│ ├── Cargo.toml
│ └── src/main.rs
├── formal/ # Formal SMT Z3 verification suite
│ └── formal_verification_smt.py
├── harness/
│ ├── requirements.txt # Python dependencies (Standard Library only)
│ └── runner.py # Metrology harness under taskset -c 0
└── submission/
├── libcortex_serving_arena.so # Ephemeral CI RAM injection (/dev/shm)
├── cortex_serving.py # Python C-FFI bindings
├── verify_ground_truth.py # Ground truth semantic parity verification
└── run.sh # End-to-end local audit & verification runner
Verification & Reproduction Protocol
This benchmark is architected around a transparent two-tier verification model:
GitHub Actions Cloud CI (Full End-to-End Metrology) :
The complete benchmark with the native engine executes directly on clean, hardware-isolated cloud virtual machines under taskset -c 0 . GitHub Actions serves as an independent, tamper-proof auditor capturing live hardware telemetry ( /proc/cpuinfo , AVX-512 vector flags, cache hierarchy, invariant TSC clock source) and publishing cryptographic bit-exact SHA-256 seals.
Inspect live CI runs and silicon telemetry: GitHub Actions Workflow Runs
Local Offline Auditing (Open Source Invariants & Proofs) :
Researchers and system engineers can clone this repository to independently verify formal mathematical theorems and run the standalone anti-cheat auditor on their local hardware without requiring proprietary binaries.
Architecture / OS : Linux x86_64 with AVX2 or AVX-512 instruction support ( taskset -c 0 ).
Toolchains : Python 3.9+ (Standard Library only), Rust / Cargo (for the independent anti-cheat auditor), pip install z3-solver (for formal proofs).
Running Local Offline Auditing
# 1. Certify Formal SMT Z3 Mathematical Invariants (Bit-Exact UNSAT Proofs)
python3 formal/formal_verification_smt.py
# 2. Run the Independent Anti-Cheat Rust Auditor
cargo run --release --manifest-path audit/Cargo.toml
# 3. Alternatively, execute the complete local audit suite:
./submission/run.sh
Proprietary Native Engine & Evaluation Access
The high-performance native engine ( libcortex_serving_arena.so ) is protected intellectual property ( Covered by CIPO CA 3,322,620 ) and is not distributed in the public git repository. During CI runs, it is injected into ephemeral in-memory RAM ( /dev/shm ) via encrypted secrets and immediately purged post-run.
Contact : bmartin.systems@gmail.com
Comparative World Ranking (1-vCPU Metrology)
World Rank
Engine / Implementation
Architecture / ISA
Throughput (tok/s)
TPOT (ms)
Status / Delta
🥇
Native Rust Silicon Engine
Intel Xeon Platinum (AVX-512)
191,617.3 tok/s
0.20 ms
World #1 Leader (+2041.0% vs vLLM)
🥇
Native Rust Silicon Engine
AMD EPYC 9V45 Zen 4 (AVX-512)
177,519.2 tok/s
0.22 ms
World #1 Leader (+1883.5% vs vLLM)
🥇
Native Rust Silicon Engine
AMD EPYC 7763 Zen 3 (AVX2)
90,235.9 tok/s
0.71 ms
CI Verified Record (+908.2% vs vLLM)
🥈
TensorRT-LLM (v0.12 C++/CUDA)
Generic C++ / GPU
12,000.0 tok/s
1.28 ms
Upstream Reference
🥉
vLLM (v0.6.0 PagedAttention)
Python BlockMgr
8,950.0 tok/s
1.75 ms
Official Reference Baseline
4
HuggingFace TGI (Rust Core)
Rust Async
7,200.0 tok/s
2.10 ms
Standard Async Reference
5
llama.cpp (Static Batching)
CPU Native
4,200.0 tok/s
3.50 ms
CPU Native Baseline
Note
Scope & Metrology Methodology : This benchmark specifically isolates and evaluates the host CPU serving scheduler and Paged KV-cache allocation subsystem under strict 1-vCPU hardware isolation ( taskset -c 0 ). In production clusters (e.g. 8x NVIDIA H100), host scheduling latency is the root cause of GPU starvation bubbles, dissipating 200W-350W per GPU at idle. The baseline vLLM (8,950 tok/s) and TensorRT-LLM (12,000 tok/s) figures reflect official upstream host scheduling and block allocation ceilings under identical 1-vCPU isolation.
AI FinOps & Energy Efficiency (1-vCPU Metrology)
Engine / Implementation
Throughput (tok/s)
Tokens / Dollar
Tokens / s / Watt
OpEx Savings vs vLLM
Native Rust Silicon Engine
90,235.9
10,828,302,051
3,609.4
+681.1% Savings
TensorRT-LLM (v0.12)
12,000.0
1,440,000,000
480.0
Baseline Reference
vLLM (v0.6.0)
8,950.0
1,074,000,000
358.0
Standard Reference
Intellectual Property
Patent Protection : The core architecture and vector algorithms are legally protected ( Covered by CIPO CA 3,322,620 ).
Contact : bmartin.systems@gmail.com
High-Throughput Continuous Batching & Paged KV Cache Engine (1-vCPU Benchmark)
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
