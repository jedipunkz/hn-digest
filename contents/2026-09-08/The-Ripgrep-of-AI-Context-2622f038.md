---
source: "https://github.com/redhat-et/ripwire"
hn_url: "https://news.ycombinator.com/item?id=49611931"
title: "The Ripgrep of AI Context"
article_title: "GitHub - redhat-et/ripwire: The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess l\n[truncated]"
image: "https://opengraph.githubassets.com/ce25400b21bf07195d4f79bd5484dd89e56bc66b1829481cb57f6dd7893ac6b0/redhat-et/ripwire"
author: "schmorptron"
captured_at: "2026-09-08T16:07:01Z"
capture_tool: "hn-digest"
hn_id: 49611931
score: 2
comments: 0
posted_at: "2026-09-08T15:39:12Z"
tags:
  - hacker-news
---

# The Ripgrep of AI Context

- HN: [49611931](https://news.ycombinator.com/item?id=49611931)
- Source: [github.com](https://github.com/redhat-et/ripwire)
- Score: 2
- Comments: 0
- Posted: 2026-09-08T15:39:12Z

## Translation

Title: The Ripgrep of AI Context
Article title: GitHub - redhat-et/ripwire: The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess l
[truncated]
Description: The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess labelled, every loss publishe
[truncated]

Article text:
GitHub - redhat-et/ripwire: The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess labelled, every loss published. Paddle out with a map. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
redhat-et
/
ripwire
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,910 Commits 1,910 Commits Folders and files
.codex-plugin .codex-plugin .github .github bench bench cmake cmake docs docs hooks hooks paper paper present present prompts prompts queries queries scripts scripts skills skills src src test test third_party third_party .clang-format .clang-format .clang-tidy .clang-tidy .git-blame-ignore-revs .git-blame-ignore-revs .gitignore .gitignore .mcp.json .mcp.json .ripwire_notes .ripwire_notes .ripwire_quality_acks .ripwire_quality_acks AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md install.sh install.sh lsan_suppressions.txt lsan_suppressions.txt View all files Repository files navigation
Rip'n Fast. Less Tokens. Better Code.
Give your coding agent a map before it reads the repo.
ripwire is the ripgrep of AI context. Point it at any repository and your agent gets a ranked,
deterministic call graph — what to touch, what it breaks, which tests to run — instead of grepping
around and reading whole files.
Beside those sits a labelled survey of 237 tools that contributed nothing and says so. The two
sets are disjoint by construction, so they add rather than nest — a tool that gave a lesson is never
counted twice.
Both halves are load-bearing, and they are doing different jobs. The settled results are what
make the quality lens trustworthy: McCabe on complexity (1976), Halstead on volume (1977), Spärck
Jones on term specificity (1972), Nagappan & Ball on churn. Fifty years of replication means those
are not opinions, and a tool that measures your code should be built on the ones that survived.
The recent work is what makes it current : seventeen of the folded papers are from 2026, seven
published in the last two months and three in the last thirty days (dates as of 2026-09-08; every
row carries its arXiv id, so the claim is checkable rather than atmospheric). Retrieval for coding
agents, context-compression cost, placebo-controlled localization — that literature is months old,
not decades, and several rows were folded within weeks of the paper appearing.
Neither half alone would be enough. A tool built only on the classics would not know what an agent
needs; one built only on last month's preprints would have nothing underneath it. And the newest row
is a result that failed when it was tested here — which is the point of writing them down. All three counts are re-derived from that document's own tables by
test/readmedriftcheck.sh on every run, which fails if this page and those tables disagree, so the
claim cannot quietly drift. The row-by-row ledger is
docs/LINEAGE.md .
Languages: Rust · C++ · Objective-C/C++ · C · Metal · CUDA · Python · Go · Swift · TypeScript ·
JavaScript · Java · Ruby · PHP · Lua · Elixir · Bash · C# · JSON · TOML · YAML · Markdown — see
language support and limits .
No API key. No embeddings. No index server. No daemon.
One process, no server — indexes this repository in 0.25 s using 6.6 MB , against 46.8 s and 391 MB for the graph-database MCP server it was measured against; warm queries answer in 197 ms to its 1,082 ms
Measured on 48 matched questions across django, webpack and this repository. Across all three,
ripwire indexes in 0.25–0.45 s and 6.6–16.5 MB against that server's 23–52 s and
391–623 MB . The full method, the wins named one by one and the losses included, is in
Against the leading graph-database code-context MCP server
and docs/EVALS.md .
One self-contained binary on your own machine, offline, installed in one line — and the same line
installs and activates the task-shaped skills that teach your agent when to reach for it, not
just how, for every agent it finds on the machine. If your
agent can run shell commands — Claude Code, Codex, Cursor, Windsurf, Gemini, opencode, aider — it is
set up the moment the install finishes; the MCP server is the optional second
interface . Install it and ask it something before you finish
reading this page:
RIPWIRE_REPO=redhat-et/ripwire bash -c " $( curl -fsSL https://raw.githubusercontent.com/redhat-et/ripwire/main/scripts/install.sh ) "
export PATH= " $HOME /.local/bin: $PATH " # where it installed; the installer prints this line if you need it
cd your-repo
ripwire . --for= " <the change you are about to make, in words> "
Reach for the CLI first — it is the cheaper interface. The MCP server is the optional second
way in, and its convenience has a cost the shell pipe does not carry: its verb schemas sit in your
agent's context every session, whether or not it calls them.
Same answer, a fraction of the tokens — read this table first if your agent is on a budget
How these ten rows were measured — 2026-08-08 , figures in ~tokens (≈ bytes/4), every ratio from a real run reproduced by the command in its row
Ten everyday moments, re-measured on this repository, 2026-08-08. Figures are ~tokens (≈ bytes/4);
every ratio comes from a real run, reproduced by the command in its row — raw byte counts and exact
commands in
docs/EVALS.md §5 .
Ordered understand → navigate → review-the-change:
These aren't summaries that gamble with information. Each row is scored
same-correct-answer-or-it-doesn't-count, and both sides were checked, not assumed: orient surfaces
this repo's own pipeline files ( ingest.cpp , graph.h , serialize.h ) in the first screen, the same
three docs/ARCHITECTURE.md names as central; the --for row lands mcpStale
( src/mcpindex.h:633 ), the actual staleness check, 5th-ranked; --recall lands the container-rule
doc ( AGENTS.md ) that states, verbatim, the same "no std::map " rule CONTRIBUTING.md explains in
full; --pack-task names the same three touch points a human would — cachelint.h , mergeCachePack
( src/main.cpp:1787 ), the lintrules.h helpers it reuses; --expand --top-k=0 hands back the
requested function's complete, unmodified body — the ranked-neighborhood addition costs the same
~22.6 KB regardless of which function you ask for, confirmed on two (a fixed floor, not per-function
variance); --callers on langOfPath names its 2 real callers, the same ones a grep hit-list
buries under 5 files of comment-only mentions; --impact + --uses on coversOrEquals names the same
2 direct call sites --uses alone would, plus (disclosed) a transitive reach --uses doesn't cover
at all; --from-trace resolves all 7 frames of a real call chain by name to the same definitions a
per-frame grep would eventually find, mixed with call sites and comments; --situ on a 2-file diff
names the same 6 real test harnesses, 2 of which a filename grep across test/ cannot find even
after opening every one of its 41 candidates — a completeness gap, not just a byte one; --pr-context
surfaces co-change partners ( test/regression.sh , src/main.cpp ) a raw git diff has no way to
know were usually touched and weren't this time. The map ranks and discloses — it never paraphrases
your code — and every truncation is disclosed in the header.
The honesty line, made concrete: the same auto-selection behind the --expand row also runs the
other way. On a small file ( pageRankDouble in src/pagerank.cpp , 5,559 B) the ranked bundle would
cost 27,916 B — nearly 5× more than the file — so ripwire serves the file itself instead, disclosed
as mode="whole-file" on the response, not silently.
docs/EVALS.md §7 lists that and the other counterexamples
this project publishes against itself.
The shape: plan the work, then run a loop that matches each task to the model that fits it. Every
thread it spawns opens with an empty context on a repository it has never seen. Orienting an empty
context to a large repository is the most repeated cost in the whole system, and the one this tool
was built for — it is also the cost that grows with the size of the tree, which is why the pattern
matters more the bigger the repository gets.
--for and --pack-task answer it in a single call, at the per-call rates in the table above,
instead of a grep-and-read tour that every lane pays over again from scratch.
The return path matters as much. --quality-delta , --test-gate and --edit-check answer what did
I make worse , which tests must run , did I change a contract — quantitative answers a lane can
hand back as a verdict, rather than a transcript the orchestrator has to read to find out what
happened.
What is and is not claimed here. Every figure on this page is a single-agent measurement. That
the saving compounds with the number of cold orientations follows from the fixed-cost mechanism, but
it is pre-registered and unrun — the reason the pattern is worth trying, not a result this
project has published.
See the map — not just the numbers
Django's migration autodetector, coloured by complexity. Thresholds are fixed, so the colour means the same thing on every repo you point it at.
ripwire path/to/django/db/migrations --rank-by=rrf --top-k=120 --color-by=cx --html=map.html
The same graph, re-coloured by git churn. 76% of these nodes move to a different band — structure and history disagree, and one run shows you both.
A dashed shaft is a guess. 31 of 183 edges here are one arm of a split the resolver could not choose between. No other tool marks which of its arrows it is unsure about.
How to read these pictures — the five lenses, the fixed thresholds, and what the renderer refuses to draw
One self-contained HTML file ( --html[=FILE] ), no server, no CDN, no external asset. --color-by=lang|community|cx|churn|tested sets the initial colour; the page embeds all five and keeps a live selector, so switching lens costs no second run.
Read from the figures above, which state their own rules in a sidecar saved beside each image:
arrow points caller → callee — the graph is directed, and the page draws it that way.
31 of 183 shafts dashed in this view = the resolver could not choose between same-name definitions and split the call over all of them — per edge , not per symbol. A symbol-level "this function makes some ambiguous calls" would mark every one of its edges, which would

[truncated]

## Original Extract

The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess labelled, every loss publishe
[truncated]

GitHub - redhat-et/ripwire: The ripgrep of AI context: a zero-dependency C++23 CLI + MCP server for coding agents. Find what you want without reading the repo, then check you built what you meant — blast radius, tests-to-run, quality deltas. Signatures at 74.7% fewer bytes than bodies; every guess labelled, every loss published. Paddle out with a map. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
redhat-et
/
ripwire
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,910 Commits 1,910 Commits Folders and files
.codex-plugin .codex-plugin .github .github bench bench cmake cmake docs docs hooks hooks paper paper present present prompts prompts queries queries scripts scripts skills skills src src test test third_party third_party .clang-format .clang-format .clang-tidy .clang-tidy .git-blame-ignore-revs .git-blame-ignore-revs .gitignore .gitignore .mcp.json .mcp.json .ripwire_notes .ripwire_notes .ripwire_quality_acks .ripwire_quality_acks AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md install.sh install.sh lsan_suppressions.txt lsan_suppressions.txt View all files Repository files navigation
Rip'n Fast. Less Tokens. Better Code.
Give your coding agent a map before it reads the repo.
ripwire is the ripgrep of AI context. Point it at any repository and your agent gets a ranked,
deterministic call graph — what to touch, what it breaks, which tests to run — instead of grepping
around and reading whole files.
Beside those sits a labelled survey of 237 tools that contributed nothing and says so. The two
sets are disjoint by construction, so they add rather than nest — a tool that gave a lesson is never
counted twice.
Both halves are load-bearing, and they are doing different jobs. The settled results are what
make the quality lens trustworthy: McCabe on complexity (1976), Halstead on volume (1977), Spärck
Jones on term specificity (1972), Nagappan & Ball on churn. Fifty years of replication means those
are not opinions, and a tool that measures your code should be built on the ones that survived.
The recent work is what makes it current : seventeen of the folded papers are from 2026, seven
published in the last two months and three in the last thirty days (dates as of 2026-09-08; every
row carries its arXiv id, so the claim is checkable rather than atmospheric). Retrieval for coding
agents, context-compression cost, placebo-controlled localization — that literature is months old,
not decades, and several rows were folded within weeks of the paper appearing.
Neither half alone would be enough. A tool built only on the classics would not know what an agent
needs; one built only on last month's preprints would have nothing underneath it. And the newest row
is a result that failed when it was tested here — which is the point of writing them down. All three counts are re-derived from that document's own tables by
test/readmedriftcheck.sh on every run, which fails if this page and those tables disagree, so the
claim cannot quietly drift. The row-by-row ledger is
docs/LINEAGE.md .
Languages: Rust · C++ · Objective-C/C++ · C · Metal · CUDA · Python · Go · Swift · TypeScript ·
JavaScript · Java · Ruby · PHP · Lua · Elixir · Bash · C# · JSON · TOML · YAML · Markdown — see
language support and limits .
No API key. No embeddings. No index server. No daemon.
One process, no server — indexes this repository in 0.25 s using 6.6 MB , against 46.8 s and 391 MB for the graph-database MCP server it was measured against; warm queries answer in 197 ms to its 1,082 ms
Measured on 48 matched questions across django, webpack and this repository. Across all three,
ripwire indexes in 0.25–0.45 s and 6.6–16.5 MB against that server's 23–52 s and
391–623 MB . The full method, the wins named one by one and the losses included, is in
Against the leading graph-database code-context MCP server
and docs/EVALS.md .
One self-contained binary on your own machine, offline, installed in one line — and the same line
installs and activates the task-shaped skills that teach your agent when to reach for it, not
just how, for every agent it finds on the machine. If your
agent can run shell commands — Claude Code, Codex, Cursor, Windsurf, Gemini, opencode, aider — it is
set up the moment the install finishes; the MCP server is the optional second
interface . Install it and ask it something before you finish
reading this page:
RIPWIRE_REPO=redhat-et/ripwire bash -c " $( curl -fsSL https://raw.githubusercontent.com/redhat-et/ripwire/main/scripts/install.sh ) "
export PATH= " $HOME /.local/bin: $PATH " # where it installed; the installer prints this line if you need it
cd your-repo
ripwire . --for= " <the change you are about to make, in words> "
Reach for the CLI first — it is the cheaper interface. The MCP server is the optional second
way in, and its convenience has a cost the shell pipe does not carry: its verb schemas sit in your
agent's context every session, whether or not it calls them.
Same answer, a fraction of the tokens — read this table first if your agent is on a budget
How these ten rows were measured — 2026-08-08 , figures in ~tokens (≈ bytes/4), every ratio from a real run reproduced by the command in its row
Ten everyday moments, re-measured on this repository, 2026-08-08. Figures are ~tokens (≈ bytes/4);
every ratio comes from a real run, reproduced by the command in its row — raw byte counts and exact
commands in
docs/EVALS.md §5 .
Ordered understand → navigate → review-the-change:
These aren't summaries that gamble with information. Each row is scored
same-correct-answer-or-it-doesn't-count, and both sides were checked, not assumed: orient surfaces
this repo's own pipeline files ( ingest.cpp , graph.h , serialize.h ) in the first screen, the same
three docs/ARCHITECTURE.md names as central; the --for row lands mcpStale
( src/mcpindex.h:633 ), the actual staleness check, 5th-ranked; --recall lands the container-rule
doc ( AGENTS.md ) that states, verbatim, the same "no std::map " rule CONTRIBUTING.md explains in
full; --pack-task names the same three touch points a human would — cachelint.h , mergeCachePack
( src/main.cpp:1787 ), the lintrules.h helpers it reuses; --expand --top-k=0 hands back the
requested function's complete, unmodified body — the ranked-neighborhood addition costs the same
~22.6 KB regardless of which function you ask for, confirmed on two (a fixed floor, not per-function
variance); --callers on langOfPath names its 2 real callers, the same ones a grep hit-list
buries under 5 files of comment-only mentions; --impact + --uses on coversOrEquals names the same
2 direct call sites --uses alone would, plus (disclosed) a transitive reach --uses doesn't cover
at all; --from-trace resolves all 7 frames of a real call chain by name to the same definitions a
per-frame grep would eventually find, mixed with call sites and comments; --situ on a 2-file diff
names the same 6 real test harnesses, 2 of which a filename grep across test/ cannot find even
after opening every one of its 41 candidates — a completeness gap, not just a byte one; --pr-context
surfaces co-change partners ( test/regression.sh , src/main.cpp ) a raw git diff has no way to
know were usually touched and weren't this time. The map ranks and discloses — it never paraphrases
your code — and every truncation is disclosed in the header.
The honesty line, made concrete: the same auto-selection behind the --expand row also runs the
other way. On a small file ( pageRankDouble in src/pagerank.cpp , 5,559 B) the ranked bundle would
cost 27,916 B — nearly 5× more than the file — so ripwire serves the file itself instead, disclosed
as mode="whole-file" on the response, not silently.
docs/EVALS.md §7 lists that and the other counterexamples
this project publishes against itself.
The shape: plan the work, then run a loop that matches each task to the model that fits it. Every
thread it spawns opens with an empty context on a repository it has never seen. Orienting an empty
context to a large repository is the most repeated cost in the whole system, and the one this tool
was built for — it is also the cost that grows with the size of the tree, which is why the pattern
matters more the bigger the repository gets.
--for and --pack-task answer it in a single call, at the per-call rates in the table above,
instead of a grep-and-read tour that every lane pays over again from scratch.
The return path matters as much. --quality-delta , --test-gate and --edit-check answer what did
I make worse , which tests must run , did I change a contract — quantitative answers a lane can
hand back as a verdict, rather than a transcript the orchestrator has to read to find out what
happened.
What is and is not claimed here. Every figure on this page is a single-agent measurement. That
the saving compounds with the number of cold orientations follows from the fixed-cost mechanism, but
it is pre-registered and unrun — the reason the pattern is worth trying, not a result this
project has published.
See the map — not just the numbers
Django's migration autodetector, coloured by complexity. Thresholds are fixed, so the colour means the same thing on every repo you point it at.
ripwire path/to/django/db/migrations --rank-by=rrf --top-k=120 --color-by=cx --html=map.html
The same graph, re-coloured by git churn. 76% of these nodes move to a different band — structure and history disagree, and one run shows you both.
A dashed shaft is a guess. 31 of 183 edges here are one arm of a split the resolver could not choose between. No other tool marks which of its arrows it is unsure about.
How to read these pictures — the five lenses, the fixed thresholds, and what the renderer refuses to draw
One self-contained HTML file ( --html[=FILE] ), no server, no CDN, no external asset. --color-by=lang|community|cx|churn|tested sets the initial colour; the page embeds all five and keeps a live selector, so switching lens costs no second run.
Read from the figures above, which state their own rules in a sidecar saved beside each image:
arrow points caller → callee — the graph is directed, and the page draws it that way.
31 of 183 shafts dashed in this view = the resolver could not choose between same-name definitions and split the call over all of them — per edge , not per symbol. A symbol-level "this function makes some ambiguous calls" would mark every one of its edges, which would

[truncated]
