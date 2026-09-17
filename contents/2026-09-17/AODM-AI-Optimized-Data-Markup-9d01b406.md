---
source: "https://github.com/FucaSpark/aodm"
hn_url: "https://news.ycombinator.com/item?id=49743019"
title: "AODM: AI Optimized Data Markup"
article_title: "GitHub - FucaSpark/aodm: An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. · GitHub"
image: "https://opengraph.githubassets.com/cd794e6d3fa5f461b3a3a6db10c0030139d14871747e00a2fc46110e10d2f84d/FucaSpark/aodm"
author: "FucaSpark"
captured_at: "2026-09-17T16:22:47Z"
capture_tool: "hn-digest"
hn_id: 49743019
score: 1
comments: 0
posted_at: "2026-09-17T16:19:55Z"
tags:
  - hacker-news
---

# AODM: AI Optimized Data Markup

- HN: [49743019](https://news.ycombinator.com/item?id=49743019)
- Source: [github.com](https://github.com/FucaSpark/aodm)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T16:19:55Z

## Translation

Title: AODM: AI Optimized Data Markup
Article title: GitHub - FucaSpark/aodm: An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. · GitHub
Description: An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. - FucaSpark/aodm

Article text:
GitHub - FucaSpark/aodm: An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. · GitHub
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
FucaSpark
/
aodm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
AODM AODM converters converters graph graph llm llm parsers parsers reference reference .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md conformance-check.py conformance-check.py View all files Repository files navigation
AODM (AI Optimized Data Markup) is a small, open vocabulary for
representing knowledge — entities, relationships, facts, rules — in a form
both humans and AI systems can read, validate, and exchange.
This is the AODM 1.2 specification and its reference implementations.
AODM 1.2 Specification (start here)
The current core: 4 knowledge primitives ( entity , relationship ,
fact , rule ) and 3 annotations ( source , confidence , value ),
plus a data-quality layer — content hashing, polarity, derivation
tracking, temporal validity, and structured measurement.
Namespace: http://fucaspark.com/aodm/1.2
python3 -m venv .venv && source .venv/bin/activate && pip install lxml jsonschema && python3 conformance-check.py
It also documents which constraints the schemas enforce and which require
processor logic — notably referential integrity, since libxml2 does not
enforce xs:IDREF for XML Schema.
Content digests are processor logic of the same kind: every reference
implementation recomputes an item's sha256 hash and rejects a document
whose digest disagrees with its text, in both the XML and JSON
serialisations.
Directory
Contents
parsers/
Parsers for Python, JavaScript, Java and C#, each parsing AODM 1.2 XML into the published JSON model and applying the validation rules no schema language can express. parsers/conformance.py runs every parser against the same suite.
graph/
Graph compiler and forward-chaining inference engine
converters/
XBRL to AODM conversion
llm/context.py
Builds LLM context windows from an AODM document
reference/
Browser scripts — a validator and the HTML profile extractor
Each parser is standard library only, with no dependency to add.
Use the HTML Embedding Profile:
AODM/AODM-1.2-HTML-PROFILE.md . It marks
up content with data-aodm attributes on the elements that display it, so
the human-readable and machine-readable versions cannot drift apart.
< p data-aodm =" fact " data-aodm-id =" license " data-aodm-about =" aodm "
data-aodm-confidence =" 1.0 " >
AODM is released under the Apache License 2.0.
</ p >
Extract and validate embedded markup with
reference/aodm-html-extract.js , or
online at https://fucaspark.com/validator.php .
The 1.2 core ontology and its formal definition are complete and tested,
with reference parsers in 4 languages, a graph compiler, an inference
engine, an LLM context builder and an XBRL converter, all covered by the
conformance suite.
Apache License 2.0 — see LICENSE and NOTICE .
An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. - FucaSpark/aodm

GitHub - FucaSpark/aodm: An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange. · GitHub
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
FucaSpark
/
aodm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
AODM AODM converters converters graph graph llm llm parsers parsers reference reference .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md conformance-check.py conformance-check.py View all files Repository files navigation
AODM (AI Optimized Data Markup) is a small, open vocabulary for
representing knowledge — entities, relationships, facts, rules — in a form
both humans and AI systems can read, validate, and exchange.
This is the AODM 1.2 specification and its reference implementations.
AODM 1.2 Specification (start here)
The current core: 4 knowledge primitives ( entity , relationship ,
fact , rule ) and 3 annotations ( source , confidence , value ),
plus a data-quality layer — content hashing, polarity, derivation
tracking, temporal validity, and structured measurement.
Namespace: http://fucaspark.com/aodm/1.2
python3 -m venv .venv && source .venv/bin/activate && pip install lxml jsonschema && python3 conformance-check.py
It also documents which constraints the schemas enforce and which require
processor logic — notably referential integrity, since libxml2 does not
enforce xs:IDREF for XML Schema.
Content digests are processor logic of the same kind: every reference
implementation recomputes an item's sha256 hash and rejects a document
whose digest disagrees with its text, in both the XML and JSON
serialisations.
Directory
Contents
parsers/
Parsers for Python, JavaScript, Java and C#, each parsing AODM 1.2 XML into the published JSON model and applying the validation rules no schema language can express. parsers/conformance.py runs every parser against the same suite.
graph/
Graph compiler and forward-chaining inference engine
converters/
XBRL to AODM conversion
llm/context.py
Builds LLM context windows from an AODM document
reference/
Browser scripts — a validator and the HTML profile extractor
Each parser is standard library only, with no dependency to add.
Use the HTML Embedding Profile:
AODM/AODM-1.2-HTML-PROFILE.md . It marks
up content with data-aodm attributes on the elements that display it, so
the human-readable and machine-readable versions cannot drift apart.
< p data-aodm =" fact " data-aodm-id =" license " data-aodm-about =" aodm "
data-aodm-confidence =" 1.0 " >
AODM is released under the Apache License 2.0.
</ p >
Extract and validate embedded markup with
reference/aodm-html-extract.js , or
online at https://fucaspark.com/validator.php .
The 1.2 core ontology and its formal definition are complete and tested,
with reference parsers in 4 languages, a graph compiler, an inference
engine, an LLM context builder and an XBRL converter, all covered by the
conformance suite.
Apache License 2.0 — see LICENSE and NOTICE .
An open vocabulary for representing knowledge — entities, relationships, facts and rules — that both people and AI systems can read, validate and exchange.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
