---
source: "https://prakhargurunani.com/blog/building-claude-observability/"
hn_url: "https://news.ycombinator.com/item?id=50053675"
title: "Building an Observability Dashboard for a Claude Subscription"
article_title: "Building an Observability Dashboard for a Claude Subscription | Prakhar Gurunani"
image: ""
author: "prakhargurunani"
captured_at: "2026-10-12T10:19:58Z"
capture_tool: "hn-digest"
hn_id: 50053675
score: 2
comments: 0
posted_at: "2026-10-12T09:27:32Z"
tags:
  - hacker-news
---

# Building an Observability Dashboard for a Claude Subscription

- HN: [50053675](https://news.ycombinator.com/item?id=50053675)
- Source: [prakhargurunani.com](https://prakhargurunani.com/blog/building-claude-observability/)
- Score: 2
- Comments: 0
- Posted: 2026-10-12T09:27:32Z

## Translation

Title: Building an Observability Dashboard for a Claude Subscription
Article title: Building an Observability Dashboard for a Claude Subscription | Prakhar Gurunani

Article text:
Building an Observability Dashboard for a Claude Subscription
Building an Observability Dashboard for a Claude Subscription
The problem What the answers look like
The collector Session files are undocumented, so parse an allowlist
Write to disk before touching the network
Privacy is a test, not a promise
Corporate networks break naive HTTPS
The server Idempotent ingest from two sources
Pinning machines to one account
Reconstructing the 5-hour blocks
How it was built Plan first, and write down the decisions
Measure the data before writing the parser
Verify against something independent
Where does a Claude subscription’s usage actually go?
A Claude Max subscription gives you a lot of Claude Code. It does not tell you
much about how you used it. /usage shows a percentage of the current window
and when it resets. That’s the whole report: no history, no sessions or models,
and nothing across machines.
API customers get an Admin API with usage and cost reports. Subscribers get
nothing like it. Meanwhile the data that would answer the question sits in
plain sight, in Claude Code’s session files, spread across every laptop,
desktop and remote dev box you use it on.
So I built claude-observability :
a small collector, claude-obs , that
runs on each machine, and a Next.js dashboard that turns what the collectors
upload into account-level analytics. It’s about 2,000 lines of TypeScript for
the collector (no runtime dependencies) and 6,500 for the app, plus a small
shared package that both depend on.
The interesting parts weren’t the charts. They were:
getting a number you can trust out of files that were never meant to be read
keeping prompts and code on the machine
running a background process on computers you don’t control
Source: github.com/FirePing32/claude-observability , MIT licensed
Collector: claude-obs on npm , Node 20+
Deploy the dashboard yourself (the README walks through it), sign in, create a
workspace for your Claude account, and generate an enrollment code. Then, on
every machine that runs Claude Code with that account:
npm install -g claude-obs
claude-obs login --code XXXX-XXXX --server https://<your-dashboard>
claude-obs sync # uploads your full history the first time
claude-obs install-agent # keeps it up to date in the background
Want to see exactly what leaves the machine before enrolling?
claude-obs sync --dry-run prints it and sends nothing.
The questions are simple to ask:
Am I about to hit the 5-hour limit? And how close did I get yesterday?
What is using it? Which sessions, which projects, which models? Is Opus
quietly doing work Sonnet could do?
Is the plan worth it? What would the same tokens have cost on the
pay-as-you-go API?
Is anything missing? Has a machine stopped reporting? Did we hit a limit
that the tracked usage doesn’t explain?
The constraints are what make it a project:
There is no API to pull from. The only source is local files, on each
machine, in a format Anthropic doesn’t document and changes at will.
The account is the unit. Limits apply to the account, not to a machine.
A dashboard that shows one laptop’s usage answers the wrong question.
Nothing sensitive leaves the machine. The session files hold every
prompt, every response, every file Claude read. The collector must upload
numbers and nothing else, and that has to be provable rather than promised.
No credentials. The collector must never touch the Claude login token.
That’s the overview for my account over the last 30 days. Each question above
gets a tile:
The 5-hour limit. The current block has used $40.67 of value with two
hours to go. That’s 25% of this account’s learned limit (about $165, from
three real limit hits). At the current burn rate of $16.80 an hour it will
reach about 45% by the reset.
Is the plan worth it? $2,289 of API-equivalent usage on a $100 plan this
cycle: 22.9×. It broke even the day after renewal.
What is using it? Opus accounts for 98% of the value. The daily chart also
shows the switch from Opus 5 to Opus 5.5 in late September, without anyone
having to remember when it happened.
Is anything missing? All five machines are reporting. But telemetry
coverage is 0%, so the totals are a few percent low, and claude.ai chat isn’t
measured at all. The page says so instead of hiding it.
Below the tiles, the first insight was the one that started the project: hit a
usage limit in 2 blocks this week, vs 1 the week before .
It’s a push model. Each machine runs claude-obs , which reads the bytes
Claude Code appended to its session files since the last run. It reduces them
to usage records and uploads them. The server merges uploads from every machine
into one account, prices each request, and serves the dashboard.
A workspace is one Claude account. Machines enroll into a workspace, and
everything shown is account-wide. There is deliberately no per-person view: the
goal is the right totals for the account, not a leaderboard.
Machines are the one breakdown the dashboard does offer, because a silent or
misconfigured collector is the most likely reason for wrong totals. For each
machine (names anonymized here), it shows:
when it was last active and last uploaded
A request reported by two machines counts once, under the first one.
The repo is an npm workspace with three packages:
packages/shared : the wire schema (zod) and the price book . The
collector validates what it sends with the same schema the server validates
what it receives, so the two can’t drift.
packages/cli : the collector, bundled into a single file with tsup .
apps/web : the Next.js 16 app. It holds the ingest endpoints, the dashboard,
the alerting and the database migrations.
Session files are undocumented, so parse an allowlist
Claude Code writes one JSON object per line to
~/.claude/projects/<project>/<session>.jsonl , plus a subagents/ folder per
session for subagent runs. There are a couple of dozen record types: assistant
messages, titles, compaction markers, turn durations, Claude Code’s own cost
snapshots, permission-mode changes and more. New ones appear between releases.
The parser dispatches on type and handles only the records it understands. It
builds every output record field by field from explicit paths : no object
spreading, no copying of message . Anything unknown is counted, not parsed. The
dashboard shows those counts under Data quality , so a format change shows up as
a number going up rather than as silently wrong totals.
The first real surprise came from counting before writing any parser logic.
Claude Code writes one line per content block of a response, and repeats
the response’s usage object on each. A reply with a text block and a tool call
is two lines with identical token counts. Sum them and you double-count. On my
own history, 3,301 assistant lines turned out to be 1,889 distinct requests.
Naive summing overcounted by about 75%.
It’s subtler than duplication. For 39 of those requests, output_tokens
differed between lines, and the last line was always the largest. Those are
streaming updates. So the rule is: group by requestId and keep the max,
never sum . The collector re-sends a request only if its output grew, which
keeps a live session accurate while a response is still streaming.
The server applies the same rule, so an out-of-order batch can’t lower a value:
// never let a concurrent, staler write lower the final streamed output
output: sql `greatest( ${ schema.apiRequests.output } , excluded.output)` ,
Other records need their own handling. Lines with model: "<synthetic>" are
not requests at all. They’re API errors and usage-limit notices that Claude Code
writes into the transcript. The collector classifies them locally into a
category ( prompt_too_long , rate_limit , a five-hour or weekly limit and so
on) and uploads only the category, never the text.
Session files grow to tens of megabytes. Re-reading them on every change would
be wasteful, so the collector keeps a byte offset per file, plus its device and
inode numbers:
It reads from the offset to EOF and consumes complete lines only . A
trailing line without a newline is still being written, so it waits for the
next pass.
It decodes each line from bytes on its own, so a multi-byte character split
across a read chunk is never mangled.
If the inode changed or the file shrank, the file was replaced, so it reads
from zero. That’s safe because the server deduplicates.
In watch mode it uses recursive fs.watch with a debounce, plus a full rescan
every 15 minutes as a safety net for missed events.
Write to disk before touching the network
The order of operations is the whole reliability story:
Write each batch to an outbox directory, one file per batch, via temp
file and rename.
Commit the new offsets to state.json , the same way.
Upload the outbox, in order, deleting each file only after the server
accepts it.
A crash between steps 2 and 3 means the lines are read again and queued twice.
The server stores each request once, so that’s harmless. A crash anywhere else
leaves the batch in the outbox for the next run. A laptop that’s offline for a
week simply builds up a queue. Every failure is a re-send, never a loss.
Upload failures are classified:
Network errors, 5xx and 429 back off and retry, honouring Retry-After .
401 (token revoked) stops the agent and asks for a new login.
422 (schema rejected) moves that one batch to a quarantine folder with
the reason, so a single bad record can’t block the queue forever.
Privacy is a test, not a promise
“We only upload usage” is easy to write in a README. The repo enforces it
three ways:
Field-by-field builders , as above. There’s no code path that copies an
object from the transcript into the output.
A canary test. The fixtures plant marker strings in message content,
e-mail addresses and file paths. The test fails if any of them appear anywhere
in the serialized upload.
claude-obs sync --dry-run prints exactly what would be sent, so anyone
can check for themselves before enrolling.
What does go up: model, timestamps, token counts (with the five-minute and
one-hour cache writes separated), effort, entry point, subagent type, the
project folder’s basename (or a salted hash of it), git branch, and session
titles (which can be switched off).
Corporate networks break naive HTTPS
The first machine outside my own failed immediately:
SELF_SIGNED_CERT_IN_CHAIN . Security agents like Netskope and Zscaler intercept
HTTPS and re-sign it with a company root certificate. The OS trusts that root.
Node, by default, doesn’t: it ships its own CA bundle.
Telling every user to set NODE_EXTRA_CA_CERTS would be a support burden. So
the collector trusts the operating system’s store automatically, the way a
browser does:
Newer Node (22.19+, 24.5+): tls.setDefaultCACertificates merges the
system certificates in-process.
Older Node on macOS: it exports the System keychain and Apple’s roots
with /usr/bin/security , caches the bundle for a week, and re-launches itself
once with NODE_EXTRA_CA_CERTS pointing at it. Node reads that variable only
at startup, hence the re-exec.
Older Node on Linux: it points at the distribution’s CA bundle the same
way.
One more trap: a background service doesn’t inherit your shell’s environment.
Proxy and certificate variables that work in a terminal are simply absent under
launchd. The agent installer copies them into the service definition.
claude-obs install-agent runs sync --watch as a launchd agent on macOS, a
systemd user unit on Linux, or a Task Scheduler job on Windows, at low priority.
A PID lock file keeps it to one instance. The lock is released on exit and on
SIGINT, SIGTERM and SIGHUP, and a lock left by a dead process is detected and
replaced. Running claude-obs sync by hand while the agent is up just nudges the
running instance.
Idempotent ingest from two sources
Session files are not the whole truth. Claude Code also makes small side
requests, such as generating session titles and running classifiers, that never
reach the transcript. That’s 5–10% of usage.
Claude C

[truncated]

## Original Extract

Building an Observability Dashboard for a Claude Subscription
Building an Observability Dashboard for a Claude Subscription
The problem What the answers look like
The collector Session files are undocumented, so parse an allowlist
Write to disk before touching the network
Privacy is a test, not a promise
Corporate networks break naive HTTPS
The server Idempotent ingest from two sources
Pinning machines to one account
Reconstructing the 5-hour blocks
How it was built Plan first, and write down the decisions
Measure the data before writing the parser
Verify against something independent
Where does a Claude subscription’s usage actually go?
A Claude Max subscription gives you a lot of Claude Code. It does not tell you
much about how you used it. /usage shows a percentage of the current window
and when it resets. That’s the whole report: no history, no sessions or models,
and nothing across machines.
API customers get an Admin API with usage and cost reports. Subscribers get
nothing like it. Meanwhile the data that would answer the question sits in
plain sight, in Claude Code’s session files, spread across every laptop,
desktop and remote dev box you use it on.
So I built claude-observability :
a small collector, claude-obs , that
runs on each machine, and a Next.js dashboard that turns what the collectors
upload into account-level analytics. It’s about 2,000 lines of TypeScript for
the collector (no runtime dependencies) and 6,500 for the app, plus a small
shared package that both depend on.
The interesting parts weren’t the charts. They were:
getting a number you can trust out of files that were never meant to be read
keeping prompts and code on the machine
running a background process on computers you don’t control
Source: github.com/FirePing32/claude-observability , MIT licensed
Collector: claude-obs on npm , Node 20+
Deploy the dashboard yourself (the README walks through it), sign in, create a
workspace for your Claude account, and generate an enrollment code. Then, on
every machine that runs Claude Code with that account:
npm install -g claude-obs
claude-obs login --code XXXX-XXXX --server https://<your-dashboard>
claude-obs sync # uploads your full history the first time
claude-obs install-agent # keeps it up to date in the background
Want to see exactly what leaves the machine before enrolling?
claude-obs sync --dry-run prints it and sends nothing.
The questions are simple to ask:
Am I about to hit the 5-hour limit? And how close did I get yesterday?
What is using it? Which sessions, which projects, which models? Is Opus
quietly doing work Sonnet could do?
Is the plan worth it? What would the same tokens have cost on the
pay-as-you-go API?
Is anything missing? Has a machine stopped reporting? Did we hit a limit
that the tracked usage doesn’t explain?
The constraints are what make it a project:
There is no API to pull from. The only source is local files, on each
machine, in a format Anthropic doesn’t document and changes at will.
The account is the unit. Limits apply to the account, not to a machine.
A dashboard that shows one laptop’s usage answers the wrong question.
Nothing sensitive leaves the machine. The session files hold every
prompt, every response, every file Claude read. The collector must upload
numbers and nothing else, and that has to be provable rather than promised.
No credentials. The collector must never touch the Claude login token.
That’s the overview for my account over the last 30 days. Each question above
gets a tile:
The 5-hour limit. The current block has used $40.67 of value with two
hours to go. That’s 25% of this account’s learned limit (about $165, from
three real limit hits). At the current burn rate of $16.80 an hour it will
reach about 45% by the reset.
Is the plan worth it? $2,289 of API-equivalent usage on a $100 plan this
cycle: 22.9×. It broke even the day after renewal.
What is using it? Opus accounts for 98% of the value. The daily chart also
shows the switch from Opus 5 to Opus 5.5 in late September, without anyone
having to remember when it happened.
Is anything missing? All five machines are reporting. But telemetry
coverage is 0%, so the totals are a few percent low, and claude.ai chat isn’t
measured at all. The page says so instead of hiding it.
Below the tiles, the first insight was the one that started the project: hit a
usage limit in 2 blocks this week, vs 1 the week before .
It’s a push model. Each machine runs claude-obs , which reads the bytes
Claude Code appended to its session files since the last run. It reduces them
to usage records and uploads them. The server merges uploads from every machine
into one account, prices each request, and serves the dashboard.
A workspace is one Claude account. Machines enroll into a workspace, and
everything shown is account-wide. There is deliberately no per-person view: the
goal is the right totals for the account, not a leaderboard.
Machines are the one breakdown the dashboard does offer, because a silent or
misconfigured collector is the most likely reason for wrong totals. For each
machine (names anonymized here), it shows:
when it was last active and last uploaded
A request reported by two machines counts once, under the first one.
The repo is an npm workspace with three packages:
packages/shared : the wire schema (zod) and the price book . The
collector validates what it sends with the same schema the server validates
what it receives, so the two can’t drift.
packages/cli : the collector, bundled into a single file with tsup .
apps/web : the Next.js 16 app. It holds the ingest endpoints, the dashboard,
the alerting and the database migrations.
Session files are undocumented, so parse an allowlist
Claude Code writes one JSON object per line to
~/.claude/projects/<project>/<session>.jsonl , plus a subagents/ folder per
session for subagent runs. There are a couple of dozen record types: assistant
messages, titles, compaction markers, turn durations, Claude Code’s own cost
snapshots, permission-mode changes and more. New ones appear between releases.
The parser dispatches on type and handles only the records it understands. It
builds every output record field by field from explicit paths : no object
spreading, no copying of message . Anything unknown is counted, not parsed. The
dashboard shows those counts under Data quality , so a format change shows up as
a number going up rather than as silently wrong totals.
The first real surprise came from counting before writing any parser logic.
Claude Code writes one line per content block of a response, and repeats
the response’s usage object on each. A reply with a text block and a tool call
is two lines with identical token counts. Sum them and you double-count. On my
own history, 3,301 assistant lines turned out to be 1,889 distinct requests.
Naive summing overcounted by about 75%.
It’s subtler than duplication. For 39 of those requests, output_tokens
differed between lines, and the last line was always the largest. Those are
streaming updates. So the rule is: group by requestId and keep the max,
never sum . The collector re-sends a request only if its output grew, which
keeps a live session accurate while a response is still streaming.
The server applies the same rule, so an out-of-order batch can’t lower a value:
// never let a concurrent, staler write lower the final streamed output
output: sql `greatest( ${ schema.apiRequests.output } , excluded.output)` ,
Other records need their own handling. Lines with model: "<synthetic>" are
not requests at all. They’re API errors and usage-limit notices that Claude Code
writes into the transcript. The collector classifies them locally into a
category ( prompt_too_long , rate_limit , a five-hour or weekly limit and so
on) and uploads only the category, never the text.
Session files grow to tens of megabytes. Re-reading them on every change would
be wasteful, so the collector keeps a byte offset per file, plus its device and
inode numbers:
It reads from the offset to EOF and consumes complete lines only . A
trailing line without a newline is still being written, so it waits for the
next pass.
It decodes each line from bytes on its own, so a multi-byte character split
across a read chunk is never mangled.
If the inode changed or the file shrank, the file was replaced, so it reads
from zero. That’s safe because the server deduplicates.
In watch mode it uses recursive fs.watch with a debounce, plus a full rescan
every 15 minutes as a safety net for missed events.
Write to disk before touching the network
The order of operations is the whole reliability story:
Write each batch to an outbox directory, one file per batch, via temp
file and rename.
Commit the new offsets to state.json , the same way.
Upload the outbox, in order, deleting each file only after the server
accepts it.
A crash between steps 2 and 3 means the lines are read again and queued twice.
The server stores each request once, so that’s harmless. A crash anywhere else
leaves the batch in the outbox for the next run. A laptop that’s offline for a
week simply builds up a queue. Every failure is a re-send, never a loss.
Upload failures are classified:
Network errors, 5xx and 429 back off and retry, honouring Retry-After .
401 (token revoked) stops the agent and asks for a new login.
422 (schema rejected) moves that one batch to a quarantine folder with
the reason, so a single bad record can’t block the queue forever.
Privacy is a test, not a promise
“We only upload usage” is easy to write in a README. The repo enforces it
three ways:
Field-by-field builders , as above. There’s no code path that copies an
object from the transcript into the output.
A canary test. The fixtures plant marker strings in message content,
e-mail addresses and file paths. The test fails if any of them appear anywhere
in the serialized upload.
claude-obs sync --dry-run prints exactly what would be sent, so anyone
can check for themselves before enrolling.
What does go up: model, timestamps, token counts (with the five-minute and
one-hour cache writes separated), effort, entry point, subagent type, the
project folder’s basename (or a salted hash of it), git branch, and session
titles (which can be switched off).
Corporate networks break naive HTTPS
The first machine outside my own failed immediately:
SELF_SIGNED_CERT_IN_CHAIN . Security agents like Netskope and Zscaler intercept
HTTPS and re-sign it with a company root certificate. The OS trusts that root.
Node, by default, doesn’t: it ships its own CA bundle.
Telling every user to set NODE_EXTRA_CA_CERTS would be a support burden. So
the collector trusts the operating system’s store automatically, the way a
browser does:
Newer Node (22.19+, 24.5+): tls.setDefaultCACertificates merges the
system certificates in-process.
Older Node on macOS: it exports the System keychain and Apple’s roots
with /usr/bin/security , caches the bundle for a week, and re-launches itself
once with NODE_EXTRA_CA_CERTS pointing at it. Node reads that variable only
at startup, hence the re-exec.
Older Node on Linux: it points at the distribution’s CA bundle the same
way.
One more trap: a background service doesn’t inherit your shell’s environment.
Proxy and certificate variables that work in a terminal are simply absent under
launchd. The agent installer copies them into the service definition.
claude-obs install-agent runs sync --watch as a launchd agent on macOS, a
systemd user unit on Linux, or a Task Scheduler job on Windows, at low priority.
A PID lock file keeps it to one instance. The lock is released on exit and on
SIGINT, SIGTERM and SIGHUP, and a lock left by a dead process is detected and
replaced. Running claude-obs sync by hand while the agent is up just nudges the
running instance.
Idempotent ingest from two sources
Session files are not the whole truth. Claude Code also makes small side
requests, such as generating session titles and running classifiers, that never
reach the transcript. That’s 5–10% of usage.
Claude C

[truncated]
