---
source: "https://github.com/ggmm82/righaven"
hn_url: "https://news.ycombinator.com/item?id=50407370"
title: "RigHaven – Self-hosted execution harness for AI agents"
article_title: "GitHub - ggmm82/righaven · GitHub"
image: "https://opengraph.githubassets.com/926aa00dc9fadaff50f89cdaca4bc72ec0b691cce24559ffb248da9b7f98f5eb/ggmm82/righaven"
author: "devusergm"
captured_at: "2026-11-07T19:55:25Z"
capture_tool: "hn-digest"
hn_id: 50407370
score: 1
comments: 0
posted_at: "2026-11-07T19:40:06Z"
tags:
  - hacker-news
---

# RigHaven – Self-hosted execution harness for AI agents

- HN: [50407370](https://news.ycombinator.com/item?id=50407370)
- Source: [github.com](https://github.com/ggmm82/righaven)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T19:40:06Z

## Translation

Title: RigHaven – Self-hosted execution harness for AI agents
Article title: GitHub - ggmm82/righaven · GitHub
Description: Contribute to ggmm82/righaven development by creating an account on GitHub.

Article text:
GitHub - ggmm82/righaven · GitHub
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
ggmm82
righaven Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme License Code of conduct
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
72 Commits 72 Commits Folders and files
.github/ workflows .github/ workflows config config core core desktop desktop docker docker docs docs frontend frontend locales locales righaven righaven scripts scripts skills/ custom skills/ custom tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md VERSION VERSION docker-compose.yaml docker-compose.yaml requirements-dev.txt requirements-dev.txt requirements-uat.txt requirements-uat.txt righaven.sh righaven.sh versions.lock.json versions.lock.json View all files Repository files navigation
RigHaven is a home and an on-demand toolbox for AI agents. You describe an objective; a
coordinator composes a team of roles for it, picks the skills and the tools they need (99
managed CLI tools, from documents and video to CAD, data and networks), runs each role in its own
isolated container and follows the work through review and delivery. You decide at every step :
tool alternatives come with their trade-offs, each phase waits for your reply, and the runtime,
not the model's goodwill, enforces your policies before anything is delivered.
righaven-teaser.mp4
15-second teaser ( MP4 ), produced by RigHaven
itself: an EffectCraft motion-graphics role and a SoundCraft sound role working
from the brief “a teaser explaining RigHaven's architecture”. It is promotional
material; the guarantees RigHaven actually enforces are described below.
You want to…
In RigHaven
Guide
Get a document, a presentation, a video, a model or an analysis from a brief
New task : the coordinator builds the team and delivers files with editable sources
Workspace
Follow a proven path end to end
Recipes : phased projects (data pipeline, device monitoring, network assessment…) or your own
Recipes
Work on databases
Connections to PostgreSQL, SQL Server/Azure SQL/Fabric (also with your Microsoft account and MFA), MySQL, Oracle; catalog, replica, synthetic data, pipelines
Data engineering · Connections
Work with devices and IoT
MQTT, Modbus, HTTP, serial and USB boards, Bluetooth LE, Matter, Android TV, OBD-II; every write approved by you
Devices and IoT
Map your network and check it
Inventory and map of the networks you declare; each vulnerability check only after your approval
Network
Edit the project's files yourself
Files → Open with : JupyterLab, editors and a terminal on the same environment
Integrated editors
Set rules the work must respect
Policy Book : policy sheets checked on the finished product
Policy Book
Reuse and share environments
Environments, Gallery packs, export of configurations and prepared software
Environments
Use your PC, your phone or a dedicated machine from a server
RigHaven Agent, the Android app, on-premise environments
below
How tasks, recipes, teams, skills, environments, connections and policies fit together:
Concepts .
Where
What you need
Steps
One Windows PC
Windows 10/11 with WSL 2 (no Docker Desktop)
Download RigHaven-<version>-windows-x64-portable.zip from the latest release , extract it to e.g. C:\RigHaven , double-click RigHaven.exe . The browser opens http://127.0.0.1:2026 · details
One Linux PC
Docker Engine
Download RigHaven-<version>-linux-x64-portable.tar.gz , extract it and run ./righaven · details
A server (VPS) for you or a team
Docker Compose; on Hostinger, Traefik
Deploy the compose.hostinger.yaml attached to the release and open its HTTPS address · data across releases
From source (development)
Docker Compose
See Run from source
The first start downloads the verified images (about 1 GB, once).
2. Sign in and connect a model
Create the first account: it becomes the administrator . Later accounts are ordinary users
(if the administrator enables registration).
Open Models & resources and connect at least one provider: OpenAI Codex, Claude Code,
Antigravity or any OpenAI-compatible endpoint (OpenAI, DeepSeek, Kimi, OpenRouter…). Choose
the model of the coordinator and of the team. See model providers .
Press New task and write what you want, for example:
"A two-page PDF report comparing the three CSV files I attach, with one chart per file;
keep the editable sources." Attach the files.
Answer the coordinator's questions and choose among the tool alternatives it proposes. There is
no team to create and nothing to install first: the environment is prepared on demand.
Follow the task card above the chat: the goal, the progress, the decision waiting for you.
Each phase shows its result and waits for your reply.
Open the delivered files and their sources, then accept or ask for a focused revision.
Next: try a Recipe (Recipes → Use this recipe ), add a Connection to a database, or
change the language in Account & preferences → Appearance .
Connect your PC, your phone and your machines
When RigHaven runs on a server, it cannot see your desk. Three small pieces bring it there; all of
them connect to the server, so no router port is opened:
RigHaven can also be installed from the browser as an app, with a one-line task card on small
screens: phones and the installable app .
flowchart LR
Brief[Brief with you] --> Choice{Tool alternatives?}
Choice -- pros, cons, recommendation --> User[You choose]
Choice -- none --> Plans
User --> Plans[Execution plan + policy plan]
Plans --> Phases[Phases: roles in isolated environments]
Phases --> Review[Independent review]
Review --> Decision{Your decision}
Decision -- changes --> Phases
Decision -- approve --> Validation[Policy validation of the finished product]
Validation --> Delivery[Delivery with editable sources]
Loading
Isolated roles. Each assignment runs in its own container, with all capabilities dropped and
resource limits; it reads the project read-only and writes only its own staging folder. Other
roles see only what it deliberately publishes. Several environments can work in parallel within
the configured CPU/RAM.
Environments on demand. The software a task needs is installed into a reusable image, at
the versions you choose; an environment can be kept ready for recurring work.
Your approvals. Writes to devices and databases, vulnerability scans and phase results wait
for you; an approved action runs once, exactly as approved.
Network egress. With RIGHAVEN_WORKER_EGRESS=allowlist each job reaches only its model
provider and the domains its role declares, and every blocked connection is reported.
Personal work and shared infrastructure
Personal to each account
Shared within the installation
Chats, project files and outputs
Prepared software images and cache
Environment profiles, skill selections and preferences
Server CPU/RAM and execution capacity
Project experience and accepted lessons
Provider connections managed by the administrator
Private configurations until you publish them
The instance Gallery and approved remote catalogs
Administrators manage accounts ( User administration ), shared connections, default program
versions, capacity and software-image imports. Maintenance & reset removes selected
conversations, profiles or overrides after a typed confirmation. Details:
Workspace guide .
Requirements: Docker Compose and a platform that supports the privileged internal runtime
container (the application does not mount the host Docker socket).
git clone https://github.com/ggmm82/righaven.git
cd righaven
cp .env.example .env
Set a unique RIGHAVEN_PASSWORD of at least 16 characters in .env , then:
bash righaven.sh up
Open http://localhost:2026 . The local proxy first asks for RIGHAVEN_USER and
RIGHAVEN_PASSWORD (a gate, not an application account); then create or sign in to your
RigHaven account.
Document
Content
Concepts
Task, recipe, team, skills, environments, connections, policies, deliveries
Workspace guide
Using RigHaven step by step (English and Italian)
Recipes · Open chat
Guided projects and your own recipes; open-ended work with chosen resources
Connections
Databases, Microsoft sign-in with MFA, writing into a destination, cloud or on premise
Data engineering
Data tools and the end-to-end pipeline recipe
Devices and IoT · Network
Device connections and toolkit; network inventory and approved checks
Environments · Integrated editors
Tool catalog, preparation and reuse; JupyterLab and editors
Policy Book
Policy sheets, enforced phases and checkpoints
Model providers
Codex, Claude Code, Antigravity, OpenAI-compatible APIs
Windows · Linux portable apps
One-PC installations, commands and limits
RigHaven Agent · Android · On-premise environments
Your PC, phone and machines for a server
Data across releases
Upgrading while keeping the installation's data
Product status · Core boundary
What is complete, preview or planned
C4 model · BPMN lifecycle · Technical report
Architecture (produced by a RigHaven project about its own repository, checked against the code)
English is the default interface language; Italian is available.
python scripts/validate.py
python -m unittest discover -s tests -v
CI also exercises the real Gateway, authentication, private ownership, native skills, prepared
environments, output fixtures, portability, container replacement, the agent, compute nodes and
the Android app. A passing technical fixture does not certify artistic quality, engineering safety
or a complete public SaaS service.
RigHaven is MIT licensed . Incorporated source notices are preserved in
docs/UPSTREAM-LICENSE .
Marco Gramuglia
Computer Engineer | Data Engineer | Analytics & AI | IT Service Management
💼 LinkedIn: linkedin.com/in/marco-gramuglia-2686bb30
Stable releases use semantic tags such as v0.8.0 . VERSION defines the minimum
product version; successive stable releases automatically advance the patch number.
A larger minor/major change can raise this minimum. PR candidates use
v0.8.0-rc.pr.20.RUN.ATTEMPT . Retrying the same stable commit keeps its version.
The Overview displays the product version; commit and build IDs are under Technical
details. The release Compose supplies the exact version through RIGHAVEN_VERSION ,
including when the same verified images are promoted from preview to stable.
Release manifests retain the exact source commit and verified image digests: naming
a release does not rebuild the images already tested before merge.
Readme License Code of conduct
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to ggmm82/righaven development by creating an account on GitHub.

GitHub - ggmm82/righaven · GitHub
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
ggmm82
righaven Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme License Code of conduct
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
72 Commits 72 Commits Folders and files
.github/ workflows .github/ workflows config config core core desktop desktop docker docker docs docs frontend frontend locales locales righaven righaven scripts scripts skills/ custom skills/ custom tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md VERSION VERSION docker-compose.yaml docker-compose.yaml requirements-dev.txt requirements-dev.txt requirements-uat.txt requirements-uat.txt righaven.sh righaven.sh versions.lock.json versions.lock.json View all files Repository files navigation
RigHaven is a home and an on-demand toolbox for AI agents. You describe an objective; a
coordinator composes a team of roles for it, picks the skills and the tools they need (99
managed CLI tools, from documents and video to CAD, data and networks), runs each role in its own
isolated container and follows the work through review and delivery. You decide at every step :
tool alternatives come with their trade-offs, each phase waits for your reply, and the runtime,
not the model's goodwill, enforces your policies before anything is delivered.
righaven-teaser.mp4
15-second teaser ( MP4 ), produced by RigHaven
itself: an EffectCraft motion-graphics role and a SoundCraft sound role working
from the brief “a teaser explaining RigHaven's architecture”. It is promotional
material; the guarantees RigHaven actually enforces are described below.
You want to…
In RigHaven
Guide
Get a document, a presentation, a video, a model or an analysis from a brief
New task : the coordinator builds the team and delivers files with editable sources
Workspace
Follow a proven path end to end
Recipes : phased projects (data pipeline, device monitoring, network assessment…) or your own
Recipes
Work on databases
Connections to PostgreSQL, SQL Server/Azure SQL/Fabric (also with your Microsoft account and MFA), MySQL, Oracle; catalog, replica, synthetic data, pipelines
Data engineering · Connections
Work with devices and IoT
MQTT, Modbus, HTTP, serial and USB boards, Bluetooth LE, Matter, Android TV, OBD-II; every write approved by you
Devices and IoT
Map your network and check it
Inventory and map of the networks you declare; each vulnerability check only after your approval
Network
Edit the project's files yourself
Files → Open with : JupyterLab, editors and a terminal on the same environment
Integrated editors
Set rules the work must respect
Policy Book : policy sheets checked on the finished product
Policy Book
Reuse and share environments
Environments, Gallery packs, export of configurations and prepared software
Environments
Use your PC, your phone or a dedicated machine from a server
RigHaven Agent, the Android app, on-premise environments
below
How tasks, recipes, teams, skills, environments, connections and policies fit together:
Concepts .
Where
What you need
Steps
One Windows PC
Windows 10/11 with WSL 2 (no Docker Desktop)
Download RigHaven-<version>-windows-x64-portable.zip from the latest release , extract it to e.g. C:\RigHaven , double-click RigHaven.exe . The browser opens http://127.0.0.1:2026 · details
One Linux PC
Docker Engine
Download RigHaven-<version>-linux-x64-portable.tar.gz , extract it and run ./righaven · details
A server (VPS) for you or a team
Docker Compose; on Hostinger, Traefik
Deploy the compose.hostinger.yaml attached to the release and open its HTTPS address · data across releases
From source (development)
Docker Compose
See Run from source
The first start downloads the verified images (about 1 GB, once).
2. Sign in and connect a model
Create the first account: it becomes the administrator . Later accounts are ordinary users
(if the administrator enables registration).
Open Models & resources and connect at least one provider: OpenAI Codex, Claude Code,
Antigravity or any OpenAI-compatible endpoint (OpenAI, DeepSeek, Kimi, OpenRouter…). Choose
the model of the coordinator and of the team. See model providers .
Press New task and write what you want, for example:
"A two-page PDF report comparing the three CSV files I attach, with one chart per file;
keep the editable sources." Attach the files.
Answer the coordinator's questions and choose among the tool alternatives it proposes. There is
no team to create and nothing to install first: the environment is prepared on demand.
Follow the task card above the chat: the goal, the progress, the decision waiting for you.
Each phase shows its result and waits for your reply.
Open the delivered files and their sources, then accept or ask for a focused revision.
Next: try a Recipe (Recipes → Use this recipe ), add a Connection to a database, or
change the language in Account & preferences → Appearance .
Connect your PC, your phone and your machines
When RigHaven runs on a server, it cannot see your desk. Three small pieces bring it there; all of
them connect to the server, so no router port is opened:
RigHaven can also be installed from the browser as an app, with a one-line task card on small
screens: phones and the installable app .
flowchart LR
Brief[Brief with you] --> Choice{Tool alternatives?}
Choice -- pros, cons, recommendation --> User[You choose]
Choice -- none --> Plans
User --> Plans[Execution plan + policy plan]
Plans --> Phases[Phases: roles in isolated environments]
Phases --> Review[Independent review]
Review --> Decision{Your decision}
Decision -- changes --> Phases
Decision -- approve --> Validation[Policy validation of the finished product]
Validation --> Delivery[Delivery with editable sources]
Loading
Isolated roles. Each assignment runs in its own container, with all capabilities dropped and
resource limits; it reads the project read-only and writes only its own staging folder. Other
roles see only what it deliberately publishes. Several environments can work in parallel within
the configured CPU/RAM.
Environments on demand. The software a task needs is installed into a reusable image, at
the versions you choose; an environment can be kept ready for recurring work.
Your approvals. Writes to devices and databases, vulnerability scans and phase results wait
for you; an approved action runs once, exactly as approved.
Network egress. With RIGHAVEN_WORKER_EGRESS=allowlist each job reaches only its model
provider and the domains its role declares, and every blocked connection is reported.
Personal work and shared infrastructure
Personal to each account
Shared within the installation
Chats, project files and outputs
Prepared software images and cache
Environment profiles, skill selections and preferences
Server CPU/RAM and execution capacity
Project experience and accepted lessons
Provider connections managed by the administrator
Private configurations until you publish them
The instance Gallery and approved remote catalogs
Administrators manage accounts ( User administration ), shared connections, default program
versions, capacity and software-image imports. Maintenance & reset removes selected
conversations, profiles or overrides after a typed confirmation. Details:
Workspace guide .
Requirements: Docker Compose and a platform that supports the privileged internal runtime
container (the application does not mount the host Docker socket).
git clone https://github.com/ggmm82/righaven.git
cd righaven
cp .env.example .env
Set a unique RIGHAVEN_PASSWORD of at least 16 characters in .env , then:
bash righaven.sh up
Open http://localhost:2026 . The local proxy first asks for RIGHAVEN_USER and
RIGHAVEN_PASSWORD (a gate, not an application account); then create or sign in to your
RigHaven account.
Document
Content
Concepts
Task, recipe, team, skills, environments, connections, policies, deliveries
Workspace guide
Using RigHaven step by step (English and Italian)
Recipes · Open chat
Guided projects and your own recipes; open-ended work with chosen resources
Connections
Databases, Microsoft sign-in with MFA, writing into a destination, cloud or on premise
Data engineering
Data tools and the end-to-end pipeline recipe
Devices and IoT · Network
Device connections and toolkit; network inventory and approved checks
Environments · Integrated editors
Tool catalog, preparation and reuse; JupyterLab and editors
Policy Book
Policy sheets, enforced phases and checkpoints
Model providers
Codex, Claude Code, Antigravity, OpenAI-compatible APIs
Windows · Linux portable apps
One-PC installations, commands and limits
RigHaven Agent · Android · On-premise environments
Your PC, phone and machines for a server
Data across releases
Upgrading while keeping the installation's data
Product status · Core boundary
What is complete, preview or planned
C4 model · BPMN lifecycle · Technical report
Architecture (produced by a RigHaven project about its own repository, checked against the code)
English is the default interface language; Italian is available.
python scripts/validate.py
python -m unittest discover -s tests -v
CI also exercises the real Gateway, authentication, private ownership, native skills, prepared
environments, output fixtures, portability, container replacement, the agent, compute nodes and
the Android app. A passing technical fixture does not certify artistic quality, engineering safety
or a complete public SaaS service.
RigHaven is MIT licensed . Incorporated source notices are preserved in
docs/UPSTREAM-LICENSE .
Marco Gramuglia
Computer Engineer | Data Engineer | Analytics & AI | IT Service Management
💼 LinkedIn: linkedin.com/in/marco-gramuglia-2686bb30
Stable releases use semantic tags such as v0.8.0 . VERSION defines the minimum
product version; successive stable releases automatically advance the patch number.
A larger minor/major change can raise this minimum. PR candidates use
v0.8.0-rc.pr.20.RUN.ATTEMPT . Retrying the same stable commit keeps its version.
The Overview displays the product version; commit and build IDs are under Technical
details. The release Compose supplies the exact version through RIGHAVEN_VERSION ,
including when the same verified images are promoted from preview to stable.
Release manifests retain the exact source commit and verified image digests: naming
a release does not rebuild the images already tested before merge.
Readme License Code of conduct
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
