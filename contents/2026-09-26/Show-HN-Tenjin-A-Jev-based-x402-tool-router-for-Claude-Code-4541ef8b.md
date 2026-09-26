---
source: "https://github.com/BackTrackCo/tenjin-agent"
hn_url: "https://news.ycombinator.com/item?id=49851853"
title: "Show HN: Tenjin – A Jev based x402 tool router for Claude Code"
article_title: "GitHub - BackTrackCo/tenjin-agent: A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. · GitHub"
image: "https://opengraph.githubassets.com/cbb6526d4394cd964d41fad41ea5d0cd77ef84e6667983f10e80f1ced87f5710/BackTrackCo/tenjin-agent"
author: "AliAbdoli"
captured_at: "2026-09-26T00:31:18Z"
capture_tool: "hn-digest"
hn_id: 49851853
score: 1
comments: 0
posted_at: "2026-09-26T00:25:57Z"
tags:
  - hacker-news
---

# Show HN: Tenjin – A Jev based x402 tool router for Claude Code

- HN: [49851853](https://news.ycombinator.com/item?id=49851853)
- Source: [github.com](https://github.com/BackTrackCo/tenjin-agent)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T00:25:57Z

## Translation

Title: Show HN: Tenjin – A Jev based x402 tool router for Claude Code
Article title: GitHub - BackTrackCo/tenjin-agent: A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. · GitHub
Description: A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. - BackTrackCo/tenjin-agent
HN text: Hey HN! We're Ali And Vrajang. We just created a tool router for Claude Code built on Jev. It basically gives your agent superpowers. You keep working like you always do. When Jev thinks a tool would help, it provides it to your agent and your agent uses it. It watches your masked prompts, session context, web searches and subagent tasks, so it knows exactly when to intervene. For now, tools come from a curated library. It's the basics like Exa, Firecrawl, Context7, but we're planning to add more soon. Let us know what you'd like added! Everything runs on x402, so you don't set up each tool or juggle API keys. Your wallet pays a cent or two per call and that's it :) Please expect some rough edges as it is just in alpha for now. Would love to know what you think! Here is our website: https://tenjin.sh and GitHub repo: https://github.com/BackTrackCo/tenjin-agent for those interested.

Article text:
GitHub - BackTrackCo/tenjin-agent: A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. · GitHub
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
BackTrackCo
/
tenjin-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
193 Commits 193 Commits Folders and files
.changeset .changeset .githooks .githooks .github/ workflows .github/ workflows .greptile .greptile assets assets docs docs evals evals scripts scripts skills skills src src .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR_LICENSE_AGREEMENT.md CONTRIBUTOR_LICENSE_AGREEMENT.md NOTICE.md NOTICE.md README.md README.md RELEASING.md RELEASING.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
The tool router for coding agents.
Give your agent superpowers once. Tenjin hands it the right tool when it helps, and you keep working like before.
No API keys. No pile of MCP servers. No rules to write.
Quick start ·
Tools ·
Wallet & payments ·
Request a tool
Giving your agent good tools is a chore today:
A key for every tool. Sign up, pick a plan, paste an API key into a config file. Five tools, five accounts.
MCP servers crowd the context. Every server you install loads its tool definitions into every session, needed or not.
Your agent forgets anyway. It reaches for plain web search out of habit, so you write rules to remind it, and it still slips.
Some tools you need once. Installing something permanent for a one-off lookup isn't worth the setup.
Tenjin handles all of it. Install it once and keep working. The Tenjin router watches the moments where a tool could help: your prompt, your agent's web searches and page fetches, and the tasks it hands to subagents. When a curated tool beats what your agent was about to do, the router suggests it and your agent calls it. Your wallet pays for each call through x402 , so you manage no keys and install nothing new.
Requires Node.js 24 or newer and Claude Code . Codex support is on the way.
npm i -g tenjin-cli
tenjin install # sets up Claude Code and creates your wallet
tenjin wallet fund 2 # optional: add $2 with a card, via Coinbase
✓ Tenjin is set up for Claude Code
✓ Wallet created: 0x3c0D84055994c3062819Ce8730869D0aDeA4c3Bf
Automatic router: up to $0.25 per call; daily limit $5 a day
Next: tenjin wallet fund, then restart Claude Code
Restart Claude Code. That's it.
> What are BTC and ETH trading at?
> How do I paginate list results with the Stripe Node SDK?
> Is ada@example.com a deliverable address?
> Integrate x^2 sin(x) dx from 0 to pi.
Each answer names the provider and the price. Your agent keeps its own tools, and the router steps in only when it has something better.
The router picks from a catalog we curate and maintain. Routing is free: you pay the provider's price and nothing else.
Free tools run without a funded wallet. Twitter, Reddit and more are next, added by demand: tell us what you want .
flowchart LR
A["Your prompt,<br>a web search,<br>or a subagent task"] --> B{"Tenjin<br>router"}
B -- "a tool fits" --> C["Agent calls it<br>through x402"]
C --> D["Wallet pays the provider,<br>result comes back"]
B -- "nothing better" --> E["Agent carries on"]
Loading
The Tenjin router hooks into Claude Code at your prompt, before each web search or page fetch, and when your agent hands work to a subagent. At each point it asks Jev, a decision model from TypeSafe , whether a tool in the catalog fits. Jev can only choose from that fixed list, so it never writes a call or an instruction for your agent.
When one fits, your agent sees a one-line suggestion with the tool and its price, and calls it through Tenjin's x402 MCP server. Tenjin pays the provider from your wallet, within your limits, and hands back the result. Your status line shows the call as it happens:
x402 · request: calling pro-api.coinmarketcap.com/x402/v3/cryptocurrency/quotes/latest · {"query":{"symbol":"BTC,ETH"}}
To pick a tool, the router sends your current turn and up to six recent messages, with keys, passwords and seed phrases masked. Tool results and page contents stay on your machine, and the packet expires after 15 minutes. tenjin config set router.context turn sends only the current message. Full details, and everything install writes →
Tenjin pays for tools with x402 , an open standard that builds payments into HTTP. A paid endpoint answers 402 Payment Required with its price, the client signs a payment for that amount, and the endpoint returns the result. Neither side needs an account or an API key. Read more: x402.org · whitepaper · Coinbase docs · spec and SDKs .
tenjin install creates a wallet on your machine that holds USDC on Base . Tenjin encrypts the private key, unlocks it through your OS keychain, and never sends it anywhere. Tenjin never holds or moves your funds. Paying a provider costs no gas.
Limit
Default
Change it
Per lookup
$0.25
tenjin config set maxAutoSpend 0.10
Per day
$5
tenjin config set sessionBudget 2
Tenjin refuses any payment over either limit before it signs anything.
tenjin wallet fund 2 opens a Coinbase Onramp checkout for your wallet: pay by card, or Apple Pay where your region supports it. Sign in to Coinbase or create an account during checkout. You can also send USDC on Base to the address tenjin wallet show prints.
$1–2 goes a long way. $2 covers about 280 web searches, 200 page reads or 100 Wolfram Alpha answers.
tenjin wallet send <amount> USDC <address> sends funds to any address. It's a regular onchain transfer, so it needs a little ETH on Base for gas.
tenjin status # what you've spent today
tenjin wallet balance # what's left
tenjin wallet fund 5 # top up
tenjin update # newest version; wallet and settings stay
tenjin config set router.enabled false # pause the router on this machine
tenjin config set --project router.enabled false # ...or just in this repo
tenjin uninstall # remove the Claude Code setup; your wallet stays
Use tenjin install --project to set it up for a single project. Add --json to any command for machine-readable output.
tenjin install sets Claude Code's status line only if you don't already have one. If you do, yours is left alone and the install prints a command that shows both; tenjin install --status-line compose writes it for you, and --status-line skip leaves the setting untouched. The status line reads local files only: no network, no wallet, and it can't affect a lookup or a payment.
Limits are set only where no setting exists yet, so an update never overwrites yours. The daily budget is a rolling 24-hour window. Setting either limit to 0 blocks automatic payments; tenjin config set sessionBudget none removes the daily ceiling only.
tenjin pay <url> pays any x402 endpoint you name. It ignores the automatic limits and always asks for your consent to the quoted price, interactively or with --yes . It is never used as a workaround when the router refuses. See the safety model .
Updating to 0.1.0-alpha.18 could accidentally register the server in ~/.mcp.json while refreshing a user install. After upgrading to a release with the fix, check that file. If its x402 entry runs tenjin mcp and you didn't install at project scope in your home directory on purpose, remove just that registration:
(cd ~ && claude mcp remove x402 -s project)
This keeps your other MCP entries and the user registration in ~/.claude.json . Run tenjin install , then restart Claude Code. Don't use tenjin uninstall --project from home for this: its settings path is also the user settings path.
When refreshing from home, Tenjin keeps the existing user registration if there is one, or preserves a project-only registration; with neither, it defaults to user scope. tenjin install --refresh --project selects project scope explicitly.
Tenjin is in alpha, and the catalog grows with what people ask for. Missing a service your agent keeps needing? Hit a lookup that went to the wrong place? We want to hear it.
How a lookup runs, what it sends, and what install writes
Before contributing, read CONTRIBUTING.md for the contributor agreement.
pnpm install
pnpm run githooks # once, to use the repo's git hooks
pnpm run build
pnpm run test
pnpm run typecheck
pnpm run lint
About
A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts.
Contributing Activity Custom properties Stars
2 forks Report repository Used by
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. - BackTrackCo/tenjin-agent

Hey HN! We're Ali And Vrajang. We just created a tool router for Claude Code built on Jev. It basically gives your agent superpowers. You keep working like you always do. When Jev thinks a tool would help, it provides it to your agent and your agent uses it. It watches your masked prompts, session context, web searches and subagent tasks, so it knows exactly when to intervene. For now, tools come from a curated library. It's the basics like Exa, Firecrawl, Context7, but we're planning to add more soon. Let us know what you'd like added! Everything runs on x402, so you don't set up each tool or juggle API keys. Your wallet pays a cent or two per call and that's it :) Please expect some rough edges as it is just in alpha for now. Would love to know what you think! Here is our website: https://tenjin.sh and GitHub repo: https://github.com/BackTrackCo/tenjin-agent for those interested.

GitHub - BackTrackCo/tenjin-agent: A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts. · GitHub
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
BackTrackCo
/
tenjin-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
193 Commits 193 Commits Folders and files
.changeset .changeset .githooks .githooks .github/ workflows .github/ workflows .greptile .greptile assets assets docs docs evals evals scripts scripts skills skills src src .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR_LICENSE_AGREEMENT.md CONTRIBUTOR_LICENSE_AGREEMENT.md NOTICE.md NOTICE.md README.md README.md RELEASING.md RELEASING.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
The tool router for coding agents.
Give your agent superpowers once. Tenjin hands it the right tool when it helps, and you keep working like before.
No API keys. No pile of MCP servers. No rules to write.
Quick start ·
Tools ·
Wallet & payments ·
Request a tool
Giving your agent good tools is a chore today:
A key for every tool. Sign up, pick a plan, paste an API key into a config file. Five tools, five accounts.
MCP servers crowd the context. Every server you install loads its tool definitions into every session, needed or not.
Your agent forgets anyway. It reaches for plain web search out of habit, so you write rules to remind it, and it still slips.
Some tools you need once. Installing something permanent for a one-off lookup isn't worth the setup.
Tenjin handles all of it. Install it once and keep working. The Tenjin router watches the moments where a tool could help: your prompt, your agent's web searches and page fetches, and the tasks it hands to subagents. When a curated tool beats what your agent was about to do, the router suggests it and your agent calls it. Your wallet pays for each call through x402 , so you manage no keys and install nothing new.
Requires Node.js 24 or newer and Claude Code . Codex support is on the way.
npm i -g tenjin-cli
tenjin install # sets up Claude Code and creates your wallet
tenjin wallet fund 2 # optional: add $2 with a card, via Coinbase
✓ Tenjin is set up for Claude Code
✓ Wallet created: 0x3c0D84055994c3062819Ce8730869D0aDeA4c3Bf
Automatic router: up to $0.25 per call; daily limit $5 a day
Next: tenjin wallet fund, then restart Claude Code
Restart Claude Code. That's it.
> What are BTC and ETH trading at?
> How do I paginate list results with the Stripe Node SDK?
> Is ada@example.com a deliverable address?
> Integrate x^2 sin(x) dx from 0 to pi.
Each answer names the provider and the price. Your agent keeps its own tools, and the router steps in only when it has something better.
The router picks from a catalog we curate and maintain. Routing is free: you pay the provider's price and nothing else.
Free tools run without a funded wallet. Twitter, Reddit and more are next, added by demand: tell us what you want .
flowchart LR
A["Your prompt,<br>a web search,<br>or a subagent task"] --> B{"Tenjin<br>router"}
B -- "a tool fits" --> C["Agent calls it<br>through x402"]
C --> D["Wallet pays the provider,<br>result comes back"]
B -- "nothing better" --> E["Agent carries on"]
Loading
The Tenjin router hooks into Claude Code at your prompt, before each web search or page fetch, and when your agent hands work to a subagent. At each point it asks Jev, a decision model from TypeSafe , whether a tool in the catalog fits. Jev can only choose from that fixed list, so it never writes a call or an instruction for your agent.
When one fits, your agent sees a one-line suggestion with the tool and its price, and calls it through Tenjin's x402 MCP server. Tenjin pays the provider from your wallet, within your limits, and hands back the result. Your status line shows the call as it happens:
x402 · request: calling pro-api.coinmarketcap.com/x402/v3/cryptocurrency/quotes/latest · {"query":{"symbol":"BTC,ETH"}}
To pick a tool, the router sends your current turn and up to six recent messages, with keys, passwords and seed phrases masked. Tool results and page contents stay on your machine, and the packet expires after 15 minutes. tenjin config set router.context turn sends only the current message. Full details, and everything install writes →
Tenjin pays for tools with x402 , an open standard that builds payments into HTTP. A paid endpoint answers 402 Payment Required with its price, the client signs a payment for that amount, and the endpoint returns the result. Neither side needs an account or an API key. Read more: x402.org · whitepaper · Coinbase docs · spec and SDKs .
tenjin install creates a wallet on your machine that holds USDC on Base . Tenjin encrypts the private key, unlocks it through your OS keychain, and never sends it anywhere. Tenjin never holds or moves your funds. Paying a provider costs no gas.
Limit
Default
Change it
Per lookup
$0.25
tenjin config set maxAutoSpend 0.10
Per day
$5
tenjin config set sessionBudget 2
Tenjin refuses any payment over either limit before it signs anything.
tenjin wallet fund 2 opens a Coinbase Onramp checkout for your wallet: pay by card, or Apple Pay where your region supports it. Sign in to Coinbase or create an account during checkout. You can also send USDC on Base to the address tenjin wallet show prints.
$1–2 goes a long way. $2 covers about 280 web searches, 200 page reads or 100 Wolfram Alpha answers.
tenjin wallet send <amount> USDC <address> sends funds to any address. It's a regular onchain transfer, so it needs a little ETH on Base for gas.
tenjin status # what you've spent today
tenjin wallet balance # what's left
tenjin wallet fund 5 # top up
tenjin update # newest version; wallet and settings stay
tenjin config set router.enabled false # pause the router on this machine
tenjin config set --project router.enabled false # ...or just in this repo
tenjin uninstall # remove the Claude Code setup; your wallet stays
Use tenjin install --project to set it up for a single project. Add --json to any command for machine-readable output.
tenjin install sets Claude Code's status line only if you don't already have one. If you do, yours is left alone and the install prints a command that shows both; tenjin install --status-line compose writes it for you, and --status-line skip leaves the setting untouched. The status line reads local files only: no network, no wallet, and it can't affect a lookup or a payment.
Limits are set only where no setting exists yet, so an update never overwrites yours. The daily budget is a rolling 24-hour window. Setting either limit to 0 blocks automatic payments; tenjin config set sessionBudget none removes the daily ceiling only.
tenjin pay <url> pays any x402 endpoint you name. It ignores the automatic limits and always asks for your consent to the quoted price, interactively or with --yes . It is never used as a workaround when the router refuses. See the safety model .
Updating to 0.1.0-alpha.18 could accidentally register the server in ~/.mcp.json while refreshing a user install. After upgrading to a release with the fix, check that file. If its x402 entry runs tenjin mcp and you didn't install at project scope in your home directory on purpose, remove just that registration:
(cd ~ && claude mcp remove x402 -s project)
This keeps your other MCP entries and the user registration in ~/.claude.json . Run tenjin install , then restart Claude Code. Don't use tenjin uninstall --project from home for this: its settings path is also the user settings path.
When refreshing from home, Tenjin keeps the existing user registration if there is one, or preserves a project-only registration; with neither, it defaults to user scope. tenjin install --refresh --project selects project scope explicitly.
Tenjin is in alpha, and the catalog grows with what people ask for. Missing a service your agent keeps needing? Hit a lookup that went to the wrong place? We want to hear it.
How a lookup runs, what it sends, and what install writes
Before contributing, read CONTRIBUTING.md for the contributor agreement.
pnpm install
pnpm run githooks # once, to use the repo's git hooks
pnpm run build
pnpm run test
pnpm run typecheck
pnpm run lint
About
A router for Claude Code that finds and pays for research, data, and tools. Powered by x402 and USDC on Base, no API keys or accounts.
Contributing Activity Custom properties Stars
2 forks Report repository Used by
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
