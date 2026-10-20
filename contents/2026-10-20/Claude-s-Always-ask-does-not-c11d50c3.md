---
source: "https://blog.baldassari.me/auto-mode-can-use-connectors/"
hn_url: "https://news.ycombinator.com/item?id=50173108"
title: "Claude's \"Always ask\" does not"
article_title: "Claude's \"Always ask\" does not"
image: ""
author: "fra"
captured_at: "2026-10-20T22:11:27Z"
capture_tool: "hn-digest"
hn_id: 50173108
score: 1
comments: 0
posted_at: "2026-10-20T21:44:22Z"
tags:
  - hacker-news
---

# Claude's "Always ask" does not

- HN: [50173108](https://news.ycombinator.com/item?id=50173108)
- Source: [blog.baldassari.me](https://blog.baldassari.me/auto-mode-can-use-connectors/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T21:44:22Z

## Translation

Title: Claude's "Always ask" does not

Article text:
.. /auto-mode-can-use-connectors
Claude's "Always ask" does not
I found out a few days ago, that regardless of your permission settings on
connectors, Claude auto-mode is able to use them. This violated my expectations,
and exposed me to
the lethal trifeca of AI agents .
I had received an email from a service provider with a link to a platform I had
not heard of before. As I often do to learn about something, I asked Claude:
"what can you tell me about [platform.com]"? Claude gave me an answer, and to my
astonishment it concluded "the service provider that emailed you at 1:59pm today
likely uses [platform] to manage projects.
I did a double take. How could Claude know this? I did have a Fastmail
connector, but I was pretty sure I had locked down permissions. I checked
Claude's settings and sure enough, all actions under the Fastmail connector were
set to "Always ask for permission".
I reported it to Anthropic, thinking it was a bug (and a serious one at that).
They responded promptly that this was expected behavior, and indeed when I
however over the "Needs Approval" icon the tooltip states: "Claude asks before
each use, except in sessions set to Automatically approve".
Why is this a problem? Because the auto mode classifier is not fool proof and
email is very sensitive (even read only). An attacker could craft content that
would prompt Claude to read a 2FA code from my email and post it to some
website. While I'm sure auto mode would block many of those instances, I am not
ready to trust it with this.
My preference would be for the "Always Ask" setting to ... always ask for
permission! Or for a fourth option to be created: Always ask, except in auto
mode.
Until then, I've removed the fastmail connector from Claude :'(.

## Original Extract

.. /auto-mode-can-use-connectors
Claude's "Always ask" does not
I found out a few days ago, that regardless of your permission settings on
connectors, Claude auto-mode is able to use them. This violated my expectations,
and exposed me to
the lethal trifeca of AI agents .
I had received an email from a service provider with a link to a platform I had
not heard of before. As I often do to learn about something, I asked Claude:
"what can you tell me about [platform.com]"? Claude gave me an answer, and to my
astonishment it concluded "the service provider that emailed you at 1:59pm today
likely uses [platform] to manage projects.
I did a double take. How could Claude know this? I did have a Fastmail
connector, but I was pretty sure I had locked down permissions. I checked
Claude's settings and sure enough, all actions under the Fastmail connector were
set to "Always ask for permission".
I reported it to Anthropic, thinking it was a bug (and a serious one at that).
They responded promptly that this was expected behavior, and indeed when I
however over the "Needs Approval" icon the tooltip states: "Claude asks before
each use, except in sessions set to Automatically approve".
Why is this a problem? Because the auto mode classifier is not fool proof and
email is very sensitive (even read only). An attacker could craft content that
would prompt Claude to read a 2FA code from my email and post it to some
website. While I'm sure auto mode would block many of those instances, I am not
ready to trust it with this.
My preference would be for the "Always Ask" setting to ... always ask for
permission! Or for a fourth option to be created: Always ask, except in auto
mode.
Until then, I've removed the fastmail connector from Claude :'(.
