---
source: "https://github.com/ShadowWalker2014/rig"
hn_url: "https://news.ycombinator.com/item?id=49834499"
title: "Show HN: Rig – Open-source cloud desktops for AI agents"
article_title: "GitHub - ShadowWalker2014/rig: Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. · GitHub"
image: "https://repository-images.githubusercontent.com/1384264491/da510f74-4251-4711-9146-0c6a7a6c72c2"
author: "fengjiabo2400"
captured_at: "2026-09-24T18:15:13Z"
capture_tool: "hn-digest"
hn_id: 49834499
score: 2
comments: 0
posted_at: "2026-09-24T18:01:51Z"
tags:
  - hacker-news
---

# Show HN: Rig – Open-source cloud desktops for AI agents

- HN: [49834499](https://news.ycombinator.com/item?id=49834499)
- Source: [github.com](https://github.com/ShadowWalker2014/rig)
- Score: 2
- Comments: 0
- Posted: 2026-09-24T18:01:51Z

## Translation

Title: Show HN: Rig – Open-source cloud desktops for AI agents
Article title: GitHub - ShadowWalker2014/rig: Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. · GitHub
Description: Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. - ShadowWalker2014/rig

Article text:
GitHub - ShadowWalker2014/rig: Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. · GitHub
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
ShadowWalker2014
/
rig
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
29 Commits 29 Commits Folders and files
assets assets bin bin config config docs docs image image scripts scripts skills/ rig skills/ rig src src test test .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Rig gives each coding agent a Linux desktop in the cloud with your code, a running dev server, tests and a signed-in Chrome, so your laptop stays fast and you can run many agents at once.
Watch the 45-second demo (with sound)
Setup ·
Logins ·
Repos ·
Saved desktops ·
Computer use ·
Quick start ·
Manage ·
Security ·
FAQ ·
Docs
Claude Code, Cursor or Codex keeps editing code on your laptop. The heavy work runs in its cloud desktop, which sleeps when nobody uses it and wakes in about two seconds.
One coding agent barely uses your laptop. What freezes it is everything around the agent: a dev server per branch, a Chrome for checking the UI, and test runs. Two or three tasks in, a 16 to 24 GB laptop starts swapping.
Rig moves that work into one cloud desktop per branch. Your laptop runs only the agents.
Sleeps when idle
A cloud desktop pauses after 15 quiet minutes with everything still running inside. It costs nothing while paused and wakes in about two seconds.
Starts signed in
One command copies your logins from Chrome, Arc, Edge, Brave, Firefox or Safari. Every new cloud desktop starts with them.
Computer use and MCP
Your agent can see the whole screen and click, type and press keys, like Claude's computer use. rig mcp gives Claude Code these as native tools that return screenshots.
You can take over
Open the desktop's screen in any browser tab to finish a login or type a 2FA code. It is the same Chrome the agent drives.
No push needed
rig sync sends your local edits, committed or not.
Your whole toolbelt
Node, Bun, Python, Playwright, Puppeteer, the Vercel, Cloudflare, Railway, Fly, AWS, Google Cloud, GitHub and Stripe CLIs, Claude Code, Codex and more.
Private by default
Ports are never public, your API key never touches the repo you work in, and cookie values are never printed.
Built for many
Filters and bulk actions work across thousands of cloud desktops, and one command cleans up the stale ones.
See it working
Stills from the demo video . Every command in it ran for real on a real cloud desktop.
Two real screenshots, taken with Rig:
rig up gives this branch a cloud desktop. It copies your code in, installs packages and starts the dev server.
rig sync sends your local edits to it.
rig exec runs tests there, and rig browser drives its signed-in Chrome.
rig desktop lets you see and control its screen.
rig save makes a signed-in cloud desktop your default desktop, so every new one starts as a copy of it.
In the CLI, a cloud desktop is called a box .
Setup takes about 20 minutes, most of it waiting for the image to build. You need Bun 1.2+ and an E2B account. The full walkthrough is in the setup guide .
npm install -g @shadowwalker2014/rig # or: bun add -g @shadowwalker2014/rig
rig help # every command; `rig help <command>` for one
2. Add your E2B key
rig login # macOS: the Keychain asks for the key, so it never shows on screen
Not on a Mac? Copy .env.example to ~/.config/rig/.env , run chmod 600 on it, and set RIG_E2B_API_KEY .
3. Build the image (once, about 10 minutes)
rig image build
4. Sign in to your tools
rig cookies push # your browser's logins, except banking and payments
For command-line tools, sign in inside a cloud desktop, then save it:
rig new # an empty cloud desktop; prints its id
rig desktop < id > # open the link; sign in to Google, `gh auth login`, `vercel login`…
rig save < id > # every new cloud desktop now starts signed in
rig saved # your saved desktops; * is the one new boxes start from
Keep 1Password signed out in a desktop you save. A saved desktop is stored with E2B, and a signed-in 1Password in it would put your vault session there too.
rig skill install # Claude Code
6. Check everything
rig doctor # every line should start with ✓
Bring your logins from your browser
Sign in to anything in the browser you already use, then copy those logins into your cloud desktops. Run it again whenever you sign in to something new.
See what is there first. Neither command reads a cookie value or asks for access:
rig cookies browsers # Chrome, Edge, Brave, Arc, Comet, Chromium, Vivaldi, Opera, Firefox, Safari
rig cookies sites --from chrome # each site's cookie count, and which ones stay out by default
How it stays safe:
Rig prints only site names and counts, never a value.
Cookies are decrypted in memory and sent over E2B's encrypted connection straight into the cloud desktop's Chrome.
For Chrome-family browsers, macOS asks you to approve access each time. Click Allow , not Always Allow.
--all refuses to run from an agent or script. Banking and payment sessions stay on your laptop unless you ask for them.
Some sites refuse a login copied from another computer, and Google accounts do. For those, sign in once inside the cloud desktop with rig desktop , then run rig save . Details are in docs/cookies.md .
cd my-repo
rig up # the first time, it works out how the repo runs and asks about env files
The first rig up in a repo runs rig init . It detects the package manager, the dev command and its port, and asks before copying gitignored env files like .env.local into the box. The answers go in a small rig.json , and rig status shows the settings in use. If the box can't read a private repo, rig up stops before anything slow and prints the exact fix. Details are in using Rig in a repo .
{ "setup" : " bun install " , "dev" : " bun run dev " , "port" : 3000 , "copy" : [ " .env.local " ], "submodules" : true }
Every task
cd my-repo
rig up # this branch's cloud desktop, with the dev server running
rig sync # after editing locally
rig exec -- bun test # any command, in its copy of the repo
rig exec -- ' bunx tsc --noEmit && bun run lint '
rig browser -- open http://localhost:3000 # drive its signed-in Chrome
rig browser -- snapshot -i # the page's buttons and fields, for an agent
rig shot # screenshot to a local file
rig port 3000 # open its app at http://localhost:3000 on your laptop
rig logs # the dev server's output
Running parallel agents on one branch? rig up --new gives each its own cloud desktop. Pass -b <id> to the other commands.
rig desktop < box > # or just `rig desktop` inside a repo
Rig prints a private link. Open it in any browser tab to see and control the cloud desktop's screen: finish a login, type a 2FA code, or watch the agent work.
It stays awake while you use it. The cloud desktop does not sleep while the tab is open, and sleeps 15 minutes after you close it.
The link keeps working after the terminal or agent that opened it has gone. rig desktop again prints the same link, and rig desktop --stop closes it.
Copy and paste go through the clipboard panel on the left edge of the view.
It is private. The link works only on your machine, and asks for a one-time password it already carries.
Set up a cloud desktop once, sign in and install what you need, then save it. Every new cloud desktop starts as a copy of your default saved desktop: its files, its logins and its already-running Chrome.
Rig reminds you to save. Closing rig desktop on a clean box prints the rig save command, and rig doctor flags a missing default.
Save a clean cloud desktop made with rig new . Rig refuses to save one that ran a repo's code, because that code could have planted something that would spread to every cloud desktop started from it.
Your agent can use the whole cloud desktop, not only its Chrome. It can take a screenshot, click, type, press keys, scroll and drag. These are the same actions as Claude's computer-use tool, for things outside a web page: a terminal window, a system dialog, a browser extension or a file picker.
rig screen # screenshot the desktop; prints the file path
rig click 640 88 # click at a point in that screenshot
rig type " hello " # type into whatever has focus
rig key ctrl+l # press keys: Enter, Tab, ctrl+shift+t, cmd+a…
rig zoom 0 0 400 200 # a 2x close-up, for small text
For Claude Code, add Rig as an MCP server. Claude then gets computer , browser , shell and boxes tools and sees screenshots directly as images:
claude mcp add rig -- rig mcp
Use rig browser for anything inside a web page. It reads the page's structure, so it is faster and more reliable than clicking pixels. More in computer use .
Rig ships with an agent skill. It tells your coding agent to run dev servers, tests and browser checks in its cloud desktop instead of on your laptop, and how to hand the screen to you for a login.
rig skill install # Claude Code
npx skills add ShadowWalker2014/rig # Claude Code, Cursor, Codex, opencode and more
rig guide # the same instructions, for any other agent's AGENTS.md
Every command also explains itself: rig help <command> .
Paste this into a fresh Claude Code (or any agent) session, inside the repo you want to work on:
Use Rig (https://github.com/ShadowWalker2014/rig) to run this repo in a cloud desktop
instead of on my laptop. Keep editing code locally; run the heavy work in the cloud.
1. Run `rig guide` and follow it. Run `rig help <command>` whenever you are unsure.
2. Start this branch's cloud desktop with `rig up`. The first time in this repo it runs
`rig init`. Ask me before copying any env file.
3. After every edit, run `rig sync`. Run tests with `rig exec -- <command>`.
4. Check the UI with `rig browser -- open http://localhost:<port>`, `rig browser -- snapshot -i`
and `rig shot`. Read logs with `rig logs`.
5. If a site needs a login or a 2FA code, run `rig desktop` and give me the link.
If a site rejects the cloud login, ask me to sign in to Chrome on my laptop,
then run `rig cookies push --site <site>`.
6. For anything outside a web page, use `rig screen`, `rig click`, `rig type` and `rig key`.
7. Never run `rig save`, `rig kill` or `rig cookies push --all` without asking me first.
8. When you finish, run `rig pause`.
The same instructions ship with Rig: rig skill install adds them as a Claude Code skill, and rig guide prints them.
rig ls # every cloud desktop: state, repo

[truncated]

## Original Extract

Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. - ShadowWalker2014/rig

GitHub - ShadowWalker2014/rig: Open-source cloud desktops for AI agents. Each agent gets a Linux desktop with a dev server, tests and a signed-in Chrome that pauses when idle, so your laptop stays fast. · GitHub
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
ShadowWalker2014
/
rig
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
29 Commits 29 Commits Folders and files
assets assets bin bin config config docs docs image image scripts scripts skills/ rig skills/ rig src src test test .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Rig gives each coding agent a Linux desktop in the cloud with your code, a running dev server, tests and a signed-in Chrome, so your laptop stays fast and you can run many agents at once.
Watch the 45-second demo (with sound)
Setup ·
Logins ·
Repos ·
Saved desktops ·
Computer use ·
Quick start ·
Manage ·
Security ·
FAQ ·
Docs
Claude Code, Cursor or Codex keeps editing code on your laptop. The heavy work runs in its cloud desktop, which sleeps when nobody uses it and wakes in about two seconds.
One coding agent barely uses your laptop. What freezes it is everything around the agent: a dev server per branch, a Chrome for checking the UI, and test runs. Two or three tasks in, a 16 to 24 GB laptop starts swapping.
Rig moves that work into one cloud desktop per branch. Your laptop runs only the agents.
Sleeps when idle
A cloud desktop pauses after 15 quiet minutes with everything still running inside. It costs nothing while paused and wakes in about two seconds.
Starts signed in
One command copies your logins from Chrome, Arc, Edge, Brave, Firefox or Safari. Every new cloud desktop starts with them.
Computer use and MCP
Your agent can see the whole screen and click, type and press keys, like Claude's computer use. rig mcp gives Claude Code these as native tools that return screenshots.
You can take over
Open the desktop's screen in any browser tab to finish a login or type a 2FA code. It is the same Chrome the agent drives.
No push needed
rig sync sends your local edits, committed or not.
Your whole toolbelt
Node, Bun, Python, Playwright, Puppeteer, the Vercel, Cloudflare, Railway, Fly, AWS, Google Cloud, GitHub and Stripe CLIs, Claude Code, Codex and more.
Private by default
Ports are never public, your API key never touches the repo you work in, and cookie values are never printed.
Built for many
Filters and bulk actions work across thousands of cloud desktops, and one command cleans up the stale ones.
See it working
Stills from the demo video . Every command in it ran for real on a real cloud desktop.
Two real screenshots, taken with Rig:
rig up gives this branch a cloud desktop. It copies your code in, installs packages and starts the dev server.
rig sync sends your local edits to it.
rig exec runs tests there, and rig browser drives its signed-in Chrome.
rig desktop lets you see and control its screen.
rig save makes a signed-in cloud desktop your default desktop, so every new one starts as a copy of it.
In the CLI, a cloud desktop is called a box .
Setup takes about 20 minutes, most of it waiting for the image to build. You need Bun 1.2+ and an E2B account. The full walkthrough is in the setup guide .
npm install -g @shadowwalker2014/rig # or: bun add -g @shadowwalker2014/rig
rig help # every command; `rig help <command>` for one
2. Add your E2B key
rig login # macOS: the Keychain asks for the key, so it never shows on screen
Not on a Mac? Copy .env.example to ~/.config/rig/.env , run chmod 600 on it, and set RIG_E2B_API_KEY .
3. Build the image (once, about 10 minutes)
rig image build
4. Sign in to your tools
rig cookies push # your browser's logins, except banking and payments
For command-line tools, sign in inside a cloud desktop, then save it:
rig new # an empty cloud desktop; prints its id
rig desktop < id > # open the link; sign in to Google, `gh auth login`, `vercel login`…
rig save < id > # every new cloud desktop now starts signed in
rig saved # your saved desktops; * is the one new boxes start from
Keep 1Password signed out in a desktop you save. A saved desktop is stored with E2B, and a signed-in 1Password in it would put your vault session there too.
rig skill install # Claude Code
6. Check everything
rig doctor # every line should start with ✓
Bring your logins from your browser
Sign in to anything in the browser you already use, then copy those logins into your cloud desktops. Run it again whenever you sign in to something new.
See what is there first. Neither command reads a cookie value or asks for access:
rig cookies browsers # Chrome, Edge, Brave, Arc, Comet, Chromium, Vivaldi, Opera, Firefox, Safari
rig cookies sites --from chrome # each site's cookie count, and which ones stay out by default
How it stays safe:
Rig prints only site names and counts, never a value.
Cookies are decrypted in memory and sent over E2B's encrypted connection straight into the cloud desktop's Chrome.
For Chrome-family browsers, macOS asks you to approve access each time. Click Allow , not Always Allow.
--all refuses to run from an agent or script. Banking and payment sessions stay on your laptop unless you ask for them.
Some sites refuse a login copied from another computer, and Google accounts do. For those, sign in once inside the cloud desktop with rig desktop , then run rig save . Details are in docs/cookies.md .
cd my-repo
rig up # the first time, it works out how the repo runs and asks about env files
The first rig up in a repo runs rig init . It detects the package manager, the dev command and its port, and asks before copying gitignored env files like .env.local into the box. The answers go in a small rig.json , and rig status shows the settings in use. If the box can't read a private repo, rig up stops before anything slow and prints the exact fix. Details are in using Rig in a repo .
{ "setup" : " bun install " , "dev" : " bun run dev " , "port" : 3000 , "copy" : [ " .env.local " ], "submodules" : true }
Every task
cd my-repo
rig up # this branch's cloud desktop, with the dev server running
rig sync # after editing locally
rig exec -- bun test # any command, in its copy of the repo
rig exec -- ' bunx tsc --noEmit && bun run lint '
rig browser -- open http://localhost:3000 # drive its signed-in Chrome
rig browser -- snapshot -i # the page's buttons and fields, for an agent
rig shot # screenshot to a local file
rig port 3000 # open its app at http://localhost:3000 on your laptop
rig logs # the dev server's output
Running parallel agents on one branch? rig up --new gives each its own cloud desktop. Pass -b <id> to the other commands.
rig desktop < box > # or just `rig desktop` inside a repo
Rig prints a private link. Open it in any browser tab to see and control the cloud desktop's screen: finish a login, type a 2FA code, or watch the agent work.
It stays awake while you use it. The cloud desktop does not sleep while the tab is open, and sleeps 15 minutes after you close it.
The link keeps working after the terminal or agent that opened it has gone. rig desktop again prints the same link, and rig desktop --stop closes it.
Copy and paste go through the clipboard panel on the left edge of the view.
It is private. The link works only on your machine, and asks for a one-time password it already carries.
Set up a cloud desktop once, sign in and install what you need, then save it. Every new cloud desktop starts as a copy of your default saved desktop: its files, its logins and its already-running Chrome.
Rig reminds you to save. Closing rig desktop on a clean box prints the rig save command, and rig doctor flags a missing default.
Save a clean cloud desktop made with rig new . Rig refuses to save one that ran a repo's code, because that code could have planted something that would spread to every cloud desktop started from it.
Your agent can use the whole cloud desktop, not only its Chrome. It can take a screenshot, click, type, press keys, scroll and drag. These are the same actions as Claude's computer-use tool, for things outside a web page: a terminal window, a system dialog, a browser extension or a file picker.
rig screen # screenshot the desktop; prints the file path
rig click 640 88 # click at a point in that screenshot
rig type " hello " # type into whatever has focus
rig key ctrl+l # press keys: Enter, Tab, ctrl+shift+t, cmd+a…
rig zoom 0 0 400 200 # a 2x close-up, for small text
For Claude Code, add Rig as an MCP server. Claude then gets computer , browser , shell and boxes tools and sees screenshots directly as images:
claude mcp add rig -- rig mcp
Use rig browser for anything inside a web page. It reads the page's structure, so it is faster and more reliable than clicking pixels. More in computer use .
Rig ships with an agent skill. It tells your coding agent to run dev servers, tests and browser checks in its cloud desktop instead of on your laptop, and how to hand the screen to you for a login.
rig skill install # Claude Code
npx skills add ShadowWalker2014/rig # Claude Code, Cursor, Codex, opencode and more
rig guide # the same instructions, for any other agent's AGENTS.md
Every command also explains itself: rig help <command> .
Paste this into a fresh Claude Code (or any agent) session, inside the repo you want to work on:
Use Rig (https://github.com/ShadowWalker2014/rig) to run this repo in a cloud desktop
instead of on my laptop. Keep editing code locally; run the heavy work in the cloud.
1. Run `rig guide` and follow it. Run `rig help <command>` whenever you are unsure.
2. Start this branch's cloud desktop with `rig up`. The first time in this repo it runs
`rig init`. Ask me before copying any env file.
3. After every edit, run `rig sync`. Run tests with `rig exec -- <command>`.
4. Check the UI with `rig browser -- open http://localhost:<port>`, `rig browser -- snapshot -i`
and `rig shot`. Read logs with `rig logs`.
5. If a site needs a login or a 2FA code, run `rig desktop` and give me the link.
If a site rejects the cloud login, ask me to sign in to Chrome on my laptop,
then run `rig cookies push --site <site>`.
6. For anything outside a web page, use `rig screen`, `rig click`, `rig type` and `rig key`.
7. Never run `rig save`, `rig kill` or `rig cookies push --all` without asking me first.
8. When you finish, run `rig pause`.
The same instructions ship with Rig: rig skill install adds them as a Claude Code skill, and rig guide prints them.
rig ls # every cloud desktop: state, repo

[truncated]
