---
source: "https://github.com/ameesme/shoebox"
hn_url: "https://news.ycombinator.com/item?id=50610136"
title: "Show HN: Crazy simple AI agent orchestration with Markdown and Containers"
article_title: "GitHub - ameesme/shoebox: Crazy simple filesystem-based AI agent orchestration. · GitHub"
image: "https://opengraph.githubassets.com/749635311da308ad760679f1db817fdd9089533d72f7de0c7221845fb80f8ab7/ameesme/shoebox"
author: "ameesdotme"
captured_at: "2026-11-23T11:53:39Z"
capture_tool: "hn-digest"
hn_id: 50610136
score: 2
comments: 0
posted_at: "2026-11-23T11:07:29Z"
tags:
  - hacker-news
---

# Show HN: Crazy simple AI agent orchestration with Markdown and Containers

- HN: [50610136](https://news.ycombinator.com/item?id=50610136)
- Source: [github.com](https://github.com/ameesme/shoebox)
- Score: 2
- Comments: 0
- Posted: 2026-11-23T11:07:29Z

## Translation

Title: Show HN: Crazy simple AI agent orchestration with Markdown and Containers
Article title: GitHub - ameesme/shoebox: Crazy simple filesystem-based AI agent orchestration. · GitHub
Description: Crazy simple filesystem-based AI agent orchestration. - ameesme/shoebox
HN text: I was sick and tired of the bloated mess that Hermes Agent has become, often dealing with hard-to-triage bugs that made managing agents a part-time job. I don't need a self-improving magical agent, I need; - A primary agent to chat to, with sub-agent delegation - Isolated sub-agent workspaces, tools, binaries and skills - Kanban and cron-tasks to manage tasks - An easy way to chat (Telegram, Web UI) I built shoebox, crazy simple filesystem-based AI agent orchestration that uses a basic folder-structure to determine which agents to run, what they should do, and what capabilities they have. Been running this on prod-level workloads (with extended product-development runs) and have been very happy with the results. Curious to hear what you think!

Article text:
GitHub - ameesme/shoebox: Crazy simple filesystem-based AI agent orchestration. · GitHub
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
ameesme
shoebox Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Crazy simple filesystem-based AI agent orchestration.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
63 Commits 63 Commits Folders and files
assets assets channels channels data data docker docker runtime runtime .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.web.yaml compose.web.yaml compose.yaml compose.yaml View all files Repository files navigation
Crazy simple filesystem-based AI agent orchestration.
A main chat-agent that delegates work and reports back results
An isolated Docker container and image per agent with a filesystem for
Kanban
Instructions like Soul.md and skills
Automatic git versioning for agent definitions, skills, schedules and memory
Terminal UI, and optional web and Telegram channels
Main restarts shoebox to apply changes to agents, after you confirm
Each agent is a folder with its own instructions, tools and Docker image, where you can define it's skills, soul and tools (by adding them to the Dockerfile). You can chat to each agent, but main can easily create and delegate tasks.
Tasks are Markdown files in that agent's kanban/ folder. Moving a task between todo , in-progress , done and blocked changes the task's status.
shoebox starts agent containers automatically as tasks arrive. When a task finishes or gets blocked, main is notified and reports back to you.
From git-clone to agent-delegation in under a minute.
Install Docker and clone this repo
Copy data/shared/.env.example to data/shared/.env
Add your provider's API key, e.g. OPENAI_API_KEY=your-api-key .
See Pi's provider setup for other providers
Run docker compose up -d --build
Run docker compose exec agents chat to open Pi's native CLI
poem is included as an agent by default. As its name suggests, it writes poems, including a cron-based daily poem.
Ask main: Have poem write a short poem about a quiet morning. Main will dispatch the task and reports back when it finishes. You can keep chatting while poem works.
Ask main to create an agent, or add a folder under data/agents/ yourself. An agent needs;
The agent's role and instructions in Markdown; Pi uses it as the system prompt.
agent.json is the configuration for Pi, defining the agent's tools, skills, MCP servers and optional model (which wins over PI_MODEL in .env ). Add skills under skills/ as needed.
MCP servers go under mcpServers , as in other MCP clients; ${NAME} takes a
value from the agent's .env .
Use FROM agent-base for Pi, Python and the task runtime. Add any other software your agent needs, for example, the included code agent can add Claude Code, tmux and Chromium by uncommenting them.
When you start shoebox, it creates each agent's workspace, memory and task folders, and builds the agent's image. It rebuilds the image when its Dockerfile or the base image definition changes.
Define recurring tasks as Markdown files in each worker's cron/ folder. Ask
main to set up a schedule, or copy poem's example .
Under data/shared , you and agents can share information with each-other that is not a task. For example;
.env : shared credentials; overidden by an agent's own .env
skills/ : shared skills; overidden by an agent's own skills if names match
Shared files, like artifacts as a result of a task
Running chats pick up changes to credentials and skills after a restart, which main asks you to confirm first. Changes to the root .env or shoebox itself need docker compose up -d --build .
The default way to interact with main is through Pi's terminal UI. However, by setting a custom CHANNEL_COMMAND , you can enable other ways.
To enable Telegram, update the root .env by setting
CHANNEL_COMMAND=sh /channels/telegram.sh
TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_BOT_TOKEN
and running docker compose up -d .
You can only chat with main directly in Telegram. Use /new or /reset starts a fresh session. Speech-to-text is enabled by including an OPENAI_API_KEY in your .env.
To enable shoebox's basic htmx web-interface, update the root .env setting
CHANNEL_COMMAND=sh /channels/web.sh
COMPOSE_FILE=compose.yaml:compose.web.yaml
and running docker compose up -d . You can now open http://localhost:8080 .
To change the port, you can set WEB_PORT . To combine with Telegram, just concatenate its channel command:
CHANNEL_COMMAND=sh /channels/web.sh sh /channels/telegram.sh
In the web-interface, you can chat with any agent. Other channels' conversations, such as Telegram, are shown read-only. Files browses agent and shared files, with upload, download and delete, including .env files.
shoebox does a lot with very little code. Docker, git and Pi do the heavy lifting, so there's no database, gateway or agent loop to maintain. That's about 1.4k lines, idling at under 100 MB. If you want an assistant in every chat app, look at OpenClaw or Hermes Agent . For a full workspace with email, documents and local models, look at Odysseus .
— means not built in or not documented, as of November 2026.
¹ Non-blank source lines from each repo, without tests, generated code or dependencies. shoebox builds on Pi, which isn't counted.
² Each project's default Docker setup, with no channels or models configured, after 3 minutes idle on Docker Desktop (arm64). Memory is from docker stats , disk is the images it pulls or builds.
Crazy simple filesystem-based AI agent orchestration.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Crazy simple filesystem-based AI agent orchestration. - ameesme/shoebox

I was sick and tired of the bloated mess that Hermes Agent has become, often dealing with hard-to-triage bugs that made managing agents a part-time job. I don't need a self-improving magical agent, I need; - A primary agent to chat to, with sub-agent delegation - Isolated sub-agent workspaces, tools, binaries and skills - Kanban and cron-tasks to manage tasks - An easy way to chat (Telegram, Web UI) I built shoebox, crazy simple filesystem-based AI agent orchestration that uses a basic folder-structure to determine which agents to run, what they should do, and what capabilities they have. Been running this on prod-level workloads (with extended product-development runs) and have been very happy with the results. Curious to hear what you think!

GitHub - ameesme/shoebox: Crazy simple filesystem-based AI agent orchestration. · GitHub
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
ameesme
shoebox Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Crazy simple filesystem-based AI agent orchestration.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
63 Commits 63 Commits Folders and files
assets assets channels channels data data docker docker runtime runtime .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.web.yaml compose.web.yaml compose.yaml compose.yaml View all files Repository files navigation
Crazy simple filesystem-based AI agent orchestration.
A main chat-agent that delegates work and reports back results
An isolated Docker container and image per agent with a filesystem for
Kanban
Instructions like Soul.md and skills
Automatic git versioning for agent definitions, skills, schedules and memory
Terminal UI, and optional web and Telegram channels
Main restarts shoebox to apply changes to agents, after you confirm
Each agent is a folder with its own instructions, tools and Docker image, where you can define it's skills, soul and tools (by adding them to the Dockerfile). You can chat to each agent, but main can easily create and delegate tasks.
Tasks are Markdown files in that agent's kanban/ folder. Moving a task between todo , in-progress , done and blocked changes the task's status.
shoebox starts agent containers automatically as tasks arrive. When a task finishes or gets blocked, main is notified and reports back to you.
From git-clone to agent-delegation in under a minute.
Install Docker and clone this repo
Copy data/shared/.env.example to data/shared/.env
Add your provider's API key, e.g. OPENAI_API_KEY=your-api-key .
See Pi's provider setup for other providers
Run docker compose up -d --build
Run docker compose exec agents chat to open Pi's native CLI
poem is included as an agent by default. As its name suggests, it writes poems, including a cron-based daily poem.
Ask main: Have poem write a short poem about a quiet morning. Main will dispatch the task and reports back when it finishes. You can keep chatting while poem works.
Ask main to create an agent, or add a folder under data/agents/ yourself. An agent needs;
The agent's role and instructions in Markdown; Pi uses it as the system prompt.
agent.json is the configuration for Pi, defining the agent's tools, skills, MCP servers and optional model (which wins over PI_MODEL in .env ). Add skills under skills/ as needed.
MCP servers go under mcpServers , as in other MCP clients; ${NAME} takes a
value from the agent's .env .
Use FROM agent-base for Pi, Python and the task runtime. Add any other software your agent needs, for example, the included code agent can add Claude Code, tmux and Chromium by uncommenting them.
When you start shoebox, it creates each agent's workspace, memory and task folders, and builds the agent's image. It rebuilds the image when its Dockerfile or the base image definition changes.
Define recurring tasks as Markdown files in each worker's cron/ folder. Ask
main to set up a schedule, or copy poem's example .
Under data/shared , you and agents can share information with each-other that is not a task. For example;
.env : shared credentials; overidden by an agent's own .env
skills/ : shared skills; overidden by an agent's own skills if names match
Shared files, like artifacts as a result of a task
Running chats pick up changes to credentials and skills after a restart, which main asks you to confirm first. Changes to the root .env or shoebox itself need docker compose up -d --build .
The default way to interact with main is through Pi's terminal UI. However, by setting a custom CHANNEL_COMMAND , you can enable other ways.
To enable Telegram, update the root .env by setting
CHANNEL_COMMAND=sh /channels/telegram.sh
TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_BOT_TOKEN
and running docker compose up -d .
You can only chat with main directly in Telegram. Use /new or /reset starts a fresh session. Speech-to-text is enabled by including an OPENAI_API_KEY in your .env.
To enable shoebox's basic htmx web-interface, update the root .env setting
CHANNEL_COMMAND=sh /channels/web.sh
COMPOSE_FILE=compose.yaml:compose.web.yaml
and running docker compose up -d . You can now open http://localhost:8080 .
To change the port, you can set WEB_PORT . To combine with Telegram, just concatenate its channel command:
CHANNEL_COMMAND=sh /channels/web.sh sh /channels/telegram.sh
In the web-interface, you can chat with any agent. Other channels' conversations, such as Telegram, are shown read-only. Files browses agent and shared files, with upload, download and delete, including .env files.
shoebox does a lot with very little code. Docker, git and Pi do the heavy lifting, so there's no database, gateway or agent loop to maintain. That's about 1.4k lines, idling at under 100 MB. If you want an assistant in every chat app, look at OpenClaw or Hermes Agent . For a full workspace with email, documents and local models, look at Odysseus .
— means not built in or not documented, as of November 2026.
¹ Non-blank source lines from each repo, without tests, generated code or dependencies. shoebox builds on Pi, which isn't counted.
² Each project's default Docker setup, with no channels or models configured, after 3 minutes idle on Docker Desktop (arm64). Memory is from docker stats , disk is the images it pulls or builds.
Crazy simple filesystem-based AI agent orchestration.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
