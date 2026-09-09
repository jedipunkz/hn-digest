---
source: "https://github.com/kolapsis/maintenant"
hn_url: "https://news.ycombinator.com/item?id=49625407"
title: "Show HN: Maintenant, Docker and Kubernetes monitoring in a single Go binary"
article_title: "GitHub - kOlapsis/maintenant: Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. · GitHub"
image: "https://repository-images.githubusercontent.com/1169465369/1f76561f-7088-493c-9798-3ca271bb7e02"
author: "benjy3379"
captured_at: "2026-09-09T12:46:36Z"
capture_tool: "hn-digest"
hn_id: 49625407
score: 3
comments: 0
posted_at: "2026-09-09T12:25:11Z"
tags:
  - hacker-news
---

# Show HN: Maintenant, Docker and Kubernetes monitoring in a single Go binary

- HN: [49625407](https://news.ycombinator.com/item?id=49625407)
- Source: [github.com](https://github.com/kolapsis/maintenant)
- Score: 3
- Comments: 0
- Posted: 2026-09-09T12:25:11Z

## Translation

Title: Show HN: Maintenant, Docker and Kubernetes monitoring in a single Go binary
Article title: GitHub - kOlapsis/maintenant: Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. · GitHub
Description: Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. - kOlapsis/maintenant

Article text:
GitHub - kOlapsis/maintenant: Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. · GitHub
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
kOlapsis
/
maintenant
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
483 Commits 483 Commits Folders and files
.github .github cmd/ maintenant cmd/ maintenant deploy deploy docs docs frontend frontend internal internal proto proto scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .golangci.yml .golangci.yml .poutine.yml .poutine.yml .trivyignore .trivyignore CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md codeql-config.yml codeql-config.yml compose.authproxy.yml compose.authproxy.yml compose.test.postgres.yml compose.test.postgres.yml compose.test.yml compose.test.yml compose.yml compose.yml docker-entrypoint.sh docker-entrypoint.sh go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml View all files Repository files navigation
Drop a container. Your stack is monitored.
Docker, Kubernetes, uptime, TLS, cron jobs, live logs, image updates, CVEs: auto-discovered, alerting on every one of them,
from a single Go binary that idles under 30 MB of RAM. No PromQL, no exporters, no dashboards to build.
Quick Start • Why maintenant • Features • Documentation • Editions • Pricing
# docker-compose.yml
services :
maintenant :
image : ghcr.io/kolapsis/maintenant:latest
ports :
# ⚠️ SECURITY: publishes the UI/API (no authentication of their own) on
# every interface; put an auth reverse proxy in front, or bind "127.0.0.1:8080:8080". See https://docs.maintenant.dev/security/#reverse-proxy-setup.
- " 8080:8080 "
read_only : true
security_opt :
- no-new-privileges:true
tmpfs :
- /tmp:noexec,nosuid,size=64m
volumes :
- /var/run/docker.sock:/var/run/docker.sock:ro
- /proc:/host/proc:ro
- maintenant-data:/data
environment :
MAINTENANT_ADDR : " 0.0.0.0:8080 "
MAINTENANT_DB : " /data/maintenant.db "
restart : unless-stopped
volumes :
maintenant-data :
docker compose up -d
Open http://localhost:8080 . Your containers are already there, with their health, restart loops, resources and logs. Nothing to configure.
Docker socket access is automatic: the entrypoint reads the mounted socket's group and grants it to the unprivileged user, on Compose and on Swarm (where docker stack deploy silently ignores group_add ). If containers do not show up, see Troubleshooting .
kubectl apply -f deploy/kubernetes/
In-cluster API auto-detected, read-only RBAC, namespace filtering, workloads (Deployments, DaemonSets, StatefulSets) as first-class citizens. Kubernetes guide .
Bare Linux, no Docker at all (systemd, amd64 and arm64, statically linked)
curl -fsSL https://install.maintenant.dev | sudo bash
Endpoints, certificates and heartbeats work without any container runtime. Container monitoring switches on by itself the moment a runtime shows up. Install documentation for pinned versions, air-gapped installs and supply-chain verification.
Cloud : one cloud-init file boots a hardened host with maintenant running on Hetzner Cloud , DigitalOcean , Scaleway , OVHcloud or Vultr .
Monitoring your own infrastructure with the standard stack means running Prometheus, Grafana, Alertmanager, node-exporter, cAdvisor, blackbox-exporter, a certificate exporter, Loki, Promtail, Trivy and something for image updates. Ten-odd components, each with its own config, upgrades and dashboards, to answer one question: is my stack up, and what is burning?
maintenant answers that question with one container.
Do I still need Prometheus?
maintenant monitors your infrastructure . Prometheus monitors your application . There is no PromQL here, no custom exporters, no panels to design: maintenant already knows what a container, a certificate, an endpoint, a cron job and a CVE are, and starts watching them the moment they appear. If you ship business metrics and write your own queries, keep Prometheus for that. The two answer different questions, and plenty of people run both.
Against the tools usually stacked up next to it:
One container. One dashboard. Everything monitored.
Dashboard: uptime, response times, resources, unified monitors
Unified alerts across every source
Security posture with CVE enrichment and risk scoring (Personal)
More screenshots
Container auto-discovery
Resources per container and per host
Endpoint monitoring
Heartbeat and cron monitoring
TLS certificate tracking
Update intelligence
Network security insights
Your AI assistant, plugged in over MCP
Status page, all operational
Status page, degraded
Features
Every section links to its full documentation.
Zero-config auto-discovery for Docker, Docker Swarm and Kubernetes. Every container is tracked the moment it starts: state changes, health checks, restart loops, live log streaming with stdout/stderr demux. Compose projects are grouped automatically. Read-only: maintenant observes, it never touches your containers.
One central server , lightweight read-only agents on your other hosts, a persistent mutually-authenticated gRPC stream between them. No shared database, no message queue, no PKI to run: an agent enrolls with a one-time token and an Ed25519 keypair generated locally, and every stream is challenge-response authenticated. Revoke it from the UI at any time.
# On each remote host, one command, generated for you in the UI
docker run -d --name maintenant-agent --restart unless-stopped \
-v /var/run/docker.sock:/var/run/docker.sock:ro \
-v /proc:/host/proc:ro \
-v maintenant-agent-data:/var/lib/maintenant \
ghcr.io/kolapsis/maintenant:latest \
--mode=agent --server=grpcs://monitoring.example.com \
--enrollment-token=mnt_enr_XXXXXXXXXXXXXXXX --label= " prod-worker-01 "
Agents detect their local runtime (Docker, Swarm or Kubernetes), stream container state, endpoints, certificates, host CPU/memory/disk, and reconnect on their own. Every entity is attributed to its host, so nothing gets mixed across machines. Personal: up to 20 remote machines. Pro: unlimited.
Scans OCI registries and compares digests, so you know which images have an update before you docker pull blindly. Compose-aware update and rollback commands, with the right --project-directory . No Diun, no Watchtower, no extra container: it is part of the monitor.
HTTP and TCP checks declared as Docker labels, picked up when the container starts. Response times, uptime history, 90-day sparklines, failure and recovery thresholds.
labels :
maintenant.endpoint.http : " https://api:3000/health "
maintenant.endpoint.interval : " 15s "
maintenant.endpoint.failure-threshold : " 3 "
Heartbeat and cron monitoring
Create a monitor, get a URL, add one curl to the job. maintenant tracks start and finish, duration, exit code, and alerts when the deadline is missed.
curl -fsS -o /dev/null https://now.example.com/ping/{uuid}/ $?
TLS certificate monitoring
Auto-detected from your HTTPS endpoints, plus standalone monitors for any domain. Full chain validation, alerts at 30, 14, 7, 3 and 1 day before expiry, OCSP stapling checks (Personal).
Real-time CPU, memory, network and disk I/O per container and per host, top-consumers view for instant triage, per-container thresholds with debounce. History: 7 days on Community, 30 on Personal, 90 on Pro.
Flags what should not be there: ports bound to 0.0.0.0 , exposed database ports, host-network mode, privileged containers, Kubernetes NodePort and LoadBalancer services without a NetworkPolicy. Each image is mapped to its software ecosystem through OCI manifest inspection. Personal adds CVE enrichment, a risk score per container and a unified security posture dashboard.
One alert pipeline for every source: container restart loops and unhealthy checks, endpoint failures, missed heartbeats, expiring or invalid certificates, CPU and memory thresholds, available updates. Channels are silent by default and routed through triggers (severity, source, scope, tags). Silence rules for planned maintenance, exponential backoff on delivery.
Channels: Discord and webhooks (Community), email and Telegram (Personal), Slack and Microsoft Teams (Pro). Pro adds escalation policies that page the on-call, then the backup, then the lead, plus per-entity routing and maintenance windows.
Real-time status page with severity aggregation across every monitor, live over SSE. Personal adds incident timelines, Pro adds subscriber notifications (email and webhook) and branding.
Built-in Model Context Protocol server. Ask your AI assistant what is burning, read a container's logs, check the alert queue, acknowledge an alert, open an incident. stdio and Streamable HTTP transports, full OAuth2 for remote clients (Claude web, mobile and Desktop).
Everything is driven by Docker labels and a handful of environment variables . No YAML to maintain.
Environment variables : bind address, database, base URL, PostgreSQL DSN, MCP, Kubernetes namespaces, license key, telemetry.
Docker labels reference : endpoints, TLS, alert severity, restart thresholds, channel routing, grouping, ignore.
REST API under /api/v1/ , plus an SSE event stream.
services :
maintenant :
image : ghcr.io/kolapsis/maintenant:latest
ports :
# ⚠️ SECURITY: publishes the UI/API (no authentication of their own) on
# every interface; put an auth reverse proxy in front, or bind "127.0.0.1:8080:8080". See https://docs.maintenant.dev/security/#reverse-proxy-setup.
- " 8080:8080 "
read_only : true
security_opt :
- no-new-privileges:true
tmpfs :
- /tmp:noexec,nosuid,size=64m
volumes :
- /var/run/docker.sock:/var/run/docker.sock:ro
- /proc:/host/proc:ro
- maintenant-data:/data
environment :
MAINTENANT_ADDR : " 0.0.0.0:8080 "
MAINTENANT_DB : " /data/maintenant.db "
api :
image : myapp:latest
labels :
maintenant.group : " production "
maintenant.endpoint.http : " http://api:3000/health "
maintenant.endpoint.interval : " 15s "
maintenant.alert.severity : " critical "
maintenant.alert.channels : " ops-webhook "
postgres :
image : postgres:16
labels :
maintenant.endpoint.tcp : " postgres:5432 "
maintenant.alert.severity : " critical "
redis :
image : redis:7-alpine
labels :
maintenant.endpoint.tcp : " redis:6379 "
volumes :
maintenant-data :
Security model
No built-in authentication, by design. Like Dozzle and Prometheus, maintenant sits behind your reverse proxy and auth middleware (Traefik or Caddy, Authelia or Authentik). /ping/{uuid} and /st

[truncated]

## Original Extract

Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. - kOlapsis/maintenant

GitHub - kOlapsis/maintenant: Self-hosted monitoring for Docker, Kubernetes and uptime. Single Go binary, no agent config, live alerts. Drop a container, your stack is monitored. · GitHub
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
kOlapsis
/
maintenant
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
483 Commits 483 Commits Folders and files
.github .github cmd/ maintenant cmd/ maintenant deploy deploy docs docs frontend frontend internal internal proto proto scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .golangci.yml .golangci.yml .poutine.yml .poutine.yml .trivyignore .trivyignore CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md codeql-config.yml codeql-config.yml compose.authproxy.yml compose.authproxy.yml compose.test.postgres.yml compose.test.postgres.yml compose.test.yml compose.test.yml compose.yml compose.yml docker-entrypoint.sh docker-entrypoint.sh go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml View all files Repository files navigation
Drop a container. Your stack is monitored.
Docker, Kubernetes, uptime, TLS, cron jobs, live logs, image updates, CVEs: auto-discovered, alerting on every one of them,
from a single Go binary that idles under 30 MB of RAM. No PromQL, no exporters, no dashboards to build.
Quick Start • Why maintenant • Features • Documentation • Editions • Pricing
# docker-compose.yml
services :
maintenant :
image : ghcr.io/kolapsis/maintenant:latest
ports :
# ⚠️ SECURITY: publishes the UI/API (no authentication of their own) on
# every interface; put an auth reverse proxy in front, or bind "127.0.0.1:8080:8080". See https://docs.maintenant.dev/security/#reverse-proxy-setup.
- " 8080:8080 "
read_only : true
security_opt :
- no-new-privileges:true
tmpfs :
- /tmp:noexec,nosuid,size=64m
volumes :
- /var/run/docker.sock:/var/run/docker.sock:ro
- /proc:/host/proc:ro
- maintenant-data:/data
environment :
MAINTENANT_ADDR : " 0.0.0.0:8080 "
MAINTENANT_DB : " /data/maintenant.db "
restart : unless-stopped
volumes :
maintenant-data :
docker compose up -d
Open http://localhost:8080 . Your containers are already there, with their health, restart loops, resources and logs. Nothing to configure.
Docker socket access is automatic: the entrypoint reads the mounted socket's group and grants it to the unprivileged user, on Compose and on Swarm (where docker stack deploy silently ignores group_add ). If containers do not show up, see Troubleshooting .
kubectl apply -f deploy/kubernetes/
In-cluster API auto-detected, read-only RBAC, namespace filtering, workloads (Deployments, DaemonSets, StatefulSets) as first-class citizens. Kubernetes guide .
Bare Linux, no Docker at all (systemd, amd64 and arm64, statically linked)
curl -fsSL https://install.maintenant.dev | sudo bash
Endpoints, certificates and heartbeats work without any container runtime. Container monitoring switches on by itself the moment a runtime shows up. Install documentation for pinned versions, air-gapped installs and supply-chain verification.
Cloud : one cloud-init file boots a hardened host with maintenant running on Hetzner Cloud , DigitalOcean , Scaleway , OVHcloud or Vultr .
Monitoring your own infrastructure with the standard stack means running Prometheus, Grafana, Alertmanager, node-exporter, cAdvisor, blackbox-exporter, a certificate exporter, Loki, Promtail, Trivy and something for image updates. Ten-odd components, each with its own config, upgrades and dashboards, to answer one question: is my stack up, and what is burning?
maintenant answers that question with one container.
Do I still need Prometheus?
maintenant monitors your infrastructure . Prometheus monitors your application . There is no PromQL here, no custom exporters, no panels to design: maintenant already knows what a container, a certificate, an endpoint, a cron job and a CVE are, and starts watching them the moment they appear. If you ship business metrics and write your own queries, keep Prometheus for that. The two answer different questions, and plenty of people run both.
Against the tools usually stacked up next to it:
One container. One dashboard. Everything monitored.
Dashboard: uptime, response times, resources, unified monitors
Unified alerts across every source
Security posture with CVE enrichment and risk scoring (Personal)
More screenshots
Container auto-discovery
Resources per container and per host
Endpoint monitoring
Heartbeat and cron monitoring
TLS certificate tracking
Update intelligence
Network security insights
Your AI assistant, plugged in over MCP
Status page, all operational
Status page, degraded
Features
Every section links to its full documentation.
Zero-config auto-discovery for Docker, Docker Swarm and Kubernetes. Every container is tracked the moment it starts: state changes, health checks, restart loops, live log streaming with stdout/stderr demux. Compose projects are grouped automatically. Read-only: maintenant observes, it never touches your containers.
One central server , lightweight read-only agents on your other hosts, a persistent mutually-authenticated gRPC stream between them. No shared database, no message queue, no PKI to run: an agent enrolls with a one-time token and an Ed25519 keypair generated locally, and every stream is challenge-response authenticated. Revoke it from the UI at any time.
# On each remote host, one command, generated for you in the UI
docker run -d --name maintenant-agent --restart unless-stopped \
-v /var/run/docker.sock:/var/run/docker.sock:ro \
-v /proc:/host/proc:ro \
-v maintenant-agent-data:/var/lib/maintenant \
ghcr.io/kolapsis/maintenant:latest \
--mode=agent --server=grpcs://monitoring.example.com \
--enrollment-token=mnt_enr_XXXXXXXXXXXXXXXX --label= " prod-worker-01 "
Agents detect their local runtime (Docker, Swarm or Kubernetes), stream container state, endpoints, certificates, host CPU/memory/disk, and reconnect on their own. Every entity is attributed to its host, so nothing gets mixed across machines. Personal: up to 20 remote machines. Pro: unlimited.
Scans OCI registries and compares digests, so you know which images have an update before you docker pull blindly. Compose-aware update and rollback commands, with the right --project-directory . No Diun, no Watchtower, no extra container: it is part of the monitor.
HTTP and TCP checks declared as Docker labels, picked up when the container starts. Response times, uptime history, 90-day sparklines, failure and recovery thresholds.
labels :
maintenant.endpoint.http : " https://api:3000/health "
maintenant.endpoint.interval : " 15s "
maintenant.endpoint.failure-threshold : " 3 "
Heartbeat and cron monitoring
Create a monitor, get a URL, add one curl to the job. maintenant tracks start and finish, duration, exit code, and alerts when the deadline is missed.
curl -fsS -o /dev/null https://now.example.com/ping/{uuid}/ $?
TLS certificate monitoring
Auto-detected from your HTTPS endpoints, plus standalone monitors for any domain. Full chain validation, alerts at 30, 14, 7, 3 and 1 day before expiry, OCSP stapling checks (Personal).
Real-time CPU, memory, network and disk I/O per container and per host, top-consumers view for instant triage, per-container thresholds with debounce. History: 7 days on Community, 30 on Personal, 90 on Pro.
Flags what should not be there: ports bound to 0.0.0.0 , exposed database ports, host-network mode, privileged containers, Kubernetes NodePort and LoadBalancer services without a NetworkPolicy. Each image is mapped to its software ecosystem through OCI manifest inspection. Personal adds CVE enrichment, a risk score per container and a unified security posture dashboard.
One alert pipeline for every source: container restart loops and unhealthy checks, endpoint failures, missed heartbeats, expiring or invalid certificates, CPU and memory thresholds, available updates. Channels are silent by default and routed through triggers (severity, source, scope, tags). Silence rules for planned maintenance, exponential backoff on delivery.
Channels: Discord and webhooks (Community), email and Telegram (Personal), Slack and Microsoft Teams (Pro). Pro adds escalation policies that page the on-call, then the backup, then the lead, plus per-entity routing and maintenance windows.
Real-time status page with severity aggregation across every monitor, live over SSE. Personal adds incident timelines, Pro adds subscriber notifications (email and webhook) and branding.
Built-in Model Context Protocol server. Ask your AI assistant what is burning, read a container's logs, check the alert queue, acknowledge an alert, open an incident. stdio and Streamable HTTP transports, full OAuth2 for remote clients (Claude web, mobile and Desktop).
Everything is driven by Docker labels and a handful of environment variables . No YAML to maintain.
Environment variables : bind address, database, base URL, PostgreSQL DSN, MCP, Kubernetes namespaces, license key, telemetry.
Docker labels reference : endpoints, TLS, alert severity, restart thresholds, channel routing, grouping, ignore.
REST API under /api/v1/ , plus an SSE event stream.
services :
maintenant :
image : ghcr.io/kolapsis/maintenant:latest
ports :
# ⚠️ SECURITY: publishes the UI/API (no authentication of their own) on
# every interface; put an auth reverse proxy in front, or bind "127.0.0.1:8080:8080". See https://docs.maintenant.dev/security/#reverse-proxy-setup.
- " 8080:8080 "
read_only : true
security_opt :
- no-new-privileges:true
tmpfs :
- /tmp:noexec,nosuid,size=64m
volumes :
- /var/run/docker.sock:/var/run/docker.sock:ro
- /proc:/host/proc:ro
- maintenant-data:/data
environment :
MAINTENANT_ADDR : " 0.0.0.0:8080 "
MAINTENANT_DB : " /data/maintenant.db "
api :
image : myapp:latest
labels :
maintenant.group : " production "
maintenant.endpoint.http : " http://api:3000/health "
maintenant.endpoint.interval : " 15s "
maintenant.alert.severity : " critical "
maintenant.alert.channels : " ops-webhook "
postgres :
image : postgres:16
labels :
maintenant.endpoint.tcp : " postgres:5432 "
maintenant.alert.severity : " critical "
redis :
image : redis:7-alpine
labels :
maintenant.endpoint.tcp : " redis:6379 "
volumes :
maintenant-data :
Security model
No built-in authentication, by design. Like Dozzle and Prometheus, maintenant sits behind your reverse proxy and auth middleware (Traefik or Caddy, Authelia or Authentik). /ping/{uuid} and /st

[truncated]
