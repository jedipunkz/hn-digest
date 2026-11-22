---
source: "https://github.com/trail-openers/qlaude"
hn_url: "https://news.ycombinator.com/item?id=50597098"
title: "Qlaude – Claude Code against your own model server"
article_title: "GitHub - trail-openers/qlaude: Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). · GitHub"
image: "https://opengraph.githubassets.com/ac9936ed677d78915dd98d7ae8e3c85730a464c63e1d90a5827b28d21820586c/trail-openers/qlaude"
author: "velmu"
captured_at: "2026-11-22T07:48:16Z"
capture_tool: "hn-digest"
hn_id: 50597098
score: 1
comments: 0
posted_at: "2026-11-22T07:20:04Z"
tags:
  - hacker-news
---

# Qlaude – Claude Code against your own model server

- HN: [50597098](https://news.ycombinator.com/item?id=50597098)
- Source: [github.com](https://github.com/trail-openers/qlaude)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T07:20:04Z

## Translation

Title: Qlaude – Claude Code against your own model server
Article title: GitHub - trail-openers/qlaude: Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). · GitHub
Description: Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). - trail-openers/qlaude

Article text:
GitHub - trail-openers/qlaude: Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). · GitHub
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
trail-openers
qlaude Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows).
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
100 Commits 100 Commits Folders and files
.github .github bin bin docs docs tests tests tools tools .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SYSTEM-PROMPT.md SYSTEM-PROMPT.md install.sh install.sh install.sha256 install.sha256 View all files Repository files navigation
qlaude — Claude Code against your own model server
Run Claude Code , Anthropic's terminal coding
agent, against your own model server. Claude Code is a command-line tool
that reads your code, edits files and runs commands. One paste and five
setup questions get you running. You need no Anthropic subscription or
account, and you do not have to own a GPU or run a server to try it: a
hosted provider that offers an Anthropic-compatible API endpoint can
stand in for your own vLLM — see Prerequisites .
Works on Linux and macOS. Windows is not supported.
Verified November 2026 on Linux and macOS: Claude Code 2.1.325 + vLLM
0.33 on the POST /v1/messages route, plus a llama-swap + llama.cpp
stack serving a Gemma MoE — what exactly was tested, and where, is in
Status .
The reference server is vLLM , the open-source
LLM inference server: since v0.11.1 it answers Anthropic's Messages API
( POST /v1/messages ) itself, so Claude Code connects directly, without
a translation server in between — vLLM's
official Claude Code integration .
Anything that speaks the same route works.
You need three things: a model server that is already running (or a
hosted provider), its URL, and its access token. This repo does not run
a server for you — if you don't have one yet, read
Prerequisites first.
qlaude is the launcher this repo ships: a single POSIX-sh script
(POSIX sh is the most portable shell dialect). It stores your
endpoint/model/token in ~/.config/qlaude/config.env (mode 0600 —
readable only by your user), runs the self-checks below as
qlaude doctor , and launches Claude Code with every setting wired for
you.
curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
That paste installs ~/.local/bin/qlaude plus
~/.config/qlaude/SYSTEM-PROMPT.md (mode 0600 , the optional behavior
prompt — see Behavior prompt ). Nothing else:
no daemon, and no shell-rc edits — the only command that ever touches rc
files is qlaude ensure-path , and only with your consent. Every
downloaded file is verified against install.sha256 before running —
corruption/drift protection, not first-run authentication (the checksum
list arrives over the same HTTPS connection; trust here is TLS plus the
host). The token is never echoed, never printed, never a command-line
argument.
In a terminal the installer prints the PATH line to copy, if needed,
then runs the setup interview. It asks five questions:
Endpoint URL — the address of your server; no trailing /v1 needed,
it normalizes.
Token — shown as * while you type, stored 0600 .
Model — picked from the models the server lists; Enter accepts a single
discovered id. If the listing is unavailable, the interview lets you
type the id yourself.
Context — how many tokens the model can see at once.
Reasoning effort — how hard the model thinks before answering.
The interview ends with the five checks below. A successful run looks
like this (output sketch — your server URL and model names will differ):
$ curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
qlaude install: fetching from https://raw.githubusercontent.com/trail-openers/qlaude/main
qlaude install: checksums verified (truncation/corruption guard; supply chain = TLS + the host)
qlaude install: installed ~/.local/bin/qlaude
qlaude install: PATH note: ~/.local/bin is not on PATH — copy-paste now: export PATH="$HOME/.local/bin:$PATH" (then: ~/.local/bin/qlaude ensure-path)
qlaude install: next: guided configuration — answer the prompts (Ctrl-C skips; resume any time: ~/.local/bin/qlaude setup)
qlaude: endpoint URL (the Anthropic-compatible server qlaude will route to — type or paste, it echoes back): https://models.example.com
qlaude: token for https://models.example.com — paste it (shown as *, stored 0600 in config.env, never printed), Enter when done: **********
qlaude: model — 1 discovered; Enter accepts my-model (or type another id):
qlaude: context window tokens (= your server's --max-model-len; Enter = 262144):
qlaude: reasoning effort low/medium/xhigh ('high' is mapped for you; Enter = medium):
PASS 0 prereqs: all found
PASS 1 messages: 200 with content text
PASS 2 claude: echoed OK
PASS 3 tools: tool call round-tripped
PASS 4 info: no informational issues
And that is the whole setup: you set no environment variables by
hand. The launcher exports everything Claude Code needs at launch; the
full reference for each one is
Environment variables , which you only need when
debugging or running claude without qlaude .
Daily use: plain qlaude opens Claude Code against your endpoint —
just type what you want done. qlaude doctor repeats the five checks
any time; qlaude config shows your saved settings (never the token).
To hand arguments straight to Claude Code — a -p probe, --resume ,
anything after -- :
qlaude -- -p " Reply with exactly: OK " --max-turns 1
Skipped or Ctrl-C'd the interview? Rerun with ~/.local/bin/qlaude setup .
In scripts and CI the installer only installs; configure explicitly,
feeding the token over stdin — never as an argument (in CI, from your
secret store, not shell history):
curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
printf ' %s ' " $YOUR_TOKEN " | ~ /.local/bin/qlaude setup -y --token-stdin \
--endpoint https://models.example.com --model my-model
Uninstall
No subcommand. Three lines. The last one is macOS only and removes the
keychain entry:
rm " $HOME /.local/bin/qlaude "
rm -rf " $HOME /.config/qlaude "
security delete-generic-password -s qlaude # macOS only
If you installed with QLAUDE_INSTALL_DIR / QLAUDE_CONFIG_DIR set, drop
those files from wherever they went instead. If you ran
qlaude ensure-path , delete the # qlaude PATH … # end qlaude PATH
marker block from the file it named ( ~/.profile by default,
~/.zprofile for zsh and macOS).
Run qlaude doctor first — it names the failing check. Then:
Troubleshooting — known symptoms, real
causes and fixes (including the misleading "model may not exist"
error, which is almost always a trailing /v1 ).
Verify, in this order — two commands that split
server problems from environment problems.
Server-side checklist — the vLLM flags
that make tools work ( --tool-call-parser , --served-model-name ,
--max-model-len ).
Prerequisites — Claude Code install, your
own vLLM or a hosted provider, and where the token lives.
How it works — the Messages route, and the
two details behind almost every failure people hit.
Environment variables — the canonical
reference for every variable (the normal path sets none).
Auto mode and the safety classifier — why the
launcher exports CLAUDE_CODE_AUTO_MODE_SERVER=0 , and when to opt
back in.
Gateways — direct is the default; when a proxy
earns its keep, and which to avoid.
Behavior prompt — what the optional
SYSTEM-PROMPT.md rules prevent, and what deliberately is not in it.
Status — exactly what has been verified live.
Design and rationale: the spec .
Questions and bugs: open an issue —
include what you ran, what came back, your OS, and which commit you
installed from (your QLAUDE_REF value; if you left it on main , the
install date helps too — the installer fetches single files, so there is
no local clone to query with git rev-parse ).
Security issues: please do not open a public issue — report privately
via SECURITY.md (GitHub's Report a vulnerability form).
Contributions: CONTRIBUTING.md — the offline test suite,
the shellcheck/dash gates, the line budgets (hard per-file size caps CI
enforces), and the README rules.
Licensed under the Apache License 2.0 — use it, fork it, ship it
in your company.
Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows).
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). - trail-openers/qlaude

GitHub - trail-openers/qlaude: Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows). · GitHub
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
trail-openers
qlaude Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows).
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
100 Commits 100 Commits Folders and files
.github .github bin bin docs docs tests tests tools tools .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SYSTEM-PROMPT.md SYSTEM-PROMPT.md install.sh install.sh install.sha256 install.sha256 View all files Repository files navigation
qlaude — Claude Code against your own model server
Run Claude Code , Anthropic's terminal coding
agent, against your own model server. Claude Code is a command-line tool
that reads your code, edits files and runs commands. One paste and five
setup questions get you running. You need no Anthropic subscription or
account, and you do not have to own a GPU or run a server to try it: a
hosted provider that offers an Anthropic-compatible API endpoint can
stand in for your own vLLM — see Prerequisites .
Works on Linux and macOS. Windows is not supported.
Verified November 2026 on Linux and macOS: Claude Code 2.1.325 + vLLM
0.33 on the POST /v1/messages route, plus a llama-swap + llama.cpp
stack serving a Gemma MoE — what exactly was tested, and where, is in
Status .
The reference server is vLLM , the open-source
LLM inference server: since v0.11.1 it answers Anthropic's Messages API
( POST /v1/messages ) itself, so Claude Code connects directly, without
a translation server in between — vLLM's
official Claude Code integration .
Anything that speaks the same route works.
You need three things: a model server that is already running (or a
hosted provider), its URL, and its access token. This repo does not run
a server for you — if you don't have one yet, read
Prerequisites first.
qlaude is the launcher this repo ships: a single POSIX-sh script
(POSIX sh is the most portable shell dialect). It stores your
endpoint/model/token in ~/.config/qlaude/config.env (mode 0600 —
readable only by your user), runs the self-checks below as
qlaude doctor , and launches Claude Code with every setting wired for
you.
curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
That paste installs ~/.local/bin/qlaude plus
~/.config/qlaude/SYSTEM-PROMPT.md (mode 0600 , the optional behavior
prompt — see Behavior prompt ). Nothing else:
no daemon, and no shell-rc edits — the only command that ever touches rc
files is qlaude ensure-path , and only with your consent. Every
downloaded file is verified against install.sha256 before running —
corruption/drift protection, not first-run authentication (the checksum
list arrives over the same HTTPS connection; trust here is TLS plus the
host). The token is never echoed, never printed, never a command-line
argument.
In a terminal the installer prints the PATH line to copy, if needed,
then runs the setup interview. It asks five questions:
Endpoint URL — the address of your server; no trailing /v1 needed,
it normalizes.
Token — shown as * while you type, stored 0600 .
Model — picked from the models the server lists; Enter accepts a single
discovered id. If the listing is unavailable, the interview lets you
type the id yourself.
Context — how many tokens the model can see at once.
Reasoning effort — how hard the model thinks before answering.
The interview ends with the five checks below. A successful run looks
like this (output sketch — your server URL and model names will differ):
$ curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
qlaude install: fetching from https://raw.githubusercontent.com/trail-openers/qlaude/main
qlaude install: checksums verified (truncation/corruption guard; supply chain = TLS + the host)
qlaude install: installed ~/.local/bin/qlaude
qlaude install: PATH note: ~/.local/bin is not on PATH — copy-paste now: export PATH="$HOME/.local/bin:$PATH" (then: ~/.local/bin/qlaude ensure-path)
qlaude install: next: guided configuration — answer the prompts (Ctrl-C skips; resume any time: ~/.local/bin/qlaude setup)
qlaude: endpoint URL (the Anthropic-compatible server qlaude will route to — type or paste, it echoes back): https://models.example.com
qlaude: token for https://models.example.com — paste it (shown as *, stored 0600 in config.env, never printed), Enter when done: **********
qlaude: model — 1 discovered; Enter accepts my-model (or type another id):
qlaude: context window tokens (= your server's --max-model-len; Enter = 262144):
qlaude: reasoning effort low/medium/xhigh ('high' is mapped for you; Enter = medium):
PASS 0 prereqs: all found
PASS 1 messages: 200 with content text
PASS 2 claude: echoed OK
PASS 3 tools: tool call round-tripped
PASS 4 info: no informational issues
And that is the whole setup: you set no environment variables by
hand. The launcher exports everything Claude Code needs at launch; the
full reference for each one is
Environment variables , which you only need when
debugging or running claude without qlaude .
Daily use: plain qlaude opens Claude Code against your endpoint —
just type what you want done. qlaude doctor repeats the five checks
any time; qlaude config shows your saved settings (never the token).
To hand arguments straight to Claude Code — a -p probe, --resume ,
anything after -- :
qlaude -- -p " Reply with exactly: OK " --max-turns 1
Skipped or Ctrl-C'd the interview? Rerun with ~/.local/bin/qlaude setup .
In scripts and CI the installer only installs; configure explicitly,
feeding the token over stdin — never as an argument (in CI, from your
secret store, not shell history):
curl -fsSL https://raw.githubusercontent.com/trail-openers/qlaude/main/install.sh | sh
printf ' %s ' " $YOUR_TOKEN " | ~ /.local/bin/qlaude setup -y --token-stdin \
--endpoint https://models.example.com --model my-model
Uninstall
No subcommand. Three lines. The last one is macOS only and removes the
keychain entry:
rm " $HOME /.local/bin/qlaude "
rm -rf " $HOME /.config/qlaude "
security delete-generic-password -s qlaude # macOS only
If you installed with QLAUDE_INSTALL_DIR / QLAUDE_CONFIG_DIR set, drop
those files from wherever they went instead. If you ran
qlaude ensure-path , delete the # qlaude PATH … # end qlaude PATH
marker block from the file it named ( ~/.profile by default,
~/.zprofile for zsh and macOS).
Run qlaude doctor first — it names the failing check. Then:
Troubleshooting — known symptoms, real
causes and fixes (including the misleading "model may not exist"
error, which is almost always a trailing /v1 ).
Verify, in this order — two commands that split
server problems from environment problems.
Server-side checklist — the vLLM flags
that make tools work ( --tool-call-parser , --served-model-name ,
--max-model-len ).
Prerequisites — Claude Code install, your
own vLLM or a hosted provider, and where the token lives.
How it works — the Messages route, and the
two details behind almost every failure people hit.
Environment variables — the canonical
reference for every variable (the normal path sets none).
Auto mode and the safety classifier — why the
launcher exports CLAUDE_CODE_AUTO_MODE_SERVER=0 , and when to opt
back in.
Gateways — direct is the default; when a proxy
earns its keep, and which to avoid.
Behavior prompt — what the optional
SYSTEM-PROMPT.md rules prevent, and what deliberately is not in it.
Status — exactly what has been verified live.
Design and rationale: the spec .
Questions and bugs: open an issue —
include what you ran, what came back, your OS, and which commit you
installed from (your QLAUDE_REF value; if you left it on main , the
install date helps too — the installer fetches single files, so there is
no local clone to query with git rev-parse ).
Security issues: please do not open a public issue — report privately
via SECURITY.md (GitHub's Report a vulnerability form).
Contributions: CONTRIBUTING.md — the offline test suite,
the shellcheck/dash gates, the line budgets (hard per-file size caps CI
enforces), and the README rules.
Licensed under the Apache License 2.0 — use it, fork it, ship it
in your company.
Run Claude Code against your own model server: one paste, guided setup, self-test. vLLM Anthropic Messages route. Linux & macOS (no Windows).
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
