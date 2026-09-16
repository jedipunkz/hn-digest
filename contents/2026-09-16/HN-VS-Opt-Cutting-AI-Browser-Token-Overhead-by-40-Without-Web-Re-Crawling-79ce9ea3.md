---
source: "https://github.com/filippocarrucciu1-lang/vs-opt-deterministic-search"
hn_url: "https://news.ycombinator.com/item?id=49723840"
title: "HN: VS-Opt – Cutting AI Browser Token Overhead by 40% Without Web Re-Crawling"
article_title: "GitHub - filippocarrucciu1-lang/vs-opt-deterministic-search · GitHub"
image: "https://opengraph.githubassets.com/a36fd12b25ee6275600c9774dfcc800dfd98b1eedaacad52898e1309a383fbfa/filippocarrucciu1-lang/vs-opt-deterministic-search"
author: "filippo10"
captured_at: "2026-09-16T09:42:35Z"
capture_tool: "hn-digest"
hn_id: 49723840
score: 1
comments: 0
posted_at: "2026-09-16T09:07:56Z"
tags:
  - hacker-news
---

# HN: VS-Opt – Cutting AI Browser Token Overhead by 40% Without Web Re-Crawling

- HN: [49723840](https://news.ycombinator.com/item?id=49723840)
- Source: [github.com](https://github.com/filippocarrucciu1-lang/vs-opt-deterministic-search)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T09:07:56Z

## Translation

Title: HN: VS-Opt – Cutting AI Browser Token Overhead by 40% Without Web Re-Crawling
Article title: GitHub - filippocarrucciu1-lang/vs-opt-deterministic-search · GitHub
Description: Contribute to filippocarrucciu1-lang/vs-opt-deterministic-search development by creating an account on GitHub.

Article text:
GitHub - filippocarrucciu1-lang/vs-opt-deterministic-search · GitHub
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
filippocarrucciu1-lang
/
vs-opt-deterministic-search
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
README.md README.md View all files Repository files navigation
vs-opt-deterministic-search # VS-OPT & GIUSTRA Engine: Deterministic Pre-Query State Verification
Abstract : Reducing LLM/Search token overhead by 30-50% via pre-query intent correction and zero-trust state governance.
Current AI search engines and browser assistants rely on probabilistic user-side iterative refinement . This creates exponential token consumption, server latency, and high OpEx.
VS-OPT (Verify Search Optimizer) replaces continuous scraping and prompt re-generation with a deterministic mathematical pipeline:
KNOWN_STATE → EVENT → REQUIRED_CONTROL → NEW_KNOWN_STATE
Query Optimizer : Rule-based intent filter (Ollama-local / Chromium sandbox compatible). Corrects the prompt before search execution.
GIUSTRA State Verification : Eliminates web re-crawling by verifying source authority, timestamp, and version hic et nunc .
Fallback Sandbox : Zero-tracking fallback matrix verified via deterministic intent mapping (e.g., whatismymovie.com engine integration).
Metric
Standard LLM RAG Pipeline
VS-OPT Pre-Query Pipeline
Token Waste
High (Multi-turn iterative)
Zero (Single-pass intent)
Client Battery Impact
Heavy (Continuous parsing)
Lightweight (Sandbox isolation)
Data Authority
Probabilistic ranking
Ledger-backed state
Integration Target
Designed as a zero-training, zero-migration layer for privacy-first Chromium browsers and distributed state architectures.
#chromium, #brave-search, #llm-optimization, #rust.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to filippocarrucciu1-lang/vs-opt-deterministic-search development by creating an account on GitHub.

GitHub - filippocarrucciu1-lang/vs-opt-deterministic-search · GitHub
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
filippocarrucciu1-lang
/
vs-opt-deterministic-search
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
README.md README.md View all files Repository files navigation
vs-opt-deterministic-search # VS-OPT & GIUSTRA Engine: Deterministic Pre-Query State Verification
Abstract : Reducing LLM/Search token overhead by 30-50% via pre-query intent correction and zero-trust state governance.
Current AI search engines and browser assistants rely on probabilistic user-side iterative refinement . This creates exponential token consumption, server latency, and high OpEx.
VS-OPT (Verify Search Optimizer) replaces continuous scraping and prompt re-generation with a deterministic mathematical pipeline:
KNOWN_STATE → EVENT → REQUIRED_CONTROL → NEW_KNOWN_STATE
Query Optimizer : Rule-based intent filter (Ollama-local / Chromium sandbox compatible). Corrects the prompt before search execution.
GIUSTRA State Verification : Eliminates web re-crawling by verifying source authority, timestamp, and version hic et nunc .
Fallback Sandbox : Zero-tracking fallback matrix verified via deterministic intent mapping (e.g., whatismymovie.com engine integration).
Metric
Standard LLM RAG Pipeline
VS-OPT Pre-Query Pipeline
Token Waste
High (Multi-turn iterative)
Zero (Single-pass intent)
Client Battery Impact
Heavy (Continuous parsing)
Lightweight (Sandbox isolation)
Data Authority
Probabilistic ranking
Ledger-backed state
Integration Target
Designed as a zero-training, zero-migration layer for privacy-first Chromium browsers and distributed state architectures.
#chromium, #brave-search, #llm-optimization, #rust.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
