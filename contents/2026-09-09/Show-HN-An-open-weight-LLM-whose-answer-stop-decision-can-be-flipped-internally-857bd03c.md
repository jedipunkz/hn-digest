---
source: "https://github.com/theonlypal/PCCG-Qwen3-4B-continuation-control"
hn_url: "https://news.ycombinator.com/item?id=49635501"
title: "Show HN: An open-weight LLM whose answer/stop decision can be flipped internally"
article_title: "GitHub - theonlypal/PCCG-Qwen3-4B-continuation-control: Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. · GitHub"
image: "https://opengraph.githubassets.com/2326f732f00f1b6760bfd8ce71cdf41ca6a84d7935f3e57eb6c2d4d9a2043444/theonlypal/PCCG-Qwen3-4B-continuation-control"
author: "rayanpal_"
captured_at: "2026-09-09T22:58:33Z"
capture_tool: "hn-digest"
hn_id: 49635501
score: 1
comments: 0
posted_at: "2026-09-09T22:32:34Z"
tags:
  - hacker-news
---

# Show HN: An open-weight LLM whose answer/stop decision can be flipped internally

- HN: [49635501](https://news.ycombinator.com/item?id=49635501)
- Source: [github.com](https://github.com/theonlypal/PCCG-Qwen3-4B-continuation-control)
- Score: 1
- Comments: 0
- Posted: 2026-09-09T22:32:34Z

## Translation

Title: Show HN: An open-weight LLM whose answer/stop decision can be flipped internally
Article title: GitHub - theonlypal/PCCG-Qwen3-4B-continuation-control: Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. · GitHub
Description: Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. - theonlypal/PCCG-Qwen3-4B-continuation-control

Article text:
GitHub - theonlypal/PCCG-Qwen3-4B-continuation-control: Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. · GitHub
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
theonlypal
/
PCCG-Qwen3-4B-continuation-control
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
protocol protocol results results source source .gitattributes .gitattributes LICENSE LICENSE METHOD.md METHOD.md README.md README.md REPRODUCE.md REPRODUCE.md SHA256SUMS SHA256SUMS reproduce.py reproduce.py requirements.txt requirements.txt verify.py verify.py View all files Repository files navigation
Prompt, weights, and generated reasoning prefix held fixed during intervention.
Rayan Pal
getswiftapi.com
Independent Researcher
Thinking-enabled PCCG-Qwen3-4B with an equality-conditioned continuation policy. One fixed 2,560-dimensional activation direction at block 29, applied after the model generated </think> . Intervention strength: 2.0. GO is followed by native EOS.
Anthropic's Jacobian Lens provides internal vocabulary readouts. Additive activation interventions measure causal effects. Method .
Clone with Git LFS installed. The fitted lens is stored in Git LFS.
git lfs pull
python verify.py
Checks package hashes and recounts the confirmation records. Runs on CPU using the Python standard library.
Commands and model requirements .
results/ : complete study records, logits, activations, fitted lens, and fixed direction.
source/ : executed study code and Anthropic's Jacobian Lens implementation.
protocol/ : prompts, case sets, model qualification record, and reproduction settings.
SHA256SUMS : package file hashes.
LICENSE : Apache 2.0. The vendored implementation retains its license in source/vendor/LICENSE .
Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. - theonlypal/PCCG-Qwen3-4B-continuation-control

GitHub - theonlypal/PCCG-Qwen3-4B-continuation-control: Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention. · GitHub
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
theonlypal
/
PCCG-Qwen3-4B-continuation-control
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
protocol protocol results results source source .gitattributes .gitattributes LICENSE LICENSE METHOD.md METHOD.md README.md README.md REPRODUCE.md REPRODUCE.md SHA256SUMS SHA256SUMS reproduce.py reproduce.py requirements.txt requirements.txt verify.py verify.py View all files Repository files navigation
Prompt, weights, and generated reasoning prefix held fixed during intervention.
Rayan Pal
getswiftapi.com
Independent Researcher
Thinking-enabled PCCG-Qwen3-4B with an equality-conditioned continuation policy. One fixed 2,560-dimensional activation direction at block 29, applied after the model generated </think> . Intervention strength: 2.0. GO is followed by native EOS.
Anthropic's Jacobian Lens provides internal vocabulary readouts. Additive activation interventions measure causal effects. Method .
Clone with Git LFS installed. The fitted lens is stored in Git LFS.
git lfs pull
python verify.py
Checks package hashes and recounts the confirmation records. Runs on CPU using the Python standard library.
Commands and model requirements .
results/ : complete study records, logits, activations, fitted lens, and fixed direction.
source/ : executed study code and Anthropic's Jacobian Lens implementation.
protocol/ : prompts, case sets, model qualification record, and reproduction settings.
SHA256SUMS : package file hashes.
LICENSE : Apache 2.0. The vendored implementation retains its license in source/vendor/LICENSE .
Causal continuation-control study for PCCG-Qwen3-4B using Jacobian Lens and activation intervention.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
