---
source: "https://github.com/rbartoli/agent-usage-guard"
hn_url: "https://news.ycombinator.com/item?id=50198936"
title: "Show HN: agent-usage-guard – a Claude Code mod to protect your usage window"
article_title: "GitHub - rbartoli/agent-usage-guard: A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. · GitHub"
image: "https://opengraph.githubassets.com/b7aff4c7138d3f737a2ee3a1483c3fea21b8d7106218cef6c83f94ef1e7d7ffb/rbartoli/agent-usage-guard"
author: "zzyyxx332211"
captured_at: "2026-10-22T18:20:39Z"
capture_tool: "hn-digest"
hn_id: 50198936
score: 4
comments: 2
posted_at: "2026-10-22T18:02:38Z"
tags:
  - hacker-news
---

# Show HN: agent-usage-guard – a Claude Code mod to protect your usage window

- HN: [50198936](https://news.ycombinator.com/item?id=50198936)
- Source: [github.com](https://github.com/rbartoli/agent-usage-guard)
- Score: 4
- Comments: 2
- Posted: 2026-10-22T18:02:38Z

## Translation

Title: Show HN: agent-usage-guard – a Claude Code mod to protect your usage window
Article title: GitHub - rbartoli/agent-usage-guard: A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. · GitHub
Description: A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. - rbartoli/agent-usage-guard

Article text:
GitHub - rbartoli/agent-usage-guard: A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. · GitHub
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
rbartoli
agent-usage-guard Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
44 Commits 44 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows core core demo demo docs docs hooks hooks tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md TESTING.md TESTING.md tsconfig.json tsconfig.json View all files Repository files navigation
Stops Claude Code from spending your five-hour usage window in ten minutes.
agent-usage-guard is a Claude Code mod : it runs inside Claude Code and holds runaway spend before it happens. It reads your plan's live usage, and it holds subagent fan-out, heavy-context churn and retry loops.
Recorded from Claude Code's interactive UI with the guard loaded. A scripted local API stands in for Claude, so the sessions and token counts are staged and no usage is spent. Every question on screen is the guard's own. Rebuild it from the Northstar fixture .
It acts before the spend, not after it: before an agent starts, a prompt is sent, or a tool call runs. Agent limits hold across every Claude Code session on your machine, and fan-out and heavy context are capped at any usage level, not only near the limit. When you are at the keyboard it asks you, in Claude Code's own question dialog. When nobody can answer, it refuses and tells Claude why, so the model can wait, batch the work, or stop. /agent-guard report shows what it did.
Plan-aware. It reads the 5-hour and weekly percentages Claude Code reports, and keeps the end of the window for you.
One answer covers a while. An approval lifts its condition until it clears or times out, so the same question does not come back call after call.
Local and privacy-minimal. No network, no processes. It stores counts and coarse buckets, never prompt text, tool inputs or model output.
Fails open. If one of its hooks fails, the event goes ahead: a broken guard never breaks your session. The one exception is an agent it could not judge in time, which is refused so that Claude starts it again.
Run these commands inside Claude Code v2.1.287 or later:
/plugin marketplace add rbartoli/agent-usage-guard
/plugin install agent-usage-guard@agent-usage-guard
/reload-plugins
Then run /agent-guard to see what it sees. Mods are on by default; the overview lists the settings that turn them off.
One afternoon I gave Claude Code a research prompt at max effort. It spawned 394 subagent calls, 341 of them running at once, and pushed 84M context tokens through the API in ten minutes. Session over: "Claude usage limit reached." Locked out until the window reset.
That wasn't a one-off. Sixty days of my own logs held 12 lockouts and 166 ten-minute windows with agent bursts. The failure mode multiplies: a large context × rapid tool loops × parallel agents × blind retries . Each protection below removes one factor.
And it's not just me: 339 subagents from a single prompt , half a weekly Max plan consumed by recursive spawning , and multi-agent runs using ~15× the tokens of a chat session .
Area
Rule
Default
What happens
Plan windows
New agents need approval once a window is this full
80%
asks; approval lasts until that window resets
New agents are refused once a window is this full
95%
refuses
New agents need approval when the 5-hour window rises this fast
20 points / 10 min
asks
Agents
Agents running at once, across every session on this machine
4
asks
Agents started or resumed per 10 minutes, across this machine
12
asks
Tokens processed by subagents per 10 minutes, across this machine
10M
asks
Subagents may not start their own agents
depth 1
refuses
Context
Claude is told once to keep the turn bounded
300k tokens
adds a note to the prompt
A prompt that would re-read this much context
500k tokens
asks: compact then send, send anyway, or cancel
Tool calls per 10 minutes while a conversation or agent carries this much
20 calls at 400k
asks: keep going, compact after this turn, or stop
A heavy session idle long enough for its prompt cache to expire
150k tokens, 1 h idle
asks before re-writing the cache
Loops
The same tool call failing again and again, with no other call in between
3 identical failures
refuses the identical retry
Refused agent calls within 10 minutes
5
pauses agent spawns for 10 min
Agents covers subagents, agent-team teammates, workflow agents, and finished subagents resumed with SendMessage . Each tool call is judged on the context of the loop that makes it, so a small subagent never inherits the main conversation's size.
When it asks. The question appears in Claude Code's own dialog, the one Claude uses to ask you something, and the call waits for your answer. Agents started in parallel while it is open are held back, and Claude starts them again once you answer. If you allow it, that condition stays allowed for a while: a plan window until it resets, a heavy context until it shrinks (for tool calls, an hour at most), the agent counts for 10 minutes. If you decline or dismiss the question, Claude is told not to start agents for the next 10 minutes ( AGENT_GUARD_WINDOW_SECONDS ), and is not asked again in that time. Typed words reach Claude, so you can answer "use haiku agents instead".
When nobody can answer. In claude -p , the Agent SDK and dontAsk mode, an agent or tool gate refuses instead of asking. A heavy prompt is held with the reason, and so is a dormant heavy resume once the machine-wide cap of one per 10 minutes is reached. Claude reads each refusal as the tool's result. Repeat refusals of the same condition are worded differently each time and escalate: the second says nothing has changed and to try something else, the third says to end the turn. If a Stop hook such as /goal keeps reopening the turn, the refusal says the user has to decide.
What it leaves alone. It never holds a background task's notification or another session's message, because that would lose the result. It no longer blocks prompts after a usage limit: since v2.1.234, Claude Code waits at a limit and continues at reset on its own ( autoContinueAtUsageLimit ). The guard's job is to keep you from reaching the limit, and its journal counts the lockouts it did not prevent.
/agent-guard [status] what the guard sees right now
/agent-guard allow [agents|context] [minutes] lift its limits for this session (default: all, 10 min)
/agent-guard pause [minutes] the same as allow, for every limit
/agent-guard resume end an allow early, lift the agent pause, and let declined questions ask again
/agent-guard report [days] what it did on this machine (default: 7 days)
/agent-guard help these commands
/agent-guard runs at once, even mid-turn, and never starts a model turn, so Claude never sees an override and cannot take one as permission. Sending [allow-usage-guard] or [allow-agent-burst] alone as a prompt, the markers of earlier versions, does the same as /agent-guard allow or allow agents .
Checked against Claude Code v2.1.295:
Mods run in claude in a terminal, in the Desktop app's Code tab, in claude -p and the Agent SDK. They don't run in a Desktop-app WSL session ( where mods run ).
Native limits still matter. Claude Code caps a session at 20 concurrent subagents ( CLAUDE_CODE_MAX_CONCURRENT_SUBAGENTS ) and three layers of nesting ( CLAUDE_CODE_MAX_SUBAGENT_SPAWN_DEPTH ). It has no limit on the total number of agents a session starts. The guard's defaults are stricter and span sessions; the native caps remain the outer wall for the paths the guard only counts.
/agent-guard report 30
The report counts usage-limit lockouts (each locked window once), plan-window threshold crossings, and every question and refusal by rule, with what you answered. It reads every session's journal on this machine and never leaves it. Set AGENT_GUARD_JOURNAL=0 to stop recording.
A mod runs inside Claude Code with your permissions; read what a mod can reach before you install one. This one is about 2,400 lines of TypeScript with no dependencies, and claude plugin validate . lists every API it calls and every environment variable it reads:
It makes no network requests and starts no processes.
It decides only what the rules above describe: whether an agent starts, a prompt is sent, or a tool call runs. It asks you whenever you can answer, and gives the reason when it refuses. It reads permission modes but never changes them, and it writes no settings.
The only prompt it submits is your own: after "Compact, then send", it sends the prompt you typed, unchanged, once compaction finishes.
Its prompt hook passes your prompt on unchanged, adds a one-time note for Claude when the context passes 300k, or holds the prompt to ask you first (when nobody can answer, it says why instead). Its SessionStart hook changes nothing.
It keeps its state in the mod store, a JSON file under ~/.claude/plugins/store/ . Each session writes its own counts (running agents, recent starts and dormant resumes, subagent tokens per minute, its latest 5-hour reading), when it was last active (kept 30 days, so a resumed session knows how long it sat idle), and its journal rows (time, a session-id prefix, rule, tool name, context bucket such as 400-500k , plan window and percentage, refusal number, and your answer).
It never stores prompt text, tool inputs, error text or model output. A tool call is reduced to a 13-character fingerprint. Agent names stay in memory and are never written.
Journal days older than 90 days are deleted ( AGENT_GUARD_JOURNAL_DAYS ).
demo/ , which builds the GIF above, is not part of what runs. It starts Claude Code in a temporary home directory with a made-up API key for a scripted local API, and reads none of your credentials.
Set AGENT_GUARD=0 to turn it off without uninstalling it.
Set these as environment variables, or under env in ~/.claude/settings.json . Times are in seconds; a value that isn't a number in range keeps its default, and /agent-guard says so.
The dormant default assumes the one-hour prompt cache that subscriptions get within plan usage. With an API key or ext

[truncated]

## Original Extract

A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. - rbartoli/agent-usage-guard

GitHub - rbartoli/agent-usage-guard: A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog. · GitHub
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
rbartoli
agent-usage-guard Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
A Claude Code mod that holds subagent fan-out, heavy-context prompts and retry loops before they spend your usage window. It reads your plan's usage, counts agents across every session, and asks in Claude Code's own dialog.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
44 Commits 44 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows core core demo demo docs docs hooks hooks tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md TESTING.md TESTING.md tsconfig.json tsconfig.json View all files Repository files navigation
Stops Claude Code from spending your five-hour usage window in ten minutes.
agent-usage-guard is a Claude Code mod : it runs inside Claude Code and holds runaway spend before it happens. It reads your plan's live usage, and it holds subagent fan-out, heavy-context churn and retry loops.
Recorded from Claude Code's interactive UI with the guard loaded. A scripted local API stands in for Claude, so the sessions and token counts are staged and no usage is spent. Every question on screen is the guard's own. Rebuild it from the Northstar fixture .
It acts before the spend, not after it: before an agent starts, a prompt is sent, or a tool call runs. Agent limits hold across every Claude Code session on your machine, and fan-out and heavy context are capped at any usage level, not only near the limit. When you are at the keyboard it asks you, in Claude Code's own question dialog. When nobody can answer, it refuses and tells Claude why, so the model can wait, batch the work, or stop. /agent-guard report shows what it did.
Plan-aware. It reads the 5-hour and weekly percentages Claude Code reports, and keeps the end of the window for you.
One answer covers a while. An approval lifts its condition until it clears or times out, so the same question does not come back call after call.
Local and privacy-minimal. No network, no processes. It stores counts and coarse buckets, never prompt text, tool inputs or model output.
Fails open. If one of its hooks fails, the event goes ahead: a broken guard never breaks your session. The one exception is an agent it could not judge in time, which is refused so that Claude starts it again.
Run these commands inside Claude Code v2.1.287 or later:
/plugin marketplace add rbartoli/agent-usage-guard
/plugin install agent-usage-guard@agent-usage-guard
/reload-plugins
Then run /agent-guard to see what it sees. Mods are on by default; the overview lists the settings that turn them off.
One afternoon I gave Claude Code a research prompt at max effort. It spawned 394 subagent calls, 341 of them running at once, and pushed 84M context tokens through the API in ten minutes. Session over: "Claude usage limit reached." Locked out until the window reset.
That wasn't a one-off. Sixty days of my own logs held 12 lockouts and 166 ten-minute windows with agent bursts. The failure mode multiplies: a large context × rapid tool loops × parallel agents × blind retries . Each protection below removes one factor.
And it's not just me: 339 subagents from a single prompt , half a weekly Max plan consumed by recursive spawning , and multi-agent runs using ~15× the tokens of a chat session .
Area
Rule
Default
What happens
Plan windows
New agents need approval once a window is this full
80%
asks; approval lasts until that window resets
New agents are refused once a window is this full
95%
refuses
New agents need approval when the 5-hour window rises this fast
20 points / 10 min
asks
Agents
Agents running at once, across every session on this machine
4
asks
Agents started or resumed per 10 minutes, across this machine
12
asks
Tokens processed by subagents per 10 minutes, across this machine
10M
asks
Subagents may not start their own agents
depth 1
refuses
Context
Claude is told once to keep the turn bounded
300k tokens
adds a note to the prompt
A prompt that would re-read this much context
500k tokens
asks: compact then send, send anyway, or cancel
Tool calls per 10 minutes while a conversation or agent carries this much
20 calls at 400k
asks: keep going, compact after this turn, or stop
A heavy session idle long enough for its prompt cache to expire
150k tokens, 1 h idle
asks before re-writing the cache
Loops
The same tool call failing again and again, with no other call in between
3 identical failures
refuses the identical retry
Refused agent calls within 10 minutes
5
pauses agent spawns for 10 min
Agents covers subagents, agent-team teammates, workflow agents, and finished subagents resumed with SendMessage . Each tool call is judged on the context of the loop that makes it, so a small subagent never inherits the main conversation's size.
When it asks. The question appears in Claude Code's own dialog, the one Claude uses to ask you something, and the call waits for your answer. Agents started in parallel while it is open are held back, and Claude starts them again once you answer. If you allow it, that condition stays allowed for a while: a plan window until it resets, a heavy context until it shrinks (for tool calls, an hour at most), the agent counts for 10 minutes. If you decline or dismiss the question, Claude is told not to start agents for the next 10 minutes ( AGENT_GUARD_WINDOW_SECONDS ), and is not asked again in that time. Typed words reach Claude, so you can answer "use haiku agents instead".
When nobody can answer. In claude -p , the Agent SDK and dontAsk mode, an agent or tool gate refuses instead of asking. A heavy prompt is held with the reason, and so is a dormant heavy resume once the machine-wide cap of one per 10 minutes is reached. Claude reads each refusal as the tool's result. Repeat refusals of the same condition are worded differently each time and escalate: the second says nothing has changed and to try something else, the third says to end the turn. If a Stop hook such as /goal keeps reopening the turn, the refusal says the user has to decide.
What it leaves alone. It never holds a background task's notification or another session's message, because that would lose the result. It no longer blocks prompts after a usage limit: since v2.1.234, Claude Code waits at a limit and continues at reset on its own ( autoContinueAtUsageLimit ). The guard's job is to keep you from reaching the limit, and its journal counts the lockouts it did not prevent.
/agent-guard [status] what the guard sees right now
/agent-guard allow [agents|context] [minutes] lift its limits for this session (default: all, 10 min)
/agent-guard pause [minutes] the same as allow, for every limit
/agent-guard resume end an allow early, lift the agent pause, and let declined questions ask again
/agent-guard report [days] what it did on this machine (default: 7 days)
/agent-guard help these commands
/agent-guard runs at once, even mid-turn, and never starts a model turn, so Claude never sees an override and cannot take one as permission. Sending [allow-usage-guard] or [allow-agent-burst] alone as a prompt, the markers of earlier versions, does the same as /agent-guard allow or allow agents .
Checked against Claude Code v2.1.295:
Mods run in claude in a terminal, in the Desktop app's Code tab, in claude -p and the Agent SDK. They don't run in a Desktop-app WSL session ( where mods run ).
Native limits still matter. Claude Code caps a session at 20 concurrent subagents ( CLAUDE_CODE_MAX_CONCURRENT_SUBAGENTS ) and three layers of nesting ( CLAUDE_CODE_MAX_SUBAGENT_SPAWN_DEPTH ). It has no limit on the total number of agents a session starts. The guard's defaults are stricter and span sessions; the native caps remain the outer wall for the paths the guard only counts.
/agent-guard report 30
The report counts usage-limit lockouts (each locked window once), plan-window threshold crossings, and every question and refusal by rule, with what you answered. It reads every session's journal on this machine and never leaves it. Set AGENT_GUARD_JOURNAL=0 to stop recording.
A mod runs inside Claude Code with your permissions; read what a mod can reach before you install one. This one is about 2,400 lines of TypeScript with no dependencies, and claude plugin validate . lists every API it calls and every environment variable it reads:
It makes no network requests and starts no processes.
It decides only what the rules above describe: whether an agent starts, a prompt is sent, or a tool call runs. It asks you whenever you can answer, and gives the reason when it refuses. It reads permission modes but never changes them, and it writes no settings.
The only prompt it submits is your own: after "Compact, then send", it sends the prompt you typed, unchanged, once compaction finishes.
Its prompt hook passes your prompt on unchanged, adds a one-time note for Claude when the context passes 300k, or holds the prompt to ask you first (when nobody can answer, it says why instead). Its SessionStart hook changes nothing.
It keeps its state in the mod store, a JSON file under ~/.claude/plugins/store/ . Each session writes its own counts (running agents, recent starts and dormant resumes, subagent tokens per minute, its latest 5-hour reading), when it was last active (kept 30 days, so a resumed session knows how long it sat idle), and its journal rows (time, a session-id prefix, rule, tool name, context bucket such as 400-500k , plan window and percentage, refusal number, and your answer).
It never stores prompt text, tool inputs, error text or model output. A tool call is reduced to a 13-character fingerprint. Agent names stay in memory and are never written.
Journal days older than 90 days are deleted ( AGENT_GUARD_JOURNAL_DAYS ).
demo/ , which builds the GIF above, is not part of what runs. It starts Claude Code in a temporary home directory with a made-up API key for a scripted local API, and reads none of your credentials.
Set AGENT_GUARD=0 to turn it off without uninstalling it.
Set these as environment variables, or under env in ~/.claude/settings.json . Times are in seconds; a value that isn't a number in range keeps its default, and /agent-guard says so.
The dormant default assumes the one-hour prompt cache that subscriptions get within plan usage. With an API key or ext

[truncated]
