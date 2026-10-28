---
source: "https://github.com/budna-marketplace/recommerce"
hn_url: "https://news.ycombinator.com/item?id=50281334"
title: "AI agent skills for better second-hand decisions"
article_title: "GitHub - budna-marketplace/recommerce: Agent skills for better second-hand decisions. · GitHub"
image: "https://opengraph.githubassets.com/b8b15dce6371375927f2859b162a8e384a0c2a86830595800cb532215f685f24/budna-marketplace/recommerce"
author: "mjovanc"
captured_at: "2026-10-28T22:16:21Z"
capture_tool: "hn-digest"
hn_id: 50281334
score: 2
comments: 0
posted_at: "2026-10-28T21:42:12Z"
tags:
  - hacker-news
---

# AI agent skills for better second-hand decisions

- HN: [50281334](https://news.ycombinator.com/item?id=50281334)
- Source: [github.com](https://github.com/budna-marketplace/recommerce)
- Score: 2
- Comments: 0
- Posted: 2026-10-28T21:42:12Z

## Translation

Title: AI agent skills for better second-hand decisions
Article title: GitHub - budna-marketplace/recommerce: Agent skills for better second-hand decisions. · GitHub
Description: Agent skills for better second-hand decisions. Contribute to budna-marketplace/recommerce development by creating an account on GitHub.

Article text:
GitHub - budna-marketplace/recommerce: Agent skills for better second-hand decisions. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
budna-marketplace
recommerce Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Agent skills for better second-hand decisions.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows checks checks examples examples skills skills .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Agent skills for better second-hand decisions.
Portable agent skills for researching second-hand purchases. Recommerce covers discovery, complete setups, inspection, collection planning, repair-versus-replacement decisions, and reassessment before commitment.
A listing is incomplete, changing evidence. The skills distinguish what a seller claims, what product documentation supports, and what still needs verification. The objective is a suitable purchase at a supported total cost—including the option to keep an existing item.
Skill
Purpose
Output
Buy
Find and compare potential purchases.
A source-backed shortlist with fit, complete costs, tradeoffs, and decisive unknowns.
Kits
Assemble a working setup from multiple items.
A compatible bill of materials, required accessories, alternatives, and one shared budget.
Inspect
Assess an item before purchase or collection.
A listing-specific inspection card separating claims, useful checks, and unresolved condition.
Collect
Determine whether collection is practical.
A route and transport plan with measurements, disassembly constraints, costs, and blockers.
Keep
Compare keeping, repairing, upgrading, or replacing.
A conditional comparison of future costs, suitability, downtime, and uncertainty.
Second Look
Reassess an earlier recommendation before committing.
A change report and updated decision: what changed, what remains supported, and what was not rechecked.
Each skill works independently with evidence the host can access: listings, specifications, measurements, quotes, or a previous assessment. Available read-only tools can supply current evidence; Budna MCP is an optional integration.
After installing, ask your assistant:
Find a complete used cordless drill delivered to Gothenburg for under
1,000 SEK. I need it for furniture assembly and already own suitable bits.
Build a two-person podcast setup from these used listings. I have a laptop
and headphones; include every required connection within 2,000 SEK.
Inspect this drill listing before tomorrow's viewing. Give me the checks
that could change my decision, distinguishing seller claims, model
documentation, and observed condition.
Can I collect this wardrobe in my van? Check both access routes, packed
dimensions, payload, disassembly requirements, and total collection cost.
Compare this laptop repair quote with these used replacements. I need it
for another year; account for downtime and costs still unconfirmed.
Take a second look at yesterday's shortlist before I commit. Recheck the
decisive facts and show what changes the recommendation.
Provide the intended use and constraints you know. The assistant should ask only for missing information that changes the decision. See an illustrative inspection card .
Install directly from GitHub with the Skills CLI :
npx skills add budna-marketplace/recommerce
The installer lets you choose the skills and supported agents. To install only Buy:
npx skills add budna-marketplace/recommerce --skill recommerce-buy
Add --global to make the skill available across projects. No manual clone or repository-level check tooling is needed.
For other Agent Skills-compatible hosts with GitHub import, select the relevant folder in skills/ . Import the complete folder, including its references. Hosts without skill support cannot install these as persistent skills.
Claims retain their sources. Model specifications do not certify a particular used unit; ratings and identity badges do not establish condition.
Compare usable totals. Required accessories, fees, delivery, and collection costs belong in the decision. A current auction bid is not a guaranteed purchase price.
Respect the buyer's constraints. Do not silently relax requirements, invent availability, or fill a shortlist with unsuitable options.
Keep assessment separate from transactions. The skills do not contact sellers, reserve items, bid, or purchase. They do not certify condition, authenticity, or future availability.
These are general decision workflows; examples rely on supplied or verified specifications and quotes. No workflow certifies physical condition, safe handling, repair success, or future availability.
Recommerce Checks contains synthetic regression cases and an offline evaluation workflow. These support development, not a claim of general reliability. See CONTRIBUTING.md for repository structure, development checks, and how to contribute cases or improvements.
Agent skills for better second-hand decisions.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Agent skills for better second-hand decisions. Contribute to budna-marketplace/recommerce development by creating an account on GitHub.

GitHub - budna-marketplace/recommerce: Agent skills for better second-hand decisions. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
budna-marketplace
recommerce Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Agent skills for better second-hand decisions.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows checks checks examples examples skills skills .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Agent skills for better second-hand decisions.
Portable agent skills for researching second-hand purchases. Recommerce covers discovery, complete setups, inspection, collection planning, repair-versus-replacement decisions, and reassessment before commitment.
A listing is incomplete, changing evidence. The skills distinguish what a seller claims, what product documentation supports, and what still needs verification. The objective is a suitable purchase at a supported total cost—including the option to keep an existing item.
Skill
Purpose
Output
Buy
Find and compare potential purchases.
A source-backed shortlist with fit, complete costs, tradeoffs, and decisive unknowns.
Kits
Assemble a working setup from multiple items.
A compatible bill of materials, required accessories, alternatives, and one shared budget.
Inspect
Assess an item before purchase or collection.
A listing-specific inspection card separating claims, useful checks, and unresolved condition.
Collect
Determine whether collection is practical.
A route and transport plan with measurements, disassembly constraints, costs, and blockers.
Keep
Compare keeping, repairing, upgrading, or replacing.
A conditional comparison of future costs, suitability, downtime, and uncertainty.
Second Look
Reassess an earlier recommendation before committing.
A change report and updated decision: what changed, what remains supported, and what was not rechecked.
Each skill works independently with evidence the host can access: listings, specifications, measurements, quotes, or a previous assessment. Available read-only tools can supply current evidence; Budna MCP is an optional integration.
After installing, ask your assistant:
Find a complete used cordless drill delivered to Gothenburg for under
1,000 SEK. I need it for furniture assembly and already own suitable bits.
Build a two-person podcast setup from these used listings. I have a laptop
and headphones; include every required connection within 2,000 SEK.
Inspect this drill listing before tomorrow's viewing. Give me the checks
that could change my decision, distinguishing seller claims, model
documentation, and observed condition.
Can I collect this wardrobe in my van? Check both access routes, packed
dimensions, payload, disassembly requirements, and total collection cost.
Compare this laptop repair quote with these used replacements. I need it
for another year; account for downtime and costs still unconfirmed.
Take a second look at yesterday's shortlist before I commit. Recheck the
decisive facts and show what changes the recommendation.
Provide the intended use and constraints you know. The assistant should ask only for missing information that changes the decision. See an illustrative inspection card .
Install directly from GitHub with the Skills CLI :
npx skills add budna-marketplace/recommerce
The installer lets you choose the skills and supported agents. To install only Buy:
npx skills add budna-marketplace/recommerce --skill recommerce-buy
Add --global to make the skill available across projects. No manual clone or repository-level check tooling is needed.
For other Agent Skills-compatible hosts with GitHub import, select the relevant folder in skills/ . Import the complete folder, including its references. Hosts without skill support cannot install these as persistent skills.
Claims retain their sources. Model specifications do not certify a particular used unit; ratings and identity badges do not establish condition.
Compare usable totals. Required accessories, fees, delivery, and collection costs belong in the decision. A current auction bid is not a guaranteed purchase price.
Respect the buyer's constraints. Do not silently relax requirements, invent availability, or fill a shortlist with unsuitable options.
Keep assessment separate from transactions. The skills do not contact sellers, reserve items, bid, or purchase. They do not certify condition, authenticity, or future availability.
These are general decision workflows; examples rely on supplied or verified specifications and quotes. No workflow certifies physical condition, safe handling, repair success, or future availability.
Recommerce Checks contains synthetic regression cases and an offline evaluation workflow. These support development, not a claim of general reliability. See CONTRIBUTING.md for repository structure, development checks, and how to contribute cases or improvements.
Agent skills for better second-hand decisions.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
