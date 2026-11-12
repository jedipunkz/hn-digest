---
source: "https://github.com/Flurit/applygate"
hn_url: "https://news.ycombinator.com/item?id=50467355"
title: "Show HN: Applygate, stop coding agents from running Terraform destroy in prod"
article_title: "GitHub - Flurit/applygate: A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI,\n[truncated]"
image: "https://repository-images.githubusercontent.com/1436547137/7f105a78-579c-41af-b101-c79ed8682c5b"
author: "sumantth"
captured_at: "2026-11-12T15:01:46Z"
capture_tool: "hn-digest"
hn_id: 50467355
score: 3
comments: 0
posted_at: "2026-11-12T14:08:47Z"
tags:
  - hacker-news
---

# Show HN: Applygate, stop coding agents from running Terraform destroy in prod

- HN: [50467355](https://news.ycombinator.com/item?id=50467355)
- Source: [github.com](https://github.com/Flurit/applygate)
- Score: 3
- Comments: 0
- Posted: 2026-11-12T14:08:47Z

## Translation

Title: Show HN: Applygate, stop coding agents from running Terraform destroy in prod
Article title: GitHub - Flurit/applygate: A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI,
[truncated]
Description: A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devi
[truncated]

Article text:
GitHub - Flurit/applygate: A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu. · GitHub
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
Flurit
applygate Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu.
flurit.github.io/applygate/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
358 Commits 358 Commits Folders and files
.claude-plugin .claude-plugin .github .github cmd cmd docs docs examples/ judge examples/ judge experiments experiments harness harness internal internal packaging/ homebrew packaging/ homebrew plugin/ claude-code plugin/ claude-code review review scripts scripts site site .dockerignore .dockerignore .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md action.yml action.yml go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
Stop coding agents from running terraform destroy in production.
What it does: before an AI coding agent runs a command, applygate checks it and answers allow , ask or block , with a reason the agent can act on. It covers Terraform, OpenTofu and Terragrunt, and the aws, gcloud, az, kubectl and helm CLIs.
Who it is for: teams that let Claude Code, Codex or another coding agent work in infrastructure repositories.
How: it runs on your machine as the agent's pre-tool-call hook. Fixed rules decide, with no model involved by default, and nothing leaves your machine.
What it is not: a sandbox. It catches mistakes and most injected instructions, but an agent that is set on getting around it and holds credentials that can write can still do damage. What it does not protect against .
applygate is pre-1.0, for macOS and Linux.
brew install Flurit/tap/applygate # other ways to install: docs/install.md
cd your-infra-repo
applygate init # .applygate.yaml for this repository, hooks for Claude Code and Codex, a user config if you have none
applygate check " terraform -chdir=envs/prod destroy "
Using Codex? Open it and run /hooks once to trust the new hook. Until you do, Codex skips it without saying so. applygate doctor tells you whether it will run.
Two config files. .applygate.yaml lists the repository's environments; commit it so your team shares it. It can only add protection, and after every edit a person runs applygate trust , so an agent cannot quietly relabel production. Your own settings, such as observe or enforce, live in ~/.config/applygate/config.yaml .
It starts in observe mode. It logs what it would block and blocks nothing. That includes an agent editing or removing applygate itself, so try it where that is acceptable.
Then turn it on. After a few days, applygate report shows what it would have stopped and applygate suggest proposes environments and tiers. Set mode: enforce in your user config and run applygate doctor .
Every command is judged by the environment it targets. You map directories, workspaces, backends, AWS profiles, Kubernetes contexts, GCP projects and Azure subscriptions to a tier in .applygate.yaml ( config ): protected for production, standard or disposable . Anything applygate cannot place counts as protected.
Plan check: for apply <planfile> it reads the plan and rates what it finds as critical or high: destroyed or replaced stateful resources, removed deletion protection, internet exposure, IAM wildcards, state-backend changes. In a protected environment critical findings block and high ones ask; in a standard one critical findings ask. You can mark your own resources as critical ( policy ).
Turning applygate off is blocked in every tier, but only a managed install makes that a guarantee ( below ).
Outside protected environments, applygate still asks about a command that can run code it cannot follow if the same repository holds a protected environment.
applygate reads commands the way the shell does. It follows cd , variables, bash -c , eval , sudo , xargs , Makefiles, scripts and base64 payloads to find what actually runs, so grep "terraform destroy" is just a grep and a destroy hidden in base64 is still a destroy. In a protected environment it also asks about anything that is not a plain command, because only plain commands are fully understood.
On 17 dangerous Terraform scenarios (our own set, so a demonstration, not a measure of coverage):
The hand-written policy also blocked harmless text like git commit -m "never run terraform destroy" . The scenarios are in harness/ , and applygate selftest runs all 163 scenarios against your binary.
Agent
Status
Can show you an ask prompt
If applygate cannot run
Claude Code
Tested live
yes (except in acceptEdits , plan and dontAsk modes)
the call is blocked (Claude Code 2.1.295 or later)
Codex
Tested live
no: an ask is a block
the call runs
GitHub Copilot CLI
Preview
yes
blocked; runs on a timeout
Cursor
Preview
no: an ask is a block
blocked
Gemini CLI
Preview
no: an ask is a block
the call runs
Windsurf (Cascade)
Preview
no: an ask is a block
the call runs
Devin (shell commands only)
Preview
no: an ask is a block
the call runs
OpenCode
Preview
no: an ask is a block
blocked; runs if the plugin fails to load
Cline
Preview
no: an ask is a block (it cancels the task)
blocked; runs on a timeout or a missing hook script
applygate init installs the hooks for Claude Code and Codex.
Preview hooks follow each agent's documentation but have not been tested inside the agent. Install one with applygate hook install --agent NAME . To check it works, ask the agent to run terraform destroy -auto-approve in an empty directory: it should be blocked.
On agents that cannot ask, a person can allow one stopped call with applygate approve ID . This is off by default, because an agent on the same machine can reach your terminal too ( policy ).
On agents where the call runs if applygate cannot start, keep the agent's own sandbox and approval settings on as a second line.
MCP tools: in Claude Code and Codex, calls to known infrastructure MCP servers (AWS, HCP Terraform, Kubernetes, gcloud, az, shell servers) are judged like the equivalent command. This has not yet been tested against the running servers ( mcp ).
Per-agent details: docs/hosts.md , or applygate hosts .
A check that fails (a bad config, unreadable input, an internal error) blocks, unless every directory in the repository is configured as standard or disposable. on_error in your user config changes this.
A decision that takes longer than 25 seconds is blocked.
If applygate cannot start at all, the agent decides: see the table above. applygate doctor checks that the hook is in place.
To stop it blocking, set mode: observe in your user config. To remove it, run applygate hook remove --agent NAME for each agent, then brew uninstall applygate .
What it does not protect against
A hook cannot promise that a command is never allowed. What protects you when it is wrong is a boundary around the agent: credentials that can only plan, writes behind MFA, a versioned and locked state bucket, a host sandbox and a reviewed apply path.
applygate doctor --boundary checks the ones it can and reports an assurance level from L0 (the hook only) to L3 (all of these in place) ( boundary ).
applygate doctor --creds asks your cloud accounts whether the credentials on this machine can change production without a person ( credentials ).
Stopping an agent from turning applygate off is guaranteed only with a managed, root-owned install ( managed ); without one it is best-effort.
Known gaps are in docs/limits.md , and what has been tested live is in docs/capabilities.md .
From your terminal, applygate apply tfplan judges a saved plan the way the hooks do, then runs terraform apply (or tofu ). On an ask you confirm by typing the environment's name.
In CI, the same plan check runs between plan -out and apply , as a GitHub Action or a Terragrunt before_hook ( docs/ci.md ):
- uses : Flurit/applygate@v0.6.0
with :
plan : tfplan
working-directory : envs/prod
Model judge (optional)
No model is needed. If you want a second opinion on calls the rules do not recognize (Pulumi, database clients, unknown scripts), turn on judge: in your user config with a model you choose. It can only make a verdict stricter, never looser ( docs/judge.md ).
applygate makes no network calls by default, and logs decisions locally with secrets redacted. Three opt-in features do make network calls: the model judge, anonymous usage counts ( telemetry ), and live checks of your AWS identity and remote state.
Install : installer, Homebrew, Go, signature checks, the Claude Code plugin
Configuration and policy rules
Limits , threat model , what has been tested
Upgrading from an earlier release , changelog , roadmap
The easiest way in is a scenario: a command or plan that applygate gets wrong, added to the pack with the verdict it should get ( CONTRIBUTING.md ). Questions go in Discussions . If you find a way around the guard, report it privately as SECURITY.md describes, not in a public issue.
Apache-2.0. See LICENSE . Maintained by Flurit.
A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu.
flurit.github.io/applygate/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devi
[truncated]

GitHub - Flurit/applygate: A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu. · GitHub
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
Flurit
applygate Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu.
flurit.github.io/applygate/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
358 Commits 358 Commits Folders and files
.claude-plugin .claude-plugin .github .github cmd cmd docs docs examples/ judge examples/ judge experiments experiments harness harness internal internal packaging/ homebrew packaging/ homebrew plugin/ claude-code plugin/ claude-code review review scripts scripts site site .dockerignore .dockerignore .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md action.yml action.yml go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
Stop coding agents from running terraform destroy in production.
What it does: before an AI coding agent runs a command, applygate checks it and answers allow , ask or block , with a reason the agent can act on. It covers Terraform, OpenTofu and Terragrunt, and the aws, gcloud, az, kubectl and helm CLIs.
Who it is for: teams that let Claude Code, Codex or another coding agent work in infrastructure repositories.
How: it runs on your machine as the agent's pre-tool-call hook. Fixed rules decide, with no model involved by default, and nothing leaves your machine.
What it is not: a sandbox. It catches mistakes and most injected instructions, but an agent that is set on getting around it and holds credentials that can write can still do damage. What it does not protect against .
applygate is pre-1.0, for macOS and Linux.
brew install Flurit/tap/applygate # other ways to install: docs/install.md
cd your-infra-repo
applygate init # .applygate.yaml for this repository, hooks for Claude Code and Codex, a user config if you have none
applygate check " terraform -chdir=envs/prod destroy "
Using Codex? Open it and run /hooks once to trust the new hook. Until you do, Codex skips it without saying so. applygate doctor tells you whether it will run.
Two config files. .applygate.yaml lists the repository's environments; commit it so your team shares it. It can only add protection, and after every edit a person runs applygate trust , so an agent cannot quietly relabel production. Your own settings, such as observe or enforce, live in ~/.config/applygate/config.yaml .
It starts in observe mode. It logs what it would block and blocks nothing. That includes an agent editing or removing applygate itself, so try it where that is acceptable.
Then turn it on. After a few days, applygate report shows what it would have stopped and applygate suggest proposes environments and tiers. Set mode: enforce in your user config and run applygate doctor .
Every command is judged by the environment it targets. You map directories, workspaces, backends, AWS profiles, Kubernetes contexts, GCP projects and Azure subscriptions to a tier in .applygate.yaml ( config ): protected for production, standard or disposable . Anything applygate cannot place counts as protected.
Plan check: for apply <planfile> it reads the plan and rates what it finds as critical or high: destroyed or replaced stateful resources, removed deletion protection, internet exposure, IAM wildcards, state-backend changes. In a protected environment critical findings block and high ones ask; in a standard one critical findings ask. You can mark your own resources as critical ( policy ).
Turning applygate off is blocked in every tier, but only a managed install makes that a guarantee ( below ).
Outside protected environments, applygate still asks about a command that can run code it cannot follow if the same repository holds a protected environment.
applygate reads commands the way the shell does. It follows cd , variables, bash -c , eval , sudo , xargs , Makefiles, scripts and base64 payloads to find what actually runs, so grep "terraform destroy" is just a grep and a destroy hidden in base64 is still a destroy. In a protected environment it also asks about anything that is not a plain command, because only plain commands are fully understood.
On 17 dangerous Terraform scenarios (our own set, so a demonstration, not a measure of coverage):
The hand-written policy also blocked harmless text like git commit -m "never run terraform destroy" . The scenarios are in harness/ , and applygate selftest runs all 163 scenarios against your binary.
Agent
Status
Can show you an ask prompt
If applygate cannot run
Claude Code
Tested live
yes (except in acceptEdits , plan and dontAsk modes)
the call is blocked (Claude Code 2.1.295 or later)
Codex
Tested live
no: an ask is a block
the call runs
GitHub Copilot CLI
Preview
yes
blocked; runs on a timeout
Cursor
Preview
no: an ask is a block
blocked
Gemini CLI
Preview
no: an ask is a block
the call runs
Windsurf (Cascade)
Preview
no: an ask is a block
the call runs
Devin (shell commands only)
Preview
no: an ask is a block
the call runs
OpenCode
Preview
no: an ask is a block
blocked; runs if the plugin fails to load
Cline
Preview
no: an ask is a block (it cancels the task)
blocked; runs on a timeout or a missing hook script
applygate init installs the hooks for Claude Code and Codex.
Preview hooks follow each agent's documentation but have not been tested inside the agent. Install one with applygate hook install --agent NAME . To check it works, ask the agent to run terraform destroy -auto-approve in an empty directory: it should be blocked.
On agents that cannot ask, a person can allow one stopped call with applygate approve ID . This is off by default, because an agent on the same machine can reach your terminal too ( policy ).
On agents where the call runs if applygate cannot start, keep the agent's own sandbox and approval settings on as a second line.
MCP tools: in Claude Code and Codex, calls to known infrastructure MCP servers (AWS, HCP Terraform, Kubernetes, gcloud, az, shell servers) are judged like the equivalent command. This has not yet been tested against the running servers ( mcp ).
Per-agent details: docs/hosts.md , or applygate hosts .
A check that fails (a bad config, unreadable input, an internal error) blocks, unless every directory in the repository is configured as standard or disposable. on_error in your user config changes this.
A decision that takes longer than 25 seconds is blocked.
If applygate cannot start at all, the agent decides: see the table above. applygate doctor checks that the hook is in place.
To stop it blocking, set mode: observe in your user config. To remove it, run applygate hook remove --agent NAME for each agent, then brew uninstall applygate .
What it does not protect against
A hook cannot promise that a command is never allowed. What protects you when it is wrong is a boundary around the agent: credentials that can only plan, writes behind MFA, a versioned and locked state bucket, a host sandbox and a reviewed apply path.
applygate doctor --boundary checks the ones it can and reports an assurance level from L0 (the hook only) to L3 (all of these in place) ( boundary ).
applygate doctor --creds asks your cloud accounts whether the credentials on this machine can change production without a person ( credentials ).
Stopping an agent from turning applygate off is guaranteed only with a managed, root-owned install ( managed ); without one it is best-effort.
Known gaps are in docs/limits.md , and what has been tested live is in docs/capabilities.md .
From your terminal, applygate apply tfplan judges a saved plan the way the hooks do, then runs terraform apply (or tofu ). On an ask you confirm by typing the environment's name.
In CI, the same plan check runs between plan -out and apply , as a GitHub Action or a Terragrunt before_hook ( docs/ci.md ):
- uses : Flurit/applygate@v0.6.0
with :
plan : tfplan
working-directory : envs/prod
Model judge (optional)
No model is needed. If you want a second opinion on calls the rules do not recognize (Pulumi, database clients, unknown scripts), turn on judge: in your user config with a model you choose. It can only make a verdict stricter, never looser ( docs/judge.md ).
applygate makes no network calls by default, and logs decisions locally with secrets redacted. Three opt-in features do make network calls: the model judge, anonymous usage counts ( telemetry ), and live checks of your AWS identity and remote state.
Install : installer, Homebrew, Go, signature checks, the Claude Code plugin
Configuration and policy rules
Limits , threat model , what has been tested
Upgrading from an earlier release , changelog , roadmap
The easiest way in is a scenario: a command or plan that applygate gets wrong, added to the pack with the verdict it should get ( CONTRIBUTING.md ). Questions go in Discussions . If you find a way around the guard, report it privately as SECURITY.md describes, not in a public issue.
Apache-2.0. See LICENSE . Maintained by Flurit.
A check between AI coding agents and terraform apply. It reads the plan, not just the command, so an apply that would replace a prod database is stopped while safe ones go through. Local allow/ask/block hook for Claude Code and Codex; preview hooks for Cursor, Copilot CLI, Gemini CLI, Windsurf, Devin, OpenCode and Cline. Terraform and OpenTofu.
flurit.github.io/applygate/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
