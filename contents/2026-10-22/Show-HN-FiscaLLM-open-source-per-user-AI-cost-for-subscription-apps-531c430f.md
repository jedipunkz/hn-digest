---
source: "https://github.com/magna56/fiscallm"
hn_url: "https://news.ycombinator.com/item?id=50195023"
title: "Show HN: FiscaLLM:open-source per-user AI cost for subscription apps"
article_title: "GitHub - magna56/fiscallm: Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. · GitHub"
image: "https://repository-images.githubusercontent.com/1414899309/f784894f-baad-473b-8419-3d9ff6677c20"
author: "theaicommit"
captured_at: "2026-10-22T13:31:14Z"
capture_tool: "hn-digest"
hn_id: 50195023
score: 1
comments: 0
posted_at: "2026-10-22T13:27:23Z"
tags:
  - hacker-news
---

# Show HN: FiscaLLM:open-source per-user AI cost for subscription apps

- HN: [50195023](https://news.ycombinator.com/item?id=50195023)
- Source: [github.com](https://github.com/magna56/fiscallm)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T13:27:23Z

## Translation

Title: Show HN: FiscaLLM:open-source per-user AI cost for subscription apps
Article title: GitHub - magna56/fiscallm: Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. · GitHub
Description: Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. - magna56/fiscallm

Article text:
GitHub - magna56/fiscallm: Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. · GitHub
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
magna56
fiscallm Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs.
Readme MIT license Code of conduct
Security policy Activity Stars
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
105 Commits 105 Commits Folders and files
.github .github apps apps docs docs examples examples packages packages .env.example .env.example .gitignore .gitignore ACCURACY.md ACCURACY.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SELF-HOSTING.md SELF-HOSTING.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json vercel.json vercel.json View all files Repository files navigation
See the AI cost of every single user. FiscaLLM adds up what each user costs you in AI and sets it next to what they pay you, so you can see who is profitable and who is losing you money.
Open source (MIT) and free. Use the hosted service at fiscallm.com (free under fair use, best effort), or run your own .
Live demo (no sign-up) · Docs · How it works · What it covers · Accuracy · Roadmap ( board ) · Changelog · Self-hosting · Contributing
Sample data from the demo, not a real customer.
Create an app at fiscallm.com (or your own copy) and copy the ingest key.
Add one line after each AI call (it reads only the model name and token counts, never prompts or answers):
import { createMeter } from "fiscallm" ; // npm install fiscallm (Python: pip install fiscallm)
const meter = createMeter ( { apiKey : process . env . FISCALLM_KEY , waitUntil } ) ;
const res = await openai . chat . completions . create ( { model , messages } ) ;
meter . trackOpenAI ( res , { userId , feature : "chat" } ) ;
Connect revenue : paste a RevenueCat or Stripe webhook URL, or send payments yourself with the revenue API.
No AI key handy? FISCALLM_KEY=fsk_... node examples/send-sample-data.mjs fills a test app with sample users so you can see it working.
AI apps charge a flat price but pay a variable cost. One user costs $0.20 a month in AI, another costs $25. Your AI provider's bill shows tokens, your subscription tool shows payments, and neither shows the user.
import { createMeter } from "fiscallm" ;
const meter = createMeter ( { apiKey : process . env . FISCALLM_KEY , waitUntil } ) ;
const res = await openai . chat . completions . create ( { model , messages } ) ;
meter . trackOpenAI ( res , { userId , feature : "chat" } ) ;
from fiscallm import Meter
meter = Meter ( api_key = os . environ [ "FISCALLM_KEY" ])
res = client . chat . completions . create ( model = model , messages = messages )
meter . track_openai ( res , user_id = user_id , feature = "chat" )
The SDK reads only the model name and token counts from the response. Never prompts or answers.
Events are batched and sent in the background. Tracking never throws and never blocks your response.
Each event has a client-generated UUID, so retries cannot double count.
Revenue comes from a RevenueCat webhook, or from your own code with trackRevenue / track_revenue .
Both are joined on the user ID. RevenueCat aliases are resolved, so anonymous and logged-in users merge.
Budgets and alerts : email, Slack or signed webhooks when a user goes over budget, when a paying user loses money, or when one user burns a huge number of tokens in 5 minutes (a runaway loop).
Spending caps : mark a per-user budget as Block and call meter.check(userId) before an AI call (TypeScript, Python). It says no once the user reaches the amount, and fails open if FiscaLLM is slow or down.
Cost is priced on arrival: exact costUsd if given, otherwise per-token (input, output, cache read and cache write priced separately) or per-unit (images, audio, video) from a price table. Unknown models are flagged as unpriced, never counted as zero.
Net revenue is price × (1 − tax − store commission) .
Path
What it is
apps/web
Next.js site and dashboard
apps/api
Hono API: ingest, RevenueCat webhook, revenue API, cron jobs
packages/shared
Drizzle schema, pricing, metrics, alerts
packages/sdk
TypeScript SDK ( fiscallm on npm)
packages/sdk-python
Python SDK ( fiscallm on PyPI)
Stack: TypeScript, Next.js, Hono, Postgres (Supabase) with Drizzle, hosted on Vercel.
It needs a backend between your app and the AI provider. Apps that call a provider straight from the phone are not covered yet.
Revenue comes from RevenueCat, Stripe (beta, USD only) or the revenue API. Direct App Store and Google Play connections are not built yet.
No AI key needed: FISCALLM_KEY=fsk_... node examples/send-sample-data.mjs fills the dashboard with sample users (one heavy user who loses money).
Copy-paste starting points for Next.js and FastAPI , see examples .
Run your own copy: (create a Supabase project first, see SELF-HOSTING.md ).
pnpm install
pnpm dev:web # http://localhost:3000
pnpm dev:api # http://localhost:8787/api/health
pnpm test
Copy .env.example to apps/web/.env.local and packages/shared/.env and fill in the values.
pnpm db:generate
pnpm db:migrate
Contributing
Issues and pull requests are welcome. Read CONTRIBUTING.md first (it includes a short contributor agreement). The plan, including what we only build if there is demand, is in ROADMAP.md , with a public board . Vote with a 👍 on roadmap issues or share ideas in Discussions .
Where help is most wanted: prices for more models and providers (data only, no code knowledge needed), connectors for other billing systems (Adapty, Superwall), examples for AI frameworks, and a comparison of computed costs with a real provider invoice. See the good first issues .
Built and maintained by Mayur as an independent, open-source side project. Questions, bugs and ideas: open an issue or email theaicommit@gmail.com .
Thank you to everyone who helps. Contribute : a good place to start is an issue labelled good first issue .
Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs.
Readme MIT license Code of conduct
Security policy Activity Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. - magna56/fiscallm

GitHub - magna56/fiscallm: Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs. · GitHub
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
magna56
fiscallm Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs.
Readme MIT license Code of conduct
Security policy Activity Stars
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
105 Commits 105 Commits Folders and files
.github .github apps apps docs docs examples examples packages packages .env.example .env.example .gitignore .gitignore ACCURACY.md ACCURACY.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SELF-HOSTING.md SELF-HOSTING.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json vercel.json vercel.json View all files Repository files navigation
See the AI cost of every single user. FiscaLLM adds up what each user costs you in AI and sets it next to what they pay you, so you can see who is profitable and who is losing you money.
Open source (MIT) and free. Use the hosted service at fiscallm.com (free under fair use, best effort), or run your own .
Live demo (no sign-up) · Docs · How it works · What it covers · Accuracy · Roadmap ( board ) · Changelog · Self-hosting · Contributing
Sample data from the demo, not a real customer.
Create an app at fiscallm.com (or your own copy) and copy the ingest key.
Add one line after each AI call (it reads only the model name and token counts, never prompts or answers):
import { createMeter } from "fiscallm" ; // npm install fiscallm (Python: pip install fiscallm)
const meter = createMeter ( { apiKey : process . env . FISCALLM_KEY , waitUntil } ) ;
const res = await openai . chat . completions . create ( { model , messages } ) ;
meter . trackOpenAI ( res , { userId , feature : "chat" } ) ;
Connect revenue : paste a RevenueCat or Stripe webhook URL, or send payments yourself with the revenue API.
No AI key handy? FISCALLM_KEY=fsk_... node examples/send-sample-data.mjs fills a test app with sample users so you can see it working.
AI apps charge a flat price but pay a variable cost. One user costs $0.20 a month in AI, another costs $25. Your AI provider's bill shows tokens, your subscription tool shows payments, and neither shows the user.
import { createMeter } from "fiscallm" ;
const meter = createMeter ( { apiKey : process . env . FISCALLM_KEY , waitUntil } ) ;
const res = await openai . chat . completions . create ( { model , messages } ) ;
meter . trackOpenAI ( res , { userId , feature : "chat" } ) ;
from fiscallm import Meter
meter = Meter ( api_key = os . environ [ "FISCALLM_KEY" ])
res = client . chat . completions . create ( model = model , messages = messages )
meter . track_openai ( res , user_id = user_id , feature = "chat" )
The SDK reads only the model name and token counts from the response. Never prompts or answers.
Events are batched and sent in the background. Tracking never throws and never blocks your response.
Each event has a client-generated UUID, so retries cannot double count.
Revenue comes from a RevenueCat webhook, or from your own code with trackRevenue / track_revenue .
Both are joined on the user ID. RevenueCat aliases are resolved, so anonymous and logged-in users merge.
Budgets and alerts : email, Slack or signed webhooks when a user goes over budget, when a paying user loses money, or when one user burns a huge number of tokens in 5 minutes (a runaway loop).
Spending caps : mark a per-user budget as Block and call meter.check(userId) before an AI call (TypeScript, Python). It says no once the user reaches the amount, and fails open if FiscaLLM is slow or down.
Cost is priced on arrival: exact costUsd if given, otherwise per-token (input, output, cache read and cache write priced separately) or per-unit (images, audio, video) from a price table. Unknown models are flagged as unpriced, never counted as zero.
Net revenue is price × (1 − tax − store commission) .
Path
What it is
apps/web
Next.js site and dashboard
apps/api
Hono API: ingest, RevenueCat webhook, revenue API, cron jobs
packages/shared
Drizzle schema, pricing, metrics, alerts
packages/sdk
TypeScript SDK ( fiscallm on npm)
packages/sdk-python
Python SDK ( fiscallm on PyPI)
Stack: TypeScript, Next.js, Hono, Postgres (Supabase) with Drizzle, hosted on Vercel.
It needs a backend between your app and the AI provider. Apps that call a provider straight from the phone are not covered yet.
Revenue comes from RevenueCat, Stripe (beta, USD only) or the revenue API. Direct App Store and Google Play connections are not built yet.
No AI key needed: FISCALLM_KEY=fsk_... node examples/send-sample-data.mjs fills the dashboard with sample users (one heavy user who loses money).
Copy-paste starting points for Next.js and FastAPI , see examples .
Run your own copy: (create a Supabase project first, see SELF-HOSTING.md ).
pnpm install
pnpm dev:web # http://localhost:3000
pnpm dev:api # http://localhost:8787/api/health
pnpm test
Copy .env.example to apps/web/.env.local and packages/shared/.env and fill in the values.
pnpm db:generate
pnpm db:migrate
Contributing
Issues and pull requests are welcome. Read CONTRIBUTING.md first (it includes a short contributor agreement). The plan, including what we only build if there is demand, is in ROADMAP.md , with a public board . Vote with a 👍 on roadmap issues or share ideas in Discussions .
Where help is most wanted: prices for more models and providers (data only, no code knowledge needed), connectors for other billing systems (Adapty, Superwall), examples for AI frameworks, and a comparison of computed costs with a real provider invoice. See the good first issues .
Built and maintained by Mayur as an independent, open-source side project. Questions, bugs and ideas: open an issue or email theaicommit@gmail.com .
Thank you to everyone who helps. Contribute : a good place to start is an issue labelled good first issue .
Open-source per-user AI cost and margin tracking for subscription apps. Matches each user's AI cost to what they pay you (RevenueCat, Stripe). TypeScript and Python SDKs.
Readme MIT license Code of conduct
Security policy Activity Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
