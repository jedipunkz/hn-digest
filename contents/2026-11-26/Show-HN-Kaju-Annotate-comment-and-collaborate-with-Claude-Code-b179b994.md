---
source: "https://github.com/aigeano/kaju"
hn_url: "https://news.ycombinator.com/item?id=50660837"
title: "Show HN: Kaju – Annotate, comment and collaborate with Claude Code"
article_title: "GitHub - aigeano/kaju: kaju: annotate, comment and collaborate with Claude Code · GitHub"
image: "https://opengraph.githubassets.com/4f5df7203ff0c53ec3b9d9123169a0bb94b62fc06f1cb0fc23fdd192b9977063/aigeano/kaju"
author: "aigeano"
captured_at: "2026-11-26T23:30:39Z"
capture_tool: "hn-digest"
hn_id: 50660837
score: 1
comments: 0
posted_at: "2026-11-26T23:03:42Z"
tags:
  - hacker-news
---

# Show HN: Kaju – Annotate, comment and collaborate with Claude Code

- HN: [50660837](https://news.ycombinator.com/item?id=50660837)
- Source: [github.com](https://github.com/aigeano/kaju)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T23:03:42Z

## Translation

Title: Show HN: Kaju – Annotate, comment and collaborate with Claude Code
Article title: GitHub - aigeano/kaju: kaju: annotate, comment and collaborate with Claude Code · GitHub
Description: kaju: annotate, comment and collaborate with Claude Code - aigeano/kaju

Article text:
GitHub - aigeano/kaju: kaju: annotate, comment and collaborate with Claude Code · GitHub
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
aigeano
kaju Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
kaju: annotate, comment and collaborate with Claude Code
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
22 Commits 22 Commits Folders and files
.github/ workflows .github/ workflows Resources Resources Sources Sources Tests/ KajuCoreTests Tests/ KajuCoreTests docs docs packaging/ homebrew packaging/ homebrew tools tools .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Annotate, comment and collaborate with Claude Code.
Select some of Claude Code's output in your terminal, press a hotkey, type a comment. kaju sends the quoted text and your comment straight back to Claude Code, which answers right away. For longer reviews of a plan, a document or an explainer, collect comments in a side panel and send them all at once.
kaju is named after a ginger cat who appears to run on a single brain cell but is exceptionally good at one thing: eating. kaju the tool is similar. It does not do much, but it is very good at one thing: comments.
Select text in Claude Code's output. Claude Code copies your selection to the clipboard.
Press ⌃⌥K (or, in iTerm2, right-click › Comment with kaju ).
A small input opens next to the selection, showing the quote it captured. Type your comment.
Press ⏎ . kaju checks that Claude Code is running in that pane, then delivers:
> the text you selected
your comment
Press ⇥ in the input to start a review. Comments then collect in a panel docked on the side of the terminal window, one panel per pane. Send them all as one numbered message with Send N to Claude Code in the panel, or with ⌘⏎ in the input. Reviews are saved to disk on every change and come back after a restart.
Key
In the comment input
⏎
Send the comment, or add it to the review
⇧⏎
New line
⇥
Start or cancel a review
⇧⇥
Tag: fix · question · nit
⌘⏎
Add this comment and send the whole review
Esc
Close and clear the comment
⌘/
Show all keys
Drag
Move the input by its top line
In the review panel: click a card to show it in full, ✎ or double-click to edit, × to remove, drag the header to move the panel, Discard needs two clicks.
The input stays open until you finish or clear it. Clicking back into the terminal leaves it in place. It shows only on its own tab (where the terminal allows it) and hides while another app is in front.
Your shell does not matter. Your terminal does, because terminals differ in what they let other apps see and do.
kitty and WezTerm can get full support with an adapter. Contributions are welcome: see CONTRIBUTING.md .
kaju types into a terminal only when it is sure Claude Code is the program in front. A message pasted into a shell instead could run as commands, because the quoted lines start with > .
iTerm2, Terminal.app and tmux: before sending, kaju checks the foreground process of the pane's terminal. If it is not Claude Code, nothing is typed: your comment goes to the clipboard, or your review stays open.
Other terminals: kaju cannot see which program runs there, so it never presses Return and always pastes the message as a single line with all control characters removed, which no shell will run. It pastes only if the window you commented on is still in front.
brew install aigeano/tap/kaju
brew services start kaju
kaju is compiled on your Mac during install, so you need the Xcode command line tools 15.3 or later ( xcode-select --install ), which Homebrew already requires. Full Xcode is not needed.
After brew upgrade kaju , run brew services restart kaju .
git clone https://github.com/aigeano/kaju.git
cd kaju
make install # builds, installs ~/Applications/kaju.app and a login item
make uninstall # removes them again
Permissions
Permission
Needed for
When macOS asks
Automation › iTerm2
Finding and typing into iTerm2 sessions
First comment in iTerm2
Automation › Terminal
Finding Terminal.app tabs
First comment in Terminal.app
Accessibility
Typing into Terminal.app and other terminals
First comment there
iTerm2 also needs its Python API: iTerm2 › Settings › General › Magic › Enable Python API .
kaju is built on your Mac and signed ad hoc, so macOS treats every upgrade as a new app. kaju clears its old approvals the first time a new build starts, so after an upgrade, approve them again when asked. If a permission looks enabled in System Settings but kaju cannot use it, choose Reset permissions in kaju's menu-bar menu.
kaju reads your selection from the clipboard, where Claude Code's full-screen mode puts it automatically when you select text. In the classic scrolling mode, press ⌘C after selecting, then the hotkey.
kaju's menu-bar menu has Open config file , which opens ~/Library/Application Support/kaju/config.json :
{
"hotkey" : " ctrl+opt+k " ,
"tags" : [ " fix " , " question " , " nit " ],
"tmuxPath" : null ,
"tmuxSocket" : null ,
"clipboardWindowSeconds" : 30
}
Choose Reload config after editing. If the hotkey is taken by another app, the menu says so.
kaju makes no network connections. It reads the clipboard to see your selection, never quotes items that password managers mark as concealed, and, in terminals other than iTerm2, Terminal.app and tmux, only uses clipboard changes made while a terminal was in front. Comments are stored on your Mac only while a review is open. The log at ~/Library/Logs/kaju.log records events, never comment text.
Apache License 2.0. If you redistribute or build on kaju, keep the NOTICE file, which credits the author and links to this repository.
The iTerm2 helper uses the iterm2 Python package (GPL-2.0-or-later), which is installed separately on your machine and not distributed with kaju.
kaju is an independent project and is not affiliated with Anthropic.
kaju: annotate, comment and collaborate with Claude Code
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

kaju: annotate, comment and collaborate with Claude Code - aigeano/kaju

GitHub - aigeano/kaju: kaju: annotate, comment and collaborate with Claude Code · GitHub
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
aigeano
kaju Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
kaju: annotate, comment and collaborate with Claude Code
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
22 Commits 22 Commits Folders and files
.github/ workflows .github/ workflows Resources Resources Sources Sources Tests/ KajuCoreTests Tests/ KajuCoreTests docs docs packaging/ homebrew packaging/ homebrew tools tools .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Annotate, comment and collaborate with Claude Code.
Select some of Claude Code's output in your terminal, press a hotkey, type a comment. kaju sends the quoted text and your comment straight back to Claude Code, which answers right away. For longer reviews of a plan, a document or an explainer, collect comments in a side panel and send them all at once.
kaju is named after a ginger cat who appears to run on a single brain cell but is exceptionally good at one thing: eating. kaju the tool is similar. It does not do much, but it is very good at one thing: comments.
Select text in Claude Code's output. Claude Code copies your selection to the clipboard.
Press ⌃⌥K (or, in iTerm2, right-click › Comment with kaju ).
A small input opens next to the selection, showing the quote it captured. Type your comment.
Press ⏎ . kaju checks that Claude Code is running in that pane, then delivers:
> the text you selected
your comment
Press ⇥ in the input to start a review. Comments then collect in a panel docked on the side of the terminal window, one panel per pane. Send them all as one numbered message with Send N to Claude Code in the panel, or with ⌘⏎ in the input. Reviews are saved to disk on every change and come back after a restart.
Key
In the comment input
⏎
Send the comment, or add it to the review
⇧⏎
New line
⇥
Start or cancel a review
⇧⇥
Tag: fix · question · nit
⌘⏎
Add this comment and send the whole review
Esc
Close and clear the comment
⌘/
Show all keys
Drag
Move the input by its top line
In the review panel: click a card to show it in full, ✎ or double-click to edit, × to remove, drag the header to move the panel, Discard needs two clicks.
The input stays open until you finish or clear it. Clicking back into the terminal leaves it in place. It shows only on its own tab (where the terminal allows it) and hides while another app is in front.
Your shell does not matter. Your terminal does, because terminals differ in what they let other apps see and do.
kitty and WezTerm can get full support with an adapter. Contributions are welcome: see CONTRIBUTING.md .
kaju types into a terminal only when it is sure Claude Code is the program in front. A message pasted into a shell instead could run as commands, because the quoted lines start with > .
iTerm2, Terminal.app and tmux: before sending, kaju checks the foreground process of the pane's terminal. If it is not Claude Code, nothing is typed: your comment goes to the clipboard, or your review stays open.
Other terminals: kaju cannot see which program runs there, so it never presses Return and always pastes the message as a single line with all control characters removed, which no shell will run. It pastes only if the window you commented on is still in front.
brew install aigeano/tap/kaju
brew services start kaju
kaju is compiled on your Mac during install, so you need the Xcode command line tools 15.3 or later ( xcode-select --install ), which Homebrew already requires. Full Xcode is not needed.
After brew upgrade kaju , run brew services restart kaju .
git clone https://github.com/aigeano/kaju.git
cd kaju
make install # builds, installs ~/Applications/kaju.app and a login item
make uninstall # removes them again
Permissions
Permission
Needed for
When macOS asks
Automation › iTerm2
Finding and typing into iTerm2 sessions
First comment in iTerm2
Automation › Terminal
Finding Terminal.app tabs
First comment in Terminal.app
Accessibility
Typing into Terminal.app and other terminals
First comment there
iTerm2 also needs its Python API: iTerm2 › Settings › General › Magic › Enable Python API .
kaju is built on your Mac and signed ad hoc, so macOS treats every upgrade as a new app. kaju clears its old approvals the first time a new build starts, so after an upgrade, approve them again when asked. If a permission looks enabled in System Settings but kaju cannot use it, choose Reset permissions in kaju's menu-bar menu.
kaju reads your selection from the clipboard, where Claude Code's full-screen mode puts it automatically when you select text. In the classic scrolling mode, press ⌘C after selecting, then the hotkey.
kaju's menu-bar menu has Open config file , which opens ~/Library/Application Support/kaju/config.json :
{
"hotkey" : " ctrl+opt+k " ,
"tags" : [ " fix " , " question " , " nit " ],
"tmuxPath" : null ,
"tmuxSocket" : null ,
"clipboardWindowSeconds" : 30
}
Choose Reload config after editing. If the hotkey is taken by another app, the menu says so.
kaju makes no network connections. It reads the clipboard to see your selection, never quotes items that password managers mark as concealed, and, in terminals other than iTerm2, Terminal.app and tmux, only uses clipboard changes made while a terminal was in front. Comments are stored on your Mac only while a review is open. The log at ~/Library/Logs/kaju.log records events, never comment text.
Apache License 2.0. If you redistribute or build on kaju, keep the NOTICE file, which credits the author and links to this repository.
The iTerm2 helper uses the iterm2 Python package (GPL-2.0-or-later), which is installed separately on your machine and not distributed with kaju.
kaju is an independent project and is not affiliated with Anthropic.
kaju: annotate, comment and collaborate with Claude Code
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
