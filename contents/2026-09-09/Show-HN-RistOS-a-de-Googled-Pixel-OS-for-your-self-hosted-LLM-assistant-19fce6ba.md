---
source: "https://github.com/rist-os/rist-os-handset"
hn_url: "https://news.ycombinator.com/item?id=49625206"
title: "Show HN: RistOS – a de-Googled Pixel OS for your self-hosted LLM assistant"
article_title: "GitHub - rist-os/rist-os-handset: A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. · GitHub"
image: "https://opengraph.githubassets.com/7f7686ee5a1b2ffa1499b4b68a5c9496d9fcc90fa9cb4000175858a6cd627bca/rist-os/rist-os-handset"
author: "a_dugan"
captured_at: "2026-09-09T12:46:51Z"
capture_tool: "hn-digest"
hn_id: 49625206
score: 3
comments: 1
posted_at: "2026-09-09T12:11:18Z"
tags:
  - hacker-news
---

# Show HN: RistOS – a de-Googled Pixel OS for your self-hosted LLM assistant

- HN: [49625206](https://news.ycombinator.com/item?id=49625206)
- Source: [github.com](https://github.com/rist-os/rist-os-handset)
- Score: 3
- Comments: 1
- Posted: 2026-09-09T12:11:18Z

## Translation

Title: Show HN: RistOS – a de-Googled Pixel OS for your self-hosted LLM assistant
Article title: GitHub - rist-os/rist-os-handset: A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. · GitHub
Description: A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. - rist-os/rist-os-handset

Article text:
GitHub - rist-os/rist-os-handset: A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. · GitHub
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
rist-os
/
rist-os-handset
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github .github aosp aosp app app docs docs gradle/ wrapper gradle/ wrapper image image ota-server ota-server releases releases tools tools .gitignore .gitignore .publish-denylist.example .publish-denylist.example .publish-never-ship.example .publish-never-ship.example BUILD.md BUILD.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SAFETY.md SAFETY.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat push.sh push.sh settings.gradle.kts settings.gradle.kts View all files Repository files navigation
RistOS: LLM assistant-first phone firmware for the Google Pixel 10a
⬇ Install RistOS on your Pixel 10a
A de-Googled phone OS built on GrapheneOS for Google Pixel 10a. It is a minimalist, assistant-first OS
designed to work with your own backend. It has a small fixed set of local apps including the phone, messages, camera,
gallery, offline maps, flashlight and settings, and nothing else. No app store, no browser, no feed,
no Google account, and no way to add anything.
Rist sells no phones and no service. If you already own a Pixel 10a, you flash this yourself.
[you] --push-to-talk / text--> [Rist app] --protobuf over HTTPS--> [your backend]
^ |
+---- speech / views / commands <---+
The app records speech only while you hold the button, and POSTs it (or typed text)
to a backend endpoint as protobuf ( docs/BACKEND.md ).
The backend replies with speech, text, views, or device commands.
Bring your own backend. A public build has no endpoint compiled in and nothing to phone
home to, so you set one on the device. See docs/BACKEND.md .
Doc
What it answers
image/INSTALL.md
installing it on a phone, step by step
SAFETY.md
what can go wrong on a phone someone depends on — read first
docs/BACKEND.md
building a backend — start here, with a working example
docs/OTA.md
how updates reach a device
CHANGELOG.md
what each build does and does not do
SECURITY.md
reporting a vulnerability; known and accepted weaknesses
BUILD.md
building the app, or the whole image, yourself
TRADEMARKS.md
the code is yours to sell; the name is not — what a fork may call itself
CONTRIBUTING.md
what PRs are accepted, the pre-PR checklist, where bugs and questions go
Status and contributing
RistOS is pre-release and experimental. It supports the Google Pixel 10a only. There is no support
channel; bugs and questions go through the issue tracker. See CONTRIBUTING.md for
what is accepted and the pre-PR checklist.
Apache-2.0 (see LICENSE , third-party attributions in NOTICE ). The platform this builds on carries
its own licences, GPLv2 among them.
The download redistributes Google's proprietary Pixel firmware, unmodified — the bootloader and
the radio/modem images, plus the proprietary vendor files listed with their hashes in
image/proprietary-files.txt . Those remain Google's property and are governed by Google's terms,
not by this project's licence. Google, Pixel and Android are trademarks of Google LLC. RistOS is not
affiliated with, sponsored by, or endorsed by Google.
A modified build must not call itself RistOS; see TRADEMARKS.md .
A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. - rist-os/rist-os-handset

GitHub - rist-os/rist-os-handset: A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend. · GitHub
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
rist-os
/
rist-os-handset
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github .github aosp aosp app app docs docs gradle/ wrapper gradle/ wrapper image image ota-server ota-server releases releases tools tools .gitignore .gitignore .publish-denylist.example .publish-denylist.example .publish-never-ship.example .publish-never-ship.example BUILD.md BUILD.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SAFETY.md SAFETY.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat push.sh push.sh settings.gradle.kts settings.gradle.kts View all files Repository files navigation
RistOS: LLM assistant-first phone firmware for the Google Pixel 10a
⬇ Install RistOS on your Pixel 10a
A de-Googled phone OS built on GrapheneOS for Google Pixel 10a. It is a minimalist, assistant-first OS
designed to work with your own backend. It has a small fixed set of local apps including the phone, messages, camera,
gallery, offline maps, flashlight and settings, and nothing else. No app store, no browser, no feed,
no Google account, and no way to add anything.
Rist sells no phones and no service. If you already own a Pixel 10a, you flash this yourself.
[you] --push-to-talk / text--> [Rist app] --protobuf over HTTPS--> [your backend]
^ |
+---- speech / views / commands <---+
The app records speech only while you hold the button, and POSTs it (or typed text)
to a backend endpoint as protobuf ( docs/BACKEND.md ).
The backend replies with speech, text, views, or device commands.
Bring your own backend. A public build has no endpoint compiled in and nothing to phone
home to, so you set one on the device. See docs/BACKEND.md .
Doc
What it answers
image/INSTALL.md
installing it on a phone, step by step
SAFETY.md
what can go wrong on a phone someone depends on — read first
docs/BACKEND.md
building a backend — start here, with a working example
docs/OTA.md
how updates reach a device
CHANGELOG.md
what each build does and does not do
SECURITY.md
reporting a vulnerability; known and accepted weaknesses
BUILD.md
building the app, or the whole image, yourself
TRADEMARKS.md
the code is yours to sell; the name is not — what a fork may call itself
CONTRIBUTING.md
what PRs are accepted, the pre-PR checklist, where bugs and questions go
Status and contributing
RistOS is pre-release and experimental. It supports the Google Pixel 10a only. There is no support
channel; bugs and questions go through the issue tracker. See CONTRIBUTING.md for
what is accepted and the pre-PR checklist.
Apache-2.0 (see LICENSE , third-party attributions in NOTICE ). The platform this builds on carries
its own licences, GPLv2 among them.
The download redistributes Google's proprietary Pixel firmware, unmodified — the bootloader and
the radio/modem images, plus the proprietary vendor files listed with their hashes in
image/proprietary-files.txt . Those remain Google's property and are governed by Google's terms,
not by this project's licence. Google, Pixel and Android are trademarks of Google LLC. RistOS is not
affiliated with, sponsored by, or endorsed by Google.
A modified build must not call itself RistOS; see TRADEMARKS.md .
A de-Googled Android OS for the Pixel 10a. Push-to-talk assistant, your own backend.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
