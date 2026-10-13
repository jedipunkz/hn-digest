---
source: "https://github.com/tinyhumansai/openscion"
hn_url: "https://news.ycombinator.com/item?id=50079329"
title: "Open source Scion – Anthropic's Git alternative in open source"
article_title: "GitHub - tinyhumansai/openscion: OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion · GitHub"
image: "https://repository-images.githubusercontent.com/1417985199/ac8229fb-6b95-45aa-b9f5-6cdd6f45b090"
author: "gafoor"
captured_at: "2026-10-13T22:32:52Z"
capture_tool: "hn-digest"
hn_id: 50079329
score: 3
comments: 0
posted_at: "2026-10-13T21:46:38Z"
tags:
  - hacker-news
---

# Open source Scion – Anthropic's Git alternative in open source

- HN: [50079329](https://news.ycombinator.com/item?id=50079329)
- Source: [github.com](https://github.com/tinyhumansai/openscion)
- Score: 3
- Comments: 0
- Posted: 2026-10-13T21:46:38Z

## Translation

Title: Open source Scion – Anthropic's Git alternative in open source
Article title: GitHub - tinyhumansai/openscion: OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion · GitHub
Description: OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion - tinyhumansai/openscion

Article text:
GitHub - tinyhumansai/openscion: OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion · GitHub
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
tinyhumansai
openscion Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion
www.anthropic.com/engineering/redesigning-source-control Resources
Readme GPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
118 Commits 118 Commits Folders and files
.claude .claude .github .github apps/ web apps/ web benchmarks benchmarks crates crates docs docs worktrees worktrees .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md clippy.toml clippy.toml deny.toml deny.toml View all files Repository files navigation
OpenScion is an open-source, server-centric source control system for very
large monorepos where most clients are AI agents and CI jobs rather than people
on laptops. It follows the design Anthropic described for its internal system,
Scion, in
Redesigning source control .
opnscn against git on synthetic repositories of 4 KiB files, measured with
openscion-bench ( cargo run -p openscion-benchmarks --release ):
OpenScion imports and clones several times faster than git at both sizes. The
one-directory clone shows the design at work: opnscn fetches only the trees
on the path to the directory and the files beneath it, so its cost tracks the
slice you ask for rather than the size of the repository. Setup, methodology,
and raw numbers are in benchmarks/README.md .
Git assumes every client is offline and keeps a full copy of history. An
agent-heavy monorepo breaks that assumption. Thousands of short-lived, always-online
jobs each want a small slice of a huge tree, and each one pays for the whole
repository anyway. Anthropic reports going from about 2,500 to over 30,000
merged changes a day in four months. At that point the main-branch history was
almost 100 GB, CI ran about 25 million checkouts a day, and they needed several
hundred git mirrors to keep up.
OpenScion takes the opposite approach:
The server does the work. History, blame, commits, rebases, and merges
run server-side. A client keeps a bounded cache and lays out files.
Changesets, not snapshots. A changeset lists only the files that changed
relative to its first parent. Creating one costs the size of the change, not
the size of the repository.
Content addressing everywhere. Blobs, trees, and changesets are named by
their BLAKE3 hash. Any cache at any tier can serve them without coordination,
and any client can verify them.
Few refs, moved by compare-and-swap. Creating a changeset never touches a
ref. Landing is a server-side rebase plus a single-record swap, retried on
contention, with no global merge lock.
Source control as an API. Log, diff, blame, ancestry, rebase, revert, and
land are RPCs, so CI and bisect pipelines need no checkout.
Fetch only what you name. Clients request specific trees and blobs
instead of negotiating packfiles, so sparse checkouts need no special support.
OpenScion is purpose-built for one workload: large monorepos where most of the
traffic comes from agents and CI jobs. Server-side operations, fetching only
what is named, and lock-free landing are what keep thousands of concurrent
clients fast at that scale.
opnscn (client) ──► openscion-daemon ──► openscion-cache ──► openscion-server
working copy one per node regional, log, refs,
fs watcher node-wide objects immutable objects derived data
│
metadata store ─┴─ object store
All tiers answer the same object RPCs, so a client talks to whichever one is
nearest. Mutable state (the log, refs, and derived indexes) lives only on the
server. Storage sits behind two ports, ObjectStore and MetadataStore . The
reference implementations are in-memory and local (filesystem, SQLite), so a
single node works without any cloud services.
Document
What it covers
Architecture
goals, principles, components, tiers, delivery phases
Data model
object ids, paths, blobs, trees, changesets, the log, refs, derived data
API
operations, fetch protocol, rebase and landing, access control, errors
Background: Scion
notes on the system OpenScion is modelled on
Background: git, GitHub, GitLab
how git stores and moves data, and how the two largest forges host it
Roadmap
what exists, what is next, what is out of scope
AGENTS.md
repository guidelines for humans and coding agents
Layout
crates/
├── openscion-proto/ # the protocol: payload types and the protocol version
├── openscion-core/ # the engine: data model, stores, derived data, rebase
├── openscion-server/ # openscion-server: the HTTP API over the engine
└── openscion/ # opnscn, the command-line client (crates.io: openscion)
benchmarks/ # openscion-bench: OpenScion vs git on synthetic repos
apps/web/ # the web GUI (Next.js + shadcn/ui)
docs/
├── research/ # source material: Scion, git, GitHub, GitLab
├── specs/ # what to build and why
├── plans/ # test-first implementation plans
└── adr/ # architecture decision records
openscion-proto holds every type that crosses a process boundary. It has no
transport, runtime, or behavior, so any tool can depend on it to speak the API.
openscion-core depends on it and re-exports it, so both crates name the same
types. The server, the CLI, and the GUI all speak that one API.
cargo run -p openscion-server # serves http://127.0.0.1:7878
cargo run -p openscion -- import ./tree # opnscn: record a tree as `main`
cargo run -p openscion -- commit -m " edit " --dir ./tree
cargo run -p openscion -- clone ./copy # or `--path dir` for one directory
cd apps/web && pnpm install && pnpm dev # GUI on http://localhost:3000
cargo run -p openscion-benchmarks --release # time the scenarios against git
Install the client from crates.io with cargo install openscion , which
provides the opnscn binary.
cargo fmt --all -- --check
cargo clippy --all-targets --all-features -- -D warnings
cargo build --all-targets --all-features
cargo test --all-features
CI runs those four checks. It also requires 90% line coverage in every source
file, builds rustdoc with -D warnings , builds with the MSRV, runs a
cargo-deny supply-chain check, and asserts that openscion-proto stays free
of transports and runtimes. To run the extra checks locally:
cargo doc --no-deps --all-features # CI builds this with RUSTDOCFLAGS="-D warnings"
cargo deny check all
cargo install cargo-llvm-cov # once
.github/scripts/check-file-coverage.sh 90 coverage.json
Releasing
Run the Release workflow from the Actions tab and pick a patch , minor ,
or major bump. The workflow re-validates the workspace, bumps the shared
[workspace.package] version, tags it, publishes openscion-proto ,
openscion-core , and openscion to crates.io (using the
CARGO_REGISTRY_TOKEN secret), and creates a GitHub release. Do not edit the
version in the root Cargo.toml by hand.
See CONTRIBUTING.md . Behavior changes start in
docs/specs/ ; each accepted spec gets a test-first
plan in docs/plans/ before any code is written.
OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion
www.anthropic.com/engineering/redesigning-source-control Resources
Readme GPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion - tinyhumansai/openscion

GitHub - tinyhumansai/openscion: OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion · GitHub
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
tinyhumansai
openscion Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion
www.anthropic.com/engineering/redesigning-source-control Resources
Readme GPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
118 Commits 118 Commits Folders and files
.claude .claude .github .github apps/ web apps/ web benchmarks benchmarks crates crates docs docs worktrees worktrees .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md clippy.toml clippy.toml deny.toml deny.toml View all files Repository files navigation
OpenScion is an open-source, server-centric source control system for very
large monorepos where most clients are AI agents and CI jobs rather than people
on laptops. It follows the design Anthropic described for its internal system,
Scion, in
Redesigning source control .
opnscn against git on synthetic repositories of 4 KiB files, measured with
openscion-bench ( cargo run -p openscion-benchmarks --release ):
OpenScion imports and clones several times faster than git at both sizes. The
one-directory clone shows the design at work: opnscn fetches only the trees
on the path to the directory and the files beneath it, so its cost tracks the
slice you ask for rather than the size of the repository. Setup, methodology,
and raw numbers are in benchmarks/README.md .
Git assumes every client is offline and keeps a full copy of history. An
agent-heavy monorepo breaks that assumption. Thousands of short-lived, always-online
jobs each want a small slice of a huge tree, and each one pays for the whole
repository anyway. Anthropic reports going from about 2,500 to over 30,000
merged changes a day in four months. At that point the main-branch history was
almost 100 GB, CI ran about 25 million checkouts a day, and they needed several
hundred git mirrors to keep up.
OpenScion takes the opposite approach:
The server does the work. History, blame, commits, rebases, and merges
run server-side. A client keeps a bounded cache and lays out files.
Changesets, not snapshots. A changeset lists only the files that changed
relative to its first parent. Creating one costs the size of the change, not
the size of the repository.
Content addressing everywhere. Blobs, trees, and changesets are named by
their BLAKE3 hash. Any cache at any tier can serve them without coordination,
and any client can verify them.
Few refs, moved by compare-and-swap. Creating a changeset never touches a
ref. Landing is a server-side rebase plus a single-record swap, retried on
contention, with no global merge lock.
Source control as an API. Log, diff, blame, ancestry, rebase, revert, and
land are RPCs, so CI and bisect pipelines need no checkout.
Fetch only what you name. Clients request specific trees and blobs
instead of negotiating packfiles, so sparse checkouts need no special support.
OpenScion is purpose-built for one workload: large monorepos where most of the
traffic comes from agents and CI jobs. Server-side operations, fetching only
what is named, and lock-free landing are what keep thousands of concurrent
clients fast at that scale.
opnscn (client) ──► openscion-daemon ──► openscion-cache ──► openscion-server
working copy one per node regional, log, refs,
fs watcher node-wide objects immutable objects derived data
│
metadata store ─┴─ object store
All tiers answer the same object RPCs, so a client talks to whichever one is
nearest. Mutable state (the log, refs, and derived indexes) lives only on the
server. Storage sits behind two ports, ObjectStore and MetadataStore . The
reference implementations are in-memory and local (filesystem, SQLite), so a
single node works without any cloud services.
Document
What it covers
Architecture
goals, principles, components, tiers, delivery phases
Data model
object ids, paths, blobs, trees, changesets, the log, refs, derived data
API
operations, fetch protocol, rebase and landing, access control, errors
Background: Scion
notes on the system OpenScion is modelled on
Background: git, GitHub, GitLab
how git stores and moves data, and how the two largest forges host it
Roadmap
what exists, what is next, what is out of scope
AGENTS.md
repository guidelines for humans and coding agents
Layout
crates/
├── openscion-proto/ # the protocol: payload types and the protocol version
├── openscion-core/ # the engine: data model, stores, derived data, rebase
├── openscion-server/ # openscion-server: the HTTP API over the engine
└── openscion/ # opnscn, the command-line client (crates.io: openscion)
benchmarks/ # openscion-bench: OpenScion vs git on synthetic repos
apps/web/ # the web GUI (Next.js + shadcn/ui)
docs/
├── research/ # source material: Scion, git, GitHub, GitLab
├── specs/ # what to build and why
├── plans/ # test-first implementation plans
└── adr/ # architecture decision records
openscion-proto holds every type that crosses a process boundary. It has no
transport, runtime, or behavior, so any tool can depend on it to speak the API.
openscion-core depends on it and re-exports it, so both crates name the same
types. The server, the CLI, and the GUI all speak that one API.
cargo run -p openscion-server # serves http://127.0.0.1:7878
cargo run -p openscion -- import ./tree # opnscn: record a tree as `main`
cargo run -p openscion -- commit -m " edit " --dir ./tree
cargo run -p openscion -- clone ./copy # or `--path dir` for one directory
cd apps/web && pnpm install && pnpm dev # GUI on http://localhost:3000
cargo run -p openscion-benchmarks --release # time the scenarios against git
Install the client from crates.io with cargo install openscion , which
provides the opnscn binary.
cargo fmt --all -- --check
cargo clippy --all-targets --all-features -- -D warnings
cargo build --all-targets --all-features
cargo test --all-features
CI runs those four checks. It also requires 90% line coverage in every source
file, builds rustdoc with -D warnings , builds with the MSRV, runs a
cargo-deny supply-chain check, and asserts that openscion-proto stays free
of transports and runtimes. To run the extra checks locally:
cargo doc --no-deps --all-features # CI builds this with RUSTDOCFLAGS="-D warnings"
cargo deny check all
cargo install cargo-llvm-cov # once
.github/scripts/check-file-coverage.sh 90 coverage.json
Releasing
Run the Release workflow from the Actions tab and pick a patch , minor ,
or major bump. The workflow re-validates the workspace, bumps the shared
[workspace.package] version, tags it, publishes openscion-proto ,
openscion-core , and openscion to crates.io (using the
CARGO_REGISTRY_TOKEN secret), and creates a GitHub release. Do not edit the
version in the root Cargo.toml by hand.
See CONTRIBUTING.md . Behavior changes start in
docs/specs/ ; each accepted spec gets a test-first
plan in docs/plans/ before any code is written.
OpenScion is an open-source, server-centric source control system for very large monorepos where most clients are AI agents and CI jobs rather than people on laptops. Based on Anthropic's Scion
www.anthropic.com/engineering/redesigning-source-control Resources
Readme GPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
