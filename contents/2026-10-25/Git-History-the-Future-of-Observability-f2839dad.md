---
source: "https://github.com/premtonx/statetrack"
hn_url: "https://news.ycombinator.com/item?id=50236197"
title: "Git History the Future of Observability"
article_title: "GitHub - premtonx/statetrack: Infrastructure change intelligence - git history for live servers · GitHub"
image: "https://opengraph.githubassets.com/8c0e6fbd51dcf718a1fd3a5538b83c245e9cd9ddfd0a0f6efec1185168b601df/premtonx/statetrack"
author: "premtonx"
captured_at: "2026-10-25T21:28:50Z"
capture_tool: "hn-digest"
hn_id: 50236197
score: 2
comments: 1
posted_at: "2026-10-25T21:24:27Z"
tags:
  - hacker-news
---

# Git History the Future of Observability

- HN: [50236197](https://news.ycombinator.com/item?id=50236197)
- Source: [github.com](https://github.com/premtonx/statetrack)
- Score: 2
- Comments: 1
- Posted: 2026-10-25T21:24:27Z

## Translation

Title: Git History the Future of Observability
Article title: GitHub - premtonx/statetrack: Infrastructure change intelligence - git history for live servers · GitHub
Description: Infrastructure change intelligence - git history for live servers - premtonx/statetrack

Article text:
GitHub - premtonx/statetrack: Infrastructure change intelligence - git history for live servers · GitHub
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
premtonx
statetrack Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Infrastructure change intelligence - git history for live servers
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github/ workflows .github/ workflows agent agent deploy deploy docs docs integrations integrations pkg pkg proto/ statetrack/ v1 proto/ statetrack/ v1 scripts scripts server server web web .cursorrules .cursorrules .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md RELEASE_NOTES_v0.1.0.md RELEASE_NOTES_v0.1.0.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum infograph.jpg infograph.jpg spec.md spec.md View all files Repository files navigation
Infrastructure Change Intelligence — git history for live servers.
StateTrack answers the only question that matters at 3am: what changed?
It records every meaningful change to your Linux and Windows hosts — who changed
what, from where, with which tool, and what the file looked like before — and
puts that change on the same timeline as your CPU, latency and error-rate graphs.
A config edit, a package upgrade, a stopped IIS app pool, or a row flipped in
settings_db is captured as a reviewable diff with an actor attached, seconds
after it happens.
StateTrack.1.1.mp4
The console — fleet pulse, change activity and the latest changes timeline.
Grafana — the same Level 1/2 changes as point annotations on your metric graphs.
An outage is almost always preceded by a change. When it is, the answer is
scattered across shell histories, /var/log/dpkg.log , an Ansible log nobody
kept, and three people's memory — or it is gone.
Existing tools answer "is it up?" . StateTrack answers "what is different, who
did it, and what did it look like before?"
┌───────────────────────── Host (Linux / Windows / K8s) ─────────────────────────┐
│ statetrack-agent < 25 MB RSS · < 1% CPU · zero inbound ports │
│ │
│ collectors filesystem · registry · SCM · systemd · packages · IIS · DB · k8s│
│ │ │
│ ▼ │
│ classifier (L0–L4) → storm breaker → shadow diff → secret scrubber (on host) │
│ │ │
│ ▼ │
│ debouncer (3 s sliding window → parent event) → SQLite spool (200 MB) │
│ │ │
└───────┼──────────────── outbound gRPC 50051 (mTLS) + discovery heartbeat ──────┘
▼
┌──────────────────────────── statetrack-server ─────────────────────────────────┐
│ gRPC ingest → Postgres (partitioned events) → S3/SeaweedFS (large diffs) │
│ → Grafana · Slack · Teams · PagerDuty · SIEM · Jira · ServiceNow │
│ → AI gateway (offline heuristics / Ollama / OpenAI / Azure / Claude)│
│ → MCP server (JSON-RPC 2.0) → Prometheus /metrics │
│ → REST + WebSocket → embedded React console on the same origin │
└────────────────────────────────────────────────────────────────────────────────┘
Table of contents
Install an agent — direct downloads
Configuration: agent.yaml & targeted watchers
Reference
5. Architecture & runtime components
6. Impact levels & anti-flood defenses
7. Secret scrubber
8. Auto-discovery & the enterprise catalog
9. The web console
10. AI provider gateway
11. Model Context Protocol (MCP) server
12. Enterprise security
13. REST, WebSocket & gRPC API
14. Metrics & alerting
15. Development, build & verification
A fresh StateTrack install is empty by design : 0 hosts, 0 events, no demo
data. The console shows 0 / 0 agents online and an onboarding card until you
enrol your first host. Demo data is opt-in ( make seed ), never a default.
The first administrator is created on first boot. Its password is generated and
printed once :
docker compose logs statetrack | grep administrator
…or choose it yourself with STATETRACK_ADMIN_EMAIL / STATETRACK_ADMIN_PASSWORD .
Option A — 1-command quickstart (no clone, no toolchain)
curl -fsSL https://raw.githubusercontent.com/premtonx/statetrack/main/deploy/docker-compose.quickstart.yml \
-o docker-compose.yml
docker compose up -d
Open http://localhost:8080 . That file is self-contained: Postgres 17 with a
health check and a persistent volume, plus the pre-built
ghcr.io/premtonx/statetrack image. Nothing is compiled, nothing else is
needed. Migrations run automatically on first boot.
Change POSTGRES_PASSWORD and set STATETRACK_SESSION_SECRET before exposing
the port beyond localhost.
Option B — single docker run (bring your own Postgres)
docker run -d --name statetrack \
-p 8080:8080 -p 50051:50051 \
-e DATABASE_URL= ' postgres://user:pass@db.internal:5432/statetrack?sslmode=disable ' \
-e STATETRACK_SESSION_SECRET= ' change-me-to-a-long-random-string ' \
-v statetrack-data:/var/lib/statetrack \
--restart unless-stopped \
ghcr.io/premtonx/statetrack:latest
8080 serves the console, the REST/WS API, /metrics and /healthz ; 50051
is agent gRPC ingestion. Large diffs are stored inline in Postgres unless you
point S3_ENDPOINT at an S3-compatible store.
Option C — clone, build and run with Docker Compose
git clone https://github.com/premtonx/statetrack.git
cd statetrack
docker compose up -d --build
The multi-stage Dockerfile builds the React console ( node:20-alpine ),
compiles a static server with the console and the SQL migrations embedded
( golang:1.26-alpine , CGO_ENABLED=0 ), and copies one binary into an
alpine:3.20 runtime with a built-in HEALTHCHECK .
Optional profiles and demo data:
# Grafana :3000, Prometheus :9090, node_exporter :9100, Checkmk :5000/cmk
docker compose --profile monitoring up -d
# or, from a checkout: seed a realistic demo timeline anchored to today
make seed # Windows: .\scripts\dev.ps1 seed
make grafana-seed # import the dashboard + push annotations
Everything the server reads from the environment:
GET /api/v1/system reports the running version, commit, build date, Go
runtime, database version and storage backend — the console's Settings page
reads it, so nothing in the UI hardcodes a build string.
Direct downloads — release artifacts
Every artifact below is attached to the latest release . The latest/download links are stable: they always resolve to the newest published version, so a script written today keeps working after an upgrade.
Verify before you run. Every archive is listed in SHA256SUMS.txt . Download both into the same directory and check them — a mismatch means the file is corrupt or was tampered with in transit:
sha256sum -c --ignore-missing SHA256SUMS.txt # Linux
# Windows
$want = ( Select-String - Path SHA256SUMS.txt - Pattern ' statetrack-agent-windows-amd64.zip ' |
ForEach-Object { $_ .Line.Split ( ' ' )[ 0 ] })
$got = ( Get-FileHash statetrack - agent - windows - amd64.zip - Algorithm SHA256).Hash.ToLower()
if ( $want -eq $got ) { ' OK ' } else { throw ' checksum mismatch ' }
The sections below install the agent as a service, which is what you want on a
monitored host.
The agent is one static binary with no runtime dependencies: no JVM, no Python,
no sidecar. It dials out to :50051 and keeps one stream open, so a
monitored host has no inbound port and a compromised server cannot use the
fleet as an ingress path.
Before installing, create an enrollment token in Settings → Agents . The
console shows the exact command for this deployment, with the server address
already filled in.
Linux — option 1: one-line installer
curl -fsSL https://raw.githubusercontent.com/premtonx/statetrack/main/deploy/install.sh \
| sudo bash -s -- \
--server statetrack.internal.example.com:50051 \
--token st_enroll_xxxxxxxxxxxxxxxx \
--org production
The script detects x86_64 / aarch64 , downloads
statetrack-agent-linux-<arch>.tar.gz from the latest GitHub release, verifies
it against the release's SHA256SUMS.txt and refuses to install on a mismatch,
then writes /etc/statetrack/agent.yaml (mode 0600 ), installs the hardened
systemd unit and starts it.
Useful flags: --version v0.1.0 to pin a release, --binary ./statetrack-agent
for an air-gapped host, --url for an internal mirror, --force to overwrite an
existing config, --no-start to stage without starting.
Linux — option 2: download the package and set up systemd yourself
# 1. Download and unpack the release archive
cd /tmp
curl -fsSLO https://github.com/premtonx/statetrack/releases/latest/download/statetrack-agent-linux-amd64.tar.gz
curl -fsSLO https://github.com/premtonx/statetrack/releases/latest/download/SHA256SUMS.txt
sha256sum -c --ignore-missing SHA256SUMS.txt # must print: OK
tar -xzf statetrack-agent-linux-amd64.tar.gz
# → statetrack-agent install.sh statetrack-agent.service agent.yaml.example
# 2. Install the binary
sudo install -m 0755 statetrack-agent /usr/local/bin/statetrack-agent
/usr/local/bin/statetrack-agent --version
# 3. Write the configuration
sudo mkdir -p /etc/statetrack
sudo install -m 0600 -o root -g root agent.yaml.example /etc/statetrack/agent.yaml
sudoedit /etc/statetrack/agent.yaml # set server.address, server.token, server.org
# 4. Install and start the service
sudo install -m 0644 statetrack-agent.service /etc/systemd/system/statetrack-agent.service
sudo systemctl daemon-reload
sudo systemctl enable --now statetrack-agent
# 5. Verify
systemctl status statetrack-agent
journalctl -u statetrack-agent -f
The unit runs as root with exactly three capabilities
( CAP_SYS_ADMIN for fanotify, CAP_AUDIT_READ for attribution,
CAP_DAC_READ_SEARCH for /etc/shadow ), ProtectSystem=full , a private
/tmp , and no new privileges.
The host appears under Servers within seconds of the first heartbeat.
Windows — option 1: one-line PowerShell installer
& ([ scriptblock ]::Create((irm https: // raw.githubusercontent.com / premtonx / statetrack / main / deploy / install.ps1))) `
- Server statetrack.internal.example.com : 50051 `
- Token st_enroll_xxxxxxxxxxxxxxxx `
- Org production
It downloads statetrack-agent-windows-amd64.zip , verifies the SHA-256, extracts
to C:\ProgramData\StateTrack , writes agent.yaml with an ACL limited to
SYSTEM and Administrators , then registers and starts the StateTrackAgent
service.
Windows — option 2: download the ZIP and install the service
# 1. Download, verify and extract
$dir = ' C:\ProgramData\StateTrack '
New-Item - ItemType Directory - Force $dir | Out-Null
$base = ' https://github.com/premtonx/statetrack/releases/latest/download '
Invoke-WebRequest " $base /statetrack-agent-windows-amd64.zip " - OutFile " $ env: TEMP \agent.zip "

[truncated]

## Original Extract

Infrastructure change intelligence - git history for live servers - premtonx/statetrack

GitHub - premtonx/statetrack: Infrastructure change intelligence - git history for live servers · GitHub
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
premtonx
statetrack Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Infrastructure change intelligence - git history for live servers
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github/ workflows .github/ workflows agent agent deploy deploy docs docs integrations integrations pkg pkg proto/ statetrack/ v1 proto/ statetrack/ v1 scripts scripts server server web web .cursorrules .cursorrules .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md RELEASE_NOTES_v0.1.0.md RELEASE_NOTES_v0.1.0.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum infograph.jpg infograph.jpg spec.md spec.md View all files Repository files navigation
Infrastructure Change Intelligence — git history for live servers.
StateTrack answers the only question that matters at 3am: what changed?
It records every meaningful change to your Linux and Windows hosts — who changed
what, from where, with which tool, and what the file looked like before — and
puts that change on the same timeline as your CPU, latency and error-rate graphs.
A config edit, a package upgrade, a stopped IIS app pool, or a row flipped in
settings_db is captured as a reviewable diff with an actor attached, seconds
after it happens.
StateTrack.1.1.mp4
The console — fleet pulse, change activity and the latest changes timeline.
Grafana — the same Level 1/2 changes as point annotations on your metric graphs.
An outage is almost always preceded by a change. When it is, the answer is
scattered across shell histories, /var/log/dpkg.log , an Ansible log nobody
kept, and three people's memory — or it is gone.
Existing tools answer "is it up?" . StateTrack answers "what is different, who
did it, and what did it look like before?"
┌───────────────────────── Host (Linux / Windows / K8s) ─────────────────────────┐
│ statetrack-agent < 25 MB RSS · < 1% CPU · zero inbound ports │
│ │
│ collectors filesystem · registry · SCM · systemd · packages · IIS · DB · k8s│
│ │ │
│ ▼ │
│ classifier (L0–L4) → storm breaker → shadow diff → secret scrubber (on host) │
│ │ │
│ ▼ │
│ debouncer (3 s sliding window → parent event) → SQLite spool (200 MB) │
│ │ │
└───────┼──────────────── outbound gRPC 50051 (mTLS) + discovery heartbeat ──────┘
▼
┌──────────────────────────── statetrack-server ─────────────────────────────────┐
│ gRPC ingest → Postgres (partitioned events) → S3/SeaweedFS (large diffs) │
│ → Grafana · Slack · Teams · PagerDuty · SIEM · Jira · ServiceNow │
│ → AI gateway (offline heuristics / Ollama / OpenAI / Azure / Claude)│
│ → MCP server (JSON-RPC 2.0) → Prometheus /metrics │
│ → REST + WebSocket → embedded React console on the same origin │
└────────────────────────────────────────────────────────────────────────────────┘
Table of contents
Install an agent — direct downloads
Configuration: agent.yaml & targeted watchers
Reference
5. Architecture & runtime components
6. Impact levels & anti-flood defenses
7. Secret scrubber
8. Auto-discovery & the enterprise catalog
9. The web console
10. AI provider gateway
11. Model Context Protocol (MCP) server
12. Enterprise security
13. REST, WebSocket & gRPC API
14. Metrics & alerting
15. Development, build & verification
A fresh StateTrack install is empty by design : 0 hosts, 0 events, no demo
data. The console shows 0 / 0 agents online and an onboarding card until you
enrol your first host. Demo data is opt-in ( make seed ), never a default.
The first administrator is created on first boot. Its password is generated and
printed once :
docker compose logs statetrack | grep administrator
…or choose it yourself with STATETRACK_ADMIN_EMAIL / STATETRACK_ADMIN_PASSWORD .
Option A — 1-command quickstart (no clone, no toolchain)
curl -fsSL https://raw.githubusercontent.com/premtonx/statetrack/main/deploy/docker-compose.quickstart.yml \
-o docker-compose.yml
docker compose up -d
Open http://localhost:8080 . That file is self-contained: Postgres 17 with a
health check and a persistent volume, plus the pre-built
ghcr.io/premtonx/statetrack image. Nothing is compiled, nothing else is
needed. Migrations run automatically on first boot.
Change POSTGRES_PASSWORD and set STATETRACK_SESSION_SECRET before exposing
the port beyond localhost.
Option B — single docker run (bring your own Postgres)
docker run -d --name statetrack \
-p 8080:8080 -p 50051:50051 \
-e DATABASE_URL= ' postgres://user:pass@db.internal:5432/statetrack?sslmode=disable ' \
-e STATETRACK_SESSION_SECRET= ' change-me-to-a-long-random-string ' \
-v statetrack-data:/var/lib/statetrack \
--restart unless-stopped \
ghcr.io/premtonx/statetrack:latest
8080 serves the console, the REST/WS API, /metrics and /healthz ; 50051
is agent gRPC ingestion. Large diffs are stored inline in Postgres unless you
point S3_ENDPOINT at an S3-compatible store.
Option C — clone, build and run with Docker Compose
git clone https://github.com/premtonx/statetrack.git
cd statetrack
docker compose up -d --build
The multi-stage Dockerfile builds the React console ( node:20-alpine ),
compiles a static server with the console and the SQL migrations embedded
( golang:1.26-alpine , CGO_ENABLED=0 ), and copies one binary into an
alpine:3.20 runtime with a built-in HEALTHCHECK .
Optional profiles and demo data:
# Grafana :3000, Prometheus :9090, node_exporter :9100, Checkmk :5000/cmk
docker compose --profile monitoring up -d
# or, from a checkout: seed a realistic demo timeline anchored to today
make seed # Windows: .\scripts\dev.ps1 seed
make grafana-seed # import the dashboard + push annotations
Everything the server reads from the environment:
GET /api/v1/system reports the running version, commit, build date, Go
runtime, database version and storage backend — the console's Settings page
reads it, so nothing in the UI hardcodes a build string.
Direct downloads — release artifacts
Every artifact below is attached to the latest release . The latest/download links are stable: they always resolve to the newest published version, so a script written today keeps working after an upgrade.
Verify before you run. Every archive is listed in SHA256SUMS.txt . Download both into the same directory and check them — a mismatch means the file is corrupt or was tampered with in transit:
sha256sum -c --ignore-missing SHA256SUMS.txt # Linux
# Windows
$want = ( Select-String - Path SHA256SUMS.txt - Pattern ' statetrack-agent-windows-amd64.zip ' |
ForEach-Object { $_ .Line.Split ( ' ' )[ 0 ] })
$got = ( Get-FileHash statetrack - agent - windows - amd64.zip - Algorithm SHA256).Hash.ToLower()
if ( $want -eq $got ) { ' OK ' } else { throw ' checksum mismatch ' }
The sections below install the agent as a service, which is what you want on a
monitored host.
The agent is one static binary with no runtime dependencies: no JVM, no Python,
no sidecar. It dials out to :50051 and keeps one stream open, so a
monitored host has no inbound port and a compromised server cannot use the
fleet as an ingress path.
Before installing, create an enrollment token in Settings → Agents . The
console shows the exact command for this deployment, with the server address
already filled in.
Linux — option 1: one-line installer
curl -fsSL https://raw.githubusercontent.com/premtonx/statetrack/main/deploy/install.sh \
| sudo bash -s -- \
--server statetrack.internal.example.com:50051 \
--token st_enroll_xxxxxxxxxxxxxxxx \
--org production
The script detects x86_64 / aarch64 , downloads
statetrack-agent-linux-<arch>.tar.gz from the latest GitHub release, verifies
it against the release's SHA256SUMS.txt and refuses to install on a mismatch,
then writes /etc/statetrack/agent.yaml (mode 0600 ), installs the hardened
systemd unit and starts it.
Useful flags: --version v0.1.0 to pin a release, --binary ./statetrack-agent
for an air-gapped host, --url for an internal mirror, --force to overwrite an
existing config, --no-start to stage without starting.
Linux — option 2: download the package and set up systemd yourself
# 1. Download and unpack the release archive
cd /tmp
curl -fsSLO https://github.com/premtonx/statetrack/releases/latest/download/statetrack-agent-linux-amd64.tar.gz
curl -fsSLO https://github.com/premtonx/statetrack/releases/latest/download/SHA256SUMS.txt
sha256sum -c --ignore-missing SHA256SUMS.txt # must print: OK
tar -xzf statetrack-agent-linux-amd64.tar.gz
# → statetrack-agent install.sh statetrack-agent.service agent.yaml.example
# 2. Install the binary
sudo install -m 0755 statetrack-agent /usr/local/bin/statetrack-agent
/usr/local/bin/statetrack-agent --version
# 3. Write the configuration
sudo mkdir -p /etc/statetrack
sudo install -m 0600 -o root -g root agent.yaml.example /etc/statetrack/agent.yaml
sudoedit /etc/statetrack/agent.yaml # set server.address, server.token, server.org
# 4. Install and start the service
sudo install -m 0644 statetrack-agent.service /etc/systemd/system/statetrack-agent.service
sudo systemctl daemon-reload
sudo systemctl enable --now statetrack-agent
# 5. Verify
systemctl status statetrack-agent
journalctl -u statetrack-agent -f
The unit runs as root with exactly three capabilities
( CAP_SYS_ADMIN for fanotify, CAP_AUDIT_READ for attribution,
CAP_DAC_READ_SEARCH for /etc/shadow ), ProtectSystem=full , a private
/tmp , and no new privileges.
The host appears under Servers within seconds of the first heartbeat.
Windows — option 1: one-line PowerShell installer
& ([ scriptblock ]::Create((irm https: // raw.githubusercontent.com / premtonx / statetrack / main / deploy / install.ps1))) `
- Server statetrack.internal.example.com : 50051 `
- Token st_enroll_xxxxxxxxxxxxxxxx `
- Org production
It downloads statetrack-agent-windows-amd64.zip , verifies the SHA-256, extracts
to C:\ProgramData\StateTrack , writes agent.yaml with an ACL limited to
SYSTEM and Administrators , then registers and starts the StateTrackAgent
service.
Windows — option 2: download the ZIP and install the service
# 1. Download, verify and extract
$dir = ' C:\ProgramData\StateTrack '
New-Item - ItemType Directory - Force $dir | Out-Null
$base = ' https://github.com/premtonx/statetrack/releases/latest/download '
Invoke-WebRequest " $base /statetrack-agent-windows-amd64.zip " - OutFile " $ env: TEMP \agent.zip "

[truncated]
