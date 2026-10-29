---
source: "https://github.com/overflowy/claude-scrub"
hn_url: "https://news.ycombinator.com/item?id=50291192"
title: "Show HN: Claude Scrub - redact secrets from Claude Code history on disk"
article_title: "GitHub - overflowy/claude-scrub: Redact secrets from Claude Code history on disk · GitHub"
image: "https://opengraph.githubassets.com/647dc1ecb004b440c8dc8ef7ab6aea2627a3e067bbf835623a4808654c21bf07/overflowy/claude-scrub"
author: "overflowy"
captured_at: "2026-10-29T16:39:48Z"
capture_tool: "hn-digest"
hn_id: 50291192
score: 1
comments: 0
posted_at: "2026-10-29T16:01:16Z"
tags:
  - hacker-news
---

# Show HN: Claude Scrub - redact secrets from Claude Code history on disk

- HN: [50291192](https://news.ycombinator.com/item?id=50291192)
- Source: [github.com](https://github.com/overflowy/claude-scrub)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T16:01:16Z

## Translation

Title: Show HN: Claude Scrub - redact secrets from Claude Code history on disk
Article title: GitHub - overflowy/claude-scrub: Redact secrets from Claude Code history on disk · GitHub
Description: Redact secrets from Claude Code history on disk. Contribute to overflowy/claude-scrub development by creating an account on GitHub.

Article text:
GitHub - overflowy/claude-scrub: Redact secrets from Claude Code history on disk · GitHub
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
overflowy
claude-scrub Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Redact secrets from Claude Code history on disk
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.gitignore .gitignore README.md README.md Taskfile.yml Taskfile.yml detect.go detect.go files.go files.go go.mod go.mod go.sum go.sum main.go main.go rewrite.go rewrite.go scan.go scan.go scrub.go scrub.go scrub_test.go scrub_test.go secrets.go secrets.go View all files Repository files navigation
Claude Code keeps everything you and the model saw in plain text under
~/.claude . That includes every cat .env , every config file it read and
every token you pasted into a prompt. claude-scrub finds the secrets in that
history and replaces each one with a label:
DB_PASSWORD=[REDACTED:db_password]
It covers session transcripts (subagents and tool results included), prompt
history, the paste cache, file history and shell snapshots. It leaves
settings, plugins and live session state alone.
Download the file for your platform from the
releases page ,
make it executable, and put it on your PATH . With the GitHub CLI, on an
Apple Silicon Mac:
gh release download -R overflowy/claude-scrub -p ' *darwin_arm64 '
chmod +x claude-scrub_ * _darwin_arm64
mv claude-scrub_ * _darwin_arm64 ~ /.local/bin/claude-scrub
The other builds are darwin_amd64 , linux_amd64 , linux_arm64 ,
windows_amd64.exe and windows_arm64.exe . checksums.txt has their
SHA-256 sums. The macOS binaries aren't signed, so if macOS refuses to run
one, clear the quarantine flag with
xattr -d com.apple.quarantine ~/.local/bin/claude-scrub .
go install github.com/overflowy/claude-scrub@latest
This puts the binary in $(go env GOPATH)/bin , usually ~/go/bin .
Needs Go 1.26 or newer and Task .
git clone https://github.com/overflowy/claude-scrub
cd claude-scrub
task install # builds ~/.local/bin/claude-scrub
task --list shows the other tasks. task check runs gofmt, go fix, go vet
and the tests without changing anything.
Start with a dry run. It prints a table of labels with how many values and
occurrences it found, and changes nothing.
claude-scrub --dry-run
To see why something was flagged, ask for the text around each value of a
label. The snippet is redacted too.
claude-scrub -n --explain token
When the report looks right, run it for real. It asks before writing.
claude-scrub
Without a terminal to confirm on, it refuses to write unless you pass
--force .
It skips files modified in the last 10 minutes, so it won't rewrite a session
that's still running. Change that with --min-age . Paths given as arguments
are scanned whatever their age.
Add a SessionEnd hook to ~/.claude/settings.json and each session gets
scrubbed when it ends. Use the path you installed the binary to:
{
"hooks" : {
"SessionEnd" : [
{ "hooks" : [{ "type" : " command " , "command" : " ~/.local/bin/claude-scrub --hook " }] }
]
}
}
The hook redacts that session's transcript, its subagent and tool result
files, and history.jsonl . It prints nothing unless it redacted something
or failed.
Your own secrets and false positives
Put values that must always go in ~/.config/claude-scrub/secrets , one
label=value per line. You can also point --env-file at a .env , shell or
YAML file and it picks out the secrets in it.
Put false positives in ~/.config/claude-scrub/allow , one value per line.
label:NAME skips a whole label.
There are four sources, from most to least trusted:
Your secrets file and any --env-file .
Values assigned to secret-looking names, such as DB_PASSWORD=... ,
"apiKey": "..." or --password=... . The value has to look like a secret:
at least 8 characters, a mix of character types, and not a path, date,
placeholder or piece of code.
Passwords in URLs ( postgres://user:pass@host ) and Authorization: Bearer
or Basic headers.
The gitleaks rules, which know about
200+ token formats (AWS, GitHub, Slack, OpenAI, private keys and more).
Every value found anywhere gets redacted everywhere. If a password shows up
once as DB_PASSWORD=... and later on its own in a psql command, both go.
Transcripts are JSON lines, and the same secret can sit in them as is,
escaped once, or escaped twice when a tool returned JSON. claude-scrub looks
for all three forms and only cuts where the cut can't break an escape
sequence. Any line that would stop being valid JSON is left alone and
reported.
Every other byte stays the same. Each file is written to a temporary file and
renamed into place, keeping its mode and modification time. If a running
session appends to a file while it's being rewritten, claude-scrub notices
and tries again.
A dry run over 1.3 GB of history takes about 5 seconds on an 18-core M5
Pro. Most lines contain none of the keywords any rule needs, so they're
dropped before a regex runs. Very long lines, like a minified bundle a tool
printed, are cut into windows around each keyword.
Detection is best effort. A password that never appears next to a
secret-looking name is missed unless it's in your secrets file. The same
goes for derived values, like a base64-encoded id:key .
--explain and --peek print text from your history. Everything found is
redacted in it, but anything the detectors missed shows as is.
Redaction is permanent. There is no backup, since a backup would hold the
secrets.
Redact secrets from Claude Code history on disk
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Redact secrets from Claude Code history on disk. Contribute to overflowy/claude-scrub development by creating an account on GitHub.

GitHub - overflowy/claude-scrub: Redact secrets from Claude Code history on disk · GitHub
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
overflowy
claude-scrub Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Redact secrets from Claude Code history on disk
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.gitignore .gitignore README.md README.md Taskfile.yml Taskfile.yml detect.go detect.go files.go files.go go.mod go.mod go.sum go.sum main.go main.go rewrite.go rewrite.go scan.go scan.go scrub.go scrub.go scrub_test.go scrub_test.go secrets.go secrets.go View all files Repository files navigation
Claude Code keeps everything you and the model saw in plain text under
~/.claude . That includes every cat .env , every config file it read and
every token you pasted into a prompt. claude-scrub finds the secrets in that
history and replaces each one with a label:
DB_PASSWORD=[REDACTED:db_password]
It covers session transcripts (subagents and tool results included), prompt
history, the paste cache, file history and shell snapshots. It leaves
settings, plugins and live session state alone.
Download the file for your platform from the
releases page ,
make it executable, and put it on your PATH . With the GitHub CLI, on an
Apple Silicon Mac:
gh release download -R overflowy/claude-scrub -p ' *darwin_arm64 '
chmod +x claude-scrub_ * _darwin_arm64
mv claude-scrub_ * _darwin_arm64 ~ /.local/bin/claude-scrub
The other builds are darwin_amd64 , linux_amd64 , linux_arm64 ,
windows_amd64.exe and windows_arm64.exe . checksums.txt has their
SHA-256 sums. The macOS binaries aren't signed, so if macOS refuses to run
one, clear the quarantine flag with
xattr -d com.apple.quarantine ~/.local/bin/claude-scrub .
go install github.com/overflowy/claude-scrub@latest
This puts the binary in $(go env GOPATH)/bin , usually ~/go/bin .
Needs Go 1.26 or newer and Task .
git clone https://github.com/overflowy/claude-scrub
cd claude-scrub
task install # builds ~/.local/bin/claude-scrub
task --list shows the other tasks. task check runs gofmt, go fix, go vet
and the tests without changing anything.
Start with a dry run. It prints a table of labels with how many values and
occurrences it found, and changes nothing.
claude-scrub --dry-run
To see why something was flagged, ask for the text around each value of a
label. The snippet is redacted too.
claude-scrub -n --explain token
When the report looks right, run it for real. It asks before writing.
claude-scrub
Without a terminal to confirm on, it refuses to write unless you pass
--force .
It skips files modified in the last 10 minutes, so it won't rewrite a session
that's still running. Change that with --min-age . Paths given as arguments
are scanned whatever their age.
Add a SessionEnd hook to ~/.claude/settings.json and each session gets
scrubbed when it ends. Use the path you installed the binary to:
{
"hooks" : {
"SessionEnd" : [
{ "hooks" : [{ "type" : " command " , "command" : " ~/.local/bin/claude-scrub --hook " }] }
]
}
}
The hook redacts that session's transcript, its subagent and tool result
files, and history.jsonl . It prints nothing unless it redacted something
or failed.
Your own secrets and false positives
Put values that must always go in ~/.config/claude-scrub/secrets , one
label=value per line. You can also point --env-file at a .env , shell or
YAML file and it picks out the secrets in it.
Put false positives in ~/.config/claude-scrub/allow , one value per line.
label:NAME skips a whole label.
There are four sources, from most to least trusted:
Your secrets file and any --env-file .
Values assigned to secret-looking names, such as DB_PASSWORD=... ,
"apiKey": "..." or --password=... . The value has to look like a secret:
at least 8 characters, a mix of character types, and not a path, date,
placeholder or piece of code.
Passwords in URLs ( postgres://user:pass@host ) and Authorization: Bearer
or Basic headers.
The gitleaks rules, which know about
200+ token formats (AWS, GitHub, Slack, OpenAI, private keys and more).
Every value found anywhere gets redacted everywhere. If a password shows up
once as DB_PASSWORD=... and later on its own in a psql command, both go.
Transcripts are JSON lines, and the same secret can sit in them as is,
escaped once, or escaped twice when a tool returned JSON. claude-scrub looks
for all three forms and only cuts where the cut can't break an escape
sequence. Any line that would stop being valid JSON is left alone and
reported.
Every other byte stays the same. Each file is written to a temporary file and
renamed into place, keeping its mode and modification time. If a running
session appends to a file while it's being rewritten, claude-scrub notices
and tries again.
A dry run over 1.3 GB of history takes about 5 seconds on an 18-core M5
Pro. Most lines contain none of the keywords any rule needs, so they're
dropped before a regex runs. Very long lines, like a minified bundle a tool
printed, are cut into windows around each keyword.
Detection is best effort. A password that never appears next to a
secret-looking name is missed unless it's in your secrets file. The same
goes for derived values, like a base64-encoded id:key .
--explain and --peek print text from your history. Everything found is
redacted in it, but anything the detectors missed shows as is.
Redaction is permanent. There is no backup, since a backup would hold the
secrets.
Redact secrets from Claude Code history on disk
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
