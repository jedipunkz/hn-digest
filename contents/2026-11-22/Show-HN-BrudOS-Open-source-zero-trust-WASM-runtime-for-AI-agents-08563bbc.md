---
source: "https://github.com/rahatarch/brud"
hn_url: "https://news.ycombinator.com/item?id=50597210"
title: "Show HN: BrudOS – Open-source zero-trust WASM runtime for AI agents"
article_title: "GitHub - rahatarch/brud: An operating system for AI agents. · GitHub"
image: "https://opengraph.githubassets.com/cc0135d0bd2c9f223dc53d19311e42f1ae7f3f7a0d0fc5e19d43e185fe2e0c3b/rahatarch/brud"
author: "rahatarch"
captured_at: "2026-11-22T07:48:06Z"
capture_tool: "hn-digest"
hn_id: 50597210
score: 1
comments: 0
posted_at: "2026-11-22T07:44:37Z"
tags:
  - hacker-news
---

# Show HN: BrudOS – Open-source zero-trust WASM runtime for AI agents

- HN: [50597210](https://news.ycombinator.com/item?id=50597210)
- Source: [github.com](https://github.com/rahatarch/brud)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T07:44:37Z

## Translation

Title: Show HN: BrudOS – Open-source zero-trust WASM runtime for AI agents
Article title: GitHub - rahatarch/brud: An operating system for AI agents. · GitHub
Description: An operating system for AI agents. Contribute to rahatarch/brud development by creating an account on GitHub.

Article text:
GitHub - rahatarch/brud: An operating system for AI agents. · GitHub
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
rahatarch
brud Public Notifications You must be signed in to change notification settings
Star 44 ( 44 ) You must be signed in to star a repository
An operating system for AI agents.
Readme Apache-2.0 license Code of conduct
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
440 Commits 440 Commits Folders and files
.github .github .vscode .vscode assets assets benchmarks/ cage-languages benchmarks/ cage-languages docs-next docs-next docs docs examples/ daemon examples/ daemon packages packages scripts scripts .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md ARCHITECTURE_RULES.md ARCHITECTURE_RULES.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.json tsconfig.json View all files Repository files navigation
A small kernel for building software out of parts that do not have to trust each other.
BrudOS is a kernel you embed in a Node.js program. You build your software as independent pieces called participants . The kernel gives every participant an identity, lets each participant decide who may use its APIs, and enforces that on every call.
The kernel has no abilities of its own: it does not read files, use the network or run commands. Everything useful is a participant. That makes it a good base for plugin systems, tool hosts and AI-agent runtimes, where you need to know exactly who was allowed to do what.
Identity you cannot forge. Every participant gets its identity from the kernel. Nobody can claim to be someone else.
Every API has an owner. The participant that offers an API decides who may call it and what each caller receives. The kernel enforces that answer.
No shared internals. Participants reach each other only through declared APIs, never by importing each other's code.
Untrusted code stays caged. A participant that is not declared official runs inside a WebAssembly sandbox with no access to the operating system.
Small enough to understand. The kernel is a handful of core services, with no runtime dependencies.
Package
Use it to
@brud/kernel
Embed the kernel and run participants
@brud/sdk
Write and test participants
@brud/daemon
Host kernels in a local background process that editors, agents and scripts connect to
@brud/protocol
The message types shared by the daemon and its clients
Requires Node.js 22 or newer. Prebuilt sandbox binaries are included for Linux (x64, arm64), macOS (x64, arm64) and Windows (x64).
npm install @brud/kernel @brud/sdk
Write a participant that offers one API, acme.greet :
// greeter.mjs
import { defineParticipant } from '@brud/sdk/participant' ;
class Greeter {
greet ( name ) {
return `hello, ${ name } ` ;
}
}
export const createService = defineParticipant ( {
id : 'acme.greeter' ,
description : 'Says hello' ,
create : ( ) => new Greeter ( ) ,
exposes : {
'acme.greet' : {
version : 1 ,
scope : 'any' ,
handle : ( caller , ctx ) => ( {
greet : ( name ) => ` ${ ctx . instance . greet ( name ) } (asked by ${ caller . callerClass } )` ,
} ) ,
} ,
} ,
} ) ;
Boot a kernel that runs it, and call the API as a host:
// boot.mjs
import { mkdtemp , writeFile } from 'node:fs/promises' ;
import { tmpdir } from 'node:os' ;
import { join } from 'node:path' ;
import {
bootKernel ,
provisionIdentityRegistryStorage ,
ExposedApiRegistryToken ,
} from '@brud/kernel' ;
const state = await mkdtemp ( join ( tmpdir ( ) , 'brud-hello-' ) ) ;
const hardcores = [ 'identityRegistry' , 'exposedApiRegistry' , 'parser' , 'router' , 'wasmCage' , 'events' , 'interception' ]
. map ( ( name ) => ( { path : `kernel:serviceModules/ ${ name } .js` , official : true } ) ) ;
const configPath = join ( state , 'brud.config.json' ) ;
await writeFile ( configPath , JSON . stringify ( {
services : [ ... hardcores , { path : join ( import . meta . dirname , 'greeter.mjs' ) , official : true } ] ,
} ) ) ;
const registryFile = join ( state , 'identity.json' ) ;
provisionIdentityRegistryStorage ( registryFile ) ;
const kernel = await bootKernel ( configPath , {
options : { 'identity.registry.path' : registryFile } ,
} ) ;
const answer = kernel . runHostCall ( { capabilities : [ ] } , ( ) => {
const registry = kernel . requestAccess (
{ callerId : 'my.host' , tier : 'core' , options : { } } ,
ExposedApiRegistryToken ,
) ;
const greeter = registry . resolve ( 'acme.greet' ) ;
return greeter === null ? 'refused' : greeter . greet ( 'world' ) ;
} ) ;
console . log ( answer ) ;
await kernel . shutdown ( ) ;
node boot.mjs
# hello, world (asked by host)
The getting-started tutorials explain every line, then show how to accept text commands, test a participant against a real kernel, and run it all behind the daemon.
The daemon boots kernels from an operator-owned composition and serves them to other programs over an authenticated, encrypted local connection. UIs and AI agents get separate caller tiers, so an agent can be given less authority than a human-facing UI.
npm install @brud/daemon
npx brud-daemon --home ~ /.brud --foreground
See Run the daemon and Connect a client .
BrudOS documentation : introduction, tutorials, concepts, guides, daemon, security and reference.
What is BrudOS? and Core ideas
ARCHITECTURE.md : layers, invariants and call routing.
PRIVACY.md : local-first guarantees and data handling.
CONTRIBUTING.md : repository layout, tests and pull requests.
An operating system for AI agents.
Readme Apache-2.0 license Code of conduct
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An operating system for AI agents. Contribute to rahatarch/brud development by creating an account on GitHub.

GitHub - rahatarch/brud: An operating system for AI agents. · GitHub
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
rahatarch
brud Public Notifications You must be signed in to change notification settings
Star 44 ( 44 ) You must be signed in to star a repository
An operating system for AI agents.
Readme Apache-2.0 license Code of conduct
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
440 Commits 440 Commits Folders and files
.github .github .vscode .vscode assets assets benchmarks/ cage-languages benchmarks/ cage-languages docs-next docs-next docs docs examples/ daemon examples/ daemon packages packages scripts scripts .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md ARCHITECTURE_RULES.md ARCHITECTURE_RULES.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.eslint.json tsconfig.eslint.json tsconfig.json tsconfig.json View all files Repository files navigation
A small kernel for building software out of parts that do not have to trust each other.
BrudOS is a kernel you embed in a Node.js program. You build your software as independent pieces called participants . The kernel gives every participant an identity, lets each participant decide who may use its APIs, and enforces that on every call.
The kernel has no abilities of its own: it does not read files, use the network or run commands. Everything useful is a participant. That makes it a good base for plugin systems, tool hosts and AI-agent runtimes, where you need to know exactly who was allowed to do what.
Identity you cannot forge. Every participant gets its identity from the kernel. Nobody can claim to be someone else.
Every API has an owner. The participant that offers an API decides who may call it and what each caller receives. The kernel enforces that answer.
No shared internals. Participants reach each other only through declared APIs, never by importing each other's code.
Untrusted code stays caged. A participant that is not declared official runs inside a WebAssembly sandbox with no access to the operating system.
Small enough to understand. The kernel is a handful of core services, with no runtime dependencies.
Package
Use it to
@brud/kernel
Embed the kernel and run participants
@brud/sdk
Write and test participants
@brud/daemon
Host kernels in a local background process that editors, agents and scripts connect to
@brud/protocol
The message types shared by the daemon and its clients
Requires Node.js 22 or newer. Prebuilt sandbox binaries are included for Linux (x64, arm64), macOS (x64, arm64) and Windows (x64).
npm install @brud/kernel @brud/sdk
Write a participant that offers one API, acme.greet :
// greeter.mjs
import { defineParticipant } from '@brud/sdk/participant' ;
class Greeter {
greet ( name ) {
return `hello, ${ name } ` ;
}
}
export const createService = defineParticipant ( {
id : 'acme.greeter' ,
description : 'Says hello' ,
create : ( ) => new Greeter ( ) ,
exposes : {
'acme.greet' : {
version : 1 ,
scope : 'any' ,
handle : ( caller , ctx ) => ( {
greet : ( name ) => ` ${ ctx . instance . greet ( name ) } (asked by ${ caller . callerClass } )` ,
} ) ,
} ,
} ,
} ) ;
Boot a kernel that runs it, and call the API as a host:
// boot.mjs
import { mkdtemp , writeFile } from 'node:fs/promises' ;
import { tmpdir } from 'node:os' ;
import { join } from 'node:path' ;
import {
bootKernel ,
provisionIdentityRegistryStorage ,
ExposedApiRegistryToken ,
} from '@brud/kernel' ;
const state = await mkdtemp ( join ( tmpdir ( ) , 'brud-hello-' ) ) ;
const hardcores = [ 'identityRegistry' , 'exposedApiRegistry' , 'parser' , 'router' , 'wasmCage' , 'events' , 'interception' ]
. map ( ( name ) => ( { path : `kernel:serviceModules/ ${ name } .js` , official : true } ) ) ;
const configPath = join ( state , 'brud.config.json' ) ;
await writeFile ( configPath , JSON . stringify ( {
services : [ ... hardcores , { path : join ( import . meta . dirname , 'greeter.mjs' ) , official : true } ] ,
} ) ) ;
const registryFile = join ( state , 'identity.json' ) ;
provisionIdentityRegistryStorage ( registryFile ) ;
const kernel = await bootKernel ( configPath , {
options : { 'identity.registry.path' : registryFile } ,
} ) ;
const answer = kernel . runHostCall ( { capabilities : [ ] } , ( ) => {
const registry = kernel . requestAccess (
{ callerId : 'my.host' , tier : 'core' , options : { } } ,
ExposedApiRegistryToken ,
) ;
const greeter = registry . resolve ( 'acme.greet' ) ;
return greeter === null ? 'refused' : greeter . greet ( 'world' ) ;
} ) ;
console . log ( answer ) ;
await kernel . shutdown ( ) ;
node boot.mjs
# hello, world (asked by host)
The getting-started tutorials explain every line, then show how to accept text commands, test a participant against a real kernel, and run it all behind the daemon.
The daemon boots kernels from an operator-owned composition and serves them to other programs over an authenticated, encrypted local connection. UIs and AI agents get separate caller tiers, so an agent can be given less authority than a human-facing UI.
npm install @brud/daemon
npx brud-daemon --home ~ /.brud --foreground
See Run the daemon and Connect a client .
BrudOS documentation : introduction, tutorials, concepts, guides, daemon, security and reference.
What is BrudOS? and Core ideas
ARCHITECTURE.md : layers, invariants and call routing.
PRIVACY.md : local-first guarantees and data handling.
CONTRIBUTING.md : repository layout, tests and pull requests.
An operating system for AI agents.
Readme Apache-2.0 license Code of conduct
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
