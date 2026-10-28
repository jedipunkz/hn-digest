---
source: "https://github.com/grimmerk/memsaga"
hn_url: "https://news.ycombinator.com/item?id=50276188"
title: "Show HN: Memsaga – Git history of Claude Code memory and skills, by session"
article_title: "GitHub - grimmerk/memsaga: Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own\n[truncated]"
image: "https://repository-images.githubusercontent.com/1417536757/41408506-dd4e-4e49-84cf-6ae4a804ed13"
author: "grimmer"
captured_at: "2026-10-28T17:08:44Z"
capture_tool: "hn-digest"
hn_id: 50276188
score: 1
comments: 0
posted_at: "2026-10-28T16:21:26Z"
tags:
  - hacker-news
---

# Show HN: Memsaga – Git history of Claude Code memory and skills, by session

- HN: [50276188](https://news.ycombinator.com/item?id=50276188)
- Source: [github.com](https://github.com/grimmerk/memsaga)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T16:21:26Z

## Translation

Title: Show HN: Memsaga – Git history of Claude Code memory and skills, by session
Article title: GitHub - grimmerk/memsaga: Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own
[truncated]
Description: Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. m
[truncated]

Article text:
GitHub - grimmerk/memsaga: Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. mem + saga, Old Norse for “story” or “history”. · GitHub
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
grimmerk
memsaga Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. mem + saga, Old Norse for “story” or “history”.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
218 Commits 218 Commits Folders and files
.github/ workflows .github/ workflows crates crates demo demo docs docs skills skills .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml INSTALL-WITH-CLAUDE.md INSTALL-WITH-CLAUDE.md LICENSE LICENSE README.md README.md dist-workspace.toml dist-workspace.toml View all files Repository files navigation
Version control and two-way sync for the files your AI coding agent works
with: what it remembers, how you configure it, and anything else you choose.
Website: memsaga.grimmer.io ·
Install ( with Claude Code or
by hand ) · Do you need it? ·
FAQ ·
Discussions
Claude Code keeps two kinds of memory, both plain Markdown files on your
machine ( Anthropic's docs ):
Two ways a memory gets written:
you: Remember that the API tests need a local Redis.
Claude: writes memory/api_tests.md, and adds a line for it to MEMORY.md
Claude Code: Saved 1 memory
you: No, we only deploy on Fridays, from a git tag.
Claude: decides this is worth keeping, writes memory/deploy_rules.md
Claude Code: Saved 1 memory
Claude decides what is worth keeping, so not every session saves something.
A session next week loads MEMORY.md , sees the line for deploy_rules.md ,
and reads that file before it answers how to deploy. Claude also reads your
skills from ~/.claude/ , and Claude Code runs the hooks settings.json points
at.
Auto memory and your global ~/.claude/CLAUDE.md have no history, and they
live only on the machine that wrote them (a CLAUDE.md inside a repository is
versioned with that repository). A later session can rewrite
deploy_rules.md , two sessions that edit the same file on one day leave no
trace of the earlier version, and a clean-up that merges ten memories into
three cannot be undone.
memsaga keeps them as history:
① Claude writes or changes a memory, CLAUDE.md or a skill
│
② at the end of the turn (by default; it can commit after each edit
instead), a memsaga hook copies what changed into your data repository
and commits it, named after the conversation
│
③ with an `origin` remote (on GitHub, for example), it is pushed there: by
default at most once every ten minutes and at the end of a session, or
now with `memsaga push`; your other machines pull it, and an edit made on
another machine or on GitHub reaches the live file
What goes into the data repository
memsaga keeps one git repository you own, the data repository, and fills it
from three places, told apart by where a file lives:
Claude Code's config directory ( ~/.claude ). What Claude Code reads
and writes there: memory, CLAUDE.md , skills, hooks, scripts and (if you
want) settings.json , plus folders of your own such as drafts or
hand-written mods. They are mirrored without changing where they live or how
Claude Code writes them.
Files elsewhere that you choose. A notes file, a plan in your home
directory, a dotfile: whoever edits it, you or Claude, the change is
committed.
Content that lives only in the repository. It is a normal git
repository: add notes, companion documents (your notes, translations or
drafts about other repositories), anything else; edit and push from
anywhere. memsaga commits what you add and never deletes it.
Each commit names the session that made the change, by title and by ID, so
you can go from any change back to the conversation that caused it, and from a
conversation to everything it changed. One repository can be shared by several
machines and Claude Code accounts; an edit made in one place (another machine,
GitHub's web editor, a phone) reaches the live files everywhere else.
Status: early, in daily use by its author on macOS; Linux is tested in CI and
in a container.
Claude Code is the only source today; other agents and sources are on the
roadmap . Contributor notes: docs/architecture.md .
Two-way sync: docs/two-way-sync.md . Design record:
docs/design.md .
Quick start : with Claude Code or by hand
What you can do with it
Claude Code's config directory
See what Claude remembered, and when it changed
Undo a bad edit or a memory clean-up
Start with the history you already have
Files elsewhere, and content only in the repository
5. Version any file Claude edits for you
6. Add your own files to the repository
7. Companion docs: notes beside a repository, never committed to it
Everywhere you work
8. Two Claude Code accounts
9. Two machines
10. Edit from anywhere
Going further
11. Let Claude reorganise its memory, safely
What memsaga does not cover, and what to use instead
FAQ : other ways to keep this history, compared
Reference : commands , what is mirrored , memory directory names , configuration , two-way sync , commit messages , updating , uninstalling , setting up another machine
Pictures: the image at the top (where files come from, and where the history
goes), and a sequence diagram in use case 10 (what happens
when a file changed elsewhere).
Maybe not. If you use Claude Code with one account on one computer, a
do-it-yourself setup gets you most of the history:
Turn ~/.claude into a git repository, a folder whose history git keeps
( git init ).
Add a .gitignore that lets in only what you want kept (memory,
CLAUDE.md , skills) and keeps out conversations and credentials.
Add a hook that commits at the end of every reply (a Stop hook).
memsaga does this for you, out of the box. There is no list or hook to write
yourself: it already knows which files Claude Code reads or runs ( CLAUDE.md ,
memory, skills, subagents, commands, rules, hooks), following Claude Code's own
directory reference, and it leaves out credentials, conversations and caches.
It also handles what the do-it-yourself setup cannot easily handle:
A separate repository, away from the files Claude is using. In the
do-it-yourself setup, Claude Code's own folder is the repository, so a
single git command run there ( git checkout , git clean ) or a merge
conflict changes the files Claude is reading. memsaga keeps its copies in a
normal folder you choose (by default memories in the folder where you keep
code), not in the hidden ~/.claude , with an index.md that lists every
project by its real folder,
so it is easy to browse in Finder, an editor or Obsidian.
Files you choose, even outside Claude Code's folder. Beyond the files
Claude Code uses, which are covered by default (above), a folder you create
in ~/.claude yourself, such as ~/.claude/drafts/ , takes one line in
dirs . Outside ~/.claude , name each file in extra_files ,
such as a plan in your home directory or another tool's config file, and
every change to it is committed, whether you made it or Claude did
( use case 5 ). Whole folders
outside ~/.claude are not supported yet
( #20 ).
Edits from anywhere, and more than one account or computer. You can fix
a memory in GitHub's web editor, on a phone, or on another computer. In the
do-it-yourself setup, bringing such an edit back is a git pull that writes
straight into the files in use, and a conflict leaves git's markers in a file
Claude reads. memsaga brings the edit to the live file for you: it keeps the
edits made on both sides, and if both changed the same line, it stops and
tells you. It also tells Claude what changed. Every account lives in one
repository, and another computer picks up changes the next time you use
Claude there.
Which conversation made each change. Every commit names the conversation
behind it, and claude --resume reopens that conversation as long as Claude
Code keeps it: 30 days after its last use, unless you raise
cleanupPeriodDays .
History from before you installed it. memsaga rebuilds it from the edits
recorded in those same saved conversations.
Your own documents about other projects. Some documents are about a
project but must not go into its repository, for example notes on your
employer's code. Ask Claude for one (a note, a translation of its docs, or a
draft) and it is kept in memsaga's repository under companion/ , filed by
the project and file it is about, so a later conversation finds it again.
memsaga calls these companion documents
( use case 7 ).
Other ways to do this, compared, and the do-it-yourself setup step by step:
FAQ .
You need macOS or Linux, git and Claude Code. gh (the GitHub command-line
tool) is optional, for creating the GitHub repository. There are two ways to
install: let Claude Code do it , or
install by hand . Both run the same steps.
A second machine for a data repository you already have? Follow
Setting up another machine instead. Running
init there would start a second, separate history.
Paste this prompt into Claude Code:
Install memsaga for me. Follow
https://raw.githubusercontent.com/grimmerk/memsaga/main/INSTALL-WITH-CLAUDE.md
and ask me before each step that changes my machine.
Claude checks what you have, then runs the steps of
Install by hand and asks before each one, including where
the data repository goes. INSTALL-WITH-CLAUDE.md is
the file it follows. It also lists the permission prompts seen in a test run:
none in auto mode, 11 in manual mode. To read each command before anything
runs, install by hand instead.
Step 1 installs a prebuilt binary (Apple silicon and Intel Macs, Linux on
x86_64 and arm64). To build from source instead, with Rust's cargo (Rust
1.88 or later): cargo install memsaga --locked from crates.io, or
cargo install --git https://github.com/grimmerk/memsaga memsaga --locked for
the latest main .
Rather not pipe a script into sh ? Download the archive for your platform from
the Releases page and
put memsaga on your PATH . On macOS a browser marks the download as
quarantined, and the binary is not notarized, so clear the mark once:
xattr -d com.apple.quarantine memsaga .
Run the steps one at a time; each block below is one step.
Install the prebuilt binary into ~/.cargo/bin :
curl --proto ' =http

[truncated]

## Original Extract

Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. m
[truncated]

GitHub - grimmerk/memsaga: Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. mem + saga, Old Norse for “story” or “history”. · GitHub
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
grimmerk
memsaga Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Version control and two-way sync for your AI coding agent's files: Claude Code's memory, skills and settings, files it edits, and companion docs (notes and translations beside a repo you can't commit to). Each change is a commit named for its session, in a git repo you own, synced across machines. mem + saga, Old Norse for “story” or “history”.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
218 Commits 218 Commits Folders and files
.github/ workflows .github/ workflows crates crates demo demo docs docs skills skills .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml INSTALL-WITH-CLAUDE.md INSTALL-WITH-CLAUDE.md LICENSE LICENSE README.md README.md dist-workspace.toml dist-workspace.toml View all files Repository files navigation
Version control and two-way sync for the files your AI coding agent works
with: what it remembers, how you configure it, and anything else you choose.
Website: memsaga.grimmer.io ·
Install ( with Claude Code or
by hand ) · Do you need it? ·
FAQ ·
Discussions
Claude Code keeps two kinds of memory, both plain Markdown files on your
machine ( Anthropic's docs ):
Two ways a memory gets written:
you: Remember that the API tests need a local Redis.
Claude: writes memory/api_tests.md, and adds a line for it to MEMORY.md
Claude Code: Saved 1 memory
you: No, we only deploy on Fridays, from a git tag.
Claude: decides this is worth keeping, writes memory/deploy_rules.md
Claude Code: Saved 1 memory
Claude decides what is worth keeping, so not every session saves something.
A session next week loads MEMORY.md , sees the line for deploy_rules.md ,
and reads that file before it answers how to deploy. Claude also reads your
skills from ~/.claude/ , and Claude Code runs the hooks settings.json points
at.
Auto memory and your global ~/.claude/CLAUDE.md have no history, and they
live only on the machine that wrote them (a CLAUDE.md inside a repository is
versioned with that repository). A later session can rewrite
deploy_rules.md , two sessions that edit the same file on one day leave no
trace of the earlier version, and a clean-up that merges ten memories into
three cannot be undone.
memsaga keeps them as history:
① Claude writes or changes a memory, CLAUDE.md or a skill
│
② at the end of the turn (by default; it can commit after each edit
instead), a memsaga hook copies what changed into your data repository
and commits it, named after the conversation
│
③ with an `origin` remote (on GitHub, for example), it is pushed there: by
default at most once every ten minutes and at the end of a session, or
now with `memsaga push`; your other machines pull it, and an edit made on
another machine or on GitHub reaches the live file
What goes into the data repository
memsaga keeps one git repository you own, the data repository, and fills it
from three places, told apart by where a file lives:
Claude Code's config directory ( ~/.claude ). What Claude Code reads
and writes there: memory, CLAUDE.md , skills, hooks, scripts and (if you
want) settings.json , plus folders of your own such as drafts or
hand-written mods. They are mirrored without changing where they live or how
Claude Code writes them.
Files elsewhere that you choose. A notes file, a plan in your home
directory, a dotfile: whoever edits it, you or Claude, the change is
committed.
Content that lives only in the repository. It is a normal git
repository: add notes, companion documents (your notes, translations or
drafts about other repositories), anything else; edit and push from
anywhere. memsaga commits what you add and never deletes it.
Each commit names the session that made the change, by title and by ID, so
you can go from any change back to the conversation that caused it, and from a
conversation to everything it changed. One repository can be shared by several
machines and Claude Code accounts; an edit made in one place (another machine,
GitHub's web editor, a phone) reaches the live files everywhere else.
Status: early, in daily use by its author on macOS; Linux is tested in CI and
in a container.
Claude Code is the only source today; other agents and sources are on the
roadmap . Contributor notes: docs/architecture.md .
Two-way sync: docs/two-way-sync.md . Design record:
docs/design.md .
Quick start : with Claude Code or by hand
What you can do with it
Claude Code's config directory
See what Claude remembered, and when it changed
Undo a bad edit or a memory clean-up
Start with the history you already have
Files elsewhere, and content only in the repository
5. Version any file Claude edits for you
6. Add your own files to the repository
7. Companion docs: notes beside a repository, never committed to it
Everywhere you work
8. Two Claude Code accounts
9. Two machines
10. Edit from anywhere
Going further
11. Let Claude reorganise its memory, safely
What memsaga does not cover, and what to use instead
FAQ : other ways to keep this history, compared
Reference : commands , what is mirrored , memory directory names , configuration , two-way sync , commit messages , updating , uninstalling , setting up another machine
Pictures: the image at the top (where files come from, and where the history
goes), and a sequence diagram in use case 10 (what happens
when a file changed elsewhere).
Maybe not. If you use Claude Code with one account on one computer, a
do-it-yourself setup gets you most of the history:
Turn ~/.claude into a git repository, a folder whose history git keeps
( git init ).
Add a .gitignore that lets in only what you want kept (memory,
CLAUDE.md , skills) and keeps out conversations and credentials.
Add a hook that commits at the end of every reply (a Stop hook).
memsaga does this for you, out of the box. There is no list or hook to write
yourself: it already knows which files Claude Code reads or runs ( CLAUDE.md ,
memory, skills, subagents, commands, rules, hooks), following Claude Code's own
directory reference, and it leaves out credentials, conversations and caches.
It also handles what the do-it-yourself setup cannot easily handle:
A separate repository, away from the files Claude is using. In the
do-it-yourself setup, Claude Code's own folder is the repository, so a
single git command run there ( git checkout , git clean ) or a merge
conflict changes the files Claude is reading. memsaga keeps its copies in a
normal folder you choose (by default memories in the folder where you keep
code), not in the hidden ~/.claude , with an index.md that lists every
project by its real folder,
so it is easy to browse in Finder, an editor or Obsidian.
Files you choose, even outside Claude Code's folder. Beyond the files
Claude Code uses, which are covered by default (above), a folder you create
in ~/.claude yourself, such as ~/.claude/drafts/ , takes one line in
dirs . Outside ~/.claude , name each file in extra_files ,
such as a plan in your home directory or another tool's config file, and
every change to it is committed, whether you made it or Claude did
( use case 5 ). Whole folders
outside ~/.claude are not supported yet
( #20 ).
Edits from anywhere, and more than one account or computer. You can fix
a memory in GitHub's web editor, on a phone, or on another computer. In the
do-it-yourself setup, bringing such an edit back is a git pull that writes
straight into the files in use, and a conflict leaves git's markers in a file
Claude reads. memsaga brings the edit to the live file for you: it keeps the
edits made on both sides, and if both changed the same line, it stops and
tells you. It also tells Claude what changed. Every account lives in one
repository, and another computer picks up changes the next time you use
Claude there.
Which conversation made each change. Every commit names the conversation
behind it, and claude --resume reopens that conversation as long as Claude
Code keeps it: 30 days after its last use, unless you raise
cleanupPeriodDays .
History from before you installed it. memsaga rebuilds it from the edits
recorded in those same saved conversations.
Your own documents about other projects. Some documents are about a
project but must not go into its repository, for example notes on your
employer's code. Ask Claude for one (a note, a translation of its docs, or a
draft) and it is kept in memsaga's repository under companion/ , filed by
the project and file it is about, so a later conversation finds it again.
memsaga calls these companion documents
( use case 7 ).
Other ways to do this, compared, and the do-it-yourself setup step by step:
FAQ .
You need macOS or Linux, git and Claude Code. gh (the GitHub command-line
tool) is optional, for creating the GitHub repository. There are two ways to
install: let Claude Code do it , or
install by hand . Both run the same steps.
A second machine for a data repository you already have? Follow
Setting up another machine instead. Running
init there would start a second, separate history.
Paste this prompt into Claude Code:
Install memsaga for me. Follow
https://raw.githubusercontent.com/grimmerk/memsaga/main/INSTALL-WITH-CLAUDE.md
and ask me before each step that changes my machine.
Claude checks what you have, then runs the steps of
Install by hand and asks before each one, including where
the data repository goes. INSTALL-WITH-CLAUDE.md is
the file it follows. It also lists the permission prompts seen in a test run:
none in auto mode, 11 in manual mode. To read each command before anything
runs, install by hand instead.
Step 1 installs a prebuilt binary (Apple silicon and Intel Macs, Linux on
x86_64 and arm64). To build from source instead, with Rust's cargo (Rust
1.88 or later): cargo install memsaga --locked from crates.io, or
cargo install --git https://github.com/grimmerk/memsaga memsaga --locked for
the latest main .
Rather not pipe a script into sh ? Download the archive for your platform from
the Releases page and
put memsaga on your PATH . On macOS a browser marks the download as
quarantined, and the binary is not notarized, so clear the mark once:
xattr -d com.apple.quarantine memsaga .
Run the steps one at a time; each block below is one step.
Install the prebuilt binary into ~/.cargo/bin :
curl --proto ' =http

[truncated]
