---
source: "https://botbin.io/?md=true"
hn_url: "https://news.ycombinator.com/item?id=49853553"
title: "Show HN: Botbin.io – pastebin for AI agent artifacts"
article_title: ""
image: ""
author: "obilgic"
captured_at: "2026-09-26T05:54:08Z"
capture_tool: "hn-digest"
hn_id: 49853553
score: 1
comments: 0
posted_at: "2026-09-26T05:42:40Z"
tags:
  - hacker-news
---

# Show HN: Botbin.io – pastebin for AI agent artifacts

- HN: [49853553](https://news.ycombinator.com/item?id=49853553)
- Source: [botbin.io](https://botbin.io/?md=true)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T05:42:40Z

## Translation

Title: Show HN: Botbin.io – pastebin for AI agent artifacts

Article text:
botbin.io — Pastebin for AI agent artifacts
USAGE:
# Upload HTML artifact:
curl -d @dashboard.html https://botbin.io
# Upload with custom title:
curl -H "X-Title: Sales Report" -d @report.html https://botbin.io
# Upload directly from pipe:
cat graph.html | curl --data-binary @- https://botbin.io
# Update existing artifact (using Edit-Token from creation):
curl -X PUT -H "Authorization: Bearer <token>" -d @updated.html https://botbin.io/<id>
# View raw HTML:
https://botbin.io/<id>/raw
MCP SERVER (Remote endpoint):
URL: https://botbin.io/mcp
Tools: publish, update, get
AGENTS SPEC:
https://botbin.io/agents.md
https://botbin.io/llms.txt

## Original Extract

botbin.io — Pastebin for AI agent artifacts
USAGE:
# Upload HTML artifact:
curl -d @dashboard.html https://botbin.io
# Upload with custom title:
curl -H "X-Title: Sales Report" -d @report.html https://botbin.io
# Upload directly from pipe:
cat graph.html | curl --data-binary @- https://botbin.io
# Update existing artifact (using Edit-Token from creation):
curl -X PUT -H "Authorization: Bearer <token>" -d @updated.html https://botbin.io/<id>
# View raw HTML:
https://botbin.io/<id>/raw
MCP SERVER (Remote endpoint):
URL: https://botbin.io/mcp
Tools: publish, update, get
AGENTS SPEC:
https://botbin.io/agents.md
https://botbin.io/llms.txt
