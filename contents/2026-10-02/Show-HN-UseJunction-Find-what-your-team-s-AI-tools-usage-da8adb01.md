---
source: "https://github.com/use-junction/usejunction"
hn_url: "https://news.ycombinator.com/item?id=49936836"
title: "Show HN: UseJunction – Find what your team's AI tools usage"
article_title: "GitHub - use-junction/usejunction: AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. · GitHub"
image: "https://opengraph.githubassets.com/a65ba5086c3ba2b78d68ce79734ae9ab8bfbe3a70152cec05eb60dd85c257fff/use-junction/usejunction"
author: "Dinuda"
captured_at: "2026-10-02T19:01:01Z"
capture_tool: "hn-digest"
hn_id: 49936836
score: 1
comments: 1
posted_at: "2026-10-02T18:26:16Z"
tags:
  - hacker-news
---

# Show HN: UseJunction – Find what your team's AI tools usage

- HN: [49936836](https://news.ycombinator.com/item?id=49936836)
- Source: [github.com](https://github.com/use-junction/usejunction)
- Score: 1
- Comments: 1
- Posted: 2026-10-02T18:26:16Z

## Translation

Title: Show HN: UseJunction – Find what your team's AI tools usage
Article title: GitHub - use-junction/usejunction: AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. · GitHub
Description: AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. - use-junction/usejunction
HN text: A few months ago, we started getting into the problem of actually seeing how our team's AI usage and its cost actually translates into a business outcome. We recently rolled out a public version open source after a few months of pilot testing the product. The product is UseJunction ( https://usejunction.dev ), and if your team is running multiple subscriptions/api from Cursor, Claude, Codex etc. this is for you to manage subscription spend, wastage and predictions. A bit of background, we had a team of 20 people running $200 Cursor subscriptions, and some were using a few others. That was over $4000 of spend. With this tool, we were able to get that to $2600 on our next billing cycle! That is a 35% decrease. We knew that people had to be facing this problem too, so we made it opensource for the community. All of these subscription providers leave behind telemetry on local machines, we use that to estimate usage. Soon we will be releasing a feature to attribute that spend to actual business outcomes that come from your project management tools. Currently considering actually adding a level of control on top of observability so that teams/orgs can actually govern AI tool usage.

Article text:
GitHub - use-junction/usejunction: AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. · GitHub
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
use-junction
/
usejunction
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
103 Commits 103 Commits Folders and files
.github/ workflows .github/ workflows agent agent apps/ admin apps/ admin docs docs experiments/ workflow-benchmark-lab experiments/ workflow-benchmark-lab infra infra packages packages scripts scripts video video .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md install.ps1 install.ps1 install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml usejunction-homepage-viewport.png usejunction-homepage-viewport.png usejunction-homepage.png usejunction-homepage.png View all files Repository files navigation
AI coding spend management for engineering teams. UseJunction tracks usage, cost, latency, plan utilization, seat waste, and configuration health across Codex, Claude Code, Cursor, GitHub Copilot, local models, and more.
Site: usejunction.dev · Solutions: AI coding spend management · Seat utilization · Plan usage · Guides: Plan usage & waste · Team AI coding insights · llms.txt
AI coding spend management for teams
Engineering and platform teams use UseJunction to answer four operational questions:
Which AI coding tools and models are developers actually using?
What is the estimated cost by developer, team, tool, and model?
Which paid seats are idle, underused, or approaching plan limits?
Which devices are missing coverage or using personal keys?
UseJunction is open source and self-hostable under the UseJunction Community License. It provides visibility across Cursor, Claude Code, Codex, GitHub Copilot, Continue, Cline, Roo Code, OpenCode, Ollama, LM Studio, and related runtimes without keystroke surveillance, browser capture, or full network interception. Start with the team AI spend solutions , then read the WakaTime comparison if you are evaluating editor-time tracking versus AI tool observability.
UseJunction is currently maintained by a single independent developer. The project is open for evaluation, feedback, and self-hosted use today, with a fuller community contribution setup coming soon.
UseJunction separates verified usage (vendor-reported charges when billable, e.g. Cursor chargedCents > 0 ) from estimated usage (local scans and rate-card pricing — including Cursor included/plan usage when chargedCents = 0 ). The combinations below have been tested on real machines and confirmed to surface usage correctly in the admin UI.
Other tools and platforms are supported by the agent collector; this table will grow as additional stacks are validated end-to-end.
Run the entire stack in Docker — admin on :3001 (host; configurable via ADMIN_HOST_PORT ), Langfuse on :3000 , LiteLLM on :4000 .
cp .env.example .env
# Optional: add provider keys to test real LiteLLM completions
# OPENAI_API_KEY=sk-...
# ANTHROPIC_API_KEY=sk-ant-...
# Without keys, full-stack E2E still passes by verifying the ingest API directly.
cd infra
docker compose build admin
docker compose up -d
docker compose ps # wait until all services are healthy
If port 3001 is taken:
ADMIN_HOST_PORT=3020 docker compose up -d
ADMIN_URL=http://localhost:3020 ./run-e2e.sh
Langfuse project keys (one-time, for traces):
Open http://localhost:3000 → create account → create project
Copy Public Key and Secret Key into root .env
Restart LiteLLM: cd infra && docker compose restart litellm
The admin container runs prisma db push and seeds seed-org plus a demo enrollment token on first start.
chmod +x scripts/full-stack-e2e.sh
./scripts/full-stack-e2e.sh
# or from infra/
./run-e2e.sh
Manual gateway request (use a user id from Developers ):
curl http://localhost:4000/v1/chat/completions \
-H " Authorization: Bearer sk-usejunction-master " \
-H " Content-Type: application/json " \
-H " x-usejunction-user: <userId> " \
-H " x-usejunction-tool: codex " \
-d ' {"model":"gpt-4o-mini","messages":[{"role":"user","content":"ping"}]} '
Service
URL
Admin UI
http://localhost:3001 ( admin@example.com / admin )
Langfuse
http://localhost:3000
LiteLLM
http://localhost:4000
Postgres (host)
localhost:5432
Hybrid local dev
cp .env.example .env
cd infra
docker compose up -d postgres langfuse-db litellm-db langfuse litellm
# Wait for DBs, then start admin locally:
cd ..
pnpm install
pnpm db:push
pnpm db:seed
pnpm dev
Admin UI: http://localhost:3001
Generate a developer-bound enrollment token after signing in and joining the organization:
curl -X POST http://localhost:3001/api/me/enrollment-token \
-H " Cookie: uj_session=... " | jq
Install the local agent
From a repo checkout (builds the Go agent locally — preferred for development):
chmod +x install.sh
./install.sh --token < token > --url http://localhost:3001
# enrolls, enables Claude OTEL, sends first report, and starts the daemon
One-liner (downloads a prebuilt binary from the control plane, or builds from source if the repo is on disk):
# optional for pnpm/dev without Docker: publish binaries into apps/admin/public
./scripts/build-agent-releases.sh 0.2.0
curl -fsSL http://localhost:3001/install.sh | sh -s -- --token < token > --url http://localhost:3001
The installer adds ~/.usejunction/bin to your shell PATH . Open a new terminal (or run export PATH="$HOME/.usejunction/bin:$PATH" ) before using usejunction commands.
Windows 10/11 PowerShell (x64 or ARM64, no administrator shell required):
powershell.exe - NoProfile - ExecutionPolicy Bypass - Command " & ([scriptblock]::Create((Invoke-RestMethod -UseBasicParsing 'http://localhost:3001/install.ps1'))) -Token '<token>' -Url 'http://localhost:3001' "
The Windows installer adds %USERPROFILE%\.usejunction\bin to your user PATH . Open a new terminal before using usejunction commands.
Teammate connect uses the shared team invite link ( /i/<token> ). After signing in there, the UI shows the install command with an enrollment --token .
The onboarding and invite screens provide a separate Windows PowerShell command. Windows installs run through a per-user Scheduled Task at logon and collect native Windows coding-tool data; WSL stores are not scanned.
cd agent && go build -o usejunction .
./usejunction enroll --token < token > --url http://localhost:3001
./usejunction doctor
./usejunction report
Two agents on one Mac (production + local dev)
UseJunction supports running production and local dev agents side by side:
Enroll production from your hosted control plane (e.g. https://usejunction.dev ).
Enroll local dev from http://localhost:3001 — the installer auto-selects the test profile for loopback URLs.
Both daemons can run at the same time without clobbering each other's enrollment.
Hot-reload the local agent (development)
After the test agent is enrolled once, rebuild and reinstall into ~/.usejunction-test whenever agent/ changes:
# admin + agent watcher (rebuilds agent on start and on agent/ changes)
pnpm dev
# or: ./scripts/dev-start.sh
# admin only (no agent rebuild/watch)
pnpm dev:admin
# one-shot rebuild + swap + daemon restart
pnpm agent:reinstall
# or: ./scripts/dev-agent-reinstall.sh
# watch agent sources and reinstall on change
pnpm dev:agent
# or: ./scripts/dev-agent-watch.sh
Requires an existing ~/.usejunction-test/config.json (from ./install.sh --token … --url http://localhost:3001 or the connect curl). This path stamps a 0.0.0-dev.<sha>.<unix> version, swaps the local binary/app bundle, and restarts launchd/systemd. It does not publish a control-plane release or enroll a new device.
Set USEJUNCTION_PROFILE=default to rebuild the production agent home ( ~/.usejunction ) instead.
When you enroll against a local control plane ( http://localhost:3001 ), /install.sh injects USEJUNCTION_ROOT and USEJUNCTION_PROFILE=test so curl | sh builds the agent from this checkout as 0.0.0-dev.* into ~/.usejunction-test instead of downloading a published release or touching production enrollment. Production hosts still serve the plain customer installer (published releases only).
Install gotchas: A prior pnpm agent:reinstall writes ~/.usejunction-test/dev-source , so later curl | sh against prod may still build 0.0.0-dev.* if a dev pin exists under the target profile home. Production customer installs also require a promoted release ( GET /api/agent-releases/latest must return 200); a GitHub agent-v* tag alone is not enough. See Install script behavior (prod vs dev) .
For faster change detection, install fswatch ( brew install fswatch ). Without it, the watcher polls every ~750ms.
┌─────────────────────────────────────────────────────────────────────────┐
│ Developer machines │
│ │
│ Codex / Claude / Cursor / Copilot / OpenCode / Antigravity / … │
│ │ │
│ ▼ │
│ Go agent (profile-isolated: ~/.usejunction or ~/.usejunction-test) │
│ • heartbeat (15m) + OTA update directives │
│ • local scans (JSONL / sqlite) → estimated_api │
│ • Cursor usage events (chargedCents) → verified_usage │
│ or rate-card estimated_api when included usage is $0 │
│ • quotas, accounts, tools inventory │
│ • optional Signals / work extraction │
│ • localhost sync endpoint (47832 default / 47833 test) │
└───────────────────────────────┬─────────────────────────────────────────┘
│ UUS sync (start → chunk → commit)
│ OTEL metrics (Claude)
│ heartbeats / agent-update events
▼
┌─────────────────────────────────────────────────────────────────────────┐
│ Control plane (apps/admin · Next.js) │
│ │
│ Ingest → UsageDaily (+ inventory / quotas) │
│ Source priority: vendor_verified > otel > device_observed > estimated │
│ Cost kinds: actual_spend · verified_usage · estimated_api │
│ │
│ Org-day snapshots → dashboard KPIs / tool detail / Models tables │
│ Sync team = wake agents to upload (does not install agent binaries) │
│ Agent OTA = tag agent-v* → promote → heartbeat directive │
│ normal = 24h staggered · critical = immediate │
└───────────────────────────────┬─────────────────────────────────────────┘
│
▼
PostgreSQL
Optional gateway path (self-hosted Docker / LiteLLM) still exists for traced proxy traffic:
Coding tools → LiteLLM → Providers
↓
Langfuse traces
↓
UseJunction callback → Admin API
Cost semantics (short): dashboard Estimated Usage = verified_usage + estimated_api . Cursor plan/bonus rows with chargedCents = 0 are estimated from the rate card after agent rematerialize — they are not labeled verified at $0. See Usage Accounting .
Reads: analytical queries go through UsageDaily , the S

[truncated]

## Original Extract

AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. - use-junction/usejunction

A few months ago, we started getting into the problem of actually seeing how our team's AI usage and its cost actually translates into a business outcome. We recently rolled out a public version open source after a few months of pilot testing the product. The product is UseJunction ( https://usejunction.dev ), and if your team is running multiple subscriptions/api from Cursor, Claude, Codex etc. this is for you to manage subscription spend, wastage and predictions. A bit of background, we had a team of 20 people running $200 Cursor subscriptions, and some were using a few others. That was over $4000 of spend. With this tool, we were able to get that to $2600 on our next billing cycle! That is a 35% decrease. We knew that people had to be facing this problem too, so we made it opensource for the community. All of these subscription providers leave behind telemetry on local machines, we use that to estimate usage. Soon we will be releasing a feature to attribute that spend to actual business outcomes that come from your project management tools. Currently considering actually adding a level of control on top of observability so that teams/orgs can actually govern AI tool usage.

GitHub - use-junction/usejunction: AI coding spend management for engineering teams. Track Cursor, Claude Code, Codex, Copilot, usage, cost, plan utilization, seat waste, and device health. · GitHub
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
use-junction
/
usejunction
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
103 Commits 103 Commits Folders and files
.github/ workflows .github/ workflows agent agent apps/ admin apps/ admin docs docs experiments/ workflow-benchmark-lab experiments/ workflow-benchmark-lab infra infra packages packages scripts scripts video video .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md install.ps1 install.ps1 install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml usejunction-homepage-viewport.png usejunction-homepage-viewport.png usejunction-homepage.png usejunction-homepage.png View all files Repository files navigation
AI coding spend management for engineering teams. UseJunction tracks usage, cost, latency, plan utilization, seat waste, and configuration health across Codex, Claude Code, Cursor, GitHub Copilot, local models, and more.
Site: usejunction.dev · Solutions: AI coding spend management · Seat utilization · Plan usage · Guides: Plan usage & waste · Team AI coding insights · llms.txt
AI coding spend management for teams
Engineering and platform teams use UseJunction to answer four operational questions:
Which AI coding tools and models are developers actually using?
What is the estimated cost by developer, team, tool, and model?
Which paid seats are idle, underused, or approaching plan limits?
Which devices are missing coverage or using personal keys?
UseJunction is open source and self-hostable under the UseJunction Community License. It provides visibility across Cursor, Claude Code, Codex, GitHub Copilot, Continue, Cline, Roo Code, OpenCode, Ollama, LM Studio, and related runtimes without keystroke surveillance, browser capture, or full network interception. Start with the team AI spend solutions , then read the WakaTime comparison if you are evaluating editor-time tracking versus AI tool observability.
UseJunction is currently maintained by a single independent developer. The project is open for evaluation, feedback, and self-hosted use today, with a fuller community contribution setup coming soon.
UseJunction separates verified usage (vendor-reported charges when billable, e.g. Cursor chargedCents > 0 ) from estimated usage (local scans and rate-card pricing — including Cursor included/plan usage when chargedCents = 0 ). The combinations below have been tested on real machines and confirmed to surface usage correctly in the admin UI.
Other tools and platforms are supported by the agent collector; this table will grow as additional stacks are validated end-to-end.
Run the entire stack in Docker — admin on :3001 (host; configurable via ADMIN_HOST_PORT ), Langfuse on :3000 , LiteLLM on :4000 .
cp .env.example .env
# Optional: add provider keys to test real LiteLLM completions
# OPENAI_API_KEY=sk-...
# ANTHROPIC_API_KEY=sk-ant-...
# Without keys, full-stack E2E still passes by verifying the ingest API directly.
cd infra
docker compose build admin
docker compose up -d
docker compose ps # wait until all services are healthy
If port 3001 is taken:
ADMIN_HOST_PORT=3020 docker compose up -d
ADMIN_URL=http://localhost:3020 ./run-e2e.sh
Langfuse project keys (one-time, for traces):
Open http://localhost:3000 → create account → create project
Copy Public Key and Secret Key into root .env
Restart LiteLLM: cd infra && docker compose restart litellm
The admin container runs prisma db push and seeds seed-org plus a demo enrollment token on first start.
chmod +x scripts/full-stack-e2e.sh
./scripts/full-stack-e2e.sh
# or from infra/
./run-e2e.sh
Manual gateway request (use a user id from Developers ):
curl http://localhost:4000/v1/chat/completions \
-H " Authorization: Bearer sk-usejunction-master " \
-H " Content-Type: application/json " \
-H " x-usejunction-user: <userId> " \
-H " x-usejunction-tool: codex " \
-d ' {"model":"gpt-4o-mini","messages":[{"role":"user","content":"ping"}]} '
Service
URL
Admin UI
http://localhost:3001 ( admin@example.com / admin )
Langfuse
http://localhost:3000
LiteLLM
http://localhost:4000
Postgres (host)
localhost:5432
Hybrid local dev
cp .env.example .env
cd infra
docker compose up -d postgres langfuse-db litellm-db langfuse litellm
# Wait for DBs, then start admin locally:
cd ..
pnpm install
pnpm db:push
pnpm db:seed
pnpm dev
Admin UI: http://localhost:3001
Generate a developer-bound enrollment token after signing in and joining the organization:
curl -X POST http://localhost:3001/api/me/enrollment-token \
-H " Cookie: uj_session=... " | jq
Install the local agent
From a repo checkout (builds the Go agent locally — preferred for development):
chmod +x install.sh
./install.sh --token < token > --url http://localhost:3001
# enrolls, enables Claude OTEL, sends first report, and starts the daemon
One-liner (downloads a prebuilt binary from the control plane, or builds from source if the repo is on disk):
# optional for pnpm/dev without Docker: publish binaries into apps/admin/public
./scripts/build-agent-releases.sh 0.2.0
curl -fsSL http://localhost:3001/install.sh | sh -s -- --token < token > --url http://localhost:3001
The installer adds ~/.usejunction/bin to your shell PATH . Open a new terminal (or run export PATH="$HOME/.usejunction/bin:$PATH" ) before using usejunction commands.
Windows 10/11 PowerShell (x64 or ARM64, no administrator shell required):
powershell.exe - NoProfile - ExecutionPolicy Bypass - Command " & ([scriptblock]::Create((Invoke-RestMethod -UseBasicParsing 'http://localhost:3001/install.ps1'))) -Token '<token>' -Url 'http://localhost:3001' "
The Windows installer adds %USERPROFILE%\.usejunction\bin to your user PATH . Open a new terminal before using usejunction commands.
Teammate connect uses the shared team invite link ( /i/<token> ). After signing in there, the UI shows the install command with an enrollment --token .
The onboarding and invite screens provide a separate Windows PowerShell command. Windows installs run through a per-user Scheduled Task at logon and collect native Windows coding-tool data; WSL stores are not scanned.
cd agent && go build -o usejunction .
./usejunction enroll --token < token > --url http://localhost:3001
./usejunction doctor
./usejunction report
Two agents on one Mac (production + local dev)
UseJunction supports running production and local dev agents side by side:
Enroll production from your hosted control plane (e.g. https://usejunction.dev ).
Enroll local dev from http://localhost:3001 — the installer auto-selects the test profile for loopback URLs.
Both daemons can run at the same time without clobbering each other's enrollment.
Hot-reload the local agent (development)
After the test agent is enrolled once, rebuild and reinstall into ~/.usejunction-test whenever agent/ changes:
# admin + agent watcher (rebuilds agent on start and on agent/ changes)
pnpm dev
# or: ./scripts/dev-start.sh
# admin only (no agent rebuild/watch)
pnpm dev:admin
# one-shot rebuild + swap + daemon restart
pnpm agent:reinstall
# or: ./scripts/dev-agent-reinstall.sh
# watch agent sources and reinstall on change
pnpm dev:agent
# or: ./scripts/dev-agent-watch.sh
Requires an existing ~/.usejunction-test/config.json (from ./install.sh --token … --url http://localhost:3001 or the connect curl). This path stamps a 0.0.0-dev.<sha>.<unix> version, swaps the local binary/app bundle, and restarts launchd/systemd. It does not publish a control-plane release or enroll a new device.
Set USEJUNCTION_PROFILE=default to rebuild the production agent home ( ~/.usejunction ) instead.
When you enroll against a local control plane ( http://localhost:3001 ), /install.sh injects USEJUNCTION_ROOT and USEJUNCTION_PROFILE=test so curl | sh builds the agent from this checkout as 0.0.0-dev.* into ~/.usejunction-test instead of downloading a published release or touching production enrollment. Production hosts still serve the plain customer installer (published releases only).
Install gotchas: A prior pnpm agent:reinstall writes ~/.usejunction-test/dev-source , so later curl | sh against prod may still build 0.0.0-dev.* if a dev pin exists under the target profile home. Production customer installs also require a promoted release ( GET /api/agent-releases/latest must return 200); a GitHub agent-v* tag alone is not enough. See Install script behavior (prod vs dev) .
For faster change detection, install fswatch ( brew install fswatch ). Without it, the watcher polls every ~750ms.
┌─────────────────────────────────────────────────────────────────────────┐
│ Developer machines │
│ │
│ Codex / Claude / Cursor / Copilot / OpenCode / Antigravity / … │
│ │ │
│ ▼ │
│ Go agent (profile-isolated: ~/.usejunction or ~/.usejunction-test) │
│ • heartbeat (15m) + OTA update directives │
│ • local scans (JSONL / sqlite) → estimated_api │
│ • Cursor usage events (chargedCents) → verified_usage │
│ or rate-card estimated_api when included usage is $0 │
│ • quotas, accounts, tools inventory │
│ • optional Signals / work extraction │
│ • localhost sync endpoint (47832 default / 47833 test) │
└───────────────────────────────┬─────────────────────────────────────────┘
│ UUS sync (start → chunk → commit)
│ OTEL metrics (Claude)
│ heartbeats / agent-update events
▼
┌─────────────────────────────────────────────────────────────────────────┐
│ Control plane (apps/admin · Next.js) │
│ │
│ Ingest → UsageDaily (+ inventory / quotas) │
│ Source priority: vendor_verified > otel > device_observed > estimated │
│ Cost kinds: actual_spend · verified_usage · estimated_api │
│ │
│ Org-day snapshots → dashboard KPIs / tool detail / Models tables │
│ Sync team = wake agents to upload (does not install agent binaries) │
│ Agent OTA = tag agent-v* → promote → heartbeat directive │
│ normal = 24h staggered · critical = immediate │
└───────────────────────────────┬─────────────────────────────────────────┘
│
▼
PostgreSQL
Optional gateway path (self-hosted Docker / LiteLLM) still exists for traced proxy traffic:
Coding tools → LiteLLM → Providers
↓
Langfuse traces
↓
UseJunction callback → Admin API
Cost semantics (short): dashboard Estimated Usage = verified_usage + estimated_api . Cursor plan/bonus rows with chargedCents = 0 are estimated from the rate card after agent rematerialize — they are not labeled verified at $0. See Usage Accounting .
Reads: analytical queries go through UsageDaily , the S

[truncated]
