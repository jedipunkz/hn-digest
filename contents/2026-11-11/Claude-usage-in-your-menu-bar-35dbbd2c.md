---
source: "https://github.com/trthomps/claude-usage-widget"
hn_url: "https://news.ycombinator.com/item?id=50460269"
title: "Claude usage in your menu bar"
article_title: "GitHub - trthomps/claude-usage-widget: macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code · GitHub"
image: "https://opengraph.githubassets.com/30e562b27ae92cfd0192fa36aa86a393f0094ac0a1b47788ba2c66933bdf0c5c/trthomps/claude-usage-widget"
author: "trthomps"
captured_at: "2026-11-11T23:20:53Z"
capture_tool: "hn-digest"
hn_id: 50460269
score: 1
comments: 1
posted_at: "2026-11-11T23:09:51Z"
tags:
  - hacker-news
---

# Claude usage in your menu bar

- HN: [50460269](https://news.ycombinator.com/item?id=50460269)
- Source: [github.com](https://github.com/trthomps/claude-usage-widget)
- Score: 1
- Comments: 1
- Posted: 2026-11-11T23:09:51Z

## Translation

Title: Claude usage in your menu bar
Article title: GitHub - trthomps/claude-usage-widget: macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code · GitHub
Description: macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code - trthomps/claude-usage-widget

Article text:
GitHub - trthomps/claude-usage-widget: macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code · GitHub
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
trthomps
claude-usage-widget Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
Sources/ ClaudeUsage Sources/ ClaudeUsage docs docs scripts scripts .gitignore .gitignore Makefile Makefile Package.swift Package.swift README.md README.md View all files Repository files navigation
A macOS menu bar item and desktop widget that show your Claude plan limits: the 5-hour session, the weekly cap, and any per-model weekly caps. These are the same numbers /usage shows in Claude Code.
Requires macOS 14+, the Xcode Command Line Tools ( xcode-select --install ), and a Claude Code login ( claude , then /login ).
make install # builds "Claude Usage.app", copies it to ~/Applications, launches it
To build without installing, run make app . The app ends up in build/ .
Don't see it in the menu bar? Menu bar managers such as Hidden Bar, Ice and Bartender put new items in their hidden section. Expand that section and ⌘-drag the item into the visible part. Opening the app again (Spotlight, Finder) always brings up Settings.
Click the menu bar item for the detail popover: each limit, its reset time, a pace marker, and a forecast ("Limit in ~1h 12m" or "~62% at reset").
Right-click for Refresh, the desktop widget toggle, Settings and Quit.
Desktop widget : turn it on in Settings or from the right-click menu. It sits on the desktop, above the icons and below your windows. Drag it anywhere; right-click it for options, including "Keep Above Other Windows".
Terminal : ClaudeUsage --print prints the same table as /usage . --json gives you the parsed data as JSON.
Claude usage · Team · 5x
Current session █████████░░░░░░░░░░░░░░░ 37% Resets in 3h 49m · 18:39 ⚠ limit in ~1h 59m at this pace
Current week (all models) ████░░░░░░░░░░░░░░░░░░░░ 17% Resets Tue 06:59 · in 5d 16h (~90% at reset)
Current week (Fable) ░░░░░░░░░░░░░░░░░░░░░░░░ 0% Resets Tue 07:00 · in 5d 16h
Menu bar design
Menu bar space is scarce, especially on notched MacBooks, where items that don't fit just disappear. The design started from three questions you want answered at a glance: how much have I used, am I burning it too fast, and when does it come back?
Pills (default), about 80pt for two limits. Each limit is a small rounded bar with its percentage printed on it, like the battery glyph. The fill knocks the digits out of itself, so the number stays readable wherever the fill edge lands, on a light or dark menu bar. Digits are monospaced and the pill is sized for "88%", so going from 19% to 20% never moves your other menu bar items.
The pace caret. The small caret under each pill marks how much of the window has passed. Its position is where your usage would be if you spread it evenly. If the fill is past the caret , you're on track to hit the limit before it resets. If it's behind, you have room. The caret takes no extra width, and it sits outside the pill so it never cuts through the digits. An earlier version drew it inside the pill, which made "47%" look like "4|7%".
Color means something. In the default Warnings mode, pills use the menu bar's own text color and look like any monochrome system item until a limit passes 75% (orange) or 90% (red). Claude mode tints them clay. Mono never colors them.
At 100% the pill shows a countdown. Once you're blocked, the percentage tells you nothing, so the pill turns red and shows the time until reset instead ( 1h12m ).
Stale data fades. If the last successful fetch is too old (offline, expired login), the item is drawn at reduced opacity instead of showing numbers you shouldn't trust. The popover says why.
Other styles, in Settings → Menu bar style:
You choose which limits appear (session, week, per-model weeks). The order always follows the server's: session first, then weekly.
Data : GET https://api.anthropic.com/api/oauth/usage , the endpoint Claude Code's /usage reads. It is undocumented and may change. The parser prefers the structured limits array, falls back to the older flat five_hour / seven_day_* keys, and ignores fields it doesn't recognize.
Login : reads Claude Code's OAuth token from the Keychain item Claude Code-credentials using /usr/bin/security . Claude Code writes that item with the same tool, so reading it doesn't trigger a Keychain prompt. It falls back to ~/.claude/.credentials.json . The app never writes credentials and never refreshes the token itself. That would rotate Claude Code's refresh token out from under it. If the token has expired, the app keeps showing the last known numbers (faded) until Claude Code renews it, which happens the next time you use it.
Polling : every 2 minutes by default (configurable), when the popover opens, and after wake. On HTTP 429 it backs off. Only percentages and reset times are cached, in ~/Library/Application Support/ClaudeUsage/snapshot.json . Tokens are never written to disk.
Reset rollover : when a window's reset time passes between fetches, it shows 0% right away instead of a stale number.
swift build # debug build
swift run ClaudeUsage --print # live usage in the terminal
swift run ClaudeUsage --render-previews docs # regenerate the images in this README
CLAUDE_USAGE_DEBUG=1 " build/Claude Usage.app/Contents/MacOS/ClaudeUsage " # log fetches and draws
File
What's in it
MenuBarRenderer.swift
All menu bar drawing (Core Graphics). Used by the live item, the Settings tiles and the preview renderer.
UsageParser.swift
API response → UsageSnapshot
Credentials.swift
Finding Claude Code's login
UsageStore.swift
Polling, caching, backoff, the clock behind countdowns
StatusItemController.swift , PopoverView.swift
Menu bar item and popover
DesktopWidget.swift
Desktop panel (Liquid Glass on macOS 26, vibrancy on older versions)
SettingsView.swift
Settings window
It's a SwiftPM executable, so you don't need Xcode. scripts/build-app.sh wraps the binary in an app bundle, generates the icon, and signs it ad hoc. Set SIGN_IDENTITY to sign with a Developer ID.
A native WidgetKit widget, the kind in Notification Center and on the desktop since Sonoma, would need an Xcode project, a signed app extension and an App Group. The desktop panel gives you the same view without any of that.
macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code - trthomps/claude-usage-widget

GitHub - trthomps/claude-usage-widget: macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code · GitHub
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
trthomps
claude-usage-widget Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
Sources/ ClaudeUsage Sources/ ClaudeUsage docs docs scripts scripts .gitignore .gitignore Makefile Makefile Package.swift Package.swift README.md README.md View all files Repository files navigation
A macOS menu bar item and desktop widget that show your Claude plan limits: the 5-hour session, the weekly cap, and any per-model weekly caps. These are the same numbers /usage shows in Claude Code.
Requires macOS 14+, the Xcode Command Line Tools ( xcode-select --install ), and a Claude Code login ( claude , then /login ).
make install # builds "Claude Usage.app", copies it to ~/Applications, launches it
To build without installing, run make app . The app ends up in build/ .
Don't see it in the menu bar? Menu bar managers such as Hidden Bar, Ice and Bartender put new items in their hidden section. Expand that section and ⌘-drag the item into the visible part. Opening the app again (Spotlight, Finder) always brings up Settings.
Click the menu bar item for the detail popover: each limit, its reset time, a pace marker, and a forecast ("Limit in ~1h 12m" or "~62% at reset").
Right-click for Refresh, the desktop widget toggle, Settings and Quit.
Desktop widget : turn it on in Settings or from the right-click menu. It sits on the desktop, above the icons and below your windows. Drag it anywhere; right-click it for options, including "Keep Above Other Windows".
Terminal : ClaudeUsage --print prints the same table as /usage . --json gives you the parsed data as JSON.
Claude usage · Team · 5x
Current session █████████░░░░░░░░░░░░░░░ 37% Resets in 3h 49m · 18:39 ⚠ limit in ~1h 59m at this pace
Current week (all models) ████░░░░░░░░░░░░░░░░░░░░ 17% Resets Tue 06:59 · in 5d 16h (~90% at reset)
Current week (Fable) ░░░░░░░░░░░░░░░░░░░░░░░░ 0% Resets Tue 07:00 · in 5d 16h
Menu bar design
Menu bar space is scarce, especially on notched MacBooks, where items that don't fit just disappear. The design started from three questions you want answered at a glance: how much have I used, am I burning it too fast, and when does it come back?
Pills (default), about 80pt for two limits. Each limit is a small rounded bar with its percentage printed on it, like the battery glyph. The fill knocks the digits out of itself, so the number stays readable wherever the fill edge lands, on a light or dark menu bar. Digits are monospaced and the pill is sized for "88%", so going from 19% to 20% never moves your other menu bar items.
The pace caret. The small caret under each pill marks how much of the window has passed. Its position is where your usage would be if you spread it evenly. If the fill is past the caret , you're on track to hit the limit before it resets. If it's behind, you have room. The caret takes no extra width, and it sits outside the pill so it never cuts through the digits. An earlier version drew it inside the pill, which made "47%" look like "4|7%".
Color means something. In the default Warnings mode, pills use the menu bar's own text color and look like any monochrome system item until a limit passes 75% (orange) or 90% (red). Claude mode tints them clay. Mono never colors them.
At 100% the pill shows a countdown. Once you're blocked, the percentage tells you nothing, so the pill turns red and shows the time until reset instead ( 1h12m ).
Stale data fades. If the last successful fetch is too old (offline, expired login), the item is drawn at reduced opacity instead of showing numbers you shouldn't trust. The popover says why.
Other styles, in Settings → Menu bar style:
You choose which limits appear (session, week, per-model weeks). The order always follows the server's: session first, then weekly.
Data : GET https://api.anthropic.com/api/oauth/usage , the endpoint Claude Code's /usage reads. It is undocumented and may change. The parser prefers the structured limits array, falls back to the older flat five_hour / seven_day_* keys, and ignores fields it doesn't recognize.
Login : reads Claude Code's OAuth token from the Keychain item Claude Code-credentials using /usr/bin/security . Claude Code writes that item with the same tool, so reading it doesn't trigger a Keychain prompt. It falls back to ~/.claude/.credentials.json . The app never writes credentials and never refreshes the token itself. That would rotate Claude Code's refresh token out from under it. If the token has expired, the app keeps showing the last known numbers (faded) until Claude Code renews it, which happens the next time you use it.
Polling : every 2 minutes by default (configurable), when the popover opens, and after wake. On HTTP 429 it backs off. Only percentages and reset times are cached, in ~/Library/Application Support/ClaudeUsage/snapshot.json . Tokens are never written to disk.
Reset rollover : when a window's reset time passes between fetches, it shows 0% right away instead of a stale number.
swift build # debug build
swift run ClaudeUsage --print # live usage in the terminal
swift run ClaudeUsage --render-previews docs # regenerate the images in this README
CLAUDE_USAGE_DEBUG=1 " build/Claude Usage.app/Contents/MacOS/ClaudeUsage " # log fetches and draws
File
What's in it
MenuBarRenderer.swift
All menu bar drawing (Core Graphics). Used by the live item, the Settings tiles and the preview renderer.
UsageParser.swift
API response → UsageSnapshot
Credentials.swift
Finding Claude Code's login
UsageStore.swift
Polling, caching, backoff, the clock behind countdowns
StatusItemController.swift , PopoverView.swift
Menu bar item and popover
DesktopWidget.swift
Desktop panel (Liquid Glass on macOS 26, vibrancy on older versions)
SettingsView.swift
Settings window
It's a SwiftPM executable, so you don't need Xcode. scripts/build-app.sh wraps the binary in an app bundle, generates the icon, and signs it ad hoc. Set SIGN_IDENTITY to sign with a Developer ID.
A native WidgetKit widget, the kind in Notification Center and on the desktop since Sonoma, would need an Xcode project, a signed app extension and an App Group. The desktop panel gives you the same view without any of that.
macOS menu bar + desktop widget for Claude plan usage (5-hour session, weekly limits), like /usage in Claude Code
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
