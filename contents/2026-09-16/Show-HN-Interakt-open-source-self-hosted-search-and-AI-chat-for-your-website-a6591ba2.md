---
source: "https://github.com/alphasolutionsrepo/interakt"
hn_url: "https://news.ycombinator.com/item?id=49732963"
title: "Show HN: Interakt – open-source self-hosted search and AI chat for your website"
article_title: "GitHub - alphasolutionsrepo/interakt: Open-source, self-hosted search and AI chat for your website, grounded in your own data · GitHub"
image: "https://opengraph.githubassets.com/98c4d8b8cd01206d500423faf771b0db302eea88d1c50c8aac9aacfa5e9cabee/alphasolutionsrepo/interakt"
author: "interakt"
captured_at: "2026-09-16T21:55:42Z"
capture_tool: "hn-digest"
hn_id: 49732963
score: 2
comments: 0
posted_at: "2026-09-16T21:01:06Z"
tags:
  - hacker-news
---

# Show HN: Interakt – open-source self-hosted search and AI chat for your website

- HN: [49732963](https://news.ycombinator.com/item?id=49732963)
- Source: [github.com](https://github.com/alphasolutionsrepo/interakt)
- Score: 2
- Comments: 0
- Posted: 2026-09-16T21:01:06Z

## Translation

Title: Show HN: Interakt – open-source self-hosted search and AI chat for your website
Article title: GitHub - alphasolutionsrepo/interakt: Open-source, self-hosted search and AI chat for your website, grounded in your own data · GitHub
Description: Open-source, self-hosted search and AI chat for your website, grounded in your own data - alphasolutionsrepo/interakt

Article text:
GitHub - alphasolutionsrepo/interakt: Open-source, self-hosted search and AI chat for your website, grounded in your own data · GitHub
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
alphasolutionsrepo
/
interakt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
40 Commits 40 Commits Folders and files
.github .github .vscode .vscode backend backend demo-site demo-site docs-site docs-site .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Open-source, self-hosted search and AI chat for your website, grounded in your own data.
Website · Documentation · API Reference · Integrations · Discussions
Interakt puts a search box and a chat window in front of your users without you assembling the pieces yourself. Point it at your data, decide how it should behave in the admin dashboard, and paste a snippet into your site.
Underneath is a curated stack: a search engine, an AI provider, a chat pipeline, a versioned prompt library and an analytics database. You see one dashboard, one API and two widgets.
Search that understands intent. Keyword, semantic or hybrid, with facets, synonyms, autocomplete and AI query understanding.
Chat that knows your data. Answers are grounded in your indexes through tools, streamed, with optional citations, and kept on topic by guardrails.
Analytics that show what works. Search events, chat sessions and execution traces land in a separate analytics database, with dashboards and an analytics assistant.
Backend setup
From git clone to localhost:3000
AI search and a sales assistant
Interakt on a Medusa storefront
Managing the integration
The admin side of the Medusa setup
Highlights
🔍 Hybrid search
Keyword and vector results fused with reciprocal rank fusion, on Elasticsearch 9 or Azure AI Search. Facets, sorting, synonyms and stop words per index.
💬 Grounded chat
A deterministic pipeline (plan → retrieve → synthesize) or an agentic loop, chosen per experience. Streaming responses, with optional inline or footnote citations.
🧩 Tools and MCP
Every index ships with search, lookup, inspect and enumerate tools. Add custom HTTP tools, or connect MCP servers over Streamable HTTP or SSE.
🧠 Bring your own model
OpenAI, OpenAI-compatible endpoints through a custom base URL, or Ollama for fully local inference. Choose the provider per experience.
📦 Drop-in widgets
Preact and Shadow DOM, loaded from a single script bundle. Search and chat widgets that work on any site, no framework required.
📝 Versioned prompts
Every pipeline step is an editable template with history and rollback. Tune behaviour without redeploying.
🛡️ Guardrails and credentials
Topic gating and greeting detection. Read-only public access tokens, server-side ingestion keys, and an encrypted secrets vault.
🏠 Self-hosted
One Docker image, Postgres with pgvector, and Elasticsearch or Azure AI Search. Runs wherever Docker runs; Alpha Solutions runs it on Azure Container Apps. MIT licensed.
How it fits together
flowchart LR
D["Your data<br/>catalog · docs · CMS · files"] --> IX["Search index"]
IX --> SE["Search experience"]
IX -- "tools" --> AE["AI chat experience"]
T["Custom HTTP tools · MCP servers"] --> AE
SE --> W["Widgets · REST API"]
AE --> W
W --> Y["Your website"]
IX -.-> ES[("Elasticsearch or<br/>Azure AI Search")]
AE -.-> LLM[("OpenAI · Ollama")]
Loading
You connect a data source, populate a search index, then build one or more experiences on top of it. The chat experience calls the search tools to answer from your data. The architecture overview has the full picture.
Needs Node.js 24 and Docker. Prefer to watch? The backend setup video walks through these steps.
git clone https://github.com/alphasolutionsrepo/interakt.git
cd interakt/backend
cp .env.example .env # fill in the three generated secrets at the top
cp setup/setup.config.example.yaml setup/setup.config.yaml # set your admin email and password
npm install
npm run infra:up # Postgres (pgvector) and Elasticsearch in Docker
npm run dev # migrates, seeds, creates the admin user → http://localhost:3000
Sign in and open Platform → Initial Setup . Connect an AI provider (Ollama is free and local), then load the Fashion Catalog demo: a populated index plus ready-made search and chat experiences with access tokens. Point the reference app in demo-site/ at it, or embed the widget below.
The full walkthrough, scripts and bring-your-own-database notes are in CONTRIBUTING.md . Hosting notes are in backend/docker/ .
< div id =" chat " > </ div >
< script src =" https://your-interakt-host/embed/v1/widgets.js " > </ script >
< script >
window . ChatDropinUI . init ( {
containerId : 'chat' ,
accessToken : '<access token from the admin UI>' ,
} ) ;
</ script >
Use SearchDropinUI for the search widget. The admin UI generates the exact snippet for each experience, and the widget docs cover launcher modes, placement and theming.
# Search through a search experience
curl -X POST https://your-interakt-host/api/v1/search/ < experience-slug > /search \
-H " Authorization: Bearer <access token> " \
-H " Content-Type: application/json " \
-d ' { "query": "waterproof trail shoes" } '
# Chat through an AI experience (streams server-sent events)
curl -N -X POST https://your-interakt-host/api/v1/ai-experiences/ < experience-slug > /chat \
-H " Authorization: Bearer <access token> " \
-H " Content-Type: application/json " \
-d ' { "message": "Which of these run true to size?" } '
Access tokens are issued per experience and are public and read-only by design. Writing documents into an index uses a server-side ingestion key instead. See the search , chat and ingestion guides, or the full API reference .
Step-by-step guides for Medusa , Storyblok and React / Next.js . The Interakt-Medusa repository is the demo storefront that guide is built on.
interakt/
├── backend/ Admin dashboard + REST APIs · Next.js 16, Drizzle, Postgres + pgvector
│ ├── widgets/ Embeddable search and chat widgets · Preact, built to public/embed/v1/widgets.js
│ ├── docker/ Production Dockerfile + local docker-compose (Postgres, pgAdmin, Elasticsearch)
│ └── src/content/docs/ The documentation, served in-app at /docs and published to docs.interakt.app
├── demo-site/ Reference consumer app on port 3001
└── docs-site/ Docusaurus + Redocusaurus shell for docs.interakt.app
Contributing
Contributions are welcome. CONTRIBUTING.md covers local setup, conventions and how review works here. Found a bug or want a feature? Open an issue . Have a question? Start a discussion . For security issues, follow SECURITY.md and never open a public issue.
This project follows the Contributor Covenant code of conduct.
Open-source, self-hosted search and AI chat for your website, grounded in your own data
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source, self-hosted search and AI chat for your website, grounded in your own data - alphasolutionsrepo/interakt

GitHub - alphasolutionsrepo/interakt: Open-source, self-hosted search and AI chat for your website, grounded in your own data · GitHub
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
alphasolutionsrepo
/
interakt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
40 Commits 40 Commits Folders and files
.github .github .vscode .vscode backend backend demo-site demo-site docs-site docs-site .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Open-source, self-hosted search and AI chat for your website, grounded in your own data.
Website · Documentation · API Reference · Integrations · Discussions
Interakt puts a search box and a chat window in front of your users without you assembling the pieces yourself. Point it at your data, decide how it should behave in the admin dashboard, and paste a snippet into your site.
Underneath is a curated stack: a search engine, an AI provider, a chat pipeline, a versioned prompt library and an analytics database. You see one dashboard, one API and two widgets.
Search that understands intent. Keyword, semantic or hybrid, with facets, synonyms, autocomplete and AI query understanding.
Chat that knows your data. Answers are grounded in your indexes through tools, streamed, with optional citations, and kept on topic by guardrails.
Analytics that show what works. Search events, chat sessions and execution traces land in a separate analytics database, with dashboards and an analytics assistant.
Backend setup
From git clone to localhost:3000
AI search and a sales assistant
Interakt on a Medusa storefront
Managing the integration
The admin side of the Medusa setup
Highlights
🔍 Hybrid search
Keyword and vector results fused with reciprocal rank fusion, on Elasticsearch 9 or Azure AI Search. Facets, sorting, synonyms and stop words per index.
💬 Grounded chat
A deterministic pipeline (plan → retrieve → synthesize) or an agentic loop, chosen per experience. Streaming responses, with optional inline or footnote citations.
🧩 Tools and MCP
Every index ships with search, lookup, inspect and enumerate tools. Add custom HTTP tools, or connect MCP servers over Streamable HTTP or SSE.
🧠 Bring your own model
OpenAI, OpenAI-compatible endpoints through a custom base URL, or Ollama for fully local inference. Choose the provider per experience.
📦 Drop-in widgets
Preact and Shadow DOM, loaded from a single script bundle. Search and chat widgets that work on any site, no framework required.
📝 Versioned prompts
Every pipeline step is an editable template with history and rollback. Tune behaviour without redeploying.
🛡️ Guardrails and credentials
Topic gating and greeting detection. Read-only public access tokens, server-side ingestion keys, and an encrypted secrets vault.
🏠 Self-hosted
One Docker image, Postgres with pgvector, and Elasticsearch or Azure AI Search. Runs wherever Docker runs; Alpha Solutions runs it on Azure Container Apps. MIT licensed.
How it fits together
flowchart LR
D["Your data<br/>catalog · docs · CMS · files"] --> IX["Search index"]
IX --> SE["Search experience"]
IX -- "tools" --> AE["AI chat experience"]
T["Custom HTTP tools · MCP servers"] --> AE
SE --> W["Widgets · REST API"]
AE --> W
W --> Y["Your website"]
IX -.-> ES[("Elasticsearch or<br/>Azure AI Search")]
AE -.-> LLM[("OpenAI · Ollama")]
Loading
You connect a data source, populate a search index, then build one or more experiences on top of it. The chat experience calls the search tools to answer from your data. The architecture overview has the full picture.
Needs Node.js 24 and Docker. Prefer to watch? The backend setup video walks through these steps.
git clone https://github.com/alphasolutionsrepo/interakt.git
cd interakt/backend
cp .env.example .env # fill in the three generated secrets at the top
cp setup/setup.config.example.yaml setup/setup.config.yaml # set your admin email and password
npm install
npm run infra:up # Postgres (pgvector) and Elasticsearch in Docker
npm run dev # migrates, seeds, creates the admin user → http://localhost:3000
Sign in and open Platform → Initial Setup . Connect an AI provider (Ollama is free and local), then load the Fashion Catalog demo: a populated index plus ready-made search and chat experiences with access tokens. Point the reference app in demo-site/ at it, or embed the widget below.
The full walkthrough, scripts and bring-your-own-database notes are in CONTRIBUTING.md . Hosting notes are in backend/docker/ .
< div id =" chat " > </ div >
< script src =" https://your-interakt-host/embed/v1/widgets.js " > </ script >
< script >
window . ChatDropinUI . init ( {
containerId : 'chat' ,
accessToken : '<access token from the admin UI>' ,
} ) ;
</ script >
Use SearchDropinUI for the search widget. The admin UI generates the exact snippet for each experience, and the widget docs cover launcher modes, placement and theming.
# Search through a search experience
curl -X POST https://your-interakt-host/api/v1/search/ < experience-slug > /search \
-H " Authorization: Bearer <access token> " \
-H " Content-Type: application/json " \
-d ' { "query": "waterproof trail shoes" } '
# Chat through an AI experience (streams server-sent events)
curl -N -X POST https://your-interakt-host/api/v1/ai-experiences/ < experience-slug > /chat \
-H " Authorization: Bearer <access token> " \
-H " Content-Type: application/json " \
-d ' { "message": "Which of these run true to size?" } '
Access tokens are issued per experience and are public and read-only by design. Writing documents into an index uses a server-side ingestion key instead. See the search , chat and ingestion guides, or the full API reference .
Step-by-step guides for Medusa , Storyblok and React / Next.js . The Interakt-Medusa repository is the demo storefront that guide is built on.
interakt/
├── backend/ Admin dashboard + REST APIs · Next.js 16, Drizzle, Postgres + pgvector
│ ├── widgets/ Embeddable search and chat widgets · Preact, built to public/embed/v1/widgets.js
│ ├── docker/ Production Dockerfile + local docker-compose (Postgres, pgAdmin, Elasticsearch)
│ └── src/content/docs/ The documentation, served in-app at /docs and published to docs.interakt.app
├── demo-site/ Reference consumer app on port 3001
└── docs-site/ Docusaurus + Redocusaurus shell for docs.interakt.app
Contributing
Contributions are welcome. CONTRIBUTING.md covers local setup, conventions and how review works here. Found a bug or want a feature? Open an issue . Have a question? Start a discussion . For security issues, follow SECURITY.md and never open a public issue.
This project follows the Contributor Covenant code of conduct.
Open-source, self-hosted search and AI chat for your website, grounded in your own data
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
