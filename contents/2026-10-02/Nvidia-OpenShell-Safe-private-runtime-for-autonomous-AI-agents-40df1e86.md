---
source: "https://github.com/NVIDIA/OpenShell"
hn_url: "https://news.ycombinator.com/item?id=49933338"
title: "Nvidia/OpenShell: Safe, private runtime for autonomous AI agents"
article_title: "GitHub - NVIDIA/OpenShell: OpenShell is the safe, private runtime for autonomous AI agents. · GitHub"
image: "https://opengraph.githubassets.com/b9fe8d35794f90406abb1a136fd600c723cc6583fe89d4a2fcd90b4d95015075/NVIDIA/OpenShell"
author: "gmays"
captured_at: "2026-10-02T13:37:51Z"
capture_tool: "hn-digest"
hn_id: 49933338
score: 1
comments: 0
posted_at: "2026-10-02T13:31:51Z"
tags:
  - hacker-news
---

# Nvidia/OpenShell: Safe, private runtime for autonomous AI agents

- HN: [49933338](https://news.ycombinator.com/item?id=49933338)
- Source: [github.com](https://github.com/NVIDIA/OpenShell)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T13:31:51Z

## Translation

Title: Nvidia/OpenShell: Safe, private runtime for autonomous AI agents
Article title: GitHub - NVIDIA/OpenShell: OpenShell is the safe, private runtime for autonomous AI agents. · GitHub
Description: OpenShell is the safe, private runtime for autonomous AI agents. - NVIDIA/OpenShell

Article text:
GitHub - NVIDIA/OpenShell: OpenShell is the safe, private runtime for autonomous AI agents. · GitHub
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
NVIDIA
/
OpenShell
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,606 Commits 1,606 Commits Folders and files
.agents/ skills .agents/ skills .cargo .cargo .claude .claude .config .config .github .github .opencode/ agents .opencode/ agents crates crates deploy deploy docs docs e2e e2e examples examples fern fern nix nix proto proto providers providers python python rfc rfc scripts scripts sdk sdk skills skills snap snap tasks tasks telemetry telemetry tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .markdownlint-cli2.jsonc .markdownlint-cli2.jsonc .packit.yaml .packit.yaml .python-version .python-version .trivyignore.yaml .trivyignore.yaml AGENTS.md AGENTS.md CI.md CI.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DCO DCO GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md SECURITY.md SECURITY.md STYLEGUIDE.md STYLEGUIDE.md TESTING.md TESTING.md THIRD-PARTY-NOTICES THIRD-PARTY-NOTICES about.toml about.toml buf.yaml buf.yaml deny.toml deny.toml flake.lock flake.lock flake.nix flake.nix install.sh install.sh mise.lock mise.lock mise.toml mise.toml openshell.spec openshell.spec pyproject.toml pyproject.toml rust-toolchain.toml rust-toolchain.toml snapcraft.yaml snapcraft.yaml uv.lock uv.lock View all files Repository files navigation
New in OpenShell 0.1.x: a stable release cadence, new isolation primitives, an expanded extension surface, and new APIs. Read the 0.1.0 upgrade guide .
OpenShell is the safe, private runtime for fleets of autonomous AI agents. Agents are most useful when they can read files, install packages, call APIs, and use credentials. OpenShell gives them that capability without giving them unrestricted access to your data, secrets, or network. You declare what each agent can touch in a policy, and OpenShell enforces it.
OpenShell governs what agents can do in two ways: it instruments the kernel to enforce policy on every file access, system call, and network connection at runtime, and it uses formal verification to check what a policy change would allow before it is applied.
Kernel-level enforcement. Each agent runs in an isolated sandbox. Kernel controls confine which files it can access and which system calls it can make, and every network connection passes through a policy check before it leaves the sandbox. Agents never see real credentials; OpenShell adds them only to requests bound for approved endpoints.
Formally verified policy changes. Before a policy change is approved, OpenShell uses formal verification to flag risky new access it would grant, such as reaching a new host with credentials or calling a new API method, so those changes wait for human review.
See Architecture for how the gateway, supervisor, and sandbox fit together.
You need Linux, macOS on Apple Silicon, or Windows with WSL 2 (experimental), plus Docker, Podman, or host virtualization. See the Support Matrix for details.
curl -LsSf https://raw.githubusercontent.com/NVIDIA/OpenShell/main/install.sh | sh
openshell sandbox create --name demo
The installer sets up the CLI and a local gateway. The default sandbox image is minimal Ubuntu with no agent installed. To run a real agent, follow Run Your First Agent : it runs OpenCode against a free OpenRouter model and shows how to approve new access as the agent needs it.
Sandboxes : images, runtimes, GPUs, and lifecycle.
Policies : filesystem, network, and process rules, with the advisor and prover for reviewing changes.
Providers : credentials that work only at approved endpoints, including inference .
Gateways : the control plane for sandboxes, policy, and access.
Kubernetes : deploy the gateway with Helm. Your CNI must enforce NetworkPolicy .
Extensibility : middleware, interceptors, and compute drivers.
Tutorials : step-by-step policy and agent walkthroughs.
Prerelease and development builds : try an upcoming release or the latest commit on main .
Install the public OpenShell skills for your coding agent:
npx skills add NVIDIA/OpenShell
The skills teach your agent to drive the OpenShell CLI, write sandbox policies, and debug gateways and inference routing. They live in skills/ and work without an OpenShell source checkout.
SDKs connect applications to an OpenShell gateway. They do not install the CLI. Use the same OpenShell release for the SDK and the gateway when possible.
Questions and discussion: GitHub Discussions
Bug reports and feature requests: GitHub Issues , using the issue templates
Security vulnerabilities: follow SECURITY.md . Do not open a GitHub issue.
Roadmap: OpenShell Roadmap and the RFC board
Try it in the cloud: Brev Launchable
OpenShell is built agent-first: it is developed with the same agent-driven workflows it enables. See CONTRIBUTING.md for development setup and the contribution workflow, and AGENTS.md for repository coding conventions.
OpenShell collects anonymous telemetry, limited to operational categories and counts, to help improve the project. It does not collect sandbox names, hostnames, file paths, prompts, credentials, provider or model names, or user content. To disable it, set OPENSHELL_TELEMETRY_ENABLED=false on the gateway, or server.telemetryEnabled=false for Helm installs. You can also compile telemetry out entirely. See Telemetry for details and the community telemetry reports for published usage trends.
This software automatically retrieves, accesses or interacts with external materials. Those retrieved materials are not distributed with this software and are governed solely by separate terms, conditions and licenses. You are solely responsible for finding, reviewing and complying with all applicable terms, conditions, and licenses, and for verifying the security, integrity and suitability of any retrieved materials for your specific use case. This software is provided "AS IS", without warranty of any kind. The author makes no representations or warranties regarding any retrieved materials, and assumes no liability for any losses, damages, liabilities or legal consequences from your use or inability to use this software or any retrieved materials. Use this software and the retrieved materials at your own risk.
This project is licensed under the Apache License 2.0 .
OpenShell is the safe, private runtime for autonomous AI agents.
docs.nvidia.com/openshell/latest/ Resources
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
1.6k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

OpenShell is the safe, private runtime for autonomous AI agents. - NVIDIA/OpenShell

GitHub - NVIDIA/OpenShell: OpenShell is the safe, private runtime for autonomous AI agents. · GitHub
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
NVIDIA
/
OpenShell
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,606 Commits 1,606 Commits Folders and files
.agents/ skills .agents/ skills .cargo .cargo .claude .claude .config .config .github .github .opencode/ agents .opencode/ agents crates crates deploy deploy docs docs e2e e2e examples examples fern fern nix nix proto proto providers providers python python rfc rfc scripts scripts sdk sdk skills skills snap snap tasks tasks telemetry telemetry tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .markdownlint-cli2.jsonc .markdownlint-cli2.jsonc .packit.yaml .packit.yaml .python-version .python-version .trivyignore.yaml .trivyignore.yaml AGENTS.md AGENTS.md CI.md CI.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DCO DCO GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md SECURITY.md SECURITY.md STYLEGUIDE.md STYLEGUIDE.md TESTING.md TESTING.md THIRD-PARTY-NOTICES THIRD-PARTY-NOTICES about.toml about.toml buf.yaml buf.yaml deny.toml deny.toml flake.lock flake.lock flake.nix flake.nix install.sh install.sh mise.lock mise.lock mise.toml mise.toml openshell.spec openshell.spec pyproject.toml pyproject.toml rust-toolchain.toml rust-toolchain.toml snapcraft.yaml snapcraft.yaml uv.lock uv.lock View all files Repository files navigation
New in OpenShell 0.1.x: a stable release cadence, new isolation primitives, an expanded extension surface, and new APIs. Read the 0.1.0 upgrade guide .
OpenShell is the safe, private runtime for fleets of autonomous AI agents. Agents are most useful when they can read files, install packages, call APIs, and use credentials. OpenShell gives them that capability without giving them unrestricted access to your data, secrets, or network. You declare what each agent can touch in a policy, and OpenShell enforces it.
OpenShell governs what agents can do in two ways: it instruments the kernel to enforce policy on every file access, system call, and network connection at runtime, and it uses formal verification to check what a policy change would allow before it is applied.
Kernel-level enforcement. Each agent runs in an isolated sandbox. Kernel controls confine which files it can access and which system calls it can make, and every network connection passes through a policy check before it leaves the sandbox. Agents never see real credentials; OpenShell adds them only to requests bound for approved endpoints.
Formally verified policy changes. Before a policy change is approved, OpenShell uses formal verification to flag risky new access it would grant, such as reaching a new host with credentials or calling a new API method, so those changes wait for human review.
See Architecture for how the gateway, supervisor, and sandbox fit together.
You need Linux, macOS on Apple Silicon, or Windows with WSL 2 (experimental), plus Docker, Podman, or host virtualization. See the Support Matrix for details.
curl -LsSf https://raw.githubusercontent.com/NVIDIA/OpenShell/main/install.sh | sh
openshell sandbox create --name demo
The installer sets up the CLI and a local gateway. The default sandbox image is minimal Ubuntu with no agent installed. To run a real agent, follow Run Your First Agent : it runs OpenCode against a free OpenRouter model and shows how to approve new access as the agent needs it.
Sandboxes : images, runtimes, GPUs, and lifecycle.
Policies : filesystem, network, and process rules, with the advisor and prover for reviewing changes.
Providers : credentials that work only at approved endpoints, including inference .
Gateways : the control plane for sandboxes, policy, and access.
Kubernetes : deploy the gateway with Helm. Your CNI must enforce NetworkPolicy .
Extensibility : middleware, interceptors, and compute drivers.
Tutorials : step-by-step policy and agent walkthroughs.
Prerelease and development builds : try an upcoming release or the latest commit on main .
Install the public OpenShell skills for your coding agent:
npx skills add NVIDIA/OpenShell
The skills teach your agent to drive the OpenShell CLI, write sandbox policies, and debug gateways and inference routing. They live in skills/ and work without an OpenShell source checkout.
SDKs connect applications to an OpenShell gateway. They do not install the CLI. Use the same OpenShell release for the SDK and the gateway when possible.
Questions and discussion: GitHub Discussions
Bug reports and feature requests: GitHub Issues , using the issue templates
Security vulnerabilities: follow SECURITY.md . Do not open a GitHub issue.
Roadmap: OpenShell Roadmap and the RFC board
Try it in the cloud: Brev Launchable
OpenShell is built agent-first: it is developed with the same agent-driven workflows it enables. See CONTRIBUTING.md for development setup and the contribution workflow, and AGENTS.md for repository coding conventions.
OpenShell collects anonymous telemetry, limited to operational categories and counts, to help improve the project. It does not collect sandbox names, hostnames, file paths, prompts, credentials, provider or model names, or user content. To disable it, set OPENSHELL_TELEMETRY_ENABLED=false on the gateway, or server.telemetryEnabled=false for Helm installs. You can also compile telemetry out entirely. See Telemetry for details and the community telemetry reports for published usage trends.
This software automatically retrieves, accesses or interacts with external materials. Those retrieved materials are not distributed with this software and are governed solely by separate terms, conditions and licenses. You are solely responsible for finding, reviewing and complying with all applicable terms, conditions, and licenses, and for verifying the security, integrity and suitability of any retrieved materials for your specific use case. This software is provided "AS IS", without warranty of any kind. The author makes no representations or warranties regarding any retrieved materials, and assumes no liability for any losses, damages, liabilities or legal consequences from your use or inability to use this software or any retrieved materials. Use this software and the retrieved materials at your own risk.
This project is licensed under the Apache License 2.0 .
OpenShell is the safe, private runtime for autonomous AI agents.
docs.nvidia.com/openshell/latest/ Resources
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
1.6k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
