---
source: "https://glyphh.ai"
hn_url: "https://news.ycombinator.com/item?id=50208345"
title: "Glyphh Ada – Precedent-based classifier and LLM fallthrough router"
article_title: "Glyphh — Ada, a classifier that learns from precedent"
image: "https://glyphh.ai/opengraph-image?abd9ec4af45c4a2e"
author: "christimmerman"
captured_at: "2026-10-23T13:07:36Z"
capture_tool: "hn-digest"
hn_id: 50208345
score: 1
comments: 1
posted_at: "2026-10-23T13:02:27Z"
tags:
  - hacker-news
---

# Glyphh Ada – Precedent-based classifier and LLM fallthrough router

- HN: [50208345](https://news.ycombinator.com/item?id=50208345)
- Source: [glyphh.ai](https://glyphh.ai)
- Score: 1
- Comments: 1
- Posted: 2026-10-23T13:02:27Z

## Translation

Title: Glyphh Ada – Precedent-based classifier and LLM fallthrough router
Article title: Glyphh — Ada, a classifier that learns from precedent
Description: Ada learns from your labelled examples. A new case is answered from its nearest precedents, with a calibrated confidence and the precedents it stands on; when Ada is unsure, your application hands the case to an LLM, and that answer becomes Ada’s next time. It runs in Ada Cloud, isolated per organis
[truncated]

Article text:
Glyphh — Ada, a classifier that learns from precedent Home Use cases Support routing Each ticket to the team that owns it, from the tickets you have already routed. Loan decisions Approve or refuse from your past decisions, by what each mistake costs you. Blog Pricing Docs Log in Sign up Home Use cases Support routing Loan decisions Blog Pricing Docs Log in Sign up Give your code the ability to reason without an LLM.
Glyphh Ada classifies anything you have labelled examples of — tickets, transactions, documents, events and more — returns a confidence, probability gates, and the precedents behind each answer, so every decision can be audited.
A case is typed JSON in the shape your spec gives it. Recorded with the outcome the world confirmed, it becomes a precedent, and counts from the next question.
A calibrated probability for each outcome, how many precedents voted, and their receipts. When nothing is close enough, Ada says so and asks.
Sharper with every graded case
Each graded case is a new precedent, and Ada learns its own weights from them: the parts of a case that predict the outcome count for more.
An answer you can check, not a guess you have to trust.
Ada does not generate. It finds the precedents closest to the case, weighs them, and reports what they say. A single precedent is not consensus, and the answer says how many stand behind it.
incidents.decide({
"error": {
"what": { "class": "TimeoutError",
"message": "payment gateway timed out",
"status": 504 }
}
}) Ada answers (example)
{
"fix": "raise circuit breaker",
"confidence": 0.82,
"from": 6,
"last_time": {
"cause": "gateway rate limit hit on retry storm",
"owner": "payments",
"runbook": "RB-12 back off retries, raise circuit breaker"
}
} Where it runs Getting started → Ada Cloud
Your models, in Glyphh’s cloud
Your models live and run in Glyphh’s cloud, isolated per organisation. Nothing to install or run. We never train on your data.
Add https://api.glyphh.ai/mcp as a connector, sign in and click Allow. The glyphh plugin’s Ada skill builds the model with you and asks it.
pip install glyphh or npm install @glyphh-ai/glyphh, with an API key from the platform, calling https://ada.glyphh.ai.
Each caller sees what they may
A spec can put layers behind a clearance. A status bot gets the fix; an engineer also gets the cause, the owner and the runbook.
Claude signs in with your Glyphh account: no key to type. Ada is in every plan, Free included, and never priced per seat, per model or per call.
/plugin marketplace add glyphh-ai/glyphh
/plugin install glyphh@glyphh
# then add the Glyphh connector, https://api.glyphh.ai/mcp,
# and ask Claude to build a model with you github.com/glyphh-ai/glyphh
Ada, a classifier that learns from precedent and hands the cases it is unsure of to an LLM.

## Original Extract

Ada learns from your labelled examples. A new case is answered from its nearest precedents, with a calibrated confidence and the precedents it stands on; when Ada is unsure, your application hands the case to an LLM, and that answer becomes Ada’s next time. It runs in Ada Cloud, isolated per organis
[truncated]

Glyphh — Ada, a classifier that learns from precedent Home Use cases Support routing Each ticket to the team that owns it, from the tickets you have already routed. Loan decisions Approve or refuse from your past decisions, by what each mistake costs you. Blog Pricing Docs Log in Sign up Home Use cases Support routing Loan decisions Blog Pricing Docs Log in Sign up Give your code the ability to reason without an LLM.
Glyphh Ada classifies anything you have labelled examples of — tickets, transactions, documents, events and more — returns a confidence, probability gates, and the precedents behind each answer, so every decision can be audited.
A case is typed JSON in the shape your spec gives it. Recorded with the outcome the world confirmed, it becomes a precedent, and counts from the next question.
A calibrated probability for each outcome, how many precedents voted, and their receipts. When nothing is close enough, Ada says so and asks.
Sharper with every graded case
Each graded case is a new precedent, and Ada learns its own weights from them: the parts of a case that predict the outcome count for more.
An answer you can check, not a guess you have to trust.
Ada does not generate. It finds the precedents closest to the case, weighs them, and reports what they say. A single precedent is not consensus, and the answer says how many stand behind it.
incidents.decide({
"error": {
"what": { "class": "TimeoutError",
"message": "payment gateway timed out",
"status": 504 }
}
}) Ada answers (example)
{
"fix": "raise circuit breaker",
"confidence": 0.82,
"from": 6,
"last_time": {
"cause": "gateway rate limit hit on retry storm",
"owner": "payments",
"runbook": "RB-12 back off retries, raise circuit breaker"
}
} Where it runs Getting started → Ada Cloud
Your models, in Glyphh’s cloud
Your models live and run in Glyphh’s cloud, isolated per organisation. Nothing to install or run. We never train on your data.
Add https://api.glyphh.ai/mcp as a connector, sign in and click Allow. The glyphh plugin’s Ada skill builds the model with you and asks it.
pip install glyphh or npm install @glyphh-ai/glyphh, with an API key from the platform, calling https://ada.glyphh.ai.
Each caller sees what they may
A spec can put layers behind a clearance. A status bot gets the fix; an engineer also gets the cause, the owner and the runbook.
Claude signs in with your Glyphh account: no key to type. Ada is in every plan, Free included, and never priced per seat, per model or per call.
/plugin marketplace add glyphh-ai/glyphh
/plugin install glyphh@glyphh
# then add the Glyphh connector, https://api.glyphh.ai/mcp,
# and ask Claude to build a model with you github.com/glyphh-ai/glyphh
Ada, a classifier that learns from precedent and hands the cases it is unsure of to an LLM.
