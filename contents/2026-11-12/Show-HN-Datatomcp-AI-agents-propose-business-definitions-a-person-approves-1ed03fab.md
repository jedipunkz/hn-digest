---
source: "https://datatomcp.com/"
hn_url: "https://news.ycombinator.com/item?id=50467571"
title: "Show HN: Datatomcp – AI agents propose business definitions, a person approves"
article_title: "datatomcp · The AI-native data layer for teams without a data team"
image: "https://datatomcp.com/og.png"
author: "Killian1603"
captured_at: "2026-11-12T15:01:40Z"
capture_tool: "hn-digest"
hn_id: 50467571
score: 3
comments: 0
posted_at: "2026-11-12T14:16:53Z"
tags:
  - hacker-news
---

# Show HN: Datatomcp – AI agents propose business definitions, a person approves

- HN: [50467571](https://news.ycombinator.com/item?id=50467571)
- Source: [datatomcp.com](https://datatomcp.com/)
- Score: 3
- Comments: 0
- Posted: 2026-11-12T14:16:53Z

## Translation

Title: Show HN: Datatomcp – AI agents propose business definitions, a person approves
Article title: datatomcp · The AI-native data layer for teams without a data team
Description: Make your team

Article text:
datatomcp Product Live demo Pricing Guides FAQ Sign in Start free For humans For agents > @agents, connect datatomcp
friends with Claude Claude Code Cursor Codex ChatGPT (soon) any MCP client
install.md how to connect datatomcp in your app
index.md this page as Markdown
llms.txt facts, docs and guides
llms-full.txt every guide in one file
mcp the MCP server (Streamable HTTP, OAuth)
# datatomcp: your team's data, ready for every AI agent
> datatomcp is the AI-native data layer for teams without a data team. It makes a team's Google Sheets, CSV and Excel files, databases (PostgreSQL, MySQL) and apps (Stripe, HubSpot, Airtable) AI-ready: one governed place every AI agent can ask, with exact answers computed by a database, shared definitions of what each number means, a full audit trail of agent questions, charts and dashboards. Agents connect over a remote MCP server (OAuth) or a REST API. Hosted in the EU.
Source: https://datatomcp.com/
## How it works
1. Connect: upload CSV or Excel files, pick Google Sheets, or connect PostgreSQL, MySQL, Stripe, HubSpot or Airtable. Each source becomes a dataset with versions and sync, and an import report of anything odd.
2. Link: datatomcp finds the ids the datasets share (orders.customer_id and customers.customer_id) and links them, so agents answer across datasets in one query.
3. Define: approve the definitions datatomcp drafts from your columns, and add your team's words to the glossary.
4. Ask: add the MCP URL https://app.datatomcp.com/mcp to Claude, ChatGPT, Cursor or any MCP client, or call the REST API. Answers are counted by a database, name the dataset version, and can come back as a chart or a dashboard.
5. Audit: every agent question is recorded, with repeated questions compared.
## Facts
- What it is: an AI-native data layer (a hosted data warehouse plus semantic layer for teams without a data team). Not a spreadsheet add-on, not a BI tool, not a SQL database.
- Who it is for: teams that work from spreadsheets and use several AI agents (operations, finance, revenue, product, agencies), and the consultants who set this up for clients.
- MCP server: https://app.datatomcp.com/mcp (Streamable HTTP, OAuth sign-in, no API key)
- MCP tools: workspace_overview, list_datasets, describe_dataset, distinct_values, query_rows, aggregate (read-only data tools); chart (draws the answer in the chat, MCP Apps vi
[truncated]
Your team's data, ready for every AI agent.
Connect your databases, apps and sheets once. datatomcp links them, writes down what each number means, and lets every agent on your team ask for exact answers. Every question is recorded, every number can be checked, and the answer can come back as a chart.
Watch the 90-second demo Free for 3 datasets and 2 people · No card needed · Read-only for agents · Hosted in the EU
Click Connect, sign in to datatomcp, then ask about your data in any chat. Desktop and mobile use the same connector.
Claude Code Paste in Claude Code Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Sound on Works with any AI agent that supports MCP, and anything that speaks REST. Listed where agents look:
Model Context Protocol MCP Registry
Don't send the sheet. Send the question.
Pasting a sheet puts every row into the model's context: it fills the window, and you pay for every token. With datatomcp the agent asks a database and gets back only the result, about 475 times fewer tokens on our test file, and the total is computed, not estimated.
How many orders came from France in September, and what was the revenue?
Each square is 1,000 tokens of Claude Opus 5.5's 1M-token context window.
Measured on our 50,000-row test orders file. Pasting sends the whole 4,118,185-byte file; a full datatomcp session (tool list, schema, query and result) sends 8,663 bytes. Tokens estimated at about 4 bytes per token. Price: Claude Opus 5.5 input, $4 per million tokens (October 2026).
Connect it once. Every agent gets the same answer.
Five steps take your team from scattered sheets, exports and databases to answers you can put in a report without checking them by hand.
Connect PostgreSQL, MySQL, Stripe, HubSpot and Airtable, pick Google Sheets in Google's own picker, or upload Excel and CSV files. Each source becomes a dataset that stays in sync, with a version for every refresh.
Every import comes with a report: lines that couldn't be read, numbers stored as text, mixed dates. Nothing is converted until someone confirms it.
Add data Northwind Studio PostgreSQL
orders PostgreSQL 48 rows Importing Ready
customers Google Sheets 10 rows Ready
sample_sales CSV 1,500 rows 1 finding
revenue looks like numbers but is stored as text, like "$712.80" . Nothing changes until you confirm.
datatomcp finds the ids and codes your datasets share, like the customer_id in orders and in customers, and draws the map. Links chain, so an agent can total orders by a customer's region in one query.
A link confirms itself only when every value matches. Otherwise you see what didn't match, and decide.
Map 3 datasets · 2 links orders 48 order_id customer_id amount status customers 10 customer_id segment region_id regions 3 region_id region orders → customers 47 of 48 matched Confirm Confirmed
customers → regions 10 of 10 matched Confirmed on its own
group_by: ["regions.region"] works from orders now
Agree on what each number means
datatomcp drafts definitions from your columns, like paid revenue as the sum of amount where the status is paid. A person approves or edits each one, and agents only use approved ones.
Every answer that uses a definition says so. A glossary explains your team's words, like "active customer", the same way.
Definitions orders paid_revenue Draft Approved sum(amount) where status = "paid"
Revenue from paid orders; refunds excluded.
Also answers to revenue and sales
Approved by you · agents use it from now on Glossary
Add one URL to Claude, ChatGPT, Cursor or any MCP client, or call the REST API. A database counts every row, so totals are exact, and each answer names the dataset version it came from.
When a picture is clearer, the agent draws a chart or makes a dashboard right in the chat, and anyone can refresh it later.
Claude ChatGPT Cursor datatomcp connected What's our paid revenue by region?
aggregate · orders · paid_revenue by regions.region 190 ms
Paid revenue is 24,360 , from your team's approved paid_revenue definition.
orders, version 1 · 1 order has no matching customer · Check this number
Every agent question is kept: who asked, from which agent, on which dataset, the query, the time it took and any error.
When two agents ask the same question, you see whether they got the same answer. Keep 90 days on Pro, and export it as a spreadsheet.
Audit 72 calls · 4 agents · 3 people · 7 days Maya · Claude 190 ms 4 groups aggregate · orders · paid_revenue by regions.region
Sam · ChatGPT 160 ms 4 groups Same answer as Claude aggregate · orders · paid_revenue by regions.region
You · Cursor 120 ms 8 rows query_rows · orders · status = refunded
Sam · ChatGPT 140 ms Error: revenue is text aggregate · sample_sales · sum(revenue)
You · Claude Code 210 ms 8 columns describe_dataset · sample_sales
I help about 300 finance people use AI agents at a Fortune 100 company. Almost all of them hit the same wall: they paste a whole CSV or Google Sheet into the chat, the context fills up, the tokens add up, and the totals are close but not exact.
datatomcp is the tool I wanted for them: read-only for agents, hosted in the EU, and honest about its limits.
One workspace for the whole team. Data, meaning and every agent, side by side.
Every dataset, its freshness, the team's definitions and what the agents asked, on one page. Open this page live Which agents asked, which definitions they used, which datasets those read, and how the datasets link. Open this page live Approved definitions with who approved them, drafts waiting for a person, and the team's glossary. Open this page live Every agent question, repeated questions and whether they got the same answer, and every error. Open this page live A board an agent made from a few questions, refreshed on the latest data whenever you open it. This is the public demo workspace, Northwind Studio, with its real numbers. Click through every page, no sign-up needed.
Made for teams without a data team. And for the people who set them up.
Invite people as editors, who add and update data, or viewers, who ask. Each person connects their own agent once and gets the same datasets at the same version, so everyone quotes the same number.
Members Invite people N Northwind Studio Owner
Give each client a workspace of their own, with their data, their definitions and their words. Limit the agent you use for a client to that client's data, show them what their agents asked, then hand it over.
Workspaces New workspace N Northwind Studio 4 datasets · 3 people Current
+ Your next client Its own data, definitions and team
Claude Code Northwind Studio only
Start free, no card needed. Upgrade in the app when your data or your team outgrow it, and cancel anytime.
Sync within minutes of an edit
Agents read through validated, read-only tools, and only the results they ask for. Everything else stays with us, in the EU, until you delete it.
Hosted in the EU A copy of each source (one file per version), backups and accounts are kept in the European Union. No AI processing unless you opt in Never used for training. AI cleaning suggestions (Mistral, EU) and the in-app assistant (your own key) stay off until you turn them on. Private by default Every request needs a sign-in or an API key. Database and app credentials are encrypted and never shown to agents. Deleted when you say so Delete a dataset or your whole account from the app. It's gone at once, and from backups within 14 days. Questions
Anything else? Write to [email protected] , the founder answers.
A single place where your team's spreadsheets, files, databases and apps live, kept up to date, linked together and described with the team's own definitions, that every AI agent can ask for exact answers. Agents don't just reach the data: they know what each number means, a database does the counting, and every question is recorded. It does the job of a data warehouse and a semantic layer for teams without a data team.
Databases and apps: PostgreSQL and MySQL (a table or a read-only query), Stripe (customers, subscriptions with MRR, invoices, payments), HubSpot (contacts, companies, deals) and Airtable. Sheets and files: Google Sheets, Excel and CSV. Each one becomes a dataset with versions, sync, links to your other data and the audit. Need another source? Ask for it in the app: we add sources on request, the most requested first.
The team's agreed meaning of a number, like "paid revenue = sum of amount where status is paid". datatomcp drafts them from your columns when you add data, and a person approves or edits each one. Agents only see approved definitions, every answer that uses one says who approved it, and an unknown name gets the closest definitions instead of a guess. A glossary explains words like "active customer" the same way.
Any client that supports remote MCP servers, including Claude, ChatGPT, Cursor and Claude Code. Agents without MCP support, n8n workflows and your own apps use the REST API with an API key. Both return the same results.
Over MCP: workspace_overview, list_datasets, describe_dataset, distinct_values, query_rows and aggregate to read and count, chart to draw the answer, create_dashboard, list_dashboards and refresh_dashboard for dashboards, report_issue to flag a data problem to the team, and propose_definition, propose_link, propose_term and propose_brief to suggest a definition, a link, a glossary term or the busin

[truncated]

## Original Extract

Make your team

datatomcp Product Live demo Pricing Guides FAQ Sign in Start free For humans For agents > @agents, connect datatomcp
friends with Claude Claude Code Cursor Codex ChatGPT (soon) any MCP client
install.md how to connect datatomcp in your app
index.md this page as Markdown
llms.txt facts, docs and guides
llms-full.txt every guide in one file
mcp the MCP server (Streamable HTTP, OAuth)
# datatomcp: your team's data, ready for every AI agent
> datatomcp is the AI-native data layer for teams without a data team. It makes a team's Google Sheets, CSV and Excel files, databases (PostgreSQL, MySQL) and apps (Stripe, HubSpot, Airtable) AI-ready: one governed place every AI agent can ask, with exact answers computed by a database, shared definitions of what each number means, a full audit trail of agent questions, charts and dashboards. Agents connect over a remote MCP server (OAuth) or a REST API. Hosted in the EU.
Source: https://datatomcp.com/
## How it works
1. Connect: upload CSV or Excel files, pick Google Sheets, or connect PostgreSQL, MySQL, Stripe, HubSpot or Airtable. Each source becomes a dataset with versions and sync, and an import report of anything odd.
2. Link: datatomcp finds the ids the datasets share (orders.customer_id and customers.customer_id) and links them, so agents answer across datasets in one query.
3. Define: approve the definitions datatomcp drafts from your columns, and add your team's words to the glossary.
4. Ask: add the MCP URL https://app.datatomcp.com/mcp to Claude, ChatGPT, Cursor or any MCP client, or call the REST API. Answers are counted by a database, name the dataset version, and can come back as a chart or a dashboard.
5. Audit: every agent question is recorded, with repeated questions compared.
## Facts
- What it is: an AI-native data layer (a hosted data warehouse plus semantic layer for teams without a data team). Not a spreadsheet add-on, not a BI tool, not a SQL database.
- Who it is for: teams that work from spreadsheets and use several AI agents (operations, finance, revenue, product, agencies), and the consultants who set this up for clients.
- MCP server: https://app.datatomcp.com/mcp (Streamable HTTP, OAuth sign-in, no API key)
- MCP tools: workspace_overview, list_datasets, describe_dataset, distinct_values, query_rows, aggregate (read-only data tools); chart (draws the answer in the chat, MCP Apps vi
[truncated]
Your team's data, ready for every AI agent.
Connect your databases, apps and sheets once. datatomcp links them, writes down what each number means, and lets every agent on your team ask for exact answers. Every question is recorded, every number can be checked, and the answer can come back as a chart.
Watch the 90-second demo Free for 3 datasets and 2 people · No card needed · Read-only for agents · Hosted in the EU
Click Connect, sign in to datatomcp, then ask about your data in any chat. Desktop and mobile use the same connector.
Claude Code Paste in Claude Code Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Set up datatomcp by reading https://datatomcp.com/install.md and following its instructions.
Sound on Works with any AI agent that supports MCP, and anything that speaks REST. Listed where agents look:
Model Context Protocol MCP Registry
Don't send the sheet. Send the question.
Pasting a sheet puts every row into the model's context: it fills the window, and you pay for every token. With datatomcp the agent asks a database and gets back only the result, about 475 times fewer tokens on our test file, and the total is computed, not estimated.
How many orders came from France in September, and what was the revenue?
Each square is 1,000 tokens of Claude Opus 5.5's 1M-token context window.
Measured on our 50,000-row test orders file. Pasting sends the whole 4,118,185-byte file; a full datatomcp session (tool list, schema, query and result) sends 8,663 bytes. Tokens estimated at about 4 bytes per token. Price: Claude Opus 5.5 input, $4 per million tokens (October 2026).
Connect it once. Every agent gets the same answer.
Five steps take your team from scattered sheets, exports and databases to answers you can put in a report without checking them by hand.
Connect PostgreSQL, MySQL, Stripe, HubSpot and Airtable, pick Google Sheets in Google's own picker, or upload Excel and CSV files. Each source becomes a dataset that stays in sync, with a version for every refresh.
Every import comes with a report: lines that couldn't be read, numbers stored as text, mixed dates. Nothing is converted until someone confirms it.
Add data Northwind Studio PostgreSQL
orders PostgreSQL 48 rows Importing Ready
customers Google Sheets 10 rows Ready
sample_sales CSV 1,500 rows 1 finding
revenue looks like numbers but is stored as text, like "$712.80" . Nothing changes until you confirm.
datatomcp finds the ids and codes your datasets share, like the customer_id in orders and in customers, and draws the map. Links chain, so an agent can total orders by a customer's region in one query.
A link confirms itself only when every value matches. Otherwise you see what didn't match, and decide.
Map 3 datasets · 2 links orders 48 order_id customer_id amount status customers 10 customer_id segment region_id regions 3 region_id region orders → customers 47 of 48 matched Confirm Confirmed
customers → regions 10 of 10 matched Confirmed on its own
group_by: ["regions.region"] works from orders now
Agree on what each number means
datatomcp drafts definitions from your columns, like paid revenue as the sum of amount where the status is paid. A person approves or edits each one, and agents only use approved ones.
Every answer that uses a definition says so. A glossary explains your team's words, like "active customer", the same way.
Definitions orders paid_revenue Draft Approved sum(amount) where status = "paid"
Revenue from paid orders; refunds excluded.
Also answers to revenue and sales
Approved by you · agents use it from now on Glossary
Add one URL to Claude, ChatGPT, Cursor or any MCP client, or call the REST API. A database counts every row, so totals are exact, and each answer names the dataset version it came from.
When a picture is clearer, the agent draws a chart or makes a dashboard right in the chat, and anyone can refresh it later.
Claude ChatGPT Cursor datatomcp connected What's our paid revenue by region?
aggregate · orders · paid_revenue by regions.region 190 ms
Paid revenue is 24,360 , from your team's approved paid_revenue definition.
orders, version 1 · 1 order has no matching customer · Check this number
Every agent question is kept: who asked, from which agent, on which dataset, the query, the time it took and any error.
When two agents ask the same question, you see whether they got the same answer. Keep 90 days on Pro, and export it as a spreadsheet.
Audit 72 calls · 4 agents · 3 people · 7 days Maya · Claude 190 ms 4 groups aggregate · orders · paid_revenue by regions.region
Sam · ChatGPT 160 ms 4 groups Same answer as Claude aggregate · orders · paid_revenue by regions.region
You · Cursor 120 ms 8 rows query_rows · orders · status = refunded
Sam · ChatGPT 140 ms Error: revenue is text aggregate · sample_sales · sum(revenue)
You · Claude Code 210 ms 8 columns describe_dataset · sample_sales
I help about 300 finance people use AI agents at a Fortune 100 company. Almost all of them hit the same wall: they paste a whole CSV or Google Sheet into the chat, the context fills up, the tokens add up, and the totals are close but not exact.
datatomcp is the tool I wanted for them: read-only for agents, hosted in the EU, and honest about its limits.
One workspace for the whole team. Data, meaning and every agent, side by side.
Every dataset, its freshness, the team's definitions and what the agents asked, on one page. Open this page live Which agents asked, which definitions they used, which datasets those read, and how the datasets link. Open this page live Approved definitions with who approved them, drafts waiting for a person, and the team's glossary. Open this page live Every agent question, repeated questions and whether they got the same answer, and every error. Open this page live A board an agent made from a few questions, refreshed on the latest data whenever you open it. This is the public demo workspace, Northwind Studio, with its real numbers. Click through every page, no sign-up needed.
Made for teams without a data team. And for the people who set them up.
Invite people as editors, who add and update data, or viewers, who ask. Each person connects their own agent once and gets the same datasets at the same version, so everyone quotes the same number.
Members Invite people N Northwind Studio Owner
Give each client a workspace of their own, with their data, their definitions and their words. Limit the agent you use for a client to that client's data, show them what their agents asked, then hand it over.
Workspaces New workspace N Northwind Studio 4 datasets · 3 people Current
+ Your next client Its own data, definitions and team
Claude Code Northwind Studio only
Start free, no card needed. Upgrade in the app when your data or your team outgrow it, and cancel anytime.
Sync within minutes of an edit
Agents read through validated, read-only tools, and only the results they ask for. Everything else stays with us, in the EU, until you delete it.
Hosted in the EU A copy of each source (one file per version), backups and accounts are kept in the European Union. No AI processing unless you opt in Never used for training. AI cleaning suggestions (Mistral, EU) and the in-app assistant (your own key) stay off until you turn them on. Private by default Every request needs a sign-in or an API key. Database and app credentials are encrypted and never shown to agents. Deleted when you say so Delete a dataset or your whole account from the app. It's gone at once, and from backups within 14 days. Questions
Anything else? Write to [email protected] , the founder answers.
A single place where your team's spreadsheets, files, databases and apps live, kept up to date, linked together and described with the team's own definitions, that every AI agent can ask for exact answers. Agents don't just reach the data: they know what each number means, a database does the counting, and every question is recorded. It does the job of a data warehouse and a semantic layer for teams without a data team.
Databases and apps: PostgreSQL and MySQL (a table or a read-only query), Stripe (customers, subscriptions with MRR, invoices, payments), HubSpot (contacts, companies, deals) and Airtable. Sheets and files: Google Sheets, Excel and CSV. Each one becomes a dataset with versions, sync, links to your other data and the audit. Need another source? Ask for it in the app: we add sources on request, the most requested first.
The team's agreed meaning of a number, like "paid revenue = sum of amount where status is paid". datatomcp drafts them from your columns when you add data, and a person approves or edits each one. Agents only see approved definitions, every answer that uses one says who approved it, and an unknown name gets the closest definitions instead of a guess. A glossary explains words like "active customer" the same way.
Any client that supports remote MCP servers, including Claude, ChatGPT, Cursor and Claude Code. Agents without MCP support, n8n workflows and your own apps use the REST API with an API key. Both return the same results.
Over MCP: workspace_overview, list_datasets, describe_dataset, distinct_values, query_rows and aggregate to read and count, chart to draw the answer, create_dashboard, list_dashboards and refresh_dashboard for dashboards, report_issue to flag a data problem to the team, and propose_definition, propose_link, propose_term and propose_brief to suggest a definition, a link, a glossary term or the busin

[truncated]
