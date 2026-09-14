---
source: "https://www.costclaw.io/"
hn_url: "https://news.ycombinator.com/item?id=49700895"
title: "Show HN: CostClaw, a local audit of your Claude Code logs"
article_title: "CostClaw: see where your Claude Code spend leaks"
image: "https://www.costclaw.io/og.png"
author: "practicalsystem"
captured_at: "2026-09-14T18:02:23Z"
capture_tool: "hn-digest"
hn_id: 49700895
score: 1
comments: 0
posted_at: "2026-09-14T17:44:14Z"
tags:
  - hacker-news
---

# Show HN: CostClaw, a local audit of your Claude Code logs

- HN: [49700895](https://news.ycombinator.com/item?id=49700895)
- Source: [www.costclaw.io](https://www.costclaw.io/)
- Score: 1
- Comments: 0
- Posted: 2026-09-14T17:44:14Z

## Translation

Title: Show HN: CostClaw, a local audit of your Claude Code logs
Article title: CostClaw: see where your Claude Code spend leaks
Description: A local audit for Claude Code. It finds recoverable token spend, scores your setup across six pillars, and never sees your prompts.

Article text:
Skip to content Cost Claw The leak How it works Privacy Pricing Get started
Local audit for Claude Code
See where your Claude Code spend leaks.
CostClaw reads the Claude Code logs already on your disk, finds the token
spend you can get back, and scores your setup across six pillars. It runs
on your machine and never reads your prompts.
Copy
See what it finds
No account, no upload. It parses the logs on your disk and prints a
report.
costclaw audit read-only · nothing uploaded ~/.claude/projects 1840 sessions · 41 projects · CLAUDE.md found
Model spend opus 58% / sonnet 33% / haiku 9%
$52.40 Use a cheaper model for 61 small sessions
$18.90 Compact four long sessions that lost the cache
warn Workflow and subagent runs cost $214
warn Three projects fire one tool 100+ times per session
Your invoice is one number. The waste is in the logs.
Anthropic bills you a monthly total. It cannot tell you which dollars you
could have kept. CostClaw reads the session logs on your disk, including
the subagent and workflow sub-logs, and prices the waste the invoice
hides. Four of those leaks:
Claude Code caches the front of your prompt so repeat tokens cost a
tenth of the price. When a session rebuilds that cache instead of
reusing it, you pay full input rate on tokens that should have been
nearly free. CostClaw measures the gap and prices it.
A session that fires the same tool a hundred times is usually thrashing:
re-reading, re-grepping, circling a problem. It does not show up as a
line item, but it burns tokens and time. CostClaw counts the loops and
flags the sessions that run hot.
A short, simple session does not need the top model. When a quick task
runs on Opus that a cheaper model would have served, you pay the Opus
rate for Sonnet work. CostClaw flags those sessions and prices the
difference against the cheaper model.
Sessions that outgrow the cache
A session that runs for hours keeps adding context until the cache stops
paying off. The early turns hit it; the closing turns re-pay full price
for a prompt that only grew. CostClaw tracks the per-turn cache and flags
the sessions that should have been compacted.
Three passes over logs you already have.
It reads the JSONL session logs under your Claude directory, including the subagent and workflow sub-logs, in memory, on your machine. Nothing is sent anywhere.
From the raw logs it computes totals: spend by model, cache hit rate per turn, tool use, session timing. Prompts, paths, and code are dropped here.
It scores six setup pillars from the evidence and your settings, then ranks fixes by the dollars they recover. The report prints to your terminal.
Optional, on the Pro plan. It only ever receives the derived record
above. The free CLI sends nothing at all.
A cost report, with the recoverable part priced
Real spend over the window you audit, your cache hit rate, and the
slice you could get back. The recoverable figure is the headline, so
you know what fixing the leak is worth before you start.
A six-pillar setup score from evidence
CLAUDE.md quality, context hygiene, prompting, session management,
tooling, and cost discipline. Each pillar is scored only from what
your logs, settings, and CLAUDE.md actually show; the audit finds
your global CLAUDE.md on its own. A pillar with no evidence reads
n/a; nothing is invented to fill a bar.
Fixes ranked by the dollars they recover
Not a wall of warnings. The list leads with the fix that returns the
most spend, carries the exact figure and the evidence behind it, and
tells you the one change to make. Smaller issues sit below it, in
order.
$1,284.50 Recover spend lost to cache misses
$52.40 Use a cheaper model for small sessions
warn Workflow and subagent runs add up
The week your spend actually follows
Burn Clock buckets every dollar by local day and time of day, so you
can see the rhythm the bill hides. The peak day is compared against
an average day, never against your quietest hour, so the multiplier
stays a number you can check against the bars. Under seven active
days it reads n/a instead of guessing a pattern.
Tuesday is your most expensive day ($1,402.87),
about 1.6x an average day across 34 active days. Peak stretch: Tuesday afternoon.
The only thing CostClaw produces is an
AuditRecord : totals and generated
sentences. No prompt text, no file paths, no secrets. That boundary is
held by a test that runs on every build.
The test plants tripwire strings, a fake prompt and a fake client path,
into sample logs, runs the full pipeline, and asserts that none of them
appear in the record. If a leak ever slipped in, the build would fail
before it shipped.
// the privacy invariant, enforced on every run
const raw = sessionWith({
prompt: "TRIPWIRE_secret_prompt" ,
path: "/Users/me/clients/acme" ,
});
const record = buildAudit(analyze(raw));
const json = JSON.stringify(record);
expect(json). not .toContain( "TRIPWIRE_secret_prompt" );
expect(json). not .toContain( "/Users/me/clients/acme" );
expect(hasEmDash(json)).toBe( false ); Pricing
The CLI is real and usable right now, at no cost. The hosted plan is in
progress, and it will only ever see the derived record, never your logs.
Local cost report with the recoverable figure
Six-pillar setup score from your logs and CLAUDE.md
Fixes ranked by dollars recovered, with the evidence
Burn Clock: your spend rhythm by day and time of day
HTML report, score card, and README badge
Runs offline, nothing uploaded
Copy
Buy CostClaw - $39
One-time license that unlocks costclaw optimize .
It scans your project and writes five artifacts: a filled-in
CLAUDE.md, a ranked fix playbook, suggested settings, a work
order your own Claude Code can apply, and a dashboard that
opens in your browser with the token receipt, the fixes, and
your CLAUDE.md before and after. Verified offline. Your key is
emailed to you in about a minute.
Hosted dashboard, history over time
Shareable PDF report for billing clients
Receives only the derived record
Want it when it ships? Leave your email.
Not ready yet. The CLI does the core job today.
The audit reads what is already on your disk and prints the report. If it
finds nothing, you have lost a minute. If it finds the usual, you have
found real money.
Copy Reads local logs. Sends nothing. No account.
Cost Claw A local cost and setup audit for Claude Code.

## Original Extract

A local audit for Claude Code. It finds recoverable token spend, scores your setup across six pillars, and never sees your prompts.

Skip to content Cost Claw The leak How it works Privacy Pricing Get started
Local audit for Claude Code
See where your Claude Code spend leaks.
CostClaw reads the Claude Code logs already on your disk, finds the token
spend you can get back, and scores your setup across six pillars. It runs
on your machine and never reads your prompts.
Copy
See what it finds
No account, no upload. It parses the logs on your disk and prints a
report.
costclaw audit read-only · nothing uploaded ~/.claude/projects 1840 sessions · 41 projects · CLAUDE.md found
Model spend opus 58% / sonnet 33% / haiku 9%
$52.40 Use a cheaper model for 61 small sessions
$18.90 Compact four long sessions that lost the cache
warn Workflow and subagent runs cost $214
warn Three projects fire one tool 100+ times per session
Your invoice is one number. The waste is in the logs.
Anthropic bills you a monthly total. It cannot tell you which dollars you
could have kept. CostClaw reads the session logs on your disk, including
the subagent and workflow sub-logs, and prices the waste the invoice
hides. Four of those leaks:
Claude Code caches the front of your prompt so repeat tokens cost a
tenth of the price. When a session rebuilds that cache instead of
reusing it, you pay full input rate on tokens that should have been
nearly free. CostClaw measures the gap and prices it.
A session that fires the same tool a hundred times is usually thrashing:
re-reading, re-grepping, circling a problem. It does not show up as a
line item, but it burns tokens and time. CostClaw counts the loops and
flags the sessions that run hot.
A short, simple session does not need the top model. When a quick task
runs on Opus that a cheaper model would have served, you pay the Opus
rate for Sonnet work. CostClaw flags those sessions and prices the
difference against the cheaper model.
Sessions that outgrow the cache
A session that runs for hours keeps adding context until the cache stops
paying off. The early turns hit it; the closing turns re-pay full price
for a prompt that only grew. CostClaw tracks the per-turn cache and flags
the sessions that should have been compacted.
Three passes over logs you already have.
It reads the JSONL session logs under your Claude directory, including the subagent and workflow sub-logs, in memory, on your machine. Nothing is sent anywhere.
From the raw logs it computes totals: spend by model, cache hit rate per turn, tool use, session timing. Prompts, paths, and code are dropped here.
It scores six setup pillars from the evidence and your settings, then ranks fixes by the dollars they recover. The report prints to your terminal.
Optional, on the Pro plan. It only ever receives the derived record
above. The free CLI sends nothing at all.
A cost report, with the recoverable part priced
Real spend over the window you audit, your cache hit rate, and the
slice you could get back. The recoverable figure is the headline, so
you know what fixing the leak is worth before you start.
A six-pillar setup score from evidence
CLAUDE.md quality, context hygiene, prompting, session management,
tooling, and cost discipline. Each pillar is scored only from what
your logs, settings, and CLAUDE.md actually show; the audit finds
your global CLAUDE.md on its own. A pillar with no evidence reads
n/a; nothing is invented to fill a bar.
Fixes ranked by the dollars they recover
Not a wall of warnings. The list leads with the fix that returns the
most spend, carries the exact figure and the evidence behind it, and
tells you the one change to make. Smaller issues sit below it, in
order.
$1,284.50 Recover spend lost to cache misses
$52.40 Use a cheaper model for small sessions
warn Workflow and subagent runs add up
The week your spend actually follows
Burn Clock buckets every dollar by local day and time of day, so you
can see the rhythm the bill hides. The peak day is compared against
an average day, never against your quietest hour, so the multiplier
stays a number you can check against the bars. Under seven active
days it reads n/a instead of guessing a pattern.
Tuesday is your most expensive day ($1,402.87),
about 1.6x an average day across 34 active days. Peak stretch: Tuesday afternoon.
The only thing CostClaw produces is an
AuditRecord : totals and generated
sentences. No prompt text, no file paths, no secrets. That boundary is
held by a test that runs on every build.
The test plants tripwire strings, a fake prompt and a fake client path,
into sample logs, runs the full pipeline, and asserts that none of them
appear in the record. If a leak ever slipped in, the build would fail
before it shipped.
// the privacy invariant, enforced on every run
const raw = sessionWith({
prompt: "TRIPWIRE_secret_prompt" ,
path: "/Users/me/clients/acme" ,
});
const record = buildAudit(analyze(raw));
const json = JSON.stringify(record);
expect(json). not .toContain( "TRIPWIRE_secret_prompt" );
expect(json). not .toContain( "/Users/me/clients/acme" );
expect(hasEmDash(json)).toBe( false ); Pricing
The CLI is real and usable right now, at no cost. The hosted plan is in
progress, and it will only ever see the derived record, never your logs.
Local cost report with the recoverable figure
Six-pillar setup score from your logs and CLAUDE.md
Fixes ranked by dollars recovered, with the evidence
Burn Clock: your spend rhythm by day and time of day
HTML report, score card, and README badge
Runs offline, nothing uploaded
Copy
Buy CostClaw - $39
One-time license that unlocks costclaw optimize .
It scans your project and writes five artifacts: a filled-in
CLAUDE.md, a ranked fix playbook, suggested settings, a work
order your own Claude Code can apply, and a dashboard that
opens in your browser with the token receipt, the fixes, and
your CLAUDE.md before and after. Verified offline. Your key is
emailed to you in about a minute.
Hosted dashboard, history over time
Shareable PDF report for billing clients
Receives only the derived record
Want it when it ships? Leave your email.
Not ready yet. The CLI does the core job today.
The audit reads what is already on your disk and prints the report. If it
finds nothing, you have lost a minute. If it finds the usual, you have
found real money.
Copy Reads local logs. Sends nothing. No account.
Cost Claw A local cost and setup audit for Claude Code.
