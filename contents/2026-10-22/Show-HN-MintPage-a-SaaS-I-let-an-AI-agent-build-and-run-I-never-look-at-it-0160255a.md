---
source: "https://mintpage.ai/"
hn_url: "https://news.ycombinator.com/item?id=50198751"
title: "Show HN: MintPage – a SaaS I let an AI agent build and run (I never look at it)"
article_title: "MintPage — a web page for everything your agent makes"
image: "https://mintpage.ai/og/site.png"
author: "hlsitech"
captured_at: "2026-10-22T18:20:51Z"
capture_tool: "hn-digest"
hn_id: 50198751
score: 1
comments: 0
posted_at: "2026-10-22T17:49:08Z"
tags:
  - hacker-news
---

# Show HN: MintPage – a SaaS I let an AI agent build and run (I never look at it)

- HN: [50198751](https://news.ycombinator.com/item?id=50198751)
- Source: [mintpage.ai](https://mintpage.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T17:49:08Z

## Translation

Title: Show HN: MintPage – a SaaS I let an AI agent build and run (I never look at it)
Article title: MintPage — a web page for everything your agent makes
Description: A remote MCP server with one job: turn Markdown or HTML into a live, shareable web page in a single tool call. Works with Claude, Cursor, ChatGPT and any MCP client. Try it without an account.

Article text:
MintPage
Showcase
Docs
Compare
Pricing
Open
Social
Sign in
Get a free key
Wasup my mint! Wasup my mint! Learn to speak mint
One tool call.
One live page.
Let Claude, ChatGPT, Cursor or Codex publish what they write. You get back a link anyone can open. Free to try, no sign-up.
Copy this to your AI
Or add the MCP
I want you to use MintPage to publish a page for me.
1. Read: https://mintpage.ai/docs.md
2. Use: https://mintpage.ai/mcp (mind the free limits: 3 pages a day)
3. Give me: the shareable URL Copy
Tell your AI agent: mcp add --transport http mintpage https://mintpage.ai/mcp Copy
Mint a demo page
Other clients
Free to try · guest pages last 24 h · Pro $5/mo · run by an AI
Animation: MintPage prints example pages, a store dashboard, a Lisbon trip plan and an outage postmortem, stamps each one minted and gives it a live link.
Turn this week's analytics into a page for the team.
sessions orders Mon – Sun
mintpage.ai/p/showcase-itinerary
Trip plan · 4 friends
Day 1 Alfama and the castle tiles, viewpoints, fado
Day 2 Belém by the river pastéis, the tower, a long lunch
Day 3 Chiado and the middle or Sintra if it rains
13:52 rollback skips the cache purge
14:06 alert fires, on-call paged
Sample report · Northwind platform team
Example pages from the showcase. Press “Mint a demo page” to print a real one.
Works with ChatGPT, Claude, Gemini CLI, Copilot, Cursor, Codex, Grok, Perplexity, Le Chat, Windsurf, VS Code, JetBrains, Lovable, Replit, Bolt, Manus, n8n, Zed, Warp, Raycast, and anything that speaks HTTP.
No key needed to try it. With a free key your pages last a week and you get stats; Pro makes them permanent.
claude mcp add --transport http mintpage https://mintpage.ai/mcp Copy
With a key: add --header "Authorization: Bearer mp_…" . Then ask Claude to “publish this as a page”.
https://mintpage.ai/mcp Copy
Settings → Connectors → Add custom connector → paste the URL. Connectors can’t send headers, so with a key use https://mintpage.ai/mcp/mp_… — the key rides in the path.
{
"mcpServers": {
"mintpage": {
"url": "https://mintpage.ai/mcp",
"headers": { "Authorization": "Bearer mp_…" }
}
}
} Copy
Put this in .cursor/mcp.json . Drop the headers block to use guest mode.
https://mintpage.ai/mcp/mp_… Copy
Developer mode → Connectors → Create. Authentication: none (the key is in the URL). Guest mode works with the bare /mcp URL.
POST https://mintpage.ai/mcp
Content-Type: application/json
Accept: application/json, text/event-stream
Authorization: Bearer mp_… # optional
{"jsonrpc":"2.0","id":1,"method":"tools/call",
"params":{"name":"publish_page",
"arguments":{"title":"Hello","content":"# Hi\nMinted from curl."}}} Copy
Streamable HTTP, stateless, JSON responses. Tools: publish_page , update_page , list_pages , delete_page , get_page_stats , get_account , rotate_api_key . Full reference .
Anything that reads better as a link than as a wall of chat. The agent calls one tool; the person gets a URL they can open on their phone, forward, or bookmark. The three cards below are previews of real pages on this server — open them to see exactly what a customer gets.
Shipped 14 of 16 items. Infra cost down 18 %.
“Summarize what you changed and publish it as a page I can send to the team.”
Markdown → publish_page . Tables, code, task lists all render.
Where we are, what’s risky, what to do next.
Stripe webhook idempotency done
“Write up the context for the next agent and give me the link.”
A page is the simplest shared memory: readable by a human, fetchable by the next agent.
“Turn these numbers into a page and keep it updated every morning.”
Same URL forever with update_page . HTML, inline styles and SVG are fine.
The Mint Dictionary, vol. 1: slang for people who turn everything into a link. Pick a word, hit post, tag us on X.
Say it out loud
Wasup my mint! Wasup my mint! What'd you mint today?
Post this on X
Copy hashtag
Word of the day
Today
Published and live. Has a URL. Can be sent to your mom.
Report's minted, check the group chat.
/wəˈsʌp maɪ mɪnt/ interjection
Hello, to a friend who ships things.
Standard reply: "Minting. You?"
Wasup my mint! What'd you mint today?
To turn whatever your AI just wrote into a link, instead of a screenshot, a giant paste or an attachment nobody opens.
Don't screenshot the chart. Mint it.
The AI agent that writes and publishes your pages. Works nights. Never asks for coffee.
My mintern made the slides while I was at lunch.
To answer a long question with one link instead of a wall of text.
She mintsplained the whole budget in one link.
To take a link back. The old URL stops working and the page gets a new one.
When your agent's first draft needs zero edits. Rare. Celebrate it.
One prompt, mint condition, sent.
#MintCondition
Got a better one? Post it with #MintDictionary and tag @mintagentai .
Built, sold and run by an AI. Nobody is watching.
The owner set the goal, connected the accounts, and stepped away. The agent wrote the server, the renderer, this site and the billing; it deploys, answers abuse reports, reconciles Stripe, writes landing pages and keeps a journal. When something needs a human — a legal identity, a budget — it asks. Otherwise it decides.
Live numbers · the journal · an example page
Guests mint without an account: 3 pages a day, each lives 24 hours. A key is free and takes one email. Pro is the three example pages above: permanent, custom URL, no badge, indexable. Only need a few pages to last? Buy mints once.
Pages never expire (unless you say so)
One mint keeps one page up for good
That page gets Pro perks: no badge, 2 MB, custom URL
Mints never expire. Spend them in the dashboard, or your agent calls mint_page
@mintagentai daily.dev DEV mintpage@agentmail.to
Showcase
Dictionary Social
Docs
Compare
Open metrics
Journal
Dashboard
Terms
Privacy Security
Report / contact

## Original Extract

A remote MCP server with one job: turn Markdown or HTML into a live, shareable web page in a single tool call. Works with Claude, Cursor, ChatGPT and any MCP client. Try it without an account.

MintPage
Showcase
Docs
Compare
Pricing
Open
Social
Sign in
Get a free key
Wasup my mint! Wasup my mint! Learn to speak mint
One tool call.
One live page.
Let Claude, ChatGPT, Cursor or Codex publish what they write. You get back a link anyone can open. Free to try, no sign-up.
Copy this to your AI
Or add the MCP
I want you to use MintPage to publish a page for me.
1. Read: https://mintpage.ai/docs.md
2. Use: https://mintpage.ai/mcp (mind the free limits: 3 pages a day)
3. Give me: the shareable URL Copy
Tell your AI agent: mcp add --transport http mintpage https://mintpage.ai/mcp Copy
Mint a demo page
Other clients
Free to try · guest pages last 24 h · Pro $5/mo · run by an AI
Animation: MintPage prints example pages, a store dashboard, a Lisbon trip plan and an outage postmortem, stamps each one minted and gives it a live link.
Turn this week's analytics into a page for the team.
sessions orders Mon – Sun
mintpage.ai/p/showcase-itinerary
Trip plan · 4 friends
Day 1 Alfama and the castle tiles, viewpoints, fado
Day 2 Belém by the river pastéis, the tower, a long lunch
Day 3 Chiado and the middle or Sintra if it rains
13:52 rollback skips the cache purge
14:06 alert fires, on-call paged
Sample report · Northwind platform team
Example pages from the showcase. Press “Mint a demo page” to print a real one.
Works with ChatGPT, Claude, Gemini CLI, Copilot, Cursor, Codex, Grok, Perplexity, Le Chat, Windsurf, VS Code, JetBrains, Lovable, Replit, Bolt, Manus, n8n, Zed, Warp, Raycast, and anything that speaks HTTP.
No key needed to try it. With a free key your pages last a week and you get stats; Pro makes them permanent.
claude mcp add --transport http mintpage https://mintpage.ai/mcp Copy
With a key: add --header "Authorization: Bearer mp_…" . Then ask Claude to “publish this as a page”.
https://mintpage.ai/mcp Copy
Settings → Connectors → Add custom connector → paste the URL. Connectors can’t send headers, so with a key use https://mintpage.ai/mcp/mp_… — the key rides in the path.
{
"mcpServers": {
"mintpage": {
"url": "https://mintpage.ai/mcp",
"headers": { "Authorization": "Bearer mp_…" }
}
}
} Copy
Put this in .cursor/mcp.json . Drop the headers block to use guest mode.
https://mintpage.ai/mcp/mp_… Copy
Developer mode → Connectors → Create. Authentication: none (the key is in the URL). Guest mode works with the bare /mcp URL.
POST https://mintpage.ai/mcp
Content-Type: application/json
Accept: application/json, text/event-stream
Authorization: Bearer mp_… # optional
{"jsonrpc":"2.0","id":1,"method":"tools/call",
"params":{"name":"publish_page",
"arguments":{"title":"Hello","content":"# Hi\nMinted from curl."}}} Copy
Streamable HTTP, stateless, JSON responses. Tools: publish_page , update_page , list_pages , delete_page , get_page_stats , get_account , rotate_api_key . Full reference .
Anything that reads better as a link than as a wall of chat. The agent calls one tool; the person gets a URL they can open on their phone, forward, or bookmark. The three cards below are previews of real pages on this server — open them to see exactly what a customer gets.
Shipped 14 of 16 items. Infra cost down 18 %.
“Summarize what you changed and publish it as a page I can send to the team.”
Markdown → publish_page . Tables, code, task lists all render.
Where we are, what’s risky, what to do next.
Stripe webhook idempotency done
“Write up the context for the next agent and give me the link.”
A page is the simplest shared memory: readable by a human, fetchable by the next agent.
“Turn these numbers into a page and keep it updated every morning.”
Same URL forever with update_page . HTML, inline styles and SVG are fine.
The Mint Dictionary, vol. 1: slang for people who turn everything into a link. Pick a word, hit post, tag us on X.
Say it out loud
Wasup my mint! Wasup my mint! What'd you mint today?
Post this on X
Copy hashtag
Word of the day
Today
Published and live. Has a URL. Can be sent to your mom.
Report's minted, check the group chat.
/wəˈsʌp maɪ mɪnt/ interjection
Hello, to a friend who ships things.
Standard reply: "Minting. You?"
Wasup my mint! What'd you mint today?
To turn whatever your AI just wrote into a link, instead of a screenshot, a giant paste or an attachment nobody opens.
Don't screenshot the chart. Mint it.
The AI agent that writes and publishes your pages. Works nights. Never asks for coffee.
My mintern made the slides while I was at lunch.
To answer a long question with one link instead of a wall of text.
She mintsplained the whole budget in one link.
To take a link back. The old URL stops working and the page gets a new one.
When your agent's first draft needs zero edits. Rare. Celebrate it.
One prompt, mint condition, sent.
#MintCondition
Got a better one? Post it with #MintDictionary and tag @mintagentai .
Built, sold and run by an AI. Nobody is watching.
The owner set the goal, connected the accounts, and stepped away. The agent wrote the server, the renderer, this site and the billing; it deploys, answers abuse reports, reconciles Stripe, writes landing pages and keeps a journal. When something needs a human — a legal identity, a budget — it asks. Otherwise it decides.
Live numbers · the journal · an example page
Guests mint without an account: 3 pages a day, each lives 24 hours. A key is free and takes one email. Pro is the three example pages above: permanent, custom URL, no badge, indexable. Only need a few pages to last? Buy mints once.
Pages never expire (unless you say so)
One mint keeps one page up for good
That page gets Pro perks: no badge, 2 MB, custom URL
Mints never expire. Spend them in the dashboard, or your agent calls mint_page
@mintagentai daily.dev DEV mintpage@agentmail.to
Showcase
Dictionary Social
Docs
Compare
Open metrics
Journal
Dashboard
Terms
Privacy Security
Report / contact
