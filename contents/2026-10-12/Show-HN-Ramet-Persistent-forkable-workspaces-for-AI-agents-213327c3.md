---
source: "https://github.com/shenli/ramet"
hn_url: "https://news.ycombinator.com/item?id=50061655"
title: "Show HN: Ramet – Persistent, forkable workspaces for AI agents"
article_title: "GitHub - shenli/ramet: Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. · GitHub"
image: "https://opengraph.githubassets.com/10c61e8b48ebb8fc3bc4c47869c1351428b2f10f57d60a2e8cb8513163763d14/shenli/ramet"
author: "shenli3514"
captured_at: "2026-10-12T20:08:29Z"
capture_tool: "hn-digest"
hn_id: 50061655
score: 2
comments: 0
posted_at: "2026-10-12T19:28:33Z"
tags:
  - hacker-news
---

# Show HN: Ramet – Persistent, forkable workspaces for AI agents

- HN: [50061655](https://news.ycombinator.com/item?id=50061655)
- Source: [github.com](https://github.com/shenli/ramet)
- Score: 2
- Comments: 0
- Posted: 2026-10-12T19:28:33Z

## Translation

Title: Show HN: Ramet – Persistent, forkable workspaces for AI agents
Article title: GitHub - shenli/ramet: Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. · GitHub
Description: Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. - shenli/ramet

Article text:
GitHub - shenli/ramet: Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. · GitHub
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
shenli
ramet Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github client client cmd/ ramet cmd/ ramet configs configs deploy/ minio deploy/ minio docs docs engine engine examples examples integration integration internal internal schemas schemas scripts scripts sdk/ python sdk/ python storage storage third_party third_party .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.integration.yaml compose.integration.yaml compose.remote.yaml compose.remote.yaml compose.yaml compose.yaml go.mod go.mod go.sum go.sum View all files Repository files navigation
Persistent, forkable workspaces for AI agents.
Save and restore agent work across sessions and sandbox replacements. Fork
isolated attempts, reset to a clean starting point, and preserve the results
you choose.
Ramet gives each session or attempt an isolated writable workspace from a shared
immutable base. Changes become persistent when your harness explicitly captures
a snapshot or promotes selected state.
Bring your own agent framework and sandbox. Your framework coordinates the
work, your sandbox runs it, and Ramet manages the workspace lifecycle.
Using LangChain / LangGraph, CrewAI, AutoGen, or SkyRL?
LangChain/LangGraph, CrewAI, and AutoGen require application adapter code;
dedicated Ramet plugins are not included. Your framework keeps conversation
state, orchestration, rewards, and training. Ramet persists filesystem state;
process memory and external services need their own recovery strategy.
Why workspace state needs different treatment
An agent workspace mixes files with very different lifetimes. A local commit
may be the only copy of useful work; a dependency directory can be rebuilt;
a scratch file may belong to just one attempt. Saving everything carries
reproducible and temporary data into every saved result. Throwing everything
away loses work you cannot reconstruct.
Ramet lets you express those differences when saving an attempt:
Your harness assigns path policies; Ramet does not infer which files are safe
to drop. These policies apply on promotion, with unclassified paths preserved
by default. Initial snapshots and checkpoints capture all supported files;
your runtime executes reconstruction recipes. See persistence policies .
Preview what will be saved before promotion, then restore through approved
reconstruction steps with input checks and an application readiness check.
For parallel agents, this means sharing a prepared base, isolating each
attempt's writes, and choosing both which attempt and which state to
keep. Disposable attempts need no saved snapshot.
Development alpha · Linux & macOS · Apache-2.0
Your harness calls Ramet through the Python SDK or Go client:
Prepare: snapshot a project or import a Git revision, then create a pool.
Acquire: get an isolated workspace and pass execution.path to your
runtime as the sandbox's workspace mount.
Run: use your existing framework to run the agent and evaluate its work.
Save or recycle: stop the sandbox's writers and unmount the workspace;
call promote() to save a selected result, then release the execution.
Release resets the workspace for reuse.
The path belongs to the Ramet worker's filesystem . Deploy the worker beside
the runtime that can mount it; a remote SDK connection does not transfer files
to a remote sandbox. Your runtime owns process isolation, resource limits, and
cleanup on lease loss.
How much integration work is involved?
Your agent loop and evaluator stay in your framework. The integration guide
shows the three call sites to change, an adapter sketch, and checks to validate
your environment. The Python SDK guide walks through
setup, Docker execution, persistence, and parallel attempts.
Execution writes stay local. Capture or promote explicitly to persist state.
Start with local storage; add PostgreSQL and S3-compatible storage when snapshots
need to survive worker-disk loss. Copy-on-write shares file data where supported,
with independent full copies as the fallback. See the
architecture and runtime security boundary .
For large snapshots where attempts touch only a subset of files, the optional
Linux FUSE backend hydrates working-tree files on first access.
The included Docker adapter uses copy/reflink; FUSE needs custom mount attachment.
This same-host setup introduces the API for local development and integration
testing. A deployed worker uses the same lifecycle beside your sandbox runtime.
Local storage keeps snapshots on that worker's disk; optional PostgreSQL and
S3-compatible storage let a replacement worker restore published snapshots.
See deployment and storage modes .
Requires Go 1.26.9+ , Python 3.10+ , Git, and Linux or macOS.
Install the CLI and Python SDK, then start a local worker:
git clone https://github.com/shenli/ramet.git
cd ramet
go build -o bin/ramet ./cmd/ramet
python3 -m venv .venv
.venv/bin/python -m pip install ./sdk/python
./bin/ramet serve --data /tmp/ramet-data --listen 127.0.0.1:7433
Leave the worker running. In another terminal, from the same checkout, run the
following with the absolute path to your project. Keep that project outside
/tmp/ramet-data and stop its writers while capturing the snapshot.
.venv/bin/python - << ' PY '
from ramet import Ramet
ramet = Ramet() # Connects to http://127.0.0.1:7433.
base = ramet.snapshot(source="/absolute/path/to/your/project")
with ramet.execution_pool(base, concurrency=4) as pool:
with pool.acquire() as execution:
# Run your tools against this directory; the source project is untouched.
print("Workspace:", execution.path)
(execution.path / "result.txt").write_text("selected result\n")
saved = execution.promote() # Persist this attempt as a new snapshot.
# Leaving the execution context resets and releases its workspace.
with ramet.execution_pool(saved, concurrency=1) as pool:
with pool.acquire() as restored:
print((restored.path / "result.txt").read_text())
PY
The restored workspace prints selected result . Snapshots persist in the
worker's data directory (use a persistent directory instead of /tmp to retain
them long term); pools and execution directories are temporary. The
SDK and worker run on the same host for this example. When connecting your
runtime, mount only execution.path and stop its writers before promotion or
release; the Docker adapter handles this
lifecycle for container commands.
For ready-made demonstrations, run ./bin/ramet demo or follow the
best-of-N coding demo .
Agent integration guide — discover capabilities, validate policies, and recover interrupted calls.
Persistence policies · Selective save and restore — inspect and reconstruct saved work.
Integration guide · Google AX assessment — scope your runtime integration.
Coding demo — run the workflow and integrate your agent.
Python SDK · Go integration — use Ramet from your harness.
Architecture · Design — components and correctness guarantees.
Operations · Security — deploy workers and connect a runtime.
Linux FUSE — configure and validate lazy workspace hydration.
Developer guide · Contributing — build, test, and extend Ramet.
Benchmarks · Validation — reproduce measurements and understand their limits.
Documentation index — API details, roadmap, and release guides.
Licensed under Apache-2.0 . See third-party notices .
Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. - shenli/ramet

GitHub - shenli/ramet: Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions. · GitHub
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
shenli
ramet Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github client client cmd/ ramet cmd/ ramet configs configs deploy/ minio deploy/ minio docs docs engine engine examples examples integration integration internal internal schemas schemas scripts scripts sdk/ python sdk/ python storage storage third_party third_party .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.integration.yaml compose.integration.yaml compose.remote.yaml compose.remote.yaml compose.yaml compose.yaml go.mod go.mod go.sum go.sum View all files Repository files navigation
Persistent, forkable workspaces for AI agents.
Save and restore agent work across sessions and sandbox replacements. Fork
isolated attempts, reset to a clean starting point, and preserve the results
you choose.
Ramet gives each session or attempt an isolated writable workspace from a shared
immutable base. Changes become persistent when your harness explicitly captures
a snapshot or promotes selected state.
Bring your own agent framework and sandbox. Your framework coordinates the
work, your sandbox runs it, and Ramet manages the workspace lifecycle.
Using LangChain / LangGraph, CrewAI, AutoGen, or SkyRL?
LangChain/LangGraph, CrewAI, and AutoGen require application adapter code;
dedicated Ramet plugins are not included. Your framework keeps conversation
state, orchestration, rewards, and training. Ramet persists filesystem state;
process memory and external services need their own recovery strategy.
Why workspace state needs different treatment
An agent workspace mixes files with very different lifetimes. A local commit
may be the only copy of useful work; a dependency directory can be rebuilt;
a scratch file may belong to just one attempt. Saving everything carries
reproducible and temporary data into every saved result. Throwing everything
away loses work you cannot reconstruct.
Ramet lets you express those differences when saving an attempt:
Your harness assigns path policies; Ramet does not infer which files are safe
to drop. These policies apply on promotion, with unclassified paths preserved
by default. Initial snapshots and checkpoints capture all supported files;
your runtime executes reconstruction recipes. See persistence policies .
Preview what will be saved before promotion, then restore through approved
reconstruction steps with input checks and an application readiness check.
For parallel agents, this means sharing a prepared base, isolating each
attempt's writes, and choosing both which attempt and which state to
keep. Disposable attempts need no saved snapshot.
Development alpha · Linux & macOS · Apache-2.0
Your harness calls Ramet through the Python SDK or Go client:
Prepare: snapshot a project or import a Git revision, then create a pool.
Acquire: get an isolated workspace and pass execution.path to your
runtime as the sandbox's workspace mount.
Run: use your existing framework to run the agent and evaluate its work.
Save or recycle: stop the sandbox's writers and unmount the workspace;
call promote() to save a selected result, then release the execution.
Release resets the workspace for reuse.
The path belongs to the Ramet worker's filesystem . Deploy the worker beside
the runtime that can mount it; a remote SDK connection does not transfer files
to a remote sandbox. Your runtime owns process isolation, resource limits, and
cleanup on lease loss.
How much integration work is involved?
Your agent loop and evaluator stay in your framework. The integration guide
shows the three call sites to change, an adapter sketch, and checks to validate
your environment. The Python SDK guide walks through
setup, Docker execution, persistence, and parallel attempts.
Execution writes stay local. Capture or promote explicitly to persist state.
Start with local storage; add PostgreSQL and S3-compatible storage when snapshots
need to survive worker-disk loss. Copy-on-write shares file data where supported,
with independent full copies as the fallback. See the
architecture and runtime security boundary .
For large snapshots where attempts touch only a subset of files, the optional
Linux FUSE backend hydrates working-tree files on first access.
The included Docker adapter uses copy/reflink; FUSE needs custom mount attachment.
This same-host setup introduces the API for local development and integration
testing. A deployed worker uses the same lifecycle beside your sandbox runtime.
Local storage keeps snapshots on that worker's disk; optional PostgreSQL and
S3-compatible storage let a replacement worker restore published snapshots.
See deployment and storage modes .
Requires Go 1.26.9+ , Python 3.10+ , Git, and Linux or macOS.
Install the CLI and Python SDK, then start a local worker:
git clone https://github.com/shenli/ramet.git
cd ramet
go build -o bin/ramet ./cmd/ramet
python3 -m venv .venv
.venv/bin/python -m pip install ./sdk/python
./bin/ramet serve --data /tmp/ramet-data --listen 127.0.0.1:7433
Leave the worker running. In another terminal, from the same checkout, run the
following with the absolute path to your project. Keep that project outside
/tmp/ramet-data and stop its writers while capturing the snapshot.
.venv/bin/python - << ' PY '
from ramet import Ramet
ramet = Ramet() # Connects to http://127.0.0.1:7433.
base = ramet.snapshot(source="/absolute/path/to/your/project")
with ramet.execution_pool(base, concurrency=4) as pool:
with pool.acquire() as execution:
# Run your tools against this directory; the source project is untouched.
print("Workspace:", execution.path)
(execution.path / "result.txt").write_text("selected result\n")
saved = execution.promote() # Persist this attempt as a new snapshot.
# Leaving the execution context resets and releases its workspace.
with ramet.execution_pool(saved, concurrency=1) as pool:
with pool.acquire() as restored:
print((restored.path / "result.txt").read_text())
PY
The restored workspace prints selected result . Snapshots persist in the
worker's data directory (use a persistent directory instead of /tmp to retain
them long term); pools and execution directories are temporary. The
SDK and worker run on the same host for this example. When connecting your
runtime, mount only execution.path and stop its writers before promotion or
release; the Docker adapter handles this
lifecycle for container commands.
For ready-made demonstrations, run ./bin/ramet demo or follow the
best-of-N coding demo .
Agent integration guide — discover capabilities, validate policies, and recover interrupted calls.
Persistence policies · Selective save and restore — inspect and reconstruct saved work.
Integration guide · Google AX assessment — scope your runtime integration.
Coding demo — run the workflow and integrate your agent.
Python SDK · Go integration — use Ramet from your harness.
Architecture · Design — components and correctness guarantees.
Operations · Security — deploy workers and connect a runtime.
Linux FUSE — configure and validate lazy workspace hydration.
Developer guide · Contributing — build, test, and extend Ramet.
Benchmarks · Validation — reproduce measurements and understand their limits.
Documentation index — API details, roadmap, and release guides.
Licensed under Apache-2.0 . See third-party notices .
Persistent, forkable workspaces for AI agents. Git-aware snapshots, verified resets, and control over what survives across sessions.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
