---
source: "https://github.com/franzenzenhofer/big-arrow-on-the-screen"
hn_url: "https://news.ycombinator.com/item?id=50004580"
title: "Show HN: Let your AI agents paint big arrows, boxes and text on your screen"
article_title: "GitHub - franzenzenhofer/big-arrow-on-the-screen: Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. · GitHub"
image: "https://opengraph.githubassets.com/68d9566b71dfd6eac473f88e9880e60e50a0e87cfcd5ec7d9ab8d20ab13f35a0/franzenzenhofer/big-arrow-on-the-screen"
author: "franze"
captured_at: "2026-10-08T11:48:23Z"
capture_tool: "hn-digest"
hn_id: 50004580
score: 1
comments: 0
posted_at: "2026-10-08T11:33:54Z"
tags:
  - hacker-news
---

# Show HN: Let your AI agents paint big arrows, boxes and text on your screen

- HN: [50004580](https://news.ycombinator.com/item?id=50004580)
- Source: [github.com](https://github.com/franzenzenhofer/big-arrow-on-the-screen)
- Score: 1
- Comments: 0
- Posted: 2026-10-08T11:33:54Z

## Translation

Title: Show HN: Let your AI agents paint big arrows, boxes and text on your screen
Article title: GitHub - franzenzenhofer/big-arrow-on-the-screen: Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. · GitHub
Description: Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. - franzenzenhofer/big-arrow-on-the-screen

Article text:
GitHub - franzenzenhofer/big-arrow-on-the-screen: Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. · GitHub
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
franzenzenhofer
big-arrow-on-the-screen Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
50 Commits 50 Commits Folders and files
.github/ workflows .github/ workflows Sources Sources Tests Tests docs docs scripts scripts skill/ big-arrow skill/ big-arrow .gitignore .gitignore .swiftlint.yml .swiftlint.yml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md View all files Repository files navigation
Let your AI agents paint big arrows, boxes and text on your screen
big-arrow-on-the-screen ( bigarrow ): one small macOS CLI and an agent skill. Click-through, never steals the focus, gone by itself. MIT.
Your AI agent can refactor a monorepo, write a migration and explain monads, but when it needs you to click one button it prints "please click Allow in the dialog" into a terminal you are not looking at. bigarrow gives it a finger.
bigarrow point --element " Allow " --app " System Settings " --text " Franz, click Allow "
A big, friendly arrow with a sign appears on top of everything, points at the thing, and goes away again. It is click-through, it never steals your focus, it works on every display and every Space, and it needs no permission at all to draw. It is one small Swift binary. There is no daemon, no menu-bar icon, no account, no telemetry, and, we checked twice, no AI inside. It is an arrow.
Fair question. Arrows have existed since roughly the Paleolithic. Here is what changed: software agents now do real work on your Mac, and they keep hitting the same wall, the part that only a human may do.
"Click Allow." macOS permission prompts, OAuth consent screens, "Open with...?" dialogs. The agent can find the button but must not, or cannot, press it for you. It can now point at it.
"Your turn." 2FA codes, CAPTCHAs, passkeys, a payment confirmation, a signature, a legal checkbox. The things an agent should never click on its own behalf. It points, you decide, it continues.
"It's this window, not that one." You have 14 Chrome windows. The agent knows which one it means: --window "Google Chrome:Pull request" --raise .
"I need you, and you're making coffee." --say reads the sign aloud. Your Mac will literally call you back to your desk.
Guided setups and onboarding. Walk a human through a settings pane step by step: start , wait until they acted, stop , next step. Like a product tour, minus the product.
Remote help. "No, the other gear icon." Point at it instead of describing it.
Demos, screencasts, docs. Highlight what matters while recording, or render the arrow straight into a PNG with --png for documentation.
Debugging coordinates. Not sure your Accessibility, screenshot or Peekaboo coordinates are right? Point at them and look. --dry-run --json tells you where it would point without drawing.
Situations we have all been in
Staged with a neutral demo dialog and recorded with the real bigarrow on a test Mac ( scripts/funny-scenes.sh ). The dialogs are fake. The feelings are real.
What it is not: a screen annotator for humans, a click bot, or a screenshot tool. It never clicks, types or captures anything. It only points. Deliberately.
brew install franzenzenhofer/tap/bigarrow
bigarrow install-skill # teaches Claude Code (~/.claude/skills) and Codex (~/.agents/skills)
From source: swift build -c release (Xcode 16 or newer, macOS 14 or newer), binary at .build/release/bigarrow .
The three commands an agent needs
bigarrow point --element " Allow " --app " System Settings " --text " Franz, click Allow " # by label
bigarrow point --at 760,500 --text " Franz, click HERE " # by coordinate
bigarrow start --window " Safari:Inbox " --raise --text " This window " && bigarrow stop # until stopped
Three ways an arrow ends, pick your level of commitment:
Targets: --at X,Y , --rect X,Y,W,H , --mouse , --window App[:title] , --element Label --app App , --peekaboo ID --snapshot see.json (from Peekaboo's see --json ). Coordinates are global top-left logical points, the space Accessibility, CGWindowList and Peekaboo report. --display N makes --at and --rect relative to one display.
bigarrow front --app X or point --raise brings the target's app to the front first, because pointing at a window hidden behind your terminal is a special kind of unhelpful. bigarrow elements --app X lists what --element can match. bigarrow doctor shows permissions, who owns them, and your displays.
Every command takes --json . Exit codes: 0 ok, 2 bad input, 3 target not found, 4 permission missing. Agents love exit codes. Humans tolerate them.
It is an arrow, so we spent an unreasonable amount of time on how it looks.
--shape bend|straight|zigzag (zigzag for when it is really urgent)
--style arrow|ring|box ; rings and boxes are border-only, so you still see what is under them
--size S|M|L , --corners round|sharp
--color red|orange|yellow|green|teal|blue|purple|pink|black|white|#RRGGBB ; light colours automatically get a dark outline and text
--follow moves with a window or element, --until-click ends on a click on the target, --say speaks the sign
Several arrows at once keep their signs out of each other's way (the HN shot above is five independent bigarrow start calls)
The shaft grows out of the sign through a flared joint that never runs into a rounded corner. scripts/gallery.py renders every combination offscreen and zooms into every joint ( junctions ), because a seam at the joint was, apparently, unacceptable.
Does it need Screen Recording or Accessibility?
Drawing needs neither. --element , elements , --until-click and front --window use Accessibility, which macOS grants to the app that runs your shell (Terminal, iTerm2, Ghostty, VS Code, Claude), never to bigarrow itself. bigarrow doctor names that app, and exit code 4 tells the agent exactly what to ask you for. --window App:title reads window titles, which macOS 26 hides without Screen Recording; --window App alone needs nothing.
Will it steal my focus while I'm typing?
No. That was the hardest bug in the project: NSApplication.run() quietly activates a process that has no terminal, so detached arrows grabbed the focus. bigarrow pumps events itself instead, and the tests check that the frontmost app never changes.
Can I click through it?
Yes, everywhere except the optional X, which is its own tiny panel that also never takes the focus.
Multiple displays? Full-screen apps? Stage Manager? Spaces?
Yes, yes, yes, yes. Displays left of or above the main one (negative coordinates) included. Unplug a display while an arrow is on it and the arrow politely leaves. See the verification matrix .
How much CPU does a pulsing arrow cost?
1.4 % measured on a CI runner. Core Animation does the work in the render server.
Does --element work inside web pages?
In Electron apps, yes. In Chrome, only when Chrome runs with --force-renderer-accessibility (or VoiceOver is on); Chrome ignores the usual request to expose page content, verified on Chrome in October 2026. Chrome's own toolbar always works. Otherwise point at the page's coordinates, which the skill explains.
Why not just use [some screen annotation app]?
Those are for humans drawing on screens. This is for programs pointing at things, from a shell, with exit codes. Twenty-six tools were checked before writing a line ( research ). None did this.
Is it AI?
No. It is the least intelligent part of your AI stack, and proud of it.
76 automated tests: geometry, placement, joint smoothness, a golden image, recorded window-server, Accessibility and Peekaboo 4.9.0 fixtures, and tests against the real window server (window level 1000, clicks pass through, focus never moves, detach and stop timing). CI runs them on macOS 15; they also passed on macOS 26 and macOS 27.
17 behaviour checks on a clean runner ( visual.yml ): real clicks on the X, --until-click , --follow , --raise , --say , full-screen apps, Stage Manager, a Space switch, a second display, a 2x display, unplugging a display mid-arrow, CPU. The demo GIF above is recorded by the same workflow, on a desktop with nothing personal on it.
A fresh agent given only the skill and "show Franz where the Reload button in Chrome is" found it by label and built the right command ( transcript ). It also found a bug, which is now a test.
For agents (and the humans who configure them)
The skill in skill/big-arrow/ works for both Claude Code and Codex (one SKILL.md , Agent Skills format, plus agents/openai.yaml for Codex). It tells the agent when to point, how to pick a target, to write a full sentence on the sign, to add --say when you are probably not looking, and to stop once you have acted.
docs/plan/PLAN.md (goal, architecture, risks), docs/plan/TICKETS.md (generated from docs/plan/tickets.json ), docs/decisions/ , docs/research/ (verified facts with links), docs/verification/ , docs/skill-tests/ , CHANGELOG.md .
Peekaboo's visualizer ( https://github.com/openclaw/Peekaboo ) and Nameplate ( https://github.com/steipete/Nameplate ) by Peter Steinberger showed the overlay window recipe and the agent-skill packaging. Neither draws a pointing arrow with a label, which is the gap this project fills. bigarrow reads Peekaboo's see --json as an optional target source.
Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. - franzenzenhofer/big-arrow-on-the-screen

GitHub - franzenzenhofer/big-arrow-on-the-screen: Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT. · GitHub
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
franzenzenhofer
big-arrow-on-the-screen Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
50 Commits 50 Commits Folders and files
.github/ workflows .github/ workflows Sources Sources Tests Tests docs docs scripts scripts skill/ big-arrow skill/ big-arrow .gitignore .gitignore .swiftlint.yml .swiftlint.yml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md View all files Repository files navigation
Let your AI agents paint big arrows, boxes and text on your screen
big-arrow-on-the-screen ( bigarrow ): one small macOS CLI and an agent skill. Click-through, never steals the focus, gone by itself. MIT.
Your AI agent can refactor a monorepo, write a migration and explain monads, but when it needs you to click one button it prints "please click Allow in the dialog" into a terminal you are not looking at. bigarrow gives it a finger.
bigarrow point --element " Allow " --app " System Settings " --text " Franz, click Allow "
A big, friendly arrow with a sign appears on top of everything, points at the thing, and goes away again. It is click-through, it never steals your focus, it works on every display and every Space, and it needs no permission at all to draw. It is one small Swift binary. There is no daemon, no menu-bar icon, no account, no telemetry, and, we checked twice, no AI inside. It is an arrow.
Fair question. Arrows have existed since roughly the Paleolithic. Here is what changed: software agents now do real work on your Mac, and they keep hitting the same wall, the part that only a human may do.
"Click Allow." macOS permission prompts, OAuth consent screens, "Open with...?" dialogs. The agent can find the button but must not, or cannot, press it for you. It can now point at it.
"Your turn." 2FA codes, CAPTCHAs, passkeys, a payment confirmation, a signature, a legal checkbox. The things an agent should never click on its own behalf. It points, you decide, it continues.
"It's this window, not that one." You have 14 Chrome windows. The agent knows which one it means: --window "Google Chrome:Pull request" --raise .
"I need you, and you're making coffee." --say reads the sign aloud. Your Mac will literally call you back to your desk.
Guided setups and onboarding. Walk a human through a settings pane step by step: start , wait until they acted, stop , next step. Like a product tour, minus the product.
Remote help. "No, the other gear icon." Point at it instead of describing it.
Demos, screencasts, docs. Highlight what matters while recording, or render the arrow straight into a PNG with --png for documentation.
Debugging coordinates. Not sure your Accessibility, screenshot or Peekaboo coordinates are right? Point at them and look. --dry-run --json tells you where it would point without drawing.
Situations we have all been in
Staged with a neutral demo dialog and recorded with the real bigarrow on a test Mac ( scripts/funny-scenes.sh ). The dialogs are fake. The feelings are real.
What it is not: a screen annotator for humans, a click bot, or a screenshot tool. It never clicks, types or captures anything. It only points. Deliberately.
brew install franzenzenhofer/tap/bigarrow
bigarrow install-skill # teaches Claude Code (~/.claude/skills) and Codex (~/.agents/skills)
From source: swift build -c release (Xcode 16 or newer, macOS 14 or newer), binary at .build/release/bigarrow .
The three commands an agent needs
bigarrow point --element " Allow " --app " System Settings " --text " Franz, click Allow " # by label
bigarrow point --at 760,500 --text " Franz, click HERE " # by coordinate
bigarrow start --window " Safari:Inbox " --raise --text " This window " && bigarrow stop # until stopped
Three ways an arrow ends, pick your level of commitment:
Targets: --at X,Y , --rect X,Y,W,H , --mouse , --window App[:title] , --element Label --app App , --peekaboo ID --snapshot see.json (from Peekaboo's see --json ). Coordinates are global top-left logical points, the space Accessibility, CGWindowList and Peekaboo report. --display N makes --at and --rect relative to one display.
bigarrow front --app X or point --raise brings the target's app to the front first, because pointing at a window hidden behind your terminal is a special kind of unhelpful. bigarrow elements --app X lists what --element can match. bigarrow doctor shows permissions, who owns them, and your displays.
Every command takes --json . Exit codes: 0 ok, 2 bad input, 3 target not found, 4 permission missing. Agents love exit codes. Humans tolerate them.
It is an arrow, so we spent an unreasonable amount of time on how it looks.
--shape bend|straight|zigzag (zigzag for when it is really urgent)
--style arrow|ring|box ; rings and boxes are border-only, so you still see what is under them
--size S|M|L , --corners round|sharp
--color red|orange|yellow|green|teal|blue|purple|pink|black|white|#RRGGBB ; light colours automatically get a dark outline and text
--follow moves with a window or element, --until-click ends on a click on the target, --say speaks the sign
Several arrows at once keep their signs out of each other's way (the HN shot above is five independent bigarrow start calls)
The shaft grows out of the sign through a flared joint that never runs into a rounded corner. scripts/gallery.py renders every combination offscreen and zooms into every joint ( junctions ), because a seam at the joint was, apparently, unacceptable.
Does it need Screen Recording or Accessibility?
Drawing needs neither. --element , elements , --until-click and front --window use Accessibility, which macOS grants to the app that runs your shell (Terminal, iTerm2, Ghostty, VS Code, Claude), never to bigarrow itself. bigarrow doctor names that app, and exit code 4 tells the agent exactly what to ask you for. --window App:title reads window titles, which macOS 26 hides without Screen Recording; --window App alone needs nothing.
Will it steal my focus while I'm typing?
No. That was the hardest bug in the project: NSApplication.run() quietly activates a process that has no terminal, so detached arrows grabbed the focus. bigarrow pumps events itself instead, and the tests check that the frontmost app never changes.
Can I click through it?
Yes, everywhere except the optional X, which is its own tiny panel that also never takes the focus.
Multiple displays? Full-screen apps? Stage Manager? Spaces?
Yes, yes, yes, yes. Displays left of or above the main one (negative coordinates) included. Unplug a display while an arrow is on it and the arrow politely leaves. See the verification matrix .
How much CPU does a pulsing arrow cost?
1.4 % measured on a CI runner. Core Animation does the work in the render server.
Does --element work inside web pages?
In Electron apps, yes. In Chrome, only when Chrome runs with --force-renderer-accessibility (or VoiceOver is on); Chrome ignores the usual request to expose page content, verified on Chrome in October 2026. Chrome's own toolbar always works. Otherwise point at the page's coordinates, which the skill explains.
Why not just use [some screen annotation app]?
Those are for humans drawing on screens. This is for programs pointing at things, from a shell, with exit codes. Twenty-six tools were checked before writing a line ( research ). None did this.
Is it AI?
No. It is the least intelligent part of your AI stack, and proud of it.
76 automated tests: geometry, placement, joint smoothness, a golden image, recorded window-server, Accessibility and Peekaboo 4.9.0 fixtures, and tests against the real window server (window level 1000, clicks pass through, focus never moves, detach and stop timing). CI runs them on macOS 15; they also passed on macOS 26 and macOS 27.
17 behaviour checks on a clean runner ( visual.yml ): real clicks on the X, --until-click , --follow , --raise , --say , full-screen apps, Stage Manager, a Space switch, a second display, a 2x display, unplugging a display mid-arrow, CPU. The demo GIF above is recorded by the same workflow, on a desktop with nothing personal on it.
A fresh agent given only the skill and "show Franz where the Reload button in Chrome is" found it by label and built the right command ( transcript ). It also found a bug, which is now a test.
For agents (and the humans who configure them)
The skill in skill/big-arrow/ works for both Claude Code and Codex (one SKILL.md , Agent Skills format, plus agents/openai.yaml for Codex). It tells the agent when to point, how to pick a target, to write a full sentence on the sign, to add --say when you are probably not looking, and to stop once you have acted.
docs/plan/PLAN.md (goal, architecture, risks), docs/plan/TICKETS.md (generated from docs/plan/tickets.json ), docs/decisions/ , docs/research/ (verified facts with links), docs/verification/ , docs/skill-tests/ , CHANGELOG.md .
Peekaboo's visualizer ( https://github.com/openclaw/Peekaboo ) and Nameplate ( https://github.com/steipete/Nameplate ) by Peter Steinberger showed the overlay window recipe and the agent-skill packaging. Neither draws a pointing arrow with a label, which is the gap this project fills. bigarrow reads Peekaboo's see --json as an optional target source.
Let your AI agents paint big arrows, boxes and text on your Mac screen. One CLI, click-through, gone by itself. Skill for Claude Code and Codex. MIT.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
