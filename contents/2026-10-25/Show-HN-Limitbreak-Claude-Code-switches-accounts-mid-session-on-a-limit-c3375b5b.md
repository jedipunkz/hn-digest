---
source: "https://github.com/radu2lupu/limitbreak"
hn_url: "https://news.ycombinator.com/item?id=50237185"
title: "Show HN: Limitbreak – Claude Code switches accounts mid-session on a limit"
article_title: "GitHub - radu2lupu/limitbreak: Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. · GitHub"
image: "https://opengraph.githubassets.com/7cfd3027541de6e19fe4242d9d7e082e97aa1075c79652b5f677ff1d029900fb/radu2lupu/limitbreak"
author: "radulupu"
captured_at: "2026-10-25T23:51:36Z"
capture_tool: "hn-digest"
hn_id: 50237185
score: 1
comments: 0
posted_at: "2026-10-25T23:43:55Z"
tags:
  - hacker-news
---

# Show HN: Limitbreak – Claude Code switches accounts mid-session on a limit

- HN: [50237185](https://news.ycombinator.com/item?id=50237185)
- Source: [github.com](https://github.com/radu2lupu/limitbreak)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T23:43:55Z

## Translation

Title: Show HN: Limitbreak – Claude Code switches accounts mid-session on a limit
Article title: GitHub - radu2lupu/limitbreak: Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. · GitHub
Description: Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. - radu2lupu/limitbreak

Article text:
GitHub - radu2lupu/limitbreak: Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. · GitHub
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
radu2lupu
limitbreak Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.github/ workflows .github/ workflows assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md cli.go cli.go cli_test.go cli_test.go go.mod go.mod install.sh install.sh main.go main.go main_test.go main_test.go shell.sh shell.sh video.go video.go View all files Repository files navigation
★ L I M I T B R E A K ★
ADA [██████████] LIMIT → spent
GRACE [░░░░░░░░░░] ready → takes the hit
Use several Claude accounts from one Claude Code install, and switch between them automatically when an account's usage limit is hit, without restarting the session or losing its history.
A small local proxy (a single Go binary, standard library only, a few MB of RAM) sits between claude and Anthropic's API. It sends each request on the active account's own login. When Anthropic answers that the account's 5-hour or weekly window is spent, the proxy remembers that account until its reset time and replays the same request on the next account before Claude Code ever sees the error. To the session nothing happened.
claude ──ANTHROPIC_BASE_URL──▶ proxy 127.0.0.1:48617 ──▶ api.anthropic.com
│ active account's token
└─ 429 "rejected"? mark spent, retry on the next account
Try it without touching any real login: limitbreak video plays a simulated switch in your terminal as a little arcade battle.
git clone https://github.com/radu2lupu/limitbreak && cd limitbreak
./install.sh # builds and starts the background service
echo ' source "$HOME/.local/bin/limitbreak-shell.sh" ' >> ~ /.zshrc # or ~/.bashrc
limitbreak add # sign in another account (browser), repeat per account
Open a new terminal and run claude as usual. When one account's limit is hit, the next one answers.
What it does and does not touch
Only /v1/messages* requests are sent on an account's token. Everything else (claude.ai connectors, MCP, artifacts, profile) goes through untouched with Claude Code's own credentials, so those keep working.
It never sets CLAUDE_CODE_OAUTH_TOKEN (that makes Claude Code drop its connectors and local MCP servers).
Accounts are the logins Claude Code already holds: the default one plus every ~/.claude-<name> config directory that has a login. Tokens are read from the macOS Keychain item Claude Code keeps per config dir, or from <dir>/.credentials.json . They are never copied, logged, or exposed by GET /_status .
An expired token is refreshed by running one tiny claude -p under that account's directory, so Claude Code stays the only thing that refreshes a login.
If no account can serve a request, it passes through with Claude Code's own login: never worse than no proxy.
Only a window rejection ( 429 with anthropic-ratelimit-unified-status: rejected ) switches accounts. Ordinary rate limits go back to Claude unchanged.
Needs Go 1.22+ to build, macOS or Linux. Prebuilt binaries are attached to each release .
git clone https://github.com/radu2lupu/limitbreak
cd limitbreak
./install.sh # builds, installs ~/.local/bin/limitbreak, starts a launchd / systemd user service
Then add the printed line to ~/.zshrc or ~/.bashrc :
source " $HOME /.local/bin/limitbreak-shell.sh "
That defines a claude function that points Claude Code at the proxy when it is running and runs plain claude when it is not. Use ./install.sh --binary <prebuilt> on a host without Go.
limitbreak add # signs in a new account in your browser, names it account-N
limitbreak add work # or pick the name
limitbreak list # accounts, who they are, which one is active (*)
limitbreak remove work # sign it out and delete its login from this machine
add runs Claude Code's own browser sign-in into its own config directory ( ~/.claude-<name> ), checks the login landed, tells you which email it is, and warns if you added the same account twice. The running proxy picks it up within 30 seconds; no restart. The default account is whatever plain claude auth login signs in. Raw status for scripts: curl -s http://127.0.0.1:48617/_status .
Every switch is logged ( ~/Library/Logs/limitbreak.log on macOS, the journal on Linux): the account that ran out, when it resets, and the one now serving.
limitbreak serve --port 48617 --upstream https://api.anthropic.com --home ~ --credentials auto|keychain|file
A request path may begin with /_c/<client>/ to give a client (an agent, a tool) its own active account; the prefix is stripped before forwarding.
Terms of service. This uses the official claude binary with accounts you own, but spreading load across subscription accounts may conflict with Anthropic's terms for your plan. Read them; you are responsible for your use. Not affiliated with or endorsed by Anthropic.
Don't copy a login between machines. Refresh tokens rotate; two machines holding the same one will log each other out. Sign in once per machine.
Headless hosts. On macOS over SSH the Keychain is locked; run the proxy as a user service in the login session (what install.sh does).
The reset time is read from Anthropic's response headers, which are not a documented contract and may change.
macOS: launchctl bootout gui/$(id -u)/dev.limitbreak; rm ~/Library/LaunchAgents/dev.limitbreak.plist .
Linux: systemctl --user disable --now limitbreak .
Then remove ~/.local/bin/limitbreak* and the source line from your shell profile. Accounts you added stay as ordinary ~/.claude-<name> logins; limitbreak remove <name> deletes one.
go vet ./... && go test ./...
License
Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. - radu2lupu/limitbreak

GitHub - radu2lupu/limitbreak: Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT. · GitHub
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
radu2lupu
limitbreak Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.github/ workflows .github/ workflows assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md cli.go cli.go cli_test.go cli_test.go go.mod go.mod install.sh install.sh main.go main.go main_test.go main_test.go shell.sh shell.sh video.go video.go View all files Repository files navigation
★ L I M I T B R E A K ★
ADA [██████████] LIMIT → spent
GRACE [░░░░░░░░░░] ready → takes the hit
Use several Claude accounts from one Claude Code install, and switch between them automatically when an account's usage limit is hit, without restarting the session or losing its history.
A small local proxy (a single Go binary, standard library only, a few MB of RAM) sits between claude and Anthropic's API. It sends each request on the active account's own login. When Anthropic answers that the account's 5-hour or weekly window is spent, the proxy remembers that account until its reset time and replays the same request on the next account before Claude Code ever sees the error. To the session nothing happened.
claude ──ANTHROPIC_BASE_URL──▶ proxy 127.0.0.1:48617 ──▶ api.anthropic.com
│ active account's token
└─ 429 "rejected"? mark spent, retry on the next account
Try it without touching any real login: limitbreak video plays a simulated switch in your terminal as a little arcade battle.
git clone https://github.com/radu2lupu/limitbreak && cd limitbreak
./install.sh # builds and starts the background service
echo ' source "$HOME/.local/bin/limitbreak-shell.sh" ' >> ~ /.zshrc # or ~/.bashrc
limitbreak add # sign in another account (browser), repeat per account
Open a new terminal and run claude as usual. When one account's limit is hit, the next one answers.
What it does and does not touch
Only /v1/messages* requests are sent on an account's token. Everything else (claude.ai connectors, MCP, artifacts, profile) goes through untouched with Claude Code's own credentials, so those keep working.
It never sets CLAUDE_CODE_OAUTH_TOKEN (that makes Claude Code drop its connectors and local MCP servers).
Accounts are the logins Claude Code already holds: the default one plus every ~/.claude-<name> config directory that has a login. Tokens are read from the macOS Keychain item Claude Code keeps per config dir, or from <dir>/.credentials.json . They are never copied, logged, or exposed by GET /_status .
An expired token is refreshed by running one tiny claude -p under that account's directory, so Claude Code stays the only thing that refreshes a login.
If no account can serve a request, it passes through with Claude Code's own login: never worse than no proxy.
Only a window rejection ( 429 with anthropic-ratelimit-unified-status: rejected ) switches accounts. Ordinary rate limits go back to Claude unchanged.
Needs Go 1.22+ to build, macOS or Linux. Prebuilt binaries are attached to each release .
git clone https://github.com/radu2lupu/limitbreak
cd limitbreak
./install.sh # builds, installs ~/.local/bin/limitbreak, starts a launchd / systemd user service
Then add the printed line to ~/.zshrc or ~/.bashrc :
source " $HOME /.local/bin/limitbreak-shell.sh "
That defines a claude function that points Claude Code at the proxy when it is running and runs plain claude when it is not. Use ./install.sh --binary <prebuilt> on a host without Go.
limitbreak add # signs in a new account in your browser, names it account-N
limitbreak add work # or pick the name
limitbreak list # accounts, who they are, which one is active (*)
limitbreak remove work # sign it out and delete its login from this machine
add runs Claude Code's own browser sign-in into its own config directory ( ~/.claude-<name> ), checks the login landed, tells you which email it is, and warns if you added the same account twice. The running proxy picks it up within 30 seconds; no restart. The default account is whatever plain claude auth login signs in. Raw status for scripts: curl -s http://127.0.0.1:48617/_status .
Every switch is logged ( ~/Library/Logs/limitbreak.log on macOS, the journal on Linux): the account that ran out, when it resets, and the one now serving.
limitbreak serve --port 48617 --upstream https://api.anthropic.com --home ~ --credentials auto|keychain|file
A request path may begin with /_c/<client>/ to give a client (an agent, a tool) its own active account; the prefix is stripped before forwarding.
Terms of service. This uses the official claude binary with accounts you own, but spreading load across subscription accounts may conflict with Anthropic's terms for your plan. Read them; you are responsible for your use. Not affiliated with or endorsed by Anthropic.
Don't copy a login between machines. Refresh tokens rotate; two machines holding the same one will log each other out. Sign in once per machine.
Headless hosts. On macOS over SSH the Keychain is locked; run the proxy as a user service in the login session (what install.sh does).
The reset time is read from Anthropic's response headers, which are not a documented contract and may change.
macOS: launchctl bootout gui/$(id -u)/dev.limitbreak; rm ~/Library/LaunchAgents/dev.limitbreak.plist .
Linux: systemctl --user disable --now limitbreak .
Then remove ~/.local/bin/limitbreak* and the source line from your shell profile. Accounts you added stay as ordinary ~/.claude-<name> logins; limitbreak remove <name> deletes one.
go vet ./... && go test ./...
License
Use several Claude accounts from one Claude Code install; switches automatically when a usage limit is hit, no restart. MIT.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
