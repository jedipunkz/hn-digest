---
source: "https://github.com/krites-dev/krites"
hn_url: "https://news.ycombinator.com/item?id=50156263"
title: "Show HN: Krites, Claude Code Plugin that forces your tests to run"
article_title: "GitHub - krites-dev/krites: Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. · GitHub"
image: "https://opengraph.githubassets.com/103c121c83b6bf2711d4037141963ce1e39164d99e858f2acc5c8357187aa6a8/krites-dev/krites"
author: "QuinnLott"
captured_at: "2026-10-19T19:03:57Z"
capture_tool: "hn-digest"
hn_id: 50156263
score: 1
comments: 0
posted_at: "2026-10-19T18:45:47Z"
tags:
  - hacker-news
---

# Show HN: Krites, Claude Code Plugin that forces your tests to run

- HN: [50156263](https://news.ycombinator.com/item?id=50156263)
- Source: [github.com](https://github.com/krites-dev/krites)
- Score: 1
- Comments: 0
- Posted: 2026-10-19T18:45:47Z

## Translation

Title: Show HN: Krites, Claude Code Plugin that forces your tests to run
Article title: GitHub - krites-dev/krites: Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. · GitHub
Description: Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. - krites-dev/krites
HN text: Hi my name is Quinn, and I have been testing and using LLMs for various coding projects for about a year now. And at this point I run all my agents overnight or for even days at a time. And I’ve sunk countless hours and tokens into coming back to a job that's half done with issues and bugs. That's why I built Krites. It blocks the edit tools on the paths you name, runs your tests when your agent tries to stop and blocks the stop until they pass (up to 3 tries by default), and prints a receipt with /krites:receipt that contains commands, exit codes, diff and commit. The known limits along with more in depth documentation are in the repo’s README. Install it with: claude plugin install krites@anthropic-plugin-directory. Comment if you have any questions, suggestions, comments or similar experiences.

Article text:
GitHub - krites-dev/krites: Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. · GitHub
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
krites-dev
krites Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.claude-plugin .claude-plugin .github .github commands commands examples/ receipts examples/ receipts hooks hooks lib lib test test .gitattributes .gitattributes LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cli.js cli.js eslint.config.js eslint.config.js package.json package.json verify-receipt.js verify-receipt.js View all files Repository files navigation
Don't ask your agent to run the tests. Force it.
Krites runs your checks itself and blocks the stop until they pass. It is a Claude Code plugin: every check is a command from your own repo and its exit code, run by the plugin's process, never narrated by the model. Open source, no account, verifies offline. Pronounced KRY-tees.
claude plugin install krites@anthropic-plugin-directory
Requires Claude Code and Node 20 or later. This installs from Anthropic's plugin directory, which Claude Code includes as the built-in marketplace anthropic-plugin-directory ; inside a session, /plugin install krites@anthropic-plugin-directory does the same. To install from this repo instead: claude plugin marketplace add krites-dev/krites , then claude plugin install krites@krites .
/krites:init shows the krites.toml it would write for this repo and the Node version it found. /krites:init-write creates it, or approves a krites.toml you edited yourself.
/krites:verify runs the configured checks now and prints their raw output, nothing else.
/krites:receipt writes the last run as a signed JSON file and a Markdown rendering under receipts/ .
/krites:telemetry-on and /krites:telemetry-off answer the weekly ping question, and /krites:telemetry shows the answer (see License and telemetry ).
Commit krites.toml . Until a repo has one, the plugin checks nothing there and says so once at session start.
A recording of the real hooks on a copy of the plugin's Rust test fixture, with tests/** protected and cargo fmt --check and cargo test as the checks. The agent lines are the tool calls and stops sent to the hooks; src/lib.rs was broken on purpose, then fixed. Every krites line is the text Claude Code hands the agent, word for word, and an empty one means the hook said nothing and the stop went through; ... marks trimmed lines. The receipt was signed with a throwaway key; the published sample is under Receipts .
agent Edit tests/total.rs
krites Blocked: tests/total.rs is protected. Protected paths are human-edited only.
agent (breaks src/lib.rs, then stops)
krites Krites refuted the stop: cargo test
...
test total_sums_the_values ... FAILED
...
assertion `left == right` failed
left: 5
right: 7
...
test result: FAILED. 0 passed; 1 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s
...
Fix it and let the checks run before stopping. Do not describe results; the gate reads them.
agent (fixes src/lib.rs, then stops)
krites
> /krites:receipt
krites receipts/2026-09-25T16-11-37Z-382ebfd1.json
receipts/2026-09-25T16-11-37Z-382ebfd1.md
Adapted from the site's How it works animation: one crystal per check in the sample receipt, with a scripted failure on the last check before it passes.
Protect. The edit tools are denied on paths you name in krites.toml , on krites.toml itself, and outside the repo, before the write lands.
Gate. When the agent tries to stop, Krites runs your checks itself. The first failure blocks it and hands back the command and its last 60 lines.
Receipt. /krites:receipt signs the last run: commands, exit codes, diff and commit. Anyone can verify it offline with verify-receipt.js .
A run ends in one of four verdicts: passed (every check exited 0), refuted (a check failed, or a protected path changed by any tool, Bash included), timed_out (the run passed timeout_seconds and its process tree was killed), or no_checks (no command is configured: the stop goes through, and the receipt says no_checks , never passed ). Krites never says verified : the checks are yours, and nothing here vouches for them.
The gate runs on every stop attempt. A failing check, a timeout or a changed protected path blocks at most max_blocks times in a row (default 3, 1 to 8); after that the stop is allowed, the run is recorded with the reason "attempt budget exhausted", and the next session start says so. Session start means a fresh startup: a resumed or compacted session keeps its baseline and its block count. The count resets when a gate passes, when a stop finds a clean tree at the session baseline, and when a new session starts. A krites.toml that cannot be read, or a repo in which git cannot list what changed, blocks every stop until it is fixed.
A real receipt ( JSON ), the one krites.dev shows. Verify it from the root of this repo:
node verify-receipt.js --key 6b11164d362ea9b5 examples/receipts/2026-09-25T17-42-54Z-4a10c564.json
It prints valid: signed by <fingerprint>, verdict <verdict> or invalid: <why> , and exits 0 or 1. verify-receipt.js is one file with no dependencies. It checks that the file's bytes are the canonical form of the receipt, that the key and signature are canonical base64, that the signature matches, and, with --key , that the signer is the one you expect (the full fingerprint or at least its first 16 characters). The fingerprint is in the footer of the Markdown rendering.
A receipt proves which commands ran, with which exit codes, on which diff, at which commit, under which krites.toml ( config_hash ). It does not prove your tests are good; it proves they ran.
A signature proves the receipt was not altered after signing. It proves who signed only if you pin the key with --key : anyone can make a key and sign a receipt they wrote by hand, and without --key the valid line ends with signer not pinned (pass --key) .
A receipt is evidence from a machine you trust, not proof against the agent that ran on it. What gets signed is .krites/last-run.json , and the key sits in your user config directory ( %APPDATA%\krites , or $XDG_CONFIG_HOME/krites , or ~/.config/krites ), never in the repo: an agent with a shell can rewrite the first before you export, or read the second and sign on its own. /krites:receipt refuses a run file the gate could not have written and refuses when the tree changed since the run, but a carefully forged run file passes, and one can also be carried into the next passing stop. Export receipts from a session you watched, or from CI.
commands and slow list every configured check, not_run names each one that did not run, and env holds the [checks.env] assignments, scrubbed. root_hash is the hash of the repo's path, so a short, guessable path can be recovered from it.
Asking for a receipt twice for one run points at the file already written. After the tree changed, run /krites:verify first.
/krites:init proposes a krites.toml for the repo it finds. A minimal one:
version = 1
[ checks ]
commands = [
" cargo fmt --check " ,
" cargo test " ,
]
timeout_seconds = 300
[ protect ]
globs = [ " tests/** " ]
[checks].commands run in order at every stop; the first failure blocks it. max_blocks sets the attempt budget.
[checks].slow runs only under /krites:verify , after commands and only when every one of them passed. A stop never starts a slow check; its receipt lists them as not run, unless its own commands all passed and the latest recorded run is a passed /krites:verify on the same tree and config with every slow check at exit 0, whose slow results it then carries.
A check runs without a shell, so GOTOOLCHAIN=go1.26.6 make vulncheck sets nothing. Put the variable in [checks.env] ( GOTOOLCHAIN = "go1.26.6" ), and every check gets it, at a stop and under /krites:verify . Names are letters, digits and _ , not starting with a digit, and two names may not differ only in case. Values are used as written, with no $VAR expansion, and may not hold a control character, " , & , | , < , > , ^ , % or ! .
[protect].globs adds to what is always protected: krites.toml , .krites/ , .git/ , and anything outside the repo except this project's Claude Code memory folder in its default place. A memory folder that is a link, or that sits inside a repo, gets no exception.
A change to krites.toml changes its config_hash and needs /krites:init to approve it.
To open folders outside the repo to the edit tools, such as a scratchpad, list them in config.toml in your Krites config directory, not in the repo, so a repo you clone cannot decide where in your home folder the agent may write:
[ protect ]
allow_outside = [ " ~/.claude-work/** " ]
Each entry is a glob that starts with ~/ , / , or a drive letter such as C:/ , matched against the target's real path after links are followed; the entry itself is not resolved, so name the real folder (on macOS, /private/tmp , not /tmp ). A relative entry is an error that denies every edit outside the repo until it is fixed. An entry naming a network share matches nothing. Matching ignores case on Windows and macOS, and ** matches what is under a folder, not the folder itself. The Krites config directory, ~/.claude , ~/.claude*.json and a non-empty $CLAUDE_CONFIG_DIR stay closed whatever the list says. A broad entry such as ~/** still opens shell startup files, which can set CLAUDE_CONFIG_DIR or move the Krites config directory for the next session, so keep entries to the folders your workflow writes.
To stop gating a repo, delete .krites/ and remove krites.toml . If a stop is blocked with "krites.toml is missing from a repo that Krites gated before", for instance after checking out a branch that predates it, restore the file or delete .krites/ .
An install from this repo follows its default branch; Claude Code leaves auto-update off for a third-party marketplace by default, so nothing changes until you update it. Every release is tagged krites--v<version> , the version in .claude-plugin/plugin.json . The release script never moves a pushed tag, but a tag can be moved by hand on GitHub, so pin the commit too when that matters. If you installed from the directory, run claude plugin uninstall krites@anthropic-plugin-directory first, so only one copy runs.
/plugin marketplace add https://github

[truncated]

## Original Extract

Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. - krites-dev/krites

Hi my name is Quinn, and I have been testing and using LLMs for various coding projects for about a year now. And at this point I run all my agents overnight or for even days at a time. And I’ve sunk countless hours and tokens into coming back to a job that's half done with issues and bugs. That's why I built Krites. It blocks the edit tools on the paths you name, runs your tests when your agent tries to stop and blocks the stop until they pass (up to 3 tries by default), and prints a receipt with /krites:receipt that contains commands, exit codes, diff and commit. The known limits along with more in depth documentation are in the repo’s README. Install it with: claude plugin install krites@anthropic-plugin-directory. Comment if you have any questions, suggestions, comments or similar experiences.

GitHub - krites-dev/krites: Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later. · GitHub
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
krites-dev
krites Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Verification plugin for Claude Code: protected tests, a stop gate that runs your checks, signed receipts. Requires Node 20 or later.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.claude-plugin .claude-plugin .github .github commands commands examples/ receipts examples/ receipts hooks hooks lib lib test test .gitattributes .gitattributes LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cli.js cli.js eslint.config.js eslint.config.js package.json package.json verify-receipt.js verify-receipt.js View all files Repository files navigation
Don't ask your agent to run the tests. Force it.
Krites runs your checks itself and blocks the stop until they pass. It is a Claude Code plugin: every check is a command from your own repo and its exit code, run by the plugin's process, never narrated by the model. Open source, no account, verifies offline. Pronounced KRY-tees.
claude plugin install krites@anthropic-plugin-directory
Requires Claude Code and Node 20 or later. This installs from Anthropic's plugin directory, which Claude Code includes as the built-in marketplace anthropic-plugin-directory ; inside a session, /plugin install krites@anthropic-plugin-directory does the same. To install from this repo instead: claude plugin marketplace add krites-dev/krites , then claude plugin install krites@krites .
/krites:init shows the krites.toml it would write for this repo and the Node version it found. /krites:init-write creates it, or approves a krites.toml you edited yourself.
/krites:verify runs the configured checks now and prints their raw output, nothing else.
/krites:receipt writes the last run as a signed JSON file and a Markdown rendering under receipts/ .
/krites:telemetry-on and /krites:telemetry-off answer the weekly ping question, and /krites:telemetry shows the answer (see License and telemetry ).
Commit krites.toml . Until a repo has one, the plugin checks nothing there and says so once at session start.
A recording of the real hooks on a copy of the plugin's Rust test fixture, with tests/** protected and cargo fmt --check and cargo test as the checks. The agent lines are the tool calls and stops sent to the hooks; src/lib.rs was broken on purpose, then fixed. Every krites line is the text Claude Code hands the agent, word for word, and an empty one means the hook said nothing and the stop went through; ... marks trimmed lines. The receipt was signed with a throwaway key; the published sample is under Receipts .
agent Edit tests/total.rs
krites Blocked: tests/total.rs is protected. Protected paths are human-edited only.
agent (breaks src/lib.rs, then stops)
krites Krites refuted the stop: cargo test
...
test total_sums_the_values ... FAILED
...
assertion `left == right` failed
left: 5
right: 7
...
test result: FAILED. 0 passed; 1 failed; 0 ignored; 0 measured; 0 filtered out; finished in 0.00s
...
Fix it and let the checks run before stopping. Do not describe results; the gate reads them.
agent (fixes src/lib.rs, then stops)
krites
> /krites:receipt
krites receipts/2026-09-25T16-11-37Z-382ebfd1.json
receipts/2026-09-25T16-11-37Z-382ebfd1.md
Adapted from the site's How it works animation: one crystal per check in the sample receipt, with a scripted failure on the last check before it passes.
Protect. The edit tools are denied on paths you name in krites.toml , on krites.toml itself, and outside the repo, before the write lands.
Gate. When the agent tries to stop, Krites runs your checks itself. The first failure blocks it and hands back the command and its last 60 lines.
Receipt. /krites:receipt signs the last run: commands, exit codes, diff and commit. Anyone can verify it offline with verify-receipt.js .
A run ends in one of four verdicts: passed (every check exited 0), refuted (a check failed, or a protected path changed by any tool, Bash included), timed_out (the run passed timeout_seconds and its process tree was killed), or no_checks (no command is configured: the stop goes through, and the receipt says no_checks , never passed ). Krites never says verified : the checks are yours, and nothing here vouches for them.
The gate runs on every stop attempt. A failing check, a timeout or a changed protected path blocks at most max_blocks times in a row (default 3, 1 to 8); after that the stop is allowed, the run is recorded with the reason "attempt budget exhausted", and the next session start says so. Session start means a fresh startup: a resumed or compacted session keeps its baseline and its block count. The count resets when a gate passes, when a stop finds a clean tree at the session baseline, and when a new session starts. A krites.toml that cannot be read, or a repo in which git cannot list what changed, blocks every stop until it is fixed.
A real receipt ( JSON ), the one krites.dev shows. Verify it from the root of this repo:
node verify-receipt.js --key 6b11164d362ea9b5 examples/receipts/2026-09-25T17-42-54Z-4a10c564.json
It prints valid: signed by <fingerprint>, verdict <verdict> or invalid: <why> , and exits 0 or 1. verify-receipt.js is one file with no dependencies. It checks that the file's bytes are the canonical form of the receipt, that the key and signature are canonical base64, that the signature matches, and, with --key , that the signer is the one you expect (the full fingerprint or at least its first 16 characters). The fingerprint is in the footer of the Markdown rendering.
A receipt proves which commands ran, with which exit codes, on which diff, at which commit, under which krites.toml ( config_hash ). It does not prove your tests are good; it proves they ran.
A signature proves the receipt was not altered after signing. It proves who signed only if you pin the key with --key : anyone can make a key and sign a receipt they wrote by hand, and without --key the valid line ends with signer not pinned (pass --key) .
A receipt is evidence from a machine you trust, not proof against the agent that ran on it. What gets signed is .krites/last-run.json , and the key sits in your user config directory ( %APPDATA%\krites , or $XDG_CONFIG_HOME/krites , or ~/.config/krites ), never in the repo: an agent with a shell can rewrite the first before you export, or read the second and sign on its own. /krites:receipt refuses a run file the gate could not have written and refuses when the tree changed since the run, but a carefully forged run file passes, and one can also be carried into the next passing stop. Export receipts from a session you watched, or from CI.
commands and slow list every configured check, not_run names each one that did not run, and env holds the [checks.env] assignments, scrubbed. root_hash is the hash of the repo's path, so a short, guessable path can be recovered from it.
Asking for a receipt twice for one run points at the file already written. After the tree changed, run /krites:verify first.
/krites:init proposes a krites.toml for the repo it finds. A minimal one:
version = 1
[ checks ]
commands = [
" cargo fmt --check " ,
" cargo test " ,
]
timeout_seconds = 300
[ protect ]
globs = [ " tests/** " ]
[checks].commands run in order at every stop; the first failure blocks it. max_blocks sets the attempt budget.
[checks].slow runs only under /krites:verify , after commands and only when every one of them passed. A stop never starts a slow check; its receipt lists them as not run, unless its own commands all passed and the latest recorded run is a passed /krites:verify on the same tree and config with every slow check at exit 0, whose slow results it then carries.
A check runs without a shell, so GOTOOLCHAIN=go1.26.6 make vulncheck sets nothing. Put the variable in [checks.env] ( GOTOOLCHAIN = "go1.26.6" ), and every check gets it, at a stop and under /krites:verify . Names are letters, digits and _ , not starting with a digit, and two names may not differ only in case. Values are used as written, with no $VAR expansion, and may not hold a control character, " , & , | , < , > , ^ , % or ! .
[protect].globs adds to what is always protected: krites.toml , .krites/ , .git/ , and anything outside the repo except this project's Claude Code memory folder in its default place. A memory folder that is a link, or that sits inside a repo, gets no exception.
A change to krites.toml changes its config_hash and needs /krites:init to approve it.
To open folders outside the repo to the edit tools, such as a scratchpad, list them in config.toml in your Krites config directory, not in the repo, so a repo you clone cannot decide where in your home folder the agent may write:
[ protect ]
allow_outside = [ " ~/.claude-work/** " ]
Each entry is a glob that starts with ~/ , / , or a drive letter such as C:/ , matched against the target's real path after links are followed; the entry itself is not resolved, so name the real folder (on macOS, /private/tmp , not /tmp ). A relative entry is an error that denies every edit outside the repo until it is fixed. An entry naming a network share matches nothing. Matching ignores case on Windows and macOS, and ** matches what is under a folder, not the folder itself. The Krites config directory, ~/.claude , ~/.claude*.json and a non-empty $CLAUDE_CONFIG_DIR stay closed whatever the list says. A broad entry such as ~/** still opens shell startup files, which can set CLAUDE_CONFIG_DIR or move the Krites config directory for the next session, so keep entries to the folders your workflow writes.
To stop gating a repo, delete .krites/ and remove krites.toml . If a stop is blocked with "krites.toml is missing from a repo that Krites gated before", for instance after checking out a branch that predates it, restore the file or delete .krites/ .
An install from this repo follows its default branch; Claude Code leaves auto-update off for a third-party marketplace by default, so nothing changes until you update it. Every release is tagged krites--v<version> , the version in .claude-plugin/plugin.json . The release script never moves a pushed tag, but a tag can be moved by hand on GitHub, so pin the commit too when that matters. If you installed from the directory, run claude plugin uninstall krites@anthropic-plugin-directory first, so only one copy runs.
/plugin marketplace add https://github

[truncated]
