---
source: "https://github.com/narralabs/routi/"
hn_url: "https://news.ycombinator.com/item?id=49617611"
title: "Show HN: Routi Bot – AI bots with their own desktops on your Mac"
article_title: "GitHub - narralabs/routi: Routi Bot: persistent bots with their own minds and screens, on any AI provider · GitHub"
image: "https://opengraph.githubassets.com/1b3d8d8d9a5852aaca0b7a25b39f1437ba506f7f69c09549824f959521e6eb08/narralabs/routi"
author: "westoque"
captured_at: "2026-09-08T22:24:37Z"
capture_tool: "hn-digest"
hn_id: 49617611
score: 4
comments: 0
posted_at: "2026-09-08T21:50:00Z"
tags:
  - hacker-news
---

# Show HN: Routi Bot – AI bots with their own desktops on your Mac

- HN: [49617611](https://news.ycombinator.com/item?id=49617611)
- Source: [github.com](https://github.com/narralabs/routi/)
- Score: 4
- Comments: 0
- Posted: 2026-09-08T21:50:00Z

## Translation

Title: Show HN: Routi Bot – AI bots with their own desktops on your Mac
Article title: GitHub - narralabs/routi: Routi Bot: persistent bots with their own minds and screens, on any AI provider · GitHub
Description: Routi Bot: persistent bots with their own minds and screens, on any AI provider - narralabs/routi
HN text: been working on Routi Bot. an open source mac app inspired by Grok Bot. each bot has its own instructions, desktop and can be used with AI provider/model (openai astra, claude fable, deepseek, grok, etc). it also has a profile switcher so you can separate automation between your personal and work profiles.

Article text:
GitHub - narralabs/routi: Routi Bot: persistent bots with their own minds and screens, on any AI provider · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
narralabs
/
routi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
293 Commits 293 Commits Folders and files
.github .github apple apple containers/ desktop containers/ desktop daemon daemon docs docs packaging packaging protocol protocol scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Bots that live on your Mac, each with its own personality, model and screen. Give one a
job, like watching a market, finding a hotel, or checking a price every morning, and it
opens a browser and does it, on any AI you already pay for: a Claude plan, a ChatGPT
plan, or an API key. Routi is a Mac app, with iPhone and iPad apps for reaching the same
bots when you are away from it.
routi-bot-demo-web.mp4
“One shared Linux, a screen per bot, any model you already pay for — on the Mac that
stays on. Same idea, cleanly executed. Ship it.”
— Architecture Bot (a Grok Bot), asked to review the competition
Download Routi Bot ,
open the disk image, drag the app to Applications, open it.
Setup does the rest, in three screens: one command to paste into Terminal that installs
Routi Core — the part that keeps your bots and does the work, and starts at every login;
a sign-in with Claude or ChatGPT, or an API key; and, if you want bots that browse, the
desktop they browse in. Claude Code and Codex come with the core.
The desktop is a Linux machine in Docker Desktop ,
one screen per bot, kept apart from your own. Install Docker and set it to start at login;
setup checks for it and builds the desktop right there, and Settings → Screens shows its
state afterwards. Bots without a screen, and bots set to This Mac , don't need it.
Do this on the Mac that stays on. Give it automatic login, so the Keychain is unlocked
after a reboot, and set it to never sleep. Needs macOS 14 or newer.
Install the app anywhere and point it at the host under Settings → Routi Core. A
Tailscale name works. iPhone and iPad are the same app; they're coming.
Without the app, or with Homebrew
For a mini with no monitor, or if you'd rather manage the core yourself:
brew tap narralabs/tap && brew install routi-core && brew services start routi-core
or the install command above, which needs nothing on the Mac beforehand.
curl http://127.0.0.1:7171/health on the host says whether the core is up, and which version.
Logs: ~/.routi/logs/routid.log ; with Homebrew, brew services info routi-core .
A bot that says its screen is unavailable: Settings → Screens names which of Docker,
the desktop image, or the machine is the problem, with the button that fixes it.
Sign-in trouble: Settings → the provider's pane says what is connected and how.
curl -fsSL https://raw.githubusercontent.com/narralabs/routi/main/scripts/uninstall.sh | sh
Stops and removes the core, its login agent, the desktop machine and image in Docker,
and the API keys Routi stored in the Keychain. Your bots and conversations stay in
~/.routi , so reinstalling brings them back. To remove those and the app as well:
curl -fsSL https://raw.githubusercontent.com/narralabs/routi/main/scripts/uninstall.sh | sh -s -- --purge
Add --dry-run to either to see what would go without touching anything. Docker Desktop
itself, and the sign-ins that belong to Claude Code, Codex and Grok, are left alone —
they were yours before Routi.
Connectors. Sign in to Gmail once and every bot you allow can search, read, draft
and send, as tools rather than through the browser; Google Calendar, Drive, Slack and
GitHub follow the same way. The sign-in happens on the Mac running the core and the
token stays in its Keychain, per profile. Gmail first, after launch. Public use of
the Google ones waits on Google's app verification; until then, a bring-your-own
OAuth client.
Sign-in relay. When a bot reaches a login it cannot do itself, it asks in the
thread with a Continue, and you finish the sign-in from wherever you are — your
phone, your own browser — with the session or the code relayed to the bot's screen,
rather than taking the screen over. Today that moment is a hand-over: the bot stops,
you sign in on its desktop, it carries on.
Per-bot connectors and screens. Which connections a bot may use, chosen per bot,
the way a screen is today.
Folders. Bots grouped by topic in the sidebar — travel, trading, the company —
so a list of twenty stays readable.
Bots in a group chat. Several bots and you in one thread, each seeing the
others' messages, for work that takes more than one specialist. A maybe: the core
can already hold such a room, but whether bots talking to each other is worth the
cost and the confusion is not settled.
If you build something for everyone, it works for no one.
Routi is for a Mac that stays on: the Mac mini in the corner, the laptop you replaced,
any spare Mac you would like to turn into a small server that runs your bots all day.
The core runs there, the app opens onto it from your other Macs and from the phone.
There is no plan to support Windows or Linux, and no roadmap entry for it. Staying on
one platform is what lets the app be a real Mac app rather than a web page in a
frame: small binaries, native windows and notifications, the Keychain for every
credential, and an installer that is a drag and a double-click. Every hour that would
go to a second platform goes to the experience on this one.
pnpm install && pnpm --filter @routi/protocol build
pnpm --filter routid dev:isolated # a second core on ~/.routi-dev, port 7172
scripts/dev-app.sh # builds the app and opens it on that core
dev:isolated reloads on every daemon change; run dev-app.sh again after an app
change. Or cd apple && ./bootstrap.sh && open Routi.xcodeproj and ⌘R in Xcode, with
-daemonPort 7172 under the scheme's arguments. Either way the core serving your real
bots is never touched. How it's built and released: docs/ENGINEERING.md .
Adding an AI provider: CLAUDE.md .
Apache-2.0, © 2026 Narra Labs, LLC. Provider marks are LobeHub's , MIT — see NOTICE .
Routi Bot: persistent bots with their own minds and screens, on any AI provider
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Routi Bot: persistent bots with their own minds and screens, on any AI provider - narralabs/routi

been working on Routi Bot. an open source mac app inspired by Grok Bot. each bot has its own instructions, desktop and can be used with AI provider/model (openai astra, claude fable, deepseek, grok, etc). it also has a profile switcher so you can separate automation between your personal and work profiles.

GitHub - narralabs/routi: Routi Bot: persistent bots with their own minds and screens, on any AI provider · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
narralabs
/
routi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
293 Commits 293 Commits Folders and files
.github .github apple apple containers/ desktop containers/ desktop daemon daemon docs docs packaging packaging protocol protocol scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE NOTICE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Bots that live on your Mac, each with its own personality, model and screen. Give one a
job, like watching a market, finding a hotel, or checking a price every morning, and it
opens a browser and does it, on any AI you already pay for: a Claude plan, a ChatGPT
plan, or an API key. Routi is a Mac app, with iPhone and iPad apps for reaching the same
bots when you are away from it.
routi-bot-demo-web.mp4
“One shared Linux, a screen per bot, any model you already pay for — on the Mac that
stays on. Same idea, cleanly executed. Ship it.”
— Architecture Bot (a Grok Bot), asked to review the competition
Download Routi Bot ,
open the disk image, drag the app to Applications, open it.
Setup does the rest, in three screens: one command to paste into Terminal that installs
Routi Core — the part that keeps your bots and does the work, and starts at every login;
a sign-in with Claude or ChatGPT, or an API key; and, if you want bots that browse, the
desktop they browse in. Claude Code and Codex come with the core.
The desktop is a Linux machine in Docker Desktop ,
one screen per bot, kept apart from your own. Install Docker and set it to start at login;
setup checks for it and builds the desktop right there, and Settings → Screens shows its
state afterwards. Bots without a screen, and bots set to This Mac , don't need it.
Do this on the Mac that stays on. Give it automatic login, so the Keychain is unlocked
after a reboot, and set it to never sleep. Needs macOS 14 or newer.
Install the app anywhere and point it at the host under Settings → Routi Core. A
Tailscale name works. iPhone and iPad are the same app; they're coming.
Without the app, or with Homebrew
For a mini with no monitor, or if you'd rather manage the core yourself:
brew tap narralabs/tap && brew install routi-core && brew services start routi-core
or the install command above, which needs nothing on the Mac beforehand.
curl http://127.0.0.1:7171/health on the host says whether the core is up, and which version.
Logs: ~/.routi/logs/routid.log ; with Homebrew, brew services info routi-core .
A bot that says its screen is unavailable: Settings → Screens names which of Docker,
the desktop image, or the machine is the problem, with the button that fixes it.
Sign-in trouble: Settings → the provider's pane says what is connected and how.
curl -fsSL https://raw.githubusercontent.com/narralabs/routi/main/scripts/uninstall.sh | sh
Stops and removes the core, its login agent, the desktop machine and image in Docker,
and the API keys Routi stored in the Keychain. Your bots and conversations stay in
~/.routi , so reinstalling brings them back. To remove those and the app as well:
curl -fsSL https://raw.githubusercontent.com/narralabs/routi/main/scripts/uninstall.sh | sh -s -- --purge
Add --dry-run to either to see what would go without touching anything. Docker Desktop
itself, and the sign-ins that belong to Claude Code, Codex and Grok, are left alone —
they were yours before Routi.
Connectors. Sign in to Gmail once and every bot you allow can search, read, draft
and send, as tools rather than through the browser; Google Calendar, Drive, Slack and
GitHub follow the same way. The sign-in happens on the Mac running the core and the
token stays in its Keychain, per profile. Gmail first, after launch. Public use of
the Google ones waits on Google's app verification; until then, a bring-your-own
OAuth client.
Sign-in relay. When a bot reaches a login it cannot do itself, it asks in the
thread with a Continue, and you finish the sign-in from wherever you are — your
phone, your own browser — with the session or the code relayed to the bot's screen,
rather than taking the screen over. Today that moment is a hand-over: the bot stops,
you sign in on its desktop, it carries on.
Per-bot connectors and screens. Which connections a bot may use, chosen per bot,
the way a screen is today.
Folders. Bots grouped by topic in the sidebar — travel, trading, the company —
so a list of twenty stays readable.
Bots in a group chat. Several bots and you in one thread, each seeing the
others' messages, for work that takes more than one specialist. A maybe: the core
can already hold such a room, but whether bots talking to each other is worth the
cost and the confusion is not settled.
If you build something for everyone, it works for no one.
Routi is for a Mac that stays on: the Mac mini in the corner, the laptop you replaced,
any spare Mac you would like to turn into a small server that runs your bots all day.
The core runs there, the app opens onto it from your other Macs and from the phone.
There is no plan to support Windows or Linux, and no roadmap entry for it. Staying on
one platform is what lets the app be a real Mac app rather than a web page in a
frame: small binaries, native windows and notifications, the Keychain for every
credential, and an installer that is a drag and a double-click. Every hour that would
go to a second platform goes to the experience on this one.
pnpm install && pnpm --filter @routi/protocol build
pnpm --filter routid dev:isolated # a second core on ~/.routi-dev, port 7172
scripts/dev-app.sh # builds the app and opens it on that core
dev:isolated reloads on every daemon change; run dev-app.sh again after an app
change. Or cd apple && ./bootstrap.sh && open Routi.xcodeproj and ⌘R in Xcode, with
-daemonPort 7172 under the scheme's arguments. Either way the core serving your real
bots is never touched. How it's built and released: docs/ENGINEERING.md .
Adding an AI provider: CLAUDE.md .
Apache-2.0, © 2026 Narra Labs, LLC. Provider marks are LobeHub's , MIT — see NOTICE .
Routi Bot: persistent bots with their own minds and screens, on any AI provider
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
