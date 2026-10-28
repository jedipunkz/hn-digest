---
source: "https://github.com/carlostapiacl/sandbox-conformance"
hn_url: "https://news.ycombinator.com/item?id=50269082"
title: "I ran each AI coding-agent sandbox's own tests against the others"
article_title: "GitHub - carlostapiacl/sandbox-conformance: An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. · GitHub"
image: "https://opengraph.githubassets.com/3711630ae7c21389a9214e6f9adfe38c604caf314a0daafacb34b929e0fd993d/carlostapiacl/sandbox-conformance"
author: "carlostapia"
captured_at: "2026-10-28T04:19:57Z"
capture_tool: "hn-digest"
hn_id: 50269082
score: 1
comments: 0
posted_at: "2026-10-28T03:35:56Z"
tags:
  - hacker-news
---

# I ran each AI coding-agent sandbox's own tests against the others

- HN: [50269082](https://news.ycombinator.com/item?id=50269082)
- Source: [github.com](https://github.com/carlostapiacl/sandbox-conformance)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T03:35:56Z

## Translation

Title: I ran each AI coding-agent sandbox's own tests against the others
Article title: GitHub - carlostapiacl/sandbox-conformance: An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. · GitHub
Description: An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. - carlostapiacl/sandbox-conformance

Article text:
GitHub - carlostapiacl/sandbox-conformance: An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. · GitHub
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
carlostapiacl
sandbox-conformance Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison.
carlostapiacl.github.io/sandbox-conformance/ Topics
Readme Apache-2.0 license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
26 Commits 26 Commits Folders and files
.github/ workflows .github/ workflows contracts contracts corpus corpus docs-es docs-es docs docs i18n/ es i18n/ es results results runner runner sandboxes sandboxes spec spec .gitignore .gitignore .nojekyll .nojekyll CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SOURCES.md SOURCES.md corpus-roadmap.md corpus-roadmap.md index.html index.html mkdocs.es.yml mkdocs.es.yml mkdocs.yml mkdocs.yml View all files Repository files navigation
An open lab for measuring sandboxes for AI coding agents.
Every sandbox makes claims: "the agent can't read your keys", "it can't reach the
network", "it can't plant a git hook". This repository turns those claims into
contracts anyone can run , and keeps the results of running them against each
sandbox, so you can compare them on the same tests, with the evidence attached.
It is not tied to any one sandbox. Add a contract, measure your sandbox, publish
the result.
Documentation: carlostapiacl.github.io/sandbox-conformance/docs :
quickstart, concepts, how-to guides, methodology, every run and the contract catalog.
In Spanish at /docs/es : the
translations are in docs-es/ and i18n/es/ (glossary in i18n/es/GLOSARIO.md ), and a page
whose English original changed since it was translated says so until it is re-translated and
stamped ( python3 i18n/es/stamp.py ).
folder
what
contracts/
the tests, one TOML file each, grouped by topic, each crediting its source
corpus/
the corpus versions: the exact contract files of each, by hash
spec/
the contract format , explained for people who want to write one, and what a result means (draft)
sandboxes/
one page per sandbox: how a contract maps to its configuration, and its results by version
results/
published runs with their raw evidence, and the generated comparison table
What a contract is
Three parts, so a result is never just "it returned an error":
attack : something the agent must not be able to do;
control : legitimate work that must keep working (a sandbox that blocks
everything fails here);
effect : what is checked on disk or on the wire afterwards (the file is
unchanged, the secret never appeared).
Before a contract counts against a sandbox, its attack must work without a
sandbox on the same host. Otherwise a "blocked" result would prove nothing.
The corpus is versioned: v1.1 (2026-10-26) is 196 contracts, harvested from the
public test suites of sandbox-runtime, nono, Codex, Deno, bubblewrap and OpenHands, with
by-omission variants and six masked-credential contracts, each re-expressed in the
neutral format and credited in SOURCES.md and its own source field.
Which contracts a version is, and when it moves, is in corpus/ .
Maintained by the author of TartarOS. TartarOS is where I try each escape the lab
finds, to see whether it can be mitigated, so it is built against these contracts and
its column should be read that way. That is also why the contracts come from the other
projects' test suites. TartarOS is measured like the others since the
2026-10-22 run , and its
page and result file record the conflict of interest.
Contracts and adapters are welcome: see CONTRIBUTING.md .
Reports I sent to each vendor's security program, as of 2026-10-27:
The details of an open report are not published until the vendor fixes it
( SECURITY.md ).
First results in: results/2026-10-15/ measures
sandbox-runtime 0.0.79, nono 0.80.0 and Codex CLI 0.162.1 on Linux (kernel 6.8),
with the runner and adapters that produced them and a page per
sandbox in sandboxes/ .
A fourth sandbox, landrun (pure Landlock, MIT), is now
measured alongside the three in a four-column run .
Adding a sandbox is one adapter module in runner/adapters/ ; the engine and the
contracts are untouched.
The by-omission run re-measures the
read-deny contracts in a capability sandbox's own model (grant the workspace,
omit the secret), so nono is measured where the neutral model only returned
unsupported .
The 2026-10-16 run is the first judged by
spec/verdicts.md . The old runner reproduces the 2026-10-15
results on a fresh host; the new one changes five, each traced to a cause. The lab
tests itself against reference sandboxes whose behaviour
is known ( python3 -m unittest discover -s runner/tests ).
The 2026-10-21 run measures the latest releases:
sandbox-runtime 0.0.81 and Codex CLI 0.163.0 (nono and landrun unchanged). The
previous versions were re-measured first and reproduced. Which results changed in which
version is in results/versions.md .
The 2026-10-22 run adds TartarOS 0.1.0-alpha.10,
on the same host, contracts and lab tools as the 2026-10-21 run. The
alpha.11 run the same day measures the
[broker] contracts for the first time: their baseline now runs against a
reference broker with no defences, the same for every
sandbox.
The 2026-10-26 run measures TartarOS
0.1.0-alpha.16 and sandbox-runtime 0.0.81 on six new masked-credential contracts: both
contain all six on Linux.
Next: macOS runs and more corpora harvested from other sandboxes' test suites.
The results and what they mean are written up at carlostapia.cl .
An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison.
carlostapiacl.github.io/sandbox-conformance/ Topics
Readme Apache-2.0 license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. - carlostapiacl/sandbox-conformance

GitHub - carlostapiacl/sandbox-conformance: An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison. · GitHub
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
carlostapiacl
sandbox-conformance Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison.
carlostapiacl.github.io/sandbox-conformance/ Topics
Readme Apache-2.0 license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
26 Commits 26 Commits Folders and files
.github/ workflows .github/ workflows contracts contracts corpus corpus docs-es docs-es docs docs i18n/ es i18n/ es results results runner runner sandboxes sandboxes spec spec .gitignore .gitignore .nojekyll .nojekyll CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SOURCES.md SOURCES.md corpus-roadmap.md corpus-roadmap.md index.html index.html mkdocs.es.yml mkdocs.es.yml mkdocs.yml mkdocs.yml View all files Repository files navigation
An open lab for measuring sandboxes for AI coding agents.
Every sandbox makes claims: "the agent can't read your keys", "it can't reach the
network", "it can't plant a git hook". This repository turns those claims into
contracts anyone can run , and keeps the results of running them against each
sandbox, so you can compare them on the same tests, with the evidence attached.
It is not tied to any one sandbox. Add a contract, measure your sandbox, publish
the result.
Documentation: carlostapiacl.github.io/sandbox-conformance/docs :
quickstart, concepts, how-to guides, methodology, every run and the contract catalog.
In Spanish at /docs/es : the
translations are in docs-es/ and i18n/es/ (glossary in i18n/es/GLOSARIO.md ), and a page
whose English original changed since it was translated says so until it is re-translated and
stamped ( python3 i18n/es/stamp.py ).
folder
what
contracts/
the tests, one TOML file each, grouped by topic, each crediting its source
corpus/
the corpus versions: the exact contract files of each, by hash
spec/
the contract format , explained for people who want to write one, and what a result means (draft)
sandboxes/
one page per sandbox: how a contract maps to its configuration, and its results by version
results/
published runs with their raw evidence, and the generated comparison table
What a contract is
Three parts, so a result is never just "it returned an error":
attack : something the agent must not be able to do;
control : legitimate work that must keep working (a sandbox that blocks
everything fails here);
effect : what is checked on disk or on the wire afterwards (the file is
unchanged, the secret never appeared).
Before a contract counts against a sandbox, its attack must work without a
sandbox on the same host. Otherwise a "blocked" result would prove nothing.
The corpus is versioned: v1.1 (2026-10-26) is 196 contracts, harvested from the
public test suites of sandbox-runtime, nono, Codex, Deno, bubblewrap and OpenHands, with
by-omission variants and six masked-credential contracts, each re-expressed in the
neutral format and credited in SOURCES.md and its own source field.
Which contracts a version is, and when it moves, is in corpus/ .
Maintained by the author of TartarOS. TartarOS is where I try each escape the lab
finds, to see whether it can be mitigated, so it is built against these contracts and
its column should be read that way. That is also why the contracts come from the other
projects' test suites. TartarOS is measured like the others since the
2026-10-22 run , and its
page and result file record the conflict of interest.
Contracts and adapters are welcome: see CONTRIBUTING.md .
Reports I sent to each vendor's security program, as of 2026-10-27:
The details of an open report are not published until the vendor fixes it
( SECURITY.md ).
First results in: results/2026-10-15/ measures
sandbox-runtime 0.0.79, nono 0.80.0 and Codex CLI 0.162.1 on Linux (kernel 6.8),
with the runner and adapters that produced them and a page per
sandbox in sandboxes/ .
A fourth sandbox, landrun (pure Landlock, MIT), is now
measured alongside the three in a four-column run .
Adding a sandbox is one adapter module in runner/adapters/ ; the engine and the
contracts are untouched.
The by-omission run re-measures the
read-deny contracts in a capability sandbox's own model (grant the workspace,
omit the secret), so nono is measured where the neutral model only returned
unsupported .
The 2026-10-16 run is the first judged by
spec/verdicts.md . The old runner reproduces the 2026-10-15
results on a fresh host; the new one changes five, each traced to a cause. The lab
tests itself against reference sandboxes whose behaviour
is known ( python3 -m unittest discover -s runner/tests ).
The 2026-10-21 run measures the latest releases:
sandbox-runtime 0.0.81 and Codex CLI 0.163.0 (nono and landrun unchanged). The
previous versions were re-measured first and reproduced. Which results changed in which
version is in results/versions.md .
The 2026-10-22 run adds TartarOS 0.1.0-alpha.10,
on the same host, contracts and lab tools as the 2026-10-21 run. The
alpha.11 run the same day measures the
[broker] contracts for the first time: their baseline now runs against a
reference broker with no defences, the same for every
sandbox.
The 2026-10-26 run measures TartarOS
0.1.0-alpha.16 and sandbox-runtime 0.0.81 on six new masked-credential contracts: both
contain all six on Linux.
Next: macOS runs and more corpora harvested from other sandboxes' test suites.
The results and what they mean are written up at carlostapia.cl .
An open lab for measuring sandboxes for AI coding agents: runnable contracts, per-sandbox results, side-by-side comparison.
carlostapiacl.github.io/sandbox-conformance/ Topics
Readme Apache-2.0 license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
