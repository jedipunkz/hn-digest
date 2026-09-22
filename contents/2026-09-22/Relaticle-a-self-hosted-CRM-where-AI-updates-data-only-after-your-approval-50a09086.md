---
source: "https://relaticle.com"
hn_url: "https://news.ycombinator.com/item?id=49808105"
title: "Relaticle, a self-hosted CRM where AI updates data only after your approval"
article_title: "Relaticle - CRM Built for People and AI-Powered Work"
image: "https://relaticle.com/images/open-graph.jpg?v=2"
author: "manukminasyan"
captured_at: "2026-09-22T21:49:16Z"
capture_tool: "hn-digest"
hn_id: 49808105
score: 2
comments: 0
posted_at: "2026-09-22T21:06:26Z"
tags:
  - hacker-news
---

# Relaticle, a self-hosted CRM where AI updates data only after your approval

- HN: [49808105](https://news.ycombinator.com/item?id=49808105)
- Source: [relaticle.com](https://relaticle.com)
- Score: 2
- Comments: 0
- Posted: 2026-09-22T21:06:26Z

## Translation

Title: Relaticle, a self-hosted CRM where AI updates data only after your approval
Article title: Relaticle - CRM Built for People and AI-Powered Work
Description: Open-source, self-hosted CRM with a built-in AI chat and 39 MCP tools for external agents. Safe approvals, custom fields, and a REST API. Free forever.

Article text:
Relaticle - CRM Built for People and AI-Powered Work
Skip to main content
el.checkVisibility())
},
trapTab(event) {
const items = this.focusables()
if (items.length === 0) {
return
}
const edge = event.shiftKey ? items[0] : items[items.length - 1]
if (document.activeElement !== edge) {
return
}
event.preventDefault()
;(event.shiftKey ? items[items.length - 1] : items[0]).focus()
},
}"
x-effect="bodyLock()"
x-init="$watch('mobileMenu', (open) => { if (open) { mobileExpanded = null } else { $refs.hamburger.focus() } })"
@resize.window="if (window.innerWidth >= 768) mobileMenu = false">
{ this.swapTo(slug) }, 90)
},
cancelHoverOpen() { clearTimeout(this.hoverTimer) },
closeOnHoverLeave(slug) {
this.closeTimer = setTimeout(() => {
if (this.openDropdown === slug) { this.closeAll() }
}, 200)
},
// A full close (not a swap) still needs its content to fade out WITH
// the card rather than vanishing the instant openDropdown clears, so
// it borrows the same outgoingDropdown mechanism a swap uses, just with
// no incoming layer to slide past.
closeAll() {
clearTimeout(this.swapTimer)
clearTimeout(this.heightTimer)
if (this.$refs.panel) { this.$refs.panel.style.height = '' }
// outgoingDropdown is set BEFORE openDropdown clears, so
// `openDropdown ?? outgoingDropdown` is never null-null on the same
// tick. A gap there sends panelStyle() an unresolvable slug, which
// drops the inline width/height and collapses the card to a stray dot.
this.outgoingDropdown = this.openDropdown
this.openDropdown = null
// Alpine's own x-show leave transition on the card (driven purely by
// openDropdown, not outgoingDropdown) already runs and finishes on
// its own. Writing to outgoingDropdown again afterwards was enough to
// make Alpine re-run that transition's bookkeeping and re-show the
// card at a collapsed size, so only clear it if a NEW open hasn't
// already claimed it (swapTo re-purposes the same field).
const closingSlug = this.outgoingDropdown
this.swapTimer = setTimeout(() => {
if
[truncated]
Help center
Guides for every part of the product
Developers
Self-hosting, MCP, and contributing
Blog
How we build an open-source CRM
Compare
Relaticle vs EspoCRM
Attio alternative
HubSpot alternative
1.7K
158
Sign In
Start for free
$refs.close.focus())"
class="p-2.5 -mr-1 text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white active:opacity-60 rounded-lg transition-[color,opacity] duration-200 cursor-pointer focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
aria-label="Close menu">
Product
Rela
Features
Self-hosted
MCP and API
Resources
Learn
Help center
Developers
Blog
Compare
1.7K
158
Sign In
Start for free
Open Source
1.7K
0
1
.
0
1
2
3
4
5
6
7
K
+ stars
The CRM built for people
and AI-powered work
Open-source, self-hosted, and human-first.
Built-in AI chat plus 39 MCP tools for external agents.
Search
⌘K
Home
People
Companies
Opportunities
Tasks
Notes
Chats
Overdue tasks this week
This week's pipeline review
Follow up with Priya Nair
Renewal prep: Daniel Okafor
All chats
Overdue tasks this week
Tasks
Title
Due date
Company
Call Sarah Chen
Aug 24, 2026
Kovra Systems
Send proposal to Trellis Labs
Aug 23, 2026
Trellis Labs
Schedule demo with Kovra Systems
Aug 22, 2026
Kovra Systems
Mark the Kovra demo as done.
Assistant
Review the proposal below to update the task.
Schedule demo with Kovra Systems
Done
Schedule demo with Kovra Systems has been marked done.
Add Sarah Chen as a contact at Kovra Systems . She's VP of Engineering.
Assistant
Review the proposal below to add her to Kovra Systems.
Sarah Chen
Created
Sarah Chen has been created and linked to
Kovra Systems .
Sarah Chen
Job title
VP of Engineering
Company
Kovra Systems
Link Sarah to an opportunity
Create a task for Sarah
Add a note about Sarah
Review before continuing
Update Task
Schedule demo with Kovra Systems
Attribute
New value
Status
To do
Done
Create Person
Sarah Chen
Attribute
New value
Name
Sarah Chen
Job title
VP of Engineering
Company
Kovra Systems
Discard
Save changes
⌘⏎
Ask anything...
Auto
Good morning, Marcus.
Recent chat · This week's pipeline review
Ask anything...
Auto
Tasks
3
Send proposal to Trellis Labs
Aug 23, 2026
Renewal prep for Daniel Okafor
Sep 2, 2026
Works with the agents your team already uses
Claude
ChatGPT
Cursor
Gemini
+ any MCP client
33,000+ Docker pulls
/ self-hosted, AGPL-3.0, shipping weekly
Work in the app, ask Rela in chat , or connect your agents. Keep your team working from the same customer data.
Give MCP-compatible agents access to your CRM through 39 tools. Build custom integrations with the REST API.
Your agents
Claude
ChatGPT
Gemini CLI
Custom
Relaticle
MCP + REST API
Your CRM
People
Companies
Deals
Tasks
Notes
Your agents and CRM, connected.
Built-in AI chat
Ask Rela about your CRM and make changes through chat. Review proposed updates and deletions before they run.
Capture the details that matter to your business. Add custom fields, connect related records, and show fields only when relevant.
See every account in context. Keep company details, people, and deals together so your team can prepare for the next conversation.
Know who you're talking to. Connect people to their companies, keep relationship notes, and find the right person with filters.
Follow every deal from first contact to close. Move deals through custom stages on a visual board that matches your sales process.
Turn follow-ups into clear next steps. Assign tasks, set due dates, and link them to people, companies, or deals.
Send proposal to
@Acme
Follow up with
@Sarah
Review Q4 pipeline
Due today
Team collaboration
Work from shared customer records. Organize teams in separate workspaces and use roles to control who can view or change data.
Bring your records into Relaticle with CSV imports. Map columns, review validation errors, and export your data when you need it.
Keep customer context close to the records it belongs to. Add linked notes and review the activity history to see what changed.
Bring your team and agents onto the same CRM.
Relaticle is AGPL-3.0 open source, so you can run it on your own server . Star the repo, join Discord, and help shape the future of agent-native CRM.
Star our repo, report issues, and contribute code. Completely open source and free to use.
Chat with developers, get help, and share ideas. Join our growing community of builders.
Self-host with Docker, build on the REST API, and connect AI agents over MCP.
Everything you need to know about Relaticle, from deployment to AI agent integration.
Yes. Relaticle has 2,000+ automated tests, 5-layer authorization, 56+ MCP-specific tests, and is used in production. The codebase is continuously tested with PHPStan static analysis and Pest mutation testing.
What can the built-in AI chat do?
Ask anything about your CRM and the chat works on your data: list and search records, draft follow-ups, summarize a deal, create a task, update or delete a record. @-mention any record (people, companies, deals, tasks, notes) to scope a question. Voice input, persistent searchable history, and dashboard insight cards are included.
Can the AI chat delete or change my CRM data without my approval?
No. Destructive operations (delete, update existing records) show an approval card with Approve and Reject buttons. Nothing happens until you click. Approved destructive actions can be undone for 5 seconds via a toast. Read-only and create operations don't require approval.
Does the built-in chat send my data to OpenAI or Anthropic?
Inference runs through whichever AI provider your team configures (Anthropic Claude, Google Gemini, or any OpenAI-compatible endpoint). Conversation history is stored only in your Relaticle database, and Relaticle never trains on your data. Self-hosted teams supply their own provider keys, so the destination is yours to choose.
What AI agents can I connect from outside?
Any agent that speaks MCP (Model Context Protocol). Claude, ChatGPT, Gemini, open-source models, or your own custom agents. Relaticle's MCP server provides 39 tools for external AI agents to read, create, update, delete, and analyze CRM data.
What is MCP?
MCP (Model Context Protocol) is an open standard that lets AI agents interact with tools and data sources. Relaticle's MCP server gives external agents 39 tools to list companies, create contacts, update opportunities, analyze pipelines, and more.
How is Relaticle different from HubSpot or Salesforce?
Relaticle is self-hosted (you own your data), open-source (AGPL-3.0), ships with both a bui
[truncated]
Self-hosted. Agent-native. Full control over your data and your AI.
The open-source CRM built for people and AI-powered work. Self-hosted. No per-seat pricing. Yours to own.
© 2026 Relaticle. All rights
reserved.

## Original Extract

Open-source, self-hosted CRM with a built-in AI chat and 39 MCP tools for external agents. Safe approvals, custom fields, and a REST API. Free forever.

Relaticle - CRM Built for People and AI-Powered Work
Skip to main content
el.checkVisibility())
},
trapTab(event) {
const items = this.focusables()
if (items.length === 0) {
return
}
const edge = event.shiftKey ? items[0] : items[items.length - 1]
if (document.activeElement !== edge) {
return
}
event.preventDefault()
;(event.shiftKey ? items[items.length - 1] : items[0]).focus()
},
}"
x-effect="bodyLock()"
x-init="$watch('mobileMenu', (open) => { if (open) { mobileExpanded = null } else { $refs.hamburger.focus() } })"
@resize.window="if (window.innerWidth >= 768) mobileMenu = false">
{ this.swapTo(slug) }, 90)
},
cancelHoverOpen() { clearTimeout(this.hoverTimer) },
closeOnHoverLeave(slug) {
this.closeTimer = setTimeout(() => {
if (this.openDropdown === slug) { this.closeAll() }
}, 200)
},
// A full close (not a swap) still needs its content to fade out WITH
// the card rather than vanishing the instant openDropdown clears, so
// it borrows the same outgoingDropdown mechanism a swap uses, just with
// no incoming layer to slide past.
closeAll() {
clearTimeout(this.swapTimer)
clearTimeout(this.heightTimer)
if (this.$refs.panel) { this.$refs.panel.style.height = '' }
// outgoingDropdown is set BEFORE openDropdown clears, so
// `openDropdown ?? outgoingDropdown` is never null-null on the same
// tick. A gap there sends panelStyle() an unresolvable slug, which
// drops the inline width/height and collapses the card to a stray dot.
this.outgoingDropdown = this.openDropdown
this.openDropdown = null
// Alpine's own x-show leave transition on the card (driven purely by
// openDropdown, not outgoingDropdown) already runs and finishes on
// its own. Writing to outgoingDropdown again afterwards was enough to
// make Alpine re-run that transition's bookkeeping and re-show the
// card at a collapsed size, so only clear it if a NEW open hasn't
// already claimed it (swapTo re-purposes the same field).
const closingSlug = this.outgoingDropdown
this.swapTimer = setTimeout(() => {
if
[truncated]
Help center
Guides for every part of the product
Developers
Self-hosting, MCP, and contributing
Blog
How we build an open-source CRM
Compare
Relaticle vs EspoCRM
Attio alternative
HubSpot alternative
1.7K
158
Sign In
Start for free
$refs.close.focus())"
class="p-2.5 -mr-1 text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white active:opacity-60 rounded-lg transition-[color,opacity] duration-200 cursor-pointer focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary"
aria-label="Close menu">
Product
Rela
Features
Self-hosted
MCP and API
Resources
Learn
Help center
Developers
Blog
Compare
1.7K
158
Sign In
Start for free
Open Source
1.7K
0
1
.
0
1
2
3
4
5
6
7
K
+ stars
The CRM built for people
and AI-powered work
Open-source, self-hosted, and human-first.
Built-in AI chat plus 39 MCP tools for external agents.
Search
⌘K
Home
People
Companies
Opportunities
Tasks
Notes
Chats
Overdue tasks this week
This week's pipeline review
Follow up with Priya Nair
Renewal prep: Daniel Okafor
All chats
Overdue tasks this week
Tasks
Title
Due date
Company
Call Sarah Chen
Aug 24, 2026
Kovra Systems
Send proposal to Trellis Labs
Aug 23, 2026
Trellis Labs
Schedule demo with Kovra Systems
Aug 22, 2026
Kovra Systems
Mark the Kovra demo as done.
Assistant
Review the proposal below to update the task.
Schedule demo with Kovra Systems
Done
Schedule demo with Kovra Systems has been marked done.
Add Sarah Chen as a contact at Kovra Systems . She's VP of Engineering.
Assistant
Review the proposal below to add her to Kovra Systems.
Sarah Chen
Created
Sarah Chen has been created and linked to
Kovra Systems .
Sarah Chen
Job title
VP of Engineering
Company
Kovra Systems
Link Sarah to an opportunity
Create a task for Sarah
Add a note about Sarah
Review before continuing
Update Task
Schedule demo with Kovra Systems
Attribute
New value
Status
To do
Done
Create Person
Sarah Chen
Attribute
New value
Name
Sarah Chen
Job title
VP of Engineering
Company
Kovra Systems
Discard
Save changes
⌘⏎
Ask anything...
Auto
Good morning, Marcus.
Recent chat · This week's pipeline review
Ask anything...
Auto
Tasks
3
Send proposal to Trellis Labs
Aug 23, 2026
Renewal prep for Daniel Okafor
Sep 2, 2026
Works with the agents your team already uses
Claude
ChatGPT
Cursor
Gemini
+ any MCP client
33,000+ Docker pulls
/ self-hosted, AGPL-3.0, shipping weekly
Work in the app, ask Rela in chat , or connect your agents. Keep your team working from the same customer data.
Give MCP-compatible agents access to your CRM through 39 tools. Build custom integrations with the REST API.
Your agents
Claude
ChatGPT
Gemini CLI
Custom
Relaticle
MCP + REST API
Your CRM
People
Companies
Deals
Tasks
Notes
Your agents and CRM, connected.
Built-in AI chat
Ask Rela about your CRM and make changes through chat. Review proposed updates and deletions before they run.
Capture the details that matter to your business. Add custom fields, connect related records, and show fields only when relevant.
See every account in context. Keep company details, people, and deals together so your team can prepare for the next conversation.
Know who you're talking to. Connect people to their companies, keep relationship notes, and find the right person with filters.
Follow every deal from first contact to close. Move deals through custom stages on a visual board that matches your sales process.
Turn follow-ups into clear next steps. Assign tasks, set due dates, and link them to people, companies, or deals.
Send proposal to
@Acme
Follow up with
@Sarah
Review Q4 pipeline
Due today
Team collaboration
Work from shared customer records. Organize teams in separate workspaces and use roles to control who can view or change data.
Bring your records into Relaticle with CSV imports. Map columns, review validation errors, and export your data when you need it.
Keep customer context close to the records it belongs to. Add linked notes and review the activity history to see what changed.
Bring your team and agents onto the same CRM.
Relaticle is AGPL-3.0 open source, so you can run it on your own server . Star the repo, join Discord, and help shape the future of agent-native CRM.
Star our repo, report issues, and contribute code. Completely open source and free to use.
Chat with developers, get help, and share ideas. Join our growing community of builders.
Self-host with Docker, build on the REST API, and connect AI agents over MCP.
Everything you need to know about Relaticle, from deployment to AI agent integration.
Yes. Relaticle has 2,000+ automated tests, 5-layer authorization, 56+ MCP-specific tests, and is used in production. The codebase is continuously tested with PHPStan static analysis and Pest mutation testing.
What can the built-in AI chat do?
Ask anything about your CRM and the chat works on your data: list and search records, draft follow-ups, summarize a deal, create a task, update or delete a record. @-mention any record (people, companies, deals, tasks, notes) to scope a question. Voice input, persistent searchable history, and dashboard insight cards are included.
Can the AI chat delete or change my CRM data without my approval?
No. Destructive operations (delete, update existing records) show an approval card with Approve and Reject buttons. Nothing happens until you click. Approved destructive actions can be undone for 5 seconds via a toast. Read-only and create operations don't require approval.
Does the built-in chat send my data to OpenAI or Anthropic?
Inference runs through whichever AI provider your team configures (Anthropic Claude, Google Gemini, or any OpenAI-compatible endpoint). Conversation history is stored only in your Relaticle database, and Relaticle never trains on your data. Self-hosted teams supply their own provider keys, so the destination is yours to choose.
What AI agents can I connect from outside?
Any agent that speaks MCP (Model Context Protocol). Claude, ChatGPT, Gemini, open-source models, or your own custom agents. Relaticle's MCP server provides 39 tools for external AI agents to read, create, update, delete, and analyze CRM data.
What is MCP?
MCP (Model Context Protocol) is an open standard that lets AI agents interact with tools and data sources. Relaticle's MCP server gives external agents 39 tools to list companies, create contacts, update opportunities, analyze pipelines, and more.
How is Relaticle different from HubSpot or Salesforce?
Relaticle is self-hosted (you own your data), open-source (AGPL-3.0), ships with both a bui
[truncated]
Self-hosted. Agent-native. Full control over your data and your AI.
The open-source CRM built for people and AI-powered work. Self-hosted. No per-seat pricing. Yours to own.
© 2026 Relaticle. All rights
reserved.
