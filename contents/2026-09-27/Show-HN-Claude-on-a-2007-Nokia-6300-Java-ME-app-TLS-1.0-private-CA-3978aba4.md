---
source: "https://github.com/emir/claude-s40"
hn_url: "https://news.ycombinator.com/item?id=49868631"
title: "Show HN: Claude on a 2007 Nokia 6300 (Java ME app, TLS 1.0, private CA)"
article_title: "GitHub - emir/claude-s40: A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. · GitHub"
image: "https://repository-images.githubusercontent.com/1388455733/af62ee35-e541-4934-ba5f-015934a49de5"
author: "emir"
captured_at: "2026-09-27T17:33:09Z"
capture_tool: "hn-digest"
hn_id: 49868631
score: 1
comments: 0
posted_at: "2026-09-27T17:13:40Z"
tags:
  - hacker-news
---

# Show HN: Claude on a 2007 Nokia 6300 (Java ME app, TLS 1.0, private CA)

- HN: [49868631](https://news.ycombinator.com/item?id=49868631)
- Source: [github.com](https://github.com/emir/claude-s40)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T17:13:40Z

## Translation

Title: Show HN: Claude on a 2007 Nokia 6300 (Java ME app, TLS 1.0, private CA)
Article title: GitHub - emir/claude-s40: A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. · GitHub
Description: A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. - emir/claude-s40

Article text:
GitHub - emir/claude-s40: A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. · GitHub
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
emir
/
claude-s40
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows app app docs docs server server .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md View all files Repository files navigation
Chat with Claude on a 2007 Nokia. Claude S40 is an unofficial Claude
client for Nokia Series 40 phones (Java ME, CLDC 1.1 / MIDP 2.0), plus the
small Go server it talks to. Type on the keypad, get Claude's answer on a
240x320 screen: with today's news, weather and exchange rates from web
search, long answers you can page through, and a UI in Turkish or English.
Video: Claude S40 on a Nokia 6300, on X
Chat with bubbles, timestamps and a typing indicator that counts the
seconds. Replies keep their paragraphs and lists (dots, numbers, hanging
indent).
Web search for news, weather, rates: Claude searches on the server,
the phone's 2007 browser is not involved. Sources are listed under the
reply.
Reading mode (key 7): one reply full width, page by page, whole lines
only, page number and progress line. It keeps its place when you change
the text size (key 9) or load the rest, and keeps the backlight on.
Long replies in parts : "0 · Show the rest" fetches the next part
from the server for free; Claude is not asked again.
Message actions : select a message with 1/3, press the centre key:
shorten, explain more simply, translate, ask about it, open it in the
editor. They only fill in the editor; nothing is sent until you press
Send.
Earlier chats from the server, continue any of them; pin the ones
you want to keep at the top (kept until unpinned), delete or search
all chats (no Turkish letters needed: "sise" finds "şişe"); optional
offline copy of the last chat.
Save to phone : a reply as a .txt file (memory card if there is one),
readable later without the network under "Saved".
Calendar and to-do : ask "add to my calendar: dentist tomorrow at 3"
and the reply carries a ready entry; the centre key opens a prefilled
form, and only your "Save" writes it to the phone's own calendar or to-do
list. Works from any message too.
Your notes for Claude (Settings): "I'm Emir, I live in Istanbul, keep
it short" is sent with every message.
Data usage : requests and approximate kilobytes today and in total.
20 quick prompts (search the web, weather, translate, reply to a
message, add to my calendar, summarize, fix my writing, ...).
Setup wizard on first start: language, server address, connection
test, pairing with a 6-digit code (no long code to type).
Keypad-first : every screen works with the keypad; a Shortcuts screen
lists every key. Retry after an error is one key and never charges twice.
Look and feel : light and dark theme, three text sizes, start-up
animation and jingle, reply chime, vibration and backlight. Texts are
fitted to the screen and checked from 128x160 to 320x240.
One Go binary / Docker image that speaks TLS 1.0 to the phone with a
certificate from your own private CA, and modern HTTPS to the Claude API.
Per-device access tokens through pairing, daily request, token and
web-search limits, no automatic retries of paid calls.
SQLite for chats (30 days, pinned ones until unpinned), search over
them, admin API on localhost only.
Home
Waiting for Claude
Lists and paragraphs
Reading mode
A calendar entry, selected
Message actions
Quick prompts
Dark theme, large text
Screenshots are from the FreeJ2ME emulator in the app's test mode (the
"[Test mode]" replies are fake and cost nothing; the emulator reports no
calendar API, so the harness turns the calendar actions on for these
screens and never saves); the cover is a drawing.
Start-up animation: docs/images/splash.gif .
Unofficial side project. Not made, endorsed or supported by Anthropic or
Nokia. See TRADEMARKS.md .
Tested on a Nokia 6300 (RM-217, firmware V06.60) . Other Series 40
phones (CLDC 1.1 / MIDP 2.0, 240x320) may work but are untested.
Nokia (Java ME app) --HTTPS: TLS 1.0, RSA, no SNI, cert from YOUR private CA--> claude-s40-server --HTTPS--> Claude API
|
SQLite (Docker volume)
A 2007 phone cannot speak modern HTTPS: the Nokia 6300 offers only TLS 1.0
with RSA key exchange, sends no SNI, and its root store stops around 2008.
Public clouds (including Cloudflare) refuse that handshake. So the server
terminates TLS itself with a certificate from a private root CA that you
put on the phone once. Details and measurements: docs/ARCHITECTURE.md .
Path
What
app/
The phone app: CLDC 1.1 / MIDP 2.0 MIDlet, ~100 KB JAR, English + Turkish UI (see Features). Reproducible build with 44 package checks.
server/
One Go binary / Docker image (~7 MB): phone-facing TLS, chat backend (official anthropic-sdk-go ), SQLite, pairing, admin API bound to localhost.
docs/
SETUP.md (step by step), ARCHITECTURE.md (protocol, TLS, design).
Get it
git clone https://github.com/emir/claude-s40.git
cd claude-s40
Quick start
Full guide: docs/SETUP.md . In short:
A Docker host with a public IPv4 and TCP 443 (a $4-6/month VPS is enough).
server/scripts/pki.sh ca ~/.config/claude-s40/pki and
server/scripts/pki.sh server ~/.config/claude-s40/pki <server-ip> .
server/deploy/push.sh root@<server-ip> --execute (starts in test mode, no API cost).
echo GATEWAY_URL=https://<server-ip> > app/app.local.properties && make -C app ,
then install app/dist/ClaudeS40.jad/.jar on the phone.
Put the root CA on the phone ( server/deploy/serve-ca.sh , compare the fingerprint).
On the phone the setup wizard walks through the connection test and
pairing; approve the code with
server/deploy/admin.sh root@<server-ip> pair <code> .
server/deploy/set-key.sh root@<server-ip> and
S40_MOCK=0 server/deploy/push.sh root@<server-ip> --execute to go live.
make test # server: go vet + go test -race; app: build + 44 checks + reproducibility
make -C app # phone app only (downloads pinned build tools to app/.deps)
make -C server test
Requirements: Go 1.26+, JDK 11+, Python 3 with Pillow, Docker (for
deployment), OpenSSL or LibreSSL.
The Claude API key lives only on your server. The phone gets a per-device,
revocable access token through pairing.
Chats are stored on your server for 30 days after the last message;
pinned chats until you unpin or delete them. The phone keeps the last
chat only if you turn on "Keep last chat on phone"; replies you save and
calendar entries you add stay on the phone. Your notes for Claude are
stored on the phone and sent with each message (never logged).
The setup (server address, access code, notes) is also kept in a file on
the memory card so a new build needs no new pairing; revoke the device
if the card leaves your hands, or use Settings > Reset setup.
Logs contain no message text.
Every message is a Claude API call billed to your key. The server enforces
per-device daily request, output-token and web-search limits and never
retries a paid call automatically. Web searches are billed per search and
their results count as input tokens. Set a spending limit in the Claude
Console.
MIT, see LICENSE . Copyright (c) 2026 Emir Karşıyakalı. Build-time tools are downloaded, not bundled
(ECJ: EPL-2.0, ProGuard: GPL-2.0, MicroEmulator API stubs: LGPL); none of
them end up in the phone app. The optional emulator harness
app/emu/EmuShot.java links against FreeJ2ME (GPL-3.0), which is not
included.
A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server.
Readme MIT license Security policy
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. - emir/claude-s40

GitHub - emir/claude-s40: A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server. · GitHub
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
emir
/
claude-s40
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows app app docs docs server server .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md View all files Repository files navigation
Chat with Claude on a 2007 Nokia. Claude S40 is an unofficial Claude
client for Nokia Series 40 phones (Java ME, CLDC 1.1 / MIDP 2.0), plus the
small Go server it talks to. Type on the keypad, get Claude's answer on a
240x320 screen: with today's news, weather and exchange rates from web
search, long answers you can page through, and a UI in Turkish or English.
Video: Claude S40 on a Nokia 6300, on X
Chat with bubbles, timestamps and a typing indicator that counts the
seconds. Replies keep their paragraphs and lists (dots, numbers, hanging
indent).
Web search for news, weather, rates: Claude searches on the server,
the phone's 2007 browser is not involved. Sources are listed under the
reply.
Reading mode (key 7): one reply full width, page by page, whole lines
only, page number and progress line. It keeps its place when you change
the text size (key 9) or load the rest, and keeps the backlight on.
Long replies in parts : "0 · Show the rest" fetches the next part
from the server for free; Claude is not asked again.
Message actions : select a message with 1/3, press the centre key:
shorten, explain more simply, translate, ask about it, open it in the
editor. They only fill in the editor; nothing is sent until you press
Send.
Earlier chats from the server, continue any of them; pin the ones
you want to keep at the top (kept until unpinned), delete or search
all chats (no Turkish letters needed: "sise" finds "şişe"); optional
offline copy of the last chat.
Save to phone : a reply as a .txt file (memory card if there is one),
readable later without the network under "Saved".
Calendar and to-do : ask "add to my calendar: dentist tomorrow at 3"
and the reply carries a ready entry; the centre key opens a prefilled
form, and only your "Save" writes it to the phone's own calendar or to-do
list. Works from any message too.
Your notes for Claude (Settings): "I'm Emir, I live in Istanbul, keep
it short" is sent with every message.
Data usage : requests and approximate kilobytes today and in total.
20 quick prompts (search the web, weather, translate, reply to a
message, add to my calendar, summarize, fix my writing, ...).
Setup wizard on first start: language, server address, connection
test, pairing with a 6-digit code (no long code to type).
Keypad-first : every screen works with the keypad; a Shortcuts screen
lists every key. Retry after an error is one key and never charges twice.
Look and feel : light and dark theme, three text sizes, start-up
animation and jingle, reply chime, vibration and backlight. Texts are
fitted to the screen and checked from 128x160 to 320x240.
One Go binary / Docker image that speaks TLS 1.0 to the phone with a
certificate from your own private CA, and modern HTTPS to the Claude API.
Per-device access tokens through pairing, daily request, token and
web-search limits, no automatic retries of paid calls.
SQLite for chats (30 days, pinned ones until unpinned), search over
them, admin API on localhost only.
Home
Waiting for Claude
Lists and paragraphs
Reading mode
A calendar entry, selected
Message actions
Quick prompts
Dark theme, large text
Screenshots are from the FreeJ2ME emulator in the app's test mode (the
"[Test mode]" replies are fake and cost nothing; the emulator reports no
calendar API, so the harness turns the calendar actions on for these
screens and never saves); the cover is a drawing.
Start-up animation: docs/images/splash.gif .
Unofficial side project. Not made, endorsed or supported by Anthropic or
Nokia. See TRADEMARKS.md .
Tested on a Nokia 6300 (RM-217, firmware V06.60) . Other Series 40
phones (CLDC 1.1 / MIDP 2.0, 240x320) may work but are untested.
Nokia (Java ME app) --HTTPS: TLS 1.0, RSA, no SNI, cert from YOUR private CA--> claude-s40-server --HTTPS--> Claude API
|
SQLite (Docker volume)
A 2007 phone cannot speak modern HTTPS: the Nokia 6300 offers only TLS 1.0
with RSA key exchange, sends no SNI, and its root store stops around 2008.
Public clouds (including Cloudflare) refuse that handshake. So the server
terminates TLS itself with a certificate from a private root CA that you
put on the phone once. Details and measurements: docs/ARCHITECTURE.md .
Path
What
app/
The phone app: CLDC 1.1 / MIDP 2.0 MIDlet, ~100 KB JAR, English + Turkish UI (see Features). Reproducible build with 44 package checks.
server/
One Go binary / Docker image (~7 MB): phone-facing TLS, chat backend (official anthropic-sdk-go ), SQLite, pairing, admin API bound to localhost.
docs/
SETUP.md (step by step), ARCHITECTURE.md (protocol, TLS, design).
Get it
git clone https://github.com/emir/claude-s40.git
cd claude-s40
Quick start
Full guide: docs/SETUP.md . In short:
A Docker host with a public IPv4 and TCP 443 (a $4-6/month VPS is enough).
server/scripts/pki.sh ca ~/.config/claude-s40/pki and
server/scripts/pki.sh server ~/.config/claude-s40/pki <server-ip> .
server/deploy/push.sh root@<server-ip> --execute (starts in test mode, no API cost).
echo GATEWAY_URL=https://<server-ip> > app/app.local.properties && make -C app ,
then install app/dist/ClaudeS40.jad/.jar on the phone.
Put the root CA on the phone ( server/deploy/serve-ca.sh , compare the fingerprint).
On the phone the setup wizard walks through the connection test and
pairing; approve the code with
server/deploy/admin.sh root@<server-ip> pair <code> .
server/deploy/set-key.sh root@<server-ip> and
S40_MOCK=0 server/deploy/push.sh root@<server-ip> --execute to go live.
make test # server: go vet + go test -race; app: build + 44 checks + reproducibility
make -C app # phone app only (downloads pinned build tools to app/.deps)
make -C server test
Requirements: Go 1.26+, JDK 11+, Python 3 with Pillow, Docker (for
deployment), OpenSSL or LibreSSL.
The Claude API key lives only on your server. The phone gets a per-device,
revocable access token through pairing.
Chats are stored on your server for 30 days after the last message;
pinned chats until you unpin or delete them. The phone keeps the last
chat only if you turn on "Keep last chat on phone"; replies you save and
calendar entries you add stay on the phone. Your notes for Claude are
stored on the phone and sent with each message (never logged).
The setup (server address, access code, notes) is also kept in a file on
the memory card so a new build needs no new pairing; revoke the device
if the card leaves your hands, or use Settings > Reset setup.
Logs contain no message text.
Every message is a Claude API call billed to your key. The server enforces
per-device daily request, output-token and web-search limits and never
retries a paid call automatically. Web searches are billed per search and
their results count as input tokens. Set a spending limit in the Claude
Console.
MIT, see LICENSE . Copyright (c) 2026 Emir Karşıyakalı. Build-time tools are downloaded, not bundled
(ECJ: EPL-2.0, ProGuard: GPL-2.0, MicroEmulator API stubs: LGPL); none of
them end up in the phone app. The optional emulator harness
app/emu/EmuShot.java links against FreeJ2ME (GPL-3.0), which is not
included.
A 2007 Nokia can't search Google anymore, so I gave it Claude. Unofficial J2ME app + tiny Go server.
Readme MIT license Security policy
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
