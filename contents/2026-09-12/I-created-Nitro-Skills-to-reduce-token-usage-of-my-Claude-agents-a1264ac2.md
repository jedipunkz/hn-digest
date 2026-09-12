---
source: "https://nitroskills.com"
hn_url: "https://news.ycombinator.com/item?id=49678364"
title: "I created Nitro Skills to reduce token usage of my Claude agents"
article_title: "Nitro Skills — Token-Saving Skills for AI Coding Agents"
image: ""
author: "dextrata"
captured_at: "2026-09-12T23:47:41Z"
capture_tool: "hn-digest"
hn_id: 49678364
score: 1
comments: 0
posted_at: "2026-09-12T23:34:51Z"
tags:
  - hacker-news
---

# I created Nitro Skills to reduce token usage of my Claude agents

- HN: [49678364](https://news.ycombinator.com/item?id=49678364)
- Source: [nitroskills.com](https://nitroskills.com)
- Score: 1
- Comments: 0
- Posted: 2026-09-12T23:34:51Z

## Translation

Title: I created Nitro Skills to reduce token usage of my Claude agents
Article title: Nitro Skills — Token-Saving Skills for AI Coding Agents
Description: Nitro Skills is a free, open-source set of Claude Code skills that cut the number of tokens an AI coding agent burns while working — smaller reads, fewer repeated edits, mechanical refactors, and more.

Article text:
Skills that cut the tokens your AI coding agent burns while working.
Each skill is a short SKILL.md — the only part that ever enters the model's context — plus a script that gets executed, never read. The clever logic costs zero tokens no matter how many times it's reused.
View on GitHub
Explore skills
Fifteen skills, one job each
The first three shrink reads, edits and repeated command output. The rest attack everything else: repeated output, long strings, reassurance reads, mechanical edits, logs, payloads, traces, diffs, search hops, re-derived facts, and unbounded floods.
Edit files by pointing at hashed line anchors instead of quoting whole files.
Cache and reuse output from tests, builds, and lints across repeat runs.
Inspect a module's API by running it, not by reading its source.
Fold already-shown output blocks into one-line pointers.
Shorten long repeated strings (hashes, UUIDs, paths) into small tokens.
Verify a claim about a file against reality without re-reading it.
Describe a mechanical multi-file edit in a sentence instead of a diff.
Cluster huge logs into a handful of templated lines with counts.
Print the schema of structured data instead of dumping the payload.
Strip library noise from stack traces down to the user-code frames.
Summarize a diff by symbol instead of printing every changed line.
Answer structural code questions from a cached symbol index.
Show a symbol's callers and callees before you change its signature.
Persist answers to "where/how is X handled" across sessions.
Track token spend per command and block unbounded floods.
Requires only Python 3. Installs into your global Claude directory and wires up the supporting hooks automatically.
git clone https://github.com/Dextrata/nitro-skills.git
cd nitro-skills
python install.py
See the full README for manual install, flags, and a CLAUDE.md example that makes usage legible to the agent.
Nitro Skills by Dextrata — github.com/Dextrata/nitro-skills

## Original Extract

Nitro Skills is a free, open-source set of Claude Code skills that cut the number of tokens an AI coding agent burns while working — smaller reads, fewer repeated edits, mechanical refactors, and more.

Skills that cut the tokens your AI coding agent burns while working.
Each skill is a short SKILL.md — the only part that ever enters the model's context — plus a script that gets executed, never read. The clever logic costs zero tokens no matter how many times it's reused.
View on GitHub
Explore skills
Fifteen skills, one job each
The first three shrink reads, edits and repeated command output. The rest attack everything else: repeated output, long strings, reassurance reads, mechanical edits, logs, payloads, traces, diffs, search hops, re-derived facts, and unbounded floods.
Edit files by pointing at hashed line anchors instead of quoting whole files.
Cache and reuse output from tests, builds, and lints across repeat runs.
Inspect a module's API by running it, not by reading its source.
Fold already-shown output blocks into one-line pointers.
Shorten long repeated strings (hashes, UUIDs, paths) into small tokens.
Verify a claim about a file against reality without re-reading it.
Describe a mechanical multi-file edit in a sentence instead of a diff.
Cluster huge logs into a handful of templated lines with counts.
Print the schema of structured data instead of dumping the payload.
Strip library noise from stack traces down to the user-code frames.
Summarize a diff by symbol instead of printing every changed line.
Answer structural code questions from a cached symbol index.
Show a symbol's callers and callees before you change its signature.
Persist answers to "where/how is X handled" across sessions.
Track token spend per command and block unbounded floods.
Requires only Python 3. Installs into your global Claude directory and wires up the supporting hooks automatically.
git clone https://github.com/Dextrata/nitro-skills.git
cd nitro-skills
python install.py
See the full README for manual install, flags, and a CLAUDE.md example that makes usage legible to the agent.
Nitro Skills by Dextrata — github.com/Dextrata/nitro-skills
