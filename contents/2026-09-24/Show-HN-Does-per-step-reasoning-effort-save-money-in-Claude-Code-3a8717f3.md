---
source: "https://github.com/ifoster01/jev-effort"
hn_url: "https://news.ycombinator.com/item?id=49824915"
title: "Show HN: Does per-step reasoning effort save money in Claude Code?"
article_title: "GitHub - ifoster01/jev-effort: Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. · GitHub"
image: "https://opengraph.githubassets.com/a4aaecdb229d7cf76e705963c6eba27063a6a1bc1bedccb5126195c3910c1782/ifoster01/jev-effort"
author: "ifoster41901"
captured_at: "2026-09-24T01:08:48Z"
capture_tool: "hn-digest"
hn_id: 49824915
score: 1
comments: 1
posted_at: "2026-09-24T01:08:08Z"
tags:
  - hacker-news
---

# Show HN: Does per-step reasoning effort save money in Claude Code?

- HN: [49824915](https://news.ycombinator.com/item?id=49824915)
- Source: [github.com](https://github.com/ifoster01/jev-effort)
- Score: 1
- Comments: 1
- Posted: 2026-09-24T01:08:08Z

## Translation

Title: Show HN: Does per-step reasoning effort save money in Claude Code?
Article title: GitHub - ifoster01/jev-effort: Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. · GitHub
Description: Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. - ifoster01/jev-effort

Article text:
GitHub - ifoster01/jev-effort: Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ifoster01
/
jev-effort
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github .github bench/ tasks bench/ tasks bin bin docs docs src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md package.json package.json View all files Repository files navigation
Does letting Jev choose Claude Code's reasoning effort, step by step, save money? At high
effort, barely: about 1%. At max effort, it cut the cost of a coding benchmark by 55% with every
test still passing. In a long real session at max , the estimated saving was 4–9%, because
re-reading the conversation's context is most of the bill. Changing effort on every step never
reset the prompt cache.
A research project. The code works and the results are reproducible. Unofficial; not affiliated
with Anthropic or TypeSafe.
Jev is TypeSafe's small, cheap decision model. After
Astra-Ares used it to choose GPT-6 Astra's reasoning
effort per step in Codex, several tools appeared doing the same for Claude. The pitch: most agent
steps are routine, so run them at low effort and keep deep reasoning for the hard ones.
I built a proxy that does this for Claude Code without breaking the prompt cache, confirmed it
works, then measured what it saves.
Effort changes how much the model thinks, and the saving follows thinking's share of the bill.
At high , Claude Opus 5.5 thinks little on routine steps, so there's little to cut. At max , it
thinks a lot even on routine steps, and Jev moves those down. But in a long session every step
re-reads hundreds of thousands of tokens of context, which effort doesn't touch.
Changing effort doesn't reset the cache
Claude Code caches the conversation so each step only pays full price for what's new. Changing
effort mid-session used to throw that cache away. It no longer has to: Anthropic's
per-message effort
beta lets an effort-only system message change the level while leaving the cached prefix intact.
It held inside Claude Code. In the 470-step session with effort chosen on every step: 0
unexpected cache misses in 411 checked steps, and a 99.1% cache hit rate. With effort
alternating low/high on every step of a test task, each request read the whole previous prefix
from cache (26,992 → 34,285 → 35,017 → 35,270 → 35,617 tokens).
The catch: Claude Code already carries its own effort setting on a system message near the
start of the conversation (beta per-turn-control-2026-07-01 ), and the latest setting in a
request wins. An injected effort change has
to be the last message, or it's silently ignored. Changing the request's top-level effort field is
ignored the same way.
Claude Code runs a whole session at one effort level ( /effort : low, medium, high, xhigh,
max).
Opus 5.5, Fable 5.1, Mythos 5.1, and Opus 5 accept a
per-message effort change :
an effort-only system message inside messages that leaves the cached prefix intact. So
effort can change on every step without re-processing the conversation.
Research on per-step effort selection ( ARES ,
TAB ) reports token savings of 35–53%. Both count tokens,
not dollars. The open question was what per-step effort does to the total cost of real Claude
Code sessions.
1. Mechanism. jev-effort runs claude behind a local proxy ( ANTHROPIC_BASE_URL ). Before
each model request it sends Jev a trimmed view of the conversation (prompts, Claude's visible
replies, the last six tool calls) and asks two questions: which effort the next step needs, and
for how many steps to keep it. It applies the answer by appending an effort-only system message,
and re-inserts earlier ones at their original positions so each request begins with the previous
one. Jev may lower effort but never exceed the session's own setting.
2. Verification. On a hard counting problem, a low marker produced 438–548 output tokens
and max produced 1,528–1,809, in line with Claude Code's own --effort (438 and 1,733). With
effort alternating on every step, each request read the full previous prefix from cache. In the
real session: 0 unexpected cache misses in 411 checked steps, 99.1% cache hit rate.
3. Controlled benchmark. Six small coding tasks, each with a visible test and hidden checks.
Each task ran from identical files twice: once at a fixed effort, once with Jev choosing under
that same effort as its ceiling. Two rounds at high and two at max , 48 headless Claude Code
sessions in all. A warm-up run came first so neither arm paid to cache the other's system prompt.
4. Shadow mode on real work. One 92-minute session of my normal work at max effort (470
model steps, Opus 5.5, Claude Code 2.1.280). For every step, Jev's choice was logged and nothing
was changed. Each request was priced at Claude API list prices from the usage the API reported:
cache reads, cache writes, output, and hidden thinking.
Task
Pass (fixed / Jev)
Cost (fixed / Jev)
Thinking tokens (fixed / Jev)
Jev's effort mix
expr-eval
2/2 / 2/2
$2.00 / $0.32
54,397 / 551
low 5, medium 4
interval-merge
2/2 / 2/2
$0.36 / $0.25
2,907 / 45
low 5, medium 4
lru-cache
2/2 / 2/2
$0.39 / $0.23
2,963 / 0
low 3, medium 4
paginate-bug
2/2 / 2/2
$0.35 / $0.28
2,788 / 45
low 4, medium 5
rename-refactor
2/2 / 2/2
$0.38 / $0.36
1,335 / 313
low 12, high 2
stats-bugs
2/2 / 2/2
$0.27 / $0.28
509 / 0
low 6, medium 4
Total
12/12 / 12/12
$3.76 / $1.70 (−55%)
64,899 / 954 (−98.5%)
Wall time fell from 913 s to 387 s. One task dominates: on expr-eval, max thought for about
27,000 tokens per run and Jev's low and medium settings passed the same hidden tests with about
280. Without expr-eval, cost still fell 21%.
Compared with just using high . The same six tasks at a fixed high cost less than fixed
max , so the fair question is whether Jev beats simply turning effort down. On these tasks it
matched or beat it. The runs were separate, so treat this as a rough comparison:
Fixed high
Jev
Change
Hidden checks passed
12/12
12/12
Thinking tokens
1,450
787
−45.7%
Output tokens
20,281
20,515
+1.2%
Cost
$1.94
$1.92
−1.1%
Wall time
393 s
417 s
+6%
Real session at max
Opus 5.5 list prices: $0.20 per million tokens read from cache, $8 per million written to the
one-hour cache, $20 per million output tokens. I'm on a subscription, so these are
API-equivalent figures, not a bill.
The session averaged 441K tokens of context per step, re-read every time.
Jev would have lowered effort on all 470 steps: to high on 261, xhigh on 149, low on 51, and
medium on 9. It never kept max .
The max benchmark's cut is likely too high for this session: there Jev mostly chose low and
medium, while here it mostly chose high and xhigh. The break-even is thin. One extra step
costs about $0.09 to re-read the context, so the net saving equals 30–67 extra steps across 470.
If lower effort makes Claude take more steps than that, it loses money. Shadow mode can't observe
that; the benchmarks found step counts roughly unchanged (92 vs 87 turns at max ).
The mechanism works. Effort changes take hold, the cache survives, and every hidden test
passed in both arms at both effort levels.
At high , it doesn't pay. Opus 5.5 at high spends little on thinking during routine
steps, so a 46% thinking cut moved total cost by about 1%.
At max , it can pay. max spends heavily on thinking even when the step doesn't need it.
With short contexts the saving is large (55% on the benchmark, and faster). In a long session
the same kind of cut is a few percent of the bill, because context dominates.
It isn't much better than turning effort down yourself. On the benchmark, Jev under a max
ceiling cost about the same as fixed high . What it adds is keeping max available for the
steps that need it, which small tasks can't show.
It adds latency, but may not cost time. Jev took 601 ms median and 793 ms at p95 per
decision, with no errors. At max , less thinking more than made up for it: the Jev arm finished
the benchmark in 42% of the time.
Quality at max on real work is untested. The benchmark tasks are small and well specified.
If you run max because your work needs it, lowering effort may cost more than it saves.
One real session, at max , 92 minutes. There's no real-session measurement at high .
The benchmark tasks are small, and two runs per task shows direction, not a precise effect
size. One task accounts for most of the max result.
Shadow mode can't measure quality or extra steps on real work, so the real-session saving is a
range, not a measurement.
Behind any proxy, Claude Code drops a few direct-connection features, including claude.ai-backed
tools such as Artifacts, so each request carried about 4K fewer tokens than a direct session
would ( details ).
Dollar figures apply API list prices to subscription usage. How plan limits weigh each token
type isn't published.
Notes on Claude Code internals
Found while building this; details and evidence in docs/how-it-works.md .
Claude Code re-sends some messages as a plain string after first sending them as text blocks.
The API caches both the same way, but anything that hashes the history must normalize them.
Any custom ANTHROPIC_BASE_URL makes Claude Code turn off MCP tool search, loading every MCP
tool definition into every request. With several MCP servers, a fresh session started at
281K–500K tokens instead of 27K. ENABLE_TOOL_SEARCH=true restores it when the proxy forwards
tool_reference blocks.
Claude Code 2.1.260+ has early-access "function hooks" ( CLAUDE_CODE_ENABLE_FUNCTION_HOOKS=1 ).
A turn.step hook can rewrite effort per request, and Claude Code then inserts the same
cache-safe markers itself.
npx jev-effort --jev-shadow # use like `claude`; logs Jev's choices, changes nothing
npx jev-effort stats # where your spend goes and what Jev would save
npx jev-effort stats --share # the same, safe to paste (no prompts or ids)
npx jev-effort bench --effort max --runs 2 # the controlled comparison (uses your Claude usage)
Setup, configuration, and privacy details: docs/usage.md . The data behind this
page: real session at max ,
benchmark at max , and
benchmark at high . If your numbers look different,
please open an issue with jev-effort stats --share .
Project
Approach
Astra-Ares
The original: Jev picks GPT-6 Astra's effort per step in Codex. jev-effort adapts its Jev prompt wording (MIT)
jev-opus
Same proxy approach for Claude Code, plus work-phase adjustments
jev-model-router
Cl

[truncated]

## Original Extract

Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. - ifoster01/jev-effort

GitHub - ifoster01/jev-effort: Per-step reasoning effort for Claude Code, chosen by Jev, without breaking the prompt cache. Unofficial. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ifoster01
/
jev-effort
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github .github bench/ tasks bench/ tasks bin bin docs docs src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md package.json package.json View all files Repository files navigation
Does letting Jev choose Claude Code's reasoning effort, step by step, save money? At high
effort, barely: about 1%. At max effort, it cut the cost of a coding benchmark by 55% with every
test still passing. In a long real session at max , the estimated saving was 4–9%, because
re-reading the conversation's context is most of the bill. Changing effort on every step never
reset the prompt cache.
A research project. The code works and the results are reproducible. Unofficial; not affiliated
with Anthropic or TypeSafe.
Jev is TypeSafe's small, cheap decision model. After
Astra-Ares used it to choose GPT-6 Astra's reasoning
effort per step in Codex, several tools appeared doing the same for Claude. The pitch: most agent
steps are routine, so run them at low effort and keep deep reasoning for the hard ones.
I built a proxy that does this for Claude Code without breaking the prompt cache, confirmed it
works, then measured what it saves.
Effort changes how much the model thinks, and the saving follows thinking's share of the bill.
At high , Claude Opus 5.5 thinks little on routine steps, so there's little to cut. At max , it
thinks a lot even on routine steps, and Jev moves those down. But in a long session every step
re-reads hundreds of thousands of tokens of context, which effort doesn't touch.
Changing effort doesn't reset the cache
Claude Code caches the conversation so each step only pays full price for what's new. Changing
effort mid-session used to throw that cache away. It no longer has to: Anthropic's
per-message effort
beta lets an effort-only system message change the level while leaving the cached prefix intact.
It held inside Claude Code. In the 470-step session with effort chosen on every step: 0
unexpected cache misses in 411 checked steps, and a 99.1% cache hit rate. With effort
alternating low/high on every step of a test task, each request read the whole previous prefix
from cache (26,992 → 34,285 → 35,017 → 35,270 → 35,617 tokens).
The catch: Claude Code already carries its own effort setting on a system message near the
start of the conversation (beta per-turn-control-2026-07-01 ), and the latest setting in a
request wins. An injected effort change has
to be the last message, or it's silently ignored. Changing the request's top-level effort field is
ignored the same way.
Claude Code runs a whole session at one effort level ( /effort : low, medium, high, xhigh,
max).
Opus 5.5, Fable 5.1, Mythos 5.1, and Opus 5 accept a
per-message effort change :
an effort-only system message inside messages that leaves the cached prefix intact. So
effort can change on every step without re-processing the conversation.
Research on per-step effort selection ( ARES ,
TAB ) reports token savings of 35–53%. Both count tokens,
not dollars. The open question was what per-step effort does to the total cost of real Claude
Code sessions.
1. Mechanism. jev-effort runs claude behind a local proxy ( ANTHROPIC_BASE_URL ). Before
each model request it sends Jev a trimmed view of the conversation (prompts, Claude's visible
replies, the last six tool calls) and asks two questions: which effort the next step needs, and
for how many steps to keep it. It applies the answer by appending an effort-only system message,
and re-inserts earlier ones at their original positions so each request begins with the previous
one. Jev may lower effort but never exceed the session's own setting.
2. Verification. On a hard counting problem, a low marker produced 438–548 output tokens
and max produced 1,528–1,809, in line with Claude Code's own --effort (438 and 1,733). With
effort alternating on every step, each request read the full previous prefix from cache. In the
real session: 0 unexpected cache misses in 411 checked steps, 99.1% cache hit rate.
3. Controlled benchmark. Six small coding tasks, each with a visible test and hidden checks.
Each task ran from identical files twice: once at a fixed effort, once with Jev choosing under
that same effort as its ceiling. Two rounds at high and two at max , 48 headless Claude Code
sessions in all. A warm-up run came first so neither arm paid to cache the other's system prompt.
4. Shadow mode on real work. One 92-minute session of my normal work at max effort (470
model steps, Opus 5.5, Claude Code 2.1.280). For every step, Jev's choice was logged and nothing
was changed. Each request was priced at Claude API list prices from the usage the API reported:
cache reads, cache writes, output, and hidden thinking.
Task
Pass (fixed / Jev)
Cost (fixed / Jev)
Thinking tokens (fixed / Jev)
Jev's effort mix
expr-eval
2/2 / 2/2
$2.00 / $0.32
54,397 / 551
low 5, medium 4
interval-merge
2/2 / 2/2
$0.36 / $0.25
2,907 / 45
low 5, medium 4
lru-cache
2/2 / 2/2
$0.39 / $0.23
2,963 / 0
low 3, medium 4
paginate-bug
2/2 / 2/2
$0.35 / $0.28
2,788 / 45
low 4, medium 5
rename-refactor
2/2 / 2/2
$0.38 / $0.36
1,335 / 313
low 12, high 2
stats-bugs
2/2 / 2/2
$0.27 / $0.28
509 / 0
low 6, medium 4
Total
12/12 / 12/12
$3.76 / $1.70 (−55%)
64,899 / 954 (−98.5%)
Wall time fell from 913 s to 387 s. One task dominates: on expr-eval, max thought for about
27,000 tokens per run and Jev's low and medium settings passed the same hidden tests with about
280. Without expr-eval, cost still fell 21%.
Compared with just using high . The same six tasks at a fixed high cost less than fixed
max , so the fair question is whether Jev beats simply turning effort down. On these tasks it
matched or beat it. The runs were separate, so treat this as a rough comparison:
Fixed high
Jev
Change
Hidden checks passed
12/12
12/12
Thinking tokens
1,450
787
−45.7%
Output tokens
20,281
20,515
+1.2%
Cost
$1.94
$1.92
−1.1%
Wall time
393 s
417 s
+6%
Real session at max
Opus 5.5 list prices: $0.20 per million tokens read from cache, $8 per million written to the
one-hour cache, $20 per million output tokens. I'm on a subscription, so these are
API-equivalent figures, not a bill.
The session averaged 441K tokens of context per step, re-read every time.
Jev would have lowered effort on all 470 steps: to high on 261, xhigh on 149, low on 51, and
medium on 9. It never kept max .
The max benchmark's cut is likely too high for this session: there Jev mostly chose low and
medium, while here it mostly chose high and xhigh. The break-even is thin. One extra step
costs about $0.09 to re-read the context, so the net saving equals 30–67 extra steps across 470.
If lower effort makes Claude take more steps than that, it loses money. Shadow mode can't observe
that; the benchmarks found step counts roughly unchanged (92 vs 87 turns at max ).
The mechanism works. Effort changes take hold, the cache survives, and every hidden test
passed in both arms at both effort levels.
At high , it doesn't pay. Opus 5.5 at high spends little on thinking during routine
steps, so a 46% thinking cut moved total cost by about 1%.
At max , it can pay. max spends heavily on thinking even when the step doesn't need it.
With short contexts the saving is large (55% on the benchmark, and faster). In a long session
the same kind of cut is a few percent of the bill, because context dominates.
It isn't much better than turning effort down yourself. On the benchmark, Jev under a max
ceiling cost about the same as fixed high . What it adds is keeping max available for the
steps that need it, which small tasks can't show.
It adds latency, but may not cost time. Jev took 601 ms median and 793 ms at p95 per
decision, with no errors. At max , less thinking more than made up for it: the Jev arm finished
the benchmark in 42% of the time.
Quality at max on real work is untested. The benchmark tasks are small and well specified.
If you run max because your work needs it, lowering effort may cost more than it saves.
One real session, at max , 92 minutes. There's no real-session measurement at high .
The benchmark tasks are small, and two runs per task shows direction, not a precise effect
size. One task accounts for most of the max result.
Shadow mode can't measure quality or extra steps on real work, so the real-session saving is a
range, not a measurement.
Behind any proxy, Claude Code drops a few direct-connection features, including claude.ai-backed
tools such as Artifacts, so each request carried about 4K fewer tokens than a direct session
would ( details ).
Dollar figures apply API list prices to subscription usage. How plan limits weigh each token
type isn't published.
Notes on Claude Code internals
Found while building this; details and evidence in docs/how-it-works.md .
Claude Code re-sends some messages as a plain string after first sending them as text blocks.
The API caches both the same way, but anything that hashes the history must normalize them.
Any custom ANTHROPIC_BASE_URL makes Claude Code turn off MCP tool search, loading every MCP
tool definition into every request. With several MCP servers, a fresh session started at
281K–500K tokens instead of 27K. ENABLE_TOOL_SEARCH=true restores it when the proxy forwards
tool_reference blocks.
Claude Code 2.1.260+ has early-access "function hooks" ( CLAUDE_CODE_ENABLE_FUNCTION_HOOKS=1 ).
A turn.step hook can rewrite effort per request, and Claude Code then inserts the same
cache-safe markers itself.
npx jev-effort --jev-shadow # use like `claude`; logs Jev's choices, changes nothing
npx jev-effort stats # where your spend goes and what Jev would save
npx jev-effort stats --share # the same, safe to paste (no prompts or ids)
npx jev-effort bench --effort max --runs 2 # the controlled comparison (uses your Claude usage)
Setup, configuration, and privacy details: docs/usage.md . The data behind this
page: real session at max ,
benchmark at max , and
benchmark at high . If your numbers look different,
please open an issue with jev-effort stats --share .
Project
Approach
Astra-Ares
The original: Jev picks GPT-6 Astra's effort per step in Codex. jev-effort adapts its Jev prompt wording (MIT)
jev-opus
Same proxy approach for Claude Code, plus work-phase adjustments
jev-model-router
Cl

[truncated]
