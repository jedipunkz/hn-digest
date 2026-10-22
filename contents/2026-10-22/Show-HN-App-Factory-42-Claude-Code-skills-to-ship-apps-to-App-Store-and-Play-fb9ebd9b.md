---
source: "https://github.com/yybd/app-factory"
hn_url: "https://news.ycombinator.com/item?id=50194964"
title: "Show HN: App Factory – 42 Claude Code skills to ship apps to App Store and Play"
article_title: "GitHub - yybd/app-factory: The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. · GitHub"
image: "https://repository-images.githubusercontent.com/1394992024/d755d542-c521-4ffd-9a40-a94e3162da36"
author: "yybd"
captured_at: "2026-10-22T13:31:22Z"
capture_tool: "hn-digest"
hn_id: 50194964
score: 2
comments: 0
posted_at: "2026-10-22T13:22:58Z"
tags:
  - hacker-news
---

# Show HN: App Factory – 42 Claude Code skills to ship apps to App Store and Play

- HN: [50194964](https://news.ycombinator.com/item?id=50194964)
- Source: [github.com](https://github.com/yybd/app-factory)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T13:22:58Z

## Translation

Title: Show HN: App Factory – 42 Claude Code skills to ship apps to App Store and Play
Article title: GitHub - yybd/app-factory: The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. · GitHub
Description: The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. - yybd/app-factory

Article text:
GitHub - yybd/app-factory: The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. · GitHub
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
yybd
app-factory Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks.
www.bdtech.app/blog/skills-and-tracks?utm_source=github Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.claude-plugin .claude-plugin .github .github .vscode .vscode docs docs factory factory plugins plugins .editorconfig .editorconfig .gitignore .gitignore .prettierignore .prettierignore APP-LIFECYCLE.he.md APP-LIFECYCLE.he.md APP-LIFECYCLE.md APP-LIFECYCLE.md CHANGELOG.md CHANGELOG.md CLAUDE.he.md CLAUDE.he.md CLAUDE.md CLAUDE.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md FAQ.he.md FAQ.he.md FAQ.md FAQ.md HOW-IT-WORKS.he.md HOW-IT-WORKS.he.md HOW-IT-WORKS.md HOW-IT-WORKS.md LICENSE LICENSE QUICKSTART.he.md QUICKSTART.he.md QUICKSTART.md QUICKSTART.md README.he.md README.he.md README.md README.md SECURITY.md SECURITY.md SKILLS.he.md SKILLS.he.md SKILLS.md SKILLS.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRUST.md TRUST.md analyze_skills.py analyze_skills.py View all files Repository files navigation
42 skills that build, describe, ship and announce apps — to the App Store and
Google Play, and outside them, on the Mac and on Windows. Each one loads only where it is wanted.
The story behind it — why one big instructions file stopped holding, and what replaced it:
the article .
They are Claude Code plugins. You enable the tracks a repo needs, and a session opened
there carries them; a session opened anywhere else does not.
Built for Claude Code, and tested only there. The scripts the skills run are
ordinary command-line tools any agent can call, but the loading — tracks, the
marketplace, enable.py — is Claude Code's plugin system, and most skills find their
scripts through ${CLAUDE_PLUGIN_ROOT} , which only Claude Code sets.
claude plugin marketplace add yybd/app-factory
cd ~ /code/my-app && python3 ~ /.claude/plugins/marketplaces/app-factory/factory/enable.py
# then the install lines it prints, and QUICKSTART for the rest
QUICKSTART.md walks the whole thing — fifteen minutes, one app,
with a check after every step.
The problem it is built around
The usual way to steer a coding agent is one large instruction file. Everything goes in:
release rules, test conventions, SEO, the way you name things. It does not hold. The
file grows past what stays in view, rules contradict each other, and the agent proposes
an Android deploy while you are editing a website.
The failure is not that the agent forgets. It is that everything was told to it at
once, whether or not it applied .
So: nothing is global. A skill is a document the agent loads when its description
matches what you asked for , and the descriptions are the only part that is always in
context. What it costs to have 42 of them available is about 8,400 tokens of
descriptions — and only in a repo where you turned them on.
Eight tracks of skills, plus one that does setup:
Each track's own README lists what is in it, what it costs, and what it needs installed.
The counts above are generated from the skills on disk — in this repo, every number
that can be counted is.
Executors do one thing: play-store-ship builds, signs and uploads an AAB.
app-icon-generator writes every icon size from one image.
Orchestrators sequence executors without doing their work. prepare-app-release
drives identity, compliance, media, copy, the version bump, the build and the upload —
and hands the website leg off explicitly rather than doing it silently.
The division matters because the second kind is where a release goes wrong: not in any
single step, but in the order, and in the step nobody remembered.
It writes files in the repo you point it at, reads credentials in place without ever
copying or printing them, and uploads to the stores at the end of skills whose whole job
that is. It never submits an app for review, never changes a price, and never pushes.
There is one directory it will delete from. TRUST.md is the full list,
written to be read before you enable anything.
Nothing, to install. Individual skills need what they need — Xcode for Apple, a JDK for
Android, fastlane for either store — and each says so and stops rather than failing
halfway. COMPATIBILITY.md has the table.
You do not need a data repo, a website, or the companion guards project. All three
are optional and the skills detect their absence rather than assuming it.
QUICKSTART.md
install to first working skill
TRUST.md
what the skills do to your repos and stores
COMPATIBILITY.md
platforms, tools, credentials, versions
SKILLS.md
every skill and where it loads
APP-LIFECYCLE.md
the whole path from finished code to a live listing
HOW-IT-WORKS.md
loading, the plugin cache, and the three traps
CONTRIBUTING.md
adding a skill, and the rules that came from real failures
FAQ.md
the glossary, and the questions installing raised
CHANGELOG.md
what changed, and why a version matters
SECURITY.md
what is handled, and how to report something
Working across several repos
Once you are shipping more than one app, a second problem appears: an app's code, its
listing text and its website are different repos, and a session stands in exactly one of
them. Work left in another is not missing — it is invisible, which is how "done" gets
said over a clean git status .
That is RepoGuards : one registry that decides
which tracks load where, and guards that refuse a write into a repo the session is not
standing in. It is a separate marketplace, it is optional, and nothing here requires it.
Licence: MIT, except two skills that came from elsewhere and keep their own —
see THIRD_PARTY_NOTICES.md .
The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks.
www.bdtech.app/blog/skills-and-tracks?utm_source=github Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. - yybd/app-factory

GitHub - yybd/app-factory: The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks. · GitHub
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
yybd
app-factory Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks.
www.bdtech.app/blog/skills-and-tracks?utm_source=github Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
52 Commits 52 Commits Folders and files
.claude-plugin .claude-plugin .github .github .vscode .vscode docs docs factory factory plugins plugins .editorconfig .editorconfig .gitignore .gitignore .prettierignore .prettierignore APP-LIFECYCLE.he.md APP-LIFECYCLE.he.md APP-LIFECYCLE.md APP-LIFECYCLE.md CHANGELOG.md CHANGELOG.md CLAUDE.he.md CLAUDE.he.md CLAUDE.md CLAUDE.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md FAQ.he.md FAQ.he.md FAQ.md FAQ.md HOW-IT-WORKS.he.md HOW-IT-WORKS.he.md HOW-IT-WORKS.md HOW-IT-WORKS.md LICENSE LICENSE QUICKSTART.he.md QUICKSTART.he.md QUICKSTART.md QUICKSTART.md README.he.md README.he.md README.md README.md SECURITY.md SECURITY.md SKILLS.he.md SKILLS.he.md SKILLS.md SKILLS.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRUST.md TRUST.md analyze_skills.py analyze_skills.py View all files Repository files navigation
42 skills that build, describe, ship and announce apps — to the App Store and
Google Play, and outside them, on the Mac and on Windows. Each one loads only where it is wanted.
The story behind it — why one big instructions file stopped holding, and what replaced it:
the article .
They are Claude Code plugins. You enable the tracks a repo needs, and a session opened
there carries them; a session opened anywhere else does not.
Built for Claude Code, and tested only there. The scripts the skills run are
ordinary command-line tools any agent can call, but the loading — tracks, the
marketplace, enable.py — is Claude Code's plugin system, and most skills find their
scripts through ${CLAUDE_PLUGIN_ROOT} , which only Claude Code sets.
claude plugin marketplace add yybd/app-factory
cd ~ /code/my-app && python3 ~ /.claude/plugins/marketplaces/app-factory/factory/enable.py
# then the install lines it prints, and QUICKSTART for the rest
QUICKSTART.md walks the whole thing — fifteen minutes, one app,
with a check after every step.
The problem it is built around
The usual way to steer a coding agent is one large instruction file. Everything goes in:
release rules, test conventions, SEO, the way you name things. It does not hold. The
file grows past what stays in view, rules contradict each other, and the agent proposes
an Android deploy while you are editing a website.
The failure is not that the agent forgets. It is that everything was told to it at
once, whether or not it applied .
So: nothing is global. A skill is a document the agent loads when its description
matches what you asked for , and the descriptions are the only part that is always in
context. What it costs to have 42 of them available is about 8,400 tokens of
descriptions — and only in a repo where you turned them on.
Eight tracks of skills, plus one that does setup:
Each track's own README lists what is in it, what it costs, and what it needs installed.
The counts above are generated from the skills on disk — in this repo, every number
that can be counted is.
Executors do one thing: play-store-ship builds, signs and uploads an AAB.
app-icon-generator writes every icon size from one image.
Orchestrators sequence executors without doing their work. prepare-app-release
drives identity, compliance, media, copy, the version bump, the build and the upload —
and hands the website leg off explicitly rather than doing it silently.
The division matters because the second kind is where a release goes wrong: not in any
single step, but in the order, and in the step nobody remembered.
It writes files in the repo you point it at, reads credentials in place without ever
copying or printing them, and uploads to the stores at the end of skills whose whole job
that is. It never submits an app for review, never changes a price, and never pushes.
There is one directory it will delete from. TRUST.md is the full list,
written to be read before you enable anything.
Nothing, to install. Individual skills need what they need — Xcode for Apple, a JDK for
Android, fastlane for either store — and each says so and stops rather than failing
halfway. COMPATIBILITY.md has the table.
You do not need a data repo, a website, or the companion guards project. All three
are optional and the skills detect their absence rather than assuming it.
QUICKSTART.md
install to first working skill
TRUST.md
what the skills do to your repos and stores
COMPATIBILITY.md
platforms, tools, credentials, versions
SKILLS.md
every skill and where it loads
APP-LIFECYCLE.md
the whole path from finished code to a live listing
HOW-IT-WORKS.md
loading, the plugin cache, and the three traps
CONTRIBUTING.md
adding a skill, and the rules that came from real failures
FAQ.md
the glossary, and the questions installing raised
CHANGELOG.md
what changed, and why a version matters
SECURITY.md
what is handled, and how to report something
Working across several repos
Once you are shipping more than one app, a second problem appears: an app's code, its
listing text and its website are different repos, and a session stands in exactly one of
them. Work left in another is not missing — it is invisible, which is how "done" gets
said over a clean git status .
That is RepoGuards : one registry that decides
which tracks load where, and guards that refuse a write into a repo the session is not
standing in. It is a separate marketplace, it is optional, and nothing here requires it.
Licence: MIT, except two skills that came from elsewhere and keep their own —
see THIRD_PARTY_NOTICES.md .
The skills that build, describe, ship and announce apps — to the App Store and Google Play, and direct on Mac and Windows. A Claude Code plugin marketplace: 42 skills in eight opt-in tracks.
www.bdtech.app/blog/skills-and-tracks?utm_source=github Topics
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
