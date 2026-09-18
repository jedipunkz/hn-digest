---
source: "https://github.com/anthropics/claude-code/commit/a92ea1cdb11ad21f9d583fad2db181dfdac918a6"
hn_url: "https://news.ycombinator.com/item?id=49759414"
title: "Claude Code 2.1.277 adds support for AGENTS.md"
article_title: "mods/agents-md: the AGENTS.md project-instructions mod (#95409) · anthropics/claude-code@a92ea1c · GitHub"
image: "https://opengraph.githubassets.com/a8bb17c111e994c717d7d9813a2aa578e09aafead23301d02a9690f29abdd672/anthropics/claude-code/commit/a92ea1cdb11ad21f9d583fad2db181dfdac918a6"
author: "skwashd"
captured_at: "2026-09-18T20:30:32Z"
capture_tool: "hn-digest"
hn_id: 49759414
score: 1
comments: 0
posted_at: "2026-09-18T19:58:58Z"
tags:
  - hacker-news
---

# Claude Code 2.1.277 adds support for AGENTS.md

- HN: [49759414](https://news.ycombinator.com/item?id=49759414)
- Source: [github.com](https://github.com/anthropics/claude-code/commit/a92ea1cdb11ad21f9d583fad2db181dfdac918a6)
- Score: 1
- Comments: 0
- Posted: 2026-09-18T19:58:58Z

## Translation

Title: Claude Code 2.1.277 adds support for AGENTS.md
Article title: mods/agents-md: the AGENTS.md project-instructions mod (#95409) · anthropics/claude-code@a92ea1c · GitHub
Description: Claude Code is an agentic coding tool that lives in your terminal, understands your codebase, and helps you code faster by executing routine tasks, explaining complex code, and handling git workflows - all through natural language commands. - mods/agents-md: the AGENTS.md project-instructions mod (#
[truncated]

Article text:
mods/agents-md: the AGENTS.md project-instructions mod (#95409) · anthropics/claude-code@a92ea1c · GitHub
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
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Browse files Browse the repository at this point in the history Browse files poteat authored mods/agents-md: the AGENTS.md project-instructions mod ( #95409 ) * mods/types: refresh the engine typings; diff's old-files fixture names isLink
* mods/agents-md: the AGENTS.md project-instructions mod
* mods/agents-md: the instructionFiles option with its legacy key, as shipped; typings at 2.1.277 1 parent 31a3b00 commit a92ea1c Copy full SHA for a92ea1c 76 file s changed
Expand file tree Collapse file tree Open diff view settings Filter options mods README.md
agents-md .claude-plugin plugin.json
is-kept-without-instructions.ts
diff/tests/fixtures old-files.ts
Dismiss banner Expand file tree Collapse file tree Open diff view settings Collapse file ‎ mods/README.md ‎
Copy file name to clipboard Expand all lines: mods/README.md + 2 - 1 Lines changed: 2 additions & 1 deletion Display the source diff
Original file line number Diff line number Diff line change @@ -2,14 +2,15 @@ 2 2
3 3 A mod is a Claude Code plugin whose behaviour lives in a hooks module: one 4 4 ` register(on, options) ` entry that hooks the engine's events as functions 5 - ` ($, e, next) ` . These three ship inside Claude Code; this folder is their 5 + ` ($, e, next) ` . These four ship inside Claude Code; this folder is their 6 6 source, published as it is built into the binary. 7 7
8 8 | Mod | What it does | Seated | 9 9 | --- | --- | --- | 10 10 | [ ` sec-default ` ] ( sec-default ) | Keeps an organization's classic hooks, prompt content, managed settings and tool policy out of reach of the plugins a person installs; adds no policy of its own. | Outermost, on a machine with managed settings or for a Team or Enterprise organization, unless managed ` prependPlugins ` says otherwise | 11 11 | [ ` diff ` ] ( diff ) | ` /diff ` : the session's uncommitted changes in a pane beside the transcript, file by file with their hunks, refreshed as Claude edits files and runs commands. | Built in | 12 12 | [ ` telemetry ` ] ( telemetry ) | Adds ` $.telemetry ` ( ` log ` , ` mark ` ) in the ` engine.create ` fold so a plugin can record an event as a first-party analytics row; sends nothing wherever Claude Code's analytics are off. | Built in | 13 + | [ ` agents-md ` ] ( agents-md ) | ` AGENTS.md ` as project instructions, by one option: loaded where the project has no ` CLAUDE.md ` of its own ( ` claude-md-or-agents-md ` , the default) or beside it ( ` claude-md-and-agents-md ` ), placed and framed exactly as the engine places ` CLAUDE.md ` , nested ones on a ` Read ` ; or the project's and the person's instruction files dropped and the organization's kept ( ` managed-only ` ); or ` CLAUDE.md ` alone, as the engine reads it ( ` claude-md ` ). | Built in | 13 14
14 15 Each folder is a complete plugin: ` .claude-plugin/plugin.json ` , a 15 16 ` hooks/hooks.json ` naming the module, and TypeScript under ` hook
[truncated]
Copy file name to clipboard + 23 Lines changed: 23 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,23 @@ 1 + { 2 + "name" : " agents-md " , 3 + "version" : " 0.1.0 " , 4 + "description" : " AGENTS.md as project instructions, by one option. Under claude-md-or-agents-md a project with no instruction files of its own gets its AGENTS.md files instead: a prompt.context hook asks $.fs.ancestors for every AGENTS.md and .claude/AGENTS.md from the filesystem root down to the working directory and hands them to the engine as project instruction files, which it renders, announces and withholds from agents exactly as it does CLAUDE.md, skipping any file the engine already loaded by import or link; a tool.call hook on Read attaches the ones under the project root as the engine attaches a nested CLAUDE.md. Under claude-md-and-agents-md every AGENTS.md is loaded beside CLAUDE.md. Under managed-only the project's and the person's instruction files are dropped and the organization's managed ones kept. Under claude-md the engine's walk stands alone. claude-md-or-agents-md is the default. " , 5 + "author" : { 6 + "name" : " Anthropic " 7 + }, 8 + "userConfig" : { 9 + "instructionFiles" : { 10 + "type" : " string " , 11 + "title" : " Project instructions " , 12 + "description" : " \" claude-md \" : CLAUDE.md only, loaded by the engine as today. \" claude-md-or-agents-md \" (default): a project with no CLAUDE.md of its own gets its AGENTS.md files instead, loaded exactly where and how CLAUDE.md would be. \" claude-md-and-agents-md \" : AGENTS.md files are loaded beside CLAUDE.md (a file CLAUDE.md already imports or links to is not loaded twice). \" managed-only \" : the project's and your own instruction files are dropped; the organization's managed CLAUDE.md and memory stay. " , 13 + "required" : false , 14 + "default" : " claude-md-or-agents-md " , 15 + "options" : [ 16 + " claude-md " , 17 + " claude-md-or-agents-md " , 18 + " claude-md-
[truncated]
Copy file name to clipboard + 163 Lines changed: 163 additions & 0 deletions Display the source diff
Original file line number Diff line number Diff line change @@ -0,0 +1,163 @@ 1 + # agents-md 2 + 3 + ` AGENTS.md ` read the way Claude Code reads ` CLAUDE.md ` , as a plugin, under 4 + one option, ` instructionFiles ` : 5 + 6 + - ` claude-md ` : only ` CLAUDE.md ` is loaded, by the engine, as today. The plugin 7 + adds nothing. 8 + - ` claude-md-or-agents-md ` (the default): a project with no instruction files 9 + of its own gets its ` AGENTS.md ` files instead, loaded exactly where and how 10 + ` CLAUDE.md ` would be. "Of its own" is read off what the engine loaded for 11 + the context: a ` CLAUDE.md ` , ` .claude/CLAUDE.md ` or ` CLAUDE.local.md ` in any 12 + directory from the root down to the working directory leaves the whole 13 + project to the engine, and the plugin stays out (the organization's managed 14 + file, the person's ` ~/.claude/CLAUDE.md ` , a ` .claude/rules ` file and an 15 + added directory's ` CLAUDE.md ` do not count, as the nested walk does not see 16 + them either). With none, every ` AGENTS.md ` and ` .claude/AGENTS.md ` on that 17 + path joins the instruction files the engine renders, and a ` Read ` under a 18 + subdirectory attaches that directory's ` AGENTS.md ` unless a ` CLAUDE.md ` 19 + there claims it. 20 + - ` claude-md-and-agents-md ` : every ` AGENTS.md ` is loaded beside ` CLAUDE.md ` , 21 + up and down the tree; a file ` CLAUDE.md ` already ` @ ` -imports, or is a link 22 + to, is not loaded a second time (compared by path, then by content). 23 + - ` managed-only ` : the project's checked-in and private instruction files and 24 + the person's own are dropped from the context; the organization's managed 25 + ` CLAUDE.md ` and the engine's memory stay. The engine's nested ` CLAUDE.md ` 26 + attachments on ` Read ` are not an event yet and still arrive. (The engine's 27 + ` claudeMdExcludes ` setting also exists, for user, project and local files, 28 + and applies to the ` AGENTS.md ` files this plugin reads too.) 29 + 30 + How th
[truncated]
Copy file name to clipboard + 28 Lines changed: 28 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,28 @@ 1 + import type { InstructionFile } from 'claude-code' 2 + 3 + /** 4 + * The path of the file at the head of an instruction file's `@`-import 5 + * chain; the file's own path when nothing imported it. 6 + * 7 + * Follows `parent` among the given files until a file nothing imported. 8 + * 9 + * @param file the instruction file 10 + * @param byPath the files it may have come through, by path 11 + * @returns the chain head's path 12 + */ 13 + export function chainRootOf ( 14 + file : InstructionFile , 15 + byPath : ReadonlyMap < string , InstructionFile > , 16 + ) : string { 17 + const seen = new Set < string > ( [ file . path ] ) 18 + let path = file . path 19 + let parent = file . parent 20 + 21 + while ( parent !== undefined && ! seen . has ( parent ) ) { 22 + seen . add ( parent ) 23 + path = parent 24 + parent = byPath . get ( parent ) ?. parent 25 + } 26 + 27 + return path 28 + } Collapse file ‎ mods/agents-md/hooks/files/dropped-kinds.ts ‎
Copy file name to clipboard + 12 Lines changed: 12 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,12 @@ 1 + import type { InstructionFileKind } from 'claude-code' 2 + 3 + /** 4 + * The kinds `managed-only` drops: the project's checked-in and private 5 + * instruction files and the person's own; what it keeps is the organization's 6 + * and memory. 7 + */ 8 + export const DROPPED_KINDS : readonly InstructionFileKind [ ] = [ 9 + 'project' , 10 + 'local' , 11 + 'user' , 12 + ] Collapse file ‎ mods/agents-md/hooks/files/files-of.ts ‎
Copy file name to clipboard + 20 Lines changed: 20 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,20 @@ 1 + import type { FsAncestor , InstructionFile } from 'claude-code' 2 + 3 + /** 4 + * The found AGENTS.md files as project instruction files, one per file and 5 + * per file it `@`-imported, in walk and load order. 6 + * 7 + * An import names the file that brought it as its parent. 8 + * 9 + * @param found what `$.fs.ancestors` found, root first 10 + * @returns the instruction files, kind `project` 11 + */ 12 + export const filesOf = ( found : readonly FsAncestor [ ] ) : InstructionFile [ ] => 13 + found . flatMap ( entry => 14 + entry . parts . map ( ( part , index ) => ( { 15 + path : part . path , 16 + kind : 'project' as const , 17 + content : part . content , 18 + ... ( index > 0 && { parent : entry . parts [ 0 ] ?. path ?? part . path } ) , 19 + } ) ) , 20 + ) Collapse file ‎ mods/agents-md/hooks/files/index.ts ‎
Copy file name to clipboard + 12 Lines changed: 12 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,12 @@ 1 + export * from './chain-root-of.js' 2 + export * from './dropped-kinds.js' 3 + export * from './files-of.js' 4 + export * from './insertion-index.js' 5 + export * from './is-claude-file-on-walk.js' 6 + export * from './is-kept-without-instructions.js' 7 + export * from './is-project-own.js' 8 + export * from './project-dir-of.js' 9 + export * from './unseen-files.js' 10 + export * from './with-project-files.js' 11 + 12 + export * as default from '.' Collapse file ‎ mods/agents-md/hooks/files/insertion-index.ts ‎
Copy file name to clipboard + 44 Lines changed: 44 additions & 0 deletions Original file line number

[truncated]

## Original Extract

Claude Code is an agentic coding tool that lives in your terminal, understands your codebase, and helps you code faster by executing routine tasks, explaining complex code, and handling git workflows - all through natural language commands. - mods/agents-md: the AGENTS.md project-instructions mod (#
[truncated]

mods/agents-md: the AGENTS.md project-instructions mod (#95409) · anthropics/claude-code@a92ea1c · GitHub
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
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Browse files Browse the repository at this point in the history Browse files poteat authored mods/agents-md: the AGENTS.md project-instructions mod ( #95409 ) * mods/types: refresh the engine typings; diff's old-files fixture names isLink
* mods/agents-md: the AGENTS.md project-instructions mod
* mods/agents-md: the instructionFiles option with its legacy key, as shipped; typings at 2.1.277 1 parent 31a3b00 commit a92ea1c Copy full SHA for a92ea1c 76 file s changed
Expand file tree Collapse file tree Open diff view settings Filter options mods README.md
agents-md .claude-plugin plugin.json
is-kept-without-instructions.ts
diff/tests/fixtures old-files.ts
Dismiss banner Expand file tree Collapse file tree Open diff view settings Collapse file ‎ mods/README.md ‎
Copy file name to clipboard Expand all lines: mods/README.md + 2 - 1 Lines changed: 2 additions & 1 deletion Display the source diff
Original file line number Diff line number Diff line change @@ -2,14 +2,15 @@ 2 2
3 3 A mod is a Claude Code plugin whose behaviour lives in a hooks module: one 4 4 ` register(on, options) ` entry that hooks the engine's events as functions 5 - ` ($, e, next) ` . These three ship inside Claude Code; this folder is their 5 + ` ($, e, next) ` . These four ship inside Claude Code; this folder is their 6 6 source, published as it is built into the binary. 7 7
8 8 | Mod | What it does | Seated | 9 9 | --- | --- | --- | 10 10 | [ ` sec-default ` ] ( sec-default ) | Keeps an organization's classic hooks, prompt content, managed settings and tool policy out of reach of the plugins a person installs; adds no policy of its own. | Outermost, on a machine with managed settings or for a Team or Enterprise organization, unless managed ` prependPlugins ` says otherwise | 11 11 | [ ` diff ` ] ( diff ) | ` /diff ` : the session's uncommitted changes in a pane beside the transcript, file by file with their hunks, refreshed as Claude edits files and runs commands. | Built in | 12 12 | [ ` telemetry ` ] ( telemetry ) | Adds ` $.telemetry ` ( ` log ` , ` mark ` ) in the ` engine.create ` fold so a plugin can record an event as a first-party analytics row; sends nothing wherever Claude Code's analytics are off. | Built in | 13 + | [ ` agents-md ` ] ( agents-md ) | ` AGENTS.md ` as project instructions, by one option: loaded where the project has no ` CLAUDE.md ` of its own ( ` claude-md-or-agents-md ` , the default) or beside it ( ` claude-md-and-agents-md ` ), placed and framed exactly as the engine places ` CLAUDE.md ` , nested ones on a ` Read ` ; or the project's and the person's instruction files dropped and the organization's kept ( ` managed-only ` ); or ` CLAUDE.md ` alone, as the engine reads it ( ` claude-md ` ). | Built in | 13 14
14 15 Each folder is a complete plugin: ` .claude-plugin/plugin.json ` , a 15 16 ` hooks/hooks.json ` naming the module, and TypeScript under ` hook
[truncated]
Copy file name to clipboard + 23 Lines changed: 23 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,23 @@ 1 + { 2 + "name" : " agents-md " , 3 + "version" : " 0.1.0 " , 4 + "description" : " AGENTS.md as project instructions, by one option. Under claude-md-or-agents-md a project with no instruction files of its own gets its AGENTS.md files instead: a prompt.context hook asks $.fs.ancestors for every AGENTS.md and .claude/AGENTS.md from the filesystem root down to the working directory and hands them to the engine as project instruction files, which it renders, announces and withholds from agents exactly as it does CLAUDE.md, skipping any file the engine already loaded by import or link; a tool.call hook on Read attaches the ones under the project root as the engine attaches a nested CLAUDE.md. Under claude-md-and-agents-md every AGENTS.md is loaded beside CLAUDE.md. Under managed-only the project's and the person's instruction files are dropped and the organization's managed ones kept. Under claude-md the engine's walk stands alone. claude-md-or-agents-md is the default. " , 5 + "author" : { 6 + "name" : " Anthropic " 7 + }, 8 + "userConfig" : { 9 + "instructionFiles" : { 10 + "type" : " string " , 11 + "title" : " Project instructions " , 12 + "description" : " \" claude-md \" : CLAUDE.md only, loaded by the engine as today. \" claude-md-or-agents-md \" (default): a project with no CLAUDE.md of its own gets its AGENTS.md files instead, loaded exactly where and how CLAUDE.md would be. \" claude-md-and-agents-md \" : AGENTS.md files are loaded beside CLAUDE.md (a file CLAUDE.md already imports or links to is not loaded twice). \" managed-only \" : the project's and your own instruction files are dropped; the organization's managed CLAUDE.md and memory stay. " , 13 + "required" : false , 14 + "default" : " claude-md-or-agents-md " , 15 + "options" : [ 16 + " claude-md " , 17 + " claude-md-or-agents-md " , 18 + " claude-md-
[truncated]
Copy file name to clipboard + 163 Lines changed: 163 additions & 0 deletions Display the source diff
Original file line number Diff line number Diff line change @@ -0,0 +1,163 @@ 1 + # agents-md 2 + 3 + ` AGENTS.md ` read the way Claude Code reads ` CLAUDE.md ` , as a plugin, under 4 + one option, ` instructionFiles ` : 5 + 6 + - ` claude-md ` : only ` CLAUDE.md ` is loaded, by the engine, as today. The plugin 7 + adds nothing. 8 + - ` claude-md-or-agents-md ` (the default): a project with no instruction files 9 + of its own gets its ` AGENTS.md ` files instead, loaded exactly where and how 10 + ` CLAUDE.md ` would be. "Of its own" is read off what the engine loaded for 11 + the context: a ` CLAUDE.md ` , ` .claude/CLAUDE.md ` or ` CLAUDE.local.md ` in any 12 + directory from the root down to the working directory leaves the whole 13 + project to the engine, and the plugin stays out (the organization's managed 14 + file, the person's ` ~/.claude/CLAUDE.md ` , a ` .claude/rules ` file and an 15 + added directory's ` CLAUDE.md ` do not count, as the nested walk does not see 16 + them either). With none, every ` AGENTS.md ` and ` .claude/AGENTS.md ` on that 17 + path joins the instruction files the engine renders, and a ` Read ` under a 18 + subdirectory attaches that directory's ` AGENTS.md ` unless a ` CLAUDE.md ` 19 + there claims it. 20 + - ` claude-md-and-agents-md ` : every ` AGENTS.md ` is loaded beside ` CLAUDE.md ` , 21 + up and down the tree; a file ` CLAUDE.md ` already ` @ ` -imports, or is a link 22 + to, is not loaded a second time (compared by path, then by content). 23 + - ` managed-only ` : the project's checked-in and private instruction files and 24 + the person's own are dropped from the context; the organization's managed 25 + ` CLAUDE.md ` and the engine's memory stay. The engine's nested ` CLAUDE.md ` 26 + attachments on ` Read ` are not an event yet and still arrive. (The engine's 27 + ` claudeMdExcludes ` setting also exists, for user, project and local files, 28 + and applies to the ` AGENTS.md ` files this plugin reads too.) 29 + 30 + How th
[truncated]
Copy file name to clipboard + 28 Lines changed: 28 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,28 @@ 1 + import type { InstructionFile } from 'claude-code' 2 + 3 + /** 4 + * The path of the file at the head of an instruction file's `@`-import 5 + * chain; the file's own path when nothing imported it. 6 + * 7 + * Follows `parent` among the given files until a file nothing imported. 8 + * 9 + * @param file the instruction file 10 + * @param byPath the files it may have come through, by path 11 + * @returns the chain head's path 12 + */ 13 + export function chainRootOf ( 14 + file : InstructionFile , 15 + byPath : ReadonlyMap < string , InstructionFile > , 16 + ) : string { 17 + const seen = new Set < string > ( [ file . path ] ) 18 + let path = file . path 19 + let parent = file . parent 20 + 21 + while ( parent !== undefined && ! seen . has ( parent ) ) { 22 + seen . add ( parent ) 23 + path = parent 24 + parent = byPath . get ( parent ) ?. parent 25 + } 26 + 27 + return path 28 + } Collapse file ‎ mods/agents-md/hooks/files/dropped-kinds.ts ‎
Copy file name to clipboard + 12 Lines changed: 12 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,12 @@ 1 + import type { InstructionFileKind } from 'claude-code' 2 + 3 + /** 4 + * The kinds `managed-only` drops: the project's checked-in and private 5 + * instruction files and the person's own; what it keeps is the organization's 6 + * and memory. 7 + */ 8 + export const DROPPED_KINDS : readonly InstructionFileKind [ ] = [ 9 + 'project' , 10 + 'local' , 11 + 'user' , 12 + ] Collapse file ‎ mods/agents-md/hooks/files/files-of.ts ‎
Copy file name to clipboard + 20 Lines changed: 20 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,20 @@ 1 + import type { FsAncestor , InstructionFile } from 'claude-code' 2 + 3 + /** 4 + * The found AGENTS.md files as project instruction files, one per file and 5 + * per file it `@`-imported, in walk and load order. 6 + * 7 + * An import names the file that brought it as its parent. 8 + * 9 + * @param found what `$.fs.ancestors` found, root first 10 + * @returns the instruction files, kind `project` 11 + */ 12 + export const filesOf = ( found : readonly FsAncestor [ ] ) : InstructionFile [ ] => 13 + found . flatMap ( entry => 14 + entry . parts . map ( ( part , index ) => ( { 15 + path : part . path , 16 + kind : 'project' as const , 17 + content : part . content , 18 + ... ( index > 0 && { parent : entry . parts [ 0 ] ?. path ?? part . path } ) , 19 + } ) ) , 20 + ) Collapse file ‎ mods/agents-md/hooks/files/index.ts ‎
Copy file name to clipboard + 12 Lines changed: 12 additions & 0 deletions Original file line number Diff line number Diff line change @@ -0,0 +1,12 @@ 1 + export * from './chain-root-of.js' 2 + export * from './dropped-kinds.js' 3 + export * from './files-of.js' 4 + export * from './insertion-index.js' 5 + export * from './is-claude-file-on-walk.js' 6 + export * from './is-kept-without-instructions.js' 7 + export * from './is-project-own.js' 8 + export * from './project-dir-of.js' 9 + export * from './unseen-files.js' 10 + export * from './with-project-files.js' 11 + 12 + export * as default from '.' Collapse file ‎ mods/agents-md/hooks/files/insertion-index.ts ‎
Copy file name to clipboard + 44 Lines changed: 44 additions & 0 deletions Original file line number

[truncated]
