---
source: "https://github.com/data-gras/savi-loop-guard"
hn_url: "https://news.ycombinator.com/item?id=49704711"
title: "A zero-dependency Python loop detector for AI agents"
article_title: "GitHub - data-gras/savi-loop-guard: Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. · GitHub"
image: "https://opengraph.githubassets.com/c1fe4ec3fcbc492789ef06442563fa7ef1765fc476c63eb3884ebc0fa4a83cc4/data-gras/savi-loop-guard"
author: "GSingGras"
captured_at: "2026-09-14T22:16:57Z"
capture_tool: "hn-digest"
hn_id: 49704711
score: 1
comments: 0
posted_at: "2026-09-14T21:54:39Z"
tags:
  - hacker-news
---

# A zero-dependency Python loop detector for AI agents

- HN: [49704711](https://news.ycombinator.com/item?id=49704711)
- Source: [github.com](https://github.com/data-gras/savi-loop-guard)
- Score: 1
- Comments: 0
- Posted: 2026-09-14T21:54:39Z

## Translation

Title: A zero-dependency Python loop detector for AI agents
Article title: GitHub - data-gras/savi-loop-guard: Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. · GitHub
Description: Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. - data-gras/savi-loop-guard

Article text:
GitHub - data-gras/savi-loop-guard: Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. · GitHub
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
data-gras
/
savi-loop-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github loop_guard loop_guard tests tests .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml setup.py setup.py View all files Repository files navigation
A zero-dependency Python library that detects when an AI agent is stuck in a
loop: calling the same tool over and over, or firing calls far faster than
any real workflow would. Drop it into your own agent code. No account, no
API key, no network call, ever.
Built by SAVI as a standalone, dependency-free
package, so you can detect these patterns in your own agent code without
an account, an API key, or a dependency on SAVI's platform.
pip install savi-loop-guard
No dependencies. Nothing else gets installed alongside it.
Record each call as your agent makes it, then check for loops whenever you
want (after every call, on a timer, whatever fits your loop):
from loop_guard import LoopGuard , CallEvent
from datetime import datetime , timezone
guard = LoopGuard ()
guard . record ( CallEvent (
span_id = "call_1" ,
agent_id = "doc-extractor" ,
timestamp = datetime . now ( timezone . utc ),
tool_call = "search_web" ,
))
issues = guard . check ()
for issue in issues :
print ( issue [ "type" ], issue [ "agent_id" ])
An issue looks like:
{ "type" : LoopType . VELOCITY , "agent_id" : "doc-extractor" , "elapsed_s" : 12.4 , "call_count" : 6 }
# or
{ "type" : LoopType . STRUCTURAL , "agent_id" : "doc-extractor" ,
"tool_call" : "search_web" , "tool_variants" : [ "search_web" , "search_web_v2" ], "call_count" : 7 }
Prevent the call instead of just observing it
check_before_call() records the event and raises immediately if it would
trip a threshold, for callers who want to stop the loop rather than find
out about it afterwards:
from loop_guard import LoopGuard , CallEvent , LoopDetected
guard = LoopGuard ()
try :
guard . check_before_call ( event )
except LoopDetected as e :
print ( f"Blocked: { e . loop_type } " ) # "velocity_loop" or "structural_loop"
print ( e . details ) # the same dict check() would have returned
What it detects
Velocity loop : more than 5 calls from the same agent_id within a
30-second window.
Structural loop : the same tool called more than 5 times, with fuzzy
matching so a broken agent can't dodge detection by alternating between
near-identical tool names ( search_web vs search_web_v2 vs web_search
count as the same tool if their name tokens overlap enough).
All four numbers are configurable:
guard = LoopGuard (
velocity_window_seconds = 30 ,
velocity_call_limit = 5 ,
structural_call_limit = 5 ,
tool_fuzzy_similarity_threshold = 0.70 ,
)
How this compares
A few other standalone Python packages exist for this: agent-loop-detector ,
agent-loop-guard , agentguard-kit . All of them, like savi-loop-guard 's
record() / check() API, are post-hoc/observational; they analyze calls
after they happen. savi-loop-guard adds check_before_call() on top for
callers who want to prevent the call rather than just observe it, which none
of those currently offer.
LoopGuard keeps every recorded event in memory for the life of the
instance; nothing is ever pruned automatically. That's deliberate:
structural-loop detection is a total call count with no time bound by
design (a tool called 6 times over 3 hours is still a loop, not just a
tool called 6 times in 30 seconds), so silently dropping "old" events
would blind it to exactly the slow, steady loops it exists to catch.
In practice this means: for a short-lived task, a single LoopGuard()
is fine as-is. For a long-running process, create a fresh LoopGuard()
per logical unit of work (e.g. per agent run) rather than holding one
open indefinitely, so memory doesn't grow without bound.
Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. - data-gras/savi-loop-guard

GitHub - data-gras/savi-loop-guard: Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block. · GitHub
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
data-gras
/
savi-loop-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.github .github loop_guard loop_guard tests tests .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml setup.py setup.py View all files Repository files navigation
A zero-dependency Python library that detects when an AI agent is stuck in a
loop: calling the same tool over and over, or firing calls far faster than
any real workflow would. Drop it into your own agent code. No account, no
API key, no network call, ever.
Built by SAVI as a standalone, dependency-free
package, so you can detect these patterns in your own agent code without
an account, an API key, or a dependency on SAVI's platform.
pip install savi-loop-guard
No dependencies. Nothing else gets installed alongside it.
Record each call as your agent makes it, then check for loops whenever you
want (after every call, on a timer, whatever fits your loop):
from loop_guard import LoopGuard , CallEvent
from datetime import datetime , timezone
guard = LoopGuard ()
guard . record ( CallEvent (
span_id = "call_1" ,
agent_id = "doc-extractor" ,
timestamp = datetime . now ( timezone . utc ),
tool_call = "search_web" ,
))
issues = guard . check ()
for issue in issues :
print ( issue [ "type" ], issue [ "agent_id" ])
An issue looks like:
{ "type" : LoopType . VELOCITY , "agent_id" : "doc-extractor" , "elapsed_s" : 12.4 , "call_count" : 6 }
# or
{ "type" : LoopType . STRUCTURAL , "agent_id" : "doc-extractor" ,
"tool_call" : "search_web" , "tool_variants" : [ "search_web" , "search_web_v2" ], "call_count" : 7 }
Prevent the call instead of just observing it
check_before_call() records the event and raises immediately if it would
trip a threshold, for callers who want to stop the loop rather than find
out about it afterwards:
from loop_guard import LoopGuard , CallEvent , LoopDetected
guard = LoopGuard ()
try :
guard . check_before_call ( event )
except LoopDetected as e :
print ( f"Blocked: { e . loop_type } " ) # "velocity_loop" or "structural_loop"
print ( e . details ) # the same dict check() would have returned
What it detects
Velocity loop : more than 5 calls from the same agent_id within a
30-second window.
Structural loop : the same tool called more than 5 times, with fuzzy
matching so a broken agent can't dodge detection by alternating between
near-identical tool names ( search_web vs search_web_v2 vs web_search
count as the same tool if their name tokens overlap enough).
All four numbers are configurable:
guard = LoopGuard (
velocity_window_seconds = 30 ,
velocity_call_limit = 5 ,
structural_call_limit = 5 ,
tool_fuzzy_similarity_threshold = 0.70 ,
)
How this compares
A few other standalone Python packages exist for this: agent-loop-detector ,
agent-loop-guard , agentguard-kit . All of them, like savi-loop-guard 's
record() / check() API, are post-hoc/observational; they analyze calls
after they happen. savi-loop-guard adds check_before_call() on top for
callers who want to prevent the call rather than just observe it, which none
of those currently offer.
LoopGuard keeps every recorded event in memory for the life of the
instance; nothing is ever pruned automatically. That's deliberate:
structural-loop detection is a total call count with no time bound by
design (a tool called 6 times over 3 hours is still a loop, not just a
tool called 6 times in 30 seconds), so silently dropping "old" events
would blind it to exactly the slow, steady loops it exists to catch.
In practice this means: for a short-lived task, a single LoopGuard()
is fine as-is. For a long-running process, create a fresh LoopGuard()
per logical unit of work (e.g. per agent run) rather than holding one
open indefinitely, so memory doesn't grow without bound.
Zero-dependency Python library that detects an AI agent stuck in a loop, velocity + fuzzy tool-name matching, with an optional pre-call block.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
