---
source: "https://github.com/jerichosiahaya/snip"
hn_url: "https://news.ycombinator.com/item?id=49936702"
title: "Show HN: Snip – Terminal notes app that Claude can read and write via MCP"
article_title: "GitHub - jerichosiahaya/snip: Lightweight Rust terminal note-taking app with SQLite full-text search · GitHub"
image: "https://opengraph.githubassets.com/2ecf493c3df65990aef192c03c83e47fb620a5d010fd0660d3c2943d991bcf31/jerichosiahaya/snip"
author: "ikigaineo"
captured_at: "2026-10-02T19:01:09Z"
capture_tool: "hn-digest"
hn_id: 49936702
score: 1
comments: 0
posted_at: "2026-10-02T18:12:13Z"
tags:
  - hacker-news
---

# Show HN: Snip – Terminal notes app that Claude can read and write via MCP

- HN: [49936702](https://news.ycombinator.com/item?id=49936702)
- Source: [github.com](https://github.com/jerichosiahaya/snip)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T18:12:13Z

## Translation

Title: Show HN: Snip – Terminal notes app that Claude can read and write via MCP
Article title: GitHub - jerichosiahaya/snip: Lightweight Rust terminal note-taking app with SQLite full-text search · GitHub
Description: Lightweight Rust terminal note-taking app with SQLite full-text search - jerichosiahaya/snip

Article text:
GitHub - jerichosiahaya/snip: Lightweight Rust terminal note-taking app with SQLite full-text search · GitHub
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
jerichosiahaya
/
snip
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.github .github packaging/ aur packaging/ aur src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md server.json server.json View all files Repository files navigation
Lightweight terminal note-taking app. Rust + SQLite (FTS5 full-text search),
with an MCP server so Claude can read and write your notes.
mcp-name: io.github.jerichosiahaya/snip
Install from the AUR (Arch Linux)
yay -S snip-notes-bin # prebuilt binary from the GitHub release
yay -S snip-notes # or build from source
Either installs the snip command (use paru or makepkg if you prefer).
The PKGBUILDs live in packaging/aur/ . Each release
updates them and pushes them to the AUR automatically (see
.github/workflows/aur-publish.yml ; it needs the AUR_SSH_PRIVATE_KEY
repository secret).
Install with Cargo (any platform)
Requires Rust/Cargo and a C compiler for the bundled SQLite:
cargo install snip-notes # installs the `snip` command into ~/.cargo/bin
Install on Arch Linux (x86-64)
Download the prebuilt executable— no Rust or Cargo required :
curl -fLO https://github.com/jerichosiahaya/snip/releases/latest/download/snip-arch-linux-x86_64.tar.gz
curl -fLO https://github.com/jerichosiahaya/snip/releases/latest/download/SHA256SUMS
sha256sum -c SHA256SUMS && \
tar -xzf snip-arch-linux-x86_64.tar.gz && \
install -Dm755 snip " $HOME /.local/bin/snip "
Run ~/.local/bin/snip , or snip if ~/.local/bin is on your PATH.
Requires an up-to-date Arch Linux x86-64 system ( glibc , gcc-libs ) and
$EDITOR or nano for editing. SQLite is bundled. This build is not for
ARM or guaranteed compatible with older Linux distributions.
Requires Rust/Cargo and a C compiler for bundled SQLite:
cargo build --release --locked
install -Dm755 target/release/snip " $HOME /.local/bin/snip "
Release executable: ~1.6 MB , below the 5 MB budget (excluding system libraries and notes).
Early release: trash/restore is not implemented; deletion is permanent after
confirmation. Keep backups of important notes.
snip # open the interactive TUI
snip add " note text " --tag work # quick capture, first line = title
snip search " meeting " # print matching notes
snip search " " --tag work
snip export ~ /notes # every note as a Markdown file (see below)
Notes are stored in $XDG_DATA_HOME/snip.db (normally ~/.local/share/snip.db ) by default. Override with snip --db /path/to.db ... .
snip / to search 17 notes
──────────────────────────────┬────────────────────────────────────────────
Meeting notes 2h │ Meeting notes
Groceries 3d │ #work · edited 2h ago
Book list 1w │
│ Discuss the new feature with the team…
──────────────────────────────┴────────────────────────────────────────────
↑↓ move ⏎ edit ^N new / search ^Q quit ^T tag Del delete
The list shows the most recently edited notes first. Terminals 72 columns
or wider also show a preview of the selected note; narrower ones show its
tags in the list instead.
↑↓ PgUp PgDn Home End move through notes
Enter edit selected note (opens $EDITOR, falls back to nano)
Ctrl-N new note
/ search as you type (title + body; every word must match,
prefixes ok: "meet" finds "meeting"); Enter keeps it, Esc clears it
Ctrl-T cycle tag filter (all → each tag → all)
Esc clear search and tag filter
Delete delete selected note (asks y/N in the footer)
Ctrl-Q quit
Note format: first line = title, the rest = body. Works in your normal editor;
$EDITOR may include arguments, e.g. EDITOR="code -w" .
snip mcp runs a Model Context Protocol
server on stdin/stdout, so Claude can save, find and update your notes. It
uses the same database as the interactive app, and both can be open at once.
claude mcp add --scope user snip -- " $( command -v snip ) " mcp
Claude Desktop: add this to claude_desktop_config.json (Settings →
Developer → Edit Config), using the absolute path to snip ( command -v snip
prints it: ~/.cargo/bin/snip after cargo install , ~/.local/bin/snip for
the prebuilt download), then restart Claude Desktop:
{
"mcpServers" : {
"snip" : { "command" : " /home/YOU/.local/bin/snip " , "args" : [ " mcp " ] }
}
}
To use a different database, put --db /path/to.db before mcp in either
command.
There is deliberately no delete tool: deleting notes stays something you do
in the app. update_note can still replace a note's body, so keep backups
(below) if Claude edits important notes.
Source of truth: the SQLite DB (WAL mode). Use SQLite's .backup command for a consistent backup, including while the app is open:
sqlite3 " ${XDG_DATA_HOME :- $HOME / .local / share} /snip.db " " .backup 'snip-backup.db' "
Markdown export: snip export [DIR] [--tag NAME] writes one .md file per
note (DIR defaults to ./snip-export ). Each file has YAML front matter that
Obsidian and other Markdown tools read, then the title as a heading and the
body:
---
title : " Meeting notes "
tags : ["work", "q4"]
created : 2026-10-02T16:49:00Z
updated : 2026-10-02T17:12:00Z
snip_id : 12
---
# Meeting notes
- ship the redesign
Files are named after the title ( meeting-notes.md ; the note id is added if
two titles clash) and their modification time is the note's last edit.
Exporting again overwrites those files and leaves anything else in DIR alone,
so a note that was renamed or deleted keeps its old file until you remove it.
cargo test # regression: FTS stays consistent across create/update/delete
About
Lightweight Rust terminal note-taking app with SQLite full-text search
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Lightweight Rust terminal note-taking app with SQLite full-text search - jerichosiahaya/snip

GitHub - jerichosiahaya/snip: Lightweight Rust terminal note-taking app with SQLite full-text search · GitHub
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
jerichosiahaya
/
snip
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.github .github packaging/ aur packaging/ aur src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md server.json server.json View all files Repository files navigation
Lightweight terminal note-taking app. Rust + SQLite (FTS5 full-text search),
with an MCP server so Claude can read and write your notes.
mcp-name: io.github.jerichosiahaya/snip
Install from the AUR (Arch Linux)
yay -S snip-notes-bin # prebuilt binary from the GitHub release
yay -S snip-notes # or build from source
Either installs the snip command (use paru or makepkg if you prefer).
The PKGBUILDs live in packaging/aur/ . Each release
updates them and pushes them to the AUR automatically (see
.github/workflows/aur-publish.yml ; it needs the AUR_SSH_PRIVATE_KEY
repository secret).
Install with Cargo (any platform)
Requires Rust/Cargo and a C compiler for the bundled SQLite:
cargo install snip-notes # installs the `snip` command into ~/.cargo/bin
Install on Arch Linux (x86-64)
Download the prebuilt executable— no Rust or Cargo required :
curl -fLO https://github.com/jerichosiahaya/snip/releases/latest/download/snip-arch-linux-x86_64.tar.gz
curl -fLO https://github.com/jerichosiahaya/snip/releases/latest/download/SHA256SUMS
sha256sum -c SHA256SUMS && \
tar -xzf snip-arch-linux-x86_64.tar.gz && \
install -Dm755 snip " $HOME /.local/bin/snip "
Run ~/.local/bin/snip , or snip if ~/.local/bin is on your PATH.
Requires an up-to-date Arch Linux x86-64 system ( glibc , gcc-libs ) and
$EDITOR or nano for editing. SQLite is bundled. This build is not for
ARM or guaranteed compatible with older Linux distributions.
Requires Rust/Cargo and a C compiler for bundled SQLite:
cargo build --release --locked
install -Dm755 target/release/snip " $HOME /.local/bin/snip "
Release executable: ~1.6 MB , below the 5 MB budget (excluding system libraries and notes).
Early release: trash/restore is not implemented; deletion is permanent after
confirmation. Keep backups of important notes.
snip # open the interactive TUI
snip add " note text " --tag work # quick capture, first line = title
snip search " meeting " # print matching notes
snip search " " --tag work
snip export ~ /notes # every note as a Markdown file (see below)
Notes are stored in $XDG_DATA_HOME/snip.db (normally ~/.local/share/snip.db ) by default. Override with snip --db /path/to.db ... .
snip / to search 17 notes
──────────────────────────────┬────────────────────────────────────────────
Meeting notes 2h │ Meeting notes
Groceries 3d │ #work · edited 2h ago
Book list 1w │
│ Discuss the new feature with the team…
──────────────────────────────┴────────────────────────────────────────────
↑↓ move ⏎ edit ^N new / search ^Q quit ^T tag Del delete
The list shows the most recently edited notes first. Terminals 72 columns
or wider also show a preview of the selected note; narrower ones show its
tags in the list instead.
↑↓ PgUp PgDn Home End move through notes
Enter edit selected note (opens $EDITOR, falls back to nano)
Ctrl-N new note
/ search as you type (title + body; every word must match,
prefixes ok: "meet" finds "meeting"); Enter keeps it, Esc clears it
Ctrl-T cycle tag filter (all → each tag → all)
Esc clear search and tag filter
Delete delete selected note (asks y/N in the footer)
Ctrl-Q quit
Note format: first line = title, the rest = body. Works in your normal editor;
$EDITOR may include arguments, e.g. EDITOR="code -w" .
snip mcp runs a Model Context Protocol
server on stdin/stdout, so Claude can save, find and update your notes. It
uses the same database as the interactive app, and both can be open at once.
claude mcp add --scope user snip -- " $( command -v snip ) " mcp
Claude Desktop: add this to claude_desktop_config.json (Settings →
Developer → Edit Config), using the absolute path to snip ( command -v snip
prints it: ~/.cargo/bin/snip after cargo install , ~/.local/bin/snip for
the prebuilt download), then restart Claude Desktop:
{
"mcpServers" : {
"snip" : { "command" : " /home/YOU/.local/bin/snip " , "args" : [ " mcp " ] }
}
}
To use a different database, put --db /path/to.db before mcp in either
command.
There is deliberately no delete tool: deleting notes stays something you do
in the app. update_note can still replace a note's body, so keep backups
(below) if Claude edits important notes.
Source of truth: the SQLite DB (WAL mode). Use SQLite's .backup command for a consistent backup, including while the app is open:
sqlite3 " ${XDG_DATA_HOME :- $HOME / .local / share} /snip.db " " .backup 'snip-backup.db' "
Markdown export: snip export [DIR] [--tag NAME] writes one .md file per
note (DIR defaults to ./snip-export ). Each file has YAML front matter that
Obsidian and other Markdown tools read, then the title as a heading and the
body:
---
title : " Meeting notes "
tags : ["work", "q4"]
created : 2026-10-02T16:49:00Z
updated : 2026-10-02T17:12:00Z
snip_id : 12
---
# Meeting notes
- ship the redesign
Files are named after the title ( meeting-notes.md ; the note id is added if
two titles clash) and their modification time is the note's last edit.
Exporting again overwrites those files and leaves anything else in DIR alone,
so a note that was renamed or deleted keeps its old file until you remove it.
cargo test # regression: FTS stays consistent across create/update/delete
About
Lightweight Rust terminal note-taking app with SQLite full-text search
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
