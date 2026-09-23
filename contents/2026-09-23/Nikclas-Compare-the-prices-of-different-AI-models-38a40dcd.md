---
source: "https://github.com/NikclasTech/nikclas"
hn_url: "https://news.ycombinator.com/item?id=49816775"
title: "Nikclas – Compare the prices of different AI models"
article_title: "GitHub - NikclasTech/nikclas: Compare the prices of different AI models before starting your larger project. · GitHub"
image: "https://opengraph.githubassets.com/e6562b8f21e0199a53d6bbafd6b6d3631865a30d0125cf4f7c805c31eca007aa/NikclasTech/nikclas"
author: "Cr12dev"
captured_at: "2026-09-23T15:19:57Z"
capture_tool: "hn-digest"
hn_id: 49816775
score: 2
comments: 2
posted_at: "2026-09-23T14:30:23Z"
tags:
  - hacker-news
---

# Nikclas – Compare the prices of different AI models

- HN: [49816775](https://news.ycombinator.com/item?id=49816775)
- Source: [github.com](https://github.com/NikclasTech/nikclas)
- Score: 2
- Comments: 2
- Posted: 2026-09-23T14:30:23Z

## Translation

Title: Nikclas – Compare the prices of different AI models
Article title: GitHub - NikclasTech/nikclas: Compare the prices of different AI models before starting your larger project. · GitHub
Description: Compare the prices of different AI models before starting your larger project. - NikclasTech/nikclas

Article text:
GitHub - NikclasTech/nikclas: Compare the prices of different AI models before starting your larger project. · GitHub
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
NikclasTech
/
nikclas
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
19 Commits 19 Commits Folders and files
.github .github components components data data docs docs public public scripts scripts src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock eslint.config.js eslint.config.js index.html index.html netlify.toml netlify.toml package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Stop overpaying for tokens. Nikclas compares $/1M token prices across
frontier AI models so you can ship the same quality for less.
Read this in: Espanol | Francais |
Deutsch | 中文
Nikclas is a single-page React application that answers one question:
which AI model does the same job for the least money?
It opens with a thesis hero (live price board sorted by input $/1M) and
continues with a head-to-head comparator: pick a provider and a model on
each side, see the price gap, and get a verdict with the exact saving per
1M input tokens.
All prices load live from the
LiteLLM Model Catalog API .
Live price board — the 4 highest-traffic flagships
( gpt-4o-mini , gpt-4o , claude-sonnet-4-6 , deepseek-chat ), sorted
by input $/1M with log-scale bars, BEST / AVOID tags and a live /
syncing / cached status indicator.
Head-to-head comparator — provider and model selectors on both
sides, swap button, per-model datasheets (input, output, context,
speed), a central savings readout and a plain-language verdict.
Live pricing with honest fallback — prices are fetched from
https://api.litellm.ai/model_catalog/{model_id} , cached in
localStorage for 24h, and fall back to a bundled snapshot (verified
2026-09-22) while loading or offline. The UI always states which source
is on screen.
Tech-styled identity — Chakra Petch display type, Inter body text,
JetBrains Mono data type; deep-navy blueprint theme with signal-cyan
(compute) and amber (savings) accents.
Accessible by default — semantic landmarks, labelled form controls,
visible keyboard focus, live regions for price updates, and
prefers-reduced-motion support.
Layer
Technology
UI
React 19 + TypeScript
Build
Vite 8
Styling
Tailwind CSS 4
Fonts
Chakra Petch, Inter, JetBrains Mono (Google Fonts)
Data
LiteLLM Model Catalog API ( api.litellm.ai )
Lint
ESLint + typescript-eslint
No backend is required. The app is a static build that calls the public
pricing API directly from the browser.
.
├── components/
│ ├── hero/
│ │ ├── HeroSection.tsx # Thesis, CTAs, stats + board layout
│ │ └── PriceBoard.tsx # Live price-board console
│ └── compare/
│ ├── CompareSection.tsx # State, savings math, verdict
│ ├── CompareForm.tsx # Provider + model selectors, swap
│ └── ModelCard.tsx # Model datasheet card
├── src/
│ ├── lib/
│ │ ├── litellm.ts # Catalog, API client, cache, hook
│ │ └── format.ts # Formatting, vendor labels, bar scales
│ ├── App.tsx # Composes HeroSection + CompareSection
│ ├── main.tsx # React entry point
│ └── index.css # Tailwind theme, fonts, animations
├── docs/
│ ├── README.es.md # Spanish translation
│ ├── README.fr.md # French translation
│ ├── README.de.md # German translation
│ └── README.zh.md # Chinese translation
├── public/ # Static assets (banner, favicon, icons)
├── index.html # Fonts, meta, title
└── LICENSE # GNU General Public License v3.0
5. Getting started
npm install
Run the dev server
npm run dev
Open http://localhost:5173/ in your browser.
npm run build
npm run preview
npm run build type-checks ( tsc -b ) and emits the static site to
dist/ , which can be served by any static host.
GET https://api.litellm.ai/model_catalog/{model_id} returns per-token
costs ( input_cost_per_token , output_cost_per_token ). The app
multiplies them by 1,000,000 to display $/1M tokens, and reads
max_input_tokens for the context window. Model metadata refreshes from
LiteLLM's catalog on every load (deduplicated, then cached 24h).
The API free tier allows 100 requests/day per IP without a key; a fresh
load of the full 15-model catalog uses 15 requests.
api.litellm.ai does not send an Access-Control-Allow-Origin header
(verified with and without an Origin header; OPTIONS returns 405),
so browsers may block the live response. In that case the app renders the
bundled snapshot and labels it as cached. Serving the app behind your own
backend or proxy that forwards to api.litellm.ai restores live data
without changing any component.
Bundled snapshot (verified 2026-09-22, $/1M)
Model
Provider
Input
Output
Context
gpt-4o-mini
OpenAI
$0.15
$0.60
128k
deepseek-chat
DeepSeek
$0.28
$0.42
131k
deepseek-v4-flash
DeepSeek
$0.30
$1.20
1M
deepinfra/nvidia/Llama-3.1-Nemotron-70B-Instruct
NVIDIA
$0.60
$0.60
131k
gemini-2.5-flash-lite
Google
$0.10
$0.40
1M
gpt-5.6-luna
OpenAI
$0.20
$1.20
922k
zai/glm-5.3
z.ai
$1.40
$4.40
1M
qwencloud/qwen-max
Qwen
$1.60
$6.40
31k
xai/grok-4.5
xAI
$2.00
$6.00
500k
gpt-4o
OpenAI
$2.50
$10.00
128k
claude-sonnet-4-6
Anthropic
$3.00
$15.00
1M
moonshot/kimi-k3
Kimi
$3.00
$15.00
1M
claude-fable-5-1
Anthropic
$10.00
$50.00
1M
gpt-6-astra
OpenAI
$10.00
$50.00
922k
claude-opus-4-1
Anthropic
$15.00
$75.00
200k
Median speeds shown on the cards are indicative static values; prices and
context windows are the live fields.
There is no database: data/pricing.json is the current dataset and
git log -- data/pricing.json is the price history. Every 6 hours the
pricing-update workflow runs the collector ( npm run collect:pricing ),
validates the result and opens a pull request titled
chore: update AI model pricing when prices moved — main is never
written to directly. Jumps of 3x or more are kept in the PR but marked
[possible-anomaly] for human review. Source links are re-checked
weekly by the link-check workflow, which opens an issue if a source
breaks.
Command
Description
npm run dev
Start the Vite dev server
npm run build
Type-check and build for production
npm run preview
Preview the production build locally
npm run lint
Run ESLint over the project
npm test
Run unit tests (vitest)
npm run collect:pricing
Refresh data/pricing.json from live sources
npm run validate:pricing
Validate data/*.json
npm run check:links
Check pricing source URLs
8. Automation
Workflow
Trigger
What it does
ci.yml
PRs, pushes to main
Install, lint, tests, build, data validation
pricing-update.yml
Every 6h, manual dispatch
Collector run, validation, pricing PR (never direct to main )
link-check.yml
Weekly, manual dispatch
Verifies source URLs, opens an issue on breakage
security.yml
PRs, pushes to main , weekly
CodeQL analysis (JavaScript/TypeScript)
Dependabot updates npm and GitHub Actions dependencies weekly
( .github/dependabot.yml , no auto-merge). No secrets are required:
the pricing API needs no key and automation uses the built-in
GITHUB_TOKEN with least-privilege permissions.
See CONTRIBUTING.md , including how to add a new
pricing provider (collector, parser, tests, official source,
validation, pull request).
This project is licensed under the GNU General Public License v3.0
or later (GPLv3+) . See LICENSE for the full text.
You are free to run, study, share and modify this software, provided
that any distributed work remains under the same license and its source
code stays available.
Compare the prices of different AI models before starting your larger project.
Readme GPL-3.0 license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Compare the prices of different AI models before starting your larger project. - NikclasTech/nikclas

GitHub - NikclasTech/nikclas: Compare the prices of different AI models before starting your larger project. · GitHub
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
NikclasTech
/
nikclas
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
19 Commits 19 Commits Folders and files
.github .github components components data data docs docs public public scripts scripts src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock eslint.config.js eslint.config.js index.html index.html netlify.toml netlify.toml package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Stop overpaying for tokens. Nikclas compares $/1M token prices across
frontier AI models so you can ship the same quality for less.
Read this in: Espanol | Francais |
Deutsch | 中文
Nikclas is a single-page React application that answers one question:
which AI model does the same job for the least money?
It opens with a thesis hero (live price board sorted by input $/1M) and
continues with a head-to-head comparator: pick a provider and a model on
each side, see the price gap, and get a verdict with the exact saving per
1M input tokens.
All prices load live from the
LiteLLM Model Catalog API .
Live price board — the 4 highest-traffic flagships
( gpt-4o-mini , gpt-4o , claude-sonnet-4-6 , deepseek-chat ), sorted
by input $/1M with log-scale bars, BEST / AVOID tags and a live /
syncing / cached status indicator.
Head-to-head comparator — provider and model selectors on both
sides, swap button, per-model datasheets (input, output, context,
speed), a central savings readout and a plain-language verdict.
Live pricing with honest fallback — prices are fetched from
https://api.litellm.ai/model_catalog/{model_id} , cached in
localStorage for 24h, and fall back to a bundled snapshot (verified
2026-09-22) while loading or offline. The UI always states which source
is on screen.
Tech-styled identity — Chakra Petch display type, Inter body text,
JetBrains Mono data type; deep-navy blueprint theme with signal-cyan
(compute) and amber (savings) accents.
Accessible by default — semantic landmarks, labelled form controls,
visible keyboard focus, live regions for price updates, and
prefers-reduced-motion support.
Layer
Technology
UI
React 19 + TypeScript
Build
Vite 8
Styling
Tailwind CSS 4
Fonts
Chakra Petch, Inter, JetBrains Mono (Google Fonts)
Data
LiteLLM Model Catalog API ( api.litellm.ai )
Lint
ESLint + typescript-eslint
No backend is required. The app is a static build that calls the public
pricing API directly from the browser.
.
├── components/
│ ├── hero/
│ │ ├── HeroSection.tsx # Thesis, CTAs, stats + board layout
│ │ └── PriceBoard.tsx # Live price-board console
│ └── compare/
│ ├── CompareSection.tsx # State, savings math, verdict
│ ├── CompareForm.tsx # Provider + model selectors, swap
│ └── ModelCard.tsx # Model datasheet card
├── src/
│ ├── lib/
│ │ ├── litellm.ts # Catalog, API client, cache, hook
│ │ └── format.ts # Formatting, vendor labels, bar scales
│ ├── App.tsx # Composes HeroSection + CompareSection
│ ├── main.tsx # React entry point
│ └── index.css # Tailwind theme, fonts, animations
├── docs/
│ ├── README.es.md # Spanish translation
│ ├── README.fr.md # French translation
│ ├── README.de.md # German translation
│ └── README.zh.md # Chinese translation
├── public/ # Static assets (banner, favicon, icons)
├── index.html # Fonts, meta, title
└── LICENSE # GNU General Public License v3.0
5. Getting started
npm install
Run the dev server
npm run dev
Open http://localhost:5173/ in your browser.
npm run build
npm run preview
npm run build type-checks ( tsc -b ) and emits the static site to
dist/ , which can be served by any static host.
GET https://api.litellm.ai/model_catalog/{model_id} returns per-token
costs ( input_cost_per_token , output_cost_per_token ). The app
multiplies them by 1,000,000 to display $/1M tokens, and reads
max_input_tokens for the context window. Model metadata refreshes from
LiteLLM's catalog on every load (deduplicated, then cached 24h).
The API free tier allows 100 requests/day per IP without a key; a fresh
load of the full 15-model catalog uses 15 requests.
api.litellm.ai does not send an Access-Control-Allow-Origin header
(verified with and without an Origin header; OPTIONS returns 405),
so browsers may block the live response. In that case the app renders the
bundled snapshot and labels it as cached. Serving the app behind your own
backend or proxy that forwards to api.litellm.ai restores live data
without changing any component.
Bundled snapshot (verified 2026-09-22, $/1M)
Model
Provider
Input
Output
Context
gpt-4o-mini
OpenAI
$0.15
$0.60
128k
deepseek-chat
DeepSeek
$0.28
$0.42
131k
deepseek-v4-flash
DeepSeek
$0.30
$1.20
1M
deepinfra/nvidia/Llama-3.1-Nemotron-70B-Instruct
NVIDIA
$0.60
$0.60
131k
gemini-2.5-flash-lite
Google
$0.10
$0.40
1M
gpt-5.6-luna
OpenAI
$0.20
$1.20
922k
zai/glm-5.3
z.ai
$1.40
$4.40
1M
qwencloud/qwen-max
Qwen
$1.60
$6.40
31k
xai/grok-4.5
xAI
$2.00
$6.00
500k
gpt-4o
OpenAI
$2.50
$10.00
128k
claude-sonnet-4-6
Anthropic
$3.00
$15.00
1M
moonshot/kimi-k3
Kimi
$3.00
$15.00
1M
claude-fable-5-1
Anthropic
$10.00
$50.00
1M
gpt-6-astra
OpenAI
$10.00
$50.00
922k
claude-opus-4-1
Anthropic
$15.00
$75.00
200k
Median speeds shown on the cards are indicative static values; prices and
context windows are the live fields.
There is no database: data/pricing.json is the current dataset and
git log -- data/pricing.json is the price history. Every 6 hours the
pricing-update workflow runs the collector ( npm run collect:pricing ),
validates the result and opens a pull request titled
chore: update AI model pricing when prices moved — main is never
written to directly. Jumps of 3x or more are kept in the PR but marked
[possible-anomaly] for human review. Source links are re-checked
weekly by the link-check workflow, which opens an issue if a source
breaks.
Command
Description
npm run dev
Start the Vite dev server
npm run build
Type-check and build for production
npm run preview
Preview the production build locally
npm run lint
Run ESLint over the project
npm test
Run unit tests (vitest)
npm run collect:pricing
Refresh data/pricing.json from live sources
npm run validate:pricing
Validate data/*.json
npm run check:links
Check pricing source URLs
8. Automation
Workflow
Trigger
What it does
ci.yml
PRs, pushes to main
Install, lint, tests, build, data validation
pricing-update.yml
Every 6h, manual dispatch
Collector run, validation, pricing PR (never direct to main )
link-check.yml
Weekly, manual dispatch
Verifies source URLs, opens an issue on breakage
security.yml
PRs, pushes to main , weekly
CodeQL analysis (JavaScript/TypeScript)
Dependabot updates npm and GitHub Actions dependencies weekly
( .github/dependabot.yml , no auto-merge). No secrets are required:
the pricing API needs no key and automation uses the built-in
GITHUB_TOKEN with least-privilege permissions.
See CONTRIBUTING.md , including how to add a new
pricing provider (collector, parser, tests, official source,
validation, pull request).
This project is licensed under the GNU General Public License v3.0
or later (GPLv3+) . See LICENSE for the full text.
You are free to run, study, share and modify this software, provided
that any distributed work remains under the same license and its source
code stays available.
Compare the prices of different AI models before starting your larger project.
Readme GPL-3.0 license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
