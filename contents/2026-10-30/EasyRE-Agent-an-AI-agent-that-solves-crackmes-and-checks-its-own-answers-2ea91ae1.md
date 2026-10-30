---
source: "https://github.com/wojtas99/EasyRE-Agent"
hn_url: "https://news.ycombinator.com/item?id=50309306"
title: "EasyRE-Agent – an AI agent that solves crackmes and checks its own answers"
article_title: "GitHub - wojtas99/EasyRE-Agent · GitHub"
image: "https://opengraph.githubassets.com/b3d1a03948b809c9d078907f05e32bf86bab6e9bc7c58b1f2ab418b3297fef0e/wojtas99/EasyRE-Agent"
author: "krawczorek"
captured_at: "2026-10-30T20:27:31Z"
capture_tool: "hn-digest"
hn_id: 50309306
score: 2
comments: 0
posted_at: "2026-10-30T19:54:48Z"
tags:
  - hacker-news
---

# EasyRE-Agent – an AI agent that solves crackmes and checks its own answers

- HN: [50309306](https://news.ycombinator.com/item?id=50309306)
- Source: [github.com](https://github.com/wojtas99/EasyRE-Agent)
- Score: 2
- Comments: 0
- Posted: 2026-10-30T19:54:48Z

## Translation

Title: EasyRE-Agent – an AI agent that solves crackmes and checks its own answers
Article title: GitHub - wojtas99/EasyRE-Agent · GitHub
Description: Contribute to wojtas99/EasyRE-Agent development by creating an account on GitHub.

Article text:
GitHub - wojtas99/EasyRE-Agent · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
wojtas99
EasyRE-Agent Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
config config skills skills toolchain toolchain .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml easyre easyre easyre.cmd easyre.cmd easyre.sh easyre.sh mockm.py mockm.py setup.cmd setup.cmd setup.ps1 setup.ps1 setup.sh setup.sh View all files Repository files navigation
Give it a binary and a task. Pick any model. Get the answer, or learn how to find it yourself.
EasyRE-Agent runs an AI agent inside a Docker container. The container holds your reverse engineering toolchain, and the agent uses it to solve a task you give it.
📁 tasks/crackme ─► 🔎 triage ─► 🤖 agent (your model) ─► ✅ verifier ─► 📄 report.md / 💡 hints.md
Feature
🤖
Any model
Claude Code, OpenAI Codex, or any OpenAI-compatible API: OpenRouter, Ollama, LM Studio, vLLM
✅
Verified answers
Checks against a regex, a success message, or your own command. The agent keeps working until a candidate passes.
🎓
Hint mode
The agent solves privately and gives you graded hints, one at a time. The answer stays hidden, including in the console output.
🧑‍🏫
Tutor mode
An interactive tutor: you tell it what you did, and it guides you step by step without spoiling the answer
💰
Stats
Totals of time, tokens and cost per model over all jobs
⏱️
Budgets
Time and token limits per job
🔁
Resumable
Continue a job later with more time or a different model
🔎
Triage
File types, runtimes and packers detected up front
📄
Reports
Answer, solution path, scripts, dead ends, exact model, time and tokens
🧰
Your toolchain
Plug-in install script plus Markdown skills (playbooks) the agent reads on demand
🎯 Why it exists
Reversing a crackme or a CTF binary is mostly the same loop every time: figure out what the file is, pick the right tools, find the check, pull out the constants, write a small solver, and confirm the answer. An AI agent can do much of that loop, but in practice three things get in the way:
The toolchain. A typical task needs a mix of Ghidra, Detect It Easy, Wine, and decompilers for .NET, Python, Java or JavaScript. Installing and wiring all of that on every machine is slow and fragile. EasyRE-Agent ships it as one Docker image that behaves the same on Windows, Linux and macOS.
Guessing. Without tools and a checker, a model tends to produce an answer that only looks plausible. Here the agent works on the real file with real tools, follows a playbook (skill) for that kind of target, and an answer counts only when the program or your checker accepts it.
Lock-in. It works with Claude, Codex or any OpenAI-compatible model, including local ones, so you can use what you already have and compare models on the same task.
It is also built for learning. Being handed the answer teaches little, so hint and tutor modes let the agent solve the task privately and then guide you toward the solution without spoiling it.
EasyRE-Agent is meant for crackmes, CTF reversing challenges and puzzle binaries, not for analyzing real malware (see ⚠️ Security ).
Pick your system: 💻 Windows · 🐧 Linux · 🍎 macOS
You need about 10 GB of free disk space (see 💾 Disk space ).
Install Docker. Get Docker Desktop , install it with the default settings (WSL 2), and start it.
Run the setup. Double-click setup.cmd in the project folder. It builds the image, asks which model to use, logs you in and checks everything (see ⚙️ What the setup does ).
Add a task. Put the files into a folder under tasks :
tasks\crackme\crackme.exe
tasks\crackme\task.md (description, optional)
Solve it. Open PowerShell or Command Prompt in the project folder:
.\ easyre.cmd run crackme -- success " Correct "
.\ easyre.cmd open
The first command solves the task. The second opens the report.
Install Docker. Install Docker Engine with the Compose plugin. Then let your user run Docker without sudo , and log out and back in:
sudo usermod -aG docker $USER
Run the setup in the project folder:
sh setup.sh
Add a task: put its files into tasks/crackme/ (plus an optional task.md with the description).
Solve it:
sh easyre.sh run crackme --success " Correct "
sh easyre.sh open
The image is x86-64 (Wine needs 32-bit x86 libraries). On ARM Linux it runs only with x86 emulation (qemu/binfmt) and is slow.
Install Docker. Get Docker Desktop for Mac and start it. On Apple Silicon (M1–M4), enable Settings → General → Use Rosetta for x86_64/amd64 emulation . The image is x86-64, so it runs emulated there: it works, but builds and runs more slowly.
Run the setup in Terminal, in the project folder:
sh setup.sh
Add a task: put its files into tasks/crackme/ (plus an optional task.md with the description).
Solve it:
sh easyre.sh run crackme --success " Correct "
sh easyre.sh open
📝 The examples below use the Windows form .\easyre.cmd . On Linux and macOS, write sh easyre.sh instead; everything else is the same.
Builds the image (the first build takes several minutes).
Asks which model you want to use and logs you in:
Runs easyre doctor to check that everything works.
--success is the text the program prints when the answer is right. Leave it out if the program prints nothing like that.
.\ easyre.cmd doctor
doctor checks your model profile, credentials, API connection, Wine and skills. It shows each problem together with how to fix it. To switch models later, run the setup again, or see 🎛️ Profiles .
Colors: the output is colored. If you see codes like ←[32m instead of colors, use Windows Terminal, or add NO_COLOR=1 to .env .
easyre run TARGET [options] solve a task (folder or file under ./tasks)
easyre resume JOB [options] continue a job
easyre hint JOB [--answer] next hint (or reveal the answer)
easyre tutor JOB continue a tutor conversation
easyre jobs list jobs
easyre open [JOB] [folder] open a job's report (or tutor chat, or folder)
easyre profiles list model profiles
easyre doctor check that the setup works
easyre stats time, tokens and cost per model
easyre close JOB mark a job as closed (resume reopens it)
easyre rm JOB... | --FILTER delete job folders; filters: --open --closed --unsolved --solved --all
easyre login claude|codex authenticate a backend
Option
Description
--desc / --desc-file
Task description. Default: task.md , opis.txt or README.md from the task folder.
--success REGEX
Output that means the answer was accepted
--verify CMD
Checker command. {} is replaced by the candidate; otherwise the candidate goes to stdin.
--flag-format REGEX
Required answer shape, e.g. "flag\{.+\}"
--profile
Model profile from config/backends.toml
--backend
claude (default), codex or openai
--model
Model name for the backend
--mode
solve (default), hint or tutor
--budget-min / --budget-tokens
Time budget in minutes (default 120) and output-token budget
--logs
Stream the agent's commands live instead of the progress bar
-i
Interactive session (Claude Code / Codex UI)
📌 Examples
.\ easyre.cmd run crackme -- success " (?i)correct "
.\ easyre.cmd run keygenme -- verify " wine /work/src/app.exe {} " -- success " Valid "
.\ easyre.cmd run ctf - rev -- flag - format " CTF\{.+\} "
.\ easyre.cmd run crackme -- mode hint
.\ easyre.cmd hint crackme
.\ easyre.cmd run crackme -- mode tutor -- success " Correct "
.\ easyre.cmd tutor crackme
.\ easyre.cmd run crackme -- backend openai -- model qwen3 - coder -- budget - min 30
.\ easyre.cmd resume crackme -- budget - min 60
📂 Opening results
.\ easyre.cmd open
.\ easyre.cmd open crackme
.\ easyre.cmd open crackme folder
Opens the newest job's report.md (or tutor.md for tutor jobs) in your default app.
Opens the newest job whose name starts with crackme .
Opens that job's folder in Explorer.
open runs on your computer, not in Docker, so it starts instantly.
.\easyre.cmd run crackme --mode tutor --success "Correct"
The agent first solves the task privately. The console shows only progress, so nothing spoils the answer.
Then a chat starts. Tell it what you did and what you see, and it replies to where you are: it confirms the right path, warns about dead ends, and suggests one small next step.
tutor> This is a small 64-bit console program. What have you tried so far?
you> I found main, it calls checkPassword. What next?
tutor> Good. Look at what checkPassword compares its argument with...
Command
What it does
/hint
Ask for a nudge
/answer
Reveal the answer (asks for confirmation)
/quit
Leave. Continue later with easyre tutor <job>
The conversation is saved to tutor.md in the job folder. If a reply accidentally contains the answer, it is replaced with [hidden] .
--flag-format must match, if set.
--verify runs, and its output is matched against --success . Without --success , exit code 0 counts as success.
If only --success is set and the task has exactly one .exe or ELF file, that file is run automatically with the candidate on stdin (through Wine for .exe ).
With no verifier, the agent must prove the answer itself, and the report marks it unverified .
easyre --help also lists two commands marked (agent) . The agent runs them inside the job's container; you do not need them:
Run from your own console, they only print a note, because there is no job to check against. Use easyre jobs and easyre open to see results.
🔬 How a job runs: a crackme, step by step
The example is a console crackme that asks for a serial and prints Correct! or Wrong :
tasks/keygenme/keygenme.exe
tasks/keygenme/task.md Find the valid serial.
.\ easyre.cmd run keygenme -- success " Correct "
Steps 1–3 and 5–7 are fixed code in easyre . Step 4 is what a model typically does: it is not scripted. The model picks its tools itself, guided by the skills and the rules in config/AGENT.md .
1. Start. easyre.cmd runs docker compose run --rm easyre easyre run … in a fresh container. tasks/ and jobs/ are mounted, and so are easyre , config/ and skills/ (live, so edits need no rebuild).
Creates jobs/keygenme-<timestamp>/ with src/ (a copy of the task, archives extracted), agent/ and .easyre/ .
Takes the description from task.md (or opis.txt , README.md ).
Links the job as /work , the agent's working directory.
Picks the verifier: --success is set and the task has exactly one .exe , so the verifier is wine /work/src/keygenme.exe , with the candidate on stdin.
3. Triage. For every file: file , byte markers (PyInstaller, .NET, UPX, Go, Rust, Delphi, AutoIt, WASM, Electron, …), repair of a broken MZ header, and diec . On the main native binary, the Ghidra

[truncated]

## Original Extract

Contribute to wojtas99/EasyRE-Agent development by creating an account on GitHub.

GitHub - wojtas99/EasyRE-Agent · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
wojtas99
EasyRE-Agent Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
config config skills skills toolchain toolchain .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml easyre easyre easyre.cmd easyre.cmd easyre.sh easyre.sh mockm.py mockm.py setup.cmd setup.cmd setup.ps1 setup.ps1 setup.sh setup.sh View all files Repository files navigation
Give it a binary and a task. Pick any model. Get the answer, or learn how to find it yourself.
EasyRE-Agent runs an AI agent inside a Docker container. The container holds your reverse engineering toolchain, and the agent uses it to solve a task you give it.
📁 tasks/crackme ─► 🔎 triage ─► 🤖 agent (your model) ─► ✅ verifier ─► 📄 report.md / 💡 hints.md
Feature
🤖
Any model
Claude Code, OpenAI Codex, or any OpenAI-compatible API: OpenRouter, Ollama, LM Studio, vLLM
✅
Verified answers
Checks against a regex, a success message, or your own command. The agent keeps working until a candidate passes.
🎓
Hint mode
The agent solves privately and gives you graded hints, one at a time. The answer stays hidden, including in the console output.
🧑‍🏫
Tutor mode
An interactive tutor: you tell it what you did, and it guides you step by step without spoiling the answer
💰
Stats
Totals of time, tokens and cost per model over all jobs
⏱️
Budgets
Time and token limits per job
🔁
Resumable
Continue a job later with more time or a different model
🔎
Triage
File types, runtimes and packers detected up front
📄
Reports
Answer, solution path, scripts, dead ends, exact model, time and tokens
🧰
Your toolchain
Plug-in install script plus Markdown skills (playbooks) the agent reads on demand
🎯 Why it exists
Reversing a crackme or a CTF binary is mostly the same loop every time: figure out what the file is, pick the right tools, find the check, pull out the constants, write a small solver, and confirm the answer. An AI agent can do much of that loop, but in practice three things get in the way:
The toolchain. A typical task needs a mix of Ghidra, Detect It Easy, Wine, and decompilers for .NET, Python, Java or JavaScript. Installing and wiring all of that on every machine is slow and fragile. EasyRE-Agent ships it as one Docker image that behaves the same on Windows, Linux and macOS.
Guessing. Without tools and a checker, a model tends to produce an answer that only looks plausible. Here the agent works on the real file with real tools, follows a playbook (skill) for that kind of target, and an answer counts only when the program or your checker accepts it.
Lock-in. It works with Claude, Codex or any OpenAI-compatible model, including local ones, so you can use what you already have and compare models on the same task.
It is also built for learning. Being handed the answer teaches little, so hint and tutor modes let the agent solve the task privately and then guide you toward the solution without spoiling it.
EasyRE-Agent is meant for crackmes, CTF reversing challenges and puzzle binaries, not for analyzing real malware (see ⚠️ Security ).
Pick your system: 💻 Windows · 🐧 Linux · 🍎 macOS
You need about 10 GB of free disk space (see 💾 Disk space ).
Install Docker. Get Docker Desktop , install it with the default settings (WSL 2), and start it.
Run the setup. Double-click setup.cmd in the project folder. It builds the image, asks which model to use, logs you in and checks everything (see ⚙️ What the setup does ).
Add a task. Put the files into a folder under tasks :
tasks\crackme\crackme.exe
tasks\crackme\task.md (description, optional)
Solve it. Open PowerShell or Command Prompt in the project folder:
.\ easyre.cmd run crackme -- success " Correct "
.\ easyre.cmd open
The first command solves the task. The second opens the report.
Install Docker. Install Docker Engine with the Compose plugin. Then let your user run Docker without sudo , and log out and back in:
sudo usermod -aG docker $USER
Run the setup in the project folder:
sh setup.sh
Add a task: put its files into tasks/crackme/ (plus an optional task.md with the description).
Solve it:
sh easyre.sh run crackme --success " Correct "
sh easyre.sh open
The image is x86-64 (Wine needs 32-bit x86 libraries). On ARM Linux it runs only with x86 emulation (qemu/binfmt) and is slow.
Install Docker. Get Docker Desktop for Mac and start it. On Apple Silicon (M1–M4), enable Settings → General → Use Rosetta for x86_64/amd64 emulation . The image is x86-64, so it runs emulated there: it works, but builds and runs more slowly.
Run the setup in Terminal, in the project folder:
sh setup.sh
Add a task: put its files into tasks/crackme/ (plus an optional task.md with the description).
Solve it:
sh easyre.sh run crackme --success " Correct "
sh easyre.sh open
📝 The examples below use the Windows form .\easyre.cmd . On Linux and macOS, write sh easyre.sh instead; everything else is the same.
Builds the image (the first build takes several minutes).
Asks which model you want to use and logs you in:
Runs easyre doctor to check that everything works.
--success is the text the program prints when the answer is right. Leave it out if the program prints nothing like that.
.\ easyre.cmd doctor
doctor checks your model profile, credentials, API connection, Wine and skills. It shows each problem together with how to fix it. To switch models later, run the setup again, or see 🎛️ Profiles .
Colors: the output is colored. If you see codes like ←[32m instead of colors, use Windows Terminal, or add NO_COLOR=1 to .env .
easyre run TARGET [options] solve a task (folder or file under ./tasks)
easyre resume JOB [options] continue a job
easyre hint JOB [--answer] next hint (or reveal the answer)
easyre tutor JOB continue a tutor conversation
easyre jobs list jobs
easyre open [JOB] [folder] open a job's report (or tutor chat, or folder)
easyre profiles list model profiles
easyre doctor check that the setup works
easyre stats time, tokens and cost per model
easyre close JOB mark a job as closed (resume reopens it)
easyre rm JOB... | --FILTER delete job folders; filters: --open --closed --unsolved --solved --all
easyre login claude|codex authenticate a backend
Option
Description
--desc / --desc-file
Task description. Default: task.md , opis.txt or README.md from the task folder.
--success REGEX
Output that means the answer was accepted
--verify CMD
Checker command. {} is replaced by the candidate; otherwise the candidate goes to stdin.
--flag-format REGEX
Required answer shape, e.g. "flag\{.+\}"
--profile
Model profile from config/backends.toml
--backend
claude (default), codex or openai
--model
Model name for the backend
--mode
solve (default), hint or tutor
--budget-min / --budget-tokens
Time budget in minutes (default 120) and output-token budget
--logs
Stream the agent's commands live instead of the progress bar
-i
Interactive session (Claude Code / Codex UI)
📌 Examples
.\ easyre.cmd run crackme -- success " (?i)correct "
.\ easyre.cmd run keygenme -- verify " wine /work/src/app.exe {} " -- success " Valid "
.\ easyre.cmd run ctf - rev -- flag - format " CTF\{.+\} "
.\ easyre.cmd run crackme -- mode hint
.\ easyre.cmd hint crackme
.\ easyre.cmd run crackme -- mode tutor -- success " Correct "
.\ easyre.cmd tutor crackme
.\ easyre.cmd run crackme -- backend openai -- model qwen3 - coder -- budget - min 30
.\ easyre.cmd resume crackme -- budget - min 60
📂 Opening results
.\ easyre.cmd open
.\ easyre.cmd open crackme
.\ easyre.cmd open crackme folder
Opens the newest job's report.md (or tutor.md for tutor jobs) in your default app.
Opens the newest job whose name starts with crackme .
Opens that job's folder in Explorer.
open runs on your computer, not in Docker, so it starts instantly.
.\easyre.cmd run crackme --mode tutor --success "Correct"
The agent first solves the task privately. The console shows only progress, so nothing spoils the answer.
Then a chat starts. Tell it what you did and what you see, and it replies to where you are: it confirms the right path, warns about dead ends, and suggests one small next step.
tutor> This is a small 64-bit console program. What have you tried so far?
you> I found main, it calls checkPassword. What next?
tutor> Good. Look at what checkPassword compares its argument with...
Command
What it does
/hint
Ask for a nudge
/answer
Reveal the answer (asks for confirmation)
/quit
Leave. Continue later with easyre tutor <job>
The conversation is saved to tutor.md in the job folder. If a reply accidentally contains the answer, it is replaced with [hidden] .
--flag-format must match, if set.
--verify runs, and its output is matched against --success . Without --success , exit code 0 counts as success.
If only --success is set and the task has exactly one .exe or ELF file, that file is run automatically with the candidate on stdin (through Wine for .exe ).
With no verifier, the agent must prove the answer itself, and the report marks it unverified .
easyre --help also lists two commands marked (agent) . The agent runs them inside the job's container; you do not need them:
Run from your own console, they only print a note, because there is no job to check against. Use easyre jobs and easyre open to see results.
🔬 How a job runs: a crackme, step by step
The example is a console crackme that asks for a serial and prints Correct! or Wrong :
tasks/keygenme/keygenme.exe
tasks/keygenme/task.md Find the valid serial.
.\ easyre.cmd run keygenme -- success " Correct "
Steps 1–3 and 5–7 are fixed code in easyre . Step 4 is what a model typically does: it is not scripted. The model picks its tools itself, guided by the skills and the rules in config/AGENT.md .
1. Start. easyre.cmd runs docker compose run --rm easyre easyre run … in a fresh container. tasks/ and jobs/ are mounted, and so are easyre , config/ and skills/ (live, so edits need no rebuild).
Creates jobs/keygenme-<timestamp>/ with src/ (a copy of the task, archives extracted), agent/ and .easyre/ .
Takes the description from task.md (or opis.txt , README.md ).
Links the job as /work , the agent's working directory.
Picks the verifier: --success is set and the task has exactly one .exe , so the verifier is wine /work/src/keygenme.exe , with the candidate on stdin.
3. Triage. For every file: file , byte markers (PyInstaller, .NET, UPX, Go, Rust, Delphi, AutoIt, WASM, Electron, …), repair of a broken MZ header, and diec . On the main native binary, the Ghidra

[truncated]
