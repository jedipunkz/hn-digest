---
source: "https://gobstopper.sh"
hn_url: "https://news.ycombinator.com/item?id=49944789"
title: "Gobstopper – context compaction proxy for Claude Code and Codex"
article_title: "Gobstopper: context compaction proxy for Claude Code and Codex"
image: "https://gobstopper.sh/opengraph-image"
author: "benzguo"
captured_at: "2026-10-03T15:03:40Z"
capture_tool: "hn-digest"
hn_id: 49944789
score: 1
comments: 0
posted_at: "2026-10-03T14:49:52Z"
tags:
  - hacker-news
---

# Gobstopper – context compaction proxy for Claude Code and Codex

- HN: [49944789](https://news.ycombinator.com/item?id=49944789)
- Source: [gobstopper.sh](https://gobstopper.sh)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T14:49:52Z

## Translation

Title: Gobstopper – context compaction proxy for Claude Code and Codex
Article title: Gobstopper: context compaction proxy for Claude Code and Codex
Description: Gobstopper makes long Claude Code and Codex sessions smaller. A local proxy compacts live requests, and saved sessions get a smaller copy beside the original.

Article text:
Gobstopper: context compaction proxy for Claude Code and Codex Skip to content Gobstopper Docs Benchmarks Compare Blog GitHub Install Theme Catppuccin Gruvbox Rosé Pine Tokyo Night Paper Appearance Light Dark System Context compaction you can undo.
Keep long coding sessions smaller, with the original available to restore.
Copy curl -fsSL https://gobstopper.sh/install.sh | sh Apple silicon Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Copy curl -fsSL https://gobstopper.sh/install.sh | sh x86_64 and ARM64, glibc 2.34+ Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Copy irm https://gobstopper.sh/install.ps1 | iex x86_64 · the proxy works; the vault, watch and provider hooks need macOS or Linux Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Start Claude Code through Gobstopper. Small requests pass through unchanged; large ones keep the system prompt, the first task, and the newest turns word for word while older turns become a summary.
gobstopper proxy run -- claude Session Without Gobstopper With Gobstopper Point in the session Early Halfway End claude Request 20 of 383 ~ 63,920 estimated tokens sent with this request threshold after calibration: ~ 64,000 no proxy: nothing is ever summarized claude Request 191 of 383 ~ 327,458 estimated tokens sent with this request threshold after calibration: ~ 77,154 no proxy: nothing is ever summarized claude Request 383 of 383 ~ 491,135 estimated tokens sent with this request threshold after calibration: ~ 82,315 no proxy: nothing is ever summarized claude, through gobstopper proxy Request 20 of 383 ~ 63,920 estimated tokens sent with this request threshold after calibration: ~ 64,000 summarized 0 times so far claude, through gobstopper proxy Request 191 of 383 ~ 38,967 estimated tokens sent with this request threshold after calibration: ~ 77,154 summarized 7 times so far claude, through gobstopper proxy Request 383 of 383 ~ 68,581 estimated tokens sent with this request threshold after calibration: ~ 82,315 summarized 10 times so far claude, through gobstopper proxy Request 191 of 383 ~ 38,967 estimated tokens sent with this request threshold after calibration: ~ 77,154 summarized 7 times so far One recorded Claude Code session replayed at the 128,000-token default. Sizes estimate four characters per token. The proxy starts on a free local port and stops when the agent exits. Its default threshold is 128,000 estimated tokens. It uses local rules to summarize, with no extra model call, and leaves your session files unchanged.
CC Claude Code C Codex O opencode C Crush A Aider G Goose Claude Code and Codex routing is live-checked. Other agents need a configurable provider address; OpenAI Chat Completions support is tested with synthetic histories. Proxy setup and options .
Try a smaller copy. Keep a way back.
Find a Claude Code or Codex session and inspect what a strategy would remove before writing anything.
Compaction creates a separate copy and saves the original bytes in a local vault. Search that vault or prepare a restored copy when you need an older detail.
Copy checks cover record links, order, and tool-call pairs. They cannot guarantee provider resume, retention of every task fact, or what you will be billed.
Step 1 of 5 Preview a saved session gobstopper detect
gobstopper plan < session > --trigger 250000 Terminal-Bench 2.1
About as many tasks solved, 29% fewer tokens sent.
One recorded run: Gobstopper at tail 0 resolved 61 of 89 tasks against 60 with no proxy, within single-trial noise, and sent 29% fewer input tokens. The benchmarks page has the setup and limits.
29% fewer input tokens in one benchmark
Gobstopper, tail 0 vs Claude Code, no proxy · about as many tasks solved, within single-trial noise
Gobstopper, tail 0 61 of 89 solved
Claude Code, no proxy 60 of 89 solved
Gobstopper, tail 40 (old default) 59 of 89 solved
Bar chart of total input tokens over 89 Terminal-Bench tasks. Gobstopper, tail 0: 84.3 million, 61 solved. Claude Code, no proxy: 118.6 million, 60 solved. Gobstopper, tail 40 (old default): 118.5 million, 59 solved. Cache reads make up most of each bar.
These results cover one model and setup; they do not establish task-quality or billing improvements for other agents. Setup, statistics and downloads .
No narration. Captions carry every line. The film never plays until you press play.
The film is also available as text below.
Your coding agent resends everything.
Every step. The whole session.
Long sessions pay for it again.
Keep the start. Keep the last three.
Long printouts out. Files stay on disk.
Same prefix, so the cache keeps matching.
Almost all of it: cache reads.
So v0.7.3 changed the default.
Your session files are never edited.
One trial. One model. Read the limits.
gobstopper proxy run -- claude
Terminal-Bench 2.1 · 89 tasks · 1 trial per arm · GLM 5.3 Flash · 45K threshold (default 128K) · Sept 27–28, 2026. Resolution is within single-trial noise. Cost is provider-reported, metered through Vercel AI Gateway, for this model, at a 45K threshold (default 128K).
Each installer downloads the latest release for your platform, checks its SHA-256, and installs it for your user only. The proxy needs curl 8.3 or newer.
Then run one Claude Code session through the proxy: gobstopper proxy run -- claude .
Latest tagged release: v 0.8.5 .
Built-in inspection makes no model call. An optional remote scorer receives selected transcript text; trusted plugins run your code. Installation and command reference .
Any agent that lets you set a custom provider address and uses Anthropic Messages, OpenAI Responses, or OpenAI Chat Completions. Claude Code and Codex routing is live-checked; Chat Completions is contract-tested on synthetic histories. File commands read Claude Code and Codex session formats.
Before writing a separate Claude Code or Codex copy, Gobstopper archives the original and prepared bytes. Search the local vault for a missing record, read its saved text, or run gobstopper undo to prepare a restored copy with a new session identity. Keep backups of the vault: recovery depends on your storage.
The proxy changes outgoing requests and leaves session files unchanged. File commands prepare separate copies. Gobstopper does not trigger Claude Code's or Codex's own compaction and does not edit session files in place. Automatic provider compaction stays disabled.
You or the agent can reserve a larger context for a difficult phase, within the model and client's supported window. The budget expires by time or request count. Optional adaptive rescue temporarily raises the threshold after repeated reads of unchanged evidence the proxy removed. Gobstopper also keeps selected original tool results and images through repeated compactions. These controls reduce evidence loss; they cannot guarantee task progress. See context budgets and retention .
Request and compaction deadlines limit stalls, while logging and metrics run in background workers. gobstopper proxy launch --client claude checks readiness before opening Claude Code and can start directly with its official provider when the configuration permits. Codex launch requires a healthy proxy. Existing sessions using a fixed proxy URL still depend on that listener. See recovery and fallback requirements .
Run gobstopper proxy install to start a user service at login and restart it after a process exit. Managed service changes wait for active inference to finish. During inference, Gobstopper requests idle-sleep prevention; closing a lid or forcing sleep can override it. Local session metrics are available through gobstopper data metrics . See startup and local visibility .
Use your agent's built-in /compact when you want a model-written summary without another tool. Gobstopper adds previews and archived originals for saved sessions, plus a proxy that applies local rules without a model call. See the Claude Code comparison or the CliffCompaction comparison .
Excalibur (xcb) xcb.sh Routes coding tasks across the Claude, Codex, and Devin plans you have
ChatGPT Claude Perplexity Grok Gobstopper Context compaction you can undo · MIT or Apache-2.0 · Preview
by Hraness Accept cookies Analytics preferences Your browser remembers your appearance setting and this choice. We use analytics to understand how the site is used. None of it is used for advertising or cross-site tracking. You can change your analytics choice here at any time. Privacy policy Accept analytics Decline analytics

## Original Extract

Gobstopper makes long Claude Code and Codex sessions smaller. A local proxy compacts live requests, and saved sessions get a smaller copy beside the original.

Gobstopper: context compaction proxy for Claude Code and Codex Skip to content Gobstopper Docs Benchmarks Compare Blog GitHub Install Theme Catppuccin Gruvbox Rosé Pine Tokyo Night Paper Appearance Light Dark System Context compaction you can undo.
Keep long coding sessions smaller, with the original available to restore.
Copy curl -fsSL https://gobstopper.sh/install.sh | sh Apple silicon Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Copy curl -fsSL https://gobstopper.sh/install.sh | sh x86_64 and ARM64, glibc 2.34+ Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Copy irm https://gobstopper.sh/install.ps1 | iex x86_64 · the proxy works; the vault, watch and provider hooks need macOS or Linux Build from source (Rust 1.85+) Copy cargo install --git https://github.com/hraness/gobstopper gobstopper --locked
Start Claude Code through Gobstopper. Small requests pass through unchanged; large ones keep the system prompt, the first task, and the newest turns word for word while older turns become a summary.
gobstopper proxy run -- claude Session Without Gobstopper With Gobstopper Point in the session Early Halfway End claude Request 20 of 383 ~ 63,920 estimated tokens sent with this request threshold after calibration: ~ 64,000 no proxy: nothing is ever summarized claude Request 191 of 383 ~ 327,458 estimated tokens sent with this request threshold after calibration: ~ 77,154 no proxy: nothing is ever summarized claude Request 383 of 383 ~ 491,135 estimated tokens sent with this request threshold after calibration: ~ 82,315 no proxy: nothing is ever summarized claude, through gobstopper proxy Request 20 of 383 ~ 63,920 estimated tokens sent with this request threshold after calibration: ~ 64,000 summarized 0 times so far claude, through gobstopper proxy Request 191 of 383 ~ 38,967 estimated tokens sent with this request threshold after calibration: ~ 77,154 summarized 7 times so far claude, through gobstopper proxy Request 383 of 383 ~ 68,581 estimated tokens sent with this request threshold after calibration: ~ 82,315 summarized 10 times so far claude, through gobstopper proxy Request 191 of 383 ~ 38,967 estimated tokens sent with this request threshold after calibration: ~ 77,154 summarized 7 times so far One recorded Claude Code session replayed at the 128,000-token default. Sizes estimate four characters per token. The proxy starts on a free local port and stops when the agent exits. Its default threshold is 128,000 estimated tokens. It uses local rules to summarize, with no extra model call, and leaves your session files unchanged.
CC Claude Code C Codex O opencode C Crush A Aider G Goose Claude Code and Codex routing is live-checked. Other agents need a configurable provider address; OpenAI Chat Completions support is tested with synthetic histories. Proxy setup and options .
Try a smaller copy. Keep a way back.
Find a Claude Code or Codex session and inspect what a strategy would remove before writing anything.
Compaction creates a separate copy and saves the original bytes in a local vault. Search that vault or prepare a restored copy when you need an older detail.
Copy checks cover record links, order, and tool-call pairs. They cannot guarantee provider resume, retention of every task fact, or what you will be billed.
Step 1 of 5 Preview a saved session gobstopper detect
gobstopper plan < session > --trigger 250000 Terminal-Bench 2.1
About as many tasks solved, 29% fewer tokens sent.
One recorded run: Gobstopper at tail 0 resolved 61 of 89 tasks against 60 with no proxy, within single-trial noise, and sent 29% fewer input tokens. The benchmarks page has the setup and limits.
29% fewer input tokens in one benchmark
Gobstopper, tail 0 vs Claude Code, no proxy · about as many tasks solved, within single-trial noise
Gobstopper, tail 0 61 of 89 solved
Claude Code, no proxy 60 of 89 solved
Gobstopper, tail 40 (old default) 59 of 89 solved
Bar chart of total input tokens over 89 Terminal-Bench tasks. Gobstopper, tail 0: 84.3 million, 61 solved. Claude Code, no proxy: 118.6 million, 60 solved. Gobstopper, tail 40 (old default): 118.5 million, 59 solved. Cache reads make up most of each bar.
These results cover one model and setup; they do not establish task-quality or billing improvements for other agents. Setup, statistics and downloads .
No narration. Captions carry every line. The film never plays until you press play.
The film is also available as text below.
Your coding agent resends everything.
Every step. The whole session.
Long sessions pay for it again.
Keep the start. Keep the last three.
Long printouts out. Files stay on disk.
Same prefix, so the cache keeps matching.
Almost all of it: cache reads.
So v0.7.3 changed the default.
Your session files are never edited.
One trial. One model. Read the limits.
gobstopper proxy run -- claude
Terminal-Bench 2.1 · 89 tasks · 1 trial per arm · GLM 5.3 Flash · 45K threshold (default 128K) · Sept 27–28, 2026. Resolution is within single-trial noise. Cost is provider-reported, metered through Vercel AI Gateway, for this model, at a 45K threshold (default 128K).
Each installer downloads the latest release for your platform, checks its SHA-256, and installs it for your user only. The proxy needs curl 8.3 or newer.
Then run one Claude Code session through the proxy: gobstopper proxy run -- claude .
Latest tagged release: v 0.8.5 .
Built-in inspection makes no model call. An optional remote scorer receives selected transcript text; trusted plugins run your code. Installation and command reference .
Any agent that lets you set a custom provider address and uses Anthropic Messages, OpenAI Responses, or OpenAI Chat Completions. Claude Code and Codex routing is live-checked; Chat Completions is contract-tested on synthetic histories. File commands read Claude Code and Codex session formats.
Before writing a separate Claude Code or Codex copy, Gobstopper archives the original and prepared bytes. Search the local vault for a missing record, read its saved text, or run gobstopper undo to prepare a restored copy with a new session identity. Keep backups of the vault: recovery depends on your storage.
The proxy changes outgoing requests and leaves session files unchanged. File commands prepare separate copies. Gobstopper does not trigger Claude Code's or Codex's own compaction and does not edit session files in place. Automatic provider compaction stays disabled.
You or the agent can reserve a larger context for a difficult phase, within the model and client's supported window. The budget expires by time or request count. Optional adaptive rescue temporarily raises the threshold after repeated reads of unchanged evidence the proxy removed. Gobstopper also keeps selected original tool results and images through repeated compactions. These controls reduce evidence loss; they cannot guarantee task progress. See context budgets and retention .
Request and compaction deadlines limit stalls, while logging and metrics run in background workers. gobstopper proxy launch --client claude checks readiness before opening Claude Code and can start directly with its official provider when the configuration permits. Codex launch requires a healthy proxy. Existing sessions using a fixed proxy URL still depend on that listener. See recovery and fallback requirements .
Run gobstopper proxy install to start a user service at login and restart it after a process exit. Managed service changes wait for active inference to finish. During inference, Gobstopper requests idle-sleep prevention; closing a lid or forcing sleep can override it. Local session metrics are available through gobstopper data metrics . See startup and local visibility .
Use your agent's built-in /compact when you want a model-written summary without another tool. Gobstopper adds previews and archived originals for saved sessions, plus a proxy that applies local rules without a model call. See the Claude Code comparison or the CliffCompaction comparison .
Excalibur (xcb) xcb.sh Routes coding tasks across the Claude, Codex, and Devin plans you have
ChatGPT Claude Perplexity Grok Gobstopper Context compaction you can undo · MIT or Apache-2.0 · Preview
by Hraness Accept cookies Analytics preferences Your browser remembers your appearance setting and this choice. We use analytics to understand how the site is used. None of it is used for advertising or cross-site tracking. You can change your analytics choice here at any time. Privacy policy Accept analytics Decline analytics
