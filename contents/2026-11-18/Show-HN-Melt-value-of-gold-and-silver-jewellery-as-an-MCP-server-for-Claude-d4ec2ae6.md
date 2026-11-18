---
source: "https://github.com/srinivasalokeshraju/Jewellery_value"
hn_url: "https://news.ycombinator.com/item?id=50547304"
title: "Show HN: Melt value of gold and silver jewellery, as an MCP server for Claude"
article_title: "GitHub - srinivasalokeshraju/Jewellery_value · GitHub"
image: "https://opengraph.githubassets.com/e363671a99e57d034da8f5b4bfff14bc698a617b3b9521e1061a7294f5712b86/srinivasalokeshraju/Jewellery_value"
author: "lokessri"
captured_at: "2026-11-18T16:14:21Z"
capture_tool: "hn-digest"
hn_id: 50547304
score: 1
comments: 0
posted_at: "2026-11-18T15:15:16Z"
tags:
  - hacker-news
---

# Show HN: Melt value of gold and silver jewellery, as an MCP server for Claude

- HN: [50547304](https://news.ycombinator.com/item?id=50547304)
- Source: [github.com](https://github.com/srinivasalokeshraju/Jewellery_value)
- Score: 1
- Comments: 0
- Posted: 2026-11-18T15:15:16Z

## Translation

Title: Show HN: Melt value of gold and silver jewellery, as an MCP server for Claude
Article title: GitHub - srinivasalokeshraju/Jewellery_value · GitHub
Description: Contribute to srinivasalokeshraju/Jewellery_value development by creating an account on GitHub.

Article text:
GitHub - srinivasalokeshraju/Jewellery_value · GitHub
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
srinivasalokeshraju
Jewellery_value Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
98 Commits 98 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .github .github docs/ plans docs/ plans plugin-dev plugin-dev plugin plugin server server .editorconfig .editorconfig .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .oxlintrc.json .oxlintrc.json .prettierignore .prettierignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Estimate the melt value of gold and silver jewellery in your own currency. A remote MCP server computes every number from the live spot price and exchange rate, and names the source and time of each rate. The model computes none.
/plugin install melt-value --marketplace srinivasalokeshraju/Jewellery_value
The plugin adds a guided valuation flow, a hallmark reference, and the /melt-value , /gold-rate , and /silver-rate commands. See plugin/README.md .
https://mcp.meltvalue.dev/mcp
It uses the streamable HTTP transport and needs no account or API key. It offers four tools: calc_melt_value , get_metal_rate , convert_weight , and list_supported .
You: I have a 22K gold chain, 20 grams, in Hyderabad. What is it worth?
22K gold chain, 20 g
Pure gold content 18.33 g
Melt value (spot) 209409.69 INR spot at 18 Nov 2026, 05:25 UTC, gold-api.com
Likely offer from buyer 146586.79 to 198939.21 INR (typical 70 to 95% of melt)
Warnings:
- No weight was given for stones or other non-metal parts, so the full weight counts as metal and the value may be too high.
Assumed: purity assumed from what you told me; weight from your scale; no stones deducted.
Not included: making charges, stones, brand or antique value.
This is an estimate from the stated weight and purity, not an appraisal or an offer to buy.
The figures above came from the hosted server on 18 Nov 2026, with gold at 3795.67 USD per troy ounce and 93.6 INR per USD. Your result uses the rates at the time you ask.
Why the number differs from a shop rate
Melt value is the international spot price of the pure metal, converted to your currency. It leaves out import duty, sales tax, and the retail margin. A shop's buying or selling rate usually includes some of these, so it is often higher than melt value. In India, shop rates on 17–18 Nov 2026 ran 19–24% above melt value. The gap differs by country and changes over time.
Every result is an estimate from the stated weight and purity, not an appraisal or an offer to buy.
PRIVACY.md : what the server receives, logs, and shares.
server/AGENTS.md : how the server is built, for contributors.
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to srinivasalokeshraju/Jewellery_value development by creating an account on GitHub.

GitHub - srinivasalokeshraju/Jewellery_value · GitHub
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
srinivasalokeshraju
Jewellery_value Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
98 Commits 98 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .github .github docs/ plans docs/ plans plugin-dev plugin-dev plugin plugin server server .editorconfig .editorconfig .git-blame-ignore-revs .git-blame-ignore-revs .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .oxlintrc.json .oxlintrc.json .prettierignore .prettierignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Estimate the melt value of gold and silver jewellery in your own currency. A remote MCP server computes every number from the live spot price and exchange rate, and names the source and time of each rate. The model computes none.
/plugin install melt-value --marketplace srinivasalokeshraju/Jewellery_value
The plugin adds a guided valuation flow, a hallmark reference, and the /melt-value , /gold-rate , and /silver-rate commands. See plugin/README.md .
https://mcp.meltvalue.dev/mcp
It uses the streamable HTTP transport and needs no account or API key. It offers four tools: calc_melt_value , get_metal_rate , convert_weight , and list_supported .
You: I have a 22K gold chain, 20 grams, in Hyderabad. What is it worth?
22K gold chain, 20 g
Pure gold content 18.33 g
Melt value (spot) 209409.69 INR spot at 18 Nov 2026, 05:25 UTC, gold-api.com
Likely offer from buyer 146586.79 to 198939.21 INR (typical 70 to 95% of melt)
Warnings:
- No weight was given for stones or other non-metal parts, so the full weight counts as metal and the value may be too high.
Assumed: purity assumed from what you told me; weight from your scale; no stones deducted.
Not included: making charges, stones, brand or antique value.
This is an estimate from the stated weight and purity, not an appraisal or an offer to buy.
The figures above came from the hosted server on 18 Nov 2026, with gold at 3795.67 USD per troy ounce and 93.6 INR per USD. Your result uses the rates at the time you ask.
Why the number differs from a shop rate
Melt value is the international spot price of the pure metal, converted to your currency. It leaves out import duty, sales tax, and the retail margin. A shop's buying or selling rate usually includes some of these, so it is often higher than melt value. In India, shop rates on 17–18 Nov 2026 ran 19–24% above melt value. The gap differs by country and changes over time.
Every result is an estimate from the stated weight and purity, not an appraisal or an offer to buy.
PRIVACY.md : what the server receives, logs, and shares.
server/AGENTS.md : how the server is built, for contributors.
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
