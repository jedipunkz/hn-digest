---
source: "https://starwell.dev/"
hn_url: "https://news.ycombinator.com/item?id=49625644"
title: "Show HN: Starwell – give AI agents real statistics instead of made up numbers"
article_title: "Starwell · The world's data, one question away"
image: "https://starwell.dev/starwell-og.png"
author: "adarsh4052"
captured_at: "2026-09-09T12:46:23Z"
capture_tool: "hn-digest"
hn_id: 49625644
score: 1
comments: 1
posted_at: "2026-09-09T12:44:36Z"
tags:
  - hacker-news
---

# Show HN: Starwell – give AI agents real statistics instead of made up numbers

- HN: [49625644](https://news.ycombinator.com/item?id=49625644)
- Source: [starwell.dev](https://starwell.dev/)
- Score: 1
- Comments: 1
- Posted: 2026-09-09T12:44:36Z

## Translation

Title: Show HN: Starwell – give AI agents real statistics instead of made up numbers
Article title: Starwell · The world's data, one question away
Description: Verified official statistics for people and AI agents: one REST API, one MCP server. Every value served with provenance, a verification status, and a citation to the official source. Free with an account.

Article text:
Skip to content
STARWELL
Chat
Catalog
Docs
Pricing
About
Earthlight
Account
Get your free key
The verified data layer for AI
AI agents
hallucinate numbers.
Starwell fixes that.
The world's official statistics, harmonized into one REST API and one MCP server. Every value cites the official table it came from.
Aug 2024 Aug 2026
United States U.S. Bureau of Labor Statistics · 12-mo change of CPI-U, computed
3.3 % Jul 2026
Canada Statistics Canada · 12-mo change of all-items CPI, computed
3.0 % Jul 2026
Euro area European Central Bank · HICP annual rate, as published
3.3 % Aug 2026
United Kingdom UK Office for National Statistics · CPI annual rate, as published
2.9 % Jul 2026
● Counts as of September 2026. Browse the catalog .
29 official sources, and counting.
Each source keeps its own authority, license, attribution and terms.
World Bank Open Data (World Development Indicators)
Eurostat (EU statistics dissemination API)
ILOSTAT (International Labour Organization)
U.S. Bureau of Economic Analysis Data API (NIPA)
Australian Bureau of Statistics Data API
Explore all 29 sources in the catalog →
Every response states its license.
Every response looks like this.
The value and the record behind it, down to the citation.
curl "https://starwell.dev/api/starwell/v1/sources/fred/series/DGS10/observations?latest=1"
Copy
DGS10 = 4.78% (2026-09-04) passing Federal Reserve Economic Data (FRED)
{
"success": true,
"data": [
{
"period": "2026-09-04",
"value": 4.78,
"status": "normal",
"decimals": null,
"revision": 1,
"provenance": {
"sourceUrl": "https://api.stlouisfed.org/fred/series/observations?series_id=DGS10&observation_start=1962-01-02&observation_end=2026-09-04&file_type=json&api_key=REDACTED",
"retrievedAt": "2026-09-09T06:03:49.443+00:00",
"connectorVersion": "1.0.0",
"releaseTime": null
}
}
],
"meta": {
"series": {
"externalId": "DGS10",
"indicator": "Market Yield on U.S. Treasury Securities at 10-Year Constant Maturity, Quoted on an Investment Basis",
"unit": "Percent",
"frequency": "daily",
"geography": "United States"
},
"verification": {
"status": "passing",
"lastVerifiedAt": "2026-09-09T06:04:45.696+00:00"
},
"license": {
"code": "public-domain",
"url": "https://fred.stlouisfed.org/legal/",
"attribution": "Source: U.S. federal statistical agencies via FRED, Federal Reserve Bank of St. Louis. This product uses the FRED API but is not endorsed or certified by the Federal Reserve Bank of St. Louis.",
"redistributable": true,
"trainingEligible": true
},
"citation": {
"source": "Federal Reserve Economic Data (FRED)",
"authority": "Federal Reserve Bank of St. Louis",
"dataset": "H.15 Selected Interest Rates",
"datasetExternalId": "18",
"url": "https://fred.stlouisfed.org/release?rid=18",
"termsUrl": "https://fred.stlouisfed.org/legal/",
"text": "H.15 Selected Interest Rates. Federal Reserve Bank of St. Louis. https://fred.stlouisfed.org/release?rid=18"
},
"returned": 1,
"hasMore": true
}
}
A real response from the API.
Connectors pull each agency's official releases on schedule and normalize periods, units and revisions. A new source ships only after its tests pass.
115 flagship indicators are re-read live from the source agency on a schedule and compared, value for value.
REST and MCP, same store. Every value carries its provenance, its license and a citation.
Versioned /v1 : sources, datasets, series, observations, computed answers. OpenAPI 3.1 at /openapi.json .
11 tools over Streamable HTTP: search the catalog, fetch observations, compute answers, manage monitors.
Ask in plain English. The result cites the data points it used.
format=csv on any observations call. Citation and license ride along as header comments.
Every indicator as an embeddable SVG chart at chart.svg . No JavaScript.
series.updated fires at your URL when a new or revised value lands.
New data announced at /changelog , machine-readable as Atom at /changelog.xml .
Every page has a machine twin.
The store speaks MCP too. Point a client at the endpoint and these statistics become native tools ( the agent-native web ).
{
"mcpServers": {
"official-statistics": {
"type": "http",
"url": "https://starwell.dev/api/starwell/mcp"
}
}
}
/llms.txt
/llms-full.txt
/openapi.json
/changelog.xml
/v1/catalog
By the world's data we mean the world's public record: official statistics published by national agencies and public institutions.
Starwell
The world's data, one question away.

## Original Extract

Verified official statistics for people and AI agents: one REST API, one MCP server. Every value served with provenance, a verification status, and a citation to the official source. Free with an account.

Skip to content
STARWELL
Chat
Catalog
Docs
Pricing
About
Earthlight
Account
Get your free key
The verified data layer for AI
AI agents
hallucinate numbers.
Starwell fixes that.
The world's official statistics, harmonized into one REST API and one MCP server. Every value cites the official table it came from.
Aug 2024 Aug 2026
United States U.S. Bureau of Labor Statistics · 12-mo change of CPI-U, computed
3.3 % Jul 2026
Canada Statistics Canada · 12-mo change of all-items CPI, computed
3.0 % Jul 2026
Euro area European Central Bank · HICP annual rate, as published
3.3 % Aug 2026
United Kingdom UK Office for National Statistics · CPI annual rate, as published
2.9 % Jul 2026
● Counts as of September 2026. Browse the catalog .
29 official sources, and counting.
Each source keeps its own authority, license, attribution and terms.
World Bank Open Data (World Development Indicators)
Eurostat (EU statistics dissemination API)
ILOSTAT (International Labour Organization)
U.S. Bureau of Economic Analysis Data API (NIPA)
Australian Bureau of Statistics Data API
Explore all 29 sources in the catalog →
Every response states its license.
Every response looks like this.
The value and the record behind it, down to the citation.
curl "https://starwell.dev/api/starwell/v1/sources/fred/series/DGS10/observations?latest=1"
Copy
DGS10 = 4.78% (2026-09-04) passing Federal Reserve Economic Data (FRED)
{
"success": true,
"data": [
{
"period": "2026-09-04",
"value": 4.78,
"status": "normal",
"decimals": null,
"revision": 1,
"provenance": {
"sourceUrl": "https://api.stlouisfed.org/fred/series/observations?series_id=DGS10&observation_start=1962-01-02&observation_end=2026-09-04&file_type=json&api_key=REDACTED",
"retrievedAt": "2026-09-09T06:03:49.443+00:00",
"connectorVersion": "1.0.0",
"releaseTime": null
}
}
],
"meta": {
"series": {
"externalId": "DGS10",
"indicator": "Market Yield on U.S. Treasury Securities at 10-Year Constant Maturity, Quoted on an Investment Basis",
"unit": "Percent",
"frequency": "daily",
"geography": "United States"
},
"verification": {
"status": "passing",
"lastVerifiedAt": "2026-09-09T06:04:45.696+00:00"
},
"license": {
"code": "public-domain",
"url": "https://fred.stlouisfed.org/legal/",
"attribution": "Source: U.S. federal statistical agencies via FRED, Federal Reserve Bank of St. Louis. This product uses the FRED API but is not endorsed or certified by the Federal Reserve Bank of St. Louis.",
"redistributable": true,
"trainingEligible": true
},
"citation": {
"source": "Federal Reserve Economic Data (FRED)",
"authority": "Federal Reserve Bank of St. Louis",
"dataset": "H.15 Selected Interest Rates",
"datasetExternalId": "18",
"url": "https://fred.stlouisfed.org/release?rid=18",
"termsUrl": "https://fred.stlouisfed.org/legal/",
"text": "H.15 Selected Interest Rates. Federal Reserve Bank of St. Louis. https://fred.stlouisfed.org/release?rid=18"
},
"returned": 1,
"hasMore": true
}
}
A real response from the API.
Connectors pull each agency's official releases on schedule and normalize periods, units and revisions. A new source ships only after its tests pass.
115 flagship indicators are re-read live from the source agency on a schedule and compared, value for value.
REST and MCP, same store. Every value carries its provenance, its license and a citation.
Versioned /v1 : sources, datasets, series, observations, computed answers. OpenAPI 3.1 at /openapi.json .
11 tools over Streamable HTTP: search the catalog, fetch observations, compute answers, manage monitors.
Ask in plain English. The result cites the data points it used.
format=csv on any observations call. Citation and license ride along as header comments.
Every indicator as an embeddable SVG chart at chart.svg . No JavaScript.
series.updated fires at your URL when a new or revised value lands.
New data announced at /changelog , machine-readable as Atom at /changelog.xml .
Every page has a machine twin.
The store speaks MCP too. Point a client at the endpoint and these statistics become native tools ( the agent-native web ).
{
"mcpServers": {
"official-statistics": {
"type": "http",
"url": "https://starwell.dev/api/starwell/mcp"
}
}
}
/llms.txt
/llms-full.txt
/openapi.json
/changelog.xml
/v1/catalog
By the world's data we mean the world's public record: official statistics published by national agencies and public institutions.
Starwell
The world's data, one question away.
