---
source: "https://github.com/ctrlb-hq/ctrlb-decompose"
hn_url: "https://news.ycombinator.com/item?id=49625553"
title: "Show HN: Ctrlb-decompose: Strip the noise before senting it to LLMs"
article_title: "GitHub - ctrlb-hq/ctrlb-decompose: LLM-ready reasoning surface over logs · GitHub"
image: "https://opengraph.githubassets.com/f6a9bac216931806cee87f07e06aa4980477d7fab38d4c516a7c9a793ce98898/ctrlb-hq/ctrlb-decompose"
author: "ruhani_grover"
captured_at: "2026-09-09T12:46:29Z"
capture_tool: "hn-digest"
hn_id: 49625553
score: 11
comments: 0
posted_at: "2026-09-09T12:36:09Z"
tags:
  - hacker-news
---

# Show HN: Ctrlb-decompose: Strip the noise before senting it to LLMs

- HN: [49625553](https://news.ycombinator.com/item?id=49625553)
- Source: [github.com](https://github.com/ctrlb-hq/ctrlb-decompose)
- Score: 11
- Comments: 0
- Posted: 2026-09-09T12:36:09Z

## Translation

Title: Show HN: Ctrlb-decompose: Strip the noise before senting it to LLMs
Article title: GitHub - ctrlb-hq/ctrlb-decompose: LLM-ready reasoning surface over logs · GitHub
Description: LLM-ready reasoning surface over logs. Contribute to ctrlb-hq/ctrlb-decompose development by creating an account on GitHub.

Article text:
GitHub - ctrlb-hq/ctrlb-decompose: LLM-ready reasoning surface over logs · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
ctrlb-hq
/
ctrlb-decompose
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
26 Commits 26 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets benches benches plugin plugin research-paper research-paper src src tests tests web web .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md dist-workspace.toml dist-workspace.toml View all files Repository files navigation
Turn millions of noisy log lines into compact patterns with typed variables, quantiles, anomalies, and LLM-ready output.
Runs as a CLI, in the browser via WASM, or as a Rust library — no logs ever leave your machine.
No install, no signup — the whole pipeline (CLP encoding, Drain3 clustering, typing, stats) runs client-side in WebAssembly.
Paste or drop in a log file (or click Generate Example to try it with sample data)
Hit Analyze and watch thousands of lines collapse into a handful of typed patterns
A real 56,482-line Apache-style error log, decomposed live in the browser demo :
Before — raw, repetitive, un-skimmable:
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating channel.jni:jni ( channel.jni, jni)
[Thu Jun 09 06:07:05 2005] [error] config.update(): Can't create channel.jni:jni
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating vm: ( vm, )
[Thu Jun 09 06:07:05 2005] [error] config.update(): Can't create vm:
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating worker.jni:onStartup ( worker.jni, onStartup)
... 56,477 more lines like this ...
After — 25 typed patterns, ranked by volume, in 246ms:
ctrlb-decompose: 56,482 lines → 25 patterns (100.0% reduction)
Time range: 06:07:04 UTC → 03:49:01 UTC
Pattern #1 [ERROR] (20,862 occurrences, 36.9%)
"<TS> [error] [client <*>] <*> <*> <*> <*> <*>"
Variables:
IPv4: 456 unique values
String: 46 unique values
Enum: does (100.0%), to (0.0%)
Enum: not (100.0%), serve (0.0%)
Enum: exist (100.0%), directory (0.0%)
String: 81 unique values
Pattern #2 (7,044 occurrences, 12.5%)
...
Output modes
HUMAN mode — colored, for terminal investigation:
LLM OPTIMIZE mode — compact markdown for feeding into an LLM:
Resource
Live browser demo
ctrlb.ai/decompose
Claude Code plugin
github.com/ctrlb-hq/ctrlb-decompose/tree/main/plugin
Research paper
ctrlb.ai/research
How It Works
ctrlb-decompose uses a two-stage normalization and clustering pipeline that processes logs in a single streaming pass with minimal memory footprint.
┌──────────────────────────────────────────────┐
│ ctrlb-decompose pipeline │
└──────────────────────────────────────────────┘
Raw Log Lines
│
▼
┌──────────────┐ Strip & parse timestamps (ISO 8601, Apache,
│ Timestamp │ syslog, Unix epoch, etc.) into normalized
│ Extraction │ <TS> markers with DateTime values.
└──────┬───────┘
│
▼
┌──────────────┐ Replace integers, floats, IPs, and strings
│ CLP │ with compact placeholder bytes. Structurally
│ Encoding │ identical lines now produce the same "logtype."
└──────┬───────┘
│
▼
┌──────────────┐ Tree-based similarity clustering (Drain3) groups
│ Drain3 │ logtypes into patterns. Differing tokens become
│ Clustering │ <*> wildcards. Incremental — no second pass needed.
└──────┬───────┘
│
▼
┌──────────────┐ Merge CLP-decoded values with Drain3 wildcard
│ Variable │ positions. Classify each variable into semantic
│ Extraction │ types: IPv4, UUID, Duration, Enum, Integer, etc.
│ & Typing │
└──────┬───────┘
│
▼
┌──────────────┐ DDSketch quantiles (p50/p99), HyperLogLog
│ Statistics │ cardinality estimation, top-k values, temporal
│ Accumulation │ bucketing, and reservoir-sampled example lines.
└──────┬───────┘
│
▼
┌──────────────┐ Frequency spikes, error cascades, low-cardinality
│ Anomaly │ flags, bimodal distributions, and clustered
│ Detection │ numeric detection.
└──────┬───────┘
│
▼
┌──────────────┐ Keyword-based severity (ERROR > WARN > INFO > DEBUG),
│ Scoring │ temporal co-occurrence, shared variable correlation,
│ & Correlation│ and error cascade detection across patterns.
└──────┬───────┘
│
▼
┌──────────────┐
│ Output │──── Human (ANSI terminal) / LLM (compact markdown) / JSON
└──────────────┘
Stage 1 — CLP Encoding
CLP (Compressed Log Processor) encoding normalizes variable tokens into typed placeholders, so structurally identical lines produce identical logtypes regardless of the actual values:
Input: "Request from 10.0.1.15 completed in 45ms status=200"
Logtype: "Request from <dict> completed in <float>ms status=<int>"
Stage 2 — Drain3 Clustering
The Drain algorithm builds a prefix tree over logtypes and groups them by token similarity (configurable threshold, default 0.4). Where tokens diverge, the template gains a <*> wildcard. This runs incrementally — each line is processed once with no second pass.
Extracted variables are classified into semantic types for richer analysis:
Drain3 clusters : O(k) with LRU eviction (default 10k max)
Quantiles : DDSketch — fixed ~200 bytes per numeric slot, no raw value storage
Cardinality : HyperLogLog++ — ~200 bytes per high-cardinality variable
Examples : Reservoir sampling — bounded buffer per pattern
brew tap ctrlb-hq/tap
brew install ctrlb-decompose
Debian / Ubuntu
curl -LO https://github.com/ctrlb-hq/ctrlb-decompose/releases/download/v0.1.0/ctrlb-decompose_0.1.0-1_amd64.deb
sudo dpkg -i ctrlb-decompose_0.1.0-1_amd64.deb
Build from source
git clone https://github.com/ctrlb-hq/ctrlb-decompose.git
cd ctrlb-decompose
cargo build --release
# Binary at target/release/ctrlb-decompose
Usage
# Pipe from stdin
cat /var/log/syslog | ctrlb-decompose
# Read from file
ctrlb-decompose server.log
# LLM-optimized output (compact, token-efficient)
ctrlb-decompose --llm app.log
# JSON output
ctrlb-decompose --json app.log
# Top 10 patterns with 3 example lines each
ctrlb-decompose --top 10 --context 3 app.log
Options
ctrlb-decompose [OPTIONS] [FILE]
Arguments:
[FILE] Log file path (reads stdin if omitted or "-")
Options:
--human Human-readable output with colors (default)
--llm LLM-optimized compact markdown
--json Structured JSON output
--top <N> Show top N patterns (default: 20)
--context <N> Example lines per pattern (default: 0)
--no-color Disable ANSI colors
--no-banner Suppress header/footer
-q, --quiet Suppress progress messages
-h, --help Show help
-V, --version Show version
Output Formats
ctrlb-decompose has three output modes, all driven off the same analysis pass — pick the one that fits where you're reading it. See the before/after screenshots above for --human and --llm side by side in the browser demo.
┌──────────────────────────────────────────────────────────────────┐
│ ctrlb-decompose: 80,000 lines -> 3 patterns (100.0% reduction) │
└──────────────────────────────────────────────────────────────────┘
Time range: 14:22:01 UTC -> 16:35:10 UTC
Pattern #1 (75,211 occurrences, 94.0%)
"<TS> INFO [<*>] Request from <*> completed in <*> status=<*>"
Variables:
HexID: 804 unique values
IPv4: 27 unique values
Duration: mean=46, p50=45, p99=116, min=3, max=169
Integer: mean=224, p50=198, p99=498, min=200, max=503
Pattern #2 [WARN] (2,773 occurrences, 3.5%)
"<TS> WARN [<*>] Connection pool exhausted, waiting <*>"
Variables:
HexID: 726 unique values
Duration: mean=526, p50=529, p99=889, min=150, max=900
LLM — ctrlb-decompose server.log --llm
Compact, token-efficient markdown designed to be pasted straight into a prompt — see the LLM OPTIMIZE screenshot above for a full real-world example against a 56K-line log.
{
"summary" : {
"total_lines" : 80000 ,
"pattern_count" : 3 ,
"patterns_shown" : 1 ,
"patterns_omitted" : 2 ,
"time_range" : {
"start" : " 2026-08-21T14:22:01.051+00:00 " ,
"end" : " 2026-08-21T16:35:10.707+00:00 "
}
},
"patterns" : [
{
"id" : 1 ,
"template" : " <TS> INFO [<*>] Request from <*> completed in <*> status=<*> " ,
"count" : 75211 ,
"frequency_pct" : 94.0 ,
"severity" : " info " ,
"variables" : [
{ "slot" : 0 , "type" : " HexID " , "unique_count" : 804 },
{ "slot" : 1 , "type" : " IPv4 " , "unique_count" : 27 },
{ "slot" : 2 , "type" : " Duration " , "unique_count" : 157 },
{ "slot" : 3 , "type" : " Integer " , "unique_count" : 4 }
]
}
]
}
Claude Code Plugin
Use ctrlb-decompose directly from Claude Code — no CLI knowledge needed. The plugin installs ctrlb-decompose automatically and lets you analyze logs just by asking.
/plugin marketplace add ctrlb-hq/ctrlb-decompose
/plugin install ctrlb-decompose@ctrlb-hq
Usage
Just describe what you want in plain language:
"Analyze the errors in /var/log/app.log "
"What are the most common patterns in this log file?"
"Summarize these logs and highlight anomalies"
Claude will check if ctrlb-decompose is installed (and walk you through installation if not), run the analysis, and explain the results — surfacing errors first, calling out anomalies, and suggesting what to investigate next.
See github.com/ctrlb-hq/ctrlb-decompose/tree/main/plugin for full details.
LLM-ready reasoning surface over logs
Readme MIT license Activity Custom properties Stars
23 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

LLM-ready reasoning surface over logs. Contribute to ctrlb-hq/ctrlb-decompose development by creating an account on GitHub.

GitHub - ctrlb-hq/ctrlb-decompose: LLM-ready reasoning surface over logs · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
ctrlb-hq
/
ctrlb-decompose
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
26 Commits 26 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets benches benches plugin plugin research-paper research-paper src src tests tests web web .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md dist-workspace.toml dist-workspace.toml View all files Repository files navigation
Turn millions of noisy log lines into compact patterns with typed variables, quantiles, anomalies, and LLM-ready output.
Runs as a CLI, in the browser via WASM, or as a Rust library — no logs ever leave your machine.
No install, no signup — the whole pipeline (CLP encoding, Drain3 clustering, typing, stats) runs client-side in WebAssembly.
Paste or drop in a log file (or click Generate Example to try it with sample data)
Hit Analyze and watch thousands of lines collapse into a handful of typed patterns
A real 56,482-line Apache-style error log, decomposed live in the browser demo :
Before — raw, repetitive, un-skimmable:
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating channel.jni:jni ( channel.jni, jni)
[Thu Jun 09 06:07:05 2005] [error] config.update(): Can't create channel.jni:jni
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating vm: ( vm, )
[Thu Jun 09 06:07:05 2005] [error] config.update(): Can't create vm:
[Thu Jun 09 06:07:05 2005] [error] env.createBean2(): Factory error creating worker.jni:onStartup ( worker.jni, onStartup)
... 56,477 more lines like this ...
After — 25 typed patterns, ranked by volume, in 246ms:
ctrlb-decompose: 56,482 lines → 25 patterns (100.0% reduction)
Time range: 06:07:04 UTC → 03:49:01 UTC
Pattern #1 [ERROR] (20,862 occurrences, 36.9%)
"<TS> [error] [client <*>] <*> <*> <*> <*> <*>"
Variables:
IPv4: 456 unique values
String: 46 unique values
Enum: does (100.0%), to (0.0%)
Enum: not (100.0%), serve (0.0%)
Enum: exist (100.0%), directory (0.0%)
String: 81 unique values
Pattern #2 (7,044 occurrences, 12.5%)
...
Output modes
HUMAN mode — colored, for terminal investigation:
LLM OPTIMIZE mode — compact markdown for feeding into an LLM:
Resource
Live browser demo
ctrlb.ai/decompose
Claude Code plugin
github.com/ctrlb-hq/ctrlb-decompose/tree/main/plugin
Research paper
ctrlb.ai/research
How It Works
ctrlb-decompose uses a two-stage normalization and clustering pipeline that processes logs in a single streaming pass with minimal memory footprint.
┌──────────────────────────────────────────────┐
│ ctrlb-decompose pipeline │
└──────────────────────────────────────────────┘
Raw Log Lines
│
▼
┌──────────────┐ Strip & parse timestamps (ISO 8601, Apache,
│ Timestamp │ syslog, Unix epoch, etc.) into normalized
│ Extraction │ <TS> markers with DateTime values.
└──────┬───────┘
│
▼
┌──────────────┐ Replace integers, floats, IPs, and strings
│ CLP │ with compact placeholder bytes. Structurally
│ Encoding │ identical lines now produce the same "logtype."
└──────┬───────┘
│
▼
┌──────────────┐ Tree-based similarity clustering (Drain3) groups
│ Drain3 │ logtypes into patterns. Differing tokens become
│ Clustering │ <*> wildcards. Incremental — no second pass needed.
└──────┬───────┘
│
▼
┌──────────────┐ Merge CLP-decoded values with Drain3 wildcard
│ Variable │ positions. Classify each variable into semantic
│ Extraction │ types: IPv4, UUID, Duration, Enum, Integer, etc.
│ & Typing │
└──────┬───────┘
│
▼
┌──────────────┐ DDSketch quantiles (p50/p99), HyperLogLog
│ Statistics │ cardinality estimation, top-k values, temporal
│ Accumulation │ bucketing, and reservoir-sampled example lines.
└──────┬───────┘
│
▼
┌──────────────┐ Frequency spikes, error cascades, low-cardinality
│ Anomaly │ flags, bimodal distributions, and clustered
│ Detection │ numeric detection.
└──────┬───────┘
│
▼
┌──────────────┐ Keyword-based severity (ERROR > WARN > INFO > DEBUG),
│ Scoring │ temporal co-occurrence, shared variable correlation,
│ & Correlation│ and error cascade detection across patterns.
└──────┬───────┘
│
▼
┌──────────────┐
│ Output │──── Human (ANSI terminal) / LLM (compact markdown) / JSON
└──────────────┘
Stage 1 — CLP Encoding
CLP (Compressed Log Processor) encoding normalizes variable tokens into typed placeholders, so structurally identical lines produce identical logtypes regardless of the actual values:
Input: "Request from 10.0.1.15 completed in 45ms status=200"
Logtype: "Request from <dict> completed in <float>ms status=<int>"
Stage 2 — Drain3 Clustering
The Drain algorithm builds a prefix tree over logtypes and groups them by token similarity (configurable threshold, default 0.4). Where tokens diverge, the template gains a <*> wildcard. This runs incrementally — each line is processed once with no second pass.
Extracted variables are classified into semantic types for richer analysis:
Drain3 clusters : O(k) with LRU eviction (default 10k max)
Quantiles : DDSketch — fixed ~200 bytes per numeric slot, no raw value storage
Cardinality : HyperLogLog++ — ~200 bytes per high-cardinality variable
Examples : Reservoir sampling — bounded buffer per pattern
brew tap ctrlb-hq/tap
brew install ctrlb-decompose
Debian / Ubuntu
curl -LO https://github.com/ctrlb-hq/ctrlb-decompose/releases/download/v0.1.0/ctrlb-decompose_0.1.0-1_amd64.deb
sudo dpkg -i ctrlb-decompose_0.1.0-1_amd64.deb
Build from source
git clone https://github.com/ctrlb-hq/ctrlb-decompose.git
cd ctrlb-decompose
cargo build --release
# Binary at target/release/ctrlb-decompose
Usage
# Pipe from stdin
cat /var/log/syslog | ctrlb-decompose
# Read from file
ctrlb-decompose server.log
# LLM-optimized output (compact, token-efficient)
ctrlb-decompose --llm app.log
# JSON output
ctrlb-decompose --json app.log
# Top 10 patterns with 3 example lines each
ctrlb-decompose --top 10 --context 3 app.log
Options
ctrlb-decompose [OPTIONS] [FILE]
Arguments:
[FILE] Log file path (reads stdin if omitted or "-")
Options:
--human Human-readable output with colors (default)
--llm LLM-optimized compact markdown
--json Structured JSON output
--top <N> Show top N patterns (default: 20)
--context <N> Example lines per pattern (default: 0)
--no-color Disable ANSI colors
--no-banner Suppress header/footer
-q, --quiet Suppress progress messages
-h, --help Show help
-V, --version Show version
Output Formats
ctrlb-decompose has three output modes, all driven off the same analysis pass — pick the one that fits where you're reading it. See the before/after screenshots above for --human and --llm side by side in the browser demo.
┌──────────────────────────────────────────────────────────────────┐
│ ctrlb-decompose: 80,000 lines -> 3 patterns (100.0% reduction) │
└──────────────────────────────────────────────────────────────────┘
Time range: 14:22:01 UTC -> 16:35:10 UTC
Pattern #1 (75,211 occurrences, 94.0%)
"<TS> INFO [<*>] Request from <*> completed in <*> status=<*>"
Variables:
HexID: 804 unique values
IPv4: 27 unique values
Duration: mean=46, p50=45, p99=116, min=3, max=169
Integer: mean=224, p50=198, p99=498, min=200, max=503
Pattern #2 [WARN] (2,773 occurrences, 3.5%)
"<TS> WARN [<*>] Connection pool exhausted, waiting <*>"
Variables:
HexID: 726 unique values
Duration: mean=526, p50=529, p99=889, min=150, max=900
LLM — ctrlb-decompose server.log --llm
Compact, token-efficient markdown designed to be pasted straight into a prompt — see the LLM OPTIMIZE screenshot above for a full real-world example against a 56K-line log.
{
"summary" : {
"total_lines" : 80000 ,
"pattern_count" : 3 ,
"patterns_shown" : 1 ,
"patterns_omitted" : 2 ,
"time_range" : {
"start" : " 2026-08-21T14:22:01.051+00:00 " ,
"end" : " 2026-08-21T16:35:10.707+00:00 "
}
},
"patterns" : [
{
"id" : 1 ,
"template" : " <TS> INFO [<*>] Request from <*> completed in <*> status=<*> " ,
"count" : 75211 ,
"frequency_pct" : 94.0 ,
"severity" : " info " ,
"variables" : [
{ "slot" : 0 , "type" : " HexID " , "unique_count" : 804 },
{ "slot" : 1 , "type" : " IPv4 " , "unique_count" : 27 },
{ "slot" : 2 , "type" : " Duration " , "unique_count" : 157 },
{ "slot" : 3 , "type" : " Integer " , "unique_count" : 4 }
]
}
]
}
Claude Code Plugin
Use ctrlb-decompose directly from Claude Code — no CLI knowledge needed. The plugin installs ctrlb-decompose automatically and lets you analyze logs just by asking.
/plugin marketplace add ctrlb-hq/ctrlb-decompose
/plugin install ctrlb-decompose@ctrlb-hq
Usage
Just describe what you want in plain language:
"Analyze the errors in /var/log/app.log "
"What are the most common patterns in this log file?"
"Summarize these logs and highlight anomalies"
Claude will check if ctrlb-decompose is installed (and walk you through installation if not), run the analysis, and explain the results — surfacing errors first, calling out anomalies, and suggesting what to investigate next.
See github.com/ctrlb-hq/ctrlb-decompose/tree/main/plugin for full details.
LLM-ready reasoning surface over logs
Readme MIT license Activity Custom properties Stars
23 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
