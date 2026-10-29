---
source: "https://flowdeck.studio/blog/run-multiple-claude-code-agents-xcode-project"
hn_url: "https://news.ycombinator.com/item?id=50296034"
title: "How to Run Multiple Claude Code Agents on the Same Xcode Project"
article_title: "How to Run Multiple Claude Code Agents on the Same Xcode Project | FlowDeck"
image: "https://framerusercontent.com/images/Adigt5w6x8kGHyGTiUSKlsNjMA.png?width=1200&height=630"
author: "afterxleep"
captured_at: "2026-10-29T21:34:15Z"
capture_tool: "hn-digest"
hn_id: 50296034
score: 1
comments: 0
posted_at: "2026-10-29T21:09:16Z"
tags:
  - hacker-news
---

# How to Run Multiple Claude Code Agents on the Same Xcode Project

- HN: [50296034](https://news.ycombinator.com/item?id=50296034)
- Source: [flowdeck.studio](https://flowdeck.studio/blog/run-multiple-claude-code-agents-xcode-project)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T21:09:16Z

## Translation

Title: How to Run Multiple Claude Code Agents on the Same Xcode Project
Article title: How to Run Multiple Claude Code Agents on the Same Xcode Project | FlowDeck
Description: Learn how to run multiple Claude Code agents on the same Xcode project without simulator conflicts.

Article text:
How to Run Multiple Claude Code Agents on the Same Xcode Project
Learn how to run multiple Claude Code agents on the same Xcode project without simulator conflicts.
Running two or more Claude Code agents on one Xcode project is the natural next step once a single agent works well. The expectation is two tasks in roughly the time of one.
Where it goes wrong is not in the code. An iOS project shares more than files: the simulator, the installed app, and the data the app saved. I wanted to see what happens when two agents reach that shared part.
I gave two Claude Code agents separate tasks on a small SwiftUI weather app, each in its own Git worktree. One added a feature. The other added a test. Neither prompt named a simulator.
Both agents picked the iPhone 18 Pro that was already booted. Both builds succeeded, and both reported success.
Then I checked what was installed. The simulator was running the second agent's build, and the first agent's build had been replaced. Neither agent's output mentioned the other, or that its build was gone.
Worktrees keep files apart. They don't keep simulators apart.
If two sessions install the same app on one simulator, they share the same build, app data, and screen. To see it, I built the app from two worktrees with different build numbers ( 111 and 222 ) and the same bundle ID. The first worktree branch (let's call it Agent A) installed and launched its build, then the second (Agent B) installed its own, both on one simulator. The commands and results are specific to this example app, and the full run is in this gist:
https://gist.github.com/amarildolucas/257229e7f3eeff3de073614a934e4b6a
Four things happened, and none of them raised an error:
The simulator held build 222 . Agent A’s build 111 was gone.
The app that A had launched was no longer running. The log shows the install sending a terminate request to the running app.
The cities that build 111 saved in UserDefaults were still there for build 222. The log shows No saved cities found; seeding defaults (2) on first launch and loadCities loaded 2 cities from UserDefaults afterward. Installing over an existing app keeps its data; only uninstalling removes it.
When both agents launched their app afterward, simctl returned the same process ID . Agent A was now running B’s build, and reading the installed build number gave 222 , not the 111 it had built.
What real Claude Code agents do
The manual test forces the collision. Real agents don't need forcing. I started one session per worktree with claude --worktree and gave each agent an ordinary task, the way you would without thinking about simulators.
Add a cloud-cover line to the humidity tile in
Add a cloud-cover line to the humidity tile in
Add a cloud-cover line to the humidity tile in
Agent B:
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Both agents picked the booted iPhone 18 Pro and installed on it in both runs I did.
To see the consequence, I made the overlap certain. I stamped Agent A's build 111 and Agent B's build 222 , added "after launching, wait 90 seconds, then read the installed build number" to each prompt, and pointed both agents at the same simulator. Both builds succeeded. Only the check told them apart:
Agent A's installed build was 222 , not 111 , and its process had changed. Agent B saw nothing wrong because it had installed last. Without the check, nothing in Agent A's output would have said so.
Run one agent at a time. There is no collision, but the second task waits for the first, which defeats the point of running two.
Give each agent its own simulator , set up by hand. It works, and the next section shows how.
Use a tool that keeps each session's simulator , worktree, and build state together. The last section shows how it does.
The fix by hand: one simulator per agent
Clone one simulator per agent, from a base simulator that is shut down. xcrun simctl clone prints each new UDID, and xcrun simctl list devices lists them all.
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
Tell each agent which simulator is its own. Clones alone change nothing, because both agents took whichever simulator was already booted. Add this to each prompt (Agent A for the first prompt and B for the second):
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Pass the same UDID to every build and test command:
xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath I ran the same two prompts again with those lines and the 90-second check. Agent A read 111 and Agent B read 222 , and neither process had changed:
This works, but you have to repeat the setup.
Let FlowDeck manage the simulator
Every new session needs a clone, a UDID, and a prompt that carries it. A mismatched UDID sends both agents back to the same simulator, and a green run will not tell you.
FlowDeck Studio moves that setup into the session, and each has its own branch, Git worktree, project configuration (scheme, destination, and build configuration), and separate DerivedData folder. The destination is part of the session, and you set it once in the app.
Here are two sessions on the same project, one per task. One works on the cloud-cover line, the other on the unknown weather codes. Session "Cloud Cover in the Humidity Card" runs on Agent A Simulator
Session "Unknown weather codes show as Clear" runs on Agent B Simulator
Each session shows the simulator it runs on, so a shared simulator is visible before it causes a bad test run.
What made this collision easy to miss is that nothing failed. Both builds succeeded, both agents reported success, and the only giveaway was a build number that read 222 instead of 111 .
A passing build tells you the code compiled. It does not tell you whose code is installed. If you run agents in parallel, you want the work done in roughly half the time, while ensuring that each build runs the correct code and produces the right results.
Cloning a simulator and carrying its UDID into every prompt was the part I had to repeat for each session, and it is massive, repetitive work. FlowDeck Studio makes the simulator part of the session setup, so you can get straight to work and focus on the feature rather than on customizing your setup.
Does this only happen with Claude Code?
Do I need to change my existing project to use FlowDeck?
How many simulators can I run at once before my Mac runs out of space?
Why your AI agent keeps failing at iOS development
MCP vs CLI for AI-powered iOS development
I’ve built scalable native iOS apps for over 12 years. I’m currently focused on learning and driving innovation at the intersection of AI-assisted software development.
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring

## Original Extract

Learn how to run multiple Claude Code agents on the same Xcode project without simulator conflicts.

How to Run Multiple Claude Code Agents on the Same Xcode Project
Learn how to run multiple Claude Code agents on the same Xcode project without simulator conflicts.
Running two or more Claude Code agents on one Xcode project is the natural next step once a single agent works well. The expectation is two tasks in roughly the time of one.
Where it goes wrong is not in the code. An iOS project shares more than files: the simulator, the installed app, and the data the app saved. I wanted to see what happens when two agents reach that shared part.
I gave two Claude Code agents separate tasks on a small SwiftUI weather app, each in its own Git worktree. One added a feature. The other added a test. Neither prompt named a simulator.
Both agents picked the iPhone 18 Pro that was already booted. Both builds succeeded, and both reported success.
Then I checked what was installed. The simulator was running the second agent's build, and the first agent's build had been replaced. Neither agent's output mentioned the other, or that its build was gone.
Worktrees keep files apart. They don't keep simulators apart.
If two sessions install the same app on one simulator, they share the same build, app data, and screen. To see it, I built the app from two worktrees with different build numbers ( 111 and 222 ) and the same bundle ID. The first worktree branch (let's call it Agent A) installed and launched its build, then the second (Agent B) installed its own, both on one simulator. The commands and results are specific to this example app, and the full run is in this gist:
https://gist.github.com/amarildolucas/257229e7f3eeff3de073614a934e4b6a
Four things happened, and none of them raised an error:
The simulator held build 222 . Agent A’s build 111 was gone.
The app that A had launched was no longer running. The log shows the install sending a terminate request to the running app.
The cities that build 111 saved in UserDefaults were still there for build 222. The log shows No saved cities found; seeding defaults (2) on first launch and loadCities loaded 2 cities from UserDefaults afterward. Installing over an existing app keeps its data; only uninstalling removes it.
When both agents launched their app afterward, simctl returned the same process ID . Agent A was now running B’s build, and reading the installed build number gave 222 , not the 111 it had built.
What real Claude Code agents do
The manual test forces the collision. Real agents don't need forcing. I started one session per worktree with claude --worktree and gave each agent an ordinary task, the way you would without thinking about simulators.
Add a cloud-cover line to the humidity tile in
Add a cloud-cover line to the humidity tile in
Add a cloud-cover line to the humidity tile in
Agent B:
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Add a test in IconMappingSwiftTests.swift for unknown weather codes.
Change WeatherIcon.swift only if
Both agents picked the booted iPhone 18 Pro and installed on it in both runs I did.
To see the consequence, I made the overlap certain. I stamped Agent A's build 111 and Agent B's build 222 , added "after launching, wait 90 seconds, then read the installed build number" to each prompt, and pointed both agents at the same simulator. Both builds succeeded. Only the check told them apart:
Agent A's installed build was 222 , not 111 , and its process had changed. Agent B saw nothing wrong because it had installed last. Without the check, nothing in Agent A's output would have said so.
Run one agent at a time. There is no collision, but the second task waits for the first, which defeats the point of running two.
Give each agent its own simulator , set up by hand. It works, and the next section shows how.
Use a tool that keeps each session's simulator , worktree, and build state together. The last section shows how it does.
The fix by hand: one simulator per agent
Clone one simulator per agent, from a base simulator that is shut down. xcrun simctl clone prints each new UDID, and xcrun simctl list devices lists them all.
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
xcrun simctl shutdown <base-udid>
xcrun simctl clone <base-udid> "Agent A"
xcrun simctl clone <base-udid> "Agent B"
Tell each agent which simulator is its own. Clones alone change nothing, because both agents took whichever simulator was already booted. Add this to each prompt (Agent A for the first prompt and B for the second):
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Use only the simulator named "Agent A" (UDID <agent-a-udid>).
Do not touch
Pass the same UDID to every build and test command:
xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath xcodebuild test -project WeatherApp.xcodeproj -scheme "Weather App" -destination "platform=iOS Simulator,id=<agent-a-udid>" -derivedDataPath I ran the same two prompts again with those lines and the 90-second check. Agent A read 111 and Agent B read 222 , and neither process had changed:
This works, but you have to repeat the setup.
Let FlowDeck manage the simulator
Every new session needs a clone, a UDID, and a prompt that carries it. A mismatched UDID sends both agents back to the same simulator, and a green run will not tell you.
FlowDeck Studio moves that setup into the session, and each has its own branch, Git worktree, project configuration (scheme, destination, and build configuration), and separate DerivedData folder. The destination is part of the session, and you set it once in the app.
Here are two sessions on the same project, one per task. One works on the cloud-cover line, the other on the unknown weather codes. Session "Cloud Cover in the Humidity Card" runs on Agent A Simulator
Session "Unknown weather codes show as Clear" runs on Agent B Simulator
Each session shows the simulator it runs on, so a shared simulator is visible before it causes a bad test run.
What made this collision easy to miss is that nothing failed. Both builds succeeded, both agents reported success, and the only giveaway was a build number that read 222 instead of 111 .
A passing build tells you the code compiled. It does not tell you whose code is installed. If you run agents in parallel, you want the work done in roughly half the time, while ensuring that each build runs the correct code and produces the right results.
Cloning a simulator and carrying its UDID into every prompt was the part I had to repeat for each session, and it is massive, repetitive work. FlowDeck Studio makes the simulator part of the session setup, so you can get straight to work and focus on the feature rather than on customizing your setup.
Does this only happen with Claude Code?
Do I need to change my existing project to use FlowDeck?
How many simulators can I run at once before my Mac runs out of space?
Why your AI agent keeps failing at iOS development
MCP vs CLI for AI-powered iOS development
I’ve built scalable native iOS apps for over 12 years. I’m currently focused on learning and driving innovation at the intersection of AI-assisted software development.
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring
Install FlowDeck, open your project and give your agent a task.
Native diff viewer and PR review
Github Integration & PR monitoring
