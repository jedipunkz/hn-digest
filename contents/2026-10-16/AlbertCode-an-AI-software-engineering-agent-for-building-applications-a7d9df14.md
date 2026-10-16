---
source: "https://www.albertcode.dev"
hn_url: "https://news.ycombinator.com/item?id=50116533"
title: "AlbertCode, an AI software engineering agent for building applications"
article_title: "AlbertCode SWE"
image: "https://www.albertcode.dev/opengraph-image?c7fc71a19567d7da"
author: "huynvic"
captured_at: "2026-10-16T10:50:19Z"
capture_tool: "hn-digest"
hn_id: 50116533
score: 2
comments: 2
posted_at: "2026-10-16T10:16:12Z"
tags:
  - hacker-news
---

# AlbertCode, an AI software engineering agent for building applications

- HN: [50116533](https://news.ycombinator.com/item?id=50116533)
- Source: [www.albertcode.dev](https://www.albertcode.dev)
- Score: 2
- Comments: 2
- Posted: 2026-10-16T10:16:12Z

## Translation

Title: AlbertCode, an AI software engineering agent for building applications
Article title: AlbertCode SWE
Description: AlbertCode SWE Agent is an AI coding agent that plans a change, builds it in an isolated copy of your repository, runs your own tests, and waits for your approval before anything reaches your files.

Article text:
AlbertCode SWE Skip to content Albert Code SWE Product
Start Building AlbertCode From idea to working software Software engineering, accelerated by AI.
Meet AlbertCode SWE Agent. A free-to-install AI coding agent, with open installers and documentation, built to help you go from ideas to working code.
Read the docs macOS · Windows · Linux / no account / your model, or one running locally
An agent that engineers the way you review.
AlbertCode takes a request in plain words and returns a change you can trust enough to accept: planned in the open, built away from your working tree, proven by your own tests, and applied only when you say so.
Every change starts as a plan you can read: the goal, the files it expects to touch, how it will test the change, and the risk.
The work happens in an isolated copy of your repository. Your own tests run there, and independent security and quality reviews read the result.
Two human gates — the plan and the diff. Nothing reaches your working tree until you accept it, and Git stays your undo.
One request, from plan to proof.
Step through a governed change exactly as the agent shows it. Every screen is a capture of the product from its README — not a mock-up.
You describe the change in plain words. AlbertCode reads the repository and proposes a plan: the files it will change, its steps, how it will be tested, a security note, how to undo it and its risk. Nothing has been written yet.
You describe the change in plain words. AlbertCode reads the repository and proposes a plan: the files it will change, its steps, how it will be tested, a security note, how to undo it and its risk. Nothing has been written yet.
Captures of the AlbertCode SWE Agent 1.54 interface from its README. The model's answers were scripted for the pictures.
One agent, in your terminal, your browser and VS Code.
Every interface talks to the same engine — the planning, the isolated copy, the checks, the reviews and the two gates — so a task started in one can be followed in another.
The coding agent that proves its work.
A coding agent that runs on your machine. It plans a change, builds it in an isolated copy of your repository, runs your own tests, reviews the result, and shows you the diff. Nothing reaches your files until you accept it.
Two clear decision points — the plan and the diff — so nothing reaches your repository by surprise.
Runs your own toolchain (pytest, npm scripts, go test, cargo, Maven, Gradle, dotnet) and shows what passed and what didn't.
Terminal, browser and VS Code, all driving the same local service and the same sessions.
A hosted provider, OpenRouter, Hugging Face, any compatible API, or a model on your own machine.
The task board: every task by where it stands, and what it waits for. Also on the AlbertCode engine
mini-AlbertCode SWE Agent , a second agent with a full-screen terminal and a browser workspace, without the VS Code panel.
Two gates. Nothing crosses them without you.
This is Governed mode, the default. Circles are work the agent does; diamonds are the two points where a person decides.
The goal, the files it expects to touch, the test plan and the risk. Nothing is written.
You approve the plan, ask for changes, or reject it.
The change is made in an isolated copy of your repository — never in your working tree.
Your project's own tests and checks run, and the change is reviewed for security and quality.
You read the diff and the results, then apply it, ask for changes, or discard it.
Everything runs on your machine.
AlbertCode runs as a small service on your own computer, with the terminal, the browser workspace and the VS Code panel on top. A task started in one can be followed in another, and closing a window never stops the work.
Rejecting or discarding a change deletes its isolated copy: your working tree never saw it.
Saved keys stay in your system's keychain or keyring, and are never handed to the commands a task runs.
Accepted changes are ordinary edits in your working tree: review them with git diff, commit them, undo them with Git.
Your computer Terminal Full-screen conversation, gates and lists Browser workspace Plans, diffs, files and your running app VS Code The same agent and gates, in your editor same sessions Local service on 127.0.0.1 Starts on first use and keeps running between windows. Answers only requests carrying this installation's access token. builds in Isolated copy of your repository Where the change is built, your tests run and the reviews read it. gate 2 · you accept Your working tree Changes land only when you accept. Commit or undo them with Git. parts a task needs The model you connect A hosted provider with your key, or a local model server — then nothing leaves your machine. No telemetry, no account. Drawn from the agents' documentation (How it works, Security model). Both run only on your machine; neither is a hosted service. Developer workflow
Built for the whole loop, not just the edit.
Each capability below is described in AlbertCode SWE Agent's own documentation.
Choose the pace per request: two gates for changes you care about, direct edits you watch, or questions that change nothing.
pytest, npm scripts, go test, cargo, Maven, Gradle, dotnet — the project's own checks decide what passed.
Run the project in a preview panel at desktop, tablet or phone width, with its errors surfaced and one click to ask for a fix.
Add MCP servers from a gallery of connectors, or your own by hand — local or remote, with sign-in.
albertcode run sends one request without the interface and answers only the gates you name with --approve.
Hosted providers, OpenRouter, Hugging Face, any compatible API — or Ollama, LM Studio and friends, so nothing leaves your machine.
The same agent and gates inside your editor, from the .vsix extension attached to each release.
albertcode.contract.toml bounds what may change; every action lands in a hash-chained ledger you can export.
Requirements, architecture, browser tests
A spec as a checklist backed by browser tests, what your app is made of read from its own files, and Playwright tests in your e2e/ folder.
Open about everything — including what isn't open.
AlbertCode SWE Agent is a proprietary, free-to-install application. Everything you need to adopt, inspect and improve it lives in a public repository you can read, fork and contribute to.
The one-line installers are plain shell and PowerShell, Apache-2.0, and verify each download's SHA-256 before running it.
Every guide, from installation to the contract format, is public and open to corrections.
Benchmarks with their evidence
A published methodology, per-task outcomes, and results labelled self-reported until someone independent checks them.
No telemetry, no account. Your code goes only to the model you choose — or nowhere, with a local model.
One command. Then describe the change.
The installer shows what it will do and asks before it changes anything. It installs for your user only — Python included — so there's no pip, virtual environment or PATH editing.
Paste the command for your system and answer Y.
Paste an API key, or point it at a model on your machine.
cd into a project, start the agent, and say what should change.
Copy $ curl -fsSL https://raw.githubusercontent.com/ huynvic/ albertcode-swe-agent/ main/ installer/ install.sh | sh Open Terminal (⌘ Space, then type Terminal), paste the command, read the plan, then type Y.
Open a new Terminal window when it finishes, then run albertcode --version.
More in the documentation, and in each repository's own FAQ.
AlbertCode SWE Agent is an AI coding agent that runs on your machine. It plans a change, builds it in an isolated copy of your repository, runs your own tests, reviews the result, and waits for you to accept the diff before anything reaches your files. Is it open source? No — and we'd rather say so plainly. AlbertCode SWE Agent is a proprietary application. Its public GitHub repository holds the installers, the documentation, examples and the benchmark methodology, under Apache-2.0 (code) and CC BY 4.0 (docs), and welcome contributions. The engine's source isn't published. What's open, and what isn't . Does it cost anything? Installing is free and needs no account. AlbertCode doesn't charge for model usage: you use your own model provider's account, or a local model at no cost. The repository notes that terms for AlbertCode itself will be published with its first public release. Does my code leave my machine? Only to the model provider you connect (and to MCP servers you add), because the model has to read the code it works on. With a local model, nothing leaves your machine. There is no telemetry and no account. Saved keys are kept by your operating system's keychain or keyring. Which models can I use? Hosted providers with an API key, Kimi, Hugging Face, OpenRouter, any OpenAI-compatible service, or a model on your own machine through Ollama, LM Studio, vLLM or llama.cpp. The model must support tool calling; albertcode models test checks one. Can it change my files without asking? Not in Governed mode, the default: you approve the plan, then the diff. Direct mode works in your real files and asks before each write, unless you explicitly choose auto-approve. Direct needs Git, because Git is the undo. Can I run it in the browser, here? No. AlbertCode runs on your own computer; it isn't offered as a hosted service, and this website never runs code for you. Its own browser workspace opens on your computer with albertcode web . The run planner here prepares the exact c
[truncated]
From idea to working software.
Install in a minute or two, connect the model you already use, and let the first change come back with its proof.
Read the docs Albert Code SWE From idea to working software. AI software engineering agents that plan, build in isolation, verify with your own tests, and wait for your approval.
Free to install · No account · No telemetry
Discussions (opens in a new tab)
Report an issue (opens in a new tab)
Report a vulnerability (opens in a new tab)
Support AlbertCode (opens in a new tab)
Licensing (opens in a new tab)
© 2026 Huy Nguyen . AlbertCode SWE Agent and mini-AlbertCode SWE Agent are proprietary software, free to install. Their public repositories' code is Apache-2.0 and documentation CC BY 4.0.
IBM Plex typefaces under the SIL Open Font License.

## Original Extract

AlbertCode SWE Agent is an AI coding agent that plans a change, builds it in an isolated copy of your repository, runs your own tests, and waits for your approval before anything reaches your files.

AlbertCode SWE Skip to content Albert Code SWE Product
Start Building AlbertCode From idea to working software Software engineering, accelerated by AI.
Meet AlbertCode SWE Agent. A free-to-install AI coding agent, with open installers and documentation, built to help you go from ideas to working code.
Read the docs macOS · Windows · Linux / no account / your model, or one running locally
An agent that engineers the way you review.
AlbertCode takes a request in plain words and returns a change you can trust enough to accept: planned in the open, built away from your working tree, proven by your own tests, and applied only when you say so.
Every change starts as a plan you can read: the goal, the files it expects to touch, how it will test the change, and the risk.
The work happens in an isolated copy of your repository. Your own tests run there, and independent security and quality reviews read the result.
Two human gates — the plan and the diff. Nothing reaches your working tree until you accept it, and Git stays your undo.
One request, from plan to proof.
Step through a governed change exactly as the agent shows it. Every screen is a capture of the product from its README — not a mock-up.
You describe the change in plain words. AlbertCode reads the repository and proposes a plan: the files it will change, its steps, how it will be tested, a security note, how to undo it and its risk. Nothing has been written yet.
You describe the change in plain words. AlbertCode reads the repository and proposes a plan: the files it will change, its steps, how it will be tested, a security note, how to undo it and its risk. Nothing has been written yet.
Captures of the AlbertCode SWE Agent 1.54 interface from its README. The model's answers were scripted for the pictures.
One agent, in your terminal, your browser and VS Code.
Every interface talks to the same engine — the planning, the isolated copy, the checks, the reviews and the two gates — so a task started in one can be followed in another.
The coding agent that proves its work.
A coding agent that runs on your machine. It plans a change, builds it in an isolated copy of your repository, runs your own tests, reviews the result, and shows you the diff. Nothing reaches your files until you accept it.
Two clear decision points — the plan and the diff — so nothing reaches your repository by surprise.
Runs your own toolchain (pytest, npm scripts, go test, cargo, Maven, Gradle, dotnet) and shows what passed and what didn't.
Terminal, browser and VS Code, all driving the same local service and the same sessions.
A hosted provider, OpenRouter, Hugging Face, any compatible API, or a model on your own machine.
The task board: every task by where it stands, and what it waits for. Also on the AlbertCode engine
mini-AlbertCode SWE Agent , a second agent with a full-screen terminal and a browser workspace, without the VS Code panel.
Two gates. Nothing crosses them without you.
This is Governed mode, the default. Circles are work the agent does; diamonds are the two points where a person decides.
The goal, the files it expects to touch, the test plan and the risk. Nothing is written.
You approve the plan, ask for changes, or reject it.
The change is made in an isolated copy of your repository — never in your working tree.
Your project's own tests and checks run, and the change is reviewed for security and quality.
You read the diff and the results, then apply it, ask for changes, or discard it.
Everything runs on your machine.
AlbertCode runs as a small service on your own computer, with the terminal, the browser workspace and the VS Code panel on top. A task started in one can be followed in another, and closing a window never stops the work.
Rejecting or discarding a change deletes its isolated copy: your working tree never saw it.
Saved keys stay in your system's keychain or keyring, and are never handed to the commands a task runs.
Accepted changes are ordinary edits in your working tree: review them with git diff, commit them, undo them with Git.
Your computer Terminal Full-screen conversation, gates and lists Browser workspace Plans, diffs, files and your running app VS Code The same agent and gates, in your editor same sessions Local service on 127.0.0.1 Starts on first use and keeps running between windows. Answers only requests carrying this installation's access token. builds in Isolated copy of your repository Where the change is built, your tests run and the reviews read it. gate 2 · you accept Your working tree Changes land only when you accept. Commit or undo them with Git. parts a task needs The model you connect A hosted provider with your key, or a local model server — then nothing leaves your machine. No telemetry, no account. Drawn from the agents' documentation (How it works, Security model). Both run only on your machine; neither is a hosted service. Developer workflow
Built for the whole loop, not just the edit.
Each capability below is described in AlbertCode SWE Agent's own documentation.
Choose the pace per request: two gates for changes you care about, direct edits you watch, or questions that change nothing.
pytest, npm scripts, go test, cargo, Maven, Gradle, dotnet — the project's own checks decide what passed.
Run the project in a preview panel at desktop, tablet or phone width, with its errors surfaced and one click to ask for a fix.
Add MCP servers from a gallery of connectors, or your own by hand — local or remote, with sign-in.
albertcode run sends one request without the interface and answers only the gates you name with --approve.
Hosted providers, OpenRouter, Hugging Face, any compatible API — or Ollama, LM Studio and friends, so nothing leaves your machine.
The same agent and gates inside your editor, from the .vsix extension attached to each release.
albertcode.contract.toml bounds what may change; every action lands in a hash-chained ledger you can export.
Requirements, architecture, browser tests
A spec as a checklist backed by browser tests, what your app is made of read from its own files, and Playwright tests in your e2e/ folder.
Open about everything — including what isn't open.
AlbertCode SWE Agent is a proprietary, free-to-install application. Everything you need to adopt, inspect and improve it lives in a public repository you can read, fork and contribute to.
The one-line installers are plain shell and PowerShell, Apache-2.0, and verify each download's SHA-256 before running it.
Every guide, from installation to the contract format, is public and open to corrections.
Benchmarks with their evidence
A published methodology, per-task outcomes, and results labelled self-reported until someone independent checks them.
No telemetry, no account. Your code goes only to the model you choose — or nowhere, with a local model.
One command. Then describe the change.
The installer shows what it will do and asks before it changes anything. It installs for your user only — Python included — so there's no pip, virtual environment or PATH editing.
Paste the command for your system and answer Y.
Paste an API key, or point it at a model on your machine.
cd into a project, start the agent, and say what should change.
Copy $ curl -fsSL https://raw.githubusercontent.com/ huynvic/ albertcode-swe-agent/ main/ installer/ install.sh | sh Open Terminal (⌘ Space, then type Terminal), paste the command, read the plan, then type Y.
Open a new Terminal window when it finishes, then run albertcode --version.
More in the documentation, and in each repository's own FAQ.
AlbertCode SWE Agent is an AI coding agent that runs on your machine. It plans a change, builds it in an isolated copy of your repository, runs your own tests, reviews the result, and waits for you to accept the diff before anything reaches your files. Is it open source? No — and we'd rather say so plainly. AlbertCode SWE Agent is a proprietary application. Its public GitHub repository holds the installers, the documentation, examples and the benchmark methodology, under Apache-2.0 (code) and CC BY 4.0 (docs), and welcome contributions. The engine's source isn't published. What's open, and what isn't . Does it cost anything? Installing is free and needs no account. AlbertCode doesn't charge for model usage: you use your own model provider's account, or a local model at no cost. The repository notes that terms for AlbertCode itself will be published with its first public release. Does my code leave my machine? Only to the model provider you connect (and to MCP servers you add), because the model has to read the code it works on. With a local model, nothing leaves your machine. There is no telemetry and no account. Saved keys are kept by your operating system's keychain or keyring. Which models can I use? Hosted providers with an API key, Kimi, Hugging Face, OpenRouter, any OpenAI-compatible service, or a model on your own machine through Ollama, LM Studio, vLLM or llama.cpp. The model must support tool calling; albertcode models test checks one. Can it change my files without asking? Not in Governed mode, the default: you approve the plan, then the diff. Direct mode works in your real files and asks before each write, unless you explicitly choose auto-approve. Direct needs Git, because Git is the undo. Can I run it in the browser, here? No. AlbertCode runs on your own computer; it isn't offered as a hosted service, and this website never runs code for you. Its own browser workspace opens on your computer with albertcode web . The run planner here prepares the exact c
[truncated]
From idea to working software.
Install in a minute or two, connect the model you already use, and let the first change come back with its proof.
Read the docs Albert Code SWE From idea to working software. AI software engineering agents that plan, build in isolation, verify with your own tests, and wait for your approval.
Free to install · No account · No telemetry
Discussions (opens in a new tab)
Report an issue (opens in a new tab)
Report a vulnerability (opens in a new tab)
Support AlbertCode (opens in a new tab)
Licensing (opens in a new tab)
© 2026 Huy Nguyen . AlbertCode SWE Agent and mini-AlbertCode SWE Agent are proprietary software, free to install. Their public repositories' code is Apache-2.0 and documentation CC BY 4.0.
IBM Plex typefaces under the SIL Open Font License.
