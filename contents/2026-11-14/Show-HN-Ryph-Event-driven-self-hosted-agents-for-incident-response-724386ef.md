---
source: "https://github.com/ryph-hq/ryph"
hn_url: "https://news.ycombinator.com/item?id=50497125"
title: "Show HN: Ryph – Event-driven, self-hosted agents for incident response"
article_title: "GitHub - ryph-hq/ryph: 🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. · GitHub"
image: "https://opengraph.githubassets.com/7437414fb57845024035fe039044950267dc5517bec2f7191ddc70a5ce5259f8/ryph-hq/ryph"
author: "lftherios"
captured_at: "2026-11-14T14:48:48Z"
capture_tool: "hn-digest"
hn_id: 50497125
score: 3
comments: 0
posted_at: "2026-11-14T14:05:05Z"
tags:
  - hacker-news
---

# Show HN: Ryph – Event-driven, self-hosted agents for incident response

- HN: [50497125](https://news.ycombinator.com/item?id=50497125)
- Source: [github.com](https://github.com/ryph-hq/ryph)
- Score: 3
- Comments: 0
- Posted: 2026-11-14T14:05:05Z

## Translation

Title: Show HN: Ryph – Event-driven, self-hosted agents for incident response
Article title: GitHub - ryph-hq/ryph: 🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. · GitHub
Description: 🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. - ryph-hq/ryph

Article text:
GitHub - ryph-hq/ryph: 🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. · GitHub
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
ryph-hq
ryph Public Notifications You must be signed in to change notification settings
Star 17 ( 17 ) You must be signed in to star a repository
🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.ryph .ryph deploy deploy docs docs images images packages packages perf perf scripts scripts types types .dockerignore .dockerignore .gitignore .gitignore .oxlintrc.json .oxlintrc.json CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml dev dev dprint.json dprint.json mise.toml mise.toml package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json View all files Repository files navigation
Homepage ·
Docs ·
Getting started
Ryph is under active and heavy development. Breaking changes are possible before v1; see
Releases & versioning .
Ryph is a complete, batteries-included agent orchestrator for teams that's easy to deploy and put
into action straight away.
Deploy an incident response bot
An alert fires on your IRM, and an automation within Ryph
immediately spins up a secure sandbox for your responder bot within your infra. The sandbox
ships with a warm build cache and dependencies pre-installed ,
so your bot can get to work right away.
The bot starts reading telemetry directly from your observability stack,
securely through the Tool Proxy , and diagnoses the problem. It
remembers details from past incidents that help it pinpoint
the root cause.
It joins your teams existing incident chat , and reports what it finds.
People in this chat can mention it to ask follow-up questions or steer it.
Once the root cause is clear, the bot builds a fix, runs your repo's checks, submits a hotfix pull
request that opens with all CI checks already green , and sends the
link to your chat. Ryph's forge safeguards
prevent the bot from accidentally making a bad push, or even sidestepping a code review.
An engineer reviews the hotfix, merges and deploys. The bot is automatically notified of the PR
being merged, and starts watching telemetry for the deploy landing.
Eventually, the alert resolves, the bot reports all green, and closes out.
Give your teammates a secure helper bot that can reason from code, and ship fixes
With the Chat integration , allow your teammates to ask a self-hosted
agent right from where your team already communicates - Slack, Discord, Telegram, or Zulip. For each
new agent run, Ryph provisions a secure sandbox,
with dependencies already pre-installed .
The bot gets specific repos checked out, can reason from their code, and answer questions. If you
want, give it access to specific tables and fields from your database through the
Tool Proxy , so it can perform data analysis for non-technical
teammates. For long-form reports, the bot creates an Artifact and
posts the link back in the chat - all self-hosted on your own infra.
When asked to fix a visual glitch in the frontend, the bot can be steered by screenshots people post
in the chat. It goes to fix the glitch, verifies it in its
in-sandbox browser , ships a fix PR for an engineer to review, and
reports back in the chat with a screenshot of the fixed UI. Ryph's
forge safeguards prevent the bot from
accidentally making a bad push, or even sidestepping a code review.
Ryph is the self-hosted platform that does all this and more: coding agents that run in secure
containers, started by an alert, a task in the UI, a /ryph comment on a pull request, a mention in
chat or a schedule, whose work leaves as a pull request, with every step on record.
Ryph ships as a single binary for trying it out on your laptop,
in Docker Compose for teams, or as a
Helm Chart for teams that require autoscaling with Kubernetes.
For smaller teams and sporadic agent use, Docker Compose is
the quickest way to get started.
For heavier use, especially when also using Ryph Checks , the
Kubernetes deployment mode unlocks automatic resource scaling.
New agent sandboxes and check pods are placed within your cluster by k8s, and in combination with
autoscaling solutions such as
GKE Autopilot ,
the cluster automatically provisions and removes nodes according to short-term resource needs.
Read more about deployment options in our docs
With your coding agent. Tell Claude Code, Codex or Cursor: read
https://raw.githubusercontent.com/ryph-hq/ryph/main/docs/guides/setup-with-an-agent.md and set
me up. It installs the binary, answers the setup questions with you, starts it, and hands you the
two steps only you can do.
By hand. docs/guides/quickstart.md : install, setup, first run,
first automation.
For a team. deploy/README.md compares the binary, Compose and the Helm
chart, and says which to pick.
ryph calls a handled alert a response : a run that stays open, bound to one conversation, while
later alerts about the same thing join it as messages.
sequenceDiagram
participant S as Alert source
participant A as Automation
participant R as Run (container)
participant C as Incident chat
participant E as Engineers
participant F as Forge
S->>A: alert webhook (signed)
Note over A: verify · read status and key<br/>open a response, or join the open one
A->>R: start the run: recipe + alert as untrusted data
A->>C: opening line: the alert, the agent, the run link
R->>C: findings, as the agent works
E->>C: "@bot check the connection pool first"
C->>R: the message lands in the run's inbox
R->>F: patch (git bundle) → branch → pull request
F-->>C: PR link, once it is open
S->>A: resolved
A->>R: joins the response as a message
R->>C: closing line: outcome, pull requests, run link
Loading
Automations take the webhook, verify it, and read what it says: Grafana, Alertmanager, Grafana
IRM, PagerDuty, Sentry, incident.io, Datadog, Opsgenie, or your own mapping
( @ryph/automations ).
Chat is where the response lives: a Slack thread, a Zulip topic or a Telegram group, and the
people in it steer the run by mentioning the bot ( @ryph/chat ).
The run is one container with one agent, holding only credentials minted for it; it reaches
models and tools through ryph's proxies and never sees a real key
( @ryph/runner ,
@ryph/model-proxy ,
@ryph/mcp-gateway ).
Memory gives it the runbooks and gotchas earlier runs left behind
( @ryph/memory ).
The forge turns a patch into a branch and a pull request, after a person approves it in ryph
or at once where a repo allows that ( @ryph/forge ).
The guide tells the whole story on one stack and gives the checklist for yours:
docs/guides/incident-response.md .
Alerting
Grafana, Alertmanager, Grafana IRM, PagerDuty, Sentry, incident.io, Datadog, Opsgenie, or a custom mapping
Chat
Slack, Zulip, Telegram
Forge
GitHub (as you, or as an App), GitLab, Radicle, or a directory on this machine
Models
Anthropic, OpenAI, a ChatGPT plan, or any OpenAI-compatible endpoint
Tools for agents
any MCP server, hosted by ryph or remote; Grafana, Sentry, Datadog and PagerDuty are in the catalogue
The checklist that walks a team through its own combination, with a coding agent doing the parts it
can:
fit it into your operational context .
Tasks from the UI , on one repo or a workspace of several.
/ryph on a pull request, merge request or issue starts a run on it
( forge ).
A mention in chat starts a run in that thread.
Schedules : a cron expression as a trigger, same recipe.
Previews of a run's dev server in your browser ( previews ).
Checks a repo declares in .ryph/checks.yml , run on ryph and reported to the forge
( checks ).
Skills and MCP servers for agents, with a policy per agent
( skills , mcp-gateway ).
ryph mcp : the deployment as MCP tools for the agent on your machine
( the binary's README ).
docs/system-flow.md has the cast of players, the happy path in one picture,
and the walk through every step with the doc that owns it. The workspace, one package per role:
images/ builds one image per service plus the agent images.
deploy/ holds Compose and the Helm chart.
CONTRIBUTING.md : the toolchain, the commands, how the repo runs on itself, and
how designs are written down.
🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. - ryph-hq/ryph

GitHub - ryph-hq/ryph: 🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge. · GitHub
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
ryph-hq
ryph Public Notifications You must be signed in to change notification settings
Star 17 ( 17 ) You must be signed in to star a repository
🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.ryph .ryph deploy deploy docs docs images images packages packages perf perf scripts scripts types types .dockerignore .dockerignore .gitignore .gitignore .oxlintrc.json .oxlintrc.json CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml dev dev dprint.json dprint.json mise.toml mise.toml package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json View all files Repository files navigation
Homepage ·
Docs ·
Getting started
Ryph is under active and heavy development. Breaking changes are possible before v1; see
Releases & versioning .
Ryph is a complete, batteries-included agent orchestrator for teams that's easy to deploy and put
into action straight away.
Deploy an incident response bot
An alert fires on your IRM, and an automation within Ryph
immediately spins up a secure sandbox for your responder bot within your infra. The sandbox
ships with a warm build cache and dependencies pre-installed ,
so your bot can get to work right away.
The bot starts reading telemetry directly from your observability stack,
securely through the Tool Proxy , and diagnoses the problem. It
remembers details from past incidents that help it pinpoint
the root cause.
It joins your teams existing incident chat , and reports what it finds.
People in this chat can mention it to ask follow-up questions or steer it.
Once the root cause is clear, the bot builds a fix, runs your repo's checks, submits a hotfix pull
request that opens with all CI checks already green , and sends the
link to your chat. Ryph's forge safeguards
prevent the bot from accidentally making a bad push, or even sidestepping a code review.
An engineer reviews the hotfix, merges and deploys. The bot is automatically notified of the PR
being merged, and starts watching telemetry for the deploy landing.
Eventually, the alert resolves, the bot reports all green, and closes out.
Give your teammates a secure helper bot that can reason from code, and ship fixes
With the Chat integration , allow your teammates to ask a self-hosted
agent right from where your team already communicates - Slack, Discord, Telegram, or Zulip. For each
new agent run, Ryph provisions a secure sandbox,
with dependencies already pre-installed .
The bot gets specific repos checked out, can reason from their code, and answer questions. If you
want, give it access to specific tables and fields from your database through the
Tool Proxy , so it can perform data analysis for non-technical
teammates. For long-form reports, the bot creates an Artifact and
posts the link back in the chat - all self-hosted on your own infra.
When asked to fix a visual glitch in the frontend, the bot can be steered by screenshots people post
in the chat. It goes to fix the glitch, verifies it in its
in-sandbox browser , ships a fix PR for an engineer to review, and
reports back in the chat with a screenshot of the fixed UI. Ryph's
forge safeguards prevent the bot from
accidentally making a bad push, or even sidestepping a code review.
Ryph is the self-hosted platform that does all this and more: coding agents that run in secure
containers, started by an alert, a task in the UI, a /ryph comment on a pull request, a mention in
chat or a schedule, whose work leaves as a pull request, with every step on record.
Ryph ships as a single binary for trying it out on your laptop,
in Docker Compose for teams, or as a
Helm Chart for teams that require autoscaling with Kubernetes.
For smaller teams and sporadic agent use, Docker Compose is
the quickest way to get started.
For heavier use, especially when also using Ryph Checks , the
Kubernetes deployment mode unlocks automatic resource scaling.
New agent sandboxes and check pods are placed within your cluster by k8s, and in combination with
autoscaling solutions such as
GKE Autopilot ,
the cluster automatically provisions and removes nodes according to short-term resource needs.
Read more about deployment options in our docs
With your coding agent. Tell Claude Code, Codex or Cursor: read
https://raw.githubusercontent.com/ryph-hq/ryph/main/docs/guides/setup-with-an-agent.md and set
me up. It installs the binary, answers the setup questions with you, starts it, and hands you the
two steps only you can do.
By hand. docs/guides/quickstart.md : install, setup, first run,
first automation.
For a team. deploy/README.md compares the binary, Compose and the Helm
chart, and says which to pick.
ryph calls a handled alert a response : a run that stays open, bound to one conversation, while
later alerts about the same thing join it as messages.
sequenceDiagram
participant S as Alert source
participant A as Automation
participant R as Run (container)
participant C as Incident chat
participant E as Engineers
participant F as Forge
S->>A: alert webhook (signed)
Note over A: verify · read status and key<br/>open a response, or join the open one
A->>R: start the run: recipe + alert as untrusted data
A->>C: opening line: the alert, the agent, the run link
R->>C: findings, as the agent works
E->>C: "@bot check the connection pool first"
C->>R: the message lands in the run's inbox
R->>F: patch (git bundle) → branch → pull request
F-->>C: PR link, once it is open
S->>A: resolved
A->>R: joins the response as a message
R->>C: closing line: outcome, pull requests, run link
Loading
Automations take the webhook, verify it, and read what it says: Grafana, Alertmanager, Grafana
IRM, PagerDuty, Sentry, incident.io, Datadog, Opsgenie, or your own mapping
( @ryph/automations ).
Chat is where the response lives: a Slack thread, a Zulip topic or a Telegram group, and the
people in it steer the run by mentioning the bot ( @ryph/chat ).
The run is one container with one agent, holding only credentials minted for it; it reaches
models and tools through ryph's proxies and never sees a real key
( @ryph/runner ,
@ryph/model-proxy ,
@ryph/mcp-gateway ).
Memory gives it the runbooks and gotchas earlier runs left behind
( @ryph/memory ).
The forge turns a patch into a branch and a pull request, after a person approves it in ryph
or at once where a repo allows that ( @ryph/forge ).
The guide tells the whole story on one stack and gives the checklist for yours:
docs/guides/incident-response.md .
Alerting
Grafana, Alertmanager, Grafana IRM, PagerDuty, Sentry, incident.io, Datadog, Opsgenie, or a custom mapping
Chat
Slack, Zulip, Telegram
Forge
GitHub (as you, or as an App), GitLab, Radicle, or a directory on this machine
Models
Anthropic, OpenAI, a ChatGPT plan, or any OpenAI-compatible endpoint
Tools for agents
any MCP server, hosted by ryph or remote; Grafana, Sentry, Datadog and PagerDuty are in the catalogue
The checklist that walks a team through its own combination, with a coding agent doing the parts it
can:
fit it into your operational context .
Tasks from the UI , on one repo or a workspace of several.
/ryph on a pull request, merge request or issue starts a run on it
( forge ).
A mention in chat starts a run in that thread.
Schedules : a cron expression as a trigger, same recipe.
Previews of a run's dev server in your browser ( previews ).
Checks a repo declares in .ryph/checks.yml , run on ryph and reported to the forge
( checks ).
Skills and MCP servers for agents, with a policy per agent
( skills , mcp-gateway ).
ryph mcp : the deployment as MCP tools for the agent on your machine
( the binary's README ).
docs/system-flow.md has the cast of players, the happy path in one picture,
and the walk through every step with the doc that owns it. The workspace, one package per role:
images/ builds one image per service plus the agent images.
deploy/ holds Compose and the Helm chart.
CONTRIBUTING.md : the toolchain, the commands, how the repo runs on itself, and
how designs are written down.
🤖 Self-hosted coding agents for your team: on call for your incidents, in your chat, on your forge.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
