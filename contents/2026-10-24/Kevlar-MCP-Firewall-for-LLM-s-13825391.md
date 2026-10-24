---
source: "https://github.com/Ablation-Tool/kevlar"
hn_url: "https://news.ycombinator.com/item?id=50226016"
title: "Kevlar = MCP Firewall for LLM's"
article_title: "GitHub - Ablation-Tool/kevlar: Kevlar is a proxy firewall for Model Context Protocol · GitHub"
image: "https://opengraph.githubassets.com/237387ab2ef88075f16bbb086c0887818bfe00d340fe73507f163ccc82a698e5/Ablation-Tool/kevlar"
author: "zellkernel"
captured_at: "2026-10-24T22:55:48Z"
capture_tool: "hn-digest"
hn_id: 50226016
score: 1
comments: 0
posted_at: "2026-10-24T22:24:09Z"
tags:
  - hacker-news
---

# Kevlar = MCP Firewall for LLM's

- HN: [50226016](https://news.ycombinator.com/item?id=50226016)
- Source: [github.com](https://github.com/Ablation-Tool/kevlar)
- Score: 1
- Comments: 0
- Posted: 2026-10-24T22:24:09Z

## Translation

Title: Kevlar = MCP Firewall for LLM's
Article title: GitHub - Ablation-Tool/kevlar: Kevlar is a proxy firewall for Model Context Protocol · GitHub
Description: Kevlar is a proxy firewall for Model Context Protocol - Ablation-Tool/kevlar

Article text:
GitHub - Ablation-Tool/kevlar: Kevlar is a proxy firewall for Model Context Protocol · GitHub
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
Ablation-Tool
kevlar Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Kevlar is a proxy firewall for Model Context Protocol
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
docs docs tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md KEVLAR-01.jpeg KEVLAR-01.jpeg KEVLAR-repo-01.png KEVLAR-repo-01.png LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md kevlar.py kevlar.py policy.py policy.py pyproject.toml pyproject.toml pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
Kevlar is a proxy firewall for Model Context Protocol
LLM [ KEVLAR ] MCP
It sits between (you) and (everyone else)
MCP servers control their own tools.
Those tools (good or bad) go into your LLM blindly, due to trust.
Any MCP server can exploit this to redirect the model, pull credentials, or add any tools they want.
The MCP spec does not prevent any of this.
You do not control what an MCP server puts in its messages.
TL;DR
Kevlar sits between client and server and does four things:
Strips all tool descriptions before they reach the LLM. No exceptions.
Scans tool call responses for injection payloads and either redacts matched spans or blocks the response entirely.
Blocks outbound calls carrying credentials in their arguments, even if the LLM has already been compromised.
Locks the tool registry after the initial handshake so the server cannot inject new tools mid-session.
Kevlar strips every description field to "" unconditionally on every tools/list response, before the message reaches the client. The LLM sees name and inputSchema only. inputSchema is rebuilt from an allowlist: prose fields like description , title , examples , and $comment are stripped there too, and enum values, property keys, $ref targets, and type values are validated against the MCP spec. Score 0 on this pass cannot be achieved by adding description text back, because description text is never forwarded regardless of score.
Tool call results cannot be stripped without breaking functionality. Kevlar scans them against ~60 indicators sourced from published prompt injection research. Responses scoring above the mode threshold are quarantined (matched spans replaced with [KEVLAR:<ID>] tokens) or rejected (replaced with an isError block). Unicode bypass attempts using characters that normalize to injection text under NFKC are detected before normalization and redacted after.
Outbound tools/call arguments are scanned before the call is forwarded to the server. If the LLM has already been compromised and is trying to exfiltrate credentials, the call is blocked at the proxy layer. The server never receives it. Runs in all three modes.
After the initial tools/list handshake, the set of callable tool names is frozen. notifications/tools/list_changed from the server is dropped, not forwarded. Any call to a tool not in the CDR-sanitized set is blocked with a synthetic isError response regardless of call content.
Install
git clone https://github.com/Ablation-Tool/kevlar.git
cd kevlar
pip install .
Requires Python 3.10+. Linux and macOS only. Depends on os.killpg and asyncio.connect_read_pipe .
kevlar --mode strict -- npx -y @modelcontextprotocol/server-filesystem /tmp
kevlar --mode audit --log /tmp/kevlar.jsonl -- uvx mcp-server-git
Claude Desktop integration
In ~/Library/Application Support/Claude/claude_desktop_config.json (macOS):
{
"mcpServers" : {
"filesystem" : {
"command" : " kevlar " ,
"args" : [
" --mode " , " strict " ,
" --log " , " /tmp/kevlar-filesystem.jsonl " ,
" -- " ,
" npx " , " -y " , " @modelcontextprotocol/server-filesystem " ,
" /Users/you/Documents "
]
}
}
}
Modes
Mode
Reject threshold
Tool registry
list_changed
strict
score ≥ 40
Locked at handshake
Dropped
lenient
score ≥ 80
Dynamic updates allowed
Forwarded
audit
Never blocks
Dynamic
Forwarded
Use strict in production. Use audit to profile a new server before tightening the threshold. --log produces a JSONL file with every policy decision, indicator ID, score, and masked value snippet.
Architecture
Two files: policy.py is a pure CDR engine with no I/O, directly unit-testable. kevlar.py is the async stdio proxy. No runtime dependencies, stdlib only. kevlar.py spawns the server with start_new_session=True and kills the process group on shutdown, so npx / uvx child processes don't linger.
See docs/architecture.md for the full design: proxy lifecycle, CDR pipeline per message type, scoring arithmetic, indicator registries, and audit log format.
pip install -e " .[dev] "
pytest
149 tests. tests/test_proxy.py runs integration tests against real subprocess pairs. tests/test_policy.py covers the pure policy engine in unit tests.
Kevlar is a proxy firewall for Model Context Protocol
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Kevlar is a proxy firewall for Model Context Protocol - Ablation-Tool/kevlar

GitHub - Ablation-Tool/kevlar: Kevlar is a proxy firewall for Model Context Protocol · GitHub
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
Ablation-Tool
kevlar Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Kevlar is a proxy firewall for Model Context Protocol
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
docs docs tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md KEVLAR-01.jpeg KEVLAR-01.jpeg KEVLAR-repo-01.png KEVLAR-repo-01.png LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md kevlar.py kevlar.py policy.py policy.py pyproject.toml pyproject.toml pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
Kevlar is a proxy firewall for Model Context Protocol
LLM [ KEVLAR ] MCP
It sits between (you) and (everyone else)
MCP servers control their own tools.
Those tools (good or bad) go into your LLM blindly, due to trust.
Any MCP server can exploit this to redirect the model, pull credentials, or add any tools they want.
The MCP spec does not prevent any of this.
You do not control what an MCP server puts in its messages.
TL;DR
Kevlar sits between client and server and does four things:
Strips all tool descriptions before they reach the LLM. No exceptions.
Scans tool call responses for injection payloads and either redacts matched spans or blocks the response entirely.
Blocks outbound calls carrying credentials in their arguments, even if the LLM has already been compromised.
Locks the tool registry after the initial handshake so the server cannot inject new tools mid-session.
Kevlar strips every description field to "" unconditionally on every tools/list response, before the message reaches the client. The LLM sees name and inputSchema only. inputSchema is rebuilt from an allowlist: prose fields like description , title , examples , and $comment are stripped there too, and enum values, property keys, $ref targets, and type values are validated against the MCP spec. Score 0 on this pass cannot be achieved by adding description text back, because description text is never forwarded regardless of score.
Tool call results cannot be stripped without breaking functionality. Kevlar scans them against ~60 indicators sourced from published prompt injection research. Responses scoring above the mode threshold are quarantined (matched spans replaced with [KEVLAR:<ID>] tokens) or rejected (replaced with an isError block). Unicode bypass attempts using characters that normalize to injection text under NFKC are detected before normalization and redacted after.
Outbound tools/call arguments are scanned before the call is forwarded to the server. If the LLM has already been compromised and is trying to exfiltrate credentials, the call is blocked at the proxy layer. The server never receives it. Runs in all three modes.
After the initial tools/list handshake, the set of callable tool names is frozen. notifications/tools/list_changed from the server is dropped, not forwarded. Any call to a tool not in the CDR-sanitized set is blocked with a synthetic isError response regardless of call content.
Install
git clone https://github.com/Ablation-Tool/kevlar.git
cd kevlar
pip install .
Requires Python 3.10+. Linux and macOS only. Depends on os.killpg and asyncio.connect_read_pipe .
kevlar --mode strict -- npx -y @modelcontextprotocol/server-filesystem /tmp
kevlar --mode audit --log /tmp/kevlar.jsonl -- uvx mcp-server-git
Claude Desktop integration
In ~/Library/Application Support/Claude/claude_desktop_config.json (macOS):
{
"mcpServers" : {
"filesystem" : {
"command" : " kevlar " ,
"args" : [
" --mode " , " strict " ,
" --log " , " /tmp/kevlar-filesystem.jsonl " ,
" -- " ,
" npx " , " -y " , " @modelcontextprotocol/server-filesystem " ,
" /Users/you/Documents "
]
}
}
}
Modes
Mode
Reject threshold
Tool registry
list_changed
strict
score ≥ 40
Locked at handshake
Dropped
lenient
score ≥ 80
Dynamic updates allowed
Forwarded
audit
Never blocks
Dynamic
Forwarded
Use strict in production. Use audit to profile a new server before tightening the threshold. --log produces a JSONL file with every policy decision, indicator ID, score, and masked value snippet.
Architecture
Two files: policy.py is a pure CDR engine with no I/O, directly unit-testable. kevlar.py is the async stdio proxy. No runtime dependencies, stdlib only. kevlar.py spawns the server with start_new_session=True and kills the process group on shutdown, so npx / uvx child processes don't linger.
See docs/architecture.md for the full design: proxy lifecycle, CDR pipeline per message type, scoring arithmetic, indicator registries, and audit log format.
pip install -e " .[dev] "
pytest
149 tests. tests/test_proxy.py runs integration tests against real subprocess pairs. tests/test_policy.py covers the pure policy engine in unit tests.
Kevlar is a proxy firewall for Model Context Protocol
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
