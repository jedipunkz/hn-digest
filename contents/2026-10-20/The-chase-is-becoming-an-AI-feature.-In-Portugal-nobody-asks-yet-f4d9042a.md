---
source: "https://gil-neto.com/blog/the-chase-is-becoming-an-ai-feature/"
hn_url: "https://news.ycombinator.com/item?id=50162912"
title: "The chase is becoming an AI feature. In Portugal, nobody asks yet"
article_title: "The chase is becoming an AI feature. In Portugal, nobody asks yet — Gil Neto"
image: "https://gil-neto.com/og-image.png"
author: "gilneto8"
captured_at: "2026-10-20T09:12:10Z"
capture_tool: "hn-digest"
hn_id: 50162912
score: 2
comments: 0
posted_at: "2026-10-20T08:24:35Z"
tags:
  - hacker-news
---

# The chase is becoming an AI feature. In Portugal, nobody asks yet

- HN: [50162912](https://news.ycombinator.com/item?id=50162912)
- Source: [gil-neto.com](https://gil-neto.com/blog/the-chase-is-becoming-an-ai-feature/)
- Score: 2
- Comments: 0
- Posted: 2026-10-20T08:24:35Z

## Translation

Title: The chase is becoming an AI feature. In Portugal, nobody asks yet
Article title: The chase is becoming an AI feature. In Portugal, nobody asks yet — Gil Neto
Description: Xero and Pennylane are automating the most hated monthly job in a small practice: asking clients for missing documents. Portuguese software receives documents but never asks for them, and the useful version of the ask starts at the bank line.

Article text:
The chase is becoming an AI feature. In Portugal, nobody asks yet — Gil Neto
← All writing accounting ai-agents portugal reconciliation The chase is becoming an AI feature. In Portugal, nobody asks yet
In a small accounting practice there’s a monthly job nobody likes. Before the books can close, someone has to chase clients for what’s missing: the invoice behind a payment, the receipt for a card charge, the bank statement that never got sent. I’ve been on the other side of it too. My own accountant chases me to confirm invoices on e-fatura, the Portuguese tax authority’s invoice portal, and I’m exactly the kind of client who answers late.
This year the chase started turning into a software feature.
Xero announced “Document Requests” in its Partner Hub on 20 August. Its AI agent, JAX, will “request documents from the client, send reminders … then match it with transactions in Xero” . Accounting Today noted the release “isn’t tied to a specific date”, and JAX is still described as beta, so I’d read it as announced and not shipped.
Pennylane already ships it in France. Reminders go out automatically on the 1st and 15th of the month, and recurring requests cover periodic documents “comme les relevés bancaires”.
Uku , an accounting-software vendor, published a survey this year where 68% of firms named chasing clients for missing documents as the task they’d most like to hand to an AI agent. It’s a vendor survey, “dozens of firms” across eight countries, with no sample size disclosed, so I read it as a direction and not a number.
So, the companies that own the ledger are taking the chase. That makes sense: the ledger already knows which transactions are missing a document.
Portugal: the software receives, but doesn’t ask
In Portugal a lot of small practices run on TOConline, which belongs to the Ordem dos Contabilistas Certificados (OCC), the accountants’ professional body. It has intake: a client emails documents to an address built from their tax number, and a filing robot matches them ( TOConline manual - the pages I found may be old). Sage has client uploads. What I couldn’t find, in TOConline, Sage or Cegid, is the other half: a feature that sends the request, or the reminder, from the accountant to the client. That was a search, not an audit, so I may have missed something.
There’s also a reason the written trail matters here that I haven’t seen anyone frame. According to the OCC’s own FAQ , an accountant whose client doesn’t hand over the documents can refuse to sign the tax returns, and the evidence named for that refusal is the written communications asking the client for them. (I read the FAQ, not the statute article it cites.) So in Portugal a chase is more than an admin chore - it’s also the accountant’s paper trail.
I build Damasqo Flow, former Kelaro, a tool that syncs a practice’s client bank accounts into the spreadsheet the practice already uses, and ties each import to the bank’s own balance. It’s one part of Damasqo, the set of tools I’m building around verified bank data. Last week I picked the chase as the first thing to build on top of it, and the reason is the bank feed.
A generic document request (“please send us your October documents”) is turning into a commodity, and whoever owns the ledger will ship it. However, if you already hold the bank lines, you never need to chase the bank statement, which is exactly what Pennylane chases twice a month. And the ask gets specific: one question for each bank line that doesn’t match anything.
bank line matched?
-43,20 € · 03 Oct · CARD PAYMENT ───────▶ │
yes ──┴── no
│ │
done ▼
"What is this 43,20 € on 3 October?
Send the invoice to <your intake address>."
│
▼
question + reply kept per client
Three things I like about that shape:
The client answers a question they can actually answer, about a line they can recognise.
The answer goes to the intake the practice already has, so nothing new sits between the client and the accountant.
Every question and every reply is written down, which is the trail the OCC asks for.
I haven’t built it. Flow is going through a test battery right now, and everything waits on that passing. This is the plan, written down while the market moves under it.
I don’t know yet what my pilot accountant actually chases. If most of their monthly chase is payroll inputs or cash receipts, the bank feed can’t see those, and the specific question covers less than I think. That’s the first thing I’m asking them.
TOConline could add requests at any time. If the ledger owner here ships the asking half, mine has to be better at the question.
The 68% is a vendor’s survey. I’m using it as a direction, and that’s how I’d want you to read it too.
One email when a post goes out. What moved in the world, and what I'm building about it.
You'll get one email to confirm. Unsubscribe from any email.
The architecture, the first working version and the handover. Remote across EU timezones, based in Lisbon.

## Original Extract

Xero and Pennylane are automating the most hated monthly job in a small practice: asking clients for missing documents. Portuguese software receives documents but never asks for them, and the useful version of the ask starts at the bank line.

The chase is becoming an AI feature. In Portugal, nobody asks yet — Gil Neto
← All writing accounting ai-agents portugal reconciliation The chase is becoming an AI feature. In Portugal, nobody asks yet
In a small accounting practice there’s a monthly job nobody likes. Before the books can close, someone has to chase clients for what’s missing: the invoice behind a payment, the receipt for a card charge, the bank statement that never got sent. I’ve been on the other side of it too. My own accountant chases me to confirm invoices on e-fatura, the Portuguese tax authority’s invoice portal, and I’m exactly the kind of client who answers late.
This year the chase started turning into a software feature.
Xero announced “Document Requests” in its Partner Hub on 20 August. Its AI agent, JAX, will “request documents from the client, send reminders … then match it with transactions in Xero” . Accounting Today noted the release “isn’t tied to a specific date”, and JAX is still described as beta, so I’d read it as announced and not shipped.
Pennylane already ships it in France. Reminders go out automatically on the 1st and 15th of the month, and recurring requests cover periodic documents “comme les relevés bancaires”.
Uku , an accounting-software vendor, published a survey this year where 68% of firms named chasing clients for missing documents as the task they’d most like to hand to an AI agent. It’s a vendor survey, “dozens of firms” across eight countries, with no sample size disclosed, so I read it as a direction and not a number.
So, the companies that own the ledger are taking the chase. That makes sense: the ledger already knows which transactions are missing a document.
Portugal: the software receives, but doesn’t ask
In Portugal a lot of small practices run on TOConline, which belongs to the Ordem dos Contabilistas Certificados (OCC), the accountants’ professional body. It has intake: a client emails documents to an address built from their tax number, and a filing robot matches them ( TOConline manual - the pages I found may be old). Sage has client uploads. What I couldn’t find, in TOConline, Sage or Cegid, is the other half: a feature that sends the request, or the reminder, from the accountant to the client. That was a search, not an audit, so I may have missed something.
There’s also a reason the written trail matters here that I haven’t seen anyone frame. According to the OCC’s own FAQ , an accountant whose client doesn’t hand over the documents can refuse to sign the tax returns, and the evidence named for that refusal is the written communications asking the client for them. (I read the FAQ, not the statute article it cites.) So in Portugal a chase is more than an admin chore - it’s also the accountant’s paper trail.
I build Damasqo Flow, former Kelaro, a tool that syncs a practice’s client bank accounts into the spreadsheet the practice already uses, and ties each import to the bank’s own balance. It’s one part of Damasqo, the set of tools I’m building around verified bank data. Last week I picked the chase as the first thing to build on top of it, and the reason is the bank feed.
A generic document request (“please send us your October documents”) is turning into a commodity, and whoever owns the ledger will ship it. However, if you already hold the bank lines, you never need to chase the bank statement, which is exactly what Pennylane chases twice a month. And the ask gets specific: one question for each bank line that doesn’t match anything.
bank line matched?
-43,20 € · 03 Oct · CARD PAYMENT ───────▶ │
yes ──┴── no
│ │
done ▼
"What is this 43,20 € on 3 October?
Send the invoice to <your intake address>."
│
▼
question + reply kept per client
Three things I like about that shape:
The client answers a question they can actually answer, about a line they can recognise.
The answer goes to the intake the practice already has, so nothing new sits between the client and the accountant.
Every question and every reply is written down, which is the trail the OCC asks for.
I haven’t built it. Flow is going through a test battery right now, and everything waits on that passing. This is the plan, written down while the market moves under it.
I don’t know yet what my pilot accountant actually chases. If most of their monthly chase is payroll inputs or cash receipts, the bank feed can’t see those, and the specific question covers less than I think. That’s the first thing I’m asking them.
TOConline could add requests at any time. If the ledger owner here ships the asking half, mine has to be better at the question.
The 68% is a vendor’s survey. I’m using it as a direction, and that’s how I’d want you to read it too.
One email when a post goes out. What moved in the world, and what I'm building about it.
You'll get one email to confirm. Unsubscribe from any email.
The architecture, the first working version and the handover. Remote across EU timezones, based in Lisbon.
