---
source: "https://bweller.me/write-all-the-code-youll-ever-need-from-your-phone-for-20-per-month/"
hn_url: "https://news.ycombinator.com/item?id=50217054"
title: "Poor Man's Claude Pro – All the code you'll ever need for $20 per month"
article_title: "Poor Man's Claude Pro"
image: "https://images.unsplash.com/photo-1744130400729-c4aee523f490?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=M3wxMTc3M3wwfDF8c2VhcmNofDR8fG1vbmV5fGVufDB8fHx8MTc4ODkyODgxNHww&ixlib=rb-4.1.0&q=80&w=2000"
author: "bweller"
captured_at: "2026-10-24T00:18:26Z"
capture_tool: "hn-digest"
hn_id: 50217054
score: 1
comments: 0
posted_at: "2026-10-24T00:09:30Z"
tags:
  - hacker-news
---

# Poor Man's Claude Pro – All the code you'll ever need for $20 per month

- HN: [50217054](https://news.ycombinator.com/item?id=50217054)
- Source: [bweller.me](https://bweller.me/write-all-the-code-youll-ever-need-from-your-phone-for-20-per-month/)
- Score: 1
- Comments: 0
- Posted: 2026-10-24T00:09:30Z

## Translation

Title: Poor Man's Claude Pro – All the code you'll ever need for $20 per month
Article title: Poor Man's Claude Pro
Description: Write all the code you'll ever need from your phone for $20 per month
A Yeti water bottle + a Nike T-shirt + Irish Spring 5 in 1: Body, Face, Shampoo, Conditioner, and 24 Hour Deodorant, or an AeroPress Go Portable Coffee Maker + a three-pack of the Field Notes memo books + the

Article text:
Sign in
Subscribe
Poor Man's Claude Pro
Write all the code you'll ever need from your phone for $20 per month
A Yeti water bottle + a Nike T-shirt + Irish Spring 5 in 1: Body, Face, Shampoo, Conditioner, and 24 Hour Deodorant , or an AeroPress Go Portable Coffee Maker + a three-pack of the Field Notes memo books + the classic Swiss Army Knife , or The Power Broker + Graza Drizzle Extra Virgin Olive Oil + a Patagonia snapback .
That's how much more you're missing out per month by spending on the Claude Max plan vs. just the Pro plan.
My coworkers bragged about how much they could code on the go via the Claude Max plan that allows them to run Claude remotely on a server Anthropic provides and push to a GitHub repo by default....jokes on them, I can now attempt to carry around my three Yeti water bottles while layered under three layers of shirts, but I don't have to worry about any BO from sweat because I used my Irish Spring 5 in 1 magic serum today; all while doing the same exact thing .
To get started with using Claude to write code/research/design documents on the go:
Set up your web server, I used Derek Sivers guide as inspiration
Set up a decent terminal on your phone
Include this alias in your phone's local dotfiles ( you do have dotfiles set up.....right? ):
function claudeonserver() {
local ts
ts="$(date +%Y%m%d-%H%M%S)"
ssh <username@server> -i <keyfile> \
"mkdir -p ~/<project_path>/${ts} && tmux new-session -d -s claude-${ts} -c ~/<project_path>/${ts} '/home/<username>/.local/bin/claude --remote-control'"
}
Then it's as simple as opening your terminal and launching a new Claude session per project. If you want Claude to have the ability to create repositories to host these projects in, you could install/set up the GitHub CLI and ask it to push (or just do it manually, you animals).
Now enjoy your coffee while staring at the blank pages of your Field Notes memo book while playing five-finger fillet with your Swiss Army knife; or whatever else you're into.
Ben Weller
16 Aug 2026
← Previous
Sign up

## Original Extract

Write all the code you'll ever need from your phone for $20 per month
A Yeti water bottle + a Nike T-shirt + Irish Spring 5 in 1: Body, Face, Shampoo, Conditioner, and 24 Hour Deodorant, or an AeroPress Go Portable Coffee Maker + a three-pack of the Field Notes memo books + the

Sign in
Subscribe
Poor Man's Claude Pro
Write all the code you'll ever need from your phone for $20 per month
A Yeti water bottle + a Nike T-shirt + Irish Spring 5 in 1: Body, Face, Shampoo, Conditioner, and 24 Hour Deodorant , or an AeroPress Go Portable Coffee Maker + a three-pack of the Field Notes memo books + the classic Swiss Army Knife , or The Power Broker + Graza Drizzle Extra Virgin Olive Oil + a Patagonia snapback .
That's how much more you're missing out per month by spending on the Claude Max plan vs. just the Pro plan.
My coworkers bragged about how much they could code on the go via the Claude Max plan that allows them to run Claude remotely on a server Anthropic provides and push to a GitHub repo by default....jokes on them, I can now attempt to carry around my three Yeti water bottles while layered under three layers of shirts, but I don't have to worry about any BO from sweat because I used my Irish Spring 5 in 1 magic serum today; all while doing the same exact thing .
To get started with using Claude to write code/research/design documents on the go:
Set up your web server, I used Derek Sivers guide as inspiration
Set up a decent terminal on your phone
Include this alias in your phone's local dotfiles ( you do have dotfiles set up.....right? ):
function claudeonserver() {
local ts
ts="$(date +%Y%m%d-%H%M%S)"
ssh <username@server> -i <keyfile> \
"mkdir -p ~/<project_path>/${ts} && tmux new-session -d -s claude-${ts} -c ~/<project_path>/${ts} '/home/<username>/.local/bin/claude --remote-control'"
}
Then it's as simple as opening your terminal and launching a new Claude session per project. If you want Claude to have the ability to create repositories to host these projects in, you could install/set up the GitHub CLI and ask it to push (or just do it manually, you animals).
Now enjoy your coffee while staring at the blank pages of your Field Notes memo book while playing five-finger fillet with your Swiss Army knife; or whatever else you're into.
Ben Weller
16 Aug 2026
← Previous
Sign up
