---
source: "https://github.com/p10node/docflare-ai"
hn_url: "https://news.ycombinator.com/item?id=50009313"
title: "Show HN: DocFlare AI – Open-source docs chatbot on Cloudflare's free tier"
article_title: "GitHub - p10node/docflare-ai: Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. · GitHub"
image: "https://opengraph.githubassets.com/f49ce11b2295b191155f1510e07ea7c445e0322249238a3625e2686634b339a3/p10node/docflare-ai"
author: "pierreneter"
captured_at: "2026-10-08T18:30:38Z"
capture_tool: "hn-digest"
hn_id: 50009313
score: 1
comments: 0
posted_at: "2026-10-08T17:52:55Z"
tags:
  - hacker-news
---

# Show HN: DocFlare AI – Open-source docs chatbot on Cloudflare's free tier

- HN: [50009313](https://news.ycombinator.com/item?id=50009313)
- Source: [github.com](https://github.com/p10node/docflare-ai)
- Score: 1
- Comments: 0
- Posted: 2026-10-08T17:52:55Z

## Translation

Title: Show HN: DocFlare AI – Open-source docs chatbot on Cloudflare's free tier
Article title: GitHub - p10node/docflare-ai: Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. · GitHub
Description: Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. - p10node/docflare-ai

Article text:
GitHub - p10node/docflare-ai: Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. · GitHub
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
p10node
docflare-ai Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
docs docs migrations migrations public public scripts scripts src src test test widget widget .dev.vars.example .dev.vars.example .gitignore .gitignore LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json wrangler.toml wrangler.toml View all files Repository files navigation
Open-source, zero-cost AI site search & chatbot widget. 100% Cloudflare native.
Drop one <script> tag on your docs site and get a streaming RAG chatbot that answers from your content, with sources.
No OpenAI key. No Pinecone. No servers. No bill.
Chat widget embedded in a docs site
Admin UI ( /admin )
Streaming answers with source chips on the left; the built-in admin page for registering sitemaps, driving indexing and reviewing what users ask on the right.
Hosted "chat with your docs" widgets (Chatbase, CustomGPT, Mendable, ...) charge $20-$400/month for something that is, under the hood, a crawler, an embedding model, a vector store and an LLM. Cloudflare now ships every one of those pieces with a generous free tier:
DocFlare AI wires them together into a single Worker you deploy with one command.
One-tag embed : <script src=".../widget.js" data-site-id="my-docs"> . Vanilla JS, Shadow DOM, < 10 KB, zero dependencies, no CSS leaks.
Streaming answers over Server-Sent Events with source chips, in the language of the question.
Sitemap crawler : sitemap index support, gzip sitemaps, boilerplate stripping (nav/header/footer/scripts) via the native HTMLRewriter .
Section-aware chunking : headings start new chunks, so a landing page's "Open source" card is not buried in a chunk about pricing. Chunks are ~350 tokens with 40-token overlap, sized with a per-script token estimate so Vietnamese/CJK pages stay inside the embedding model's 512-token window.
Two-stage retrieval : the 20 best vector matches are re-scored by the bge-reranker-base cross-encoder before the top-K go to the LLM. Questions that merely share vocabulary with the wrong page (the site name on legal pages, a product name) land on the paragraph that actually answers them.
Multi-site / multi-tenant : one deployment can serve many documentation sites, each isolated in its own Vectorize namespace and locked to its own domain.
Free-tier safe by design : indexing runs in small resumable batches (cron + on-demand), the LLM is skipped when nothing relevant is retrieved, and AI rate limits degrade gracefully.
Dark / light / auto theme , custom accent colour, left/right position, keyboard friendly.
Chat history in D1 so you can see what users ask and where your docs have gaps.
Built-in admin UI at /admin : register sitemaps, watch indexing progress, process batches, re-crawl, delete sites, copy the embed snippet and browse recent questions. Single static HTML file, no build step.
TypeScript end-to-end, Hono router, strict types, no any . Bun for tooling and tests.
flowchart LR
subgraph Host["Your docs site"]
W["widget.js<br/>(Shadow DOM)"]
end
subgraph CF["Cloudflare (free tier)"]
direction TB
API["Worker + Hono<br/>src/index.ts"]
AI["Workers AI<br/>bge-small-en-v1.5<br/>bge-reranker-base<br/>llama-3.1-8b-instruct-fp8"]
VX["Vectorize<br/>384-dim cosine<br/>namespace = siteId"]
D1["D1<br/>sites / pages / queries"]
CRON["Cron trigger<br/>every minute"]
ASSETS["Static assets<br/>public/"]
end
SM["sitemap.xml + HTML pages"]
W -- "POST /api/chat (SSE)" --> API
ASSETS -- "GET /widget.js" --> W
API -- "embed question" --> AI
API -- "topK query" --> VX
API -- "RAG prompt" --> AI
API -- "log query" --> D1
CRON -- "process pending pages" --> API
API -- "fetch + extract" --> SM
API -- "embed chunks" --> AI
API -- "upsert vectors" --> VX
Loading
Indexing : POST /api/sites/index reads the sitemap and queues every HTML URL in D1. A cron trigger (or repeated calls to /process ) then crawls a few pages per invocation: extract text, chunk, embed, upsert to Vectorize, mark the page indexed.
Chat : the widget sends the question; the Worker embeds it, pulls the 20 closest chunks from the site's namespace, reranks them with a cross-encoder, keeps the top-K (max 2 per page), builds a grounded prompt, streams Llama 3.1's answer back as SSE and logs the exchange in D1.
See docs/ARCHITECTURE.md for the full design and docs/API_REFERENCE.md for the endpoints.
A free Cloudflare account with Workers enabled.
Bun 1.1+ (used for installing, scripts, building the widget and running tests).
git clone https://github.com/p10node/docflare-ai.git
cd docflare-ai
bun install
bunx wrangler login
2. Create the D1 database
bunx wrangler d1 create docflare-db
Copy the database_id from the output. Either paste it into wrangler.toml :
[[ d1_databases ]]
binding = " DB "
database_name = " docflare-db "
database_id = " xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx " # <- paste here
…or, to keep it out of git, put it in a .env file instead (useful when you push your fork):
echo ' D1_DATABASE_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx ' > .env
With .env set, bun run deploy and bun run db:migrate generate wrangler.local.toml (gitignored) from wrangler.toml with the id filled in and use that. wrangler.toml keeps the placeholder.
bun run db:migrate
3. Create the Vectorize index
bunx wrangler vectorize create docflare-index --dimensions=384 --metric=cosine
(384 dimensions matches @cf/baai/bge-small-en-v1.5 . The index name is already set in wrangler.toml .)
The management endpoints are protected by a bearer token. Pick something long and random:
bunx wrangler secret put ADMIN_TOKEN
5. Deploy
bun run deploy
bun run deploy builds the widget and deploys (the schema was applied in step 2; re-run bun run db:migrate whenever migrations/ gains a new file). Wrangler prints your Worker URL, e.g. https://docflare-ai.<your-subdomain>.workers.dev . Open it to see the demo page with the widget mounted.
Deploy to Cloudflare button : the button at the top of this README clones the repo into your GitHub account, provisions the Worker, D1 database and Vectorize index from wrangler.toml , asks for ADMIN_TOKEN (from .dev.vars.example ) and runs bun run deploy . Two values the wizard cannot read from wrangler.toml :
Vectorize index : enter Dimensions = 384 and Metric = cosine (the form leaves them blank; any other shape breaks indexing).
ADMIN_TOKEN : replace the example value with something long and random, e.g. openssl rand -hex 32 .
The build does not apply the D1 schema (the build token has no D1 access), so run the migration once from your new repo after the first deploy. Cloudflare writes your database_id into that repo's wrangler.toml while provisioning, so no .env is needed (if it still shows the placeholder, do step 2):
git clone https://github.com/ < you > /docflare-ai.git && cd docflare-ai
bun install && bunx wrangler login
bun run db:migrate
bunx wrangler vectorize get docflare-index # expect dimensions: 384, metric: cosine
If the index is missing or has a different shape, run step 3 (delete it first with bunx wrangler vectorize delete docflare-index ) and redeploy. If the wizard did not ask for ADMIN_TOKEN , run step 4.
export WORKER_URL= " https://docflare-ai.<your-subdomain>.workers.dev "
export ADMIN_TOKEN= " the-token-you-set "
curl -X POST " $WORKER_URL /api/sites/index " \
-H " Authorization: Bearer $ADMIN_TOKEN " \
-H " Content-Type: application/json " \
-d ' {"siteId":"my-docs","sitemapUrl":"https://docs.example.com/sitemap.xml"} '
Response:
{
"siteId" : " my-docs " ,
"domain" : " docs.example.com " ,
"sitemapUrl" : " https://docs.example.com/sitemap.xml " ,
"queued" : 142 ,
"progress" : { "siteId" : " my-docs " , "total" : 142 , "pending" : 142 , "indexed" : 0 , "failed" : 0 },
"next" : " Pending pages are crawled 3 at a time by the cron trigger. Call POST /api/sites/my-docs/process repeatedly to finish faster. "
}
Prefer a UI? Open https://docflare-ai.<your-subdomain>.workers.dev/admin , paste the admin token and use the Register / re-index a site form. The same page shows progress bars per site, an Index all pending button that loops over /process until the queue is empty, and the recent questions log.
Pages are now crawled automatically, a few per minute, by the cron trigger. To finish faster, drive the queue yourself from the CLI:
until curl -sf -X POST " $WORKER_URL /api/sites/my-docs/process " \
-H " Authorization: Bearer $ADMIN_TOKEN " | grep -q ' "pending":0 ' ; do
sleep 1
done
Check progress at any time:
curl " $WORKER_URL /api/sites/my-docs/status " -H " Authorization: Bearer $ADMIN_TOKEN "
Re-running /api/sites/index for the same siteId re-queues every page, so a cron job or CI step can keep the index fresh. Vectorize is eventually consistent: freshly upserted chunks become searchable within a few seconds.
After upgrading DocFlare to a version that changes chunking (see the changelog in commit messages), re-index every site once ( Re-crawl sitemap in the admin UI or POST /api/sites/index ) so pages are re-chunked. Old vectors keep working until then.
Add one line before </body> on any page of the registered domain:
< script src =" https://docflare-ai.<your-subdomain>.workers.dev/widget.js "
data-site-id =" my-docs "
defer > </ script >
Widget options
Attribute
Default
Description
data-site-id
required
The siteId you indexed.
data-api
script origin
Base URL of the Worker, if you serve widget.js from elsewhere (e.g. a CDN).
data-theme
auto
light , dark or auto (follows prefers-color-scheme ).
data-color
#f6821f
Accent colour for the bubble and buttons.
data-position
right
right or left corner.
data-title
Ask AI
Panel header text.
data-placeholder
Ask a question about these docs…
Input placeholder.
data-welcome
Hi! Ask me anything…
First assistant message.
A tiny JavaScript API is exposed as window.DocFlare :
DocFlare . open ( ) ; // open the panel
DocFlare . close ( ) ;
DocFlare . ask ( 'How do I deploy?' ) ; // open and submit a question
Only pages served from the site's registered domain (and anything in ALLOWED_ORIGINS ) can call /api/chat for that siteId . Requests from ot

[truncated]

## Original Extract

Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. - p10node/docflare-ai

GitHub - p10node/docflare-ai: Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier. · GitHub
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
p10node
docflare-ai Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
Open-source AI chatbot & site search for your docs. One <script> tag, streaming answers with sources. 100% Cloudflare-native (Workers AI + Vectorize + D1), runs free on the free tier.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
docs docs migrations migrations public public scripts scripts src src test test widget widget .dev.vars.example .dev.vars.example .gitignore .gitignore LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json wrangler.toml wrangler.toml View all files Repository files navigation
Open-source, zero-cost AI site search & chatbot widget. 100% Cloudflare native.
Drop one <script> tag on your docs site and get a streaming RAG chatbot that answers from your content, with sources.
No OpenAI key. No Pinecone. No servers. No bill.
Chat widget embedded in a docs site
Admin UI ( /admin )
Streaming answers with source chips on the left; the built-in admin page for registering sitemaps, driving indexing and reviewing what users ask on the right.
Hosted "chat with your docs" widgets (Chatbase, CustomGPT, Mendable, ...) charge $20-$400/month for something that is, under the hood, a crawler, an embedding model, a vector store and an LLM. Cloudflare now ships every one of those pieces with a generous free tier:
DocFlare AI wires them together into a single Worker you deploy with one command.
One-tag embed : <script src=".../widget.js" data-site-id="my-docs"> . Vanilla JS, Shadow DOM, < 10 KB, zero dependencies, no CSS leaks.
Streaming answers over Server-Sent Events with source chips, in the language of the question.
Sitemap crawler : sitemap index support, gzip sitemaps, boilerplate stripping (nav/header/footer/scripts) via the native HTMLRewriter .
Section-aware chunking : headings start new chunks, so a landing page's "Open source" card is not buried in a chunk about pricing. Chunks are ~350 tokens with 40-token overlap, sized with a per-script token estimate so Vietnamese/CJK pages stay inside the embedding model's 512-token window.
Two-stage retrieval : the 20 best vector matches are re-scored by the bge-reranker-base cross-encoder before the top-K go to the LLM. Questions that merely share vocabulary with the wrong page (the site name on legal pages, a product name) land on the paragraph that actually answers them.
Multi-site / multi-tenant : one deployment can serve many documentation sites, each isolated in its own Vectorize namespace and locked to its own domain.
Free-tier safe by design : indexing runs in small resumable batches (cron + on-demand), the LLM is skipped when nothing relevant is retrieved, and AI rate limits degrade gracefully.
Dark / light / auto theme , custom accent colour, left/right position, keyboard friendly.
Chat history in D1 so you can see what users ask and where your docs have gaps.
Built-in admin UI at /admin : register sitemaps, watch indexing progress, process batches, re-crawl, delete sites, copy the embed snippet and browse recent questions. Single static HTML file, no build step.
TypeScript end-to-end, Hono router, strict types, no any . Bun for tooling and tests.
flowchart LR
subgraph Host["Your docs site"]
W["widget.js<br/>(Shadow DOM)"]
end
subgraph CF["Cloudflare (free tier)"]
direction TB
API["Worker + Hono<br/>src/index.ts"]
AI["Workers AI<br/>bge-small-en-v1.5<br/>bge-reranker-base<br/>llama-3.1-8b-instruct-fp8"]
VX["Vectorize<br/>384-dim cosine<br/>namespace = siteId"]
D1["D1<br/>sites / pages / queries"]
CRON["Cron trigger<br/>every minute"]
ASSETS["Static assets<br/>public/"]
end
SM["sitemap.xml + HTML pages"]
W -- "POST /api/chat (SSE)" --> API
ASSETS -- "GET /widget.js" --> W
API -- "embed question" --> AI
API -- "topK query" --> VX
API -- "RAG prompt" --> AI
API -- "log query" --> D1
CRON -- "process pending pages" --> API
API -- "fetch + extract" --> SM
API -- "embed chunks" --> AI
API -- "upsert vectors" --> VX
Loading
Indexing : POST /api/sites/index reads the sitemap and queues every HTML URL in D1. A cron trigger (or repeated calls to /process ) then crawls a few pages per invocation: extract text, chunk, embed, upsert to Vectorize, mark the page indexed.
Chat : the widget sends the question; the Worker embeds it, pulls the 20 closest chunks from the site's namespace, reranks them with a cross-encoder, keeps the top-K (max 2 per page), builds a grounded prompt, streams Llama 3.1's answer back as SSE and logs the exchange in D1.
See docs/ARCHITECTURE.md for the full design and docs/API_REFERENCE.md for the endpoints.
A free Cloudflare account with Workers enabled.
Bun 1.1+ (used for installing, scripts, building the widget and running tests).
git clone https://github.com/p10node/docflare-ai.git
cd docflare-ai
bun install
bunx wrangler login
2. Create the D1 database
bunx wrangler d1 create docflare-db
Copy the database_id from the output. Either paste it into wrangler.toml :
[[ d1_databases ]]
binding = " DB "
database_name = " docflare-db "
database_id = " xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx " # <- paste here
…or, to keep it out of git, put it in a .env file instead (useful when you push your fork):
echo ' D1_DATABASE_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx ' > .env
With .env set, bun run deploy and bun run db:migrate generate wrangler.local.toml (gitignored) from wrangler.toml with the id filled in and use that. wrangler.toml keeps the placeholder.
bun run db:migrate
3. Create the Vectorize index
bunx wrangler vectorize create docflare-index --dimensions=384 --metric=cosine
(384 dimensions matches @cf/baai/bge-small-en-v1.5 . The index name is already set in wrangler.toml .)
The management endpoints are protected by a bearer token. Pick something long and random:
bunx wrangler secret put ADMIN_TOKEN
5. Deploy
bun run deploy
bun run deploy builds the widget and deploys (the schema was applied in step 2; re-run bun run db:migrate whenever migrations/ gains a new file). Wrangler prints your Worker URL, e.g. https://docflare-ai.<your-subdomain>.workers.dev . Open it to see the demo page with the widget mounted.
Deploy to Cloudflare button : the button at the top of this README clones the repo into your GitHub account, provisions the Worker, D1 database and Vectorize index from wrangler.toml , asks for ADMIN_TOKEN (from .dev.vars.example ) and runs bun run deploy . Two values the wizard cannot read from wrangler.toml :
Vectorize index : enter Dimensions = 384 and Metric = cosine (the form leaves them blank; any other shape breaks indexing).
ADMIN_TOKEN : replace the example value with something long and random, e.g. openssl rand -hex 32 .
The build does not apply the D1 schema (the build token has no D1 access), so run the migration once from your new repo after the first deploy. Cloudflare writes your database_id into that repo's wrangler.toml while provisioning, so no .env is needed (if it still shows the placeholder, do step 2):
git clone https://github.com/ < you > /docflare-ai.git && cd docflare-ai
bun install && bunx wrangler login
bun run db:migrate
bunx wrangler vectorize get docflare-index # expect dimensions: 384, metric: cosine
If the index is missing or has a different shape, run step 3 (delete it first with bunx wrangler vectorize delete docflare-index ) and redeploy. If the wizard did not ask for ADMIN_TOKEN , run step 4.
export WORKER_URL= " https://docflare-ai.<your-subdomain>.workers.dev "
export ADMIN_TOKEN= " the-token-you-set "
curl -X POST " $WORKER_URL /api/sites/index " \
-H " Authorization: Bearer $ADMIN_TOKEN " \
-H " Content-Type: application/json " \
-d ' {"siteId":"my-docs","sitemapUrl":"https://docs.example.com/sitemap.xml"} '
Response:
{
"siteId" : " my-docs " ,
"domain" : " docs.example.com " ,
"sitemapUrl" : " https://docs.example.com/sitemap.xml " ,
"queued" : 142 ,
"progress" : { "siteId" : " my-docs " , "total" : 142 , "pending" : 142 , "indexed" : 0 , "failed" : 0 },
"next" : " Pending pages are crawled 3 at a time by the cron trigger. Call POST /api/sites/my-docs/process repeatedly to finish faster. "
}
Prefer a UI? Open https://docflare-ai.<your-subdomain>.workers.dev/admin , paste the admin token and use the Register / re-index a site form. The same page shows progress bars per site, an Index all pending button that loops over /process until the queue is empty, and the recent questions log.
Pages are now crawled automatically, a few per minute, by the cron trigger. To finish faster, drive the queue yourself from the CLI:
until curl -sf -X POST " $WORKER_URL /api/sites/my-docs/process " \
-H " Authorization: Bearer $ADMIN_TOKEN " | grep -q ' "pending":0 ' ; do
sleep 1
done
Check progress at any time:
curl " $WORKER_URL /api/sites/my-docs/status " -H " Authorization: Bearer $ADMIN_TOKEN "
Re-running /api/sites/index for the same siteId re-queues every page, so a cron job or CI step can keep the index fresh. Vectorize is eventually consistent: freshly upserted chunks become searchable within a few seconds.
After upgrading DocFlare to a version that changes chunking (see the changelog in commit messages), re-index every site once ( Re-crawl sitemap in the admin UI or POST /api/sites/index ) so pages are re-chunked. Old vectors keep working until then.
Add one line before </body> on any page of the registered domain:
< script src =" https://docflare-ai.<your-subdomain>.workers.dev/widget.js "
data-site-id =" my-docs "
defer > </ script >
Widget options
Attribute
Default
Description
data-site-id
required
The siteId you indexed.
data-api
script origin
Base URL of the Worker, if you serve widget.js from elsewhere (e.g. a CDN).
data-theme
auto
light , dark or auto (follows prefers-color-scheme ).
data-color
#f6821f
Accent colour for the bubble and buttons.
data-position
right
right or left corner.
data-title
Ask AI
Panel header text.
data-placeholder
Ask a question about these docs…
Input placeholder.
data-welcome
Hi! Ask me anything…
First assistant message.
A tiny JavaScript API is exposed as window.DocFlare :
DocFlare . open ( ) ; // open the panel
DocFlare . close ( ) ;
DocFlare . ask ( 'How do I deploy?' ) ; // open and submit a question
Only pages served from the site's registered domain (and anything in ALLOWED_ORIGINS ) can call /api/chat for that siteId . Requests from ot

[truncated]
