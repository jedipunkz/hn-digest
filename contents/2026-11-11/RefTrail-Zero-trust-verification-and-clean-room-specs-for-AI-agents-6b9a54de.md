---
source: "https://github.com/kidimusicofficial-prog/reftrail"
hn_url: "https://news.ycombinator.com/item?id=50455326"
title: "RefTrail – Zero-trust verification and clean-room specs for AI agents"
article_title: "GitHub - kidimusicofficial-prog/reftrail: Zero-trust verification & clean-room spec protocol for coding agents. · GitHub"
image: "https://repository-images.githubusercontent.com/1453836476/87cfb6fe-9d78-4c84-8740-6077ddab73c9"
author: "SEPSolutions"
captured_at: "2026-11-11T16:55:53Z"
capture_tool: "hn-digest"
hn_id: 50455326
score: 1
comments: 0
posted_at: "2026-11-11T16:33:17Z"
tags:
  - hacker-news
---

# RefTrail – Zero-trust verification and clean-room specs for AI agents

- HN: [50455326](https://news.ycombinator.com/item?id=50455326)
- Source: [github.com](https://github.com/kidimusicofficial-prog/reftrail)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T16:33:17Z

## Translation

Title: RefTrail – Zero-trust verification and clean-room specs for AI agents
Article title: GitHub - kidimusicofficial-prog/reftrail: Zero-trust verification & clean-room spec protocol for coding agents. · GitHub
Description: Zero-trust verification & clean-room spec protocol for coding agents. - kidimusicofficial-prog/reftrail

Article text:
GitHub - kidimusicofficial-prog/reftrail: Zero-trust verification & clean-room spec protocol for coding agents. · GitHub
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
kidimusicofficial-prog
reftrail Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Zero-trust verification & clean-room spec protocol for coding agents.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples integration integration prompts prompts showcase/ uniq showcase/ uniq templates templates tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.de.md README.de.md README.en.md README.en.md README.md README.md package.json package.json View all files Repository files navigation
Build an independent implementation from documented behavior. Set the acceptance criteria before coding, then show which workflows pass and where the evidence ends.
A coding agent reports a finished feature within minutes. RefTrail makes that report checkable. Each requirement gets a source and a measurable outcome, the criteria are frozen before the first line of product code, and the final report states what was proven, by which check, and what is still open.
The tools need Node.js 22.16 or newer and no further packages.
Line Runs reproduces four stdin/stdout modes of GNU uniq . Given bolt, bolt, nut, bolt as four consecutive lines, its count mode returns three groups of 2, 1 and 1 . Run the acceptance cases from the repository root:
node showcase/uniq/verify.mjs
node tools/catalog.mjs report showcase/uniq/COVERAGE.csv --gates G01
The first command replays the acceptance cases against the frozen expectations. The second calculates the result from the catalog: 100% functional coverage, 1/1 mandatory checks passed . To compare both programs on your machine, install GNU Coreutils separately and append --reference to the first command.
Every recorded output comparison matched GNU uniq byte for byte: 16 of 16 under Linux with version 9.4 for the first criteria revision, and for revision 2 16 of 16 under Windows with version 8.32 and 16 of 16 under Linux with version 9.4. Revision 2 adds a quiet stop when a reader closes the output early. It was frozen before the program changed, and 2 of 2 early-exit checks passed for both programs on both systems. These checks cover termination and stderr, not output bytes. All four local invalid-input checks passed. An additional reviewer tested 40 more inputs across all four modes. Inspect the criteria fixed before coding , working program , filled catalog , report with timings and cost status and delivery decision .
Open the English agent brief . Supply a product and its official website. Add the desired edition, workflows, limits and budget if you know them.
Give the completed brief to your coding agent in the target directory. With only the two inputs, the first run covers one bounded end-to-end workflow on one platform, with a 60-minute limit. A larger product goal stays visible in the backlog.
Record each requirement, source and measurable outcome. Freeze the criteria before implementation and retain the digest independently. Keep progress and evidence in the coverage catalog . The brief runs the tools for these steps, so keep a downloaded copy of this repository where your agent can read it, or copy the tools folder into your project.
Run the finished workflow through its real interface. Report weighted feature coverage alongside passed mandatory checks, unresolved critical defects and the delivery decision. The catalog tool prints the first two together. A missing mandatory backup can block acceptance even at 100% weighted coverage.
Access to the reference. RefTrail organizes what you are allowed to observe. It does not analyze or open a product.
Proof of the sequence. The digest shows that the criteria did not change. That they existed before the code follows from where you kept the digest. Use a place the implementer cannot rewrite, such as a pushed commit or a message to the requester.
The reach of a result. A passed catalog covers the listed cases on the tested systems. Equivalence with the whole reference product needs more cases.
The legal review. The notes below are general information.
REA can help an agent investigate a specific question about an accessible application, package or executable. RefTrail can use an authorized, reproducible REA finding as an observation: record the target version, procedure, evidence and limitations. Then test the independent implementation against that observation. Pass on the observed behavior in your own words and keep decompiled code away from the agent that writes the implementation. The included example uses the GNU CLI directly. REA project
The core method consists of the brief, the coverage catalog and the criteria baseline. The technical examples apply when you want to bind tested sources to a deliverable. The integration runner freezes the checked tree, reruns the tests and verifies the showcase from that copy and packages the candidate and its license. For an installable Windows product, use the target-system acceptance guide .
node --test " tests/*.test.mjs "
node integration/verify.mjs
The direct suite uses the ambient temporary directory. The packaging runner runs its source tests with a separate temporary directory. The GitHub workflow exercises both and includes a Windows path-alias check. npm test runs the direct suite as well.
Reference material and license
Use sources and accounts you are authorized to access. Check the terms for each source and file you plan to reuse. A license or trial agreement can restrict uses that copyright law alone would allow. Describe the required behavior in your own words and create your own screen layouts, text, icons and artwork. Give the result its own name.
In the EU, the functionality of a program, its programming language and its data file formats are not protected by the copyright in computer programs. A licensed user may observe, study and test how a program works while using it as the license allows. Copying text from a manual can still infringe. CJEU, C-406/10 An original GUI can qualify for copyright protection. CJEU, C-393/09, paragraphs 44–51 Decompiling a program is permitted only in narrow cases such as interoperability, and its results must not serve to build a program that is substantially similar in its expression. Directive 2009/24/EC, Article 6 Record the basis for any necessary interface identifiers and separately licensed material. Trademark, unfair-competition, trade-secret and contract law apply separately. These notes are general information and no legal advice. Have a concrete project reviewed before you publish or sell it.
RefTrail's own briefs, templates, example code and documentation are available under the MIT License . The reference products have their own terms. Changes · Contributing
Zero-trust verification & clean-room spec protocol for coding agents.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Zero-trust verification & clean-room spec protocol for coding agents. - kidimusicofficial-prog/reftrail

GitHub - kidimusicofficial-prog/reftrail: Zero-trust verification & clean-room spec protocol for coding agents. · GitHub
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
kidimusicofficial-prog
reftrail Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Zero-trust verification & clean-room spec protocol for coding agents.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples integration integration prompts prompts showcase/ uniq showcase/ uniq templates templates tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.de.md README.de.md README.en.md README.en.md README.md README.md package.json package.json View all files Repository files navigation
Build an independent implementation from documented behavior. Set the acceptance criteria before coding, then show which workflows pass and where the evidence ends.
A coding agent reports a finished feature within minutes. RefTrail makes that report checkable. Each requirement gets a source and a measurable outcome, the criteria are frozen before the first line of product code, and the final report states what was proven, by which check, and what is still open.
The tools need Node.js 22.16 or newer and no further packages.
Line Runs reproduces four stdin/stdout modes of GNU uniq . Given bolt, bolt, nut, bolt as four consecutive lines, its count mode returns three groups of 2, 1 and 1 . Run the acceptance cases from the repository root:
node showcase/uniq/verify.mjs
node tools/catalog.mjs report showcase/uniq/COVERAGE.csv --gates G01
The first command replays the acceptance cases against the frozen expectations. The second calculates the result from the catalog: 100% functional coverage, 1/1 mandatory checks passed . To compare both programs on your machine, install GNU Coreutils separately and append --reference to the first command.
Every recorded output comparison matched GNU uniq byte for byte: 16 of 16 under Linux with version 9.4 for the first criteria revision, and for revision 2 16 of 16 under Windows with version 8.32 and 16 of 16 under Linux with version 9.4. Revision 2 adds a quiet stop when a reader closes the output early. It was frozen before the program changed, and 2 of 2 early-exit checks passed for both programs on both systems. These checks cover termination and stderr, not output bytes. All four local invalid-input checks passed. An additional reviewer tested 40 more inputs across all four modes. Inspect the criteria fixed before coding , working program , filled catalog , report with timings and cost status and delivery decision .
Open the English agent brief . Supply a product and its official website. Add the desired edition, workflows, limits and budget if you know them.
Give the completed brief to your coding agent in the target directory. With only the two inputs, the first run covers one bounded end-to-end workflow on one platform, with a 60-minute limit. A larger product goal stays visible in the backlog.
Record each requirement, source and measurable outcome. Freeze the criteria before implementation and retain the digest independently. Keep progress and evidence in the coverage catalog . The brief runs the tools for these steps, so keep a downloaded copy of this repository where your agent can read it, or copy the tools folder into your project.
Run the finished workflow through its real interface. Report weighted feature coverage alongside passed mandatory checks, unresolved critical defects and the delivery decision. The catalog tool prints the first two together. A missing mandatory backup can block acceptance even at 100% weighted coverage.
Access to the reference. RefTrail organizes what you are allowed to observe. It does not analyze or open a product.
Proof of the sequence. The digest shows that the criteria did not change. That they existed before the code follows from where you kept the digest. Use a place the implementer cannot rewrite, such as a pushed commit or a message to the requester.
The reach of a result. A passed catalog covers the listed cases on the tested systems. Equivalence with the whole reference product needs more cases.
The legal review. The notes below are general information.
REA can help an agent investigate a specific question about an accessible application, package or executable. RefTrail can use an authorized, reproducible REA finding as an observation: record the target version, procedure, evidence and limitations. Then test the independent implementation against that observation. Pass on the observed behavior in your own words and keep decompiled code away from the agent that writes the implementation. The included example uses the GNU CLI directly. REA project
The core method consists of the brief, the coverage catalog and the criteria baseline. The technical examples apply when you want to bind tested sources to a deliverable. The integration runner freezes the checked tree, reruns the tests and verifies the showcase from that copy and packages the candidate and its license. For an installable Windows product, use the target-system acceptance guide .
node --test " tests/*.test.mjs "
node integration/verify.mjs
The direct suite uses the ambient temporary directory. The packaging runner runs its source tests with a separate temporary directory. The GitHub workflow exercises both and includes a Windows path-alias check. npm test runs the direct suite as well.
Reference material and license
Use sources and accounts you are authorized to access. Check the terms for each source and file you plan to reuse. A license or trial agreement can restrict uses that copyright law alone would allow. Describe the required behavior in your own words and create your own screen layouts, text, icons and artwork. Give the result its own name.
In the EU, the functionality of a program, its programming language and its data file formats are not protected by the copyright in computer programs. A licensed user may observe, study and test how a program works while using it as the license allows. Copying text from a manual can still infringe. CJEU, C-406/10 An original GUI can qualify for copyright protection. CJEU, C-393/09, paragraphs 44–51 Decompiling a program is permitted only in narrow cases such as interoperability, and its results must not serve to build a program that is substantially similar in its expression. Directive 2009/24/EC, Article 6 Record the basis for any necessary interface identifiers and separately licensed material. Trademark, unfair-competition, trade-secret and contract law apply separately. These notes are general information and no legal advice. Have a concrete project reviewed before you publish or sell it.
RefTrail's own briefs, templates, example code and documentation are available under the MIT License . The reference products have their own terms. Changes · Contributing
Zero-trust verification & clean-room spec protocol for coding agents.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
