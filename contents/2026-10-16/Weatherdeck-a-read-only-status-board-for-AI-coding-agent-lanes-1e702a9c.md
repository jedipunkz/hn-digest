---
source: "https://github.com/Moshik21/weatherdeck"
hn_url: "https://news.ycombinator.com/item?id=50122767"
title: "Weatherdeck – a read-only status board for AI coding agent lanes"
article_title: "GitHub - Moshik21/weatherdeck: A read-only, at-a-glance status board for work reported by many sources. · GitHub"
image: "https://opengraph.githubassets.com/7c144eb1edca26ec96ca5af93ebe5c5ac5f52e61bfd1589e01b96513fdc53d86/Moshik21/weatherdeck"
author: "Konnerm"
captured_at: "2026-10-16T19:22:06Z"
capture_tool: "hn-digest"
hn_id: 50122767
score: 1
comments: 0
posted_at: "2026-10-16T18:54:22Z"
tags:
  - hacker-news
---

# Weatherdeck – a read-only status board for AI coding agent lanes

- HN: [50122767](https://news.ycombinator.com/item?id=50122767)
- Source: [github.com](https://github.com/Moshik21/weatherdeck)
- Score: 1
- Comments: 0
- Posted: 2026-10-16T18:54:22Z

## Translation

Title: Weatherdeck – a read-only status board for AI coding agent lanes
Article title: GitHub - Moshik21/weatherdeck: A read-only, at-a-glance status board for work reported by many sources. · GitHub
Description: A read-only, at-a-glance status board for work reported by many sources. - Moshik21/weatherdeck

Article text:
GitHub - Moshik21/weatherdeck: A read-only, at-a-glance status board for work reported by many sources. · GitHub
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
Moshik21
weatherdeck Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A read-only, at-a-glance status board for work reported by many sources.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github .github bin bin concepts concepts docs docs examples examples scripts scripts src/ weatherdeck src/ weatherdeck tests tests web web .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md pip_not_supported.py pip_not_supported.py pyproject.toml pyproject.toml View all files Repository files navigation
Weatherdeck is a status board for machines and the work running on them (a TV or phone page fed by small pushers), not a weather app and not a monitoring or alerting system.
What you see. Every picture here is the demo ( weatherdeck run --demo , with #theme=ship-deck on the page address). Its sources are invented and nothing in it is real.
Waiting on you. The amber plate counts the calls that need a person and says how long each has waited.
Honest progress. Progress shows as a percent range where the source gives one (for example 38-42%). A lane the board cannot read is counted as "can't tell" (the "+7 more" line says 1), not shown as fine.
A quiet source is named. The lamp by the clock counts live sources. When one stops reporting it reads "2/3 live", and that machine's ship is drawn in fog and labelled silent.
One board for TV and phone. The same page is the open sea on a wide screen and one column on a phone.
In the clip, Cinder (one of the invented sources) stops reporting. Still frame .
The phone page from a few minutes into the demo, cropped after the ALL FINE tile: the calls, then the problems, then that tile, which is lit only when no call or problem is open and every feed reports.
It is an at-a-glance board for work reported by many sources, meant to be left on a TV or a
phone: what is waiting on you, what is moving and how far along it honestly is, what got done, what
is next, and whether you can trust what you are looking at.
Any tool can leave a card. One small JSON file in a drop directory becomes a status tile; the board only reads it ( docs/ADDINS.md ).
Any tool can feed it. A source is anything that can send one JSON document over HTTP. Sources are discovered, never configured; machines and items appear on their own.
Checkable. Every push a host receives is one line in its received log ( weatherdeck received --diff ): what each source sent and what changed on the board.
Honest about gaps. A source that stops reporting stays on screen, marked silent and fogged. The board never shows stale data as fresh, and says so when a source is missing. It claims no value for itself or any tool until a number is measured ( docs/honesty.md ).
Read-only, local, tiny. Python standard library only, one SQLite file, one plain HTML/JS/CSS page. No build step, no CDN, no telemetry, no root.
Runs anywhere. The same command natively on Windows, macOS and Linux, no admin needed. Run several hosts at once; every page says which host serves it.
Deterministic. No model or LLM call anywhere; the only outbound traffic is an adapter pushing to a host you name.
What this is not. It is not a monitoring or alerting system: it never pages anyone, and it shows only what sources push to
it. It is not a hosted service: you run it yourself, on machines you choose. And it is pre-1.0 hobby software, open-sourced
because it solved a problem the author had: details may change, and support is best effort.
You need Python 3.9 to 3.14 (an older one stops at startup with a one-line message; the Microsoft Store one is fine for trying it; to run it at logon or run the
beat, see docs/runbook-windows.md ). Nothing to install: run it from a source checkout or the release zip. Get the code by cloning this repository
( git clone it, then cd into the folder), or by unzipping the zip that scripts/build-package builds from a clone. pip install is not supported yet and stops with a
message saying so. Below, weatherdeck means python3 bin/weatherdeck on macOS/Linux and bin\weatherdeck.cmd on Windows (cmd or PowerShell).
weatherdeck run --demo
Open http://127.0.0.1:8787/ . Three invented sources report; one stops reporting after about
two and a half minutes, shows late from about 2:40 and silent about three minutes after start, so you can see the offline state. Ctrl-C stops it. ( --port N to change the port.)
weatherdeck on # background; prints the URLs
weatherdeck status
weatherdeck off
The data lives in your per-user data directory ( --data DIR or $WEATHERDECK_DATA to choose another).
On Windows, weatherdeck startup install also starts it at every logon, with no admin rights.
Send it something. Any tool, one POST (from the same machine, no setup needed):
curl -X POST http://127.0.0.1:8787/api/ingest -H 'Content-Type: application/json' -d '{
"schema": "lookout.feed", "schema_version": "1.0",
"source": { "id": "my-laptop" },
"items": [{ "id": "build", "title": "Nightly build", "state": "running" }]
}'
The schema id reads lookout.feed on purpose: it is the wire name every adapter already speaks, kept unchanged
when the product was renamed (see docs/feed-spec.md , "Schema ids").
To feed it continuously, run the example adapter, which turns a hand-edited list plus this
machine's vitals into a feed and pushes it to every host in a list:
cp examples/adapter/work.example.json work.json
cp examples/adapter/sinks.example.txt sinks.txt # edit the hosts
python3 examples/adapter/push_example.py --work work.json --sinks sinks.txt
To feed it from your own tool, see docs/adapters.md , "Write a pusher for your tool".
Let other devices see it. Add addresses to hosts.bind in your config file (a LAN address, a
VPN range, or both), create a token per source
( weatherdeck token new my-laptop ), and open the page from your phone. Once any token
exists, pushes need Authorization: Bearer <token> . For a TV, open the page full-screen on any
machine that can reach the host, or in kiosk mode on the host itself.
The config file is <data>/config.json , in the data directory from step 2 ( weatherdeck status prints that
directory; --config FILE or $WEATHERDECK_CONFIG names another file). It does not exist until you create it. A minimal one:
{ "hosts": { "bind": ["localhost", "auto:192.168.1.0/24"] } }
Restart the host after editing it ( weatherdeck restart ). Every key is in docs/hosting.md .
Remove it. Stop the host with weatherdeck off , then delete the folder you unpacked or cloned and the data directory
(where weatherdeck status says it is). On Windows, run weatherdeck startup remove first if you ran startup install .
Nothing else is installed: no service, no registry entry, no files elsewhere.
Windows 11
macOS (Apple Silicon)
Linux (Ubuntu, WSL2)
Host ( weatherdeck run , on , ingest )
supported
supported
supported
Python
3.9 to 3.14 (3.8 is not supported)
3.9 to 3.14 (3.8 is not supported)
3.9 to 3.14 (3.8 is not supported)
Test suite on GitHub's hosted runners (first run, 2026-10-16; 433 tests, .github/workflows/test.yml , run by hand)
10 to 13 tests fail on every Python version. On the hosted Windows image nine tests that start Node fail because Windows rejects the command line as too long; one doctor check, one push-time check, one text-cutting check and one scratch-folder cleanup also fail there (the Python 3.9 and 3.10 legs fail 10, the others 13). The host itself runs natively on Windows 11, where the maintainer runs it, and this is not yet fixed
all pass
all pass
Page: Chrome, Edge
supported
supported
supported
Page: Safari (macOS), iPhone
headless WebKit (layout, overflow, console errors); the dashboard depth scene was checked and signed off on Safari, iPhone and a 4K TV on 2026-10-10 by the maintainer, and nothing else on those devices has been ( docs/device-checklist.md )
same
same
Page: Firefox
expected to work, not checked each release
Since you last looked (browser storage)
per browser; a browser that blocks site data shows no since plate
same
same
Feed schema read
1.0 to 1.11
1.0 to 1.11
1.0 to 1.11
One ship per machine (0.4.0)
TV: up to six ships, then a convoy; phone: the convoy strip scrolls inside itself
same
same
The open sea (0.4.1)
TV (ship deck): with nothing waiting, one banner line and the sea across the whole upper area; checked at 1920 x 1080 and 3840 x 2160
same
same
Versions per machine (0.5.0)
reporter rows, the watch strip, the trust line and the phone's SOFTWARE section say each source's adapter, or "version unknown"; only a gap that changes meaning is flagged
same
same
Last seen and the replay key (0.4.1)
per browser tab (session storage); a tab that blocks it starts with "not yet on this screen"
same
same
Slowed on purpose
[truncated]
What each release was actually run on is in CHANGELOG.md , with the checks only a person can make
(a real Safari, iPhone and TV; only the dashboard depth scene has been signed off so far, see
docs/device-checklist.md ) listed as such. Tests: python3 -m unittest discover -s tests (see CONTRIBUTING.md).
docs/honesty.md : the honesty contract on one page: what the board never does (green when silent, zero for unknown, a made-up event), where every number, creature and light comes from in the feed, the machine-alert debounce (two feeds to raise, two to clear, one keepalive of latency), and why no value number is shown until it is measured.
docs/ledger-spec.md : the tool value ledger: one append-only file per tool, unknown written as unknown with its reason, the honest verbs (surfaced, refused, placed, landed, not dispatched), and the reading rule that turns fourteen days of lines into KEEP, QUIET, CUT-CANDIDATE or UNMEASURED. No value number is shown until it is measured.
docs/feed-spec.md : the feed schema (1.11), how sources merge, the five lane classes, machine rest and strain, lane state (one word and one sentence per lane), goals (the goal each lane serves, with a count under it), allowance chips, black holes (no agent at the work, outside waits, queues, coming up, the shipping gap, freshness, burn rate), the library (one honest age and word per mirrored folder), the view at a glance (since you last looked, grouping by project), and what silence looks like.
CHANGELOG.md : what changed in each release, and what it was checked on; docs/device-checklist.md : w

[truncated]

## Original Extract

A read-only, at-a-glance status board for work reported by many sources. - Moshik21/weatherdeck

GitHub - Moshik21/weatherdeck: A read-only, at-a-glance status board for work reported by many sources. · GitHub
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
Moshik21
weatherdeck Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A read-only, at-a-glance status board for work reported by many sources.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github .github bin bin concepts concepts docs docs examples examples scripts scripts src/ weatherdeck src/ weatherdeck tests tests web web .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md pip_not_supported.py pip_not_supported.py pyproject.toml pyproject.toml View all files Repository files navigation
Weatherdeck is a status board for machines and the work running on them (a TV or phone page fed by small pushers), not a weather app and not a monitoring or alerting system.
What you see. Every picture here is the demo ( weatherdeck run --demo , with #theme=ship-deck on the page address). Its sources are invented and nothing in it is real.
Waiting on you. The amber plate counts the calls that need a person and says how long each has waited.
Honest progress. Progress shows as a percent range where the source gives one (for example 38-42%). A lane the board cannot read is counted as "can't tell" (the "+7 more" line says 1), not shown as fine.
A quiet source is named. The lamp by the clock counts live sources. When one stops reporting it reads "2/3 live", and that machine's ship is drawn in fog and labelled silent.
One board for TV and phone. The same page is the open sea on a wide screen and one column on a phone.
In the clip, Cinder (one of the invented sources) stops reporting. Still frame .
The phone page from a few minutes into the demo, cropped after the ALL FINE tile: the calls, then the problems, then that tile, which is lit only when no call or problem is open and every feed reports.
It is an at-a-glance board for work reported by many sources, meant to be left on a TV or a
phone: what is waiting on you, what is moving and how far along it honestly is, what got done, what
is next, and whether you can trust what you are looking at.
Any tool can leave a card. One small JSON file in a drop directory becomes a status tile; the board only reads it ( docs/ADDINS.md ).
Any tool can feed it. A source is anything that can send one JSON document over HTTP. Sources are discovered, never configured; machines and items appear on their own.
Checkable. Every push a host receives is one line in its received log ( weatherdeck received --diff ): what each source sent and what changed on the board.
Honest about gaps. A source that stops reporting stays on screen, marked silent and fogged. The board never shows stale data as fresh, and says so when a source is missing. It claims no value for itself or any tool until a number is measured ( docs/honesty.md ).
Read-only, local, tiny. Python standard library only, one SQLite file, one plain HTML/JS/CSS page. No build step, no CDN, no telemetry, no root.
Runs anywhere. The same command natively on Windows, macOS and Linux, no admin needed. Run several hosts at once; every page says which host serves it.
Deterministic. No model or LLM call anywhere; the only outbound traffic is an adapter pushing to a host you name.
What this is not. It is not a monitoring or alerting system: it never pages anyone, and it shows only what sources push to
it. It is not a hosted service: you run it yourself, on machines you choose. And it is pre-1.0 hobby software, open-sourced
because it solved a problem the author had: details may change, and support is best effort.
You need Python 3.9 to 3.14 (an older one stops at startup with a one-line message; the Microsoft Store one is fine for trying it; to run it at logon or run the
beat, see docs/runbook-windows.md ). Nothing to install: run it from a source checkout or the release zip. Get the code by cloning this repository
( git clone it, then cd into the folder), or by unzipping the zip that scripts/build-package builds from a clone. pip install is not supported yet and stops with a
message saying so. Below, weatherdeck means python3 bin/weatherdeck on macOS/Linux and bin\weatherdeck.cmd on Windows (cmd or PowerShell).
weatherdeck run --demo
Open http://127.0.0.1:8787/ . Three invented sources report; one stops reporting after about
two and a half minutes, shows late from about 2:40 and silent about three minutes after start, so you can see the offline state. Ctrl-C stops it. ( --port N to change the port.)
weatherdeck on # background; prints the URLs
weatherdeck status
weatherdeck off
The data lives in your per-user data directory ( --data DIR or $WEATHERDECK_DATA to choose another).
On Windows, weatherdeck startup install also starts it at every logon, with no admin rights.
Send it something. Any tool, one POST (from the same machine, no setup needed):
curl -X POST http://127.0.0.1:8787/api/ingest -H 'Content-Type: application/json' -d '{
"schema": "lookout.feed", "schema_version": "1.0",
"source": { "id": "my-laptop" },
"items": [{ "id": "build", "title": "Nightly build", "state": "running" }]
}'
The schema id reads lookout.feed on purpose: it is the wire name every adapter already speaks, kept unchanged
when the product was renamed (see docs/feed-spec.md , "Schema ids").
To feed it continuously, run the example adapter, which turns a hand-edited list plus this
machine's vitals into a feed and pushes it to every host in a list:
cp examples/adapter/work.example.json work.json
cp examples/adapter/sinks.example.txt sinks.txt # edit the hosts
python3 examples/adapter/push_example.py --work work.json --sinks sinks.txt
To feed it from your own tool, see docs/adapters.md , "Write a pusher for your tool".
Let other devices see it. Add addresses to hosts.bind in your config file (a LAN address, a
VPN range, or both), create a token per source
( weatherdeck token new my-laptop ), and open the page from your phone. Once any token
exists, pushes need Authorization: Bearer <token> . For a TV, open the page full-screen on any
machine that can reach the host, or in kiosk mode on the host itself.
The config file is <data>/config.json , in the data directory from step 2 ( weatherdeck status prints that
directory; --config FILE or $WEATHERDECK_CONFIG names another file). It does not exist until you create it. A minimal one:
{ "hosts": { "bind": ["localhost", "auto:192.168.1.0/24"] } }
Restart the host after editing it ( weatherdeck restart ). Every key is in docs/hosting.md .
Remove it. Stop the host with weatherdeck off , then delete the folder you unpacked or cloned and the data directory
(where weatherdeck status says it is). On Windows, run weatherdeck startup remove first if you ran startup install .
Nothing else is installed: no service, no registry entry, no files elsewhere.
Windows 11
macOS (Apple Silicon)
Linux (Ubuntu, WSL2)
Host ( weatherdeck run , on , ingest )
supported
supported
supported
Python
3.9 to 3.14 (3.8 is not supported)
3.9 to 3.14 (3.8 is not supported)
3.9 to 3.14 (3.8 is not supported)
Test suite on GitHub's hosted runners (first run, 2026-10-16; 433 tests, .github/workflows/test.yml , run by hand)
10 to 13 tests fail on every Python version. On the hosted Windows image nine tests that start Node fail because Windows rejects the command line as too long; one doctor check, one push-time check, one text-cutting check and one scratch-folder cleanup also fail there (the Python 3.9 and 3.10 legs fail 10, the others 13). The host itself runs natively on Windows 11, where the maintainer runs it, and this is not yet fixed
all pass
all pass
Page: Chrome, Edge
supported
supported
supported
Page: Safari (macOS), iPhone
headless WebKit (layout, overflow, console errors); the dashboard depth scene was checked and signed off on Safari, iPhone and a 4K TV on 2026-10-10 by the maintainer, and nothing else on those devices has been ( docs/device-checklist.md )
same
same
Page: Firefox
expected to work, not checked each release
Since you last looked (browser storage)
per browser; a browser that blocks site data shows no since plate
same
same
Feed schema read
1.0 to 1.11
1.0 to 1.11
1.0 to 1.11
One ship per machine (0.4.0)
TV: up to six ships, then a convoy; phone: the convoy strip scrolls inside itself
same
same
The open sea (0.4.1)
TV (ship deck): with nothing waiting, one banner line and the sea across the whole upper area; checked at 1920 x 1080 and 3840 x 2160
same
same
Versions per machine (0.5.0)
reporter rows, the watch strip, the trust line and the phone's SOFTWARE section say each source's adapter, or "version unknown"; only a gap that changes meaning is flagged
same
same
Last seen and the replay key (0.4.1)
per browser tab (session storage); a tab that blocks it starts with "not yet on this screen"
same
same
Slowed on purpose
[truncated]
What each release was actually run on is in CHANGELOG.md , with the checks only a person can make
(a real Safari, iPhone and TV; only the dashboard depth scene has been signed off so far, see
docs/device-checklist.md ) listed as such. Tests: python3 -m unittest discover -s tests (see CONTRIBUTING.md).
docs/honesty.md : the honesty contract on one page: what the board never does (green when silent, zero for unknown, a made-up event), where every number, creature and light comes from in the feed, the machine-alert debounce (two feeds to raise, two to clear, one keepalive of latency), and why no value number is shown until it is measured.
docs/ledger-spec.md : the tool value ledger: one append-only file per tool, unknown written as unknown with its reason, the honest verbs (surfaced, refused, placed, landed, not dispatched), and the reading rule that turns fourteen days of lines into KEEP, QUIET, CUT-CANDIDATE or UNMEASURED. No value number is shown until it is measured.
docs/feed-spec.md : the feed schema (1.11), how sources merge, the five lane classes, machine rest and strain, lane state (one word and one sentence per lane), goals (the goal each lane serves, with a count under it), allowance chips, black holes (no agent at the work, outside waits, queues, coming up, the shipping gap, freshness, burn rate), the library (one honest age and word per mirrored folder), the view at a glance (since you last looked, grouping by project), and what silence looks like.
CHANGELOG.md : what changed in each release, and what it was checked on; docs/device-checklist.md : w

[truncated]
