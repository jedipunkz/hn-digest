---
source: "https://github.com/RunanywhereAI/trnrt"
hn_url: "https://news.ycombinator.com/item?id=50533270"
title: "Show HN: TrnRT – An open-source LLM inference engine for AWS Trainium"
article_title: "GitHub - RunanywhereAI/trnrt: The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. · GitHub"
image: "https://repository-images.githubusercontent.com/1440096236/8057fa03-91df-46f8-9d49-309f457392f0"
author: "sanchitmonga22"
captured_at: "2026-11-17T16:44:06Z"
capture_tool: "hn-digest"
hn_id: 50533270
score: 1
comments: 0
posted_at: "2026-11-17T16:01:04Z"
tags:
  - hacker-news
---

# Show HN: TrnRT – An open-source LLM inference engine for AWS Trainium

- HN: [50533270](https://news.ycombinator.com/item?id=50533270)
- Source: [github.com](https://github.com/RunanywhereAI/trnrt)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T16:01:04Z

## Translation

Title: Show HN: TrnRT – An open-source LLM inference engine for AWS Trainium
Article title: GitHub - RunanywhereAI/trnrt: The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. · GitHub
Description: The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. - RunanywhereAI/trnrt
HN text: Hi HN, Sanchit here from RunAnywhere (YC W26). We just open-sourced TrnRT: Trainium RunTime, our inference engine for AWS Trainium. We started 3 weeks back, and currently we’re able to serve and run GLM 5.3 flash (FP8) on AWS trn2.48xlarge, and will bring more models support soon. It’s open to contributions :) What it does today: - Serves 2 instances of GLM-5.3 flash on a single box, 8 trn2 with a cache aware router. - Can process 3.57M TPM (tokens per minute) on SemiAnalysis's AgentX benchmarks. - Runs decode as one NKI kernel per layer half, or for the 11 sparse-attention layers fuses both blocks into a single kernel, for faster performance. We started working on Trn2 because we wanted to explore it and see if we can come close to beating it’s performance in terms of cost efficiency compared to H100/H200. A big MoE with 1M context window such as GLM-5.3 flash requires a lot of memory, and memory/tb on a trn2.48xlarge costs less than half compared to 8x h200 on AWS, and hence we wanted to experiment with it and now we’re sharing some early results. But AWS's own stack didn't support the model: vLLM Neuron's latest release did not register any GLM model, and NxDI, AWS's own distributed inference framework, is no longer under active development. We also saw the drop of SGLang Neuron implementation yesterday, but there is no good support for Trn2, which is the only chip available for public use right now. We ran semiAnalysis's AgentX benchmarks, which replays real coding agent sessions: - 1 box served 28.9 RPM(requests per minute) with 32 agents at batch 12 (also known as concurrency) -The P50 TTFT with a cached prompt at 16 agents is ~1.24 seconds with a decode of 35.9 tps. Initially, the recurrence in one kda (linear attention) layer was ~2% of the layer's attention time at batch 1, and the rest was projections, transposes and XLA glue between small kernels. So we wrote the megakernel so that it covers the whole layer in 2 kernel dispatch, which made decode 25% faster at batch 8. The sparse attention kernel keep a layer’s working weights in SBUF, distributing weighted streaming across the cores and handling the interchip communication within There’s an approximation in the sparse-attention prefill, affecting about 1.2% of block selections. You can refer to the write-up for the profiling results: https://runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-mega... Please give it a try yourself on a trn2.48xlarge and reproduce the AgentX results by following the steps in docs/RESULTS.md. Happy to collaborate and hear any feedback. Deploying high performant dedicated inference at lab like performance and half the price for open models such as the GLM-5.3 family, Kimi K3, Deepseek flash etc. is how we make money. You can talk to us regarding the same at founders@runanywhere.ai. We hope you find this useful. If you are working with Trainium or have experience with NKI kernels, we'd love to get your honest feedback. Feel free to share the models you want support for next, and we welcome any contributions to the repository.

Article text:
GitHub - RunanywhereAI/trnrt: The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. · GitHub
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
RunanywhereAI
trnrt Public Notifications You must be signed in to change notification settings
Star 1.6k ( 1.6k ) You must be signed in to star a repository
The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere.
runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-megakernels Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
19 forks Report repository development Branches Tags Go to file Code Open more actions menu Latest commit
315 Commits 315 Commits Folders and files
.github .github docker docker docs docs tests tests tools tools trnrt trnrt .coderabbit.yaml .coderabbit.yaml .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md cubic.yaml cubic.yaml greptile.json greptile.json pyproject.toml pyproject.toml View all files Repository files navigation
Quickstart ·
Benchmarks ·
Docs ·
Kernels ·
Blog post ·
runanywhere.ai
TrnRT is an open inference engine for large language models on AWS Trainium. It serves an OpenAI-compatible API
from one trn2.48xlarge , runs decode as NKI megakernels (about one kernel per layer half), and keeps a prefix cache
across device and host memory. GLM-5.3-Flash is the first supported model; more open models are on the way.
How we built it, with the AgentX results: Serving GLM-5.3-Flash on one Trainium2 box with NKI layer megakernels .
Feature
What it does
NKI megakernels
One fused kernel per layer half for decode: KDA attention, the fused DSA layer, mHC and the MoE FFN.
1M-token context
Sparse attention (top 2,048) and linear attention, with fixed-size state.
Prefix cache
A radix tree over FP8 KV pages, plus a host-memory tier for evicted pages and state checkpoints.
Two engines per box
Each engine is tensor-parallel over 32 logical NeuronCores, behind a session-affine router.
OpenAI API
Chat and completions, streaming, tool calls and reasoning output.
Offline compile
Every graph is built on a CPU host; the serving box loads them and compiles nothing by default.
Quickstart
You need Docker, a trn2.48xlarge to serve on, and any x86_64 host to compile on. New to Trainium? Start with
docs/AWS.md .
1. Build the image (Neuron SDK 2.32 (neuronx-cc 2.27), Python 3.13)
git clone https://github.com/RunanywhereAI/trnrt.git && cd trnrt
docker build -f docker/Dockerfile \
--build-arg TRNRT_VERSION= $( python3 -c ' import trnrt; print(trnrt.__version__) ' ) \
--build-arg TRNRT_REVISION= $( git rev-parse HEAD ) -t trnrt .
2. Build the compile cache (any x86_64 host, no NeuronCore needed; takes hours)
OUT=/data/trnrt REV=eb9eb208eb0d988989d07a6a12d0fdeb5f52574a
mkdir -p $OUT /nki $OUT /work
trnrt () { docker run --rm -e HF_TOKEN -v $OUT : $OUT \
-v $OUT /nki:/var/tmp/nki-intermediate-cache trnrt " $@ " ; }
trnrt compile shapes zai-org/GLM-5.3-Flash --revision $REV --out $OUT
trnrt compile decisions --model zai-org/GLM-5.3-Flash --revision $REV --out $OUT
trnrt compile capture --out $OUT -- --model $OUT /shapes
trnrt compile compile --out $OUT --work-root $OUT /work
trnrt compile strip --out $OUT
3. Serve (on the trn2.48xlarge , with $OUT/cache and $OUT/nki copied over)
hf download zai-org/GLM-5.3-Flash --revision eb9eb208eb0d988989d07a6a12d0fdeb5f52574a \
--local-dir /models/GLM-5.3-Flash
docker run --rm --privileged --network host --ipc=host \
-v /data/trnrt/cache:/cache \
-v /data/trnrt/nki:/var/tmp/nki-intermediate-cache \
-v /models/GLM-5.3-Flash:/model \
-e NEURON_LIBTORCH_CACHE_ROOT=/cache \
trnrt serve --model /model --port 8000
4. Call it with any OpenAI client
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8000/v1" , api_key = "unused" )
reply = client . chat . completions . create (
model = "glm" ,
messages = [{ "role" : "user" , "content" : "Write a haiku about Trainium." }],
extra_body = { "session_id" : "my-session" }, # keeps a conversation on one engine's prefix cache
)
print ( reply . choices [ 0 ]. message . content )
The full guide covers what each step needs, start-up checks, the batch 4, 8 and 12 configurations and every flag:
deploy · API · configuration .
GLM-5.3-Flash is where TrnRT starts. The engine, router, caches and compile pipeline are built to carry more model
families, and docs/ARCHITECTURE.md describes what a new one needs.
SemiAnalysis AgentX replays real coding-agent sessions. One
trn2.48xlarge ran it at a 1M-token context with 32 concurrent agents, for one hour per point:
Tokens count input (cached input included) and output. Spot price: $8.60/h, us-east-2, 3 Nov 2026. On a capacity block
at $35.76/h, the batch 12 cost is $0.167 per million tokens. Full tables, cold-prompt latencies, long prompts and
methodology are in docs/RESULTS.md .
These numbers were measured with an earlier build of TrnRT. This repository passes these checks against that build
( docs/KNOWN_ISSUES.md ):
Graph equivalence: EQUAL on all 32 ranks, for all 157 graphs that build compiles in the default batch-12
configuration.
Outputs: byte-identical to that build's on 384 of 384 requests.
Serving: the two-engine serve in the Quickstart passes end to end.
Decode is not limited by memory bandwidth here; small instructions, kernel boundaries and exposed collectives dominate it. So
each decode megakernel follows three rules: keep the working set in SBUF, use both physical cores symmetrically, and move
the collectives into the kernel. docs/MEGAKERNELS.md walks through every kernel.
trnrt/
├── __main__.py CLI: serve, engine, compile
├── serve/ OpenAI API, session-affine router, output parsers
├── engine/ scheduler, prefix caches, tensor parallelism
├── models/ GLM-5.3-Flash: KDA, DSA, mHC, MoE, FP8 loading
├── kernels/ NKI decode megakernels and prefill kernels
└── compile/ offline compile: shapes, capture, neuronx-cc, strip
docker/ the serving and compile image
tools/ graph-equivalence gate, device checks, doc checks
tests/ CPU test suite (no NeuronCore needed)
docs/ guides, results, kernel notes
More in docs/ARCHITECTURE.md .
make dev # editable install with test and lint tools (Python 3.13)
make check # lint, type-check, then the CPU test suite
See docs/DEVELOPMENT.md for the test tiers and CI, CONTRIBUTING.md to send a
change (pull requests target development ), SUPPORT.md to get help, and SECURITY.md to
report a vulnerability.
Resource
Link
Trainium2 and trn2 instances
aws.amazon.com/ec2/instance-types/trn2
AWS Neuron SDK documentation
awsdocs-neuron.readthedocs-hosted.com
NKI (Neuron Kernel Interface)
NKI documentation
Neuron deep learning containers
github.com/aws-neuron/deep-learning-containers
EC2 Capacity Blocks for ML
aws.amazon.com/ec2/capacityblocks
Acknowledgments
Thanks to the AWS Neuron team for NKI and its documentation, to Z.ai for releasing GLM-5.3-Flash, and to SemiAnalysis
for AgentX and its public traces. TrnRT includes code derived from other Apache-2.0 projects: see NOTICE
and THIRD_PARTY_NOTICES.md .
AWS, Amazon Web Services, Trainium, Inferentia and Neuron are trademarks of Amazon.com, Inc. or its affiliates. The
AWS logo appears here as published in AWS's Architecture Icons. TrnRT is an independent open-source project, not
affiliated with or endorsed by Amazon Web Services.
The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere.
runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-megakernels Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
19 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. - RunanywhereAI/trnrt

Hi HN, Sanchit here from RunAnywhere (YC W26). We just open-sourced TrnRT: Trainium RunTime, our inference engine for AWS Trainium. We started 3 weeks back, and currently we’re able to serve and run GLM 5.3 flash (FP8) on AWS trn2.48xlarge, and will bring more models support soon. It’s open to contributions :) What it does today: - Serves 2 instances of GLM-5.3 flash on a single box, 8 trn2 with a cache aware router. - Can process 3.57M TPM (tokens per minute) on SemiAnalysis's AgentX benchmarks. - Runs decode as one NKI kernel per layer half, or for the 11 sparse-attention layers fuses both blocks into a single kernel, for faster performance. We started working on Trn2 because we wanted to explore it and see if we can come close to beating it’s performance in terms of cost efficiency compared to H100/H200. A big MoE with 1M context window such as GLM-5.3 flash requires a lot of memory, and memory/tb on a trn2.48xlarge costs less than half compared to 8x h200 on AWS, and hence we wanted to experiment with it and now we’re sharing some early results. But AWS's own stack didn't support the model: vLLM Neuron's latest release did not register any GLM model, and NxDI, AWS's own distributed inference framework, is no longer under active development. We also saw the drop of SGLang Neuron implementation yesterday, but there is no good support for Trn2, which is the only chip available for public use right now. We ran semiAnalysis's AgentX benchmarks, which replays real coding agent sessions: - 1 box served 28.9 RPM(requests per minute) with 32 agents at batch 12 (also known as concurrency) -The P50 TTFT with a cached prompt at 16 agents is ~1.24 seconds with a decode of 35.9 tps. Initially, the recurrence in one kda (linear attention) layer was ~2% of the layer's attention time at batch 1, and the rest was projections, transposes and XLA glue between small kernels. So we wrote the megakernel so that it covers the whole layer in 2 kernel dispatch, which made decode 25% faster at batch 8. The sparse attention kernel keep a layer’s working weights in SBUF, distributing weighted streaming across the cores and handling the interchip communication within There’s an approximation in the sparse-attention prefill, affecting about 1.2% of block selections. You can refer to the write-up for the profiling results: https://runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-mega... Please give it a try yourself on a trn2.48xlarge and reproduce the AgentX results by following the steps in docs/RESULTS.md. Happy to collaborate and hear any feedback. Deploying high performant dedicated inference at lab like performance and half the price for open models such as the GLM-5.3 family, Kimi K3, Deepseek flash etc. is how we make money. You can talk to us regarding the same at founders@runanywhere.ai. We hope you find this useful. If you are working with Trainium or have experience with NKI kernels, we'd love to get your honest feedback. Feel free to share the models you want support for next, and we welcome any contributions to the repository.

GitHub - RunanywhereAI/trnrt: The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere. · GitHub
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
RunanywhereAI
trnrt Public Notifications You must be signed in to change notification settings
Star 1.6k ( 1.6k ) You must be signed in to star a repository
The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere.
runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-megakernels Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
19 forks Report repository development Branches Tags Go to file Code Open more actions menu Latest commit
315 Commits 315 Commits Folders and files
.github .github docker docker docs docs tests tests tools tools trnrt trnrt .coderabbit.yaml .coderabbit.yaml .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md cubic.yaml cubic.yaml greptile.json greptile.json pyproject.toml pyproject.toml View all files Repository files navigation
Quickstart ·
Benchmarks ·
Docs ·
Kernels ·
Blog post ·
runanywhere.ai
TrnRT is an open inference engine for large language models on AWS Trainium. It serves an OpenAI-compatible API
from one trn2.48xlarge , runs decode as NKI megakernels (about one kernel per layer half), and keeps a prefix cache
across device and host memory. GLM-5.3-Flash is the first supported model; more open models are on the way.
How we built it, with the AgentX results: Serving GLM-5.3-Flash on one Trainium2 box with NKI layer megakernels .
Feature
What it does
NKI megakernels
One fused kernel per layer half for decode: KDA attention, the fused DSA layer, mHC and the MoE FFN.
1M-token context
Sparse attention (top 2,048) and linear attention, with fixed-size state.
Prefix cache
A radix tree over FP8 KV pages, plus a host-memory tier for evicted pages and state checkpoints.
Two engines per box
Each engine is tensor-parallel over 32 logical NeuronCores, behind a session-affine router.
OpenAI API
Chat and completions, streaming, tool calls and reasoning output.
Offline compile
Every graph is built on a CPU host; the serving box loads them and compiles nothing by default.
Quickstart
You need Docker, a trn2.48xlarge to serve on, and any x86_64 host to compile on. New to Trainium? Start with
docs/AWS.md .
1. Build the image (Neuron SDK 2.32 (neuronx-cc 2.27), Python 3.13)
git clone https://github.com/RunanywhereAI/trnrt.git && cd trnrt
docker build -f docker/Dockerfile \
--build-arg TRNRT_VERSION= $( python3 -c ' import trnrt; print(trnrt.__version__) ' ) \
--build-arg TRNRT_REVISION= $( git rev-parse HEAD ) -t trnrt .
2. Build the compile cache (any x86_64 host, no NeuronCore needed; takes hours)
OUT=/data/trnrt REV=eb9eb208eb0d988989d07a6a12d0fdeb5f52574a
mkdir -p $OUT /nki $OUT /work
trnrt () { docker run --rm -e HF_TOKEN -v $OUT : $OUT \
-v $OUT /nki:/var/tmp/nki-intermediate-cache trnrt " $@ " ; }
trnrt compile shapes zai-org/GLM-5.3-Flash --revision $REV --out $OUT
trnrt compile decisions --model zai-org/GLM-5.3-Flash --revision $REV --out $OUT
trnrt compile capture --out $OUT -- --model $OUT /shapes
trnrt compile compile --out $OUT --work-root $OUT /work
trnrt compile strip --out $OUT
3. Serve (on the trn2.48xlarge , with $OUT/cache and $OUT/nki copied over)
hf download zai-org/GLM-5.3-Flash --revision eb9eb208eb0d988989d07a6a12d0fdeb5f52574a \
--local-dir /models/GLM-5.3-Flash
docker run --rm --privileged --network host --ipc=host \
-v /data/trnrt/cache:/cache \
-v /data/trnrt/nki:/var/tmp/nki-intermediate-cache \
-v /models/GLM-5.3-Flash:/model \
-e NEURON_LIBTORCH_CACHE_ROOT=/cache \
trnrt serve --model /model --port 8000
4. Call it with any OpenAI client
from openai import OpenAI
client = OpenAI ( base_url = "http://localhost:8000/v1" , api_key = "unused" )
reply = client . chat . completions . create (
model = "glm" ,
messages = [{ "role" : "user" , "content" : "Write a haiku about Trainium." }],
extra_body = { "session_id" : "my-session" }, # keeps a conversation on one engine's prefix cache
)
print ( reply . choices [ 0 ]. message . content )
The full guide covers what each step needs, start-up checks, the batch 4, 8 and 12 configurations and every flag:
deploy · API · configuration .
GLM-5.3-Flash is where TrnRT starts. The engine, router, caches and compile pipeline are built to carry more model
families, and docs/ARCHITECTURE.md describes what a new one needs.
SemiAnalysis AgentX replays real coding-agent sessions. One
trn2.48xlarge ran it at a 1M-token context with 32 concurrent agents, for one hour per point:
Tokens count input (cached input included) and output. Spot price: $8.60/h, us-east-2, 3 Nov 2026. On a capacity block
at $35.76/h, the batch 12 cost is $0.167 per million tokens. Full tables, cold-prompt latencies, long prompts and
methodology are in docs/RESULTS.md .
These numbers were measured with an earlier build of TrnRT. This repository passes these checks against that build
( docs/KNOWN_ISSUES.md ):
Graph equivalence: EQUAL on all 32 ranks, for all 157 graphs that build compiles in the default batch-12
configuration.
Outputs: byte-identical to that build's on 384 of 384 requests.
Serving: the two-engine serve in the Quickstart passes end to end.
Decode is not limited by memory bandwidth here; small instructions, kernel boundaries and exposed collectives dominate it. So
each decode megakernel follows three rules: keep the working set in SBUF, use both physical cores symmetrically, and move
the collectives into the kernel. docs/MEGAKERNELS.md walks through every kernel.
trnrt/
├── __main__.py CLI: serve, engine, compile
├── serve/ OpenAI API, session-affine router, output parsers
├── engine/ scheduler, prefix caches, tensor parallelism
├── models/ GLM-5.3-Flash: KDA, DSA, mHC, MoE, FP8 loading
├── kernels/ NKI decode megakernels and prefill kernels
└── compile/ offline compile: shapes, capture, neuronx-cc, strip
docker/ the serving and compile image
tools/ graph-equivalence gate, device checks, doc checks
tests/ CPU test suite (no NeuronCore needed)
docs/ guides, results, kernel notes
More in docs/ARCHITECTURE.md .
make dev # editable install with test and lint tools (Python 3.13)
make check # lint, type-check, then the CPU test suite
See docs/DEVELOPMENT.md for the test tiers and CI, CONTRIBUTING.md to send a
change (pull requests target development ), SUPPORT.md to get help, and SECURITY.md to
report a vulnerability.
Resource
Link
Trainium2 and trn2 instances
aws.amazon.com/ec2/instance-types/trn2
AWS Neuron SDK documentation
awsdocs-neuron.readthedocs-hosted.com
NKI (Neuron Kernel Interface)
NKI documentation
Neuron deep learning containers
github.com/aws-neuron/deep-learning-containers
EC2 Capacity Blocks for ML
aws.amazon.com/ec2/capacityblocks
Acknowledgments
Thanks to the AWS Neuron team for NKI and its documentation, to Z.ai for releasing GLM-5.3-Flash, and to SemiAnalysis
for AgentX and its public traces. TrnRT includes code derived from other Apache-2.0 projects: see NOTICE
and THIRD_PARTY_NOTICES.md .
AWS, Amazon Web Services, Trainium, Inferentia and Neuron are trademarks of Amazon.com, Inc. or its affiliates. The
AWS logo appears here as published in AWS's Architecture Icons. TrnRT is an independent open-source project, not
affiliated with or endorsed by Amazon Web Services.
The open LLM inference engine for AWS Trainium: NKI megakernels, 1M-token context, OpenAI-compatible API. By RunAnywhere.
runanywhere.ai/blog/glm-5-3-flash-trainium2-nki-megakernels Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
19 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
