---
source: "https://github.com/davertor/jev-slop-guard"
hn_url: "https://news.ycombinator.com/item?id=49801462"
title: "A Chrome extension that blurs AI slop on X and LinkedIn"
article_title: "GitHub - davertor/jev-slop-guard: Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll · GitHub"
image: "https://repository-images.githubusercontent.com/1379509422/6f7b569a-3099-49e1-b78e-22ce31b52519"
author: "davertor"
captured_at: "2026-09-22T14:23:31Z"
capture_tool: "hn-digest"
hn_id: 49801462
score: 1
comments: 0
posted_at: "2026-09-22T14:01:20Z"
tags:
  - hacker-news
---

# A Chrome extension that blurs AI slop on X and LinkedIn

- HN: [49801462](https://news.ycombinator.com/item?id=49801462)
- Source: [github.com](https://github.com/davertor/jev-slop-guard)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T14:01:20Z

## Translation

Title: A Chrome extension that blurs AI slop on X and LinkedIn
Article title: GitHub - davertor/jev-slop-guard: Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll · GitHub
Description: Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll - davertor/jev-slop-guard

Article text:
GitHub - davertor/jev-slop-guard: Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll · GitHub
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
davertor
/
jev-slop-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
76 Commits 76 Commits Folders and files
.github/ workflows .github/ workflows chrome-mv3 chrome-mv3 docs docs entrypoints entrypoints lib lib preview preview public public test test .cz.toml .cz.toml .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json wxt.config.ts wxt.config.ts View all files Repository files navigation
A Chrome extension that stands between you and the slop on X and LinkedIn.
Every post gets a small pill with a slop probability. Posts over your threshold
get blurred and stamped SLOP , with a button to show them anyway. The verdict
comes from Jev , TypeSafe's System One model: one typed
{ slop, not_slop } choice per post, no free-form LLM text, and you bring your
own API key.
What it does ·
Requirements ·
Try it in your Chrome ·
Settings ·
Privacy ·
References
Works on x.com and linkedin.com .
Every post gets a pill with its slop score: green Slop below your
threshold, red Stop at or above it.
At or above the threshold the post is also blurred and stamped SLOP .
Show the post clears that for the current session.
Scrolling is never blocked. Classification runs in the background with a
small concurrency limit and a per-post cache, so a post is scored once.
Promoted and sponsored posts and "Who to follow" widgets are skipped.
Chrome , or any browser that loads Manifest V3 extensions.
An API key for one of the two providers. Every post is classified by
Jev on that account, so usage is billed to you.
TypeSafe AI , the default provider.
OpenRouter . Allow TypeSafe under
OpenRouter Settings → Privacy, or the decisions endpoint refuses the call.
Clone or download this repo. The built extension is in chrome-mv3/ .
Turn on Developer mode (top-right toggle).
Click Load unpacked and pick the chrome-mv3/ folder.
Pin Jev Slop Guard in the toolbar, open the popup, paste your API key, click Save .
Open x.com or linkedin.com/feed
and scroll.
If you find this useful, consider dropping a ⭐.
Stored in chrome.storage.local only. Nothing leaves your browser except the
classification request described below.
For each post the extension sends the post text and the author handle to the
provider you picked, and nothing else. Your key and settings stay in
chrome.storage.local . The extension asks for access to x.com
and linkedin.com to read posts and draw badges, and to
api.typesafe.ai and openrouter.ai to send the classification requests.
Robin Bilgil's real-time slop detector demo ,
the idea this extension copies: a pill under the post, then blur and a SLOP stamp.
Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll - davertor/jev-slop-guard

GitHub - davertor/jev-slop-guard: Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll · GitHub
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
davertor
/
jev-slop-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
76 Commits 76 Commits Folders and files
.github/ workflows .github/ workflows chrome-mv3 chrome-mv3 docs docs entrypoints entrypoints lib lib preview preview public public test test .cz.toml .cz.toml .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json wxt.config.ts wxt.config.ts View all files Repository files navigation
A Chrome extension that stands between you and the slop on X and LinkedIn.
Every post gets a small pill with a slop probability. Posts over your threshold
get blurred and stamped SLOP , with a button to show them anyway. The verdict
comes from Jev , TypeSafe's System One model: one typed
{ slop, not_slop } choice per post, no free-form LLM text, and you bring your
own API key.
What it does ·
Requirements ·
Try it in your Chrome ·
Settings ·
Privacy ·
References
Works on x.com and linkedin.com .
Every post gets a pill with its slop score: green Slop below your
threshold, red Stop at or above it.
At or above the threshold the post is also blurred and stamped SLOP .
Show the post clears that for the current session.
Scrolling is never blocked. Classification runs in the background with a
small concurrency limit and a per-post cache, so a post is scored once.
Promoted and sponsored posts and "Who to follow" widgets are skipped.
Chrome , or any browser that loads Manifest V3 extensions.
An API key for one of the two providers. Every post is classified by
Jev on that account, so usage is billed to you.
TypeSafe AI , the default provider.
OpenRouter . Allow TypeSafe under
OpenRouter Settings → Privacy, or the decisions endpoint refuses the call.
Clone or download this repo. The built extension is in chrome-mv3/ .
Turn on Developer mode (top-right toggle).
Click Load unpacked and pick the chrome-mv3/ folder.
Pin Jev Slop Guard in the toolbar, open the popup, paste your API key, click Save .
Open x.com or linkedin.com/feed
and scroll.
If you find this useful, consider dropping a ⭐.
Stored in chrome.storage.local only. Nothing leaves your browser except the
classification request described below.
For each post the extension sends the post text and the author handle to the
provider you picked, and nothing else. Your key and settings stay in
chrome.storage.local . The extension asks for access to x.com
and linkedin.com to read posts and draw badges, and to
api.typesafe.ai and openrouter.ai to send the classification requests.
Robin Bilgil's real-time slop detector demo ,
the idea this extension copies: a pill under the post, then blur and a SLOP stamp.
Jev Slop Guard — a Chrome extension that scores and stamps AI slop on your X and LinkedIn feeds as you scroll
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
