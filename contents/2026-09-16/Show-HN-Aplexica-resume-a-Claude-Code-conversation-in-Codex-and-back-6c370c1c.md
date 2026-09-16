---
source: "https://github.com/Aplexica/Aplexica"
hn_url: "https://news.ycombinator.com/item?id=49727534"
title: "Show HN: Aplexica – resume a Claude Code conversation in Codex, and back"
article_title: "GitHub - Aplexica/Aplexica: Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents · GitHub"
image: "https://opengraph.githubassets.com/c15f96d20f9197695c3c4854972b836c83b0b030907af73715e824184ff26c4a/Aplexica/Aplexica"
author: "yruzin"
captured_at: "2026-09-16T14:39:14Z"
capture_tool: "hn-digest"
hn_id: 49727534
score: 1
comments: 0
posted_at: "2026-09-16T14:25:31Z"
tags:
  - hacker-news
---

# Show HN: Aplexica – resume a Claude Code conversation in Codex, and back

- HN: [49727534](https://news.ycombinator.com/item?id=49727534)
- Source: [github.com](https://github.com/Aplexica/Aplexica)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T14:25:31Z

## Translation

Title: Show HN: Aplexica – resume a Claude Code conversation in Codex, and back
Article title: GitHub - Aplexica/Aplexica: Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents · GitHub
Description: Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents - Aplexica/Aplexica

Article text:
GitHub - Aplexica/Aplexica: Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents · GitHub
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
Aplexica
/
Aplexica
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.github .github assets assets cmd cmd docker docker docs docs internal internal packaging packaging pkg/ adapterplugin pkg/ adapterplugin scripts scripts tools tools .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml .magiclint-allow .magiclint-allow CCLA.md CCLA.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-EXCEPTIONS.md LICENSE-EXCEPTIONS.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md aplexica-release.pub aplexica-release.pub go.mod go.mod go.sum go.sum View all files Repository files navigation
Use Claude Code and Codex together. Same memory, same skills, and resume either one's session in the other.
Aplexica is a local daemon that keeps your AI coding agents in sync. It watches where each agent stores its memory, skills, MCP config, and conversation history, translates them into one canonical event log, and writes them back out in every other agent's native format. Deterministic translation, not an LLM summary: the next agent doesn't read about your last session, it continues it.
Works with Claude Code, Codex, Hermes, OpenClaw, and Kilo Code. Runs entirely on your machine. No account, no telemetry, no network required.
You need two supported agents installed. The example uses Codex and Claude Code.
brew install Aplexica/tap/aplexica
aplexica setup --yes --install
Then turn on sync. Aplexica imports from your agents right away but writes nothing to any of them until you say so:
Open the local web UI with aplexica web open (the tray icon opens it too).
Go to Routing rules , choose Add from preset , and pick Sync everything everywhere .
aplexica sync enable --all
aplexica daemon reload
Now ask Codex anything:
What is the largest planet in our solar system?
Open a second terminal, run claude , type /resume , and pick the conversation named Codex: what is the largest planet in our solar system? . Claude Code continues the conversation from its own native history. Ask a follow-up there, and it appears in Codex. For Hermes, close and reopen it before the refreshed conversation list appears.
Other install channels (Debian/Ubuntu .deb , Windows .zip , verified archives, build from source) are under Install .
One memory for every agent. Write a memory in Claude Code, and Codex, Hermes, OpenClaw, and Kilo Code have it within seconds. Skills and MCP tool config sync the same way.
Resume any conversation in any agent. Session history is translated into each agent's own session format, so /resume just works.
Fork a conversation across agents. Branch a session into a second agent to try a different direction, keep both, compare.
Choose what syncs where. Tags and TOML routing rules send work memories to work agents and keep private ones where they started.
Back up and restore. Snapshot everything your agents know, restore after a lost laptop or a reinstall.
Not a Cursor, Gemini CLI, Copilot, or Windsurf integration yet. Five agents are supported today. Adding one is a ~80-line adapter; see docs/adapters/ . Open an issue for the one you want.
Not a hosted service. The daemon is complete on its own. An optional, end-to-end encrypted relay for syncing between your own machines is at aplexica.com .
Not a summarizer. Nothing is paraphrased by a model. State is replicated exactly.
Picking an AI coding agent is a one-way door. After a few weeks the memories, skills, MCP servers, and conversations you've built up are what make the agent useful, and they only exist in that agent's format. Switching means starting over. Running two side by side means maintaining two copies by hand. Trying the agent that shipped yesterday means abandoning everything.
Aplexica turns that state into something you own and that follows you to every agent you use.
Uninstall test. A portability tool should survive its own removal. See the Uninstall Test .
Aplexica installs three executables together:
aplexica — CLI, daemon, local web UI server, and setup wizard.
aplexica-status — status helper the tray spawns, so process monitors can
tell the watcher apart from the daemon.
aplexicatray — menu bar / system-tray companion.
The local web UI is compiled into aplexica itself in release builds; there is
no separate UI package to install.
macOS / Linux — Homebrew: Homebrew is a live channel at 1.0.74.
Full steps: brew.md .
brew install Aplexica/tap/aplexica
aplexica setup --yes --install
Debian / Ubuntu — .deb : follow the exact-version download and
verification steps in apt.md , then:
aplexica setup --yes --install
Windows (release archive): download aplexica-1.0.74-windows-amd64.zip or
aplexica-1.0.74-windows-arm64.zip from the
v1.0.74 release ,
verify it following verify.md , unzip it into a folder
you own, and run aplexica.exe or aplexicatray.exe from that folder. Full
steps, including start-at-logon: windows.md . Then:
aplexica setup --yes --install
That one setup command registers and starts the daemon, installs and launches
the tray when it is enabled (the default on every supported desktop OS), and
brings up the local web UI. Adapters are built in and auto-discover your
agents.
Aplexica release authority is a non-exportable AWS KMS key—not GitHub, a CDN, a package registry, or a maintainer workstation. Each release ships a KMS-backed cosign signature over SHA256SUMS and a KMS-backed, public-policy-checked SLSA v1 provenance bundle. GitHub Actions receives only a short-lived AWS session for the isolated signing job; the publication job has no signing authority.
Verify SHA256SUMS with the independently distributed aplexica-release.pub , check the selected artifact digest, and verify its provenance and semantic policy. GitHub Releases, a NAS, USB, or a local HTTP directory may all transport the same authenticated bytes. The exact commands are in Verify a release . Never pipe downloaded text to a shell or PowerShell expression.
Development builds use the public Go source directly:
git clone https://github.com/Aplexica/Aplexica.git
cd Aplexica
make build
make tray
The local web UI is replaced by an explanatory placeholder unless
internal/web/embed/dist-local/ contains a Portal bundle. Release-style builds
add -tags release , which makes the daemon fail closed when that bundle is
missing; development builds can omit the tag. See the
source-build guide for tests, cross-compilation, and
local installation.
# 1. Install Aplexica using one of the available methods in docs/install/.
# 2. Register and start the daemon, install tray autostart, launch the tray, and
# bring up the local web UI. The tray is enabled by default on every OS.
aplexica setup --yes --install
aplexica status
# (Run `aplexica setup` with no flags for an interactive walkthrough.)
# 3. Watch the canonical store
aplexica list # nothing yet — agent state hasn't been imported
# 4. Add at least one matching routing rule first; docs/quickstart.md gives
# the exact TOML and `aplexica rules add` command. Then enable receivers.
aplexica sync enable --all # or name specific agents
aplexica daemon reload
# 5. Let your agents create artifacts naturally. Once Claude Code,
# Codex, or any other supported agent writes its native file, the
# daemon imports it within ~2 seconds and fans it out to every
# enabled agent's native location.
# 6. Anywhere along the way:
aplexica list # see every artifact across kinds
aplexica show < artifact-id > # inspect one
aplexica conflicts list # any divergent writes that need your decision
aplexica branch list < artifact-id > # branch topology if you've forked
A guided walkthrough from install to working two-agent sync lives at docs/quickstart.md .
Five V1 agents, four artifact kinds, full bidirectional translation:
Adding a new agent is one ~80-line adapter package. See docs/adapters/ for per-agent specs and internal/adapter/openclaw/ for a reference implementation.
Each supported agent has an adapter that translates between its native format and the Aplexica Canonical Format (ACF) — an append-only event log per artifact, hash-chained for integrity, plus a typed payload. State lives in the user-owned canonical store at ~/.aplexica/store/ . The daemon, run through aplexica daemon , watches each agent's native location and runs the import → fan-out cycle on every change. Secrets live in a separate protected store and never enter the canonical event log. The whole stack is deterministic and lossless — no LLM summarization, no consolidation, no briefing layer.
The public BRDs in docs/ describe the canonical format and local
product requirements.
Selective sync, forking, branching
Power-user workflows are first-class:
Branches. aplexica fork <artifact> --from <event> --to-agent codex creates a divergent branch; the source agent's view is unchanged. aplexica log --graph <artifact> renders the topology. aplexica merge <artifact> --from <branch> --into <branch> reconciles.
Selective sync. aplexica tag add <artifact> work then a TOML rule like route.agents = ["claude-code", "codex"] routes that artifact only to your work agents. aplexica rules test <artifact> explains every routing decision.
Stays-local content. Tags are metadata, not policy, and no routing rules
are active by default. To keep a private memory on its originating agent,
add an explicit rule matching that tag with
route.agents = ["__originatingAgent__"] and route.remote = "exclude" ;
the quickstart gives the complete rule
and safe replacement sequence.
Full reference: docs/04-brd-forking-and-merging.md , docs/05-brd-selective-sync-and-routing.md .
aplexicatray is a cross-platform menubar / system-tray / SNI indicator that reflects daemon state at a glance:
Right-click for Pause / Resume , Show Conflicts , Open Logs , Pending Projects , and Open Aplexica . Packaged installs include it, and aplexica daemon install wires it to start with Aplexica on every desktop OS unless you explicitly opt out.
Path
Purpose
docs/00-vision.md
North-star vision and scope
docs/user-guide.md
Complete CLI, tray, and local web UI user guide
docs/quickstart.md
10-minute install + sync walkthrough
docs/install/
Per-platform / per-channel installation
docs/01-brd-backup-restore.md
Local export, import, convert
docs/02-brd-format-adapters.

[truncated]

## Original Extract

Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents - Aplexica/Aplexica

GitHub - Aplexica/Aplexica: Aplexica - end AI agent lock-in. Start conversation with one agent - continue with another in real time. Cross-agent state portability for AI agents · GitHub
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
Aplexica
/
Aplexica
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.github .github assets assets cmd cmd docker docker docs docs internal internal packaging packaging pkg/ adapterplugin pkg/ adapterplugin scripts scripts tools tools .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml .magiclint-allow .magiclint-allow CCLA.md CCLA.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-EXCEPTIONS.md LICENSE-EXCEPTIONS.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md aplexica-release.pub aplexica-release.pub go.mod go.mod go.sum go.sum View all files Repository files navigation
Use Claude Code and Codex together. Same memory, same skills, and resume either one's session in the other.
Aplexica is a local daemon that keeps your AI coding agents in sync. It watches where each agent stores its memory, skills, MCP config, and conversation history, translates them into one canonical event log, and writes them back out in every other agent's native format. Deterministic translation, not an LLM summary: the next agent doesn't read about your last session, it continues it.
Works with Claude Code, Codex, Hermes, OpenClaw, and Kilo Code. Runs entirely on your machine. No account, no telemetry, no network required.
You need two supported agents installed. The example uses Codex and Claude Code.
brew install Aplexica/tap/aplexica
aplexica setup --yes --install
Then turn on sync. Aplexica imports from your agents right away but writes nothing to any of them until you say so:
Open the local web UI with aplexica web open (the tray icon opens it too).
Go to Routing rules , choose Add from preset , and pick Sync everything everywhere .
aplexica sync enable --all
aplexica daemon reload
Now ask Codex anything:
What is the largest planet in our solar system?
Open a second terminal, run claude , type /resume , and pick the conversation named Codex: what is the largest planet in our solar system? . Claude Code continues the conversation from its own native history. Ask a follow-up there, and it appears in Codex. For Hermes, close and reopen it before the refreshed conversation list appears.
Other install channels (Debian/Ubuntu .deb , Windows .zip , verified archives, build from source) are under Install .
One memory for every agent. Write a memory in Claude Code, and Codex, Hermes, OpenClaw, and Kilo Code have it within seconds. Skills and MCP tool config sync the same way.
Resume any conversation in any agent. Session history is translated into each agent's own session format, so /resume just works.
Fork a conversation across agents. Branch a session into a second agent to try a different direction, keep both, compare.
Choose what syncs where. Tags and TOML routing rules send work memories to work agents and keep private ones where they started.
Back up and restore. Snapshot everything your agents know, restore after a lost laptop or a reinstall.
Not a Cursor, Gemini CLI, Copilot, or Windsurf integration yet. Five agents are supported today. Adding one is a ~80-line adapter; see docs/adapters/ . Open an issue for the one you want.
Not a hosted service. The daemon is complete on its own. An optional, end-to-end encrypted relay for syncing between your own machines is at aplexica.com .
Not a summarizer. Nothing is paraphrased by a model. State is replicated exactly.
Picking an AI coding agent is a one-way door. After a few weeks the memories, skills, MCP servers, and conversations you've built up are what make the agent useful, and they only exist in that agent's format. Switching means starting over. Running two side by side means maintaining two copies by hand. Trying the agent that shipped yesterday means abandoning everything.
Aplexica turns that state into something you own and that follows you to every agent you use.
Uninstall test. A portability tool should survive its own removal. See the Uninstall Test .
Aplexica installs three executables together:
aplexica — CLI, daemon, local web UI server, and setup wizard.
aplexica-status — status helper the tray spawns, so process monitors can
tell the watcher apart from the daemon.
aplexicatray — menu bar / system-tray companion.
The local web UI is compiled into aplexica itself in release builds; there is
no separate UI package to install.
macOS / Linux — Homebrew: Homebrew is a live channel at 1.0.74.
Full steps: brew.md .
brew install Aplexica/tap/aplexica
aplexica setup --yes --install
Debian / Ubuntu — .deb : follow the exact-version download and
verification steps in apt.md , then:
aplexica setup --yes --install
Windows (release archive): download aplexica-1.0.74-windows-amd64.zip or
aplexica-1.0.74-windows-arm64.zip from the
v1.0.74 release ,
verify it following verify.md , unzip it into a folder
you own, and run aplexica.exe or aplexicatray.exe from that folder. Full
steps, including start-at-logon: windows.md . Then:
aplexica setup --yes --install
That one setup command registers and starts the daemon, installs and launches
the tray when it is enabled (the default on every supported desktop OS), and
brings up the local web UI. Adapters are built in and auto-discover your
agents.
Aplexica release authority is a non-exportable AWS KMS key—not GitHub, a CDN, a package registry, or a maintainer workstation. Each release ships a KMS-backed cosign signature over SHA256SUMS and a KMS-backed, public-policy-checked SLSA v1 provenance bundle. GitHub Actions receives only a short-lived AWS session for the isolated signing job; the publication job has no signing authority.
Verify SHA256SUMS with the independently distributed aplexica-release.pub , check the selected artifact digest, and verify its provenance and semantic policy. GitHub Releases, a NAS, USB, or a local HTTP directory may all transport the same authenticated bytes. The exact commands are in Verify a release . Never pipe downloaded text to a shell or PowerShell expression.
Development builds use the public Go source directly:
git clone https://github.com/Aplexica/Aplexica.git
cd Aplexica
make build
make tray
The local web UI is replaced by an explanatory placeholder unless
internal/web/embed/dist-local/ contains a Portal bundle. Release-style builds
add -tags release , which makes the daemon fail closed when that bundle is
missing; development builds can omit the tag. See the
source-build guide for tests, cross-compilation, and
local installation.
# 1. Install Aplexica using one of the available methods in docs/install/.
# 2. Register and start the daemon, install tray autostart, launch the tray, and
# bring up the local web UI. The tray is enabled by default on every OS.
aplexica setup --yes --install
aplexica status
# (Run `aplexica setup` with no flags for an interactive walkthrough.)
# 3. Watch the canonical store
aplexica list # nothing yet — agent state hasn't been imported
# 4. Add at least one matching routing rule first; docs/quickstart.md gives
# the exact TOML and `aplexica rules add` command. Then enable receivers.
aplexica sync enable --all # or name specific agents
aplexica daemon reload
# 5. Let your agents create artifacts naturally. Once Claude Code,
# Codex, or any other supported agent writes its native file, the
# daemon imports it within ~2 seconds and fans it out to every
# enabled agent's native location.
# 6. Anywhere along the way:
aplexica list # see every artifact across kinds
aplexica show < artifact-id > # inspect one
aplexica conflicts list # any divergent writes that need your decision
aplexica branch list < artifact-id > # branch topology if you've forked
A guided walkthrough from install to working two-agent sync lives at docs/quickstart.md .
Five V1 agents, four artifact kinds, full bidirectional translation:
Adding a new agent is one ~80-line adapter package. See docs/adapters/ for per-agent specs and internal/adapter/openclaw/ for a reference implementation.
Each supported agent has an adapter that translates between its native format and the Aplexica Canonical Format (ACF) — an append-only event log per artifact, hash-chained for integrity, plus a typed payload. State lives in the user-owned canonical store at ~/.aplexica/store/ . The daemon, run through aplexica daemon , watches each agent's native location and runs the import → fan-out cycle on every change. Secrets live in a separate protected store and never enter the canonical event log. The whole stack is deterministic and lossless — no LLM summarization, no consolidation, no briefing layer.
The public BRDs in docs/ describe the canonical format and local
product requirements.
Selective sync, forking, branching
Power-user workflows are first-class:
Branches. aplexica fork <artifact> --from <event> --to-agent codex creates a divergent branch; the source agent's view is unchanged. aplexica log --graph <artifact> renders the topology. aplexica merge <artifact> --from <branch> --into <branch> reconciles.
Selective sync. aplexica tag add <artifact> work then a TOML rule like route.agents = ["claude-code", "codex"] routes that artifact only to your work agents. aplexica rules test <artifact> explains every routing decision.
Stays-local content. Tags are metadata, not policy, and no routing rules
are active by default. To keep a private memory on its originating agent,
add an explicit rule matching that tag with
route.agents = ["__originatingAgent__"] and route.remote = "exclude" ;
the quickstart gives the complete rule
and safe replacement sequence.
Full reference: docs/04-brd-forking-and-merging.md , docs/05-brd-selective-sync-and-routing.md .
aplexicatray is a cross-platform menubar / system-tray / SNI indicator that reflects daemon state at a glance:
Right-click for Pause / Resume , Show Conflicts , Open Logs , Pending Projects , and Open Aplexica . Packaged installs include it, and aplexica daemon install wires it to start with Aplexica on every desktop OS unless you explicitly opt out.
Path
Purpose
docs/00-vision.md
North-star vision and scope
docs/user-guide.md
Complete CLI, tray, and local web UI user guide
docs/quickstart.md
10-minute install + sync walkthrough
docs/install/
Per-platform / per-channel installation
docs/01-brd-backup-restore.md
Local export, import, convert
docs/02-brd-format-adapters.

[truncated]
