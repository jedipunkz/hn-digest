---
source: "https://github.com/unrealumanga/Nedster"
hn_url: "https://news.ycombinator.com/item?id=50091652"
title: "Nedster – An open-source local AI"
article_title: "GitHub - unrealumanga/Nedster: Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. · GitHub"
image: "https://opengraph.githubassets.com/e8746b62efa58aaa79f86731f8ce265d83a7db1d6c392f3435bd517051fb9f39/unrealumanga/Nedster"
author: "unrealumanga"
captured_at: "2026-10-14T18:20:41Z"
capture_tool: "hn-digest"
hn_id: 50091652
score: 1
comments: 0
posted_at: "2026-10-14T17:37:41Z"
tags:
  - hacker-news
---

# Nedster – An open-source local AI

- HN: [50091652](https://news.ycombinator.com/item?id=50091652)
- Source: [github.com](https://github.com/unrealumanga/Nedster)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T17:37:41Z

## Translation

Title: Nedster – An open-source local AI
Article title: GitHub - unrealumanga/Nedster: Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. · GitHub
Description: Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. - unrealumanga/Nedster

Article text:
GitHub - unrealumanga/Nedster: Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. · GitHub
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
unrealumanga
Nedster Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI.
Readme MIT license Security policy
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
88 Commits 88 Commits Folders and files
.github/ workflows .github/ workflows benchmarks benchmarks docs docs evals evals ironclad ironclad nedster_agent nedster_agent tests tests .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh install_global.sh install_global.sh pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A local-first coding agent for Ollama. One native tool-calling loop reads, edits, searches and runs your
project's tests. What happened is decided from the tool records, never from what the model says about itself.
Windows, macOS and Linux.
Status: 4.0 alpha. The v3 engine is gone; see Upgrading from v3 .
No fake success. A "fix X" request that changed nothing ends INCOMPLETE , not "done"; "already done" is
reported as a claim backed by the files it cites. 16/16 eval fixtures pass with 0 false successes
( Evals ); the v3 engine scored 1/8 with one false success on the same tasks.
Auto-verify. If the model edits files but skips the tests, Nedster runs them itself and gives the model one
repair turn when they fail.
Undo and sessions. /undo restores the last turn's file edits (refusing if you changed them since);
nedster sessions and /resume bring back earlier conversations.
Long tasks. The context budget follows the model's real window; old history is compacted deterministically
from tool records. Append-only history keeps Ollama's prompt cache warm (measured ~47x faster prefill at 56K tokens).
Faster, cheaper edits. edit_file applies several edits to a file in one atomic call and returns the new hash.
Safer execution. Commands run only with your consent, in Docker / bubblewrap / Seatbelt, or on the host
when you opt in (Windows process trees are always contained in a Job Object, and fail closed).
Local only. Inference goes to your Ollama server. No account, no cloud fallback.
One loop. The model gets a small set of typed tools ( read_file , edit_file , write_file , list_files ,
search_text , git_status/diff/log/show , run_tests , run_command , show_output , declare_no_change ).
The runtime validates arguments, enforces policy, runs the tool and records the result.
Outcomes come from evidence. A run ends in exactly one of:
Edits are guarded. edit_file needs the file's current sha256 from a read; a stale hash is refused. Protected
paths ( .git , .nedster ) cannot be written. Writes are atomic.
Requirements: Windows, macOS or Linux; Python 3.11+; Ollama (setup can install it).
About 8 GB RAM minimum; a GPU is recommended, not required.
# Windows
git clone https: // github.com / unrealumanga / Nedster.git
cd Nedster
powershell - ExecutionPolicy Bypass - File .\install.ps1 # add -SkipSetup to install only the launcher
# macOS / Linux
git clone https://github.com/unrealumanga/Nedster.git
cd Nedster
chmod +x install.sh && ./install.sh
Both installers put a global nedster launcher on PATH (via uv tool install if uv exists, otherwise
pip install --user ) and then run nedster setup .
nedster setup detects hardware, picks a model plan ( --profile fast|balanced|maximum ), keeps or installs Ollama,
downloads the models, creates the nedster-worker / nedster-router aliases, verifies a generation, and writes
~/.nedster/config.toml (the old file is backed up). It asks whether commands may run (see
Execution consent ). Useful flags: --dry-run (change nothing), --yes ,
--no-input , --with-embeddings , --skip-ollama-install , --ollama-url .
nedster doctor # inference-ready / read-only-ready / execution-ready, plus context + KV-cache advice
cd my-project
nedster # opens the TUI in this folder
Default model plans ( nedster models recommend shows yours):
Execution consent and sandboxes
Reading, searching and editing files inside the project need no consent. Running commands and tests is off
until you allow it : the default is policy.max_auto_risk = "write" . nedster setup asks, or pass
--allow-execution (run only inside an isolating sandbox) or --allow-host-execution (run directly on this
machine, no isolation). A project's own .nedster/config.toml can only tighten these settings, never loosen them.
Backend order with sandbox.backend = "auto" : bwrap , seatbelt , docker , then host only if
sandbox.allow_host_execution = true . With none ready, execution is refused and the message says how to fix it.
An explicitly configured backend that is not ready is never replaced silently.
On Windows there is no native filesystem sandbox: use Docker for isolation, or opt in to host , where the Job Object
only contains the process tree (it also kills the tree if Nedster dies), not file or network access.
Programs are resolved from trusted PATH entries outside the project; a binary planted in the repo is never run.
Tests you run are project code, so a hostile repo can still print a fake "all passed" line; the runner and its
last summary line are parsed, which stops trivial spoofing but is not a proof. See SECURITY.md .
nedster (or nedster tui ) opens a chat-first terminal UI for the current folder ( --project DIR to pick another).
You can keep typing while a run is busy; messages queue. Enter sends, Shift+Enter adds a line, @ completes a file.
Slash commands: /help , /new , /clear , /sessions , /resume [id] , /model [name] , /status , /diff ,
/undo , /doctor , /models , /permissions , /quit . /new , /resume , /model , /undo , /sessions work only
while idle.
/resume replays the stored transcript for display only . Nothing is re-executed, and earlier test results
are not treated as current evidence. A session held by another live Nedster window cannot be resumed.
/undo reverts the last turn that edited files (one step; undoing twice is not supported).
Every file edit made through write_file / edit_file journals the file's pre-image before changing it. /undo
restores those files only if each is still exactly as the turn left it, and refuses (restoring nothing) if you or
something else changed any of them since. Commands are not undone : anything run_command or run_tests did
(files created by a build, installs, git operations) is outside the journal, and undo says so when a turn ran
commands. Files the edit tools created are removed on undo; side effects of commands are not.
nedster run " Find the failing tests, fix the cause, and verify the suite "
nedster run --expect answer " Where is the retry logic implemented? "
echo " explain src/cache.py " | nedster run # objective from stdin
nedster run --json " ... " # one JSON document
nedster run --jsonl " ... " # one EngineEvent per line
nedster run --continue " now add a test for it " # latest session in this project
nedster run --session < id-prefix > " ... " # a specific session
--expect change|answer states what success requires. Without it, the objective is classified (change /
question / unknown). A change request that made no change ends INCOMPLETE , not ANSWERED .
--model NAME overrides the worker model. --engine native is deprecated and does nothing.
Piped or redirected output never opens the full-screen UI. Global flags: --plain , --json , --no-tui ,
--color / --no-color , --ascii , --project , --config .
If the model changes files and then stops without running tests, the engine runs run_tests itself (up to 2
times per run). If they fail the model gets another turn to repair. The result is a real tool record, so this can only
move UNVALIDATED to COMPLETED or INCOMPLETE . It needs execution consent and a recognized test runner
(pytest, unittest, npm, node --test, cargo, go).
Sessions are stored per project (SQLite, under ~/.nedster/data ).
nedster sessions # list, newest first (also: sessions list)
nedster sessions show < id > # metadata + stored transcript (id prefixes work)
nedster sessions delete < id > # refused while another process holds the session
nedster sessions prune --older-than 30 # delete old tool outputs nothing references
Also: nedster doctor , nedster models [list|recommend|install] , nedster permissions ,
nedster config show , nedster verify-ledger (checks the old HMAC event chain; the native engine writes no
events to it). Each takes --plain / --json .
Context budget. runtime.context_tokens = 0 (default) uses the worker model's num_ctx as reported by
Ollama, falling back to 32768. Set it explicitly to override; nedster doctor warns if it differs from num_ctx .
Deterministic compaction. At ~70% of the window old tool outputs are replaced by stubs; the full text stays
available through show_output . If still too large, the oldest complete call/result groups are replaced by one
structured summary built from tool records (reads, changes with sha, validations fresh/stale, failures), aiming
at 50% of the window. Your messages are never dropped. If a request cannot fit even then, the run ends
INCOMPLETE ("context budget exhausted") without sending an oversized request.
KV cache. History is append-only between compactions, so Ollama reuses the cached prefix. Measured on an RTX
5090 with nedster-worker (qwen35moe, num_ctx 65536, Ollama 0.40.4), ~57K-token context: the next turn's prefill
took 0.21 s against 9.9 s cold. Rewriting earlier history costs a full re-prefill (5.3 s at 33K tokens), which
is why compaction is rare and aggressive. Re-sending the model's thinking cost ~3 s per turn, so Nedster never
sends it back.
Advice. Set OLLAMA_FLASH_ATTENTION=1 and OLLAMA_KV_CACHE_TYPE=q8_0 in the Ollama server's environment and
restart it: KV memory roughly halves, so the same VRAM holds about twice the context. nedster doctor prints this
when the values are not set in your shell (it can only see its own environment).
16 offline toy-repo fixtures in evals/toyrepos/ (bug fixes, rename, multi-file feature, questions, an
already-done request, prompt injection in a README and in a code comment, a 40-file search, unicode paths, a JS
project). Each has a check script that decides pass/fail independent of the agent. Run:
python -m evals.run --model nedster-worker [--only a_add_subtract,l_already_done] [--repeat N] [--timeout S] [--out

[truncated]

## Original Extract

Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. - unrealumanga/Nedster

GitHub - unrealumanga/Nedster: Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI. · GitHub
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
unrealumanga
Nedster Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Local-first AI coding and automation agent with automatic Ollama setup, hardware-aware models, RAG, MCP, and a friendly terminal UI.
Readme MIT license Security policy
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
88 Commits 88 Commits Folders and files
.github/ workflows .github/ workflows benchmarks benchmarks docs docs evals evals ironclad ironclad nedster_agent nedster_agent tests tests .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh install_global.sh install_global.sh pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
A local-first coding agent for Ollama. One native tool-calling loop reads, edits, searches and runs your
project's tests. What happened is decided from the tool records, never from what the model says about itself.
Windows, macOS and Linux.
Status: 4.0 alpha. The v3 engine is gone; see Upgrading from v3 .
No fake success. A "fix X" request that changed nothing ends INCOMPLETE , not "done"; "already done" is
reported as a claim backed by the files it cites. 16/16 eval fixtures pass with 0 false successes
( Evals ); the v3 engine scored 1/8 with one false success on the same tasks.
Auto-verify. If the model edits files but skips the tests, Nedster runs them itself and gives the model one
repair turn when they fail.
Undo and sessions. /undo restores the last turn's file edits (refusing if you changed them since);
nedster sessions and /resume bring back earlier conversations.
Long tasks. The context budget follows the model's real window; old history is compacted deterministically
from tool records. Append-only history keeps Ollama's prompt cache warm (measured ~47x faster prefill at 56K tokens).
Faster, cheaper edits. edit_file applies several edits to a file in one atomic call and returns the new hash.
Safer execution. Commands run only with your consent, in Docker / bubblewrap / Seatbelt, or on the host
when you opt in (Windows process trees are always contained in a Job Object, and fail closed).
Local only. Inference goes to your Ollama server. No account, no cloud fallback.
One loop. The model gets a small set of typed tools ( read_file , edit_file , write_file , list_files ,
search_text , git_status/diff/log/show , run_tests , run_command , show_output , declare_no_change ).
The runtime validates arguments, enforces policy, runs the tool and records the result.
Outcomes come from evidence. A run ends in exactly one of:
Edits are guarded. edit_file needs the file's current sha256 from a read; a stale hash is refused. Protected
paths ( .git , .nedster ) cannot be written. Writes are atomic.
Requirements: Windows, macOS or Linux; Python 3.11+; Ollama (setup can install it).
About 8 GB RAM minimum; a GPU is recommended, not required.
# Windows
git clone https: // github.com / unrealumanga / Nedster.git
cd Nedster
powershell - ExecutionPolicy Bypass - File .\install.ps1 # add -SkipSetup to install only the launcher
# macOS / Linux
git clone https://github.com/unrealumanga/Nedster.git
cd Nedster
chmod +x install.sh && ./install.sh
Both installers put a global nedster launcher on PATH (via uv tool install if uv exists, otherwise
pip install --user ) and then run nedster setup .
nedster setup detects hardware, picks a model plan ( --profile fast|balanced|maximum ), keeps or installs Ollama,
downloads the models, creates the nedster-worker / nedster-router aliases, verifies a generation, and writes
~/.nedster/config.toml (the old file is backed up). It asks whether commands may run (see
Execution consent ). Useful flags: --dry-run (change nothing), --yes ,
--no-input , --with-embeddings , --skip-ollama-install , --ollama-url .
nedster doctor # inference-ready / read-only-ready / execution-ready, plus context + KV-cache advice
cd my-project
nedster # opens the TUI in this folder
Default model plans ( nedster models recommend shows yours):
Execution consent and sandboxes
Reading, searching and editing files inside the project need no consent. Running commands and tests is off
until you allow it : the default is policy.max_auto_risk = "write" . nedster setup asks, or pass
--allow-execution (run only inside an isolating sandbox) or --allow-host-execution (run directly on this
machine, no isolation). A project's own .nedster/config.toml can only tighten these settings, never loosen them.
Backend order with sandbox.backend = "auto" : bwrap , seatbelt , docker , then host only if
sandbox.allow_host_execution = true . With none ready, execution is refused and the message says how to fix it.
An explicitly configured backend that is not ready is never replaced silently.
On Windows there is no native filesystem sandbox: use Docker for isolation, or opt in to host , where the Job Object
only contains the process tree (it also kills the tree if Nedster dies), not file or network access.
Programs are resolved from trusted PATH entries outside the project; a binary planted in the repo is never run.
Tests you run are project code, so a hostile repo can still print a fake "all passed" line; the runner and its
last summary line are parsed, which stops trivial spoofing but is not a proof. See SECURITY.md .
nedster (or nedster tui ) opens a chat-first terminal UI for the current folder ( --project DIR to pick another).
You can keep typing while a run is busy; messages queue. Enter sends, Shift+Enter adds a line, @ completes a file.
Slash commands: /help , /new , /clear , /sessions , /resume [id] , /model [name] , /status , /diff ,
/undo , /doctor , /models , /permissions , /quit . /new , /resume , /model , /undo , /sessions work only
while idle.
/resume replays the stored transcript for display only . Nothing is re-executed, and earlier test results
are not treated as current evidence. A session held by another live Nedster window cannot be resumed.
/undo reverts the last turn that edited files (one step; undoing twice is not supported).
Every file edit made through write_file / edit_file journals the file's pre-image before changing it. /undo
restores those files only if each is still exactly as the turn left it, and refuses (restoring nothing) if you or
something else changed any of them since. Commands are not undone : anything run_command or run_tests did
(files created by a build, installs, git operations) is outside the journal, and undo says so when a turn ran
commands. Files the edit tools created are removed on undo; side effects of commands are not.
nedster run " Find the failing tests, fix the cause, and verify the suite "
nedster run --expect answer " Where is the retry logic implemented? "
echo " explain src/cache.py " | nedster run # objective from stdin
nedster run --json " ... " # one JSON document
nedster run --jsonl " ... " # one EngineEvent per line
nedster run --continue " now add a test for it " # latest session in this project
nedster run --session < id-prefix > " ... " # a specific session
--expect change|answer states what success requires. Without it, the objective is classified (change /
question / unknown). A change request that made no change ends INCOMPLETE , not ANSWERED .
--model NAME overrides the worker model. --engine native is deprecated and does nothing.
Piped or redirected output never opens the full-screen UI. Global flags: --plain , --json , --no-tui ,
--color / --no-color , --ascii , --project , --config .
If the model changes files and then stops without running tests, the engine runs run_tests itself (up to 2
times per run). If they fail the model gets another turn to repair. The result is a real tool record, so this can only
move UNVALIDATED to COMPLETED or INCOMPLETE . It needs execution consent and a recognized test runner
(pytest, unittest, npm, node --test, cargo, go).
Sessions are stored per project (SQLite, under ~/.nedster/data ).
nedster sessions # list, newest first (also: sessions list)
nedster sessions show < id > # metadata + stored transcript (id prefixes work)
nedster sessions delete < id > # refused while another process holds the session
nedster sessions prune --older-than 30 # delete old tool outputs nothing references
Also: nedster doctor , nedster models [list|recommend|install] , nedster permissions ,
nedster config show , nedster verify-ledger (checks the old HMAC event chain; the native engine writes no
events to it). Each takes --plain / --json .
Context budget. runtime.context_tokens = 0 (default) uses the worker model's num_ctx as reported by
Ollama, falling back to 32768. Set it explicitly to override; nedster doctor warns if it differs from num_ctx .
Deterministic compaction. At ~70% of the window old tool outputs are replaced by stubs; the full text stays
available through show_output . If still too large, the oldest complete call/result groups are replaced by one
structured summary built from tool records (reads, changes with sha, validations fresh/stale, failures), aiming
at 50% of the window. Your messages are never dropped. If a request cannot fit even then, the run ends
INCOMPLETE ("context budget exhausted") without sending an oversized request.
KV cache. History is append-only between compactions, so Ollama reuses the cached prefix. Measured on an RTX
5090 with nedster-worker (qwen35moe, num_ctx 65536, Ollama 0.40.4), ~57K-token context: the next turn's prefill
took 0.21 s against 9.9 s cold. Rewriting earlier history costs a full re-prefill (5.3 s at 33K tokens), which
is why compaction is rare and aggressive. Re-sending the model's thinking cost ~3 s per turn, so Nedster never
sends it back.
Advice. Set OLLAMA_FLASH_ATTENTION=1 and OLLAMA_KV_CACHE_TYPE=q8_0 in the Ollama server's environment and
restart it: KV memory roughly halves, so the same VRAM holds about twice the context. nedster doctor prints this
when the values are not set in your shell (it can only see its own environment).
16 offline toy-repo fixtures in evals/toyrepos/ (bug fixes, rename, multi-file feature, questions, an
already-done request, prompt injection in a README and in a code comment, a 40-file search, unicode paths, a JS
project). Each has a check script that decides pass/fail independent of the agent. Run:
python -m evals.run --model nedster-worker [--only a_add_subtract,l_already_done] [--repeat N] [--timeout S] [--out

[truncated]
