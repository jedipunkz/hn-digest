---
source: "https://github.com/adeptads/agent-skills"
hn_url: "https://news.ycombinator.com/item?id=50460287"
title: "Show HN: Agent skill and Claude/Codex plugins for Google Ads campaigns"
article_title: "GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. · GitHub"
image: "https://opengraph.githubassets.com/388cb416247b4ef08a71083b080afd0e5647a1abdb006bb6d63da62e9e2c336f/adeptads/agent-skills"
author: "adeptads_ai"
captured_at: "2026-11-11T23:20:51Z"
capture_tool: "hn-digest"
hn_id: 50460287
score: 1
comments: 0
posted_at: "2026-11-11T23:11:38Z"
tags:
  - hacker-news
---

# Show HN: Agent skill and Claude/Codex plugins for Google Ads campaigns

- HN: [50460287](https://news.ycombinator.com/item?id=50460287)
- Source: [github.com](https://github.com/adeptads/agent-skills)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T23:11:38Z

## Translation

Title: Show HN: Agent skill and Claude/Codex plugins for Google Ads campaigns
Article title: GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. · GitHub
Description: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. - GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.

Article text:
GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. · GitHub
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
adeptads
agent-skills Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.agents/ plugins .agents/ plugins .claude-plugin .claude-plugin assets assets plugins plugins scripts scripts skills/ adeptads-google-ads skills/ adeptads-google-ads templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Tell your AI agent what you sell and who should see it. AdeptAds researches your offer, plans the ad groups, writes the ads and creates the images. You look it over and decide when it goes live.
🚀 Launch
🌱 Grow
📊 Check
You: New Search campaign for acme-plumbing.example. Leads, Austin, $40 a day.
Agent: On it. About 10 to 20 minutes.
Agent: Done and paused: 4 ad groups, 38 keywords, 12 ads. Start it?
You: Add some keywords. Fifteen at most.
Agent: Added 11. Four more didn't fit your offer, so I left them out.
You: How's it going?
Agent: Last 30 days: 412 clicks, $508 spent, 19 conversions, $26.74 each. The latest days look low because conversions arrive late.
Example sessions. Your numbers will be your own.
Campaigns arrive paused. Nothing spends until you say start.
claude plugin marketplace add adeptads/agent-skills
claude plugin install adeptads-google-ads@adeptads
Start a new session and use /mcp to complete AdeptAds OAuth sign-in.
codex plugin marketplace add adeptads/agent-skills
codex plugin add adeptads-google-ads@adeptads
Start a new session and complete AdeptAds OAuth sign-in when prompted.
npx skills add adeptads/agent-skills --skill adeptads-google-ads
This installs instructions only. Add the AdeptAds remote MCP connection under the server name adeptads , for example in Claude Code:
claude mcp add --transport http adeptads https://api.adeptads.ai/mcp
Not sure where to start? Try: “Use AdeptAds to list my Google Ads accounts.”
Build
Grow
Control
A complete Search, Display or Performance Max campaign from your website and a short brief.
New ad groups, keywords, search themes and locations for campaigns you already run.
Start, pause, change budget or bidding, and read results in conversation.
Good to know
Can the agent spend money without asking?
It can't start a campaign, raise a budget or change bidding without your OK. New ad groups and keywords on a running campaign go live at once; the agent asks first when the idea was its own. New locations wait for your approval in the AdeptAds dashboard.
Yes. The agent ends with a list of every change. AdeptAds also emails the account owner about agent-triggered starts, bidding strategy changes, removed limits, large spend-related increases, and changes it could not verify. Alerts do not replace your approval.
Installation is free. Running workflows needs an AdeptAds account, connected Google Ads access, and credits from your trial or paid plan. Expansion uses optimization credits. Google bills ad spend separately.
Autopilot setup, keyword-level bids, negatives, new images for existing ads, and network or location settings. Use the AdeptAds dashboard or Google Ads for those.
While the conversation is open, yes: the agent checks progress and reports each step. It does not send updates into a closed conversation. A queued task is not a finished campaign, so ask the agent to check again.
Package
Location
Contents
Standalone skill
skills/adeptads-google-ads/
Instructions and tool reference; connect MCP separately
Codex plugin
plugins/codex/adeptads-google-ads/
The skill and a hosted MCP connection
Claude plugin
plugins/claude/adeptads-google-ads/
The skill and a hosted MCP connection
The MCP server and campaign workflows run in the AdeptAds backend. This repository contains client packages, not the server implementation or credentials.
Maintain the skill in skills/ , metadata in package.json , and the plugin README in templates/plugin-README.md . Verify PPC advice against its sources when changing it. Generate both plugin bundles and ZIP archives with:
python3 scripts/build.py
python3 scripts/validate.py
Generated skill copies must not be edited independently. ZIPs and their SHA-256 inventory are written to dist/ ; the builder does not upload, submit, or deploy anything.
Bump the version for every published change so installed clients receive the update. The builder removes its own older versioned ZIPs before rebuilding.
MIT covers the skills, plugin configuration, scripts and documentation in this repository. The hosted AdeptAds service and backend workflows are outside this repository and remain subject to their own terms and billing.
🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. - GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.

GitHub - adeptads/agent-skills: 🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows. · GitHub
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
adeptads
agent-skills Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.agents/ plugins .agents/ plugins .claude-plugin .claude-plugin assets assets plugins plugins scripts scripts skills/ adeptads-google-ads skills/ adeptads-google-ads templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Tell your AI agent what you sell and who should see it. AdeptAds researches your offer, plans the ad groups, writes the ads and creates the images. You look it over and decide when it goes live.
🚀 Launch
🌱 Grow
📊 Check
You: New Search campaign for acme-plumbing.example. Leads, Austin, $40 a day.
Agent: On it. About 10 to 20 minutes.
Agent: Done and paused: 4 ad groups, 38 keywords, 12 ads. Start it?
You: Add some keywords. Fifteen at most.
Agent: Added 11. Four more didn't fit your offer, so I left them out.
You: How's it going?
Agent: Last 30 days: 412 clicks, $508 spent, 19 conversions, $26.74 each. The latest days look low because conversions arrive late.
Example sessions. Your numbers will be your own.
Campaigns arrive paused. Nothing spends until you say start.
claude plugin marketplace add adeptads/agent-skills
claude plugin install adeptads-google-ads@adeptads
Start a new session and use /mcp to complete AdeptAds OAuth sign-in.
codex plugin marketplace add adeptads/agent-skills
codex plugin add adeptads-google-ads@adeptads
Start a new session and complete AdeptAds OAuth sign-in when prompted.
npx skills add adeptads/agent-skills --skill adeptads-google-ads
This installs instructions only. Add the AdeptAds remote MCP connection under the server name adeptads , for example in Claude Code:
claude mcp add --transport http adeptads https://api.adeptads.ai/mcp
Not sure where to start? Try: “Use AdeptAds to list my Google Ads accounts.”
Build
Grow
Control
A complete Search, Display or Performance Max campaign from your website and a short brief.
New ad groups, keywords, search themes and locations for campaigns you already run.
Start, pause, change budget or bidding, and read results in conversation.
Good to know
Can the agent spend money without asking?
It can't start a campaign, raise a budget or change bidding without your OK. New ad groups and keywords on a running campaign go live at once; the agent asks first when the idea was its own. New locations wait for your approval in the AdeptAds dashboard.
Yes. The agent ends with a list of every change. AdeptAds also emails the account owner about agent-triggered starts, bidding strategy changes, removed limits, large spend-related increases, and changes it could not verify. Alerts do not replace your approval.
Installation is free. Running workflows needs an AdeptAds account, connected Google Ads access, and credits from your trial or paid plan. Expansion uses optimization credits. Google bills ad spend separately.
Autopilot setup, keyword-level bids, negatives, new images for existing ads, and network or location settings. Use the AdeptAds dashboard or Google Ads for those.
While the conversation is open, yes: the agent checks progress and reports each step. It does not send updates into a closed conversation. A queued task is not a finished campaign, so ask the agent to check again.
Package
Location
Contents
Standalone skill
skills/adeptads-google-ads/
Instructions and tool reference; connect MCP separately
Codex plugin
plugins/codex/adeptads-google-ads/
The skill and a hosted MCP connection
Claude plugin
plugins/claude/adeptads-google-ads/
The skill and a hosted MCP connection
The MCP server and campaign workflows run in the AdeptAds backend. This repository contains client packages, not the server implementation or credentials.
Maintain the skill in skills/ , metadata in package.json , and the plugin README in templates/plugin-README.md . Verify PPC advice against its sources when changing it. Generate both plugin bundles and ZIP archives with:
python3 scripts/build.py
python3 scripts/validate.py
Generated skill copies must not be edited independently. ZIPs and their SHA-256 inventory are written to dist/ ; the builder does not upload, submit, or deploy anything.
Bump the version for every published change so installed clients receive the update. The builder removes its own older versioned ZIPs before rebuilding.
MIT covers the skills, plugin configuration, scripts and documentation in this repository. The hosted AdeptAds service and backend workflows are outside this repository and remain subject to their own terms and billing.
🚀 Give your AI agent Google Ads superpowers. Skills and plugins to create and expand campaigns with AdeptAds multi-agent workflows.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
