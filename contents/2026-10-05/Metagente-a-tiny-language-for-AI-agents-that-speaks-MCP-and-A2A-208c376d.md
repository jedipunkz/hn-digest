---
source: "https://github.com/cleuton/MetaAgent"
hn_url: "https://news.ycombinator.com/item?id=49967468"
title: "Metagente, a tiny language for AI agents that speaks MCP and A2A"
article_title: "GitHub - cleuton/MetaAgent: Ai Agent programming language · GitHub"
image: "https://opengraph.githubassets.com/a725d525f41a1cb0474dfc09e57f2ecf98032f1dca2a9cf63eca1afc10a7fff3/cleuton/MetaAgent"
author: "cleuton"
captured_at: "2026-10-05T17:59:37Z"
capture_tool: "hn-digest"
hn_id: 49967468
score: 1
comments: 0
posted_at: "2026-10-05T17:03:19Z"
tags:
  - hacker-news
---

# Metagente, a tiny language for AI agents that speaks MCP and A2A

- HN: [49967468](https://news.ycombinator.com/item?id=49967468)
- Source: [github.com](https://github.com/cleuton/MetaAgent)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T17:03:19Z

## Translation

Title: Metagente, a tiny language for AI agents that speaks MCP and A2A
Article title: GitHub - cleuton/MetaAgent: Ai Agent programming language · GitHub
Description: Ai Agent programming language. Contribute to cleuton/MetaAgent development by creating an account on GitHub.

Article text:
GitHub - cleuton/MetaAgent: Ai Agent programming language · GitHub
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
cleuton
/
MetaAgent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples multi-agent-samples multi-agent-samples scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md hello.ag hello.ag logo.jpg logo.jpg metagente.toml metagente.toml rustfmt.toml rustfmt.toml testapi.ag testapi.ag View all files Repository files navigation
Build AI agents in a few lines, not a few hundred.
Metagente is a small language for building AI agents. You describe what the agent is for, which tools it may
use and what it answers; the interpreter does the rest. It speaks MCP (Model Context Protocol) and A2A
(Agent to Agent) natively, so your agents can use any MCP tool and talk to other agents out of the box.
This demonstration is in multi-agent-samples
The interpreter is a single binary written in Rust. No Python environment, no framework to learn.
agent Weather
goal "Answer questions about the weather"
tool weather from mcp "npx -y weather-mcp"
accepts ask city
on ask
forecast = weather.forecast city: city
reply "In {city} it will be {forecast.summary}"
That is a complete agent: a goal, an MCP tool, an input and a reply.
Download the zip for your system from the latest release :
metagente-linux-amd64.zip , metagente-macos-arm64.zip or metagente-windows-amd64.zip .
Unzip it and open a terminal in the folder.
./bin/metagente run samples/clock.ag now
On Windows use .\bin\metagente.exe instead. Then create your own agent:
./bin/metagente new hello
./bin/metagente run hello.ag greet name=World
Prefer to build from source? You need Rust :
git clone https://github.com/cleuton/MetaAgent
cd MetaAgent
cargo build --release
target/release/metagente run examples/clock.ag now
Why Metagente?
Frameworks such as LangChain or CrewAI are powerful, but they assume you are a programmer working inside a
Python project. Metagente makes a different trade:
If you need fine control in Python, use a framework. If you want a working agent quickly, or you want to give
non-programmers a way to build agents, Metagente is for you.
Language and interpreter: goals, tools, inputs, replies, if , loops and results.
Built in tools and MCP tools: plug in any MCP server.
Safe agents: each agent declares which folders, environment variables and other agents it may use.
Dynamic link: agents call agents by name or path, with interface checks and cycle detection.
A2A and MCP server: expose an agent as an MCP server or with an A2A Agent Card, and send or receive A2A tasks.
CLI: new , check , run and serve .
City Briefing puts it all together: a Concierge agent asks a Researcher
agent over A2A, and the Researcher reads a web page through an MCP tool. Both use Claude. Bring your own API key.
Topic
Where
Tutorial, step by step
docs/tutorial.md
Programming guide (if, loops, results)
docs/guide.md
Language syntax
docs/syntax.md
Samples
samples/
Changelog
CHANGELOG.md
Roadmap
Stage
State
Core language, built in and MCP tools, CLI, tutorial
Built
Safe agents (per-agent permissions)
Built
Dynamic link between agents
Built
v1.0: MCP server, A2A Agent Card, A2A tasks, serve
Built, release pending
Project Barracuda: an agent server invoked via A2A or a frontend, with a queue for batch triggering
Next
Long running tasks, authentication for served agents, agent registry, scheduling, A2A streaming
Ideas
How it is built
Metagente is developed with spec-driven development using GitHub Spec Kit .
Every feature starts as a spec, so you can read why things are the way they are:
Issues, ideas and pull requests are welcome. If you build an agent with Metagente, open an issue and show it;
good ones become samples.
If Metagente is useful to you, a star on the repository helps other people find it.
v0.1.1. Semantic versioning. The version is kept in this file, in CHANGELOG.md and in Cargo.toml .
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Ai Agent programming language. Contribute to cleuton/MetaAgent development by creating an account on GitHub.

GitHub - cleuton/MetaAgent: Ai Agent programming language · GitHub
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
cleuton
/
MetaAgent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples multi-agent-samples multi-agent-samples scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md hello.ag hello.ag logo.jpg logo.jpg metagente.toml metagente.toml rustfmt.toml rustfmt.toml testapi.ag testapi.ag View all files Repository files navigation
Build AI agents in a few lines, not a few hundred.
Metagente is a small language for building AI agents. You describe what the agent is for, which tools it may
use and what it answers; the interpreter does the rest. It speaks MCP (Model Context Protocol) and A2A
(Agent to Agent) natively, so your agents can use any MCP tool and talk to other agents out of the box.
This demonstration is in multi-agent-samples
The interpreter is a single binary written in Rust. No Python environment, no framework to learn.
agent Weather
goal "Answer questions about the weather"
tool weather from mcp "npx -y weather-mcp"
accepts ask city
on ask
forecast = weather.forecast city: city
reply "In {city} it will be {forecast.summary}"
That is a complete agent: a goal, an MCP tool, an input and a reply.
Download the zip for your system from the latest release :
metagente-linux-amd64.zip , metagente-macos-arm64.zip or metagente-windows-amd64.zip .
Unzip it and open a terminal in the folder.
./bin/metagente run samples/clock.ag now
On Windows use .\bin\metagente.exe instead. Then create your own agent:
./bin/metagente new hello
./bin/metagente run hello.ag greet name=World
Prefer to build from source? You need Rust :
git clone https://github.com/cleuton/MetaAgent
cd MetaAgent
cargo build --release
target/release/metagente run examples/clock.ag now
Why Metagente?
Frameworks such as LangChain or CrewAI are powerful, but they assume you are a programmer working inside a
Python project. Metagente makes a different trade:
If you need fine control in Python, use a framework. If you want a working agent quickly, or you want to give
non-programmers a way to build agents, Metagente is for you.
Language and interpreter: goals, tools, inputs, replies, if , loops and results.
Built in tools and MCP tools: plug in any MCP server.
Safe agents: each agent declares which folders, environment variables and other agents it may use.
Dynamic link: agents call agents by name or path, with interface checks and cycle detection.
A2A and MCP server: expose an agent as an MCP server or with an A2A Agent Card, and send or receive A2A tasks.
CLI: new , check , run and serve .
City Briefing puts it all together: a Concierge agent asks a Researcher
agent over A2A, and the Researcher reads a web page through an MCP tool. Both use Claude. Bring your own API key.
Topic
Where
Tutorial, step by step
docs/tutorial.md
Programming guide (if, loops, results)
docs/guide.md
Language syntax
docs/syntax.md
Samples
samples/
Changelog
CHANGELOG.md
Roadmap
Stage
State
Core language, built in and MCP tools, CLI, tutorial
Built
Safe agents (per-agent permissions)
Built
Dynamic link between agents
Built
v1.0: MCP server, A2A Agent Card, A2A tasks, serve
Built, release pending
Project Barracuda: an agent server invoked via A2A or a frontend, with a queue for batch triggering
Next
Long running tasks, authentication for served agents, agent registry, scheduling, A2A streaming
Ideas
How it is built
Metagente is developed with spec-driven development using GitHub Spec Kit .
Every feature starts as a spec, so you can read why things are the way they are:
Issues, ideas and pull requests are welcome. If you build an agent with Metagente, open an issue and show it;
good ones become samples.
If Metagente is useful to you, a star on the repository helps other people find it.
v0.1.1. Semantic versioning. The version is kept in this file, in CHANGELOG.md and in Cargo.toml .
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
