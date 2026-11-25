---
source: "https://github.com/mikewolfd/fast-silicon"
hn_url: "https://news.ycombinator.com/item?id=50637917"
title: "Skip cloud GPUs: batch ML locally on Mac GPU and Neural Engine (up to 11x MLX)"
article_title: "GitHub - mikewolfd/fast-silicon: Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace · GitHub"
image: "https://opengraph.githubassets.com/48c1daa293e6a8e17a41b8bfa26c84a8c5e6dc5837d6ca851f940501ac84a568/mikewolfd/fast-silicon"
author: "mikewolfd"
captured_at: "2026-11-25T03:29:34Z"
capture_tool: "hn-digest"
hn_id: 50637917
score: 2
comments: 0
posted_at: "2026-11-25T03:08:03Z"
tags:
  - hacker-news
---

# Skip cloud GPUs: batch ML locally on Mac GPU and Neural Engine (up to 11x MLX)

- HN: [50637917](https://news.ycombinator.com/item?id=50637917)
- Source: [github.com](https://github.com/mikewolfd/fast-silicon)
- Score: 2
- Comments: 0
- Posted: 2026-11-25T03:08:03Z

## Translation

Title: Skip cloud GPUs: batch ML locally on Mac GPU and Neural Engine (up to 11x MLX)
Article title: GitHub - mikewolfd/fast-silicon: Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace · GitHub
Description: Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace - mikewolfd/fast-silicon

Article text:
GitHub - mikewolfd/fast-silicon: Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
mikewolfd
fast-silicon Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
96 Commits 96 Commits Folders and files
.agents/ plugins .agents/ plugins configs configs docs docs packages packages scripts scripts tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Local Apple Silicon engines for bulk embeddings, document OCR, decision scoring, and schema-constrained JSON extraction. Built to replace cloud ML API bills at batch scale by running the Metal GPU and 16-core Apple Neural Engine (ANE) concurrently, with zero telemetry, no account, and 100% offline execution after weight download.
Supported: Apple Silicon Macs (verified on 48 GB M4 Pro macOS 26.6 and 24 GB base M5 macOS 27.2.1), macOS 15+ for Core ML ANE bundles and native runners ( macOS 14+ for GPU-only MLX), Python 3.12+, and uv .
Unsupported: Windows, Linux, and Intel Macs (ANE runners and Apple Vision require macOS Apple Silicon frameworks). Xcode Command Line Tools are required to compile native ANE .dylib runners ( make -C packages/fast-bge-small native , optional when passing default_mode="gpu" ); inference needs neither torch nor transformers .
Quickstart ( fast-bge-small , ~0.15-0.35 GB prebuilt download)
git clone https://github.com/mikewolfd/fast-silicon
cd fast-silicon
uv sync --python 3.13 --package fast-bge-small --inexact
make -C packages/fast-bge-small native
.venv/bin/bge-small setup-weights
from fast_bge_small import FastBgeSmall
passages = [
"Ocean currents carry heat from the tropics toward the poles." ,
"A piano has 88 keys." ,
]
engine = FastBgeSmall ( default_mode = "hetero" )
query = engine . encode_texts ( "How does heat move through the ocean?" , add_query_prefix = True )
vectors = engine . encode_texts ( passages )
print ( passages [ int (( vectors @ query ). argmax ())])
To install multiple packages or build tools, run uv sync --extra dev --extra extract --extra build and preview weight downloads (~36-38 GB full across all 8 model packages, or ~25.6 GB with --skip-test-extras ) with .venv/bin/python scripts/setup_all.py --dry-run .
Stock MLX uses the Metal GPU and leaves the 16-core ANE idle, while Core ML uses the ANE and leaves the GPU idle. Most local runners are also tuned for single-user B=1 chat instead of offline batch ETL jobs that share the same system prompt and JSON schema across thousands of rows.
Variable-length batch (sorted by token count)
├── longest sequences ──▶ Metal GPU (packed varlen / custom kernels) ──┐
│ ├──▶ merged outputs in input order
└── shortest sequences ◀─ 16-core ANE (multifunction Core ML ring) ───┘
Pair that dual-accelerator scheduler with zero-padding sequence packing, shared system-prompt KV caches ( PrefixTree / Hydragen v2 ), INT8/binary vector indexes, and schema-enforced fast-forward decoding, and throughput scales cleanly inside unified memory.
Workload
Package
What you get
Prebuilt download
Search short English passages
fast-bge-small
384d vectors; 0.49 ms warm query (M4 Pro), 200.6k tok/s bulk ( 11.3x stock MLX).
~0.15-0.35 GB
Index 2K passages & page images
fast-embeddinggemma-2
Up to 2,048 tokens (128-768d); 48.8k tok/s ( 4.3-4.8x MLX, 3.3x Core ML), 19.0 img/s lazy vision.
~0.30-0.65 GB
Contextual passage & 32K search
fast-pplx
Contextual 4K chunking lifts legal nDCG from 0.359 (isolated) to 0.703, matching stock pplx-context-mlx at 4.4x speed; MaxSim covers 4K-32K docs.
~1.8-2.4 GB
Score decisions & rerank choices
fast-decision-models
100% Pure-ANE 421M Laya ( 5.1 ms query, 1,400 items in 37.9 s hetero) + Kev-0.8B/4B with prefix cache.
~1.3-4.9 GB
High-throughput JSON & tools (2B)
fast-gemma4
10.5 valid docs/s at B=128 ( 5.2x mlx-vlm valid rate across B=128 vs B=64; 150/150 valid at 0.946 F1; 4.8-6.3 GB MLX).
~4.2-5.4 GB (~10.4-17.0 GB with test extras)
Reasoning, tools & 16K JSON (4B/9B)
fast-qwen3
Hybrid GatedDeltaNet + 4-bit QuaRot KV ( 2.1x faster at 16K); 300/300 valid on 300 JSON/tool tasks.
~3.3 GB
Extract Markdown & tables from PDFs
fast-ocr
Clean PDF bypass ( 139-220 pp/s ), parallel Apple Vision ANE ( 5-9.6 pp/s ), and OvisOCR2-8bit VLM ( 1.5-3.5x ).
~2.60 GB
Batch 20B MXFP4 MoE generation
fast-gpt-oss
gpt-oss-20b verified across 24-48 GB+ Macs ( 103 tok/s @ B=16 on 24 GB M5; 444 tok/s @ B=128 on 48 GB; auto-clamps on 16 GB).
~12.10 GB
Shared runtime & compact index
fast-shared
Dynamic chip dispatch, CompactRetrievalIndex (0.50 ms @ 100k vecs), PrefixTree , KV quant, and Parquet ETL.
-
Cross-silicon benchmarks ( M4 Pro 48 GB vs. Base M5 24 GB )
All runs execute locally after warmup with deterministic accuracy verification on pinned public benchmarks. See docs/BENCHMARKS.md for datasets, stock configs, hardware ceilings, and CLI reproducers.
Sub-millisecond vector search: CompactRetrievalIndex pairs 1-bit ARM64 NEON Hamming prefiltering with symmetric INT8 reranking ( 4-6x smaller than FP32) to search 100,000 vectors in 0.50-0.66 ms/query ( 3.0-4.4x faster than FP32 BLAS) with zero ranking loss ( index details ).
Local post-training: Packed 8-bit QLoRA runs at 80.5-81.6 supervised tok/s on 24 GB M5 ( pack=1024 , 5.2-7.3 GB peak) and 546 (packed) / 1,385 (shared-prefix) logical input tok/s on 48 GB M4 Pro ( pack=2048 ), while disaggregated GRPO runs Gemma 4 rollouts on Metal GPU concurrently with Laya critic scoring on the ANE ( 2.6-3.4 rollouts/s M5, <0.4 ms exposed critic time) ( training benchmarks ).
Prebuilt Hugging Face artifacts
Pre-compiled .mlmodelc Core ML bundles and pre-quantized MLX checkpoints skip local model compilation:
Single-model RAM vs. standalone Core ML: Running Metal GPU and ANE together consumes more unified memory than a standalone 8-bit Core ML export (for instance, 3.43 GB vs. 0.84 GB on EmbeddingGemma-2). Use mode="ane" for minimal RAM when latency allows.
Batch HTTP serving: OpenAI-compatible HTTP servers collect microbatches and return completed responses rather than token-by-token SSE streams.
Large-batch 20B memory: fast-gpt-oss runs cleanly on 24 GB Macs ( 13.4-14.3 GB peak, auto-clamped to B=16 ), while large B=64-128 batch sweeps want 32-48 GB+ RAM. There is no matched-prompt/batch B=128 stock mlx-lm throughput baseline, and on 25 tool tasks the recorded MXFP4-Q4 run trailed unconstrained mlx-lm accuracy ( 9/25 vs 14/25 ).
Grammar lift vs. kernel speed: On structured JSON benchmarks against unconstrained stock mlx-lm / mlx-vlm , llguidance grammar enforcement eliminates malformed JSON alongside kernel speedups; package READMEs report raw speed and validity separately.
Hardware coverage: Verified continuously on M4 Pro and base M5; M1-M3 Macs work via the same MLX/Core ML paths but are not in our hardware CI matrix.
Prior Art, Development & License
Compared to mlx-lm / mlx-vlm and llama.cpp : Those excel at general interactive chat and broad GGUF/MLX model coverage; fast-silicon targets high-volume Apple Silicon batch pipelines with ANE offloading, shared-prefix KV reuse, and compact retrieval.
Compared to public Core ML exports & Apple CoreAI: Standalone Core ML models save RAM on single models, while fast-silicon adds long-context GPU fallback ( 2K-32K ), batch work-stealing, and unified Python/Parquet CLIs.
Acknowledgments: zm2231/pplx-context-mlx (MIT) inspired fast-pplx memmap embeddings, tile-aligned slicing, and parity testing. Built on mlx , mlx-lm , coremltools , llguidance , pypdfium2 , and upstream weights from BAAI, Google, Qwen, Perplexity, OpenAI, and ConvAI.
Code is MIT licensed (downloaded upstream model weights remain subject to their original licenses). For contributor workflows, sibling repo syncing, and test suites, see CONTRIBUTING.md .
Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace - mikewolfd/fast-silicon

GitHub - mikewolfd/fast-silicon: Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
mikewolfd
fast-silicon Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
96 Commits 96 Commits Folders and files
.agents/ plugins .agents/ plugins configs configs docs docs packages packages scripts scripts tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Local Apple Silicon engines for bulk embeddings, document OCR, decision scoring, and schema-constrained JSON extraction. Built to replace cloud ML API bills at batch scale by running the Metal GPU and 16-core Apple Neural Engine (ANE) concurrently, with zero telemetry, no account, and 100% offline execution after weight download.
Supported: Apple Silicon Macs (verified on 48 GB M4 Pro macOS 26.6 and 24 GB base M5 macOS 27.2.1), macOS 15+ for Core ML ANE bundles and native runners ( macOS 14+ for GPU-only MLX), Python 3.12+, and uv .
Unsupported: Windows, Linux, and Intel Macs (ANE runners and Apple Vision require macOS Apple Silicon frameworks). Xcode Command Line Tools are required to compile native ANE .dylib runners ( make -C packages/fast-bge-small native , optional when passing default_mode="gpu" ); inference needs neither torch nor transformers .
Quickstart ( fast-bge-small , ~0.15-0.35 GB prebuilt download)
git clone https://github.com/mikewolfd/fast-silicon
cd fast-silicon
uv sync --python 3.13 --package fast-bge-small --inexact
make -C packages/fast-bge-small native
.venv/bin/bge-small setup-weights
from fast_bge_small import FastBgeSmall
passages = [
"Ocean currents carry heat from the tropics toward the poles." ,
"A piano has 88 keys." ,
]
engine = FastBgeSmall ( default_mode = "hetero" )
query = engine . encode_texts ( "How does heat move through the ocean?" , add_query_prefix = True )
vectors = engine . encode_texts ( passages )
print ( passages [ int (( vectors @ query ). argmax ())])
To install multiple packages or build tools, run uv sync --extra dev --extra extract --extra build and preview weight downloads (~36-38 GB full across all 8 model packages, or ~25.6 GB with --skip-test-extras ) with .venv/bin/python scripts/setup_all.py --dry-run .
Stock MLX uses the Metal GPU and leaves the 16-core ANE idle, while Core ML uses the ANE and leaves the GPU idle. Most local runners are also tuned for single-user B=1 chat instead of offline batch ETL jobs that share the same system prompt and JSON schema across thousands of rows.
Variable-length batch (sorted by token count)
├── longest sequences ──▶ Metal GPU (packed varlen / custom kernels) ──┐
│ ├──▶ merged outputs in input order
└── shortest sequences ◀─ 16-core ANE (multifunction Core ML ring) ───┘
Pair that dual-accelerator scheduler with zero-padding sequence packing, shared system-prompt KV caches ( PrefixTree / Hydragen v2 ), INT8/binary vector indexes, and schema-enforced fast-forward decoding, and throughput scales cleanly inside unified memory.
Workload
Package
What you get
Prebuilt download
Search short English passages
fast-bge-small
384d vectors; 0.49 ms warm query (M4 Pro), 200.6k tok/s bulk ( 11.3x stock MLX).
~0.15-0.35 GB
Index 2K passages & page images
fast-embeddinggemma-2
Up to 2,048 tokens (128-768d); 48.8k tok/s ( 4.3-4.8x MLX, 3.3x Core ML), 19.0 img/s lazy vision.
~0.30-0.65 GB
Contextual passage & 32K search
fast-pplx
Contextual 4K chunking lifts legal nDCG from 0.359 (isolated) to 0.703, matching stock pplx-context-mlx at 4.4x speed; MaxSim covers 4K-32K docs.
~1.8-2.4 GB
Score decisions & rerank choices
fast-decision-models
100% Pure-ANE 421M Laya ( 5.1 ms query, 1,400 items in 37.9 s hetero) + Kev-0.8B/4B with prefix cache.
~1.3-4.9 GB
High-throughput JSON & tools (2B)
fast-gemma4
10.5 valid docs/s at B=128 ( 5.2x mlx-vlm valid rate across B=128 vs B=64; 150/150 valid at 0.946 F1; 4.8-6.3 GB MLX).
~4.2-5.4 GB (~10.4-17.0 GB with test extras)
Reasoning, tools & 16K JSON (4B/9B)
fast-qwen3
Hybrid GatedDeltaNet + 4-bit QuaRot KV ( 2.1x faster at 16K); 300/300 valid on 300 JSON/tool tasks.
~3.3 GB
Extract Markdown & tables from PDFs
fast-ocr
Clean PDF bypass ( 139-220 pp/s ), parallel Apple Vision ANE ( 5-9.6 pp/s ), and OvisOCR2-8bit VLM ( 1.5-3.5x ).
~2.60 GB
Batch 20B MXFP4 MoE generation
fast-gpt-oss
gpt-oss-20b verified across 24-48 GB+ Macs ( 103 tok/s @ B=16 on 24 GB M5; 444 tok/s @ B=128 on 48 GB; auto-clamps on 16 GB).
~12.10 GB
Shared runtime & compact index
fast-shared
Dynamic chip dispatch, CompactRetrievalIndex (0.50 ms @ 100k vecs), PrefixTree , KV quant, and Parquet ETL.
-
Cross-silicon benchmarks ( M4 Pro 48 GB vs. Base M5 24 GB )
All runs execute locally after warmup with deterministic accuracy verification on pinned public benchmarks. See docs/BENCHMARKS.md for datasets, stock configs, hardware ceilings, and CLI reproducers.
Sub-millisecond vector search: CompactRetrievalIndex pairs 1-bit ARM64 NEON Hamming prefiltering with symmetric INT8 reranking ( 4-6x smaller than FP32) to search 100,000 vectors in 0.50-0.66 ms/query ( 3.0-4.4x faster than FP32 BLAS) with zero ranking loss ( index details ).
Local post-training: Packed 8-bit QLoRA runs at 80.5-81.6 supervised tok/s on 24 GB M5 ( pack=1024 , 5.2-7.3 GB peak) and 546 (packed) / 1,385 (shared-prefix) logical input tok/s on 48 GB M4 Pro ( pack=2048 ), while disaggregated GRPO runs Gemma 4 rollouts on Metal GPU concurrently with Laya critic scoring on the ANE ( 2.6-3.4 rollouts/s M5, <0.4 ms exposed critic time) ( training benchmarks ).
Prebuilt Hugging Face artifacts
Pre-compiled .mlmodelc Core ML bundles and pre-quantized MLX checkpoints skip local model compilation:
Single-model RAM vs. standalone Core ML: Running Metal GPU and ANE together consumes more unified memory than a standalone 8-bit Core ML export (for instance, 3.43 GB vs. 0.84 GB on EmbeddingGemma-2). Use mode="ane" for minimal RAM when latency allows.
Batch HTTP serving: OpenAI-compatible HTTP servers collect microbatches and return completed responses rather than token-by-token SSE streams.
Large-batch 20B memory: fast-gpt-oss runs cleanly on 24 GB Macs ( 13.4-14.3 GB peak, auto-clamped to B=16 ), while large B=64-128 batch sweeps want 32-48 GB+ RAM. There is no matched-prompt/batch B=128 stock mlx-lm throughput baseline, and on 25 tool tasks the recorded MXFP4-Q4 run trailed unconstrained mlx-lm accuracy ( 9/25 vs 14/25 ).
Grammar lift vs. kernel speed: On structured JSON benchmarks against unconstrained stock mlx-lm / mlx-vlm , llguidance grammar enforcement eliminates malformed JSON alongside kernel speedups; package READMEs report raw speed and validity separately.
Hardware coverage: Verified continuously on M4 Pro and base M5; M1-M3 Macs work via the same MLX/Core ML paths but are not in our hardware CI matrix.
Prior Art, Development & License
Compared to mlx-lm / mlx-vlm and llama.cpp : Those excel at general interactive chat and broad GGUF/MLX model coverage; fast-silicon targets high-volume Apple Silicon batch pipelines with ANE offloading, shared-prefix KV reuse, and compact retrieval.
Compared to public Core ML exports & Apple CoreAI: Standalone Core ML models save RAM on single models, while fast-silicon adds long-context GPU fallback ( 2K-32K ), batch work-stealing, and unified Python/Parquet CLIs.
Acknowledgments: zm2231/pplx-context-mlx (MIT) inspired fast-pplx memmap embeddings, tile-aligned slicing, and parity testing. Built on mlx , mlx-lm , coremltools , llguidance , pypdfium2 , and upstream weights from BAAI, Google, Qwen, Perplexity, OpenAI, and ConvAI.
Code is MIT licensed (downloaded upstream model weights remain subject to their original licenses). For contributor workflows, sibling repo syncing, and test suites, see CONTRIBUTING.md .
Hardware-limit Apple Silicon (M4 / M5+ Metal 4 MPP + 16-core ANE + CPU SME) ML inference and embedding workspace
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
