---
source: "https://github.com/AndyCross/claude-hot-path"
hn_url: "https://news.ycombinator.com/item?id=50610333"
title: "Show HN: Claude Hot Path, one-file Claude clients in 7 languages, no SDK"
article_title: "GitHub - AndyCross/claude-hot-path: The fastest way from your code to Claude. One file, no SDK. · GitHub"
image: "https://opengraph.githubassets.com/6e5b75c7d789e1fbdac35de25cccd4708982c94ad3b31ed795e53686dd914ac3/AndyCross/claude-hot-path"
author: "stepintodev"
captured_at: "2026-11-23T11:53:34Z"
capture_tool: "hn-digest"
hn_id: 50610333
score: 1
comments: 0
posted_at: "2026-11-23T11:33:58Z"
tags:
  - hacker-news
---

# Show HN: Claude Hot Path, one-file Claude clients in 7 languages, no SDK

- HN: [50610333](https://news.ycombinator.com/item?id=50610333)
- Source: [github.com](https://github.com/AndyCross/claude-hot-path)
- Score: 1
- Comments: 0
- Posted: 2026-11-23T11:33:58Z

## Translation

Title: Show HN: Claude Hot Path, one-file Claude clients in 7 languages, no SDK
Article title: GitHub - AndyCross/claude-hot-path: The fastest way from your code to Claude. One file, no SDK. · GitHub
Description: The fastest way from your code to Claude. One file, no SDK. - AndyCross/claude-hot-path

Article text:
GitHub - AndyCross/claude-hot-path: The fastest way from your code to Claude. One file, no SDK. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AndyCross
claude-hot-path Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
The fastest way from your code to Claude. One file, no SDK.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.github/ workflows .github/ workflows bench bench clients clients fixtures fixtures mock mock results results runner runner .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md compose.yaml compose.yaml go.work go.work View all files Repository files navigation
The fastest way from your code to Claude. One file, no SDK.
The hot path to Claude is one HTTPS POST and a stream of six Server-Sent Event
types. Each client here writes it by hand in one source file, in one language,
and the harness measures what it adds on top of the API: cold start, the
turnaround between the two requests of a tool call, memory, and size. The
leaderboard sorts by the milliseconds added, fastest first.
Posts about each one are at https://stepinto.dev/claude-hot-path/ .
Median of 100 cold starts for each client, under Docker, against the mock server. Client time is every millisecond the mock was not streaming: cold start, turnaround and exit added together.
"Ships" is the artefact plus any runtime it needs. The Python SDK row is the reference and is not an entry: it uses the official anthropic package the way its documentation says. The full board, with p95 figures, toolchain versions and the commit each row was measured at, is results/leaderboard.csv .
./bench/bench.sh bash # build the bash client, check it, measure 100 cold starts
./bench/bench.sh all # every client, then results/leaderboard.csv
You need Docker and Python 3. You do not need an API key: the mock server in
mock/ replays recorded exchanges with their original timing. The only step
that talks to Anthropic is recording, which proxies one real exchange through
the mock and saves it:
ANTHROPIC_API_KEY=... ./bench/bench.sh record bash
The key is read from your shell by the proxy. The client under test is given a
placeholder and never sees it, and nothing writes it to disk.
Every client must do the following, and the bench checks each point rather
than trusting the code.
Read the key from ANTHROPIC_API_KEY and the base URL from
ANTHROPIC_BASE_URL (default https://api.anthropic.com ). Nothing else
from the environment.
POST /v1/messages with headers x-api-key , anthropic-version: 2023-06-01 and content-type: application/json , and a body with model ,
max_tokens , stream: true , messages and tools .
Print each text_delta as it arrives. The bench checks that the first
visible byte reached stdout before the mock finished sending the stream
that carried it, so buffering the reply, or flushing only on newlines,
fails.
Declare one tool, local_time , taking an IANA timezone . When the stream
ends with stop_reason: "tool_use" , assemble the input_json_delta
fragments, run the tool, send the tool_result in a second request with the
assistant turn echoed back, and stream that reply too.
On a non-2xx response, print the body's error.message to stderr and exit
non-zero.
One source file. HTTP and TLS from the platform where it has them. Where it
does not, one named library, which counts towards size.
clients/csharp-tuned/ is the same spec written for speed rather than the
way the templates say: libc's socket calls and hand-written HTTP/1.1 in place
of HttpClient , Utf8JsonReader in place of JsonNode , bare file
descriptors in place of Console , _exit in place of runtime shutdown, and
a static musl build from scratch . It bends rule 7, so it sits on the board
as its own row next to the stock csharp client rather than replacing it.
Column
Meaning
artefact_bytes
The stripped binary, or the source file for an interpreted language.
runtime_bytes
Zero for a static binary. The interpreter and any shared runtime otherwise, measured inside the image.
cold_start_ms
Process spawn to the mock receiving the first request. Median and p95.
turnaround_ms
First reply finished to second request received: parse the stream, run the tool, build and send the next request.
tail_ms
Second reply finished to exit.
client_ms
The sum of the three: every millisecond the mock was not streaming. Median and p95.
max_rss_kb
Peak resident memory, from the kernel's rusage for the child.
source_lines
Non-blank, non-comment lines in the one source file. Manifests (Cargo.toml, the csproj) are not counted.
build_s
Clean image build, reported and not ranked.
The streaming check uses the same clocks: the runner notes when the first
non-whitespace byte reached stdout and the mock notes when it wrote each chunk,
and the first must come before the last. On the recorded fixture a streaming
client has about 11 ms in hand and a buffering one is 0.3 to 1 ms late, which
the check catches. All timings are differences between timestamps the mock recorded and the
runner's spawn and exit times, which share a clock because the containers share
a kernel. Nothing is subtracted, so the mock's own sleep overshoot never counts
against a client. The leaderboard sorts by client_ms , fastest first. Nothing is
measured over the network: the one real exchange is for conformance, and no
timing from it is published.
Make clients/<name>/ with the source file, a Dockerfile whose ENTRYPOINT
runs it, and a meta.env :
NAME=zig
KIND=binary # interpreted, or sdk for an official-SDK baseline row
ARTIFACT=/app/client # path inside the image
RUNTIME= ' ' # interpreter binaries, for interpreted and sdk
RUNTIME_DIRS= ' ' # directories priced whole, e.g. a stdlib and site-packages
SOURCE=client.zig # the one file, for the lines column
Write /toolchain.txt in the image during the build; the bench reads it.
Then ./bench/bench.sh <name> . CI runs the mock bench on every pull request.
The fastest way from your code to Claude. One file, no SDK.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The fastest way from your code to Claude. One file, no SDK. - AndyCross/claude-hot-path

GitHub - AndyCross/claude-hot-path: The fastest way from your code to Claude. One file, no SDK. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AndyCross
claude-hot-path Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
The fastest way from your code to Claude. One file, no SDK.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.github/ workflows .github/ workflows bench bench clients clients fixtures fixtures mock mock results results runner runner .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md compose.yaml compose.yaml go.work go.work View all files Repository files navigation
The fastest way from your code to Claude. One file, no SDK.
The hot path to Claude is one HTTPS POST and a stream of six Server-Sent Event
types. Each client here writes it by hand in one source file, in one language,
and the harness measures what it adds on top of the API: cold start, the
turnaround between the two requests of a tool call, memory, and size. The
leaderboard sorts by the milliseconds added, fastest first.
Posts about each one are at https://stepinto.dev/claude-hot-path/ .
Median of 100 cold starts for each client, under Docker, against the mock server. Client time is every millisecond the mock was not streaming: cold start, turnaround and exit added together.
"Ships" is the artefact plus any runtime it needs. The Python SDK row is the reference and is not an entry: it uses the official anthropic package the way its documentation says. The full board, with p95 figures, toolchain versions and the commit each row was measured at, is results/leaderboard.csv .
./bench/bench.sh bash # build the bash client, check it, measure 100 cold starts
./bench/bench.sh all # every client, then results/leaderboard.csv
You need Docker and Python 3. You do not need an API key: the mock server in
mock/ replays recorded exchanges with their original timing. The only step
that talks to Anthropic is recording, which proxies one real exchange through
the mock and saves it:
ANTHROPIC_API_KEY=... ./bench/bench.sh record bash
The key is read from your shell by the proxy. The client under test is given a
placeholder and never sees it, and nothing writes it to disk.
Every client must do the following, and the bench checks each point rather
than trusting the code.
Read the key from ANTHROPIC_API_KEY and the base URL from
ANTHROPIC_BASE_URL (default https://api.anthropic.com ). Nothing else
from the environment.
POST /v1/messages with headers x-api-key , anthropic-version: 2023-06-01 and content-type: application/json , and a body with model ,
max_tokens , stream: true , messages and tools .
Print each text_delta as it arrives. The bench checks that the first
visible byte reached stdout before the mock finished sending the stream
that carried it, so buffering the reply, or flushing only on newlines,
fails.
Declare one tool, local_time , taking an IANA timezone . When the stream
ends with stop_reason: "tool_use" , assemble the input_json_delta
fragments, run the tool, send the tool_result in a second request with the
assistant turn echoed back, and stream that reply too.
On a non-2xx response, print the body's error.message to stderr and exit
non-zero.
One source file. HTTP and TLS from the platform where it has them. Where it
does not, one named library, which counts towards size.
clients/csharp-tuned/ is the same spec written for speed rather than the
way the templates say: libc's socket calls and hand-written HTTP/1.1 in place
of HttpClient , Utf8JsonReader in place of JsonNode , bare file
descriptors in place of Console , _exit in place of runtime shutdown, and
a static musl build from scratch . It bends rule 7, so it sits on the board
as its own row next to the stock csharp client rather than replacing it.
Column
Meaning
artefact_bytes
The stripped binary, or the source file for an interpreted language.
runtime_bytes
Zero for a static binary. The interpreter and any shared runtime otherwise, measured inside the image.
cold_start_ms
Process spawn to the mock receiving the first request. Median and p95.
turnaround_ms
First reply finished to second request received: parse the stream, run the tool, build and send the next request.
tail_ms
Second reply finished to exit.
client_ms
The sum of the three: every millisecond the mock was not streaming. Median and p95.
max_rss_kb
Peak resident memory, from the kernel's rusage for the child.
source_lines
Non-blank, non-comment lines in the one source file. Manifests (Cargo.toml, the csproj) are not counted.
build_s
Clean image build, reported and not ranked.
The streaming check uses the same clocks: the runner notes when the first
non-whitespace byte reached stdout and the mock notes when it wrote each chunk,
and the first must come before the last. On the recorded fixture a streaming
client has about 11 ms in hand and a buffering one is 0.3 to 1 ms late, which
the check catches. All timings are differences between timestamps the mock recorded and the
runner's spawn and exit times, which share a clock because the containers share
a kernel. Nothing is subtracted, so the mock's own sleep overshoot never counts
against a client. The leaderboard sorts by client_ms , fastest first. Nothing is
measured over the network: the one real exchange is for conformance, and no
timing from it is published.
Make clients/<name>/ with the source file, a Dockerfile whose ENTRYPOINT
runs it, and a meta.env :
NAME=zig
KIND=binary # interpreted, or sdk for an official-SDK baseline row
ARTIFACT=/app/client # path inside the image
RUNTIME= ' ' # interpreter binaries, for interpreted and sdk
RUNTIME_DIRS= ' ' # directories priced whole, e.g. a stdlib and site-packages
SOURCE=client.zig # the one file, for the lines column
Write /toolchain.txt in the image during the build; the bench reads it.
Then ./bench/bench.sh <name> . CI runs the mock bench on every pull request.
The fastest way from your code to Claude. One file, no SDK.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
