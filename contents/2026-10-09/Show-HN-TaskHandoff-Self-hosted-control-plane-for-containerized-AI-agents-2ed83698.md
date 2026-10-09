---
source: "https://github.com/edgestorage/task-handoff"
hn_url: "https://news.ycombinator.com/item?id=50023550"
title: "Show HN: TaskHandoff – Self-hosted control plane for containerized AI agents"
article_title: "GitHub - edgestorage/task-handoff: TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines · GitHub"
image: "https://opengraph.githubassets.com/6d763b10745e9adc893d27043751806df4a44810c7e82ea21e7a65d9c7d98a59/edgestorage/task-handoff"
author: "huadream5827"
captured_at: "2026-10-09T18:02:16Z"
capture_tool: "hn-digest"
hn_id: 50023550
score: 1
comments: 0
posted_at: "2026-10-09T17:09:06Z"
tags:
  - hacker-news
---

# Show HN: TaskHandoff – Self-hosted control plane for containerized AI agents

- HN: [50023550](https://news.ycombinator.com/item?id=50023550)
- Source: [github.com](https://github.com/edgestorage/task-handoff)
- Score: 1
- Comments: 0
- Posted: 2026-10-09T17:09:06Z

## Translation

Title: Show HN: TaskHandoff – Self-hosted control plane for containerized AI agents
Article title: GitHub - edgestorage/task-handoff: TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines · GitHub
Description: TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines - edgestorage/task-handoff

Article text:
GitHub - edgestorage/task-handoff: TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines · GitHub
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
edgestorage
task-handoff Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines
Readme Apache-2.0 license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
194 Commits 194 Commits Folders and files
.agents/ skills .agents/ skills .github/ workflows .github/ workflows apps apps bin bin build build builtin-projects/ assistant builtin-projects/ assistant docker docker docs docs packages packages patches patches scripts scripts shared shared test test .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .npmrc .npmrc .tmp-web-cap-inspect.js .tmp-web-cap-inspect.js LICENSE LICENSE NOTICE NOTICE README.md README.md README_CN.md README_CN.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rollup.config.mjs rollup.config.mjs runtime-packages.config.mjs runtime-packages.config.mjs tsconfig.controlled-instance.json tsconfig.controlled-instance.json tsconfig.json tsconfig.json tsconfig.workspace.json tsconfig.workspace.json View all files Repository files navigation
A unified control plane for running, managing, and collaborating with AI workspaces across local and remote machines
TaskHandoff brings Codex and other AI development work into one control plane. It connects AI sessions spread across machines, workspaces, and chat platforms while managing node enrollment, instance lifecycles, sessions, applications, and message routing.
Beta: Task Handoff is under active development. Breaking changes may land between releases.
Dark mode, with a Story session and a browser running in the container:
For more interface screenshots, see the Interface and Workbench guide .
Multi-node management — Connect local and remote nodes and inspect their resources and managed instances from one place.
Managed workspaces — Create, start, stop, and restore isolated workspaces, with Docker as the primary runtime today.
Image market and custom images — Choose from a read-only built-in catalog or separately managed custom images through one instance creation flow.
Environment templates — Save a Docker instance's installed tools and container configuration as a node-local reusable environment, then combine it with any project or local-folder workspace.
AI session center — View and control sessions across instances with real-time state delivered over WebSocket.
Repository workflows — Inspect files, changes, branches, and worktrees, with conservative remote delivery for Git repositories.
Managed Git credentials — Scope HTTPS tokens or pinned SSH keys to remotes, use them for one-time provisioning, or retain them for Agent, Terminal, App, and Repository Git commands.
Chat integrations — Route messages, approvals, and actions from Telegram, DingTalk, WeChat, and Feishu/Lark to a selected instance.
Application management — Install, remove, and run applications on target instances through a trusted built-in catalog.
Mobile client — Connect an iOS or Android device directly to a user-managed Control Plane for AI sessions, instance operations, applications, and terminals.
Desktop and server deployment — Run TaskHandoff as a mobile or desktop application, or as systemd services on Debian and Ubuntu.
English and Chinese UI — Switch languages instantly or follow the browser language automatically.
Browser / Desktop / Mobile / Chat platforms
│
▼
Control Plane
UI, API, and chat gateway
│
▼
Node Agent
Node resources and instance lifecycle
│
▼
Controlled Instance
Workspace, applications, and AI sessions
TaskHandoff is organized into three runtime layers:
Control Plane provides the Web/API management surface and owns the node inventory, instance board, chat gateway, and cross-instance AI session views.
Node Agent runs on each managed machine and owns node-local configuration, runtime resources, folder inventory, and controlled instance lifecycles. Instances continue running when the control plane is stopped or restarted.
Controlled Instance hosts a workspace, applications, AI sessions, triggers, and metadata. It can run standalone; in a managed deployment, its lifecycle and access are owned by the Node Agent.
Chat and AI Session state form a cross-layer path: the Control Plane owns chat credentials, bindings, command parsing, and routing, while each target AI Session remains the source of truth for conversation state.
Docker is the primary isolated runtime and supports multiple instances on one node. A built-in Local Runtime is also available on supported non-Windows nodes for one controlled instance per host user. Runtime capabilities and adapters keep the same model extensible to Kubernetes without creating a separate UI flow.
An environment template is a node-local Docker image created from an existing instance with docker commit . Registry images and environment templates are peer environment sources in the instance creation flow. Workspace selection remains independent, so either source can be combined with a Git project or a local-folder workspace.
Templates capture only the container writable layer, such as installed system packages and tools. They exclude /workspace , /data , /home/agent , every other bind mount or volume, memory, processes, and network state. Derived instances always receive a new identity, registration token, port, and managed volumes. The node agent briefly pauses the source container during commit and rejects a template if Docker Config contains instance-private credentials.
Every Docker instance has managed volumes for /data and /home/agent ; Git workspaces also have a managed /workspace volume, while local folders use an external bind mount. The instance deletion dialog uses one option, selected by default, to delete all managed data. Clearing it retains every managed volume and reports its name; retained volumes are never attached automatically to another instance.
The source node owns both the template record and its Docker image, so a template can be used only on that node while it is ready. Deleting a template removes its internal template tag. A content-addressed internal lease keeps the image recoverable while derived instances reference it, and the image is garbage-collected after the final reference is removed.
Docker, when using Docker Runtime, building container images, or running the standalone Compose profile
pnpm install
pnpm run build:all
pnpm cli help
Start the Control Plane API and development UI in separate terminals. The disabled authentication mode is intended only for loopback development:
pnpm cli control-plane --auth-mode disabled
pnpm run control-plane-ui:dev
Common development commands:
# Start the control-plane UI
pnpm run control-plane-ui:dev
# Type-check and build
pnpm run typecheck
pnpm run web:typecheck
pnpm run build:all
# Run tests
pnpm test
# Inspect the npm package contents
pnpm run pack:dry
To run a standalone Browser-profile controlled instance instead of the Control Plane development stack:
docker compose up -d --build
The current directory is mounted at /workspace by default. Set TASK_HANDOFF_WORKSPACE_HOST to mount a different host directory. This Compose service is a standalone controlled instance, not a Control Plane and Node Agent deployment.
Server deployments install the Control Plane and the server-local Node Agent as independent systemd services. The control plane can stop or restart without terminating instances managed by the agent.
On a Debian or Ubuntu server running systemd, run the latest stable installer as root:
curl -fsSL https://github.com/edgestorage/task-handoff/releases/latest/download/install-server.sh | sudo sh
The script checks the host, installs Node.js 24 and Docker when needed, installs the latest stable @task-handoff/server package from npm, and then creates and starts the Control Plane and Node Agent systemd services. The default auto source profile uses Tsinghua APT mirrors and npmmirror for Chinese locale or timezone environments, and also falls back to those mirrors when the official Node.js source is unreachable. Its temporary APT source list does not overwrite the host's source configuration. By default, the control plane listens on port 8081 with password authentication enabled. Installer options can change the port, authentication mode, release channel, and other service settings.
Use --install-source china or --install-source official to select a source profile explicitly:
curl -fsSL https://github.com/edgestorage/task-handoff/releases/latest/download/install-server.sh | sudo sh -s -- --install-source china
Install with Node.js and Docker already available
sudo npm install -g @task-handoff/server@latest
sudo task-handoff install
Manage services and updates:
sudo task-handoff start
sudo task-handoff stop
sudo task-handoff restart
task-handoff check
sudo task-handoff update
The installation creates:
task-handoff-node-agent.service
task-handoff-control-plane.service
See scripts/install-server.sh for supported installer options.
Remote machines only need the Node Agent. Generate a one-time join token in the Control Plane and prefer the exact installation command shown there. Its package version is resolved from the running Control Plane release. The equivalent form is:
curl -fsSL https://CONTROL_PLANE_HOST/install-node-agent.sh | sudo sh -s -- \
--control-plane https://CONTROL_PLANE_HOST \
--join-token JOIN_TOKEN \
--npm-package @task-handoff/node-agent \
--controlled-instance-package @task-handoff/controlled-instance \
--version RELEASE_VERSION
On Debian and Ubuntu, the remote-node installer bootstraps the required Node.js
24 and npm on a fresh host. It uses the same automatic source selection; append
--install-source china to force Chinese mirrors. The selected npm registry is
preserved in the Node Agent service environment for subsequent managed updates.
Replace RELEASE_VERSION with the Control Plane's runtime package version so the Node Agent and controlled-instance runtime use the same release.
An installed Node Agent can also generate a one-time invitation directly on the node:
sudo task-handoff-node-agent invite --ipc-path /run/task-handoff/node-agent.sock
Add --json for automation-friendly output. Remote TCP access still requires an invitation and paired HMAC authentication.
To remove a standalone Node Agent installation:
sudo task

[truncated]

## Original Extract

TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines - edgestorage/task-handoff

GitHub - edgestorage/task-handoff: TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines · GitHub
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
edgestorage
task-handoff Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
TaskHandoff is a unified control plane for running, managing, and collaborating with docker-based Codex workspaces across local and remote machines
Readme Apache-2.0 license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
194 Commits 194 Commits Folders and files
.agents/ skills .agents/ skills .github/ workflows .github/ workflows apps apps bin bin build build builtin-projects/ assistant builtin-projects/ assistant docker docker docs docs packages packages patches patches scripts scripts shared shared test test .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .npmrc .npmrc .tmp-web-cap-inspect.js .tmp-web-cap-inspect.js LICENSE LICENSE NOTICE NOTICE README.md README.md README_CN.md README_CN.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rollup.config.mjs rollup.config.mjs runtime-packages.config.mjs runtime-packages.config.mjs tsconfig.controlled-instance.json tsconfig.controlled-instance.json tsconfig.json tsconfig.json tsconfig.workspace.json tsconfig.workspace.json View all files Repository files navigation
A unified control plane for running, managing, and collaborating with AI workspaces across local and remote machines
TaskHandoff brings Codex and other AI development work into one control plane. It connects AI sessions spread across machines, workspaces, and chat platforms while managing node enrollment, instance lifecycles, sessions, applications, and message routing.
Beta: Task Handoff is under active development. Breaking changes may land between releases.
Dark mode, with a Story session and a browser running in the container:
For more interface screenshots, see the Interface and Workbench guide .
Multi-node management — Connect local and remote nodes and inspect their resources and managed instances from one place.
Managed workspaces — Create, start, stop, and restore isolated workspaces, with Docker as the primary runtime today.
Image market and custom images — Choose from a read-only built-in catalog or separately managed custom images through one instance creation flow.
Environment templates — Save a Docker instance's installed tools and container configuration as a node-local reusable environment, then combine it with any project or local-folder workspace.
AI session center — View and control sessions across instances with real-time state delivered over WebSocket.
Repository workflows — Inspect files, changes, branches, and worktrees, with conservative remote delivery for Git repositories.
Managed Git credentials — Scope HTTPS tokens or pinned SSH keys to remotes, use them for one-time provisioning, or retain them for Agent, Terminal, App, and Repository Git commands.
Chat integrations — Route messages, approvals, and actions from Telegram, DingTalk, WeChat, and Feishu/Lark to a selected instance.
Application management — Install, remove, and run applications on target instances through a trusted built-in catalog.
Mobile client — Connect an iOS or Android device directly to a user-managed Control Plane for AI sessions, instance operations, applications, and terminals.
Desktop and server deployment — Run TaskHandoff as a mobile or desktop application, or as systemd services on Debian and Ubuntu.
English and Chinese UI — Switch languages instantly or follow the browser language automatically.
Browser / Desktop / Mobile / Chat platforms
│
▼
Control Plane
UI, API, and chat gateway
│
▼
Node Agent
Node resources and instance lifecycle
│
▼
Controlled Instance
Workspace, applications, and AI sessions
TaskHandoff is organized into three runtime layers:
Control Plane provides the Web/API management surface and owns the node inventory, instance board, chat gateway, and cross-instance AI session views.
Node Agent runs on each managed machine and owns node-local configuration, runtime resources, folder inventory, and controlled instance lifecycles. Instances continue running when the control plane is stopped or restarted.
Controlled Instance hosts a workspace, applications, AI sessions, triggers, and metadata. It can run standalone; in a managed deployment, its lifecycle and access are owned by the Node Agent.
Chat and AI Session state form a cross-layer path: the Control Plane owns chat credentials, bindings, command parsing, and routing, while each target AI Session remains the source of truth for conversation state.
Docker is the primary isolated runtime and supports multiple instances on one node. A built-in Local Runtime is also available on supported non-Windows nodes for one controlled instance per host user. Runtime capabilities and adapters keep the same model extensible to Kubernetes without creating a separate UI flow.
An environment template is a node-local Docker image created from an existing instance with docker commit . Registry images and environment templates are peer environment sources in the instance creation flow. Workspace selection remains independent, so either source can be combined with a Git project or a local-folder workspace.
Templates capture only the container writable layer, such as installed system packages and tools. They exclude /workspace , /data , /home/agent , every other bind mount or volume, memory, processes, and network state. Derived instances always receive a new identity, registration token, port, and managed volumes. The node agent briefly pauses the source container during commit and rejects a template if Docker Config contains instance-private credentials.
Every Docker instance has managed volumes for /data and /home/agent ; Git workspaces also have a managed /workspace volume, while local folders use an external bind mount. The instance deletion dialog uses one option, selected by default, to delete all managed data. Clearing it retains every managed volume and reports its name; retained volumes are never attached automatically to another instance.
The source node owns both the template record and its Docker image, so a template can be used only on that node while it is ready. Deleting a template removes its internal template tag. A content-addressed internal lease keeps the image recoverable while derived instances reference it, and the image is garbage-collected after the final reference is removed.
Docker, when using Docker Runtime, building container images, or running the standalone Compose profile
pnpm install
pnpm run build:all
pnpm cli help
Start the Control Plane API and development UI in separate terminals. The disabled authentication mode is intended only for loopback development:
pnpm cli control-plane --auth-mode disabled
pnpm run control-plane-ui:dev
Common development commands:
# Start the control-plane UI
pnpm run control-plane-ui:dev
# Type-check and build
pnpm run typecheck
pnpm run web:typecheck
pnpm run build:all
# Run tests
pnpm test
# Inspect the npm package contents
pnpm run pack:dry
To run a standalone Browser-profile controlled instance instead of the Control Plane development stack:
docker compose up -d --build
The current directory is mounted at /workspace by default. Set TASK_HANDOFF_WORKSPACE_HOST to mount a different host directory. This Compose service is a standalone controlled instance, not a Control Plane and Node Agent deployment.
Server deployments install the Control Plane and the server-local Node Agent as independent systemd services. The control plane can stop or restart without terminating instances managed by the agent.
On a Debian or Ubuntu server running systemd, run the latest stable installer as root:
curl -fsSL https://github.com/edgestorage/task-handoff/releases/latest/download/install-server.sh | sudo sh
The script checks the host, installs Node.js 24 and Docker when needed, installs the latest stable @task-handoff/server package from npm, and then creates and starts the Control Plane and Node Agent systemd services. The default auto source profile uses Tsinghua APT mirrors and npmmirror for Chinese locale or timezone environments, and also falls back to those mirrors when the official Node.js source is unreachable. Its temporary APT source list does not overwrite the host's source configuration. By default, the control plane listens on port 8081 with password authentication enabled. Installer options can change the port, authentication mode, release channel, and other service settings.
Use --install-source china or --install-source official to select a source profile explicitly:
curl -fsSL https://github.com/edgestorage/task-handoff/releases/latest/download/install-server.sh | sudo sh -s -- --install-source china
Install with Node.js and Docker already available
sudo npm install -g @task-handoff/server@latest
sudo task-handoff install
Manage services and updates:
sudo task-handoff start
sudo task-handoff stop
sudo task-handoff restart
task-handoff check
sudo task-handoff update
The installation creates:
task-handoff-node-agent.service
task-handoff-control-plane.service
See scripts/install-server.sh for supported installer options.
Remote machines only need the Node Agent. Generate a one-time join token in the Control Plane and prefer the exact installation command shown there. Its package version is resolved from the running Control Plane release. The equivalent form is:
curl -fsSL https://CONTROL_PLANE_HOST/install-node-agent.sh | sudo sh -s -- \
--control-plane https://CONTROL_PLANE_HOST \
--join-token JOIN_TOKEN \
--npm-package @task-handoff/node-agent \
--controlled-instance-package @task-handoff/controlled-instance \
--version RELEASE_VERSION
On Debian and Ubuntu, the remote-node installer bootstraps the required Node.js
24 and npm on a fresh host. It uses the same automatic source selection; append
--install-source china to force Chinese mirrors. The selected npm registry is
preserved in the Node Agent service environment for subsequent managed updates.
Replace RELEASE_VERSION with the Control Plane's runtime package version so the Node Agent and controlled-instance runtime use the same release.
An installed Node Agent can also generate a one-time invitation directly on the node:
sudo task-handoff-node-agent invite --ipc-path /run/task-handoff/node-agent.sock
Add --json for automation-friendly output. Remote TCP access still requires an invitation and paired HMAC authentication.
To remove a standalone Node Agent installation:
sudo task

[truncated]
