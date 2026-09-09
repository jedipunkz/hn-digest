---
source: "https://github.com/quazardous/jobbox"
hn_url: "https://news.ycombinator.com/item?id=49622255"
title: "AI agent spends too much time waiting"
article_title: "GitHub - quazardous/jobbox: Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. · GitHub"
image: "https://opengraph.githubassets.com/affc4a85e3697edf953c20760f17faacb521cd9cefe41ec026691909fd293517/quazardous/jobbox"
author: "quazarzero"
captured_at: "2026-09-09T07:07:35Z"
capture_tool: "hn-digest"
hn_id: 49622255
score: 1
comments: 0
posted_at: "2026-09-09T06:54:34Z"
tags:
  - hacker-news
---

# AI agent spends too much time waiting

- HN: [49622255](https://news.ycombinator.com/item?id=49622255)
- Source: [github.com](https://github.com/quazardous/jobbox)
- Score: 1
- Comments: 0
- Posted: 2026-09-09T06:54:34Z

## Translation

Title: AI agent spends too much time waiting
Article title: GitHub - quazardous/jobbox: Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. · GitHub
Description: Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. - quazardous/jobbox

Article text:
GitHub - quazardous/jobbox: Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. · GitHub
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
quazardous
/
jobbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
126 Commits 126 Commits Folders and files
.claude-plugin .claude-plugin .github .github bin bin demo demo plugin plugin src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CLI-AI.md CLI-AI.md CODE-SIGNING-POLICY.md CODE-SIGNING-POLICY.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE MACOS.md MACOS.md README.md README.md TROUBLESHOOTING.md TROUBLESHOOTING.md USAGE.md USAGE.md WINDOWS.md WINDOWS.md build.rs build.rs install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Time is money. Your agent spends both, standing still.
A five-minute build runs. The agent waits. You wait. Nothing else
happens — and you are billed for all of it twice: your hour, and the
tokens burning in a session doing nothing.
No single wait is worth stopping for. It is their sum that costs , and
the sum is invisible until something counts it.
$ jbx gain
jbx gain — since the beginning
commands wrapped 142
of those, detached 11 (7.7% of them)
they took 3h04m
you stood still 35m12s
given back 2h29m ███████████████████░░░░░ 81.0%
project calls detached elapsed waited saved impact
acme 142 11 3h04m 35m12s 2h29m (81%) ██████████
api 96 7 2h11m 18m03s 1h53m (86%) ████████░░
front 34 4 39m 3m17s 36m (92%) ██░░░░░░░░
last hour 6 calls · 1 detached · 8m12s saved (73%)
last day 38 calls · 4 detached · 40m05s saved (77%)
all 142 calls · 11 detached · 2h29m saved (81%)
That is one week. Put your own rate on it.
JobBox wraps every command your agent runs. The quick ones come back
untouched — output as written, exit code unchanged, as though nothing
were there. The slow ones detach themselves , say so, and tell whoever
needs to know when they end.
Nobody judges in advance which is which. That judgement is the thing
everybody gets wrong, so JobBox does not make it: it runs the line and
finds out.
One binary. Rust, serde_json , nothing else. Linux, macOS, Windows.
$ jbx hook --list
claude Bash PreToolUse ~/.claude/settings.json
gemini run_shell_command BeforeTool ~/.gemini/settings.json
droid Execute PreToolUse declare by hand
cursor Shell preToolUse declare by hand (no unasked endings)
copilot bash preToolUse declare by hand
They agree on almost nothing — not the name of the shell tool, not the
event, not the shape of the answer, not the file it is declared in.
jbx init handles the first two; the rest take one block of JSON, and
CLI-AI.md has the exact block for each , along with the
two it deliberately does not support and why.
As a Claude Code plugin — the hooks, the binary and a background watch in
one thing:
$ claude plugin marketplace add quazardous/jobbox
$ claude plugin install jbx@jobbox
Or as a command, anywhere:
$ curl -fsSL https://raw.githubusercontent.com/quazardous/jobbox/main/install.sh | sh
It checks the download against the sums published with the release, puts
it on your PATH , and asks before declaring its hooks — they go in a
settings file other tools share. On Windows, irm https://raw.githubusercontent.com/quazardous/jobbox/main/install.ps1 | iex , and WINDOWS.md has the rest.
Run jbx init as well if you have rtk .
A plugin declares hooks; it cannot displace somebody else's, and two
hooks rewriting one field is a race no harness documents. init settles
that by calling rtk itself.
Nothing changes — until something is slow:
$ npm run build
> building…
jbx: this passed 30s, so it is in the BACKGROUND as j7f3a91c — nothing lost.
DO NOT WAIT FOR IT, DO SOMETHING ELSE. With nothing else: Monitor
`jbx wait j7f3a91c`, which ends when the job does.
jbx help j7f3a91c
The build output arrived as it was written , not replayed at the end.
The ending reaches you two ways, and one is better. Left alone, it is
announced on the next turn — free, and it costs that delay. Monitored, it
arrives the moment it happens: jbx wait <id> ends exactly when the job
does, so anything watching it is woken then. jbx watch does that for
every job at once, one line each.
That is the difference between waiting and being told. Polling is
neither — it is waiting with extra steps.
The one judgement left to make
The old answer to "when should this go to the background?" was a document
telling an agent to estimate how long a command would take. Agents get
that wrong, and so do people. jbx removes the question and leaves a
smaller one:
Do you need this result before you can do anything else ?
Almost always, no. When the answer is yes, say so — jbx fg -- '<line>'
runs without ever letting go, and jbx gain counts what that cost.
USAGE.md — every verb, every setting, how saved is
counted and why it is a ceiling, and what was deliberately left out.
CLI-AI.md — every agent CLI jbx answers, what to paste
where, and the two it leaves alone.
MACOS.md — why the install command never trips
Gatekeeper, and what to do if you downloaded by hand.
WINDOWS.md — Smart App Control, which shell runs your
commands, and the two things that differ there.
TROUBLESHOOTING.md — why nothing is
detaching, which is usually on purpose.
CONTRIBUTING.md — what belongs in a test, and
what has already been ruled out.
CODE-SIGNING-POLICY.md — who may release
this, what it never sends anywhere, and how to undo everything it does.
CHANGELOG.md — what changed, and why it mattered.
Made with simai-cli , from
demo/detach.txt .
Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. - quazardous/jobbox

GitHub - quazardous/jobbox: Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows. · GitHub
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
quazardous
/
jobbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
126 Commits 126 Commits Folders and files
.claude-plugin .claude-plugin .github .github bin bin demo demo plugin plugin src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CLI-AI.md CLI-AI.md CODE-SIGNING-POLICY.md CODE-SIGNING-POLICY.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE MACOS.md MACOS.md README.md README.md TROUBLESHOOTING.md TROUBLESHOOTING.md USAGE.md USAGE.md WINDOWS.md WINDOWS.md build.rs build.rs install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Time is money. Your agent spends both, standing still.
A five-minute build runs. The agent waits. You wait. Nothing else
happens — and you are billed for all of it twice: your hour, and the
tokens burning in a session doing nothing.
No single wait is worth stopping for. It is their sum that costs , and
the sum is invisible until something counts it.
$ jbx gain
jbx gain — since the beginning
commands wrapped 142
of those, detached 11 (7.7% of them)
they took 3h04m
you stood still 35m12s
given back 2h29m ███████████████████░░░░░ 81.0%
project calls detached elapsed waited saved impact
acme 142 11 3h04m 35m12s 2h29m (81%) ██████████
api 96 7 2h11m 18m03s 1h53m (86%) ████████░░
front 34 4 39m 3m17s 36m (92%) ██░░░░░░░░
last hour 6 calls · 1 detached · 8m12s saved (73%)
last day 38 calls · 4 detached · 40m05s saved (77%)
all 142 calls · 11 detached · 2h29m saved (81%)
That is one week. Put your own rate on it.
JobBox wraps every command your agent runs. The quick ones come back
untouched — output as written, exit code unchanged, as though nothing
were there. The slow ones detach themselves , say so, and tell whoever
needs to know when they end.
Nobody judges in advance which is which. That judgement is the thing
everybody gets wrong, so JobBox does not make it: it runs the line and
finds out.
One binary. Rust, serde_json , nothing else. Linux, macOS, Windows.
$ jbx hook --list
claude Bash PreToolUse ~/.claude/settings.json
gemini run_shell_command BeforeTool ~/.gemini/settings.json
droid Execute PreToolUse declare by hand
cursor Shell preToolUse declare by hand (no unasked endings)
copilot bash preToolUse declare by hand
They agree on almost nothing — not the name of the shell tool, not the
event, not the shape of the answer, not the file it is declared in.
jbx init handles the first two; the rest take one block of JSON, and
CLI-AI.md has the exact block for each , along with the
two it deliberately does not support and why.
As a Claude Code plugin — the hooks, the binary and a background watch in
one thing:
$ claude plugin marketplace add quazardous/jobbox
$ claude plugin install jbx@jobbox
Or as a command, anywhere:
$ curl -fsSL https://raw.githubusercontent.com/quazardous/jobbox/main/install.sh | sh
It checks the download against the sums published with the release, puts
it on your PATH , and asks before declaring its hooks — they go in a
settings file other tools share. On Windows, irm https://raw.githubusercontent.com/quazardous/jobbox/main/install.ps1 | iex , and WINDOWS.md has the rest.
Run jbx init as well if you have rtk .
A plugin declares hooks; it cannot displace somebody else's, and two
hooks rewriting one field is a race no harness documents. init settles
that by calling rtk itself.
Nothing changes — until something is slow:
$ npm run build
> building…
jbx: this passed 30s, so it is in the BACKGROUND as j7f3a91c — nothing lost.
DO NOT WAIT FOR IT, DO SOMETHING ELSE. With nothing else: Monitor
`jbx wait j7f3a91c`, which ends when the job does.
jbx help j7f3a91c
The build output arrived as it was written , not replayed at the end.
The ending reaches you two ways, and one is better. Left alone, it is
announced on the next turn — free, and it costs that delay. Monitored, it
arrives the moment it happens: jbx wait <id> ends exactly when the job
does, so anything watching it is woken then. jbx watch does that for
every job at once, one line each.
That is the difference between waiting and being told. Polling is
neither — it is waiting with extra steps.
The one judgement left to make
The old answer to "when should this go to the background?" was a document
telling an agent to estimate how long a command would take. Agents get
that wrong, and so do people. jbx removes the question and leaves a
smaller one:
Do you need this result before you can do anything else ?
Almost always, no. When the answer is yes, say so — jbx fg -- '<line>'
runs without ever letting go, and jbx gain counts what that cost.
USAGE.md — every verb, every setting, how saved is
counted and why it is a ceiling, and what was deliberately left out.
CLI-AI.md — every agent CLI jbx answers, what to paste
where, and the two it leaves alone.
MACOS.md — why the install command never trips
Gatekeeper, and what to do if you downloaded by hand.
WINDOWS.md — Smart App Control, which shell runs your
commands, and the two things that differ there.
TROUBLESHOOTING.md — why nothing is
detaching, which is usually on purpose.
CONTRIBUTING.md — what belongs in a test, and
what has already been ruled out.
CODE-SIGNING-POLICY.md — who may release
this, what it never sends anywhere, and how to undo everything it does.
CHANGELOG.md — what changed, and why it mattered.
Made with simai-cli , from
demo/detach.txt .
Every command your agent runs is wrapped; the slow ones put themselves in the BACKGROUND and tell you when they end. `jbx stats` says how much time that saved. One Rust binary — Linux, macOS, Windows.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
