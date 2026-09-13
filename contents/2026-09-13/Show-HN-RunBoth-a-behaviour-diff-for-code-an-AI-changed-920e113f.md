---
source: "https://github.com/runboth/runboth"
hn_url: "https://news.ycombinator.com/item?id=49686484"
title: "Show HN: RunBoth, a behaviour diff for code an AI changed"
article_title: "GitHub - runboth/runboth: An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. · GitHub"
image: "https://repository-images.githubusercontent.com/1368617230/c4600b93-1493-487e-9a7f-5ce920dfc677"
author: "kyleclouthier"
captured_at: "2026-09-13T17:56:06Z"
capture_tool: "hn-digest"
hn_id: 49686484
score: 1
comments: 0
posted_at: "2026-09-13T17:40:16Z"
tags:
  - hacker-news
---

# Show HN: RunBoth, a behaviour diff for code an AI changed

- HN: [49686484](https://news.ycombinator.com/item?id=49686484)
- Source: [github.com](https://github.com/runboth/runboth)
- Score: 1
- Comments: 0
- Posted: 2026-09-13T17:40:16Z

## Translation

Title: Show HN: RunBoth, a behaviour diff for code an AI changed
Article title: GitHub - runboth/runboth: An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. · GitHub
Description: An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. - runboth/runboth

Article text:
GitHub - runboth/runboth: An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
runboth
/
runboth
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows brand brand results results runboth runboth scripts scripts tests tests .gitattributes .gitattributes .gitignore .gitignore CAUGHT.md CAUGHT.md COVERAGE_2026-09-12.md COVERAGE_2026-09-12.md HOW_IT_WORKS.md HOW_IT_WORKS.md LICENSE.md LICENSE.md README.md README.md RED_TEAM_2026-09-12.md RED_TEAM_2026-09-12.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
An AI changed your code. RunBoth runs both versions and tells you what actually behaves
differently, including the functions nobody touched.
runboth.dev ·
How it works ·
Red team results
$ git commit -m "refactor: tidy up the rates module"
BLOCKED: the behaviour changed and your message does not say so.
rate(100)
used to: return 0.1
now: return 0.0
AND 1 function you did NOT touch now behaves differently,
because it calls what you changed:
total(2.5, 100) pkg/invoice.py
used to: return 225.0
now: return 250.0
What it does
It checks out both versions of your code, generates inputs for every changed function from its
signature and from the constants mined out of its own bytecode, runs both versions in separate
sandboxed subprocesses, and compares seven observation channels. When they disagree it hands you
the exact input that separates them.
No test suite required. No network calls. No AI model. No dependencies.
pip install runboth # once the first release is on PyPI
pip install git+https://github.com/runboth/runboth # works today
runboth install-hook # a commit-msg gate, silent unless behaviour moved
As a GitHub Action, running on your own runners:
- uses : runboth/runboth@v0.1.0
with :
budget : 60
Three verdicts, never two
verdict
meaning
changed
with a witness: the arguments, the old result, the new result
no_change at budget N
N generated inputs found no difference across seven channels
abstained
it could not be checked, and here is the reason
"Cannot tell" and "no difference" are different claims, and collapsing them into a green check is
how tools end up lying. RunBoth never says safe.
Return value · exception raised · warnings · stdout · stderr · argument mutation · object state.
A narrow definition of behaviour does not under-report, it lies, because whatever sits outside the
definition comes back as no_change .
Not a model checker. Kani and CBMC translate code into logic, let inputs be unconstrained
symbols, and ask a solver whether a bad state is reachable within a bound. They return a proof.
RunBoth executes real code on concrete values. It finds differences and reproduces them; it
cannot prove absence, and never claims to.
Not mutation testing. Mutation testing damages your code to score your test suite. RunBoth
damages nothing; both versions come from your git history, and no test suite is needed.
Red-teamed against eight public repositories it had never been tuned on, with an automated oracle
built to catch the tool lying. 2,548 functions, zero false positives. Full method and numbers
in RED_TEAM_2026-09-12.md .
An adversarial corpus of 22 functions written specifically to induce false positives (object
addresses in default repr , datetime.now , unseeded random , uuid4 , os.getpid , set
iteration order, mutable defaults, generators, __file__ paths) produced none.
variable
effect
RUNBOTH_SKIP=1
let a commit through without checking it
RUNBOTH_BUDGET
generated inputs per function (gate default 80)
RUNBOTH_WORKERS
parallel adjudications, default 4
RUNBOTH_ALL_PATHS=1
also check tests, benchmarks, docs and task runners
RUNBOTH_ENGINE
engine directory, if the hook cannot resolve it
git commit --no-verify also bypasses the gate, and the gate says so itself when it blocks.
Function-level checking is Python only . Changed files in other languages are named
explicitly rather than passed over quietly.
Sampling finds differences; it cannot prove their absence.
Nondeterministic, too-slow, or unconstructible functions abstain with a reason , and are
never counted as passing.
The sandbox contains accidents: resource limits, network blocked, filesystem writes blocked. It
is not a security boundary against hostile code, and no pure-Python sandbox is.
pip install -e .
pytest tests/ -q
runboth selftest # the control suites, half of which must fail
Licence
FSL-1.1-Apache-2.0 . Free for every use except building a competing product, and it
converts to plain Apache 2.0 two years after each release.
Built by Kyle Clouthier at Clouthier Simulation Labs.
An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. - runboth/runboth

GitHub - runboth/runboth: An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
runboth
/
runboth
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows brand brand results results runboth runboth scripts scripts tests tests .gitattributes .gitattributes .gitignore .gitignore CAUGHT.md CAUGHT.md COVERAGE_2026-09-12.md COVERAGE_2026-09-12.md HOW_IT_WORKS.md HOW_IT_WORKS.md LICENSE.md LICENSE.md README.md README.md RED_TEAM_2026-09-12.md RED_TEAM_2026-09-12.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
An AI changed your code. RunBoth runs both versions and tells you what actually behaves
differently, including the functions nobody touched.
runboth.dev ·
How it works ·
Red team results
$ git commit -m "refactor: tidy up the rates module"
BLOCKED: the behaviour changed and your message does not say so.
rate(100)
used to: return 0.1
now: return 0.0
AND 1 function you did NOT touch now behaves differently,
because it calls what you changed:
total(2.5, 100) pkg/invoice.py
used to: return 225.0
now: return 250.0
What it does
It checks out both versions of your code, generates inputs for every changed function from its
signature and from the constants mined out of its own bytecode, runs both versions in separate
sandboxed subprocesses, and compares seven observation channels. When they disagree it hands you
the exact input that separates them.
No test suite required. No network calls. No AI model. No dependencies.
pip install runboth # once the first release is on PyPI
pip install git+https://github.com/runboth/runboth # works today
runboth install-hook # a commit-msg gate, silent unless behaviour moved
As a GitHub Action, running on your own runners:
- uses : runboth/runboth@v0.1.0
with :
budget : 60
Three verdicts, never two
verdict
meaning
changed
with a witness: the arguments, the old result, the new result
no_change at budget N
N generated inputs found no difference across seven channels
abstained
it could not be checked, and here is the reason
"Cannot tell" and "no difference" are different claims, and collapsing them into a green check is
how tools end up lying. RunBoth never says safe.
Return value · exception raised · warnings · stdout · stderr · argument mutation · object state.
A narrow definition of behaviour does not under-report, it lies, because whatever sits outside the
definition comes back as no_change .
Not a model checker. Kani and CBMC translate code into logic, let inputs be unconstrained
symbols, and ask a solver whether a bad state is reachable within a bound. They return a proof.
RunBoth executes real code on concrete values. It finds differences and reproduces them; it
cannot prove absence, and never claims to.
Not mutation testing. Mutation testing damages your code to score your test suite. RunBoth
damages nothing; both versions come from your git history, and no test suite is needed.
Red-teamed against eight public repositories it had never been tuned on, with an automated oracle
built to catch the tool lying. 2,548 functions, zero false positives. Full method and numbers
in RED_TEAM_2026-09-12.md .
An adversarial corpus of 22 functions written specifically to induce false positives (object
addresses in default repr , datetime.now , unseeded random , uuid4 , os.getpid , set
iteration order, mutable defaults, generators, __file__ paths) produced none.
variable
effect
RUNBOTH_SKIP=1
let a commit through without checking it
RUNBOTH_BUDGET
generated inputs per function (gate default 80)
RUNBOTH_WORKERS
parallel adjudications, default 4
RUNBOTH_ALL_PATHS=1
also check tests, benchmarks, docs and task runners
RUNBOTH_ENGINE
engine directory, if the hook cannot resolve it
git commit --no-verify also bypasses the gate, and the gate says so itself when it blocks.
Function-level checking is Python only . Changed files in other languages are named
explicitly rather than passed over quietly.
Sampling finds differences; it cannot prove their absence.
Nondeterministic, too-slow, or unconstructible functions abstain with a reason , and are
never counted as passing.
The sandbox contains accidents: resource limits, network blocked, filesystem writes blocked. It
is not a security boundary against hostile code, and no pure-Python sandbox is.
pip install -e .
pytest tests/ -q
runboth selftest # the control suites, half of which must fail
Licence
FSL-1.1-Apache-2.0 . Free for every use except building a competing product, and it
converts to plain Apache 2.0 two years after each release.
Built by Kyle Clouthier at Clouthier Simulation Labs.
An AI changed your code. RunBoth runs both versions and tells you what actually behaves differently, including the functions nobody touched.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
