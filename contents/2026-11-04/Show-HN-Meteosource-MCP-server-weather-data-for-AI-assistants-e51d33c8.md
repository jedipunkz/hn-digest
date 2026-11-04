---
source: "https://mcp.meteosource.com/"
hn_url: "https://news.ycombinator.com/item?id=50362615"
title: "Show HN: Meteosource MCP server, weather data for AI assistants"
article_title: "Weather MCP Server by Meteosource: weather for AI agents"
image: "https://www.meteosource.com/static/img/mcp-og.png"
author: "Thunder_rumbles"
captured_at: "2026-11-04T08:43:30Z"
capture_tool: "hn-digest"
hn_id: 50362615
score: 2
comments: 1
posted_at: "2026-11-04T08:02:25Z"
tags:
  - hacker-news
---

# Show HN: Meteosource MCP server, weather data for AI assistants

- HN: [50362615](https://news.ycombinator.com/item?id=50362615)
- Source: [mcp.meteosource.com](https://mcp.meteosource.com/)
- Score: 2
- Comments: 1
- Posted: 2026-11-04T08:02:25Z

## Translation

Title: Show HN: Meteosource MCP server, weather data for AI assistants
Article title: Weather MCP Server by Meteosource: weather for AI agents
Description: Hosted MCP server for AI agents: weather forecasts, severe-weather alerts, air quality, 20 years of history and solar and wind production from Meteosource.

Article text:
-->
Get a free API key
Meteosource
Connecting
Tools
Documentation
Weather data for AI agents, ready to connect
Forecasts, alerts, air quality, 20 years of history and renewable production over the Model Context Protocol
A hosted MCP server by Meteosource. Point Claude, Codex, Antigravity or any other MCP client at
one address, add your API key, and the agent answers weather questions with real data.
https://mcp.meteosource.com/mcp
Authorization: Bearer YOUR_METEOSOURCE_API_KEY
Claude, Codex, Antigravity and any other MCP client
Two values, and the agent is online
The server speaks the Model Context Protocol over Streamable HTTP. Every client needs the same two
things: the address and your Meteosource API key in a header. The plan the key is on decides which
tools the agent sees.
Get a Meteosource API key Sign up at meteosource.com ; the free plan is enough to try the forecast tools.
Add the server in your client Address https://mcp.meteosource.com/mcp , transport Streamable HTTP.
Add the header Authorization: Bearer YOUR_METEOSOURCE_API_KEY . Clients with a plain API-key field can send X-API-Key instead.
Ask about the weather The agent picks the right tools itself, from place lookup to forecast.
The key travels in a request header, is used only to call Meteosource on your behalf and is never stored in full. Caches are kept per key.
Step-by-step setup for Claude Code, Claude Desktop and claude.ai, Codex, Antigravity and other
clients, including team setups that keep the key out of version control, is in the
Connect section of the README .
The weather tools are for the Free, Startup and Standard plans; the production tools need the
Renewables plan. See plans and pricing .
Place name, postal code or coordinates to a place ID for the other tools.
Current conditions, hourly and daily forecast, minute-by-minute precipitation, up to 20 locations per call.
Official severe-weather warnings for the USA, EU and Canada.
Past weather, hourly, daily or monthly, with long-term normals and records.
Hourly air quality about four days ahead.
Output of a PV system, 20 years back to four days ahead.
Output of a wind farm, 20 years back to four days ahead.
Parameters, limits and the shape of what the agent receives, in the README.
The README is the reference for the server. Jump straight to the part you need.
Connect setup for each client, including team setups that keep the key out of version control
Tools what each tool accepts and returns
Try it example questions to ask the agent
Plans and limits quotas, per-call limits and the cache
Your key and your data how the key is handled
Troubleshooting what the agent reports and what to do
Questions about your key or plan go to Meteosource support. For the server itself,
open an issue on GitHub.

## Original Extract

Hosted MCP server for AI agents: weather forecasts, severe-weather alerts, air quality, 20 years of history and solar and wind production from Meteosource.

-->
Get a free API key
Meteosource
Connecting
Tools
Documentation
Weather data for AI agents, ready to connect
Forecasts, alerts, air quality, 20 years of history and renewable production over the Model Context Protocol
A hosted MCP server by Meteosource. Point Claude, Codex, Antigravity or any other MCP client at
one address, add your API key, and the agent answers weather questions with real data.
https://mcp.meteosource.com/mcp
Authorization: Bearer YOUR_METEOSOURCE_API_KEY
Claude, Codex, Antigravity and any other MCP client
Two values, and the agent is online
The server speaks the Model Context Protocol over Streamable HTTP. Every client needs the same two
things: the address and your Meteosource API key in a header. The plan the key is on decides which
tools the agent sees.
Get a Meteosource API key Sign up at meteosource.com ; the free plan is enough to try the forecast tools.
Add the server in your client Address https://mcp.meteosource.com/mcp , transport Streamable HTTP.
Add the header Authorization: Bearer YOUR_METEOSOURCE_API_KEY . Clients with a plain API-key field can send X-API-Key instead.
Ask about the weather The agent picks the right tools itself, from place lookup to forecast.
The key travels in a request header, is used only to call Meteosource on your behalf and is never stored in full. Caches are kept per key.
Step-by-step setup for Claude Code, Claude Desktop and claude.ai, Codex, Antigravity and other
clients, including team setups that keep the key out of version control, is in the
Connect section of the README .
The weather tools are for the Free, Startup and Standard plans; the production tools need the
Renewables plan. See plans and pricing .
Place name, postal code or coordinates to a place ID for the other tools.
Current conditions, hourly and daily forecast, minute-by-minute precipitation, up to 20 locations per call.
Official severe-weather warnings for the USA, EU and Canada.
Past weather, hourly, daily or monthly, with long-term normals and records.
Hourly air quality about four days ahead.
Output of a PV system, 20 years back to four days ahead.
Output of a wind farm, 20 years back to four days ahead.
Parameters, limits and the shape of what the agent receives, in the README.
The README is the reference for the server. Jump straight to the part you need.
Connect setup for each client, including team setups that keep the key out of version control
Tools what each tool accepts and returns
Try it example questions to ask the agent
Plans and limits quotas, per-call limits and the cache
Your key and your data how the key is handled
Troubleshooting what the agent reports and what to do
Questions about your key or plan go to Meteosource support. For the server itself,
open an issue on GitHub.
