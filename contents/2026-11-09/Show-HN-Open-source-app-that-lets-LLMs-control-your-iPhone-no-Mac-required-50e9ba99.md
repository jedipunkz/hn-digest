---
source: "https://github.com/EliotAndres/iphone-use-app"
hn_url: "https://news.ycombinator.com/item?id=50430860"
title: "Show HN: Open-source app that lets LLMs control your iPhone (no Mac required)"
article_title: "GitHub - EliotAndres/iphone-use-app: An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) · GitHub"
image: "https://opengraph.githubassets.com/c8b891e55edf390bbe162719836b3bb3ba0cab5f6a6ba9ca156709cb9e4226c3/EliotAndres/iphone-use-app"
author: "ea016"
captured_at: "2026-11-09T20:28:27Z"
capture_tool: "hn-digest"
hn_id: 50430860
score: 2
comments: 0
posted_at: "2026-11-09T19:47:29Z"
tags:
  - hacker-news
---

# Show HN: Open-source app that lets LLMs control your iPhone (no Mac required)

- HN: [50430860](https://news.ycombinator.com/item?id=50430860)
- Source: [github.com](https://github.com/EliotAndres/iphone-use-app)
- Score: 2
- Comments: 0
- Posted: 2026-11-09T19:47:29Z

## Translation

Title: Show HN: Open-source app that lets LLMs control your iPhone (no Mac required)
Article title: GitHub - EliotAndres/iphone-use-app: An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) · GitHub
Description: An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) - EliotAndres/iphone-use-app

Article text:
GitHub - EliotAndres/iphone-use-app: An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) · GitHub
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
EliotAndres
iphone-use-app Public Notifications You must be signed in to change notification settings
Star 33 ( 33 ) You must be signed in to star a repository
An iPhone app where an LLM controls the phone it runs on (no Mac required after setup)
Readme AGPL-3.0 license Activity Stars
5 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
897 Commits 897 Commits Folders and files
AgentWidget AgentWidget Shared Shared TunnelProvider TunnelProvider docs docs evals evals iPhoneUse.xcodeproj iPhoneUse.xcodeproj iPhoneUse iPhoneUse idevice-patch idevice-patch scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md TODO.md TODO.md source.json source.json View all files Repository files navigation
Let an AI model use your iPhone, from the iPhone itself. No Mac at run time (one-time pairing only), no extra hardware, no jailbreak.
Left: "Select a Lyft to the Eiffel Tower (don't book yet)" (4× speed): the agent leaves its own app, opens Lyft,
types the destination, picks it and stops at the ride options. Middle: the finished run, each step with its screenshot. Right: a new chat.
No computer at run time: the agent runs on the iPhone and controls that same iPhone, no extra hardware or jailbreak.
A Mac is only needed once, to make the pairing file
Start it from anywhere: the chat, Siri, the "Run Agent" Shortcut or an iphoneuse://run?prompt=… link.
iphone-use-app
TapKit
esp32-iphone-agent
iPhone-use
Computer at run time
None (a Mac once, for pairing)
Mac app, phone plugged in
None (a computer once, to flash the board)
Mac (MCP server)
Extra hardware
None
None
ESP32-S3 board (~$7) on USB-C
None
How it sees and acts
CoreDevice screenshots and touches, through a loopback VPN
Apple accessibility features
AirPlay mirroring to the board, USB mouse/keyboard via AssistiveTouch
WebDriverAgent (XCTest) over USB
Requirements
Part
Notes
iPhone
iOS 26 or later, Developer Mode on.
Apple account
Paid team: the app uses its own built-in VPN (needs the Network Extension entitlement).
Free Apple ID: build with BUILTIN_VPN=NO (Xcode build setting or xcodebuild argument) and install LocalDevVPN ; re-sign every 7 days.
Pairing file
Made once from a Mac over USB (below).
Vision model
An OpenRouter key, or any OpenAI-compatible endpoint whose model takes images and tool calls.
Mac
Only to make the pairing file (and to build the app, unless you install the release). Not needed at run time.
Limitations
Shortcut needed on cellular: iOS refuses new loopback VPN connections on cellular data, so the app runs a
Cellular Toggle shortcut that turns data off while it connects, then back on (or turn it off by hand).
No other VPN at the same time: it needs a loopback VPN (the built-in one, or
LocalDevVPN ), and iOS allows one VPN at a time, so
Tailscale and similar stop while it is on.
"Viewed remotely from Mac" notice on every run: iOS shows "Device Hub: This device was viewed remotely from
Mac" when its developer services are used, even though the "Mac" here is the app itself. Nothing leaves the phone (except calls to LLMs).
Permissions: requires Developer Mode, a VPN configuration, Local Network access (to find the phone's
pairing service), and Location set to "Always" (keeps runs alive in the background, alongside silent audio).
Microphone only for voice prompts.
Pairing files expire: make a new one when the app says so.
Sideloaded build only: it needs Developer Mode and developer services, so it can't go on the App Store.
A few seconds per step: most of it is the model.
flowchart LR
subgraph App["iphone-use-app"]
L["LLM calling"]
D["Mount developer image"]
O["Orchestrator"]
end
L <-- "screenshot + history /<br/>tool call" --> R["OpenRouter"]
D -- "DDI mount" --> V["Loopback VPN<br/>10.7.0.1"]
O <-- "screenshot / tap, swipe, type" --> V
Loading
The app mounts a developer image, then uses a loopback VPN to screenshot and control the iPhone it runs on.
Each step, the screenshot goes to the model, which answers with a tool call (tap, long press, swipe, type, open app) until it's done.
A Live Activity shows progress with a Stop button, and every run is saved as a chat.
One decision is one screenshot → one tool call. scripts/eval.sh replays 11 recorded steps (open Lyft, search a
destination, pick a ride, stop before paying; write a note) against each model through OpenRouter, run twice
(22 decisions per row). November 2026; times are full round trips from a Mac.
This is the one step that needs a Mac. Plug the phone into the Mac, unlock it (tap Trust if asked), then:
git clone https://github.com/jkcoxson/idevice && cd idevice
cargo build --release -p idevice-tools --bin idevice-tools
./target/release/idevice-tools rppairing pair my-mac pairingFile.plist
The file keeps working until the phone stops trusting that host (Reset Location & Privacy, Developer Mode turned off).
Make a new one if the app says the pairing file is invalid or expired.
Without building: add this source in SideStore or AltStore :
https://raw.githubusercontent.com/EliotAndres/iphone-use-app/main/source.json
(or sideload the .ipa from Releases ).
It is the free Apple ID build: install LocalDevVPN too.
From source (needed for the built-in VPN, which takes a paid team):
# once: the prebuilt Rust lib (45 MB, gitignored) from the latest release...
curl -L -o iPhoneUse/idevice/libidevice_ffi.a https://github.com/EliotAndres/iphone-use-app/releases/latest/download/libidevice_ffi.a
# ...or build it yourself (needs rustup + the aarch64-apple-ios target)
scripts/build-idevice.sh
open iPhoneUse.xcodeproj # set your team, run the iPhoneUse scheme on the phone
3. Set up (once, in the app)
Settings (gear, top right) → Import Pairing File . 1
Allow the Add VPN Configuration prompt. It only routes 10.7.0.1 back to the phone.
With LocalDevVPN instead: turn it on, and pick LocalDevVPN in Settings → Phone Connection
(picked for you if the built-in VPN fails).
Settings → Model : pick a provider, paste the key, pick a model.
Wait for the "not connected" line above the input to go away (the developer disk image downloads on first run).
1 Or copy it over once the app is installed:
xcrun devicectl device copy to --device <udid> --domain-type appDataContainer --domain-identifier blog.eliot.iphoneuse --source pairingFile.plist --destination Documents/pairingFile.plist
Type a goal, for example "Open Settings and turn on Dark Mode" , and send it. Stop it from the Dynamic Island
or the chat.
Scripted runs (no LLM)
DEV= < udid > scripts/runscript.sh " open com.apple.Preferences; wait 1.5; tap 0.5 0.9; shot a " a
Launches the app with -agentScript , waits for "script done" in Documents/agent.log and pulls the screenshots.
-agentPrompt "..." as a launch argument starts a real run.
The simulator can't reach a device, but it shows the UI. -demoChat YES plays a fake run
(placeholder screenshots, a question to answer), Debug builds only.
rustup target add aarch64-apple-ios-sim && scripts/build-idevice.sh sim : the libs are
single-arch, so the simulator needs its own, in iPhoneUse/idevice/sim/ .
Build Debug (the scheme defaults to Release, which drops -demoChat ) against that lib:
xcodebuild -project iPhoneUse.xcodeproj -scheme iPhoneUse -configuration Debug -destination id=<sim udid> LIBRARY_SEARCH_PATHS='$(PROJECT_DIR)/iPhoneUse/idevice/sim' build
Install the .app from DerivedData and launch it with -demoChat YES .
Symptom (before the fix): a run froze on a step after about a minute in the background, and the
Live Activity stayed on that step. The app had been killed: the phone's crash logs held an
iPhoneUse.cpu_resource_fatal-*.ips report ("exceeding limit of 80% cpu over 60 seconds", "Process killed").
Cause: the heartbeat thread ( Heartbeat.run() in DeviceSession.swift ) keeps CoreDevice services
alive during a run. Its connection is closed right after the handshake; idevice's userspace TCP stack
(jktcp) then answers every read with BrokenPipe in microseconds, and the loop retried with no delay,
about 45k times a second. iOS kills background apps above 80% CPU over 60 seconds.
The idevice log was written to Documents/idevice_log.txt at Debug level, which prints every XPC
message, screenshots included, as decimal byte lists: 178 MB of log and ~1 GB of memory in a minute.
The loop matched errors by the strings HeartbeatTimeout / HeartbeatSleepyTime , which never
appear in idevice's messages ("heartbeat timeout", "device went to sleep").
Fix: timeouts are matched on the error code (100, sub-code 2) and the loop waits again; any other
error ends the loop. The file log is at Info.
Still open: why the heartbeat connection closes. Runs work without it (a 17-step, 2.5-minute run
finished fine); if CoreDevice services start dropping on long runs, reconnecting the heartbeat
(with backoff) is the next step.
Debugging a killed run: idevicecrashreport -u <udid> -k <dir> pulls the crash logs. Symbolicate
with atos against the dSYM of the exact build (the UUID in the report must match
dwarfdump --uuid ).
Feature scope: docs/FEATURES.md .
Only automate devices and accounts you own or may use, and respect the terms of the apps involved. The agent acts
with your hands: watch it, and keep Stop within reach.
Project
Repo
Used for
StikDebug, by Stephen and contributors
StikDebug/StikDebug
The tunnel, pairing file and developer disk image code is based on it
idevice, by jkcoxson
jkcoxson/idevice
The Rust library behind every device call
StosVPN, by the SideStore Team and Stossy11
SideStore/StosVPN
The loopback VPN ( TunnelProvider/ ) is based on it
LocalDevVPN, by jkcoxson
jkcoxson/LocalDevVPN
The VPN app for free Apple IDs; our tunnel came from StosVPN through it
esp32-iphone-agent, by Henry Yuan
onehumanbeing/esp32-iphone-agent
The model provider and model picker design, and the layout of this README
DeveloperDiskImage, by doronz88
doronz88/DeveloperDiskImage
Developer disk images
License
An iPhone app where an LLM controls the phone it runs on (no Mac required after setup)
Readme AGPL-3.0 license Activity Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) - EliotAndres/iphone-use-app

GitHub - EliotAndres/iphone-use-app: An iPhone app where an LLM controls the phone it runs on (no Mac required after setup) · GitHub
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
EliotAndres
iphone-use-app Public Notifications You must be signed in to change notification settings
Star 33 ( 33 ) You must be signed in to star a repository
An iPhone app where an LLM controls the phone it runs on (no Mac required after setup)
Readme AGPL-3.0 license Activity Stars
5 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
897 Commits 897 Commits Folders and files
AgentWidget AgentWidget Shared Shared TunnelProvider TunnelProvider docs docs evals evals iPhoneUse.xcodeproj iPhoneUse.xcodeproj iPhoneUse iPhoneUse idevice-patch idevice-patch scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md TODO.md TODO.md source.json source.json View all files Repository files navigation
Let an AI model use your iPhone, from the iPhone itself. No Mac at run time (one-time pairing only), no extra hardware, no jailbreak.
Left: "Select a Lyft to the Eiffel Tower (don't book yet)" (4× speed): the agent leaves its own app, opens Lyft,
types the destination, picks it and stops at the ride options. Middle: the finished run, each step with its screenshot. Right: a new chat.
No computer at run time: the agent runs on the iPhone and controls that same iPhone, no extra hardware or jailbreak.
A Mac is only needed once, to make the pairing file
Start it from anywhere: the chat, Siri, the "Run Agent" Shortcut or an iphoneuse://run?prompt=… link.
iphone-use-app
TapKit
esp32-iphone-agent
iPhone-use
Computer at run time
None (a Mac once, for pairing)
Mac app, phone plugged in
None (a computer once, to flash the board)
Mac (MCP server)
Extra hardware
None
None
ESP32-S3 board (~$7) on USB-C
None
How it sees and acts
CoreDevice screenshots and touches, through a loopback VPN
Apple accessibility features
AirPlay mirroring to the board, USB mouse/keyboard via AssistiveTouch
WebDriverAgent (XCTest) over USB
Requirements
Part
Notes
iPhone
iOS 26 or later, Developer Mode on.
Apple account
Paid team: the app uses its own built-in VPN (needs the Network Extension entitlement).
Free Apple ID: build with BUILTIN_VPN=NO (Xcode build setting or xcodebuild argument) and install LocalDevVPN ; re-sign every 7 days.
Pairing file
Made once from a Mac over USB (below).
Vision model
An OpenRouter key, or any OpenAI-compatible endpoint whose model takes images and tool calls.
Mac
Only to make the pairing file (and to build the app, unless you install the release). Not needed at run time.
Limitations
Shortcut needed on cellular: iOS refuses new loopback VPN connections on cellular data, so the app runs a
Cellular Toggle shortcut that turns data off while it connects, then back on (or turn it off by hand).
No other VPN at the same time: it needs a loopback VPN (the built-in one, or
LocalDevVPN ), and iOS allows one VPN at a time, so
Tailscale and similar stop while it is on.
"Viewed remotely from Mac" notice on every run: iOS shows "Device Hub: This device was viewed remotely from
Mac" when its developer services are used, even though the "Mac" here is the app itself. Nothing leaves the phone (except calls to LLMs).
Permissions: requires Developer Mode, a VPN configuration, Local Network access (to find the phone's
pairing service), and Location set to "Always" (keeps runs alive in the background, alongside silent audio).
Microphone only for voice prompts.
Pairing files expire: make a new one when the app says so.
Sideloaded build only: it needs Developer Mode and developer services, so it can't go on the App Store.
A few seconds per step: most of it is the model.
flowchart LR
subgraph App["iphone-use-app"]
L["LLM calling"]
D["Mount developer image"]
O["Orchestrator"]
end
L <-- "screenshot + history /<br/>tool call" --> R["OpenRouter"]
D -- "DDI mount" --> V["Loopback VPN<br/>10.7.0.1"]
O <-- "screenshot / tap, swipe, type" --> V
Loading
The app mounts a developer image, then uses a loopback VPN to screenshot and control the iPhone it runs on.
Each step, the screenshot goes to the model, which answers with a tool call (tap, long press, swipe, type, open app) until it's done.
A Live Activity shows progress with a Stop button, and every run is saved as a chat.
One decision is one screenshot → one tool call. scripts/eval.sh replays 11 recorded steps (open Lyft, search a
destination, pick a ride, stop before paying; write a note) against each model through OpenRouter, run twice
(22 decisions per row). November 2026; times are full round trips from a Mac.
This is the one step that needs a Mac. Plug the phone into the Mac, unlock it (tap Trust if asked), then:
git clone https://github.com/jkcoxson/idevice && cd idevice
cargo build --release -p idevice-tools --bin idevice-tools
./target/release/idevice-tools rppairing pair my-mac pairingFile.plist
The file keeps working until the phone stops trusting that host (Reset Location & Privacy, Developer Mode turned off).
Make a new one if the app says the pairing file is invalid or expired.
Without building: add this source in SideStore or AltStore :
https://raw.githubusercontent.com/EliotAndres/iphone-use-app/main/source.json
(or sideload the .ipa from Releases ).
It is the free Apple ID build: install LocalDevVPN too.
From source (needed for the built-in VPN, which takes a paid team):
# once: the prebuilt Rust lib (45 MB, gitignored) from the latest release...
curl -L -o iPhoneUse/idevice/libidevice_ffi.a https://github.com/EliotAndres/iphone-use-app/releases/latest/download/libidevice_ffi.a
# ...or build it yourself (needs rustup + the aarch64-apple-ios target)
scripts/build-idevice.sh
open iPhoneUse.xcodeproj # set your team, run the iPhoneUse scheme on the phone
3. Set up (once, in the app)
Settings (gear, top right) → Import Pairing File . 1
Allow the Add VPN Configuration prompt. It only routes 10.7.0.1 back to the phone.
With LocalDevVPN instead: turn it on, and pick LocalDevVPN in Settings → Phone Connection
(picked for you if the built-in VPN fails).
Settings → Model : pick a provider, paste the key, pick a model.
Wait for the "not connected" line above the input to go away (the developer disk image downloads on first run).
1 Or copy it over once the app is installed:
xcrun devicectl device copy to --device <udid> --domain-type appDataContainer --domain-identifier blog.eliot.iphoneuse --source pairingFile.plist --destination Documents/pairingFile.plist
Type a goal, for example "Open Settings and turn on Dark Mode" , and send it. Stop it from the Dynamic Island
or the chat.
Scripted runs (no LLM)
DEV= < udid > scripts/runscript.sh " open com.apple.Preferences; wait 1.5; tap 0.5 0.9; shot a " a
Launches the app with -agentScript , waits for "script done" in Documents/agent.log and pulls the screenshots.
-agentPrompt "..." as a launch argument starts a real run.
The simulator can't reach a device, but it shows the UI. -demoChat YES plays a fake run
(placeholder screenshots, a question to answer), Debug builds only.
rustup target add aarch64-apple-ios-sim && scripts/build-idevice.sh sim : the libs are
single-arch, so the simulator needs its own, in iPhoneUse/idevice/sim/ .
Build Debug (the scheme defaults to Release, which drops -demoChat ) against that lib:
xcodebuild -project iPhoneUse.xcodeproj -scheme iPhoneUse -configuration Debug -destination id=<sim udid> LIBRARY_SEARCH_PATHS='$(PROJECT_DIR)/iPhoneUse/idevice/sim' build
Install the .app from DerivedData and launch it with -demoChat YES .
Symptom (before the fix): a run froze on a step after about a minute in the background, and the
Live Activity stayed on that step. The app had been killed: the phone's crash logs held an
iPhoneUse.cpu_resource_fatal-*.ips report ("exceeding limit of 80% cpu over 60 seconds", "Process killed").
Cause: the heartbeat thread ( Heartbeat.run() in DeviceSession.swift ) keeps CoreDevice services
alive during a run. Its connection is closed right after the handshake; idevice's userspace TCP stack
(jktcp) then answers every read with BrokenPipe in microseconds, and the loop retried with no delay,
about 45k times a second. iOS kills background apps above 80% CPU over 60 seconds.
The idevice log was written to Documents/idevice_log.txt at Debug level, which prints every XPC
message, screenshots included, as decimal byte lists: 178 MB of log and ~1 GB of memory in a minute.
The loop matched errors by the strings HeartbeatTimeout / HeartbeatSleepyTime , which never
appear in idevice's messages ("heartbeat timeout", "device went to sleep").
Fix: timeouts are matched on the error code (100, sub-code 2) and the loop waits again; any other
error ends the loop. The file log is at Info.
Still open: why the heartbeat connection closes. Runs work without it (a 17-step, 2.5-minute run
finished fine); if CoreDevice services start dropping on long runs, reconnecting the heartbeat
(with backoff) is the next step.
Debugging a killed run: idevicecrashreport -u <udid> -k <dir> pulls the crash logs. Symbolicate
with atos against the dSYM of the exact build (the UUID in the report must match
dwarfdump --uuid ).
Feature scope: docs/FEATURES.md .
Only automate devices and accounts you own or may use, and respect the terms of the apps involved. The agent acts
with your hands: watch it, and keep Stop within reach.
Project
Repo
Used for
StikDebug, by Stephen and contributors
StikDebug/StikDebug
The tunnel, pairing file and developer disk image code is based on it
idevice, by jkcoxson
jkcoxson/idevice
The Rust library behind every device call
StosVPN, by the SideStore Team and Stossy11
SideStore/StosVPN
The loopback VPN ( TunnelProvider/ ) is based on it
LocalDevVPN, by jkcoxson
jkcoxson/LocalDevVPN
The VPN app for free Apple IDs; our tunnel came from StosVPN through it
esp32-iphone-agent, by Henry Yuan
onehumanbeing/esp32-iphone-agent
The model provider and model picker design, and the layout of this README
DeveloperDiskImage, by doronz88
doronz88/DeveloperDiskImage
Developer disk images
License
An iPhone app where an LLM controls the phone it runs on (no Mac required after setup)
Readme AGPL-3.0 license Activity Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
