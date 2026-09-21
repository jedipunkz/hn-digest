---
source: "https://github.com/goulinkh/omp-semantic-policy"
hn_url: "https://news.ycombinator.com/item?id=49788876"
title: "Using Jev as an LLM Linter for OMP (Pi) Agent"
article_title: "GitHub - goulinkh/omp-semantic-policy: Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. · GitHub"
image: "https://opengraph.githubassets.com/317bc7c9ba32c7ccf1f4c7f9725b91f2d0a554d44672c805f2e5255d175a4f1f/goulinkh/omp-semantic-policy"
author: "goulinkh"
captured_at: "2026-09-21T16:15:23Z"
capture_tool: "hn-digest"
hn_id: 49788876
score: 1
comments: 0
posted_at: "2026-09-21T15:52:56Z"
tags:
  - hacker-news
---

# Using Jev as an LLM Linter for OMP (Pi) Agent

- HN: [49788876](https://news.ycombinator.com/item?id=49788876)
- Source: [github.com](https://github.com/goulinkh/omp-semantic-policy)
- Score: 1
- Comments: 0
- Posted: 2026-09-21T15:52:56Z

## Translation

Title: Using Jev as an LLM Linter for OMP (Pi) Agent
Article title: GitHub - goulinkh/omp-semantic-policy: Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. · GitHub
Description: Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. - goulinkh/omp-semantic-policy

Article text:
GitHub - goulinkh/omp-semantic-policy: Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. · GitHub
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
goulinkh
/
omp-semantic-policy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
23 Commits 23 Commits Folders and files
.omp-plugin .omp-plugin dist dist docs docs scripts scripts src src testing testing .gitignore .gitignore README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
An Oh My Pi (OMP) plugin that checks agent actions against your project instructions and authorization. TypeSafe AI provides semantic evaluation when remote checks are enabled.
A policy gate, not a sandbox. Uncertainty is automatically accepted; grounded denials block.
Requires OMP 18.1.19 or newer. Remote evaluation also requires a TypeSafe AI token and explicit consent.
Add the marketplace and install the plugin:
omp plugin marketplace add goulinkh/omp-semantic-policy
omp plugin install omp-semantic-policy@goulinkh
Restart OMP inside a Git worktree, then:
/login typesafe-ai
/policy consent on
/policy status
Projects onboard automatically. Use TYPESAFE_API_KEY instead of /login if preferred. Without credentials or consent, local checks and configured fallback remain active.
Refresh the marketplace catalog before upgrading the installed plugin:
omp plugin marketplace update goulinkh
omp plugin upgrade omp-semantic-policy@goulinkh
The package, marketplace catalog, and prepared release tag use version 0.1.0 ( v0.1.0 ).
/policy status : Show the active snapshot and evaluator state.
/policy coverage : Inspect runtime enforcement coverage.
/policy review : Review the compiled rules.
/policy audit [1–100] : Inspect recent redacted decisions and evidence.
/policy onboard : Refresh the project's policy snapshot.
/policy maintenance : Review an uncertain maintenance proposal.
/policy maintenance approve <action-id> : Approve one exact maintenance retry.
/policy maintenance revoke : Clear the maintenance approval.
Maintenance approval applies to one exact retry, not blanket authorization.
/policy consent on / /policy consent off : Enable or disable remote evaluation.
/logout typesafe-ai : Remove the stored credential.
showStatus (default: true ): Show policy state in the status bar.
showViolationFeedback (default: true ): Show direct-action and headless workflow outcome notifications. Tool denials stay attached to their own result cards; confirmation dialogs identify the exact action.
confirmationDefault (default: approve ): Automatically accept confirmation requests, including ungrounded model denials and unavailable-evaluator fallbacks. Set deny for fail-closed automatic resolution. Grounded policy denials remain blocked.
confirmationThreshold (default: 1 ): Automatic mode, without dialogs, including headless sessions. Set below 1 to enable interactive confirmation at or above that confidence; 0 prompts for every uncertain checked action. Explicitly enabled prompts deny without an interactive UI.
Interactive sessions skip post-response workflow evaluation entirely so extension work cannot race or consume the next draft's first keystroke. Headless session-stop checks remain automatic and use confirmationDefault .
The model's allow-confidence cutoff is 50% , while deny confidence requires 80% . A hard-violation probability of 80% can also trigger denial, but every semantic denial requires a matched applicable rule . Citation-choice confidence of 65% establishes attribution directly. If an otherwise blocking assessment selects a known rule below that cutoff, one additional check can establish that specific violation at 80% probability; competing valid citations need not concentrate on one choice. No selected rule, failed verification, or unresolved evidence does not establish a violation. Automatic mode deliberately accepts uncertainty, not compliance. Deterministic local protections and incomplete-action guards remain blocking. Existing explicit settings are preserved.
Blocked tool cards separate the action, readable rule, source file, and next step. Prohibitions retain their Don't meaning. IDs, probability scores, full provenance, and confirmation diagnostics remain in /policy audit , not the compact card. Execution feedback includes the working directory and environment override names, never their values. If a source is incorrectly scoped, correct it and run /policy onboard ; a different tool is not an approval workaround.
Remote mutation checks include proposed content, bounded original source, and explicit before/after hunks for supported edits. The hunks describe proposed changes, not executed results. Stale anchors, unsupported operations, missing files, and omitted evidence are marked unavailable or partial rather than guessed. Inspection respects local read prohibitions and canonical project boundaries.
Shell checks also include bounded recent user text and the preceding assistant proposal, so a short confirmation or refusal retains its subject. Assistant text is not permission; conversational approval never overrides an absolute prohibition. Recognized credentials are redacted, but arbitrary sensitive text may remain. Automatic acceptance of uncertainty is unchanged: better evidence does not guarantee detection.
enabledToolCalls : A comma-separated allowlist of exact tool names for remote evaluation. An empty value selects all tool calls.
Default: bash,eval,python,write,edit,task,hub,browser,computer,debug .
disabledToolCalls (default: empty): A comma-separated list of exact tool names to exclude from semantic evaluation. Takes precedence over enabledToolCalls ; excluded unknown tools are not rejected for lacking a built-in adapter.
toolOperations (default: empty): Classify custom or MCP tools with comma-separated name=operation entries, for example launchpad=read,launchpad_write=execute . Supported operations are read , write , execute , delegate , network , workflow , internal , and unknown . JSON object syntax is also accepted. Built-in classifications cannot be overridden.
Coverage filters do not bypass grounded local path protections. Completeness checks apply to semantically covered calls; an enabled but unclassified custom tool fails closed. Direct shell, Python, and headless workflow gates remain independent.
Host hooks gate tool calls. Broad execution and restricted subagents are dispatch-only; session-stop workflow enforcement is headless-only because interactive post-response hooks race the next draft.
Nested effects, utility commands, and trusted extensions are not contained.
Local path checks cover a narrow literal-path grammar, not all prose, filesystem races, or unseen LSP effects.
See Architecture for scope, authorization, fallback, and maintenance guarantees and OMP integration for host boundaries.
Development · Roadmap · Design decisions
Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. - goulinkh/omp-semantic-policy

GitHub - goulinkh/omp-semantic-policy: Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation. · GitHub
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
goulinkh
/
omp-semantic-policy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
23 Commits 23 Commits Folders and files
.omp-plugin .omp-plugin dist dist docs docs scripts scripts src src testing testing .gitignore .gitignore README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
An Oh My Pi (OMP) plugin that checks agent actions against your project instructions and authorization. TypeSafe AI provides semantic evaluation when remote checks are enabled.
A policy gate, not a sandbox. Uncertainty is automatically accepted; grounded denials block.
Requires OMP 18.1.19 or newer. Remote evaluation also requires a TypeSafe AI token and explicit consent.
Add the marketplace and install the plugin:
omp plugin marketplace add goulinkh/omp-semantic-policy
omp plugin install omp-semantic-policy@goulinkh
Restart OMP inside a Git worktree, then:
/login typesafe-ai
/policy consent on
/policy status
Projects onboard automatically. Use TYPESAFE_API_KEY instead of /login if preferred. Without credentials or consent, local checks and configured fallback remain active.
Refresh the marketplace catalog before upgrading the installed plugin:
omp plugin marketplace update goulinkh
omp plugin upgrade omp-semantic-policy@goulinkh
The package, marketplace catalog, and prepared release tag use version 0.1.0 ( v0.1.0 ).
/policy status : Show the active snapshot and evaluator state.
/policy coverage : Inspect runtime enforcement coverage.
/policy review : Review the compiled rules.
/policy audit [1–100] : Inspect recent redacted decisions and evidence.
/policy onboard : Refresh the project's policy snapshot.
/policy maintenance : Review an uncertain maintenance proposal.
/policy maintenance approve <action-id> : Approve one exact maintenance retry.
/policy maintenance revoke : Clear the maintenance approval.
Maintenance approval applies to one exact retry, not blanket authorization.
/policy consent on / /policy consent off : Enable or disable remote evaluation.
/logout typesafe-ai : Remove the stored credential.
showStatus (default: true ): Show policy state in the status bar.
showViolationFeedback (default: true ): Show direct-action and headless workflow outcome notifications. Tool denials stay attached to their own result cards; confirmation dialogs identify the exact action.
confirmationDefault (default: approve ): Automatically accept confirmation requests, including ungrounded model denials and unavailable-evaluator fallbacks. Set deny for fail-closed automatic resolution. Grounded policy denials remain blocked.
confirmationThreshold (default: 1 ): Automatic mode, without dialogs, including headless sessions. Set below 1 to enable interactive confirmation at or above that confidence; 0 prompts for every uncertain checked action. Explicitly enabled prompts deny without an interactive UI.
Interactive sessions skip post-response workflow evaluation entirely so extension work cannot race or consume the next draft's first keystroke. Headless session-stop checks remain automatic and use confirmationDefault .
The model's allow-confidence cutoff is 50% , while deny confidence requires 80% . A hard-violation probability of 80% can also trigger denial, but every semantic denial requires a matched applicable rule . Citation-choice confidence of 65% establishes attribution directly. If an otherwise blocking assessment selects a known rule below that cutoff, one additional check can establish that specific violation at 80% probability; competing valid citations need not concentrate on one choice. No selected rule, failed verification, or unresolved evidence does not establish a violation. Automatic mode deliberately accepts uncertainty, not compliance. Deterministic local protections and incomplete-action guards remain blocking. Existing explicit settings are preserved.
Blocked tool cards separate the action, readable rule, source file, and next step. Prohibitions retain their Don't meaning. IDs, probability scores, full provenance, and confirmation diagnostics remain in /policy audit , not the compact card. Execution feedback includes the working directory and environment override names, never their values. If a source is incorrectly scoped, correct it and run /policy onboard ; a different tool is not an approval workaround.
Remote mutation checks include proposed content, bounded original source, and explicit before/after hunks for supported edits. The hunks describe proposed changes, not executed results. Stale anchors, unsupported operations, missing files, and omitted evidence are marked unavailable or partial rather than guessed. Inspection respects local read prohibitions and canonical project boundaries.
Shell checks also include bounded recent user text and the preceding assistant proposal, so a short confirmation or refusal retains its subject. Assistant text is not permission; conversational approval never overrides an absolute prohibition. Recognized credentials are redacted, but arbitrary sensitive text may remain. Automatic acceptance of uncertainty is unchanged: better evidence does not guarantee detection.
enabledToolCalls : A comma-separated allowlist of exact tool names for remote evaluation. An empty value selects all tool calls.
Default: bash,eval,python,write,edit,task,hub,browser,computer,debug .
disabledToolCalls (default: empty): A comma-separated list of exact tool names to exclude from semantic evaluation. Takes precedence over enabledToolCalls ; excluded unknown tools are not rejected for lacking a built-in adapter.
toolOperations (default: empty): Classify custom or MCP tools with comma-separated name=operation entries, for example launchpad=read,launchpad_write=execute . Supported operations are read , write , execute , delegate , network , workflow , internal , and unknown . JSON object syntax is also accepted. Built-in classifications cannot be overridden.
Coverage filters do not bypass grounded local path protections. Completeness checks apply to semantically covered calls; an enabled but unclassified custom tool fails closed. Direct shell, Python, and headless workflow gates remain independent.
Host hooks gate tool calls. Broad execution and restricted subagents are dispatch-only; session-stop workflow enforcement is headless-only because interactive post-response hooks race the next draft.
Nested effects, utility commands, and trusted extensions are not contained.
Local path checks cover a narrow literal-path grammar, not all prose, filesystem races, or unseen LSP effects.
See Architecture for scope, authorization, fallback, and maintenance guarantees and OMP integration for host boundaries.
Development · Roadmap · Design decisions
Semantic policy enforcement for coding agents, beginning with OMP and TypeSafe model-backed evaluation.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
