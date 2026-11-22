---
source: "https://github.com/relaticle/relaticle"
hn_url: "https://news.ycombinator.com/item?id=50600315"
title: "Show HN: Relaticle, an open-source CRM that Claude and ChatGPT can work in"
article_title: "GitHub - relaticle/relaticle: Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. · GitHub"
image: "https://opengraph.githubassets.com/395d9c9dc17454f794f33d191d5d8a47d87e7b0c35fa658165121cd59bc339e9/relaticle/relaticle"
author: "manukminasyan"
captured_at: "2026-11-22T15:58:25Z"
capture_tool: "hn-digest"
hn_id: 50600315
score: 2
comments: 1
posted_at: "2026-11-22T14:59:49Z"
tags:
  - hacker-news
---

# Show HN: Relaticle, an open-source CRM that Claude and ChatGPT can work in

- HN: [50600315](https://news.ycombinator.com/item?id=50600315)
- Source: [github.com](https://github.com/relaticle/relaticle)
- Score: 2
- Comments: 1
- Posted: 2026-11-22T14:59:49Z

## Translation

Title: Show HN: Relaticle, an open-source CRM that Claude and ChatGPT can work in
Article title: GitHub - relaticle/relaticle: Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. · GitHub
Description: Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. - relaticle/relaticle
HN text: I asked rela (the assistant inside my open source crm relaticle) once for a task and a note (in a test workspace). It shows me a plan, one line per change, one approve button. There was no task or note until I clicked approve. Relaticle (my open source CRM) has an MCP server for AI agents right in the core app. Claude and ChatGPT connect to it, and over the MCP protocol, they don't need a plan, they don't need an approve button, the writes just land right away. Not sure yet if those writes should wait for approval too. I started writing this in August 2024. It's written in Laravel + PostgreSQL. I'm releasing it under the AGPL-3.0 license. If you want to try it out you can run it using the compose file in the README. I'm very interested in hearing from people who let an agent write to a real system and what they have to say about it.

Article text:
GitHub - relaticle/relaticle: Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. · GitHub
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
relaticle
relaticle Public Notifications You must be signed in to change notification settings
Star 1.9k ( 1.9k ) You must be signed in to star a repository
Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
239 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
7,001 Commits 7,001 Commits Folders and files
.agents .agents .ai .ai .amp .amp .claude/ skills .claude/ skills .conductor .conductor .githooks .githooks .github .github app app bin bin bootstrap bootstrap config config database database docs docs lang lang packages packages public public resources resources routes routes storage storage stubs stubs tests tests .dockerignore .dockerignore .editorconfig .editorconfig .env.ci .env.ci .env.example .env.example .env.testing .env.testing .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GEMINI.md GEMINI.md INTEGRATION-EXCEPTION.md INTEGRATION-EXCEPTION.md LICENSE LICENSE README.md README.md artisan artisan boost.json boost.json compose.dev.yml compose.dev.yml compose.yml compose.yml composer.json composer.json composer.lock composer.lock opencode.json opencode.json package.json package.json paseo.json paseo.json phpstan-method-length.php phpstan-method-length.php phpstan.neon phpstan.neon phpunit.ci.xml phpunit.ci.xml phpunit.xml phpunit.xml pint.json pint.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml polyscope.json polyscope.json rector-tests.php rector-tests.php rector.php rector.php server.json server.json vite.config.js vite.config.js View all files Repository files navigation
The Open-Source AI CRM for Selling Without a Sales Team
Website ·
Documentation ·
MCP Server ·
Roadmap ·
Discord
One request to Rela, the built-in AI assistant: three changes, applied after you approve them.
Relaticle is a self-hosted CRM your AI agents can work in. Connect Claude, ChatGPT, or an open-source model over MCP to read, write, and analyze your CRM data. Custom fields, REST API, and multi-workspace isolation.
Built for: people who sell to a named list of companies without a sales team. Connect Claude, ChatGPT, or your own agent to work in the CRM.
Agent-Native Infrastructure - MCP server and REST API with full CRUD, schema access, activity history, and pipeline analysis
Customizable Data Model - Custom fields for text, dates, currency, selects, and links between records, with per-field encryption. No migrations needed.
Multi-Workspace Isolation - Workspace-scoped data with role-based permissions
Modern Tech Stack - Laravel 13, Filament 5, PHP 8.5, with an automated test suite on every change
Privacy-First - Self-hosted, AGPL-3.0, your data stays on your server
Relaticle ships an MCP server, so the agents you already use can work in your CRM.
Claude : add a connector with the URL https://mcp.relaticle.com (hosted) or https://your-host/mcp (self-hosted).
ChatGPT : install Relaticle from the plugin directory .
Claude Code, Cursor, Gemini CLI : plugins and the extension live in relaticle/agent-plugins .
Anything else that speaks MCP : Streamable HTTP, OAuth 2.1 or a personal access token, one workspace per connection. Setup in the MCP guide .
The fastest way to self-host. Six containers: the app, a queue worker, a scheduler, a websocket server, PostgreSQL, and Redis.
curl -o compose.yml https://raw.githubusercontent.com/Relaticle/relaticle/main/compose.yml
Generate an application key:
echo " APP_KEY=base64: $( openssl rand -base64 32 ) "
Create a .env file with your settings:
APP_KEY=base64:your-generated-key-here
DB_PASSWORD=your-secure-database-password
APP_URL=https://crm.example.com
Start the containers:
docker compose up -d
Create your admin account:
docker compose exec app php artisan make:filament-user
Your CRM is now available at {APP_URL}/app . The Self-Hosting Guide covers configuration, updates, backups, and a manual deployment. The sections below are for running from source.
Redis for queues (optional for development)
git clone https://github.com/Relaticle/relaticle.git
cd relaticle && composer app-install
Development
# Start everything (server, queue, vite)
composer dev
# Run tests
composer test
# Format code
composer lint
Self-Hosting
See the Self-Hosting Guide for Docker and manual deployment instructions.
Relaticle ships in English, German, Spanish and French. Each user picks a language under Profile, Language. See docs/i18n.md for what follows the language and how to fix a translation.
Visit our documentation for guides on business usage, technical architecture, MCP server setup, REST API integration, and more.
Star us on GitHub to support the project
Relaticle is open-source software licensed under the AGPL-3.0 license . The one exception is packages/Cloud , the code behind relaticle.com, which its own license governs. Everything else stays AGPL-3.0.
Software that talks to Relaticle only through its REST API or MCP server can use any license, including a proprietary one. See the Integration Exception .
The Relaticle name and logo are trademarks of Relaticle, Inc. If you host Relaticle for others or ship a modified version, use your own name. See the Trademark Policy .
Contributors sign a Contributor License Agreement on their first pull request.
Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
239 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. - relaticle/relaticle

I asked rela (the assistant inside my open source crm relaticle) once for a task and a note (in a test workspace). It shows me a plan, one line per change, one approve button. There was no task or note until I clicked approve. Relaticle (my open source CRM) has an MCP server for AI agents right in the core app. Claude and ChatGPT connect to it, and over the MCP protocol, they don't need a plan, they don't need an approve button, the writes just land right away. Not sure yet if those writes should wait for approval too. I started writing this in August 2024. It's written in Laravel + PostgreSQL. I'm releasing it under the AGPL-3.0 license. If you want to try it out you can run it using the compose file in the README. I'm very interested in hearing from people who let an agent write to a real system and what they have to say about it.

GitHub - relaticle/relaticle: Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud. · GitHub
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
relaticle
relaticle Public Notifications You must be signed in to change notification settings
Star 1.9k ( 1.9k ) You must be signed in to star a repository
Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
239 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
7,001 Commits 7,001 Commits Folders and files
.agents .agents .ai .ai .amp .amp .claude/ skills .claude/ skills .conductor .conductor .githooks .githooks .github .github app app bin bin bootstrap bootstrap config config database database docs docs lang lang packages packages public public resources resources routes routes storage storage stubs stubs tests tests .dockerignore .dockerignore .editorconfig .editorconfig .env.ci .env.ci .env.example .env.example .env.testing .env.testing .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GEMINI.md GEMINI.md INTEGRATION-EXCEPTION.md INTEGRATION-EXCEPTION.md LICENSE LICENSE README.md README.md artisan artisan boost.json boost.json compose.dev.yml compose.dev.yml compose.yml compose.yml composer.json composer.json composer.lock composer.lock opencode.json opencode.json package.json package.json paseo.json paseo.json phpstan-method-length.php phpstan-method-length.php phpstan.neon phpstan.neon phpunit.ci.xml phpunit.ci.xml phpunit.xml phpunit.xml pint.json pint.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml polyscope.json polyscope.json rector-tests.php rector-tests.php rector.php rector.php server.json server.json vite.config.js vite.config.js View all files Repository files navigation
The Open-Source AI CRM for Selling Without a Sales Team
Website ·
Documentation ·
MCP Server ·
Roadmap ·
Discord
One request to Rela, the built-in AI assistant: three changes, applied after you approve them.
Relaticle is a self-hosted CRM your AI agents can work in. Connect Claude, ChatGPT, or an open-source model over MCP to read, write, and analyze your CRM data. Custom fields, REST API, and multi-workspace isolation.
Built for: people who sell to a named list of companies without a sales team. Connect Claude, ChatGPT, or your own agent to work in the CRM.
Agent-Native Infrastructure - MCP server and REST API with full CRUD, schema access, activity history, and pipeline analysis
Customizable Data Model - Custom fields for text, dates, currency, selects, and links between records, with per-field encryption. No migrations needed.
Multi-Workspace Isolation - Workspace-scoped data with role-based permissions
Modern Tech Stack - Laravel 13, Filament 5, PHP 8.5, with an automated test suite on every change
Privacy-First - Self-hosted, AGPL-3.0, your data stays on your server
Relaticle ships an MCP server, so the agents you already use can work in your CRM.
Claude : add a connector with the URL https://mcp.relaticle.com (hosted) or https://your-host/mcp (self-hosted).
ChatGPT : install Relaticle from the plugin directory .
Claude Code, Cursor, Gemini CLI : plugins and the extension live in relaticle/agent-plugins .
Anything else that speaks MCP : Streamable HTTP, OAuth 2.1 or a personal access token, one workspace per connection. Setup in the MCP guide .
The fastest way to self-host. Six containers: the app, a queue worker, a scheduler, a websocket server, PostgreSQL, and Redis.
curl -o compose.yml https://raw.githubusercontent.com/Relaticle/relaticle/main/compose.yml
Generate an application key:
echo " APP_KEY=base64: $( openssl rand -base64 32 ) "
Create a .env file with your settings:
APP_KEY=base64:your-generated-key-here
DB_PASSWORD=your-secure-database-password
APP_URL=https://crm.example.com
Start the containers:
docker compose up -d
Create your admin account:
docker compose exec app php artisan make:filament-user
Your CRM is now available at {APP_URL}/app . The Self-Hosting Guide covers configuration, updates, backups, and a manual deployment. The sections below are for running from source.
Redis for queues (optional for development)
git clone https://github.com/Relaticle/relaticle.git
cd relaticle && composer app-install
Development
# Start everything (server, queue, vite)
composer dev
# Run tests
composer test
# Format code
composer lint
Self-Hosting
See the Self-Hosting Guide for Docker and manual deployment instructions.
Relaticle ships in English, German, Spanish and French. Each user picks a language under Profile, Language. See docs/i18n.md for what follows the language and how to fix a translation.
Visit our documentation for guides on business usage, technical architecture, MCP server setup, REST API integration, and more.
Star us on GitHub to support the project
Relaticle is open-source software licensed under the AGPL-3.0 license . The one exception is packages/Cloud , the code behind relaticle.com, which its own license governs. Everything else stays AGPL-3.0.
Software that talks to Relaticle only through its REST API or MCP server can use any license, including a proprietary one. See the Integration Exception .
The Relaticle name and logo are trademarks of Relaticle, Inc. If you host Relaticle for others or ship a modified version, use your own name. See the Trademark Policy .
Contributors sign a Contributor License Agreement on their first pull request.
Open-source CRM your AI agents can read and update: companies, people, opportunities, tasks and notes. REST API, self-hosted or cloud.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
239 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
