---
source: "https://beamfar.com/guides/web-security-and-ai-crawler-benchmark-2026"
hn_url: "https://news.ycombinator.com/item?id=50641988"
title: "We scanned 482 live sites: 35% publish llms.txt, 58% ship zero security headers"
article_title: "2026 Web Security Headers, DMARC and AI Crawler Benchmark (482 Sites) | Beamfar"
image: "https://beamfar.com/static/site/img/og.png"
author: "corchesrares"
captured_at: "2026-11-25T13:56:26Z"
capture_tool: "hn-digest"
hn_id: 50641988
score: 1
comments: 0
posted_at: "2026-11-25T13:23:38Z"
tags:
  - hacker-news
---

# We scanned 482 live sites: 35% publish llms.txt, 58% ship zero security headers

- HN: [50641988](https://news.ycombinator.com/item?id=50641988)
- Source: [beamfar.com](https://beamfar.com/guides/web-security-and-ai-crawler-benchmark-2026)
- Score: 1
- Comments: 0
- Posted: 2026-11-25T13:23:38Z

## Translation

Title: We scanned 482 live sites: 35% publish llms.txt, 58% ship zero security headers
Article title: 2026 Web Security Headers, DMARC and AI Crawler Benchmark (482 Sites) | Beamfar
Description: Live passive measurements across 482 public business, agency and SaaS websites on 25 November 2026: 58.5% ship zero security headers, 78.4% leave DMARC unenforced, while 35.3% already serve llms.txt.

Article text:
Skip to content
Beamfar
Features Monitoring
Uptime monitoring Checked from several regions; one alert per outage, with what each saw.
SSL certificates Expiry, issuer and replacement, read on every page check.
Domain expiry New Registration dates from RDAP or WHOIS, with alerts as they get close.
DNS changes Record changes, with the incident that follows.
SEO monitoring Noindex tags, robots.txt blocks, off-host canonicals and site crawl.
Heartbeats Cron jobs and backups that ping after each run.
Page speed New Load time from every region; Google’s Web Vitals where published.
Server monitoring New CPU, memory, disk and load from a small agent.
Exploit Watch New Active exploit alerts, fix cards, connected layers and Exploit Watch+ per site.
Security checks and headers New Script hosts, hidden links and missing headers, read from outside.
AI security watch New DNS hijacks and exposed services, through your providers.
Stack and vulnerability watch New The software and versions a site shows, checked against published advisories.
Approved auto-repair Coming soon Fixes you approve per site, after backup verification.
Alert channels New E-mail and push, Slack and nine more chat and on-call tools, signed webhooks from Team.
Two-region confirmation Down only when two regions agree: no alert for a blip.
Why this happened The likely cause, its evidence and what changed just before.
Acknowledge and escalate Show that someone is on it; escalate from Team.
Maintenance windows Planned work holds alerts back and stays out of uptime.
Status pages A public page that never shows a day without data as green.
Monthly client reports One report per client each month, from stored checks.
Client portal Read-only access for a client contact, without a seat.
White label Your name, logo and colour on reports and status pages.
MCP server New Fifteen read-only tools for your AI assistant, with your key.
REST API New Your sites, incidents and reports over HTTPS; writing from Agency.
Ask Beam Answers from the product’s own pages, or says it does not know.
Every feature, the plans that include it and what is coming next. All features
Web agencies Every client site on one screen, a report per client.
Freelancers A handful of client sites, looked after without a pager.
Online shops The shop, its certificate and its domain, watched from outside.
SaaS and developers Heartbeats, API checks, and an API of your own.
Small businesses Get alerted with the likely cause when your site or SSL breaks.
Security checks and headers New
Stack and vulnerability watch New
Approved auto-repair Coming soon
2026 Web Security Headers, DMARC and AI Crawler Benchmark
By the Beamfar team. Published 25 November 2026 . Updated 25 November 2026 . 7 min read.
On 25 November 2026 we ran a strictly passive, read-only scan of 500 unique public business, web-agency, WordPress and SaaS domains across Europe and North America (482 responsive 2xx/3xx websites). The headline gap is stark: 35.3% of live sites already publish /llms.txt to court AI search engines, yet 58.5% ship zero of the five core browser security headers , 78.4% fail to enforce DMARC , and only 2.5% publish /.well-known/security.txt .
Key findings at a glance (482 live sites, 25 November 2026)
1. HTTPS is universal, but 58.5% ship zero browser security headers
2. 35.3% already publish llms.txt , while only 2.3% block AI crawlers
3. Email authentication: 92.7% have SPF, but 78.4% do not enforce DMARC
4. security.txt adoption sits at 2.5%, and slow TTFB affects 40.9%
Practical 15-minute hardening checklist
Key findings at a glance (482 live sites, 25 November 2026)
Every percentage below is measured directly from live HTTP responses and public DNS TXT records on 25 November 2026. All findings are aggregated and anonymized; no individual domain is named.
Check it now. The security headers checker reads a page’s HSTS, CSP, X-Content-Type-Options, frame protection, Referrer-Policy and Permissions-Policy; the indexability checker reads a page’s noindex, robots.txt, canonical address and redirects. No account, and the address is not stored.
1. HTTPS is universal, but 58.5% ship zero browser security headers
All 482 responsive websites redirected or loaded over HTTPS ( 100.0% ). Free automated TLS certificates have solved baseline transport encryption, yet response-level browser hardening has barely moved outside dedicated engineering teams.
We checked five baseline headers recommended by the OWASP Secure Headers Project : Strict-Transport-Security (HSTS), Content-Security-Policy (CSP), X-Content-Type-Options: nosniff , clickjacking protection ( X-Frame-Options or CSP frame-ancestors ), and Referrer-Policy .
Per-header gaps across the 482 sites:
Missing Content-Security-Policy : 84.9% ( 409 sites)
Missing Referrer-Policy : 78.8% ( 380 sites)
Missing Strict-Transport-Security (HSTS): 75.9% ( 366 sites)
Missing clickjacking protection ( X-Frame-Options or frame-ancestors ): 74.3% ( 358 sites)
Missing X-Content-Type-Options: nosniff : 69.3% ( 334 sites)
Version banner disclosure ( Server: nginx/x.y or X-Powered-By ): 31.3% ( 151 sites)
Because 63.3% ( 305 ) of the cohort runs WordPress on managed or shared hosting, default server blocks rarely add HSTS, nosniff , or Referrer-Policy unless an agency configures them at the reverse proxy or CDN edge. You can test your own live response headers in seconds with our free security headers checker .
2. 35.3% already publish llms.txt , while only 2.3% block AI crawlers
The most surprising result of our November 2026 scan is how quickly /llms.txt has spread compared with explicit robots.txt crawler controls:
Valid robots.txt present: 94.8% ( 457 sites); 82.8% declare a Sitemap: directive; 5.2% return 404 or an HTML page.
Homepage SEO hygiene: 11.8% omit <link rel="canonical"> on their homepage, and 1 live production site served an active noindex directive on its homepage.
Explicit AI crawler rules in robots.txt : only 10.4% ( 50 sites) name any AI bot ( GPTBot , OAI-SearchBot , ClaudeBot , anthropic-ai , PerplexityBot , Google-Extended , CCBot , or Bytespider ).
Sites blocking at least one AI crawler: only 2.3% ( 11 sites). Among those 11 sites, the most frequently blocked user-agents were Bytespider ( 9 ), CCBot ( 7 ), GPTBot ( 7 ), anthropic-ai ( 6 ), ClaudeBot ( 5 ), Google-Extended ( 5 ), and PerplexityBot ( 1 ).
Valid /llms.txt present: 35.3% ( 170 sites) return a real plain-text or Markdown /llms.txt file. Spot checks of the responses show rapid adoption driven largely by WordPress SEO plugins (Rank Math SEO, All in One SEO, and Yoast) alongside custom Markdown files on agency and SaaS sites.
In short, commercial websites in late 2026 overwhelmingly want to be discovered by AI assistants rather than blocked — yet 89.6% still leave robots.txt silent on AI crawlers while 58.5% omit every browser security header. Verify your own robots.txt , canonical tags, and noindex status with our free indexability checker .
3. Email authentication: 92.7% have SPF, but 78.4% do not enforce DMARC
Domain spoofing hurts deliverability for transactional receipts, alerts, and client outreach alike. Across the 482 responsive domains, public DNS TXT records showed a classic half-finished rollout of RFC 7489 DMARC :
Nearly half of all domains ( 47.5% ) published a p=none record to pass bulk-sender checks in 2024–2025 and never came back to tighten it to p=quarantine or p=reject . Check your current DNS and mail records with our free DNS lookup tool .
4. security.txt adoption sits at 2.5%, and slow TTFB affects 40.9%
Two operational metrics stood out alongside headers and DNS:
/.well-known/security.txt ( RFC 9116 ): only 2.5% ( 12 / 482 ) publish a valid security contact file. If a researcher finds an exposed key, broken redirect, or vulnerable plugin on the remaining 97.5% , there is no standard machine-readable address to report it to.
Public /openapi.json schemas: 1 live site returned a publicly readable OpenAPI specification with route definitions. While rare on marketing sites, unauthenticated API documentation is far more common on newly deployed SaaS apps; see our Cursor, Lovable and Bolt security checklist for hardening API routes and hiding source maps.
First-byte response speed (measured from Europe): median TTFB was 704 ms , while the 90th percentile reached 3,382 ms . Overall, 40.9% of sites took 1,000 ms or longer to send the first byte, and 19.1% took 2,000 ms or longer — typically uncached dynamic PHP pages or distant single-region origins without edge caching.
Practical 15-minute hardening checklist
If you run a SaaS product or manage client websites at an agency, five low-risk changes move any site into the top 8% of this benchmark:
Add four safe baseline headers at the edge or reverse proxy (Caddy, Nginx, Cloudflare Transform Rules, or Vercel headers ):
- Strict-Transport-Security: max-age=31536000; includeSubDomains (once every subdomain uses HTTPS) - X-Content-Type-Options: nosniff - X-Frame-Options: SAMEORIGIN - Referrer-Policy: strict-origin-when-cross-origin
Roll out Content-Security-Policy in report-only mode first ( Content-Security-Policy-Report-Only ) so third-party embeds, fonts, or payment frames do not break when you enforce it.
Strip version banners: disable server_tokens in Nginx, unset X-Powered-By in PHP/Express/Next.js, and verify responses no longer broadcast patch versions.
Move DMARC past p=none : review aggregate reports ( rua= ) for 14 days to confirm every legitimate sender passes SPF or DKIM, then step up to p=quarantine; pct=25 , p=quarantine; pct=100 , and finally p=reject .
Publish /.well-known/security.txt and audit robots.txt : add a two-line Contact: + Expires: file under /.well-known/security.txt , confirm robots.txt declares your sitemap and reflects your actual AI crawler policy, and check that no deploy left a noindex tag on production.
Date: 25 November 2026 (single pass, strictly read-only GET requests + public DNS TXT queries).
Sample: 500 unique public domains drawn from European and North American web agencies, WordPress care studios, SMB commercial sites, and SaaS teams; 18 timed out or returned non-2xx/3xx errors, leaving 482 responsive websites ( 305 WordPress, 23 Next.js, 154 other stacks).
Checks performed: landing-page HTTPS status, five OWASP browser security headers, Server / X-Powered-By version banners, homepage noindex and canonical tags, /robots.txt AI crawler directives, /llms.txt plain-text/Markdown presence, /.well-known/security.txt ( Contact: line), /openapi.json JSON schema check, DNS SPF and _dmarc TXT records, and first-byte timing from a European vantage point.
Limitations: this cohort skews toward active agency, WordPress, and SMB sites ( 63.3% WordPress) rather than the Alexa/Tranco top 1,000, so enterprise header adoption is higher than this sample. Passive checks cannot see private repositories, authenticated dashboards, or server-side WAF rules.
Beamfar runs these external checks continuously after every deploy: multi-region uptime confirmation, TLS and domain expiry warnings, DNS and DMARC change alerts, SEO-risk alerts ( noindex , robots.txt , canonical drift), passive frontend secret leak detection across 12 credential families, exposed API schema checks, and continuous lockfile/SBOM CISA KEV Exploit Watch across eight ecosystems. Run a one-off check with zero storage using our free site scan , security headers checker , frontend secret leak checker , or lockfile CVE checker — or start on the Free plan (one site, no card, no time limit) and connect Claude, Cursor, Codex, or Gemini CLI via our 16-tool read-only MCP server .
Read on 25 November 2026 . Standards and vendor pages change; confirm the current text before you rely on it.
OWASP Secure Headers Project (owasp.org)
RFC 9116 security.txt specification (www.rfc-editor.org)
RFC 7489 DMARC specification (www.rfc-editor.org)
OpenAI crawler documentation (GPTBot and OAI-SearchBot) (platform.openai.com)
Anthropic crawler document

[truncated]

## Original Extract

Live passive measurements across 482 public business, agency and SaaS websites on 25 November 2026: 58.5% ship zero security headers, 78.4% leave DMARC unenforced, while 35.3% already serve llms.txt.

Skip to content
Beamfar
Features Monitoring
Uptime monitoring Checked from several regions; one alert per outage, with what each saw.
SSL certificates Expiry, issuer and replacement, read on every page check.
Domain expiry New Registration dates from RDAP or WHOIS, with alerts as they get close.
DNS changes Record changes, with the incident that follows.
SEO monitoring Noindex tags, robots.txt blocks, off-host canonicals and site crawl.
Heartbeats Cron jobs and backups that ping after each run.
Page speed New Load time from every region; Google’s Web Vitals where published.
Server monitoring New CPU, memory, disk and load from a small agent.
Exploit Watch New Active exploit alerts, fix cards, connected layers and Exploit Watch+ per site.
Security checks and headers New Script hosts, hidden links and missing headers, read from outside.
AI security watch New DNS hijacks and exposed services, through your providers.
Stack and vulnerability watch New The software and versions a site shows, checked against published advisories.
Approved auto-repair Coming soon Fixes you approve per site, after backup verification.
Alert channels New E-mail and push, Slack and nine more chat and on-call tools, signed webhooks from Team.
Two-region confirmation Down only when two regions agree: no alert for a blip.
Why this happened The likely cause, its evidence and what changed just before.
Acknowledge and escalate Show that someone is on it; escalate from Team.
Maintenance windows Planned work holds alerts back and stays out of uptime.
Status pages A public page that never shows a day without data as green.
Monthly client reports One report per client each month, from stored checks.
Client portal Read-only access for a client contact, without a seat.
White label Your name, logo and colour on reports and status pages.
MCP server New Fifteen read-only tools for your AI assistant, with your key.
REST API New Your sites, incidents and reports over HTTPS; writing from Agency.
Ask Beam Answers from the product’s own pages, or says it does not know.
Every feature, the plans that include it and what is coming next. All features
Web agencies Every client site on one screen, a report per client.
Freelancers A handful of client sites, looked after without a pager.
Online shops The shop, its certificate and its domain, watched from outside.
SaaS and developers Heartbeats, API checks, and an API of your own.
Small businesses Get alerted with the likely cause when your site or SSL breaks.
Security checks and headers New
Stack and vulnerability watch New
Approved auto-repair Coming soon
2026 Web Security Headers, DMARC and AI Crawler Benchmark
By the Beamfar team. Published 25 November 2026 . Updated 25 November 2026 . 7 min read.
On 25 November 2026 we ran a strictly passive, read-only scan of 500 unique public business, web-agency, WordPress and SaaS domains across Europe and North America (482 responsive 2xx/3xx websites). The headline gap is stark: 35.3% of live sites already publish /llms.txt to court AI search engines, yet 58.5% ship zero of the five core browser security headers , 78.4% fail to enforce DMARC , and only 2.5% publish /.well-known/security.txt .
Key findings at a glance (482 live sites, 25 November 2026)
1. HTTPS is universal, but 58.5% ship zero browser security headers
2. 35.3% already publish llms.txt , while only 2.3% block AI crawlers
3. Email authentication: 92.7% have SPF, but 78.4% do not enforce DMARC
4. security.txt adoption sits at 2.5%, and slow TTFB affects 40.9%
Practical 15-minute hardening checklist
Key findings at a glance (482 live sites, 25 November 2026)
Every percentage below is measured directly from live HTTP responses and public DNS TXT records on 25 November 2026. All findings are aggregated and anonymized; no individual domain is named.
Check it now. The security headers checker reads a page’s HSTS, CSP, X-Content-Type-Options, frame protection, Referrer-Policy and Permissions-Policy; the indexability checker reads a page’s noindex, robots.txt, canonical address and redirects. No account, and the address is not stored.
1. HTTPS is universal, but 58.5% ship zero browser security headers
All 482 responsive websites redirected or loaded over HTTPS ( 100.0% ). Free automated TLS certificates have solved baseline transport encryption, yet response-level browser hardening has barely moved outside dedicated engineering teams.
We checked five baseline headers recommended by the OWASP Secure Headers Project : Strict-Transport-Security (HSTS), Content-Security-Policy (CSP), X-Content-Type-Options: nosniff , clickjacking protection ( X-Frame-Options or CSP frame-ancestors ), and Referrer-Policy .
Per-header gaps across the 482 sites:
Missing Content-Security-Policy : 84.9% ( 409 sites)
Missing Referrer-Policy : 78.8% ( 380 sites)
Missing Strict-Transport-Security (HSTS): 75.9% ( 366 sites)
Missing clickjacking protection ( X-Frame-Options or frame-ancestors ): 74.3% ( 358 sites)
Missing X-Content-Type-Options: nosniff : 69.3% ( 334 sites)
Version banner disclosure ( Server: nginx/x.y or X-Powered-By ): 31.3% ( 151 sites)
Because 63.3% ( 305 ) of the cohort runs WordPress on managed or shared hosting, default server blocks rarely add HSTS, nosniff , or Referrer-Policy unless an agency configures them at the reverse proxy or CDN edge. You can test your own live response headers in seconds with our free security headers checker .
2. 35.3% already publish llms.txt , while only 2.3% block AI crawlers
The most surprising result of our November 2026 scan is how quickly /llms.txt has spread compared with explicit robots.txt crawler controls:
Valid robots.txt present: 94.8% ( 457 sites); 82.8% declare a Sitemap: directive; 5.2% return 404 or an HTML page.
Homepage SEO hygiene: 11.8% omit <link rel="canonical"> on their homepage, and 1 live production site served an active noindex directive on its homepage.
Explicit AI crawler rules in robots.txt : only 10.4% ( 50 sites) name any AI bot ( GPTBot , OAI-SearchBot , ClaudeBot , anthropic-ai , PerplexityBot , Google-Extended , CCBot , or Bytespider ).
Sites blocking at least one AI crawler: only 2.3% ( 11 sites). Among those 11 sites, the most frequently blocked user-agents were Bytespider ( 9 ), CCBot ( 7 ), GPTBot ( 7 ), anthropic-ai ( 6 ), ClaudeBot ( 5 ), Google-Extended ( 5 ), and PerplexityBot ( 1 ).
Valid /llms.txt present: 35.3% ( 170 sites) return a real plain-text or Markdown /llms.txt file. Spot checks of the responses show rapid adoption driven largely by WordPress SEO plugins (Rank Math SEO, All in One SEO, and Yoast) alongside custom Markdown files on agency and SaaS sites.
In short, commercial websites in late 2026 overwhelmingly want to be discovered by AI assistants rather than blocked — yet 89.6% still leave robots.txt silent on AI crawlers while 58.5% omit every browser security header. Verify your own robots.txt , canonical tags, and noindex status with our free indexability checker .
3. Email authentication: 92.7% have SPF, but 78.4% do not enforce DMARC
Domain spoofing hurts deliverability for transactional receipts, alerts, and client outreach alike. Across the 482 responsive domains, public DNS TXT records showed a classic half-finished rollout of RFC 7489 DMARC :
Nearly half of all domains ( 47.5% ) published a p=none record to pass bulk-sender checks in 2024–2025 and never came back to tighten it to p=quarantine or p=reject . Check your current DNS and mail records with our free DNS lookup tool .
4. security.txt adoption sits at 2.5%, and slow TTFB affects 40.9%
Two operational metrics stood out alongside headers and DNS:
/.well-known/security.txt ( RFC 9116 ): only 2.5% ( 12 / 482 ) publish a valid security contact file. If a researcher finds an exposed key, broken redirect, or vulnerable plugin on the remaining 97.5% , there is no standard machine-readable address to report it to.
Public /openapi.json schemas: 1 live site returned a publicly readable OpenAPI specification with route definitions. While rare on marketing sites, unauthenticated API documentation is far more common on newly deployed SaaS apps; see our Cursor, Lovable and Bolt security checklist for hardening API routes and hiding source maps.
First-byte response speed (measured from Europe): median TTFB was 704 ms , while the 90th percentile reached 3,382 ms . Overall, 40.9% of sites took 1,000 ms or longer to send the first byte, and 19.1% took 2,000 ms or longer — typically uncached dynamic PHP pages or distant single-region origins without edge caching.
Practical 15-minute hardening checklist
If you run a SaaS product or manage client websites at an agency, five low-risk changes move any site into the top 8% of this benchmark:
Add four safe baseline headers at the edge or reverse proxy (Caddy, Nginx, Cloudflare Transform Rules, or Vercel headers ):
- Strict-Transport-Security: max-age=31536000; includeSubDomains (once every subdomain uses HTTPS) - X-Content-Type-Options: nosniff - X-Frame-Options: SAMEORIGIN - Referrer-Policy: strict-origin-when-cross-origin
Roll out Content-Security-Policy in report-only mode first ( Content-Security-Policy-Report-Only ) so third-party embeds, fonts, or payment frames do not break when you enforce it.
Strip version banners: disable server_tokens in Nginx, unset X-Powered-By in PHP/Express/Next.js, and verify responses no longer broadcast patch versions.
Move DMARC past p=none : review aggregate reports ( rua= ) for 14 days to confirm every legitimate sender passes SPF or DKIM, then step up to p=quarantine; pct=25 , p=quarantine; pct=100 , and finally p=reject .
Publish /.well-known/security.txt and audit robots.txt : add a two-line Contact: + Expires: file under /.well-known/security.txt , confirm robots.txt declares your sitemap and reflects your actual AI crawler policy, and check that no deploy left a noindex tag on production.
Date: 25 November 2026 (single pass, strictly read-only GET requests + public DNS TXT queries).
Sample: 500 unique public domains drawn from European and North American web agencies, WordPress care studios, SMB commercial sites, and SaaS teams; 18 timed out or returned non-2xx/3xx errors, leaving 482 responsive websites ( 305 WordPress, 23 Next.js, 154 other stacks).
Checks performed: landing-page HTTPS status, five OWASP browser security headers, Server / X-Powered-By version banners, homepage noindex and canonical tags, /robots.txt AI crawler directives, /llms.txt plain-text/Markdown presence, /.well-known/security.txt ( Contact: line), /openapi.json JSON schema check, DNS SPF and _dmarc TXT records, and first-byte timing from a European vantage point.
Limitations: this cohort skews toward active agency, WordPress, and SMB sites ( 63.3% WordPress) rather than the Alexa/Tranco top 1,000, so enterprise header adoption is higher than this sample. Passive checks cannot see private repositories, authenticated dashboards, or server-side WAF rules.
Beamfar runs these external checks continuously after every deploy: multi-region uptime confirmation, TLS and domain expiry warnings, DNS and DMARC change alerts, SEO-risk alerts ( noindex , robots.txt , canonical drift), passive frontend secret leak detection across 12 credential families, exposed API schema checks, and continuous lockfile/SBOM CISA KEV Exploit Watch across eight ecosystems. Run a one-off check with zero storage using our free site scan , security headers checker , frontend secret leak checker , or lockfile CVE checker — or start on the Free plan (one site, no card, no time limit) and connect Claude, Cursor, Codex, or Gemini CLI via our 16-tool read-only MCP server .
Read on 25 November 2026 . Standards and vendor pages change; confirm the current text before you rely on it.
OWASP Secure Headers Project (owasp.org)
RFC 9116 security.txt specification (www.rfc-editor.org)
RFC 7489 DMARC specification (www.rfc-editor.org)
OpenAI crawler documentation (GPTBot and OAI-SearchBot) (platform.openai.com)
Anthropic crawler document

[truncated]
