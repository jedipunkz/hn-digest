---
source: "https://github.com/IshekKhal/bitrun"
hn_url: "https://news.ycombinator.com/item?id=50106895"
title: "Show HN: BitRun v0.2 – Run 8B ternary LLMs on CPU with live wattage telemetry"
article_title: "GitHub - IshekKhal/bitrun: Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) · GitHub"
image: "https://opengraph.githubassets.com/03ff1231f944c4fd192d9490eff33994f5f9eff9ecdf28fdcc3f738f2d86563d/IshekKhal/bitrun"
author: "ishekofficial"
captured_at: "2026-10-15T18:02:51Z"
capture_tool: "hn-digest"
hn_id: 50106895
score: 2
comments: 0
posted_at: "2026-10-15T17:15:08Z"
tags:
  - hacker-news
---

# Show HN: BitRun v0.2 – Run 8B ternary LLMs on CPU with live wattage telemetry

- HN: [50106895](https://news.ycombinator.com/item?id=50106895)
- Source: [github.com](https://github.com/IshekKhal/bitrun)
- Score: 2
- Comments: 0
- Posted: 2026-10-15T17:15:08Z

## Translation

Title: Show HN: BitRun v0.2 – Run 8B ternary LLMs on CPU with live wattage telemetry
Article title: GitHub - IshekKhal/bitrun: Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) · GitHub
Description: Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) - IshekKhal/bitrun

Article text:
GitHub - IshekKhal/bitrun: Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) · GitHub
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
IshekKhal
bitrun Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux)
Readme MIT license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
examples examples patches patches pkg pkg scripts scripts .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh main.go main.go View all files Repository files navigation
Zero-compile runner, model manager, and live terminal dashboard for 1-bit and 1.58-bit ternary LLMs on Apple Silicon and x86_64 CPUs.
Microsoft's BitNet b1.58 demonstrated that ternary (1-bit / 1.58-bit) language models can run at high speed on commodity CPUs with 70–80% lower energy consumption by replacing floating-point multiplications with integer additions.
Ternary weights ( {-1, 0, +1} ) are engineered for local CPU inference, battery-powered laptops, and edge devices where running standard FP16 or INT4 models drains battery or demands an expensive GPU.
However, testing research implementations directly often requires:
Cloning recursive C++ submodules and managing compiler toolchains, CMake, or virtual environments.
Manually configuring CPU vector instruction flags (AVX2, AVX-512, ARM NEON).
Dealing with multi-gigabyte Docker images that defeat the point of low-power computing.
Debugging memory stride crashes and unhandled tensor formats on commodity hardware.
BitRun packages the native inference engine into a standalone runner with zero external dependencies and a clean developer workflow:
# 1. Pull a 1-bit model directly from Hugging Face
bitrun pull falcon3-1b
# 2. Launch the BitRun Cockpit full-screen dashboard
bitrun cockpit falcon3-1b
# 3. Or spin up a local OpenAI-compatible REST server
bitrun serve falcon3-1b
The BitRun Cockpit
BitRun v0.2.0-beta includes Cockpit , an interactive full-screen terminal interface built with Go and Bubble Tea:
◆ BitRun Cockpit • model: bonsai-1.7b • quant: PrismML Q1_0 • port: 11435 ⚡ Pure CPU SIMD (-ngl 0)
╭────────────────────────────────────────────────────────╮╭────────────────────────────────────────────╮
│❯ User 12:11:37 ││ LIVE HARDWARE TELEMETRY │
│Explain how 1-bit models run on commodity CPUs. ││ Gen Speed : 73.7 tok/s (peak: 77.9) │
│ ││ Prompt Speed : 433.5 tok/s │
│◆ BitRun 12:11:37 ││ Prompt TTFT : 138 ms │
│Ternary models constrain weight values to {-1, 0, +1}. ││ Context : 124/4096 (3%) │
│This transforms matrix multiplication from floating- ││ [■□□□□□□□□□□□□□□□] │
│point math into simple integer additions and ││ P-Cores ( 8) : [■■■■■■■■■□□□] 75% │
│subtractions. Because weights take only 1.58 bits, ││ E-Cores ( 4) : [■□□□□□□□□□□□] 11% │
│memory bandwidth drops by roughly 10x, allowing laptops ││ RAM Available : 24.0 GB (Host) │
│to run inference at high speed while drawing minimal ││ SIMD Engine : ARM NEON DotProd (Native) │
│power... ││ │
│ │╰────────────────────────────────────────────╯
│ │╭────────────────────────────────────────────╮
│ ││ POWER & ENERGY TELEMETRY │
│ ││ Package Power : 11.7 W [■■■□□□□□□□□□] │
│ ││ Energy/Token : 0.1587 Joules │
│ ││ Efficiency : 6.3 tokens/Joule │
│ ││ Session Energy: 216.125 mWh (778.0 J) │
│ ││ Power Source : AC Attached │
│ ││ Battery Draw : 0.0 W (AC Powered) │
│ ││ REST Endpoint : http://127.0.0.1:11435 │
│ ││ Status : [READY] │
│ │╰────────────────────────────────────────────╯
╰────────────────────────────────────────────────────────╯
[Enter] Send Prompt • [↑/↓/PgUp/PgDn] Scroll Chat • [Ctrl+L] Clear Screen • [Esc/Ctrl+C] Quit Cockpit
Cockpit Capabilities
Dual-Pane Layout : The left viewport maintains multi-turn conversation history with independent keyboard and mouse wheel scrolling.
Real-Time Streaming : Tokens stream character-by-character via server-sent events ( /v1/chat/completions ) with an active cursor indicator.
Live Hardware Telemetry : Samples prompt time-to-first-token (TTFT), instantaneous and peak generation speed, and P-core vs E-core utilization.
Energy Gauges : Tracks package power in Watts, energy efficiency in Joules per token, cumulative milliwatt-hours consumed, and estimated battery runtime.
Clean Terminal Buffer : Background inference engine logs redirect to ~/.bitrun/engine.log , keeping the terminal display free of coordinate corruption.
macOS (Apple Silicon) & Linux (x86_64) :
curl -fsSL https://raw.githubusercontent.com/IshekKhal/bitrun/main/install.sh | sh
Windows (PowerShell) :
irm https: // raw.githubusercontent.com / IshekKhal / bitrun / main / install.ps1 | iex
Pre-compiled standalone archives are also available on the Releases page for offline setup.
# Verify system topology and SIMD extensions
bitrun hardware
# Download a model from the registry
bitrun pull bonsai-1.7b
# Launch interactive Cockpit interface
bitrun cockpit bonsai-1.7b
# Or run a standard interactive terminal session
bitrun run bonsai-1.7b
# Run a one-shot query directly
bitrun run bonsai-1.7b " Summarize the advantages of ternary weight quantization. "
Supported Model Catalog
BitRun includes pre-configured registries for 6 ternary and compact models:
The figures below reflect physical hardware runs using native binaries with pure CPU SIMD execution ( --tokens 64 ). Zero GPU offloading ( -ngl 0 ).
1. Apple Silicon (M4 Pro CPU, 12 Cores: 8P + 4E, 24 GB Unified RAM)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
11,549.2 t/s
1,619.8 t/s
78 ms
11.70 W
0.043 J/tok
bonsai-1.7b
237 MB
411.7 t/s
72.8 t/s
576 ms
11.70 W
0.316 J/tok
falcon3-1b
1.30 GB
170.8 t/s
73.2 t/s
628 ms
11.70 W
0.345 J/tok
bitnet-2b
1.13 GB
94.4 t/s
31.6 t/s
881 ms
11.70 W
0.484 J/tok
falcon3-3b
2.07 GB
80.6 t/s
33.1 t/s
1.08 s
11.70 W
0.592 J/tok
bonsai-8b
2.08 GB
30.1 t/s
21.0 t/s
1.62 s
11.70 W
0.886 J/tok
2. Windows 10 Native (Intel Core i5-12600K, 16 Threads, 32 GB RAM, MSVC AVX2)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
1,766.2 t/s
597.8 t/s
595 ms
22.60 W
0.038 J/tok
bonsai-1.7b
237 MB
152.4 t/s
85.6 t/s
1.42 s
22.60 W
0.264 J/tok
falcon3-1b
1.30 GB
14.9 t/s
11.1 t/s
5.48 s
22.60 W
2.036 J/tok
bitnet-2b
1.13 GB
8.0 t/s
6.7 t/s
9.67 s
22.60 W
3.373 J/tok
falcon3-3b
2.07 GB
7.5 t/s
6.0 t/s
10.34 s
22.60 W
3.766 J/tok
bonsai-8b
2.08 GB
4.7 t/s
4.0 t/s
15.38 s
22.60 W
5.650 J/tok
3. Linux WSL2 (Intel Core i5-12600K, 4 vCPUs Allocated, GCC 13.3 AVX2)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
333.1 t/s
134.1 t/s
159 ms
8.50 W
0.063 J/tok
bonsai-1.7b
237 MB
41.8 t/s
16.8 t/s
1.27 s
8.50 W
0.505 J/tok
falcon3-1b
1.30 GB
16.7 t/s
6.7 t/s
3.17 s
8.50 W
1.261 J/tok
bitnet-2b
1.13 GB
10.9 t/s
4.4 t/s
4.86 s
8.50 W
1.938 J/tok
falcon3-3b
2.07 GB
8.7 t/s
3.5 t/s
6.12 s
8.50 W
2.439 J/tok
bonsai-8b
2.08 GB
1.3 t/s
0.5 t/s
40.30 s
8.50 W
16.058 J/tok
Energy & Telemetry Monitoring
Ternary models run at low electrical power (8W to 22W total package draw) rather than 250W to 400W on dedicated graphics cards. BitRun incorporates hardware telemetry hooks across:
macOS : Queries Apple Silicon SoC power counters and battery voltage/amperage via ioreg .
Linux : Reads Intel and AMD processor energy directly from RAPL microjoule counters ( /sys/class/powercap/intel-rapl ).
Windows : Measures CPU package power via Windows Processor Power Management (PPM).
# Run a prompt with real-time energy telemetry
bitrun run bonsai-1.7b " Explain quantum key distribution " --power
# Run a synthetic benchmark and report Joules/token
bitrun bench bonsai-1.7b --tokens 64 --threads 8
Example terminal telemetry box:
┌────────────────────────────────────────────────────────┐
│ BitRun Performance & Power │
├───────────────────────┬────────────────────────────────┤
│ Prompt Evaluation │ 411.7 tokens/sec (12 tokens) │
│ Token Generation │ 72.8 tokens/sec (64 tokens) │
│ Time to First Token │ 576 ms │
│ Average Power │ 11.70 W │
│ Total Energy Consumed │ 10.28 Joules │
│ Energy Efficiency │ 0.316 Joules/token │
│ Battery Drain Delta │ 0.00% │
└───────────────────────┴────────────────────────────────┘
Modelfiles & Custom Personas
You can define custom personas, temperature presets, and context windows using declarative Modelfiles:
# Modelfile
FROM falcon3-1b
SYSTEM You are an expert systems engineer. Answer concisely with production-grade code.
PARAMETER temperature 0.2
PARAMETER context_size 4096
Build and run your custom persona:
# Create custom persona manifest
bitrun create gopher -f Modelfile
# Chat with your persona
bitrun run gopher
# Or launch it directly inside the Cockpit
bitrun cockpit gopher
OpenAI-Compatible REST API
bitrun serve starts a background server on port 11435 (avoiding port collisions with Ollama on 11434 ):
bitrun serve falcon3-1b --port 11435
Direct HTTP Request
curl -s http://127.0.0.1:11435/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "falcon3-1b",
"messages": [
{"role": "user", "content": "How do 1-bit models conserve memory bandwidth?"}
],
"temperature": 0.2
} '
Python SDK Example
from openai import OpenAI
client = OpenAI (
base_url = "http://127.0.0.1:11435/v1" ,
api_key = "bitrun" # Any string
)
response = client . chat . completions . create (
model = "falcon3-1b" ,
messages = [{ "role" : "user" , "content" : "Write a concurrent queue in Go." }],
stream = True
)
for chunk in response :
content = chunk . choices [ 0 ]. delta . content
if content :
print ( content , end = "" , flush = True )
print ()
Connects with Cursor, Continue.dev, Open WebUI, and LangChain by setting the base URL to http://localhost:11435/v1 .
Command
Usage
Description
bitrun cockpit <model>
bitrun cockpit bonsai-1.7b
Opens the full-screen terminal dashboard with live hardware and energy gauges.
bitrun run <model> [prompt]
bitrun run falcon3-1b
Starts an interactive chat session or evaluates a single prompt.
bitrun bench <model>
bitrun bench bitnet-2b --tokens 64
Measures prompt speed, generation throughput, and Joules/token.
bitrun pull <model>
bitrun pull bonsai-8b
Downloads a model checkpoint from Hugging Face with resume support.
bitrun serve <model>
bitrun serve falcon3-1b --port 11435
Runs the local OpenAI-compatible REST server daemon.
bitr

[truncated]

## Original Extract

Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) - IshekKhal/bitrun

GitHub - IshekKhal/bitrun: Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux) · GitHub
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
IshekKhal
bitrun Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Zero-compile runner and terminal Cockpit for 1-bit / 1.58-bit ternary LLMs on CPU. Real-time wattage telemetry, 70+ t/s on laptops, and 8B model support (macOS, Windows, Linux)
Readme MIT license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
examples examples patches patches pkg pkg scripts scripts .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh main.go main.go View all files Repository files navigation
Zero-compile runner, model manager, and live terminal dashboard for 1-bit and 1.58-bit ternary LLMs on Apple Silicon and x86_64 CPUs.
Microsoft's BitNet b1.58 demonstrated that ternary (1-bit / 1.58-bit) language models can run at high speed on commodity CPUs with 70–80% lower energy consumption by replacing floating-point multiplications with integer additions.
Ternary weights ( {-1, 0, +1} ) are engineered for local CPU inference, battery-powered laptops, and edge devices where running standard FP16 or INT4 models drains battery or demands an expensive GPU.
However, testing research implementations directly often requires:
Cloning recursive C++ submodules and managing compiler toolchains, CMake, or virtual environments.
Manually configuring CPU vector instruction flags (AVX2, AVX-512, ARM NEON).
Dealing with multi-gigabyte Docker images that defeat the point of low-power computing.
Debugging memory stride crashes and unhandled tensor formats on commodity hardware.
BitRun packages the native inference engine into a standalone runner with zero external dependencies and a clean developer workflow:
# 1. Pull a 1-bit model directly from Hugging Face
bitrun pull falcon3-1b
# 2. Launch the BitRun Cockpit full-screen dashboard
bitrun cockpit falcon3-1b
# 3. Or spin up a local OpenAI-compatible REST server
bitrun serve falcon3-1b
The BitRun Cockpit
BitRun v0.2.0-beta includes Cockpit , an interactive full-screen terminal interface built with Go and Bubble Tea:
◆ BitRun Cockpit • model: bonsai-1.7b • quant: PrismML Q1_0 • port: 11435 ⚡ Pure CPU SIMD (-ngl 0)
╭────────────────────────────────────────────────────────╮╭────────────────────────────────────────────╮
│❯ User 12:11:37 ││ LIVE HARDWARE TELEMETRY │
│Explain how 1-bit models run on commodity CPUs. ││ Gen Speed : 73.7 tok/s (peak: 77.9) │
│ ││ Prompt Speed : 433.5 tok/s │
│◆ BitRun 12:11:37 ││ Prompt TTFT : 138 ms │
│Ternary models constrain weight values to {-1, 0, +1}. ││ Context : 124/4096 (3%) │
│This transforms matrix multiplication from floating- ││ [■□□□□□□□□□□□□□□□] │
│point math into simple integer additions and ││ P-Cores ( 8) : [■■■■■■■■■□□□] 75% │
│subtractions. Because weights take only 1.58 bits, ││ E-Cores ( 4) : [■□□□□□□□□□□□] 11% │
│memory bandwidth drops by roughly 10x, allowing laptops ││ RAM Available : 24.0 GB (Host) │
│to run inference at high speed while drawing minimal ││ SIMD Engine : ARM NEON DotProd (Native) │
│power... ││ │
│ │╰────────────────────────────────────────────╯
│ │╭────────────────────────────────────────────╮
│ ││ POWER & ENERGY TELEMETRY │
│ ││ Package Power : 11.7 W [■■■□□□□□□□□□] │
│ ││ Energy/Token : 0.1587 Joules │
│ ││ Efficiency : 6.3 tokens/Joule │
│ ││ Session Energy: 216.125 mWh (778.0 J) │
│ ││ Power Source : AC Attached │
│ ││ Battery Draw : 0.0 W (AC Powered) │
│ ││ REST Endpoint : http://127.0.0.1:11435 │
│ ││ Status : [READY] │
│ │╰────────────────────────────────────────────╯
╰────────────────────────────────────────────────────────╯
[Enter] Send Prompt • [↑/↓/PgUp/PgDn] Scroll Chat • [Ctrl+L] Clear Screen • [Esc/Ctrl+C] Quit Cockpit
Cockpit Capabilities
Dual-Pane Layout : The left viewport maintains multi-turn conversation history with independent keyboard and mouse wheel scrolling.
Real-Time Streaming : Tokens stream character-by-character via server-sent events ( /v1/chat/completions ) with an active cursor indicator.
Live Hardware Telemetry : Samples prompt time-to-first-token (TTFT), instantaneous and peak generation speed, and P-core vs E-core utilization.
Energy Gauges : Tracks package power in Watts, energy efficiency in Joules per token, cumulative milliwatt-hours consumed, and estimated battery runtime.
Clean Terminal Buffer : Background inference engine logs redirect to ~/.bitrun/engine.log , keeping the terminal display free of coordinate corruption.
macOS (Apple Silicon) & Linux (x86_64) :
curl -fsSL https://raw.githubusercontent.com/IshekKhal/bitrun/main/install.sh | sh
Windows (PowerShell) :
irm https: // raw.githubusercontent.com / IshekKhal / bitrun / main / install.ps1 | iex
Pre-compiled standalone archives are also available on the Releases page for offline setup.
# Verify system topology and SIMD extensions
bitrun hardware
# Download a model from the registry
bitrun pull bonsai-1.7b
# Launch interactive Cockpit interface
bitrun cockpit bonsai-1.7b
# Or run a standard interactive terminal session
bitrun run bonsai-1.7b
# Run a one-shot query directly
bitrun run bonsai-1.7b " Summarize the advantages of ternary weight quantization. "
Supported Model Catalog
BitRun includes pre-configured registries for 6 ternary and compact models:
The figures below reflect physical hardware runs using native binaries with pure CPU SIMD execution ( --tokens 64 ). Zero GPU offloading ( -ngl 0 ).
1. Apple Silicon (M4 Pro CPU, 12 Cores: 8P + 4E, 24 GB Unified RAM)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
11,549.2 t/s
1,619.8 t/s
78 ms
11.70 W
0.043 J/tok
bonsai-1.7b
237 MB
411.7 t/s
72.8 t/s
576 ms
11.70 W
0.316 J/tok
falcon3-1b
1.30 GB
170.8 t/s
73.2 t/s
628 ms
11.70 W
0.345 J/tok
bitnet-2b
1.13 GB
94.4 t/s
31.6 t/s
881 ms
11.70 W
0.484 J/tok
falcon3-3b
2.07 GB
80.6 t/s
33.1 t/s
1.08 s
11.70 W
0.592 J/tok
bonsai-8b
2.08 GB
30.1 t/s
21.0 t/s
1.62 s
11.70 W
0.886 J/tok
2. Windows 10 Native (Intel Core i5-12600K, 16 Threads, 32 GB RAM, MSVC AVX2)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
1,766.2 t/s
597.8 t/s
595 ms
22.60 W
0.038 J/tok
bonsai-1.7b
237 MB
152.4 t/s
85.6 t/s
1.42 s
22.60 W
0.264 J/tok
falcon3-1b
1.30 GB
14.9 t/s
11.1 t/s
5.48 s
22.60 W
2.036 J/tok
bitnet-2b
1.13 GB
8.0 t/s
6.7 t/s
9.67 s
22.60 W
3.373 J/tok
falcon3-3b
2.07 GB
7.5 t/s
6.0 t/s
10.34 s
22.60 W
3.766 J/tok
bonsai-8b
2.08 GB
4.7 t/s
4.0 t/s
15.38 s
22.60 W
5.650 J/tok
3. Linux WSL2 (Intel Core i5-12600K, 4 vCPUs Allocated, GCC 13.3 AVX2)
Model Tag
Size
Prompt Speed
Generation Speed
TTFT
Package Power
Energy / Token
bitnet-70m
79 MB
333.1 t/s
134.1 t/s
159 ms
8.50 W
0.063 J/tok
bonsai-1.7b
237 MB
41.8 t/s
16.8 t/s
1.27 s
8.50 W
0.505 J/tok
falcon3-1b
1.30 GB
16.7 t/s
6.7 t/s
3.17 s
8.50 W
1.261 J/tok
bitnet-2b
1.13 GB
10.9 t/s
4.4 t/s
4.86 s
8.50 W
1.938 J/tok
falcon3-3b
2.07 GB
8.7 t/s
3.5 t/s
6.12 s
8.50 W
2.439 J/tok
bonsai-8b
2.08 GB
1.3 t/s
0.5 t/s
40.30 s
8.50 W
16.058 J/tok
Energy & Telemetry Monitoring
Ternary models run at low electrical power (8W to 22W total package draw) rather than 250W to 400W on dedicated graphics cards. BitRun incorporates hardware telemetry hooks across:
macOS : Queries Apple Silicon SoC power counters and battery voltage/amperage via ioreg .
Linux : Reads Intel and AMD processor energy directly from RAPL microjoule counters ( /sys/class/powercap/intel-rapl ).
Windows : Measures CPU package power via Windows Processor Power Management (PPM).
# Run a prompt with real-time energy telemetry
bitrun run bonsai-1.7b " Explain quantum key distribution " --power
# Run a synthetic benchmark and report Joules/token
bitrun bench bonsai-1.7b --tokens 64 --threads 8
Example terminal telemetry box:
┌────────────────────────────────────────────────────────┐
│ BitRun Performance & Power │
├───────────────────────┬────────────────────────────────┤
│ Prompt Evaluation │ 411.7 tokens/sec (12 tokens) │
│ Token Generation │ 72.8 tokens/sec (64 tokens) │
│ Time to First Token │ 576 ms │
│ Average Power │ 11.70 W │
│ Total Energy Consumed │ 10.28 Joules │
│ Energy Efficiency │ 0.316 Joules/token │
│ Battery Drain Delta │ 0.00% │
└───────────────────────┴────────────────────────────────┘
Modelfiles & Custom Personas
You can define custom personas, temperature presets, and context windows using declarative Modelfiles:
# Modelfile
FROM falcon3-1b
SYSTEM You are an expert systems engineer. Answer concisely with production-grade code.
PARAMETER temperature 0.2
PARAMETER context_size 4096
Build and run your custom persona:
# Create custom persona manifest
bitrun create gopher -f Modelfile
# Chat with your persona
bitrun run gopher
# Or launch it directly inside the Cockpit
bitrun cockpit gopher
OpenAI-Compatible REST API
bitrun serve starts a background server on port 11435 (avoiding port collisions with Ollama on 11434 ):
bitrun serve falcon3-1b --port 11435
Direct HTTP Request
curl -s http://127.0.0.1:11435/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "falcon3-1b",
"messages": [
{"role": "user", "content": "How do 1-bit models conserve memory bandwidth?"}
],
"temperature": 0.2
} '
Python SDK Example
from openai import OpenAI
client = OpenAI (
base_url = "http://127.0.0.1:11435/v1" ,
api_key = "bitrun" # Any string
)
response = client . chat . completions . create (
model = "falcon3-1b" ,
messages = [{ "role" : "user" , "content" : "Write a concurrent queue in Go." }],
stream = True
)
for chunk in response :
content = chunk . choices [ 0 ]. delta . content
if content :
print ( content , end = "" , flush = True )
print ()
Connects with Cursor, Continue.dev, Open WebUI, and LangChain by setting the base URL to http://localhost:11435/v1 .
Command
Usage
Description
bitrun cockpit <model>
bitrun cockpit bonsai-1.7b
Opens the full-screen terminal dashboard with live hardware and energy gauges.
bitrun run <model> [prompt]
bitrun run falcon3-1b
Starts an interactive chat session or evaluates a single prompt.
bitrun bench <model>
bitrun bench bitnet-2b --tokens 64
Measures prompt speed, generation throughput, and Joules/token.
bitrun pull <model>
bitrun pull bonsai-8b
Downloads a model checkpoint from Hugging Face with resume support.
bitrun serve <model>
bitrun serve falcon3-1b --port 11435
Runs the local OpenAI-compatible REST server daemon.
bitr

[truncated]
