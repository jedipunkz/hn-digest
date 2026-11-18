---
source: "https://github.com/gabrielvfonseca/replica"
hn_url: "https://news.ycombinator.com/item?id=50553493"
title: "Show HN: 100ms copy-on-write workspaces for AI agents on macOS"
article_title: "GitHub - gabrielvfonseca/replica: Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. · GitHub"
image: "https://opengraph.githubassets.com/52b17e9b0edf97874bdb9c1b50be35888c431d173c1061e655f88fc74d82c77f/gabrielvfonseca/replica"
author: "gabrielfonseca"
captured_at: "2026-11-18T23:17:37Z"
capture_tool: "hn-digest"
hn_id: 50553493
score: 2
comments: 0
posted_at: "2026-11-18T22:51:02Z"
tags:
  - hacker-news
---

# Show HN: 100ms copy-on-write workspaces for AI agents on macOS

- HN: [50553493](https://news.ycombinator.com/item?id=50553493)
- Source: [github.com](https://github.com/gabrielvfonseca/replica)
- Score: 2
- Comments: 0
- Posted: 2026-11-18T22:51:02Z

## Translation

Title: Show HN: 100ms copy-on-write workspaces for AI agents on macOS
Article title: GitHub - gabrielvfonseca/replica: Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. · GitHub
Description: Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. - gabrielvfonseca/replica

Article text:
GitHub - gabrielvfonseca/replica: Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gabrielvfonseca
replica Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing.
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
65 Commits 65 Commits Folders and files
.github .github .zed .zed Formula Formula benchmarks benchmarks docs docs packaging packaging scripts scripts skills skills src src templates templates tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml HOMEBREW.md HOMEBREW.md LICENSE.md LICENSE.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md banner.jpg banner.jpg clippy.toml clippy.toml deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
Isolated replicas of your workspace for AI agents
Replica gives every agent run its own fully installed copy of your project, created in milliseconds with APFS copy-on-write. Each copy is a real git checkout, and automatic cleanup keeps total disk use bounded.
cd " $( replica new --ttl 2h --name fix-login ) " # a ready workspace; the path goes to stdout
git checkout -b agent/fix-login && git commit -am " fix login " && git push -u origin agent/fix-login
replica rm < id > # done; or let the TTL expire
Why
Agent workloads are dominated by creating and deleting small files (a checkout,
node_modules , build output), and that's where macOS is slowest. A fresh
install takes seconds to minutes before the agent does any work.
Replica builds a base once: a git clone with its dependencies installed.
Each replica new clones that base with one clonefile(2) call. The data
blocks stay shared with the base until a file changes, so the new workspace is
ready almost at once. Measured on an Apple M1 Max, replica new of a 1,700-file
project takes about 100 ms at N = 1 and about 230 ms at N = 10 concurrent
workspaces, versus seconds for git clone plus an install.
Numbers and methodology .
Fast workspaces. A clone of a 1,700-file project takes about 100 ms at N = 1 (about 230 ms with 10 concurrent workspaces), against seconds for git clone plus an install. Exact numbers vary with machine load; see docs/BENCHMARKS.md for the full table and methodology.
Real git. Each workspace is a full git checkout. It can branch, commit, and push, and the base stays untouched. Details below .
Bounded disk. TTL expiry, least-recently-used eviction, a byte cap, and an optional dedicated APFS volume with a quota.
Cheap updates. replica base refresh brings the base up to date with its remote without a full rebuild. Details below .
Agent-ready. A stable CLI with --json output, documented exit codes, and skills for Claude Code, Cursor, Copilot, Zed and Aider.
Concurrency tested. tests/concurrency.rs runs 10 simultaneous replica new /use/ rm cycles with zero SQLite busy errors, no corruption, and all workspaces created and removed. replica gc skips workspaces with a fresh lease ( .replica/lease ), so a launchd gc every 10 minutes never deletes a workspace an agent is using.
macOS on an APFS volume. Replica uses clonefile(2) , which Linux does not provide. Tested on macOS 26.6 (Apple M1 Max).
The base tree and every workspace must be on the same APFS volume . Clones cannot cross volumes, and Replica refuses rather than silently copying.
Binaries are published for macOS on Apple Silicon ( arm64 ) and Intel ( x86_64 ). A universal binary covers both.
# Homebrew is not available yet; see HOMEBREW.md.
# Installer script: verifies the download against sha256sums.txt
curl -fsSL https://raw.githubusercontent.com/gabrielvfonseca/replica/main/scripts/get-replica.sh | sh
# From source (needs a Rust toolchain, https://rustup.rs)
git clone https://github.com/gabrielvfonseca/replica.git && cd replica
make release && PREFIX= $HOME /.local make install
Run replica help for a guided overview, or replica help quickstart to get started.
1. Choose a volume. A dedicated APFS volume with a quota is the safest setup, because clones can never fill your main disk.
diskutil apfs list # find your APFS container
./scripts/setup-volume.sh disk3 100g Work # creates /Volumes/Work, turns off Spotlight indexing there
export REPLICA_ROOT=/Volumes/Work/replica
Any APFS directory works too. Add REPLICA_MAX_BYTES to cap its size. See DISK_SAFETY for the trade-offs.
replica base build --repo https://github.com/you/proj.git --setup " npm ci "
replica base check # find build output that breaks when cloned
3. Create, use, and remove workspaces.
cd " $( replica new --ttl 4h --name feature-x ) " # path on stdout, status on stderr
eval " $( replica env < id > ) " # per-workspace TMPDIR and shared caches
replica touch < id > # keep a long session out of LRU eviction
replica list # table; add --json for scripts
replica rm < id > # remove one (idempotent)
Branches with different lockfiles. If you create a workspace from a branch
whose lockfile differs from the base, use --ref :
replica new --ref my-feature-branch
Replica compares the branch's lockfile hash with the base's recorded hash. If
they match, the workspace is cloned directly ( cloned directly ). If they differ,
Replica clones the base, checks out my-feature-branch , then runs the setup
command (an incremental install) so the workspace gets correct dependencies, and
reports lockfile drift on branch my-feature-branch; running incremental install .
The CLI always tells you which path it took.
4. Give your agents the skill.
npx skills add gabrielvfonseca/replica # skills.sh CLI: all six skills
replica setup --dry-run # see what setup would write for Claude Code, Cursor, Copilot, Zed, Aider
replica setup # write it
Git: branch, commit, and push from a workspace
Yes. A workspace is a complete git repository, not a read-only copy:
Remote credentials. push uses whatever credentials git uses on this machine (SSH keys, a credential helper). Replica does not add or store any.
The remote is the one you built from. origin in a workspace is the --repo you passed to base build . If that is a local path, pushes go to that path.
Workspaces are independent. Two workspaces on the same branch will conflict on push, the same as two clones would. Use one branch per workspace.
Verified. tests/git_workspace.rs builds a base from a bare repo, commits and pushes from a workspace, and checks that the origin and the base are as expected.
For agents that open pull requests, the flow is: replica new , work and commit on a branch, push, open the PR with gh pr create , then replica rm .
When the remote moves on, replica base refresh updates the base. It does not rebuild from scratch:
Clonefiles the current base. This copies metadata only, so it is quick and uses almost no extra disk.
Runs git fetch and resets the copy to origin/HEAD . Only files that changed upstream are rewritten.
Re-runs setup only if a lockfile changed . If none did, setup is skipped entirely.
Swaps the new base in. Existing workspaces are unaffected, because each one owns its blocks.
So an upstream commit that does not touch dependencies costs a fetch and a checkout. A lockfile change costs about as much as the setup command itself. If the fast path fails for any reason, refresh falls back to a full rebuild.
Check freshness with replica base status , and use replica new --require-fresh to refuse a stale base.
Disk usage: what a workspace really costs
At creation , a workspace adds APFS metadata for every file and directory it contains. For the benchmark fixture (about 1,700 files) that is about 1.4 MB. Data blocks are shared with the base.
After the agent works , every file it changes takes real space, because APFS copies the blocks that change. Rewriting 500 MB of build output costs 500 MB, and so does a npm rebuild that rewrites node_modules .
Deleting a workspace frees only the blocks no other file shares. A fresh workspace frees little, and a heavily changed one frees much more.
This is why the TTL, LRU eviction, and byte cap exist. DISK_SAFETY explains each one.
Spotlight can also add writes on a busy volume. scripts/setup-volume.sh turns indexing off for the dedicated volume only (it needs sudo ). We have not yet measured how much that changes benchmark results.
Path-dependent build output. Some tools write absolute paths into their output: Python venvs, node_modules/.bin scripts, TypeScript .tsbuildinfo . These break in a clone. Run replica base check to find them, then fix them with --post-clone or by regenerating them. See the replica-base skill .
Cargo target/ clones safely but adds cost. A cloned target/ gives a correct incremental no-op build (~20 ms, no recompile) when nothing changed, and recompiles only changed crates after a source touch — it does not force a full rebuild. However, target/ is created by the setup command, so .gitignore does not keep it out of the base. Exclude it another way: either don't build target/ during setup, or add an exclusion option to your base build. A cloned target/ of a large project can add tens of MB of APFS metadata per clone, so exclude it if disk cost matters. A shared CARGO_TARGET_DIR under $REPLICA_ROOT/cache keeps incremental builds fast but reintroduces Cargo's file-lock waiting, so parallel agents can block on Blocking waiting for file lock . For parallel use, consider sccache instead of a shared target dir.
Cost scales with file count. Each clone writes metadata for every file. Replica is much faster than copying or cloning file by file, but it is not free, and very large trees (hundreds of thousands of files) have not been measured.
Same volume only. Source, base, and workspaces must share one APFS volume.
Not a speedup inside a running agent. Replica makes workspace creation fast. It does not speed up reads or writes during a run, does not reduce the RAM or CPU used by running many agents at once, and does not fix Spotlight or XProtect scanning.
macOS only. Linux and Windows are out of scope.
The suite in benchmarks/ compares replica new with
git clone plus install, git worktree plus install, pnpm with the global
virtual store, and cp -c , and it also measures replica rm and
replica base refresh , at concurrency N = 1, 5, and 10.
make release # target/release/replica
make bench # bundled fixture (~15 min)
m

[truncated]

## Original Extract

Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. - gabrielvfonseca/replica

GitHub - gabrielvfonseca/replica: Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gabrielvfonseca
replica Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Instant, disk-safe workspaces for AI agents on macOS. Clones a prebuilt project tree with APFS clonefile instead of checking out and installing.
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
65 Commits 65 Commits Folders and files
.github .github .zed .zed Formula Formula benchmarks benchmarks docs docs packaging packaging scripts scripts skills skills src src templates templates tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml HOMEBREW.md HOMEBREW.md LICENSE.md LICENSE.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md banner.jpg banner.jpg clippy.toml clippy.toml deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
Isolated replicas of your workspace for AI agents
Replica gives every agent run its own fully installed copy of your project, created in milliseconds with APFS copy-on-write. Each copy is a real git checkout, and automatic cleanup keeps total disk use bounded.
cd " $( replica new --ttl 2h --name fix-login ) " # a ready workspace; the path goes to stdout
git checkout -b agent/fix-login && git commit -am " fix login " && git push -u origin agent/fix-login
replica rm < id > # done; or let the TTL expire
Why
Agent workloads are dominated by creating and deleting small files (a checkout,
node_modules , build output), and that's where macOS is slowest. A fresh
install takes seconds to minutes before the agent does any work.
Replica builds a base once: a git clone with its dependencies installed.
Each replica new clones that base with one clonefile(2) call. The data
blocks stay shared with the base until a file changes, so the new workspace is
ready almost at once. Measured on an Apple M1 Max, replica new of a 1,700-file
project takes about 100 ms at N = 1 and about 230 ms at N = 10 concurrent
workspaces, versus seconds for git clone plus an install.
Numbers and methodology .
Fast workspaces. A clone of a 1,700-file project takes about 100 ms at N = 1 (about 230 ms with 10 concurrent workspaces), against seconds for git clone plus an install. Exact numbers vary with machine load; see docs/BENCHMARKS.md for the full table and methodology.
Real git. Each workspace is a full git checkout. It can branch, commit, and push, and the base stays untouched. Details below .
Bounded disk. TTL expiry, least-recently-used eviction, a byte cap, and an optional dedicated APFS volume with a quota.
Cheap updates. replica base refresh brings the base up to date with its remote without a full rebuild. Details below .
Agent-ready. A stable CLI with --json output, documented exit codes, and skills for Claude Code, Cursor, Copilot, Zed and Aider.
Concurrency tested. tests/concurrency.rs runs 10 simultaneous replica new /use/ rm cycles with zero SQLite busy errors, no corruption, and all workspaces created and removed. replica gc skips workspaces with a fresh lease ( .replica/lease ), so a launchd gc every 10 minutes never deletes a workspace an agent is using.
macOS on an APFS volume. Replica uses clonefile(2) , which Linux does not provide. Tested on macOS 26.6 (Apple M1 Max).
The base tree and every workspace must be on the same APFS volume . Clones cannot cross volumes, and Replica refuses rather than silently copying.
Binaries are published for macOS on Apple Silicon ( arm64 ) and Intel ( x86_64 ). A universal binary covers both.
# Homebrew is not available yet; see HOMEBREW.md.
# Installer script: verifies the download against sha256sums.txt
curl -fsSL https://raw.githubusercontent.com/gabrielvfonseca/replica/main/scripts/get-replica.sh | sh
# From source (needs a Rust toolchain, https://rustup.rs)
git clone https://github.com/gabrielvfonseca/replica.git && cd replica
make release && PREFIX= $HOME /.local make install
Run replica help for a guided overview, or replica help quickstart to get started.
1. Choose a volume. A dedicated APFS volume with a quota is the safest setup, because clones can never fill your main disk.
diskutil apfs list # find your APFS container
./scripts/setup-volume.sh disk3 100g Work # creates /Volumes/Work, turns off Spotlight indexing there
export REPLICA_ROOT=/Volumes/Work/replica
Any APFS directory works too. Add REPLICA_MAX_BYTES to cap its size. See DISK_SAFETY for the trade-offs.
replica base build --repo https://github.com/you/proj.git --setup " npm ci "
replica base check # find build output that breaks when cloned
3. Create, use, and remove workspaces.
cd " $( replica new --ttl 4h --name feature-x ) " # path on stdout, status on stderr
eval " $( replica env < id > ) " # per-workspace TMPDIR and shared caches
replica touch < id > # keep a long session out of LRU eviction
replica list # table; add --json for scripts
replica rm < id > # remove one (idempotent)
Branches with different lockfiles. If you create a workspace from a branch
whose lockfile differs from the base, use --ref :
replica new --ref my-feature-branch
Replica compares the branch's lockfile hash with the base's recorded hash. If
they match, the workspace is cloned directly ( cloned directly ). If they differ,
Replica clones the base, checks out my-feature-branch , then runs the setup
command (an incremental install) so the workspace gets correct dependencies, and
reports lockfile drift on branch my-feature-branch; running incremental install .
The CLI always tells you which path it took.
4. Give your agents the skill.
npx skills add gabrielvfonseca/replica # skills.sh CLI: all six skills
replica setup --dry-run # see what setup would write for Claude Code, Cursor, Copilot, Zed, Aider
replica setup # write it
Git: branch, commit, and push from a workspace
Yes. A workspace is a complete git repository, not a read-only copy:
Remote credentials. push uses whatever credentials git uses on this machine (SSH keys, a credential helper). Replica does not add or store any.
The remote is the one you built from. origin in a workspace is the --repo you passed to base build . If that is a local path, pushes go to that path.
Workspaces are independent. Two workspaces on the same branch will conflict on push, the same as two clones would. Use one branch per workspace.
Verified. tests/git_workspace.rs builds a base from a bare repo, commits and pushes from a workspace, and checks that the origin and the base are as expected.
For agents that open pull requests, the flow is: replica new , work and commit on a branch, push, open the PR with gh pr create , then replica rm .
When the remote moves on, replica base refresh updates the base. It does not rebuild from scratch:
Clonefiles the current base. This copies metadata only, so it is quick and uses almost no extra disk.
Runs git fetch and resets the copy to origin/HEAD . Only files that changed upstream are rewritten.
Re-runs setup only if a lockfile changed . If none did, setup is skipped entirely.
Swaps the new base in. Existing workspaces are unaffected, because each one owns its blocks.
So an upstream commit that does not touch dependencies costs a fetch and a checkout. A lockfile change costs about as much as the setup command itself. If the fast path fails for any reason, refresh falls back to a full rebuild.
Check freshness with replica base status , and use replica new --require-fresh to refuse a stale base.
Disk usage: what a workspace really costs
At creation , a workspace adds APFS metadata for every file and directory it contains. For the benchmark fixture (about 1,700 files) that is about 1.4 MB. Data blocks are shared with the base.
After the agent works , every file it changes takes real space, because APFS copies the blocks that change. Rewriting 500 MB of build output costs 500 MB, and so does a npm rebuild that rewrites node_modules .
Deleting a workspace frees only the blocks no other file shares. A fresh workspace frees little, and a heavily changed one frees much more.
This is why the TTL, LRU eviction, and byte cap exist. DISK_SAFETY explains each one.
Spotlight can also add writes on a busy volume. scripts/setup-volume.sh turns indexing off for the dedicated volume only (it needs sudo ). We have not yet measured how much that changes benchmark results.
Path-dependent build output. Some tools write absolute paths into their output: Python venvs, node_modules/.bin scripts, TypeScript .tsbuildinfo . These break in a clone. Run replica base check to find them, then fix them with --post-clone or by regenerating them. See the replica-base skill .
Cargo target/ clones safely but adds cost. A cloned target/ gives a correct incremental no-op build (~20 ms, no recompile) when nothing changed, and recompiles only changed crates after a source touch — it does not force a full rebuild. However, target/ is created by the setup command, so .gitignore does not keep it out of the base. Exclude it another way: either don't build target/ during setup, or add an exclusion option to your base build. A cloned target/ of a large project can add tens of MB of APFS metadata per clone, so exclude it if disk cost matters. A shared CARGO_TARGET_DIR under $REPLICA_ROOT/cache keeps incremental builds fast but reintroduces Cargo's file-lock waiting, so parallel agents can block on Blocking waiting for file lock . For parallel use, consider sccache instead of a shared target dir.
Cost scales with file count. Each clone writes metadata for every file. Replica is much faster than copying or cloning file by file, but it is not free, and very large trees (hundreds of thousands of files) have not been measured.
Same volume only. Source, base, and workspaces must share one APFS volume.
Not a speedup inside a running agent. Replica makes workspace creation fast. It does not speed up reads or writes during a run, does not reduce the RAM or CPU used by running many agents at once, and does not fix Spotlight or XProtect scanning.
macOS only. Linux and Windows are out of scope.
The suite in benchmarks/ compares replica new with
git clone plus install, git worktree plus install, pnpm with the global
virtual store, and cp -c , and it also measures replica rm and
replica base refresh , at concurrency N = 1, 5, and 10.
make release # target/release/replica
make bench # bundled fixture (~15 min)
m

[truncated]
