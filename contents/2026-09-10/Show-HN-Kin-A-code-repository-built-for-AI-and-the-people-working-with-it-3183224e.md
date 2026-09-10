---
source: "https://github.com/firelock-ai/kin"
hn_url: "https://news.ycombinator.com/item?id=49636442"
title: "Show HN: Kin – A code repository built for AI and the people working with it"
article_title: "GitHub - firelock-ai/kin: The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. · GitHub"
image: "https://repository-images.githubusercontent.com/1178256335/67051cfd-f2f4-4e3e-b401-a4f4b9fd6d93"
author: "troyjr4103"
captured_at: "2026-09-10T00:58:45Z"
capture_tool: "hn-digest"
hn_id: 49636442
score: 1
comments: 0
posted_at: "2026-09-10T00:10:33Z"
tags:
  - hacker-news
---

# Show HN: Kin – A code repository built for AI and the people working with it

- HN: [49636442](https://news.ycombinator.com/item?id=49636442)
- Source: [github.com](https://github.com/firelock-ai/kin)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T00:10:33Z

## Translation

Title: Show HN: Kin – A code repository built for AI and the people working with it
Article title: GitHub - firelock-ai/kin: The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. · GitHub
Description: The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. - firelock-ai/kin

Article text:
GitHub - firelock-ai/kin: The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. · GitHub
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
firelock-ai
/
kin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,982 Commits 2,982 Commits Folders and files
.agents/ plugins .agents/ plugins .cargo .cargo .claude-plugin .claude-plugin .config .config .cursor-plugin .cursor-plugin .github .github assets/ mcp assets/ mcp brand brand completions completions crates crates docs docs fuzz fuzz packages packages plugins plugins scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore .mailmap .mailmap .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md cloudbuild.yaml cloudbuild.yaml deny.toml deny.toml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh gemini-extension.json gemini-extension.json llms-install.md llms-install.md llms.txt llms.txt lychee.toml lychee.toml marketplace.json marketplace.json package.json package.json pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml server.json server.json View all files Repository files navigation
AI changed who writes code.
Kin changes what they build on.
A graph-native code repository for people and AI agents.
AI writes a change in seconds. Working out what it touches has not gotten any
faster.
Every agent reads a repository the way a person would. Search, open files,
follow callers, build the picture, throw it away when the session ends. The
next one starts over. The reviewer starts over again. Most AI tools rebuild
context for each task. Kin keeps a durable semantic record across tasks,
agents, and changes.
Git shows which lines changed. Kin shows what the change affects.
Kin makes the graph the repository. Entities, relationships, exact source, and
change history are what you commit, branch, and merge, and files stay a
projection so ordinary tools keep working.
Agents stop rebuilding context and start editing code. Reviewers see what a
change touches before it merges.
Kin is a public alpha. It runs today as a local CLI, a daemon, an MCP server for
agents, a review surface, and a graph-backed filesystem projection. It is
pre-1.0, so expect rough edges and breaking changes. See the
latest stable release and
what is real today and what is alpha before
you put it in a critical workflow.
Point it at a repository you know and ask it something you already know the
answer to. Or watch it run on Kin's own repositories at
kinlab.ai/demo .
A one-line signature change in ripgrep looks harmless in the diff. Ask
kin impact about it, before any compiler runs, and it names what the edit
reaches. The callers of the changed signature come first, then everything
those callers pull in behind them.
Recorded against a prepared graph at ripgrep commit
e89fff89ac9af12e8d4ce9d5fd07beb408ca730f . A one-line signature edit, and Kin
surfaces the entities it affects before a compiler runs. The graph was built
beforehand. No compiler ran. The two commands are the ones in the quickstart
below: kin init . to build the graph, then kin impact on the entity you
changed.
The raw run directory for this capture is not public yet, so treat it as a
recipe you can re-run rather than a trace you can audit.
Kin surfaces what the change touches. Whether the change is correct stays with
your compiler, tests, and review. The graph is built beforehand by kin init ,
and building it is the expensive part; after that, impact questions are
answered from graph truth, not from re-reading the tree.
Install, admit your repository, check the graph, ask it something. Wire your
agent last. The agent tools answer from the graph, so a client pointed at a
repository with no graph gets a tool surface with nothing behind it.
Run the installer on its own and finish any setup prompts:
curl -fsSL https://get.kinlab.dev/install | sh
Once it finishes, reload your shell with this separate command:
exec " $SHELL " -l
At the new prompt, replace the path below with your repository's location and
run this block:
cd /path/to/your/repository &&
kin init . &&
kin overview
kin overview prints what the graph now holds: entity counts by kind, by
language, and the files carrying the most of them. If it prints counts, the
graph is real and the commands below have something to answer from.
The rest of this section is the same path with the detail behind each step.
On macOS or Linux, run the installer on its own and finish any setup prompts:
curl -fsSL https://get.kinlab.dev/install | sh
Once it finishes, reload your shell with this separate command:
exec " $SHELL " -l
The installer resolves the latest stable release ,
verifies its published SHA-256 checksum, installs the managed binaries under
~/.kin , and launches setup. Running the explicit agent intent, which
step 5 does once the graph exists, configures the built-in
MCP server for detected supported clients. Use --intent local for CLI and
filesystem use without MCP configuration, or --intent editor for the VS Code
path.
npm, Homebrew, and a manual archive resolve that same public release channel:
npm install -g @kinlab/kin@latest
brew install firelock-ai/kin/kin
Each archive and its .sha256 file is published under
https://github.com/firelock-ai/kin/releases/latest/download/ , and the release
page lists the asset names.
Confirm what you installed with kin --version , whichever path you took.
The quickstart doc carries the operator detail:
the asset matrix, what to do when a global npm install hits EACCES , how the
Homebrew formula is regenerated from each release rather than hand-maintained,
and kin setup uninstall when you want the integrations gone.
On Windows, run irm https://get.kinlab.dev/install.ps1 | iex in PowerShell.
Native Windows x86_64 support is early. Repository admission works: kin init imports a Git repository and publishes graph authority, and graph, lexical, and daemon-backed queries answer natively. Transparent filesystem projection is not shipped on Windows, and the end-to-end install proof does not yet cover MCP or review workflows there, so WSL2 remains the recommended path for the full Kin experience.
Read Platform and maturity below before choosing a
Windows install path.
2. Admit your repository as graph truth
Replace the path below with your repository's location:
cd /path/to/your/repository && kin init .
In a detected Git repository, kin init atomically admits complete reachable
history, refs, raw objects, the exact workspace tree, and admission policy into
repository-v6 graph authority. A worktree with uncommitted edits, staged
changes, or untracked files still admits: kin init admits the committed state
and discloses what it did not admit. It never substitutes an exact-HEAD snapshot or
raw-filesystem semantic rebuild. Supported repository-local remote URLs,
refspecs, branch tracking, and push defaults are sealed into Kin's Git
coexistence configuration; unsafe, ambiguous, or unsupported transfer settings
fail closed before publication.
Admission also derives the semantic entity and relation layer for every
supported entity-source file in that history, and kin init reports the durable,
generation-bound counts it committed. kin status reports that repository
authority view; kin graph status separately reports the daemon's mutable live
query graph, which may include later derived enrichment.
Query surfaces consume graph-owned enrichment when it exists and report its
absence instead of hiding the gap behind raw file search.
kin init is the slow step and the one that earns the rest. It admits your Git
history into the graph, and every answer after it comes from that graph rather
than from re-reading the tree. Measured on a fresh Debian 12 container with 4
CPUs and 8 GiB against the release npm serves today, the installer took 4
seconds, kin init took 139 seconds on a 503-file repository with 1,983
commits, and the first kin locate answered in 6.7 seconds while the daemon
cold-started, then in 71 milliseconds warm. Those are separately measured legs
of one sitting, not one timed run, and a repository with deeper history takes
longer.
"Supported entity-source file" means a file one of Kin's language adapters
claims. The adapter registry is the whole set, and every file in a repository
resolves through it:
A .h header is read as C++ when its contents say so, so a C++ project does not
lose namespaces and templates to the C grammar.
Everything else is admitted as content and stays queryable as history and text,
but is not parsed into entities and relations. That includes Markdown, HTML and
CSS, SQL, YAML, JSON and TOML, shell scripts, Objective-C, Scala, Elixir, Dart,
Lua, R, Zig, Haskell, and Nix. If your language is on that list, locate and
refs will not find symbols in it.
kin graph status
kin embed
kin graph status reports the daemon's live query graph and its coverage.
Admission derives the semantic entities, not their vectors, so run kin embed
to add local vector similarity over them and confirm coverage with
kin graph status again.
One thing to expect on a small repository: kin init starts a background
download of the roughly 523 MB embedding model, and a conversion that finishes
in seconds can beat it. When that happens the first kin locate ranks on
lexical and graph signals alone and says why on the line beneath its rows. If
that line reports the model still downloading, run the query again once it
lands. If it reports that none of it arrived, do not wait on it: kin embed
fetches the rest.
4. Ask it something you already know the answer to
This is the honest way to judge it. Pick a helper you know the callers of, or a
subsystem you could describe from memory, and see whether Kin agrees with you.
A question about code you've never read tells you nothing about whether the
answer is right.
kin locate " <something you already know is in this repository> "
kin refs ExactEntityName
kin trace ExactEntityName
kin impact ExactEntityName
Replace ExactEntityName with a symbol returned by locate . locate finds the
entities relevant to an intent, refs shows graph-owned callers/importers and
references, and trace returns the

[truncated]

## Original Extract

The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. - firelock-ai/kin

GitHub - firelock-ai/kin: The system of record for AI-written software. A persistent graph of entities, relationships, changes, and provenance, so humans and AI agents see what a change touches before it merges. Beside Git today. · GitHub
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
firelock-ai
/
kin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,982 Commits 2,982 Commits Folders and files
.agents/ plugins .agents/ plugins .cargo .cargo .claude-plugin .claude-plugin .config .config .cursor-plugin .cursor-plugin .github .github assets/ mcp assets/ mcp brand brand completions completions crates crates docs docs fuzz fuzz packages packages plugins plugins scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore .mailmap .mailmap .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md cloudbuild.yaml cloudbuild.yaml deny.toml deny.toml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh gemini-extension.json gemini-extension.json llms-install.md llms-install.md llms.txt llms.txt lychee.toml lychee.toml marketplace.json marketplace.json package.json package.json pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml server.json server.json View all files Repository files navigation
AI changed who writes code.
Kin changes what they build on.
A graph-native code repository for people and AI agents.
AI writes a change in seconds. Working out what it touches has not gotten any
faster.
Every agent reads a repository the way a person would. Search, open files,
follow callers, build the picture, throw it away when the session ends. The
next one starts over. The reviewer starts over again. Most AI tools rebuild
context for each task. Kin keeps a durable semantic record across tasks,
agents, and changes.
Git shows which lines changed. Kin shows what the change affects.
Kin makes the graph the repository. Entities, relationships, exact source, and
change history are what you commit, branch, and merge, and files stay a
projection so ordinary tools keep working.
Agents stop rebuilding context and start editing code. Reviewers see what a
change touches before it merges.
Kin is a public alpha. It runs today as a local CLI, a daemon, an MCP server for
agents, a review surface, and a graph-backed filesystem projection. It is
pre-1.0, so expect rough edges and breaking changes. See the
latest stable release and
what is real today and what is alpha before
you put it in a critical workflow.
Point it at a repository you know and ask it something you already know the
answer to. Or watch it run on Kin's own repositories at
kinlab.ai/demo .
A one-line signature change in ripgrep looks harmless in the diff. Ask
kin impact about it, before any compiler runs, and it names what the edit
reaches. The callers of the changed signature come first, then everything
those callers pull in behind them.
Recorded against a prepared graph at ripgrep commit
e89fff89ac9af12e8d4ce9d5fd07beb408ca730f . A one-line signature edit, and Kin
surfaces the entities it affects before a compiler runs. The graph was built
beforehand. No compiler ran. The two commands are the ones in the quickstart
below: kin init . to build the graph, then kin impact on the entity you
changed.
The raw run directory for this capture is not public yet, so treat it as a
recipe you can re-run rather than a trace you can audit.
Kin surfaces what the change touches. Whether the change is correct stays with
your compiler, tests, and review. The graph is built beforehand by kin init ,
and building it is the expensive part; after that, impact questions are
answered from graph truth, not from re-reading the tree.
Install, admit your repository, check the graph, ask it something. Wire your
agent last. The agent tools answer from the graph, so a client pointed at a
repository with no graph gets a tool surface with nothing behind it.
Run the installer on its own and finish any setup prompts:
curl -fsSL https://get.kinlab.dev/install | sh
Once it finishes, reload your shell with this separate command:
exec " $SHELL " -l
At the new prompt, replace the path below with your repository's location and
run this block:
cd /path/to/your/repository &&
kin init . &&
kin overview
kin overview prints what the graph now holds: entity counts by kind, by
language, and the files carrying the most of them. If it prints counts, the
graph is real and the commands below have something to answer from.
The rest of this section is the same path with the detail behind each step.
On macOS or Linux, run the installer on its own and finish any setup prompts:
curl -fsSL https://get.kinlab.dev/install | sh
Once it finishes, reload your shell with this separate command:
exec " $SHELL " -l
The installer resolves the latest stable release ,
verifies its published SHA-256 checksum, installs the managed binaries under
~/.kin , and launches setup. Running the explicit agent intent, which
step 5 does once the graph exists, configures the built-in
MCP server for detected supported clients. Use --intent local for CLI and
filesystem use without MCP configuration, or --intent editor for the VS Code
path.
npm, Homebrew, and a manual archive resolve that same public release channel:
npm install -g @kinlab/kin@latest
brew install firelock-ai/kin/kin
Each archive and its .sha256 file is published under
https://github.com/firelock-ai/kin/releases/latest/download/ , and the release
page lists the asset names.
Confirm what you installed with kin --version , whichever path you took.
The quickstart doc carries the operator detail:
the asset matrix, what to do when a global npm install hits EACCES , how the
Homebrew formula is regenerated from each release rather than hand-maintained,
and kin setup uninstall when you want the integrations gone.
On Windows, run irm https://get.kinlab.dev/install.ps1 | iex in PowerShell.
Native Windows x86_64 support is early. Repository admission works: kin init imports a Git repository and publishes graph authority, and graph, lexical, and daemon-backed queries answer natively. Transparent filesystem projection is not shipped on Windows, and the end-to-end install proof does not yet cover MCP or review workflows there, so WSL2 remains the recommended path for the full Kin experience.
Read Platform and maturity below before choosing a
Windows install path.
2. Admit your repository as graph truth
Replace the path below with your repository's location:
cd /path/to/your/repository && kin init .
In a detected Git repository, kin init atomically admits complete reachable
history, refs, raw objects, the exact workspace tree, and admission policy into
repository-v6 graph authority. A worktree with uncommitted edits, staged
changes, or untracked files still admits: kin init admits the committed state
and discloses what it did not admit. It never substitutes an exact-HEAD snapshot or
raw-filesystem semantic rebuild. Supported repository-local remote URLs,
refspecs, branch tracking, and push defaults are sealed into Kin's Git
coexistence configuration; unsafe, ambiguous, or unsupported transfer settings
fail closed before publication.
Admission also derives the semantic entity and relation layer for every
supported entity-source file in that history, and kin init reports the durable,
generation-bound counts it committed. kin status reports that repository
authority view; kin graph status separately reports the daemon's mutable live
query graph, which may include later derived enrichment.
Query surfaces consume graph-owned enrichment when it exists and report its
absence instead of hiding the gap behind raw file search.
kin init is the slow step and the one that earns the rest. It admits your Git
history into the graph, and every answer after it comes from that graph rather
than from re-reading the tree. Measured on a fresh Debian 12 container with 4
CPUs and 8 GiB against the release npm serves today, the installer took 4
seconds, kin init took 139 seconds on a 503-file repository with 1,983
commits, and the first kin locate answered in 6.7 seconds while the daemon
cold-started, then in 71 milliseconds warm. Those are separately measured legs
of one sitting, not one timed run, and a repository with deeper history takes
longer.
"Supported entity-source file" means a file one of Kin's language adapters
claims. The adapter registry is the whole set, and every file in a repository
resolves through it:
A .h header is read as C++ when its contents say so, so a C++ project does not
lose namespaces and templates to the C grammar.
Everything else is admitted as content and stays queryable as history and text,
but is not parsed into entities and relations. That includes Markdown, HTML and
CSS, SQL, YAML, JSON and TOML, shell scripts, Objective-C, Scala, Elixir, Dart,
Lua, R, Zig, Haskell, and Nix. If your language is on that list, locate and
refs will not find symbols in it.
kin graph status
kin embed
kin graph status reports the daemon's live query graph and its coverage.
Admission derives the semantic entities, not their vectors, so run kin embed
to add local vector similarity over them and confirm coverage with
kin graph status again.
One thing to expect on a small repository: kin init starts a background
download of the roughly 523 MB embedding model, and a conversion that finishes
in seconds can beat it. When that happens the first kin locate ranks on
lexical and graph signals alone and says why on the line beneath its rows. If
that line reports the model still downloading, run the query again once it
lands. If it reports that none of it arrived, do not wait on it: kin embed
fetches the rest.
4. Ask it something you already know the answer to
This is the honest way to judge it. Pick a helper you know the callers of, or a
subsystem you could describe from memory, and see whether Kin agrees with you.
A question about code you've never read tells you nothing about whether the
answer is right.
kin locate " <something you already know is in this repository> "
kin refs ExactEntityName
kin trace ExactEntityName
kin impact ExactEntityName
Replace ExactEntityName with a symbol returned by locate . locate finds the
entities relevant to an intent, refs shows graph-owned callers/importers and
references, and trace returns the

[truncated]
