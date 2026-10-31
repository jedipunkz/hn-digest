---
source: "https://github.com/alepoll/NichiHab"
hn_url: "https://news.ycombinator.com/item?id=50318196"
title: "NichiHab Android app blocker with Gemini AI task verification"
article_title: "GitHub - alepoll/NichiHab · GitHub"
image: "https://opengraph.githubassets.com/32c862e87acad2b6615712177782807c93dd09479cb8f32f3617cac89e5e8edf/alepoll/NichiHab"
author: "alepoll"
captured_at: "2026-10-31T18:14:15Z"
capture_tool: "hn-digest"
hn_id: 50318196
score: 1
comments: 0
posted_at: "2026-10-31T17:37:54Z"
tags:
  - hacker-news
---

# NichiHab Android app blocker with Gemini AI task verification

- HN: [50318196](https://news.ycombinator.com/item?id=50318196)
- Source: [github.com](https://github.com/alepoll/NichiHab)
- Score: 1
- Comments: 0
- Posted: 2026-10-31T17:37:54Z

## Translation

Title: NichiHab Android app blocker with Gemini AI task verification
Article title: GitHub - alepoll/NichiHab · GitHub
Description: Contribute to alepoll/NichiHab development by creating an account on GitHub.

Article text:
GitHub - alepoll/NichiHab · GitHub
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
alepoll
NichiHab Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
18 Commits 18 Commits Folders and files
app app gradle gradle .gitignore .gitignore README.md README.md build.gradle.kts build.gradle.kts contextoCode.md contextoCode.md gradle.properties gradle.properties gradlew gradlew settings.gradle.kts settings.gradle.kts View all files Repository files navigation
📦 Getting Started
Prerequisites
NichiHab is a high-strictness habit-building and focus application for Android. It uses AI verification and system-level blocking to ensure you stay committed to your goals before accessing distracting applications.
Smart App Blocking: Uses Android's Accessibility Service to detect and block distracting apps during your focus sessions.
AI Task Verification: Powered by Google Gemini AI . Complete a real-world task (like making your bed or reading) and prove it with a photo. The AI verifies the photo in real-time.
Hard Mode: Prevents access to phone settings and app uninstallation during active sessions to eliminate the temptation to bypass blocks.
Earned Unlock Time: Complete tasks to "earn" minutes of unlocked screen time.
Privacy First: Photos are processed ephemerally. They are encrypted, sent for analysis, and never saved to your gallery or stored on our servers.
Language: Kotlin (100%) & C++ (Security Layer)
UI: Jetpack Compose with Material Design 3
Architecture: MVVM with Dagger Hilt (Dependency Injection)
Asynchrony: Kotlin Coroutines & Flow
Database: Room (Session and Task persistence)
Storage: Jetpack DataStore (Settings)
Android Studio Ladybug (2024.2.1) or newer
git clone https://github.com/your-username/NichiHab.git
Open the project in Android Studio.
Sync Project with Gradle Files.
Configure your Gemini API Key in app/src/main/cpp/security-lib.cpp .
(Optional) Configure your Gemini API Key in app/src/main/cpp/security-lib.cpp .
The application includes a native C++ layer ( security-lib.cpp ) to handle sensitive keys and integrity checks, providing a robust defense against common APK modification tools.
This project is licensed under the MIT License - see the LICENSE file for details.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to alepoll/NichiHab development by creating an account on GitHub.

GitHub - alepoll/NichiHab · GitHub
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
alepoll
NichiHab Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
18 Commits 18 Commits Folders and files
app app gradle gradle .gitignore .gitignore README.md README.md build.gradle.kts build.gradle.kts contextoCode.md contextoCode.md gradle.properties gradle.properties gradlew gradlew settings.gradle.kts settings.gradle.kts View all files Repository files navigation
📦 Getting Started
Prerequisites
NichiHab is a high-strictness habit-building and focus application for Android. It uses AI verification and system-level blocking to ensure you stay committed to your goals before accessing distracting applications.
Smart App Blocking: Uses Android's Accessibility Service to detect and block distracting apps during your focus sessions.
AI Task Verification: Powered by Google Gemini AI . Complete a real-world task (like making your bed or reading) and prove it with a photo. The AI verifies the photo in real-time.
Hard Mode: Prevents access to phone settings and app uninstallation during active sessions to eliminate the temptation to bypass blocks.
Earned Unlock Time: Complete tasks to "earn" minutes of unlocked screen time.
Privacy First: Photos are processed ephemerally. They are encrypted, sent for analysis, and never saved to your gallery or stored on our servers.
Language: Kotlin (100%) & C++ (Security Layer)
UI: Jetpack Compose with Material Design 3
Architecture: MVVM with Dagger Hilt (Dependency Injection)
Asynchrony: Kotlin Coroutines & Flow
Database: Room (Session and Task persistence)
Storage: Jetpack DataStore (Settings)
Android Studio Ladybug (2024.2.1) or newer
git clone https://github.com/your-username/NichiHab.git
Open the project in Android Studio.
Sync Project with Gradle Files.
Configure your Gemini API Key in app/src/main/cpp/security-lib.cpp .
(Optional) Configure your Gemini API Key in app/src/main/cpp/security-lib.cpp .
The application includes a native C++ layer ( security-lib.cpp ) to handle sensitive keys and integrity checks, providing a robust defense against common APK modification tools.
This project is licensed under the MIT License - see the LICENSE file for details.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
