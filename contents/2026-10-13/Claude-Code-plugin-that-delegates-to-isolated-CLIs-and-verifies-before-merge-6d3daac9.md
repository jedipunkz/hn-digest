---
source: "https://github.com/PyModel/claude-architect"
hn_url: "https://news.ycombinator.com/item?id=50075321"
title: "Claude Code plugin that delegates to isolated CLIs and verifies before merge"
article_title: "GitHub - PyModel/claude-architect: Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. · GitHub"
image: "https://repository-images.githubusercontent.com/1298757994/221c051b-1b4f-4af7-9348-da8dfb920fe4"
author: "adobe"
captured_at: "2026-10-13T17:58:17Z"
capture_tool: "hn-digest"
hn_id: 50075321
score: 1
comments: 0
posted_at: "2026-10-13T17:20:50Z"
tags:
  - hacker-news
---

# Claude Code plugin that delegates to isolated CLIs and verifies before merge

- HN: [50075321](https://news.ycombinator.com/item?id=50075321)
- Source: [github.com](https://github.com/PyModel/claude-architect)
- Score: 1
- Comments: 0
- Posted: 2026-10-13T17:20:50Z

## Translation

Title: Claude Code plugin that delegates to isolated CLIs and verifies before merge
Article title: GitHub - PyModel/claude-architect: Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. · GitHub
Description: Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. - PyModel/claude-architect

Article text:
GitHub - PyModel/claude-architect: Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. · GitHub
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
PyModel
claude-architect Public Notifications You must be signed in to change notification settings
Star 31 ( 31 ) You must be signed in to star a repository
Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves.
github.com/PyModel/claude-architect Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
458 Commits 458 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .githooks .githooks .github .github agents agents assets assets docs docs native native runtime runtime scripts scripts skills skills src src tests tests vendor/ superpowers vendor/ superpowers .coderabbit.yaml .coderabbit.yaml .gitignore .gitignore .mcp.json .mcp.json .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md esbuild.config.mjs esbuild.config.mjs opencode.json opencode.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Verified coding-agent delegation for Claude Code. Claude stays the architect and reviewer: it writes the spec, judges the evidence, and reports what landed. Implementation goes to a fresh-context implementer on the coding CLI you choose ( Codex, OpenCode, Pi, Pythinker, Antigravity CLI, or a headless Claude Code session ), each invocation starting clean inside an isolated Git worktree. The work comes back as a frozen, hash-anchored candidate that Claude reviews against independent verification evidence before a single byte can reach your checkout.
Three guarantees are enforced in host code, not in prompts:
Isolation. Every Producer runs in a detached worktree with a sanitized environment, an explicit write allowlist, and OS sandboxing where certified. Out-of-scope changes are rejected at freeze time.
Evidence over claims. A Producer saying "tests pass" is never accepted; the runtime reruns your authorized verification commands in a clean worktree and records the real output.
Separated authority. Implementers cannot approve their own work. Review, decision, and hash-gated integration are separate steps, and integration stages the reviewed tree without committing it. Only a reviewed candidate carrying commit-bound gate clearance is accepted without prompting; anything less requires a human ( decision authority ).
Public beta. Do not use Claude Architect unattended for production, destructive, or security-sensitive work. Review the complete candidate and verification evidence before integration. Producer availability depends on host OS, CLI version, authentication, requested lane, and proven execution capabilities.
flowchart LR
A[Versioned spec] --> B[Producer in isolated worktree]
B --> C[Frozen candidate]
C --> D[Independent verification]
D --> E[Adversarial review]
E --> F{Configured decision authority}
F -->|accept| G[Guarded integration]
F -->|reject or revise| H[Discard or fresh attempt]
Loading
All agent output is an untrusted candidate. A candidate that fails independent verification, or whose review gate refused it, can only be accepted by a human.
Claude Code requires Node.js 22 or newer. The runtime needs Git 2.40 or newer: review diffs read attributes from the trusted base with --attr-source , and doctor reports git-too-old below that.
claude plugin marketplace add PyModel/claude-architect
claude plugin install architect@claude-architect
claude plugin list --json
Restart Claude Code after installing or updating. Install and authenticate at least one supported Producer CLI ( codex , opencode , pi , pythinker , agy , or claude ); Claude Architect reports unavailable lanes rather than silently substituting another agent.
Open Claude Code in a Git repository and name the Producer you want:
/architect:delegate Use Codex to add rate limiting to the public API, run the tests, and show me the independently reviewed candidate before integration.
If no Producer is named, the skill asks you to choose one. Model selection inside a lane is optional and defers to that CLI's configured default:
Non-trivial work runs through the fresh-context review pipeline. Read the exact patch, findings, and verification output before deciding whether to accept.
The direct, unverified Codex CLI lane runs codex exec against your current checkout without an isolated worktree, frozen Candidate Artifact, or independent verification:
/architect:codex Review this checkout with gpt-6.1-sol at high reasoning.
Use it for direct Codex assistance when those controls are not required. Use /architect:delegate when changes need the verified lane and its isolated worktree, frozen Candidate Artifact, independent verification, and guarded integration.
Producers do not inherit host-side Superpowers skills. Each edit attempt is offered only three task-scoped procedures compatible with the trust model: test-driven-development , systematic-debugging , and verification-before-completion , vendored from obra/superpowers 6.2.0 under MIT. Skills that assume nested delegation, self-review, branch acceptance, or an interactive human stay with the architect, the configured decision authority, and the human.
Dispatch a delegation through the delegation-lane agent to watch it as a native Claude Code subagent row instead of a long-running MCP call. The lane agent is a courier whose only tools are delegate and delegatePipeline ; it cannot read the repository, run commands, review, decide, or integrate. Lanes against different repositories run in parallel; lanes against the same repository are serialized by the runtime's repository lock. All reviewable evidence comes from reviewCandidate , and every acceptance is gated on independent verification with its provenance recorded. Details live in the delegate skill.
By default, decideCandidate records accepted without prompting for an independently verified candidate that produces no advisory warnings from a readable archive: either a delegatePipeline candidate carrying a durable pipelineGateCleared record that names the archived candidate commit and does not require a human, or a plain delegate result judged on its independent verification alone. Autonomy also requires that project verification ran under an OS confinement backend, so off macOS every decision needs a human, and a plain delegate candidate that edits its own verification inputs (tests, test or build configuration, dependency manifests) needs one too. Setting CLAUDE_ARCHITECT_JEV=on with a TYPESAFE_API_KEY adds an opt-in Jev screen of the patch that can route an otherwise autonomous acceptance to a human, never the reverse. Gate-refused, review-incomplete, malformed, commit-mismatched, human-required, unverified, or unreadable cases still require a human, as does every non-accept verdict. Set CLAUDE_ARCHITECT_DECISION_AUTHORITY=human to require confirmation for every decision; an unrecognized value fails closed to human with a warning.
This never relaxes the gates themselves: independent verification decides what may be accepted at all, integration refuses any acceptance whose provenance is unknown, and it aborts on a moved HEAD , a dirty tree, or a hash that does not match the reviewed artifact. Every decision records its provenance.
Kind
Name
Purpose
Skill
/architect:delegate
Builds a versioned spec and drives delegation, review, decision, and guarded integration.
Skill
/architect:codex
Runs Codex CLI directly against the current checkout without the verified delegation lifecycle.
Skill
/architect:subagent-driven-delegation
Executes a multi-task plan with the Superpowers subagent-driven-development loop, using a verified Producer as the implementer for every task.
Agent
advisor
Strictly read-only commitment-boundary advisor.
Agent
candidate-reviewer
Read-only Opus reviewer for one frozen candidate: reads the exact bytes through reviewCandidate , returns two verdicts and a recommendation, never decides or integrates.
Agent
delegation-lane
Courier that dispatches one delegation as a native subagent row.
MCP
validateDelegationSpec
Validates a spec without starting a Producer and returns its canonical correlation digest.
MCP
delegate
Runs one validated, isolated, independently verified attempt.
MCP
delegatePipeline
Runs the fresh-context implement / review / repair pipeline.
MCP
reviewCandidate
Returns the exact frozen patch and verification evidence.
MCP
decideCandidate
Records accepted, rejected, or revision-requested.
MCP
integrateCandidate
Applies an accepted hash-matched candidate under safety guards.
MCP
doctor
Reports runtime, Git, platform, and Producer diagnostics.
MCP
gitStatus , gitDiff , gitLog , gitChangedFiles
Bounded, redacted, read-only Git evidence for advisors.
Security and trust model
Authority is separated across roles and artifacts. Producers receive bounded write scope in isolated worktrees. Candidate bytes are frozen and identified by hashes before independent verification. Reviewers operate in fresh context, and read-only roles lack mutation tools. The runtime rejects nested delegation, scope escapes, changed bases, mismatched anchors or trees, and unaccepted candidates. Integration stages reviewed bytes; it does not commit them.
The central rule: all agent output is an untrusted candidate; implementers cannot approve their own work; and only an independently verified pipeline candidate with commit-bound gate clearance can be accepted without a human. Verification reduces risk but does not establish that a change is safe for your particular deployment. See docs/SECURITY_MODEL.md , docs/THREAT_MODEL.md , and docs/TRUST_BOUNDARIES.md .
Permissions and external commands
The plugin starts its MCP server with ${CLAUDE_PLUGIN_ROOT}/runtime/bootstrap.mjs . It may invoke Git, Node.js, configured verification executables, and a selected Producer CLI. Producer processes can edit only through an eligible isolated lane; verification commands are Host-authorized and their confinement and network enforcement is reported honestly. The runtime uses executable-plus-argv invocation, sanitized environments, bounded timeouts, process-tree termination, executa

[truncated]

## Original Extract

Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. - PyModel/claude-architect

GitHub - PyModel/claude-architect: Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves. · GitHub
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
PyModel
claude-architect Public Notifications You must be signed in to change notification settings
Star 31 ( 31 ) You must be signed in to star a repository
Claude Code delegates coding to isolated CLI agents it doesn't trust (Codex, OpenCode, Pi, Pythinker), freezes what they produce, verifies it independently, and merges only what a human approves.
github.com/PyModel/claude-architect Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
458 Commits 458 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .githooks .githooks .github .github agents agents assets assets docs docs native native runtime runtime scripts scripts skills skills src src tests tests vendor/ superpowers vendor/ superpowers .coderabbit.yaml .coderabbit.yaml .gitignore .gitignore .mcp.json .mcp.json .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md esbuild.config.mjs esbuild.config.mjs opencode.json opencode.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Verified coding-agent delegation for Claude Code. Claude stays the architect and reviewer: it writes the spec, judges the evidence, and reports what landed. Implementation goes to a fresh-context implementer on the coding CLI you choose ( Codex, OpenCode, Pi, Pythinker, Antigravity CLI, or a headless Claude Code session ), each invocation starting clean inside an isolated Git worktree. The work comes back as a frozen, hash-anchored candidate that Claude reviews against independent verification evidence before a single byte can reach your checkout.
Three guarantees are enforced in host code, not in prompts:
Isolation. Every Producer runs in a detached worktree with a sanitized environment, an explicit write allowlist, and OS sandboxing where certified. Out-of-scope changes are rejected at freeze time.
Evidence over claims. A Producer saying "tests pass" is never accepted; the runtime reruns your authorized verification commands in a clean worktree and records the real output.
Separated authority. Implementers cannot approve their own work. Review, decision, and hash-gated integration are separate steps, and integration stages the reviewed tree without committing it. Only a reviewed candidate carrying commit-bound gate clearance is accepted without prompting; anything less requires a human ( decision authority ).
Public beta. Do not use Claude Architect unattended for production, destructive, or security-sensitive work. Review the complete candidate and verification evidence before integration. Producer availability depends on host OS, CLI version, authentication, requested lane, and proven execution capabilities.
flowchart LR
A[Versioned spec] --> B[Producer in isolated worktree]
B --> C[Frozen candidate]
C --> D[Independent verification]
D --> E[Adversarial review]
E --> F{Configured decision authority}
F -->|accept| G[Guarded integration]
F -->|reject or revise| H[Discard or fresh attempt]
Loading
All agent output is an untrusted candidate. A candidate that fails independent verification, or whose review gate refused it, can only be accepted by a human.
Claude Code requires Node.js 22 or newer. The runtime needs Git 2.40 or newer: review diffs read attributes from the trusted base with --attr-source , and doctor reports git-too-old below that.
claude plugin marketplace add PyModel/claude-architect
claude plugin install architect@claude-architect
claude plugin list --json
Restart Claude Code after installing or updating. Install and authenticate at least one supported Producer CLI ( codex , opencode , pi , pythinker , agy , or claude ); Claude Architect reports unavailable lanes rather than silently substituting another agent.
Open Claude Code in a Git repository and name the Producer you want:
/architect:delegate Use Codex to add rate limiting to the public API, run the tests, and show me the independently reviewed candidate before integration.
If no Producer is named, the skill asks you to choose one. Model selection inside a lane is optional and defers to that CLI's configured default:
Non-trivial work runs through the fresh-context review pipeline. Read the exact patch, findings, and verification output before deciding whether to accept.
The direct, unverified Codex CLI lane runs codex exec against your current checkout without an isolated worktree, frozen Candidate Artifact, or independent verification:
/architect:codex Review this checkout with gpt-6.1-sol at high reasoning.
Use it for direct Codex assistance when those controls are not required. Use /architect:delegate when changes need the verified lane and its isolated worktree, frozen Candidate Artifact, independent verification, and guarded integration.
Producers do not inherit host-side Superpowers skills. Each edit attempt is offered only three task-scoped procedures compatible with the trust model: test-driven-development , systematic-debugging , and verification-before-completion , vendored from obra/superpowers 6.2.0 under MIT. Skills that assume nested delegation, self-review, branch acceptance, or an interactive human stay with the architect, the configured decision authority, and the human.
Dispatch a delegation through the delegation-lane agent to watch it as a native Claude Code subagent row instead of a long-running MCP call. The lane agent is a courier whose only tools are delegate and delegatePipeline ; it cannot read the repository, run commands, review, decide, or integrate. Lanes against different repositories run in parallel; lanes against the same repository are serialized by the runtime's repository lock. All reviewable evidence comes from reviewCandidate , and every acceptance is gated on independent verification with its provenance recorded. Details live in the delegate skill.
By default, decideCandidate records accepted without prompting for an independently verified candidate that produces no advisory warnings from a readable archive: either a delegatePipeline candidate carrying a durable pipelineGateCleared record that names the archived candidate commit and does not require a human, or a plain delegate result judged on its independent verification alone. Autonomy also requires that project verification ran under an OS confinement backend, so off macOS every decision needs a human, and a plain delegate candidate that edits its own verification inputs (tests, test or build configuration, dependency manifests) needs one too. Setting CLAUDE_ARCHITECT_JEV=on with a TYPESAFE_API_KEY adds an opt-in Jev screen of the patch that can route an otherwise autonomous acceptance to a human, never the reverse. Gate-refused, review-incomplete, malformed, commit-mismatched, human-required, unverified, or unreadable cases still require a human, as does every non-accept verdict. Set CLAUDE_ARCHITECT_DECISION_AUTHORITY=human to require confirmation for every decision; an unrecognized value fails closed to human with a warning.
This never relaxes the gates themselves: independent verification decides what may be accepted at all, integration refuses any acceptance whose provenance is unknown, and it aborts on a moved HEAD , a dirty tree, or a hash that does not match the reviewed artifact. Every decision records its provenance.
Kind
Name
Purpose
Skill
/architect:delegate
Builds a versioned spec and drives delegation, review, decision, and guarded integration.
Skill
/architect:codex
Runs Codex CLI directly against the current checkout without the verified delegation lifecycle.
Skill
/architect:subagent-driven-delegation
Executes a multi-task plan with the Superpowers subagent-driven-development loop, using a verified Producer as the implementer for every task.
Agent
advisor
Strictly read-only commitment-boundary advisor.
Agent
candidate-reviewer
Read-only Opus reviewer for one frozen candidate: reads the exact bytes through reviewCandidate , returns two verdicts and a recommendation, never decides or integrates.
Agent
delegation-lane
Courier that dispatches one delegation as a native subagent row.
MCP
validateDelegationSpec
Validates a spec without starting a Producer and returns its canonical correlation digest.
MCP
delegate
Runs one validated, isolated, independently verified attempt.
MCP
delegatePipeline
Runs the fresh-context implement / review / repair pipeline.
MCP
reviewCandidate
Returns the exact frozen patch and verification evidence.
MCP
decideCandidate
Records accepted, rejected, or revision-requested.
MCP
integrateCandidate
Applies an accepted hash-matched candidate under safety guards.
MCP
doctor
Reports runtime, Git, platform, and Producer diagnostics.
MCP
gitStatus , gitDiff , gitLog , gitChangedFiles
Bounded, redacted, read-only Git evidence for advisors.
Security and trust model
Authority is separated across roles and artifacts. Producers receive bounded write scope in isolated worktrees. Candidate bytes are frozen and identified by hashes before independent verification. Reviewers operate in fresh context, and read-only roles lack mutation tools. The runtime rejects nested delegation, scope escapes, changed bases, mismatched anchors or trees, and unaccepted candidates. Integration stages reviewed bytes; it does not commit them.
The central rule: all agent output is an untrusted candidate; implementers cannot approve their own work; and only an independently verified pipeline candidate with commit-bound gate clearance can be accepted without a human. Verification reduces risk but does not establish that a change is safe for your particular deployment. See docs/SECURITY_MODEL.md , docs/THREAT_MODEL.md , and docs/TRUST_BOUNDARIES.md .
Permissions and external commands
The plugin starts its MCP server with ${CLAUDE_PLUGIN_ROOT}/runtime/bootstrap.mjs . It may invoke Git, Node.js, configured verification executables, and a selected Producer CLI. Producer processes can edit only through an eligible isolated lane; verification commands are Host-authorized and their confinement and network enforcement is reported honestly. The runtime uses executable-plus-argv invocation, sanitized environments, bounded timeouts, process-tree termination, executa

[truncated]
