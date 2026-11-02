---
source: "https://github.com/shipvo/shipvo"
hn_url: "https://news.ycombinator.com/item?id=50333957"
title: "Show HN: Shipvo – open-source RN OTA updates and an AI-drivable simulator"
article_title: "GitHub - shipvo/shipvo: Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. · GitHub"
image: "https://repository-images.githubusercontent.com/1438951977/e93161df-c008-490f-9a7d-3bac9a298355"
author: "hansade"
captured_at: "2026-11-02T07:18:57Z"
capture_tool: "hn-digest"
hn_id: 50333957
score: 1
comments: 0
posted_at: "2026-11-02T07:13:44Z"
tags:
  - hacker-news
---

# Show HN: Shipvo – open-source RN OTA updates and an AI-drivable simulator

- HN: [50333957](https://news.ycombinator.com/item?id=50333957)
- Source: [github.com](https://github.com/shipvo/shipvo)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T07:13:44Z

## Translation

Title: Show HN: Shipvo – open-source RN OTA updates and an AI-drivable simulator
Article title: GitHub - shipvo/shipvo: Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. · GitHub
Description: Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. - shipvo/shipvo

Article text:
GitHub - shipvo/shipvo: Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
shipvo
shipvo Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
190 Commits 190 Commits Folders and files
.github .github apps/ tide-and-tackle apps/ tide-and-tackle blog blog docs docs legal legal packages packages scripts scripts services services skills/ shipvo-design skills/ shipvo-design .gitignore .gitignore .npmrc .npmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md glama.json glama.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
Shipvo: build once, publish for people and AI agents
One React Native app, two places it runs. People get a native Android or iOS app from the stores, kept current with signed over-the-air updates. AI agents get the same release in a phone-sized simulator they can see and operate over MCP, with no device, including native features (camera, location, biometrics, payments, notifications). Every claim about the simulator is measured against real Android and iOS and published.
Shipvo was briefly called PureNative. The hosted control plane is api.shipvo.dev and the hosted simulator sim.shipvo.dev ; the older purenative.workers.dev addresses keep answering for apps whose shell was built against them.
Publish. Signed OTA updates (64 KB patches, crash-loop rollback, staged rollouts), a built-in owner admin at /app/<id>/admin , and a control plane you can self-host on Cloudflare.
Simulate. @shipvo/mcp opens any Shipvo app for an agent: accessibility tree with refs, tap/type/scroll/back, screenshots. @shipvo/device gives apps one device API, real on phones and simulated in the simulator, with Android- and iOS-worded system sheets the agent answers like a person. Simulation is explicit: every simulated event is logged as simulated and no simulated payment or fingerprint is ever presented as real.
Verify. The parity harness runs the same flows on the simulator, an Android emulator and an iPhone simulator and compares every screen. Current result: 100% actionable parity on Android and iOS across the 13-screen purchase flows; on the native-feature flow, 18 of 21 screens match at 100% on Android (the three gaps are listed) and 100% on iOS, with the orders themselves compared ( docs/parity-results.md , which also says exactly what the score does and does not measure).
Website · Docs · Blog · Shipvo vs Expo · Pricing
Status: pre-launch, measured. OTA numbers come from a real low-end phone (Redmi 9A, 2 GB RAM): docs/spike-results.md . Agent results: docs/agent-gate-results.md . Simulator-vs-device parity: docs/parity-results.md . Native-feature parity (camera, location, permissions, payment prompt) is measured on Android; iOS is next. The iOS OTA runtime is verified on the iOS Simulator in CI (fetch, verify, stage, reload, cold start); physical iPhones are not yet measured.
Every change to a native mobile app normally means a new build, a store submission, a review, and then waiting for users to update. Many never do. Shipvo separates what has to go through the store from what doesn't:
Updates are interpreted JavaScript running on native capabilities the shell declares up front. The app never asks for a permission it hasn't declared.
Signed OTA updates: Ed25519 signatures, verified natively on the device against a public key built into the app. An update can't replace its own trust root.
Integrity checks: SHA-256 and size are verified before a bundle is staged; staging is atomic.
Crash-loop protection: a release that fails to boot 3 times is discarded automatically, and you can roll back to the previous release in one call.
Staged rollouts: cohort bands (1%, 10%, everyone) served as static signed pointers from the edge, with no device registry.
Telemetry for every release: applied, reverted, cold-start time, download size and time.
Control plane on Cloudflare Workers: KV, D1 and R2. It's self-hostable, with a Docker target planned.
Admin per app at /app/<id>/admin : live content, prices, offers, orders, history, publish to every phone in seconds; the same typed owner commands are available to agents with a token.
Agent-operable apps: every app exposes typed commands and a UI simulator; @shipvo/mcp plugs both into Claude, Cursor, Claude Code or PiPilot.
Simulated device: @shipvo/device runs camera, location, biometrics, payments, notifications, share and network on phones (Expo modules) and in the simulator (system sheets in the accessibility tree, state set by the agent). Each capability declares its support level: real, simulated or unavailable.
Design skill: a runnable theme generator, screen checker and anti-generic lint in skills/shipvo-design .
npx -y @shipvo/cli init myshop && cd myshop && npm install && npx playwright install chromium
npm run sim # the app in the simulator; agents: open_app("myshop") through @shipvo/mcp
npx expo run:android # the native shell; `shipvo publish` ships changes over the air
Guides: Build your first app · Device API · Agents and MCP · Self-hosting the control plane · measured results: OTA on a 2 GB phone , simulator vs device parity , agent gate .
Use an app as an AI agent (no phone)
Add the MCP server to Claude, Claude Code, Cursor or PiPilot:
{ "mcpServers" : { "shipvo" : { "command" : " npx " , "args" : [ " -y " , " @shipvo/mcp " ] } } }
Then: open_app("tide@latest") returns a phone-sized session and the first screen as an accessibility tree with refs; tap , type , scroll , back and screenshot act on it; set_device decides what the phone has (a barcode in front of the camera, a location, whether the fingerprint matches, whether a payment is approved, network off); device_log reports every device event, all marked simulated; end_session returns the steps taken and the orders the app sent. tide@r13 opens exactly the release phones are running; tide opens a local build. Full tool list: packages/mcp/README.md .
What the agent sees is what a person with a screen reader would: only the items of a rail that are on screen (swipe to see more), nothing behind a modal sheet, system dialogs worded as Android or iOS word them.
How a release reaches phones and agents
shipvo publish Build Hermes bundle → SHA-256 → bsdiff patches → sign pointer (Ed25519)
│ Build web artifact of the same commit (platform adapters only)
▼
Cloudflare edge /p/{app}/{shell_major}/{band}.json (signed envelope, ETag)
│ /b/{app}/r{release}.hbc (immutable bundle, bsdiff patches)
│ /w/{app}/r{release}/ (the same release for agents)
├──────────────────────────────┐
▼ ▼
Device runtime (Kotlin) Simulator / sim.shipvo.dev
verify → download → verify open_app("tide@r13") → accessibility tree
→ stage → reload → healthy → tap/type/scroll → device sheets → report
or revert after 3 failed boots
▼
Telemetry POST /t → D1; orders POST /o/{app} → D1
Repository layout
Path
What it is
packages/schema
Contracts: shell manifest, add-on manifest, release pointer, signed envelope, gate report, change log, telemetry (Zod)
packages/cli
shipvo CLI: init a new app, sim it, publish / rollback signed releases, signing keys
packages/runtime
@shipvo/runtime : the native OTA runtime (verify, patch, stage, apply, crash-loop rollback) and the config plugin that pins the trust root
packages/sim
@shipvo/sim : the simulator (phone-sized headless browser, accessibility tree, simulated device, mock backend)
packages/mcp
@shipvo/mcp : the MCP server agents use to open, see and operate apps, control the simulated device, and run owner commands
packages/device
@shipvo/device : one device API (camera, location, biometrics, payments, notifications, share, network), real on phones, simulated in the simulator, plus Metro aliases for raw Expo imports
services/control-plane
Hono worker on Cloudflare: pointers, bundles, web artifacts, content, orders, telemetry, admin, owner API
services/sim
Hosted simulator Worker: one sandbox per API key, the same MCP tools over HTTPS (sim.shipvo.dev)
services/site
shipvo.dev: landing page and the public app marketplace
apps/tide-and-tackle
Sample shop app built on the packages above, plus the device drivers and the parity harness ( sim/ )
docs/
Measured results: spike, agent gate, parity
skills/shipvo-design
Design skill with runnable checks
Try the sample app
Download the latest tide-and-tackle-shell.apk from Releases and install it on an Android 7+ phone.
Edit apps/tide-and-tackle/src/release.ts and run shipvo publish (or push to main and let the publish-ota workflow do it). It signs and publishes the release.
The app checks every 10 seconds while open, verifies the update, applies it, and logs the timings on screen.
Change prices or the home screen at /app/tide/admin on your control plane; the phone shows it in about 10 seconds.
Open the same release as an agent: cd apps/tide-and-tackle && pnpm sim , then open_app("tide@latest") through @shipvo/mcp , or scripts/ui screen from a terminal.
pnpm install
pnpm -r test
pnpm --filter @shipvo/control-plane dev
Requires Node 22 and pnpm 10.
Shipvo sits alongside tools like Expo Updates (EAS Update) and self-hosted CodePush servers, which replaced Microsoft App Center CodePush after it was retired in 2025. What Shipvo adds: a declared-capability shell manifest, signed add-ons, a per-app admin, cohort rollouts from static edge files, and apps that AI agents can operate.
Framework: Apache-2.0 . The hosted control plane may be offered under a source-available license; self-hosting stays permitted.
Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. - shipvo/shipvo

GitHub - shipvo/shipvo: Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
shipvo
shipvo Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
190 Commits 190 Commits Folders and files
.github .github apps/ tide-and-tackle apps/ tide-and-tackle blog blog docs docs legal legal packages packages scripts scripts services services skills/ shipvo-design skills/ shipvo-design .gitignore .gitignore .npmrc .npmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md glama.json glama.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
Shipvo: build once, publish for people and AI agents
One React Native app, two places it runs. People get a native Android or iOS app from the stores, kept current with signed over-the-air updates. AI agents get the same release in a phone-sized simulator they can see and operate over MCP, with no device, including native features (camera, location, biometrics, payments, notifications). Every claim about the simulator is measured against real Android and iOS and published.
Shipvo was briefly called PureNative. The hosted control plane is api.shipvo.dev and the hosted simulator sim.shipvo.dev ; the older purenative.workers.dev addresses keep answering for apps whose shell was built against them.
Publish. Signed OTA updates (64 KB patches, crash-loop rollback, staged rollouts), a built-in owner admin at /app/<id>/admin , and a control plane you can self-host on Cloudflare.
Simulate. @shipvo/mcp opens any Shipvo app for an agent: accessibility tree with refs, tap/type/scroll/back, screenshots. @shipvo/device gives apps one device API, real on phones and simulated in the simulator, with Android- and iOS-worded system sheets the agent answers like a person. Simulation is explicit: every simulated event is logged as simulated and no simulated payment or fingerprint is ever presented as real.
Verify. The parity harness runs the same flows on the simulator, an Android emulator and an iPhone simulator and compares every screen. Current result: 100% actionable parity on Android and iOS across the 13-screen purchase flows; on the native-feature flow, 18 of 21 screens match at 100% on Android (the three gaps are listed) and 100% on iOS, with the orders themselves compared ( docs/parity-results.md , which also says exactly what the score does and does not measure).
Website · Docs · Blog · Shipvo vs Expo · Pricing
Status: pre-launch, measured. OTA numbers come from a real low-end phone (Redmi 9A, 2 GB RAM): docs/spike-results.md . Agent results: docs/agent-gate-results.md . Simulator-vs-device parity: docs/parity-results.md . Native-feature parity (camera, location, permissions, payment prompt) is measured on Android; iOS is next. The iOS OTA runtime is verified on the iOS Simulator in CI (fetch, verify, stage, reload, cold start); physical iPhones are not yet measured.
Every change to a native mobile app normally means a new build, a store submission, a review, and then waiting for users to update. Many never do. Shipvo separates what has to go through the store from what doesn't:
Updates are interpreted JavaScript running on native capabilities the shell declares up front. The app never asks for a permission it hasn't declared.
Signed OTA updates: Ed25519 signatures, verified natively on the device against a public key built into the app. An update can't replace its own trust root.
Integrity checks: SHA-256 and size are verified before a bundle is staged; staging is atomic.
Crash-loop protection: a release that fails to boot 3 times is discarded automatically, and you can roll back to the previous release in one call.
Staged rollouts: cohort bands (1%, 10%, everyone) served as static signed pointers from the edge, with no device registry.
Telemetry for every release: applied, reverted, cold-start time, download size and time.
Control plane on Cloudflare Workers: KV, D1 and R2. It's self-hostable, with a Docker target planned.
Admin per app at /app/<id>/admin : live content, prices, offers, orders, history, publish to every phone in seconds; the same typed owner commands are available to agents with a token.
Agent-operable apps: every app exposes typed commands and a UI simulator; @shipvo/mcp plugs both into Claude, Cursor, Claude Code or PiPilot.
Simulated device: @shipvo/device runs camera, location, biometrics, payments, notifications, share and network on phones (Expo modules) and in the simulator (system sheets in the accessibility tree, state set by the agent). Each capability declares its support level: real, simulated or unavailable.
Design skill: a runnable theme generator, screen checker and anti-generic lint in skills/shipvo-design .
npx -y @shipvo/cli init myshop && cd myshop && npm install && npx playwright install chromium
npm run sim # the app in the simulator; agents: open_app("myshop") through @shipvo/mcp
npx expo run:android # the native shell; `shipvo publish` ships changes over the air
Guides: Build your first app · Device API · Agents and MCP · Self-hosting the control plane · measured results: OTA on a 2 GB phone , simulator vs device parity , agent gate .
Use an app as an AI agent (no phone)
Add the MCP server to Claude, Claude Code, Cursor or PiPilot:
{ "mcpServers" : { "shipvo" : { "command" : " npx " , "args" : [ " -y " , " @shipvo/mcp " ] } } }
Then: open_app("tide@latest") returns a phone-sized session and the first screen as an accessibility tree with refs; tap , type , scroll , back and screenshot act on it; set_device decides what the phone has (a barcode in front of the camera, a location, whether the fingerprint matches, whether a payment is approved, network off); device_log reports every device event, all marked simulated; end_session returns the steps taken and the orders the app sent. tide@r13 opens exactly the release phones are running; tide opens a local build. Full tool list: packages/mcp/README.md .
What the agent sees is what a person with a screen reader would: only the items of a rail that are on screen (swipe to see more), nothing behind a modal sheet, system dialogs worded as Android or iOS word them.
How a release reaches phones and agents
shipvo publish Build Hermes bundle → SHA-256 → bsdiff patches → sign pointer (Ed25519)
│ Build web artifact of the same commit (platform adapters only)
▼
Cloudflare edge /p/{app}/{shell_major}/{band}.json (signed envelope, ETag)
│ /b/{app}/r{release}.hbc (immutable bundle, bsdiff patches)
│ /w/{app}/r{release}/ (the same release for agents)
├──────────────────────────────┐
▼ ▼
Device runtime (Kotlin) Simulator / sim.shipvo.dev
verify → download → verify open_app("tide@r13") → accessibility tree
→ stage → reload → healthy → tap/type/scroll → device sheets → report
or revert after 3 failed boots
▼
Telemetry POST /t → D1; orders POST /o/{app} → D1
Repository layout
Path
What it is
packages/schema
Contracts: shell manifest, add-on manifest, release pointer, signed envelope, gate report, change log, telemetry (Zod)
packages/cli
shipvo CLI: init a new app, sim it, publish / rollback signed releases, signing keys
packages/runtime
@shipvo/runtime : the native OTA runtime (verify, patch, stage, apply, crash-loop rollback) and the config plugin that pins the trust root
packages/sim
@shipvo/sim : the simulator (phone-sized headless browser, accessibility tree, simulated device, mock backend)
packages/mcp
@shipvo/mcp : the MCP server agents use to open, see and operate apps, control the simulated device, and run owner commands
packages/device
@shipvo/device : one device API (camera, location, biometrics, payments, notifications, share, network), real on phones, simulated in the simulator, plus Metro aliases for raw Expo imports
services/control-plane
Hono worker on Cloudflare: pointers, bundles, web artifacts, content, orders, telemetry, admin, owner API
services/sim
Hosted simulator Worker: one sandbox per API key, the same MCP tools over HTTPS (sim.shipvo.dev)
services/site
shipvo.dev: landing page and the public app marketplace
apps/tide-and-tackle
Sample shop app built on the packages above, plus the device drivers and the parity harness ( sim/ )
docs/
Measured results: spike, agent gate, parity
skills/shipvo-design
Design skill with runnable checks
Try the sample app
Download the latest tide-and-tackle-shell.apk from Releases and install it on an Android 7+ phone.
Edit apps/tide-and-tackle/src/release.ts and run shipvo publish (or push to main and let the publish-ota workflow do it). It signs and publishes the release.
The app checks every 10 seconds while open, verifies the update, applies it, and logs the timings on screen.
Change prices or the home screen at /app/tide/admin on your control plane; the phone shows it in about 10 seconds.
Open the same release as an agent: cd apps/tide-and-tackle && pnpm sim , then open_app("tide@latest") through @shipvo/mcp , or scripts/ui screen from a terminal.
pnpm install
pnpm -r test
pnpm --filter @shipvo/control-plane dev
Requires Node 22 and pnpm 10.
Shipvo sits alongside tools like Expo Updates (EAS Update) and self-hosted CodePush servers, which replaced Microsoft App Center CodePush after it was retired in 2025. What Shipvo adds: a declared-capability shell manifest, signed add-ons, a per-app admin, cohort rollouts from static edge files, and apps that AI agents can operate.
Framework: Apache-2.0 . The hosted control plane may be offered under a source-available license; self-hosting stays permitted.
Shipvo: build once, publish for people and AI agents. Open-source React Native framework with signed OTA updates, a built-in admin, an agent-usable simulator over MCP, and measured simulator-vs-device parity.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
