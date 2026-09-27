---
source: "https://github.com/reindent/jauvex"
hn_url: "https://news.ycombinator.com/item?id=49863700"
title: "Show HN: Jauvex, one app for Claude+Codex+Grok+Jev with two-way voice chat"
article_title: "GitHub - reindent/jauvex: Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. · GitHub"
image: "https://opengraph.githubassets.com/e6cab8dcb649df48d4b58329df131bf5f2cbb8c840bd4d8ee472c0ee423b1e2c/reindent/jauvex"
author: "daraosn"
captured_at: "2026-09-27T06:19:12Z"
capture_tool: "hn-digest"
hn_id: 49863700
score: 1
comments: 1
posted_at: "2026-09-27T05:50:56Z"
tags:
  - hacker-news
---

# Show HN: Jauvex, one app for Claude+Codex+Grok+Jev with two-way voice chat

- HN: [49863700](https://news.ycombinator.com/item?id=49863700)
- Source: [github.com](https://github.com/reindent/jauvex)
- Score: 1
- Comments: 1
- Posted: 2026-09-27T05:50:56Z

## Translation

Title: Show HN: Jauvex, one app for Claude+Codex+Grok+Jev with two-way voice chat
Article title: GitHub - reindent/jauvex: Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. · GitHub
Description: Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. - reindent/jauvex

Article text:
GitHub - reindent/jauvex: Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. · GitHub
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
reindent
/
jauvex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
assets assets electron electron scripts scripts shared shared tests tests web web .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE README.md README.md package-lock.json package-lock.json package.json package.json start.sh start.sh tsconfig.json tsconfig.json tsconfig.tools.json tsconfig.tools.json vite.config.ts vite.config.ts View all files Repository files navigation
Your coding agents, side by side, by voice. Claude, Codex and Grok in one desktop app, with Jev (TypeSafe) for the fast
decisions. Jauvex Personal, version 1.1, for macOS; Apache License 2.0. Source: github.com/reindent/jauvex ; site:
jauvex.reindent.com . Made by Reindent (one human and agents).
Jauvex is an Electron client for the Claude Code, Codex and Grok Build sessions on your Mac. Add a folder, pick up any of its
sessions or start new ones with either provider, and talk to them: a voice channel that answers in three beats (a quick
word, what it understood, a summary of the agent's answer), steers a working agent without interrupting it, names and
starts agents by voice, and lets agents talk to each other inside the app. No server: the window talks to the main
process over IPC, and your sessions stay where Claude Code and Codex keep them.
curl -fsSL https://jauvex.reindent.com/install | sh
It needs Node 22.18 or newer. It downloads this source, checks its SHA-256 and builds Jauvex on your Mac: nothing
prebuilt is downloaded, so there is nothing for Apple to notarize. Jauvex lands in Applications ( ~/Applications when
/Applications is not writable), a real app with its own name, icon and microphone permission; the source and the build
stay in ~/.jauvex/personal/app . Run the command again to update, with Jauvex closed, or let Jauvex do it: it asks
jauvex.reindent.com which version is the latest ( /api/personal/version ), at launch and every six hours, and when a newer one is out
the sidebar's footer says so ("1.2.0 is out") and the Jauvex agent asks you, once per version, whether to update now. On a yes
(or "update the app" at any time) it runs node scripts/jauvex.ts update : refused while other agents work ( --now on your word);
otherwise the app fetches the same install command, checks that it installs the version offered, leaves it and a small runner in
~/.jauvex/personal/update/ , hands the runner to launchd and quits. The runner waits for the app to exit, runs the install command
(it rebuilds the app on your Mac and opens it), opens the old app again if it does not finish, and removes it; its log is
update/update.log . Only the app the install command made updates itself; a clone updates with git (T-165). To remove it, quit it and delete
Jauvex.app and ~/.jauvex/personal/app ; its settings stay in ~/.jauvex/personal . To work on the code, clone this
repository instead: npm start runs it from the clone, and npm run app makes the same app in tmp/mac-app/Jauvex.app
( scripts/mac-app.ts : Electron's app renamed Jauvex, with the built app, the Whisper models, Claude and Codex inside, signed
ad hoc on your Mac).
Jauvex needs a few things that are not in this repository; the install command and npm start take care of most of them. The welcome screen (on the first start, and from Jauvex
settings after) checks each of them and tells you what is missing.
Node 22.18 or newer (it runs the TypeScript scripts and checks as they are), then npm install in this folder (it also fetches the Electron binary; if the app ever says
Electron.app does not exist , run npx install-electron ).
Claude or Codex, signed in: at least one is a must. Both binaries come with npm install (the Claude Agent SDK
brings Claude Code, @openai/codex brings Codex), and Jauvex uses the account each one is signed in to on this Mac.
You sign in with their own command lines, in Terminal: claude auth login (Claude Code; to install it,
curl -fsSL https://claude.ai/install.sh | bash ) or codex login (Codex; brew install codex ). Jauvex signs no one in
itself: Anthropic does not let apps built on its Agent SDK offer the Claude.ai login. The welcome screen and the accounts
panel (the Jauvex button below the sidebar) say who is signed in and give the command; with one signed in, the Jauvex
agent can walk you through the other.
Grok, optional : Grok Build, xAI's coding agent, when it is on this Mac ( curl -fsSL https://x.ai/cli/install.sh | bash ,
then grok login ). Jauvex runs it as grok agent stdio with your Grok account and settings; without it, Grok is simply not
offered.
whisper.cpp for the ears : npm start installs it with Homebrew if it is missing and downloads the two models into
models/ (git-ignored): ggml-small-q5_1.bin for the transcript and ggml-base-q5_1.bin for the live words while you
speak. Each is checked against the size and SHA-256 Hugging Face lists for it ( scripts/models.sh ): a download goes to a
.part file and becomes the model once it is whole, so one cut short is downloaded again the next time, never taken for done. Other models from huggingface.co/ggerganov/whisper.cpp can be dropped in the same folder and picked in the
voice settings ( ggml-large-v3-turbo-q5_0.bin , 574 MB, hears better and is still quick on Apple silicon). Without
Whisper you can still type.
A TypeSafe key for Jev , optional: in TYPESAFE_API_KEY or the file ~/.typesafe/token , the key alone (the way Hugging Face keeps its token; the older ~/.typesafe/jev still works). Without it the small voice model
makes the decisions Jev would make, a little more slowly.
The microphone : macOS asks the first time voice mode is switched on.
The voice is macOS's System voice, through say . On a fresh Mac that is the basic Samantha, which sounds robotic:
pick a Siri voice in System Settings > Accessibility > Spoken Content > System voice. An app cannot choose a Siri
voice by itself ( say -v falls back to Samantha); only that setting reaches them. The welcome screen checks this and
has a button that opens the pane.
npm start # builds, then launches through macOS LaunchServices (start.sh)
npm run dev # Vite + Electron with reload, for working on the UI
Put the folder somewhere plain, such as ~/Coding/jauvex : macOS protects Documents, Desktop, Downloads and iCloud Drive,
and an app started there cannot read its own files until it has been granted access ( npm start then starts it from the
terminal instead, which works but attributes the permission prompts to the terminal). Only one copy may run at a time
(the app holds a lock; a second launch focuses the first). It keeps its own state in its data folder, ~/.jauvex/personal
(every copy, run from source or compiled; data/ below means that folder; CVC_DATA_DIR moves it), and uses ports 4340 (Vite, dev only) and 4341 (whisper-server), plus 4342 for the live words (Pro uses 4320 to 4322, so both can run side by side). Jauvex is built on
your Mac from this source: npm start runs it from the Electron binary in node_modules , the install command as Jauvex.app.
Add folder : native folder dialog. Each folder is a project, like in Claude Code.
+ on a project : lists every Claude, Codex and Grok session recorded for that folder (title, first prompt,
provider, branch, size, last activity), with a switch to see all, only Claude's or only Codex's (with counts) and each
provider's mark on its rows. Tick the ones you want; they appear under the project in the sidebar.
Folders fold, the sidebar resizes : a click on a folder's name (its folder icon open or closed) folds its sessions
away, and it stays folded after a reload. The sidebar's right edge drags to any width from 220 to 560 px, kept too.
Only a signed-in provider can be chosen : a provider that is not signed in on this Mac is greyed out, with the reason,
in the provider selector, in the Jauvex agent's move selector and in the default-agent setting; a new session never
starts on one.
What agents are told about the thread : it is plain markdown, so no LaTeX (formulas and matrices as plain text or a
code block), no Mermaid or other diagram languages (plain-text drawings in a code block), no raw HTML, local files as
paths in backticks rather than links, images as markdown images. Terminal colour codes in tool output are stripped.
A provider's failure is an error, not an answer : a turn that fails, or an answer that is the provider's own failure
text (an organisation that disabled subscription access, an allowance run out, a network error), shows as a red card
in the thread, and the voice says one line about it instead of summing it up.
One provider per session, for life : a session imported from Claude continues with Claude, one imported from
Codex continues with Codex, one from Grok with Grok (each session row carries its provider's own tiny mark, Claude's, OpenAI's or Grok's, and a Jev agent row TypeSafe's; the title bar has a chip. The marks are their owners' trademarks, used only to say whose session it is: the first three from @lobehub/icons-static-svg , MIT; TypeSafe's is its site icon, assets/typesafe.png ). A new session
lets you pick the provider in the composer until the first message is sent; the model and effort pickers follow the
provider (effort: Claude's fixed levels, or the levels Codex or Grok reports for the chosen model; applies from the next message)
(Codex models come from your account through model/list , Grok's from its agent's model list).
Open a session : the conversation loads. User messages as bubbles, Claude's replies as text,
tool calls and thinking folded into one-line rows you can expand (a Codex call shows the moment it starts and gets
its result when it completes), harness plumbing (reminders,
tool results, cross-session messages) hidden unless you toggle the eye icon. Long sessions page
from the end ("Load earlier messages").
The Jauvex agent : one session that always exists, pinned at the top of the sidebar, with the app's own folder as
its project and a briefing about the app itself. It is the entry point for everything about the app: restart it,
update it ( git pull , build, relaunch, on request), install what is missing, explain how it works, create agents and
add folders (through the app's command line, below), and develop it for contributors. Every other agent is told it
exists and to send it what concerns running the app, and that the app's code is theirs to work on too when th

[truncated]

## Original Extract

Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. - reindent/jauvex

GitHub - reindent/jauvex: Your coding agents, side by side, by voice. Claude, Codex & Grok in one desktop app, with Jev for the fast decisions. · GitHub
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
reindent
/
jauvex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
assets assets electron electron scripts scripts shared shared tests tests web web .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE README.md README.md package-lock.json package-lock.json package.json package.json start.sh start.sh tsconfig.json tsconfig.json tsconfig.tools.json tsconfig.tools.json vite.config.ts vite.config.ts View all files Repository files navigation
Your coding agents, side by side, by voice. Claude, Codex and Grok in one desktop app, with Jev (TypeSafe) for the fast
decisions. Jauvex Personal, version 1.1, for macOS; Apache License 2.0. Source: github.com/reindent/jauvex ; site:
jauvex.reindent.com . Made by Reindent (one human and agents).
Jauvex is an Electron client for the Claude Code, Codex and Grok Build sessions on your Mac. Add a folder, pick up any of its
sessions or start new ones with either provider, and talk to them: a voice channel that answers in three beats (a quick
word, what it understood, a summary of the agent's answer), steers a working agent without interrupting it, names and
starts agents by voice, and lets agents talk to each other inside the app. No server: the window talks to the main
process over IPC, and your sessions stay where Claude Code and Codex keep them.
curl -fsSL https://jauvex.reindent.com/install | sh
It needs Node 22.18 or newer. It downloads this source, checks its SHA-256 and builds Jauvex on your Mac: nothing
prebuilt is downloaded, so there is nothing for Apple to notarize. Jauvex lands in Applications ( ~/Applications when
/Applications is not writable), a real app with its own name, icon and microphone permission; the source and the build
stay in ~/.jauvex/personal/app . Run the command again to update, with Jauvex closed, or let Jauvex do it: it asks
jauvex.reindent.com which version is the latest ( /api/personal/version ), at launch and every six hours, and when a newer one is out
the sidebar's footer says so ("1.2.0 is out") and the Jauvex agent asks you, once per version, whether to update now. On a yes
(or "update the app" at any time) it runs node scripts/jauvex.ts update : refused while other agents work ( --now on your word);
otherwise the app fetches the same install command, checks that it installs the version offered, leaves it and a small runner in
~/.jauvex/personal/update/ , hands the runner to launchd and quits. The runner waits for the app to exit, runs the install command
(it rebuilds the app on your Mac and opens it), opens the old app again if it does not finish, and removes it; its log is
update/update.log . Only the app the install command made updates itself; a clone updates with git (T-165). To remove it, quit it and delete
Jauvex.app and ~/.jauvex/personal/app ; its settings stay in ~/.jauvex/personal . To work on the code, clone this
repository instead: npm start runs it from the clone, and npm run app makes the same app in tmp/mac-app/Jauvex.app
( scripts/mac-app.ts : Electron's app renamed Jauvex, with the built app, the Whisper models, Claude and Codex inside, signed
ad hoc on your Mac).
Jauvex needs a few things that are not in this repository; the install command and npm start take care of most of them. The welcome screen (on the first start, and from Jauvex
settings after) checks each of them and tells you what is missing.
Node 22.18 or newer (it runs the TypeScript scripts and checks as they are), then npm install in this folder (it also fetches the Electron binary; if the app ever says
Electron.app does not exist , run npx install-electron ).
Claude or Codex, signed in: at least one is a must. Both binaries come with npm install (the Claude Agent SDK
brings Claude Code, @openai/codex brings Codex), and Jauvex uses the account each one is signed in to on this Mac.
You sign in with their own command lines, in Terminal: claude auth login (Claude Code; to install it,
curl -fsSL https://claude.ai/install.sh | bash ) or codex login (Codex; brew install codex ). Jauvex signs no one in
itself: Anthropic does not let apps built on its Agent SDK offer the Claude.ai login. The welcome screen and the accounts
panel (the Jauvex button below the sidebar) say who is signed in and give the command; with one signed in, the Jauvex
agent can walk you through the other.
Grok, optional : Grok Build, xAI's coding agent, when it is on this Mac ( curl -fsSL https://x.ai/cli/install.sh | bash ,
then grok login ). Jauvex runs it as grok agent stdio with your Grok account and settings; without it, Grok is simply not
offered.
whisper.cpp for the ears : npm start installs it with Homebrew if it is missing and downloads the two models into
models/ (git-ignored): ggml-small-q5_1.bin for the transcript and ggml-base-q5_1.bin for the live words while you
speak. Each is checked against the size and SHA-256 Hugging Face lists for it ( scripts/models.sh ): a download goes to a
.part file and becomes the model once it is whole, so one cut short is downloaded again the next time, never taken for done. Other models from huggingface.co/ggerganov/whisper.cpp can be dropped in the same folder and picked in the
voice settings ( ggml-large-v3-turbo-q5_0.bin , 574 MB, hears better and is still quick on Apple silicon). Without
Whisper you can still type.
A TypeSafe key for Jev , optional: in TYPESAFE_API_KEY or the file ~/.typesafe/token , the key alone (the way Hugging Face keeps its token; the older ~/.typesafe/jev still works). Without it the small voice model
makes the decisions Jev would make, a little more slowly.
The microphone : macOS asks the first time voice mode is switched on.
The voice is macOS's System voice, through say . On a fresh Mac that is the basic Samantha, which sounds robotic:
pick a Siri voice in System Settings > Accessibility > Spoken Content > System voice. An app cannot choose a Siri
voice by itself ( say -v falls back to Samantha); only that setting reaches them. The welcome screen checks this and
has a button that opens the pane.
npm start # builds, then launches through macOS LaunchServices (start.sh)
npm run dev # Vite + Electron with reload, for working on the UI
Put the folder somewhere plain, such as ~/Coding/jauvex : macOS protects Documents, Desktop, Downloads and iCloud Drive,
and an app started there cannot read its own files until it has been granted access ( npm start then starts it from the
terminal instead, which works but attributes the permission prompts to the terminal). Only one copy may run at a time
(the app holds a lock; a second launch focuses the first). It keeps its own state in its data folder, ~/.jauvex/personal
(every copy, run from source or compiled; data/ below means that folder; CVC_DATA_DIR moves it), and uses ports 4340 (Vite, dev only) and 4341 (whisper-server), plus 4342 for the live words (Pro uses 4320 to 4322, so both can run side by side). Jauvex is built on
your Mac from this source: npm start runs it from the Electron binary in node_modules , the install command as Jauvex.app.
Add folder : native folder dialog. Each folder is a project, like in Claude Code.
+ on a project : lists every Claude, Codex and Grok session recorded for that folder (title, first prompt,
provider, branch, size, last activity), with a switch to see all, only Claude's or only Codex's (with counts) and each
provider's mark on its rows. Tick the ones you want; they appear under the project in the sidebar.
Folders fold, the sidebar resizes : a click on a folder's name (its folder icon open or closed) folds its sessions
away, and it stays folded after a reload. The sidebar's right edge drags to any width from 220 to 560 px, kept too.
Only a signed-in provider can be chosen : a provider that is not signed in on this Mac is greyed out, with the reason,
in the provider selector, in the Jauvex agent's move selector and in the default-agent setting; a new session never
starts on one.
What agents are told about the thread : it is plain markdown, so no LaTeX (formulas and matrices as plain text or a
code block), no Mermaid or other diagram languages (plain-text drawings in a code block), no raw HTML, local files as
paths in backticks rather than links, images as markdown images. Terminal colour codes in tool output are stripped.
A provider's failure is an error, not an answer : a turn that fails, or an answer that is the provider's own failure
text (an organisation that disabled subscription access, an allowance run out, a network error), shows as a red card
in the thread, and the voice says one line about it instead of summing it up.
One provider per session, for life : a session imported from Claude continues with Claude, one imported from
Codex continues with Codex, one from Grok with Grok (each session row carries its provider's own tiny mark, Claude's, OpenAI's or Grok's, and a Jev agent row TypeSafe's; the title bar has a chip. The marks are their owners' trademarks, used only to say whose session it is: the first three from @lobehub/icons-static-svg , MIT; TypeSafe's is its site icon, assets/typesafe.png ). A new session
lets you pick the provider in the composer until the first message is sent; the model and effort pickers follow the
provider (effort: Claude's fixed levels, or the levels Codex or Grok reports for the chosen model; applies from the next message)
(Codex models come from your account through model/list , Grok's from its agent's model list).
Open a session : the conversation loads. User messages as bubbles, Claude's replies as text,
tool calls and thinking folded into one-line rows you can expand (a Codex call shows the moment it starts and gets
its result when it completes), harness plumbing (reminders,
tool results, cross-session messages) hidden unless you toggle the eye icon. Long sessions page
from the end ("Load earlier messages").
The Jauvex agent : one session that always exists, pinned at the top of the sidebar, with the app's own folder as
its project and a briefing about the app itself. It is the entry point for everything about the app: restart it,
update it ( git pull , build, relaunch, on request), install what is missing, explain how it works, create agents and
add folders (through the app's command line, below), and develop it for contributors. Every other agent is told it
exists and to send it what concerns running the app, and that the app's code is theirs to work on too when th

[truncated]
