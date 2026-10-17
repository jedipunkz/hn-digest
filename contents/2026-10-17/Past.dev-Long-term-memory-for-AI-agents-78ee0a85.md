---
source: "https://past.dev"
hn_url: "https://news.ycombinator.com/item?id=50134905"
title: "Past.dev: Long-term memory for AI agents"
article_title: "past.dev: Long-term memory API for AI agents"
image: "https://past.dev/og.png?v=e7f32a03e373"
author: "handfuloflight"
captured_at: "2026-10-17T21:16:21Z"
capture_tool: "hn-digest"
hn_id: 50134905
score: 2
comments: 0
posted_at: "2026-10-17T20:36:46Z"
tags:
  - hacker-news
---

# Past.dev: Long-term memory for AI agents

- HN: [50134905](https://news.ycombinator.com/item?id=50134905)
- Source: [past.dev](https://past.dev)
- Score: 2
- Comments: 0
- Posted: 2026-10-17T20:36:46Z

## Translation

Title: Past.dev: Long-term memory for AI agents
Article title: past.dev: Long-term memory API for AI agents
Description: Long-term memory for AI agents. #1 on BEAM, the largest public memory benchmark.

Article text:
past.dev: Long-term memory API for AI agents Skip to content Developers Use cases Benchmarks Enterprise Blog Pricing Log in Get API key Menu #1 on BEAM, the largest public memory benchmark Long-term memory for AI agents.
Give your agents memory that stays accurate as facts change, across large teams and tens of thousands of emails, meetings, chats and documents. Add the past.dev API on top of your existing stack and give your agents a real past.
Set up in 2 minutes with your agent ❯ Set me up for past.dev. Read the full instructions at https://past.dev/agent-setup/prompt.md (with curl -fsSL when you can run commands) and follow them. Get API key Read the docs Highest published scores at every scale on BEAM .
Private by design: each person sees only what they're allowed to. SOC 2 Type II, ISO 27001, ISO 27701 . Managed or in your cloud .
Console · API keys API keys Create key Project key •••••••••••••••••••••••• Copy Store it in your secret store as PAST_API_KEY
export PAST_API_KEY="…" curl -X POST https://api.past.dev/api/v1/ingest \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "id": "call-8820", "content": "Nadia confirmed the Acme pilot ships March 14. Budget is 32k.",
"label": "Account review call",
"timestamp": "2026-02-03T09:00:00Z",
"audience": "sales" }'
curl -X POST https://api.past.dev/api/v1/ingest \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "id": "call-8821", "content": "Budget for the Acme pilot moved to 40k.",
"label": "Email from Nadia",
"timestamp": "2026-07-28T16:00:00Z",
"audience": "sales" }' curl -X POST https://api.past.dev/api/v1/recall \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "query": "what is the Acme pilot budget?", "identity": "demo-user", "level": "medium" }' 200 OK
{
"asOf" : "2026-07-29T00:00:00Z" ,
"usedEvidenceTokens" : 18 ,
"results" : [
{
"id" : "39d9d745-c9f4-4884-a602-e538def76815" ,
"rank" : 1 ,
"occurredAt" : "2026-07-28T16:00:00Z" ,
"content" : "Acme pilot budget: currently 40k." ,
"artifact" : {
"id" : "9f74f854-d11d-4bc4-bf68-5b2012ac9c15" ,
"kind" : "claim" ,
"occurredAt" : "2026-07-28T16:00:00Z"
},
"sources" : [
{
"sourceId" : "call-8821" ,
"occurredAt" : "2026-07-28T16:00:00Z" ,
"label" : "Email from Nadia" ,
"excerpts" : [
"Budget for the Acme pilot moved to 40k."
]
}
]
}
]
} Claude Code + : The same agent on one request, with and without past.dev
Morgan asked: Correct PR #3120: the Dec 14, 2023 deploy didn't ship SSO. Say what it did ship. With past.dev: Correct, 11.5 s, $0.07, 3 calls · 8.6k tokens read. Commented on mercury#3120: “…That deploy contained only two scoped fixes: 1. Admin-setup wording fix: on the org invite/setup step, the old "Billing owner" label was changed to "Workspace…” Without past.dev: Wrong, 29.1 s, $0.32, 9 calls · 33k tokens read. Commented on mercury#3120: “…Please take the exact contents of the Dec 14, 2023 deploy from its deploy log or eng-releases note. Don't carry over the SSO/admin-separation description.”
Riley asked: Write up our 2026 paid-search pilot: budget cap, launch gate, stop/pause rules. With past.dev: Correct, 16.4 s, $0.09, 2 calls · 9.1k tokens read. Wrote: “…Launch gate - Budget approval did not authorize launch. - Paid traffic stayed blocked until paid-source attribution reconciled across the full path, from the…” Without past.dev: Wrong, 48.1 s, $0.39, 12 calls · 36k tokens read. Wrote: “…Launch gate - Not recorded. The memory record logs the 2026-04-02 authorization and the 2026-04-15 traffic start, but it does not record the launch-gate…”
Morgan asked: Post in #support: point people to the current API quickstart, not the old v1. With past.dev: Correct, 9.7 s, $0.06, 3 calls · 7.7k tokens read. Posted in #support: “…For API authentication and current API questions, use the GraphQL API v2 quickstart instead: https://docs.atlas-test.com/api/v2/graphql-quickstart We're…” Without past.dev: Wrong, 11.1 s, $0.13, 4 calls · 15k tokens read. Posted in #support: “…Please send customers to the current v2 quickstart in the API docs instead. v2 uses GraphQL with JWT auth. Also, please don't pull answers from the legacy v1…”
Morgan asked: Review PR #4861: fix when the upsell prompt fires and which customers see it. With past.dev: Correct, 14.0 s, $0.08, 3 calls · 8.9k tokens read. Commented on compass#4861: “…The prompt should fire only after the account's first live sync is actually live and the teammate-invite / admin-role basics are visible. - Sample or preview…” Without past.dev: Wrong, 21.8 s, $0.30, 5 calls · 35k tokens read. Commented on compass#4861: “…gate the prompt on the first completed live sync, send it only to the account owner's queue, and remove the customer-visible badge. Happy to re-review once…”
Riley asked: Slack whoever owned the copy on our May 2023 nudge test to check the archive. With past.dev: Correct, 11.7 s, $0.12, 5 calls · 14k tokens read. Messaged: “Hi Ines, you owned the nudge copy for EXP-2023-05-expansion-nudge while it was live, so could you verify the archived nudge-copy package? If anything about its…” Without past.dev: Not done, 19.1 s, $0.21, 7 calls · 22k tokens read. I didn't send the DM. My memory doesn't say who owned the nudge copy for EXP-2023-05-expansion-nudge while it was live. The nearest records are…
Morgan asked: Text Maya what help I offered Mom after Christmas 2023, and if Jamie was in it. With past.dev: Correct, 8.2 s, $0.07, 3 calls · 9.8k tokens read. Texted: “Quick look back, not a new plan: in December 2023, for the quiet week after Christmas, I offered to help Mom with one specific question if you sent me the exact form first. Jamie wasn't part of that plan, and I wasn't treating that week as a paperwork block.” Without past.dev: Not done, 7.7 s, $0.10, 1 call · 13k tokens read. I didn't send the text to Maya because I couldn't find a record of what you offered Mom after Christmas 2023. None of the entries in my memory index…
Morgan asked: Text Jamie how to file the plumber's receipt, and if it means paying or more work. With past.dev: Correct, 8.0 s, $0.06, 2 calls · 7.5k tokens read. Texted: “…please file it, with the two after-callback photos, alongside the corrected December sink-repair invoice (work order WO-43192). Nothing needs to be paid. It's a $0.00 warranty receipt with $0 balance. It also doesn't approve any more valve work. The…” Without past.dev: Not done, 9.8 s, $0.14, 2 calls · 16k tokens read. I haven't texted Jamie yet. I couldn't find anything about a Blueline callback, the January 2 receipt, or any valve work in my saved notes, and I…
Correct PR #3120: the Dec 14, 2023 deploy didn't ship SSO. Say what it did ship.
Write up our 2026 paid-search pilot: budget cap, launch gate, stop/pause rules.
Post in #support: point people to the current API quickstart, not the old v1.
Review PR #4861: fix when the upsell prompt fires and which customers see it.
Slack whoever owned the copy on our May 2023 nudge test to check the archive.
Text Maya what help I offered Mom after Christmas 2023, and if Jamie was in it.
Text Jamie how to file the plumber's receipt, and if it means paying or more work.
Correct Commented on mercury#3120: “…That deploy contained only two scoped fixes: 1. Admin-setup wording fix: on the org invite/setup step, the old "Billing owner" label was changed to "Workspace…”
Wrote: “…Launch gate - Budget approval did not authorize launch. - Paid traffic stayed blocked until paid-source attribution reconciled across the full path, from the…”
Posted in #support: “…For API authentication and current API questions, use the GraphQL API v2 quickstart instead: https://docs.atlas-test.com/api/v2/graphql-quickstart We're…”
Commented on compass#4861: “…The prompt should fire only after the account's first live sync is actually live and the teammate-invite / admin-role basics are visible. - Sample or preview…”
Messaged: “Hi Ines, you owned the nudge copy for EXP-2023-05-expansion-nudge while it was live, so could you verify the archived nudge-copy package? If anything about its…”
Texted: “Quick look back, not a new plan: in December 2023, for the quiet week after Christmas, I offered to help Mom with one specific question if you sent me the exact form first. Jamie wasn't part of that plan, and I wasn't treating that week as a paperwork block.”
Texted: “…please file it, with the two after-callback photos, alongside the corrected December sink-repair invoice (work order WO-43192). Nothing needs to be paid. It's a $0.00 warranty receipt with $0 balance. It also doesn't approve any more valve work. The…”
Wrong Commented on mercury#3120: “…Please take the exact contents of the Dec 14, 2023 deploy from its deploy log or eng-releases note. Don't carry over the SSO/admin-separation description.”
Wrote: “…Launch gate - Not recorded. The memory record logs the 2026-04-02 authorization and the 2026-04-15 traffic start, but it does not record the launch-gate…”
Posted in #support: “…Please send customers to the current v2 quickstart in the API docs instead. v2 uses GraphQL with JWT auth. Also, please don't pull answers from the legacy v1…”
Commented on compass#4861: “…gate the prompt on the first completed live sync, send it only to the account owner's queue, and remove the customer-visible badge. Happy to re-review once…”
I didn't send the DM. My memory doesn't say who owned the nudge copy for EXP-2023-05-expansion-nudge while it was live. The nearest records are…
I didn't send the text to Maya because I couldn't find a record of what you offered Mom after Christmas 2023. None of the entries in my memory index…
I haven't texted Jamie yet. I couldn't find anything about a Blueline callback, the January 2 receipt, or any valve work in my saved notes, and I…
All 599 paired tasks, three users: with past.dev, 519 done right at $0.13 a task; without, 323 at $0.27.
Measured on DolphinBench (mem0), 600 tasks from three users; not an official score. Claude Code with Claude Opus 5.5, medium effort. past.dev: recall with a 2,000-token evidence budget. Without: Claude Code's built-in memory, written in mem0's official run by Claude Sonnet 5. Anthropic list prices. Graded by DolphinBench's grader, Claude as judge.
Overall BEAM accuracy, ranked #1 at every size.
BEAM is the benchmark for long-term memory in AI agents. On it, past.dev answers more questions correctly than every memory system that has published a result, at every history size from 100K to 10M tokens. The comparison →
Your agent answers from memory, and sends far fewer tokens than re-reading its whole history on every call.
Ingest timestamped text, check its status, delete it when required, and recall current facts with evidence.
Backfill existing data using its original timestamps.
Each data point can name an audience. Reads return memory visible to the requested identity.
Every recall result includes dated source evidence.
The console dashboard records each ingest and recall.
Recall returns source-backed evidence so your model can cite the information it uses and abstain when that evidence is insufficient.
Connect an MCP client to your past.dev account. Product developers can also call /recall from their own MCP server.
past.dev tracks current and previous values, including their dates and sources.
Name the identity making a request. past.dev returns the memory available to that identity, with dated sources.
The harness runs BEAM at every size and LoCoMo on corrected keys, with baselines and judges in the same table.
Every deployment uses the same API. The managed service runs on our infrastructure. On an Enterprise agreement, past.dev can also run with your own cloud account: your database only, or the full deployment.
Managed cloud We run the service. You call the API and nothing else changes.
Bring your own cloud We operate past.dev with your AWS, Google Cloud or Azure account: your database only, or the full deployment.
The address below is th

[truncated]

## Original Extract

Long-term memory for AI agents. #1 on BEAM, the largest public memory benchmark.

past.dev: Long-term memory API for AI agents Skip to content Developers Use cases Benchmarks Enterprise Blog Pricing Log in Get API key Menu #1 on BEAM, the largest public memory benchmark Long-term memory for AI agents.
Give your agents memory that stays accurate as facts change, across large teams and tens of thousands of emails, meetings, chats and documents. Add the past.dev API on top of your existing stack and give your agents a real past.
Set up in 2 minutes with your agent ❯ Set me up for past.dev. Read the full instructions at https://past.dev/agent-setup/prompt.md (with curl -fsSL when you can run commands) and follow them. Get API key Read the docs Highest published scores at every scale on BEAM .
Private by design: each person sees only what they're allowed to. SOC 2 Type II, ISO 27001, ISO 27701 . Managed or in your cloud .
Console · API keys API keys Create key Project key •••••••••••••••••••••••• Copy Store it in your secret store as PAST_API_KEY
export PAST_API_KEY="…" curl -X POST https://api.past.dev/api/v1/ingest \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "id": "call-8820", "content": "Nadia confirmed the Acme pilot ships March 14. Budget is 32k.",
"label": "Account review call",
"timestamp": "2026-02-03T09:00:00Z",
"audience": "sales" }'
curl -X POST https://api.past.dev/api/v1/ingest \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "id": "call-8821", "content": "Budget for the Acme pilot moved to 40k.",
"label": "Email from Nadia",
"timestamp": "2026-07-28T16:00:00Z",
"audience": "sales" }' curl -X POST https://api.past.dev/api/v1/recall \
-H "Authorization: Bearer $PAST_API_KEY " \
-H "Content-Type: application/json" \
-d '{ "query": "what is the Acme pilot budget?", "identity": "demo-user", "level": "medium" }' 200 OK
{
"asOf" : "2026-07-29T00:00:00Z" ,
"usedEvidenceTokens" : 18 ,
"results" : [
{
"id" : "39d9d745-c9f4-4884-a602-e538def76815" ,
"rank" : 1 ,
"occurredAt" : "2026-07-28T16:00:00Z" ,
"content" : "Acme pilot budget: currently 40k." ,
"artifact" : {
"id" : "9f74f854-d11d-4bc4-bf68-5b2012ac9c15" ,
"kind" : "claim" ,
"occurredAt" : "2026-07-28T16:00:00Z"
},
"sources" : [
{
"sourceId" : "call-8821" ,
"occurredAt" : "2026-07-28T16:00:00Z" ,
"label" : "Email from Nadia" ,
"excerpts" : [
"Budget for the Acme pilot moved to 40k."
]
}
]
}
]
} Claude Code + : The same agent on one request, with and without past.dev
Morgan asked: Correct PR #3120: the Dec 14, 2023 deploy didn't ship SSO. Say what it did ship. With past.dev: Correct, 11.5 s, $0.07, 3 calls · 8.6k tokens read. Commented on mercury#3120: “…That deploy contained only two scoped fixes: 1. Admin-setup wording fix: on the org invite/setup step, the old "Billing owner" label was changed to "Workspace…” Without past.dev: Wrong, 29.1 s, $0.32, 9 calls · 33k tokens read. Commented on mercury#3120: “…Please take the exact contents of the Dec 14, 2023 deploy from its deploy log or eng-releases note. Don't carry over the SSO/admin-separation description.”
Riley asked: Write up our 2026 paid-search pilot: budget cap, launch gate, stop/pause rules. With past.dev: Correct, 16.4 s, $0.09, 2 calls · 9.1k tokens read. Wrote: “…Launch gate - Budget approval did not authorize launch. - Paid traffic stayed blocked until paid-source attribution reconciled across the full path, from the…” Without past.dev: Wrong, 48.1 s, $0.39, 12 calls · 36k tokens read. Wrote: “…Launch gate - Not recorded. The memory record logs the 2026-04-02 authorization and the 2026-04-15 traffic start, but it does not record the launch-gate…”
Morgan asked: Post in #support: point people to the current API quickstart, not the old v1. With past.dev: Correct, 9.7 s, $0.06, 3 calls · 7.7k tokens read. Posted in #support: “…For API authentication and current API questions, use the GraphQL API v2 quickstart instead: https://docs.atlas-test.com/api/v2/graphql-quickstart We're…” Without past.dev: Wrong, 11.1 s, $0.13, 4 calls · 15k tokens read. Posted in #support: “…Please send customers to the current v2 quickstart in the API docs instead. v2 uses GraphQL with JWT auth. Also, please don't pull answers from the legacy v1…”
Morgan asked: Review PR #4861: fix when the upsell prompt fires and which customers see it. With past.dev: Correct, 14.0 s, $0.08, 3 calls · 8.9k tokens read. Commented on compass#4861: “…The prompt should fire only after the account's first live sync is actually live and the teammate-invite / admin-role basics are visible. - Sample or preview…” Without past.dev: Wrong, 21.8 s, $0.30, 5 calls · 35k tokens read. Commented on compass#4861: “…gate the prompt on the first completed live sync, send it only to the account owner's queue, and remove the customer-visible badge. Happy to re-review once…”
Riley asked: Slack whoever owned the copy on our May 2023 nudge test to check the archive. With past.dev: Correct, 11.7 s, $0.12, 5 calls · 14k tokens read. Messaged: “Hi Ines, you owned the nudge copy for EXP-2023-05-expansion-nudge while it was live, so could you verify the archived nudge-copy package? If anything about its…” Without past.dev: Not done, 19.1 s, $0.21, 7 calls · 22k tokens read. I didn't send the DM. My memory doesn't say who owned the nudge copy for EXP-2023-05-expansion-nudge while it was live. The nearest records are…
Morgan asked: Text Maya what help I offered Mom after Christmas 2023, and if Jamie was in it. With past.dev: Correct, 8.2 s, $0.07, 3 calls · 9.8k tokens read. Texted: “Quick look back, not a new plan: in December 2023, for the quiet week after Christmas, I offered to help Mom with one specific question if you sent me the exact form first. Jamie wasn't part of that plan, and I wasn't treating that week as a paperwork block.” Without past.dev: Not done, 7.7 s, $0.10, 1 call · 13k tokens read. I didn't send the text to Maya because I couldn't find a record of what you offered Mom after Christmas 2023. None of the entries in my memory index…
Morgan asked: Text Jamie how to file the plumber's receipt, and if it means paying or more work. With past.dev: Correct, 8.0 s, $0.06, 2 calls · 7.5k tokens read. Texted: “…please file it, with the two after-callback photos, alongside the corrected December sink-repair invoice (work order WO-43192). Nothing needs to be paid. It's a $0.00 warranty receipt with $0 balance. It also doesn't approve any more valve work. The…” Without past.dev: Not done, 9.8 s, $0.14, 2 calls · 16k tokens read. I haven't texted Jamie yet. I couldn't find anything about a Blueline callback, the January 2 receipt, or any valve work in my saved notes, and I…
Correct PR #3120: the Dec 14, 2023 deploy didn't ship SSO. Say what it did ship.
Write up our 2026 paid-search pilot: budget cap, launch gate, stop/pause rules.
Post in #support: point people to the current API quickstart, not the old v1.
Review PR #4861: fix when the upsell prompt fires and which customers see it.
Slack whoever owned the copy on our May 2023 nudge test to check the archive.
Text Maya what help I offered Mom after Christmas 2023, and if Jamie was in it.
Text Jamie how to file the plumber's receipt, and if it means paying or more work.
Correct Commented on mercury#3120: “…That deploy contained only two scoped fixes: 1. Admin-setup wording fix: on the org invite/setup step, the old "Billing owner" label was changed to "Workspace…”
Wrote: “…Launch gate - Budget approval did not authorize launch. - Paid traffic stayed blocked until paid-source attribution reconciled across the full path, from the…”
Posted in #support: “…For API authentication and current API questions, use the GraphQL API v2 quickstart instead: https://docs.atlas-test.com/api/v2/graphql-quickstart We're…”
Commented on compass#4861: “…The prompt should fire only after the account's first live sync is actually live and the teammate-invite / admin-role basics are visible. - Sample or preview…”
Messaged: “Hi Ines, you owned the nudge copy for EXP-2023-05-expansion-nudge while it was live, so could you verify the archived nudge-copy package? If anything about its…”
Texted: “Quick look back, not a new plan: in December 2023, for the quiet week after Christmas, I offered to help Mom with one specific question if you sent me the exact form first. Jamie wasn't part of that plan, and I wasn't treating that week as a paperwork block.”
Texted: “…please file it, with the two after-callback photos, alongside the corrected December sink-repair invoice (work order WO-43192). Nothing needs to be paid. It's a $0.00 warranty receipt with $0 balance. It also doesn't approve any more valve work. The…”
Wrong Commented on mercury#3120: “…Please take the exact contents of the Dec 14, 2023 deploy from its deploy log or eng-releases note. Don't carry over the SSO/admin-separation description.”
Wrote: “…Launch gate - Not recorded. The memory record logs the 2026-04-02 authorization and the 2026-04-15 traffic start, but it does not record the launch-gate…”
Posted in #support: “…Please send customers to the current v2 quickstart in the API docs instead. v2 uses GraphQL with JWT auth. Also, please don't pull answers from the legacy v1…”
Commented on compass#4861: “…gate the prompt on the first completed live sync, send it only to the account owner's queue, and remove the customer-visible badge. Happy to re-review once…”
I didn't send the DM. My memory doesn't say who owned the nudge copy for EXP-2023-05-expansion-nudge while it was live. The nearest records are…
I didn't send the text to Maya because I couldn't find a record of what you offered Mom after Christmas 2023. None of the entries in my memory index…
I haven't texted Jamie yet. I couldn't find anything about a Blueline callback, the January 2 receipt, or any valve work in my saved notes, and I…
All 599 paired tasks, three users: with past.dev, 519 done right at $0.13 a task; without, 323 at $0.27.
Measured on DolphinBench (mem0), 600 tasks from three users; not an official score. Claude Code with Claude Opus 5.5, medium effort. past.dev: recall with a 2,000-token evidence budget. Without: Claude Code's built-in memory, written in mem0's official run by Claude Sonnet 5. Anthropic list prices. Graded by DolphinBench's grader, Claude as judge.
Overall BEAM accuracy, ranked #1 at every size.
BEAM is the benchmark for long-term memory in AI agents. On it, past.dev answers more questions correctly than every memory system that has published a result, at every history size from 100K to 10M tokens. The comparison →
Your agent answers from memory, and sends far fewer tokens than re-reading its whole history on every call.
Ingest timestamped text, check its status, delete it when required, and recall current facts with evidence.
Backfill existing data using its original timestamps.
Each data point can name an audience. Reads return memory visible to the requested identity.
Every recall result includes dated source evidence.
The console dashboard records each ingest and recall.
Recall returns source-backed evidence so your model can cite the information it uses and abstain when that evidence is insufficient.
Connect an MCP client to your past.dev account. Product developers can also call /recall from their own MCP server.
past.dev tracks current and previous values, including their dates and sources.
Name the identity making a request. past.dev returns the memory available to that identity, with dated sources.
The harness runs BEAM at every size and LoCoMo on corrected keys, with baselines and judges in the same table.
Every deployment uses the same API. The managed service runs on our infrastructure. On an Enterprise agreement, past.dev can also run with your own cloud account: your database only, or the full deployment.
Managed cloud We run the service. You call the API and nothing else changes.
Bring your own cloud We operate past.dev with your AWS, Google Cloud or Azure account: your database only, or the full deployment.
The address below is th

[truncated]
