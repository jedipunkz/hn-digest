---
source: "https://github.com/michaeldmueller/hue-claude"
hn_url: "https://news.ycombinator.com/item?id=50245886"
title: "Show HN: Flash your Hue lights when Claude finishes a turn or needs permission"
article_title: "GitHub - michaeldmueller/hue-claude: Claude Code plugin for controlling Philips Hue lights on status changes · GitHub"
image: "https://opengraph.githubassets.com/f9bbcfbcefd1ecebfd610530992d3cc7e004ede5da6aaab22bcb6afebcd4c772/michaeldmueller/hue-claude"
author: "mdmueller10"
captured_at: "2026-10-26T16:33:22Z"
capture_tool: "hn-digest"
hn_id: 50245886
score: 1
comments: 0
posted_at: "2026-10-26T15:53:47Z"
tags:
  - hacker-news
---

# Show HN: Flash your Hue lights when Claude finishes a turn or needs permission

- HN: [50245886](https://news.ycombinator.com/item?id=50245886)
- Source: [github.com](https://github.com/michaeldmueller/hue-claude)
- Score: 1
- Comments: 0
- Posted: 2026-10-26T15:53:47Z

## Translation

Title: Show HN: Flash your Hue lights when Claude finishes a turn or needs permission
Article title: GitHub - michaeldmueller/hue-claude: Claude Code plugin for controlling Philips Hue lights on status changes · GitHub
Description: Claude Code plugin for controlling Philips Hue lights on status changes - michaeldmueller/hue-claude

Article text:
GitHub - michaeldmueller/hue-claude: Claude Code plugin for controlling Philips Hue lights on status changes · GitHub
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
michaeldmueller
hue-claude Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Claude Code plugin for controlling Philips Hue lights on status changes
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.claude-plugin .claude-plugin assets assets bin bin hooks hooks skills/ init skills/ init .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Flash your Philips Hue lights when Claude Code finishes a turn or is waiting
on a permission.
Doesn't permanently change lights - Uses the signaling API so that your lights return to their original state after flashing.
Local only - Talks to the bridge over your LAN.
Quick setup - Claude will guide you through initialization, which will require you to press the Hue button on your bridge once.
done.mp4
Claude needs your feedback
/plugin marketplace add michaeldmueller/hue-claude
/plugin install hue@hue-claude
/hue:init
/hue:init will search for the bridge and prompt you to press the round link button on the top. After this, it will ask which lights you'd like to control (and write it to ~/.config/hue/config.env ).
claude plugin marketplace add michaeldmueller/hue-claude
claude plugin install hue@hue-claude
Requirements
macOS or Linux. Bash 3.2 or later (the stock macOS /bin/bash works),
curl , and jq ( brew install jq / apt install jq ).
Bridge discovery uses dns-sd on macOS or avahi-browse on Linux
( apt install avahi-utils ). Without either it falls back to Hue's cloud
lookup, or you can type the IP during /hue:init .
Windows works through WSL or Git Bash, with the cloud lookup or a typed IP.
A square Hue Bridge (v2) on firmware with the CLIP v2 API (2021 or later).
Color bulbs that support signaling. hue lights shows which do.
By default, this plugin will flash green to indicate a finished turn, and flash amber and red to indicate that Claude is waiting for your feedback. These defaults can be changed to any color or duration. Create a config file at ~/.config/hue/events.conf (or tell Claude to) with your desired settings. Example:
claude-done on_off_color 2000 blue
claude-wait alternating 8000 red white
Signals are alternating (two colors), on_off_color (one), on_off (none),
and breathe (one smooth pulse in the current color, no duration or colors).
Claude Code plugin for controlling Philips Hue lights on status changes
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Claude Code plugin for controlling Philips Hue lights on status changes - michaeldmueller/hue-claude

GitHub - michaeldmueller/hue-claude: Claude Code plugin for controlling Philips Hue lights on status changes · GitHub
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
michaeldmueller
hue-claude Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Claude Code plugin for controlling Philips Hue lights on status changes
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.claude-plugin .claude-plugin assets assets bin bin hooks hooks skills/ init skills/ init .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Flash your Philips Hue lights when Claude Code finishes a turn or is waiting
on a permission.
Doesn't permanently change lights - Uses the signaling API so that your lights return to their original state after flashing.
Local only - Talks to the bridge over your LAN.
Quick setup - Claude will guide you through initialization, which will require you to press the Hue button on your bridge once.
done.mp4
Claude needs your feedback
/plugin marketplace add michaeldmueller/hue-claude
/plugin install hue@hue-claude
/hue:init
/hue:init will search for the bridge and prompt you to press the round link button on the top. After this, it will ask which lights you'd like to control (and write it to ~/.config/hue/config.env ).
claude plugin marketplace add michaeldmueller/hue-claude
claude plugin install hue@hue-claude
Requirements
macOS or Linux. Bash 3.2 or later (the stock macOS /bin/bash works),
curl , and jq ( brew install jq / apt install jq ).
Bridge discovery uses dns-sd on macOS or avahi-browse on Linux
( apt install avahi-utils ). Without either it falls back to Hue's cloud
lookup, or you can type the IP during /hue:init .
Windows works through WSL or Git Bash, with the cloud lookup or a typed IP.
A square Hue Bridge (v2) on firmware with the CLIP v2 API (2021 or later).
Color bulbs that support signaling. hue lights shows which do.
By default, this plugin will flash green to indicate a finished turn, and flash amber and red to indicate that Claude is waiting for your feedback. These defaults can be changed to any color or duration. Create a config file at ~/.config/hue/events.conf (or tell Claude to) with your desired settings. Example:
claude-done on_off_color 2000 blue
claude-wait alternating 8000 red white
Signals are alternating (two colors), on_off_color (one), on_off (none),
and breathe (one smooth pulse in the current color, no duration or colors).
Claude Code plugin for controlling Philips Hue lights on status changes
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
