---
source: "https://notibuddy.app/blog/what-is-notibuddy"
hn_url: "https://news.ycombinator.com/item?id=50620802"
title: "What Is NotiBuddy? Approve AI Agents from Your iPhone with Face ID"
article_title: "What is NotiBuddy? Approve AI Agents from Your iPhone with Face ID"
image: "https://notibuddy.app/og-card.png"
author: "deegles"
captured_at: "2026-11-23T23:27:55Z"
capture_tool: "hn-digest"
hn_id: 50620802
score: 1
comments: 0
posted_at: "2026-11-23T23:09:43Z"
tags:
  - hacker-news
---

# What Is NotiBuddy? Approve AI Agents from Your iPhone with Face ID

- HN: [50620802](https://news.ycombinator.com/item?id=50620802)
- Source: [notibuddy.app](https://notibuddy.app/blog/what-is-notibuddy)
- Score: 1
- Comments: 0
- Posted: 2026-11-23T23:09:43Z

## Translation

Title: What Is NotiBuddy? Approve AI Agents from Your iPhone with Face ID
Article title: What is NotiBuddy? Approve AI Agents from Your iPhone with Face ID
Description: Approve AI coding agents, scripts, and CI deploys from your iPhone or Apple Watch with Face ID, code diffs, and native screens—end-to-end encrypted, no account.

Article text:
Skip to content NotiBuddy How it works Screens Setup guides by tool Plans Setup Blog Help English English Español (México) Deutsch Français (France) Français (Canada) Español (España) Download on the App Store ← All blog posts Overview Guide · November 2026 · notibuddy 0.4.0+ · 5 min read What is NotiBuddy? Approve AI Agents from Your iPhone with Face ID
Approve AI coding agents, scripts, and CI deploys from your iPhone or Apple Watch with Face ID, code diffs, and native screens—end-to-end encrypted, no account.
You start a big refactor in Claude Code, Cursor, Codex, or Antigravity and walk away to make coffee. Ten minutes later you come back to find it has been sitting there the whole time, waiting for you to type “yes.”
Whether you’re tinkering on a weekend side project, running a home-server script, or shipping a GitHub Actions deploy, every automated workflow hits the same wall eventually: at some point, it needs a human decision.
NotiBuddy sends that question to your phone. It is a native iPhone, iPad, and Apple Watch app paired with a local CLI and MCP server ( notibuddy@^0.4.0 ). Your agent or script asks, your phone buzzes, you tap Approve or Reject —or confirm with Face ID for risky actions—and your run keeps moving.
No account to sign up for. No chat bot or webhook server to host. And only your computer and your devices can read what passes between them.
1. Pair your phone (about a minute)
In NotiBuddy on your iPhone, tap Connect → Add a connection → Create setup code → Share Setup . Keep the app open, copy the one-time setup message to your computer (via Universal Clipboard or AirDrop), and run (Node.js 20+):
pbpaste | npx -y notibuddy@^0.4.0 pair --to personal
That’s it—your computer and phone now trust each other, with no login, password, or shared API key to manage. (Under the hood, the CLI creates a local key in ~/.notibuddy/ and swaps public keys with your phone. Want the short nb command on your PATH ? Install once via brew install yourbit-network/tap/notibuddy or npm i -g notibuddy@^0.4.0 .)
2. Connect your agent or script
Pick whatever fits how you work:
AI coding agents: add NotiBuddy as a local MCP server ( npx -y notibuddy@^0.4.0 mcp --to personal ), or install the official Claude Code plugin.
Unattended Claude Code prompts: nb hook claude PermissionRequest gives you 10 seconds to answer at your keyboard, then forwards unanswered prompts to your phone.
Shell scripts, cron jobs, and CI: call nb gate (or npx -y notibuddy@^0.4.0 gate ) to pause any command until you decide.
3. Answer from wherever you are
Your request arrives as an interactive notification on your iPhone, iPad, or paired Apple Watch—or opens as a full native screen on iPhone and iPad. Tap a button, fill in a quick form, or confirm with Face ID (or Touch ID). Your answer goes straight back to the waiting agent or script that asked.
No answer means no. If a request times out, never arrives, or isn’t confirmed with Face ID when required, NotiBuddy fails closed and treats it as a rejection. Your agent or script will never mistake silence for a “yes.”
What you can send to your phone
NotiBuddy goes well beyond a one-way push ping:
Face ID & Touch ID approvals: Require biometric confirmation before anything risky—production deploys, database migrations, or destructive shell commands ( requires_auth: true in MCP, --requires-auth in the CLI).
Inline code diffs: Read original vs. modified patches line by line on your phone before letting an agent edit files or apply SQL changes.
Custom Native Screens: Render real native SwiftUI screens from simple JSON—KPI strips, status alerts, checklists, forms, diffs, and gated action buttons. Start from 8 built-in templates (like Deployment Gate, Database Migration, Incident Triage, and Code Review) or mix and match 12 native SwiftUI blocks .
Live Activities, widgets & Apple Watch: Track long builds, test suites, or multi-step runs on your Lock Screen, Dynamic Island, StandBy widgets, and Apple Watch ( notibuddy_activity )—and progress updates don’t use extra free-tier requests.
Send files from iPhone to your agent: Share a UI bug screenshot, PDF, or log (up to 10 MB) from the iOS Share Sheet straight to an approved computer or agent ( nb files list / nb files receive or notibuddy_files ), saved locally without auto-opening untrusted files.
You don’t need to know anything about cryptography to use NotiBuddy, but here is what happens behind the scenes:
Only your devices can read it. Every question, diff, custom screen, file, and reply is end-to-end encrypted. NotiBuddy’s delivery service only passes sealed messages along—it cannot read your prompts, code, or answers.
Replies can’t be faked. Every answer is signed on your phone and verified locally by your computer against the phone you paired before anything runs.
No accounts. No email, no password, and no shared bot tokens sitting on a server.
How NotiBuddy compares to push apps and chat bots
Example: a deploy that waits for your Face ID
You don’t need an AI agent to use NotiBuddy. Any shell script or CI job can pause for your approval with one command:
npx -y notibuddy@^0.4.0 gate "Deploy v2.4.1 to production?" \
--to personal \
--message "Canary healthy · 128/128 checks passed" \
--approve "Promote" \
--reject "Abort" \
--requires-auth \
--timeout 600
Your phone shows the question with Promote and Abort buttons. The command then exits with:
Only 0 means “go,” so a missed notification can never ship your code.
NotiBuddy is free to start on the App Store (iOS 17+, iPadOS 17+, watchOS 10+). The free tier includes five requests per rolling 24 hours with every feature unlocked—Face ID approvals, Custom Native Screens, inline diffs, Live Activities, widgets, and Apple Watch—with optional Pro plans when you need higher volume or multiple devices.
Using an AI coding agent? Follow the step-by-step guide for Claude Code , Cursor , OpenAI Codex , or Google Antigravity .
Automating builds and deploys? Add approvals to GitHub Actions or shell scripts & webhooks .
Want custom screens? Copy-paste examples from the Screens & Templates gallery .
Want the full reference? Read the complete setup guide , or point your agent at /llms.txt .
Try NotiBuddy on your iPhone & Apple Watch
Free to start on the App Store (iOS 17+, watchOS 10+) with Face ID approvals, Custom Native Screens, inline code diffs, and local MCP/CLI support.
Agent-readable Markdown (includes a “For AI Agents” technical reference): notibuddy.app/blog/what-is-notibuddy.md · RSS feed
Your AI pauses. You decide from anywhere.
© 2026 NotiBuddy · A product by Yourbit

## Original Extract

Approve AI coding agents, scripts, and CI deploys from your iPhone or Apple Watch with Face ID, code diffs, and native screens—end-to-end encrypted, no account.

Skip to content NotiBuddy How it works Screens Setup guides by tool Plans Setup Blog Help English English Español (México) Deutsch Français (France) Français (Canada) Español (España) Download on the App Store ← All blog posts Overview Guide · November 2026 · notibuddy 0.4.0+ · 5 min read What is NotiBuddy? Approve AI Agents from Your iPhone with Face ID
Approve AI coding agents, scripts, and CI deploys from your iPhone or Apple Watch with Face ID, code diffs, and native screens—end-to-end encrypted, no account.
You start a big refactor in Claude Code, Cursor, Codex, or Antigravity and walk away to make coffee. Ten minutes later you come back to find it has been sitting there the whole time, waiting for you to type “yes.”
Whether you’re tinkering on a weekend side project, running a home-server script, or shipping a GitHub Actions deploy, every automated workflow hits the same wall eventually: at some point, it needs a human decision.
NotiBuddy sends that question to your phone. It is a native iPhone, iPad, and Apple Watch app paired with a local CLI and MCP server ( notibuddy@^0.4.0 ). Your agent or script asks, your phone buzzes, you tap Approve or Reject —or confirm with Face ID for risky actions—and your run keeps moving.
No account to sign up for. No chat bot or webhook server to host. And only your computer and your devices can read what passes between them.
1. Pair your phone (about a minute)
In NotiBuddy on your iPhone, tap Connect → Add a connection → Create setup code → Share Setup . Keep the app open, copy the one-time setup message to your computer (via Universal Clipboard or AirDrop), and run (Node.js 20+):
pbpaste | npx -y notibuddy@^0.4.0 pair --to personal
That’s it—your computer and phone now trust each other, with no login, password, or shared API key to manage. (Under the hood, the CLI creates a local key in ~/.notibuddy/ and swaps public keys with your phone. Want the short nb command on your PATH ? Install once via brew install yourbit-network/tap/notibuddy or npm i -g notibuddy@^0.4.0 .)
2. Connect your agent or script
Pick whatever fits how you work:
AI coding agents: add NotiBuddy as a local MCP server ( npx -y notibuddy@^0.4.0 mcp --to personal ), or install the official Claude Code plugin.
Unattended Claude Code prompts: nb hook claude PermissionRequest gives you 10 seconds to answer at your keyboard, then forwards unanswered prompts to your phone.
Shell scripts, cron jobs, and CI: call nb gate (or npx -y notibuddy@^0.4.0 gate ) to pause any command until you decide.
3. Answer from wherever you are
Your request arrives as an interactive notification on your iPhone, iPad, or paired Apple Watch—or opens as a full native screen on iPhone and iPad. Tap a button, fill in a quick form, or confirm with Face ID (or Touch ID). Your answer goes straight back to the waiting agent or script that asked.
No answer means no. If a request times out, never arrives, or isn’t confirmed with Face ID when required, NotiBuddy fails closed and treats it as a rejection. Your agent or script will never mistake silence for a “yes.”
What you can send to your phone
NotiBuddy goes well beyond a one-way push ping:
Face ID & Touch ID approvals: Require biometric confirmation before anything risky—production deploys, database migrations, or destructive shell commands ( requires_auth: true in MCP, --requires-auth in the CLI).
Inline code diffs: Read original vs. modified patches line by line on your phone before letting an agent edit files or apply SQL changes.
Custom Native Screens: Render real native SwiftUI screens from simple JSON—KPI strips, status alerts, checklists, forms, diffs, and gated action buttons. Start from 8 built-in templates (like Deployment Gate, Database Migration, Incident Triage, and Code Review) or mix and match 12 native SwiftUI blocks .
Live Activities, widgets & Apple Watch: Track long builds, test suites, or multi-step runs on your Lock Screen, Dynamic Island, StandBy widgets, and Apple Watch ( notibuddy_activity )—and progress updates don’t use extra free-tier requests.
Send files from iPhone to your agent: Share a UI bug screenshot, PDF, or log (up to 10 MB) from the iOS Share Sheet straight to an approved computer or agent ( nb files list / nb files receive or notibuddy_files ), saved locally without auto-opening untrusted files.
You don’t need to know anything about cryptography to use NotiBuddy, but here is what happens behind the scenes:
Only your devices can read it. Every question, diff, custom screen, file, and reply is end-to-end encrypted. NotiBuddy’s delivery service only passes sealed messages along—it cannot read your prompts, code, or answers.
Replies can’t be faked. Every answer is signed on your phone and verified locally by your computer against the phone you paired before anything runs.
No accounts. No email, no password, and no shared bot tokens sitting on a server.
How NotiBuddy compares to push apps and chat bots
Example: a deploy that waits for your Face ID
You don’t need an AI agent to use NotiBuddy. Any shell script or CI job can pause for your approval with one command:
npx -y notibuddy@^0.4.0 gate "Deploy v2.4.1 to production?" \
--to personal \
--message "Canary healthy · 128/128 checks passed" \
--approve "Promote" \
--reject "Abort" \
--requires-auth \
--timeout 600
Your phone shows the question with Promote and Abort buttons. The command then exits with:
Only 0 means “go,” so a missed notification can never ship your code.
NotiBuddy is free to start on the App Store (iOS 17+, iPadOS 17+, watchOS 10+). The free tier includes five requests per rolling 24 hours with every feature unlocked—Face ID approvals, Custom Native Screens, inline diffs, Live Activities, widgets, and Apple Watch—with optional Pro plans when you need higher volume or multiple devices.
Using an AI coding agent? Follow the step-by-step guide for Claude Code , Cursor , OpenAI Codex , or Google Antigravity .
Automating builds and deploys? Add approvals to GitHub Actions or shell scripts & webhooks .
Want custom screens? Copy-paste examples from the Screens & Templates gallery .
Want the full reference? Read the complete setup guide , or point your agent at /llms.txt .
Try NotiBuddy on your iPhone & Apple Watch
Free to start on the App Store (iOS 17+, watchOS 10+) with Face ID approvals, Custom Native Screens, inline code diffs, and local MCP/CLI support.
Agent-readable Markdown (includes a “For AI Agents” technical reference): notibuddy.app/blog/what-is-notibuddy.md · RSS feed
Your AI pauses. You decide from anywhere.
© 2026 NotiBuddy · A product by Yourbit
