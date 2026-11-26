---
source: "https://github.com/pierridotite/mobile-agent-testers"
hn_url: "https://news.ycombinator.com/item?id=50655803"
title: "Show HN: AI agents explore your mobile app, then their paths replay as tests"
article_title: "GitHub - pierridotite/mobile-agent-testers: 20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs · GitHub"
image: "https://repository-images.githubusercontent.com/1460452633/e9e71cf3-4c59-4b76-bd47-c15c7010b355"
author: "pierridotite"
captured_at: "2026-11-26T14:45:08Z"
capture_tool: "hn-digest"
hn_id: 50655803
score: 1
comments: 0
posted_at: "2026-11-26T14:39:23Z"
tags:
  - hacker-news
---

# Show HN: AI agents explore your mobile app, then their paths replay as tests

- HN: [50655803](https://news.ycombinator.com/item?id=50655803)
- Source: [github.com](https://github.com/pierridotite/mobile-agent-testers)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T14:39:23Z

## Translation

Title: Show HN: AI agents explore your mobile app, then their paths replay as tests
Article title: GitHub - pierridotite/mobile-agent-testers: 20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs · GitHub
Description: 20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs - pierridotite/mobile-agent-testers

Article text:
GitHub - pierridotite/mobile-agent-testers: 20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs · GitHub
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
pierridotite
mobile-agent-testers Public template Notifications You must be signed in to change notification settings
Star 2.5k ( 2.5k ) You must be signed in to star a repository
20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs
Readme MIT license Contributing
Contributing Cite this repository Activity Stars
252 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
182 Commits 182 Commits Folders and files
.claude-plugin .claude-plugin .github .github assets assets briefs briefs dashboard dashboard demo-app demo-app docs docs i18n i18n plugins/ mobile-agent-testers plugins/ mobile-agent-testers scripts scripts skills/ mobile-agent-testers skills/ mobile-agent-testers tests tests traces traces .gitattributes .gitattributes .gitignore .gitignore AGENT_BRIEF.md AGENT_BRIEF.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md personas.json personas.json results_public.pem results_public.pem setup.ps1 setup.ps1 setup.sh setup.sh View all files Repository files navigation
20 AI testers on 20 different phones, in parallel, on GitHub's free runners.
Examples · Dashboard demo · Compare · FAQ · Personas · Contributing
English · 简体中文 · 日本語 · Español · Português · Tiếng Việt · Türkçe
The phones are real screen recordings of the LocalSend run; the map is the dashboard on the demo app . Video with sound (MP4) · The 20 phones, full recording · The run, device by device
No test scripts to write. Each tester is a Claude Code agent with its own iOS
Simulator or Android emulator on GitHub Actions. It plays a persona (iPhone SE
in Swedish at the largest text size, Pixel 7 at 200%, iPad in German, Arabic
right to left...), uses your app like a new user (or runs through the feature you
point it at) and writes a bug report with screenshots. Flutter, React Native, Expo, native Kotlin and Swift, Kotlin
Multiplatform, Capacitor and .NET MAUI.
GitHub's runners are free for public repos. You only pay for the model, with your
Claude subscription or API key: about $0.15 to $0.55 per phone in these runs .
gh repo create my-testers --public --template pierridotite/mobile-agent-testers --clone
cd my-testers && ./setup.sh # key pair + secrets + Claude token (Windows: .\setup.ps1)
gh workflow run agent-testers.yml # demo: 2 Android + 2 iOS testers, ~20 min
bash scripts/collect.sh < run id > # decrypts reports, writes runs/<id>/SUMMARY.md
Then on your own app:
gh workflow run agent-testers.yml -f source_repo=you/app -f app_id=com.you.app \
-f android_personas=anna,malik -f ios_personas=clara,jonas
Or let Claude set it up for you: /plugin marketplace add pierridotite/mobile-agent-testers , then /plugin install mobile-agent-testers@mobile-agent-testers .
Every input (builds, briefs, bug reports, cheaper models, build arguments, video) and the supported frameworks: docs/configuration.md .
The demo app has 11 deliberate bugs . In one run , the 4 demo testers found 9 of them (2 to 7 minutes per agent, about 21 minutes for the whole run, builds included), plus real issues nobody planted (a button under the Android nav bar, no data kept after a relaunch). Model usage: $0.87 for the 4.
Each persona explores your app like a new user and reports what gets in the way, with numbered screenshots. A run can also follow one feature end to end, reproduce a bug report, write an improvement report, compare two versions or test a website on phones. One use per run, set with the inputs above.
Any Anthropic-compatible model works, from Claude Haiku to OpenRouter, DeepSeek, Kimi, GLM, MiniMax or Qwen; findings are only as good as the model's reading of screenshots ( docs/cheaper-models.md ). Maestro or Appium run the tests you write; these agents look for what nobody thought to test ( how it compares ).
Every screen the testers reached, with the path of each one in its own color and a heat map of where they tapped. Pick a device in the Device menu to see the screens in its shape (phone or tablet) and its own heat map. Comment a screen, or turn what is wrong on it into a check.
On your own run: python3 scripts/dashboard.py runs/<id> . It runs on your machine and the results never leave it ( docs/dashboard.md ).
A tester that finds a bug took a path no script would have written. Each Android tester keeps that path as a trace, with each tap named after the element under the finger. With -f replay_path=true , the path replays on the same phone right after the agent, with no model. A check added on the map makes the bug part of the test: the replay says bug while it is there and fixed once it is gone ( docs/replay.md ).
The traces follow Explore-Compile-Replay, my open format that compiles an agent's path into a Maestro flow. Android only for now; ECR is an early draft.
Open-source apps, with nothing planted. On LocalSend (Flutter, 100k stars), 20 phones in 14 languages found 46 distinct bugs in 69 minutes, for $4.62 of model usage. One is a blocker: on an iPhone SE in Swedish at the largest text size, you can't send a text message at all.
34 of LocalSend's 46 bugs had nothing to do with text size or accessibility: broken flows, silent failures, missing translations, tablet layouts. Screenshots, every finding and the commands to run them again: docs/examples.md .
Issues are ready to take , some of them good first issues : new personas, another agent CLI, GitLab CI, translations. Comment on the one you take so two people don't build the same thing. CONTRIBUTING.md says where each one starts and how to test it offline. Merged a pull request, or ran it on your app? Show your work . Questions: Discord .
Hosted version for teams (exploring)
The workflow here stays free, open source and complete. I'm exploring a hosted version for teams with private apps: macOS runners included (on a private repo, GitHub counts each macOS minute ten times) and a web dashboard to share the reports with the whole team, product managers included.
Would that help your team? Answer three questions or write to me on LinkedIn . Already ran it on your app? Tell me what it found and what got in the way .
Claude Code ·
AXe ·
android-emulator-runner
Thanks to everyone who has contributed. Want to help? Start here .
Made by Pierre Raffalli · LinkedIn · GitHub · Discord · MIT license
20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs
Readme MIT license Contributing
Contributing Cite this repository Activity Stars
252 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs - pierridotite/mobile-agent-testers

GitHub - pierridotite/mobile-agent-testers: 20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs · GitHub
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
pierridotite
mobile-agent-testers Public template Notifications You must be signed in to change notification settings
Star 2.5k ( 2.5k ) You must be signed in to star a repository
20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs
Readme MIT license Contributing
Contributing Cite this repository Activity Stars
252 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
182 Commits 182 Commits Folders and files
.claude-plugin .claude-plugin .github .github assets assets briefs briefs dashboard dashboard demo-app demo-app docs docs i18n i18n plugins/ mobile-agent-testers plugins/ mobile-agent-testers scripts scripts skills/ mobile-agent-testers skills/ mobile-agent-testers tests tests traces traces .gitattributes .gitattributes .gitignore .gitignore AGENT_BRIEF.md AGENT_BRIEF.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md personas.json personas.json results_public.pem results_public.pem setup.ps1 setup.ps1 setup.sh setup.sh View all files Repository files navigation
20 AI testers on 20 different phones, in parallel, on GitHub's free runners.
Examples · Dashboard demo · Compare · FAQ · Personas · Contributing
English · 简体中文 · 日本語 · Español · Português · Tiếng Việt · Türkçe
The phones are real screen recordings of the LocalSend run; the map is the dashboard on the demo app . Video with sound (MP4) · The 20 phones, full recording · The run, device by device
No test scripts to write. Each tester is a Claude Code agent with its own iOS
Simulator or Android emulator on GitHub Actions. It plays a persona (iPhone SE
in Swedish at the largest text size, Pixel 7 at 200%, iPad in German, Arabic
right to left...), uses your app like a new user (or runs through the feature you
point it at) and writes a bug report with screenshots. Flutter, React Native, Expo, native Kotlin and Swift, Kotlin
Multiplatform, Capacitor and .NET MAUI.
GitHub's runners are free for public repos. You only pay for the model, with your
Claude subscription or API key: about $0.15 to $0.55 per phone in these runs .
gh repo create my-testers --public --template pierridotite/mobile-agent-testers --clone
cd my-testers && ./setup.sh # key pair + secrets + Claude token (Windows: .\setup.ps1)
gh workflow run agent-testers.yml # demo: 2 Android + 2 iOS testers, ~20 min
bash scripts/collect.sh < run id > # decrypts reports, writes runs/<id>/SUMMARY.md
Then on your own app:
gh workflow run agent-testers.yml -f source_repo=you/app -f app_id=com.you.app \
-f android_personas=anna,malik -f ios_personas=clara,jonas
Or let Claude set it up for you: /plugin marketplace add pierridotite/mobile-agent-testers , then /plugin install mobile-agent-testers@mobile-agent-testers .
Every input (builds, briefs, bug reports, cheaper models, build arguments, video) and the supported frameworks: docs/configuration.md .
The demo app has 11 deliberate bugs . In one run , the 4 demo testers found 9 of them (2 to 7 minutes per agent, about 21 minutes for the whole run, builds included), plus real issues nobody planted (a button under the Android nav bar, no data kept after a relaunch). Model usage: $0.87 for the 4.
Each persona explores your app like a new user and reports what gets in the way, with numbered screenshots. A run can also follow one feature end to end, reproduce a bug report, write an improvement report, compare two versions or test a website on phones. One use per run, set with the inputs above.
Any Anthropic-compatible model works, from Claude Haiku to OpenRouter, DeepSeek, Kimi, GLM, MiniMax or Qwen; findings are only as good as the model's reading of screenshots ( docs/cheaper-models.md ). Maestro or Appium run the tests you write; these agents look for what nobody thought to test ( how it compares ).
Every screen the testers reached, with the path of each one in its own color and a heat map of where they tapped. Pick a device in the Device menu to see the screens in its shape (phone or tablet) and its own heat map. Comment a screen, or turn what is wrong on it into a check.
On your own run: python3 scripts/dashboard.py runs/<id> . It runs on your machine and the results never leave it ( docs/dashboard.md ).
A tester that finds a bug took a path no script would have written. Each Android tester keeps that path as a trace, with each tap named after the element under the finger. With -f replay_path=true , the path replays on the same phone right after the agent, with no model. A check added on the map makes the bug part of the test: the replay says bug while it is there and fixed once it is gone ( docs/replay.md ).
The traces follow Explore-Compile-Replay, my open format that compiles an agent's path into a Maestro flow. Android only for now; ECR is an early draft.
Open-source apps, with nothing planted. On LocalSend (Flutter, 100k stars), 20 phones in 14 languages found 46 distinct bugs in 69 minutes, for $4.62 of model usage. One is a blocker: on an iPhone SE in Swedish at the largest text size, you can't send a text message at all.
34 of LocalSend's 46 bugs had nothing to do with text size or accessibility: broken flows, silent failures, missing translations, tablet layouts. Screenshots, every finding and the commands to run them again: docs/examples.md .
Issues are ready to take , some of them good first issues : new personas, another agent CLI, GitLab CI, translations. Comment on the one you take so two people don't build the same thing. CONTRIBUTING.md says where each one starts and how to test it offline. Merged a pull request, or ran it on your app? Show your work . Questions: Discord .
Hosted version for teams (exploring)
The workflow here stays free, open source and complete. I'm exploring a hosted version for teams with private apps: macOS runners included (on a private repo, GitHub counts each macOS minute ten times) and a web dashboard to share the reports with the whole team, product managers included.
Would that help your team? Answer three questions or write to me on LinkedIn . Already ran it on your app? Tell me what it found and what got in the way .
Claude Code ·
AXe ·
android-emulator-runner
Thanks to everyone who has contributed. Want to help? Start here .
Made by Pierre Raffalli · LinkedIn · GitHub · Discord · MIT license
20 AI testers on 20 different phones, in parallel, on GitHub's free runners: Claude Code agents on iOS Simulators and Android emulators that find broken flows, translation, layout and accessibility bugs
Readme MIT license Contributing
Contributing Cite this repository Activity Stars
252 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
