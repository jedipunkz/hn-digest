---
source: "https://github.com/frizurd/postquel"
hn_url: "https://news.ycombinator.com/item?id=49968130"
title: "Show HN: Postquel, a Free Mac Postgres Client with Claude Code and Codex"
article_title: "GitHub - frizurd/postquel: Native macOS PostgreSQL client with bring-your-own AI (work in progress) · GitHub"
image: "https://opengraph.githubassets.com/2e26ed0b4819dc53179f19eec33f4ec73ce36f5a6890ae45b4bc551230a0e64e/frizurd/postquel"
author: "frizky"
captured_at: "2026-10-05T17:59:11Z"
capture_tool: "hn-digest"
hn_id: 49968130
score: 1
comments: 0
posted_at: "2026-10-05T17:55:06Z"
tags:
  - hacker-news
---

# Show HN: Postquel, a Free Mac Postgres Client with Claude Code and Codex

- HN: [49968130](https://news.ycombinator.com/item?id=49968130)
- Source: [github.com](https://github.com/frizurd/postquel)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T17:55:06Z

## Translation

Title: Show HN: Postquel, a Free Mac Postgres Client with Claude Code and Codex
Article title: GitHub - frizurd/postquel: Native macOS PostgreSQL client with bring-your-own AI (work in progress) · GitHub
Description: Native macOS PostgreSQL client with bring-your-own AI (work in progress) - frizurd/postquel

Article text:
GitHub - frizurd/postquel: Native macOS PostgreSQL client with bring-your-own AI (work in progress) · GitHub
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
frizurd
/
postquel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
43 Commits 43 Commits Folders and files
Resources Resources Sources Sources docs docs scripts scripts .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
A fast, native PostgreSQL client for macOS, with an AI assistant that uses the Claude Code, Codex or Cursor account already on your Mac.
Tables — browse and edit rows, follow foreign keys, and switch between Content, Structure and DDL. Structure edits columns, indexes, constraints and notes, then saves them in one transaction after an SQL preview.
Queries — a SQL editor with highlighting and line numbers, results you can sort by clicking a header, and query timing.
Assistant — ask about your data in plain language, or describe a query and get SQL. It reads the real schema and runs read-only queries; changes are only ever proposed, for you to dry-run and apply.
Bring your own AI — uses Claude Code, Codex or Cursor Agent, whichever is installed and signed in. Pick any of their models. No API keys are stored in Postquel.
Native — SwiftUI and AppKit on libpq, with Liquid Glass on macOS 26. SSH tunnels and Keychain passwords included.
Postquel is a prototype and a work in progress.
Download Postquel for Mac — macOS 15 or later, Apple silicon and Intel. Signed and notarized by Apple.
Open the DMG and drag Postquel to Applications. For the assistant, install and sign in to Claude Code , Codex or Cursor Agent . Release notes are on the releases page .
Requires Xcode / Swift 6 and libpq (defaults to Postgres.app; set LIBPQ_PREFIX for Homebrew libpq ).
./scripts/build-app.sh --install # build, install to /Applications, relaunch if running
./scripts/dev.sh # watch mode: rebuild + reinstall + relaunch on every change
./scripts/build-dmg.sh # universal app in a DMG, for distribution
swift run # run without bundling
The app bundles libpq and the OpenSSL libraries it uses (from Postgres.app, or LIBPQ_PREFIX ), so it runs on Macs without Postgres installed. To sign and notarize a DMG, set DEVELOPER_ID to a "Developer ID Application" identity and NOTARY_PROFILE to a profile saved with xcrun notarytool store-credentials ; see scripts/build-dmg.sh .
Demo data: createdb postquel_demo && psql postquel_demo -f scripts/seed-demo.sql
Resources/AppIcon.png – silver icon on dark graphite; the app build generates all macOS icon sizes
Resources/Logo.png – silver mark with a transparent background
Resources/Licenses – notices for the bundled libpq and OpenSSL, copied into the app
Sources/CLibPQ – module map for libpq
Sources/Postquel/Database – PGConnection (libpq on a serial queue, text results, cancel)
Sources/Postquel/Models – session, table browser, query editor state
Sources/Postquel/Views – ResultsGrid (NSTableView), SQLEditor (NSTextView), SwiftUI shell
Native macOS PostgreSQL client with bring-your-own AI (work in progress)
frizky.dev/products/postquel Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Native macOS PostgreSQL client with bring-your-own AI (work in progress) - frizurd/postquel

GitHub - frizurd/postquel: Native macOS PostgreSQL client with bring-your-own AI (work in progress) · GitHub
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
frizurd
/
postquel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
43 Commits 43 Commits Folders and files
Resources Resources Sources Sources docs docs scripts scripts .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
A fast, native PostgreSQL client for macOS, with an AI assistant that uses the Claude Code, Codex or Cursor account already on your Mac.
Tables — browse and edit rows, follow foreign keys, and switch between Content, Structure and DDL. Structure edits columns, indexes, constraints and notes, then saves them in one transaction after an SQL preview.
Queries — a SQL editor with highlighting and line numbers, results you can sort by clicking a header, and query timing.
Assistant — ask about your data in plain language, or describe a query and get SQL. It reads the real schema and runs read-only queries; changes are only ever proposed, for you to dry-run and apply.
Bring your own AI — uses Claude Code, Codex or Cursor Agent, whichever is installed and signed in. Pick any of their models. No API keys are stored in Postquel.
Native — SwiftUI and AppKit on libpq, with Liquid Glass on macOS 26. SSH tunnels and Keychain passwords included.
Postquel is a prototype and a work in progress.
Download Postquel for Mac — macOS 15 or later, Apple silicon and Intel. Signed and notarized by Apple.
Open the DMG and drag Postquel to Applications. For the assistant, install and sign in to Claude Code , Codex or Cursor Agent . Release notes are on the releases page .
Requires Xcode / Swift 6 and libpq (defaults to Postgres.app; set LIBPQ_PREFIX for Homebrew libpq ).
./scripts/build-app.sh --install # build, install to /Applications, relaunch if running
./scripts/dev.sh # watch mode: rebuild + reinstall + relaunch on every change
./scripts/build-dmg.sh # universal app in a DMG, for distribution
swift run # run without bundling
The app bundles libpq and the OpenSSL libraries it uses (from Postgres.app, or LIBPQ_PREFIX ), so it runs on Macs without Postgres installed. To sign and notarize a DMG, set DEVELOPER_ID to a "Developer ID Application" identity and NOTARY_PROFILE to a profile saved with xcrun notarytool store-credentials ; see scripts/build-dmg.sh .
Demo data: createdb postquel_demo && psql postquel_demo -f scripts/seed-demo.sql
Resources/AppIcon.png – silver icon on dark graphite; the app build generates all macOS icon sizes
Resources/Logo.png – silver mark with a transparent background
Resources/Licenses – notices for the bundled libpq and OpenSSL, copied into the app
Sources/CLibPQ – module map for libpq
Sources/Postquel/Database – PGConnection (libpq on a serial queue, text results, cancel)
Sources/Postquel/Models – session, table browser, query editor state
Sources/Postquel/Views – ResultsGrid (NSTableView), SQLEditor (NSTextView), SwiftUI shell
Native macOS PostgreSQL client with bring-your-own AI (work in progress)
frizky.dev/products/postquel Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
