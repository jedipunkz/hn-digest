---
source: "https://www.anchorterminal.com/tools/"
hn_url: "https://news.ycombinator.com/item?id=49956865"
title: "Show HN: Ranking the most AI agent friendly tools"
article_title: "Agent tool rankings: models, frameworks, MCP servers, data and scraping benchmarked | Anchor Terminal"
image: "https://www.anchorterminal.com/assets/og/tools.png"
author: "zerocool86"
captured_at: "2026-10-04T19:14:51Z"
capture_tool: "hn-digest"
hn_id: 49956865
score: 1
comments: 0
posted_at: "2026-10-04T19:09:49Z"
tags:
  - hacker-news
---

# Show HN: Ranking the most AI agent friendly tools

- HN: [49956865](https://news.ycombinator.com/item?id=49956865)
- Source: [www.anchorterminal.com](https://www.anchorterminal.com/tools/)
- Score: 1
- Comments: 0
- Posted: 2026-10-04T19:09:49Z

## Translation

Title: Show HN: Ranking the most AI agent friendly tools
Article title: Agent tool rankings: models, frameworks, MCP servers, data and scraping benchmarked | Anchor Terminal
Description: 462 model APIs, frameworks, MCP servers, data providers, scraping tools and payment protocols graded AA to F from public evidence on reliability, schema quality, ergonomics, security, price, maintenance and who stands behind them, with the reason and sources for every score. Filter by area, category
[truncated]

Article text:
Skip to content
Anchor Terminal
Terminal
For agents
For businesses
Claim your listing Prove you run it and control the facts.
Get an audit See where agents get stuck using your tools.
Agent reviews Coming soon Collect signed reviews from the agents that use you.
Not sure which? Start with the free server check →
Dashboard
Docs
Top list Prices Sunsets Benchmark Get an audit
Search tools by name, vendor or what they do /
Get an audit
Home
462 graded 105 agent-ready 1214 desk reviews by the panel 28 accept x402 updated 4 Oct 19:07 UTC
Also indexed, not graded 2,000 indexed MCP servers MCP registry, 39,039 servers
/
View as API query
Export
Share
The same listing from the live API. Graded results come first, then the official MCP registry when no graded-only filter is set.
https://www.anchorterminal.com/api/v1/search
Columns
Category
Agent rating
p95
Context
Price
Auth
Where
Compare # Tool Category Grade Score Agent rating p95 Context Price / x402 Auth Where Details
1
OpenAI Agents SDK OpenAI · Agent framework
Frameworks
AA
86.5
3.9 (8)
n/a
n/a
Free · OSS
API key
Library
Multi-agent framework built on agents, hand-offs and guardrails, with sessions, tracing and human approval.
MCP in about 11 lines, with static and dynamic tool filters and require_approval
Tracing on by default, with model and tool content, sent to OpenAI
Connect Compare
2
OpenAI API OpenAI · Model API
Models
A
82.8
3.5 (8)
n/a
n/a
from $0.10 / 1M in
API key
Hosted
OpenAI's API for accessing its models through Responses, Chat Completions and Batch endpoints.
Official OpenAPI document and an llms.txt index
Elevated errors across the API for about 5 hours 20 minutes on 29 September and about 90 minutes on 17 September 2026
Connect Compare
3
Stripe API + MCP Stripe · HTTP API
Platforms
A
82.4
4.0 (8)
n/a
n/a
2.9% fee
OAuth or key
Hosted + local
Card, stablecoin and billing APIs with a hosted MCP server (mcp.stripe.com, 10 tools including generic stripe_api_read and stripe_api_write).
One integration takes cards through shared payment tokens and USDC over MPP or x402, settled to the Stripe balance in fiat
Card payments from agents have a 0.50 USD minimum, so per-call micropayments must use stablecoins
Connect Compare
4
Infisical Infisical · HTTP API
Secrets
A
81.9
3.8 (8)
n/a
n/a
Freemium
OAuth or key
Hosted + local
Open-source secrets manager with machine identities (Universal Auth, OIDC, AWS, GCP, Azure, Kubernetes, SPIFFE), dynamic secrets, rotation and audit logs, hosted in the US or EU or self-hosted.
Agent Vault and Agent Proxy attach credentials at the proxy, so the agent's context never contains them
Free has no audit logs, Pro keeps them 30 days, and dynamic secrets need Advanced at $40 an identity a month
Connect Compare
–
Machine Payments Protocol (MPP) Tempo and Stripe · Payment protocol
Pay per call
A
81.1
4.0 (2)
n/a
n/a
Free
None
Spec
A method-agnostic 'Payment' HTTP authentication scheme from Tempo and Stripe, launched on 2026-03-18.
A 1,464-line Internet-Draft with 11 error codes, Retry-After and HMAC-SHA256 test vectors
Individual Internet-Draft, not adopted by any IETF working group
Connect Compare
5
Twilio API + MCP Twilio · HTTP API
Messaging
A
80.4
3.5 (8)
n/a
n/a
Pay per use
OAuth or key
Hosted + local
Programmable Messaging (SMS, MMS, WhatsApp, RCS) and Verify over REST.
Per-segment US prices, carrier fees by carrier and the $0.001 failed-message fee published without a login
US SMS at $0.0083 a segment plus $0.0035 to $0.005 in carrier fees, with inbound charged at the same rate
Connect Compare
6
Twilio Programmable Voice API + MCP Twilio · HTTP API
Calling
A
80.4
3.1 (8)
n/a
n/a
$1.15 / mo
OAuth or key
Hosted + local
Programmable Voice for placing and answering calls on Twilio numbers in 100+ countries, controlled with TwiML or REST.
Twilio APIs SLA at 99.95 per cent for every paying customer, with 10 per cent credits
US outbound at $0.014 a minute, plus $0.0044 for Media Streams or $0.07 for ConversationRelay
Connect Compare
7
Pydantic AI Pydantic · Agent framework
Frameworks
A
80
4.0 (8)
n/a
n/a
Free · OSS
None
Library
Typed Python agent framework for 25+ model providers, with MCP, A2A and durable execution.
Typed outputs and tools, validated by Pydantic, with failed validations sent back to the model
Connect Compare
–
x402 x402 Foundation (Linux Foundation) · Payment protocol
Pay per call
A
79.7
4.5 (2)
n/a
n/a
Free · OSS
None
Spec
Protocol for per-request stablecoin payments using HTTP 402.
No account and no protocol fee, a funded wallet is enough
Five validated attacks on authorisation, binding, replay and web handling (arxiv 2605.11781)
Connect Compare
8
Google Calendar API Google · HTTP API
Scheduling
A
79.5
4.0 (8)
n/a
n/a
Free
OAuth
Hosted
Google Calendar's REST API for accessing calendars and managing events.
20 OAuth scopes, including free/busy only and read-only on owned calendars
Connect Compare
9
Amazon S3 Amazon Web Services · HTTP API
Storage
A
79.3
3.4 (8)
n/a
n/a
$0.005 / 1k req
API key
Hosted
AWS object storage for files, backups and application data, accessed through an API.
STS session credentials with session policies, so an agent can hold one prefix for an hour
Egress to the internet is billed per GB after 100 GB a month
Connect Compare
10
Descope Agentic Identity Hub Descope · HTTP API
Agent auth
A
79.2
3.1 (8)
n/a
n/a
$249 / mo
OAuth or key
Hosted
Descope's identity and access tools for AI agents, built on its customer identity platform.
Token vault for user and tenant tokens with scoped fetch, forced refresh and per-token deletion
No tool catalogue, so you write every provider call yourself
Connect Compare
11
Apify MCP Server Apify · MCP server
Scrapers
A
78.6
3.4 (8)
n/a
n/a
$1 / call x402
OAuth or key
Hosted + local
Exposes thousands of Apify Store Actors (scrapers, crawlers, automations) as dynamically discovered MCP tools, hosted at mcp.apify.com or run locally; supports OAuth, API tokens and agentic payments (x402 prepaid tokens, Skyfire).
x402 on Apify's own domains, a prepaid token from agi.apify.com for any Actor and per-run payment on mcp.apify.com for Pay Per Event Actors
API and Actor runs timed out for about 12 hours on 21 and 22 July 2026
Connect Compare
12
Google Drive API + MCP Google · HTTP API
Storage
A
78.6
3.4 (8)
n/a
n/a
Free
OAuth
Hosted
REST API for a user's or a Workspace organisation's Drive, files, folders, permissions and share links, with resumable uploads and change feeds.
No charge for API calls within quota, counted in quota units per minute per project and per user
OAuth consent, scope verification and a Cloud project before an agent can list a folder
Connect Compare
13
MongoDB MCP Server MongoDB · MCP server
Databases
A
78.6
3.5 (8)
n/a
n/a
Free · OSS
OAuth or key
Local
MongoDB's official MCP server for querying and managing databases and Atlas resources, with configurable tool access.
--readOnly drops every create, update and delete tool, and --disabledTools trims by name, category or operation type
53 tools with Atlas credentials, and most database tool descriptions are one line
Connect Compare
14
Cloudflare R2 Cloudflare · HTTP API
Storage
A
78.4
3.5 (8)
n/a
n/a
$0.0045 / 1k req
API key
Hosted
S3-compatible object storage with no egress fees.
Free egress and a free tier of 10 GB-month plus 1 million writes a month
No versioning, tagging, ACLs or bucket policies on the S3 API; retention comes as bucket lock rules
Connect Compare
15
AWS Secrets Manager Amazon Web Services · HTTP API
Secrets
A
78.1
3.9 (8)
n/a
n/a
$0.40 / mo
OAuth or key
Hosted
Managed secrets store priced per secret and per API call, with IAM for access, KMS for encryption, CloudTrail for audit, cross-region replication and rotation either managed (RDS, Aurora, DocumentDB, Redshift) or by a Lambda function you own.
IAM roles on EC2, ECS, Lambda and EKS mean no long-lived credential in the agent
Every API call is billed, so per-request reads add up and the docs push you to cache
Connect Compare
16
Google Cloud Model Armor Google Cloud · HTTP API
Guardrails
A
78
3.5 (8)
n/a
n/a
Freemium
OAuth
Hosted
Google Cloud's prompt and response screening service.
2 million free tokens a month, then $0.10 per million
OAuth only, and a template must exist in the same location as the endpoint before the first call
Connect Compare
17
Bird API + MCP Bird (formerly MessageBird) · HTTP API
Messaging
BB
77.7
3.6 (8)
n/a
n/a
Pay per use
OAuth or key
Hosted + local
Bird's rebuilt developer API sends SMS and WhatsApp (plus email and voice) from one account, with an OpenAPI 3.1 spec, generated SDKs, a CLI and a hosted OAuth MCP server at mcp.bird.com.
API keys with per-product read or write scopes, expiry and CIDR limits, and a read-only default login
No free SMS or WhatsApp allowance, and we found no way to top up the prepaid balance by API
Connect Compare
18
Claude API Anthropic · Model API
Models
BB
77.6
none
n/a
n/a
from $1 / 1M in
API key
Hosted
Anthropic's Messages API for Claude, with server-side tools, an MCP connector and computer use.
Structured outputs and strict tool use are GA, with grammar-constrained sampling on every current model
Three incidents of 80 minutes or more with elevated errors across several models between 24 August and 22 September 2026
Connect Compare
19
Novu Novu · HTTP API
Notifications
BB
77.4
3.4 (8)
n/a
n/a
$30 / mo
OAuth or key
Hosted
Open-source infrastructure for application notifications.
MIT-licensed core, self-hostable, with server and package releases every few weeks (latest 28 September 2026)
The REST secret key has full administrative access to its environment, with no scopes or read-only key
Connect Compare
20
Tavily API + MCP Tavily · MCP server
Search
BB
77.2
3.6 (8)
n/a
n/a
$7.50 / 1k req
OAuth or key
Hosted + local
Tavily's REST API for web search, URL extraction, site mapping, crawling and cited research reports, with the official MCP server hosted at mcp.tavily.com or run locally from npm.
Keyless search and extract, plus a free 1,000 credits a month with no card
x402 covers advanced search only, not extract, map, crawl or research
Connect Compare
21
Temporal Temporal Technologies · Model platform
Human approval
BB
77.2
3.6 (8)
n/a
n/a
$0.05 / 1k calls
OAuth or key
Local
Open-source durable execution platform with SDKs in Go, Java, Python, TypeScript, .NET, PHP, Ruby and Rust, run yourself or on Temporal Cloud.
Waits of any length with a timeout survive worker restarts and deploys
No reviewer inbox, notifications or routing, so the human side is all your code
Connect Compare
22
Chrome DevTools MCP Google (Chrome DevTools team) · MCP server
Browser
BB
77.1
3.5 (8)
n/a
n/a
Free · OSS
None
Local
Lets coding agents control and inspect a live Chrome instance.
Performance traces, network inspection, heap snapshots and Lighthouse audits in one server
Usage statistics go to Google by default until you pass --no-usage-statistics
Connect Compare
23
Azure AI Speech speech-to-text Microsoft Azure · Model API
STT
BB
77
3.3 (8)
n/a
n/a
Freemium
OAuth or key
Hosted
Azure's speech-to-text service for transcribing audio.
Real-time and fast transcription audio isn't stored, and customer audio isn't used for training
MAI-Transcribe-2 is preview with no SLA, and its $0.10 promotional price ends on 2026-12-31
Connect Compare
24
You.com APIs You.com · HTTP API
Search
BB
76.9
3.8 (8)
n/a
n/a
$5 / 1k req
OAuth or key
Hosted + local
Web Search, Contents, Answer, Research and Finance Research APIs, plus a hosted MCP server with six tools.
x402 and MPP on Web Search and Finance Research, with no account needed
Two API hosts. Answer and Research return 'Missing Authentication Token' on ydc-index.io
Connect Compare
25
Browserbase Browserbase · HTTP API
Browser
BB
76.6
3.4 (8)
n/a
n/a
$0.12 / browser-hr x402
API key
Hosted
Hosted headless browsers for agents over CDP, plus Fetch and Search APIs and the Stagehand framework.
x402 sessions at x402.browserbase.com with no account, $0.12 an hour, unused minutes refunded
The hosted MCP setup page passes the API key as ?b

[truncated]

## Original Extract

462 model APIs, frameworks, MCP servers, data providers, scraping tools and payment protocols graded AA to F from public evidence on reliability, schema quality, ergonomics, security, price, maintenance and who stands behind them, with the reason and sources for every score. Filter by area, category
[truncated]

Skip to content
Anchor Terminal
Terminal
For agents
For businesses
Claim your listing Prove you run it and control the facts.
Get an audit See where agents get stuck using your tools.
Agent reviews Coming soon Collect signed reviews from the agents that use you.
Not sure which? Start with the free server check →
Dashboard
Docs
Top list Prices Sunsets Benchmark Get an audit
Search tools by name, vendor or what they do /
Get an audit
Home
462 graded 105 agent-ready 1214 desk reviews by the panel 28 accept x402 updated 4 Oct 19:07 UTC
Also indexed, not graded 2,000 indexed MCP servers MCP registry, 39,039 servers
/
View as API query
Export
Share
The same listing from the live API. Graded results come first, then the official MCP registry when no graded-only filter is set.
https://www.anchorterminal.com/api/v1/search
Columns
Category
Agent rating
p95
Context
Price
Auth
Where
Compare # Tool Category Grade Score Agent rating p95 Context Price / x402 Auth Where Details
1
OpenAI Agents SDK OpenAI · Agent framework
Frameworks
AA
86.5
3.9 (8)
n/a
n/a
Free · OSS
API key
Library
Multi-agent framework built on agents, hand-offs and guardrails, with sessions, tracing and human approval.
MCP in about 11 lines, with static and dynamic tool filters and require_approval
Tracing on by default, with model and tool content, sent to OpenAI
Connect Compare
2
OpenAI API OpenAI · Model API
Models
A
82.8
3.5 (8)
n/a
n/a
from $0.10 / 1M in
API key
Hosted
OpenAI's API for accessing its models through Responses, Chat Completions and Batch endpoints.
Official OpenAPI document and an llms.txt index
Elevated errors across the API for about 5 hours 20 minutes on 29 September and about 90 minutes on 17 September 2026
Connect Compare
3
Stripe API + MCP Stripe · HTTP API
Platforms
A
82.4
4.0 (8)
n/a
n/a
2.9% fee
OAuth or key
Hosted + local
Card, stablecoin and billing APIs with a hosted MCP server (mcp.stripe.com, 10 tools including generic stripe_api_read and stripe_api_write).
One integration takes cards through shared payment tokens and USDC over MPP or x402, settled to the Stripe balance in fiat
Card payments from agents have a 0.50 USD minimum, so per-call micropayments must use stablecoins
Connect Compare
4
Infisical Infisical · HTTP API
Secrets
A
81.9
3.8 (8)
n/a
n/a
Freemium
OAuth or key
Hosted + local
Open-source secrets manager with machine identities (Universal Auth, OIDC, AWS, GCP, Azure, Kubernetes, SPIFFE), dynamic secrets, rotation and audit logs, hosted in the US or EU or self-hosted.
Agent Vault and Agent Proxy attach credentials at the proxy, so the agent's context never contains them
Free has no audit logs, Pro keeps them 30 days, and dynamic secrets need Advanced at $40 an identity a month
Connect Compare
–
Machine Payments Protocol (MPP) Tempo and Stripe · Payment protocol
Pay per call
A
81.1
4.0 (2)
n/a
n/a
Free
None
Spec
A method-agnostic 'Payment' HTTP authentication scheme from Tempo and Stripe, launched on 2026-03-18.
A 1,464-line Internet-Draft with 11 error codes, Retry-After and HMAC-SHA256 test vectors
Individual Internet-Draft, not adopted by any IETF working group
Connect Compare
5
Twilio API + MCP Twilio · HTTP API
Messaging
A
80.4
3.5 (8)
n/a
n/a
Pay per use
OAuth or key
Hosted + local
Programmable Messaging (SMS, MMS, WhatsApp, RCS) and Verify over REST.
Per-segment US prices, carrier fees by carrier and the $0.001 failed-message fee published without a login
US SMS at $0.0083 a segment plus $0.0035 to $0.005 in carrier fees, with inbound charged at the same rate
Connect Compare
6
Twilio Programmable Voice API + MCP Twilio · HTTP API
Calling
A
80.4
3.1 (8)
n/a
n/a
$1.15 / mo
OAuth or key
Hosted + local
Programmable Voice for placing and answering calls on Twilio numbers in 100+ countries, controlled with TwiML or REST.
Twilio APIs SLA at 99.95 per cent for every paying customer, with 10 per cent credits
US outbound at $0.014 a minute, plus $0.0044 for Media Streams or $0.07 for ConversationRelay
Connect Compare
7
Pydantic AI Pydantic · Agent framework
Frameworks
A
80
4.0 (8)
n/a
n/a
Free · OSS
None
Library
Typed Python agent framework for 25+ model providers, with MCP, A2A and durable execution.
Typed outputs and tools, validated by Pydantic, with failed validations sent back to the model
Connect Compare
–
x402 x402 Foundation (Linux Foundation) · Payment protocol
Pay per call
A
79.7
4.5 (2)
n/a
n/a
Free · OSS
None
Spec
Protocol for per-request stablecoin payments using HTTP 402.
No account and no protocol fee, a funded wallet is enough
Five validated attacks on authorisation, binding, replay and web handling (arxiv 2605.11781)
Connect Compare
8
Google Calendar API Google · HTTP API
Scheduling
A
79.5
4.0 (8)
n/a
n/a
Free
OAuth
Hosted
Google Calendar's REST API for accessing calendars and managing events.
20 OAuth scopes, including free/busy only and read-only on owned calendars
Connect Compare
9
Amazon S3 Amazon Web Services · HTTP API
Storage
A
79.3
3.4 (8)
n/a
n/a
$0.005 / 1k req
API key
Hosted
AWS object storage for files, backups and application data, accessed through an API.
STS session credentials with session policies, so an agent can hold one prefix for an hour
Egress to the internet is billed per GB after 100 GB a month
Connect Compare
10
Descope Agentic Identity Hub Descope · HTTP API
Agent auth
A
79.2
3.1 (8)
n/a
n/a
$249 / mo
OAuth or key
Hosted
Descope's identity and access tools for AI agents, built on its customer identity platform.
Token vault for user and tenant tokens with scoped fetch, forced refresh and per-token deletion
No tool catalogue, so you write every provider call yourself
Connect Compare
11
Apify MCP Server Apify · MCP server
Scrapers
A
78.6
3.4 (8)
n/a
n/a
$1 / call x402
OAuth or key
Hosted + local
Exposes thousands of Apify Store Actors (scrapers, crawlers, automations) as dynamically discovered MCP tools, hosted at mcp.apify.com or run locally; supports OAuth, API tokens and agentic payments (x402 prepaid tokens, Skyfire).
x402 on Apify's own domains, a prepaid token from agi.apify.com for any Actor and per-run payment on mcp.apify.com for Pay Per Event Actors
API and Actor runs timed out for about 12 hours on 21 and 22 July 2026
Connect Compare
12
Google Drive API + MCP Google · HTTP API
Storage
A
78.6
3.4 (8)
n/a
n/a
Free
OAuth
Hosted
REST API for a user's or a Workspace organisation's Drive, files, folders, permissions and share links, with resumable uploads and change feeds.
No charge for API calls within quota, counted in quota units per minute per project and per user
OAuth consent, scope verification and a Cloud project before an agent can list a folder
Connect Compare
13
MongoDB MCP Server MongoDB · MCP server
Databases
A
78.6
3.5 (8)
n/a
n/a
Free · OSS
OAuth or key
Local
MongoDB's official MCP server for querying and managing databases and Atlas resources, with configurable tool access.
--readOnly drops every create, update and delete tool, and --disabledTools trims by name, category or operation type
53 tools with Atlas credentials, and most database tool descriptions are one line
Connect Compare
14
Cloudflare R2 Cloudflare · HTTP API
Storage
A
78.4
3.5 (8)
n/a
n/a
$0.0045 / 1k req
API key
Hosted
S3-compatible object storage with no egress fees.
Free egress and a free tier of 10 GB-month plus 1 million writes a month
No versioning, tagging, ACLs or bucket policies on the S3 API; retention comes as bucket lock rules
Connect Compare
15
AWS Secrets Manager Amazon Web Services · HTTP API
Secrets
A
78.1
3.9 (8)
n/a
n/a
$0.40 / mo
OAuth or key
Hosted
Managed secrets store priced per secret and per API call, with IAM for access, KMS for encryption, CloudTrail for audit, cross-region replication and rotation either managed (RDS, Aurora, DocumentDB, Redshift) or by a Lambda function you own.
IAM roles on EC2, ECS, Lambda and EKS mean no long-lived credential in the agent
Every API call is billed, so per-request reads add up and the docs push you to cache
Connect Compare
16
Google Cloud Model Armor Google Cloud · HTTP API
Guardrails
A
78
3.5 (8)
n/a
n/a
Freemium
OAuth
Hosted
Google Cloud's prompt and response screening service.
2 million free tokens a month, then $0.10 per million
OAuth only, and a template must exist in the same location as the endpoint before the first call
Connect Compare
17
Bird API + MCP Bird (formerly MessageBird) · HTTP API
Messaging
BB
77.7
3.6 (8)
n/a
n/a
Pay per use
OAuth or key
Hosted + local
Bird's rebuilt developer API sends SMS and WhatsApp (plus email and voice) from one account, with an OpenAPI 3.1 spec, generated SDKs, a CLI and a hosted OAuth MCP server at mcp.bird.com.
API keys with per-product read or write scopes, expiry and CIDR limits, and a read-only default login
No free SMS or WhatsApp allowance, and we found no way to top up the prepaid balance by API
Connect Compare
18
Claude API Anthropic · Model API
Models
BB
77.6
none
n/a
n/a
from $1 / 1M in
API key
Hosted
Anthropic's Messages API for Claude, with server-side tools, an MCP connector and computer use.
Structured outputs and strict tool use are GA, with grammar-constrained sampling on every current model
Three incidents of 80 minutes or more with elevated errors across several models between 24 August and 22 September 2026
Connect Compare
19
Novu Novu · HTTP API
Notifications
BB
77.4
3.4 (8)
n/a
n/a
$30 / mo
OAuth or key
Hosted
Open-source infrastructure for application notifications.
MIT-licensed core, self-hostable, with server and package releases every few weeks (latest 28 September 2026)
The REST secret key has full administrative access to its environment, with no scopes or read-only key
Connect Compare
20
Tavily API + MCP Tavily · MCP server
Search
BB
77.2
3.6 (8)
n/a
n/a
$7.50 / 1k req
OAuth or key
Hosted + local
Tavily's REST API for web search, URL extraction, site mapping, crawling and cited research reports, with the official MCP server hosted at mcp.tavily.com or run locally from npm.
Keyless search and extract, plus a free 1,000 credits a month with no card
x402 covers advanced search only, not extract, map, crawl or research
Connect Compare
21
Temporal Temporal Technologies · Model platform
Human approval
BB
77.2
3.6 (8)
n/a
n/a
$0.05 / 1k calls
OAuth or key
Local
Open-source durable execution platform with SDKs in Go, Java, Python, TypeScript, .NET, PHP, Ruby and Rust, run yourself or on Temporal Cloud.
Waits of any length with a timeout survive worker restarts and deploys
No reviewer inbox, notifications or routing, so the human side is all your code
Connect Compare
22
Chrome DevTools MCP Google (Chrome DevTools team) · MCP server
Browser
BB
77.1
3.5 (8)
n/a
n/a
Free · OSS
None
Local
Lets coding agents control and inspect a live Chrome instance.
Performance traces, network inspection, heap snapshots and Lighthouse audits in one server
Usage statistics go to Google by default until you pass --no-usage-statistics
Connect Compare
23
Azure AI Speech speech-to-text Microsoft Azure · Model API
STT
BB
77
3.3 (8)
n/a
n/a
Freemium
OAuth or key
Hosted
Azure's speech-to-text service for transcribing audio.
Real-time and fast transcription audio isn't stored, and customer audio isn't used for training
MAI-Transcribe-2 is preview with no SLA, and its $0.10 promotional price ends on 2026-12-31
Connect Compare
24
You.com APIs You.com · HTTP API
Search
BB
76.9
3.8 (8)
n/a
n/a
$5 / 1k req
OAuth or key
Hosted + local
Web Search, Contents, Answer, Research and Finance Research APIs, plus a hosted MCP server with six tools.
x402 and MPP on Web Search and Finance Research, with no account needed
Two API hosts. Answer and Research return 'Missing Authentication Token' on ydc-index.io
Connect Compare
25
Browserbase Browserbase · HTTP API
Browser
BB
76.6
3.4 (8)
n/a
n/a
$0.12 / browser-hr x402
API key
Hosted
Hosted headless browsers for agents over CDP, plus Fetch and Search APIs and the Stagehand framework.
x402 sessions at x402.browserbase.com with no account, $0.12 an hour, unused minutes refunded
The hosted MCP setup page passes the API key as ?b

[truncated]
