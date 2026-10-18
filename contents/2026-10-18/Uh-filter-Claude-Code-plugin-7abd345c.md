---
source: "https://github.com/amirzak/claude-code-voice-uh-filter"
hn_url: "https://news.ycombinator.com/item?id=50140542"
title: "Uh filter Claude Code plugin"
article_title: "GitHub - amirzak/claude-code-voice-uh-filter: Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts · GitHub"
image: "https://opengraph.githubassets.com/5a647338145cbf3d1c70b1d8f0b502bb03e1b3905fe24d8a881b339235346700/amirzak/claude-code-voice-uh-filter"
author: "amirzak"
captured_at: "2026-10-18T13:00:37Z"
capture_tool: "hn-digest"
hn_id: 50140542
score: 1
comments: 0
posted_at: "2026-10-18T12:33:55Z"
tags:
  - hacker-news
---

# Uh filter Claude Code plugin

- HN: [50140542](https://news.ycombinator.com/item?id=50140542)
- Source: [github.com](https://github.com/amirzak/claude-code-voice-uh-filter)
- Score: 1
- Comments: 0
- Posted: 2026-10-18T12:33:55Z

## Translation

Title: Uh filter Claude Code plugin
Article title: GitHub - amirzak/claude-code-voice-uh-filter: Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts · GitHub
Description: Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts - amirzak/claude-code-voice-uh-filter

Article text:
GitHub - amirzak/claude-code-voice-uh-filter: Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts · GitHub
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
amirzak
claude-code-voice-uh-filter Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A Claude Code mod that removes spoken filler words from your prompts, made for voice dictation.
On send: strips uh , um , uhm , erm , hmm , mm (and their trailing commas) plus stutter repeats like "check check" or "where you where you" before the prompt reaches the model.
In the terminal prompt box: dictated or pasted text is cleaned as it lands, and a filler you type disappears as soon as you type the space after it.
Leaves alone: words like "album" or "uh-huh", common real doubles ("that that", "had had"), anything inside backticks or code blocks, pasted code, and slash commands.
To keep a literal "um" in a prompt, put it in backticks.
The desktop app draws its own prompt box, so there the cleanup happens only when you send.
The mod is one hooks module, hooks/register.ts , with two hooks:
prompt.submit : rewrites the text of every prompt you send, removing fillers and stutter repeats. It skips slash commands and prompts Claude Code sends on its own. Since prompt.submit is also a call other plugins can make ( $.prompt.submit ), prompts that other plugins submit pass through this cleanup too.
prompt.edit : in the terminal prompt box, cleans text that lands in one go (dictation, a paste of plain prose) and drops a filler word when you type the space after it. Pasted code and anything inside backticks is left as is.
It reads no files, runs no commands, and sends nothing anywhere.
At the prompt of a terminal Claude Code session:
/plugin install uh-filter --marketplace amirzak/claude-code-voice-uh-filter
Answer y to add the marketplace, then pick a scope.
claude plugin install uh-filter --marketplace amirzak/claude-code-voice-uh-filter
In the desktop app: click + next to the prompt box, then Plugins , Add marketplace with amirzak/claude-code-voice-uh-filter , then Add plugin .
claude plugin validate .
claude plugin test .
claude --plugin-dir .
License
Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts - amirzak/claude-code-voice-uh-filter

GitHub - amirzak/claude-code-voice-uh-filter: Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts · GitHub
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
amirzak
claude-code-voice-uh-filter Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A Claude Code mod that removes spoken filler words from your prompts, made for voice dictation.
On send: strips uh , um , uhm , erm , hmm , mm (and their trailing commas) plus stutter repeats like "check check" or "where you where you" before the prompt reaches the model.
In the terminal prompt box: dictated or pasted text is cleaned as it lands, and a filler you type disappears as soon as you type the space after it.
Leaves alone: words like "album" or "uh-huh", common real doubles ("that that", "had had"), anything inside backticks or code blocks, pasted code, and slash commands.
To keep a literal "um" in a prompt, put it in backticks.
The desktop app draws its own prompt box, so there the cleanup happens only when you send.
The mod is one hooks module, hooks/register.ts , with two hooks:
prompt.submit : rewrites the text of every prompt you send, removing fillers and stutter repeats. It skips slash commands and prompts Claude Code sends on its own. Since prompt.submit is also a call other plugins can make ( $.prompt.submit ), prompts that other plugins submit pass through this cleanup too.
prompt.edit : in the terminal prompt box, cleans text that lands in one go (dictation, a paste of plain prose) and drops a filler word when you type the space after it. Pasted code and anything inside backticks is left as is.
It reads no files, runs no commands, and sends nothing anywhere.
At the prompt of a terminal Claude Code session:
/plugin install uh-filter --marketplace amirzak/claude-code-voice-uh-filter
Answer y to add the marketplace, then pick a scope.
claude plugin install uh-filter --marketplace amirzak/claude-code-voice-uh-filter
In the desktop app: click + next to the prompt box, then Plugins , Add marketplace with amirzak/claude-code-voice-uh-filter , then Add plugin .
claude plugin validate .
claude plugin test .
claude --plugin-dir .
License
Claude Code mod that strips spoken filler words (uh, um, mm) from voice-dictated prompts
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
