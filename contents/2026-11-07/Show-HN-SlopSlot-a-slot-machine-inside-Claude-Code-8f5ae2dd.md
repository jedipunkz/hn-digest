---
source: "https://github.com/gajus/slopslot"
hn_url: "https://news.ycombinator.com/item?id=50400511"
title: "Show HN: SlopSlot – a slot machine inside Claude Code"
article_title: "GitHub - gajus/slopslot: A slot machine inside Claude Code · GitHub"
image: "https://opengraph.githubassets.com/293eab3cb2f9df3101496c5417cdc3035d95e1cb9719da35dfdf409c9ef3e905/gajus/slopslot"
author: "gajus"
captured_at: "2026-11-07T01:39:48Z"
capture_tool: "hn-digest"
hn_id: 50400511
score: 1
comments: 0
posted_at: "2026-11-07T00:48:29Z"
tags:
  - hacker-news
---

# Show HN: SlopSlot – a slot machine inside Claude Code

- HN: [50400511](https://news.ycombinator.com/item?id=50400511)
- Source: [github.com](https://github.com/gajus/slopslot)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T00:48:29Z

## Translation

Title: Show HN: SlopSlot – a slot machine inside Claude Code
Article title: GitHub - gajus/slopslot: A slot machine inside Claude Code · GitHub
Description: A slot machine inside Claude Code. Contribute to gajus/slopslot development by creating an account on GitHub.
HN text: Because that's how it feels like. claude plugin install slopslot --marketplace gajus/slopslot

Article text:
GitHub - gajus/slopslot: A slot machine inside Claude Code · GitHub
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
gajus
slopslot Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A slot machine inside Claude Code
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks previews previews tests tests .gitignore .gitignore README.md README.md package.json package.json View all files Repository files navigation
A slot machine inside Claude Code.
Every prompt pulls the lever: three reels
spin while Claude works, then stop one at a time. Collect points, climb the
ranks, unlock badges, and occasionally get paid for absolutely no skill.
You need Claude Code 2.1.319 or later . No dependencies, API keys, or build
step are required.
claude plugin install slopslot --marketplace gajus/slopslot
Or run this inside Claude Code:
/plugin install slopslot --marketplace gajus/slopslot
Start a new Claude Code session after installation. Accept any workspace or
plugin trust prompts Claude Code shows. Send a normal
prompt and the reels will spin automatically. For a spin without asking Claude
to do anything, run:
/slopslot
Want the full machine?
/slopslot cabinet
The cabinet opens in a pane, with a Spin button on wider displays. Close it
with Escape or its Close control to return to the display above your prompt.
Check /plugin → Installed if SlopSlot does not appear.
claude plugin update slopslot@gajus-slopslot
SlopSlot is distributed through the marketplace in
gajus/slopslot . It does not require npm.
A little casino for your context window
Terminal: original pixel sprites, scrolling reels, chasing lights, a
moving lever, and reward sparkles.
Claude Desktop renderer: a shaded SVG cabinet, illustrated symbols,
and bouncing gold coins. Jackpots get a bigger celebration.
Progress: an XP bar, rank titles, win streaks, a gift countdown, and six
collectible badges.
Small screens: a compact text display when there is not enough room for
the artwork. /slopslot cabinet opens a separate view.
These images are artwork previews with example symbols and scores. The terminal
has been checked in a live Claude Code session. The Desktop renderer has been
checked through rendering tests and standalone previews, rather than a live
Desktop session. Both were developed against Claude Code 2.1.319 .
Command
What it does
/slopslot or /slopslot spin
Spin without a model request; also enables the machine.
/slopslot cabinet
Open the larger interactive cabinet.
/slopslot rewards
Show payout rules, badges, and unlock conditions.
/slopslot status
Show points, level, XP, badge count, and best spin.
/slopslot off
Hide the machine, close the cabinet, and stop automatic spins.
/slopslot on
Show the machine and resume automatic spins.
/slopslot reset
Clear session progress and cancel an unfinished spin.
The extremely serious economy
Each reel independently picks one of six equally likely symbols: cherries,
lemon, grapes, bell, star, or seven.
Consecutive matching spins multiply the base payout: ×1 → ×2 → ×3 . The
multiplier stays at ×3 until a miss resets the streak. It applies to base
points only; bonuses are added afterward.
Every completed spin earns 20 XP , plus 10 extra XP for matching reels.
Every 100 XP earns a new level and 25 bonus points . Every fifth completed
spin adds a 25-point gift . Even a miss moves you forward.
Rank titles advance every two levels:
Rookie → Regular → High roller → VIP → Legend .
Each badge pays once per session. One spin can unlock multiple badges.
Spins are free. Points have no monetary value. Your only real prize is having
installed a slot machine in your coding assistant.
What happens to your prompts and progress?
SlopSlot runs locally. It makes no network requests, model calls, or file
writes, and adds no text to your prompts or conversation. A spin starts
immediately without waiting for the reels to finish before letting your prompt
continue.
Progress lives in memory and resets when the plugin reloads or the session
restarts. Sending another prompt during a spin settles that result before
starting the next one. Reset cancels an unfinished spin without awarding it.
Reels stop after two seconds. Reward animations last another 1.2 seconds, then
stop repainting. Turning SlopSlot off cancels its timers and leaves other mods'
prompt content visible.
Clone the repository and load your working copy for one session:
git clone https://github.com/gajus/slopslot.git
cd slopslot
claude --plugin-dir .
Validate the plugin and run its tests:
claude plugin validate .
claude plugin test
The test suite covers scoring, streaks, bonuses, badge unlocks, prompt
passthrough, timer cleanup, controls, and both graphical renderers.
Claude Code watches a folder loaded with --plugin-dir and reloads changes.
Reloading also resets SlopSlot's session progress.
A slot machine inside Claude Code
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A slot machine inside Claude Code. Contribute to gajus/slopslot development by creating an account on GitHub.

Because that's how it feels like. claude plugin install slopslot --marketplace gajus/slopslot

GitHub - gajus/slopslot: A slot machine inside Claude Code · GitHub
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
gajus
slopslot Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A slot machine inside Claude Code
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks previews previews tests tests .gitignore .gitignore README.md README.md package.json package.json View all files Repository files navigation
A slot machine inside Claude Code.
Every prompt pulls the lever: three reels
spin while Claude works, then stop one at a time. Collect points, climb the
ranks, unlock badges, and occasionally get paid for absolutely no skill.
You need Claude Code 2.1.319 or later . No dependencies, API keys, or build
step are required.
claude plugin install slopslot --marketplace gajus/slopslot
Or run this inside Claude Code:
/plugin install slopslot --marketplace gajus/slopslot
Start a new Claude Code session after installation. Accept any workspace or
plugin trust prompts Claude Code shows. Send a normal
prompt and the reels will spin automatically. For a spin without asking Claude
to do anything, run:
/slopslot
Want the full machine?
/slopslot cabinet
The cabinet opens in a pane, with a Spin button on wider displays. Close it
with Escape or its Close control to return to the display above your prompt.
Check /plugin → Installed if SlopSlot does not appear.
claude plugin update slopslot@gajus-slopslot
SlopSlot is distributed through the marketplace in
gajus/slopslot . It does not require npm.
A little casino for your context window
Terminal: original pixel sprites, scrolling reels, chasing lights, a
moving lever, and reward sparkles.
Claude Desktop renderer: a shaded SVG cabinet, illustrated symbols,
and bouncing gold coins. Jackpots get a bigger celebration.
Progress: an XP bar, rank titles, win streaks, a gift countdown, and six
collectible badges.
Small screens: a compact text display when there is not enough room for
the artwork. /slopslot cabinet opens a separate view.
These images are artwork previews with example symbols and scores. The terminal
has been checked in a live Claude Code session. The Desktop renderer has been
checked through rendering tests and standalone previews, rather than a live
Desktop session. Both were developed against Claude Code 2.1.319 .
Command
What it does
/slopslot or /slopslot spin
Spin without a model request; also enables the machine.
/slopslot cabinet
Open the larger interactive cabinet.
/slopslot rewards
Show payout rules, badges, and unlock conditions.
/slopslot status
Show points, level, XP, badge count, and best spin.
/slopslot off
Hide the machine, close the cabinet, and stop automatic spins.
/slopslot on
Show the machine and resume automatic spins.
/slopslot reset
Clear session progress and cancel an unfinished spin.
The extremely serious economy
Each reel independently picks one of six equally likely symbols: cherries,
lemon, grapes, bell, star, or seven.
Consecutive matching spins multiply the base payout: ×1 → ×2 → ×3 . The
multiplier stays at ×3 until a miss resets the streak. It applies to base
points only; bonuses are added afterward.
Every completed spin earns 20 XP , plus 10 extra XP for matching reels.
Every 100 XP earns a new level and 25 bonus points . Every fifth completed
spin adds a 25-point gift . Even a miss moves you forward.
Rank titles advance every two levels:
Rookie → Regular → High roller → VIP → Legend .
Each badge pays once per session. One spin can unlock multiple badges.
Spins are free. Points have no monetary value. Your only real prize is having
installed a slot machine in your coding assistant.
What happens to your prompts and progress?
SlopSlot runs locally. It makes no network requests, model calls, or file
writes, and adds no text to your prompts or conversation. A spin starts
immediately without waiting for the reels to finish before letting your prompt
continue.
Progress lives in memory and resets when the plugin reloads or the session
restarts. Sending another prompt during a spin settles that result before
starting the next one. Reset cancels an unfinished spin without awarding it.
Reels stop after two seconds. Reward animations last another 1.2 seconds, then
stop repainting. Turning SlopSlot off cancels its timers and leaves other mods'
prompt content visible.
Clone the repository and load your working copy for one session:
git clone https://github.com/gajus/slopslot.git
cd slopslot
claude --plugin-dir .
Validate the plugin and run its tests:
claude plugin validate .
claude plugin test
The test suite covers scoring, streaks, bonuses, badge unlocks, prompt
passthrough, timer cleanup, controls, and both graphical renderers.
Claude Code watches a folder loaded with --plugin-dir and reloads changes.
Reloading also resets SlopSlot's session progress.
A slot machine inside Claude Code
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
