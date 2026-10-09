---
source: "https://github.com/SpectraSI/SpectraAdamw"
hn_url: "https://news.ycombinator.com/item?id=50018712"
title: "Show HN: Spectra – A drop-in optimizer that cuts LLM training VRAM by 50%"
article_title: "GitHub - SpectraSI/SpectraAdamw: Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression · GitHub"
image: "https://opengraph.githubassets.com/cca4cc16a8c492ac6087aa323a188bd461c8fc672bdc3a678a05fc41f754a57d/SpectraSI/SpectraAdamw"
author: "KanishkIndia"
captured_at: "2026-10-09T11:41:09Z"
capture_tool: "hn-digest"
hn_id: 50018712
score: 2
comments: 0
posted_at: "2026-10-09T10:51:42Z"
tags:
  - hacker-news
---

# Show HN: Spectra – A drop-in optimizer that cuts LLM training VRAM by 50%

- HN: [50018712](https://news.ycombinator.com/item?id=50018712)
- Source: [github.com](https://github.com/SpectraSI/SpectraAdamw)
- Score: 2
- Comments: 0
- Posted: 2026-10-09T10:51:42Z

## Translation

Title: Show HN: Spectra – A drop-in optimizer that cuts LLM training VRAM by 50%
Article title: GitHub - SpectraSI/SpectraAdamw: Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression · GitHub
Description: Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression - SpectraSI/SpectraAdamw

Article text:
GitHub - SpectraSI/SpectraAdamw: Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression · GitHub
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
SpectraSI
SpectraAdamw Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
LICENSE LICENSE README.md README.md View all files Repository files navigation
Authors: Spectra Labs Research
Status: Closed Beta (Python POC Available)
The primary memory bottleneck when fine-tuning Large Language Models (LLMs) on consumer hardware is not the model weights, but the optimizer states. Standard AdamW requires storing two massive state tensors (momentum and variance), costing 2x the model parameters in VRAM. Unstructured sparsity (magnitude pruning) and aggressive 8-bit quantization often lead to severe degradation in convergence stability.
We introduce SpectraAdamW , a hybrid optimizer utilizing factorized variance approximations and frequency-domain momentum compression to reduce the optimizer state overhead by exactly 50% (~1x parameters) while maintaining full-precision convergence behavior.
1. Factorized Variance Approximation (FVA)
Standard AdamW tracks the second moment (uncentered variance) of the gradients at full precision. SpectraAdamW decomposes this second moment into row and column means for 2D weight matrices (e.g., Linear layers).
By storing only $O(N \times 1)$ and $O(1 \times M)$ tensors instead of the full $O(N \times M)$ matrix, we effectively eliminate the memory footprint of the variance state.
To prevent the Gibbs ringing artifacts commonly associated with aggressive variance compression, the factorized variance is strictly reconstructed via scalar multiplication before the gradient update, ensuring the denominator remains strictly positive and bounded.
2. Frequency-Domain Momentum Compression (Native Backend)
While the variance is factorized, the first moment (momentum) retains heavy directional information. In our proprietary C++ backend, momentum is compressed by transforming gradients into the frequency domain via a Fast Fourier Transform (FFT). This isolates the high-energy signal from the noise, allowing us to dynamically mask low-impact frequencies and compress the state without losing structural gradient integrity.
In simulated 8192-dimension Transformer layer updates (Batch Size 128):
SpectraAdamW State VRAM: 128.16 MB
Loss curves demonstrate perfect alignment with AdamW convergence by step 40, avoiding the delayed convergence penalty seen in aggressive quantization methods.
We are currently distributing the PyTorch Minimum Viable Product (MVP) to verify the factorized variance math and VRAM reduction. The optimizer acts as a 1-line drop-in replacement for torch.optim.AdamW :
from spectra_optim import SpectraAdamW
# Example: Fine-tuning an 8B model
model = AutoModelForCausalLM . from_pretrained ( "meta-llama/Llama-3-8B" )
# Standard AdamW State VRAM: ~32GB
# SpectraAdamW State VRAM: ~16GB
optimizer = SpectraAdamW ( model . parameters (), lr = 1e-4 , weight_decay = 0.01 )
loss . backward ()
optimizer . step ()
Beta Access
The fully optimized C++ CUDA backend is currently in closed development. If you are conducting research on constrained hardware and wish to verify the Python POC in our private Google Colab environment, please reach out to the Spectra Labs team on Reddit at r/SpectraLabs .
Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression - SpectraSI/SpectraAdamw

GitHub - SpectraSI/SpectraAdamw: Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression · GitHub
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
SpectraSI
SpectraAdamw Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
LICENSE LICENSE README.md README.md View all files Repository files navigation
Authors: Spectra Labs Research
Status: Closed Beta (Python POC Available)
The primary memory bottleneck when fine-tuning Large Language Models (LLMs) on consumer hardware is not the model weights, but the optimizer states. Standard AdamW requires storing two massive state tensors (momentum and variance), costing 2x the model parameters in VRAM. Unstructured sparsity (magnitude pruning) and aggressive 8-bit quantization often lead to severe degradation in convergence stability.
We introduce SpectraAdamW , a hybrid optimizer utilizing factorized variance approximations and frequency-domain momentum compression to reduce the optimizer state overhead by exactly 50% (~1x parameters) while maintaining full-precision convergence behavior.
1. Factorized Variance Approximation (FVA)
Standard AdamW tracks the second moment (uncentered variance) of the gradients at full precision. SpectraAdamW decomposes this second moment into row and column means for 2D weight matrices (e.g., Linear layers).
By storing only $O(N \times 1)$ and $O(1 \times M)$ tensors instead of the full $O(N \times M)$ matrix, we effectively eliminate the memory footprint of the variance state.
To prevent the Gibbs ringing artifacts commonly associated with aggressive variance compression, the factorized variance is strictly reconstructed via scalar multiplication before the gradient update, ensuring the denominator remains strictly positive and bounded.
2. Frequency-Domain Momentum Compression (Native Backend)
While the variance is factorized, the first moment (momentum) retains heavy directional information. In our proprietary C++ backend, momentum is compressed by transforming gradients into the frequency domain via a Fast Fourier Transform (FFT). This isolates the high-energy signal from the noise, allowing us to dynamically mask low-impact frequencies and compress the state without losing structural gradient integrity.
In simulated 8192-dimension Transformer layer updates (Batch Size 128):
SpectraAdamW State VRAM: 128.16 MB
Loss curves demonstrate perfect alignment with AdamW convergence by step 40, avoiding the delayed convergence penalty seen in aggressive quantization methods.
We are currently distributing the PyTorch Minimum Viable Product (MVP) to verify the factorized variance math and VRAM reduction. The optimizer acts as a 1-line drop-in replacement for torch.optim.AdamW :
from spectra_optim import SpectraAdamW
# Example: Fine-tuning an 8B model
model = AutoModelForCausalLM . from_pretrained ( "meta-llama/Llama-3-8B" )
# Standard AdamW State VRAM: ~32GB
# SpectraAdamW State VRAM: ~16GB
optimizer = SpectraAdamW ( model . parameters (), lr = 1e-4 , weight_decay = 0.01 )
loss . backward ()
optimizer . step ()
Beta Access
The fully optimized C++ CUDA backend is currently in closed development. If you are conducting research on constrained hardware and wish to verify the Python POC in our private Google Colab environment, please reach out to the Spectra Labs team on Reddit at r/SpectraLabs .
Halving Optimizer State Memory via Factorized Variance and Frequency-Domain Compression
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
