---
source: "https://github.com/crawlerpeek/crawlerpeek"
hn_url: "https://news.ycombinator.com/item?id=50236098"
title: "Crawlerpeek – check which AI crawlers your site lets in (CLI, MCP)"
article_title: "GitHub - crawlerpeek/crawlerpeek: Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. · GitHub"
image: "https://opengraph.githubassets.com/d0ee170751f9a8d6c7f3a736a7c65fe1e988ead8be077f12aa91b50121b08659/crawlerpeek/crawlerpeek"
author: "crawlerpeek"
captured_at: "2026-10-25T21:28:56Z"
capture_tool: "hn-digest"
hn_id: 50236098
score: 1
comments: 0
posted_at: "2026-10-25T21:11:05Z"
tags:
  - hacker-news
---

# Crawlerpeek – check which AI crawlers your site lets in (CLI, MCP)

- HN: [50236098](https://news.ycombinator.com/item?id=50236098)
- Source: [github.com](https://github.com/crawlerpeek/crawlerpeek)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T21:11:05Z

## Translation

Title: Crawlerpeek – check which AI crawlers your site lets in (CLI, MCP)
Article title: GitHub - crawlerpeek/crawlerpeek: Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. · GitHub
Description: Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. - crawlerpeek/crawlerpeek

Article text:
GitHub - crawlerpeek/crawlerpeek: Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. · GitHub
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
crawlerpeek
crawlerpeek Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github .github action action src src test test .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build.mjs build.mjs package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.cjs.json tsconfig.cjs.json tsconfig.json tsconfig.json vitest.config.mts vitest.config.mts View all files Repository files navigation
Check whether AI crawlers can access your site. For each bot (GPTBot, OAI-SearchBot, ChatGPT-User, ClaudeBot, Claude-SearchBot, PerplexityBot, Google-Extended, Googlebot, Bingbot, CCBot and more) it tells you:
whether robots.txt allows the URL's path for that bot (RFC 9309: user-agent groups, wildcards, $ , longest match wins), and
what the page actually returns when fetched with that bot's user-agent string: status, final URL after redirects, content type, size, title, and whether the response looks like a block or challenge page.
It also generates a robots.txt that welcomes AI crawlers. A library, a CLI for CI, and an MCP server. Node 20 or newer, no runtime dependencies, ESM and CommonJS, TypeScript types included.
The same checks as a web tool: crawlerpeek.com . Documentation with examples for CI and MCP clients: crawlerpeek.com/cli .
npm install crawlerpeek # library
npx crawlerpeek --help # CLI, no install needed
CLI
npx crawlerpeek check https://example.com
https://example.com
robots.txt: ok (HTTP 200) at https://example.com/robots.txt
Bot robots HTTP
GPTBot blocked 200 FAIL
ClaudeBot allowed 403 (block page) FAIL
Google-Extended allowed -
npx crawlerpeek check https://example.com/pricing --bots GPTBot,ClaudeBot --fail-on blocked
npx crawlerpeek check https://example.com --bots all --json
npx crawlerpeek check https://example.com --no-http # robots.txt only
npx crawlerpeek see https://example.com --bot ClaudeBot # one bot, full detail
npx crawlerpeek robots --disallow /admin/,/cart/ --sitemap https://example.com/sitemap.xml
Exit codes: 0 ok, 1 the --fail-on condition was met, 2 usage error (including an unknown option) or request error. see exits 0 when the page was fetched, whatever its HTTP status, and 2 when the request failed or was refused.
Without --fail-on the exit code is 0 whatever the result. --bots takes tokens or all ; the default is 15 bots: the main AI crawlers plus Googlebot and Bingbot. A robots.txt that cannot be fetched (timeout, 5xx, a redirect that is refused or has no target) is reported as unknown , with the reason in robotsTxt.error , and never counts as blocked. Note: this deliberately differs from RFC 9309 §2.3.1.4, where crawlers treat an unreachable robots.txt (5xx) as "disallow everything" (Google does too), so a site whose robots.txt returns 5xx is effectively closed to many bots even though --fail-on blocked passes. Treat unknown as a problem to fix.
import { checkAiAccess , seeAsBot , generateRobotsTxt , failingBots } from "crawlerpeek" ;
const report = await checkAiAccess ( "https://example.com/pricing" , { bots : [ "GPTBot" , "ClaudeBot" ] } ) ;
report . bots [ 0 ] . robots ; // { access: "allowed", pathAllowed: true, matchedGroup: "*" }
report . bots [ 0 ] . http ; // { status: 200, finalUrl, redirects, contentType, bytes, title, looksBlocked, outcome: "ok", ... }
failingBots ( report , "blocked" ) ; // ["ClaudeBot"] when it is closed or challenged
await seeAsBot ( "https://example.com" , "PerplexityBot" ) ;
generateRobotsTxt ( { disallow : [ "/admin/" ] , blockBots : [ "CCBot" ] , sitemap : "https://example.com/sitemap.xml" } ) ;
generateRobotsTxt names each AI bot explicitly with Allow: / . Disallowed paths are repeated in that group on purpose: a bot with its own group ignores the * group.
crawlerpeek mcp speaks the Model Context Protocol over stdio and exposes three read-only tools: check_ai_access , see_as_bot , generate_robots_txt .
{ "mcpServers" : { "crawlerpeek" : { "command" : " npx " , "args" : [ " -y " , " crawlerpeek " , " mcp " ] } } }
Results are cached for 10 minutes per call. No LLM is called; the tools only fetch the URL you give them.
A composite action lives in action/ ; see action/README.md .
- uses : crawlerpeek/crawlerpeek/action@v0
with : { url: "https://example.com", fail-on: blocked }
What you should know
see-as-bot sends the bot's own user-agent string from your machine. It is not the real bot and not from the operator's IP ranges. Sites that verify bot IPs (Googlebot, GPTBot and others publish theirs) may treat a request that only claims to be the bot differently, often by blocking it. A 403 for Googlebot here can therefore be normal; a 403 for an unusual bot with a 200 for a browser usually points at a WAF or CDN rule. robots.txt itself is fetched with an honest crawlerpeek/<version> user agent.
Google-Extended and Applebot-Extended are robots.txt control tokens, not crawlers: they get a robots.txt decision but are never fetched.
Pages are fetched as raw HTML. No JavaScript runs.
robots.txt is a request, not access control. Some user-triggered fetchers (ChatGPT-User, Perplexity-User) may not follow it; the operators' documentation says which.
The bot list and user-agent strings come from the operators' documentation and can drift; the list was last compared with the operators' documentation on 2026-10-18.
One request per bot plus one for robots.txt, sequential with a 300 ms pause ( --delay ); 10 s timeout for the whole request chain; at most 5 redirects, each re-checked (a redirect to a refused target is reported for that bot or for robots.txt, it does not stop the run); response bodies capped at 2 MB; no cookies. Private, loopback, link-local and metadata addresses are refused after DNS resolution (and again at connect time), as are ports other than 80 and 443 and URLs with credentials. --allow-private (CLI and crawlerpeek mcp ) or allowPrivate: true lifts these limits for local testing.
Issues and pull requests are welcome, see CONTRIBUTING.md . Report security problems privately, see SECURITY.md .
Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. - crawlerpeek/crawlerpeek

GitHub - crawlerpeek/crawlerpeek: Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action. · GitHub
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
crawlerpeek
crawlerpeek Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github .github action action src src test test .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build.mjs build.mjs package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.cjs.json tsconfig.cjs.json tsconfig.json tsconfig.json vitest.config.mts vitest.config.mts View all files Repository files navigation
Check whether AI crawlers can access your site. For each bot (GPTBot, OAI-SearchBot, ChatGPT-User, ClaudeBot, Claude-SearchBot, PerplexityBot, Google-Extended, Googlebot, Bingbot, CCBot and more) it tells you:
whether robots.txt allows the URL's path for that bot (RFC 9309: user-agent groups, wildcards, $ , longest match wins), and
what the page actually returns when fetched with that bot's user-agent string: status, final URL after redirects, content type, size, title, and whether the response looks like a block or challenge page.
It also generates a robots.txt that welcomes AI crawlers. A library, a CLI for CI, and an MCP server. Node 20 or newer, no runtime dependencies, ESM and CommonJS, TypeScript types included.
The same checks as a web tool: crawlerpeek.com . Documentation with examples for CI and MCP clients: crawlerpeek.com/cli .
npm install crawlerpeek # library
npx crawlerpeek --help # CLI, no install needed
CLI
npx crawlerpeek check https://example.com
https://example.com
robots.txt: ok (HTTP 200) at https://example.com/robots.txt
Bot robots HTTP
GPTBot blocked 200 FAIL
ClaudeBot allowed 403 (block page) FAIL
Google-Extended allowed -
npx crawlerpeek check https://example.com/pricing --bots GPTBot,ClaudeBot --fail-on blocked
npx crawlerpeek check https://example.com --bots all --json
npx crawlerpeek check https://example.com --no-http # robots.txt only
npx crawlerpeek see https://example.com --bot ClaudeBot # one bot, full detail
npx crawlerpeek robots --disallow /admin/,/cart/ --sitemap https://example.com/sitemap.xml
Exit codes: 0 ok, 1 the --fail-on condition was met, 2 usage error (including an unknown option) or request error. see exits 0 when the page was fetched, whatever its HTTP status, and 2 when the request failed or was refused.
Without --fail-on the exit code is 0 whatever the result. --bots takes tokens or all ; the default is 15 bots: the main AI crawlers plus Googlebot and Bingbot. A robots.txt that cannot be fetched (timeout, 5xx, a redirect that is refused or has no target) is reported as unknown , with the reason in robotsTxt.error , and never counts as blocked. Note: this deliberately differs from RFC 9309 §2.3.1.4, where crawlers treat an unreachable robots.txt (5xx) as "disallow everything" (Google does too), so a site whose robots.txt returns 5xx is effectively closed to many bots even though --fail-on blocked passes. Treat unknown as a problem to fix.
import { checkAiAccess , seeAsBot , generateRobotsTxt , failingBots } from "crawlerpeek" ;
const report = await checkAiAccess ( "https://example.com/pricing" , { bots : [ "GPTBot" , "ClaudeBot" ] } ) ;
report . bots [ 0 ] . robots ; // { access: "allowed", pathAllowed: true, matchedGroup: "*" }
report . bots [ 0 ] . http ; // { status: 200, finalUrl, redirects, contentType, bytes, title, looksBlocked, outcome: "ok", ... }
failingBots ( report , "blocked" ) ; // ["ClaudeBot"] when it is closed or challenged
await seeAsBot ( "https://example.com" , "PerplexityBot" ) ;
generateRobotsTxt ( { disallow : [ "/admin/" ] , blockBots : [ "CCBot" ] , sitemap : "https://example.com/sitemap.xml" } ) ;
generateRobotsTxt names each AI bot explicitly with Allow: / . Disallowed paths are repeated in that group on purpose: a bot with its own group ignores the * group.
crawlerpeek mcp speaks the Model Context Protocol over stdio and exposes three read-only tools: check_ai_access , see_as_bot , generate_robots_txt .
{ "mcpServers" : { "crawlerpeek" : { "command" : " npx " , "args" : [ " -y " , " crawlerpeek " , " mcp " ] } } }
Results are cached for 10 minutes per call. No LLM is called; the tools only fetch the URL you give them.
A composite action lives in action/ ; see action/README.md .
- uses : crawlerpeek/crawlerpeek/action@v0
with : { url: "https://example.com", fail-on: blocked }
What you should know
see-as-bot sends the bot's own user-agent string from your machine. It is not the real bot and not from the operator's IP ranges. Sites that verify bot IPs (Googlebot, GPTBot and others publish theirs) may treat a request that only claims to be the bot differently, often by blocking it. A 403 for Googlebot here can therefore be normal; a 403 for an unusual bot with a 200 for a browser usually points at a WAF or CDN rule. robots.txt itself is fetched with an honest crawlerpeek/<version> user agent.
Google-Extended and Applebot-Extended are robots.txt control tokens, not crawlers: they get a robots.txt decision but are never fetched.
Pages are fetched as raw HTML. No JavaScript runs.
robots.txt is a request, not access control. Some user-triggered fetchers (ChatGPT-User, Perplexity-User) may not follow it; the operators' documentation says which.
The bot list and user-agent strings come from the operators' documentation and can drift; the list was last compared with the operators' documentation on 2026-10-18.
One request per bot plus one for robots.txt, sequential with a 300 ms pause ( --delay ); 10 s timeout for the whole request chain; at most 5 redirects, each re-checked (a redirect to a refused target is reported for that bot or for robots.txt, it does not stop the run); response bodies capped at 2 MB; no cookies. Private, loopback, link-local and metadata addresses are refused after DNS resolution (and again at connect time), as are ports other than 80 and 443 and URLs with credentials. --allow-private (CLI and crawlerpeek mcp ) or allowPrivate: true lifts these limits for local testing.
Issues and pull requests are welcome, see CONTRIBUTING.md . Report security problems privately, see SECURITY.md .
Check whether AI crawlers (GPTBot, ClaudeBot, PerplexityBot and more) can access your site. CLI, library, MCP server and GitHub Action.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
