---
source: "https://github.com/DevDizzle/gammarips-harness"
hn_url: "https://news.ycombinator.com/item?id=50594625"
title: "Open-source harness and MCP dataset for options trading with AI agents"
article_title: "GitHub - DevDizzle/gammarips-harness: Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. · GitHub"
image: "https://opengraph.githubassets.com/21b7ae2cabec6318914b0332856e12d3df49567d888c7d169909e022fe0f2390/DevDizzle/gammarips-harness"
author: "gammarips"
captured_at: "2026-11-22T00:12:46Z"
capture_tool: "hn-digest"
hn_id: 50594625
score: 1
comments: 0
posted_at: "2026-11-21T23:20:48Z"
tags:
  - hacker-news
---

# Open-source harness and MCP dataset for options trading with AI agents

- HN: [50594625](https://news.ycombinator.com/item?id=50594625)
- Source: [github.com](https://github.com/DevDizzle/gammarips-harness)
- Score: 1
- Comments: 0
- Posted: 2026-11-21T23:20:48Z

## Translation

Title: Open-source harness and MCP dataset for options trading with AI agents
Article title: GitHub - DevDizzle/gammarips-harness: Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. · GitHub
Description: Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. - DevDizzle/gammarips-harness

Article text:
GitHub - DevDizzle/gammarips-harness: Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. · GitHub
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
DevDizzle
gammarips-harness Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude assets assets docs docs eval eval scripts scripts traces traces wiki wiki .gitignore .gitignore .mcp.json .mcp.json CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
An open-source agent workflow harness for the GammaRips MCP curated options dataset.
GammaRips is a data layer, not a signal service. Every trading day, the engine filters ~3,500 optionable US stocks down to a curated cohort of liquid, tradable, high-edge option contracts and serves that dataset over MCP. Your agent pulls the surviving cohort, layers its own live catalyst research (X/Twitter cashtags, filings, web search, parallel subagents), and builds a deterministic trade plan with historical MFE/MAE excursion data.
The Product: Curated Options Dataset via MCP
Instead of dumping noisy raw chains into an LLM context window, GammaRips filters the US options universe down to the daily cohort of surviving contracts that pass strict quantitative and liquidity gates:
Volatility Edge: RV/IV >= 0.85 (realized volatility vs. implied volatility)
Connect any MCP-compatible agent (Claude Code, Grok, Codex, Cursor, or custom runners):
Pro (OAuth): https://mcp.gammarips.com/pro
API Key (Bearer): https://mcp.gammarips.com/mcp with header Authorization: Bearer gr_live_...
Free Preview (No Key): https://mcp.gammarips.com/mcp (serves the daily report, regime context, market calendar, playbooks, and cohort preview)
1. Pull the Daily Cohort ( get_pool )
Fetch today's surviving tradable option contracts and their point-in-time volatility and liquidity metrics ( rv_iv_ratio , open_interest , mid_price , delta , dte , and opportunity surfaces), alongside market calendar status ( get_market_calendar_status ) and VIX term-structure regime ( get_regime_context ).
2. Layer Your Own Catalyst Search (Bring Your Own X / Web Agent)
GammaRips handles the quantitative options filtering; you bring the catalyst layer. Have your agent (Claude Code, Grok, Codex, Cursor) spin up parallel subagents or live X/Twitter and web search — for example, Grok searching live $TICKER cashtag chatter, SEC/FDA regulatory filings, or unpriced catalysts — to debate and rank the surviving contracts.
3. Build the Trade Plan / Execute
Size the position against your risk budget, set deterministic bracket exits (profit target, stop loss, and mandatory time-stop) using historical MFE/MAE excursion distributions ( query_outcomes and replay_contract ), pre-register the decision in eval/funnel-log.jsonl , and optionally route the order to your own broker MCP.
Try the free preview (no key, 10 seconds)
claude mcp add --transport http gammarips https://mcp.gammarips.com/mcp
Ask your agent for today's brief to inspect the market calendar, VIX regime context, daily report, and cohort preview.
Connect with Pro (OAuth or API Key)
Subscribe at gammarips.com/pricing and create an API key at gammarips.com/account (or authenticate directly via OAuth at https://mcp.gammarips.com/pro ).
Export your key for .mcp.json :
export GAMMARIPS_MCP_KEY= " gr_live_... "
Or install as a Claude Code plugin:
/plugin marketplace add DevDizzle/gammarips-harness
/plugin install gammarips@gammarips
Run the daily skills:
/trade Step 1–3 morning workflow: pull cohort, layer catalysts, build exit plan
/review After the close: score every cohort contract with MFE/MAE outcomes
/coach On demand: referee trading behavior against your written rules
Path
Description
CLAUDE.md
Agent instructions: dataset specification, 3-step workflow, and core rules
.mcp.json
Streamable HTTP MCP config ( https://mcp.gammarips.com/mcp with Bearer ${GAMMARIPS_MCP_KEY} )
docs/TRADING-DOCTRINE.md
Risk rules, hard exclusions, sizing framework, and deterministic exit discipline
docs/OPERATIONS.md
MCP endpoints ( /pro OAuth and /mcp Bearer key), session timing, and failure handling
.claude/skills/
Agent procedures: trade , review , coach , wiki-distill , ste100
wiki/
Claim-tagged research notes ( findings/ , literature/ ) and index registries
eval/
Your local dataset: funnel-log.jsonl , trades.jsonl , behavior-ledger.jsonl
scripts/
Validators ( funnel_log.py , behavior_log.py , lint.py ) and evaluators ( edge.py , skill_eval.py )
Not Investment Advice
GammaRips is a financial data vendor. Everything in this repository is for educational and research purposes, configured for paper trading by default. Nothing here constitutes investment advice. Trade your own capital at your own risk.
Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. - DevDizzle/gammarips-harness

GitHub - DevDizzle/gammarips-harness: Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice. · GitHub
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
DevDizzle
gammarips-harness Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
24 Commits 24 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude assets assets docs docs eval eval scripts scripts traces traces wiki wiki .gitignore .gitignore .mcp.json .mcp.json CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
An open-source agent workflow harness for the GammaRips MCP curated options dataset.
GammaRips is a data layer, not a signal service. Every trading day, the engine filters ~3,500 optionable US stocks down to a curated cohort of liquid, tradable, high-edge option contracts and serves that dataset over MCP. Your agent pulls the surviving cohort, layers its own live catalyst research (X/Twitter cashtags, filings, web search, parallel subagents), and builds a deterministic trade plan with historical MFE/MAE excursion data.
The Product: Curated Options Dataset via MCP
Instead of dumping noisy raw chains into an LLM context window, GammaRips filters the US options universe down to the daily cohort of surviving contracts that pass strict quantitative and liquidity gates:
Volatility Edge: RV/IV >= 0.85 (realized volatility vs. implied volatility)
Connect any MCP-compatible agent (Claude Code, Grok, Codex, Cursor, or custom runners):
Pro (OAuth): https://mcp.gammarips.com/pro
API Key (Bearer): https://mcp.gammarips.com/mcp with header Authorization: Bearer gr_live_...
Free Preview (No Key): https://mcp.gammarips.com/mcp (serves the daily report, regime context, market calendar, playbooks, and cohort preview)
1. Pull the Daily Cohort ( get_pool )
Fetch today's surviving tradable option contracts and their point-in-time volatility and liquidity metrics ( rv_iv_ratio , open_interest , mid_price , delta , dte , and opportunity surfaces), alongside market calendar status ( get_market_calendar_status ) and VIX term-structure regime ( get_regime_context ).
2. Layer Your Own Catalyst Search (Bring Your Own X / Web Agent)
GammaRips handles the quantitative options filtering; you bring the catalyst layer. Have your agent (Claude Code, Grok, Codex, Cursor) spin up parallel subagents or live X/Twitter and web search — for example, Grok searching live $TICKER cashtag chatter, SEC/FDA regulatory filings, or unpriced catalysts — to debate and rank the surviving contracts.
3. Build the Trade Plan / Execute
Size the position against your risk budget, set deterministic bracket exits (profit target, stop loss, and mandatory time-stop) using historical MFE/MAE excursion distributions ( query_outcomes and replay_contract ), pre-register the decision in eval/funnel-log.jsonl , and optionally route the order to your own broker MCP.
Try the free preview (no key, 10 seconds)
claude mcp add --transport http gammarips https://mcp.gammarips.com/mcp
Ask your agent for today's brief to inspect the market calendar, VIX regime context, daily report, and cohort preview.
Connect with Pro (OAuth or API Key)
Subscribe at gammarips.com/pricing and create an API key at gammarips.com/account (or authenticate directly via OAuth at https://mcp.gammarips.com/pro ).
Export your key for .mcp.json :
export GAMMARIPS_MCP_KEY= " gr_live_... "
Or install as a Claude Code plugin:
/plugin marketplace add DevDizzle/gammarips-harness
/plugin install gammarips@gammarips
Run the daily skills:
/trade Step 1–3 morning workflow: pull cohort, layer catalysts, build exit plan
/review After the close: score every cohort contract with MFE/MAE outcomes
/coach On demand: referee trading behavior against your written rules
Path
Description
CLAUDE.md
Agent instructions: dataset specification, 3-step workflow, and core rules
.mcp.json
Streamable HTTP MCP config ( https://mcp.gammarips.com/mcp with Bearer ${GAMMARIPS_MCP_KEY} )
docs/TRADING-DOCTRINE.md
Risk rules, hard exclusions, sizing framework, and deterministic exit discipline
docs/OPERATIONS.md
MCP endpoints ( /pro OAuth and /mcp Bearer key), session timing, and failure handling
.claude/skills/
Agent procedures: trade , review , coach , wiki-distill , ste100
wiki/
Claim-tagged research notes ( findings/ , literature/ ) and index registries
eval/
Your local dataset: funnel-log.jsonl , trades.jsonl , behavior-ledger.jsonl
scripts/
Validators ( funnel_log.py , behavior_log.py , lint.py ) and evaluators ( edge.py , skill_eval.py )
Not Investment Advice
GammaRips is a financial data vendor. Everything in this repository is for educational and research purposes, configured for paper trading by default. Nothing here constitutes investment advice. Trade your own capital at your own risk.
Open-source GammaRips options-flow trading harness: clone the exact morning workflow (agent picks a contract at run time from live liquidity). Live data via the GammaRips MCP (subscription). Educational, not investment advice.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
