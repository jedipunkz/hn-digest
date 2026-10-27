---
source: "https://github.com/pgquiles/opencode-muse"
hn_url: "https://news.ycombinator.com/item?id=50267038"
title: "OpenCode-Muse: Use Meta AI's Muse Code Subscription in OpenCode"
article_title: "GitHub - pgquiles/opencode-muse: Muse Code subscription provider plugin for OpenCode · GitHub"
image: "https://opengraph.githubassets.com/0f6d160caf73c29ee78bf43d6abf4b3c25fabfc6bf9acc26910db5d16d0f38ed/pgquiles/opencode-muse"
author: "pgquiles"
captured_at: "2026-10-27T23:32:42Z"
capture_tool: "hn-digest"
hn_id: 50267038
score: 2
comments: 0
posted_at: "2026-10-27T22:56:05Z"
tags:
  - hacker-news
---

# OpenCode-Muse: Use Meta AI's Muse Code Subscription in OpenCode

- HN: [50267038](https://news.ycombinator.com/item?id=50267038)
- Source: [github.com](https://github.com/pgquiles/opencode-muse)
- Score: 2
- Comments: 0
- Posted: 2026-10-27T22:56:05Z

## Translation

Title: OpenCode-Muse: Use Meta AI's Muse Code Subscription in OpenCode
Article title: GitHub - pgquiles/opencode-muse: Muse Code subscription provider plugin for OpenCode · GitHub
Description: Muse Code subscription provider plugin for OpenCode - pgquiles/opencode-muse

Article text:
GitHub - pgquiles/opencode-muse: Muse Code subscription provider plugin for OpenCode · GitHub
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
pgquiles
opencode-muse Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Muse Code subscription provider plugin for OpenCode
www.npmjs.com/package/opencode-muse Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
113 Commits 113 Commits Folders and files
.github/ workflows .github/ workflows docs docs etc/ msp-schema etc/ msp-schema src src test test .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
An OpenCode plugin that lets OpenCode (and therefore
OpenChamber ) use a Muse Code subscription with
auth and inference staying entirely inside the official Muse Code CLI
( muse serve , via @muse-code/sdk ).
No subscription credential is ever extracted, cached, or injected into
OpenCode.
Add the published package to your project-level opencode.jsonc (or
opencode.json ):
{
"$schema" : " https://opencode.ai/config.json " ,
"plugins" : [ " opencode-muse@0.7.0 " ],
}
OpenCode downloads and manages the public package, so users do not need to run
npm install or create an npm account. Sign in with muse login (see below),
restart OpenCode or OpenChamber if it was already running, and choose a Muse
model.
Compatibility fallback: Current OpenCode V2 documentation supports package
entries in plugins , but our earlier probes did not observe config-array
loading on OpenCode 2.1.3; the embedded OpenCode 2.0.32 flow was verified with a
project-local plugin file. If your OpenCode version ignores the package entry,
install the package locally and add this file under .opencode/plugins/ :
npm install opencode-muse
mkdir -p .opencode/plugins
cat > .opencode/plugins/opencode-muse.js << ' EOF '
export { default } from "opencode-muse";
EOF
The locally packed package passed provider discovery and a text-turn smoke
test. Muse models and a text turn were also confirmed in OpenChamber after
restarting it. Verified Muse CLI versions are 1.4.12 and 1.4.13 (the SDK
stays 1.4.12; a newer CLI fingerprint starts with a recorded warning
instead of failing). See
#21 .
Sign in through the official Muse Code CLI before using the provider:
muse login
OpenChamber's current provider UI does not expose OpenCode plugin command
authentication methods, so this plugin does not promise a sign-in button. It
uses the Muse CLI's existing account/configuration through muse serve ; it
never reads Muse credential files, exports tokens, or changes subscription
mode. If the CLI has no usable model catalog, the provider remains a disabled
setup placeholder rather than advertising a model that will fail on use.
The billing lane is a plain process environment variable read once when the
plugin loads (it is not an opencode.json field — probed on OpenCode 2.1.4,
plugin options arrive empty):
# One session on the Standard lane:
OPENCODE_MUSE_CONTRIBUTOR_MODE=off opencode
# Persist it by exporting the variable in your shell profile
# (e.g. ~/.bashrc, ~/.zshrc) before launching OpenCode/OpenChamber.
Restart OpenCode/OpenChamber after changing it; the lane is fixed when the
plugin loads. Anything else fails setup with a clear error instead of
billing the wrong lane. Confirm the requested lane in the startup log and
the billed lane with muse usage . See
#29 .
Use Muse Code models from OpenCode or OpenChamber after adding the package to
your OpenCode plugin configuration.
Keep Muse sign-in and inference inside the official Muse CLI; the plugin
never copies or stores your subscription credentials.
Bill through the Contributor lane by default, like the Muse TUI, with a
documented opt-out to the Standard lane. See Configuration
and #29 .
Continue text conversations in a mapped Muse session, including resuming
supported sessions and replaying available session updates after reconnects.
Attach supported images (PNG, JPEG, GIF, WebP) and videos (MP4, MOV) to a
Muse turn.
Attach audio clips (WAV, MP3, OGG, and anything else the CLI's local
decoder accepts) to a Muse turn: the plugin transcribes them through the
official muse voice transcribe route — voice mode is included in every
subscription plan — and the transcript enters the turn as declared text.
A clip that cannot be transcribed fails the turn loudly instead of being
silently dropped. See
#32 .
Attach PDF files to a Muse turn: the plugin stages the bytes under the
session workspace (generated name, never a user-supplied path) and tells
the model where to read them, so the agent parses them with its own
tools. The staging directory is purged when the provider unloads. See
#33 .
Pick a reasoning effort per model through OpenCode's own controls: an
untouched model runs its catalog default (which rotates over time), Ctrl+T
cycles the concrete variants ( minimal … max , model-dependent, one way
until the model is reselected), and the TUI /effort command ("Switch model
variant", verified end to end) or a #variant model suffix
( opencode run --model muse/…#high ) selects one explicitly. Unknown variants
fail the turn instead of guessing. See
#28 .
Answer Muse's permission prompts conversationally: Muse's exact options are
shown as a lettered list and your next message picks one (fuzzy match on
letter, number, or name). Anything unrecognized is asked once more, then
denied — ambiguity never approves — and your message still runs as a
fresh turn instead of being eaten, so a re-asking model can never trap
the conversation. Queued prompts are answered
first-in-first-out with an echo of what was granted. See
#30 .
Answer Muse's structured option-questions conversationally: the oldest
unanswered question renders as a lettered list and your next message
answers it the same forgiving-match way, with the question echoed back on
success. See #31 .
Work each OpenCode session on its own server-managed worktree instead of
sharing the checkout: concurrent sessions stop serializing at the
workspace lock, and the human merges worktrees back through normal
source control (the bridge never auto-merges). Requires a trusted
workspace and a CLI serving the experimental worktree surface;
otherwise the single-writer lock applies with a notice. Empty
worktrees are removed when the provider unloads; anything with
unique work is left for you. See
#34 .
Stay working across Muse CLI updates: capability checks against the pinned
CLI mean a newer CLI fingerprint starts with a recorded warning instead of
failing, and only a missing/malformed fingerprint or an unsupported
envelope version fails closed. See
#21 .
Let Muse run its own agent and tools, while the plugin prevents overlapping
writes to the same workspace.
Approvals: conversational verdicts skip OpenCode permission rules until
core ships a native verdict channel (tracked in
#22 ). The
conversational flow above stands in meanwhile.
Structured input: single-select option-questions are answered
conversationally; multi-select and option-less questions cancel with a
reason until native structured input lands (tracked in
#6 ).
Files and response formats: The bridge accepts text, PNG/JPEG/GIF/WebP
images, MP4/MOV video, audio clips (transcribed to labeled text), and
PDF files (staged into the workspace for the agent to read) — see
Features.
Strict JSON output and some OpenCode generation settings (such as
temperature and token limits) are not supported; requests that require a
specific OpenCode tool are rejected. Muse's separate CLI JSON-output
commands ( muse exec --json , transcript muse export ) are not exposed
as OpenCode features; muse exec --output-schema answer-shaping stays
unwired because it would fork the turn pipeline (see #35).
Tools and agent instructions: Muse runs its own agent and tools; OpenCode
tools are not executed for a Muse turn. Recognized custom-agent instructions
may be forwarded, but Muse otherwise follows its own workspace rules.
Rewind and session names: Rewind/fork is unavailable on the currently
tested Linux Muse runtime, and custom session names may not work. Start a new
session after rewinding or editing the OpenCode conversation. Rewinding does
not undo files Muse changed. See
#20 .
Recovery after missed events: Basic session resume and cursor replay have
been tested, but recovery from a real Muse event gap has only been tested
with a simulated gap. See
#23 .
Billing: opencode-muse is subscription-only; for pay-as-you-go use
OpenCode's official OpenAI-compatible provider ( metamate ). Only
muse usage confirms the billed lane, and unknown Contributor settings fail
closed instead of guessing. Contributor catalog rows bill the requested lane
like every other row. See Configuration and
#29 .
Simultaneous work: Each OpenCode session gets its own worktree where
the CLI supports it (see Features); where worktrees are unavailable the
plugin allows only one active writer per workspace, so a new turn may be
reported as busy while another is still finishing.
Contributor mode is fully supported and is the default (see
Configuration ). If something labeled "Contributor" fails,
check which provider it belongs to first:
"Muse Spark 1.3 Contributor" never connects after a model switch: that
entry is a retired OpenCode-gateway model ( opencode/muse-spark-1.3-contributor-free ),
not this plugin's model. It no longer resolves anywhere, so OpenCode keeps
retrying and the session appears stuck on "connecting". Pick the model under
provider Muse Code — muse/muse-profile-tbh-model-muse-spark-1.3 —
instead of any historical "contributor-free" entry in recent-model lists. See
#26 .
opencode-claude — a
reference for OpenCode v2 plugin setup, model/session hooks, bridge
lifecycle, and the local-proxy fallback. This plugin uses Muse-native tools;
auth and inference stay in the Muse CLI.
muse-code-sdk — official
Muse Code SDK and the Muse Session Protocol used by this plugin.
Paseo's Muse Code integration — a
protocol-level reference implementation.
opencode-muse-auth is another credential-extraction/direct-API approach; see its issue #12 for its discussion of subscription-term concerns.
opencode-muse-code is an unofficial direct-API alternative. It reads the Muse CLI login, requests an API key from Meta, and sends model requests to Meta's API from OpenCode instead of routing inference through muse serve and MSP. Its subscription eligibility is separate from this plugin's CLI-mediated approach and has not been established by this project.
Official Muse Code documentation describes using Muse models in OpenCode thr

[truncated]

## Original Extract

Muse Code subscription provider plugin for OpenCode - pgquiles/opencode-muse

GitHub - pgquiles/opencode-muse: Muse Code subscription provider plugin for OpenCode · GitHub
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
pgquiles
opencode-muse Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Muse Code subscription provider plugin for OpenCode
www.npmjs.com/package/opencode-muse Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
113 Commits 113 Commits Folders and files
.github/ workflows .github/ workflows docs docs etc/ msp-schema etc/ msp-schema src src test test .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
An OpenCode plugin that lets OpenCode (and therefore
OpenChamber ) use a Muse Code subscription with
auth and inference staying entirely inside the official Muse Code CLI
( muse serve , via @muse-code/sdk ).
No subscription credential is ever extracted, cached, or injected into
OpenCode.
Add the published package to your project-level opencode.jsonc (or
opencode.json ):
{
"$schema" : " https://opencode.ai/config.json " ,
"plugins" : [ " opencode-muse@0.7.0 " ],
}
OpenCode downloads and manages the public package, so users do not need to run
npm install or create an npm account. Sign in with muse login (see below),
restart OpenCode or OpenChamber if it was already running, and choose a Muse
model.
Compatibility fallback: Current OpenCode V2 documentation supports package
entries in plugins , but our earlier probes did not observe config-array
loading on OpenCode 2.1.3; the embedded OpenCode 2.0.32 flow was verified with a
project-local plugin file. If your OpenCode version ignores the package entry,
install the package locally and add this file under .opencode/plugins/ :
npm install opencode-muse
mkdir -p .opencode/plugins
cat > .opencode/plugins/opencode-muse.js << ' EOF '
export { default } from "opencode-muse";
EOF
The locally packed package passed provider discovery and a text-turn smoke
test. Muse models and a text turn were also confirmed in OpenChamber after
restarting it. Verified Muse CLI versions are 1.4.12 and 1.4.13 (the SDK
stays 1.4.12; a newer CLI fingerprint starts with a recorded warning
instead of failing). See
#21 .
Sign in through the official Muse Code CLI before using the provider:
muse login
OpenChamber's current provider UI does not expose OpenCode plugin command
authentication methods, so this plugin does not promise a sign-in button. It
uses the Muse CLI's existing account/configuration through muse serve ; it
never reads Muse credential files, exports tokens, or changes subscription
mode. If the CLI has no usable model catalog, the provider remains a disabled
setup placeholder rather than advertising a model that will fail on use.
The billing lane is a plain process environment variable read once when the
plugin loads (it is not an opencode.json field — probed on OpenCode 2.1.4,
plugin options arrive empty):
# One session on the Standard lane:
OPENCODE_MUSE_CONTRIBUTOR_MODE=off opencode
# Persist it by exporting the variable in your shell profile
# (e.g. ~/.bashrc, ~/.zshrc) before launching OpenCode/OpenChamber.
Restart OpenCode/OpenChamber after changing it; the lane is fixed when the
plugin loads. Anything else fails setup with a clear error instead of
billing the wrong lane. Confirm the requested lane in the startup log and
the billed lane with muse usage . See
#29 .
Use Muse Code models from OpenCode or OpenChamber after adding the package to
your OpenCode plugin configuration.
Keep Muse sign-in and inference inside the official Muse CLI; the plugin
never copies or stores your subscription credentials.
Bill through the Contributor lane by default, like the Muse TUI, with a
documented opt-out to the Standard lane. See Configuration
and #29 .
Continue text conversations in a mapped Muse session, including resuming
supported sessions and replaying available session updates after reconnects.
Attach supported images (PNG, JPEG, GIF, WebP) and videos (MP4, MOV) to a
Muse turn.
Attach audio clips (WAV, MP3, OGG, and anything else the CLI's local
decoder accepts) to a Muse turn: the plugin transcribes them through the
official muse voice transcribe route — voice mode is included in every
subscription plan — and the transcript enters the turn as declared text.
A clip that cannot be transcribed fails the turn loudly instead of being
silently dropped. See
#32 .
Attach PDF files to a Muse turn: the plugin stages the bytes under the
session workspace (generated name, never a user-supplied path) and tells
the model where to read them, so the agent parses them with its own
tools. The staging directory is purged when the provider unloads. See
#33 .
Pick a reasoning effort per model through OpenCode's own controls: an
untouched model runs its catalog default (which rotates over time), Ctrl+T
cycles the concrete variants ( minimal … max , model-dependent, one way
until the model is reselected), and the TUI /effort command ("Switch model
variant", verified end to end) or a #variant model suffix
( opencode run --model muse/…#high ) selects one explicitly. Unknown variants
fail the turn instead of guessing. See
#28 .
Answer Muse's permission prompts conversationally: Muse's exact options are
shown as a lettered list and your next message picks one (fuzzy match on
letter, number, or name). Anything unrecognized is asked once more, then
denied — ambiguity never approves — and your message still runs as a
fresh turn instead of being eaten, so a re-asking model can never trap
the conversation. Queued prompts are answered
first-in-first-out with an echo of what was granted. See
#30 .
Answer Muse's structured option-questions conversationally: the oldest
unanswered question renders as a lettered list and your next message
answers it the same forgiving-match way, with the question echoed back on
success. See #31 .
Work each OpenCode session on its own server-managed worktree instead of
sharing the checkout: concurrent sessions stop serializing at the
workspace lock, and the human merges worktrees back through normal
source control (the bridge never auto-merges). Requires a trusted
workspace and a CLI serving the experimental worktree surface;
otherwise the single-writer lock applies with a notice. Empty
worktrees are removed when the provider unloads; anything with
unique work is left for you. See
#34 .
Stay working across Muse CLI updates: capability checks against the pinned
CLI mean a newer CLI fingerprint starts with a recorded warning instead of
failing, and only a missing/malformed fingerprint or an unsupported
envelope version fails closed. See
#21 .
Let Muse run its own agent and tools, while the plugin prevents overlapping
writes to the same workspace.
Approvals: conversational verdicts skip OpenCode permission rules until
core ships a native verdict channel (tracked in
#22 ). The
conversational flow above stands in meanwhile.
Structured input: single-select option-questions are answered
conversationally; multi-select and option-less questions cancel with a
reason until native structured input lands (tracked in
#6 ).
Files and response formats: The bridge accepts text, PNG/JPEG/GIF/WebP
images, MP4/MOV video, audio clips (transcribed to labeled text), and
PDF files (staged into the workspace for the agent to read) — see
Features.
Strict JSON output and some OpenCode generation settings (such as
temperature and token limits) are not supported; requests that require a
specific OpenCode tool are rejected. Muse's separate CLI JSON-output
commands ( muse exec --json , transcript muse export ) are not exposed
as OpenCode features; muse exec --output-schema answer-shaping stays
unwired because it would fork the turn pipeline (see #35).
Tools and agent instructions: Muse runs its own agent and tools; OpenCode
tools are not executed for a Muse turn. Recognized custom-agent instructions
may be forwarded, but Muse otherwise follows its own workspace rules.
Rewind and session names: Rewind/fork is unavailable on the currently
tested Linux Muse runtime, and custom session names may not work. Start a new
session after rewinding or editing the OpenCode conversation. Rewinding does
not undo files Muse changed. See
#20 .
Recovery after missed events: Basic session resume and cursor replay have
been tested, but recovery from a real Muse event gap has only been tested
with a simulated gap. See
#23 .
Billing: opencode-muse is subscription-only; for pay-as-you-go use
OpenCode's official OpenAI-compatible provider ( metamate ). Only
muse usage confirms the billed lane, and unknown Contributor settings fail
closed instead of guessing. Contributor catalog rows bill the requested lane
like every other row. See Configuration and
#29 .
Simultaneous work: Each OpenCode session gets its own worktree where
the CLI supports it (see Features); where worktrees are unavailable the
plugin allows only one active writer per workspace, so a new turn may be
reported as busy while another is still finishing.
Contributor mode is fully supported and is the default (see
Configuration ). If something labeled "Contributor" fails,
check which provider it belongs to first:
"Muse Spark 1.3 Contributor" never connects after a model switch: that
entry is a retired OpenCode-gateway model ( opencode/muse-spark-1.3-contributor-free ),
not this plugin's model. It no longer resolves anywhere, so OpenCode keeps
retrying and the session appears stuck on "connecting". Pick the model under
provider Muse Code — muse/muse-profile-tbh-model-muse-spark-1.3 —
instead of any historical "contributor-free" entry in recent-model lists. See
#26 .
opencode-claude — a
reference for OpenCode v2 plugin setup, model/session hooks, bridge
lifecycle, and the local-proxy fallback. This plugin uses Muse-native tools;
auth and inference stay in the Muse CLI.
muse-code-sdk — official
Muse Code SDK and the Muse Session Protocol used by this plugin.
Paseo's Muse Code integration — a
protocol-level reference implementation.
opencode-muse-auth is another credential-extraction/direct-API approach; see its issue #12 for its discussion of subscription-term concerns.
opencode-muse-code is an unofficial direct-API alternative. It reads the Muse CLI login, requests an API key from Meta, and sends model requests to Meta's API from OpenCode instead of routing inference through muse serve and MSP. Its subscription eligibility is separate from this plugin's CLI-mediated approach and has not been established by this project.
Official Muse Code documentation describes using Muse models in OpenCode thr

[truncated]
