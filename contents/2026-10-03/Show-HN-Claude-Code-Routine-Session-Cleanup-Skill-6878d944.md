---
source: "https://github.com/naw103/claude-routine-cleanup"
hn_url: "https://news.ycombinator.com/item?id=49946507"
title: "Show HN: Claude Code Routine Session Cleanup Skill"
article_title: "GitHub - naw103/claude-routine-cleanup: Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. · GitHub"
image: "https://opengraph.githubassets.com/d57697c9737f4f5c7dddaa6bcdb8580d698c1f9525d0bb203aee63ac0df7ef67/naw103/claude-routine-cleanup"
author: "naw103"
captured_at: "2026-10-03T19:02:08Z"
capture_tool: "hn-digest"
hn_id: 49946507
score: 2
comments: 1
posted_at: "2026-10-03T18:20:21Z"
tags:
  - hacker-news
---

# Show HN: Claude Code Routine Session Cleanup Skill

- HN: [49946507](https://news.ycombinator.com/item?id=49946507)
- Source: [github.com](https://github.com/naw103/claude-routine-cleanup)
- Score: 2
- Comments: 1
- Posted: 2026-10-03T18:20:21Z

## Translation

Title: Show HN: Claude Code Routine Session Cleanup Skill
Article title: GitHub - naw103/claude-routine-cleanup: Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. · GitHub
Description: Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. - naw103/claude-routine-cleanup
HN text: Claude Code's scheduled tasks (routines) leaves a session behind on every run and in the desktop app you can only delete them one at a time. We had 618 open sessions in just one routine. The skill plans the clean up (keeps the newest N and protect any sessions where you typed a reply - optional) and then has Claude delete through the apps own session tool, 25 per call each behind an approval card. I tried just deleting the sessions on disk manually first but deleting the files directly dosnt work because the app keeps the list in memory and writes records back. The apps delete tool is turned off inside routine runs so the clean up command has to run from a normal conversation session. The planner is a single python file that dosnt delete by itself so you can run it by hand to see where everything exists.

Article text:
GitHub - naw103/claude-routine-cleanup: Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. · GitHub
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
naw103
/
claude-routine-cleanup
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows skills/ routine-cleanup skills/ routine-cleanup tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Routine Cleanup for Claude Code
Scheduled tasks (routines) in the Claude desktop app leave a session behind every
time they run. An hourly routine makes 700+ sessions a month. They pile up in the
routine's Runs list, take disk space, and slow the app down when it loads them.
Until now, the only way to remove them was one click at a time.
Routine Cleanup is a Claude Code skill that prunes a routine's old runs in bulk:
Keeps the newest N runs you choose.
Never touches a run you replied in. It reads each transcript and only
counts what a person actually typed, not skill loads, tool output, or background
notifications.
Shows the full plan first and deletes nothing until you approve it.
Deletes through the desktop app itself , so the Runs list updates
immediately, instead of leaving "Session not found on disk" ghosts behind.
First: stop the 30-day expiry from deleting your history
Claude Code already deletes old session data on its own. Anything under ~/.claude
older than cleanupPeriodDays
is swept: transcripts, subagent transcripts, checkpoint snapshots, plans. The default is
30 days , so a terminal session you want to resume or search next month is gone.
Desktop app sessions (including routine runs) are kept at any age since Claude Code
v2.1.248, unless you set desktopSessionCleanupPeriodDays . Earlier versions deleted
them after cleanupPeriodDays too, and the sweep removes only the transcript, not the
app's Runs entry. That leaves runs you can no longer open, which this skill reports as
"unverifiable".
An age cutoff can't tell the sessions you care about from routine noise, so it deletes
both. We recommend turning the expiry up so nothing is lost to age, and pruning routine
runs with this skill instead:
// ~/.claude/settings.json
{ "cleanupPeriodDays" : 3650 }
Leave desktopSessionCleanupPeriodDays unset. ( 0 is not "keep forever": Claude Code
rejects it. The minimum is 1.) The planner prints a reminder when your setting is
below a year.
Routine : Nightly report
scheduledTaskId: nightly-report
Total runs : 618
Plan (keeping newest 4):
AUTONOMOUS (no human reply) 332 runs 403.3 MB WILL DELETE 2026-07-16 .. 2026-10-02
ENGAGED (a human replied) 43 runs 120.0 MB PROTECTED (needs --include-engaged)
UNVERIFIABLE (transcript gone) 239 runs 47.2 MB PROTECTED (needs --include-unverifiable)
Selected for deletion: 332 runs, 403.3 MB, in 14 batches of <=25
PLAN ONLY: nothing was deleted. Deletion goes through the app's delete_session tool.
Install
As a plugin (Claude Code CLI or the desktop app's Code tab):
/plugin marketplace add naw103/claude-routine-cleanup
/plugin install routine-cleanup@routine-cleanup
As a plain skill: copy skills/routine-cleanup/ to ~/.claude/skills/routine-cleanup/ .
git clone https://github.com/naw103/claude-routine-cleanup
cp -r claude-routine-cleanup/skills/routine-cleanup ~ /.claude/skills/
Requires Python 3.9+ (standard library only) and the Claude desktop app.
In an ordinary conversation in the Code tab (not inside a routine run):
/routine-cleanup nightly-report keep 10
or just ask: "clean up my bug-fix routine's old runs, keep the last 10" .
Claude lists your routines, confirms which one you mean, shows the plan, and waits
for your yes. The app then shows its own approval card for each batch of 25 runs.
Also never selected: the newest N runs, the conversation you are in, and any run
active in the last 15 minutes.
If your old transcripts were pruned by Claude Code's cleanupPeriodDays and you are
happy to delete those runs too, either say so when asked or make it the default:
// ~/.claude/routine-cleanup.json
{ "include_unverifiable" : true }
Why it deletes through the app
Two things learned the hard way, so you don't have to:
Deleting the files directly does not work. The desktop app keeps the Runs list
in memory and writes a run's record back when it saves or you click the run. You
get "Session not found on disk" entries that won't go away. So the bundled script
only plans ; Claude deletes through the app's own delete_session tool.
That tool is off inside routine runs , and accepts 25 sessions per call. So the
skill runs from a normal conversation and works in batches of 25.
Known limitation: 25 runs per approval
The desktop app's delete tool accepts at most 25 sessions per call, and every call
shows its own approval card. That limit is the app's, not this skill's, and there is
no way around it: deleting the files directly is exactly what leaves the ghost entries
described above.
To make it less tedious, Claude sends several batches at once, so the cards arrive
together and you can approve them one after another without waiting between them.
The plan tells you the batch count up front, so you know how many clicks it will be
before you start. Runs only pile up this far once; after the first cleanup, running
it every few weeks keeps it to a card or two.
If you know a way to delete more per call, or Anthropic raises the limit, please
open an issue.
skills/routine-cleanup/cleanup_runs.py never deletes anything, so it is safe to run
by hand to see where your disk went:
python3 skills/routine-cleanup/cleanup_runs.py # list routines
python3 skills/routine-cleanup/cleanup_runs.py --routine ID --keep 10 --brief
python3 skills/routine-cleanup/cleanup_runs.py --routine ID --keep 10 --json
Flag
Meaning
--routine ID
exact scheduledTaskId (omit to list routines)
--keep N / --session ID
bulk, or one run by id or unique prefix
--include-engaged
also select runs you replied in
--include-unverifiable / --exclude-unverifiable
override the config default
--live-window-minutes M
skip runs active this recently (default 15)
--brief
counts and date ranges instead of one line per run
--ids / --ids-file PATH
session ids in batches of 25
--json
machine-readable plan
Environment overrides: CLAUDE_CONFIG_DIR (default ~/.claude ) and
ROUTINE_CLEANUP_APP_DIR (default: the app's data dir for your OS).
Built and tested on macOS with the Claude desktop app. The app data paths for
Windows ( %APPDATA%\Claude ) and Linux ( ~/.config/Claude ) follow the app's usual
layout but are untested; ROUTINE_CLEANUP_APP_DIR overrides them. Issues and PRs welcome.
python3 -m unittest discover -s tests -v
claude plugin validate . --strict
License
MIT. Not affiliated with Anthropic.
Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin.
github.com/naw103/claude-routine-cleanup Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. - naw103/claude-routine-cleanup

Claude Code's scheduled tasks (routines) leaves a session behind on every run and in the desktop app you can only delete them one at a time. We had 618 open sessions in just one routine. The skill plans the clean up (keeps the newest N and protect any sessions where you typed a reply - optional) and then has Claude delete through the apps own session tool, 25 per call each behind an approval card. I tried just deleting the sessions on disk manually first but deleting the files directly dosnt work because the app keeps the list in memory and writes records back. The apps delete tool is turned off inside routine runs so the clean up command has to run from a normal conversation session. The planner is a single python file that dosnt delete by itself so you can run it by hand to see where everything exists.

GitHub - naw103/claude-routine-cleanup: Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin. · GitHub
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
naw103
/
claude-routine-cleanup
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows skills/ routine-cleanup skills/ routine-cleanup tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Routine Cleanup for Claude Code
Scheduled tasks (routines) in the Claude desktop app leave a session behind every
time they run. An hourly routine makes 700+ sessions a month. They pile up in the
routine's Runs list, take disk space, and slow the app down when it loads them.
Until now, the only way to remove them was one click at a time.
Routine Cleanup is a Claude Code skill that prunes a routine's old runs in bulk:
Keeps the newest N runs you choose.
Never touches a run you replied in. It reads each transcript and only
counts what a person actually typed, not skill loads, tool output, or background
notifications.
Shows the full plan first and deletes nothing until you approve it.
Deletes through the desktop app itself , so the Runs list updates
immediately, instead of leaving "Session not found on disk" ghosts behind.
First: stop the 30-day expiry from deleting your history
Claude Code already deletes old session data on its own. Anything under ~/.claude
older than cleanupPeriodDays
is swept: transcripts, subagent transcripts, checkpoint snapshots, plans. The default is
30 days , so a terminal session you want to resume or search next month is gone.
Desktop app sessions (including routine runs) are kept at any age since Claude Code
v2.1.248, unless you set desktopSessionCleanupPeriodDays . Earlier versions deleted
them after cleanupPeriodDays too, and the sweep removes only the transcript, not the
app's Runs entry. That leaves runs you can no longer open, which this skill reports as
"unverifiable".
An age cutoff can't tell the sessions you care about from routine noise, so it deletes
both. We recommend turning the expiry up so nothing is lost to age, and pruning routine
runs with this skill instead:
// ~/.claude/settings.json
{ "cleanupPeriodDays" : 3650 }
Leave desktopSessionCleanupPeriodDays unset. ( 0 is not "keep forever": Claude Code
rejects it. The minimum is 1.) The planner prints a reminder when your setting is
below a year.
Routine : Nightly report
scheduledTaskId: nightly-report
Total runs : 618
Plan (keeping newest 4):
AUTONOMOUS (no human reply) 332 runs 403.3 MB WILL DELETE 2026-07-16 .. 2026-10-02
ENGAGED (a human replied) 43 runs 120.0 MB PROTECTED (needs --include-engaged)
UNVERIFIABLE (transcript gone) 239 runs 47.2 MB PROTECTED (needs --include-unverifiable)
Selected for deletion: 332 runs, 403.3 MB, in 14 batches of <=25
PLAN ONLY: nothing was deleted. Deletion goes through the app's delete_session tool.
Install
As a plugin (Claude Code CLI or the desktop app's Code tab):
/plugin marketplace add naw103/claude-routine-cleanup
/plugin install routine-cleanup@routine-cleanup
As a plain skill: copy skills/routine-cleanup/ to ~/.claude/skills/routine-cleanup/ .
git clone https://github.com/naw103/claude-routine-cleanup
cp -r claude-routine-cleanup/skills/routine-cleanup ~ /.claude/skills/
Requires Python 3.9+ (standard library only) and the Claude desktop app.
In an ordinary conversation in the Code tab (not inside a routine run):
/routine-cleanup nightly-report keep 10
or just ask: "clean up my bug-fix routine's old runs, keep the last 10" .
Claude lists your routines, confirms which one you mean, shows the plan, and waits
for your yes. The app then shows its own approval card for each batch of 25 runs.
Also never selected: the newest N runs, the conversation you are in, and any run
active in the last 15 minutes.
If your old transcripts were pruned by Claude Code's cleanupPeriodDays and you are
happy to delete those runs too, either say so when asked or make it the default:
// ~/.claude/routine-cleanup.json
{ "include_unverifiable" : true }
Why it deletes through the app
Two things learned the hard way, so you don't have to:
Deleting the files directly does not work. The desktop app keeps the Runs list
in memory and writes a run's record back when it saves or you click the run. You
get "Session not found on disk" entries that won't go away. So the bundled script
only plans ; Claude deletes through the app's own delete_session tool.
That tool is off inside routine runs , and accepts 25 sessions per call. So the
skill runs from a normal conversation and works in batches of 25.
Known limitation: 25 runs per approval
The desktop app's delete tool accepts at most 25 sessions per call, and every call
shows its own approval card. That limit is the app's, not this skill's, and there is
no way around it: deleting the files directly is exactly what leaves the ghost entries
described above.
To make it less tedious, Claude sends several batches at once, so the cards arrive
together and you can approve them one after another without waiting between them.
The plan tells you the batch count up front, so you know how many clicks it will be
before you start. Runs only pile up this far once; after the first cleanup, running
it every few weeks keeps it to a card or two.
If you know a way to delete more per call, or Anthropic raises the limit, please
open an issue.
skills/routine-cleanup/cleanup_runs.py never deletes anything, so it is safe to run
by hand to see where your disk went:
python3 skills/routine-cleanup/cleanup_runs.py # list routines
python3 skills/routine-cleanup/cleanup_runs.py --routine ID --keep 10 --brief
python3 skills/routine-cleanup/cleanup_runs.py --routine ID --keep 10 --json
Flag
Meaning
--routine ID
exact scheduledTaskId (omit to list routines)
--keep N / --session ID
bulk, or one run by id or unique prefix
--include-engaged
also select runs you replied in
--include-unverifiable / --exclude-unverifiable
override the config default
--live-window-minutes M
skip runs active this recently (default 15)
--brief
counts and date ranges instead of one line per run
--ids / --ids-file PATH
session ids in batches of 25
--json
machine-readable plan
Environment overrides: CLAUDE_CONFIG_DIR (default ~/.claude ) and
ROUTINE_CLEANUP_APP_DIR (default: the app's data dir for your OS).
Built and tested on macOS with the Claude desktop app. The app data paths for
Windows ( %APPDATA%\Claude ) and Linux ( ~/.config/Claude ) follow the app's usual
layout but are untested; ROUTINE_CLEANUP_APP_DIR overrides them. Issues and PRs welcome.
python3 -m unittest discover -s tests -v
claude plugin validate . --strict
License
MIT. Not affiliated with Anthropic.
Bulk-delete old Claude Code scheduled-task (routine) runs safely: keeps the newest N and every run you replied in. A Claude Code skill + plugin.
github.com/naw103/claude-routine-cleanup Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
