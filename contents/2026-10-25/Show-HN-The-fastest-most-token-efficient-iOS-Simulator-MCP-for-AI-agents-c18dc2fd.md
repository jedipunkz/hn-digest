---
source: "https://github.com/NoMac-app/simulator-use"
hn_url: "https://news.ycombinator.com/item?id=50231724"
title: "Show HN: The fastest, most token-efficient iOS Simulator MCP for AI agents"
article_title: "GitHub - NoMac-app/simulator-use: Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. · GitHub"
image: "https://opengraph.githubassets.com/c24d35b42580b28830f11d297a516927ecb4d0760690d4ce9040d93dc15cc25d/NoMac-app/simulator-use"
author: "garymiklos"
captured_at: "2026-10-25T14:52:03Z"
capture_tool: "hn-digest"
hn_id: 50231724
score: 1
comments: 0
posted_at: "2026-10-25T14:05:18Z"
tags:
  - hacker-news
---

# Show HN: The fastest, most token-efficient iOS Simulator MCP for AI agents

- HN: [50231724](https://news.ycombinator.com/item?id=50231724)
- Source: [github.com](https://github.com/NoMac-app/simulator-use)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T14:05:18Z

## Translation

Title: Show HN: The fastest, most token-efficient iOS Simulator MCP for AI agents
Article title: GitHub - NoMac-app/simulator-use: Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. · GitHub
Description: Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. - NoMac-app/simulator-use

Article text:
GitHub - NoMac-app/simulator-use: Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. · GitHub
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
NoMac-app
simulator-use Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line.
www.npmjs.com/package/simulator-use Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
85 Commits 85 Commits Folders and files
.github/ workflows .github/ workflows bin bin examples examples native native scripts scripts skill skill src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md THIRD_PARTY_LICENSES THIRD_PARTY_LICENSES action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
The fastest, most token-efficient way for an AI agent to use the iOS Simulator. simulator-use gives Claude Code, Codex, Cursor and any other MCP client a phone: it builds and runs your Xcode project, reads the screen as a few lines of text, taps and types by label, and runs your tests case by case. Its own native engine talks to the Simulator directly, so reading the whole screen takes about 250 ms and a tap goes out in under a millisecond, where tools that start a helper for every call take 1.5 to 3 seconds. Free and open source (MIT), on any Mac with Xcode, and from Linux, Windows or a Mac without Xcode through a cloud iPhone.
Fast: a whole-screen read in about 250 ms, a tap sent in under 1 ms, a screenshot in about 60 ms (measured on an M6 Mac mini, October 2026).
Token-efficient: one line per element the agent can use, nothing else; a dozen tools with short descriptions.
Complete: build, install, launch, XCTest and XCUITest per case, logs and crashes, video, permission prompts, iPhone and iPad.
Continuously improved: weekly releases driven by thousands of daily sessions.
Anywhere: npx simulator-use cloud start gives a cloud iPhone with Xcode to any machine.
It builds, runs and tests Xcode projects, reads and uses the screen of any app (SwiftUI, UIKit, React Native), answers iOS's permission prompts, reads the app's logs and crashes, and records video.
Demo · Sign In · iPhone 17 Pro (iOS 26.5)
textfield "Email" #email @201,194
securefield "Password" #password @201,246
switch "Remember me" #remember off @341,298
button "Sign In" #signIn @201,350
text "Not signed in" #status @201,402
tabs: Login (selected) @158,822, Items @244,822
One line per element: kind, label, #id (the app's accessibility identifier), state and @x,y , the center in points.
Groups, icons inside buttons, scroll bars, the status bar and the keyboard's keys are folded away. The tab bar, scrolling and the keyboard are one line each.
Only what can be tapped is listed. An alert, a sheet, a menu or a full-screen view in front hides the screen behind it, and iOS's own permission prompts read as alerts over the app. Rows under the navigation bar, the tab bar or the keyboard are counted as hidden (above, below, to the left or right), not listed.
A web view in UIKit (Capacitor, React Native's WebView, a sign-in page inside most apps) is read point by point, and its headings, texts, fields, checkboxes, buttons and links are listed like the rest. Web views hosted by SwiftUI, Safari's own pages, in-app browsers, maps, games and iOS's out-of-process pickers have nothing for accessibility to read. They come back as one unreadable line with their place, so the agent knows to take a screenshot and tap by point.
A screenshot, when asked for, is sized to points, so a point in the picture is a point to tap.
The sign-in screen above is about 70 tokens. The full accessibility tree it comes from is about 40,000.
ui_act takes steps in the same words ui_look prints, runs them in order, and answers with what changed:
ui_act steps: ["fill #email gary@example.com", "fill #password secret1", "tap #signIn", "wait \"Welcome\""]
4 steps done in 4.5 s
changed: #email "gary@example.com" · #password "•••••••" · #status "Welcome, gary@example.com"
Steps:
Every step is checked, so nothing succeeds silently:
Typed text is read back from the field. When iOS's autocorrect or smart punctuation changes it ( -- to — , ' to ’ ), the text is pasted instead, which iOS takes as is. A field that still shows something else stops the call and says what it shows.
Text in any language (Café, 日本語, emoji) can be typed: what a US keyboard lacks goes through the Simulator's clipboard.
fill refuses anything that isn't a text field, so keys never land in whichever field still has focus.
A tap that changes nothing is reported, after waiting 2.5 s for late answers such as a permission prompt. A tap inside an area that can't be read says it can't be checked.
A target that isn't on screen yet gets two seconds to appear, then fails with the next step: "2 elements are below: scroll down", or, when an alert came up in front, "Answer it first: tap "Cancel" or "Open"".
The first failing step stops the rest and returns the screen as it is now.
A new screen comes back whole; otherwise only what changed.
Tool
CLI
Does
app_run
simulator-use run
Builds the Xcode project here, installs and launches it, answers with the first screen; errors as file:line ; --xcode 26.6 builds with another installed Xcode
devices
simulator-use devices
Lists the Simulators on this Mac with their runtime, UDID and state, and its Xcodes
device_use
simulator-use use "iPhone 17 Pro"
Chooses the Simulator the other tools use, booting it (one boot at a time); dark mode, a clean status bar
app_control
simulator-use app launch dev.example.app
install , launch (answers with the first screen), terminate , uninstall , open_url ; grant , revoke and reset_permissions ; location ; push ; the phone's appearance , locale , timezone , text_size and a clean status_bar
ui_look
simulator-use look
Reads the screen; find keeps matching lines, screenshot adds a picture, full saves the whole tree
ui_act
simulator-use act "tap #signIn" ...
Does steps in order and answers with what changed
app_logs
simulator-use logs
The app's print output and its own system log lines; grep , wait_for , crashes
tests_run
simulator-use test
Runs the unit and UI tests; each case as it finishes, counts, each failure as file:line and the slowest; "0 tests ran" is an error
screen_record
simulator-use record start
Starts or stops a video of the screen; stop answers with the MP4
job_status
simulator-use job
Waits up to 40 s for a long build and answers with its result, or its progress
job_cancel
simulator-use cancel
Stops a job
cloud
simulator-use cloud start
start , status , stop a NoMac cloud phone that every tool then runs on; run builds, tests and stops in one call (see Cloud)
The MCP server and the CLI are built from one list, so they always match. The tool list and instructions stay under 4,000 tokens; a test enforces it.
app_run does the whole loop in one call: it finds the project (a workspace first, then a project, an XcodeGen spec, or ios/ ), picks the scheme, builds for the Simulator, installs, launches, and answers with the first screen.
The app's permissions are granted when it is installed and its first launch's prompts (notifications too) are answered, so no prompt stops the agent; permissions: "ask" keeps iOS's prompts, to test them. open_url answers iOS's "Open in …?" for a link that opens another app.
Built Demo in 4.2 s, installed dev.simtap.demo.
Launched Demo in 3.9 s.
Demo · Sign In · iPhone 17 Pro (iOS 26.5)
textfield "Email" #email @201,194
...
A failed build answers with the errors only, as Sources/App.swift:49:37: error: ... , and the path to the full log.
Each folder (each git worktree) builds into its own DerivedData, so rebuilds are incremental and worktrees never share products. SIMULATOR_USE_CACHE_DIR puts it on a disk you keep, so the next machine's build is incremental too.
An XcodeGen project is regenerated from project.yml before each build (skipped when nothing changed), so new files are always in it.
The next call in the same folder needs no arguments.
env: "API_URL=http://127.0.0.1:8080 MODE=test" gives the app those variables (it reads them from ProcessInfo ), as an Xcode scheme does; they stay for its later launches, and env: "" clears them.
An MCP call answers within 40 s; a longer build answers with a job, and job_status waits for it. From the CLI, simulator-use run simply waits.
Two builds of one project never run at once, also from two processes: the second says which process is building it.
Expo, React Native, Flutter and Capacitor builds are not done yet; build them with their own tools, then app_control install .
tests_run runs the scheme's tests and reads the answer from Xcode's result bundle, not from the scrolling output:
Tests failed: 1 passed, 2 failed (3 in 35.3 s). Result bundle: .../tests-1792370825323.xcresult
DemoTests/SignInTests/testRejectsAShortPassword(): Tests/SignInTests.swift:12: XCTAssertEqual failed: ("Welcome, gary@example.com") is not equal to ("Invalid email or password")
DemoTests/rejectsAnEmailWithoutAt(): Tests/SignInTests.swift:18: Expectation failed: ...
XCTest and Swift Testing both.
Each failing test is named the way only takes it, to rerun just that one.
"0 tests ran" (no test target, or a filter that matched nothing) is an error, never a pass.
app_logs shows what the app itself said, not the thousands of lines Apple's frameworks log from inside it:
Its print() output, kept from the moment app_control launches it.
Its own system log lines (its subsystem, or logged by its own code), streamed from launch. macOS keeps info lines only in memory and drops them within seconds when the Simulator is busy, so reading the log store afterwards misses them.
grep also searches Apple's lines, for things like a missing image or a layout warning.
wait_for waits up to 25 seconds for a line, for example after a tap that starts a network call.
crashes lists crash reports with the crashing thread down to your code ( ItemsView.body.getter (App.swift:64) ), next to the app's last output, which holds Swift's own message ( App.swift:64: Fatal error: ... ). macOS can take up to a minute to write a report.
Each agent runs its own MCP server and holds its own Simulator through device_use , and its own cloud phone: one agent's cloud start never moves another's tools, its cloud stop never ends another's phone, and the phone stops when the agent's server ends. A cloud start typed in a terminal is the command line's own.
Boo

[truncated]

## Original Extract

Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. - NoMac-app/simulator-use

GitHub - NoMac-app/simulator-use: Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line. · GitHub
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
NoMac-app
simulator-use Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Simulator use for agents: read and drive the iOS Simulator from any MCP client or the command line.
www.npmjs.com/package/simulator-use Topics
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
85 Commits 85 Commits Folders and files
.github/ workflows .github/ workflows bin bin examples examples native native scripts scripts skill skill src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md THIRD_PARTY_LICENSES THIRD_PARTY_LICENSES action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
The fastest, most token-efficient way for an AI agent to use the iOS Simulator. simulator-use gives Claude Code, Codex, Cursor and any other MCP client a phone: it builds and runs your Xcode project, reads the screen as a few lines of text, taps and types by label, and runs your tests case by case. Its own native engine talks to the Simulator directly, so reading the whole screen takes about 250 ms and a tap goes out in under a millisecond, where tools that start a helper for every call take 1.5 to 3 seconds. Free and open source (MIT), on any Mac with Xcode, and from Linux, Windows or a Mac without Xcode through a cloud iPhone.
Fast: a whole-screen read in about 250 ms, a tap sent in under 1 ms, a screenshot in about 60 ms (measured on an M6 Mac mini, October 2026).
Token-efficient: one line per element the agent can use, nothing else; a dozen tools with short descriptions.
Complete: build, install, launch, XCTest and XCUITest per case, logs and crashes, video, permission prompts, iPhone and iPad.
Continuously improved: weekly releases driven by thousands of daily sessions.
Anywhere: npx simulator-use cloud start gives a cloud iPhone with Xcode to any machine.
It builds, runs and tests Xcode projects, reads and uses the screen of any app (SwiftUI, UIKit, React Native), answers iOS's permission prompts, reads the app's logs and crashes, and records video.
Demo · Sign In · iPhone 17 Pro (iOS 26.5)
textfield "Email" #email @201,194
securefield "Password" #password @201,246
switch "Remember me" #remember off @341,298
button "Sign In" #signIn @201,350
text "Not signed in" #status @201,402
tabs: Login (selected) @158,822, Items @244,822
One line per element: kind, label, #id (the app's accessibility identifier), state and @x,y , the center in points.
Groups, icons inside buttons, scroll bars, the status bar and the keyboard's keys are folded away. The tab bar, scrolling and the keyboard are one line each.
Only what can be tapped is listed. An alert, a sheet, a menu or a full-screen view in front hides the screen behind it, and iOS's own permission prompts read as alerts over the app. Rows under the navigation bar, the tab bar or the keyboard are counted as hidden (above, below, to the left or right), not listed.
A web view in UIKit (Capacitor, React Native's WebView, a sign-in page inside most apps) is read point by point, and its headings, texts, fields, checkboxes, buttons and links are listed like the rest. Web views hosted by SwiftUI, Safari's own pages, in-app browsers, maps, games and iOS's out-of-process pickers have nothing for accessibility to read. They come back as one unreadable line with their place, so the agent knows to take a screenshot and tap by point.
A screenshot, when asked for, is sized to points, so a point in the picture is a point to tap.
The sign-in screen above is about 70 tokens. The full accessibility tree it comes from is about 40,000.
ui_act takes steps in the same words ui_look prints, runs them in order, and answers with what changed:
ui_act steps: ["fill #email gary@example.com", "fill #password secret1", "tap #signIn", "wait \"Welcome\""]
4 steps done in 4.5 s
changed: #email "gary@example.com" · #password "•••••••" · #status "Welcome, gary@example.com"
Steps:
Every step is checked, so nothing succeeds silently:
Typed text is read back from the field. When iOS's autocorrect or smart punctuation changes it ( -- to — , ' to ’ ), the text is pasted instead, which iOS takes as is. A field that still shows something else stops the call and says what it shows.
Text in any language (Café, 日本語, emoji) can be typed: what a US keyboard lacks goes through the Simulator's clipboard.
fill refuses anything that isn't a text field, so keys never land in whichever field still has focus.
A tap that changes nothing is reported, after waiting 2.5 s for late answers such as a permission prompt. A tap inside an area that can't be read says it can't be checked.
A target that isn't on screen yet gets two seconds to appear, then fails with the next step: "2 elements are below: scroll down", or, when an alert came up in front, "Answer it first: tap "Cancel" or "Open"".
The first failing step stops the rest and returns the screen as it is now.
A new screen comes back whole; otherwise only what changed.
Tool
CLI
Does
app_run
simulator-use run
Builds the Xcode project here, installs and launches it, answers with the first screen; errors as file:line ; --xcode 26.6 builds with another installed Xcode
devices
simulator-use devices
Lists the Simulators on this Mac with their runtime, UDID and state, and its Xcodes
device_use
simulator-use use "iPhone 17 Pro"
Chooses the Simulator the other tools use, booting it (one boot at a time); dark mode, a clean status bar
app_control
simulator-use app launch dev.example.app
install , launch (answers with the first screen), terminate , uninstall , open_url ; grant , revoke and reset_permissions ; location ; push ; the phone's appearance , locale , timezone , text_size and a clean status_bar
ui_look
simulator-use look
Reads the screen; find keeps matching lines, screenshot adds a picture, full saves the whole tree
ui_act
simulator-use act "tap #signIn" ...
Does steps in order and answers with what changed
app_logs
simulator-use logs
The app's print output and its own system log lines; grep , wait_for , crashes
tests_run
simulator-use test
Runs the unit and UI tests; each case as it finishes, counts, each failure as file:line and the slowest; "0 tests ran" is an error
screen_record
simulator-use record start
Starts or stops a video of the screen; stop answers with the MP4
job_status
simulator-use job
Waits up to 40 s for a long build and answers with its result, or its progress
job_cancel
simulator-use cancel
Stops a job
cloud
simulator-use cloud start
start , status , stop a NoMac cloud phone that every tool then runs on; run builds, tests and stops in one call (see Cloud)
The MCP server and the CLI are built from one list, so they always match. The tool list and instructions stay under 4,000 tokens; a test enforces it.
app_run does the whole loop in one call: it finds the project (a workspace first, then a project, an XcodeGen spec, or ios/ ), picks the scheme, builds for the Simulator, installs, launches, and answers with the first screen.
The app's permissions are granted when it is installed and its first launch's prompts (notifications too) are answered, so no prompt stops the agent; permissions: "ask" keeps iOS's prompts, to test them. open_url answers iOS's "Open in …?" for a link that opens another app.
Built Demo in 4.2 s, installed dev.simtap.demo.
Launched Demo in 3.9 s.
Demo · Sign In · iPhone 17 Pro (iOS 26.5)
textfield "Email" #email @201,194
...
A failed build answers with the errors only, as Sources/App.swift:49:37: error: ... , and the path to the full log.
Each folder (each git worktree) builds into its own DerivedData, so rebuilds are incremental and worktrees never share products. SIMULATOR_USE_CACHE_DIR puts it on a disk you keep, so the next machine's build is incremental too.
An XcodeGen project is regenerated from project.yml before each build (skipped when nothing changed), so new files are always in it.
The next call in the same folder needs no arguments.
env: "API_URL=http://127.0.0.1:8080 MODE=test" gives the app those variables (it reads them from ProcessInfo ), as an Xcode scheme does; they stay for its later launches, and env: "" clears them.
An MCP call answers within 40 s; a longer build answers with a job, and job_status waits for it. From the CLI, simulator-use run simply waits.
Two builds of one project never run at once, also from two processes: the second says which process is building it.
Expo, React Native, Flutter and Capacitor builds are not done yet; build them with their own tools, then app_control install .
tests_run runs the scheme's tests and reads the answer from Xcode's result bundle, not from the scrolling output:
Tests failed: 1 passed, 2 failed (3 in 35.3 s). Result bundle: .../tests-1792370825323.xcresult
DemoTests/SignInTests/testRejectsAShortPassword(): Tests/SignInTests.swift:12: XCTAssertEqual failed: ("Welcome, gary@example.com") is not equal to ("Invalid email or password")
DemoTests/rejectsAnEmailWithoutAt(): Tests/SignInTests.swift:18: Expectation failed: ...
XCTest and Swift Testing both.
Each failing test is named the way only takes it, to rerun just that one.
"0 tests ran" (no test target, or a filter that matched nothing) is an error, never a pass.
app_logs shows what the app itself said, not the thousands of lines Apple's frameworks log from inside it:
Its print() output, kept from the moment app_control launches it.
Its own system log lines (its subsystem, or logged by its own code), streamed from launch. macOS keeps info lines only in memory and drops them within seconds when the Simulator is busy, so reading the log store afterwards misses them.
grep also searches Apple's lines, for things like a missing image or a layout warning.
wait_for waits up to 25 seconds for a line, for example after a tap that starts a network call.
crashes lists crash reports with the crashing thread down to your code ( ItemsView.body.getter (App.swift:64) ), next to the app's last output, which holds Swift's own message ( App.swift:64: Fatal error: ... ). macOS can take up to a minute to write a report.
Each agent runs its own MCP server and holds its own Simulator through device_use , and its own cloud phone: one agent's cloud start never moves another's tools, its cloud stop never ends another's phone, and the phone stops when the agent's server ends. A cloud start typed in a terminal is the command line's own.
Boo

[truncated]
