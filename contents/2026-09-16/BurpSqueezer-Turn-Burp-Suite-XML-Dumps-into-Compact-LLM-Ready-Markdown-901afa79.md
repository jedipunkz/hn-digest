---
source: "https://github.com/vaginskii/BurpSqueezer"
hn_url: "https://news.ycombinator.com/item?id=49727521"
title: "BurpSqueezer – Turn Burp Suite XML Dumps into Compact LLM-Ready Markdown"
article_title: "GitHub - vaginskii/BurpSqueezer: Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. · GitHub"
image: "https://opengraph.githubassets.com/e79704eab3a0f300fd6af1d99eca09d9058f8a266a7b48bde8b8b9ed4a49676f/vaginskii/BurpSqueezer"
author: "vaginskii"
captured_at: "2026-09-16T14:39:17Z"
capture_tool: "hn-digest"
hn_id: 49727521
score: 1
comments: 1
posted_at: "2026-09-16T14:24:51Z"
tags:
  - hacker-news
---

# BurpSqueezer – Turn Burp Suite XML Dumps into Compact LLM-Ready Markdown

- HN: [49727521](https://news.ycombinator.com/item?id=49727521)
- Source: [github.com](https://github.com/vaginskii/BurpSqueezer)
- Score: 1
- Comments: 1
- Posted: 2026-09-16T14:24:51Z

## Translation

Title: BurpSqueezer – Turn Burp Suite XML Dumps into Compact LLM-Ready Markdown
Article title: GitHub - vaginskii/BurpSqueezer: Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. · GitHub
Description: Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. - vaginskii/BurpSqueezer

Article text:
GitHub - vaginskii/BurpSqueezer: Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. · GitHub
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
vaginskii
/
BurpSqueezer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.github/ workflows .github/ workflows src src tests tests Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports.
BurpSqueezer is a security research tool that transforms large Burp Suite HTTP traffic dumps into highly compact, structured Markdown representations designed to be consumed by LLMs.
Instead of giving an LLM a large amount of raw HTTP traffic and expecting it to parse, filter, and connect everything itself, BurpSqueezer performs the initial structural analysis and removes a significant amount of redundant and low-value data.
Large Burp Suite XML dump
│
▼
BurpSqueezer
│
▼
Compact structured Markdown
│
▼
LLM analysis
Why BurpSqueezer?
Large Burp Suite dumps can contain hundreds or thousands of HTTP transactions, including repeated requests, infrastructure noise, dynamic values, headers, responses, and other data that is expensive or impractical to provide directly to an LLM.
BurpSqueezer is designed to reduce this dataset while preserving useful information about the application's underlying structure.
The resulting report can be substantially smaller than the original dump, making it more practical for LLM-based analysis, especially in environments with file-size, context-window, or token-budget limitations.
BurpSqueezer is not an autonomous pentester . It prepares and compresses application traffic into a representation that can be further analyzed by humans or LLMs. Final security conclusions and verification still require manual testing.
A test on a real Burp Suite XML dump demonstrated substantial size reduction:
In the standard test, an approximately 26.7 MB Burp Suite XML dump containing 323 transactions was reduced to approximately 35 KB of structured Markdown.
The original traffic is not included in this repository because real Burp captures may contain sensitive application data, credentials, tokens, or other private information.
Compression results naturally vary depending on the structure and contents of the input dataset.
BurpSqueezer is designed as a universal tool with no hardcoded endpoints or application-specific patterns. Instead of assuming how an API is structured, it uses statistical and heuristic analysis to identify potentially meaningful relationships within the observed traffic.
A key goal is information-dense compression : reducing large HTTP request dumps into much smaller representations while retaining useful structural, relational, and data-flow information.
BurpSqueezer is designed primarily for APIs and applications with meaningful business logic . Simple websites with little or no backend logic may produce significantly less useful results because there may be insufficient structure and relationships for the tool to analyze.
Universal Analysis — No hardcoded endpoints, application patterns, or assumptions about API structure
High Compression — Reduces large HTTP request dumps into significantly smaller reports while preserving relevant analytical information
Noise Filtering — Uses statistical and heuristic techniques to reduce redundant and low-value traffic
Data Flow Tracking — Identifies value propagation and relationships between requests
Structural Analysis — Extracts relationships, sequences, states, and other signals from observed traffic
LLM-Optimized Output — Produces compact Markdown designed to be used as context for LLM-based analysis
Multiple Modes — Adjustable selectivity depending on whether completeness or maximum compression is preferred
BurpSqueezer is written in Rust. Make sure you have the Rust toolchain installed before continuing.
If Rust is not installed, install it from the official Rust website.
Then clone the repository and install BurpSqueezer:
git clone https://github.com/vaginskii/BurpSqueezer.git
cd BurpSqueezer
cargo install --path .
After installation, verify that BurpSqueezer is available:
burpsqueezer --help
If the command displays the available options, the installation was successful.
# Basic usage
burpsqueezer solve burp_dump.xml --output analysis.md
# Maximum selectivity
burpsqueezer solve large_dump.xml --output focused.md --mode apocalyptic
# Lowest selectivity
burpsqueezer solve large_dump.xml --output full.md --mode peaceful
# Verbose output for debugging
burpsqueezer solve test.xml --output report.md --verbose
Modes
peaceful — lower thresholds and a fuller report; preserves more potentially useful information
standard — balanced default mode between information retention and compression
apocalyptic — maximum selectivity; focuses on the strongest structural signals
--output — destination for the Markdown report (required)
--mode — analysis mode (default: standard )
--quiet — silence all progress output
--verbose — emit per-stage detail
BurpSqueezer transforms raw Burp Suite XML traffic into a structured and highly compact Markdown representation intended for LLM-based analysis.
The generated report can contain information about:
API endpoints and their relationships
Parameter and value propagation
Sequences and structural relationships
Statistical relationships between endpoints
Relevant signals identified after noise reduction
The exact output depends on the input dataset and selected analysis mode.
BurpSqueezer is not intended to simply summarize HTTP traffic . Its goal is to produce a compact, security-oriented representation of the underlying API structure and relationships while removing a significant amount of redundant and low-value data.
The original Burp dump can still be useful for manual verification or retrieving information that was intentionally omitted during compression.
BurpSqueezer is an independent security research tool and is not affiliated with, endorsed by, or developed by PortSwigger .
It does not contain or distribute Burp Suite software.
BurpSqueezer operates on HTTP traffic exported by the user from Burp Suite. The input is user-provided Burp Suite XML data; BurpSqueezer does not interact with or send requests to the target application.
This is an experimental security research tool. Results may vary across different Burp collections.
BurpSqueezer is primarily designed for large datasets and applications with meaningful API structure or business logic. Small, highly specialized, or structurally sparse datasets may provide less information for the statistical analysis to work with.
Statistical and heuristic analysis cannot guarantee that every relevant relationship, data flow, or security signal will be identified.
Aggressive compression modes may intentionally discard information in exchange for a smaller output.
BurpSqueezer is intended to assist human and LLM-based analysis, not replace manual security testing or verification .
BurpSqueezer is intended for authorized security testing, penetration testing, bug bounty programs, and security research.
Only analyze HTTP traffic that you are authorized to access. Do not use BurpSqueezer with data obtained from systems or applications without appropriate permission.
The author is not responsible for misuse of the software.
See the LICENSE file for the terms under which BurpSqueezer is distributed.
Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. - vaginskii/BurpSqueezer

GitHub - vaginskii/BurpSqueezer: Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports. · GitHub
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
vaginskii
/
BurpSqueezer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.github/ workflows .github/ workflows src src tests tests Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports.
BurpSqueezer is a security research tool that transforms large Burp Suite HTTP traffic dumps into highly compact, structured Markdown representations designed to be consumed by LLMs.
Instead of giving an LLM a large amount of raw HTTP traffic and expecting it to parse, filter, and connect everything itself, BurpSqueezer performs the initial structural analysis and removes a significant amount of redundant and low-value data.
Large Burp Suite XML dump
│
▼
BurpSqueezer
│
▼
Compact structured Markdown
│
▼
LLM analysis
Why BurpSqueezer?
Large Burp Suite dumps can contain hundreds or thousands of HTTP transactions, including repeated requests, infrastructure noise, dynamic values, headers, responses, and other data that is expensive or impractical to provide directly to an LLM.
BurpSqueezer is designed to reduce this dataset while preserving useful information about the application's underlying structure.
The resulting report can be substantially smaller than the original dump, making it more practical for LLM-based analysis, especially in environments with file-size, context-window, or token-budget limitations.
BurpSqueezer is not an autonomous pentester . It prepares and compresses application traffic into a representation that can be further analyzed by humans or LLMs. Final security conclusions and verification still require manual testing.
A test on a real Burp Suite XML dump demonstrated substantial size reduction:
In the standard test, an approximately 26.7 MB Burp Suite XML dump containing 323 transactions was reduced to approximately 35 KB of structured Markdown.
The original traffic is not included in this repository because real Burp captures may contain sensitive application data, credentials, tokens, or other private information.
Compression results naturally vary depending on the structure and contents of the input dataset.
BurpSqueezer is designed as a universal tool with no hardcoded endpoints or application-specific patterns. Instead of assuming how an API is structured, it uses statistical and heuristic analysis to identify potentially meaningful relationships within the observed traffic.
A key goal is information-dense compression : reducing large HTTP request dumps into much smaller representations while retaining useful structural, relational, and data-flow information.
BurpSqueezer is designed primarily for APIs and applications with meaningful business logic . Simple websites with little or no backend logic may produce significantly less useful results because there may be insufficient structure and relationships for the tool to analyze.
Universal Analysis — No hardcoded endpoints, application patterns, or assumptions about API structure
High Compression — Reduces large HTTP request dumps into significantly smaller reports while preserving relevant analytical information
Noise Filtering — Uses statistical and heuristic techniques to reduce redundant and low-value traffic
Data Flow Tracking — Identifies value propagation and relationships between requests
Structural Analysis — Extracts relationships, sequences, states, and other signals from observed traffic
LLM-Optimized Output — Produces compact Markdown designed to be used as context for LLM-based analysis
Multiple Modes — Adjustable selectivity depending on whether completeness or maximum compression is preferred
BurpSqueezer is written in Rust. Make sure you have the Rust toolchain installed before continuing.
If Rust is not installed, install it from the official Rust website.
Then clone the repository and install BurpSqueezer:
git clone https://github.com/vaginskii/BurpSqueezer.git
cd BurpSqueezer
cargo install --path .
After installation, verify that BurpSqueezer is available:
burpsqueezer --help
If the command displays the available options, the installation was successful.
# Basic usage
burpsqueezer solve burp_dump.xml --output analysis.md
# Maximum selectivity
burpsqueezer solve large_dump.xml --output focused.md --mode apocalyptic
# Lowest selectivity
burpsqueezer solve large_dump.xml --output full.md --mode peaceful
# Verbose output for debugging
burpsqueezer solve test.xml --output report.md --verbose
Modes
peaceful — lower thresholds and a fuller report; preserves more potentially useful information
standard — balanced default mode between information retention and compression
apocalyptic — maximum selectivity; focuses on the strongest structural signals
--output — destination for the Markdown report (required)
--mode — analysis mode (default: standard )
--quiet — silence all progress output
--verbose — emit per-stage detail
BurpSqueezer transforms raw Burp Suite XML traffic into a structured and highly compact Markdown representation intended for LLM-based analysis.
The generated report can contain information about:
API endpoints and their relationships
Parameter and value propagation
Sequences and structural relationships
Statistical relationships between endpoints
Relevant signals identified after noise reduction
The exact output depends on the input dataset and selected analysis mode.
BurpSqueezer is not intended to simply summarize HTTP traffic . Its goal is to produce a compact, security-oriented representation of the underlying API structure and relationships while removing a significant amount of redundant and low-value data.
The original Burp dump can still be useful for manual verification or retrieving information that was intentionally omitted during compression.
BurpSqueezer is an independent security research tool and is not affiliated with, endorsed by, or developed by PortSwigger .
It does not contain or distribute Burp Suite software.
BurpSqueezer operates on HTTP traffic exported by the user from Burp Suite. The input is user-provided Burp Suite XML data; BurpSqueezer does not interact with or send requests to the target application.
This is an experimental security research tool. Results may vary across different Burp collections.
BurpSqueezer is primarily designed for large datasets and applications with meaningful API structure or business logic. Small, highly specialized, or structurally sparse datasets may provide less information for the statistical analysis to work with.
Statistical and heuristic analysis cannot guarantee that every relevant relationship, data flow, or security signal will be identified.
Aggressive compression modes may intentionally discard information in exchange for a smaller output.
BurpSqueezer is intended to assist human and LLM-based analysis, not replace manual security testing or verification .
BurpSqueezer is intended for authorized security testing, penetration testing, bug bounty programs, and security research.
Only analyze HTTP traffic that you are authorized to access. Do not use BurpSqueezer with data obtained from systems or applications without appropriate permission.
The author is not responsible for misuse of the software.
See the LICENSE file for the terms under which BurpSqueezer is distributed.
Turn Burp Suite XML dumps into compact, LLM-ready Markdown reports.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
