---
source: "https://github.com/ob4cl/configlens"
hn_url: "https://news.ycombinator.com/item?id=50061910"
title: "ConfigLens: Figure out why your AI tool is ignore your settings"
article_title: "GitHub - ob4cl/configlens: Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. · GitHub"
image: "https://opengraph.githubassets.com/7e37b779cda2447005d6749da197a1205863dd6da2d3205ec146db12eec0ab4b/ob4cl/configlens"
author: "ob4cl"
captured_at: "2026-10-12T20:08:14Z"
capture_tool: "hn-digest"
hn_id: 50061910
score: 1
comments: 0
posted_at: "2026-10-12T19:43:44Z"
tags:
  - hacker-news
---

# ConfigLens: Figure out why your AI tool is ignore your settings

- HN: [50061910](https://news.ycombinator.com/item?id=50061910)
- Source: [github.com](https://github.com/ob4cl/configlens)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T19:43:44Z

## Translation

Title: ConfigLens: Figure out why your AI tool is ignore your settings
Article title: GitHub - ob4cl/configlens: Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. · GitHub
Description: Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. - ob4cl/configlens

Article text:
GitHub - ob4cl/configlens: Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. · GitHub
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
ob4cl
configlens Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated.
ob4cl.github.io/configlens/ Topics
Readme Apache-2.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
.github .github benches benches compatibility compatibility demo demo docs docs fuzz fuzz scripts scripts src src tests tests website website .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Website & interactive demo · Releases · Documentation
Find out why your AI coding tool is ignoring your configuration.
ConfigLens is a small, local, read-only Rust CLI that traces a supported configuration key for
Claude Code and OpenCode across supported configuration sources, and shows the observed
winner — with the file, line and reason for each contribution, including the ones that lost.
It never executes your configuration, never runs a client binary, opens no socket and reads no
credential. Values are redacted unless an adapter explicitly allowlists them.
Pre-release, 0.1.0-beta.3 . A working prototype, not a finished product: both adapters are
fixture-verified rather than runtime-conformant, and unreadable layers are reported as gaps instead
of guessed. Status says exactly what that means.
Prebuilt binaries are attached to every release. Three targets ship today:
Each release also carries a SHA256SUMS file and a CycloneDX SBOM per target. The binaries are
unsigned and there is no installer or package manager; replace <tag> below with the release you
want (currently v0.1.0-beta.3 ).
tag=v0.1.0-beta.3
target=x86_64-unknown-linux-gnu # or aarch64-apple-darwin
base=https://github.com/ob4cl/configlens/releases/download/ $tag
curl -LO $base /configlens- $tag - $target .tar.gz
curl -LO $base /SHA256SUMS
sha256sum -c --ignore-missing SHA256SUMS # macOS: use shasum -a 256 -c
tar xzf configlens- $tag - $target .tar.gz
./configlens- $tag - $target /configlens --version
Windows (PowerShell):
$tag = ' v0.1.0-beta.3 '
$repo = ' ob4cl/configlens '
# With the GitHub CLI:
gh release download $tag -- repo $repo -- pattern ' *windows-msvc.zip '
gh release download $tag -- repo $repo -- pattern ' SHA256SUMS '
# Without it: download both files from the release page in a browser.
$zip = Get-ChildItem * windows - msvc.zip | Select-Object - First 1
$expected = ( Get-Content .\SHA256SUMS | Select-String $zip .Name ).Line.Split()[ 0 ]
$actual = ( Get-FileHash - Algorithm SHA256 $zip .FullName ).Hash.ToLower()
if ( $actual -ne $expected ) { throw ' SHA256 mismatch - do not run this file ' }
Expand-Archive $zip .FullName - DestinationPath unpacked
& .\unpacked\configlens -* \ configlens.exe -- version
Expect configlens 0.1.0-beta.3 . On macOS the archive is not notarized, so Gatekeeper warns on first
run: right-click → Open, or xattr -d com.apple.quarantine <binary> .
Any platform with Rust 1.85 or newer:
cargo install --locked --git https://github.com/ob4cl/configlens
Or from a checkout: cargo build --locked --release , which leaves the binary at
target/release/configlens .
Not available yet: macOS Intel, Linux arm64, crates.io, Homebrew, Scoop/winget. There is no
mandatory curl-to-shell installer, and nothing here contacts the network at run time.
configlens explain /model --client claude --client-version 2.1.296
That prints the observed winner for model with the file, line and reason for every contribution.
A real terminal example shows the full output and how to read it.
You set a model, a permission, an MCP flag. The tool ignores it. You now have five plausible files,
an environment variable, a session flag and a managed policy that might all be involved — and the
tool just says "using the other one". ConfigLens answers the four questions that follow:
Which value won? The observed winner among everything it could read.
Where did each value come from? File, line and column for every contribution, in precedence
order, including the definitions that were replaced or skipped.
Which file should I edit? The winning source is named, so you edit the right one.
Can I paste this into an issue? Reports redact protected values, and --paths share
reduces identifying details; configlens support-report --preview
builds a share-mode report for exactly this. Review the output and the privacy limitations below
before sharing; a redacted report is not a guarantee of anonymity.
Until a package is published, build from this checkout. No client needs to be installed, and no
network access is needed at run time.
cargo build --locked --release
./target/release/configlens --help
Prefer to install the binary into your PATH from the checkout:
cargo install --locked --path .
Run the isolated demo
demo/demo.sh reads no real configuration . It creates a throwaway home and project under a
temporary directory, writes only obviously-fake values, clears the environment, and runs the CLI
against that sandbox. Nothing outside the temporary directory is written, and no network or client
subprocess is used.
cargo build --locked --release
./demo/demo.sh
A real terminal example
Captured from the isolated demo above (synthetic values only). The long coverage line and the two
trailing trust_unknown diagnostics are elided — the tool prints every gap and every diagnostic.
$ configlens explain /model --client claude --client-version 2.1.296
/model = "local-model" (observed)
Observed winner: $PROJECT/.claude/settings.local.json:1:12 "local-model" observed_winner
Losers:
"project-model" $PROJECT/.claude/settings.json:2:12 replaced_by_later_layer
Rule: claude.scalar.layer-override
Coverage: partial — cc.managed.hkcu — registry_not_read (expected_unobservable); … (3 more gaps)
Support: fixture_verified; confidence: evidence_class:fixture_verified
Conclusion: "local-model" wins among the observed applicable sources; effective value withheld.
Next inspection:
- supply the managed settings file with --input-context to include it
- remote/MDM layers are not observable here; re-run where the managed source is readable
Read it as: the project file set model to "project-model" , a later layer replaced it with
"local-model" , and "local-model" is what ConfigLens observed to win. It deliberately does not
call that the runtime value, because a higher layer it cannot read might still change it — see
bounded semantics .
configlens <command> [options]
Command
Example
What it does
clients
configlens clients
List implemented adapters with their tested versions, the layers that are always incomplete, and the clients explicitly unsupported.
scan
configlens scan --client claude
Inventory candidate sources with applicability and readability, plus coverage gaps. No recursive crawl.
explain
configlens explain /model --client claude --client-version 2.1.296
Explain one RFC 6901 JSON Pointer: the ordered trace, the observed winner (or an unresolved result) and coverage gaps.
doctor
configlens doctor --client opencode --format json --strict
Parse / schema / applicability / coverage diagnostics. Never certifies a configuration as safe.
snapshot
configlens snapshot --client claude --output before.json
Write the redacted report JSON to an explicit path. Create-new by default; --overwrite to replace.
diff
configlens diff before.json after.json
Compare two reports from the same client and schema; protected values compare as unknown_redacted .
schema
configlens schema --kind report
Print the bundled report JSON Schema. No network.
support-report
configlens support-report --preview
Build a share-mode report for pasting into an issue. Always share mode (regardless of --paths ). --preview prints it to stdout and writes nothing; --output <PATH> writes it ( 0600 on Unix). There is no upload.
completions
configlens completions --shell bash
Print a shell completion script for bash , zsh or fish , generated from the live command tree. Nothing is written to disk.
Run configlens --help (and configlens <command> --help ) for the current, authoritative list — the
table above covers the core commands.
--client <opencode|claude> is required for resolution. Shared options:
--cwd <PATH> (default . ), --client-version <VERSION> (a claimed version — ConfigLens never
reads a client binary), --format text|json , --strict , --no-env , --input-context <PATH> ,
--color auto|always|never , --paths local|share , --allow-read-root <PATH> (repeatable),
--project-root <PATH> . In addition, a global --quiet (accepted on every subcommand, as
configlens <cmd> --quiet ) suppresses informational stderr notes only — the snapshot
wrote N bytes line and the support-report write/ --preview notes. Errors are always printed to
stderr, stdout stays machine-readable (JSON), and exit codes are unchanged. snapshot adds
--output / --overwrite ; doctor adds
--fail-on warning|error ; support-report takes --preview or --output <PATH> ; supplying neither is a usage
error (exit 2 ), and if both are given --preview wins and nothing is written; completions takes --shell bash|zsh|fish .
Code
Meaning
0
Report produced. Non-fatal findings are allowed.
1
doctor findings at or above --fail-on (default error ).
2
Usage error or incompatible input/report.
3
Requested key/client semantics unsupported, or --strict found an incomplete context.
4
A required file was inaccessible, malformed or over a resource limit.
When several apply, the priority is 2 → 4 → 3 → 1 → 0 .
Three terms carry the weight in every report:
observed winner — the highest-precedence definition among the sources ConfigLens could read.
fixture_verified — the rule was checked against this project's fixtures and upstream
documentation, not against a running client.
coverage gap — a source that exists but could not be read, so the effective value is withheld.
ConfigLens is a static analyser of the inputs it can read. It does not inspect a running session and
it never assumes an unreadable layer is absent. That has three visible consequences:
Observed, not effective. When a higher-priority layer exists but is

[truncated]

## Original Extract

Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. - ob4cl/configlens

GitHub - ob4cl/configlens: Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated. · GitHub
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
ob4cl
configlens Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Trace which Claude Code or OpenCode configuration value actually wins - offline, read-only, and with every redaction and coverage gap stated.
ob4cl.github.io/configlens/ Topics
Readme Apache-2.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
.github .github benches benches compatibility compatibility demo demo docs docs fuzz fuzz scripts scripts src src tests tests website website .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Website & interactive demo · Releases · Documentation
Find out why your AI coding tool is ignoring your configuration.
ConfigLens is a small, local, read-only Rust CLI that traces a supported configuration key for
Claude Code and OpenCode across supported configuration sources, and shows the observed
winner — with the file, line and reason for each contribution, including the ones that lost.
It never executes your configuration, never runs a client binary, opens no socket and reads no
credential. Values are redacted unless an adapter explicitly allowlists them.
Pre-release, 0.1.0-beta.3 . A working prototype, not a finished product: both adapters are
fixture-verified rather than runtime-conformant, and unreadable layers are reported as gaps instead
of guessed. Status says exactly what that means.
Prebuilt binaries are attached to every release. Three targets ship today:
Each release also carries a SHA256SUMS file and a CycloneDX SBOM per target. The binaries are
unsigned and there is no installer or package manager; replace <tag> below with the release you
want (currently v0.1.0-beta.3 ).
tag=v0.1.0-beta.3
target=x86_64-unknown-linux-gnu # or aarch64-apple-darwin
base=https://github.com/ob4cl/configlens/releases/download/ $tag
curl -LO $base /configlens- $tag - $target .tar.gz
curl -LO $base /SHA256SUMS
sha256sum -c --ignore-missing SHA256SUMS # macOS: use shasum -a 256 -c
tar xzf configlens- $tag - $target .tar.gz
./configlens- $tag - $target /configlens --version
Windows (PowerShell):
$tag = ' v0.1.0-beta.3 '
$repo = ' ob4cl/configlens '
# With the GitHub CLI:
gh release download $tag -- repo $repo -- pattern ' *windows-msvc.zip '
gh release download $tag -- repo $repo -- pattern ' SHA256SUMS '
# Without it: download both files from the release page in a browser.
$zip = Get-ChildItem * windows - msvc.zip | Select-Object - First 1
$expected = ( Get-Content .\SHA256SUMS | Select-String $zip .Name ).Line.Split()[ 0 ]
$actual = ( Get-FileHash - Algorithm SHA256 $zip .FullName ).Hash.ToLower()
if ( $actual -ne $expected ) { throw ' SHA256 mismatch - do not run this file ' }
Expand-Archive $zip .FullName - DestinationPath unpacked
& .\unpacked\configlens -* \ configlens.exe -- version
Expect configlens 0.1.0-beta.3 . On macOS the archive is not notarized, so Gatekeeper warns on first
run: right-click → Open, or xattr -d com.apple.quarantine <binary> .
Any platform with Rust 1.85 or newer:
cargo install --locked --git https://github.com/ob4cl/configlens
Or from a checkout: cargo build --locked --release , which leaves the binary at
target/release/configlens .
Not available yet: macOS Intel, Linux arm64, crates.io, Homebrew, Scoop/winget. There is no
mandatory curl-to-shell installer, and nothing here contacts the network at run time.
configlens explain /model --client claude --client-version 2.1.296
That prints the observed winner for model with the file, line and reason for every contribution.
A real terminal example shows the full output and how to read it.
You set a model, a permission, an MCP flag. The tool ignores it. You now have five plausible files,
an environment variable, a session flag and a managed policy that might all be involved — and the
tool just says "using the other one". ConfigLens answers the four questions that follow:
Which value won? The observed winner among everything it could read.
Where did each value come from? File, line and column for every contribution, in precedence
order, including the definitions that were replaced or skipped.
Which file should I edit? The winning source is named, so you edit the right one.
Can I paste this into an issue? Reports redact protected values, and --paths share
reduces identifying details; configlens support-report --preview
builds a share-mode report for exactly this. Review the output and the privacy limitations below
before sharing; a redacted report is not a guarantee of anonymity.
Until a package is published, build from this checkout. No client needs to be installed, and no
network access is needed at run time.
cargo build --locked --release
./target/release/configlens --help
Prefer to install the binary into your PATH from the checkout:
cargo install --locked --path .
Run the isolated demo
demo/demo.sh reads no real configuration . It creates a throwaway home and project under a
temporary directory, writes only obviously-fake values, clears the environment, and runs the CLI
against that sandbox. Nothing outside the temporary directory is written, and no network or client
subprocess is used.
cargo build --locked --release
./demo/demo.sh
A real terminal example
Captured from the isolated demo above (synthetic values only). The long coverage line and the two
trailing trust_unknown diagnostics are elided — the tool prints every gap and every diagnostic.
$ configlens explain /model --client claude --client-version 2.1.296
/model = "local-model" (observed)
Observed winner: $PROJECT/.claude/settings.local.json:1:12 "local-model" observed_winner
Losers:
"project-model" $PROJECT/.claude/settings.json:2:12 replaced_by_later_layer
Rule: claude.scalar.layer-override
Coverage: partial — cc.managed.hkcu — registry_not_read (expected_unobservable); … (3 more gaps)
Support: fixture_verified; confidence: evidence_class:fixture_verified
Conclusion: "local-model" wins among the observed applicable sources; effective value withheld.
Next inspection:
- supply the managed settings file with --input-context to include it
- remote/MDM layers are not observable here; re-run where the managed source is readable
Read it as: the project file set model to "project-model" , a later layer replaced it with
"local-model" , and "local-model" is what ConfigLens observed to win. It deliberately does not
call that the runtime value, because a higher layer it cannot read might still change it — see
bounded semantics .
configlens <command> [options]
Command
Example
What it does
clients
configlens clients
List implemented adapters with their tested versions, the layers that are always incomplete, and the clients explicitly unsupported.
scan
configlens scan --client claude
Inventory candidate sources with applicability and readability, plus coverage gaps. No recursive crawl.
explain
configlens explain /model --client claude --client-version 2.1.296
Explain one RFC 6901 JSON Pointer: the ordered trace, the observed winner (or an unresolved result) and coverage gaps.
doctor
configlens doctor --client opencode --format json --strict
Parse / schema / applicability / coverage diagnostics. Never certifies a configuration as safe.
snapshot
configlens snapshot --client claude --output before.json
Write the redacted report JSON to an explicit path. Create-new by default; --overwrite to replace.
diff
configlens diff before.json after.json
Compare two reports from the same client and schema; protected values compare as unknown_redacted .
schema
configlens schema --kind report
Print the bundled report JSON Schema. No network.
support-report
configlens support-report --preview
Build a share-mode report for pasting into an issue. Always share mode (regardless of --paths ). --preview prints it to stdout and writes nothing; --output <PATH> writes it ( 0600 on Unix). There is no upload.
completions
configlens completions --shell bash
Print a shell completion script for bash , zsh or fish , generated from the live command tree. Nothing is written to disk.
Run configlens --help (and configlens <command> --help ) for the current, authoritative list — the
table above covers the core commands.
--client <opencode|claude> is required for resolution. Shared options:
--cwd <PATH> (default . ), --client-version <VERSION> (a claimed version — ConfigLens never
reads a client binary), --format text|json , --strict , --no-env , --input-context <PATH> ,
--color auto|always|never , --paths local|share , --allow-read-root <PATH> (repeatable),
--project-root <PATH> . In addition, a global --quiet (accepted on every subcommand, as
configlens <cmd> --quiet ) suppresses informational stderr notes only — the snapshot
wrote N bytes line and the support-report write/ --preview notes. Errors are always printed to
stderr, stdout stays machine-readable (JSON), and exit codes are unchanged. snapshot adds
--output / --overwrite ; doctor adds
--fail-on warning|error ; support-report takes --preview or --output <PATH> ; supplying neither is a usage
error (exit 2 ), and if both are given --preview wins and nothing is written; completions takes --shell bash|zsh|fish .
Code
Meaning
0
Report produced. Non-fatal findings are allowed.
1
doctor findings at or above --fail-on (default error ).
2
Usage error or incompatible input/report.
3
Requested key/client semantics unsupported, or --strict found an incomplete context.
4
A required file was inaccessible, malformed or over a resource limit.
When several apply, the priority is 2 → 4 → 3 → 1 → 0 .
Three terms carry the weight in every report:
observed winner — the highest-precedence definition among the sources ConfigLens could read.
fixture_verified — the rule was checked against this project's fixtures and upstream
documentation, not against a running client.
coverage gap — a source that exists but could not be read, so the effective value is withheld.
ConfigLens is a static analyser of the inputs it can read. It does not inspect a running session and
it never assumes an unreadable layer is absent. That has three visible consequences:
Observed, not effective. When a higher-priority layer exists but is

[truncated]
