---
source: "https://github.com/defenseunicorns/compose-bridge-uds"
hn_url: "https://news.ycombinator.com/item?id=49664258"
title: "Show HN: Package your compose app for airgap K8s"
article_title: "GitHub - defenseunicorns/compose-bridge-uds: a Compose Bridge template for transforming your Compose file to a UDS Package · GitHub"
image: "https://opengraph.githubassets.com/c11ca997b9ce95de536ad6f7840f5a8881a25f523d7de3dba15da57f89a0e1a9/defenseunicorns/compose-bridge-uds"
author: "willswire"
captured_at: "2026-09-11T20:33:53Z"
capture_tool: "hn-digest"
hn_id: 49664258
score: 1
comments: 0
posted_at: "2026-09-11T19:41:48Z"
tags:
  - hacker-news
---

# Show HN: Package your compose app for airgap K8s

- HN: [49664258](https://news.ycombinator.com/item?id=49664258)
- Source: [github.com](https://github.com/defenseunicorns/compose-bridge-uds)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T19:41:48Z

## Translation

Title: Show HN: Package your compose app for airgap K8s
Article title: GitHub - defenseunicorns/compose-bridge-uds: a Compose Bridge template for transforming your Compose file to a UDS Package · GitHub
Description: a Compose Bridge template for transforming your Compose file to a UDS Package - defenseunicorns/compose-bridge-uds

Article text:
GitHub - defenseunicorns/compose-bridge-uds: a Compose Bridge template for transforming your Compose file to a UDS Package · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
defenseunicorns
/
compose-bridge-uds
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
77 Commits 77 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples internal internal scripts scripts tasks tasks tests tests .dockerignore .dockerignore .gitignore .gitignore .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go release-please-config.json release-please-config.json renovate.json renovate.json tasks.yaml tasks.yaml View all files Repository files navigation
Convert a Docker Compose application into a deployable UDS package using Docker Compose Bridge . The transformation consumes a fully-resolved Compose model and emits a Helm chart tailored for UDS, ready for zarf package create and zarf package deploy .
This is not a supported product pathway. It's an experimental transformation we're sharing to gather feedback. Please enagage with our team on the project discussions page with any questions or suggestions.
Docker with Docker Compose (v5.5.0 or later is recommended for build-only services)
This walkthrough deploys UDS Core Slim Dev on k3d, then packages and deploys WordPress and MySQL from a Compose file.
# 1. Create a local k3d cluster with UDS Core Slim Dev
uds zarf package deploy oci://ghcr.io/defenseunicorns/dev/uds/checkpoints/k3d-core-slim-dev:1.11.1
# 2. Transform the Compose application into a UDS Package
cd examples/simple
docker compose bridge convert -t ghcr.io/defenseunicorns/compose-bridge-uds
# 3. Build and deploy the package
zarf package create out/ --flavor upstream
zarf package deploy zarf-package-wordpress- * .tar.zst
The transformation writes these artifacts to out/ :
Is Compose Bridge a good fit for your application? Use the Awesome Compose compatibility matrix as a practical benchmark. If your application depends on unsupported Compose features or needs deeper package customization, start with the UDS reference package and build the package directly.
Compose support describes supported configuration, known limitations, development-service exclusions, and local Dockerfile builds.
UDS package generation describes generated resources, x-uds extensions, secrets, and policy exemptions.
examples/full/compose.yaml demonstrates the complete supported configuration.
The end-to-end smoke test builds this repository as a Compose Bridge transformation, converts examples/simple/compose.yaml , creates and deploys the generated Zarf package on UDS Core Slim in k3d, and runs a Playwright journey through UDS SSO and WordPress installation.
With Docker Compose, k3d, and the UDS CLI installed, run:
uds run test-install
The task uses the standard uds-common package creation, deployment, Core Slim setup, and Keycloak test-user tasks before running the Playwright journey.
a Compose Bridge template for transforming your Compose file to a UDS Package
Readme Apache-2.0 license Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

a Compose Bridge template for transforming your Compose file to a UDS Package - defenseunicorns/compose-bridge-uds

GitHub - defenseunicorns/compose-bridge-uds: a Compose Bridge template for transforming your Compose file to a UDS Package · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
defenseunicorns
/
compose-bridge-uds
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
77 Commits 77 Commits Folders and files
.github/ workflows .github/ workflows docs docs examples examples internal internal scripts scripts tasks tasks tests tests .dockerignore .dockerignore .gitignore .gitignore .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum main.go main.go main_test.go main_test.go release-please-config.json release-please-config.json renovate.json renovate.json tasks.yaml tasks.yaml View all files Repository files navigation
Convert a Docker Compose application into a deployable UDS package using Docker Compose Bridge . The transformation consumes a fully-resolved Compose model and emits a Helm chart tailored for UDS, ready for zarf package create and zarf package deploy .
This is not a supported product pathway. It's an experimental transformation we're sharing to gather feedback. Please enagage with our team on the project discussions page with any questions or suggestions.
Docker with Docker Compose (v5.5.0 or later is recommended for build-only services)
This walkthrough deploys UDS Core Slim Dev on k3d, then packages and deploys WordPress and MySQL from a Compose file.
# 1. Create a local k3d cluster with UDS Core Slim Dev
uds zarf package deploy oci://ghcr.io/defenseunicorns/dev/uds/checkpoints/k3d-core-slim-dev:1.11.1
# 2. Transform the Compose application into a UDS Package
cd examples/simple
docker compose bridge convert -t ghcr.io/defenseunicorns/compose-bridge-uds
# 3. Build and deploy the package
zarf package create out/ --flavor upstream
zarf package deploy zarf-package-wordpress- * .tar.zst
The transformation writes these artifacts to out/ :
Is Compose Bridge a good fit for your application? Use the Awesome Compose compatibility matrix as a practical benchmark. If your application depends on unsupported Compose features or needs deeper package customization, start with the UDS reference package and build the package directly.
Compose support describes supported configuration, known limitations, development-service exclusions, and local Dockerfile builds.
UDS package generation describes generated resources, x-uds extensions, secrets, and policy exemptions.
examples/full/compose.yaml demonstrates the complete supported configuration.
The end-to-end smoke test builds this repository as a Compose Bridge transformation, converts examples/simple/compose.yaml , creates and deploys the generated Zarf package on UDS Core Slim in k3d, and runs a Playwright journey through UDS SSO and WordPress installation.
With Docker Compose, k3d, and the UDS CLI installed, run:
uds run test-install
The task uses the standard uds-common package creation, deployment, Core Slim setup, and Keycloak test-user tasks before running the Playwright journey.
a Compose Bridge template for transforming your Compose file to a UDS Package
Readme Apache-2.0 license Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
