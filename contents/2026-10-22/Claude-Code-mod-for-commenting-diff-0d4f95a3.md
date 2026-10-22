---
source: "https://github.com/formstep/claude-mod-comment-diff"
hn_url: "https://news.ycombinator.com/item?id=50195022"
title: "Claude Code mod for commenting diff"
article_title: "GitHub - formstep/claude-mod-comment-diff: A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review · GitHub"
image: "https://opengraph.githubassets.com/bfba161a74a8e485891765bb91b882d3dc63cb65e15f5b0205ec3a23dac7f8f5/formstep/claude-mod-comment-diff"
author: "mellosty"
captured_at: "2026-10-22T13:31:16Z"
capture_tool: "hn-digest"
hn_id: 50195022
score: 1
comments: 0
posted_at: "2026-10-22T13:27:19Z"
tags:
  - hacker-news
---

# Claude Code mod for commenting diff

- HN: [50195022](https://news.ycombinator.com/item?id=50195022)
- Source: [github.com](https://github.com/formstep/claude-mod-comment-diff)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T13:27:19Z

## Translation

Title: Claude Code mod for commenting diff
Article title: GitHub - formstep/claude-mod-comment-diff: A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review · GitHub
Description: A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review - formstep/claude-mod-comment-diff

Article text:
GitHub - formstep/claude-mod-comment-diff: A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review · GitHub
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
formstep
claude-mod-comment-diff Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks tests tests types types .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A Claude Code mod that lets you review Claude's edits the way you review a pull request on GitHub: comment on individual lines of a diff, batch the comments into a review, send it, and watch Claude work through and resolve each one.
Line comments on every diff. Each Edit and Write result, and each shell command that changes files, is redrawn in the transcript as a commentable diff with a + next to every line. Click it to write a comment on that line.
Reviews or single comments. Collect several comments into a review and send them together, or send one comment straight away.
Overview above the prompt. A "Diff comments" list above the prompt box shows every comment with its status, file and line, plus a Submit review button.
Threads. Reply to a comment to add to it, edit your last message, or reopen a resolved comment by replying to it.
Claude resolves what it fixed. The mod gives Claude a resolve tool. After it handles a comment it marks it resolved with a one-line note on what it did, which shows next to the comment.
Works mid-turn. A comment sent while Claude is busy is delivered with its next tool result, so it does not have to wait for the turn to end.
Claude Code 2.1.302 or later , in a terminal session. The diff view is drawn in the terminal only; other surfaces keep the normal diff.
Mouse clicks enabled in your terminal (the + , edit , reply and the buttons are clickable).
At the prompt of a Claude Code terminal session:
/plugin install comment-diff --marketplace formstep/claude-mod-comment-diff
Answer y to add the formstep marketplace, then press Enter to pick the scope (user scope, the default, enables it in every session). The mod is active immediately, with no restart.
claude plugin install comment-diff --marketplace formstep/claude-mod-comment-diff
In the desktop app: click + next to the prompt box, then Plugins , Add marketplace with formstep/claude-mod-comment-diff , then Add plugin .
claude plugin update comment-diff
claude plugin uninstall comment-diff
Usage
When Claude edits or writes a file, or runs a command that changes one, the diff in the transcript shows a dim + beside each line. Click it and a box opens above the prompt with the line and the lines around it. Type the comment in the prompt box itself:
Whatever you had typed in the prompt is set aside while you write the comment, and put back after. Slash commands still run as usual.
Pending comments collect in the Diff comments list above the prompt. Press Submit review (n) to send them all to Claude in one message. Each comment includes the file, the line and a few lines of surrounding context, so Claude knows exactly what you mean.
Open a comment: click it in the list. The box opens with its diff context and the comment's last message in the prompt, ready to rewrite, plus Resolve and Delete . A resolved comment opens as a reply, which reopens it.
Reply: click reply under a comment in the diff.
Edit: click edit on an unresolved comment. If Claude already read it, the edited text is sent again.
Resolve / Delete: use the buttons in the list or the box. Resolved comments cannot be deleted; they stay as a record.
Reopen: reply to a resolved comment and it becomes pending again.
Comments last for the session. They are not saved between sessions.
Very large diffs (past about 90,000 characters of text, roughly 800 lines or more) and errored tool calls keep Claude Code's normal diff view, without comment buttons. Long diffs below that are drawn in several parts, each with its own + buttons. Edits a subagent makes are not drawn as diffs in the transcript, so they have no comment buttons either.
Run the mod from a local clone instead of the installed copy:
git clone https://github.com/formstep/claude-mod-comment-diff
claude --plugin-dir ./claude-mod-comment-diff
Check and test it:
claude plugin validate .
claude plugin test .
tsc -p . # after the mod has loaded once, so .claude-plugin/types/ exists
Layout:
A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review - formstep/claude-mod-comment-diff

GitHub - formstep/claude-mod-comment-diff: A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review · GitHub
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
formstep
claude-mod-comment-diff Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.claude-plugin .claude-plugin hooks hooks tests tests types types .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
A Claude Code mod that lets you review Claude's edits the way you review a pull request on GitHub: comment on individual lines of a diff, batch the comments into a review, send it, and watch Claude work through and resolve each one.
Line comments on every diff. Each Edit and Write result, and each shell command that changes files, is redrawn in the transcript as a commentable diff with a + next to every line. Click it to write a comment on that line.
Reviews or single comments. Collect several comments into a review and send them together, or send one comment straight away.
Overview above the prompt. A "Diff comments" list above the prompt box shows every comment with its status, file and line, plus a Submit review button.
Threads. Reply to a comment to add to it, edit your last message, or reopen a resolved comment by replying to it.
Claude resolves what it fixed. The mod gives Claude a resolve tool. After it handles a comment it marks it resolved with a one-line note on what it did, which shows next to the comment.
Works mid-turn. A comment sent while Claude is busy is delivered with its next tool result, so it does not have to wait for the turn to end.
Claude Code 2.1.302 or later , in a terminal session. The diff view is drawn in the terminal only; other surfaces keep the normal diff.
Mouse clicks enabled in your terminal (the + , edit , reply and the buttons are clickable).
At the prompt of a Claude Code terminal session:
/plugin install comment-diff --marketplace formstep/claude-mod-comment-diff
Answer y to add the formstep marketplace, then press Enter to pick the scope (user scope, the default, enables it in every session). The mod is active immediately, with no restart.
claude plugin install comment-diff --marketplace formstep/claude-mod-comment-diff
In the desktop app: click + next to the prompt box, then Plugins , Add marketplace with formstep/claude-mod-comment-diff , then Add plugin .
claude plugin update comment-diff
claude plugin uninstall comment-diff
Usage
When Claude edits or writes a file, or runs a command that changes one, the diff in the transcript shows a dim + beside each line. Click it and a box opens above the prompt with the line and the lines around it. Type the comment in the prompt box itself:
Whatever you had typed in the prompt is set aside while you write the comment, and put back after. Slash commands still run as usual.
Pending comments collect in the Diff comments list above the prompt. Press Submit review (n) to send them all to Claude in one message. Each comment includes the file, the line and a few lines of surrounding context, so Claude knows exactly what you mean.
Open a comment: click it in the list. The box opens with its diff context and the comment's last message in the prompt, ready to rewrite, plus Resolve and Delete . A resolved comment opens as a reply, which reopens it.
Reply: click reply under a comment in the diff.
Edit: click edit on an unresolved comment. If Claude already read it, the edited text is sent again.
Resolve / Delete: use the buttons in the list or the box. Resolved comments cannot be deleted; they stay as a record.
Reopen: reply to a resolved comment and it becomes pending again.
Comments last for the session. They are not saved between sessions.
Very large diffs (past about 90,000 characters of text, roughly 800 lines or more) and errored tool calls keep Claude Code's normal diff view, without comment buttons. Long diffs below that are drawn in several parts, each with its own + buttons. Edits a subagent makes are not drawn as diffs in the transcript, so they have no comment buttons either.
Run the mod from a local clone instead of the installed copy:
git clone https://github.com/formstep/claude-mod-comment-diff
claude --plugin-dir ./claude-mod-comment-diff
Check and test it:
claude plugin validate .
claude plugin test .
tsc -p . # after the mod has loaded once, so .claude-plugin/types/ exists
Layout:
A Claude Code mod to comment on lines of Claude's diffs, like a GitHub review
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
