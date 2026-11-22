---
source: "https://github.com/V0val0x/DroidPath/blob/main/docs/benchmarks.md"
hn_url: "https://news.ycombinator.com/item?id=50600460"
title: "Bypassing ADB: 0.32ms Android UI trees and A* routing for AI agents (Rust/Kotlin"
article_title: "DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath · GitHub"
image: "https://opengraph.githubassets.com/8db6038f14c2d63eb98d85efffbaae12cb67fddac7485cb4fcad4d7497fa368b/V0val0x/DroidPath"
author: "Ak4shik"
captured_at: "2026-11-22T15:58:20Z"
capture_tool: "hn-digest"
hn_id: 50600460
score: 1
comments: 0
posted_at: "2026-11-22T15:19:19Z"
tags:
  - hacker-news
---

# Bypassing ADB: 0.32ms Android UI trees and A* routing for AI agents (Rust/Kotlin

- HN: [50600460](https://news.ycombinator.com/item?id=50600460)
- Source: [github.com](https://github.com/V0val0x/DroidPath/blob/main/docs/benchmarks.md)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T15:19:19Z

## Translation

Title: Bypassing ADB: 0.32ms Android UI trees and A* routing for AI agents (Rust/Kotlin
Article title: DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath · GitHub
Description: Cross-platform Rust + Tauri v2 + Kotlin Shadow DOM MCP server & persistent A* symbolic UI router for autonomous Android control (Windows, Linux, macOS) - DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath

Article text:
DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath · GitHub
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
V0val0x
Copy path Blame More file actions Blame More file actions Latest commit
History History 81 lines (59 loc) · 6.23 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions DroidPath Agentic Benchmarks
This benchmark suite compares DroidPath (Shadow DOM + Layer 3 A Spatial Graph Router) * against the two dominant Android automation paradigms used by autonomous AI agents:
Pure Vision (Screenshot + VLM loop) — captures a full-resolution PNG screenshot on every step and sends it to a Vision-Language Model to predict (x, y) coordinates.
Appium / uiautomator dump XML loop — runs uiautomator dump over ADB on every step and feeds the raw XML hierarchy (~25–35 KB) to an LLM.
DroidPath (Shadow DOM + Layer 3 A ) * — streams incremental NDJSON accessibility diffs over a persistent loopback TCP socket ( android-daemon ) and executes learned multi-hop paths locally in Rust via A* without intermediate LLM turns.
Measured Results (Live Unlocked Pixel 10a Hardware + Replay Suite)
Transparency Note: Hardware primitive latencies and End-to-End (E2E) Settings navigation timings below were measured live over USB ADB on an unlocked Google Pixel 10a ( 61061JEA307110 , Android com.android.settings root screen, 54 live UI nodes) using node scripts/benchmark_agentic.mjs --mode live --json against compiled target/release/droidpath-mcp . Multi-step Pure Vision and Appium/XML comparisons in the replay table run real droidpath-mcp stdio JSON-RPC calls against a 60-node loopback Shadow DOM daemon paired with explicit token and cost formulas documented under Methodology & Token Model .
1. Live Hardware Probe (Unlocked Google Pixel 10a · com.android.settings , 54 Nodes)
Primitive / Method
Status
Measured Latency
Payload Size
Notes
DroidPath android_get_ui_tree (steady-state)
ok
0.32 ms
10,742 bytes (~2,686 tok)
54 live UI nodes from local Shadow DOM stream ( 7,575× faster than uiautomator )
DroidPath android_get_ui_tree (cold connect)
ok
41.89 ms
10,742 bytes (~2,686 tok)
First TCP connect + live snapshot ( 58× faster than uiautomator )
DroidPath Hybrid Vision ( include_vision: true )
ok
452.33 ms
38,581 bytes JPEG (540×1212)
54 live UI nodes + compressed JPEG (q45, 6.1× smaller than raw PNG)
adb exec-out screencap -p (raw PNG)
ok
658.1 ms
235,258 bytes PNG
Visual capture only — no structured UI tree
adb exec-out uiautomator dump /dev/tty (XML)
ok
2,423.9 ms
29,357 bytes (~7,340 tok)
Requires stopping UiAutomation daemon; 2.7× more text tokens than Shadow DOM
2. Live End-to-End Hardware Navigation (Google Pixel 10a · Real Seed Hydration → Warm A*)
Task ( Settings → Target )
Cold-Start (Seed Hydration)
Cold Status
Warm A* Graph
Warm Status
Live Success Rate
Settings → Display
378.6 ms
arrived
296.2 ms (1 step)
arrived
100%
Settings → Battery
655.9 ms
arrived
642.8 ms (1 step)
arrived
100%
Settings → About phone (scroll required)
775.7 ms
arrived
744.1 ms (1 step)
arrived
100%
Average across live suite
603 ms
100% (3/3)
561 ms
100% (3/3)
100%
3. Local MCP Runtime Micro-Benchmarks (Measured)
Metric
Measured Value
tools/list JSON-RPC latency (9 tools)
0.25 ms
android_graph_stats JSON-RPC latency
0.42 ms
Pre-hydrated Seed Graph loaded at boot
20 vertices / 19 edges
Compact Shadow DOM tree (60-node mock Settings screen)
4,537 bytes (~1,135 tokens)
4. Architecture & Cost Comparison (5 WebArena-Style Multi-Step Tasks · Replay + API Cost Model)
Architecture
Avg LLM Calls / Task
Avg Tokens / Task
Avg Est. Task Time
Est. Cost / 1,000 Tasks
Advantage of DroidPath
DroidPath (Shadow DOM + A Router) *
1.0
39 (warm A )* / ~1,135 (cold tree)
464 ms (561 ms live warm)
$0.12
100% live E2E success · 192× cheaper than Vision
Pure Vision (Screenshots + VLM)
4.2
7,673
9,954 ms
$23.02
Baseline
Appium / uiautomator XML
4.2
29,400
8,946 ms
$88.20
735× more expensive than Warm DroidPath
Per-Task Breakdown
Task
Hops
DroidPath (Calls / Tokens / Time)
Pure Vision (Calls / Tokens / Time)
Appium XML (Calls / Tokens / Time)
Enable Night Light (Display)
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Open Wi-Fi Settings (Network)
4
1 call · 39 tok · 580 ms
5 calls · 9,135 tok · 11,850 ms
5 calls · 35,000 tok · 10,650 ms
Navigate to Battery Saver
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Open App Notifications
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Scroll to About Phone
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Methodology & Token Model
Text tokens: ceil(utf8_bytes / 4) (standard BPE approximation for JSON/XML).
Cost / 1,000 tasks: Computed transparently at $3.00 / 1M input tokens (standard Claude Sonnet-class API pricing): (avg_tokens_per_task * 1,000 / 1,000,000) * $3.00 .
Vision image tokens: Anthropic's published (W × H) / 750 formula for a 1080×2400 Android screenshot scaled to 1568px long edge ( 706 × 1568 / 750 = 1,477 image tokens + ~350 prompt/tool tokens = 1,827 tokens/step ).
Opt-in Hybrid Vision ( include_vision: true ): Downscaled 540px-wide JPEG ( quality: 45 , measured 38.6 KB on unlocked Pixel 10a Settings vs 235.3 KB raw PNG) is strictly opt-in and only used when an agent explicitly needs visual disambiguation (e.g., canvas/icons), remaining strictly blocked on FLAG_SECURE screens .
# Build release binary first (optional — falls back to target/debug)
cargo build --release -p droidpath-mcp
# Replay benchmark (zero device required)
node scripts/benchmark_agentic.mjs
# Live physical device / emulator probe (measures real screencap -p & uiautomator dump)
node scripts/benchmark_agentic.mjs --mode live
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Cross-platform Rust + Tauri v2 + Kotlin Shadow DOM MCP server & persistent A* symbolic UI router for autonomous Android control (Windows, Linux, macOS) - DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath

DroidPath/docs/benchmarks.md at main · V0val0x/DroidPath · GitHub
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
V0val0x
Copy path Blame More file actions Blame More file actions Latest commit
History History 81 lines (59 loc) · 6.23 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions DroidPath Agentic Benchmarks
This benchmark suite compares DroidPath (Shadow DOM + Layer 3 A Spatial Graph Router) * against the two dominant Android automation paradigms used by autonomous AI agents:
Pure Vision (Screenshot + VLM loop) — captures a full-resolution PNG screenshot on every step and sends it to a Vision-Language Model to predict (x, y) coordinates.
Appium / uiautomator dump XML loop — runs uiautomator dump over ADB on every step and feeds the raw XML hierarchy (~25–35 KB) to an LLM.
DroidPath (Shadow DOM + Layer 3 A ) * — streams incremental NDJSON accessibility diffs over a persistent loopback TCP socket ( android-daemon ) and executes learned multi-hop paths locally in Rust via A* without intermediate LLM turns.
Measured Results (Live Unlocked Pixel 10a Hardware + Replay Suite)
Transparency Note: Hardware primitive latencies and End-to-End (E2E) Settings navigation timings below were measured live over USB ADB on an unlocked Google Pixel 10a ( 61061JEA307110 , Android com.android.settings root screen, 54 live UI nodes) using node scripts/benchmark_agentic.mjs --mode live --json against compiled target/release/droidpath-mcp . Multi-step Pure Vision and Appium/XML comparisons in the replay table run real droidpath-mcp stdio JSON-RPC calls against a 60-node loopback Shadow DOM daemon paired with explicit token and cost formulas documented under Methodology & Token Model .
1. Live Hardware Probe (Unlocked Google Pixel 10a · com.android.settings , 54 Nodes)
Primitive / Method
Status
Measured Latency
Payload Size
Notes
DroidPath android_get_ui_tree (steady-state)
ok
0.32 ms
10,742 bytes (~2,686 tok)
54 live UI nodes from local Shadow DOM stream ( 7,575× faster than uiautomator )
DroidPath android_get_ui_tree (cold connect)
ok
41.89 ms
10,742 bytes (~2,686 tok)
First TCP connect + live snapshot ( 58× faster than uiautomator )
DroidPath Hybrid Vision ( include_vision: true )
ok
452.33 ms
38,581 bytes JPEG (540×1212)
54 live UI nodes + compressed JPEG (q45, 6.1× smaller than raw PNG)
adb exec-out screencap -p (raw PNG)
ok
658.1 ms
235,258 bytes PNG
Visual capture only — no structured UI tree
adb exec-out uiautomator dump /dev/tty (XML)
ok
2,423.9 ms
29,357 bytes (~7,340 tok)
Requires stopping UiAutomation daemon; 2.7× more text tokens than Shadow DOM
2. Live End-to-End Hardware Navigation (Google Pixel 10a · Real Seed Hydration → Warm A*)
Task ( Settings → Target )
Cold-Start (Seed Hydration)
Cold Status
Warm A* Graph
Warm Status
Live Success Rate
Settings → Display
378.6 ms
arrived
296.2 ms (1 step)
arrived
100%
Settings → Battery
655.9 ms
arrived
642.8 ms (1 step)
arrived
100%
Settings → About phone (scroll required)
775.7 ms
arrived
744.1 ms (1 step)
arrived
100%
Average across live suite
603 ms
100% (3/3)
561 ms
100% (3/3)
100%
3. Local MCP Runtime Micro-Benchmarks (Measured)
Metric
Measured Value
tools/list JSON-RPC latency (9 tools)
0.25 ms
android_graph_stats JSON-RPC latency
0.42 ms
Pre-hydrated Seed Graph loaded at boot
20 vertices / 19 edges
Compact Shadow DOM tree (60-node mock Settings screen)
4,537 bytes (~1,135 tokens)
4. Architecture & Cost Comparison (5 WebArena-Style Multi-Step Tasks · Replay + API Cost Model)
Architecture
Avg LLM Calls / Task
Avg Tokens / Task
Avg Est. Task Time
Est. Cost / 1,000 Tasks
Advantage of DroidPath
DroidPath (Shadow DOM + A Router) *
1.0
39 (warm A )* / ~1,135 (cold tree)
464 ms (561 ms live warm)
$0.12
100% live E2E success · 192× cheaper than Vision
Pure Vision (Screenshots + VLM)
4.2
7,673
9,954 ms
$23.02
Baseline
Appium / uiautomator XML
4.2
29,400
8,946 ms
$88.20
735× more expensive than Warm DroidPath
Per-Task Breakdown
Task
Hops
DroidPath (Calls / Tokens / Time)
Pure Vision (Calls / Tokens / Time)
Appium XML (Calls / Tokens / Time)
Enable Night Light (Display)
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Open Wi-Fi Settings (Network)
4
1 call · 39 tok · 580 ms
5 calls · 9,135 tok · 11,850 ms
5 calls · 35,000 tok · 10,650 ms
Navigate to Battery Saver
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Open App Notifications
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Scroll to About Phone
3
1 call · 39 tok · 435 ms
4 calls · 7,308 tok · 9,480 ms
4 calls · 28,000 tok · 8,520 ms
Methodology & Token Model
Text tokens: ceil(utf8_bytes / 4) (standard BPE approximation for JSON/XML).
Cost / 1,000 tasks: Computed transparently at $3.00 / 1M input tokens (standard Claude Sonnet-class API pricing): (avg_tokens_per_task * 1,000 / 1,000,000) * $3.00 .
Vision image tokens: Anthropic's published (W × H) / 750 formula for a 1080×2400 Android screenshot scaled to 1568px long edge ( 706 × 1568 / 750 = 1,477 image tokens + ~350 prompt/tool tokens = 1,827 tokens/step ).
Opt-in Hybrid Vision ( include_vision: true ): Downscaled 540px-wide JPEG ( quality: 45 , measured 38.6 KB on unlocked Pixel 10a Settings vs 235.3 KB raw PNG) is strictly opt-in and only used when an agent explicitly needs visual disambiguation (e.g., canvas/icons), remaining strictly blocked on FLAG_SECURE screens .
# Build release binary first (optional — falls back to target/debug)
cargo build --release -p droidpath-mcp
# Replay benchmark (zero device required)
node scripts/benchmark_agentic.mjs
# Live physical device / emulator probe (measures real screencap -p & uiautomator dump)
node scripts/benchmark_agentic.mjs --mode live
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
