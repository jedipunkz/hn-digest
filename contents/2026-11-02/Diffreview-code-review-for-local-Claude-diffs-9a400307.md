---
source: "https://github.com/chalmagean/diffreview"
hn_url: "https://news.ycombinator.com/item?id=50333783"
title: "Diffreview – code review for local Claude diffs"
article_title: "GitHub - chalmagean/diffreview · GitHub"
image: "https://opengraph.githubassets.com/f17cd0cbabb63e1da1173409f32c5717d0596f3bd77bace8c7210fc420acf9fb/chalmagean/diffreview"
author: "mixandgo"
captured_at: "2026-11-02T07:19:00Z"
capture_tool: "hn-digest"
hn_id: 50333783
score: 1
comments: 0
posted_at: "2026-11-02T06:41:40Z"
tags:
  - hacker-news
---

# Diffreview – code review for local Claude diffs

- HN: [50333783](https://news.ycombinator.com/item?id=50333783)
- Source: [github.com](https://github.com/chalmagean/diffreview)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T06:41:40Z

## Translation

Title: Diffreview – code review for local Claude diffs
Article title: GitHub - chalmagean/diffreview · GitHub
Description: Contribute to chalmagean/diffreview development by creating an account on GitHub.

Article text:
GitHub - chalmagean/diffreview · GitHub
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
chalmagean
diffreview Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude/ commands .claude/ commands Sources Sources Tests/ DiffReviewCoreTests Tests/ DiffReviewCoreTests docs docs scripts scripts .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
A native macOS tool for reviewing local code changes GitHub-PR style, built to
send that review back to Claude Code (or anything
else that reads markdown).
Run diffreview in a git repository and a window opens with the diff. Leave inline
comments on lines or ranges, file-level comments and an overall summary, then click
Submit review . The review is printed to stdout as markdown and the process exits.
Swift 6 toolchain (Xcode or the Command Line Tools)
git clone < this repo > diffreview && cd diffreview
swift build -c release
install -m 755 .build/release/diffreview ~ /.local/bin/diffreview # or any directory on your PATH
Usage
diffreview # uncommitted changes: `git diff HEAD` plus untracked files
diffreview --base main # commits on the current branch since its merge-base with main
diffreview --help
Result
stdout
exit code
Submit review
the review as markdown
0
Close window / quit without submitting
nothing
1
No changes to review
nothing (message on stderr)
1
Not a git repository, unknown ref, bad arguments
nothing (message on stderr)
2
--base compares committed work only ( merge-base..HEAD ), like a pull request.
Uncommitted changes are not included in that mode.
Click a line to comment on it. Shift-click another line in the same file to extend the comment to a range.
⌘↩ adds the comment, Esc cancels. Existing comments can be edited or deleted.
Comment on file (in each file header) adds a file-level comment.
Finish review (⌘⇧↩) opens the summary box and the Submit review button.
Files can be collapsed. Files with more than 1,500 changed lines start collapsed.
Drafts are saved automatically to .git/diffreview/comments.json , so closing the
window doesn't lose work: the next run picks up where you left off. The file is
deleted when you submit. If you close the window while a comment is still being
typed, that text is kept as a comment.
# Code Review
Overall summary.
## path/to/file.swift
A file-level comment.
### Lines 42-45 (new)
``` swift
let the = " selected lines "
```
The comment.
Line numbers refer to the new version of the file, or to the old version ( (old) )
when only deleted lines are selected. Selections that mix added and removed lines
are shown as a diff block.
The repo ships a /diffreview command in
.claude/commands/diffreview.md . Install it for all your
projects:
mkdir -p ~ /.claude/commands
cp .claude/commands/diffreview.md ~ /.claude/commands/
It contains:
---
description : Open diffreview to review changes GitHub-style, then address the comments
argument-hint : [--base <ref>]
allowed-tools : Bash(diffreview) Bash(diffreview *)
---
! ` diffreview $ARGUMENTS `
Address each review comment above. If something is unclear, ask before changing it.
Then run /diffreview (or /diffreview --base main ) in any repository. Cancelling the
review makes the command abort, so nothing is sent to Claude.
Timeout: Claude Code runs ! commands under the Bash tool's default timeout of
2 minutes, and reviews usually take longer. Raise it in ~/.claude/settings.json :
{
"env" : {
"BASH_DEFAULT_TIMEOUT_MS" : " 3600000 " ,
"BASH_MAX_TIMEOUT_MS" : " 3600000 "
}
}
This also changes the default timeout for every other Bash command Claude Code runs.
swift build
scripts/test.sh
Use scripts/test.sh rather than swift test . With only the Command Line Tools
installed (no Xcode), swift test builds the swift-testing tests but runs none of
them and still reports success. The script detects that setup and runs the tests
through a small runner ( scripts/TestRunner.swift ). Under Xcode it just calls swift test .
Sources/DiffReviewCore : no UI. Runs git, parses unified diffs (renames, copies,
binary files, mode changes, quoted paths, "\ No newline at end of file"), models
comments, renders the markdown review and stores drafts.
Sources/diffreview : the AppKit/SwiftUI app.
Tests/DiffReviewCoreTests : parser and renderer tests, plus integration tests that
run real git in temporary repositories.
Syntax highlighting isn't implemented yet. Sources/diffreview/SyntaxHighlighter.swift
is the place to add it: each line is passed through SyntaxHighlighter.highlight(_:language:) .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to chalmagean/diffreview development by creating an account on GitHub.

GitHub - chalmagean/diffreview · GitHub
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
chalmagean
diffreview Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.claude/ commands .claude/ commands Sources Sources Tests/ DiffReviewCoreTests Tests/ DiffReviewCoreTests docs docs scripts scripts .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
A native macOS tool for reviewing local code changes GitHub-PR style, built to
send that review back to Claude Code (or anything
else that reads markdown).
Run diffreview in a git repository and a window opens with the diff. Leave inline
comments on lines or ranges, file-level comments and an overall summary, then click
Submit review . The review is printed to stdout as markdown and the process exits.
Swift 6 toolchain (Xcode or the Command Line Tools)
git clone < this repo > diffreview && cd diffreview
swift build -c release
install -m 755 .build/release/diffreview ~ /.local/bin/diffreview # or any directory on your PATH
Usage
diffreview # uncommitted changes: `git diff HEAD` plus untracked files
diffreview --base main # commits on the current branch since its merge-base with main
diffreview --help
Result
stdout
exit code
Submit review
the review as markdown
0
Close window / quit without submitting
nothing
1
No changes to review
nothing (message on stderr)
1
Not a git repository, unknown ref, bad arguments
nothing (message on stderr)
2
--base compares committed work only ( merge-base..HEAD ), like a pull request.
Uncommitted changes are not included in that mode.
Click a line to comment on it. Shift-click another line in the same file to extend the comment to a range.
⌘↩ adds the comment, Esc cancels. Existing comments can be edited or deleted.
Comment on file (in each file header) adds a file-level comment.
Finish review (⌘⇧↩) opens the summary box and the Submit review button.
Files can be collapsed. Files with more than 1,500 changed lines start collapsed.
Drafts are saved automatically to .git/diffreview/comments.json , so closing the
window doesn't lose work: the next run picks up where you left off. The file is
deleted when you submit. If you close the window while a comment is still being
typed, that text is kept as a comment.
# Code Review
Overall summary.
## path/to/file.swift
A file-level comment.
### Lines 42-45 (new)
``` swift
let the = " selected lines "
```
The comment.
Line numbers refer to the new version of the file, or to the old version ( (old) )
when only deleted lines are selected. Selections that mix added and removed lines
are shown as a diff block.
The repo ships a /diffreview command in
.claude/commands/diffreview.md . Install it for all your
projects:
mkdir -p ~ /.claude/commands
cp .claude/commands/diffreview.md ~ /.claude/commands/
It contains:
---
description : Open diffreview to review changes GitHub-style, then address the comments
argument-hint : [--base <ref>]
allowed-tools : Bash(diffreview) Bash(diffreview *)
---
! ` diffreview $ARGUMENTS `
Address each review comment above. If something is unclear, ask before changing it.
Then run /diffreview (or /diffreview --base main ) in any repository. Cancelling the
review makes the command abort, so nothing is sent to Claude.
Timeout: Claude Code runs ! commands under the Bash tool's default timeout of
2 minutes, and reviews usually take longer. Raise it in ~/.claude/settings.json :
{
"env" : {
"BASH_DEFAULT_TIMEOUT_MS" : " 3600000 " ,
"BASH_MAX_TIMEOUT_MS" : " 3600000 "
}
}
This also changes the default timeout for every other Bash command Claude Code runs.
swift build
scripts/test.sh
Use scripts/test.sh rather than swift test . With only the Command Line Tools
installed (no Xcode), swift test builds the swift-testing tests but runs none of
them and still reports success. The script detects that setup and runs the tests
through a small runner ( scripts/TestRunner.swift ). Under Xcode it just calls swift test .
Sources/DiffReviewCore : no UI. Runs git, parses unified diffs (renames, copies,
binary files, mode changes, quoted paths, "\ No newline at end of file"), models
comments, renders the markdown review and stores drafts.
Sources/diffreview : the AppKit/SwiftUI app.
Tests/DiffReviewCoreTests : parser and renderer tests, plus integration tests that
run real git in temporary repositories.
Syntax highlighting isn't implemented yet. Sources/diffreview/SyntaxHighlighter.swift
is the place to add it: each line is passed through SyntaxHighlighter.highlight(_:language:) .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
