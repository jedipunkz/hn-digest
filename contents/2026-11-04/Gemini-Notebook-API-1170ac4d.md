---
source: "https://developers.google.com/notebook/docs"
hn_url: "https://news.ycombinator.com/item?id=50360411"
title: "Gemini Notebook API"
article_title: "Gemini Notebook API | Google for Developers"
image: "https://www.gstatic.com/devrel-devsite/prod/veed85ba0e3651062c793a40796e84fb34afbf8489308d45f92d5448e587e3f03/developers/images/opengraph/google-blue.png"
author: "jonbaer"
captured_at: "2026-11-04T01:56:05Z"
capture_tool: "hn-digest"
hn_id: 50360411
score: 1
comments: 0
posted_at: "2026-11-04T01:14:49Z"
tags:
  - hacker-news
---

# Gemini Notebook API

- HN: [50360411](https://news.ycombinator.com/item?id=50360411)
- Source: [developers.google.com](https://developers.google.com/notebook/docs)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T01:14:49Z

## Translation

Title: Gemini Notebook API
Article title: Gemini Notebook API | Google for Developers
Description: Programmatically create and manage notebooks and sources with the Gemini Notebook API.

Article text:
Gemini Notebook API | Google for Developers
Skip to main content
Gemini Notebook API
Gemini Notebook API
Stay organized with collections
Save and categorize content based on your preferences.
The Gemini Notebook API is a REST interface for programmatically creating and
managing notebooks and sources in Gemini Notebook. Ingest web pages, YouTube
videos, plain text, and Markdown into persistent notebooks.
Early Access Public Preview: The Gemini Notebook API and
MCP server are in an early access Public
Preview, and we're actively working on our
roadmap . We'd love to hear what you want
the API and MCP server to do. Please
file
feature requests in our Issue Tracker . We're excited to build this out,
and there's more to come!
During the preview, the API is available for personal Google Accounts.
Google Workspace accounts are not yet supported.
Just want to get started? Go to the
Quickstart , copy the setup prompt,
and give it to your coding agent.
import requests
API = "https://notebook.googleapis.com/v1"
headers = {
"Authorization" : f "Bearer { ACCESS_TOKEN } " ,
"x-goog-user-project" : PROJECT_ID ,
"Content-Type" : "application/json" ,
}
# 1. Create a notebook
nb = requests . post ( f " { API } /notebooks" , headers = headers , json = {
"display_name" : "My Research"
}) . json ()
# 2. Add a source
requests . post ( f " { API } / { nb [ 'name' ] } /sources" , headers = headers , json = {
"display_name" : "Gemini API Guide" ,
"web_content" : { "url" : "https://ai.google.dev/gemini-api/docs" }
})
JavaScript
const API = "https://notebook.googleapis.com/v1" ;
const headers = {
"Authorization" : `Bearer ${ ACCESS_TOKEN } ` ,
"x-goog-user-project" : PROJECT_ID ,
"Content-Type" : "application/json"
};
// 1. Create a notebook
const nb = await ( await fetch ( ` ${ API } /notebooks` , {
method : "POST" ,
headers ,
body : JSON . stringify ({ display_name : "My Research" })
})). json ();
// 2. Add a source
await fetch ( ` ${ API } / ${ nb . name } /sources` , {
method : "POST" ,
headers ,
body : JSON . stringify ({
display_name : "Gemini API Guide" ,
web_content : { url : "https://ai.google.dev/gemini-api/docs" }
})
});
REST
# 1. Create a notebook
NOTEBOOK = $( curl -s -X POST "https://notebook.googleapis.com/v1/notebooks" \
-H "Authorization: Bearer $ACCESS_TOKEN " \
-H "x-goog-user-project: $PROJECT_ID " \
-H "Content-Type: application/json" \
-d '{"display_name": "My Research"}' )
NOTEBOOK_NAME = $( echo " $NOTEBOOK " | python3 -c "import sys, json; print(json.load(sys.stdin)['name'])" )
# 2. Add a source
curl -s -X POST "https://notebook.googleapis.com/v1/ ${ NOTEBOOK_NAME } /sources" \
-H "Authorization: Bearer $ACCESS_TOKEN " \
-H "x-goog-user-project: $PROJECT_ID " \
-H "Content-Type: application/json" \
-d '{
"display_name": "Gemini API Guide",
"web_content": {"url": "https://ai.google.dev/gemini-api/docs"}
}'
Guides and resources
Create, list, and update notebooks.
Add web pages, YouTube videos, plain text, or Markdown to a notebook.
hub
Model Context Protocol (MCP)
Connect AI coding agents and IDEs to the remote Gemini Notebook MCP server.
Configure OAuth 2.0 access tokens and scopes for API requests.
Review request rate limits, per-source word limits, and tier capacity.
See items on the Gemini Notebook API roadmap.
We'd love to hear what you want the Gemini Notebook API and
MCP server to do. See our
roadmap for upcoming capabilities, and
file feature requests or bug reports in our Issue Tracker .
Except as otherwise noted, the content of this page is licensed under the Creative Commons Attribution 4.0 License , and code samples are licensed under the Apache 2.0 License . For details, see the Google Developers Site Policies . Java is a registered trademark of Oracle and/or its affiliates.

## Original Extract

Programmatically create and manage notebooks and sources with the Gemini Notebook API.

Gemini Notebook API | Google for Developers
Skip to main content
Gemini Notebook API
Gemini Notebook API
Stay organized with collections
Save and categorize content based on your preferences.
The Gemini Notebook API is a REST interface for programmatically creating and
managing notebooks and sources in Gemini Notebook. Ingest web pages, YouTube
videos, plain text, and Markdown into persistent notebooks.
Early Access Public Preview: The Gemini Notebook API and
MCP server are in an early access Public
Preview, and we're actively working on our
roadmap . We'd love to hear what you want
the API and MCP server to do. Please
file
feature requests in our Issue Tracker . We're excited to build this out,
and there's more to come!
During the preview, the API is available for personal Google Accounts.
Google Workspace accounts are not yet supported.
Just want to get started? Go to the
Quickstart , copy the setup prompt,
and give it to your coding agent.
import requests
API = "https://notebook.googleapis.com/v1"
headers = {
"Authorization" : f "Bearer { ACCESS_TOKEN } " ,
"x-goog-user-project" : PROJECT_ID ,
"Content-Type" : "application/json" ,
}
# 1. Create a notebook
nb = requests . post ( f " { API } /notebooks" , headers = headers , json = {
"display_name" : "My Research"
}) . json ()
# 2. Add a source
requests . post ( f " { API } / { nb [ 'name' ] } /sources" , headers = headers , json = {
"display_name" : "Gemini API Guide" ,
"web_content" : { "url" : "https://ai.google.dev/gemini-api/docs" }
})
JavaScript
const API = "https://notebook.googleapis.com/v1" ;
const headers = {
"Authorization" : `Bearer ${ ACCESS_TOKEN } ` ,
"x-goog-user-project" : PROJECT_ID ,
"Content-Type" : "application/json"
};
// 1. Create a notebook
const nb = await ( await fetch ( ` ${ API } /notebooks` , {
method : "POST" ,
headers ,
body : JSON . stringify ({ display_name : "My Research" })
})). json ();
// 2. Add a source
await fetch ( ` ${ API } / ${ nb . name } /sources` , {
method : "POST" ,
headers ,
body : JSON . stringify ({
display_name : "Gemini API Guide" ,
web_content : { url : "https://ai.google.dev/gemini-api/docs" }
})
});
REST
# 1. Create a notebook
NOTEBOOK = $( curl -s -X POST "https://notebook.googleapis.com/v1/notebooks" \
-H "Authorization: Bearer $ACCESS_TOKEN " \
-H "x-goog-user-project: $PROJECT_ID " \
-H "Content-Type: application/json" \
-d '{"display_name": "My Research"}' )
NOTEBOOK_NAME = $( echo " $NOTEBOOK " | python3 -c "import sys, json; print(json.load(sys.stdin)['name'])" )
# 2. Add a source
curl -s -X POST "https://notebook.googleapis.com/v1/ ${ NOTEBOOK_NAME } /sources" \
-H "Authorization: Bearer $ACCESS_TOKEN " \
-H "x-goog-user-project: $PROJECT_ID " \
-H "Content-Type: application/json" \
-d '{
"display_name": "Gemini API Guide",
"web_content": {"url": "https://ai.google.dev/gemini-api/docs"}
}'
Guides and resources
Create, list, and update notebooks.
Add web pages, YouTube videos, plain text, or Markdown to a notebook.
hub
Model Context Protocol (MCP)
Connect AI coding agents and IDEs to the remote Gemini Notebook MCP server.
Configure OAuth 2.0 access tokens and scopes for API requests.
Review request rate limits, per-source word limits, and tier capacity.
See items on the Gemini Notebook API roadmap.
We'd love to hear what you want the Gemini Notebook API and
MCP server to do. See our
roadmap for upcoming capabilities, and
file feature requests or bug reports in our Issue Tracker .
Except as otherwise noted, the content of this page is licensed under the Creative Commons Attribution 4.0 License , and code samples are licensed under the Apache 2.0 License . For details, see the Google Developers Site Policies . Java is a registered trademark of Oracle and/or its affiliates.
