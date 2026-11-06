---
source: "https://github.com/zalando-incubator/agentic-identity-broker"
hn_url: "https://news.ycombinator.com/item?id=50393235"
title: "Show HN: Agentic Identity Broker, OAuth Delegation for AI Agents"
article_title: "GitHub - zalando-incubator/agentic-identity-broker: Identity broker for AI agents with OAuth2 delegation and consent management · GitHub"
image: "https://repository-images.githubusercontent.com/1295060661/bc15dfb3-27b6-410b-ab72-cdbe13fe9c03"
author: "jbspeakr"
captured_at: "2026-11-06T15:42:32Z"
capture_tool: "hn-digest"
hn_id: 50393235
score: 2
comments: 0
posted_at: "2026-11-06T15:07:47Z"
tags:
  - hacker-news
---

# Show HN: Agentic Identity Broker, OAuth Delegation for AI Agents

- HN: [50393235](https://news.ycombinator.com/item?id=50393235)
- Source: [github.com](https://github.com/zalando-incubator/agentic-identity-broker)
- Score: 2
- Comments: 0
- Posted: 2026-11-06T15:07:47Z

## Translation

Title: Show HN: Agentic Identity Broker, OAuth Delegation for AI Agents
Article title: GitHub - zalando-incubator/agentic-identity-broker: Identity broker for AI agents with OAuth2 delegation and consent management · GitHub
Description: Identity broker for AI agents with OAuth2 delegation and consent management - zalando-incubator/agentic-identity-broker
HN text: Hi HN, I’m Jan. I am co-author of the Agentic Identity Broker (AIB), an MIT-licensed OAuth2 broker in the Zalando Incubator, that lets a user give an AI agent scoped, revocable access to MCP services. The broker solves three major challenges that are especially relevant in corporate environments: (A) Enable users to delegate workloads to agents without providing the agents all the user’s tokens. (B) Connect fragmented identity and authorization systems. (C) Centralize tool approvals and audit logging. We built AIB to fill an infrastructure gap in our own agent platform. At this moment in time, IETF working groups are still finalizing standards for agent identity, and vendors are still turning the necessary infrastructure into off-the-shelf products. We needed that infrastructure today. How it works: the user consents once per agent and service. The broker stores the third-party tokens encrypted. At request time a gateway (e.g. Agentgateway) does RFC 8693 token exchange. Agents never see user tokens and MCP servers only implement tools and don’t need to care about all the OAuth 2.1 plumbing. AIB is v0.1 and runs in production, bringing together hundreds of our agents with dozens of internal and external MCP servers. To try it locally: clone the repo, run "just compose-up-detached". The current features focus on on-behalf-of (OBO) use cases, where human actors delegate to agents. Next we are looking to also support headless agents, where agents are invoked by machines. Intent-based authorization is yet another package on our roadmap. If you’ve built either, or have feedback on the UI/UX, I’d like to hear from you!

Article text:
GitHub - zalando-incubator/agentic-identity-broker: Identity broker for AI agents with OAuth2 delegation and consent management · GitHub
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
zalando-incubator
agentic-identity-broker Public Notifications You must be signed in to change notification settings
Star 34 ( 34 ) You must be signed in to star a repository
Identity broker for AI agents with OAuth2 delegation and consent management
agenticidentitybroker.dev Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
262 Commits 262 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills .githooks .githooks .github .github .specify .specify .vscode .vscode adrs adrs api api assets/ docusaurus assets/ docusaurus build build charts/ agentic-identity-broker charts/ agentic-identity-broker cmd cmd coding-standards coding-standards configs configs docs docs examples/ config examples/ config infra infra internal internal migrations migrations mocks mocks scripts scripts specs specs tests tests tools tools web web .air.toml .air.toml .dockerignore .dockerignore .env.compose.example .env.compose.example .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .zappr.yaml .zappr.yaml AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS MAINTAINERS README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md VERSION VERSION config.yaml config.yaml delivery.yaml delivery.yaml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum justfile justfile osv-scanner.toml osv-scanner.toml View all files Repository files navigation
Delegate scoped, revocable access to AI agents — without handing them your credentials.
Agentic Identity Broker manages delegation chains for on-behalf-of flows in Agentic AI, bridges distinct OAuth2 infrastructures, and provides a secure token vault. Deployed within the call path via gateways like Agentgateway , it offloads auth burdens from both agent and MCP server creators.
For users, it provides a simple way to manage consent, granting agents granular permissions across services such as Google, GitHub, Databricks, and Linear without compromising control.
Secure Identity Management
Agent registration, principal authentication, and session management
OAuth2 Delegation & Consent
Manage OAuth2 scopes and user consent for third-party service integrations
Security-First Design
Sensitive value redaction, command injection prevention, cryptographic validation, and fail-closed architecture
Rich Consent Experience
Modern UX for managing OAuth2 delegations with design system compliance
Comprehensive Audit Logging
Structured JSON logging for compliance and security monitoring
Frameworks like FastMCP can implement the whole OAuth2.1 ceremonies as mandated by the MCP spec but this means that you have to configure it for every MCP Server you deploy and maintain it. You also have to configure secure storage and basically run a multitude of OAuth2 authorization servers. With the identity broker and a gateway you can dumb down MCP server development to only provide tools on the target technology you want. OAuth2 ceremonies, token vaulting and token translation are done transparently. Because token exchange is done centrally via a gateway, this allows agents to interact with hundreds of MCP servers via one channel.
If you want to offer an MCP server to your customers as part of your SaaS, the identity broker can provide an additional consent surface that records which agents were actually used. With centralized tool authorization, you can keep an audit trail of approvals on your side regardless of which agent the customer runs.
The agentic identity broker can simplify agents by requiring only a single user token that can be used with an arbitrary number of MCP servers. When paired with a portal for user interactions, token procurement can be front-loaded in the portal, and the agent's responsibility is to forward this token to upstream MCP servers via a gateway.
Delegation Chains Without User Presence
The impersonation feature allows to start delegation chains via privileged clients, like a Google Chat or Slack bridge. The broker can be instructed to trust for two independent machine identities to assume the session for a particular user. This allows to use the permissions of a user on the company chat in another third-party system.
The tool approval support in the broker allows to enforce tool calls independent of the agent interface in use. This allows a last resort check before emails get sent or documents shared with a central audit trail.
Quickstart — Get started with Agentic Identity Broker in minutes.
agenticidentitybroker.dev – Underlying concepts, guides and API specs.
The architecture overview describes the system design, component boundaries, and domain glossary.
Agentic Identity Broker has a built-in consent management UI:
For contribution rules, test commands, and coverage reports, see CONTRIBUTING.md .
Agentic Identity Broker is currently in active development. We use this project in production internally for several MCP servers and have seen good potential. We are coding in the open to gauge our approach. If you find it useful, let us know. If you'd like a feature that's missing, open an issue in our GitHub repo .
While this project was created by professionals with decades of experience in the identity space (including building multiple OIDC and OAuth2 providers), it was developed with a strong focus on agentic engineering and spec-driven development. We review every PR and prioritize sound engineering practices over convenience.
Major roadmap topics include: (a) formalizing token exchange for agent-to-agent calls, (b) intent based authorization and (c) user facing acticity logs.
Thanks to all contributors who are helping to make the Agentic Identity Broker better.
Agentic Identity Broker is a Zalando Engineering project.
Identity broker for AI agents with OAuth2 delegation and consent management
agenticidentitybroker.dev Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Identity broker for AI agents with OAuth2 delegation and consent management - zalando-incubator/agentic-identity-broker

Hi HN, I’m Jan. I am co-author of the Agentic Identity Broker (AIB), an MIT-licensed OAuth2 broker in the Zalando Incubator, that lets a user give an AI agent scoped, revocable access to MCP services. The broker solves three major challenges that are especially relevant in corporate environments: (A) Enable users to delegate workloads to agents without providing the agents all the user’s tokens. (B) Connect fragmented identity and authorization systems. (C) Centralize tool approvals and audit logging. We built AIB to fill an infrastructure gap in our own agent platform. At this moment in time, IETF working groups are still finalizing standards for agent identity, and vendors are still turning the necessary infrastructure into off-the-shelf products. We needed that infrastructure today. How it works: the user consents once per agent and service. The broker stores the third-party tokens encrypted. At request time a gateway (e.g. Agentgateway) does RFC 8693 token exchange. Agents never see user tokens and MCP servers only implement tools and don’t need to care about all the OAuth 2.1 plumbing. AIB is v0.1 and runs in production, bringing together hundreds of our agents with dozens of internal and external MCP servers. To try it locally: clone the repo, run "just compose-up-detached". The current features focus on on-behalf-of (OBO) use cases, where human actors delegate to agents. Next we are looking to also support headless agents, where agents are invoked by machines. Intent-based authorization is yet another package on our roadmap. If you’ve built either, or have feedback on the UI/UX, I’d like to hear from you!

GitHub - zalando-incubator/agentic-identity-broker: Identity broker for AI agents with OAuth2 delegation and consent management · GitHub
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
zalando-incubator
agentic-identity-broker Public Notifications You must be signed in to change notification settings
Star 34 ( 34 ) You must be signed in to star a repository
Identity broker for AI agents with OAuth2 delegation and consent management
agenticidentitybroker.dev Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
262 Commits 262 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills .githooks .githooks .github .github .specify .specify .vscode .vscode adrs adrs api api assets/ docusaurus assets/ docusaurus build build charts/ agentic-identity-broker charts/ agentic-identity-broker cmd cmd coding-standards coding-standards configs configs docs docs examples/ config examples/ config infra infra internal internal migrations migrations mocks mocks scripts scripts specs specs tests tests tools tools web web .air.toml .air.toml .dockerignore .dockerignore .env.compose.example .env.compose.example .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .zappr.yaml .zappr.yaml AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS MAINTAINERS README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md VERSION VERSION config.yaml config.yaml delivery.yaml delivery.yaml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum justfile justfile osv-scanner.toml osv-scanner.toml View all files Repository files navigation
Delegate scoped, revocable access to AI agents — without handing them your credentials.
Agentic Identity Broker manages delegation chains for on-behalf-of flows in Agentic AI, bridges distinct OAuth2 infrastructures, and provides a secure token vault. Deployed within the call path via gateways like Agentgateway , it offloads auth burdens from both agent and MCP server creators.
For users, it provides a simple way to manage consent, granting agents granular permissions across services such as Google, GitHub, Databricks, and Linear without compromising control.
Secure Identity Management
Agent registration, principal authentication, and session management
OAuth2 Delegation & Consent
Manage OAuth2 scopes and user consent for third-party service integrations
Security-First Design
Sensitive value redaction, command injection prevention, cryptographic validation, and fail-closed architecture
Rich Consent Experience
Modern UX for managing OAuth2 delegations with design system compliance
Comprehensive Audit Logging
Structured JSON logging for compliance and security monitoring
Frameworks like FastMCP can implement the whole OAuth2.1 ceremonies as mandated by the MCP spec but this means that you have to configure it for every MCP Server you deploy and maintain it. You also have to configure secure storage and basically run a multitude of OAuth2 authorization servers. With the identity broker and a gateway you can dumb down MCP server development to only provide tools on the target technology you want. OAuth2 ceremonies, token vaulting and token translation are done transparently. Because token exchange is done centrally via a gateway, this allows agents to interact with hundreds of MCP servers via one channel.
If you want to offer an MCP server to your customers as part of your SaaS, the identity broker can provide an additional consent surface that records which agents were actually used. With centralized tool authorization, you can keep an audit trail of approvals on your side regardless of which agent the customer runs.
The agentic identity broker can simplify agents by requiring only a single user token that can be used with an arbitrary number of MCP servers. When paired with a portal for user interactions, token procurement can be front-loaded in the portal, and the agent's responsibility is to forward this token to upstream MCP servers via a gateway.
Delegation Chains Without User Presence
The impersonation feature allows to start delegation chains via privileged clients, like a Google Chat or Slack bridge. The broker can be instructed to trust for two independent machine identities to assume the session for a particular user. This allows to use the permissions of a user on the company chat in another third-party system.
The tool approval support in the broker allows to enforce tool calls independent of the agent interface in use. This allows a last resort check before emails get sent or documents shared with a central audit trail.
Quickstart — Get started with Agentic Identity Broker in minutes.
agenticidentitybroker.dev – Underlying concepts, guides and API specs.
The architecture overview describes the system design, component boundaries, and domain glossary.
Agentic Identity Broker has a built-in consent management UI:
For contribution rules, test commands, and coverage reports, see CONTRIBUTING.md .
Agentic Identity Broker is currently in active development. We use this project in production internally for several MCP servers and have seen good potential. We are coding in the open to gauge our approach. If you find it useful, let us know. If you'd like a feature that's missing, open an issue in our GitHub repo .
While this project was created by professionals with decades of experience in the identity space (including building multiple OIDC and OAuth2 providers), it was developed with a strong focus on agentic engineering and spec-driven development. We review every PR and prioritize sound engineering practices over convenience.
Major roadmap topics include: (a) formalizing token exchange for agent-to-agent calls, (b) intent based authorization and (c) user facing acticity logs.
Thanks to all contributors who are helping to make the Agentic Identity Broker better.
Agentic Identity Broker is a Zalando Engineering project.
Identity broker for AI agents with OAuth2 delegation and consent management
agenticidentitybroker.dev Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
