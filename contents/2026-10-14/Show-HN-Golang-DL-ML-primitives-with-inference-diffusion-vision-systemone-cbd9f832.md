---
source: "https://github.com/computerex/dlgo"
hn_url: "https://news.ycombinator.com/item?id=50091939"
title: "Show HN: Golang DL/ML primitives with inference, diffusion, vision, systemone"
article_title: "GitHub - computerex/dlgo: Golang inference engine and deep learning primitives · GitHub"
image: "https://opengraph.githubassets.com/4b7247f4e0bc4e3a35bb54a1d87e0d2e4609ffe93aa01bdcb27387032858692c/computerex/dlgo"
author: "computerex"
captured_at: "2026-10-14T18:20:22Z"
capture_tool: "hn-digest"
hn_id: 50091939
score: 2
comments: 0
posted_at: "2026-10-14T17:53:56Z"
tags:
  - hacker-news
---

# Show HN: Golang DL/ML primitives with inference, diffusion, vision, systemone

- HN: [50091939](https://news.ycombinator.com/item?id=50091939)
- Source: [github.com](https://github.com/computerex/dlgo)
- Score: 2
- Comments: 0
- Posted: 2026-10-14T17:53:56Z

## Translation

Title: Show HN: Golang DL/ML primitives with inference, diffusion, vision, systemone
Article title: GitHub - computerex/dlgo: Golang inference engine and deep learning primitives · GitHub
Description: Golang inference engine and deep learning primitives - computerex/dlgo
HN text: I created a DL inference engine that slowly grew to have more and functionality. - Inference for a bunch of open weights architectures - Diffusion image generation fir Z ai turbo - Vision input into LLM for inference - Vulkan/fast CPU, optimized kernels - Playground for testing - Last but not least, system one fast inference mode with same API semantics as Jev It's MIT, so go nuts

Article text:
GitHub - computerex/dlgo: Golang inference engine and deep learning primitives · GitHub
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
computerex
dlgo Public Notifications You must be signed in to change notification settings
Star 40 ( 40 ) You must be signed in to star a repository
Golang inference engine and deep learning primitives
Readme Apache-2.0 license Activity Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
114 Commits 114 Commits Folders and files
adapters/ z-agent adapters/ z-agent audio audio bench bench blas blas cmd cmd core core decode decode docs docs examples examples format format frontend frontend gpu gpu grammar grammar layers layers memory memory mmap mmap models models ops ops quant quant scripts scripts server server systemone-train systemone-train testdata testdata vision vision .gitignore .gitignore LICENSE LICENSE README.md README.md coherence_all_models.json coherence_all_models.json dlgo.go dlgo.go go.mod go.mod go.sum go.sum jev-plan.md jev-plan.md jev-work-checklist.md jev-work-checklist.md qwen35_chat_template.ginja qwen35_chat_template.ginja qwen36_chat_template.ginja qwen36_chat_template.ginja regression_results.json regression_results.json show_results.py show_results.py test_coherence_all_models.py test_coherence_all_models.py test_coherency_long.py test_coherency_long.py test_regression.py test_regression.py work-summary.md work-summary.md View all files Repository files navigation
Fast LLM inference in pure Go. Run GGUF models locally on CPU or GPU — no Python, CUDA, or external inference runtime required.
# Install
go install github.com/computerex/dlgo/cmd/dlgo@latest
# Chat with a model (like ollama)
dlgo run model.gguf
# Chat with GPU acceleration
dlgo run model.gguf --gpu
# Start web server with UI
dlgo server --model model.gguf --gpu
Vision / multimodal inference
The repository includes a local multimodal path for an mmproj vision projector plus a compatible text GGUF backbone. dlgo-vision preprocesses an image, executes the projector, inserts the resulting image-embedding block at <image> , and generates directly from the combined sequence.
# Build (CPU path)
go build -ldflags " -linkmode internal " -o dlgo-vision ./cmd/dlgo-vision
# Describe an image. The projector output embedding width must equal the
# language model embedding width.
dlgo-vision \
-mmproj C: \m odels \m mproj.gguf \
-model C: \m odels \t ext-model.gguf \
-image C: \i mages \e xample.png \
-prompt " Describe this image in detail:\n<image> " \
-max-tokens 128
-image also accepts deterministic synthetic fixtures ( white , black , gray , cb , red , green , blue ) for projector and image-token-placement checks. cmd/mmproj-encode is a companion utility for encoding/inspecting projector output. The vision stack lives in vision/ and supports mmproj loading, CLIP-style preprocessing, patch/grid encoding, and projector output insertion into the language-model prompt.
The current implementation validates projector/text embedding-dimension compatibility at runtime and supports image grids needed by the embedding-generation path. It is an inference implementation, not a claim of compatibility with every vendor chat template or multimodal model family; validate the projector, text backbone, image prompt marker, and visual quality together for each model pairing.
In addition to language and vision-language inference, dlgo includes a local text-to-image diffusion pipeline. The current implementation targets Z-Image / Lumina2-style DiT checkpoints rather than claiming universal compatibility with every Stable Diffusion checkpoint family. It runs the complete pipeline locally:
prompt -> Qwen text encoder -> quantized DiT denoising -> VAE decode -> PNG
The implementation includes a GGUF DiT loader, scheduler/sampler, 3D RoPE, quantized-weight CPU inference, Vulkan GPU DiT execution, SafeTensors VAE decoding, memory-budget estimation, and PNG output.
# CPU inference
go build -ldflags " -linkmode internal " -o dlgo-image ./cmd/dlgo-image
dlgo-image \
--diffusion-model C: \m odels \z -image-turbo-Q4_K_M.gguf \
--vae C: \m odels \a e.safetensors \
--llm C: \m odels \q wen-text-encoder.gguf \
-p " a cinematic photograph of a red fox in snow " \
-o output.png -W 1024 -H 1024 --steps 8
# Vulkan-accelerated DiT inference
dlgo-image --gpu \
--diffusion-model C: \m odels \z -image-turbo-Q4_K_M.gguf \
--vae C: \m odels \a e.safetensors \
--llm C: \m odels \q wen-text-encoder.gguf \
-p " a cinematic photograph of a red fox in snow " -o output.png
dlgo-image validates dimensions and estimates peak memory before generation. The text encoder is released after prompt embedding so the DiT and VAE have the available memory budget. GPU acceleration applies to the DiT execution path; model compatibility, image quality, determinism, VRAM use, and throughput must be validated for each checkpoint/quantization/resolution combination. The diffusion implementation is in models/diffusion/ , with CPU/GPU correctness and integration tests alongside it.
Start an Ollama-style interactive chat session:
$ dlgo run qwen3.5-0.8b-q8_0.gguf --gpu
Loading qwen3.5-0.8b-q8_0.gguf (2.3s)
Model: qwen3.5
Params: 24 layers, 1024 dim, 16 heads, vocab 151936
Context: 8192 tokens
Backend: GPU (NVIDIA GeForce RTX 4070 Ti SUPER)
Sampling: temp=0.70 top-k=40 top-p=0.90
Type /help for commands, or start chatting.
>>> What is the capital of France?
The capital of France is Paris.
45.2 tok/s | 32 tokens | 0.7s
>>> /help
/help Show this help
/info Show model info
/clear Clear conversation history
/set temp N Set temperature
/set tokens N Set max tokens
/exit Quit
Run flags:
Start an HTTP server with a built-in chat web interface and an OpenAI-compatible API:
$ dlgo server --model model.gguf --gpu --port 8080
dlgo server v0.1.0
linux/amd64, 16 cores
Web UI: http://localhost:8080
API: http://localhost:8080/v1/chat/completions
Health: http://localhost:8080/health
The web UI lets you:
Load and unload models dynamically
Toggle between CPU and GPU backends
Adjust temperature, top-p, top-k, max tokens
Stream responses with live performance metrics
Use dlgo as a Go package for embedding inference in your applications:
model , err := dlgo . LoadLLM ( "model.gguf" )
if err != nil {
log . Fatal ( err )
}
// Single-turn chat
response , _ := model . Chat ( "You are helpful." , "What is Go?" )
fmt . Println ( response )
// Streaming
model . ChatStream ( "" , "Write a poem." , func ( token string ) {
fmt . Print ( token )
}, dlgo . WithMaxTokens ( 256 ))
Options: WithMaxTokens(n) , WithTemperature(t) , WithTopK(k) , WithTopP(p) , WithSeed(s) , WithGreedy()
# CPU only (portable static binary with internal linking)
go build -ldflags " -linkmode internal " -o dlgo ./cmd/dlgo/
# With Vulkan GPU support
go build -tags vulkan -ldflags " -linkmode internal " -o dlgo ./cmd/dlgo/
# Build the web frontend (requires Node.js)
cd frontend && npm install && npm run build && cd ..
# Run server with frontend
dlgo server --model model.gguf --gpu --frontend frontend/dist
Prerequisites
Vulkan SDK (optional, for GPU support) — install from vulkan.lunarg.com
Node.js 18+ (optional, for building the web frontend)
25+ quantization formats — Q4_0 through Q8_0, K-quants (Q2_K–Q8_K), I-quants (IQ1_S–IQ4_XS), MXFP4, F16, BF16, F32
Vulkan GPU inference — full Vulkan compute backend with quantized MatVec shaders, fused attention, dp4a integer dot products
Never-OOM GPU — automatic VRAM budget with graceful fallback to partial GPU + CPU
Mixture of Experts (MoE) — fused multi-expert GPU dispatch, GPU-side top-K routing, zero CPU-GPU sync
Hybrid SSM+Attention — Gated Delta Net recurrent layers (Qwen3.5, Qwen3-Coder-Next)
Multi-head Latent Attention (MLA) — compressed KV cache (DeepSeek-V2, GLM-4.7)
Fast CPU — AVX2/FMA/VNNI SIMD, parallel worker pools, batch prefill GEMM
Speech-to-text — Whisper transcription
Voice activity detection — Silero VAD
Architecture
Example Models
CPU tok/s
GPU tok/s
LLaMA
Llama 3.2 1B, TinyLlama 1.1B
52–65
314–422
Qwen2/3
Qwen 2.5 0.5B, Qwen3 0.6B
60–98
241–411
Qwen3 MoE
Qwen3-Coder-30B-A3B (128 experts)
~5.2
~40
Qwen3.5
Qwen3.5 0.8B–27B (hybrid GDN+attention)
2.4–34
19–287
Qwen3.5 MoE
Qwen3.5 35B-A3B, 122B-A10B
1.4–4.1
2–11
GLM-4.7
GLM-4.7 Flash (MLA + MoE)
~5.6
~15
gpt-oss
gpt-oss-20b (MoE, attention sinks)
4.5–5.6
33–52
Gemma 2/3
Gemma 3 270M–1B
44–154
249–530
SmolLM2
SmolLM2 360M–1.7B
42–96
177–411
Phi
Phi-2, Phi-4-mini
9–20
~125
Whisper
Tiny, Base, Small
~1x RT
—
GPU benchmarks on NVIDIA RTX 4070 Ti SUPER. CPU with AVX2+FMA SIMD.
Same GGUF file, temperature=0 , seed=42 , max_tokens=64 .
Model
Quant
dlgo
Ollama
Delta
TinyLlama 1.1B
Q4_0
423 tok/s
187 tok/s
+126%
Gemma 3 1B
Q4_K_M
245 tok/s
116 tok/s
+111%
Qwen 2.5 0.5B
Q4_K_M
394 tok/s
237 tok/s
+66%
GPU (Vulkan vs Ollama CUDA)
Model
Quant
dlgo Vulkan
Ollama CUDA
Delta
Qwen3.5 0.8B
Q8_0
287 tok/s
250 tok/s
+15%
Qwen3.5 27B
Q3_K_M
6.4 tok/s
4.0 tok/s
+60%
CPU
Model
Quant
dlgo
Ollama
Delta
Qwen3.5 0.8B
Q8_0
33.2 tok/s
27.3 tok/s
+22%
Qwen3.5 9B
Q3_K_M
7.5 tok/s
7.2 tok/s
+4%
Qwen3.5 27B
Q3_K_M
2.5 tok/s
2.4 tok/s
+4%
Qwen3.5 35B MoE
Q3_K_M
8.1 tok/s
7.6 tok/s
+7%
System One: Fast Typed Decisions
Alongside chat generation, dlgo has a local System One API for deterministic closed-set decisions: choice selection, binary verification ( noul ), and ordered scores. It is designed to mirror the typed request/response shape used by TypeSafe Jev while keeping inference local.
Scope and calibration: System One head accuracy is only established for the evaluated distribution. A high binary probability is a preference between the two decision candidates, not a universal guarantee of truth. Evaluate and calibrate each application domain before automating consequential actions.
POST /v1/systemone takes a loaded model, an on-disk readout artifact, a state, and one or more questions. The response contains typed answers and Server-Timing ( queue , decision ) measurements.
curl http://127.0.0.1:8080/v1/systemone \
-H " Content-Type: application/json " \
-d ' {
"model": "qwen35-9b",
"readout": "C:\\projects\\dlgo\\systemone-train\\artifacts\\qwen35-9b-q3-readout-v2-uni.json",
"state": "A standard Pokemon type chart applies. Moltres is a Fire/Flying-type Pokemon.",
"questions": {
"moltres_water": {
"type": "noul",
"instructions": "Is Moltres, a Fire/Flying-type, weak to Water-type attacks?",
"criteria": {
"true": "Water attacks deal super-effective damage to Moltres.",
"false": "Water attacks deal neutral, reduced, or no damage to Moltres."
}
}
}
} '
For noul , use semantic true / false criteria whenever possible. They are compiled as the candidate meanings while the public response remains the probability of yes :
{ "answers" :{ "moltres_water" :{ "type" : " noul " ,

[truncated]

## Original Extract

Golang inference engine and deep learning primitives - computerex/dlgo

I created a DL inference engine that slowly grew to have more and functionality. - Inference for a bunch of open weights architectures - Diffusion image generation fir Z ai turbo - Vision input into LLM for inference - Vulkan/fast CPU, optimized kernels - Playground for testing - Last but not least, system one fast inference mode with same API semantics as Jev It's MIT, so go nuts

GitHub - computerex/dlgo: Golang inference engine and deep learning primitives · GitHub
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
computerex
dlgo Public Notifications You must be signed in to change notification settings
Star 40 ( 40 ) You must be signed in to star a repository
Golang inference engine and deep learning primitives
Readme Apache-2.0 license Activity Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
114 Commits 114 Commits Folders and files
adapters/ z-agent adapters/ z-agent audio audio bench bench blas blas cmd cmd core core decode decode docs docs examples examples format format frontend frontend gpu gpu grammar grammar layers layers memory memory mmap mmap models models ops ops quant quant scripts scripts server server systemone-train systemone-train testdata testdata vision vision .gitignore .gitignore LICENSE LICENSE README.md README.md coherence_all_models.json coherence_all_models.json dlgo.go dlgo.go go.mod go.mod go.sum go.sum jev-plan.md jev-plan.md jev-work-checklist.md jev-work-checklist.md qwen35_chat_template.ginja qwen35_chat_template.ginja qwen36_chat_template.ginja qwen36_chat_template.ginja regression_results.json regression_results.json show_results.py show_results.py test_coherence_all_models.py test_coherence_all_models.py test_coherency_long.py test_coherency_long.py test_regression.py test_regression.py work-summary.md work-summary.md View all files Repository files navigation
Fast LLM inference in pure Go. Run GGUF models locally on CPU or GPU — no Python, CUDA, or external inference runtime required.
# Install
go install github.com/computerex/dlgo/cmd/dlgo@latest
# Chat with a model (like ollama)
dlgo run model.gguf
# Chat with GPU acceleration
dlgo run model.gguf --gpu
# Start web server with UI
dlgo server --model model.gguf --gpu
Vision / multimodal inference
The repository includes a local multimodal path for an mmproj vision projector plus a compatible text GGUF backbone. dlgo-vision preprocesses an image, executes the projector, inserts the resulting image-embedding block at <image> , and generates directly from the combined sequence.
# Build (CPU path)
go build -ldflags " -linkmode internal " -o dlgo-vision ./cmd/dlgo-vision
# Describe an image. The projector output embedding width must equal the
# language model embedding width.
dlgo-vision \
-mmproj C: \m odels \m mproj.gguf \
-model C: \m odels \t ext-model.gguf \
-image C: \i mages \e xample.png \
-prompt " Describe this image in detail:\n<image> " \
-max-tokens 128
-image also accepts deterministic synthetic fixtures ( white , black , gray , cb , red , green , blue ) for projector and image-token-placement checks. cmd/mmproj-encode is a companion utility for encoding/inspecting projector output. The vision stack lives in vision/ and supports mmproj loading, CLIP-style preprocessing, patch/grid encoding, and projector output insertion into the language-model prompt.
The current implementation validates projector/text embedding-dimension compatibility at runtime and supports image grids needed by the embedding-generation path. It is an inference implementation, not a claim of compatibility with every vendor chat template or multimodal model family; validate the projector, text backbone, image prompt marker, and visual quality together for each model pairing.
In addition to language and vision-language inference, dlgo includes a local text-to-image diffusion pipeline. The current implementation targets Z-Image / Lumina2-style DiT checkpoints rather than claiming universal compatibility with every Stable Diffusion checkpoint family. It runs the complete pipeline locally:
prompt -> Qwen text encoder -> quantized DiT denoising -> VAE decode -> PNG
The implementation includes a GGUF DiT loader, scheduler/sampler, 3D RoPE, quantized-weight CPU inference, Vulkan GPU DiT execution, SafeTensors VAE decoding, memory-budget estimation, and PNG output.
# CPU inference
go build -ldflags " -linkmode internal " -o dlgo-image ./cmd/dlgo-image
dlgo-image \
--diffusion-model C: \m odels \z -image-turbo-Q4_K_M.gguf \
--vae C: \m odels \a e.safetensors \
--llm C: \m odels \q wen-text-encoder.gguf \
-p " a cinematic photograph of a red fox in snow " \
-o output.png -W 1024 -H 1024 --steps 8
# Vulkan-accelerated DiT inference
dlgo-image --gpu \
--diffusion-model C: \m odels \z -image-turbo-Q4_K_M.gguf \
--vae C: \m odels \a e.safetensors \
--llm C: \m odels \q wen-text-encoder.gguf \
-p " a cinematic photograph of a red fox in snow " -o output.png
dlgo-image validates dimensions and estimates peak memory before generation. The text encoder is released after prompt embedding so the DiT and VAE have the available memory budget. GPU acceleration applies to the DiT execution path; model compatibility, image quality, determinism, VRAM use, and throughput must be validated for each checkpoint/quantization/resolution combination. The diffusion implementation is in models/diffusion/ , with CPU/GPU correctness and integration tests alongside it.
Start an Ollama-style interactive chat session:
$ dlgo run qwen3.5-0.8b-q8_0.gguf --gpu
Loading qwen3.5-0.8b-q8_0.gguf (2.3s)
Model: qwen3.5
Params: 24 layers, 1024 dim, 16 heads, vocab 151936
Context: 8192 tokens
Backend: GPU (NVIDIA GeForce RTX 4070 Ti SUPER)
Sampling: temp=0.70 top-k=40 top-p=0.90
Type /help for commands, or start chatting.
>>> What is the capital of France?
The capital of France is Paris.
45.2 tok/s | 32 tokens | 0.7s
>>> /help
/help Show this help
/info Show model info
/clear Clear conversation history
/set temp N Set temperature
/set tokens N Set max tokens
/exit Quit
Run flags:
Start an HTTP server with a built-in chat web interface and an OpenAI-compatible API:
$ dlgo server --model model.gguf --gpu --port 8080
dlgo server v0.1.0
linux/amd64, 16 cores
Web UI: http://localhost:8080
API: http://localhost:8080/v1/chat/completions
Health: http://localhost:8080/health
The web UI lets you:
Load and unload models dynamically
Toggle between CPU and GPU backends
Adjust temperature, top-p, top-k, max tokens
Stream responses with live performance metrics
Use dlgo as a Go package for embedding inference in your applications:
model , err := dlgo . LoadLLM ( "model.gguf" )
if err != nil {
log . Fatal ( err )
}
// Single-turn chat
response , _ := model . Chat ( "You are helpful." , "What is Go?" )
fmt . Println ( response )
// Streaming
model . ChatStream ( "" , "Write a poem." , func ( token string ) {
fmt . Print ( token )
}, dlgo . WithMaxTokens ( 256 ))
Options: WithMaxTokens(n) , WithTemperature(t) , WithTopK(k) , WithTopP(p) , WithSeed(s) , WithGreedy()
# CPU only (portable static binary with internal linking)
go build -ldflags " -linkmode internal " -o dlgo ./cmd/dlgo/
# With Vulkan GPU support
go build -tags vulkan -ldflags " -linkmode internal " -o dlgo ./cmd/dlgo/
# Build the web frontend (requires Node.js)
cd frontend && npm install && npm run build && cd ..
# Run server with frontend
dlgo server --model model.gguf --gpu --frontend frontend/dist
Prerequisites
Vulkan SDK (optional, for GPU support) — install from vulkan.lunarg.com
Node.js 18+ (optional, for building the web frontend)
25+ quantization formats — Q4_0 through Q8_0, K-quants (Q2_K–Q8_K), I-quants (IQ1_S–IQ4_XS), MXFP4, F16, BF16, F32
Vulkan GPU inference — full Vulkan compute backend with quantized MatVec shaders, fused attention, dp4a integer dot products
Never-OOM GPU — automatic VRAM budget with graceful fallback to partial GPU + CPU
Mixture of Experts (MoE) — fused multi-expert GPU dispatch, GPU-side top-K routing, zero CPU-GPU sync
Hybrid SSM+Attention — Gated Delta Net recurrent layers (Qwen3.5, Qwen3-Coder-Next)
Multi-head Latent Attention (MLA) — compressed KV cache (DeepSeek-V2, GLM-4.7)
Fast CPU — AVX2/FMA/VNNI SIMD, parallel worker pools, batch prefill GEMM
Speech-to-text — Whisper transcription
Voice activity detection — Silero VAD
Architecture
Example Models
CPU tok/s
GPU tok/s
LLaMA
Llama 3.2 1B, TinyLlama 1.1B
52–65
314–422
Qwen2/3
Qwen 2.5 0.5B, Qwen3 0.6B
60–98
241–411
Qwen3 MoE
Qwen3-Coder-30B-A3B (128 experts)
~5.2
~40
Qwen3.5
Qwen3.5 0.8B–27B (hybrid GDN+attention)
2.4–34
19–287
Qwen3.5 MoE
Qwen3.5 35B-A3B, 122B-A10B
1.4–4.1
2–11
GLM-4.7
GLM-4.7 Flash (MLA + MoE)
~5.6
~15
gpt-oss
gpt-oss-20b (MoE, attention sinks)
4.5–5.6
33–52
Gemma 2/3
Gemma 3 270M–1B
44–154
249–530
SmolLM2
SmolLM2 360M–1.7B
42–96
177–411
Phi
Phi-2, Phi-4-mini
9–20
~125
Whisper
Tiny, Base, Small
~1x RT
—
GPU benchmarks on NVIDIA RTX 4070 Ti SUPER. CPU with AVX2+FMA SIMD.
Same GGUF file, temperature=0 , seed=42 , max_tokens=64 .
Model
Quant
dlgo
Ollama
Delta
TinyLlama 1.1B
Q4_0
423 tok/s
187 tok/s
+126%
Gemma 3 1B
Q4_K_M
245 tok/s
116 tok/s
+111%
Qwen 2.5 0.5B
Q4_K_M
394 tok/s
237 tok/s
+66%
GPU (Vulkan vs Ollama CUDA)
Model
Quant
dlgo Vulkan
Ollama CUDA
Delta
Qwen3.5 0.8B
Q8_0
287 tok/s
250 tok/s
+15%
Qwen3.5 27B
Q3_K_M
6.4 tok/s
4.0 tok/s
+60%
CPU
Model
Quant
dlgo
Ollama
Delta
Qwen3.5 0.8B
Q8_0
33.2 tok/s
27.3 tok/s
+22%
Qwen3.5 9B
Q3_K_M
7.5 tok/s
7.2 tok/s
+4%
Qwen3.5 27B
Q3_K_M
2.5 tok/s
2.4 tok/s
+4%
Qwen3.5 35B MoE
Q3_K_M
8.1 tok/s
7.6 tok/s
+7%
System One: Fast Typed Decisions
Alongside chat generation, dlgo has a local System One API for deterministic closed-set decisions: choice selection, binary verification ( noul ), and ordered scores. It is designed to mirror the typed request/response shape used by TypeSafe Jev while keeping inference local.
Scope and calibration: System One head accuracy is only established for the evaluated distribution. A high binary probability is a preference between the two decision candidates, not a universal guarantee of truth. Evaluate and calibrate each application domain before automating consequential actions.
POST /v1/systemone takes a loaded model, an on-disk readout artifact, a state, and one or more questions. The response contains typed answers and Server-Timing ( queue , decision ) measurements.
curl http://127.0.0.1:8080/v1/systemone \
-H " Content-Type: application/json " \
-d ' {
"model": "qwen35-9b",
"readout": "C:\\projects\\dlgo\\systemone-train\\artifacts\\qwen35-9b-q3-readout-v2-uni.json",
"state": "A standard Pokemon type chart applies. Moltres is a Fire/Flying-type Pokemon.",
"questions": {
"moltres_water": {
"type": "noul",
"instructions": "Is Moltres, a Fire/Flying-type, weak to Water-type attacks?",
"criteria": {
"true": "Water attacks deal super-effective damage to Moltres.",
"false": "Water attacks deal neutral, reduced, or no damage to Moltres."
}
}
}
} '
For noul , use semantic true / false criteria whenever possible. They are compiled as the candidate meanings while the public response remains the probability of yes :
{ "answers" :{ "moltres_water" :{ "type" : " noul " ,

[truncated]
