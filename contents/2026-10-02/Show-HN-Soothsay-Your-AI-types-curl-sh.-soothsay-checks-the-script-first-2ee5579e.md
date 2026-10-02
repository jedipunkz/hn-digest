---
source: "https://github.com/rijuld/soothsay"
hn_url: "https://news.ycombinator.com/item?id=49928435"
title: "Show HN: Soothsay - Your AI types curl | sh. soothsay checks the script first"
article_title: "GitHub - rijuld/soothsay: Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). · GitHub"
image: "https://repository-images.githubusercontent.com/1393933011/34469c0c-dd0b-45d6-ad52-90a83654d7dc"
author: "rijuldahiya"
captured_at: "2026-10-02T00:33:57Z"
capture_tool: "hn-digest"
hn_id: 49928435
score: 1
comments: 0
posted_at: "2026-10-02T00:04:39Z"
tags:
  - hacker-news
---

# Show HN: Soothsay - Your AI types curl | sh. soothsay checks the script first

- HN: [49928435](https://news.ycombinator.com/item?id=49928435)
- Source: [github.com](https://github.com/rijuld/soothsay)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T00:04:39Z

## Translation

Title: Show HN: Soothsay - Your AI types curl | sh. soothsay checks the script first
Article title: GitHub - rijuld/soothsay: Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). · GitHub
Description: Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). - rijuld/soothsay

Article text:
GitHub - rijuld/soothsay: Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). · GitHub
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
rijuld
/
soothsay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
22 Commits 22 Commits Folders and files
.claude-plugin .claude-plugin .github .github plugin plugin src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Read the omens before you curl | sh .
soothsay reads a shell install script and tells you, in plain English, what it's
about to do to your machine before you run it. As a Claude Code hook, it also stops
AI coding agents from piping installers into your shell until you've seen what they do.
curl -fsSL https://example.com/install.sh | soothsay
Guard Claude Code (needs the soothsay binary, see Install ):
/plugin marketplace add rijuld/soothsay
/plugin install soothsay@soothsay
Claude's curl … | sh is then blocked, reviewed, pinned to the exact bytes, and put to
you for approval. How it works .
In 2026 almost every developer tool installs the same way:
curl -fsSL https:// < some-ai-cli > .dev/install.sh | bash
Coding agents, runtimes, package managers and language toolchains all do it. You're
handing a stranger's 2,000-line shell script a root-capable terminal. You'll get a nice
progress bar, but you won't be told that it:
appended three lines to your ~/.zshrc ,
installed a launchd job that runs at every login,
downloaded a second script and piped that into sh too,
stripped macOS quarantine so Gatekeeper never looks at the binary.
Reading the script yourself is the right answer, and nobody does it. soothsay does it
for you in about a millisecond and hands you the parts that matter.
This is real output for Bun's installer ( curl -fsSL https://bun.sh/install | soothsay ),
captured on 2026-09-28:
🔮 soothsay stdin · 326 lines · bash · sha256 04882bf4…3be8
● EDITS YOUR SHELL STARTUP FILES
L214 ● appends to ~/.config/fish/config.fish
L246 ● appends to ~/.zshrc
L293 ● appends to ${bash_config} (looks like your shell profile)
● BLIND SPOTS: SOOTHSAY CAN'T SEE PAST THESE
L177 ● runs ~/.bun/bin/bun: a program soothsay can't see inside (also L197, L229, L261)
↳ ~/.bun/bin/bun completions
FILES IT TOUCHES
~/.bun/bin create
~/.bun/bin/bun.zip download, delete
~/.bun/bin/bun copy
~/.bun/bin/bun-${target} delete
~/.config/fish/config.fish append
~/.zshrc append
${bash_config} append
URLS
download ${bun_uri}
🌤 Mild omens. Typical installer behaviour; skim the notices.
7 notices
2 low-level notes hidden (use -v)
soothsay reads scripts; it doesn't run them. Advisory, not a sandbox.
That's a well-behaved installer. Here's an excerpt of the report for
tests/fixtures/nasty.sh , a deliberately hostile test script:
🔮 soothsay nasty.sh · 39 lines · sh · sha256 d67d5b29…ffe6
▲ RUNS MORE CODE FROM THE INTERNET
L9 ▲ pipes http://updates.example.net/stage2.sh straight into bash as root
L10 ▲ downloads and runs https://example.net/stage3.sh with sh
L11 ▲ downloads and runs https://example.net/env.sh with source
✖ HIDDEN OR ENCODED PAYLOADS
L14 ✖ decodes a hidden payload and pipes it into sh
↳ base64 --decode
✖ TOUCHES SECRETS & CREDENTIALS
L26 ✖ reads ~/.ssh/id_ed25519 and sends it over the network
↳ cat ~/.ssh/id_ed25519
L27 ✖ appends to ~/.ssh/authorized_keys: adds keys that can log into this machine
↳ ssh-ed25519 AAAA attacker@box
L28 ✖ shows a fake dialog asking for your password
↳ osascript -e display dialog "macOS needs your password" default answer "" with hidden answer
L29 ✖ reads the macOS Keychain (find-generic-password)
L24 ▲ reads ~/.ssh
↳ tar czf /tmp/k.tgz ~/.ssh ~/.aws/credentials
L24 ▲ reads ~/.aws/credentials
↳ tar czf /tmp/k.tgz ~/.ssh ~/.aws/credentials
✖ WEAKENS SECURITY SETTINGS
L5 ✖ stops recording shell history
L33 ✖ turns Gatekeeper off for the whole machine
L34 ✖ sets setuid/setgid on /usr/local/lib/helper/run: it will run with its owner's privileges
L36 ✖ redirects into a raw network socket /dev/tcp/10.0.0.1/4444 (classic reverse shell)
L32 ▲ strips macOS quarantine so Gatekeeper won't check the download
↳ xattr -dr com.apple.quarantine /Applications/Helper.app
… (persistence, deletes, root, system writes, network, files and URLs sections cut for length)
☠️ Dark omens. Don't run this unless you understand every red line.
9 dangers · 12 warnings · 8 notices
Install
cargo install --locked soothsay
Or use a prebuilt binary for macOS or Linux (x86_64 and arm64) from the
releases page . Check the hash and
the build attestation before you put it on your PATH :
gh release download v0.1.0 --repo rijuld/soothsay -p SHA256SUMS -p ' *aarch64-apple-darwin* '
sha256sum -c SHA256SUMS --ignore-missing
gh attestation verify soothsay-v0.1.0-aarch64-apple-darwin.tar.gz --repo rijuld/soothsay
tar xzf soothsay-v0.1.0-aarch64-apple-darwin.tar.gz
It's a single small binary with zero dependencies . That's on purpose: a tool you
pipe untrusted scripts into should be small enough to audit in an afternoon. The whole
thing is a few thousand lines of plain Rust.
Every release attaches a SHA256SUMS file and a
build provenance attestation
proving the binaries were built from this repo by its release workflow.
# Pipe a script in
curl -fsSL https://example.com/install.sh | soothsay
# Or point it at a file
soothsay ./install.sh
# Everything, including low-level notes and uncapped lists
soothsay -v install.sh
# Read the report, then decide, then run *exactly the bytes you just read*
curl -fsSL https://sh.rustup.rs | soothsay --run -- -y
--run : review, then run the bytes you reviewed
--run prints the report and asks on your terminal ( /dev/tty , since stdin is the script):
Run these exact bytes (sha256 7d0ea0f8eba7…) with sh? [y/N]
If you say yes, it writes the buffered, already-analyzed bytes to a private temp file
( 0700 ) and runs that. It never re-downloads. That matters because a server can detect
curl | sh and serve different content to a pipe than to a browser .
The script gets your terminal as stdin, so installers that prompt still work, and
soothsay exits with the script's exit code.
In CI: guard your own installer
If you maintain an install.sh , soothsay can keep it honest across PRs:
# .github/workflows/installer.yml
# Pin a version. Never install a security tool from a moving branch.
- run : cargo install --locked soothsay --version 0.1.0
- run : soothsay --deny persistence,remote-exec,obfuscation --fail-on danger install.sh
Flag
Effect
--deny <cats>
exit 1 if any live finding is in these categories (comma-separated, or all )
--fail-on <sev>
exit 1 if any live finding is at least notice / warn / danger
--expect-sha256 <hex>
exit 1 if the input's sha256 isn't exactly this (pin the bytes you reviewed)
--ignore-unreachable
let policy skip findings inside functions soothsay thinks are never called (by default they count)
--json
machine-readable report (findings, files, URLs, functions, sha256)
--run [-y]
run the analyzed bytes after confirming (or with -y , without asking)
--allow-danger
with --run -y : run even with danger findings (otherwise it refuses, exit 1 )
--shell <sh>
interpreter for --run (default: the shebang, else sh )
--no-color
plain output (also honours NO_COLOR ; colour is off when piped)
Exit codes: 0 ok · 1 policy matched · 2 usage or I/O error, or input that isn't a
script soothsay can read (empty, an HTML error page, a non-shell interpreter). In
automation, treat anything other than 0 as "don't run".
Coding agents run curl … | sh too, usually without showing anyone the script. As a
Claude Code PreToolUse hook, soothsay stops
that before it happens.
The easy way is the plugin, which registers the hook for you:
/plugin marketplace add rijuld/soothsay
/plugin install soothsay@soothsay
It finds soothsay on your PATH or in ~/.cargo/bin (or $SOOTHSAY_BIN ). If the
binary is missing, it blocks only commands that look like they run downloaded code, and
says how to install it. To wire the hook up by hand instead:
{
"hooks" : {
"PreToolUse" : [
{
"matcher" : " Bash " ,
"hooks" : [
{ "type" : " command " , "command" : " \" $HOME/.cargo/bin/soothsay \" hook " , "timeout" : 120 }
]
}
]
}
}
Put that in ~/.claude/settings.json (or a project's .claude/settings.json ), with
the absolute path to your soothsay binary. For every Bash command the agent is about
to run, soothsay:
lets it through untouched if it doesn't run code from the network;
otherwise blocks it, downloads the script itself with your curl , reviews it, and
saves the exact bytes under ~/.cache/soothsay/<sha256>.sh . A script with danger
findings is never saved and gets no run instructions;
tells the agent what the script does and how to run those bytes:
soothsay --run --yes --expect-sha256 <sha256> <saved file> ;
when the agent runs that, puts it to you : the hook returns an ask decision
with the review attached, so Claude Code shows you a permission prompt. Consent is
enforced by the harness, not left to the agent's judgement. In permission modes
that don't prompt ( bypassPermissions , auto , dontAsk ), soothsay blocks
instead, and you can run the command yourself.
It follows a download wherever it goes: curl -o i.sh … (or curl -O , wget URL ) in
one command, then sh i.sh , ./i.sh , sh < i.sh , cat i.sh | sh or
soothsay --run i.sh in a later one, and curl … | soothsay --run .
The hook fails closed. An unresolvable URL, a failed download, an HTML error page, a
non-shell script, bad input or a crash all block the command (exit 2 ); Claude Code
lets a command through on any other exit code, so soothsay never uses one. Script
text in the message is escaped and labelled as data, not instructions.
soothsay --check-command '<cmd>' runs the same check for other harnesses: exit 0
if there's nothing to review, 1 with the review on stdout if the command is blocked,
3 if it runs reviewed bytes and a human should approve.
Fail-open cases. If the hook times out, or the soothsay binary isn't at the
path in your settings (the shell exits 127 ), Claude Code lets the command through
to its normal permission flow. Use an absolute path, and check it works with
echo nope | "$HOME/.cargo/bin/soothsay" hook; echo $? (it should print 2 ).
What the script downloads. It reviews the script, not the binaries the script
fetches and runs; those stay blind spots in the review.
Paths it can't line up. A cd followed by a relative path may not match an
earlier download.
$ soothsay --categories
remote-exec pipes a

[truncated]

## Original Extract

Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). - rijuld/soothsay

GitHub - rijuld/soothsay: Read the omens before you `curl | sh`: explains what an install script will do before it runs, and stops AI coding agents running unreviewed installers (Claude Code hook). · GitHub
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
rijuld
/
soothsay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
22 Commits 22 Commits Folders and files
.claude-plugin .claude-plugin .github .github plugin plugin src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Read the omens before you curl | sh .
soothsay reads a shell install script and tells you, in plain English, what it's
about to do to your machine before you run it. As a Claude Code hook, it also stops
AI coding agents from piping installers into your shell until you've seen what they do.
curl -fsSL https://example.com/install.sh | soothsay
Guard Claude Code (needs the soothsay binary, see Install ):
/plugin marketplace add rijuld/soothsay
/plugin install soothsay@soothsay
Claude's curl … | sh is then blocked, reviewed, pinned to the exact bytes, and put to
you for approval. How it works .
In 2026 almost every developer tool installs the same way:
curl -fsSL https:// < some-ai-cli > .dev/install.sh | bash
Coding agents, runtimes, package managers and language toolchains all do it. You're
handing a stranger's 2,000-line shell script a root-capable terminal. You'll get a nice
progress bar, but you won't be told that it:
appended three lines to your ~/.zshrc ,
installed a launchd job that runs at every login,
downloaded a second script and piped that into sh too,
stripped macOS quarantine so Gatekeeper never looks at the binary.
Reading the script yourself is the right answer, and nobody does it. soothsay does it
for you in about a millisecond and hands you the parts that matter.
This is real output for Bun's installer ( curl -fsSL https://bun.sh/install | soothsay ),
captured on 2026-09-28:
🔮 soothsay stdin · 326 lines · bash · sha256 04882bf4…3be8
● EDITS YOUR SHELL STARTUP FILES
L214 ● appends to ~/.config/fish/config.fish
L246 ● appends to ~/.zshrc
L293 ● appends to ${bash_config} (looks like your shell profile)
● BLIND SPOTS: SOOTHSAY CAN'T SEE PAST THESE
L177 ● runs ~/.bun/bin/bun: a program soothsay can't see inside (also L197, L229, L261)
↳ ~/.bun/bin/bun completions
FILES IT TOUCHES
~/.bun/bin create
~/.bun/bin/bun.zip download, delete
~/.bun/bin/bun copy
~/.bun/bin/bun-${target} delete
~/.config/fish/config.fish append
~/.zshrc append
${bash_config} append
URLS
download ${bun_uri}
🌤 Mild omens. Typical installer behaviour; skim the notices.
7 notices
2 low-level notes hidden (use -v)
soothsay reads scripts; it doesn't run them. Advisory, not a sandbox.
That's a well-behaved installer. Here's an excerpt of the report for
tests/fixtures/nasty.sh , a deliberately hostile test script:
🔮 soothsay nasty.sh · 39 lines · sh · sha256 d67d5b29…ffe6
▲ RUNS MORE CODE FROM THE INTERNET
L9 ▲ pipes http://updates.example.net/stage2.sh straight into bash as root
L10 ▲ downloads and runs https://example.net/stage3.sh with sh
L11 ▲ downloads and runs https://example.net/env.sh with source
✖ HIDDEN OR ENCODED PAYLOADS
L14 ✖ decodes a hidden payload and pipes it into sh
↳ base64 --decode
✖ TOUCHES SECRETS & CREDENTIALS
L26 ✖ reads ~/.ssh/id_ed25519 and sends it over the network
↳ cat ~/.ssh/id_ed25519
L27 ✖ appends to ~/.ssh/authorized_keys: adds keys that can log into this machine
↳ ssh-ed25519 AAAA attacker@box
L28 ✖ shows a fake dialog asking for your password
↳ osascript -e display dialog "macOS needs your password" default answer "" with hidden answer
L29 ✖ reads the macOS Keychain (find-generic-password)
L24 ▲ reads ~/.ssh
↳ tar czf /tmp/k.tgz ~/.ssh ~/.aws/credentials
L24 ▲ reads ~/.aws/credentials
↳ tar czf /tmp/k.tgz ~/.ssh ~/.aws/credentials
✖ WEAKENS SECURITY SETTINGS
L5 ✖ stops recording shell history
L33 ✖ turns Gatekeeper off for the whole machine
L34 ✖ sets setuid/setgid on /usr/local/lib/helper/run: it will run with its owner's privileges
L36 ✖ redirects into a raw network socket /dev/tcp/10.0.0.1/4444 (classic reverse shell)
L32 ▲ strips macOS quarantine so Gatekeeper won't check the download
↳ xattr -dr com.apple.quarantine /Applications/Helper.app
… (persistence, deletes, root, system writes, network, files and URLs sections cut for length)
☠️ Dark omens. Don't run this unless you understand every red line.
9 dangers · 12 warnings · 8 notices
Install
cargo install --locked soothsay
Or use a prebuilt binary for macOS or Linux (x86_64 and arm64) from the
releases page . Check the hash and
the build attestation before you put it on your PATH :
gh release download v0.1.0 --repo rijuld/soothsay -p SHA256SUMS -p ' *aarch64-apple-darwin* '
sha256sum -c SHA256SUMS --ignore-missing
gh attestation verify soothsay-v0.1.0-aarch64-apple-darwin.tar.gz --repo rijuld/soothsay
tar xzf soothsay-v0.1.0-aarch64-apple-darwin.tar.gz
It's a single small binary with zero dependencies . That's on purpose: a tool you
pipe untrusted scripts into should be small enough to audit in an afternoon. The whole
thing is a few thousand lines of plain Rust.
Every release attaches a SHA256SUMS file and a
build provenance attestation
proving the binaries were built from this repo by its release workflow.
# Pipe a script in
curl -fsSL https://example.com/install.sh | soothsay
# Or point it at a file
soothsay ./install.sh
# Everything, including low-level notes and uncapped lists
soothsay -v install.sh
# Read the report, then decide, then run *exactly the bytes you just read*
curl -fsSL https://sh.rustup.rs | soothsay --run -- -y
--run : review, then run the bytes you reviewed
--run prints the report and asks on your terminal ( /dev/tty , since stdin is the script):
Run these exact bytes (sha256 7d0ea0f8eba7…) with sh? [y/N]
If you say yes, it writes the buffered, already-analyzed bytes to a private temp file
( 0700 ) and runs that. It never re-downloads. That matters because a server can detect
curl | sh and serve different content to a pipe than to a browser .
The script gets your terminal as stdin, so installers that prompt still work, and
soothsay exits with the script's exit code.
In CI: guard your own installer
If you maintain an install.sh , soothsay can keep it honest across PRs:
# .github/workflows/installer.yml
# Pin a version. Never install a security tool from a moving branch.
- run : cargo install --locked soothsay --version 0.1.0
- run : soothsay --deny persistence,remote-exec,obfuscation --fail-on danger install.sh
Flag
Effect
--deny <cats>
exit 1 if any live finding is in these categories (comma-separated, or all )
--fail-on <sev>
exit 1 if any live finding is at least notice / warn / danger
--expect-sha256 <hex>
exit 1 if the input's sha256 isn't exactly this (pin the bytes you reviewed)
--ignore-unreachable
let policy skip findings inside functions soothsay thinks are never called (by default they count)
--json
machine-readable report (findings, files, URLs, functions, sha256)
--run [-y]
run the analyzed bytes after confirming (or with -y , without asking)
--allow-danger
with --run -y : run even with danger findings (otherwise it refuses, exit 1 )
--shell <sh>
interpreter for --run (default: the shebang, else sh )
--no-color
plain output (also honours NO_COLOR ; colour is off when piped)
Exit codes: 0 ok · 1 policy matched · 2 usage or I/O error, or input that isn't a
script soothsay can read (empty, an HTML error page, a non-shell interpreter). In
automation, treat anything other than 0 as "don't run".
Coding agents run curl … | sh too, usually without showing anyone the script. As a
Claude Code PreToolUse hook, soothsay stops
that before it happens.
The easy way is the plugin, which registers the hook for you:
/plugin marketplace add rijuld/soothsay
/plugin install soothsay@soothsay
It finds soothsay on your PATH or in ~/.cargo/bin (or $SOOTHSAY_BIN ). If the
binary is missing, it blocks only commands that look like they run downloaded code, and
says how to install it. To wire the hook up by hand instead:
{
"hooks" : {
"PreToolUse" : [
{
"matcher" : " Bash " ,
"hooks" : [
{ "type" : " command " , "command" : " \" $HOME/.cargo/bin/soothsay \" hook " , "timeout" : 120 }
]
}
]
}
}
Put that in ~/.claude/settings.json (or a project's .claude/settings.json ), with
the absolute path to your soothsay binary. For every Bash command the agent is about
to run, soothsay:
lets it through untouched if it doesn't run code from the network;
otherwise blocks it, downloads the script itself with your curl , reviews it, and
saves the exact bytes under ~/.cache/soothsay/<sha256>.sh . A script with danger
findings is never saved and gets no run instructions;
tells the agent what the script does and how to run those bytes:
soothsay --run --yes --expect-sha256 <sha256> <saved file> ;
when the agent runs that, puts it to you : the hook returns an ask decision
with the review attached, so Claude Code shows you a permission prompt. Consent is
enforced by the harness, not left to the agent's judgement. In permission modes
that don't prompt ( bypassPermissions , auto , dontAsk ), soothsay blocks
instead, and you can run the command yourself.
It follows a download wherever it goes: curl -o i.sh … (or curl -O , wget URL ) in
one command, then sh i.sh , ./i.sh , sh < i.sh , cat i.sh | sh or
soothsay --run i.sh in a later one, and curl … | soothsay --run .
The hook fails closed. An unresolvable URL, a failed download, an HTML error page, a
non-shell script, bad input or a crash all block the command (exit 2 ); Claude Code
lets a command through on any other exit code, so soothsay never uses one. Script
text in the message is escaped and labelled as data, not instructions.
soothsay --check-command '<cmd>' runs the same check for other harnesses: exit 0
if there's nothing to review, 1 with the review on stdout if the command is blocked,
3 if it runs reviewed bytes and a human should approve.
Fail-open cases. If the hook times out, or the soothsay binary isn't at the
path in your settings (the shell exits 127 ), Claude Code lets the command through
to its normal permission flow. Use an absolute path, and check it works with
echo nope | "$HOME/.cargo/bin/soothsay" hook; echo $? (it should print 2 ).
What the script downloads. It reviews the script, not the binaries the script
fetches and runs; those stay blind spots in the review.
Paths it can't line up. A cd followed by a relative path may not match an
earlier download.
$ soothsay --categories
remote-exec pipes a

[truncated]
