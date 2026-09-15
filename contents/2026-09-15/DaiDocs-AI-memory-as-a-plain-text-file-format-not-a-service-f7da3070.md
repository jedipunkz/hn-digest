---
source: "https://github.com/Kerneta/daidocs"
hn_url: "https://news.ycombinator.com/item?id=49715672"
title: "DaiDocs: AI memory as a plain-text file format, not a service"
article_title: "GitHub - Kerneta/daidocs: Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% Lo\n[truncated]"
image: "https://repository-images.githubusercontent.com/1368190389/5233a508-a87e-44e2-a8ea-387cbc18977a"
author: "Amin_Rigi"
captured_at: "2026-09-15T17:58:14Z"
capture_tool: "hn-digest"
hn_id: 49715672
score: 3
comments: 1
posted_at: "2026-09-15T17:18:06Z"
tags:
  - hacker-news
---

# DaiDocs: AI memory as a plain-text file format, not a service

- HN: [49715672](https://news.ycombinator.com/item?id=49715672)
- Source: [github.com](https://github.com/Kerneta/daidocs)
- Score: 3
- Comments: 1
- Posted: 2026-09-15T17:18:06Z

## Translation

Title: DaiDocs: AI memory as a plain-text file format, not a service
Article title: GitHub - Kerneta/daidocs: Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% Lo
[truncated]
Description: Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% LongMemEval-S (GPT-4o), 92%
[truncated]

Article text:
GitHub - Kerneta/daidocs: Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% LongMemEval-S (GPT-4o), 92% (Claude Fable 5), 10x fewer tokens. · GitHub
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
Kerneta
/
daidocs
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github assets assets benchmark benchmark docs docs experiments/ recall-sweep experiments/ recall-sweep lib lib prompts prompts readers/ python readers/ python run-artifacts run-artifacts spec spec tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.sha256 MANIFEST.sha256 NOTICE NOTICE QUICKSTART.md QUICKSTART.md QUICKSTART.zh-CN.md QUICKSTART.zh-CN.md README.md README.md README.zh-CN.md README.zh-CN.md RESULTS-ACTORS.md RESULTS-ACTORS.md RESULTS-SUMMARY.md RESULTS-SUMMARY.md RUNBOOK.md RUNBOOK.md SECURITY.md SECURITY.md daidocs.js daidocs.js keyless_check.mjs keyless_check.mjs lock.js lock.js mcp_selftest.mjs mcp_selftest.mjs mcp_server.mjs mcp_server.mjs package-lock.json package-lock.json package.json package.json session_archiver.mjs session_archiver.mjs session_autosave.mjs session_autosave.mjs session_context.mjs session_context.mjs setup.js setup.js verify_surfaces.mjs verify_surfaces.mjs View all files Repository files navigation
An open plain-text format for AI memory (launched Sep 2026).
Your assistant's memory becomes files on your disk that you can open, grep and keep.
Second on the LongMemEval-S leaderboard among memory systems anyone can re-run, 22.40 points above the same model with no memory , reading 10x fewer tokens per question.
Every number here ships with its per-question judge verdicts and a sha256 manifest.
Quickstart ·
The guide ·
Benchmark ·
The format ·
Reproduce it
Launch release V4.4n32, 12 September 2026. What is in it.
If .dai is useful to you, a star helps other people find it.
The whole product, in one loop. Install, build, remember, recall: one store, every model, about a tenth of the tokens. The recording uses invented data.
▶ For higher quality, watch the demo live in your browser
Language independent, model independent
A .dai file is three plain-text zones: a YAML header, a fenced JSON block, and the text. No binary, no database, no SDK required to read it.
Any programming language. The reference engine is Node. A reader in Python, Rust or Go is an afternoon's work, and the spec is normative, written so that two independent implementations agree.
Any model. The store is written once by a cheap observer model and read by whichever model answers. The same store measured with five answering models: 78% to 92%. Change the model, keep the memory.
Any tool. grep , git log , diff , your editor, a shell script. Memory that answers to ordinary tools.
Build a reader in another language and open a PR: that is the contribution that matters most.
What is in this repository: the Kerneta Engine V4.4n that reads and writes .dai files,
the MCP server that connects it to your assistants, and the complete evidence for every number
quoted below: the benchmark run, the judge's verdict on each of the 500 questions, and the
five-model comparison. Each evidence file is hashed in MANIFEST.sha256
so you can check that what is described is what was measured; how to do that is in
docs/PROVENANCE.md .
Second on the leaderboard. One setup for every row: LongMemEval-S, GPT-4o answering, all 500 questions, micro-averaged, and only configurations somebody who does not work for the vendor could re-run. The bottom row is that same GPT-4o with no memory system, reading the whole history pasted into its context: 22.40 points below us . Every figure here carries a caveat and the caveat travels with it, in docs/RESULTS.md .
Not in that table? Graphify, Hindsight, Mem0 and the others publish figures measured on a different answering model, a different denominator or a different benchmark, so they cannot be set beside a GPT-4o 500/500 row in either direction. Every one of them is at daidocs.com/results.html , with what its number actually measures and where ours sits against it.
One memory layer, five answering models, 500 questions each. Retrieval identical for every row (proved by a byte-identical diagnostics file). Details and caveats in RESULTS-ACTORS.md .
Every memory product on the market keeps your history inside its own service and hands it back
through its own API. .dai takes the opposite bet: memory is a file format , the way a
photo is a JPEG. Three plain-text zones per conversation, a small derived index beside them,
and any model, any tool, or grep can read it.
The store is built once by a cheap observer model and read by any actor model. Convert
with a good model, then answer with whatever is cheapest, fastest or local. Numbers below.
One store, connected over MCP, read and written by the tools you already use. node setup.js detects and configures each of these and backs up what it touches; Install has the per-tool commands.
Any MCP-capable runtime, too. The server is a plain stdio MCP server, so frameworks that speak MCP call save_memory and recall_memory with no adapter to write: the OpenAI Agents SDK, the Vercel AI SDK, LangGraph, LangChain, CrewAI and LlamaIndex all consume an MCP server as a tool source. Point them at node mcp_server.mjs .
Bring your history. Claude Code sessions on this machine convert automatically. From any other tool, export a folder of .txt , .md or .jsonl and run node daidocs.js convert . Native history import from more tools is on the roadmap .
npx daidocs setup
One command. It detects Claude Desktop, Claude Code, Cursor, Windsurf, Codex, Cline,
Continue and Zed, configures all of them, installs the session hooks, the reading
protocol and the .dai icon, and backs up every file it touches. On a Claude
subscription there is no API key and nothing to pay.
Want the source and the benchmark artifacts too? Clone it and run setup from there
instead:
git clone https://github.com/Kerneta/daidocs daidocs-app
cd daidocs-app
node setup.js
The clone is named daidocs-app on purpose. git clone would otherwise make a folder
called daidocs , and the default memory store is DaiDocs : on Windows and macOS those
are the same folder, so a clone made from your home directory would land on top of your
own memory. Setup refuses to run from inside the store if it ever happens.
Prefer Python? Read your .dai stores from code, and drive the engine from a
daidocs command:
pip install daidocs
from daidocs import Store
store = Store ( "~/DaiDocs" ) # your memory store
for entry in store . manifest (): # every document
print ( entry [ "id" ], entry [ "title" ])
doc = store . read ( store . ids ()[ 0 ]) # one document, fully parsed
print ( doc [ "understanding" ][ "summary" ])
print ( store . search ( "deploy" )) # find documents by keyword
Two things in one install:
Reader (pure Python, no Node): from daidocs import Store reads the
manifest, any document, and the facts / events / profile indexes.
daidocs command: drives the Node engine, so daidocs setup and
daidocs convert behave like npx daidocs . This needs Node 18+; if Node is
missing it says so and offers to install it. Full guide:
readers/python/ .
setup.js installs the dependencies on its first run and then configures everything.
npm run setup does the same thing, but node setup.js is the one to reach for on
Windows: PowerShell refuses to run npm at all until you change its execution policy,
and node is not affected by that. Each line above is its own command, because Windows
PowerShell 5.1 has no && .
Setup asks nothing. It detects what you have and configures all of it: Claude Desktop,
Claude Code, the session hooks, Cursor, Windsurf, Codex, Cline, Continue, Zed, the reading
protocol and the .dai file icon. It backs up every file it touches.
node setup.js --status what is on, and the command that changes each one
node setup.js --ask choose each surface yourself instead
node setup.js --restore put the machine back exactly as it was
The one thing it never does on its own is convert the history you already have, because
that can run for a while and, with an API key, it spends money. It is one command when
you want it, and it is worth wanting: see Bring what you already have .
Another MCP client? One command each, rather than a config to edit:
node setup.js --client codex (or cursor , windsurf , cline , continue , zed ), and
node setup.js --client generic --config <that client's config file> for anything else.
node setup.js --client list shows the names and where each one keeps its config.
In Claude Code every session saves itself as you work, so there is nothing to remember. On any other connected assistant, say "save this chat to memory" . To give a folder its own project memory, say "make this folder a project" in it, or in a subfolder to make a sub-project.
Most people installing this have months of conversations sitting on disk already. One
command turns them into memory, which is the difference between a store that is useful
this afternoon and one that fills up slowly from here:
node daidocs.js convert
It reads three kinds of history, all the same way:
Claude Code sessions on this machine , from ~/.claude/projects
Sessions already captured but not yet converted , the _pending markers with their
text in _unconverted/
A folder of exports from anywhere else: .txt , .md or .jsonl
It lists what it found with dates, projects and sizes, asks which ones ( all , or
1,3,5-8 ), asks where the store goes, and quotes the worst-case token count and cost
before any paid call . On a Claude subscription that cost is nothing: the assistant in
the session writes the extraction itself. Each item is written the moment it finishes, so
a crash loses at most one document and re-running skips whatever is already converted.
Sessions land in the folder they came from, so a project's history ends up in that
project's own store rather than in one pile.
Every question has a flag, so it scripts:
node daidocs.js convert --source claude --project atlas-api --pick 1-5 --to ~ /DaiDocs --yes
The guide lists every flag.
You do not have to remember to save anything. Once the hook is installed, every
session saves itself:
Every 4,000 new tokens , the session so far is converte

[truncated]

## Original Extract

Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% LongMemEval-S (GPT-4o), 92%
[truncated]

GitHub - Kerneta/daidocs: Open plain-text file format for AI memory. Your assistant's long-term memory as .dai files on your disk: readable by Claude, GPT, Gemini, Cursor, local models and grep (all LLM models work). MCP server + hooks for Claude Code, Claude Desktop, Cursor, Windsurf, Codex. 83% LongMemEval-S (GPT-4o), 92% (Claude Fable 5), 10x fewer tokens. · GitHub
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
Kerneta
/
daidocs
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github assets assets benchmark benchmark docs docs experiments/ recall-sweep experiments/ recall-sweep lib lib prompts prompts readers/ python readers/ python run-artifacts run-artifacts spec spec tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.sha256 MANIFEST.sha256 NOTICE NOTICE QUICKSTART.md QUICKSTART.md QUICKSTART.zh-CN.md QUICKSTART.zh-CN.md README.md README.md README.zh-CN.md README.zh-CN.md RESULTS-ACTORS.md RESULTS-ACTORS.md RESULTS-SUMMARY.md RESULTS-SUMMARY.md RUNBOOK.md RUNBOOK.md SECURITY.md SECURITY.md daidocs.js daidocs.js keyless_check.mjs keyless_check.mjs lock.js lock.js mcp_selftest.mjs mcp_selftest.mjs mcp_server.mjs mcp_server.mjs package-lock.json package-lock.json package.json package.json session_archiver.mjs session_archiver.mjs session_autosave.mjs session_autosave.mjs session_context.mjs session_context.mjs setup.js setup.js verify_surfaces.mjs verify_surfaces.mjs View all files Repository files navigation
An open plain-text format for AI memory (launched Sep 2026).
Your assistant's memory becomes files on your disk that you can open, grep and keep.
Second on the LongMemEval-S leaderboard among memory systems anyone can re-run, 22.40 points above the same model with no memory , reading 10x fewer tokens per question.
Every number here ships with its per-question judge verdicts and a sha256 manifest.
Quickstart ·
The guide ·
Benchmark ·
The format ·
Reproduce it
Launch release V4.4n32, 12 September 2026. What is in it.
If .dai is useful to you, a star helps other people find it.
The whole product, in one loop. Install, build, remember, recall: one store, every model, about a tenth of the tokens. The recording uses invented data.
▶ For higher quality, watch the demo live in your browser
Language independent, model independent
A .dai file is three plain-text zones: a YAML header, a fenced JSON block, and the text. No binary, no database, no SDK required to read it.
Any programming language. The reference engine is Node. A reader in Python, Rust or Go is an afternoon's work, and the spec is normative, written so that two independent implementations agree.
Any model. The store is written once by a cheap observer model and read by whichever model answers. The same store measured with five answering models: 78% to 92%. Change the model, keep the memory.
Any tool. grep , git log , diff , your editor, a shell script. Memory that answers to ordinary tools.
Build a reader in another language and open a PR: that is the contribution that matters most.
What is in this repository: the Kerneta Engine V4.4n that reads and writes .dai files,
the MCP server that connects it to your assistants, and the complete evidence for every number
quoted below: the benchmark run, the judge's verdict on each of the 500 questions, and the
five-model comparison. Each evidence file is hashed in MANIFEST.sha256
so you can check that what is described is what was measured; how to do that is in
docs/PROVENANCE.md .
Second on the leaderboard. One setup for every row: LongMemEval-S, GPT-4o answering, all 500 questions, micro-averaged, and only configurations somebody who does not work for the vendor could re-run. The bottom row is that same GPT-4o with no memory system, reading the whole history pasted into its context: 22.40 points below us . Every figure here carries a caveat and the caveat travels with it, in docs/RESULTS.md .
Not in that table? Graphify, Hindsight, Mem0 and the others publish figures measured on a different answering model, a different denominator or a different benchmark, so they cannot be set beside a GPT-4o 500/500 row in either direction. Every one of them is at daidocs.com/results.html , with what its number actually measures and where ours sits against it.
One memory layer, five answering models, 500 questions each. Retrieval identical for every row (proved by a byte-identical diagnostics file). Details and caveats in RESULTS-ACTORS.md .
Every memory product on the market keeps your history inside its own service and hands it back
through its own API. .dai takes the opposite bet: memory is a file format , the way a
photo is a JPEG. Three plain-text zones per conversation, a small derived index beside them,
and any model, any tool, or grep can read it.
The store is built once by a cheap observer model and read by any actor model. Convert
with a good model, then answer with whatever is cheapest, fastest or local. Numbers below.
One store, connected over MCP, read and written by the tools you already use. node setup.js detects and configures each of these and backs up what it touches; Install has the per-tool commands.
Any MCP-capable runtime, too. The server is a plain stdio MCP server, so frameworks that speak MCP call save_memory and recall_memory with no adapter to write: the OpenAI Agents SDK, the Vercel AI SDK, LangGraph, LangChain, CrewAI and LlamaIndex all consume an MCP server as a tool source. Point them at node mcp_server.mjs .
Bring your history. Claude Code sessions on this machine convert automatically. From any other tool, export a folder of .txt , .md or .jsonl and run node daidocs.js convert . Native history import from more tools is on the roadmap .
npx daidocs setup
One command. It detects Claude Desktop, Claude Code, Cursor, Windsurf, Codex, Cline,
Continue and Zed, configures all of them, installs the session hooks, the reading
protocol and the .dai icon, and backs up every file it touches. On a Claude
subscription there is no API key and nothing to pay.
Want the source and the benchmark artifacts too? Clone it and run setup from there
instead:
git clone https://github.com/Kerneta/daidocs daidocs-app
cd daidocs-app
node setup.js
The clone is named daidocs-app on purpose. git clone would otherwise make a folder
called daidocs , and the default memory store is DaiDocs : on Windows and macOS those
are the same folder, so a clone made from your home directory would land on top of your
own memory. Setup refuses to run from inside the store if it ever happens.
Prefer Python? Read your .dai stores from code, and drive the engine from a
daidocs command:
pip install daidocs
from daidocs import Store
store = Store ( "~/DaiDocs" ) # your memory store
for entry in store . manifest (): # every document
print ( entry [ "id" ], entry [ "title" ])
doc = store . read ( store . ids ()[ 0 ]) # one document, fully parsed
print ( doc [ "understanding" ][ "summary" ])
print ( store . search ( "deploy" )) # find documents by keyword
Two things in one install:
Reader (pure Python, no Node): from daidocs import Store reads the
manifest, any document, and the facts / events / profile indexes.
daidocs command: drives the Node engine, so daidocs setup and
daidocs convert behave like npx daidocs . This needs Node 18+; if Node is
missing it says so and offers to install it. Full guide:
readers/python/ .
setup.js installs the dependencies on its first run and then configures everything.
npm run setup does the same thing, but node setup.js is the one to reach for on
Windows: PowerShell refuses to run npm at all until you change its execution policy,
and node is not affected by that. Each line above is its own command, because Windows
PowerShell 5.1 has no && .
Setup asks nothing. It detects what you have and configures all of it: Claude Desktop,
Claude Code, the session hooks, Cursor, Windsurf, Codex, Cline, Continue, Zed, the reading
protocol and the .dai file icon. It backs up every file it touches.
node setup.js --status what is on, and the command that changes each one
node setup.js --ask choose each surface yourself instead
node setup.js --restore put the machine back exactly as it was
The one thing it never does on its own is convert the history you already have, because
that can run for a while and, with an API key, it spends money. It is one command when
you want it, and it is worth wanting: see Bring what you already have .
Another MCP client? One command each, rather than a config to edit:
node setup.js --client codex (or cursor , windsurf , cline , continue , zed ), and
node setup.js --client generic --config <that client's config file> for anything else.
node setup.js --client list shows the names and where each one keeps its config.
In Claude Code every session saves itself as you work, so there is nothing to remember. On any other connected assistant, say "save this chat to memory" . To give a folder its own project memory, say "make this folder a project" in it, or in a subfolder to make a sub-project.
Most people installing this have months of conversations sitting on disk already. One
command turns them into memory, which is the difference between a store that is useful
this afternoon and one that fills up slowly from here:
node daidocs.js convert
It reads three kinds of history, all the same way:
Claude Code sessions on this machine , from ~/.claude/projects
Sessions already captured but not yet converted , the _pending markers with their
text in _unconverted/
A folder of exports from anywhere else: .txt , .md or .jsonl
It lists what it found with dates, projects and sizes, asks which ones ( all , or
1,3,5-8 ), asks where the store goes, and quotes the worst-case token count and cost
before any paid call . On a Claude subscription that cost is nothing: the assistant in
the session writes the extraction itself. Each item is written the moment it finishes, so
a crash loses at most one document and re-running skips whatever is already converted.
Sessions land in the folder they came from, so a project's history ends up in that
project's own store rather than in one pile.
Every question has a flag, so it scripts:
node daidocs.js convert --source claude --project atlas-api --pick 1-5 --to ~ /DaiDocs --yes
The guide lists every flag.
You do not have to remember to save anything. Once the hook is installed, every
session saves itself:
Every 4,000 new tokens , the session so far is converte

[truncated]
