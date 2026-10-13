---
source: "https://github.com/maharshi365/deslop"
hn_url: "https://news.ycombinator.com/item?id=50075809"
title: "Show HN: Deslop – Oxlint rules for AI-generated TypeScript"
article_title: "GitHub - maharshi365/deslop · GitHub"
image: "https://opengraph.githubassets.com/04ca49e62c6a03158078754fe7b3946966d2d22326bbf67132d3362fbaf3f511/maharshi365/deslop"
author: "maharshi365"
captured_at: "2026-10-13T17:58:01Z"
capture_tool: "hn-digest"
hn_id: 50075809
score: 2
comments: 0
posted_at: "2026-10-13T17:52:57Z"
tags:
  - hacker-news
---

# Show HN: Deslop – Oxlint rules for AI-generated TypeScript

- HN: [50075809](https://news.ycombinator.com/item?id=50075809)
- Source: [github.com](https://github.com/maharshi365/deslop)
- Score: 2
- Comments: 0
- Posted: 2026-10-13T17:52:57Z

## Translation

Title: Show HN: Deslop – Oxlint rules for AI-generated TypeScript
Article title: GitHub - maharshi365/deslop · GitHub
Description: Contribute to maharshi365/deslop development by creating an account on GitHub.
HN text: Catches unsafe types, duplicate code, and low-value tests.

Article text:
GitHub - maharshi365/deslop · GitHub
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
maharshi365
deslop Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
31 Commits 31 Commits Folders and files
docs docs src src tests tests .gitignore .gitignore .oxlintrc.json .oxlintrc.json README.md README.md bench-project.mjs bench-project.mjs bench.mjs bench.mjs package-lock.json package-lock.json package.json package.json stitch-factory.yaml stitch-factory.yaml tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
Opinionated Oxlint rules that reject low-evidence, low-signal, AI-slop code patterns.
Rules are vendored from dmmulroy/anti-slop (generic rules) alongside additional deslop rules, all renamed under the deslop/* namespace. The vendored source is MIT licensed; see Credits below.
npm install --save-dev ai-deslop
# @oxlint/plugins comes along as a runtime dependency — keep it in sync with your oxlint version
Ships compiled JavaScript — works with any oxlint JS-plugin runtime (Bun, Deno, or Node.js).
import { defineConfig } from "oxlint" ;
import { configs as deslopConfigs } from "ai-deslop" ;
export default defineConfig ( {
jsPlugins : [ "ai-deslop" ] ,
options : deslopConfigs . recommended . options ,
rules : deslopConfigs . recommended . rules
} ) ;
Rules
Rule
Status
Description
no-chained-type-assertions
Recommended
Rejects nested as / angle-bracket assertions that fabricate evidence.
no-conditional-empty-object-spread
Recommended
Reports object spreads using a conditional {} branch to omit fields.
no-known-value-widening
Recommended
Rejects known expressions flowing into explicit unknown , object , or open dictionaries.
no-object-parameters
Recommended
Rejects object on function inputs.
no-runtime-typeof
Recommended
Requires boundary parsing instead of ad hoc typeof narrowing.
no-unknown-parameters
Recommended
Rejects unknown on function inputs.
no-unknown-returns
Recommended
Rejects return contracts resolving to unknown .
no-unknown-type-aliases
Recommended
Rejects aliases whose resolved type is unknown .
no-unsafe-dictionary-type
Recommended
Rejects dictionary values based on unsafe escape hatches.
no-widen-then-assert
Recommended
Rejects widening known evidence before narrowing assertions.
require-safety-comment-for-type-assertion
Recommended
Requires an invariant justification for non-const assertions.
no-pass-through-type-alias
Recommended
Disallows aliases that only rename another type.
Code quality and APIs
Rule
Status
Description
no-array-filter-map
Recommended
Rejects adjacent eager array filter/map passes; allows lazy iterator pipelines.
no-reduce-accumulator-copy
Recommended
Rejects non-spread accumulator copies inside reducers. Enable native oxc/no-accumulating-spread alongside it.
no-reflect-apply
Recommended
Rejects global Reflect.apply .
no-reflect-get
Recommended
Rejects global Reflect.get .
no-shape-in-symbol-names
Recommended
Rejects shape in locally owned symbol names.
Tests
Rule
Status
Description
no-call-only-assertions
Recommended
Flags tests whose assertions only check mock calls.
no-module-mocking
Recommended
Rejects Vitest and Jest module mocks.
no-poor-substring-tests
Recommended
Rejects testing that a constant string contains substrings.
no-repeated-expect-call
Recommended
Flags adjacent assertions that repeat the same call in test files. Reuse the result instead. No fixer.
Opt-in rules
Rule
Status
Description
no-duplicate-consts
Opt-in
Reports all-caps constants with the same static value in different files.
no-duplicate-types
Opt-in
Reports exact duplicate object contracts. No fixer.
no-renamed-imports
Opt-in
Fixes unnecessary named-import aliases when the original name has no module-level conflict.
canonical-class-names
Opt-in
Fixes non-canonical Tailwind CSS v4 class spellings.
Standard rules enabled by the recommended config
Standard rule
Status
Description
typescript/no-deprecated
Recommended
Disallows using code marked @deprecated . Requires type-aware linting.
typescript/no-require-imports
Recommended
Disallows CommonJS require() imports.
typescript/no-floating-promises
Recommended
Disallows floating promises. Requires type-aware linting.
typescript/no-unnecessary-condition
Recommended
Disallows always truthy/falsy/nullish conditions. Requires type-aware linting.
no-unreachable
Recommended
Disallows unreachable code.
Guides
Guide
Contents
No duplicate consts
Configuration and exact static-value comparison behavior.
No duplicate types
Exact comparison behavior, worker scope, complete TypeScript/JSON examples, and every supported option.
No renamed imports
When named-import aliases are fixed and how conflicts are handled.
Tailwind canonical class names
Tailwind v4 setup, complete TypeScript/JSON examples, and every supported option.
Credits
Generic rules vendored verbatim from dmmulroy/anti-slop (MIT, © Dillon Mulroy), including src/shared/* .
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to maharshi365/deslop development by creating an account on GitHub.

Catches unsafe types, duplicate code, and low-value tests.

GitHub - maharshi365/deslop · GitHub
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
maharshi365
deslop Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
31 Commits 31 Commits Folders and files
docs docs src src tests tests .gitignore .gitignore .oxlintrc.json .oxlintrc.json README.md README.md bench-project.mjs bench-project.mjs bench.mjs bench.mjs package-lock.json package-lock.json package.json package.json stitch-factory.yaml stitch-factory.yaml tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
Opinionated Oxlint rules that reject low-evidence, low-signal, AI-slop code patterns.
Rules are vendored from dmmulroy/anti-slop (generic rules) alongside additional deslop rules, all renamed under the deslop/* namespace. The vendored source is MIT licensed; see Credits below.
npm install --save-dev ai-deslop
# @oxlint/plugins comes along as a runtime dependency — keep it in sync with your oxlint version
Ships compiled JavaScript — works with any oxlint JS-plugin runtime (Bun, Deno, or Node.js).
import { defineConfig } from "oxlint" ;
import { configs as deslopConfigs } from "ai-deslop" ;
export default defineConfig ( {
jsPlugins : [ "ai-deslop" ] ,
options : deslopConfigs . recommended . options ,
rules : deslopConfigs . recommended . rules
} ) ;
Rules
Rule
Status
Description
no-chained-type-assertions
Recommended
Rejects nested as / angle-bracket assertions that fabricate evidence.
no-conditional-empty-object-spread
Recommended
Reports object spreads using a conditional {} branch to omit fields.
no-known-value-widening
Recommended
Rejects known expressions flowing into explicit unknown , object , or open dictionaries.
no-object-parameters
Recommended
Rejects object on function inputs.
no-runtime-typeof
Recommended
Requires boundary parsing instead of ad hoc typeof narrowing.
no-unknown-parameters
Recommended
Rejects unknown on function inputs.
no-unknown-returns
Recommended
Rejects return contracts resolving to unknown .
no-unknown-type-aliases
Recommended
Rejects aliases whose resolved type is unknown .
no-unsafe-dictionary-type
Recommended
Rejects dictionary values based on unsafe escape hatches.
no-widen-then-assert
Recommended
Rejects widening known evidence before narrowing assertions.
require-safety-comment-for-type-assertion
Recommended
Requires an invariant justification for non-const assertions.
no-pass-through-type-alias
Recommended
Disallows aliases that only rename another type.
Code quality and APIs
Rule
Status
Description
no-array-filter-map
Recommended
Rejects adjacent eager array filter/map passes; allows lazy iterator pipelines.
no-reduce-accumulator-copy
Recommended
Rejects non-spread accumulator copies inside reducers. Enable native oxc/no-accumulating-spread alongside it.
no-reflect-apply
Recommended
Rejects global Reflect.apply .
no-reflect-get
Recommended
Rejects global Reflect.get .
no-shape-in-symbol-names
Recommended
Rejects shape in locally owned symbol names.
Tests
Rule
Status
Description
no-call-only-assertions
Recommended
Flags tests whose assertions only check mock calls.
no-module-mocking
Recommended
Rejects Vitest and Jest module mocks.
no-poor-substring-tests
Recommended
Rejects testing that a constant string contains substrings.
no-repeated-expect-call
Recommended
Flags adjacent assertions that repeat the same call in test files. Reuse the result instead. No fixer.
Opt-in rules
Rule
Status
Description
no-duplicate-consts
Opt-in
Reports all-caps constants with the same static value in different files.
no-duplicate-types
Opt-in
Reports exact duplicate object contracts. No fixer.
no-renamed-imports
Opt-in
Fixes unnecessary named-import aliases when the original name has no module-level conflict.
canonical-class-names
Opt-in
Fixes non-canonical Tailwind CSS v4 class spellings.
Standard rules enabled by the recommended config
Standard rule
Status
Description
typescript/no-deprecated
Recommended
Disallows using code marked @deprecated . Requires type-aware linting.
typescript/no-require-imports
Recommended
Disallows CommonJS require() imports.
typescript/no-floating-promises
Recommended
Disallows floating promises. Requires type-aware linting.
typescript/no-unnecessary-condition
Recommended
Disallows always truthy/falsy/nullish conditions. Requires type-aware linting.
no-unreachable
Recommended
Disallows unreachable code.
Guides
Guide
Contents
No duplicate consts
Configuration and exact static-value comparison behavior.
No duplicate types
Exact comparison behavior, worker scope, complete TypeScript/JSON examples, and every supported option.
No renamed imports
When named-import aliases are fixed and how conflicts are handled.
Tailwind canonical class names
Tailwind v4 setup, complete TypeScript/JSON examples, and every supported option.
Credits
Generic rules vendored verbatim from dmmulroy/anti-slop (MIT, © Dillon Mulroy), including src/shared/* .
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
