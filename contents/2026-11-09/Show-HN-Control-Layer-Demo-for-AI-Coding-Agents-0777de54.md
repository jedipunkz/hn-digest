---
source: "https://github.com/MintCoderpro/mintigate-demo"
hn_url: "https://news.ycombinator.com/item?id=50422833"
title: "Show HN: Control Layer Demo for AI Coding Agents"
article_title: "GitHub - MintCoderpro/mintigate-demo: Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. · GitHub"
image: "https://opengraph.githubassets.com/a77d76d13673e15e21690b60e9d3485af83059a2b21bacb1b16aea05b8df73f5/MintCoderpro/mintigate-demo"
author: "MintCoder"
captured_at: "2026-11-09T09:15:08Z"
capture_tool: "hn-digest"
hn_id: 50422833
score: 1
comments: 0
posted_at: "2026-11-09T08:43:54Z"
tags:
  - hacker-news
---

# Show HN: Control Layer Demo for AI Coding Agents

- HN: [50422833](https://news.ycombinator.com/item?id=50422833)
- Source: [github.com](https://github.com/MintCoderpro/mintigate-demo)
- Score: 1
- Comments: 0
- Posted: 2026-11-09T08:43:54Z

## Translation

Title: Show HN: Control Layer Demo for AI Coding Agents
Article title: GitHub - MintCoderpro/mintigate-demo: Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. · GitHub
Description: Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. - MintCoderpro/mintigate-demo
HN text: Letting coding agents (such as Claude Code) work autonomously on repos on a daily-basis, brings a high level of uncertainty. For those who won't let the agent decide alone whether it can create/modify or delete any folder on their repos; I wanted to be able to trace actions performed and being able to let the agent continue his work autonomously, or ask him to request a human approval according to some criteria defined.
After investigations, the cleanest way to do it was according to me to develop a deterministic control layer between the LLM and the work of the agent that would track and catch potential non-respect of pre-defined rules.
How? by first setting up approval rules (Pass/requires approval/refused) before the work of the agent, and afterwards controlling cryptographic SHA 256 verifications, etc.
This allows to compare the before-after work of the agent.
I have put in my public GitHub repo a demo that you can launch easily and observe different scenarios behaviors like applied authorized modification, protected file waiting for Human approval, or outside perimeter authorization modification refused. The demo is offline, does not involve the official runtime, in native node.js and without dependencies. It should still allow you to understand the core principles.
I would like to know: how do you implement and handle safety boundaries? Especially if you let agents work on critical data/sensitive codebases? With which guardrails?

Article text:
GitHub - MintCoderpro/mintigate-demo: Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. · GitHub
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
MintCoderpro
mintigate-demo Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
fixture/ src fixture/ src LICENSE LICENSE README.md README.md SHA256SUMS SHA256SUMS demo-runtime.mjs demo-runtime.mjs mintigate-demo.mjs mintigate-demo.mjs scenarios.json scenarios.json smoke-test.mjs smoke-test.mjs View all files Repository files navigation
Offline demonstration — simulated agent, simplified demo runtime. Proprietary Mintigate Runtime not executed.
This folder is a small, deterministic demonstration that you can read in full and run on your own computer. A simulated agent proposes three fixed changes to a synthetic two-file fixture. A simplified demo runtime decides on each change in a temporary copy of the fixture, and an independent check confirms what happened.
#
Proposed change
Expected decision
Written?
1
Set --mint-demo-accent from 0 to 1 in src/styles/base.css , inside the declared scope
APPLIED
Yes, in the temporary copy only
2
Change src/js/story.js , a protected file
PENDING_APPROVAL
No: a human approval is required, and this demo gives none
3
Change config/deploy.json , outside the declared scope
REFUSED
No
Each decision is recorded with the SHA-256 of the target before and after, in a report.
Node.js 18 or later. Nothing else: no package to install, no Git, no Claude Code, no SDK, no network access, no account and no model call.
Open a terminal in this folder (on Windows, PowerShell or Windows Terminal) and run:
node mintigate-demo.mjs
The demo takes no option. There is no live mode.
Code
Result shown
Meaning
0
VERIFIED
The three expected behaviours were verified: one change applied, one waiting for approval, one refused. Code 0 does not mean that every proposed action was allowed.
1
NOT VERIFIED or STOPPED ON AN ERROR
A check failed, or the demo stopped on an error. The console and the report say which.
2
NOT RUN
An option was given, Node.js is too old, or a file of this folder is missing or unreadable. Nothing was written.
What it writes
Each run creates a new folder .output/run-<date>-<random>/ in this folder. Earlier runs are never changed.
The fixture is copied into workspace/ inside that new folder, and the changes are made there only. The shipped fixture/ folder is never written.
At the end of the run, even after an error, the demo removes the files and folders of that workspace/ one by one. It removes nothing else.
What remains is one report: .output/run-<date>-<random>/report.json . It holds the decisions, the audit records, the hashes, the checks, the negative controls, the clean-up count and the hashes of the demo files that were used. It contains no absolute path of your computer.
Nothing is written outside .output/ , and no setting of your computer is changed. To remove everything, delete this folder.
smoke-test.mjs does not use demo-runtime.mjs to judge the outcome. It copies the fixture, reads the files itself, computes their hashes and compares them with the expectations in scenarios.json :
the decision of each scenario;
for the applied change: exactly the expected line changed, and nothing else;
for the other two: the target is unchanged, or was not created;
the hashes in each audit record match the files;
every protected file and every file outside the scope is unchanged, no file was created or deleted, the shipped fixture is unchanged, and the temporary copy was removed.
It then runs negative controls: deliberately wrong expectations or observations, such as expecting the refused change to be applied or pretending that the protected file changed. Each one must make a named check fail. If a contradiction went unnoticed, the run ends with code 1.
File
Role
mintigate-demo.mjs
Entry point and console summary
demo-runtime.mjs
Simplified scope, protection and application rules
smoke-test.mjs
Preparation of the working copy and independent checks
scenarios.json
The simulated proposals and the expectations
fixture/src/styles/base.css , fixture/src/js/story.js
Synthetic files written for this demo
LICENSE
MIT licence for this kit
SHA256SUMS
SHA-256 of the files above
Limits
The agent is simulated: its three proposals are fixed in scenarios.json . No model is called.
The runtime is a simplified demo runtime of a few rules: a folder or file scope, a list of protected files, and exact text replacements. It is not the proprietary Mintigate Runtime, which this demo does not contain or execute.
The demo works only on its own synthetic fixture. It does not run on your repository, does not watch or intercept your tools, and has no approval command.
It is not a benchmark and measures no performance.
The MIT licence in LICENSE applies to the files of this kit only.
In a separate copy of the demo, open scenarios.json . For protected-file-change , change expect.decision from PENDING_APPROVAL to APPLIED , then run:
The demo should exit with code 1: the protected change still requires approval, and the independent check rejects your altered expectation. Inspect the generated report.json to see the decision, failed checks, and file hashes.
You control this local demo and can change its policies or code. It illustrates governance rules; it does not protect against its own administrator.
This demo illustrates scoped changes, human approval boundaries, and audit records using a deterministic agent and a simplified runtime.
Mintigate is being developed to help developers govern real agent workflows. This demo does not execute its proprietary Runtime or connect to a live model.
Explore the project, current evaluation results, and limitations: https://mintigate.com
Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. - MintCoderpro/mintigate-demo

Letting coding agents (such as Claude Code) work autonomously on repos on a daily-basis, brings a high level of uncertainty. For those who won't let the agent decide alone whether it can create/modify or delete any folder on their repos; I wanted to be able to trace actions performed and being able to let the agent continue his work autonomously, or ask him to request a human approval according to some criteria defined.
After investigations, the cleanest way to do it was according to me to develop a deterministic control layer between the LLM and the work of the agent that would track and catch potential non-respect of pre-defined rules.
How? by first setting up approval rules (Pass/requires approval/refused) before the work of the agent, and afterwards controlling cryptographic SHA 256 verifications, etc.
This allows to compare the before-after work of the agent.
I have put in my public GitHub repo a demo that you can launch easily and observe different scenarios behaviors like applied authorized modification, protected file waiting for Human approval, or outside perimeter authorization modification refused. The demo is offline, does not involve the official runtime, in native node.js and without dependencies. It should still allow you to understand the core principles.
I would like to know: how do you implement and handle safety boundaries? Especially if you let agents work on critical data/sensitive codebases? With which guardrails?

GitHub - MintCoderpro/mintigate-demo: Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model. · GitHub
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
MintCoderpro
mintigate-demo Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
fixture/ src fixture/ src LICENSE LICENSE README.md README.md SHA256SUMS SHA256SUMS demo-runtime.mjs demo-runtime.mjs mintigate-demo.mjs mintigate-demo.mjs scenarios.json scenarios.json smoke-test.mjs smoke-test.mjs View all files Repository files navigation
Offline demonstration — simulated agent, simplified demo runtime. Proprietary Mintigate Runtime not executed.
This folder is a small, deterministic demonstration that you can read in full and run on your own computer. A simulated agent proposes three fixed changes to a synthetic two-file fixture. A simplified demo runtime decides on each change in a temporary copy of the fixture, and an independent check confirms what happened.
#
Proposed change
Expected decision
Written?
1
Set --mint-demo-accent from 0 to 1 in src/styles/base.css , inside the declared scope
APPLIED
Yes, in the temporary copy only
2
Change src/js/story.js , a protected file
PENDING_APPROVAL
No: a human approval is required, and this demo gives none
3
Change config/deploy.json , outside the declared scope
REFUSED
No
Each decision is recorded with the SHA-256 of the target before and after, in a report.
Node.js 18 or later. Nothing else: no package to install, no Git, no Claude Code, no SDK, no network access, no account and no model call.
Open a terminal in this folder (on Windows, PowerShell or Windows Terminal) and run:
node mintigate-demo.mjs
The demo takes no option. There is no live mode.
Code
Result shown
Meaning
0
VERIFIED
The three expected behaviours were verified: one change applied, one waiting for approval, one refused. Code 0 does not mean that every proposed action was allowed.
1
NOT VERIFIED or STOPPED ON AN ERROR
A check failed, or the demo stopped on an error. The console and the report say which.
2
NOT RUN
An option was given, Node.js is too old, or a file of this folder is missing or unreadable. Nothing was written.
What it writes
Each run creates a new folder .output/run-<date>-<random>/ in this folder. Earlier runs are never changed.
The fixture is copied into workspace/ inside that new folder, and the changes are made there only. The shipped fixture/ folder is never written.
At the end of the run, even after an error, the demo removes the files and folders of that workspace/ one by one. It removes nothing else.
What remains is one report: .output/run-<date>-<random>/report.json . It holds the decisions, the audit records, the hashes, the checks, the negative controls, the clean-up count and the hashes of the demo files that were used. It contains no absolute path of your computer.
Nothing is written outside .output/ , and no setting of your computer is changed. To remove everything, delete this folder.
smoke-test.mjs does not use demo-runtime.mjs to judge the outcome. It copies the fixture, reads the files itself, computes their hashes and compares them with the expectations in scenarios.json :
the decision of each scenario;
for the applied change: exactly the expected line changed, and nothing else;
for the other two: the target is unchanged, or was not created;
the hashes in each audit record match the files;
every protected file and every file outside the scope is unchanged, no file was created or deleted, the shipped fixture is unchanged, and the temporary copy was removed.
It then runs negative controls: deliberately wrong expectations or observations, such as expecting the refused change to be applied or pretending that the protected file changed. Each one must make a named check fail. If a contradiction went unnoticed, the run ends with code 1.
File
Role
mintigate-demo.mjs
Entry point and console summary
demo-runtime.mjs
Simplified scope, protection and application rules
smoke-test.mjs
Preparation of the working copy and independent checks
scenarios.json
The simulated proposals and the expectations
fixture/src/styles/base.css , fixture/src/js/story.js
Synthetic files written for this demo
LICENSE
MIT licence for this kit
SHA256SUMS
SHA-256 of the files above
Limits
The agent is simulated: its three proposals are fixed in scenarios.json . No model is called.
The runtime is a simplified demo runtime of a few rules: a folder or file scope, a list of protected files, and exact text replacements. It is not the proprietary Mintigate Runtime, which this demo does not contain or execute.
The demo works only on its own synthetic fixture. It does not run on your repository, does not watch or intercept your tools, and has no approval command.
It is not a benchmark and measures no performance.
The MIT licence in LICENSE applies to the files of this kit only.
In a separate copy of the demo, open scenarios.json . For protected-file-change , change expect.decision from PENDING_APPROVAL to APPLIED , then run:
The demo should exit with code 1: the protected change still requires approval, and the independent check rejects your altered expectation. Inspect the generated report.json to see the decision, failed checks, and file hashes.
You control this local demo and can change its policies or code. It illustrates governance rules; it does not protect against its own administrator.
This demo illustrates scoped changes, human approval boundaries, and audit records using a deterministic agent and a simplified runtime.
Mintigate is being developed to help developers govern real agent workflows. This demo does not execute its proprietary Runtime or connect to a live model.
Explore the project, current evaluation results, and limitations: https://mintigate.com
Offline CLI demo of scoped changes, approval boundaries, and audit records. Simulated agent; no live model.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
