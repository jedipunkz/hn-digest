---
source: "https://parserail.thecompound.tech"
hn_url: "https://news.ycombinator.com/item?id=49801589"
title: "Show HN: ParseRail, production AI endpoints behind one key and a credit wallet"
article_title: "ParseRail, The AI back-end for your product."
image: "https://parserail.thecompound.tech/og.jpg"
author: "kyisaiah47"
captured_at: "2026-09-22T14:23:27Z"
capture_tool: "hn-digest"
hn_id: 49801589
score: 1
comments: 0
posted_at: "2026-09-22T14:11:04Z"
tags:
  - hacker-news
---

# Show HN: ParseRail, production AI endpoints behind one key and a credit wallet

- HN: [49801589](https://news.ycombinator.com/item?id=49801589)
- Source: [parserail.thecompound.tech](https://parserail.thecompound.tech)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T14:11:04Z

## Translation

Title: Show HN: ParseRail, production AI endpoints behind one key and a credit wallet
Article title: ParseRail, The AI back-end for your product.
Description: Production AI endpoints behind one key and a pay-per-call credit wallet, the same engine the apps Compound ships run on, solved once and run in production. Parse documents, extract fields, redact PII, analyze contracts, fight chargebacks, enrich companies, live in minutes, no subscription.

Article text:
ParseRail, The AI back-end for your product. ParseRail api.thecompound.tech/v1 39 endpoints 1 paused 242 / 242 checks passing credits from $ 20 Docs Playground Endpoints Pricing Benchmarks Sign in Menu Docs Playground Endpoints Pricing Benchmarks Sign in api.thecompound.tech/v1 39 endpoints 1 paused 242 / 242 checks passing credits from $ 20 POST /v1/ parse $ 0 . 1 0 /document $0.10/document 10 credits, charged only on success. 60 calls a minute per key. Async capable. Quickstart Get your key Endpoints Pick an endpoint
Any invoice, receipt, EOB, ERA, or COI, PDF or image, into structured, validated JSON.
A production call, recorded 2026-09-01
charge credits ON SUCCESS ONLY
This call cost 10 credits, $0.10/document, and returned in 2.55 s.
INVOICE INV-10428 Acme Supplies Inc. 123 Warehouse Rd, Columbus OH Date: 2026-06-30 PO Number: PO-88 Bill to: Northwind Coffee LLC 12x Widget A @ $100.00 ... $1,200.00 Freight ... $55.50 TOTAL DUE: $1,255.50 Terms: Net 30 The response schema-valid · req_7e88e5dfd8ebeb94dd5530d8 // sending the call
// 0 credits charged so far curl TypeScript Python MCP Copy curl https://api.thecompound.tech/v1/parse \
-H "Authorization: Bearer $COMPOUND_API_KEY" \
-H "Content-Type: application/json" \
-d '{"fileUrl":"https://…/invoice.pdf"}' Recorded calls
Each response is the one production returned for its sample document on 2026-09-01 .
req_7e88e5dfd8ebeb94dd5530d8 Put it on the request line POST /v1/invoice Invoice extraction 8 cr 2.18 s req_52b8e5a9084ee8cac4751b0f Put it on the request line POST /v1/receipt Receipt extraction 6 cr 1.98 s req_69f649faea40aba962831a6d Put it on the request line POST /v1/statement Statement parsing 12 cr 2.36 s req_5100df980ffa3bd358606505 Put it on the request line POST /v1/contract Contract analysis 12 cr 2.42 s req_df46abf4ab51fbe0c190ea69 Put it on the request line POST /v1/resume Resume parsing 8 cr 2.23 s req_dceefd3c61e43e4c9c30890f Put it on the request line POST /v1/tables Table extraction 8 cr 1.42 s req_7c889ae7d0d0073c671a084a Put it on the request line POST /v1/split Document splitting 10 cr 2.08 s req_470184e8e02bcdb4d60d01f7 Put it on the request line Credits
Credits are bought up front, packs from $20, or a plan from $29/mo. 1 credit = $0.01 and a call is charged only when it succeeds, so a failed call always costs nothing.
Auto-recharge buys your chosen pack when the wallet falls below 200 credits.
8 calls on this page were recorded from production on 2026-09-01 . The playground on this site does not call the API.
Credits are charged on success only.
An API key is stored as a hash plus a short prefix.
The eval suite last ran against production on 2026-07-08 . Benchmarks
Whether it is up right now: Status
200 charged The call succeeded and burned its credits.
queued, running An async job in flight. It is charged when it succeeds.
failed Something failed on our side. Never charged.
503 paused Safe to retry with backoff, the call was never charged.
402 refused Your balance can't cover this call. Nothing was charged.
1 credit = 1¢. Every endpoint has a flat per-task price, an invoice parse is $0.08, and you're only charged when the task succeeds.
No. Credits are bought up front, a $20 pack or a plan from $29/mo, and a call is charged only when it succeeds, so a failed call costs nothing.
Is this a wrapper around someone else's API?
No. ParseRail is the same inference layer that runs every Compound Labs product in production, the receptionist that answers calls, the desk that fights chargebacks, the agent that chases invoices. We open our own engine to developers.
How accurate is document parsing?
Scored in the open. Every document engine has a public benchmark at /benchmarks, accuracy per field, per document kind, updated as the engine improves.
Submit as a job and get a webhook when it finishes. The console shows Requests, Jobs, and Webhooks with full payloads and replay.
Never. One wallet spends on any endpoint, packs and plans top up the same balance, and auto-recharge keeps unattended agents running.
ParseRail serves 39 endpoints and charges credits only on success. Endpoints are added, paused and repriced, and a paused one returns 503 and charges nothing. Leave an address and ParseRail writes when one changes.
The AI back-end for your product. A Compound Labs product.
Every document extraction API, side by side
Invoice OCR API pricing compared (September 2026)
How much does an invoice OCR API cost per month? (2026)
The best document parsing and extraction APIs in 2026
Receipt OCR API for developers: 7 compared (2026)
Free invoice and receipt OCR APIs: 16 free tiers compared (2026)

## Original Extract

Production AI endpoints behind one key and a pay-per-call credit wallet, the same engine the apps Compound ships run on, solved once and run in production. Parse documents, extract fields, redact PII, analyze contracts, fight chargebacks, enrich companies, live in minutes, no subscription.

ParseRail, The AI back-end for your product. ParseRail api.thecompound.tech/v1 39 endpoints 1 paused 242 / 242 checks passing credits from $ 20 Docs Playground Endpoints Pricing Benchmarks Sign in Menu Docs Playground Endpoints Pricing Benchmarks Sign in api.thecompound.tech/v1 39 endpoints 1 paused 242 / 242 checks passing credits from $ 20 POST /v1/ parse $ 0 . 1 0 /document $0.10/document 10 credits, charged only on success. 60 calls a minute per key. Async capable. Quickstart Get your key Endpoints Pick an endpoint
Any invoice, receipt, EOB, ERA, or COI, PDF or image, into structured, validated JSON.
A production call, recorded 2026-09-01
charge credits ON SUCCESS ONLY
This call cost 10 credits, $0.10/document, and returned in 2.55 s.
INVOICE INV-10428 Acme Supplies Inc. 123 Warehouse Rd, Columbus OH Date: 2026-06-30 PO Number: PO-88 Bill to: Northwind Coffee LLC 12x Widget A @ $100.00 ... $1,200.00 Freight ... $55.50 TOTAL DUE: $1,255.50 Terms: Net 30 The response schema-valid · req_7e88e5dfd8ebeb94dd5530d8 // sending the call
// 0 credits charged so far curl TypeScript Python MCP Copy curl https://api.thecompound.tech/v1/parse \
-H "Authorization: Bearer $COMPOUND_API_KEY" \
-H "Content-Type: application/json" \
-d '{"fileUrl":"https://…/invoice.pdf"}' Recorded calls
Each response is the one production returned for its sample document on 2026-09-01 .
req_7e88e5dfd8ebeb94dd5530d8 Put it on the request line POST /v1/invoice Invoice extraction 8 cr 2.18 s req_52b8e5a9084ee8cac4751b0f Put it on the request line POST /v1/receipt Receipt extraction 6 cr 1.98 s req_69f649faea40aba962831a6d Put it on the request line POST /v1/statement Statement parsing 12 cr 2.36 s req_5100df980ffa3bd358606505 Put it on the request line POST /v1/contract Contract analysis 12 cr 2.42 s req_df46abf4ab51fbe0c190ea69 Put it on the request line POST /v1/resume Resume parsing 8 cr 2.23 s req_dceefd3c61e43e4c9c30890f Put it on the request line POST /v1/tables Table extraction 8 cr 1.42 s req_7c889ae7d0d0073c671a084a Put it on the request line POST /v1/split Document splitting 10 cr 2.08 s req_470184e8e02bcdb4d60d01f7 Put it on the request line Credits
Credits are bought up front, packs from $20, or a plan from $29/mo. 1 credit = $0.01 and a call is charged only when it succeeds, so a failed call always costs nothing.
Auto-recharge buys your chosen pack when the wallet falls below 200 credits.
8 calls on this page were recorded from production on 2026-09-01 . The playground on this site does not call the API.
Credits are charged on success only.
An API key is stored as a hash plus a short prefix.
The eval suite last ran against production on 2026-07-08 . Benchmarks
Whether it is up right now: Status
200 charged The call succeeded and burned its credits.
queued, running An async job in flight. It is charged when it succeeds.
failed Something failed on our side. Never charged.
503 paused Safe to retry with backoff, the call was never charged.
402 refused Your balance can't cover this call. Nothing was charged.
1 credit = 1¢. Every endpoint has a flat per-task price, an invoice parse is $0.08, and you're only charged when the task succeeds.
No. Credits are bought up front, a $20 pack or a plan from $29/mo, and a call is charged only when it succeeds, so a failed call costs nothing.
Is this a wrapper around someone else's API?
No. ParseRail is the same inference layer that runs every Compound Labs product in production, the receptionist that answers calls, the desk that fights chargebacks, the agent that chases invoices. We open our own engine to developers.
How accurate is document parsing?
Scored in the open. Every document engine has a public benchmark at /benchmarks, accuracy per field, per document kind, updated as the engine improves.
Submit as a job and get a webhook when it finishes. The console shows Requests, Jobs, and Webhooks with full payloads and replay.
Never. One wallet spends on any endpoint, packs and plans top up the same balance, and auto-recharge keeps unattended agents running.
ParseRail serves 39 endpoints and charges credits only on success. Endpoints are added, paused and repriced, and a paused one returns 503 and charges nothing. Leave an address and ParseRail writes when one changes.
The AI back-end for your product. A Compound Labs product.
Every document extraction API, side by side
Invoice OCR API pricing compared (September 2026)
How much does an invoice OCR API cost per month? (2026)
The best document parsing and extraction APIs in 2026
Receipt OCR API for developers: 7 compared (2026)
Free invoice and receipt OCR APIs: 16 free tiers compared (2026)
