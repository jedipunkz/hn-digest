---
source: "https://layerfs.ai/"
hn_url: "https://news.ycombinator.com/item?id=50179857"
title: "LayerFS: The filesystem for AI coding agents"
article_title: "LayerFS: the filesystem for AI coding agents"
image: "https://layerfs.ai/github-preview.png"
author: "l2dy"
captured_at: "2026-10-21T13:30:12Z"
capture_tool: "hn-digest"
hn_id: 50179857
score: 10
comments: 6
posted_at: "2026-10-21T12:34:46Z"
tags:
  - hacker-news
---

# LayerFS: The filesystem for AI coding agents

- HN: [50179857](https://news.ycombinator.com/item?id=50179857)
- Source: [layerfs.ai](https://layerfs.ai/)
- Score: 10
- Comments: 6
- Posted: 2026-10-21T12:34:46Z

## Translation

Title: LayerFS: The filesystem for AI coding agents
Article title: LayerFS: the filesystem for AI coding agents
Description: LayerFS is an open-source (MIT) filesystem for AI coding agents: each tool call gets a complete, writable Workspace over shared history.

Article text:
Layer FS v1.0.0 → releases Docs Arch Architecture Bench Benchmarks Blog Docs Architecture Benchmarks Blog Releases v1.0.0 GitHub +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 Ephemerality: scaling multi-agent development beyond imagination
Ephemeral Workspaces.
Durable Shared History.
LayerFS is the substrate under your agents. Each one mounts a complete, writable Workspace over shared committed history, does its work, commits what it changed, and goes away.
WATCH VIDEO TRY IT LIVE ↓ Ephemeral workspaces
LayerFS × and DSH Copy to play with LayerFS on DSH LayerFS × and Pi soon
Compute exists only while a command runs.
fs · atomic write → Store write edit rm sandbox · workspace per bash mount exec commit unmount mount exec commit unmount mount exec commit unmount mount exec commit unmount mount exec commit unmount >_ mkdir out >_ cat log.txt >_ python train.py >_ node build.js UpToDate Store C41 C42 C43 C44 C45 C46 Compute: Agents spend ~80% thinking, ~10% in Bash.
FS: Each file tool call is one atomic Store op.
Sandbox: Each Bash tool call mounts an ephemeral workspace, then commits.
Agents on different computers share one serverless Store.
store.db C41 C42 C43 C44 C45 C46 C47 computer 1 computer 2 computer 3 computer 4 Shared storage: Every agent commits to one Store, with no DB server.
Distributed computing: Each agent runs on its own computer.
One project, many branches. One branch, many agents.
main auth ui login.ts api.ts session.ts app.tsx token.ts app.tsx nav.tsx Conflict: app.tsx app.tsx · merge markers <<<<<<< head C44 <Cart /> ======= <Search /> >>>>>>> workspace blame app.tsx C12 import Nav C44 <Header /> C44 <Cart /> C46 <Search /> L3 · publish L4 · publish Commits: Every tool call is a commit on its branch.
Conflict: A same-file clash is refused with Conflict.
Blame: Each line names the step that wrote it.
Fork recursively any commit, at any depth.
publish 0.62 0.48 0.21 0.40 0.78 0.66 0.30 0.71 0.91 0.69 pruned pruned pruned pruned pruned ★ best 0.91 explorers ×3 explorers ×3 explorers ×4 explorers ×3 explorers ×1 fork: no files copied Zero cost: A fork is metadata and copies no files.
Recursive: A filesystem for RSI agents and recursive exploration for RL training.
Dedup: Each content chunk is stored once.
Compression: The stored chunks shrink further.
Ratio: 10 copies → 1.05 dedup → 0.31 compression wtf???
Dependencies come with every mount. No wasting time for re-installation.
Git / other AgentFS: Drops ignored files; reinstalls.
LayerFS: Ignored files stay in every commit.
Comparison: ~5 min to install vs 0.05 s to mount.
Auto Checkpoints 07 → Every step an agent takes is a checkpoint to return to. Every call: Each file call is a commit; each Bash call commits when it ends. No noise: A call that changes nothing commits nothing. Rewind: Fork from any checkpoint and read any file as it was.
Append-only Objects 08 → Saved objects are never rewritten. The Store only grows. Content-addressed: An object is named by the hash of its bytes. Sealed: A written pack is never updated or deleted. Safe to share: Branches share objects; readers never wait for the writer.
Conflict Rules (.lww) 09 → Name the files that should never conflict. The newest edit wins. /.lww: Gitignore-style rules at the root of the tree. Newer wins: A matched file changed on both sides takes the newer edit. For: Lockfiles, build output and caches.
Big File, Small Edit 10 → An edit costs the size of the edit, not the size of the file. Overlay: A write keeps only the bytes it wrote. Commit: Only the edited extents and their tree path are rebuilt. Storage: The rest of the file is shared with its old version.
Agent tool calls and their history

## Original Extract

LayerFS is an open-source (MIT) filesystem for AI coding agents: each tool call gets a complete, writable Workspace over shared history.

Layer FS v1.0.0 → releases Docs Arch Architecture Bench Benchmarks Blog Docs Architecture Benchmarks Blog Releases v1.0.0 GitHub +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 +1 Ephemerality: scaling multi-agent development beyond imagination
Ephemeral Workspaces.
Durable Shared History.
LayerFS is the substrate under your agents. Each one mounts a complete, writable Workspace over shared committed history, does its work, commits what it changed, and goes away.
WATCH VIDEO TRY IT LIVE ↓ Ephemeral workspaces
LayerFS × and DSH Copy to play with LayerFS on DSH LayerFS × and Pi soon
Compute exists only while a command runs.
fs · atomic write → Store write edit rm sandbox · workspace per bash mount exec commit unmount mount exec commit unmount mount exec commit unmount mount exec commit unmount mount exec commit unmount >_ mkdir out >_ cat log.txt >_ python train.py >_ node build.js UpToDate Store C41 C42 C43 C44 C45 C46 Compute: Agents spend ~80% thinking, ~10% in Bash.
FS: Each file tool call is one atomic Store op.
Sandbox: Each Bash tool call mounts an ephemeral workspace, then commits.
Agents on different computers share one serverless Store.
store.db C41 C42 C43 C44 C45 C46 C47 computer 1 computer 2 computer 3 computer 4 Shared storage: Every agent commits to one Store, with no DB server.
Distributed computing: Each agent runs on its own computer.
One project, many branches. One branch, many agents.
main auth ui login.ts api.ts session.ts app.tsx token.ts app.tsx nav.tsx Conflict: app.tsx app.tsx · merge markers <<<<<<< head C44 <Cart /> ======= <Search /> >>>>>>> workspace blame app.tsx C12 import Nav C44 <Header /> C44 <Cart /> C46 <Search /> L3 · publish L4 · publish Commits: Every tool call is a commit on its branch.
Conflict: A same-file clash is refused with Conflict.
Blame: Each line names the step that wrote it.
Fork recursively any commit, at any depth.
publish 0.62 0.48 0.21 0.40 0.78 0.66 0.30 0.71 0.91 0.69 pruned pruned pruned pruned pruned ★ best 0.91 explorers ×3 explorers ×3 explorers ×4 explorers ×3 explorers ×1 fork: no files copied Zero cost: A fork is metadata and copies no files.
Recursive: A filesystem for RSI agents and recursive exploration for RL training.
Dedup: Each content chunk is stored once.
Compression: The stored chunks shrink further.
Ratio: 10 copies → 1.05 dedup → 0.31 compression wtf???
Dependencies come with every mount. No wasting time for re-installation.
Git / other AgentFS: Drops ignored files; reinstalls.
LayerFS: Ignored files stay in every commit.
Comparison: ~5 min to install vs 0.05 s to mount.
Auto Checkpoints 07 → Every step an agent takes is a checkpoint to return to. Every call: Each file call is a commit; each Bash call commits when it ends. No noise: A call that changes nothing commits nothing. Rewind: Fork from any checkpoint and read any file as it was.
Append-only Objects 08 → Saved objects are never rewritten. The Store only grows. Content-addressed: An object is named by the hash of its bytes. Sealed: A written pack is never updated or deleted. Safe to share: Branches share objects; readers never wait for the writer.
Conflict Rules (.lww) 09 → Name the files that should never conflict. The newest edit wins. /.lww: Gitignore-style rules at the root of the tree. Newer wins: A matched file changed on both sides takes the newer edit. For: Lockfiles, build output and caches.
Big File, Small Edit 10 → An edit costs the size of the edit, not the size of the file. Overlay: A write keeps only the bytes it wrote. Commit: Only the edited extents and their tree path are rebuilt. Storage: The rest of the file is shared with its old version.
Agent tool calls and their history
