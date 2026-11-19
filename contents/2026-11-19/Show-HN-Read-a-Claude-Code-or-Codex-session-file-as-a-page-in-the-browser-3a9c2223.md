---
source: "https://keepplain.com/transcript"
hn_url: "https://news.ycombinator.com/item?id=50561908"
title: "Show HN: Read a Claude Code or Codex session file as a page, in the browser"
article_title: "Claude Code transcript viewer: read, redact and share a .jsonl session"
image: "https://keepplain.com/images/og-home-v3.png"
author: "keepplain"
captured_at: "2026-11-19T16:21:28Z"
capture_tool: "hn-digest"
hn_id: 50561908
score: 1
comments: 0
posted_at: "2026-11-19T16:06:11Z"
tags:
  - hacker-news
---

# Show HN: Read a Claude Code or Codex session file as a page, in the browser

- HN: [50561908](https://news.ycombinator.com/item?id=50561908)
- Source: [keepplain.com](https://keepplain.com/transcript)
- Score: 1
- Comments: 0
- Posted: 2026-11-19T16:06:11Z

## Translation

Title: Show HN: Read a Claude Code or Codex session file as a page, in the browser
Article title: Claude Code transcript viewer: read, redact and share a .jsonl session
Description: Drop a Claude Code, Codex, Cursor or Pi session file. It is parsed in your browser, secrets are replaced before anything leaves your machine, and you get a page to read, download as HTML or share.

Article text:
KeepPlain
Docs
Security
Teams
Pricing
GitHub
Sign in
Transcript viewer for Claude Code, Codex, Cursor and Pi
Read a Claude Code or Codex session as a page
Drop the .jsonl file. It is parsed in your browser, secrets are replaced before anything leaves your machine, and you get a page you can read, download or share.
No account, no upload. Up to 150 MB. Or paste its text .
The download is one file with no script, ready for an issue or a pull request. A link needs a GitHub or Google sign-in: the session becomes a private draft only you see, and nothing is published until you publish it.
Every agent writes each session to a file on your computer. The folders are hidden: paste the path into your file manager rather than looking for it.
~/.claude/projects/<project>/<session id>.jsonl . The terminal, the desktop app and the editor extensions all write there. Sessions older than 30 days are deleted unless cleanupPeriodDays is set.
Copy the folder’s path and choose the file
~/.codex/sessions/<year>/<month>/<day>/rollout-….jsonl . The file name starts with the date and time the session began.
Copy the folder’s path and choose the file
~/.cursor/projects/<project>/agent-transcripts/<chat id>/<chat id>.jsonl . Cursor’s file has no times, so the page shows none.
Copy the folder’s path and choose the file
~/.pi/agent/sessions/<project>/<date>_….jsonl .
Copy the folder’s path and choose the file
On Windows ~ is %USERPROFILE% . Antigravity’s transcript_full.jsonl , a Hermes Agent export and a .jsonl.gz the KeepPlain plugin prepared open the same way.
The page opens the file, drops screenshots and long command output, and draws the conversation: prompts, replies, tool calls, changed files. No request carries any of it. You can check in your browser’s network tab.
API keys, tokens, passwords in connection strings, email addresses and public IP addresses become [REDACTED] , and your home folder becomes ~ . The page lists what it found by its first characters, and you choose what to show as it is.
A link only when you ask for one
“Get a link” signs you in and puts the checked session into your drafts on KeepPlain, private to you. From there you cut it to the moments that mattered and publish it, or keep it to yourself.
Stop finding the file by hand.
One command, and your sessions are kept as they are written: searchable from any agent, with keepplain blame for every line and rules made from your corrections. Free for personal use.
keepplain enable
Install the keepplain command first · What leaves your computer · Plugin on GitHub
The file you drop is read in your browser. Nothing is uploaded unless you press Get a link and sign in.

## Original Extract

Drop a Claude Code, Codex, Cursor or Pi session file. It is parsed in your browser, secrets are replaced before anything leaves your machine, and you get a page to read, download as HTML or share.

KeepPlain
Docs
Security
Teams
Pricing
GitHub
Sign in
Transcript viewer for Claude Code, Codex, Cursor and Pi
Read a Claude Code or Codex session as a page
Drop the .jsonl file. It is parsed in your browser, secrets are replaced before anything leaves your machine, and you get a page you can read, download or share.
No account, no upload. Up to 150 MB. Or paste its text .
The download is one file with no script, ready for an issue or a pull request. A link needs a GitHub or Google sign-in: the session becomes a private draft only you see, and nothing is published until you publish it.
Every agent writes each session to a file on your computer. The folders are hidden: paste the path into your file manager rather than looking for it.
~/.claude/projects/<project>/<session id>.jsonl . The terminal, the desktop app and the editor extensions all write there. Sessions older than 30 days are deleted unless cleanupPeriodDays is set.
Copy the folder’s path and choose the file
~/.codex/sessions/<year>/<month>/<day>/rollout-….jsonl . The file name starts with the date and time the session began.
Copy the folder’s path and choose the file
~/.cursor/projects/<project>/agent-transcripts/<chat id>/<chat id>.jsonl . Cursor’s file has no times, so the page shows none.
Copy the folder’s path and choose the file
~/.pi/agent/sessions/<project>/<date>_….jsonl .
Copy the folder’s path and choose the file
On Windows ~ is %USERPROFILE% . Antigravity’s transcript_full.jsonl , a Hermes Agent export and a .jsonl.gz the KeepPlain plugin prepared open the same way.
The page opens the file, drops screenshots and long command output, and draws the conversation: prompts, replies, tool calls, changed files. No request carries any of it. You can check in your browser’s network tab.
API keys, tokens, passwords in connection strings, email addresses and public IP addresses become [REDACTED] , and your home folder becomes ~ . The page lists what it found by its first characters, and you choose what to show as it is.
A link only when you ask for one
“Get a link” signs you in and puts the checked session into your drafts on KeepPlain, private to you. From there you cut it to the moments that mattered and publish it, or keep it to yourself.
Stop finding the file by hand.
One command, and your sessions are kept as they are written: searchable from any agent, with keepplain blame for every line and rules made from your corrections. Free for personal use.
keepplain enable
Install the keepplain command first · What leaves your computer · Plugin on GitHub
The file you drop is read in your browser. Nothing is uploaded unless you press Get a link and sign in.
