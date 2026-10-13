---
source: "https://github.com/orvi2014/hatswap"
hn_url: "https://news.ycombinator.com/item?id=50075273"
title: "Show HN: Hatswap – one Claude Code account per client, switch mid-chat"
article_title: "GitHub - orvi2014/hatswap: Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. · GitHub"
image: "https://opengraph.githubassets.com/d560898cbcc60c8adc623fa1fdfe644fa49c1ad0cfffa810454fb4ff24c4d5af/orvi2014/hatswap"
author: "robatdas2018"
captured_at: "2026-10-13T17:58:21Z"
capture_tool: "hn-digest"
hn_id: 50075273
score: 1
comments: 0
posted_at: "2026-10-13T17:17:53Z"
tags:
  - hacker-news
---

# Show HN: Hatswap – one Claude Code account per client, switch mid-chat

- HN: [50075273](https://news.ycombinator.com/item?id=50075273)
- Source: [github.com](https://github.com/orvi2014/hatswap)
- Score: 1
- Comments: 0
- Posted: 2026-10-13T17:17:53Z

## Translation

Title: Show HN: Hatswap – one Claude Code account per client, switch mid-chat
Article title: GitHub - orvi2014/hatswap: Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. · GitHub
Description: Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. - orvi2014/hatswap

Article text:
GitHub - orvi2014/hatswap: Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. · GitHub
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
orvi2014
hatswap Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
25 Commits 25 Commits Folders and files
.github/ workflows .github/ workflows bin bin demo demo docs docs tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Every client's Claude Code account, signed in at once. Swap mid-conversation.
hatswap keeps each Claude Code account signed in at the same time and moves work between them — by resuming a conversation, or by starting a fresh session from a compact, repository-aware handoff brief. The same brief can restart an overgrown session on the same account, or seed an OpenAI Codex session.
Recorded in a sandbox with example accounts and a stand-in for claude ; reproduce with bash demo/setup.sh && vhs demo/demo.tape .
How this differs from other account switchers
Concurrent accounts. Every account stays signed in; run two accounts in two terminals at once.
Shared environment. Settings, agents, skills, plugins, hooks, rules, output styles, keybindings, memory and conversation history are shared across accounts. Only the login differs.
Switch in place. switch moves the current conversation to another account. In tmux (including desktop apps that host tabs in tmux), Terminal.app and iTerm2, the same tab restarts on the new login with its history intact.
Default account. default (or --default ) picks the account plain claude starts on; setup-shell installs the one-line hook for zsh, bash or fish.
Per-folder accounts. pin acme in a client's folder and every claude started there uses that client's account.
Account picker. Type hatswap on its own for a list of every account with its email, limits, default and pin; pick one with the arrow keys and start Claude, set the default or pin the folder with one key.
Usage per account. usage shows every account's 5-hour and weekly limits in one table, the numbers /usage shows inside Claude.
Account in the status line. statusline --install puts the account and email at the bottom of every session, ahead of any status line you already have.
VS Code, Cursor and Claude Desktop. vscode acme makes the Claude extension use acme's login in that folder; desktop acme opens Claude Desktop signed in as acme, in its own window.
Bots follow the default. exec (or the claude / codex shims from shim ) looks the account up on every call, so a long-running bot that starts claude -p per job moves to a new default without a restart.
Codex per account. Each account can have its own Codex login; handoff codex uses the one that belongs to the session's account.
Private history. Accounts can keep their conversations and memory to themselves, for client work that must stay separate.
One-line install and update . The installer sets up PATH , the shell hook and the slash commands; outdated slash commands are flagged.
Session handoff. Continue a conversation on another account with --resume , or fork a copy of it.
Brief handoff. Start a fresh session from a structured summary instead of replaying the full history.
Refresh. Restart a session that has grown too long, on the same account, from its brief.
Repository-aware briefs. Branch, uncommitted changes, recent commits and your open pull requests are collected at handoff time and treated as ground truth.
Review before launch. Open the brief in your editor and correct it before the new agent sees it.
In-session /handoff and /switch . A running Claude Code session writes its own brief, from full context, at no extra cost, or moves itself to another account.
Codex handoff. Seed an OpenAI Codex session with the same brief.
Live-session protection. Refuses to resume a conversation that is still open in another window.
Secret redaction. Passwords, API keys, tokens and private keys are removed before a transcript is summarised.
Operating system
macOS or Linux
Shell
Bash 3.2 or later (the macOS default works)
Runtime
Python 3
Claude Code
claude on your PATH
Optional
gh for open pull requests in briefs · codex for handoff codex · tmux for switching a session in place
Setup
curl -fsSL https://raw.githubusercontent.com/orvi2014/hatswap/main/install.sh | bash
This installs hatswap to ~/.local/bin , adds that folder to your PATH if needed, installs the default-account shell hook and the /handoff and /switch commands. Open a new terminal afterwards. Every change is marked and reversible (see Uninstall ).
Options go after bash -s -- , for example curl -fsSL …/install.sh | bash -s -- --no-shell :
Update later with hatswap update . It also refreshes the slash commands.
mkdir -p ~ /.local/bin
curl -fsSL https://raw.githubusercontent.com/orvi2014/hatswap/main/bin/hatswap \
-o ~ /.local/bin/hatswap
chmod +x ~ /.local/bin/hatswap
If ~/.local/bin is not on your PATH , add it to your shell profile.
git clone https://github.com/orvi2014/hatswap.git
ln -s " $PWD /hatswap/bin/hatswap " ~ /.local/bin/hatswap
git pull in the clone (or hatswap update ) then updates the command.
hatswap list
Your existing Claude Code login appears as main . It is used as-is and never modified.
hatswap add work
A browser window opens. Sign in with the other account — if the browser is already signed in to your main account, use a private window. Pass --email you@company.com to prefill the address, or --console to use an Anthropic Console (API billing) login.
hatswap list
# main you@personal.com (~/.claude)
# work you@company.com
4. Let plain claude follow a default account (optional)
hatswap setup-shell # adds a guarded hook to your shell's startup file
hatswap default work # plain `claude` now starts on "work"
setup-shell detects zsh, bash or fish (override with --shell ), is safe to run twice, and is undone with setup-shell --remove . Open a new terminal afterwards.
5. Enable /handoff inside Claude Code (optional)
hatswap install-command
This adds ~/.claude/commands/handoff.md and switch.md , which every account shares. In any Claude Code session you can then run /switch work , /handoff work , /handoff codex or /handoff .
Formerly claude-account . Renamed in 0.5.0 to avoid confusion with an unrelated project of the same name. Old links redirect here, and the claude-account command keeps working as an alias.
claude-account update # installs hatswap next to it; claude-account becomes an alias
Or re-run the installer. Nothing else changes: your accounts, logins, briefs, default account, pins and shell hook keep working. The data folder is still ~/.claude-accounts , because Claude Code ties each account's Keychain login to that path; moving it would sign every account out. Older names are still honoured:
Run hatswap install-command once to update /handoff and /switch (they now call hatswap ); hatswap list reminds you.
hatswap setup-shell --remove # if you added the shell hook
hatswap remove work # repeat for each added account (also signs out its Codex login)
rm -f ~ /.local/bin/hatswap ~ /.local/bin/claude-account ~ /.claude/commands/handoff.md ~ /.claude/commands/switch.md
rm -r ~ /.claude-accounts # saved briefs
The installer's PATH line is marked # hatswap PATH in your shell profile; delete that line too.
hatswap use work # or simply: hatswap work
hatswap main # back to the main account (~/.claude)
Run each in its own terminal to use both at once. hatswap whoami shows which account the current shell or Claude session is on, and which one a new claude starts on (inside Claude Code, ! hatswap whoami or /status ).
Switch the account of the current conversation
! hatswap switch work # or /switch work (after install-command)
! hatswap switch work --default # also make "work" the default for new sessions
A running Claude Code process cannot change its login, so switch restarts it: the same conversation (full history, same ID) resumes on the other account, in the same tab :
Run from a different terminal, switch restarts the session in its own tab when it can, and otherwise closes it there and resumes it in the terminal you ran it from.
hatswap default work # plain `claude` now starts on "work"
hatswap use work --default # set it and start a session
hatswap default main # back to ~/.claude
hatswap default # show the current default
The first time you set a default, hatswap offers to add a one-line hook to your shell's startup file ( ~/.zshenv , ~/.bash_profile / ~/.bashrc , or fish conf.d ) so plain claude follows it. For zsh it goes in ~/.zshenv rather than ~/.zshrc so that sessions an app reopens after a restart (with zsh -l -c claude … ) also get it; an older hook in ~/.zshrc moves itself there the next time a terminal opens. The hook reads the default on every call, so changing it later needs no reload. An explicit CLAUDE_CONFIG_DIR always wins.
cd ~ /clients/acme && hatswap pin acme # plain `claude` here, and in every subfolder, starts on "acme"
hatswap pin # show which pin applies here
hatswap pin main # this folder uses main, whatever the default
hatswap pin --clear # remove this folder's pin
pin writes a one-line .hatswap file. Inside a git repository it is added to the repository's local ignore ( .git/info/exclude ), so it is never committed. The nearest pin wins, so a subfolder can pin a different account. Order of precedence for a plain claude : an explicit CLAUDE_CONFIG_DIR , then the nearest pin, then the default account, then main . whoami shows which one applies and why, and refresh and handoff codex follow it too. Pins work through the same shell hook as the default.
Move a task to another account
hatswap sessions # conversations in this folder, and where each is open
hatswap handoff work --brief # fresh session on "work", seeded with a brief
hatswap handoff work # or resume the full conversation (close the original first)
Restart a session that has grown too long
hatswap refresh # brief + fresh session on the same account
hatswap refresh --edit # review the brief in $EDITOR first
Hand off from inside a running session
/handoff work
Claude writes the brief from its own context, saves it, and prints the command to continue — for example hatswap handoff work --from ~/.claude-accounts/.briefs/….md .
hatswap handof

[truncated]

## Original Extract

Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. - orvi2014/hatswap

GitHub - orvi2014/hatswap: Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex. · GitHub
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
orvi2014
hatswap Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Wear a different hat for each client: run several Claude Code accounts side by side, move a live conversation between them, pin folders to accounts, and hand work to Codex.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
25 Commits 25 Commits Folders and files
.github/ workflows .github/ workflows bin bin demo demo docs docs tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Every client's Claude Code account, signed in at once. Swap mid-conversation.
hatswap keeps each Claude Code account signed in at the same time and moves work between them — by resuming a conversation, or by starting a fresh session from a compact, repository-aware handoff brief. The same brief can restart an overgrown session on the same account, or seed an OpenAI Codex session.
Recorded in a sandbox with example accounts and a stand-in for claude ; reproduce with bash demo/setup.sh && vhs demo/demo.tape .
How this differs from other account switchers
Concurrent accounts. Every account stays signed in; run two accounts in two terminals at once.
Shared environment. Settings, agents, skills, plugins, hooks, rules, output styles, keybindings, memory and conversation history are shared across accounts. Only the login differs.
Switch in place. switch moves the current conversation to another account. In tmux (including desktop apps that host tabs in tmux), Terminal.app and iTerm2, the same tab restarts on the new login with its history intact.
Default account. default (or --default ) picks the account plain claude starts on; setup-shell installs the one-line hook for zsh, bash or fish.
Per-folder accounts. pin acme in a client's folder and every claude started there uses that client's account.
Account picker. Type hatswap on its own for a list of every account with its email, limits, default and pin; pick one with the arrow keys and start Claude, set the default or pin the folder with one key.
Usage per account. usage shows every account's 5-hour and weekly limits in one table, the numbers /usage shows inside Claude.
Account in the status line. statusline --install puts the account and email at the bottom of every session, ahead of any status line you already have.
VS Code, Cursor and Claude Desktop. vscode acme makes the Claude extension use acme's login in that folder; desktop acme opens Claude Desktop signed in as acme, in its own window.
Bots follow the default. exec (or the claude / codex shims from shim ) looks the account up on every call, so a long-running bot that starts claude -p per job moves to a new default without a restart.
Codex per account. Each account can have its own Codex login; handoff codex uses the one that belongs to the session's account.
Private history. Accounts can keep their conversations and memory to themselves, for client work that must stay separate.
One-line install and update . The installer sets up PATH , the shell hook and the slash commands; outdated slash commands are flagged.
Session handoff. Continue a conversation on another account with --resume , or fork a copy of it.
Brief handoff. Start a fresh session from a structured summary instead of replaying the full history.
Refresh. Restart a session that has grown too long, on the same account, from its brief.
Repository-aware briefs. Branch, uncommitted changes, recent commits and your open pull requests are collected at handoff time and treated as ground truth.
Review before launch. Open the brief in your editor and correct it before the new agent sees it.
In-session /handoff and /switch . A running Claude Code session writes its own brief, from full context, at no extra cost, or moves itself to another account.
Codex handoff. Seed an OpenAI Codex session with the same brief.
Live-session protection. Refuses to resume a conversation that is still open in another window.
Secret redaction. Passwords, API keys, tokens and private keys are removed before a transcript is summarised.
Operating system
macOS or Linux
Shell
Bash 3.2 or later (the macOS default works)
Runtime
Python 3
Claude Code
claude on your PATH
Optional
gh for open pull requests in briefs · codex for handoff codex · tmux for switching a session in place
Setup
curl -fsSL https://raw.githubusercontent.com/orvi2014/hatswap/main/install.sh | bash
This installs hatswap to ~/.local/bin , adds that folder to your PATH if needed, installs the default-account shell hook and the /handoff and /switch commands. Open a new terminal afterwards. Every change is marked and reversible (see Uninstall ).
Options go after bash -s -- , for example curl -fsSL …/install.sh | bash -s -- --no-shell :
Update later with hatswap update . It also refreshes the slash commands.
mkdir -p ~ /.local/bin
curl -fsSL https://raw.githubusercontent.com/orvi2014/hatswap/main/bin/hatswap \
-o ~ /.local/bin/hatswap
chmod +x ~ /.local/bin/hatswap
If ~/.local/bin is not on your PATH , add it to your shell profile.
git clone https://github.com/orvi2014/hatswap.git
ln -s " $PWD /hatswap/bin/hatswap " ~ /.local/bin/hatswap
git pull in the clone (or hatswap update ) then updates the command.
hatswap list
Your existing Claude Code login appears as main . It is used as-is and never modified.
hatswap add work
A browser window opens. Sign in with the other account — if the browser is already signed in to your main account, use a private window. Pass --email you@company.com to prefill the address, or --console to use an Anthropic Console (API billing) login.
hatswap list
# main you@personal.com (~/.claude)
# work you@company.com
4. Let plain claude follow a default account (optional)
hatswap setup-shell # adds a guarded hook to your shell's startup file
hatswap default work # plain `claude` now starts on "work"
setup-shell detects zsh, bash or fish (override with --shell ), is safe to run twice, and is undone with setup-shell --remove . Open a new terminal afterwards.
5. Enable /handoff inside Claude Code (optional)
hatswap install-command
This adds ~/.claude/commands/handoff.md and switch.md , which every account shares. In any Claude Code session you can then run /switch work , /handoff work , /handoff codex or /handoff .
Formerly claude-account . Renamed in 0.5.0 to avoid confusion with an unrelated project of the same name. Old links redirect here, and the claude-account command keeps working as an alias.
claude-account update # installs hatswap next to it; claude-account becomes an alias
Or re-run the installer. Nothing else changes: your accounts, logins, briefs, default account, pins and shell hook keep working. The data folder is still ~/.claude-accounts , because Claude Code ties each account's Keychain login to that path; moving it would sign every account out. Older names are still honoured:
Run hatswap install-command once to update /handoff and /switch (they now call hatswap ); hatswap list reminds you.
hatswap setup-shell --remove # if you added the shell hook
hatswap remove work # repeat for each added account (also signs out its Codex login)
rm -f ~ /.local/bin/hatswap ~ /.local/bin/claude-account ~ /.claude/commands/handoff.md ~ /.claude/commands/switch.md
rm -r ~ /.claude-accounts # saved briefs
The installer's PATH line is marked # hatswap PATH in your shell profile; delete that line too.
hatswap use work # or simply: hatswap work
hatswap main # back to the main account (~/.claude)
Run each in its own terminal to use both at once. hatswap whoami shows which account the current shell or Claude session is on, and which one a new claude starts on (inside Claude Code, ! hatswap whoami or /status ).
Switch the account of the current conversation
! hatswap switch work # or /switch work (after install-command)
! hatswap switch work --default # also make "work" the default for new sessions
A running Claude Code process cannot change its login, so switch restarts it: the same conversation (full history, same ID) resumes on the other account, in the same tab :
Run from a different terminal, switch restarts the session in its own tab when it can, and otherwise closes it there and resumes it in the terminal you ran it from.
hatswap default work # plain `claude` now starts on "work"
hatswap use work --default # set it and start a session
hatswap default main # back to ~/.claude
hatswap default # show the current default
The first time you set a default, hatswap offers to add a one-line hook to your shell's startup file ( ~/.zshenv , ~/.bash_profile / ~/.bashrc , or fish conf.d ) so plain claude follows it. For zsh it goes in ~/.zshenv rather than ~/.zshrc so that sessions an app reopens after a restart (with zsh -l -c claude … ) also get it; an older hook in ~/.zshrc moves itself there the next time a terminal opens. The hook reads the default on every call, so changing it later needs no reload. An explicit CLAUDE_CONFIG_DIR always wins.
cd ~ /clients/acme && hatswap pin acme # plain `claude` here, and in every subfolder, starts on "acme"
hatswap pin # show which pin applies here
hatswap pin main # this folder uses main, whatever the default
hatswap pin --clear # remove this folder's pin
pin writes a one-line .hatswap file. Inside a git repository it is added to the repository's local ignore ( .git/info/exclude ), so it is never committed. The nearest pin wins, so a subfolder can pin a different account. Order of precedence for a plain claude : an explicit CLAUDE_CONFIG_DIR , then the nearest pin, then the default account, then main . whoami shows which one applies and why, and refresh and handoff codex follow it too. Pins work through the same shell hook as the default.
Move a task to another account
hatswap sessions # conversations in this folder, and where each is open
hatswap handoff work --brief # fresh session on "work", seeded with a brief
hatswap handoff work # or resume the full conversation (close the original first)
Restart a session that has grown too long
hatswap refresh # brief + fresh session on the same account
hatswap refresh --edit # review the brief in $EDITOR first
Hand off from inside a running session
/handoff work
Claude writes the brief from its own context, saves it, and prints the command to continue — for example hatswap handoff work --from ~/.claude-accounts/.briefs/….md .
hatswap handof

[truncated]
