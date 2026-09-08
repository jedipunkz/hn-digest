---
source: "https://github.com/aminueza/Keyfence"
hn_url: "https://news.ycombinator.com/item?id=49615382"
title: "Show HN: Keyfence is a local proxy that stops secrets from reaching LLM APIs"
article_title: "GitHub - aminueza/Keyfence: Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. · GitHub"
image: "https://opengraph.githubassets.com/39e12d5b8729d1f9b5bf7fbbfb681b35d2b0bf60bba5905d6f66c12092034ae8/aminueza/Keyfence"
author: "felladrin"
captured_at: "2026-09-08T19:45:52Z"
capture_tool: "hn-digest"
hn_id: 49615382
score: 1
comments: 0
posted_at: "2026-09-08T19:16:46Z"
tags:
  - hacker-news
---

# Show HN: Keyfence is a local proxy that stops secrets from reaching LLM APIs

- HN: [49615382](https://news.ycombinator.com/item?id=49615382)
- Source: [github.com](https://github.com/aminueza/Keyfence)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T19:16:46Z

## Translation

Title: Show HN: Keyfence is a local proxy that stops secrets from reaching LLM APIs
Article title: GitHub - aminueza/Keyfence: Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. · GitHub
Description: Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. - aminueza/Keyfence

Article text:
GitHub - aminueza/Keyfence: Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. · GitHub
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
aminueza
/
Keyfence
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows bench bench docs docs keyfence keyfence tests tests .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.yaml compose.yaml config.example.yaml config.example.yaml docker-entrypoint.sh docker-entrypoint.sh pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A local proxy that stops secrets from reaching LLM APIs. It checks every
request to an AI provider before it leaves your machine and blocks, redacts
or placeholder-swaps API keys, passwords and other secrets. Works with Claude
Code, Cursor, Codex, Aider, curl and anything else that speaks HTTP.
pip install keyfence
Python 3.12 or newer. mitmproxy comes as a dependency.
keyfence import # register your secrets from .env and credential files (hashes only)
keyfence exec -- claude # run a tool through the proxy
keyfence canary .env # plant a fake secret; if a tool ever sends it, you will know
On first run mitmproxy creates a CA certificate in ~/.mitmproxy/ . Trust it
once so HTTPS can be inspected (macOS shown, other systems in the
setup guide ):
sudo security add-trusted-cert -d -p ssl \
-k /Library/Keychains/System.keychain ~ /.mitmproxy/mitmproxy-ca-cert.pem
Modes
mode
behaviour
block
request gets a 403 and is not sent
redact (default)
secret becomes [REDACTED:<kind>]
placeholder
secret becomes <<SECRET_id>> and the real value is restored in the response, streaming included
Documentation
Setup : CA
certificate, manual proxy setup, Docker, all commands.
Detection :
the vault, pattern rules, entropy check, what is excluded, the audit log.
Configuration :
every option, environment variables, the system prompt notice.
Benchmark :
recall by secret format and false positive rate by content type, against
gitleaks, reproducible with python bench/run.py .
Limitations :
what keyfence does not cover and what to combine it with.
Development :
tests, coverage gate, integration script.
MIT. Bundled detection rules come from gitleaks ,
also MIT.
Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. - aminueza/Keyfence

GitHub - aminueza/Keyfence: Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool. · GitHub
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
aminueza
/
Keyfence
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows bench bench docs docs keyfence keyfence tests tests .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.yaml compose.yaml config.example.yaml config.example.yaml docker-entrypoint.sh docker-entrypoint.sh pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A local proxy that stops secrets from reaching LLM APIs. It checks every
request to an AI provider before it leaves your machine and blocks, redacts
or placeholder-swaps API keys, passwords and other secrets. Works with Claude
Code, Cursor, Codex, Aider, curl and anything else that speaks HTTP.
pip install keyfence
Python 3.12 or newer. mitmproxy comes as a dependency.
keyfence import # register your secrets from .env and credential files (hashes only)
keyfence exec -- claude # run a tool through the proxy
keyfence canary .env # plant a fake secret; if a tool ever sends it, you will know
On first run mitmproxy creates a CA certificate in ~/.mitmproxy/ . Trust it
once so HTTPS can be inspected (macOS shown, other systems in the
setup guide ):
sudo security add-trusted-cert -d -p ssl \
-k /Library/Keychains/System.keychain ~ /.mitmproxy/mitmproxy-ca-cert.pem
Modes
mode
behaviour
block
request gets a 403 and is not sent
redact (default)
secret becomes [REDACTED:<kind>]
placeholder
secret becomes <<SECRET_id>> and the real value is restored in the response, streaming included
Documentation
Setup : CA
certificate, manual proxy setup, Docker, all commands.
Detection :
the vault, pattern rules, entropy check, what is excluded, the audit log.
Configuration :
every option, environment variables, the system prompt notice.
Benchmark :
recall by secret format and false positive rate by content type, against
gitleaks, reproducible with python bench/run.py .
Limitations :
what keyfence does not cover and what to combine it with.
Development :
tests, coverage gate, integration script.
MIT. Bundled detection rules come from gitleaks ,
also MIT.
Local proxy that keeps your API keys and secrets out of LLM requests. Works with Claude Code, Cursor, Codex and any tool.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
