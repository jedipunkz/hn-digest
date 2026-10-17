---
source: "https://github.com/trypando/pando"
hn_url: "https://news.ycombinator.com/item?id=50131518"
title: "Ask HN: How are you all ensuring your AI code isn't slop?"
article_title: "GitHub - trypando/pando: Deploy and share apps · GitHub"
image: "https://opengraph.githubassets.com/4d9927296d0b1a6fd4dc77e1a7e3ec975805bfc827b3023cdfdfa37e36f3f151/trypando/pando"
author: "bemeeker"
captured_at: "2026-10-17T15:04:44Z"
capture_tool: "hn-digest"
hn_id: 50131518
score: 1
comments: 1
posted_at: "2026-10-17T14:29:57Z"
tags:
  - hacker-news
---

# Ask HN: How are you all ensuring your AI code isn't slop?

- HN: [50131518](https://news.ycombinator.com/item?id=50131518)
- Source: [github.com](https://github.com/trypando/pando)
- Score: 1
- Comments: 1
- Posted: 2026-10-17T14:29:57Z

## Translation

Title: Ask HN: How are you all ensuring your AI code isn't slop?
Article title: GitHub - trypando/pando: Deploy and share apps · GitHub
Description: Deploy and share apps. Contribute to trypando/pando development by creating an account on GitHub.

Article text:
GitHub - trypando/pando: Deploy and share apps · GitHub
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
trypando
pando Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
828 Commits 828 Commits Folders and files
.claude/ skills .claude/ skills .github .github cmd cmd console console deploy/ kubernetes deploy/ kubernetes docs docs internal internal migrations migrations scripts scripts test test tools tools .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR_AGREEMENT.md CONTRIBUTOR_AGREEMENT.md Dockerfile Dockerfile LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md codecov.yml codecov.yml docker-compose.registry.yml docker-compose.registry.yml docker-compose.yml docker-compose.yml entrypoint.sh entrypoint.sh go.mod go.mod go.sum go.sum osv-scanner.toml osv-scanner.toml sonar-project.properties sonar-project.properties sqlc.yaml sqlc.yaml View all files Repository files navigation
trypando.ai · Quickstart · Docs
Pando is an open-source, self-hosted platform for running the apps your team builds, including ones
written with AI. Point it at a git repo or image and it builds the app, scans it for
vulnerabilities, runs it on its own isolated network, and puts it behind your company's SSO, shared
only with the people who need it.
Built for IT and platform teams who want employees to deploy internal tools without those tools
living on laptops or personal cloud accounts. Also works on a home server.
Status: alpha. Expect breaking changes; see the CHANGELOG before upgrading.
Needs Docker and Docker Compose. Quickstart →
Install Pando with the quickstart .
docs/identity-providers.md : connecting Okta, Microsoft Entra ID,
Google Workspace, Keycloak, Authentik or another provider, and SCIM.
docs/rootless.md : running Pando on rootless Docker or Podman on Linux.
docs/reference.md : running on several hosts or on Kubernetes, host policy at
startup, and what an app receives from Pando.
CHANGELOG.md : what changed in each release, and whether you need to act.
A bug , or something that doesn't work as documented: open an
issue .
A security vulnerability : privately, through the advisory
form , not as an issue. See
SECURITY.md for what to include and what response to expect.
Bug reports and pull requests are welcome. See CONTRIBUTING.md for how the
project is organized, how to build it, and what a change needs before it can be merged.
CODE_OF_CONDUCT.md applies everywhere the project happens.
AGPL-3.0, with a commercial license available for embedding Pando in a proprietary product or
offering it as a hosted service without publishing modifications. See
LICENSING.md for which applies to you, and LICENSE for the full text.
Contributors sign the Contributor License Agreement once, by posting a
comment on their first pull request.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Deploy and share apps. Contribute to trypando/pando development by creating an account on GitHub.

GitHub - trypando/pando: Deploy and share apps · GitHub
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
trypando
pando Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
828 Commits 828 Commits Folders and files
.claude/ skills .claude/ skills .github .github cmd cmd console console deploy/ kubernetes deploy/ kubernetes docs docs internal internal migrations migrations scripts scripts test test tools tools .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR_AGREEMENT.md CONTRIBUTOR_AGREEMENT.md Dockerfile Dockerfile LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile README.md README.md SECURITY.md SECURITY.md codecov.yml codecov.yml docker-compose.registry.yml docker-compose.registry.yml docker-compose.yml docker-compose.yml entrypoint.sh entrypoint.sh go.mod go.mod go.sum go.sum osv-scanner.toml osv-scanner.toml sonar-project.properties sonar-project.properties sqlc.yaml sqlc.yaml View all files Repository files navigation
trypando.ai · Quickstart · Docs
Pando is an open-source, self-hosted platform for running the apps your team builds, including ones
written with AI. Point it at a git repo or image and it builds the app, scans it for
vulnerabilities, runs it on its own isolated network, and puts it behind your company's SSO, shared
only with the people who need it.
Built for IT and platform teams who want employees to deploy internal tools without those tools
living on laptops or personal cloud accounts. Also works on a home server.
Status: alpha. Expect breaking changes; see the CHANGELOG before upgrading.
Needs Docker and Docker Compose. Quickstart →
Install Pando with the quickstart .
docs/identity-providers.md : connecting Okta, Microsoft Entra ID,
Google Workspace, Keycloak, Authentik or another provider, and SCIM.
docs/rootless.md : running Pando on rootless Docker or Podman on Linux.
docs/reference.md : running on several hosts or on Kubernetes, host policy at
startup, and what an app receives from Pando.
CHANGELOG.md : what changed in each release, and whether you need to act.
A bug , or something that doesn't work as documented: open an
issue .
A security vulnerability : privately, through the advisory
form , not as an issue. See
SECURITY.md for what to include and what response to expect.
Bug reports and pull requests are welcome. See CONTRIBUTING.md for how the
project is organized, how to build it, and what a change needs before it can be merged.
CODE_OF_CONDUCT.md applies everywhere the project happens.
AGPL-3.0, with a commercial license available for embedding Pando in a proprietary product or
offering it as a hosted service without publishing modifications. See
LICENSING.md for which applies to you, and LICENSE for the full text.
Contributors sign the Contributor License Agreement once, by posting a
comment on their first pull request.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
