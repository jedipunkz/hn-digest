---
source: "https://github.com/shashtag-ventures/linkedin-ops"
hn_url: "https://news.ycombinator.com/item?id=50322021"
title: "Show HN: LinkedIn-ops – Zero-dependency CDP automation built for AI agents"
article_title: "GitHub - shashtag-ventures/linkedin-ops · GitHub"
image: "https://opengraph.githubassets.com/e3dee70c792f3edb6d6763c2315dff2909d14413f6f460531cdb2791d8526d3a/shashtag-ventures/linkedin-ops"
author: "shashtag"
captured_at: "2026-11-01T01:54:17Z"
capture_tool: "hn-digest"
hn_id: 50322021
score: 1
comments: 0
posted_at: "2026-11-01T01:46:21Z"
tags:
  - hacker-news
---

# Show HN: LinkedIn-ops – Zero-dependency CDP automation built for AI agents

- HN: [50322021](https://news.ycombinator.com/item?id=50322021)
- Source: [github.com](https://github.com/shashtag-ventures/linkedin-ops)
- Score: 1
- Comments: 0
- Posted: 2026-11-01T01:46:21Z

## Translation

Title: Show HN: LinkedIn-ops – Zero-dependency CDP automation built for AI agents
Article title: GitHub - shashtag-ventures/linkedin-ops · GitHub
Description: Contribute to shashtag-ventures/linkedin-ops development by creating an account on GitHub.

Article text:
GitHub - shashtag-ventures/linkedin-ops · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
shashtag-ventures
linkedin-ops Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme License Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.github/ workflows .github/ workflows campaigns campaigns data data docs docs scratch scratch src src test test .gitignore .gitignore AGENTS.md AGENTS.md DATA_CONTRACT.md DATA_CONTRACT.md LICENSE LICENSE README.md README.md index.mjs index.mjs package.json package.json View all files Repository files navigation
High-reliability, multi-agent-safe LinkedIn automation engine and CLI built on Chrome DevTools Protocol (CDP).
Disclaimer & Responsible Use : linkedin-ops is an independent open-source tool built for developer automation research, testing, and workflow automation via local Chrome DevTools Protocol (CDP). This project is not affiliated with, endorsed by, or sponsored by LinkedIn Corporation. Automated interaction with LinkedIn must comply with applicable platform terms and policies. Always maintain realistic human delays and respect LinkedIn rate limits to safeguard account health.
npm install
2. Run automated setup
node index.mjs setup
The setup wizard automatically:
Locates your Google Chrome binary (macOS, Linux, Windows).
Discovers your existing Chrome profiles from system User Data so you can pick your logged-in account (or create a dedicated project session).
Saves settings to .linkedin-opsrc (no manual path pasting required).
Generates a one-click launcher ( start-chrome.command / start-chrome.bat ).
3. Verify session & connection
npm run doctor
# or: node index.mjs doctor
Checks CDP port 9222 reachability, browser version, user agent, and verifies an authenticated LinkedIn session.
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post on distributed systems... "
(Dry-run mode navigates and fills the note without submitting. Add --send only when ready to live submit).
Automated Runtime Fallback : Daily commands ping port 9222. If Chrome is closed, linkedin-ops automatically launches it using your configured profile. If Chrome is already running without remote debugging enabled, it detects the profile lock and prompts to restart cleanly.
Cross-Process Mutex Locking : Serializes execution across concurrent agents and processes ( linkedin.lock ). Prevents session collisions, focus-stealing, and rate-limiting.
Isolated Ephemeral Tabs : Each operation spawns its own dedicated Chrome tab via PUT /json/new and connects directly via target ID.
Auto-Cleanup : Guarantees tab destruction ( /json/close/<targetId> ) in a finally block on success or failure, eliminating zombie tab leaks.
# Automated setup wizard
node index.mjs setup
# Pre-flight health check
node index.mjs doctor
# Connection invitation with personalized note (dry-run)
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post... "
# Live connection invitation
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post... " --send
# Direct message to 1st-degree connection
node index.mjs dm --profile= " https://www.linkedin.com/in/username " --message= " Checking in on our conversation. " --send
# Comment on a LinkedIn post
node index.mjs comment --url= " https://www.linkedin.com/posts/... " --text= " Great insights on distributed systems! " --send
# Company people search (e.g. recruiters)
node index.mjs search --company= " Google " --keywords= " recruiter " --limit=10 --json
# Inspect or force-release active mutex lock
node index.mjs lock
node index.mjs lock release
Programmatic Library Usage
import { withLock , withDedicatedTab , sendConnectionInvite } from 'linkedin-ops' ;
await withLock ( 'outreach-task' , async ( ) => {
await withDedicatedTab ( { port : 9222 } , async ( cdp ) => {
const result = await sendConnectionInvite ( cdp , {
linkedinUrl : 'https://www.linkedin.com/in/username' ,
note : 'Hi, loved your work on distributed systems.' ,
send : true
} ) ;
console . log ( result ) ;
} ) ;
} ) ;
Batch Campaigns
For multi-profile outreach and candidate sweeps, see campaigns/batch-connect.example.mjs :
# Copy template to your personal user campaigns folder
cp campaigns/batch-connect.example.mjs campaigns/my-outreach.mjs
# Run the campaign safely with mutex locking and isolated tab lifecycles
node campaigns/my-outreach.mjs
(User campaigns and harvested datasets in data/ and campaigns/ are gitignored and never committed to source control).
DATA_CONTRACT.md : Strict definition of the System Layer vs User Layer boundaries.
Architecture Blueprint : Deep dive into CDP socket lifecycle, locking mechanisms, and tab management.
MIT © 2026 Karada.ai . See LICENSE for details.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to shashtag-ventures/linkedin-ops development by creating an account on GitHub.

GitHub - shashtag-ventures/linkedin-ops · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
shashtag-ventures
linkedin-ops Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme License Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.github/ workflows .github/ workflows campaigns campaigns data data docs docs scratch scratch src src test test .gitignore .gitignore AGENTS.md AGENTS.md DATA_CONTRACT.md DATA_CONTRACT.md LICENSE LICENSE README.md README.md index.mjs index.mjs package.json package.json View all files Repository files navigation
High-reliability, multi-agent-safe LinkedIn automation engine and CLI built on Chrome DevTools Protocol (CDP).
Disclaimer & Responsible Use : linkedin-ops is an independent open-source tool built for developer automation research, testing, and workflow automation via local Chrome DevTools Protocol (CDP). This project is not affiliated with, endorsed by, or sponsored by LinkedIn Corporation. Automated interaction with LinkedIn must comply with applicable platform terms and policies. Always maintain realistic human delays and respect LinkedIn rate limits to safeguard account health.
npm install
2. Run automated setup
node index.mjs setup
The setup wizard automatically:
Locates your Google Chrome binary (macOS, Linux, Windows).
Discovers your existing Chrome profiles from system User Data so you can pick your logged-in account (or create a dedicated project session).
Saves settings to .linkedin-opsrc (no manual path pasting required).
Generates a one-click launcher ( start-chrome.command / start-chrome.bat ).
3. Verify session & connection
npm run doctor
# or: node index.mjs doctor
Checks CDP port 9222 reachability, browser version, user agent, and verifies an authenticated LinkedIn session.
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post on distributed systems... "
(Dry-run mode navigates and fills the note without submitting. Add --send only when ready to live submit).
Automated Runtime Fallback : Daily commands ping port 9222. If Chrome is closed, linkedin-ops automatically launches it using your configured profile. If Chrome is already running without remote debugging enabled, it detects the profile lock and prompts to restart cleanly.
Cross-Process Mutex Locking : Serializes execution across concurrent agents and processes ( linkedin.lock ). Prevents session collisions, focus-stealing, and rate-limiting.
Isolated Ephemeral Tabs : Each operation spawns its own dedicated Chrome tab via PUT /json/new and connects directly via target ID.
Auto-Cleanup : Guarantees tab destruction ( /json/close/<targetId> ) in a finally block on success or failure, eliminating zombie tab leaks.
# Automated setup wizard
node index.mjs setup
# Pre-flight health check
node index.mjs doctor
# Connection invitation with personalized note (dry-run)
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post... "
# Live connection invitation
node index.mjs connect --profile= " https://www.linkedin.com/in/username " --note= " Hi, loved your post... " --send
# Direct message to 1st-degree connection
node index.mjs dm --profile= " https://www.linkedin.com/in/username " --message= " Checking in on our conversation. " --send
# Comment on a LinkedIn post
node index.mjs comment --url= " https://www.linkedin.com/posts/... " --text= " Great insights on distributed systems! " --send
# Company people search (e.g. recruiters)
node index.mjs search --company= " Google " --keywords= " recruiter " --limit=10 --json
# Inspect or force-release active mutex lock
node index.mjs lock
node index.mjs lock release
Programmatic Library Usage
import { withLock , withDedicatedTab , sendConnectionInvite } from 'linkedin-ops' ;
await withLock ( 'outreach-task' , async ( ) => {
await withDedicatedTab ( { port : 9222 } , async ( cdp ) => {
const result = await sendConnectionInvite ( cdp , {
linkedinUrl : 'https://www.linkedin.com/in/username' ,
note : 'Hi, loved your work on distributed systems.' ,
send : true
} ) ;
console . log ( result ) ;
} ) ;
} ) ;
Batch Campaigns
For multi-profile outreach and candidate sweeps, see campaigns/batch-connect.example.mjs :
# Copy template to your personal user campaigns folder
cp campaigns/batch-connect.example.mjs campaigns/my-outreach.mjs
# Run the campaign safely with mutex locking and isolated tab lifecycles
node campaigns/my-outreach.mjs
(User campaigns and harvested datasets in data/ and campaigns/ are gitignored and never committed to source control).
DATA_CONTRACT.md : Strict definition of the System Layer vs User Layer boundaries.
Architecture Blueprint : Deep dive into CDP socket lifecycle, locking mechanisms, and tab management.
MIT © 2026 Karada.ai . See LICENSE for details.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
