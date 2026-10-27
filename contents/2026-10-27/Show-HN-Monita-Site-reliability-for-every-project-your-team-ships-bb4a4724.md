---
source: "https://github.com/PT-Perkasa-Pilar-Utama/monita"
hn_url: "https://news.ycombinator.com/item?id=50259287"
title: "Show HN: Monita – Site reliability for every project your team ships"
article_title: "GitHub - PT-Perkasa-Pilar-Utama/monita: Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. · GitHub"
image: "https://opengraph.githubassets.com/e3b5ac631f0721c0c5ee8e96d9fd031c74226fdf952289b9c2603173f6dcd5e0/PT-Perkasa-Pilar-Utama/monita"
author: "maarut"
captured_at: "2026-10-27T14:02:16Z"
capture_tool: "hn-digest"
hn_id: 50259287
score: 1
comments: 0
posted_at: "2026-10-27T13:48:55Z"
tags:
  - hacker-news
---

# Show HN: Monita – Site reliability for every project your team ships

- HN: [50259287](https://news.ycombinator.com/item?id=50259287)
- Source: [github.com](https://github.com/PT-Perkasa-Pilar-Utama/monita)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T13:48:55Z

## Translation

Title: Show HN: Monita – Site reliability for every project your team ships
Article title: GitHub - PT-Perkasa-Pilar-Utama/monita: Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. · GitHub
Description: Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. - PT-Perkasa-Pilar-Utama/monita

Article text:
GitHub - PT-Perkasa-Pilar-Utama/monita: Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. · GitHub
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
PT-Perkasa-Pilar-Utama
Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file.
pt-perkasa-pilar-utama.github.io/monita/ Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
492 Commits 492 Commits Folders and files
.claude .claude .githooks .githooks .github .github .impeccable .impeccable deploy deploy docs docs e2e e2e licenses licenses migrations migrations scripts scripts seeds seeds site site src src static static tests tests web web .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile PRODUCT.md PRODUCT.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md build.rs build.rs deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Site reliability for every project your team ships.
Monita checks your team's websites and scheduled jobs, alerts the people you choose when one fails, and shows the outage on a public status page. It runs on your own server as one program with one SQLite file.
Website
Try the demo
Docs
Install
Monitors check a URL, a TCP port, a DNS name or a host by ping every 30 s to 5 min, and warn before an HTTPS certificate expires.
Heartbeats wait for your cron jobs to call in.
Incidents open after two failed checks or a missed heartbeat, and alert by email, Slack, Google Chat, Discord, Telegram, ntfy or PagerDuty.
Status pages show 90 days of uptime and the updates your team posts.
Tickets collect bug reports from the status page or from your app.
Maintenance windows, pause and locations keep planned work quiet and check from more than one place.
Roles, two-factor sign-in and scoped API keys let one Monita serve the whole IT department.
Click through the demo first. It runs on sample data and saves nothing.
On a 64-bit Linux server with systemd, 512 MB of memory and 5 GB of disk. On Arm, use monita_linux_arm64.tar.gz .
curl -fsSLO https://github.com/PT-Perkasa-Pilar-Utama/monita/releases/latest/download/monita_linux_amd64.tar.gz
tar -xzf monita_linux_amd64.tar.gz
sudo install -m 0755 monita /usr/local/bin/monita
sudo monita setup
sudo monita service install
monita setup asks for the public URL, the mail server and the first super admin. Put Caddy or nginx in front of 127.0.0.1:7466 for HTTPS.
To try it on a laptop with Docker, create the first account (it asks for a password of at least 8 characters):
docker volume create monita-data
docker run --rm -it -v monita-data:/data -e MONITA_ENV=dev ghcr.io/pt-perkasa-pilar-utama/monita:1 admin create --email you@company.co.id
Then start Monita and open http://localhost:7466 :
docker run -d --name monita -p 127.0.0.1:7466:7466 -v monita-data:/data -e MONITA_ENV=dev ghcr.io/pt-perkasa-pilar-utama/monita:1
The install guide covers a Docker server, macOS, Windows, upgrades and building from source.
URL=https://monita.example.com/ping/ < token >
if ! out= $( backup.sh 2>&1 ) ; then
printf ' %s ' " $out " | curl -fsS -m 10 --data-binary @- " $URL /fail "
exit 1
fi
curl -fsS -m 10 " $URL "
A missed deadline opens an incident. A call to /fail opens one at once, with the first 1 KB of the body as its cause.
Each tool checked 500 URLs every 60 s against one local nginx, alone in Docker on an Apple M1 (colima, 2 vCPU). Memory and CPU come from docker stats every 15 s for 5 minutes, after a 5-minute warm-up, on 2026-10-27.
The documentation is built from docs/guide/ .
You need Rust, Bun 1.4, a C compiler and Google Chrome. The Rust version in rust-toolchain.toml installs itself.
make setup # install the git hooks and the browser-test packages
make run # start Monita on http://127.0.0.1:7466
make check # check formatting and lints, then run the Rust tests
make e2e # run the browser tests on port 18080
Read CONTRIBUTING.md before you send a change. CLAUDE.md holds the project rules and docs/GLOSSARY.md the terms.
Report a security problem privately, as SECURITY.md describes. Everyone who takes part follows the Code of Conduct . GOVERNANCE.md names the maintainers, ROADMAP.md the direction and CHANGELOG.md each release.
Monita is released under the MIT license. Copyright PT. Perkasa Pilar Utama.
The binary embeds the Inter and Geist Mono fonts, which are under the SIL Open Font License 1.1. Their licences are in licenses/ and in every release archive.
Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file.
pt-perkasa-pilar-utama.github.io/monita/ Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. - PT-Perkasa-Pilar-Utama/monita

GitHub - PT-Perkasa-Pilar-Utama/monita: Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file. · GitHub
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
PT-Perkasa-Pilar-Utama
Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file.
pt-perkasa-pilar-utama.github.io/monita/ Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
492 Commits 492 Commits Folders and files
.claude .claude .githooks .githooks .github .github .impeccable .impeccable deploy deploy docs docs e2e e2e licenses licenses migrations migrations scripts scripts seeds seeds site site src src static static tests tests web web .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile PRODUCT.md PRODUCT.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md build.rs build.rs deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Site reliability for every project your team ships.
Monita checks your team's websites and scheduled jobs, alerts the people you choose when one fails, and shows the outage on a public status page. It runs on your own server as one program with one SQLite file.
Website
Try the demo
Docs
Install
Monitors check a URL, a TCP port, a DNS name or a host by ping every 30 s to 5 min, and warn before an HTTPS certificate expires.
Heartbeats wait for your cron jobs to call in.
Incidents open after two failed checks or a missed heartbeat, and alert by email, Slack, Google Chat, Discord, Telegram, ntfy or PagerDuty.
Status pages show 90 days of uptime and the updates your team posts.
Tickets collect bug reports from the status page or from your app.
Maintenance windows, pause and locations keep planned work quiet and check from more than one place.
Roles, two-factor sign-in and scoped API keys let one Monita serve the whole IT department.
Click through the demo first. It runs on sample data and saves nothing.
On a 64-bit Linux server with systemd, 512 MB of memory and 5 GB of disk. On Arm, use monita_linux_arm64.tar.gz .
curl -fsSLO https://github.com/PT-Perkasa-Pilar-Utama/monita/releases/latest/download/monita_linux_amd64.tar.gz
tar -xzf monita_linux_amd64.tar.gz
sudo install -m 0755 monita /usr/local/bin/monita
sudo monita setup
sudo monita service install
monita setup asks for the public URL, the mail server and the first super admin. Put Caddy or nginx in front of 127.0.0.1:7466 for HTTPS.
To try it on a laptop with Docker, create the first account (it asks for a password of at least 8 characters):
docker volume create monita-data
docker run --rm -it -v monita-data:/data -e MONITA_ENV=dev ghcr.io/pt-perkasa-pilar-utama/monita:1 admin create --email you@company.co.id
Then start Monita and open http://localhost:7466 :
docker run -d --name monita -p 127.0.0.1:7466:7466 -v monita-data:/data -e MONITA_ENV=dev ghcr.io/pt-perkasa-pilar-utama/monita:1
The install guide covers a Docker server, macOS, Windows, upgrades and building from source.
URL=https://monita.example.com/ping/ < token >
if ! out= $( backup.sh 2>&1 ) ; then
printf ' %s ' " $out " | curl -fsS -m 10 --data-binary @- " $URL /fail "
exit 1
fi
curl -fsS -m 10 " $URL "
A missed deadline opens an incident. A call to /fail opens one at once, with the first 1 KB of the body as its cause.
Each tool checked 500 URLs every 60 s against one local nginx, alone in Docker on an Apple M1 (colima, 2 vCPU). Memory and CPU come from docker stats every 15 s for 5 minutes, after a 5-minute warm-up, on 2026-10-27.
The documentation is built from docs/guide/ .
You need Rust, Bun 1.4, a C compiler and Google Chrome. The Rust version in rust-toolchain.toml installs itself.
make setup # install the git hooks and the browser-test packages
make run # start Monita on http://127.0.0.1:7466
make check # check formatting and lints, then run the Rust tests
make e2e # run the browser tests on port 18080
Read CONTRIBUTING.md before you send a change. CLAUDE.md holds the project rules and docs/GLOSSARY.md the terms.
Report a security problem privately, as SECURITY.md describes. Everyone who takes part follows the Code of Conduct . GOVERNANCE.md names the maintainers, ROADMAP.md the direction and CHANGELOG.md each release.
Monita is released under the MIT license. Copyright PT. Perkasa Pilar Utama.
The binary embeds the Inter and Geist Mono fonts, which are under the SIL Open Font License 1.1. Their licences are in licenses/ and in every release archive.
Site reliability for every project your team ships. Self-hosted uptime checks, cron heartbeats, incidents with email and webhook alerts, public status pages, bug tickets and changelogs, in one Rust binary and one SQLite file.
pt-perkasa-pilar-utama.github.io/monita/ Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
